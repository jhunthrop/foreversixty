package rankings

import (
	"encoding/json"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/engine"
	"github.com/jhunthrop/foreversixty/api/internal/metrics"
)

// Task L6.12: a report flagged tampered or unverified counts nowhere a player or guild looks.
// The public guild page (GET /v1/guilds/{region}/{ruleset}/{name}) reads its progression, its
// roster bests and its report list through presence.EvidenceReport.

func (h *harness) flag(reportID string) {
	h.t.Helper()
	if _, err := h.pool.Exec(h.t.Context(), `update reports set flagged = 'tampered' where id = $1`, reportID); err != nil {
		h.t.Fatal(err)
	}
}

// A flagged report's kill is not progression.
func TestGuildProgressionIgnoresAFlaggedReport(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedReport("report-clean")
	h.seedFightRecord("report-clean", 1, engine.FixtureBase, false)
	h.seedReport("report-flagged")
	h.seedFightRecord("report-flagged", 1, engine.FixtureBase, true)
	h.flag("report-flagged")

	var g Guild
	h.data(h.get("/v1/guilds/us/hardcore/forever-sixty"), &g)
	if len(g.Progression) != 1 || g.Progression[0].Kills != 0 || g.Progression[0].PullCount != 1 ||
		g.Progression[0].FirstKillAt != nil {
		t.Fatalf("progression = %+v, want the clean report's one wipe only (the flagged kill is no kill)", g.Progression)
	}
}

// A flagged report's parse is no roster best.
func TestGuildRosterBestIgnoresAFlaggedReport(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedReport("report-clean")
	h.seedFight("report-clean", 1, engine.FixtureBase, nil)
	var clean Guild
	h.data(h.get("/v1/guilds/us/hardcore/forever-sixty"), &clean)
	if len(clean.RosterBest) == 0 {
		t.Fatal("the clean report gives no roster best")
	}

	h.seedReport("report-flagged")
	h.seedFight("report-flagged", 1, engine.FixtureBase, func(rows []metrics.Row) {
		for i := range rows {
			rows[i].MetricDPS *= 10
			rows[i].MetricHPS *= 10
			rows[i].DamageTaken *= 10
		}
	})
	h.flag("report-flagged")

	var g Guild
	h.data(h.get("/v1/guilds/us/hardcore/forever-sixty"), &g)
	got, _ := json.Marshal(g.RosterBest)
	want, _ := json.Marshal(clean.RosterBest)
	if string(got) != string(want) {
		t.Fatalf("roster bests = %s\nwant the clean report's alone: %s", got, want)
	}
}

// A flagged report is not on the guild page's report list.
func TestGuildReportListLeavesOutAFlaggedReport(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedReport("report-clean")
	h.seedReport("report-flagged")
	h.flag("report-flagged")

	var g Guild
	h.data(h.get("/v1/guilds/us/hardcore/forever-sixty"), &g)
	if len(g.Reports) != 1 || g.Reports[0].ID != "report-clean" {
		t.Fatalf("reports = %+v, want only report-clean", g.Reports)
	}
}
