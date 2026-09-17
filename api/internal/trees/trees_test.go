package trees

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadFixtureReadsTheTwoTreeClass(t *testing.T) {
	d, err := LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	if got := d.Versions(); len(got) != 1 || got[0] != "test-1" {
		t.Fatalf("versions = %v, want [test-1]", got)
	}
	b, ok := d.Build("test-1")
	if !ok {
		t.Fatal("build test-1 missing")
	}
	if b.Version != "test-1" {
		t.Fatalf("version = %q", b.Version)
	}
	if _, ok := d.Build("nope"); ok {
		t.Fatal("an unknown tree version must not resolve")
	}

	c, ok := b.Class(1)
	if !ok || c.Name != "Warrior" || c.Color != "#c69b6d" {
		t.Fatalf("class 1 = %+v ok=%v", c, ok)
	}
	r, ok := b.Race(3)
	if !ok || r.Name != "Dwarf" || r.Faction != "alliance" {
		t.Fatalf("race 3 = %+v ok=%v", r, ok)
	}
	if !b.ComboAllowed(1, 1) {
		t.Fatal("Human Warrior is in combos.json and must be allowed")
	}
	if b.ComboAllowed(1, 2) {
		t.Fatal("Human Paladin is not in combos.json and must not be allowed")
	}

	tt := b.Trees(1)
	if len(tt) != 2 || tt[0].Name != "Arms" || tt[1].Name != "Fury" {
		t.Fatalf("trees = %+v, want Arms then Fury by position", tt)
	}
	tal, ok := b.Talent(1, 103)
	if !ok || tal.Name != "Improved Rend" || tal.TreeName != "Arms" || tal.TreeID != 161 || tal.Tier != 1 || tal.MaxRank != 3 {
		t.Fatalf("talent 103 = %+v ok=%v", tal, ok)
	}
	if tal.PrereqTalentID == nil || *tal.PrereqTalentID != 101 || tal.PrereqRank == nil || *tal.PrereqRank != 5 {
		t.Fatalf("talent 103 prereq = %v %v", tal.PrereqTalentID, tal.PrereqRank)
	}
	if len(tal.Ranks) != 3 || tal.Ranks[0].SpellID != 12286 {
		t.Fatalf("talent 103 ranks = %+v", tal.Ranks)
	}
	if _, ok := b.Talent(1, 301); ok {
		t.Fatal("a Paladin talent must not resolve for the Warrior")
	}

	item, ok := b.Item(1, 12640)
	if !ok || item.Slot != "head" || item.Stats["strength"] != 18 || item.Unique {
		t.Fatalf("item 12640 = %+v ok=%v", item, ok)
	}
	if _, ok := b.Item(2, 12640); ok {
		t.Fatal("the Paladin has no items file, so no item may resolve for it")
	}
	sets := b.Sets()
	if len(sets) != 1 || sets[0].ID != 901 || len(sets[0].Bonuses) != 1 || sets[0].Bonuses[0].Pieces != 2 {
		t.Fatalf("sets = %+v", sets)
	}
}

const (
	oneClass = `[{"id":1,"name":"Warrior","slug":"warrior","color":"#c69b6d"}]`
	oneRace  = `[{"id":1,"name":"Human","slug":"human","faction":"alliance"}]`
	oneCombo = `[{"race_id":1,"class_id":1,"new_in_forever":false}]`
	oneTree  = `{"build":"b","class_id":1,"class_slug":"warrior","trees":[{"id":161,"name":"Arms","position":0,` +
		`"talents":[{"id":101,"name":"A","icon":"i","max_rank":2,"tier":0,"column":0,"prereq_talent_id":null,` +
		`"prereq_rank":null,"ranks":[{"spell_id":1,"description":"one"},{"spell_id":2,"description":"two"}]}]}]}`
)

// writeBuild writes files (keyed by slash-separated path relative to a fresh
// temp root) and returns the root, for the loader's failure cases.
func writeBuild(t *testing.T, files map[string]string) string {
	t.Helper()
	root := t.TempDir()
	for name, body := range files {
		path := filepath.Join(root, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return root
}

func TestLoadSkipsBuildDirectoriesWithoutPhase1Files(t *testing.T) {
	root := writeBuild(t, map[string]string{"b/classes.json": oneClass, "b/races.json": oneRace})
	d, err := Load(root)
	if err != nil {
		t.Fatalf("a pre-Phase-1 build directory must be skipped, not fail the load: %v", err)
	}
	if len(d.Versions()) != 0 {
		t.Fatalf("versions = %v, want none", d.Versions())
	}
	if len(d.Skipped()) != 1 || !strings.Contains(d.Skipped()[0], "b") {
		t.Fatalf("skipped = %v", d.Skipped())
	}
}

func TestLoadReportsAMissingDirectoryAsSkippedNotAnError(t *testing.T) {
	d, err := Load(filepath.Join(t.TempDir(), "absent"))
	if err != nil {
		t.Fatalf("a missing TREE_DATA_DIR must not fail the load: %v", err)
	}
	if len(d.Skipped()) != 1 {
		t.Fatalf("skipped = %v, want one entry", d.Skipped())
	}
}

func TestLoadRejectsARankCountThatDisagreesWithMaxRank(t *testing.T) {
	bad := strings.Replace(oneTree, `"max_rank":2`, `"max_rank":3`, 1)
	root := writeBuild(t, map[string]string{
		"b/classes.json": oneClass, "b/races.json": oneRace, "b/combos.json": oneCombo,
		"b/talents/warrior.json": bad,
	})
	if _, err := Load(root); err == nil || !strings.Contains(err.Error(), "ranks") {
		t.Fatalf("err = %v, want a ranks/max_rank mismatch", err)
	}
}

func TestLoadRejectsCombosReferencingAnUnknownClass(t *testing.T) {
	root := writeBuild(t, map[string]string{
		"b/classes.json": oneClass, "b/races.json": oneRace,
		"b/combos.json":          `[{"race_id":1,"class_id":9,"new_in_forever":false}]`,
		"b/talents/warrior.json": oneTree,
	})
	if _, err := Load(root); err == nil || !strings.Contains(err.Error(), "unknown class 9") {
		t.Fatalf("err = %v, want unknown class 9", err)
	}
}

func TestLoadRejectsAMissingTalentsFileForAKnownClass(t *testing.T) {
	root := writeBuild(t, map[string]string{
		"b/classes.json": oneClass, "b/races.json": oneRace, "b/combos.json": oneCombo,
		"b/talents/other.json": oneTree,
	})
	if _, err := Load(root); err == nil || !strings.Contains(err.Error(), "warrior.json") {
		t.Fatalf("err = %v, want the missing warrior.json named", err)
	}
}

func TestLoadRejectsDuplicateTalentIDs(t *testing.T) {
	dup := strings.Replace(oneTree, `"talents":[`, `"talents":[{"id":101,"name":"A","icon":"i","max_rank":1,"tier":0,`+
		`"column":2,"prereq_talent_id":null,"prereq_rank":null,"ranks":[{"spell_id":9,"description":"x"}]},`, 1)
	root := writeBuild(t, map[string]string{
		"b/classes.json": oneClass, "b/races.json": oneRace, "b/combos.json": oneCombo,
		"b/talents/warrior.json": dup,
	})
	if _, err := Load(root); err == nil || !strings.Contains(err.Error(), "duplicate talent id 101") {
		t.Fatalf("err = %v, want duplicate talent id 101", err)
	}
}

func TestLoadRejectsDuplicateTalentSpellIDs(t *testing.T) {
	dup := `{"build":"b","class_id":1,"class_slug":"warrior","trees":[{"id":161,"name":"Arms","position":0,` +
		`"talents":[` +
		`{"id":101,"name":"A","icon":"i","max_rank":1,"tier":0,"column":0,"prereq_talent_id":null,"prereq_rank":null,` +
		`"spell_id":9,"ranks":[{"spell_id":9,"description":"one"}]},` +
		`{"id":102,"name":"B","icon":"i","max_rank":1,"tier":0,"column":1,"prereq_talent_id":null,"prereq_rank":null,` +
		`"spell_id":9,"ranks":[{"spell_id":9,"description":"two"}]}` +
		`]}]}`
	root := writeBuild(t, map[string]string{
		"b/classes.json": oneClass, "b/races.json": oneRace, "b/combos.json": oneCombo,
		"b/talents/warrior.json": dup,
	})
	if _, err := Load(root); err == nil || !strings.Contains(err.Error(), "duplicate talent spell id 9") {
		t.Fatalf("err = %v, want duplicate talent spell id 9", err)
	}
}

func TestClassesAreListedByID(t *testing.T) {
	data, err := LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	b, ok := data.Build("test-1")
	if !ok {
		t.Fatal("the fixture build should load")
	}
	classes := b.Classes()
	if len(classes) == 0 {
		t.Fatal("the fixture has classes")
	}
	for i := 1; i < len(classes); i++ {
		if classes[i-1].ID >= classes[i].ID {
			t.Fatalf("classes are not ordered by id: %v", classes)
		}
	}
}

func TestLatestIsTheNewestBuild(t *testing.T) {
	data, err := LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	b, ok := data.Latest()
	if !ok {
		t.Fatal("the fixture build should be the latest")
	}
	versions := data.Versions()
	if b.Version != versions[len(versions)-1] {
		t.Fatalf("latest = %q, want %q", b.Version, versions[len(versions)-1])
	}
	if _, ok := (&Data{}).Latest(); ok {
		t.Fatal("no builds means no latest")
	}
}

func TestLatestComparesVersionsNumericallyNotLexicographically(t *testing.T) {
	// Lexicographically these sort "1.10" < "1.15.10.1" < "1.15.9.69722" <
	// "1.2" < "1.9", which would pick "1.9" as "latest" - wrong, since the
	// project's own build directories are dotted version numbers like
	// "1.15.9.69722" and a numeric comparison must win.
	files := map[string]string{}
	for _, v := range []string{"1.2", "1.9", "1.10", "1.15.9.69722", "1.15.10.1"} {
		files[v+"/classes.json"] = oneClass
		files[v+"/races.json"] = oneRace
		files[v+"/combos.json"] = oneCombo
		files[v+"/talents/warrior.json"] = oneTree
	}
	root := writeBuild(t, files)
	d, err := Load(root)
	if err != nil {
		t.Fatal(err)
	}
	b, ok := d.Latest()
	if !ok {
		t.Fatal("expected a latest build")
	}
	if b.Version != "1.15.10.1" {
		t.Fatalf("latest = %q, want the numerically newest 1.15.10.1", b.Version)
	}
}

func TestTalentCarriesItsSpellIDAndTreeItsBackground(t *testing.T) {
	data, err := LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	b, ok := data.Build("test-1")
	if !ok {
		t.Fatal("fixture build test-1 did not load")
	}
	arms := b.Trees(1)[0]
	if arms.Background != "warriorarms" {
		t.Fatalf("Arms background = %q, want %q", arms.Background, "warriorarms")
	}
	ref, ok := b.Talent(1, 101)
	if !ok {
		t.Fatal("talent 101 is missing")
	}
	if ref.SpellID != 12281 {
		t.Fatalf("talent 101 spell id = %d, want 12281", ref.SpellID)
	}
}

func TestTalentBySpellIDFindsTheTalentTheClientWrites(t *testing.T) {
	data, err := LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	b, _ := data.Build("test-1")
	ref, ok := b.TalentBySpellID(1, 12281)
	if !ok || ref.ID != 101 {
		t.Fatalf("TalentBySpellID(1, 12281) = %+v, %v; want talent 101", ref, ok)
	}
	if _, ok := b.TalentBySpellID(1, 999999); ok {
		t.Fatal("an unknown spell id must report false")
	}
	if _, ok := b.TalentBySpellID(99, 12281); ok {
		t.Fatal("an unknown class must report false")
	}
}

func TestLatestPrefersAClientBuildOverANamedDataSet(t *testing.T) {
	for _, c := range []struct {
		versions []string
		want     string
	}{
		{[]string{"1.15.9.69722", "1.60.1.69893", "forever-prebeta"}, "1.60.1.69893"},
		{[]string{"1.15.9.69722", "1.9.1.1"}, "1.15.9.69722"},
		{[]string{"forever-prebeta"}, "forever-prebeta"},
	} {
		if got := newestVersion(c.versions); got != c.want {
			t.Fatalf("newestVersion(%v) = %q, want %q", c.versions, got, c.want)
		}
	}
}
