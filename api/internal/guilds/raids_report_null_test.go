package guilds

import "testing"

// fight_metrics.class and the metrics are nullable: a player with neither must not 500 the
// raid page, and a missing metric is no change.
func TestAReportRendersALogPlayerWithNullClassAndMetrics(t *testing.T) {
	w, night, _ := reportNight(t)
	w.exec(t, `update fight_metrics set class = null where report_id = 'night' and player_key = '`+drifterKey+`'`)
	w.exec(t, `update fight_metrics set metric_dps = null where report_id = 'night' and player_key = '`+officerKey+`'`)
	w.as(w.officer)
	rep := w.getRaid(t, night.ID).Report
	if rep == nil {
		t.Fatal("report is null")
	}
	var found bool
	for _, d := range rep.Deaths {
		if d.CharacterKey == drifterKey {
			found = true
			if d.Class != "" {
				t.Fatalf("drifter = %+v, want an empty class", d)
			}
		}
	}
	if !found {
		t.Fatalf("the null-class player is missing from deaths: %+v", rep.Deaths)
	}
	for _, m := range rep.Moves.Rows {
		if m.CharacterKey == officerKey {
			t.Fatalf("a null metric became a move: %+v", m)
		}
	}
}

// A raider who raided as a healer last week and a dps now has no comparable metric.
func TestARoleSwitchBetweenRaidsIsNotAMove(t *testing.T) {
	w, night, _ := reportNight(t)
	w.exec(t, `update fight_metrics set role = 'healer', metric_hps = 400 where report_id = 'night' and kill and player_key = '`+officerKey+`'`)
	w.exec(t, `update fight_metrics set role = 'dps' where report_id = 'night' and kill and player_key = '`+member2Key+`'`)
	w.as(w.officer)
	for _, m := range w.getRaid(t, night.ID).Report.Moves.Rows {
		if m.CharacterKey == officerKey || m.CharacterKey == member2Key {
			t.Fatalf("a role switch produced a move: %+v", m)
		}
	}
}
