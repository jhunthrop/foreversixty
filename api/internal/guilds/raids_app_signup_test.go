package guilds

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// appSignup PUTs a sign-up as the app does: a device token's actor, through a mux mounted
// with the service's AppSignups as it stands.
func (w raidWorld) appSignup(t *testing.T, raidID int64, key, body string) int {
	t.Helper()
	return w.deviceSignup(t, raidID, key, body, auth.DeviceKindApp)
}

// deviceSignup is appSignup as a device of the given kind.
func (w raidWorld) deviceSignup(t *testing.T, raidID int64, key, body, kind string) int {
	t.Helper()
	mux := http.NewServeMux()
	Mount(mux, w.h.svc, 0)
	req := httptest.NewRequest(http.MethodPut, w.path("/raids/%d/signups/%s", raidID, key), strings.NewReader(body))
	req = req.WithContext(auth.WithActor(req.Context(), auth.Actor{UserID: w.member, Role: "user", Method: "device", DeviceID: "phone", DeviceKind: kind}))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	return rec.Code
}

// A sign-up from the app is the site's PUT: it marks the raid's Discord post out of date, as
// a Discord press does (both run signUp, which marks the raid), so the drainer edits it.
func TestAnAppSignupRefreshesTheDiscordPost(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	marker := &fakeMarker{}
	w.h.svc.RaidPosts = marker
	w.h.svc.AppSignups = true
	if code := w.appSignup(t, raid.ID, memberKeyA, `{"state":"in"}`); code != http.StatusOK {
		t.Fatalf("app sign-up = %d", code)
	}
	if marker.count() != 1 || marker.ids[0] != raid.ID {
		t.Fatalf("marks = %v", marker.ids)
	}
	// The bot's press runs the same SignUp and marks the same way.
	marker.reset()
	if _, err := w.svcSignUp(w.cmd(raid.ID, w.member, memberKeyA, SignupLate)); err != nil {
		t.Fatal(err)
	}
	if marker.count() != 1 || marker.ids[0] != raid.ID {
		t.Fatalf("press marks = %v", marker.ids)
	}
}

func TestAnAppTokenSignsUpOnlyWithTheAppFlag(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	if code := w.appSignup(t, raid.ID, memberKeyA, `{"state":"in"}`); code != http.StatusForbidden {
		t.Fatalf("flag off: app sign-up = %d", code)
	}
	w.h.svc.AppSignups = true
	if code := w.appSignup(t, raid.ID, memberKeyA, `{"state":"in"}`); code != http.StatusOK {
		t.Fatalf("flag on: app sign-up = %d", code)
	}
	// A companion's token uploads; it never signs anyone up.
	if code := w.deviceSignup(t, raid.ID, memberKeyA, `{"state":"in"}`, auth.DeviceKindCompanion); code != http.StatusUnauthorized {
		t.Fatalf("flag on: companion sign-up = %d", code)
	}
}
