package rankings

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

// Task L6.10: a report flagged tampered or unverified is evidence of nothing, so the score
// history never shows a snapshot rolled up onto one of its fights.
func TestCharacterScoreHistorySkipsAFlaggedReport(t *testing.T) {
	x := newExtras(t)
	key := x.char("cx66-l610-history", "Cx Flagged", nil)
	x.report("cx66-l610-clean", "public")
	x.report("cx66-l610-flagged", "public")
	x.run(`update reports set flagged = 'tampered' where id = 'cx66-l610-flagged'`)
	score := func(n int) *int { return &n }
	tier := func(s string) *string { return &s }
	x.rated("cx66-l610-clean", 1, key, "Cx Flagged", time.Date(2026, 3, 2, 20, 0, 0, 0, time.UTC), score(400), tier("common"))
	x.rated("cx66-l610-flagged", 1, key, "Cx Flagged", time.Date(2026, 3, 5, 20, 0, 0, 0, time.UTC), score(1900), tier("epic"))

	status, c, raw, _, body := x.page("cx66-l610-history")
	if status != http.StatusOK {
		t.Fatalf("status %d: %s", status, body)
	}
	want := []ScorePoint{{Date: "2026-03-02", Score: 400, Tier: "common"}}
	if fmt.Sprint(c.ScoreHistory) != fmt.Sprint(want) {
		t.Fatalf("score_history = %+v (raw %s), want only the clean report's %+v", c.ScoreHistory, raw["score_history"], want)
	}
}
