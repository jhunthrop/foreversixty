package adapter

import (
	"encoding/json"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// goldenSpecs are the specs a fixture and a golden are checked in for.
var goldenSpecs = []struct{ spec, race, class string }{
	{"warrior-fury", "orc", "warrior"},
	{"mage-frost", "gnome", "mage"},
}

// regenEnv regenerates the goldens instead of comparing against them. The
// same variable name the logs engine uses, so one habit covers both.
const regenEnv = "FOREVER_UPDATE_GOLDEN"

// goldenEngineVersion is deliberately not enginever.Version: the golden
// pins the adapter's output shape, and an engine pin bump must not churn
// every golden file.
const goldenEngineVersion = "golden"

// To regenerate a fixture after a spec's abilities change, from the
// SITE module (never from the engine checkout):
//
//	make artifacts                                   # at the pinned sha
//	cd sim
//	./../artifacts/forever-sim \
//	  -in adapter/testdata/<spec>.request.json \
//	  -out-proto adapter/testdata/<spec>.result.pb
//	FOREVER_UPDATE_GOLDEN=1 go test ./adapter/
//
// and read the diff before committing it.
//
// The .request.json beside each fixture is the input: a real SimRequest,
// so a fixture is a real sim of a request the product could send rather
// than a second opinion about what a fury warrior is. Its buffs are the
// field names of core.Full{Raid,Individual,Debuffs} - the same set,
// though a tristate lands on its plain form because the settings bar has
// no id for the improved one (see request/buffs.go). Its gear starts
// from the engine's own phase-one preset, with every id Forever
// re-itemised away replaced by the nearest row in the same slot.
//
// The gear is real Forever item ids, resolved from the active build's
// database, which sim/internal/simdb embeds: run `make simdb` first if
// the working tree has never had it (`make artifacts` depends on it).
// Do NOT build with --tags=with_db - that is the engine's own vanilla
// table, which resolves almost none of a Forever gear set and which
// makes the engine panic at init over item effects this build has no
// rows for. Each slot's provenance is in <spec>.request.notes.md.
func TestGoldenSummaries(t *testing.T) {
	for _, tc := range goldenSpecs {
		t.Run(tc.spec, func(t *testing.T) {
			res, err := Fixture(tc.spec)
			if err != nil {
				t.Fatal(err)
			}
			req := api.SimRequest{
				EngineVersion: goldenEngineVersion,
				Spec:          tc.spec,
				Character:     api.CharacterSpec{Name: "Sim", Race: tc.race, Class: tc.class, Level: 60},
				Encounter:     api.DefaultEncounter(),
				Iterations:    3000,
			}
			got, err := Summarize(res, req)
			if err != nil {
				t.Fatal(err)
			}
			// The stamp is a FACT about the binary, never the
			// request's claim, or a cached request pinned to an old
			// sha would come back labelled with it and
			// SimResult.Stale could never fire. It is asserted here
			// and then normalised away, so the golden keeps pinning
			// the adapter's shape without churning on every pin.
			if want := "sim:" + enginever.Version; got.EngineVersion != want {
				t.Errorf("EngineVersion = %q, want %q", got.EngineVersion, want)
			}
			got.EngineVersion = "sim:" + goldenEngineVersion
			b, err := json.MarshalIndent(got, "", "  ")
			if err != nil {
				t.Fatal(err)
			}
			b = append(b, '\n')

			path := filepath.Join("testdata", tc.spec+".summary.json.golden")
			if os.Getenv(regenEnv) != "" {
				if err := os.WriteFile(path, b, 0o644); err != nil {
					t.Fatal(err)
				}
				t.Logf("regenerated %s", path)
				return
			}
			want, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("%v (run with %s=1 to create it)", err, regenEnv)
			}
			if string(b) != string(want) {
				t.Errorf("summary for %s differs from the golden.\n--- got ---\n%s\n--- want ---\n%s", tc.spec, b, want)
			}
		})
	}
}

// The golden must carry every key logs engine 0.5.3 puts in a summary, or
// the report page reads a field the sim never set. The list is the
// Summary struct's json tags, in its own order.
func TestGoldenCarriesEverySummaryKey(t *testing.T) {
	b, err := os.ReadFile(filepath.Join("testdata", "warrior-fury.summary.json.golden"))
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatal(err)
	}
	want := []string{
		"engine_version", "fight_index", "duration_ms",
		"damage_done", "damage_taken", "healing", "healing_taken",
		"deaths", "auras", "casts", "interrupts", "dispels", "resources",
		"threat", "threat_by_target", "taunts", "combatants", "roster",
		"mechanics", "phases",
	}
	for _, k := range want {
		if _, ok := m[k]; !ok {
			t.Errorf("the golden summary is missing %q", k)
		}
	}
	if len(m) != len(want) {
		t.Errorf("the golden has %d keys, the Summary struct has %d; logs engine 0.5.3 has changed shape and this adapter needs a look", len(m), len(want))
	}
}

// Summarizing is deterministic: the same result and request must produce
// byte-identical JSON, or the report page would churn between loads.
func TestSummarizeIsDeterministic(t *testing.T) {
	res, err := Fixture("warrior-fury")
	if err != nil {
		t.Fatal(err)
	}
	req := api.SimRequest{EngineVersion: goldenEngineVersion, Spec: "warrior-fury", Encounter: api.DefaultEncounter(), Iterations: 3000}
	var first string
	for i := 0; i < 20; i++ {
		s, err := Summarize(res, req)
		if err != nil {
			t.Fatal(err)
		}
		b, err := json.Marshal(s)
		if err != nil {
			t.Fatal(err)
		}
		if i == 0 {
			first = string(b)
			continue
		}
		if string(b) != first {
			t.Fatalf("run %d differs from run 0; a map is being ranged without sorting", i)
		}
	}
}

// Every row of a golden must carry a key no other row of its table
// carries. The logs engine's own maps make a duplicate impossible for a
// real fight, and the report components key their {#each} blocks on
// exactly these tuples, where Svelte 5 raises each_key_duplicate at
// runtime: ActorRow.svelte and AbilityBar.svelte on
// abilityKey = `${spell_id}|${via}` within one actor's rows,
// CastTable.svelte on `${guid}-${spell_id}`, AuraTable.svelte on
// `${target_guid}-${spell_id}`. This walks the committed goldens, which
// is the artefact the web lane renders.
func TestGoldenRowKeysAreUnique(t *testing.T) {
	for _, tc := range goldenSpecs {
		t.Run(tc.spec, func(t *testing.T) {
			b, err := os.ReadFile(filepath.Join("testdata", tc.spec+".summary.json.golden"))
			if err != nil {
				t.Fatal(err)
			}
			var got summary.Summary
			if err := json.Unmarshal(b, &got); err != nil {
				t.Fatal(err)
			}
			for _, a := range got.DamageDone {
				seen := map[string]bool{}
				for _, ab := range a.Abilities {
					k := key(ab.SpellID, a.GUID, ab.Via)
					if seen[k] {
						t.Errorf("actor %q has two ability rows keyed %s (%q)", a.Name, k, ab.Name)
					}
					seen[k] = true
				}
			}
			seenCasts := map[string]bool{}
			for _, c := range got.Casts {
				k := key(c.SpellID, c.GUID, "")
				if seenCasts[k] {
					t.Errorf("two cast rows keyed %s (%q)", k, c.SpellName)
				}
				seenCasts[k] = true
			}
			seenAuras := map[string]bool{}
			for _, au := range got.Auras {
				k := key(au.SpellID, au.TargetGUID, "")
				if seenAuras[k] {
					t.Errorf("two aura rows keyed %s (%q)", k, au.Name)
				}
				seenAuras[k] = true
			}
			// A row id is either a real spell id the web can resolve or
			// one of ours from the reserved space; nothing in between.
			for _, a := range got.DamageDone {
				for _, ab := range a.Abilities {
					if ab.SpellID <= 0 {
						t.Errorf("ability %q has no id; every row must be identifiable", ab.Name)
					}
				}
			}
		})
	}
}

func key(id int64, scope, via string) string {
	return scope + "|" + via + "|" + strconv.FormatInt(id, 10)
}

// durationTolerance is how far "table damage / table duration" may sit
// from the headline DPS(res).Mean. It is not zero because DurationMS is
// an integer number of milliseconds and the division that recovers a
// rate from it re-introduces sub-millisecond rounding, but it must stay
// tiny - well under the 0.1 DPS a reader could ever notice on a card -
// or the derivation in duration.go is not doing its job. 0.01 DPS is
// three orders of magnitude below the smallest gap TestGoldenSummaries'
// fixtures show before this task's fix (0.01-0.7% of a few hundred
// DPS), and two orders of magnitude below what the tightest pre-fix
// fixture (simarms-3t, 0.003%) already achieved by coincidence.
const durationTolerance = 0.01

// This is the test the persona reviews are asking for: one card, one
// number. The headline is DPS(res).Mean; the table's implied rate is
// the summary's own total damage - every actor in DamageDone, player
// and pets - divided by its own duration. Task 3 exists because those
// two used to disagree by up to ~1%; this proves they no longer do, for
// every checked-in fixture, to a tolerance so tight only the
// integer-millisecond field itself could produce it.
func TestHeadlineEqualsTable(t *testing.T) {
	for _, tc := range goldenSpecs {
		t.Run(tc.spec, func(t *testing.T) {
			res, err := Fixture(tc.spec)
			if err != nil {
				t.Fatal(err)
			}
			req := api.SimRequest{
				EngineVersion: goldenEngineVersion,
				Spec:          tc.spec,
				Character:     api.CharacterSpec{Name: "Sim", Race: tc.race, Class: tc.class, Level: 60},
				Encounter:     api.DefaultEncounter(),
				Iterations:    3000,
			}
			got, err := Summarize(res, req)
			if err != nil {
				t.Fatal(err)
			}
			meanDPS := DPS(res).Mean

			var tableTotal int64
			for _, a := range got.DamageDone {
				tableTotal += a.Total
			}
			if got.DurationMS <= 0 {
				t.Fatalf("DurationMS = %d, want a positive duration", got.DurationMS)
			}
			tablePerSec := float64(tableTotal) / (float64(got.DurationMS) / 1000)

			if diff := math.Abs(tablePerSec - meanDPS); diff > durationTolerance {
				t.Errorf("table damage/sec = %v, headline DPS(res).Mean = %v, disagree by %v (want <= %v)",
					tablePerSec, meanDPS, diff, durationTolerance)
			}
		})
	}
}

// D4: the per-target sub-table must sum to exactly the row it belongs
// to, for every actor of every fixture - the tank review's other
// complaint ("30,147 the table total ... 30,144 the sub-table sums to").
func TestPerTargetTotalsSumToTheActorTotal(t *testing.T) {
	for _, tc := range goldenSpecs {
		t.Run(tc.spec, func(t *testing.T) {
			res, err := Fixture(tc.spec)
			if err != nil {
				t.Fatal(err)
			}
			req := api.SimRequest{
				EngineVersion: goldenEngineVersion,
				Spec:          tc.spec,
				Character:     api.CharacterSpec{Name: "Sim", Race: tc.race, Class: tc.class, Level: 60},
				Encounter:     api.DefaultEncounter(),
				Iterations:    3000,
			}
			got, err := Summarize(res, req)
			if err != nil {
				t.Fatal(err)
			}
			for _, a := range got.DamageDone {
				var sum int64
				for _, target := range a.Targets {
					sum += target.Total
				}
				if sum != a.Total {
					t.Errorf("actor %q: sum(Targets[].Total) = %d, want Total = %d", a.Name, sum, a.Total)
				}
			}
		})
	}
}
