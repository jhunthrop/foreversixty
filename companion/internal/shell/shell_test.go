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
	if got.Title != "Forever Sixty" || got.Width != 800 || got.Height != 620 {
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

func TestTheTooltipLeadsWithTheProductAndTheMenuLineDoesNot(t *testing.T) {
	if got := tooltip("Logging"); got != "Forever Sixty · Logging" {
		t.Fatalf("tooltip = %q", got)
	}
	long := tooltip(strings.Repeat("x", 400))
	if n := len([]rune(long)); n != TooltipMax {
		t.Fatalf("a long tooltip is %d characters, Windows shows %d", n, TooltipMax)
	}
}

func TestTheTrayIsRewrittenOnlyForWhatChanged(t *testing.T) {
	var calls []string
	a := trayApply{
		Tooltip: func(s string) { calls = append(calls, "tip:"+s) },
		Title:   func(s string) { calls = append(calls, "title:"+s) },
		Icon:    func(s string) { calls = append(calls, "icon:"+s) },
	}
	step := func(prev *trayView, now trayView) string {
		calls = nil
		applyTray(prev, now, a)
		return strings.Join(calls, ",")
	}
	logging := trayView{Line: "Logging", Tone: "on"}
	if got := step(nil, logging); got != "tip:Forever Sixty · Logging,title:Logging,icon:on" {
		t.Errorf("first call = %s", got)
	}
	if got := step(&logging, logging); got != "" {
		t.Errorf("nothing changed, but: %s", got)
	}
	if got := step(&logging, trayView{Line: "Logging", Tone: "warn"}); got != "icon:warn" {
		t.Errorf("tone only = %s", got)
	}
	if got := step(&logging, trayView{Line: "Ready", Tone: "on"}); got != "tip:Forever Sixty · Ready,title:Ready" {
		t.Errorf("words only = %s", got)
	}
}

func TestTheTrayToneDefaultsToTheSealAlone(t *testing.T) {
	if got := (Options{}).withDefaults().Tone(); got != "off" {
		t.Fatalf("tone = %q", got)
	}
}

func TestTheWindowOpensAtTheSpecSizeAndNeverShrinksBelowTheMinimum(t *testing.T) {
	if MinWidth != 560 || MinHeight != 480 {
		t.Fatalf("minimum = %dx%d", MinWidth, MinHeight)
	}
	got := Options{}.withDefaults()
	if got.Width < MinWidth || got.Height < MinHeight {
		t.Fatalf("default %dx%d is under the minimum", got.Width, got.Height)
	}
}
