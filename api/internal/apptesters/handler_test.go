package apptesters

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

type fakeTesters struct {
	rows []Tester
	err  error
}

func (f *fakeTesters) Add(_ context.Context, email, platform string) (bool, error) {
	if f.err != nil {
		return false, f.err
	}
	for _, r := range f.rows {
		if r.Email == email && r.Platform == platform {
			return false, nil
		}
	}
	f.rows = append(f.rows, Tester{Email: email, Platform: platform, CreatedAt: time.Date(2026, 11, 8, 12, 0, 0, 0, time.UTC)})
	return true, nil
}

func (f *fakeTesters) List(context.Context) ([]Tester, error) { return f.rows, f.err }

var (
	anonymous = auth.Actor{}
	member    = auth.Actor{UserID: 5, Role: "user", Method: "session"}
	owner     = auth.Actor{UserID: 78, Role: "admin", Method: "session"}
)

func newMux(f *fakeTesters) *http.ServeMux {
	mux := http.NewServeMux()
	Mount(mux, &Service{Testers: f})
	return mux
}

func serveOn(mux *http.ServeMux, method, path, body, ip string, actor auth.Actor) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, path, strings.NewReader(body))
	r.RemoteAddr = ip + ":1234"
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, r.WithContext(auth.WithActor(r.Context(), actor)))
	return rec
}

func serve(f *fakeTesters, method, path, body string, actor auth.Actor) *httptest.ResponseRecorder {
	return serveOn(newMux(f), method, path, body, "203.0.113.1", actor)
}

type envelope struct {
	OK    bool            `json:"ok"`
	Data  json.RawMessage `json:"data"`
	Error *struct {
		Code   string            `json:"code"`
		Fields map[string]string `json:"fields"`
	} `json:"error"`
}

func decode(t *testing.T, rec *httptest.ResponseRecorder) envelope {
	t.Helper()
	var e envelope
	if err := json.Unmarshal(rec.Body.Bytes(), &e); err != nil {
		t.Fatalf("body %q: %v", rec.Body.String(), err)
	}
	return e
}

// A new address and one already on the list get the same answer, byte for byte, so the form
// cannot be used to learn who signed up; the repeat stores nothing.
func TestSignupAnswersTheSameForANewAndARepeatedAddress(t *testing.T) {
	f := &fakeTesters{}
	mux := newMux(f)
	body := `{"email":" Thandrel@Example.com ","platform":"ios"}`
	first := serveOn(mux, "POST", SignupPath, body, "203.0.113.1", anonymous)
	again := serveOn(mux, "POST", SignupPath, body, "203.0.113.1", anonymous)
	for _, rec := range []*httptest.ResponseRecorder{first, again} {
		if rec.Code != http.StatusCreated {
			t.Fatalf("status %d %s", rec.Code, rec.Body)
		}
		if got := string(decode(t, rec).Data); got != `{"status":"received"}` {
			t.Fatalf("data %s", got)
		}
	}
	strip := func(rec *httptest.ResponseRecorder) string {
		var e map[string]any
		_ = json.Unmarshal(rec.Body.Bytes(), &e)
		delete(e, "request_id")
		b, _ := json.Marshal(e)
		return string(b)
	}
	if strip(first) != strip(again) {
		t.Fatalf("answers differ: %s vs %s", first.Body, again.Body)
	}
	if len(f.rows) != 1 || f.rows[0].Email != "thandrel@example.com" || f.rows[0].Platform != PlatformIOS {
		t.Fatalf("stored %+v", f.rows)
	}
}

func TestSignupNeverEchoesTheStoredRows(t *testing.T) {
	f := &fakeTesters{rows: []Tester{{Email: "someone@example.com", Platform: PlatformAndroid}}}
	rec := serve(f, "POST", SignupPath, `{"email":"someone@example.com","platform":"android"}`, anonymous)
	if strings.Contains(rec.Body.String(), "someone@") {
		t.Fatalf("the answer carries an address: %s", rec.Body)
	}
}

func TestSignupRefusesABadEmailOrPlatformWith422AndTheField(t *testing.T) {
	cases := map[string]string{
		`{"email":"nope","platform":"ios"}`:         "email",
		`{"email":"a@example.com","platform":"pc"}`: "platform",
		`{"email":"a@example.com"}`:                 "platform",
		`{"platform":"android"}`:                    "email",
	}
	for body, field := range cases {
		f := &fakeTesters{}
		rec := serve(f, "POST", SignupPath, body, anonymous)
		if rec.Code != http.StatusUnprocessableEntity {
			t.Errorf("%s: %d", body, rec.Code)
			continue
		}
		e := decode(t, rec)
		if e.Error == nil || e.Error.Code != "invalid" || e.Error.Fields[field] == "" {
			t.Errorf("%s: %s", body, rec.Body)
		}
		if len(f.rows) != 0 {
			t.Errorf("%s stored %+v", body, f.rows)
		}
	}
}

func TestSignupRefusesABodyThatIsNotJSON(t *testing.T) {
	rec := serve(&fakeTesters{}, "POST", SignupPath, `email=a@example.com`, anonymous)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("%d", rec.Code)
	}
}

func TestSignupIsLimitedPerIPPerHour(t *testing.T) {
	mux := newMux(&fakeTesters{})
	for i := range SignupsPerIPPerHour {
		body := fmt.Sprintf(`{"email":"t%d@example.com","platform":"ios"}`, i)
		if rec := serveOn(mux, "POST", SignupPath, body, "198.51.100.7", anonymous); rec.Code != http.StatusCreated {
			t.Fatalf("signup %d: %d", i, rec.Code)
		}
	}
	over := `{"email":"z@example.com","platform":"ios"}`
	if rec := serveOn(mux, "POST", SignupPath, over, "198.51.100.7", anonymous); rec.Code != http.StatusTooManyRequests {
		t.Fatalf("over the budget: %d", rec.Code)
	}
	if rec := serveOn(mux, "POST", SignupPath, over, "198.51.100.8", anonymous); rec.Code != http.StatusCreated {
		t.Fatalf("another address: %d", rec.Code)
	}
}

func TestSignupStoreFailureIs500WithoutDetail(t *testing.T) {
	f := &fakeTesters{err: errors.New("db down")}
	rec := serve(f, "POST", SignupPath, `{"email":"a@example.com","platform":"ios"}`, anonymous)
	if rec.Code != http.StatusInternalServerError || strings.Contains(rec.Body.String(), "db down") {
		t.Fatalf("%d %s", rec.Code, rec.Body)
	}
}

func TestTheListIsTheOwnersOnly(t *testing.T) {
	f := &fakeTesters{rows: []Tester{{Email: "a@example.com", Platform: PlatformIOS, CreatedAt: time.Date(2026, 11, 8, 12, 0, 0, 0, time.UTC)}}}
	if rec := serve(f, "GET", AdminPath, "", anonymous); rec.Code != http.StatusUnauthorized {
		t.Fatalf("anonymous: %d", rec.Code)
	}
	if rec := serve(f, "GET", AdminPath, "", member); rec.Code != http.StatusForbidden {
		t.Fatalf("member: %d", rec.Code)
	}
	rec := serve(f, "GET", AdminPath, "", owner)
	if rec.Code != http.StatusOK {
		t.Fatalf("owner: %d", rec.Code)
	}
	if cc := rec.Header().Get("Cache-Control"); !strings.Contains(cc, "no-store") {
		t.Fatalf("cache-control %q", cc)
	}
	want := `{"testers":[{"email":"a@example.com","platform":"ios","created_at":"2026-11-08T12:00:00Z"}]}`
	if got := string(decode(t, rec).Data); got != want {
		t.Fatalf("got %s", got)
	}
}

func TestAnEmptyListIsAnEmptyArray(t *testing.T) {
	f := &fakeTesters{}
	rec := serve(f, "GET", AdminPath, "", owner)
	if got := string(decode(t, rec).Data); got != `{"testers":[]}` {
		t.Fatalf("got %s", got)
	}
}
