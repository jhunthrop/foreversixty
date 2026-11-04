package linkedrole

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

const publicBase = "https://foreversixty.gg"

// fakeStates hands out one fixed state and keeps the registered callback.
type fakeStates struct {
	callback notify.StateHandler
	kind     string
	userID   int64
}

func (f *fakeStates) StartState(_ context.Context, kind string, userID int64, _ *int64) (string, error) {
	f.kind, f.userID = kind, userID
	return "st4te", nil
}

func (f *fakeStates) OnState(kind string, h notify.StateHandler) {
	if kind == notify.StateKindLinkedRole {
		f.callback = h
	}
}

func newHandler(t *testing.T) (*Handler, *fakeStates, *memStore, *fakeDiscord) {
	t.Helper()
	s, store, discord := newSyncer()
	states := &fakeStates{}
	h := NewHandler(s, states, publicBase, nil)
	if states.callback == nil {
		t.Fatal("the callback was not registered for the linked_role state")
	}
	return h, states, store, discord
}

func serve(handler http.HandlerFunc, target string, actor auth.Actor) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, target, nil)
	req = req.WithContext(auth.WithActor(req.Context(), actor))
	rec := httptest.NewRecorder()
	handler(rec, req)
	return rec
}

func TestStartSendsASignedOutVisitorThroughSignInBackToTheAccountPage(t *testing.T) {
	h, _, _, _ := newHandler(t)
	rec := serve(h.start, StartPath+"?next=https://evil.example", auth.Actor{})
	want := publicBase + "/login?next=" + url.QueryEscape("/account?discord_role=start#you")
	if got := rec.Header().Get("Location"); rec.Code != http.StatusFound || got != want {
		t.Fatalf("got %d %q, want 302 %q", rec.Code, got, want)
	}
}

func TestStartSendsAMemberToDiscordWithAStateBoundToTheAccount(t *testing.T) {
	h, states, _, _ := newHandler(t)
	rec := serve(h.start, StartPath, auth.Actor{UserID: 7})
	if got := rec.Header().Get("Location"); got != "https://discord.test/authorize?state=st4te" {
		t.Fatalf("location = %q", got)
	}
	if states.kind != notify.StateKindLinkedRole || states.userID != 7 {
		t.Fatalf("state kind %q for user %d", states.kind, states.userID)
	}
	if cc := rec.Header().Get("Cache-Control"); !strings.Contains(cc, "no-store") {
		t.Fatalf("cache-control = %q", cc)
	}
}

func TestStartRefusesADeviceToken(t *testing.T) {
	h, _, _, _ := newHandler(t)
	rec := serve(h.start, StartPath, auth.Actor{UserID: 7, DeviceID: "dev", DeviceKind: auth.DeviceKindApp})
	if got := rec.Header().Get("Location"); got != publicBase+"/account?discord_role=failed#you" {
		t.Fatalf("location = %q", got)
	}
}

func TestCallbackOutcomes(t *testing.T) {
	cases := []struct {
		name  string
		query string
		actor auth.Actor
		state notify.ConnectState
		want  string
	}{
		{"linked", "?code=c", auth.Actor{UserID: 7}, notify.ConnectState{UserID: 7}, "linked"},
		{"another account's state", "?code=c", auth.Actor{UserID: 8}, notify.ConnectState{UserID: 7}, "failed"},
		{"signed out", "?code=c", auth.Actor{}, notify.ConnectState{UserID: 7}, "failed"},
		{"expired", "?code=c", auth.Actor{UserID: 7}, notify.ConnectState{UserID: 7, Expired: true}, "failed"},
		{"denied", "?error=access_denied", auth.Actor{UserID: 7}, notify.ConnectState{UserID: 7}, "denied"},
		{"discord error", "?error=server_error&code=c", auth.Actor{UserID: 7}, notify.ConnectState{UserID: 7}, "failed"},
		{"no code", "", auth.Actor{UserID: 7}, notify.ConnectState{UserID: 7}, "denied"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, states, store, discord := newHandler(t)
			discord.exchange, discord.userID = Tokens{Access: "a0", Refresh: "r0"}, "d7"
			rec := serve(func(w http.ResponseWriter, r *http.Request) { states.callback(w, r, tc.state) },
				"/v1/notifications/discord/callback"+tc.query, tc.actor)
			if got := rec.Header().Get("Location"); got != publicBase+"/account?discord_role="+tc.want+"#you" {
				t.Fatalf("location = %q, want outcome %s", got, tc.want)
			}
			_, linked := store.links[7]
			if linked != (tc.want == "linked") {
				t.Fatalf("linked = %v for outcome %s", linked, tc.want)
			}

		})
	}
}
