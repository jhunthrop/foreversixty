package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

func TestFormatWeights(t *testing.T) {
	got := formatWeights([]string{"agility", "crit"}, map[string]float64{"agility": 1.5, "crit": 0.75})
	want := "agility=1.500, crit=0.750"
	if got != want {
		t.Errorf("formatWeights = %q, want %q", got, want)
	}
}

func TestFormatWeightsMissingStatDefaultsToZero(t *testing.T) {
	got := formatWeights([]string{"agility", "hit"}, map[string]float64{"agility": 1})
	want := "agility=1.000, hit=0.000"
	if got != want {
		t.Errorf("formatWeights = %q, want %q", got, want)
	}
}

func TestTalentPointsSpent(t *testing.T) {
	cases := map[string]int{
		"0500000":  5,
		"5-3-0":    8,
		"":         0,
		"000":      0,
		"10-20-30": 6, // digits summed independently: 1+0+2+0+3+0
	}
	for talents, want := range cases {
		if got := talentPointsSpent(talents); got != want {
			t.Errorf("talentPointsSpent(%q) = %d, want %d", talents, got, want)
		}
	}
}

func TestParseBands(t *testing.T) {
	got, err := parseBands("10,20, 30 ,40")
	if err != nil {
		t.Fatalf("parseBands: %v", err)
	}
	want := []int{10, 20, 30, 40}
	if len(got) != len(want) {
		t.Fatalf("parseBands = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("parseBands = %v, want %v", got, want)
		}
	}
}

func TestParseBandsInvalidNumber(t *testing.T) {
	if _, err := parseBands("10,abc,30"); err == nil {
		t.Fatal("parseBands(10,abc,30): want an error, got nil")
	}
}

func TestParseBandsDefaultFlag(t *testing.T) {
	got, err := parseBands(defaultBandsFlag)
	if err != nil {
		t.Fatalf("parseBands(defaultBandsFlag): %v", err)
	}
	if len(got) != 11 || got[0] != 10 || got[len(got)-1] != 60 {
		t.Fatalf("parseBands(defaultBandsFlag) = %v, want 11 bands from 10 to 60", got)
	}
}

func TestReadActiveBuild(t *testing.T) {
	build, err := readActiveBuild(repoRootFixture)
	if err != nil {
		t.Fatalf("readActiveBuild: %v", err)
	}
	if build != "testbuild" {
		t.Errorf("build = %q, want testbuild", build)
	}
}

func TestReadActiveBuildMissingFile(t *testing.T) {
	if _, err := readActiveBuild(t.TempDir()); err == nil {
		t.Fatal("readActiveBuild on an empty dir: want an error, got nil")
	}
}

func TestReadActiveBuildEmptyBuildField(t *testing.T) {
	dir := t.TempDir()
	if err := writeFile(t, filepath.Join(dir, "web", "src", "data", "active-build.json"), `{"build":""}`); err != nil {
		t.Fatal(err)
	}
	if _, err := readActiveBuild(dir); err == nil {
		t.Fatal("readActiveBuild with an empty build field: want an error, got nil")
	}
}

func TestReadActiveBuildMalformedJSON(t *testing.T) {
	dir := t.TempDir()
	if err := writeFile(t, filepath.Join(dir, "web", "src", "data", "active-build.json"), `{not json`); err != nil {
		t.Fatal(err)
	}
	if _, err := readActiveBuild(dir); err == nil {
		t.Fatal("readActiveBuild with malformed JSON: want an error, got nil")
	}
}

func TestWriteHeapProfileWritesAFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "heap.prof")
	writeHeapProfile(path)
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat %s: %v", path, err)
	}
	if info.Size() == 0 {
		t.Error("heap profile file is empty")
	}
}

func TestWriteHeapProfileBadPathLogsRatherThanPanics(t *testing.T) {
	// A path under a directory that cannot exist (a file, not a
	// directory, in the middle of it) makes os.Create fail - this must
	// be logged, not fatal (writeHeapProfile's own doc: "a failure to
	// write one is logged, not fatal"). The test's assertion is simply
	// that this does not panic.
	dir := t.TempDir()
	blocker := filepath.Join(dir, "not-a-directory")
	if err := os.WriteFile(blocker, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	writeHeapProfile(filepath.Join(blocker, "heap.prof"))
}

func TestRunSpecWritesReportsForEveryBandAndFaction(t *testing.T) {
	outDir := t.TempDir()
	fake := &fakeEngine{
		DefaultDPS: 500,
		WeightsResult: map[string]api.StatWeight{
			"ranged_attack_power": {Stat: "ranged_attack_power", Weight: 1.0},
			"agility":             {Stat: "agility", Weight: 1.8},
		},
	}
	buildDir := buildDirFixture()
	err := runSpec(fake, repoRootFixture, buildDir, "testbuild", outDir, "hunter-marksmanship", []int{20, 30}, 5)
	if err != nil {
		t.Fatalf("runSpec: %v", err)
	}
	jsonPath := filepath.Join(outDir, "hunter-marksmanship.json")
	if _, err := os.Stat(jsonPath); err != nil {
		t.Fatalf("stat %s: %v", jsonPath, err)
	}
	mdPath := filepath.Join(outDir, "hunter-marksmanship.md")
	b, err := os.ReadFile(mdPath)
	if err != nil {
		t.Fatalf("reading %s: %v", mdPath, err)
	}
	content := string(b)
	if !strings.Contains(content, "## Alliance") || !strings.Contains(content, "## Horde") {
		t.Errorf("markdown missing both faction sections:\n%s", content)
	}
	if !strings.Contains(content, "Band 20") || !strings.Contains(content, "Band 30") {
		t.Errorf("markdown missing both bands:\n%s", content)
	}
}

func TestRunSpecUnknownSpecErrors(t *testing.T) {
	fake := &fakeEngine{}
	err := runSpec(fake, repoRootFixture, buildDirFixture(), "testbuild", t.TempDir(), "no-such-spec", []int{20}, 5)
	if err == nil {
		t.Fatal("runSpec(no-such-spec): want an error, got nil")
	}
}

func TestRunSpecWeightsFailurePropagates(t *testing.T) {
	fake := &fakeEngine{FailWeights: true}
	err := runSpec(fake, repoRootFixture, buildDirFixture(), "testbuild", t.TempDir(), "hunter-marksmanship", []int{20}, 5)
	if err == nil {
		t.Fatal("runSpec with a failing weights run: want an error, got nil")
	}
}
