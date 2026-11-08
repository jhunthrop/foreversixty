package guilds

import (
	"context"
	"slices"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/events"
)

// fakeNudges answers scripted unanswered accounts and their settings.
type fakeNudges struct {
	users []int64
	prefs map[int64]events.Prefs
}

func (f fakeNudges) Recipients(_ context.Context, a events.Audience) ([]int64, error) {
	if a.Kind != events.ToUnanswered {
		return nil, nil
	}
	return f.users, nil
}

func (f fakeNudges) PrefsOf(_ context.Context, ids []int64) (map[int64]events.Prefs, error) {
	out := map[int64]events.Prefs{}
	for _, id := range ids {
		if p, ok := f.prefs[id]; ok {
			out[id] = p
		} else {
			out[id] = events.DefaultPrefs()
		}
	}
	return out, nil
}

func prefsAt(tz string, reminder int) events.Prefs {
	p := events.DefaultPrefs()
	p.TimeZone, p.RaidReminder = tz, reminder
	return p
}

// The nudge is due at the account's reminder hour on the raid's day in its own zone; a reminder
// at or after the start falls back to two hours before it.
func TestNudgeDue(t *testing.T) {
	start := time.Date(2026, 11, 5, 2, 0, 0, 0, time.UTC) // Wed 4 Nov, 8:00 pm CST
	cases := []struct {
		p    events.Prefs
		want time.Time
	}{
		{prefsAt("America/Chicago", 15*60), time.Date(2026, 11, 4, 21, 0, 0, 0, time.UTC)},
		{prefsAt("America/Chicago", 9*60+15), time.Date(2026, 11, 4, 15, 15, 0, 0, time.UTC)},
		{prefsAt("UTC", 15*60), start.Add(-nudgeFallback)},
		{prefsAt("Not/AZone", 0), time.Date(2026, 11, 5, 0, 0, 0, 0, time.UTC)},
	}
	for _, c := range cases {
		if got := nudgeDue(start, c.p); !got.Equal(c.want) {
			t.Errorf("%s %d: %v, want %v", c.p.TimeZone, c.p.RaidReminder, got, c.want)
		}
	}
}

// The run nudges each unanswered account whose reminder came due, and only those.
func TestRunEventScheduleNudgesAtEachReminder(t *testing.T) {
	w, pub := eventsWorld(t)
	// Tomorrow at 20:00 UTC; the run is at 14:00 UTC that day.
	start := time.Now().UTC().Truncate(24 * time.Hour).Add(44 * time.Hour)
	raid := w.createRaids(t, raidBody(start.Format(time.RFC3339), 0))[0]
	pub.waitFor(t, events.RaidPosted, 1)
	now := raid.StartsAt.Add(-6 * time.Hour)
	w.h.svc.Nudges = fakeNudges{users: []int64{w.member, w.member2}, prefs: map[int64]events.Prefs{
		w.member:  prefsAt("UTC", 14*60),
		w.member2: prefsAt("UTC", 16*60),
	}}
	run, err := w.h.svc.RunEventSchedule(context.Background(), now)
	if err != nil || run.Nudged != 1 {
		t.Fatalf("run %+v %v", run, err)
	}
	e := pub.ofKind(events.RaidSignupNudge)
	if len(e) != 1 || !slices.Equal(e[0].Audience.UserIDs, []int64{w.member}) || len(e[0].Actions) != 4 {
		t.Fatalf("nudge %+v", e)
	}
}
