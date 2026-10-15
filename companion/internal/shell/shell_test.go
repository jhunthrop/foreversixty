package shell

import (
	"context"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestDefaultsFillInTheWindow(t *testing.T) {
	got := Options{}.withDefaults()
	if got.Title != "Forever Sixty" || got.Width != 960 || got.Height != 680 {
		t.Fatalf("defaults = %+v", got)
	}
	kept := Options{Title: "x", Width: 1, Height: 2}.withDefaults()
	if kept.Title != "x" || kept.Width != 1 || kept.Height != 2 {
		t.Fatalf("defaults overrode the caller: %+v", kept)
	}
}

func TestWaitReturnsWhenTheContextEnds(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Millisecond)
	defer cancel()
	if err := wait(ctx); err == nil {
		t.Fatal("wait returned no error on a cancelled context")
	}
}

// TestHeadlessRunReturns covers the path CI takes. In the real build
// Run is headless here too, so no window is opened by the test suite.
func TestHeadlessRunReturns(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	if err := Run(ctx, Options{Headless: true, URL: "http://127.0.0.1:1/"}); err == nil {
		t.Fatal("Run returned no error on a cancelled context")
	}
}

func TestTheTrayLineIsAppliedOnceAndThenOnlyWhenItChanges(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()
	var mu sync.Mutex
	lines := []string{"Not connected"}
	var applied []string
	read := func() string { mu.Lock(); defer mu.Unlock(); return lines[len(lines)-1] }
	set := func(s string) { mu.Lock(); lines = append(lines, s); mu.Unlock() }
	done := make(chan struct{})
	go func() {
		defer close(done)
		watchLine(ctx, time.Millisecond, read, func(s string) { mu.Lock(); applied = append(applied, s); mu.Unlock() })
	}()
	time.Sleep(15 * time.Millisecond) // many ticks, one line
	set("Ready")
	time.Sleep(15 * time.Millisecond)
	set("Ready")
	set("Logging")
	time.Sleep(15 * time.Millisecond)
	cancel()
	<-done
	mu.Lock()
	defer mu.Unlock()
	if got := strings.Join(applied, ","); got != "Not connected,Ready,Logging" {
		t.Fatalf("applied = %s", got)
	}
}

func TestTheTrayStateDefaultsToTheTitle(t *testing.T) {
	if got := (Options{}).withDefaults().State(); got != "Forever Sixty" {
		t.Fatalf("state = %q", got)
	}
}
