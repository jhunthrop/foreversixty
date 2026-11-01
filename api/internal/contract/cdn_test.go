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

// W3-C10: the planner's Race sheet reads races.json and combos.json.
func TestCDNRacesAndCombosFixturesMatchTheGenerator(t *testing.T) {
	dir := activeBuildDir(t)
	for _, f := range []struct{ fixture, file string }{
		{"cdn.races", "races.json"},
		{"cdn.combos", "combos.json"},
	} {
		got := readJSON(t, filepath.Join(dir, f.file))
		if diffs := Compare(Load(t, f.fixture), got, Options{}); len(diffs) > 0 {
			t.Fatalf("%s: %v", f.file, diffs)
		}
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

// itemFactsPending are the spec's C6 minimum standing and the weapon's subclass and slot words
// (lane F2's fix round). The hunter fixture carries them, read from the client's Item and
// ItemSparse tables, so the apps' tooltips draw them; the generator does not publish them yet,
// so they are optional here. `bonding` and `armor_type` are published (lane F12) and required.
var itemFactsPending = []string{
	"items[].min_standing", "items[].subclass", "items[].weapon_slot",
}

func TestCDNHunterItemsFixtureMatchesTheGenerator(t *testing.T) {
	got := readJSON(t, filepath.Join(activeBuildDir(t), "items", "hunter.json"))
	opts := Options{Open: []string{"items[].stats", "items[].percent_stats"}, Optional: itemFactsPending}
	if diffs := Compare(Load(t, "cdn.items.hunter"), got, opts); len(diffs) > 0 {
		t.Fatalf("items/hunter.json: %v", diffs)
	}
}

func TestCDNSetsFixtureMatchesTheGenerator(t *testing.T) {
	got := readJSON(t, filepath.Join(activeBuildDir(t), "sets.json"))
	if diffs := Compare(Load(t, "cdn.sets"), got, Options{}); len(diffs) > 0 {
		t.Fatalf("sets.json: %v", diffs)
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

func TestCDNHunterBisFixtureMatchesTheRanker(t *testing.T) {
	got := readJSON(t, filepath.Join(activeBuildDir(t), "bis", "hunter-beast-mastery.json"))
	// The formatted source line and drop chance are pending in the ranker (lane F2's fix round).
	pending := append([]string{
		"bands[].slots[].source_line", "bands[].slots[].drop_chance",
		"bands[].slots[].alternatives[].source_line", "bands[].slots[].alternatives[].drop_chance",
	}, bisOptional...)
	if diffs := Compare(Load(t, "cdn.bis.hunter-beast-mastery"), got, Options{Optional: pending}); len(diffs) > 0 {
		t.Fatalf("bis/hunter-beast-mastery.json: %v", diffs)
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

// lootOptional are the source, boss and quest fields loot.json writes only where they apply.
var lootOptional = []string{
	"sources[].items", "sources[].trash", "sources[].bosses", "sources[].item_chances", "sources[].zone_id",
	"sources[].npc_id", "sources[].opens", "sources[].profession", "sources[].faction_id", "sources[].standing",
	"sources[].rank", "sources[].faction", "sources[].faction_source", "sources[].title", "sources[].level_min",
	"sources[].level_max", "sources[].rare_spawn", "sources[].source_origin", "sources[].source_note",
	"sources[].reitemised_from", "sources[].bosses[].item_chances", "sources[].bosses[].item_source_origin",
	"sources[].bosses[].reitemised_from", "sources[].bosses[].rank", "sources[].bosses[].rare_spawn",
}

// W3-C8: the loot file the site already publishes, now an app contract.
func TestCDNLootFixtureMatchesTheGenerator(t *testing.T) {
	got := readJSON(t, filepath.Join(activeBuildDir(t), "loot.json"))
	opts := Options{
		Optional: lootOptional,
		Open:     []string{"quests", "factions", "sources[].item_chances", "sources[].reitemised_from", "sources[].bosses[].item_chances", "sources[].bosses[].item_source_origin", "sources[].bosses[].reitemised_from"},
	}
	if diffs := Compare(Load(t, "cdn.loot"), got, opts); len(diffs) > 0 {
		t.Fatalf("loot.json: %v", diffs)
	}
}

// minBondingShare is the floor on the share of item rows whose Binds line is known (spec F12.4:
// the tooltip prints the client's Binds line). 70707 publishes it for 97% of rows; the rest are
// wowhead-only rows with no classic-db twin, which print no line rather than a guess.
const minBondingShare = 0.9

func TestCDNItemsCarryTheBindsLine(t *testing.T) {
	files, err := filepath.Glob(filepath.Join(activeBuildDir(t), "items", "*.json"))
	if err != nil || len(files) == 0 {
		t.Fatalf("no items files: %v", err)
	}
	for _, file := range files {
		var doc struct {
			Items []struct {
				Bonding *string `json:"bonding"`
			} `json:"items"`
		}
		raw, err := os.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(raw, &doc); err != nil {
			t.Fatalf("%s: %v", file, err)
		}
		known := 0
		for _, item := range doc.Items {
			if item.Bonding != nil {
				known++
			}
		}
		if share := float64(known) / float64(len(doc.Items)); share < minBondingShare {
			t.Errorf("%s: %.0f%% of items state a binding, under the %.0f%% floor", filepath.Base(file), share*100, minBondingShare*100)
		}
	}
}
