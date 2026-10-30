package contract

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/pb33f/libopenapi"
)

// repoRoot is the repository root, found from this file's compiled-in path.
func repoRoot() string {
	_, self, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(self), "..", "..", "..")
}

// activeBuildDir is the data directory of the active client build, the generator's output the
// CDN files are copied from. A checkout without the data skips the CDN tests.
func activeBuildDir(t *testing.T) string {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join(repoRoot(), "web", "src", "data", "active-build.json"))
	if err != nil {
		t.Skipf("no active-build.json: %v", err)
	}
	var active struct {
		Build string `json:"build"`
	}
	if err := json.Unmarshal(raw, &active); err != nil || active.Build == "" {
		t.Fatalf("active-build.json: %v", err)
	}
	dir := filepath.Join(repoRoot(), "data", "builds", active.Build)
	if _, err := os.Stat(dir); err != nil {
		t.Skipf("no data for build %s: %v", active.Build, err)
	}
	return dir
}

func readJSON(t *testing.T, path string) any {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	v, err := decode(raw)
	if err != nil {
		t.Fatalf("%s: %v", path, err)
	}
	return v
}

// The fixtures of the static files are held to what the data generator really writes: every
// field a fixture names must be in every element of the generator's own file, with the same
// kind. A field the generator writes only sometimes is listed Optional.

func TestCDNActiveBuildFixtureMatchesTheFileTheSiteReads(t *testing.T) {
	want := Load(t, "cdn.active-build")
	got := readJSON(t, filepath.Join(repoRoot(), "web", "src", "data", "active-build.json"))
	if diffs := Compare(want, got, Options{}); len(diffs) > 0 {
		t.Fatalf("active-build.json: %v", diffs)
	}
}

func TestCDNDatesFixtureMatchesTheSitesDates(t *testing.T) {
	got := readJSON(t, filepath.Join(repoRoot(), "web", "src", "data", "dates.json"))
	opts := Options{Optional: []string{
		"[].note", "[].iso", "[].endIso", "[].time", "[].running", "[].role", "[].highlight", "[].at",
	}}
	if diffs := Compare(Load(t, "cdn.dates"), got, opts); len(diffs) > 0 {
		t.Fatalf("dates.json: %v", diffs)
	}
}

func TestCDNClassesFixtureMatchesTheGenerator(t *testing.T) {
	got := readJSON(t, filepath.Join(activeBuildDir(t), "classes.json"))
	if diffs := Compare(Load(t, "cdn.classes"), got, Options{}); len(diffs) > 0 {
		t.Fatalf("classes.json: %v", diffs)
	}
}

func TestCDNTalentsFixtureMatchesTheGenerator(t *testing.T) {
	got := readJSON(t, filepath.Join(activeBuildDir(t), "talents", "warrior.json"))
	if diffs := Compare(Load(t, "cdn.talents.warrior"), got, Options{}); len(diffs) > 0 {
		t.Fatalf("talents/warrior.json: %v", diffs)
	}
}

func TestCDNItemsFixtureMatchesTheGenerator(t *testing.T) {
	got := readJSON(t, filepath.Join(activeBuildDir(t), "items", "warrior.json"))
	if diffs := Compare(Load(t, "cdn.items.warrior"), got,
		Options{Open: []string{"items[].stats", "items[].percent_stats"}}); len(diffs) > 0 {
		t.Fatalf("items/warrior.json: %v", diffs)
	}
}

// bisOptional are the band and slot fields the ranker writes only when they apply.
var bisOptional = []string{
	"bands[].role", "bands[].slots[].item_id", "bands[].slots[].item_name", "bands[].slots[].source",
	"bands[].slots[].source_kind", "bands[].slots[].score", "bands[].slots[].alternatives",
	"bands[].slots[].chance", "bands[].slots[].sim_dps", "bands[].slots[].set_bonus",
	"bands[].slots[].swap_note", "bands[].slots[].dps_delta", "bands[].slots[].empty_reason",
	"bands[].slots[].effect_unmodelled", "bands[].slots[].ties", "bands[].slots[].faction_note",
	"bands[].slots[].low_value", "bands[].slots[].alternatives[].verified",
	"bands[].haste_scale_factor", "bands[].scale_reference_stat",
}

func TestCDNBisFixtureMatchesTheRanker(t *testing.T) {
	got := readJSON(t, filepath.Join(activeBuildDir(t), "bis", "warrior-fury.json"))
	if diffs := Compare(Load(t, "cdn.bis.warrior-fury"), got, Options{Optional: bisOptional}); len(diffs) > 0 {
		t.Fatalf("bis/warrior-fury.json: %v", diffs)
	}
}

// The API's own loader reads the BiS fixture, so the fixture is a file the real reader accepts.
func TestCDNBisFixtureLoadsThroughTheAPIReader(t *testing.T) {
	dir := t.TempDir()
	raw, err := os.ReadFile(FixturePath("cdn.bis.warrior-fury"))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(dir, "b1", "bis"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "b1", "bis", "warrior-fury.json"), raw, 0o644); err != nil {
		t.Fatal(err)
	}
	band, ok, err := bis.LoadBand(dir, "b1", "warrior-fury", 60, "horde")
	if err != nil || !ok {
		t.Fatalf("LoadBand = %v, %v", ok, err)
	}
	if len(band.Slots) == 0 || band.ScoreUnit == "" {
		t.Fatalf("band = %+v", band)
	}
}

// The apps decode the real files with models generated from api/openapi.yaml, so every BiS and
// items file the generator writes for the active build must match its schema, not only the
// fixture's slice of one spec: a kind the spec gets wrong (a tie's object read as an id) fails
// every app's decode of that file.
func TestCDNGeneratorFilesMatchTheOpenAPISpec(t *testing.T) {
	dir := activeBuildDir(t)
	raw, _ := loadSpec(t)
	document, err := libopenapi.NewDocument(raw)
	if err != nil {
		t.Fatal(err)
	}
	model, buildErr := document.BuildV3Model()
	if buildErr != nil {
		t.Fatal(buildErr)
	}
	for _, set := range []struct{ glob, schema string }{
		{"bis/*.json", "BisFile"},
		{"items/*.json", "ItemsFile"},
	} {
		proxy, ok := model.Model.Components.Schemas.Get(set.schema)
		if !ok {
			t.Fatalf("no schema %s", set.schema)
		}
		files, err := filepath.Glob(filepath.Join(dir, set.glob))
		if err != nil || len(files) == 0 {
			t.Fatalf("no %s in %s: %v", set.glob, dir, err)
		}
		for _, file := range files {
			t.Run(filepath.Join(filepath.Base(filepath.Dir(file)), filepath.Base(file)), func(t *testing.T) {
				body, err := os.ReadFile(file)
				if err != nil {
					t.Fatal(err)
				}
				assertSchema(t, proxy.Schema(), body)
			})
		}
	}
}
