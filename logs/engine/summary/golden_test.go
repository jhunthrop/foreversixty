// logs/engine/summary/golden_test.go
package summary

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/layout"
	"github.com/jhunthrop/foreversixty/logs/engine/lexer"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// goldenFile is the committed summary for the hand-written v16 fixture.
// The in-process determinism tests catch map-iteration nondeterminism but
// cannot see cross-machine or cross-dependency drift: a schema reorder, a
// Kind rename, a changed rounding rule all pass them and all change this
// file. JSON is portable text, so it behaves the same on every runner.
const goldenFile = "testdata/v16.summary.json.golden"

// regenEnv regenerates the golden instead of comparing against it.
const regenEnv = "FOREVER_UPDATE_GOLDEN"

// goldenEngineVersion is deliberately not session.Version: the golden
// pins decoded output, and a version bump must not churn it.
const goldenEngineVersion = "golden"

// fixtureSummary runs the hand-written v16 fixture through the real
// pipeline — lexer, retail v16 decoder, unit registry, segmenter — and
// returns the summary of the first fight that closes.
func fixtureSummary(t *testing.T) (fight.Fight, Summary) {
	t.Helper()
	text, err := os.ReadFile(filepath.Join("..", "event", "testdata", "v16.log"))
	if err != nil {
		t.Fatal(err)
	}
	// The fixture's timestamps carry no year, so the clock is seeded the
	// way a batch parse seeds it from the file's modification time.
	base := time.Date(2026, 9, 26, 0, 0, 0, 0, time.UTC)
	dec := event.NewDecoder(layout.RetailV16(), base)
	reg := units.NewRegistry(units.Options{ClassBySpec: units.RetailSpecClass})
	seg := fight.NewSegmenter(fight.DefaultOptions())

	o, _ := opts(t)
	o.Registry = reg
	var acc *Accumulator
	var got *Summary
	var closed fight.Fight

	// The same order session.handle uses: a fight that closes is
	// summarised from the accumulator that was open for it, and only then
	// is a new one started.
	finish := func(f fight.Fight) {
		if acc != nil && got == nil {
			s := acc.Snapshot(f, goldenEngineVersion)
			got, closed = &s, f
		}
		acc = nil
	}
	emit := func(ln lexer.Line) error {
		e := dec.Decode(ln)
		reg.Observe(e)
		step := seg.Feed(e)
		if step.Closed != nil {
			finish(*step.Closed)
		}
		if step.Fight == nil {
			return nil
		}
		if step.Opened {
			acc = New(o)
			acc.Start(step.Fight.Start)
		}
		if acc != nil {
			acc.Add(e)
		}
		return nil
	}
	lx := lexer.New()
	if err := lx.Feed(text, 0, emit); err != nil {
		t.Fatal(err)
	}
	if err := lx.Flush(emit); err != nil {
		t.Fatal(err)
	}
	if f := seg.Flush(dec.Time()); f != nil {
		finish(*f)
	}
	if got == nil {
		t.Fatal("the fixture produced no fight to summarise")
	}
	return closed, *got
}

func TestTheFixtureSummaryMatchesTheCommittedGolden(t *testing.T) {
	_, s := fixtureSummary(t)
	got := append([]byte(jsonIndentOf(t, s)), '\n')

	if os.Getenv(regenEnv) != "" {
		if err := os.MkdirAll(filepath.Dir(goldenFile), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(goldenFile, got, 0o644); err != nil {
			t.Fatal(err)
		}
		t.Logf("regenerated %s (%d bytes)", goldenFile, len(got))
		return
	}

	want, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatalf("%v\nregenerate it with: %s=1 go test ./engine/summary/ -run %s", err, regenEnv, t.Name())
	}
	if string(got) != string(want) {
		t.Fatalf("the fixture's summary no longer matches %s.\n"+
			"If the change is intended, regenerate with:\n"+
			"    %s=1 go test ./engine/summary/ -run %s\n"+
			"and review the diff in the commit.\ngot  %d bytes\nwant %d bytes",
			goldenFile, regenEnv, t.Name(), len(got), len(want))
	}
}

// TestTheGoldenSummaryIsStableAcrossRuns is the in-process half: the same
// fixture rendered twice must be byte-identical, which the golden alone
// would not prove if it happened to be regenerated from a lucky run.
func TestTheGoldenSummaryIsStableAcrossRuns(t *testing.T) {
	_, first := fixtureSummary(t)
	for i := 0; i < 10; i++ {
		_, again := fixtureSummary(t)
		if jsonIndentOf(t, again) != jsonIndentOf(t, first) {
			t.Fatalf("run %d of the fixture produced a different summary", i)
		}
	}
}
