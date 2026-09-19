package rankings

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/engine"
)

func TestExecutionIsASortableMetricButNotADigestedOne(t *testing.T) {
	if !ValidMetric(MetricExecution) {
		t.Fatal("execution is not a valid metric")
	}
	if ValidDigestMetric(MetricExecution) {
		t.Fatal("execution has no percentile digest; the percentile route must refuse it")
	}
	for _, m := range DigestedMetrics {
		if !ValidMetric(m) {
			t.Errorf("%q is digested but not rankable", m)
		}
	}
	var found bool
	for _, m := range Metrics {
		if m == MetricExecution {
			found = true
		}
	}
	if !found {
		t.Fatal("execution is missing from Metrics")
	}
	if got := metricColumn(MetricExecution); got != "m.execution_score" {
		t.Fatalf("metricColumn(execution) = %q", got)
	}
}

func TestThePercentileRouteRefusesExecution(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedRankedFight(t)
	res := h.get("/v1/rankings/percentile?encounter=9001&difficulty=8&value=1000&metric=execution")
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status %d, want 400: there is no execution digest to place a value on", res.StatusCode)
	}
	// And the metric that does have a digest still works.
	ok := h.get("/v1/rankings/percentile?encounter=9001&difficulty=8&value=1000&metric=dps")
	defer ok.Body.Close()
	if ok.StatusCode != http.StatusOK {
		t.Fatalf("status %d for dps, want 200", ok.StatusCode)
	}
}

func TestAnExecutionScoreIsClampedAndReadBack(t *testing.T) {
	h := newHarness(t)
	rep, index, key := h.seedRankedFight(t)

	for _, c := range []struct {
		wrote float64
		want  float64
	}{
		{0.92, 0.92},
		{-3, ExecutionMin}, // a sim that went wrong, not a player who did
		{17, ExecutionMax},
	} {
		if err := h.store.SetExecutionScore(t.Context(), rep, index, key, c.wrote); err != nil {
			t.Fatal(err)
		}
		var got float64
		if err := h.pool.QueryRow(t.Context(),
			`select execution_score from fight_metrics
			 where report_id = $1 and fight_index = $2 and player_key = $3`,
			rep, index, key).Scan(&got); err != nil {
			t.Fatal(err)
		}
		if got != c.want {
			t.Errorf("wrote %v, read %v, want %v", c.wrote, got, c.want)
		}
	}
}

func TestAnUnscoredFightIsNullEverywhereItAppears(t *testing.T) {
	h := newHarness(t)
	_, _, key := h.seedRankedFight(t)

	page, err := h.store.Rankings(t.Context(),
		Query{EncounterID: executionEncounterID, Metric: MetricDPS}, engine.FixtureBase)
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Rows) == 0 {
		t.Fatal("no rows to check")
	}
	for _, r := range page.Rows {
		if r.ExecutionScore != nil {
			t.Errorf("%s has a score before anything scored it", r.Player.Key)
		}
	}
	// And it is present in the JSON as null rather than absent: the
	// web reads the field's presence to know the column exists.
	b, err := json.Marshal(page.Rows[0])
	if err != nil {
		t.Fatal(err)
	}
	var raw map[string]any
	if err := json.Unmarshal(b, &raw); err != nil {
		t.Fatal(err)
	}
	if _, ok := raw["execution_score"]; !ok {
		t.Errorf("a ranking row has no execution_score field: %s", b)
	}

	name := key[strings.LastIndex(key, "/")+1:]
	c, ok, err := h.store.Character(t.Context(), "us", "hardcore", name)
	if err != nil || !ok {
		t.Fatalf("character: ok=%v err=%v", ok, err)
	}
	if len(c.History) == 0 {
		t.Fatal("no history to check")
	}
	if c.History[0].ExecutionScore != nil {
		t.Error("an unscored fight has a score on the character page")
	}
}

func TestTheExecutionLeaderboardOnlyListsScoredFights(t *testing.T) {
	h := newHarness(t)
	rep, index, key := h.seedRankedFight(t)

	// Nothing is scored yet, so the execution board is empty even
	// though the DPS board is not.
	board, err := h.store.Rankings(t.Context(),
		Query{EncounterID: executionEncounterID, Metric: MetricExecution}, engine.FixtureBase)
	if err != nil {
		t.Fatal(err)
	}
	if board.Total != 0 || len(board.Rows) != 0 {
		t.Fatalf("unscored fights appeared on the execution board: %+v", board)
	}

	if err := h.store.SetExecutionScore(t.Context(), rep, index, key, 0.92); err != nil {
		t.Fatal(err)
	}
	board, err = h.store.Rankings(t.Context(),
		Query{EncounterID: executionEncounterID, Metric: MetricExecution}, engine.FixtureBase)
	if err != nil {
		t.Fatal(err)
	}
	if board.Total != 1 || len(board.Rows) != 1 {
		t.Fatalf("a scored fight is not on the execution board: %+v", board)
	}
	row := board.Rows[0]
	if row.Value != 0.92 {
		t.Errorf("value %v, want the execution score", row.Value)
	}
	if row.ExecutionScore == nil || *row.ExecutionScore != 0.92 {
		t.Errorf("execution_score %v", row.ExecutionScore)
	}
}

func TestSetExecutionScoreOnAMissingFightReturnsErrNoFight(t *testing.T) {
	h := newHarness(t)
	rep, index, key := h.seedRankedFight(t)

	// A coordinate that names no row - a typo'd player key here - is
	// silently discarded by a bare UPDATE with no rows matched, which
	// the caller can't tell apart from a real write. It must surface
	// as the sentinel instead.
	err := h.store.SetExecutionScore(t.Context(), rep, index, key+"-typo", 0.5)
	if !errors.Is(err, ErrNoFight) {
		t.Fatalf("err = %v, want ErrNoFight", err)
	}

	// And a real coordinate is unaffected: the miss above wrote
	// nothing, and a hit still updates the row.
	if err := h.store.SetExecutionScore(t.Context(), rep, index, key, 0.5); err != nil {
		t.Fatal(err)
	}
	var got float64
	if err := h.pool.QueryRow(t.Context(),
		`select execution_score from fight_metrics
		 where report_id = $1 and fight_index = $2 and player_key = $3`,
		rep, index, key).Scan(&got); err != nil {
		t.Fatal(err)
	}
	if got != 0.5 {
		t.Fatalf("execution_score = %v, want 0.5", got)
	}
}
