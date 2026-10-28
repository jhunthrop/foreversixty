package coach

import (
	"context"
	"net/http"
	"testing"
)

// History and the recent list read only evidence reports (presence.EvidenceReport): a
// flagged or private report's reviews are listed nowhere, as the fight route answers 404.
func TestHistoryAndRecent_LeaveOutFlaggedAndPrivateReports(t *testing.T) {
	f := newT31(t)
	f.owners[f.premium] = map[string]bool{f.furyKey: true}
	f.linkCharacter(t, f.premium, true)
	f.seedReviews(t, 3)
	listed := func() (history, recent int) {
		t.Helper()
		code, _, body := f.call(t, user(f.premium), f.historyPath(""))
		if code != http.StatusOK {
			t.Fatalf("history %d %v", code, body)
		}
		return len(asMap(t, body)["reviews"].([]any)), len(f.recent(t, f.premium)["reviews"].([]any))
	}
	if h, r := listed(); h != 3 || r != 3 {
		t.Fatalf("a public report's reviews are listed: history %d, recent %d", h, r)
	}
	for name, update := range map[string]string{
		"flagged": `update reports set flagged = 'spam', visibility = 'public' where id = $1`,
		"private": `update reports set flagged = null, visibility = 'private' where id = $1`,
	} {
		if _, err := f.pool.Exec(context.Background(), update, t30Report); err != nil {
			t.Fatal(err)
		}
		if h, r := listed(); h != 0 || r != 0 {
			t.Fatalf("%s: history %d, recent %d; want none", name, h, r)
		}
	}
}
