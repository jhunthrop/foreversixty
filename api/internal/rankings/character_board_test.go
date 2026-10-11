package rankings

import (
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/engine"
	"github.com/jhunthrop/foreversixty/api/internal/metrics"
)

const boardPath = "/v1/rankings?encounter=9001"

// seedTwoKills writes the same three characters twice: the second report
// raises the first character's damage and lowers the second's, so each
// character has a best kill that is not simply the latest.
func seedTwoKills(h *harness) {
	h.seedReport("report-one")
	first := h.seedFight("report-one", 1, engine.FixtureBase, nil)
	h.seedReport("report-two")
	h.seedFight("report-two", 1, engine.FixtureBase.Add(2*time.Hour), func(rows []metrics.Row) {
		rows[0].MetricDPS = first[0].MetricDPS + 500
		rows[1].MetricDPS = first[1].MetricDPS / 2
	})
}

func TestRankingsDefaultStillRankEveryKill(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	seedTwoKills(h)

	var page Page
	h.data(h.get(boardPath), &page)
	if page.Deduped || page.Total != 6 || len(page.Rows) != 6 || page.CharacterRank != nil {
		t.Fatalf("page = deduped %v total %d rows %d", page.Deduped, page.Total, len(page.Rows))
	}
}

func TestRankingsDistinctCharacterKeepsEachCharactersBestKill(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	seedTwoKills(h)

	var page Page
	h.data(h.get(boardPath+"&distinct=character"), &page)
	if !page.Deduped || page.Total != 3 || len(page.Rows) != 3 {
		t.Fatalf("page = deduped %v total %d rows %d", page.Deduped, page.Total, len(page.Rows))
	}
	seen := map[string]bool{}
	for i, row := range page.Rows {
		if seen[row.Player.Key] {
			t.Fatalf("%s is on the board twice", row.Player.Key)
		}
		seen[row.Player.Key] = true
		if row.Rank != i+1 {
			t.Fatalf("row %d has rank %d", i, row.Rank)
		}
		if i > 0 && page.Rows[i-1].Value < row.Value {
			t.Fatalf("rows are not ordered by value: %+v", page.Rows)
		}
	}
	if page.Rows[0].ReportID != "report-two" {
		t.Fatalf("the top character's best kill is in %s, want report-two", page.Rows[0].ReportID)
	}
}

func TestRankingsPlaceACharacterOnTheDedupedBoard(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	seedTwoKills(h)

	var board Page
	h.data(h.get(boardPath+"&distinct=character"), &board)
	want := board.Rows[len(board.Rows)-1]

	var page Page
	h.data(h.get(boardPath+"&distinct=character&character="+want.Player.Key), &page)
	if page.CharacterRank == nil {
		t.Fatal("no character_rank for a ranked character")
	}
	if page.CharacterRank.Rank != want.Rank || page.CharacterRank.Row.Player.Key != want.Player.Key ||
		page.CharacterRank.Row.Value != want.Value {
		t.Fatalf("character_rank = %+v, want the board's own row %+v", page.CharacterRank, want)
	}

	var none Page
	h.data(h.get(boardPath+"&distinct=character&character=us/hardcore/nobody"), &none)
	if none.CharacterRank != nil {
		t.Fatalf("an unranked character got %+v", none.CharacterRank)
	}
	h.data(h.get(boardPath+"&distinct=character&character="+want.Player.Key+"&class=no-such-class"), &none)
	if none.CharacterRank != nil {
		t.Fatalf("a character outside the filtered cohort got %+v", none.CharacterRank)
	}
}

func TestRankingsRejectABadCharacterBoard(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	seedTwoKills(h)

	for name, query := range map[string]string{
		"unknown distinct":            "&distinct=guild",
		"character without distinct":  "&character=us/hardcore/baelgrim",
		"character that is not a key": "&distinct=character&character=baelgrim",
	} {
		res := h.get(boardPath + query)
		res.Body.Close()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("%s = %d, want 400", name, res.StatusCode)
		}
	}
}
