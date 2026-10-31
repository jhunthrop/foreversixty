package sims

import (
	"testing"

	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

func TestLiveReadsCountRunningAndOwnRuns(t *testing.T) {
	h := newHarness(t)
	ctx := t.Context()
	req := browserResult("warrior-fury", 0).Request
	for _, id := range []string{"aaaaaaaaaaaa", "bbbbbbbbbbbb"} {
		if err := h.store.Queue(ctx, id, h.owner, req); err != nil {
			t.Fatal(err)
		}
	}
	if n, err := h.store.RunningServerRuns(ctx, h.owner); err != nil || n != 2 {
		t.Fatalf("running = %d, %v; want 2", n, err)
	}
	done := browserResult("warrior-fury", 1000)
	done.Lane = simapi.LaneServer
	if err := h.store.Finish(ctx, "aaaaaaaaaaaa", done); err != nil {
		t.Fatal(err)
	}
	if n, err := h.store.RunningServerRuns(ctx, h.owner); err != nil || n != 1 {
		t.Fatalf("running after a finish = %d, %v; want 1", n, err)
	}
	if owns, err := h.store.OwnsServerRun(ctx, h.owner, "aaaaaaaaaaaa"); err != nil || !owns {
		t.Fatalf("owns = %v, %v", owns, err)
	}
	if owns, err := h.store.OwnsServerRun(ctx, h.owner+1, "aaaaaaaaaaaa"); err != nil || owns {
		t.Fatalf("someone else owns it: %v, %v", owns, err)
	}
}
