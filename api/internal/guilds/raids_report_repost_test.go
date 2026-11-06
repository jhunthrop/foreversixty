package guilds

import (
	"net/http"
	"testing"
	"time"
)

// The post names the night's best parses and the loot winners, counts the deaths, and a
// posted report goes out again only when the officer asks.
func TestTheReportPostSummarisesAndRepostsOnlyWhenAsked(t *testing.T) {
	w, night, events := postWorld(t)
	w.as(w.officer)
	path := w.path("/raids/%d/report/post", night.ID)
	expectError(t, w.h.do(http.MethodPost, path, `{"repost":true}`), http.StatusConflict, "not_posted")
	if res := w.h.do(http.MethodPost, path, ""); res.StatusCode != http.StatusOK {
		t.Fatalf("first post = %d: %s", res.StatusCode, readBody(res))
	}
	e := events.reports[0]
	if len(e.Loot) != 2 || e.Loot[0].Winner == "" || e.Loot[0].Item == "" {
		t.Fatalf("loot = %+v", e.Loot)
	}
	if len(e.TopParses) == 0 || len(e.TopParses) > 3 || e.TopParses[0].Metric != "dps" || e.TopParses[0].Boss == "" {
		t.Fatalf("top parses = %+v", e.TopParses)
	}
	if e.Time <= 0 {
		t.Fatalf("time = %v", e.Time)
	}
	first := w.postedAt(t, night.ID)
	expectError(t, w.h.do(http.MethodPost, path, `{"repost":false}`), http.StatusConflict, "already_posted")
	expectError(t, w.h.do(http.MethodPost, path, `junk`), http.StatusBadRequest, "invalid")

	events.notPosted = true
	expectError(t, w.h.do(http.MethodPost, path, `{"repost":true}`), http.StatusBadGateway, "post_failed")
	if at := w.postedAt(t, night.ID); at == nil || !at.Equal(*first) {
		t.Fatalf("a failed repost left posted_at %v, want the first %v", at, first)
	}
	events.notPosted = false
	time.Sleep(2 * time.Millisecond)
	if res := w.h.do(http.MethodPost, path, `{"repost":true}`); res.StatusCode != http.StatusOK {
		t.Fatalf("repost = %d: %s", res.StatusCode, readBody(res))
	}
	if n := len(events.reports); n != 3 {
		t.Fatalf("posts = %d, want the first, the failed repost and the repost", n)
	}
	if at := w.postedAt(t, night.ID); at == nil || !at.After(*first) {
		t.Fatalf("repost stamp = %v, want after %v", at, first)
	}
}
