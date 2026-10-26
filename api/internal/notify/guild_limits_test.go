package notify

import (
	"context"
	"net/http"
	"testing"
)

// cutPoster ends the request's budget mid-post, as the shared raid budget does.
type cutPoster struct {
	*FakePoster
	cut context.CancelFunc
}

func (p cutPoster) Post(ctx context.Context, url string, embed Embed) error {
	_ = p.FakePoster.Post(ctx, url, embed)
	p.cut()
	return context.DeadlineExceeded
}

// A post cut by the request's shared budget says nothing about the channel, so it is not
// one of the five failures.
func TestGuildDelivererBudgetCutOffIsNotCounted(t *testing.T) {
	r := newGuildRig(t)
	ctx, cut := context.WithCancel(context.Background())
	defer cut()
	r.d.Poster = cutPoster{FakePoster: r.poster, cut: cut}
	if err := r.d.Deliver(ctx, r.gid, eventRaidPosted, Embed{}); err != nil {
		t.Fatalf("a budget cut-off returned %v", err)
	}
	if w := r.hook(t); w.Failures != 0 || w.LastError != nil {
		t.Fatalf("a budget cut-off was recorded: %+v", w)
	}
}

// The post's own timeout, with the request still alive, is the channel's failure.
func TestGuildDelivererOwnTimeoutStillCounts(t *testing.T) {
	r := newGuildRig(t)
	r.poster.PostErr = context.DeadlineExceeded
	if err := r.d.Deliver(context.Background(), r.gid, eventRaidPosted, Embed{}); err != nil {
		t.Fatal(err)
	}
	if w := r.hook(t); w.Failures != 1 {
		t.Fatalf("failures = %d, want 1", w.Failures)
	}
}

func TestGuildConnectIsLimitedToTwentyPerMinute(t *testing.T) {
	g := newGuildAPI(t, officerWithPlan())
	for i := 1; i <= connectsPerMinute; i++ {
		if res, _ := g.get(g.route("/connect")); res.StatusCode != http.StatusFound {
			t.Fatalf("call %d status %d", i, res.StatusCode)
		}
	}
	if res, _ := g.get(g.route("/connect")); res.StatusCode != http.StatusTooManyRequests {
		t.Fatalf("21st call status %d, want 429", res.StatusCode)
	}
}
