package guilds

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

func confirmRaid() RaidView {
	tz := "America/Chicago"
	return RaidView{ZoneName: "Onyxia's Lair", StartsAt: time.Date(2027, 1, 7, 2, 0, 0, 0, time.UTC), GuildTimezone: &tz}
}

func minutes(m int16) *int16 { return &m }

// The confirmation the app posts after a notification answer names the state, the zone, the
// character and, for a Late with minutes, the arrival in the guild's raid time zone.
func TestSignupConfirmation(t *testing.T) {
	moved := "Sunderfel"
	cases := []struct {
		name, state string
		late        *int16
		moved       *string
		title, body string
	}{
		{"in", SignupIn, nil, nil, "You're In for Onyxia's Lair", "Kraggor, pull at 8 pm CST."},
		{"late with minutes", SignupLate, minutes(30), nil, "You're Late (30 min) for Onyxia's Lair", "Kraggor, arriving 8:30 pm CST."},
		{"late", SignupLate, nil, nil, "You're Late for Onyxia's Lair", "Kraggor, pull at 8 pm CST."},
		{"bench", SignupBench, nil, nil, "You're on the Bench for Onyxia's Lair", "Kraggor."},
		{"out", SignupOut, nil, nil, "You're Out for Onyxia's Lair", "Kraggor."},
		{"moved", SignupIn, nil, &moved, "You're In for Onyxia's Lair", "Kraggor, pull at 8 pm CST. Sunderfel moved to Bench."},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := signupConfirmation(confirmRaid(), "Kraggor", c.state, c.late, c.moved)
			if got.Title != c.title || got.Body != c.body {
				t.Fatalf("got %+v, want %q / %q", got, c.title, c.body)
			}
		})
	}
}

// A guild with no time zone words the time in UTC.
func TestSignupConfirmationWithoutAZone(t *testing.T) {
	raid := confirmRaid()
	raid.GuildTimezone = nil
	if got := signupConfirmation(raid, "Kraggor", SignupLate, minutes(15), nil); got.Body != "Kraggor, arriving 2:15 am UTC." {
		t.Fatalf("body = %q", got.Body)
	}
}

// The PUT's answer carries the confirmation, so a background action has its line to post.
func TestSignupAnswerCarriesTheConfirmation(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.h.svc.AppSignups = true
	mux := http.NewServeMux()
	Mount(mux, w.h.svc, 0)
	req := httptest.NewRequest(http.MethodPut, w.path("/raids/%d/signups/%s", raid.ID, memberKeyA),
		strings.NewReader(`{"state":"late","late_minutes":60}`))
	req = req.WithContext(auth.WithActor(req.Context(), auth.Actor{UserID: w.member, Role: "user", Method: "device", DeviceID: "phone", DeviceKind: auth.DeviceKindApp}))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("put = %d: %s", rec.Code, rec.Body)
	}
	var env struct {
		Data struct {
			Confirmation SignupConfirmation `json:"confirmation"`
		} `json:"data"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
		t.Fatal(err)
	}
	c := env.Data.Confirmation
	if c.Title != "You're Late (60 min) for Onyxia's Lair" || !strings.HasPrefix(c.Body, "Kraggor, arriving ") {
		t.Fatalf("confirmation = %+v", c)
	}
}
