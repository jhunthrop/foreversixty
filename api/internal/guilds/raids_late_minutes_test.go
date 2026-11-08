package guilds

import (
	"net/http"
	"testing"
)

// signupRow is the raid's listed row for key, failing when the raid lists none.
func signupRow(t *testing.T, v RaidView, key string) SignupView {
	t.Helper()
	for _, sg := range v.Signups {
		if sg.CharacterKey == key {
			return sg
		}
	}
	t.Fatalf("no sign-up row for %s in %+v", key, v.Signups)
	return SignupView{}
}

// lateRow reads key's row as the member after an app answer of body, which must stand.
func (w raidWorld) lateRow(t *testing.T, raidID int64, body string) SignupView {
	t.Helper()
	if code := w.appSignup(t, raidID, memberKeyA, body); code != http.StatusOK {
		t.Fatalf("%s = %d", body, code)
	}
	w.as(w.member)
	return signupRow(t, w.getRaid(t, raidID), memberKeyA)
}

// A notification's Late chip answers Late with its minutes; the raid lists them, and any
// other answer, or a Late with no chip, clears them (spec 2026-11-08-app-flows F1).
func TestLateCarriesItsMinutes(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.h.svc.AppSignups = true
	if row := w.lateRow(t, raid.ID, `{"state":"late","late_minutes":30}`); row.State != SignupLate || row.LateMinutes == nil || *row.LateMinutes != 30 {
		t.Fatalf("late 30 row = %+v", row)
	}
	// The same answer again (an action retried from the offline queue) keeps the minutes.
	if row := w.lateRow(t, raid.ID, `{"state":"late","late_minutes":30}`); row.LateMinutes == nil || *row.LateMinutes != 30 {
		t.Fatalf("retried row = %+v", row)
	}
	if row := w.lateRow(t, raid.ID, `{"state":"late"}`); row.State != SignupLate || row.LateMinutes != nil {
		t.Fatalf("no-chip row = %+v", row)
	}
	w.lateRow(t, raid.ID, `{"state":"late","late_minutes":15}`)
	if row := w.lateRow(t, raid.ID, `{"state":"in"}`); row.State != SignupIn || row.LateMinutes != nil {
		t.Fatalf("in row = %+v", row)
	}
}

func TestLateMinutesAreValidated(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.member)
	for _, body := range []string{
		`{"state":"late","late_minutes":20}`,
		`{"state":"late","late_minutes":0}`,
		`{"state":"in","late_minutes":15}`,
	} {
		res := w.putSignup(raid.ID, memberKeyA, body)
		if res.StatusCode != http.StatusUnprocessableEntity {
			t.Fatalf("%s = %d: %s", body, res.StatusCode, readBody(res))
		}
	}
}
