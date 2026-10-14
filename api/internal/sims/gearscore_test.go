package sims

import (
	"encoding/json"
	"errors"
	"math"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// repoBuilds is the repository's own data/builds, the directory the API's
// image ships as TreeDataDir.
const repoBuilds = "../../../data/builds"

// gearScoreFixture is docs/contracts/fixtures/gearscore.json, the number the
// addon lane pins Gear.score against too.
type gearScoreFixture struct {
	Build    string             `json:"build"`
	Spec     string             `json:"spec"`
	Weights  map[string]float64 `json:"weights"`
	GearCode string             `json:"gear_code"`
	Items    []struct {
		Slot   string         `json:"slot"`
		ItemID int            `json:"item_id"`
		Stats  map[string]int `json:"stats"`
		Score  float64        `json:"score"`
	} `json:"items"`
	Expected float64 `json:"expected_score"`
}

func readGearScoreFixture(t *testing.T) gearScoreFixture {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("..", "..", "..", "docs", "contracts", "fixtures", "gearscore.json"))
	if err != nil {
		t.Fatal(err)
	}
	var f gearScoreFixture
	if err := json.Unmarshal(b, &f); err != nil {
		t.Fatal(err)
	}
	return f
}

var (
	repoTreesOnce sync.Once
	repoTrees     *trees.Data
	repoTreesErr  error
)

// useRepoGearData points GearScore at the repository's real data/builds, once
// per test binary: loading every build's item tables takes a few seconds.
func useRepoGearData(t *testing.T) {
	t.Helper()
	repoTreesOnce.Do(func() { repoTrees, repoTreesErr = trees.Load(repoBuilds) })
	if repoTreesErr != nil {
		t.Fatal(repoTreesErr)
	}
	SetGearData(repoTrees, repoBuilds)
}

const scoreTolerance = 1e-9

// Gear.score (addon/ForeverSixty/Gear.lua) is a linear sum of each stat
// times the spec's weight; the fixture's per-item lines are that sum done
// by hand, so the fixture must agree with itself before it can pin anything.
func TestTheGearScoreFixtureAddsUp(t *testing.T) {
	f := readGearScoreFixture(t)
	total := 0.0
	for _, it := range f.Items {
		score := 0.0
		for stat, value := range it.Stats {
			score += float64(value) * f.Weights[stat]
		}
		if math.Abs(score-it.Score) > scoreTolerance {
			t.Errorf("%s %d scores %v by Gear.score's rule, the fixture says %v", it.Slot, it.ItemID, score, it.Score)
		}
		total += score
	}
	if math.Abs(total-f.Expected) > scoreTolerance {
		t.Fatalf("the items sum to %v, expected_score is %v", total, f.Expected)
	}
}

// The fixture's weights and stat lines are copies; they must be the ones the
// API itself reads, or the number pins nothing.
func TestTheGearScoreFixtureMatchesTheShippedData(t *testing.T) {
	f := readGearScoreFixture(t)
	weights, err := loadSpecWeights(repoBuilds, f.Build)
	if err != nil {
		t.Fatal(err)
	}
	got := weights[f.Spec]
	if len(got) != len(f.Weights) {
		t.Fatalf("shipped %s weights = %v, fixture = %v", f.Spec, got, f.Weights)
	}
	for stat, w := range f.Weights {
		if got[stat] != w {
			t.Errorf("shipped %s weight for %s = %v, fixture = %v", f.Spec, stat, got[stat], w)
		}
	}
	useRepoGearData(t)
	b, ok := repoTrees.Build(f.Build)
	if !ok {
		t.Fatalf("build %s not loaded", f.Build)
	}
	warrior, ok := classIDBySlug(b, "warrior")
	if !ok {
		t.Fatal("no warrior class")
	}
	for _, it := range f.Items {
		item, ok := b.Item(warrior, it.ItemID)
		if !ok {
			t.Errorf("item %d is not in the build's warrior items", it.ItemID)
			continue
		}
		if len(item.Stats) != len(it.Stats) {
			t.Errorf("item %d stats = %v, fixture = %v", it.ItemID, item.Stats, it.Stats)
			continue
		}
		for stat, v := range it.Stats {
			if item.Stats[stat] != v {
				t.Errorf("item %d %s = %d, fixture = %d", it.ItemID, stat, item.Stats[stat], v)
			}
		}
	}
}

func TestGearScoreOfTheBand20ArmsSetIsTheFixturesNumber(t *testing.T) {
	f := readGearScoreFixture(t)
	useRepoGearData(t)
	got, err := GearScore(t.Context(), f.GearCode, f.Spec)
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(got-f.Expected) > scoreTolerance {
		t.Fatalf("GearScore = %v, want %v", got, f.Expected)
	}
}

func TestGearScoreCountsAnItemTheDataDoesNotHaveAsNothing(t *testing.T) {
	useRepoGearData(t)
	// 4706 alone scores 8 (strength 4 * 2); 999999999 is no item at all.
	got, err := GearScore(t.Context(), "FS1:1.60.1.70291:warrior:human:0/0/0:back=4706,chest=999999999", "warrior-arms")
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(got-8) > scoreTolerance {
		t.Fatalf("GearScore = %v, want 8", got)
	}
}

func TestGearScoreRefusesWhatItCannotScore(t *testing.T) {
	useRepoGearData(t)
	for name, tc := range map[string]struct{ code, spec string }{
		"not an FS1 code": {"hello", "warrior-arms"},
		"unknown build":   {"FS1:0.0.0.1:warrior:human:0/0/0:back=4706", "warrior-arms"},
		"unknown class":   {"FS1:1.60.1.70291:bard:human:0/0/0:back=4706", "warrior-arms"},
		"unknown spec":    {"FS1:1.60.1.70291:warrior:human:0/0/0:back=4706", "warrior-bard"},
	} {
		if _, err := GearScore(t.Context(), tc.code, tc.spec); !errors.Is(err, ErrGearUnscorable) {
			t.Errorf("%s: err = %v, want ErrGearUnscorable", name, err)
		}
	}
}

func TestGearScoreWithNoDataWiredSaysSo(t *testing.T) {
	SetGearData(nil, "")
	t.Cleanup(func() { SetGearData(nil, "") })
	if _, err := GearScore(t.Context(), readGearScoreFixture(t).GearCode, "warrior-arms"); !errors.Is(err, ErrGearUnscorable) {
		t.Fatalf("err = %v, want ErrGearUnscorable", err)
	}
}

// The full-gear mark is the spec's own best band-60 set scored with the same build's
// weights: 4767 for warrior-arms on 1.60.1.70291 since the data lane's b388b399 rewrote the
// band-60 picks (4838, the hand-derived table's figure, before it).
func TestFullGearScoreIsTheSpecsBestBand60Set(t *testing.T) {
	useRepoGearData(t)
	got, err := FullGearScore(t.Context(), "1.60.1.70291", "warrior-arms", FullGearBand)
	if err != nil {
		t.Fatal(err)
	}
	if math.Round(got) != 4767 {
		t.Fatalf("FullGearScore(warrior-arms) = %v, want 4767 once rounded", got)
	}
	again, err := FullGearScore(t.Context(), "1.60.1.70291", "warrior-arms", FullGearBand)
	if err != nil || again != got {
		t.Fatalf("the cached mark = %v, %v; want %v", again, err, got)
	}
}

func TestFullGearScoreOfASpecWithNoBestInSlotSetIsAnError(t *testing.T) {
	useRepoGearData(t)
	// The same items and weights, in a data directory with no bis/ at all: the spec can
	// be weighted but has no set to be measured against.
	const build = "1.60.1.70291"
	dir := t.TempDir()
	weights, err := os.ReadFile(filepath.Join(repoBuilds, build, weightsFile))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(dir, build), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, build, weightsFile), weights, 0o644); err != nil {
		t.Fatal(err)
	}
	SetGearData(repoTrees, dir)
	t.Cleanup(func() { SetGearData(repoTrees, repoBuilds) })
	if _, err := FullGearScore(t.Context(), build, "warrior-arms", FullGearBand); !errors.Is(err, ErrNoBestInSlot) {
		t.Errorf("no best-in-slot set: err = %v, want ErrNoBestInSlot", err)
	}
	if _, err := FullGearScore(t.Context(), build, "warrior-gladiator", FullGearBand); !errors.Is(err, ErrGearUnscorable) {
		t.Errorf("a spec with no weights: err = %v, want ErrGearUnscorable", err)
	}
	if _, err := FullGearScore(t.Context(), "0.0.0.1", "warrior-arms", FullGearBand); !errors.Is(err, ErrGearUnscorable) {
		t.Errorf("an unknown build: err = %v, want ErrGearUnscorable", err)
	}
}

// A character's band is Gear.bandFor's: the highest band at or below its level, rounded down
// to the ten-level step and capped at 60; below 20 there is none.
func TestGearBandIsTheHighestBandAtOrBelowTheLevel(t *testing.T) {
	useRepoGearData(t)
	for level, want := range map[int]int{20: 20, 24: 20, 30: 30, 39: 30, 59: 50, 60: 60} {
		got, ok, err := GearBand(t.Context(), "1.60.1.70291", "warrior-arms", level)
		if err != nil || !ok || got != want {
			t.Errorf("GearBand(level %d) = %d, %v, %v; want %d", level, got, ok, err, want)
		}
	}
	for _, level := range []int{1, 19} {
		if got, ok, err := GearBand(t.Context(), "1.60.1.70291", "warrior-arms", level); err != nil || ok {
			t.Errorf("GearBand(level %d) = %d, %v, %v; want no band below 20", level, got, ok, err)
		}
	}
}

// A spec whose bis file skips a band is measured in the next band down, as the addon does.
func TestGearBandStepsDownPastABandTheSpecHasNoSetFor(t *testing.T) {
	useRepoGearData(t)
	const build = "1.60.1.70291"
	dir := t.TempDir()
	weights, err := os.ReadFile(filepath.Join(repoBuilds, build, weightsFile))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(dir, build, "bis"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, build, weightsFile), weights, 0o644); err != nil {
		t.Fatal(err)
	}
	bis := `{"bands": [{"band": 20, "slots": [{"item_id": 6627}]}, {"band": 40, "slots": [{"item_id": 6627}]}]}`
	if err := os.WriteFile(filepath.Join(dir, build, "bis", "warrior-arms.json"), []byte(bis), 0o644); err != nil {
		t.Fatal(err)
	}
	SetGearData(repoTrees, dir)
	t.Cleanup(func() { SetGearData(repoTrees, repoBuilds) })
	for level, want := range map[int]int{35: 20, 45: 40, 60: 40} {
		if got, ok, err := GearBand(t.Context(), build, "warrior-arms", level); err != nil || !ok || got != want {
			t.Errorf("GearBand(level %d) = %d, %v, %v; want %d", level, got, ok, err, want)
		}
	}
	if _, err := FullGearScore(t.Context(), build, "warrior-arms", 30); !errors.Is(err, ErrNoBestInSlot) {
		t.Errorf("a band with no set: err = %v, want ErrNoBestInSlot", err)
	}
}

// Each band's mark is that band's own best set: the fixture's set is warrior-arms' band-20
// alliance pick, so the band-20 mark is at least its score, and each higher band's mark is
// higher still.
func TestFullGearScoreOfABandIsThatBandsBestSet(t *testing.T) {
	useRepoGearData(t)
	f := readGearScoreFixture(t)
	previous := 0.0
	for _, band := range []int{20, 30, 40, 50, 60} {
		got, err := FullGearScore(t.Context(), f.Build, f.Spec, band)
		if err != nil {
			t.Fatalf("band %d: %v", band, err)
		}
		if band == 20 && got < f.Expected {
			t.Errorf("the band-20 mark %v is below the band-20 alliance set's own %v", got, f.Expected)
		}
		if got <= previous {
			t.Errorf("the band-%d mark %v is not above the band below's %v", band, got, previous)
		}
		previous = got
	}
}
