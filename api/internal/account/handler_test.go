package account

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// fakeSessions records ClearSession.
type fakeSessions struct{ cleared int }

func (f *fakeSessions) ClearSession(w http.ResponseWriter) {
	f.cleared++
	http.SetCookie(w, &http.Cookie{Name: auth.SessionCookie, Value: "", MaxAge: -1})
}

type answer struct {
	OK    bool `json:"ok"`
	Data  map[string]string
	Error *struct {
		Code    string            `json:"code"`
		Message string            `json:"message"`
		Fields  map[string]string `json:"fields"`
	} `json:"error"`
}

// call sends DELETE /v1/me through Mount, as actor.
func call(t *testing.T, h *Handler, actor auth.Actor, body string) (int, answer, *httptest.ResponseRecorder) {
	t.Helper()
	mux := http.NewServeMux()
	Mount(mux, h)
	req := httptest.NewRequest(http.MethodDelete, Path, strings.NewReader(body))
	req = req.WithContext(auth.WithActor(context.Background(), actor))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	var a answer
	if err := json.Unmarshal(rec.Body.Bytes(), &a); err != nil {
		t.Fatalf("body %q: %v", rec.Body.String(), err)
	}
	return rec.Code, a, rec
}

func TestTheRouteAnswersEveryCase(t *testing.T) {
	r := newRig(t)
	sessions := &fakeSessions{}
	h := &Handler{Service: r.svc, Sessions: sessions}
	victim := seedUser(t, r.pool, "Felsnap#1234", "user")
	gm := seedUser(t, r.pool, "Thandrel#1894", "user")
	seedGuild(t, r.pool, "Olympus XXVII", gm)
	browser := auth.Actor{UserID: victim, Method: "bnet"}

	cases := []struct {
		name   string
		actor  auth.Actor
		body   string
		status int
		code   string
	}{
		{"signed out", auth.Actor{}, `{"confirm":"delete"}`, http.StatusUnauthorized, "unauthorized"},
		{"a companion", auth.Actor{UserID: victim, DeviceID: "d1", DeviceKind: auth.DeviceKindCompanion}, `{"confirm":"delete"}`, http.StatusUnauthorized, "unauthorized"},
		{"a broken body", browser, `{`, http.StatusBadRequest, "invalid"},
		{"nothing typed", browser, `{"confirm":"  "}`, http.StatusBadRequest, "invalid"},
		{"the wrong words", browser, `{"confirm":"Felsnap"}`, http.StatusBadRequest, "confirm_mismatch"},
		{"the guild master", auth.Actor{UserID: gm, Method: "bnet"}, `{"confirm":"delete"}`, http.StatusConflict, "guild_obligation"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			status, a, _ := call(t, h, c.actor, c.body)
			if status != c.status || a.Error == nil || a.Error.Code != c.code {
				t.Fatalf("got %d %+v, want %d %s", status, a.Error, c.status, c.code)
			}
		})
	}
	if _, a, _ := call(t, h, auth.Actor{UserID: gm, Method: "bnet"}, `{"confirm":"delete"}`); a.Error.Fields["guilds"] != "Olympus XXVII" ||
		!strings.Contains(a.Error.Message, "Release the claim on your account page first") {
		t.Fatalf("the guild refusal = %+v, want the guild named and the fix", a.Error)
	}
	if sessions.cleared != 0 || !userExists(t, r.pool, victim) {
		t.Fatal("a refusal signed out or deleted")
	}

	t.Run("an app deletes its account", func(t *testing.T) {
		app := auth.Actor{UserID: victim, DeviceID: "a1", DeviceKind: auth.DeviceKindApp}
		status, a, _ := call(t, h, app, `{"confirm":"FELSNAP#1234"}`)
		if status != http.StatusOK || !a.OK || a.Data["deleted_at"] != "2026-11-07T21:41:00Z" {
			t.Fatalf("got %d %+v", status, a)
		}
		if sessions.cleared != 1 || userExists(t, r.pool, victim) {
			t.Fatal("the account was not deleted and signed out")
		}
	})
	t.Run("a second call finds no account", func(t *testing.T) {
		if status, a, _ := call(t, h, browser, `{"confirm":"delete"}`); status != http.StatusUnauthorized || a.Error.Code != "unauthorized" {
			t.Fatalf("got %d %+v, want 401", status, a.Error)
		}
	})
}

func TestTheRouteIsBudgetedPerAccount(t *testing.T) {
	r := newRig(t)
	h := &Handler{Service: r.svc, Sessions: &fakeSessions{}}
	victim := seedUser(t, r.pool, "Felsnap#1234", "user")
	mux := http.NewServeMux()
	Mount(mux, h)
	last := 0
	for range deletesPerHour + 1 {
		req := httptest.NewRequest(http.MethodDelete, Path, strings.NewReader(`{"confirm":"nope"}`))
		req = req.WithContext(auth.WithActor(context.Background(), auth.Actor{UserID: victim, Method: "bnet"}))
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		last = rec.Code
	}
	if last != http.StatusTooManyRequests {
		t.Fatalf("call %d answered %d, want 429", deletesPerHour+1, last)
	}
}
