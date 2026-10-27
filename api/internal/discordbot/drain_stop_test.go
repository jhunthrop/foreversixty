package discordbot

import (
	"context"
	"testing"
	"time"
)

// A stop lets the drain's last work land, but only for StopWithin: a store write that would
// get its whole writeBudget is cut at what is left of the stop.
func TestAStopCapsTheDetachedWrites(t *testing.T) {
	parent, stop := context.WithCancel(context.Background())
	const within = 100 * time.Millisecond
	land, end := landing(parent, within)
	defer end()
	row, cancelRow := context.WithTimeout(land, rowBudget)
	defer cancelRow()

	stop()
	if row.Err() != nil {
		t.Fatal("the row in flight ended with the stop signal")
	}
	start := time.Now()
	var ended time.Duration
	(&Drainer{}).write(row, "post_finish", 1, func(ctx context.Context) error {
		<-ctx.Done()
		ended = time.Since(start)
		return nil
	})
	if ended < within/2 || ended >= writeBudget {
		t.Fatalf("write ran %v after the stop, want about %v", ended, within)
	}
	if row.Err() == nil {
		t.Fatal("the row outlived the stop's budget")
	}
}

// Without a stop the landing context never ends on its own, and end releases it.
func TestLandingEndsOnlyWithTheStopOrEnd(t *testing.T) {
	land, end := landing(context.Background(), time.Millisecond)
	time.Sleep(10 * time.Millisecond)
	if land.Err() != nil {
		t.Fatal("landing ended with no stop")
	}
	end()
	if land.Err() == nil {
		t.Fatal("end did not release the landing context")
	}
}
