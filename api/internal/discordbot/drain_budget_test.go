package discordbot

import (
	"context"
	"net/http"
	"testing"
	"time"
)

// A store write after Discord answered runs on its own writeBudget context, so a row whose
// budget is spent (or a drain told to stop) still records what happened.
func TestAStoreWriteGetsItsOwnBudget(t *testing.T) {
	parent, cancel := context.WithCancel(context.Background())
	cancel()
	var deadline time.Time
	var hasDeadline bool
	var err error
	start := time.Now()
	(&Drainer{}).write(parent, "post_fail", 1, func(ctx context.Context) error {
		deadline, hasDeadline = ctx.Deadline()
		err = ctx.Err()
		return nil
	})
	if err != nil || !hasDeadline {
		t.Fatalf("write ctx: err = %v, deadline = %v", err, hasDeadline)
	}
	const slack = 200 * time.Millisecond
	if left := deadline.Sub(start); left < writeBudget || left > writeBudget+slack {
		t.Fatalf("write budget = %v, want %v", left, writeBudget)
	}
}

// claimOne leases the world's one due row.
func (w *postWorld) claimOne(t *testing.T) Post {
	t.Helper()
	posts, err := w.st.ClaimDue(context.Background(), w.clock.now(), 0, defaultDrainLease, drainBatch)
	if err != nil || len(posts) != 1 {
		t.Fatalf("claim = %+v, %v", posts, err)
	}
	return posts[0]
}

// A channel refusal reads the server only with probeNeeds of the row budget left; with less
// the row is retried and the server is not read.
func TestAChannelRefusalProbesOnlyWithBudgetLeft(t *testing.T) {
	refusal := &APIError{Status: http.StatusForbidden, Code: codeMissingPerms}
	for _, c := range []struct {
		name   string
		left   time.Duration
		probes bool
	}{
		{"under the probe budget", probeNeeds - time.Second, false},
		{"over the probe budget", probeNeeds + 2*time.Second, true},
	} {
		t.Run(c.name, func(t *testing.T) {
			w := newPostWorld(t)
			w.enqueue(t, w.raidID)
			p := w.claimOne(t)
			ctx, cancel := context.WithTimeout(context.Background(), c.left)
			defer cancel()
			w.drainer().channelRefused(ctx, p, w.clock.now(), refusal)
			if probed := len(w.callsOf("ServerLayout")) == 1; probed != c.probes {
				t.Fatalf("server read = %v, want %v", probed, c.probes)
			}
			if r, _ := w.row(t, w.raidID); !c.probes && r.attempts != 1 {
				t.Fatalf("row = %+v, want a retry", r)
			}
		})
	}
}
