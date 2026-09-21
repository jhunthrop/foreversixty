package rankings

import (
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/engine"
)

// TestEncountersRouteListsWhatTheStoreDoes is the handler-level
// counterpart to TestEncountersAreListedWithTheirSlugs: the /rankings
// picker reads this route, not the store directly, so the route itself
// needs a passing test.
func TestEncountersRouteListsWhatTheStoreDoes(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedReport("report-one")
	h.seedFight("report-one", 1, engine.FixtureBase, nil)

	var out struct {
		Rows []Encounter `json:"rows"`
	}
	h.data(h.get("/v1/encounters"), &out)
	if len(out.Rows) != 1 || out.Rows[0].Name != "Warden Kelthas" || out.Rows[0].Slug != "warden-kelthas" {
		t.Fatalf("rows = %+v", out.Rows)
	}
}

// TestEncountersRouteAnswersAnEmptyListNotNull is what the /rankings
// picker's "no encounter data yet" copy depends on: a fresh deployment
// with no ranked fight must answer `"rows": []`, not `"rows": null` -
// json.Unmarshal leaves a slice nil for the latter, so the picker's own
// `rows.length === 0` check would need a null guard the contract never
// promised it.
func TestEncountersRouteAnswersAnEmptyListNotNull(t *testing.T) {
	h := newHarness(t)
	serve(t, h)

	var out struct {
		Rows []Encounter `json:"rows"`
	}
	h.data(h.get("/v1/encounters"), &out)
	if out.Rows == nil {
		t.Fatal("rows should be an empty array, not null")
	}
	if len(out.Rows) != 0 {
		t.Fatalf("rows = %+v, want none seeded", out.Rows)
	}
}
