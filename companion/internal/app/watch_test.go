package app_test

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/app"
	"github.com/jhunthrop/foreversixty/companion/internal/watch"
)

// played marks a game folder as played at a moment: the game
// rewrites a log under Logs at launch.
func played(t *testing.T, game string, at time.Time) {
	t.Helper()
	p := filepath.Join(game, "Logs", "General.log")
	if err := os.WriteFile(p, []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.Chtimes(p, at, at); err != nil {
		t.Fatal(err)
	}
}

func watchedPath(t *testing.T, h *harness) string {
	t.Helper()
	for _, in := range h.app.Snapshot().Installs {
		if in.Watching {
			return in.Path
		}
	}
	return ""
}

// twoGames saves both folders, the harness's first.
func twoGames(t *testing.T, h *harness) string {
	t.Helper()
	second := newGame(t)
	if code, body := h.call(http.MethodPost, h.url("/api/settings"), app.Settings{
		Visibility: "public", WoWPaths: []string{h.game, second},
	}); code != http.StatusOK {
		t.Fatalf("settings = %d %v", code, body)
	}
	return second
}

func TestTheWatchedInstallIsTheOnePlayedMostRecently(t *testing.T) {
	h := newHarness(t)
	second := twoGames(t, h)
	played(t, h.game, t0.Add(-48*time.Hour))
	played(t, second, t0.Add(-time.Hour))
	h.app.Step(t.Context(), t0)
	if got := watchedPath(t, h); got != second {
		t.Fatalf("watching %q, not the install played last %q", got, second)
	}
	for _, in := range h.app.Snapshot().Installs {
		if in.LastPlayed == nil {
			t.Errorf("%s has no last_played", in.Path)
		}
	}
}

func TestNoActivityAnywhereWatchesTheFirstInstall(t *testing.T) {
	h := newHarness(t)
	twoGames(t, h)
	h.app.Step(t.Context(), t0)
	if got := watchedPath(t, h); got != h.game {
		t.Fatalf("watching %q, not the first install %q", got, h.game)
	}
	for _, in := range h.app.Snapshot().Installs {
		if in.LastPlayed != nil {
			t.Errorf("%s claims last_played %v", in.Path, in.LastPlayed)
		}
	}
}

func TestTheWatchedInstallFollowsTheNewestActivityOnLaterTicks(t *testing.T) {
	h := newHarness(t)
	second := twoGames(t, h)
	played(t, h.game, t0.Add(-time.Hour))
	h.app.Step(t.Context(), t0)
	if got := watchedPath(t, h); got != h.game {
		t.Fatalf("watching %q, want %q", got, h.game)
	}
	played(t, second, t0.Add(time.Minute))
	h.app.Step(t.Context(), t0.Add(2*time.Second))
	if got := watchedPath(t, h); got != second {
		t.Fatalf("watching %q after the other install was played, want %q", got, second)
	}
}

func TestTheSwapWaitsForTheOpenReport(t *testing.T) {
	h := newHarness(t)
	second := twoGames(t, h)
	played(t, h.game, t0.Add(-time.Hour))
	h.app.Step(t.Context(), t0)
	if err := os.WriteFile(h.log, []byte(aNight()), 0o600); err != nil {
		t.Fatal(err)
	}
	h.app.Step(t.Context(), t0.Add(time.Second))
	if p := h.app.Snapshot().Pipeline; !p.Logging {
		t.Fatalf("no report is open: %+v", p)
	}
	played(t, second, t0.Add(time.Hour))
	h.app.Step(t.Context(), t0.Add(2*time.Second))
	if got := watchedPath(t, h); got != h.game {
		t.Fatalf("the swap cut a night in progress: watching %q", got)
	}
	quiet := t0.Add(2*time.Second + watch.CompleteIdle + time.Minute)
	h.app.Step(t.Context(), quiet)
	h.app.Step(t.Context(), quiet.Add(time.Second))
	if got := watchedPath(t, h); got != second {
		t.Fatalf("after the report closed, watching %q, want %q", got, second)
	}
}

func TestANearTieDoesNotFlapTheWatchedInstall(t *testing.T) {
	h := newHarness(t)
	second := twoGames(t, h)
	base := t0.Add(-time.Hour)
	played(t, h.game, base)
	played(t, second, base.Add(time.Minute))
	h.app.Step(t.Context(), t0)
	if got := watchedPath(t, h); got != second {
		t.Fatalf("watching %q, want %q", got, second)
	}
	// The other install is touched half a second later than the watched
	// one: not enough to swap.
	played(t, h.game, base.Add(time.Minute+500*time.Millisecond))
	h.app.Step(t.Context(), t0.Add(time.Second))
	if got := watchedPath(t, h); got != second {
		t.Fatalf("a sub-second lead swapped the watch to %q", got)
	}
	// Two seconds later it does.
	played(t, h.game, base.Add(time.Minute+2*time.Second))
	h.app.Step(t.Context(), t0.Add(2*time.Second))
	if got := watchedPath(t, h); got != h.game {
		t.Fatalf("a clear lead did not swap the watch: %q", got)
	}
}

func TestARemovedWatchedFolderRetargetsOnTheNextTick(t *testing.T) {
	h := newHarness(t)
	second := twoGames(t, h)
	played(t, h.game, t0.Add(-time.Hour))
	played(t, second, t0)
	h.app.Step(t.Context(), t0)
	if got := watchedPath(t, h); got != second {
		t.Fatalf("watching %q, want %q", got, second)
	}
	if err := h.app.RemoveFolder(second); err != nil {
		t.Fatal(err)
	}
	h.app.Step(t.Context(), t0.Add(time.Second))
	if got := watchedPath(t, h); got != h.game {
		t.Fatalf("after removing the watched folder, watching %q, want %q", got, h.game)
	}
}
