// logs/engine/rating/golden_test.go
// Golden-fixture regression tests, reusing logs/engine/summary's own
// committed fixtures (they are exactly the shape rating.Score consumes)
// rather than inventing a second harness (spec §8's own instruction).
package rating

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/utility"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const regenEnv = "FOREVER_UPDATE_GOLDEN"

// noOpPercentiles simulates a source with no digest data at all -- every
// query reports ok=false -- so every component either falls back to its
// own absolute standard (RULING R5) or is excluded, deterministically and
// without a database. This exercises those fallback/exclusion paths
// against the real embedded curated data (logs/engine/mechanics's own
// tables, and this spec's new utility/consumables packages), which the
// hand-built worked-example fixtures in score_test.go do not.
type noOpPercentiles struct{}

func (noOpPercentiles) Placement(Bracket, float64) (float64, int64, bool) { return 0, 0, false }
func (noOpPercentiles) KillTimeBand(int64, int64, int64) (string, bool)   { return "", false }

type goldenDialect struct {
	name   string
	source string
}

// goldenDialects points at logs/engine/summary/testdata's own two fixture
// dialects (v16, v22), read as already-decoded summary.Summary JSON.
func goldenDialects() []goldenDialect {
	return []goldenDialect{
		{name: "v16", source: filepath.Join("..", "summary", "testdata", "v16.summary.json.golden")},
		{name: "v22", source: filepath.Join("..", "summary", "testdata", "v22.summary.json.golden")},
	}
}

func cardsGoldenFile(name string) string {
	return filepath.Join("testdata", name+".cards.json.golden")
}

// TestTheFixtureCardsMatchTheCommittedGolden scores every roster player of
// both summary fixtures and compares the resulting cards, as JSON, against
// a committed golden file -- the same regen-env-var, JSON-comparison
// pattern logs/engine/summary/golden_test.go established.
func TestTheFixtureCardsMatchTheCommittedGolden(t *testing.T) {
	for _, d := range goldenDialects() {
		t.Run(d.name, func(t *testing.T) {
			fight := loadFixtureSummary(t, d.source)
			cards := scoreEveryRosterPlayer(fight)
			got := marshalIndent(t, cards)

			path := cardsGoldenFile(d.name)
			if os.Getenv(regenEnv) != "" {
				if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(path, got, 0o644); err != nil {
					t.Fatal(err)
				}
				t.Logf("regenerated %s (%d bytes)", path, len(got))
				return
			}
			want, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("%v\nregenerate it with: %s=1 go test ./engine/rating/... -run %s", err, regenEnv, t.Name())
			}
			if string(got) != string(want) {
				t.Fatalf("the %s fixture's cards no longer match %s.\n"+
					"If the change is intended, regenerate with:\n"+
					"    %s=1 go test ./engine/rating/... -run %s\n"+
					"and review the diff in the commit.\ngot  %d bytes\nwant %d bytes",
					d.name, path, regenEnv, t.Name(), len(got), len(want))
			}
		})
	}
}

// TestTheGoldenCardsAreStableAcrossRuns is the in-process half: the same
// fixture scored twice must be byte-identical.
func TestTheGoldenCardsAreStableAcrossRuns(t *testing.T) {
	for _, d := range goldenDialects() {
		t.Run(d.name, func(t *testing.T) {
			fight := loadFixtureSummary(t, d.source)
			first := marshalIndent(t, scoreEveryRosterPlayer(fight))
			for i := 0; i < 5; i++ {
				again := marshalIndent(t, scoreEveryRosterPlayer(fight))
				if string(again) != string(first) {
					t.Fatalf("run %d of the fixture produced different cards", i)
				}
			}
		})
	}
}

func loadFixtureSummary(t *testing.T, path string) summary.Summary {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var s summary.Summary
	if err := json.Unmarshal(data, &s); err != nil {
		t.Fatal(err)
	}
	return s
}

// scoreEveryRosterPlayer builds real CuratedTables per player (the
// embedded mechanics table for the fixture's own encounter, if any; the
// embedded utility table for the player's own class/spec, if any; the
// embedded consumable catalogue's row for the player's own role) and scores
// every roster row against a source with no digest data at all.
func scoreEveryRosterPlayer(fight summary.Summary) map[string]Card {
	out := map[string]Card{}

	var mechPtr *mechanics.Table
	if t, ok := mechanics.Load(fight.EncounterID); ok {
		mechPtr = &t
	}
	catalogue := consumables.Load()

	for _, row := range fight.Roster {
		var utilPtr *utility.Table
		if row.Class != "" && row.Spec != "" {
			if t, ok := utility.Load(specSlug(row.Class, row.Spec)); ok {
				utilPtr = &t
			}
		}
		var catPtr *consumables.RoleCatalogue
		if rc, ok := catalogue.For(row.Role); ok {
			catPtr = &rc
		}
		tables := CuratedTables{Mechanics: mechPtr, Utility: utilPtr, Consumables: catPtr}
		out[row.GUID] = Score(fight, row.GUID, tables, noOpPercentiles{}, nil, DefaultModelInfo())
	}
	return out
}

func marshalIndent(t *testing.T, v any) []byte {
	t.Helper()
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	return append(b, '\n')
}
