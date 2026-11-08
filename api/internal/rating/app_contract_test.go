package rating

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/contract"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// The character rating keeps its app fixture: GET /v1/characters/{region}/{ruleset}/{name}/rating.
func TestAppContractCharacterRating(t *testing.T) {
	svc, rs := newTestService(t)
	mustCreateReport(t, rs, "contract-char-1", reports.Public)
	f := fightFixture("contract-char-1", true,
		summary.RosterRow{GUID: "g1", Name: "Trendy", Class: "Shaman", Spec: "Elemental", Role: "dps",
			ExecutionScore: execScorePtr(0.9)},
	)
	if err := svc.Store.RateFight(context.Background(), f); err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	Mount(mux, svc)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/v1/characters/us/normal/Trendy/rating", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d: %s", rec.Code, rec.Body)
	}
	contract.AssertFixture(t, "character.rating", rec.Body.Bytes(), contract.Options{
		Optional: []string{"data.next_cursor", "data.best_component", "data.worst_component",
			"data.latest.insufficient_reason", "data.latest.components[].percentile", "data.latest.components[].bracket_n",
			"data.latest.components[].reason", "data.latest.components[].moments[].spell_id",
			"data.latest.components[].moments[].spell_name"},
	})
}
