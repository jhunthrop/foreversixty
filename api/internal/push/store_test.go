package push

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/events"
)

const (
	iosToken     = "a1b2c3d4e5f60718293a4b5c6d7e8f90a1b2c3d4e5f60718293a4b5c6d7e8f90"
	androidToken = "fcm:APA91bHPRgkF3JUikC4ENAHEeMrd41Zxv3hVZjC9KtT8OvPVGJ-hQMRKRrZuJAEcl7B338qju59zJMjw2DeLhzEebdJjc6iL2lJHpEVpI8V-v_-"
)

func testPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set; start docker-compose.test.yml")
	}
	if err := db.Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := db.Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	if _, err := pool.Exec(context.Background(), `truncate users cascade`); err != nil {
		t.Fatal(err)
	}
	return pool
}

// world is one account with two paired app devices, and the device bearer tokens.
type world struct {
	pool     *pgxpool.Pool
	store    *Store
	accounts *auth.Store
	uid      int64
	devices  map[string]string // device id -> bearer token
}

func newWorld(t *testing.T) world {
	t.Helper()
	pool := testPool(t)
	w := world{pool: pool, store: &Store{Pool: pool}, accounts: &auth.Store{Pool: pool}, devices: map[string]string{}}
	if err := pool.QueryRow(context.Background(), `insert into users (email) values ($1) returning id`,
		"push-"+time.Now().Format("150405.000000000")+"@example.com").Scan(&w.uid); err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{"phone", "tablet"} {
		bearer := "fsd_" + id + "-secret-0123456789"
		if _, err := pool.Exec(context.Background(), `insert into devices (id, user_id, platform, token_hash) values ($1, $2, 'ios', $3)`,
			id, w.uid, auth.TokenHash(bearer)); err != nil {
			t.Fatal(err)
		}
		w.devices[id] = bearer
	}
	return w
}

// serve sends a request through the real device-token middleware and the routes.
func (w world) serve(t *testing.T, method, path, bearer, body string) int {
	t.Helper()
	mux := http.NewServeMux()
	Mount(mux, &Handler{Store: w.store})
	authn := &auth.Authenticator{Store: w.accounts}
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	if bearer != "" {
		req.Header.Set("Authorization", "Bearer "+bearer)
	}
	rec := httptest.NewRecorder()
	authn.Middleware(mux).ServeHTTP(rec, req)
	return rec.Code
}

func (w world) tokens(t *testing.T) []Token {
	t.Helper()
	out, err := w.store.TokensOf(context.Background(), w.uid)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

func TestRegisterTiesTheTokenToTheCallingDevice(t *testing.T) {
	w := newWorld(t)
	body := `{"platform":"ios","token":"` + iosToken + `","app_version":"1.0.0 (12)"}`
	if code := w.serve(t, http.MethodPost, "/v1/me/push-tokens", w.devices["phone"], body); code != http.StatusCreated {
		t.Fatalf("register %d", code)
	}
	got := w.tokens(t)
	if len(got) != 1 || got[0].DeviceID != "phone" || got[0].Platform != IOS || got[0].AppVersion != "1.0.0 (12)" {
		t.Fatalf("tokens %+v", got)
	}
	// The same token from the tablet moves; a new token on the tablet replaces its old one.
	if code := w.serve(t, http.MethodPost, "/v1/me/push-tokens", w.devices["tablet"], body); code != http.StatusCreated {
		t.Fatal(code)
	}
	android := `{"platform":"android","token":"` + androidToken + `"}`
	if code := w.serve(t, http.MethodPost, "/v1/me/push-tokens", w.devices["tablet"], android); code != http.StatusCreated {
		t.Fatal(code)
	}
	got = w.tokens(t)
	if len(got) != 1 || got[0].DeviceID != "tablet" || got[0].Platform != Android {
		t.Fatalf("after moves %+v", got)
	}
}

func TestTokenRoutesNeedADeviceToken(t *testing.T) {
	w := newWorld(t)
	body := `{"platform":"ios","token":"` + iosToken + `"}`
	if code := w.serve(t, http.MethodPost, "/v1/me/push-tokens", "", body); code != http.StatusUnauthorized {
		t.Fatalf("anonymous %d", code)
	}
	if code := w.serve(t, http.MethodPost, "/v1/me/push-tokens", "fsd_unknown-0123456789", body); code != http.StatusUnauthorized {
		t.Fatalf("unknown device %d", code)
	}
	if code := w.serve(t, http.MethodDelete, "/v1/me/push-tokens/"+iosToken, "", ""); code != http.StatusUnauthorized {
		t.Fatalf("anonymous delete %d", code)
	}
}

func TestRegisterValidates(t *testing.T) {
	w := newWorld(t)
	for _, body := range []string{
		`{"platform":"web","token":"` + iosToken + `"}`,
		`{"platform":"ios","token":"not-hex"}`,
		`{"platform":"android","token":"short"}`,
		`{"platform":"ios","token":"` + iosToken + `","app_version":"<script>"}`,
	} {
		if code := w.serve(t, http.MethodPost, "/v1/me/push-tokens", w.devices["phone"], body); code != http.StatusUnprocessableEntity {
			t.Errorf("%s: %d", body, code)
		}
	}
	if code := w.serve(t, http.MethodPost, "/v1/me/push-tokens", w.devices["phone"], `{`); code != http.StatusBadRequest {
		t.Fatalf("bad json %d", code)
	}
}

func TestUnregisterOnlyTouchesTheCallingDevice(t *testing.T) {
	w := newWorld(t)
	ctx := context.Background()
	if err := w.store.Register(ctx, Token{Token: iosToken, Platform: IOS, UserID: w.uid, DeviceID: "phone"}); err != nil {
		t.Fatal(err)
	}
	if code := w.serve(t, http.MethodDelete, "/v1/me/push-tokens/"+iosToken, w.devices["tablet"], ""); code != http.StatusNoContent {
		t.Fatal(code)
	}
	if len(w.tokens(t)) != 1 {
		t.Fatal("another device removed the token")
	}
	if code := w.serve(t, http.MethodDelete, "/v1/me/push-tokens/"+iosToken, w.devices["phone"], ""); code != http.StatusNoContent {
		t.Fatal(code)
	}
	if len(w.tokens(t)) != 0 {
		t.Fatal("token kept")
	}
}

func TestRevokingTheDeviceDropsItsToken(t *testing.T) {
	w := newWorld(t)
	ctx := context.Background()
	for dev, tok := range map[string]string{"phone": iosToken, "tablet": androidToken} {
		p := IOS
		if dev == "tablet" {
			p = Android
		}
		if err := w.store.Register(ctx, Token{Token: tok, Platform: p, UserID: w.uid, DeviceID: dev}); err != nil {
			t.Fatal(err)
		}
	}
	if ok, err := w.accounts.RevokeDevice(ctx, w.uid, "phone"); err != nil || !ok {
		t.Fatalf("revoke %v %v", ok, err)
	}
	got := w.tokens(t)
	if len(got) != 1 || got[0].DeviceID != "tablet" {
		t.Fatalf("after revoke %+v", got)
	}
	// A revoked device cannot register again.
	if err := w.store.Register(ctx, Token{Token: iosToken, Platform: IOS, UserID: w.uid, DeviceID: "phone"}); err != nil {
		t.Fatal(err)
	}
	if len(w.tokens(t)) != 1 {
		t.Fatal("a revoked device registered a token")
	}
}

func TestSinkSendsEveryTokenAndDropsTheInvalid(t *testing.T) {
	w := newWorld(t)
	ctx := context.Background()
	_ = w.store.Register(ctx, Token{Token: iosToken, Platform: IOS, UserID: w.uid, DeviceID: "phone"})
	_ = w.store.Register(ctx, Token{Token: androidToken, Platform: Android, UserID: w.uid, DeviceID: "tablet"})
	apns, fcm := &Fake{Invalid: map[string]bool{iosToken: true}}, &Fake{}
	sink := &Sink{Store: w.store, Senders: map[Platform]Sender{IOS: apns, Android: fcm}}
	e := events.Event{Kind: events.RaidPosted, Title: "MC", Body: "Thu", Link: "https://foreversixty.gg/x", Thread: "raid-1",
		Collapse: "raid.posted-1", Actions: []events.Action{{ID: "signup_in"}}}
	if err := sink.Deliver(ctx, w.uid, e); err != nil {
		t.Fatal(err)
	}
	if len(apns.Sent()) != 1 || len(fcm.Sent()) != 1 || fcm.Sent()[0].Category != RaidSignupCategory {
		t.Fatalf("apns %v fcm %v", apns.Sent(), fcm.Sent())
	}
	got := w.tokens(t)
	if len(got) != 1 || got[0].Platform != Android {
		t.Fatalf("invalid token kept: %+v", got)
	}
}

type memTokens struct{ tokens []Token }

func (m *memTokens) Register(context.Context, Token) error                    { return nil }
func (m *memTokens) Unregister(context.Context, string, string) (bool, error) { return false, nil }
func (m *memTokens) TokensOf(context.Context, int64) ([]Token, error)         { return m.tokens, nil }
func (m *memTokens) Delete(context.Context, string) error                     { return nil }

func TestSinkFailsOnlyWhenNoTokenTookIt(t *testing.T) {
	down := &Fake{Err: errors.New("down")}
	up := &Fake{}
	store := &memTokens{tokens: []Token{{Token: "a", Platform: IOS}, {Token: "b", Platform: Android}}}
	e := events.Event{Kind: events.SimFinished}
	if err := (&Sink{Store: store, Senders: map[Platform]Sender{IOS: down, Android: up}}).Deliver(context.Background(), 1, e); err != nil {
		t.Fatalf("partial: %v", err)
	}
	if err := (&Sink{Store: store, Senders: map[Platform]Sender{IOS: down, Android: down}}).Deliver(context.Background(), 1, e); err == nil {
		t.Fatal("all failed but no error")
	}
	// No sender for the platform (keys absent) sends nothing and fails nothing.
	if err := (&Sink{Store: store, Senders: map[Platform]Sender{}}).Deliver(context.Background(), 1, e); err != nil {
		t.Fatal(err)
	}
}
