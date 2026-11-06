package main

import (
	"encoding/json"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	simscore "github.com/jhunthrop/foreversixty/sim/score"
)

const raceSweepBuildDir = "../../../data/builds/1.60.1.70790"

// The legal races come from the client's race-class combinations, Forever's
// added ones (Skyborne, Human hunter, Troll mage, Gnome shaman) included.
func TestLegalRacesReadTheClientCombinationsPerFaction(t *testing.T) {
	cases := []struct {
		class, faction string
		want           []string
	}{
		{"warrior", "horde", []string{"orc", "undead", "tauren", "troll", "windshaper-skyborne"}},
		{"druid", "alliance", []string{"night-elf", "high-order-skyborne"}},
		{"hunter", "alliance", []string{"human", "dwarf", "night-elf", "high-order-skyborne"}},
		{"shaman", "alliance", []string{"dwarf"}},
		{"paladin", "horde", []string{"undead"}},
	}
	for _, c := range cases {
		got, err := legalRaces(raceSweepBuildDir, c.class, c.faction)
		if err != nil {
			t.Fatalf("legalRaces(%s, %s): %v", c.class, c.faction, err)
		}
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("legalRaces(%s, %s) = %v, want %v", c.class, c.faction, got, c.want)
		}
	}
}

func raceScoringEngine(scores map[string]float64, stdErr float64) *fakeEngine {
	return &fakeEngine{
		DPSFunc:       func(req api.SimRequest) (float64, error) { return scores[req.Character.Race], nil },
		DefaultStdErr: stdErr,
	}
}

// One spec, one band, low iterations: every legal race runs the same gear,
// talents, seed and count, and the best race, its margin over the guide race
// and the margin's error are published.
func TestSweepRacesNamesTheBestRaceAndItsMarginOverTheGuideRace(t *testing.T) {
	scores := map[string]float64{"orc": 1000, "undead": 980, "tauren": 990, "troll": 1030, "windshaper-skyborne": 1010}
	engine := raceScoringEngine(scores, 3)
	var seen []api.SimRequest
	inner := engine.DPSFunc
	engine.DPSFunc = func(req api.SimRequest) (float64, error) { seen = append(seen, req); return inner(req) }
	spec := specInfo{Spec: "warrior-fury", ClassSlug: "warrior", Role: "dps"}
	gear := []api.GearSlot{{Slot: "main_hand", ItemID: 19352}}
	races := []string{"orc", "undead", "tauren", "troll", "windshaper-skyborne"}

	got, err := sweepRaces(engine, spec, "orc", races, 60, "talents", gear, 300, 7)
	if err != nil {
		t.Fatalf("sweepRaces: %v", err)
	}
	if got.BestRace != "troll" || got.GuideRace != "orc" {
		t.Fatalf("best %q guide %q, want troll over orc", got.BestRace, got.GuideRace)
	}
	if got.Margin != 30 || math.Abs(got.MarginPct-3) > 1e-9 {
		t.Errorf("margin %v (%v%%), want 30 (3%%)", got.Margin, got.MarginPct)
	}
	if want := math.Sqrt(18); math.Abs(got.Error-want) > 1e-9 || !got.BeyondError {
		t.Errorf("error %v beyond %v, want %v and beyond", got.Error, got.BeyondError, want)
	}
	if got.ScoreUnit != raceScoreDPS || got.Iterations != 300 || got.Seed != 7 || len(got.Races) != len(races) {
		t.Errorf("unit %q iterations %d seed %d races %d", got.ScoreUnit, got.Iterations, got.Seed, len(got.Races))
	}
	if len(seen) != len(races) {
		t.Fatalf("%d runs, want one per race", len(seen))
	}
	for _, req := range seen {
		if req.Iterations != 300 || req.RandomSeed != 7 || req.Character.Talents != "talents" || gearKey(req.Character.Gear) != gearKey(gear) || req.Character.Level != 60 {
			t.Errorf("race %s ran %d iterations seed %d talents %q gear %s", req.Character.Race, req.Iterations, req.RandomSeed, req.Character.Talents, gearKey(req.Character.Gear))
		}
	}
}

// The guide race keeps the top line when nothing beats it, and a lead inside
// the combined error is not beyond error.
func TestSweepRacesKeepsTheGuideRaceWhenItIsBestAndFlagsALeadInsideError(t *testing.T) {
	spec := specInfo{Spec: "mage-fire", ClassSlug: "mage", Role: "dps"}
	got, err := sweepRaces(raceScoringEngine(map[string]float64{"gnome": 900, "human": 890}, 1), spec, "gnome", []string{"human", "gnome"}, 60, "", nil, 300, 1)
	if err != nil || got.BestRace != "gnome" || got.Margin != 0 || got.BeyondError {
		t.Fatalf("sweep = %+v, %v; want gnome kept, no margin", got, err)
	}
	got, err = sweepRaces(raceScoringEngine(map[string]float64{"gnome": 900, "human": 902}, 2), spec, "gnome", []string{"human", "gnome"}, 60, "", nil, 300, 1)
	if err != nil || got.BestRace != "human" || got.BeyondError {
		t.Fatalf("sweep = %+v, %v; want human ahead inside error", got, err)
	}
}

func TestSweepRacesRefusesAGuideRaceOutsideTheLegalList(t *testing.T) {
	spec := specInfo{Spec: "mage-fire", ClassSlug: "mage", Role: "dps"}
	if _, err := sweepRaces(raceScoringEngine(nil, 0), spec, "orc", []string{"human", "gnome"}, 60, "", nil, 300, 1); err == nil {
		t.Fatal("sweepRaces accepted a guide race the class cannot be in this faction")
	}
}

// A healer asks for the iteration quotient its runner scales back up, so the
// sweep lands at the count it publishes.
func TestSweepRacesAsksAHealerRunnerForTheScaledQuotient(t *testing.T) {
	var asked []int
	engine := &fakeEngine{DPSFunc: func(req api.SimRequest) (float64, error) { asked = append(asked, req.Iterations); return 1, nil }}
	spec := specInfo{Spec: "priest-holy", ClassSlug: "priest", Role: healerRole}
	got, err := sweepRaces(engine, spec, "human", []string{"human"}, 60, "", nil, 10000, 1)
	if err != nil {
		t.Fatal(err)
	}
	if want := (10000 + healIterationScale - 1) / healIterationScale; asked[0] != want || got.Iterations != want*healIterationScale || got.ScoreUnit != raceScoreHeal {
		t.Fatalf("asked %v, published %d unit %q; want %d asked", asked, got.Iterations, got.ScoreUnit, want)
	}
}

type fakeTankFigures struct {
	*fakeEngine
	tps map[string]float64
}

func (f fakeTankFigures) RunTankFigures(req api.SimRequest) (simscore.TankFigures, error) {
	return simscore.TankFigures{TPS: f.tps[req.Character.Race], TPSError: 2, Score: 1}, nil
}

// Tanks are judged on threat (owner rule), not on the tank score.
func TestSweepRacesJudgesATankOnThreat(t *testing.T) {
	engine := fakeTankFigures{fakeEngine: &fakeEngine{DefaultDPS: 1}, tps: map[string]float64{"human": 500, "dwarf": 520}}
	spec := specInfo{Spec: "warrior-protection", ClassSlug: "warrior", Role: roleTank}
	got, err := sweepRaces(newMemoRunner(engine), spec, "human", []string{"human", "dwarf"}, 60, "", nil, 300, 1)
	if err != nil {
		t.Fatal(err)
	}
	if got.BestRace != "dwarf" || got.Margin != 20 || got.ScoreUnit != raceScoreThreat {
		t.Fatalf("sweep = %+v; want dwarf ahead by 20 threat", got)
	}
}

// The standalone mode sweeps a frozen published report without ranking gear:
// one band entry, its published gear and talents, at the asked count and seed.
func TestRunRaceSweepSweepsTheFrozenPublishedSet(t *testing.T) {
	bisDir, outDir := t.TempDir(), t.TempDir()
	frozen := map[string]any{"spec": "warrior-fury", "bands": []map[string]any{{
		"band": 60, "preset": "bare", "faction": "horde", "race": "orc",
		"talents": "", "slots": []map[string]any{{"slot": "main_hand", "item_id": 19352}, {"slot": "off_hand", "item_id": 0}},
	}}}
	raw, _ := json.Marshal(frozen)
	if err := os.WriteFile(filepath.Join(bisDir, "warrior-fury.json"), raw, 0o644); err != nil {
		t.Fatal(err)
	}
	scores := map[string]float64{"orc": 1000, "undead": 990, "tauren": 995, "troll": 1020, "windshaper-skyborne": 1001}
	engine := raceScoringEngine(scores, 1)
	err := runRaceSweep(engine, raceSweepConfig{
		RepoRoot: integrationRepoRoot, BuildDir: raceSweepBuildDir, BISDir: bisDir, OutDir: outDir,
		Spec: "warrior-fury", Iterations: 500, Seed: 50001,
	}, identityTalentLayout)
	if err != nil {
		t.Fatalf("runRaceSweep: %v", err)
	}
	b, err := os.ReadFile(filepath.Join(outDir, "warrior-fury.races.json"))
	if err != nil {
		t.Fatal(err)
	}
	var out []bandRaceSweep
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatal(err)
	}
	if len(out) != 1 || out[0].Band != 60 || out[0].Faction != "horde" || out[0].Sweep.BestRace != "troll" || out[0].Sweep.Seed != 50001 {
		t.Fatalf("out = %+v", out)
	}
	// The guide's current race line is what the sweep judges (fury.md:
	// undead), whatever race the frozen report was ranked on.
	if out[0].Sweep.GuideRace != "undead" || out[0].PublishedRace != "orc" || out[0].Sweep.Margin != 30 {
		t.Fatalf("guide %q published %q margin %v", out[0].Sweep.GuideRace, out[0].PublishedRace, out[0].Sweep.Margin)
	}
	for _, g := range engine.Calls {
		if g != gearKey([]api.GearSlot{{Slot: "main_hand", ItemID: 19352}}) {
			t.Errorf("ran gear %s, want the frozen main hand only", g)
		}
	}
}
