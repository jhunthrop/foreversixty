package guilds

import (
	"net/http"
	"testing"
)

// Posting the report of a raid whose matched log went private or flagged answers 409 conflict
// before claiming: Discord is never tried and the claim stays free.
func TestPostingARaidWhoseLogIsNotPublicIsAConflict(t *testing.T) {
	w, raid, events := postWorld(t)
	w.exec(t, `update reports set flagged = 'tampered' where id = 'night'`)
	w.as(w.officer)
	expectError(t, w.h.do(http.MethodPost, w.path("/raids/%d/report/post", raid.ID), ""), http.StatusConflict, "conflict")
	if at := w.postedAt(t, raid.ID); at != nil {
		t.Fatalf("report_posted_at = %v, want no claim", at)
	}
	if n := len(events.reports); n != 0 {
		t.Fatalf("%d posts, want none", n)
	}
}
