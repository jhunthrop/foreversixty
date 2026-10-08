package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

const realRoot = "../../.."

func writeFile(t *testing.T, path, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
}

func TestLoadSpecFindsBySlugAndReportsFailures(t *testing.T) {
	root := t.TempDir()
	if _, err := loadSpec(root, "a"); err == nil {
		t.Fatal("missing file accepted")
	}
	path := filepath.Join(root, "data", "curated", "specs.json")
	writeFile(t, path, `[{"spec":"a","class_slug":"mage","role":"dps","tree_index":1},{"spec":"b","role":"healer"}]`)
	if got, err := loadSpec(root, "a"); err != nil || got.ClassSlug != "mage" || got.TreeIndex != 1 {
		t.Fatalf("got %+v %v", got, err)
	}
	if _, err := loadSpec(root, "zzz"); err == nil || !strings.Contains(err.Error(), `no spec "zzz"`) {
		t.Fatalf("unknown: %v", err)
	}
	writeFile(t, path, `{`)
	if _, err := loadSpec(root, "a"); err == nil {
		t.Fatal("bad json accepted")
	}
}

func TestLoadCuratedAPLChecksTheFilesOwnSpec(t *testing.T) {
	root := t.TempDir()
	if _, err := loadCuratedAPL(root, "a"); err == nil {
		t.Fatal("missing file accepted")
	}
	path := filepath.Join(root, "data", "curated", "apl", "a.json")
	writeFile(t, path, `{"spec":"a","state":"written","rotation":{"type":"TypeAPL"}}`)
	got, err := loadCuratedAPL(root, "a")
	if err != nil || got.State != writtenAPL || !strings.Contains(string(got.Rotation), "TypeAPL") {
		t.Fatalf("got %+v %v", got, err)
	}
	writeFile(t, path, `{"spec":"other","state":"written"}`)
	if _, err := loadCuratedAPL(root, "a"); err == nil || !strings.Contains(err.Error(), `declares spec "other"`) {
		t.Fatalf("mislabelled file: %v", err)
	}
	writeFile(t, path, `[`)
	if _, err := loadCuratedAPL(root, "a"); err == nil {
		t.Fatal("bad json accepted")
	}
}

func TestLoadBISBandMatchesLevelFactionAndPreset(t *testing.T) {
	dir := t.TempDir()
	if _, err := loadBISBand(dir, "a", 60, "alliance", "bare"); err == nil {
		t.Fatal("missing file accepted")
	}
	path := filepath.Join(dir, "bis", "a.json")
	writeFile(t, path, `{"bands":[
	  {"band":60,"preset":"bare","faction":"horde","race":"Orc"},
	  {"band":60,"preset":"raid","faction":"alliance","race":"Gnome"},
	  {"band":60,"preset":"bare","faction":"alliance","race":"Human","slots":[{"slot":"head","item_id":5},{"slot":"neck","item_id":0}]}]}`)
	band, err := loadBISBand(dir, "a", 60, "alliance", "bare")
	if err != nil || band.Race != "Human" {
		t.Fatalf("got %+v %v", band, err)
	}
	if gear := band.gear(); len(gear) != 1 || gear[0].Slot != "head" || gear[0].ItemID != 5 {
		t.Fatalf("empty slots must be dropped, got %+v", gear)
	}
	if _, err := loadBISBand(dir, "a", 40, "alliance", "bare"); err == nil || !strings.Contains(err.Error(), "no level-40") {
		t.Fatalf("absent band: %v", err)
	}
	writeFile(t, path, `x`)
	if _, err := loadBISBand(dir, "a", 60, "alliance", "bare"); err == nil {
		t.Fatal("bad json accepted")
	}
}

func realOptions(spec string) options {
	return options{repoRoot: realRoot, spec: spec, level: 60, faction: "alliance", preset: "bare", seed: 7, iterations: 2, confirmIterations: 2, rounds: 1}
}

func TestPrepareSkipsHealersAndFailsOnUnknownSpecs(t *testing.T) {
	if _, err := prepare(realOptions("priest-holy")); !errors.Is(err, errSkipped) {
		t.Errorf("healer: %v", err)
	}
	if _, err := prepare(realOptions("no-such-spec")); err == nil || errors.Is(err, errSkipped) {
		t.Errorf("unknown spec: %v", err)
	}
}

func TestPrepareReadsTheCuratedRotationAndTheGuidesTalents(t *testing.T) {
	in, err := prepare(realOptions("rogue-combat"))
	if err != nil {
		t.Fatal(err)
	}
	if len(in.base.PriorityList) == 0 || in.clientBuild == "" || in.setup.engineTalents == "" || len(in.setup.band.Slots) == 0 {
		t.Fatalf("incomplete inputs: %+v", in.setup)
	}
	if authoredCastIDs(in.base)[0] {
		t.Fatal("authored ids must be real spell ids")
	}
	for _, c := range in.candidates {
		if authoredCastIDs(in.base)[c.ID] {
			t.Errorf("candidate %s is already cast by the rotation", c.Name)
		}
	}
}

func TestPrepareWearsAnOverrideBuildCodeInsteadOfTheGuide(t *testing.T) {
	o := realOptions("rogue-combat")
	guide, err := prepare(o)
	if err != nil {
		t.Fatal(err)
	}
	o.buildCode = "FS1:" + guide.clientBuild + ":rogue:human:0/0/0:"
	bare, err := prepare(o)
	if err != nil {
		t.Fatal(err)
	}
	if bare.setup.engineTalents == guide.setup.engineTalents {
		t.Fatal("an empty override build produced the guide's talents")
	}
	o.buildCode = "garbage"
	if _, err := prepare(o); err == nil {
		t.Fatal("a malformed build code was accepted")
	}
}

func TestBuildSpellNamesAndTickSecondsComeFromTheLearnedTable(t *testing.T) {
	in, err := prepare(realOptions("mage-fire"))
	if err != nil {
		t.Fatal(err)
	}
	names, err := buildSpellNames(realRoot, in.clientBuild, "mage", 60)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, n := range names {
		found = found || n == "Fireball"
	}
	if !found {
		t.Fatal("Fireball is not among a mage's learned abilities")
	}
	if len(in.tickSeconds) == 0 {
		t.Fatal("a mage learns damage-over-time spells")
	}
	for id, s := range in.tickSeconds {
		if s <= 0 {
			t.Errorf("spell %d has tick %v", id, s)
		}
	}
	if _, err := buildSpellNames(t.TempDir(), in.clientBuild, "mage", 60); err == nil {
		t.Fatal("a repo without ability data accepted")
	}
}

func TestSetupCharacterPutsCastersOutOfMeleeRange(t *testing.T) {
	in, err := prepare(realOptions("mage-fire"))
	if err != nil {
		t.Fatal(err)
	}
	ch := in.setup.character()
	if ch.Class != "mage" || ch.Level != 60 || len(ch.Gear) == 0 || ch.Talents != in.setup.engineTalents || ch.DistanceFromTarget == 0 {
		t.Fatalf("character = %+v", ch)
	}
	melee, err := prepare(realOptions("rogue-combat"))
	if err != nil {
		t.Fatal(err)
	}
	if melee.setup.character().DistanceFromTarget != 0 {
		t.Fatal("a melee spec must stand at its target")
	}
}

func TestEngineRunMemoisesByRotationAndIterations(t *testing.T) {
	in, err := prepare(realOptions("rogue-combat"))
	if err != nil {
		t.Fatal(err)
	}
	runDPS, err := engineRun(in.setup)
	if err != nil {
		t.Fatal(err)
	}
	a, err := runDPS(in.base, 2)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := runDPS(in.base, 2)
	if a != b || a.Mean <= 0 {
		t.Fatalf("%+v vs %+v", a, b)
	}
	if _, _, err := simRotation(in.setup, rotation{Type: "NotAnAPLType"}, 2); err == nil {
		t.Fatal("a rotation the engine's proto cannot parse ran")
	}
}

func TestTallyFinishersCountsTheRogueFinishersAndSkipsOtherSpecs(t *testing.T) {
	in, err := prepare(realOptions("rogue-combat"))
	if err != nil {
		t.Fatal(err)
	}
	names, err := buildSpellNames(realRoot, in.clientBuild, "rogue", 60)
	if err != nil {
		t.Fatal(err)
	}
	got, err := tallyFinishers(in, in.base, in.base.withPriorityList(nil), 2, names)
	if err != nil || got == nil {
		t.Fatalf("got %v %v", got, err)
	}
	if got.baseline[abilityEviscerate] <= 0 {
		t.Fatalf("the combat rotation casts Eviscerate: %v", got.baseline)
	}
	if got.best[abilityEviscerate] != 0 || len(got.best.notAdoptable()) != len(requiredFinishers) {
		t.Fatalf("an empty rotation casts no finishers: %v", got.best)
	}
	in.setup.spec.Spec = "mage-fire"
	if none, err := tallyFinishers(in, in.base, in.base, 2, names); none != nil || err != nil {
		t.Fatalf("a spec without a finisher list: %v %v", none, err)
	}
	if empty := tallyCasts(&proto.UnitMetrics{}, 2, names, []string{abilityRupture}); empty[abilityRupture] != 0 {
		t.Fatalf("no casts: %v", empty)
	}
}

func TestRunWritesAProbeOnlyReport(t *testing.T) {
	out := t.TempDir()
	args := []string{"-repo-root", realRoot, "-spec", "rogue-combat", "-iterations", "2", "-confirm", "2", "-probe-only", "-out", out}
	if err := run(args); err != nil {
		t.Fatal(err)
	}
	body, err := os.ReadFile(filepath.Join(out, "rogue-combat.md"))
	if err != nil {
		t.Fatal(err)
	}
	text := string(body)
	if !strings.Contains(text, "# Rotation search:") || !strings.Contains(text, "## Action probe") || strings.Contains(text, "## Rotation, baseline vs winner") {
		t.Fatalf("report:\n%s", text)
	}
	if err := run([]string{"-repo-root", realRoot, "-spec", "priest-holy", "-out", out}); !errors.Is(err, errSkipped) {
		t.Fatalf("a healer is skipped, got %v", err)
	}
}

func TestRunWritesTheFullSearchReport(t *testing.T) {
	out := t.TempDir()
	args := []string{"-repo-root", realRoot, "-spec", "rogue-combat", "-preset", "raid", "-iterations", "2", "-confirm", "2", "-rounds", "1", "-out", out}
	if err := run(args); err != nil {
		t.Fatal(err)
	}
	body, err := os.ReadFile(filepath.Join(out, "rogue-combat-raid.md"))
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"**Verdict:", "## Finisher casts", "## Rotation, baseline vs winner", "## Winning rotation, engine APL JSON"} {
		if !strings.Contains(string(body), want) {
			t.Errorf("full report missing %q", want)
		}
	}
}
