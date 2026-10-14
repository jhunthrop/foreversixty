package reports

import (
	"testing"
	"time"
)

// Task L6.12: a flagged report (tampered or unverified) is evidence of nothing, so the public
// recent-reports feed does not list it; a clean report beside it is listed.
func TestStoreRecentLeavesOutAFlaggedReport(t *testing.T) {
	h := newHarness(t)
	base := time.Now().UTC().Truncate(time.Second)
	h.seedReport("recentclean1", "Clean night", "Molten Core", Public, StatusComplete, base)
	h.seedReport("recentflag01", "Forged night", "Molten Core", Public, StatusComplete, base.Add(time.Minute))
	if _, err := h.store.Pool.Exec(h.t.Context(), `update reports set flagged = 'tampered' where id = 'recentflag01'`); err != nil {
		t.Fatal(err)
	}

	rows, err := h.store.Recent(h.t.Context(), nil, RecentPerPage)
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 || rows[0].ID != "recentclean1" {
		t.Fatalf("rows = %+v, want only recentclean1 (the flagged report is not listed)", rows)
	}
}
