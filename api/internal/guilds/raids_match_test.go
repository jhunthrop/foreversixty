package guilds

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

const onyLogZone = "Onyxia's Lair"

// pastStart is a raid start two days back, on the minute.
func pastStart() time.Time { return time.Now().UTC().Add(-48 * time.Hour).Truncate(time.Minute) }

func (w raidWorld) raidAt(t *testing.T, zone string, starts time.Time) Raid {
	t.Helper()
	r, err := w.h.store.CreateRaid(t.Context(), RaidInput{GuildID: w.gid, Zone: zone, Title: "Night",
		StartsAt: starts, DurationMinutes: 180, CreatedBy: w.officer})
	if err != nil {
		t.Fatal(err)
	}
	return r
}

// oneFightNight seeds a report with a single Onyxia pull by the officer's character.
func (w raidWorld) oneFightNight(t *testing.T, id, zone string, at time.Time) {
	t.Helper()
	seedNight(t, w.h.pool, w.gid, id, zone, at, []seedFight{{encounterID: 1084, kill: true, duration: 5 * time.Minute,
		players: []seedPlayer{{key: officerKey, name: "Thornhide", class: "druid", role: "dps"}}}})
}

func (w raidWorld) raid(t *testing.T, id int64) Raid {
	t.Helper()
	r, ok, err := w.h.store.GetRaid(t.Context(), w.gid, id)
	if err != nil || !ok {
		t.Fatalf("get raid: %v %v", ok, err)
	}
	return r
}

func (w raidWorld) exec(t *testing.T, sql string) {
	t.Helper()
	if _, err := w.h.pool.Exec(t.Context(), sql); err != nil {
		t.Fatal(err)
	}
}

func TestAutoMatchTiesTheLogToItsRaid(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	raid := w.raidAt(t, raidZoneOny, start)
	w.oneFightNight(t, "log1", onyLogZone, start.Add(30*time.Minute))
	got, ok, err := w.h.store.AutoMatch(t.Context(), "log1")
	if err != nil || !ok {
		t.Fatalf("auto match = %v %v", ok, err)
	}
	if got.ID != raid.ID || got.Status != RaidDone || got.ReportID == nil || *got.ReportID != "log1" {
		t.Fatalf("matched raid = %+v", got)
	}
	var pre, by string
	var at *time.Time
	if err := w.h.pool.QueryRow(t.Context(),
		`select pre_done_status, matched_by, matched_at from guild_raids where id = $1`, raid.ID).Scan(&pre, &by, &at); err != nil {
		t.Fatal(err)
	}
	if pre != RaidPlanned || by != MatchedAuto || at == nil {
		t.Fatalf("pre %q by %q at %v", pre, by, at)
	}
}

func TestAutoMatchPicksOnlyTheNearestOfTwoOverlappingWindows(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	first := w.raidAt(t, raidZoneOny, start)
	second := w.raidAt(t, raidZoneOny, start.Add(8*time.Hour))
	w.oneFightNight(t, "log1", onyLogZone, start.Add(7*time.Hour+30*time.Minute))
	got, ok, err := w.h.store.AutoMatch(t.Context(), "log1")
	if err != nil || !ok || got.ID != second.ID {
		t.Fatalf("matched %+v %v %v, want raid %d", got, ok, err, second.ID)
	}
	if r := w.raid(t, first.ID); r.Status != RaidPlanned || r.ReportID != nil {
		t.Fatalf("first raid = %+v, want untouched", r)
	}
}

func TestAutoMatchRefusals(t *testing.T) {
	start := pastStart()
	cases := []struct {
		name string
		fix  string
		zone string
		at   time.Time
	}{
		{"zone differs", "", "Molten Core", start},
		{"before the window", "", onyLogZone, start.Add(-90 * time.Minute)},
		{"after the window", "", onyLogZone, start.Add(10 * time.Hour)},
		{"private log", `update reports set visibility = 'private' where id = 'log1'`, onyLogZone, start},
		{"flagged log", `update reports set flagged = 'tampered' where id = 'log1'`, onyLogZone, start},
		{"no guild", `update reports set guild_id = null where id = 'log1'`, onyLogZone, start},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			w := newRaidWorld(t, false)
			raid := w.raidAt(t, raidZoneOny, start)
			w.oneFightNight(t, "log1", c.zone, c.at)
			if c.fix != "" {
				w.exec(t, c.fix)
			}
			_, ok, err := w.h.store.AutoMatch(t.Context(), "log1")
			if err != nil || ok {
				t.Fatalf("auto match = %v %v, want no match and no error", ok, err)
			}
			if r := w.raid(t, raid.ID); r.Status != RaidPlanned {
				t.Fatalf("status = %s", r.Status)
			}
		})
	}
}

func TestAutoMatchLeavesAZoneWithNoLogNameAlone(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	w.raidAt(t, "raid:barrow-deeps", start)
	w.oneFightNight(t, "log1", onyLogZone, start)
	if _, ok, err := w.h.store.AutoMatch(t.Context(), "log1"); err != nil || ok {
		t.Fatalf("auto match = %v %v", ok, err)
	}
}

func TestAMatchedLogMatchesNoSecondRaid(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	w.raidAt(t, raidZoneOny, start)
	w.raidAt(t, raidZoneOny, start.Add(time.Hour))
	w.oneFightNight(t, "log1", onyLogZone, start.Add(30*time.Minute))
	if _, ok, err := w.h.store.AutoMatch(t.Context(), "log1"); err != nil || !ok {
		t.Fatalf("first = %v %v", ok, err)
	}
	if _, ok, err := w.h.store.AutoMatch(t.Context(), "log1"); err != nil || ok {
		t.Fatalf("second = %v %v, want no match", ok, err)
	}
}

func TestSetMatchRules(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	future := w.raidAt(t, raidZoneOny, time.Now().Add(72*time.Hour))
	raid := w.raidAt(t, raidZoneOny, start)
	other := w.raidAt(t, raidZoneOny, start.Add(-24*time.Hour))
	w.oneFightNight(t, "log1", onyLogZone, start)
	w.oneFightNight(t, "log2", onyLogZone, start.Add(time.Hour))
	ctx := t.Context()
	if _, err := w.h.store.SetMatch(ctx, w.gid, future.ID, "log1", MatchedHand); !errors.Is(err, ErrRaidNotEnded) {
		t.Fatalf("future raid: %v", err)
	}
	if _, err := w.h.store.SetMatch(ctx, w.gid, raid.ID, "nope", MatchedHand); !errors.Is(err, ErrNotFound) {
		t.Fatalf("unknown report: %v", err)
	}
	if _, err := w.h.store.SetMatch(ctx, w.gid, raid.ID, "log1", MatchedHand); err != nil {
		t.Fatal(err)
	}
	if _, err := w.h.store.SetMatch(ctx, w.gid, other.ID, "log1", MatchedHand); !errors.Is(err, ErrReportMatched) {
		t.Fatalf("log held by another raid: %v", err)
	}
	r, err := w.h.store.SetMatch(ctx, w.gid, raid.ID, "log2", MatchedHand)
	if err != nil || r.Status != RaidDone || *r.ReportID != "log2" {
		t.Fatalf("re-match = %+v %v", r, err)
	}
	var pre string
	if err := w.h.pool.QueryRow(ctx, `select pre_done_status from guild_raids where id = $1`, raid.ID).Scan(&pre); err != nil {
		t.Fatal(err)
	}
	if pre != RaidPlanned {
		t.Fatalf("pre_done_status = %q after a re-match, want planned", pre)
	}
}

func TestClearMatchRestoresStatusRejectsTheLogAndDropsAttendance(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	raid := w.raidAt(t, raidZoneOny, start)
	if _, err := w.h.store.SetRaidStatus(t.Context(), w.gid, raid.ID, []string{RaidPlanned}, RaidLocked); err != nil {
		t.Fatal(err)
	}
	w.oneFightNight(t, "log1", onyLogZone, start)
	ctx := t.Context()
	if _, err := w.h.store.SetMatch(ctx, w.gid, raid.ID, "log1", MatchedHand); err != nil {
		t.Fatal(err)
	}
	if rows, _ := w.h.store.AttendanceRows(ctx, raid.ID); len(rows) == 0 {
		t.Fatal("no attendance rows after the match")
	}
	r, err := w.h.store.ClearMatch(ctx, w.gid, raid.ID, w.officer)
	if err != nil || r.Status != RaidLocked || r.ReportID != nil {
		t.Fatalf("clear = %+v %v", r, err)
	}
	if rows, _ := w.h.store.AttendanceRows(ctx, raid.ID); len(rows) != 0 {
		t.Fatalf("attendance after clear = %d rows", len(rows))
	}
	if _, ok, err := w.h.store.AutoMatch(ctx, "log1"); err != nil || ok {
		t.Fatalf("auto match after clear = %v %v, want none", ok, err)
	}
	if _, err := w.h.store.ClearMatch(ctx, w.gid, raid.ID, w.officer); !errors.Is(err, ErrRaidStatus) {
		t.Fatalf("clear twice: %v", err)
	}
}

func TestMatchCandidates(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	raid := w.raidAt(t, raidZoneOny, start)
	other := w.raidAt(t, raidZoneOny, start.Add(-24*time.Hour))
	w.oneFightNight(t, "near", onyLogZone, start.Add(20*time.Minute))
	w.oneFightNight(t, "far", onyLogZone, start.Add(-20*time.Hour))
	w.oneFightNight(t, "held", onyLogZone, start.Add(-24*time.Hour))
	w.oneFightNight(t, "elsewhere", "Molten Core", start)
	w.oneFightNight(t, "old", onyLogZone, time.Now().Add(-20*24*time.Hour))
	w.oneFightNight(t, "hidden", onyLogZone, start)
	w.exec(t, `update reports set visibility = 'private' where id = 'hidden'`)
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, other.ID, "held", MatchedHand); err != nil {
		t.Fatal(err)
	}
	res, err := w.h.store.MatchCandidates(t.Context(), w.gid, raid.ID, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	got := res.Rows
	if !res.ZoneKnown || len(got) != 3 || got[0].ReportID != "near" || got[1].ReportID != "far" || got[2].ReportID != "held" {
		t.Fatalf("order/zone_known: %+v known=%v", got, res.ZoneKnown)
	}
	byID := map[string]Candidate{}
	for _, c := range got {
		byID[c.ReportID] = c
	}
	if len(got) != 3 || byID["near"].ReportID == "" || byID["far"].ReportID == "" || byID["held"].ReportID == "" {
		t.Fatalf("candidates = %+v", got)
	}
	if !byID["near"].Closest || byID["far"].Closest || byID["held"].Closest {
		t.Fatalf("closest flags: %+v", got)
	}
	if m := byID["held"].MatchedRaid; m == nil || m.ID != other.ID {
		t.Fatalf("held matched_raid = %+v", m)
	}
	n := byID["near"]
	if n.Pulls != 1 || n.Kills != 1 || n.DurationMS != 300000 || n.ZoneName != onyLogZone || n.MatchedRaid != nil {
		t.Fatalf("near = %+v", n)
	}
}

func TestReportCompletedMatchesAndComputesAttendance(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	raid := w.raidAt(t, raidZoneOny, start)
	w.oneFightNight(t, "log1", onyLogZone, start.Add(10*time.Minute))
	w.h.svc.ReportCompleted(t.Context(), "log1")
	if r := w.raid(t, raid.ID); r.Status != RaidDone {
		t.Fatalf("status = %s", r.Status)
	}
	rows, err := w.h.store.AttendanceRows(t.Context(), raid.ID)
	if err != nil || len(rows) != 1 || rows[0].CharacterKey != officerKey || rows[0].Outcome != OutcomePresent {
		t.Fatalf("rows = %+v %v", rows, err)
	}
	w.h.svc.ReportCompleted(t.Context(), "unknown-report") // best effort: no panic
}

func TestMatchCandidatesForAnUnknownZoneListEveryZone(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	raid := w.raidAt(t, "raid:barrow-deeps", start)
	w.oneFightNight(t, "ony", onyLogZone, start)
	w.oneFightNight(t, "mc", "Molten Core", start.Add(time.Hour))
	res, err := w.h.store.MatchCandidates(t.Context(), w.gid, raid.ID, time.Now())
	if err != nil || res.ZoneKnown || len(res.Rows) != 2 || !res.Rows[0].Closest || res.Rows[0].ReportID != "ony" {
		t.Fatalf("result = %+v %v", res, err)
	}
}

func TestMatchCandidatesCapAtTwentyNearestFirst(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	raid := w.raidAt(t, raidZoneOny, start)
	for i := 0; i < candidateCap+3; i++ {
		w.oneFightNight(t, fmt.Sprintf("log%02d", i), onyLogZone, start.Add(time.Duration(i)*time.Minute))
	}
	res, err := w.h.store.MatchCandidates(t.Context(), w.gid, raid.ID, time.Now())
	if err != nil || len(res.Rows) != candidateCap || res.Rows[0].ReportID != "log00" || !res.Rows[0].Closest {
		t.Fatalf("rows = %d %v", len(res.Rows), err)
	}
	for _, c := range res.Rows[1:] {
		if c.Closest {
			t.Fatal("more than one closest")
		}
	}
}

func TestAutoMatchBreaksATieToTheLowerRaidID(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	first := w.raidAt(t, raidZoneOny, start)
	second := w.raidAt(t, raidZoneOny, start.Add(2*time.Hour))
	w.oneFightNight(t, "log1", onyLogZone, start.Add(time.Hour))
	got, ok, err := w.h.store.AutoMatch(t.Context(), "log1")
	if err != nil || !ok || got.ID != first.ID {
		t.Fatalf("matched %+v %v %v, want raid %d (not %d)", got, ok, err, first.ID, second.ID)
	}
}
