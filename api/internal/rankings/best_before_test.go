package rankings

import (
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/engine"
)

func TestBestBeforeReadsTheEarlierKillsOnly(t *testing.T) {
	h := newHarness(t)
	reportID, _, key := h.seedRankedFight(t)
	var dps float64
	if err := h.pool.QueryRow(t.Context(),
		`select metric_dps::float8 from fight_metrics where report_id = $1 and player_key = $2`,
		reportID, key).Scan(&dps); err != nil {
		t.Fatal(err)
	}
	later := engine.FixtureBase.Add(time.Hour)

	best, found, err := h.store.BestBefore(t.Context(), key, executionEncounterID, 8, MetricDPS, later, "other")
	if err != nil || !found || best != dps {
		t.Errorf("best = %v, %v, %v; want the earlier kill's %v", best, found, err, dps)
	}
	if _, found, _ := h.store.BestBefore(t.Context(), key, executionEncounterID, 8, MetricDPS, later, reportID); found {
		t.Error("the report being read is never its own previous best")
	}
	if _, found, _ := h.store.BestBefore(t.Context(), key, executionEncounterID, 8, MetricDPS, engine.FixtureBase, "other"); found {
		t.Error("a kill at the same moment is not earlier")
	}
	if _, found, _ := h.store.BestBefore(t.Context(), key, executionEncounterID, 8, MetricHPS, later, "other"); !found {
		t.Error("the healing column is read for hps")
	}
}
