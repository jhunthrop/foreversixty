// api/internal/guilds/parse_source_test.go
package guilds

import (
	"context"
	"testing"
)

// Guild round 2, WP-13: the Raids tab, the Progression tab and the roster read one figure for
// the same raider. The clean night of seedL612Guild gets a whelp-pack kill (no encounter) on
// which the member posts 500 DPS against their 100 on Onyxia: none of the three may show 500.
func TestTheThreeGuildFiguresReadTheBossKillOnly(t *testing.T) {
	g := seedL612Guild(t)
	ctx := context.Background()
	if _, err := g.pool.Exec(ctx,
		`insert into fights (report_id, fight_index, name, encounter_id, kill, start_ms, players)
		 select id, 1, 'Onyxian Whelp', null, true, 0, '{}'::text[] from reports where id = $1`,
		g.cleanReport); err != nil {
		t.Fatal(err)
	}
	if _, err := g.pool.Exec(ctx,
		`insert into fight_metrics (report_id, fight_index, player_key, player_name, class, role, metric_dps, kill, fought_at)
		 select report_id, 1, player_key, player_name, class, role, 500, true, fought_at
		 from fight_metrics where report_id = $1 and fight_index = 0`, g.cleanReport); err != nil {
		t.Fatal(err)
	}

	page, err := g.store.Raids(ctx, g.guildID, g.userID, true, RaidsDefaultLimit, nil)
	if err != nil {
		t.Fatal(err)
	}
	var raidsTop float64 = -1
	for _, row := range page.Rows {
		if row.ID == g.cleanReport && row.TopDPS != nil {
			raidsTop = row.TopDPS.Value
		}
	}
	view, err := g.store.Progression(ctx, g.guildID, true)
	if err != nil {
		t.Fatal(err)
	}
	var progressionBest float64 = -1
	for _, encounter := range view.Encounters {
		if dps := encounter.BestByRole["dps"]; dps != nil {
			progressionBest = dps.Value
		}
	}
	row := g.rosterRow(t)
	var rosterBest float64 = -1
	if row.BestParse != nil {
		rosterBest = row.BestParse.Value
	}
	if raidsTop != 100 || progressionBest != 100 || rosterBest != 100 {
		t.Fatalf("raids %v, progression %v, roster %v: want the Onyxia kill's 100 on all three, never the whelps' 500",
			raidsTop, progressionBest, rosterBest)
	}
}
