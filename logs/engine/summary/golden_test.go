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

// goldenFileV22 is the committed summary for the v22 excerpt. A second
// dialect earns a second golden: the v16 one cannot catch a v22 field
// landing in the wrong column.
const goldenFileV22 = "testdata/v22.summary.json.golden"

// regenEnv regenerates the golden instead of comparing against it.
const regenEnv = "FOREVER_UPDATE_GOLDEN"

// goldenEngineVersion is deliberately not session.Version: the golden
// pins decoded output, and a version bump must not churn it.
const goldenEngineVersion = "golden"

// goldenDialect is one row of the table golden tests run over both
// dialects: the fixture to parse, the layout to parse it with, the base
// time the decoder is seeded with, and the golden file its summary must
// match.
type goldenDialect struct {
	name    string
	excerpt string
	layout  layout.Layout
	base    time.Time
	golden  string
}

// goldenDialects is the table both golden tests below run: one v16 case
// and one v22 case, kept as data rather than as four near-identical test
// bodies. Version 22's timestamps carry a year, so its base is never
// consulted; a deliberately wrong one is passed so that a regression that
// starts consulting it shows up as a wildly wrong date rather than as a
// plausible one.
func goldenDialects() []goldenDialect {
	return []goldenDialect{
		{
			name:    "v16",
			excerpt: filepath.Join("..", "event", "testdata", "v16.log"),
			layout:  layout.RetailV16(),
			base:    time.Date(2026, 9, 26, 0, 0, 0, 0, time.UTC),
			golden:  goldenFile,
		},
		{
			name:    "v22",
			excerpt: filepath.Join("..", "event", "testdata", "v22.log"),
			layout:  layout.RetailV22(),
			base:    time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			golden:  goldenFileV22,
		},
	}
}

// excerptSummary runs the given fixture through the real pipeline — lexer,
// decoder, unit registry, segmenter — and returns the summary of the first
// fight that closes.
func excerptSummary(t *testing.T, path string, lay layout.Layout, base time.Time) (fight.Fight, Summary) {
	t.Helper()
	text, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	dec := event.NewDecoder(lay, base)
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

// TestTheFixtureSummaryMatchesTheCommittedGolden runs both dialects' golden
// comparisons as subtests of one table: a schema reorder, a Kind rename, or
// a changed rounding rule passes the in-process determinism tests below but
// changes this committed file, on either dialect.
func TestTheFixtureSummaryMatchesTheCommittedGolden(t *testing.T) {
	for _, d := range goldenDialects() {
		t.Run(d.name, func(t *testing.T) {
			_, s := excerptSummary(t, d.excerpt, d.layout, d.base)
			got := append([]byte(jsonIndentOf(t, s)), '\n')

			if os.Getenv(regenEnv) != "" {
				if err := os.MkdirAll(filepath.Dir(d.golden), 0o755); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(d.golden, got, 0o644); err != nil {
					t.Fatal(err)
				}
				t.Logf("regenerated %s (%d bytes)", d.golden, len(got))
				return
			}

			want, err := os.ReadFile(d.golden)
			if err != nil {
				t.Fatalf("%v\nregenerate it with: %s=1 go test ./engine/summary/ -run %s", err, regenEnv, t.Name())
			}
			if string(got) != string(want) {
				t.Fatalf("the %s fixture's summary no longer matches %s.\n"+
					"If the change is intended, regenerate with:\n"+
					"    %s=1 go test ./engine/summary/ -run %s\n"+
					"and review the diff in the commit.\ngot  %d bytes\nwant %d bytes",
					d.name, d.golden, regenEnv, t.Name(), len(got), len(want))
			}
		})
	}
}

// TestTheGoldenSummaryIsStableAcrossRuns is the in-process half, over both
// dialects: the same fixture rendered twice must be byte-identical, which
// the golden alone would not prove if it happened to be regenerated from a
// lucky run.
func TestTheGoldenSummaryIsStableAcrossRuns(t *testing.T) {
	for _, d := range goldenDialects() {
		t.Run(d.name, func(t *testing.T) {
			_, first := excerptSummary(t, d.excerpt, d.layout, d.base)
			for i := 0; i < 10; i++ {
				_, again := excerptSummary(t, d.excerpt, d.layout, d.base)
				if jsonIndentOf(t, again) != jsonIndentOf(t, first) {
					t.Fatalf("run %d of the fixture produced a different summary", i)
				}
			}
		})
	}
}

// TestTheV22SummaryIsDatedFromTheLineNotTheBase proves the excerpt's own
// timestamps are used: the base handed to the decoder is in 2001 and the
// log is from 2026.
func TestTheV22SummaryIsDatedFromTheLineNotTheBase(t *testing.T) {
	f, _ := fixtureSummaryV22(t)
	if f.Start.Year() != 2026 {
		t.Fatalf("fight starts in %d, want 2026: the year came from the base, not the line", f.Start.Year())
	}
	if _, off := f.Start.Zone(); off != -5*3600 {
		t.Errorf("zone offset = %d seconds, want -18000", off)
	}
}

// fixtureSummaryV22 runs the hand-written v22 fixture through the real
// pipeline. It exists alongside goldenDialects for the one v22-only test
// above that is not part of the tabled golden comparisons.
func fixtureSummaryV22(t *testing.T) (fight.Fight, Summary) {
	t.Helper()
	for _, d := range goldenDialects() {
		if d.name == "v22" {
			return excerptSummary(t, d.excerpt, d.layout, d.base)
		}
	}
	t.Fatal("no v22 entry in goldenDialects")
	return fight.Fight{}, Summary{}
}
