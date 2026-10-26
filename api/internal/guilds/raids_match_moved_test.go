package guilds

import "testing"

// A report the officer PATCH moves to another guild leaves its old guild's raid: the match is
// cleared through ClearMatch, so the old raid never re-picks the log.
func TestMovingAMatchedReportToAnotherGuildClearsItsMatch(t *testing.T) {
	w, raid := matchedNight(t)
	other := seedGuild(t, w.h.pool, "Elsewhere")
	if _, err := w.h.pool.Exec(t.Context(), `update reports set guild_id = $1 where id = 'night'`, other); err != nil {
		t.Fatal(err)
	}
	w.h.svc.ReportGuildMoved(t.Context(), "night", w.officer)
	r := w.raid(t, raid.ID)
	if r.ReportID != nil || r.Status != RaidPlanned {
		t.Fatalf("raid = %+v, want the log released and the status restored", r)
	}
	var rejectedBy int64
	if err := w.h.pool.QueryRow(t.Context(),
		`select rejected_by from guild_raid_rejected_reports where raid_id = $1 and report_id = 'night'`, raid.ID).
		Scan(&rejectedBy); err != nil || rejectedBy != w.officer {
		t.Fatalf("rejected row = %d %v, want one by the officer", rejectedBy, err)
	}
}

// A move within the same guild (or a guild set where none was) leaves the match alone.
func TestAGuildPatchToTheSameGuildKeepsTheMatch(t *testing.T) {
	w, raid := matchedNight(t)
	w.h.svc.ReportGuildMoved(t.Context(), "night", w.officer)
	if r := w.raid(t, raid.ID); r.ReportID == nil || *r.ReportID != "night" {
		t.Fatalf("raid = %+v, want the match kept", r)
	}
}
