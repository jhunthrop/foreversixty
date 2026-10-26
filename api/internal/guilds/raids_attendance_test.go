package guilds

import (
	"errors"
	"testing"
	"time"
)

const (
	ghostKey    = "us/hardcore/ghost"
	drifterKey  = "us/hardcore/drifter"
	strangerKey = "us/hardcore/stranger"
)

func (w raidWorld) signUp(t *testing.T, raidID int64, key, name, class string, owner int64, state string) {
	t.Helper()
	if _, err := w.h.store.UpsertSignup(t.Context(), Signup{RaidID: raidID, CharacterKey: key, Name: name, Class: class,
		Role: RoleDPS, State: state, SignedBy: owner, ActedBy: owner, AsOfficer: true}); err != nil {
		t.Fatal(err)
	}
}

func rowsByKey(rows []AttendanceRow) map[string]AttendanceRow {
	out := map[string]AttendanceRow{}
	for _, r := range rows {
		out[r.CharacterKey] = r
	}
	return out
}

func (w raidWorld) attendance(t *testing.T, raidID int64) []AttendanceRow {
	t.Helper()
	rows, err := w.h.store.AttendanceRows(t.Context(), raidID)
	if err != nil {
		t.Fatal(err)
	}
	return rows
}

// attendanceNight builds the raid, its sign-ups and a four-pull night after a trash pull.
func attendanceNight(t *testing.T) (raidWorld, Raid) {
	t.Helper()
	w := newRaidWorld(t, false)
	ghost := w.seedRaider(t, "ghost@example.com", "member", raidChar{ghostKey, "Ghost", "rogue", true})
	w.seedRaider(t, "drifter@example.com", "member", raidChar{drifterKey, "Drifter", "hunter", true})
	start := pastStart()
	raid := w.raidAt(t, raidZoneOny, start)
	w.signUp(t, raid.ID, officerKey, "Thornhide", "druid", w.officer, SignupIn)
	w.signUp(t, raid.ID, memberKeyB, "Sunderfel", "mage", w.member, SignupBench)
	w.signUp(t, raid.ID, memberKeyC, "Unverifiedalt", "rogue", w.member, SignupBench)
	w.signUp(t, raid.ID, memberKeyA, "Kraggor", "warrior", w.member, SignupIn)
	w.signUp(t, raid.ID, member2Key, "Mirelle", "priest", w.member2, SignupLate)
	w.signUp(t, raid.ID, ghostKey, "Ghost", "rogue", ghost, SignupIn)
	pl := func(key, name, class string, deaths int) seedPlayer {
		return seedPlayer{key: key, name: name, class: class, role: "dps", deaths: deaths}
	}
	core := []seedPlayer{pl(officerKey, "Thornhide", "druid", 1), pl(memberKeyB, "Sunderfel", "mage", 0),
		pl(drifterKey, "Drifter", "hunter", 2), pl(strangerKey, "Stranger", "mage", 0)}
	withMirelle := append(append([]seedPlayer{}, core...), pl(member2Key, "Mirelle", "priest", 0))
	m := time.Minute
	seedNight(t, w.h.pool, w.gid, "night", onyLogZone, start.Add(time.Hour), []seedFight{
		{offset: 0, duration: 2 * m},
		{encounterID: 1084, offset: 2 * m, duration: 5 * m, players: core},
		{encounterID: 1084, offset: 12 * m, duration: 5 * m, players: core},
		{encounterID: 1084, offset: 22 * m, duration: 5 * m, players: core},
		{encounterID: 1084, kill: true, offset: 33 * m, duration: 5 * m, players: withMirelle},
	})
	return w, raid
}

func TestAttendanceFollowsTheSignupsAgainstTheLog(t *testing.T) {
	w, raid := attendanceNight(t)
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, raid.ID, "night", MatchedHand); err != nil {
		t.Fatal(err)
	}
	rows := w.attendance(t, raid.ID)
	by := rowsByKey(rows)
	if len(rows) != 7 {
		t.Fatalf("rows = %d: %+v", len(rows), rows)
	}
	if r := by[officerKey]; r.Outcome != OutcomePresent || r.PullsPresent != 4 || r.PullsTotal != 4 ||
		r.FirstSeenOffsetS == nil || *r.FirstSeenOffsetS != 120 || *r.Signed != SignupIn {
		t.Fatalf("thornhide = %+v", r)
	}
	k := by[memberKeyA]
	if k.Outcome != OutcomePresent || k.CameAs == nil || *k.CameAs != "Sunderfel" || k.PullsPresent != 4 {
		t.Fatalf("kraggor (signed in, came on Sunderfel) = %+v", k)
	}
	if r := by[memberKeyB]; r.Outcome != OutcomePresent || *r.Signed != SignupBench {
		t.Fatalf("sunderfel = %+v", r)
	}
	if r := by[ghostKey]; r.Outcome != OutcomeNoShow || r.PullsPresent != 0 || r.FirstSeenOffsetS != nil {
		t.Fatalf("ghost = %+v", r)
	}
	if r := by[drifterKey]; r.Outcome != OutcomePresent || r.Signed != nil || r.Role != nil {
		t.Fatalf("drifter (unsigned) = %+v", r)
	}
	if r := by[member2Key]; r.Outcome != OutcomeLate || r.PullsPresent != 1 || *r.FirstSeenOffsetS != 33*60 {
		t.Fatalf("mirelle (late) = %+v", r)
	}
	if r := by[memberKeyC]; r.Outcome != OutcomeAbsent {
		t.Fatalf("unverifiedalt (bench, not there) = %+v", r)
	}
	if _, ok := by[strangerKey]; ok {
		t.Fatal("a character outside the roster got a row")
	}
}

func TestARematchRecomputesAndKeepsExcused(t *testing.T) {
	w, raid := attendanceNight(t)
	ctx := t.Context()
	if _, err := w.h.store.SetMatch(ctx, w.gid, raid.ID, "night", MatchedHand); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{ghostKey, drifterKey} {
		if err := w.h.store.SetExcused(ctx, raid.ID, key, true); err != nil {
			t.Fatal(err)
		}
	}
	// A second log where Ghost shows up and Drifter does not.
	seedNight(t, w.h.pool, w.gid, "night2", onyLogZone, raid.StartsAt.Add(2*time.Hour), []seedFight{{encounterID: 1084,
		kill: true, duration: 5 * time.Minute,
		players: []seedPlayer{{key: ghostKey, name: "Ghost", class: "rogue", role: "dps"}}}})
	if _, err := w.h.store.SetMatch(ctx, w.gid, raid.ID, "night2", MatchedHand); err != nil {
		t.Fatal(err)
	}
	by := rowsByKey(w.attendance(t, raid.ID))
	if r := by[ghostKey]; r.Outcome != OutcomePresent || !r.Excused || r.PullsTotal != 1 {
		t.Fatalf("ghost after re-match = %+v", r)
	}
	if _, ok := by[drifterKey]; ok {
		t.Fatal("drifter kept a row though absent from the new log and unsigned")
	}
	if r := by[officerKey]; r.Outcome != OutcomeNoShow || r.Excused {
		t.Fatalf("thornhide after re-match = %+v", r)
	}
}

func TestSetExcusedNeedsARow(t *testing.T) {
	w, raid := attendanceNight(t)
	if err := w.h.store.SetExcused(t.Context(), raid.ID, ghostKey, true); !errors.Is(err, ErrNotFound) {
		t.Fatalf("excuse with no rows = %v", err)
	}
}

func TestAPrivateLogShowsNobody(t *testing.T) {
	w, raid := attendanceNight(t)
	w.exec(t, `update reports set visibility = 'private' where id = 'night'`)
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, raid.ID, "night", MatchedHand); err != nil {
		t.Fatal(err)
	}
	for _, r := range w.attendance(t, raid.ID) {
		if r.Outcome == OutcomePresent || r.Outcome == OutcomeLate {
			t.Fatalf("%s present in a private log", r.CharacterKey)
		}
	}
}
