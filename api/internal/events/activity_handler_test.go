package events

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/contract"
)

// activityCall serves one Activity request as uid (0: anonymous) and answers its status and
// body.
func activityCall(t *testing.T, h *Handler, method, target, body string, uid int64) (int, []byte) {
	t.Helper()
	mux := http.NewServeMux()
	Mount(mux, h)
	req := httptest.NewRequest(method, target, strings.NewReader(body))
	if uid != 0 {
		req = req.WithContext(auth.WithActor(req.Context(), auth.Actor{UserID: uid, DeviceID: "d1"}))
	}
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	raw, err := io.ReadAll(rec.Body)
	if err != nil {
		t.Fatal(err)
	}
	if method == http.MethodGet && rec.Code == http.StatusOK && !strings.Contains(rec.Header().Get("Cache-Control"), "no-store") {
		t.Fatalf("cache-control %q", rec.Header().Get("Cache-Control"))
	}
	return rec.Code, raw
}

func activityHandler(t *testing.T) (*Handler, *Store, int64) {
	t.Helper()
	pool := testPool(t)
	st := &Store{Pool: pool}
	return &Handler{Store: st, Activity: st}, st, seedUser(t, pool, "inbox")
}

func TestActivityAnswersTheFixtures(t *testing.T) {
	h, st, uid := activityHandler(t)
	code, body := activityCall(t, h, http.MethodGet, "/v1/me/activity", "", uid)
	if code != http.StatusOK {
		t.Fatalf("%d %s", code, body)
	}
	contract.AssertFixture(t, "activity.list.empty", body, contract.Options{})
	recordAll(t, st, uid, SimFinished, RaidPosted, ReportReady)
	code, body = activityCall(t, h, http.MethodGet, "/v1/me/activity?limit=2", "", uid)
	if code != http.StatusOK {
		t.Fatalf("%d %s", code, body)
	}
	contract.AssertFixture(t, "activity.list", body, contract.Options{})
	if !strings.Contains(string(body), `"unread":3`) || strings.Contains(string(body), `"next_before":null`) {
		t.Fatalf("page %s", body)
	}
	code, body = activityCall(t, h, http.MethodGet, "/v1/me/activity?group=raids", "", uid)
	if code != http.StatusOK || strings.Count(string(body), `"kind"`) != 1 || !strings.Contains(string(body), `"unread":3`) {
		t.Fatalf("raids %d %s", code, body)
	}
	rows, _, _ := st.Activity(context.Background(), uid, ActivityQuery{Limit: 30})
	code, body = activityCall(t, h, http.MethodPost, "/v1/me/activity/read", fmt.Sprintf(`{"through":%d}`, rows[0].ID), uid)
	if code != http.StatusOK {
		t.Fatalf("%d %s", code, body)
	}
	contract.AssertFixture(t, "activity.read", body, contract.Options{})
	if !strings.Contains(string(body), `"unread":0`) {
		t.Fatalf("read %s", body)
	}
}

func TestActivityRefusesWhatIsNotValid(t *testing.T) {
	h := &Handler{Store: &memPrefs{p: map[int64]Prefs{}}, Activity: &Store{}}
	for _, target := range []string{
		"/v1/me/activity?group=pvp", "/v1/me/activity?limit=0", "/v1/me/activity?limit=51",
		"/v1/me/activity?limit=x", "/v1/me/activity?before=0", "/v1/me/activity?before=x",
	} {
		if code, body := activityCall(t, h, http.MethodGet, target, "", 5); code != http.StatusBadRequest || !strings.Contains(string(body), `"invalid"`) {
			t.Errorf("%s: %d %s", target, code, body)
		}
	}
	for _, body := range []string{``, `nope`, `{}`, `{"through":0}`, `{"through":"5"}`} {
		if code, raw := activityCall(t, h, http.MethodPost, "/v1/me/activity/read", body, 5); code != http.StatusBadRequest {
			t.Errorf("%q: %d %s", body, code, raw)
		}
	}
	if code, _ := activityCall(t, h, http.MethodGet, "/v1/me/activity", "", 0); code != http.StatusUnauthorized {
		t.Fatalf("anonymous %d", code)
	}
	if code, _ := activityCall(t, h, http.MethodPost, "/v1/me/activity/read", `{"through":1}`, 0); code != http.StatusUnauthorized {
		t.Fatalf("anonymous read %d", code)
	}
}

func TestActivityRoutesNeedTheInbox(t *testing.T) {
	h := &Handler{Store: &memPrefs{p: map[int64]Prefs{}}}
	if code, _ := activityCall(t, h, http.MethodGet, "/v1/me/activity", "", 5); code != http.StatusNotFound {
		t.Fatalf("without an inbox %d", code)
	}
}
