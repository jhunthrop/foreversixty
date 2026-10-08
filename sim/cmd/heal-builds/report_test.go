package main

import (
	"encoding/json"
	"io"
	"math"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

const realRoot = "../../.."

func writeFile(t *testing.T, path, body string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()
	orig := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w
	fn()
	os.Stdout = orig
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	return string(out)
}

func TestRunRequiresSpecBisAndBuilds(t *testing.T) {
	for name, args := range map[string][]string{
		"nothing":    {},
		"no builds":  {"-spec", "priest-holy", "-bis", "x"},
		"no bis":     {"-spec", "priest-holy", "-builds", "x"},
		"no spec":    {"-bis", "x", "-builds", "y"},
		"bad flag":   {"-nope"},
		"bad number": {"-iterations", "many"},
	} {
		if err := run(args); err == nil {
			t.Errorf("%s: accepted", name)
		}
	}
}

func TestReportRefusesNonHealers(t *testing.T) {
	for _, spec := range []string{"mage-fire", "warrior-protection", "no-such-spec"} {
		if err := report(options{spec: spec}); err == nil || !strings.Contains(err.Error(), "not a healing spec") {
			t.Errorf("%s: err = %v", spec, err)
		}
	}
}

func TestLoadBandWantsTheLevel60RaidBandOfTheFaction(t *testing.T) {
	path := filepath.Join(t.TempDir(), "bis.json")
	if _, err := loadBand(path, "alliance"); err == nil {
		t.Fatal("missing file accepted")
	}
	writeFile(t, path, `{"bands":[
	  {"preset":"raid","band":50,"faction":"alliance","race":"Dwarf"},
	  {"preset":"bare","band":60,"faction":"alliance","race":"Gnome"},
	  {"preset":"raid","band":60,"faction":"horde","race":"Orc"},
	  {"preset":"raid","band":60,"faction":"alliance","race":"Human","slots":[{"slot":"head","item_id":9}]}]}`)
	band, err := loadBand(path, "alliance")
	if err != nil || band.Race != "Human" || len(band.Slots) != 1 {
		t.Fatalf("got %+v %v", band, err)
	}
	if horde, _ := loadBand(path, "horde"); horde.Race != "Orc" {
		t.Fatalf("horde = %+v", horde)
	}
	if _, err := loadBand(path, "neutral"); err == nil || !strings.Contains(err.Error(), "no level-60 neutral raid band") {
		t.Fatalf("absent faction: %v", err)
	}
	writeFile(t, path, `{`)
	if _, err := loadBand(path, "alliance"); err == nil || !strings.Contains(err.Error(), "decoding") {
		t.Fatalf("bad json: %v", err)
	}
}

func TestLoadBuildsAndSortedKeys(t *testing.T) {
	path := filepath.Join(t.TempDir(), "builds.json")
	if _, err := loadBuilds(path); err == nil {
		t.Fatal("missing file accepted")
	}
	writeFile(t, path, `{"zeta":{"Alpha":5},"alpha":{"Beta":2}}`)
	builds, err := loadBuilds(path)
	if err != nil {
		t.Fatal(err)
	}
	if got := strings.Join(sortedKeys(builds), ","); got != "alpha,zeta" {
		t.Fatalf("sortedKeys = %s", got)
	}
	if builds["zeta"]["Alpha"] != 5 {
		t.Fatalf("builds = %v", builds)
	}
	writeFile(t, path, `{"x":{"Alpha":"five"}}`)
	if _, err := loadBuilds(path); err == nil {
		t.Fatal("non-numeric rank accepted")
	}
}

func TestBuildRequestWearsTheBandsNonEmptySlotsAndTheCasterDistance(t *testing.T) {
	spec := specs.ByKey["priest-holy"]
	band := bisBand{Race: "Gnome"}
	band.Slots = append(band.Slots,
		struct {
			Slot   string `json:"slot"`
			ItemID int    `json:"item_id"`
		}{"head", 5},
		struct {
			Slot   string `json:"slot"`
			ItemID int    `json:"item_id"`
		}{"neck", 0})
	req := buildRequest(options{iterations: 12}, spec, band, "ENGINE", []string{"b"}, []string{"c"})
	c := req.Character
	if len(c.Gear) != 1 || c.Gear[0] != (api.GearSlot{Slot: "head", ItemID: 5}) {
		t.Errorf("gear = %+v", c.Gear)
	}
	if c.Race != "Gnome" || c.Class != "priest" || c.Level != level || c.Talents != "ENGINE" {
		t.Errorf("character = %+v", c)
	}
	if c.DistanceFromTarget != leveling.CasterDistanceFromTarget || c.Buffs[0] != "b" || c.Consumes[0] != "c" {
		t.Errorf("loadout = %+v", c)
	}
	if req.Iterations != 12 || req.RandomSeed != seed || req.Spec != "priest-holy" {
		t.Errorf("request = %+v", req)
	}
}

func TestRaidLoadoutLayersThePresetOverTheClassKit(t *testing.T) {
	spec := specs.ByKey["priest-holy"]
	buffs, consumes, err := raidLoadout(realRoot, spec)
	if err != nil {
		t.Fatal(err)
	}
	kit := leveling.KitBuffs(spec.Spec, level)
	if len(buffs) < len(kit) || len(consumes) == 0 {
		t.Fatalf("buffs %v consumes %v", buffs, consumes)
	}
	if _, _, err := raidLoadout(t.TempDir(), spec); err == nil {
		t.Fatal("a repo without presets accepted")
	}
}

func TestActiveBuildDirIsUnderDataBuilds(t *testing.T) {
	dir, err := activeBuildDir(realRoot)
	if err != nil {
		t.Fatal(err)
	}
	build, _ := leveling.ReadActiveBuild(realRoot)
	if dir != filepath.Join(realRoot, "data", "builds", build) {
		t.Fatalf("dir = %s", dir)
	}
	if _, err := activeBuildDir(t.TempDir()); err == nil {
		t.Fatal("a repo with no active build accepted")
	}
}

func TestLegalRequiresExactly51PointsAndSiteStringWidth(t *testing.T) {
	ranks := map[int]int{1: 5, 2: 5, 3: 5, 4: 5}
	if err := legal(trees(), ranks); err == nil || !strings.Contains(err.Error(), "spends 20 points, want 51") {
		t.Fatalf("got %v", err)
	}
}

func TestRowScoreSquaresTheShareOfTheFightManaLasted(t *testing.T) {
	mk := func(manaLasts float64) row {
		r := row{fight: 100}
		r.result.Effective.Mean = 400
		r.result.ManaLastsSec = manaLasts
		return r
	}
	if got := mk(50).score(); math.Abs(got-100) > 1e-9 { // 400 * 0.5^2
		t.Errorf("half the fight: %v, want 100", got)
	}
	if got := mk(250).score(); got != 400 {
		t.Errorf("mana to spare must cap the share at 1: %v", got)
	}
}

func TestPrintRowsSortsByScoreAndFormatsTheColumns(t *testing.T) {
	low, high := row{label: "low", site: "s1", engine: "e1", fight: 100}, row{label: "high", site: "s2", engine: "e2", fight: 100}
	low.result.Effective.Mean, low.result.ManaLastsSec = 100, 100
	high.result.Effective.Mean, high.result.ManaLastsSec = 300, 100
	high.result.Raw.Mean = 400
	rows := []row{low, high}
	out := captureStdout(t, func() { printRows(rows) })
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 3 || !strings.HasPrefix(lines[0], "build") {
		t.Fatalf("output:\n%s", out)
	}
	if !strings.HasPrefix(lines[1], "high") || !strings.HasPrefix(lines[2], "low") {
		t.Fatalf("not sorted best first:\n%s", out)
	}
	if !strings.Contains(lines[1], "300.0") || !strings.Contains(lines[1], "25.0%") || !strings.Contains(lines[1], "s2") {
		t.Fatalf("columns wrong: %s", lines[1])
	}
}

// guideBuildByName reads the priest-holy guide's talents as {name: rank}, the
// shape -builds takes.
func guideBuildByName(t *testing.T) map[string]int {
	t.Helper()
	client, digits, err := leveling.GuideBuildTalents(realRoot, "priest", "holy")
	if err != nil {
		t.Fatal(err)
	}
	trees, err := leveling.LoadTalentTrees(realRoot, client, "priest")
	if err != nil {
		t.Fatal(err)
	}
	out := map[string]int{}
	for ti, tree := range trees {
		for i, node := range tree.Talents {
			if i < len(digits[ti]) && digits[ti][i] > '0' {
				out[node.Name] = int(digits[ti][i] - '0')
			}
		}
	}
	return out
}

func TestRunSimsEachBuildAndRejectsIllegalOnes(t *testing.T) {
	dir := t.TempDir()
	build, err := leveling.ReadActiveBuild(realRoot)
	if err != nil {
		t.Fatal(err)
	}
	bis := filepath.Join(leveling.BisDir(filepath.Join(realRoot, "data", "builds", build)), "priest-holy.json")
	buildsPath := filepath.Join(dir, "builds.json")
	guide := guideBuildByName(t)
	raw, _ := json.Marshal(map[string]map[string]int{"guide": guide})
	writeFile(t, buildsPath, string(raw))

	args := []string{"-repo-root", realRoot, "-spec", "priest-holy", "-bis", bis, "-builds", buildsPath, "-iterations", "3"}
	out := captureStdout(t, func() {
		if err := run(args); err != nil {
			t.Error(err)
		}
	})
	if !strings.Contains(out, "guide") || !strings.Contains(out, "build") {
		t.Fatalf("output:\n%s", out)
	}

	illegal := map[string]int{}
	for name := range guide {
		illegal[name] = 1
	}
	raw, _ = json.Marshal(map[string]map[string]int{"short": illegal})
	writeFile(t, buildsPath, string(raw))
	if err := run(args); err == nil || !strings.Contains(err.Error(), `build "short"`) {
		t.Fatalf("an underspent build must fail by label, got %v", err)
	}

	raw, _ = json.Marshal(map[string]map[string]int{"typo": {"No Such Talent": 5}})
	writeFile(t, buildsPath, string(raw))
	if err := run(args); err == nil || !strings.Contains(err.Error(), "no talent named") {
		t.Fatalf("unknown talent: %v", err)
	}
}
