// companion/integration/integration_test.go
// The whole companion against a fake game and a fake ingest: a raid
// night written in bursts, a network that drops in the middle of it,
// and a companion that is stopped and started again. The one thing
// this asserts is the one thing a raid leader notices: every fight
// arrives, once, in the order it was fought.
package integration_test

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/app"
	"github.com/jhunthrop/foreversixty/companion/internal/client"
	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/jhunthrop/foreversixty/companion/internal/fakeapi"
	"github.com/jhunthrop/foreversixty/companion/internal/fixture"
	"github.com/jhunthrop/foreversixty/companion/internal/paths"
	"github.com/jhunthrop/foreversixty/companion/internal/secret"
	"github.com/jhunthrop/foreversixty/companion/internal/watch"
)

var t0 = time.Date(2026, 12, 9, 20, 0, 0, 0, time.UTC)

// game is the fake game writer: it appends to a combat log the way
// the client does, a burst at a time, and never rewrites what it has
// written.
type game struct {
	t    *testing.T
	path string
	now  time.Time
}

func (g *game) burst(text string) {
	g.t.Helper()
	f, err := os.OpenFile(g.path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		g.t.Fatal(err)
	}
	if _, err := f.WriteString(text); err != nil {
		g.t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		g.t.Fatal(err)
	}
	g.now = g.now.Add(time.Second)
	if err := os.Chtimes(g.path, g.now, g.now); err != nil {
		g.t.Fatal(err)
	}
}

// companion is the app over one set of directories, rebuildable so a
// restart can be simulated.
type companion struct {
	t    *testing.T
	dirs paths.Dirs
	cfg  config.Config
	app  *app.App
}

func (c *companion) start() {
	c.t.Helper()
	if c.app != nil {
		c.app.Close()
	}
	a, err := app.New(app.Options{
		Dirs: c.dirs, Config: c.cfg,
		Secret: secret.ConfigFile{Path: c.dirs.ConfigFile()},
		// One attempt with no wait: the outage in this test is
		// simulated, so there is nothing to be patient about.
		Retry: client.Retry{MaxAttempts: 1, Base: time.Millisecond, Max: time.Millisecond},
	})
	if err != nil {
		c.t.Fatal(err)
	}
	c.app = a
}

func TestARaidNightSurvivesADropAndARestart(t *testing.T) {
	home := t.TempDir()
	t.Setenv(paths.HomeEnv, home)
	dirs, err := paths.Resolve()
	if err != nil {
		t.Fatal(err)
	}
	install := filepath.Join(t.TempDir(), "_classic_era_")
	logs := filepath.Join(install, "Logs")
	for _, p := range []string{logs, filepath.Join(install, "WTF", "Account", "ACCOUNT#1", "SavedVariables")} {
		if err := os.MkdirAll(p, 0o755); err != nil {
			t.Fatal(err)
		}
	}
	srv := fakeapi.New()
	defer srv.Close()

	cfg := config.Default()
	cfg.APIBaseURL, cfg.SiteBaseURL = srv.URL, srv.URL
	cfg.WoWPaths = []string{install}
	cfg.DeviceToken = fakeapi.Token // already paired
	if err := config.Save(dirs.ConfigFile(), cfg); err != nil {
		t.Fatal(err)
	}

	g := &game{t: t, path: filepath.Join(logs, "WoWCombatLog.txt"), now: t0}
	// Last night is already in the file. The companion begins at the
	// end of what it did not watch, so tonight's report starts at a
	// non-zero file offset and every offset it sends has to be
	// counted from there rather than from the start of the file. It
	// is a different pull from tonight's, so the two nights do not
	// share bytes at the same offsets and an offset counted from the
	// wrong end addresses the wrong fight.
	prelude := fixture.Header + fixture.Zone + fixture.Encounter(9) + fixture.Heartbeat(9)
	g.burst(prelude)

	c := &companion{t: t, dirs: dirs, cfg: cfg}
	c.start()
	defer func() { c.app.Close() }()

	clock := t0
	step := func() {
		clock = clock.Add(time.Second)
		c.app.Step(context.Background(), clock)
	}

	step() // the companion is running before the player logs in
	g.burst(fixture.Header + fixture.Zone)
	step()

	const fights = 6
	for i := range fights {
		// Each encounter goes out in two bursts, so a fight can
		// straddle a poll the way a real one does.
		body := fixture.Encounter(i) + fixture.Heartbeat(i)
		half := strings.Index(body[len(body)/2:], "\n") + len(body)/2 + 1
		g.burst(body[:half])
		step()
		g.burst(body[half:])
		step()

		switch i {
		case 2:
			// The network drops for one fight's worth of bursts.
			srv.Offline(true)
		case 3:
			srv.Offline(false)
		case 4:
			// The player quits the companion and starts it again.
			c.start()
		}
	}

	// The raid ends: fifteen quiet minutes complete the report.
	clock = clock.Add(watch.CompleteIdle + time.Minute)
	c.app.Step(context.Background(), clock)
	for range 3 {
		step()
	}

	order := srv.Order()
	var arrived []int
	completed := 0
	for _, e := range order {
		switch {
		case strings.HasPrefix(e, "fight "):
			n, err := strconv.Atoi(strings.TrimPrefix(e, "fight "))
			if err != nil {
				t.Fatalf("unreadable event %q", e)
			}
			arrived = append(arrived, n)
		case strings.HasPrefix(e, "complete "):
			completed++
		}
	}
	if len(arrived) != fights {
		t.Fatalf("%d fights arrived, want %d: %v", len(arrived), fights, order)
	}
	for i := 1; i < len(arrived); i++ {
		if arrived[i-1] >= arrived[i] {
			t.Fatalf("fights arrived out of order: %v", arrived)
		}
	}
	// Ascending is not enough on its own: the indices are 1-based and
	// there are no gaps, so the night's six fights are exactly 1..6.
	for i, n := range arrived {
		if n != i+1 {
			t.Fatalf("fights arrived as %v, want 1..%d", arrived, fights)
		}
	}
	if completed != 1 {
		t.Fatalf("the report was completed %d times: %v", completed, order)
	}
	if len(srv.Reports()) != 1 {
		t.Fatalf("the night produced %d reports, want one", len(srv.Reports()))
	}

	written, err := os.ReadFile(g.path)
	if err != nil {
		t.Fatal(err)
	}
	tonight := written[len(prelude):]

	for id, rep := range srv.Reports() {
		if rep.Complete == nil {
			t.Fatalf("report %s was never completed", id)
		}
		if len(rep.Fights) != fights {
			t.Errorf("report %s holds %d fights", id, len(rep.Fights))
		}
		// The raw stream must be contiguous from zero to the final
		// offset the completion declared, because the server
		// reassembles it by offset.
		var offsets []int
		for o := range rep.Raw {
			offsets = append(offsets, int(o))
		}
		sort.Ints(offsets)
		if len(offsets) == 0 || offsets[0] != 0 {
			t.Fatalf("raw chunks start at %v", offsets)
		}
		var stream []byte
		for _, o := range offsets {
			if o != len(stream) {
				t.Fatalf("a raw chunk is missing before offset %d: %v", o, offsets)
			}
			decoded, err := decompress(rep.Raw[int64(o)])
			if err != nil {
				t.Fatal(err)
			}
			stream = append(stream, decoded...)
		}
		if int64(len(stream)) != rep.Complete.FinalOffset {
			t.Errorf("the raw stream is %d bytes, the completion declared %d",
				len(stream), rep.Complete.FinalOffset)
		}
		// Reassembled, the stream is tonight's log and nothing else:
		// last night's bytes are before the report's first byte and
		// must not be in it.
		if !bytes.Equal(stream, tonight) {
			t.Errorf("the reassembled stream is %d bytes and the night was %d",
				len(stream), len(tonight))
		}
		if rep.Complete.Health.ParseErrors != 0 {
			t.Errorf("health = %+v", rep.Complete.Health)
		}
		for n, f := range rep.Fights {
			if len(f.Events) == 0 || len(f.Metrics) == 0 || len(f.RawRange.SHA256) != 64 {
				t.Errorf("fight %d is incomplete: %d event bytes, %d metrics, hash %q",
					n, len(f.Events), len(f.Metrics), f.RawRange.SHA256)
				continue
			}
			// A fight's range addresses the report's own stream, so
			// the hash it declares is the hash of those bytes of it.
			r := f.RawRange
			if r.StartOffset < 0 || r.EndOffset > int64(len(stream)) || r.EndOffset <= r.StartOffset {
				t.Errorf("fight %d claims %d..%d of a %d-byte stream",
					n, r.StartOffset, r.EndOffset, len(stream))
				continue
			}
			sum := sha256.Sum256(stream[r.StartOffset:r.EndOffset])
			if got := hex.EncodeToString(sum[:]); got != r.SHA256 {
				t.Errorf("fight %d hashes %d..%d as %s, the bytes there hash to %s",
					n, r.StartOffset, r.EndOffset, r.SHA256, got)
			}
		}
	}
}
