package adapter

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

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
//	go run --tags=with_db ./internal/genfixture -spec <spec> -out adapter/testdata/<spec>.result.pb
//
// or, once Task 13 has built it and deleted genfixture:
//
//	./artifacts/forever-sim -in <spec>.request.json -out-proto adapter/testdata/<spec>.result.pb
//
// then run
//
//	FOREVER_UPDATE_GOLDEN=1 go test ./adapter/
//
// and read the diff before committing it.
func TestGoldenSummaries(t *testing.T) {
	for _, tc := range []struct{ spec, race, class string }{
		{"warrior-fury", "orc", "warrior"},
		{"mage-frost", "gnome", "mage"},
	} {
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
