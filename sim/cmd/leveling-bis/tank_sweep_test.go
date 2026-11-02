package main

import (
	"errors"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// tankSweepRequestForTest is the band-20 weights request runSpec builds for
// Protection warrior: the guide's race, the band's talents and the tank
// ladder character, on the active build's real items.
func tankSweepRequestForTest(t *testing.T, iterations int) (specInfo, api.SimRequest) {
	t.Helper()
	const band = 20
	si, err := loadSpec(publishedRepoRoot, "warrior-protection")
	if err != nil {
		t.Fatal(err)
	}
	build, err := leveling.ReadActiveBuild(publishedRepoRoot)
	if err != nil {
		t.Fatal(err)
	}
	buildDir := filepath.Join(publishedRepoRoot, "data", "builds", build)
	targets, err := loadPassTargets(publishedRepoRoot, build, si)
	if err != nil {
		t.Fatal(err)
	}
	trees, err := leveling.LoadTalentTrees(publishedRepoRoot, build, si.ClassSlug)
	if err != nil {
		t.Fatal(err)
	}
	layout, err := resolveTalentLayout(publishedRepoRoot, si.ClassSlug)
	if err != nil {
		t.Fatal(err)
	}
	items, _, _, err := loadCandidates(buildDir, si.ClassSlug)
	if err != nil {
		t.Fatal(err)
	}
	factors, err := loadRatingFactors(buildDir)
	if err != nil {
		t.Fatal(err)
	}
	items = withEngineWeaponDamage(markNotInSimDB(convertCandidateRatings(items, factors)), simdb.Weapon)
	guide, err := loadGuideRaces(publishedRepoRoot, buildDir, si.ClassSlug, si.SpecSlug)
	if err != nil {
		t.Fatal(err)
	}
	_, engineTalents, err := bandTalentStrings(trees, targets.forPass("bare"), si.TreeIndex, band, layout)
	if err != nil {
		t.Fatal(err)
	}
	ch := ladderCharacter(guide.AllianceRace, si.ClassSlug, band, engineTalents, ladderWeapon(items, band), ladderMeleeWeapons(items, band, si.Spec)...)
	ch = tankLadderCharacter(si, items, band, ch)
	return si, weightsRequest(si, ch, iterations, weightsSeed)
}

// sweepFigures flattens a sweep's weights, errors and per-point scale into
// one comparable map.
func sweepFigures(runner engineRunner, req api.SimRequest) (map[string]float64, error) {
	w, perPoint, err := runner.RunWeights(req)
	if err != nil {
		return nil, err
	}
	out := map[string]float64{"per_point": perPoint}
	for stat, row := range w {
		out[stat] = row.Weight
		out[stat+"_error"] = row.Error
	}
	return out, nil
}

// Round 9: the tank sweep is reproducible, and it values stamina.
//
// Every sweep point shares one seed with labelled random streams (common
// random numbers across the stat deltas), so two runs give the same weights
// to the bit, and the picks they rank are the same.
//
// Round 8's open note had several tank sweeps "fail outright: the stamina
// reference weighs nothing", and two seeds disagreeing in up to 15 armor
// slots. Its probe swept the tanks on the damage engine, where stamina is
// worth no damage. The ranker sweeps them on the tank engine (roleRunner),
// where stamina's health is worth score, and the damage engine refuses.
func TestTheTankSweepIsReproducibleAndValuesStamina(t *testing.T) {
	if testing.Short() {
		t.Skip("runs the engine")
	}
	si, req := tankSweepRequestForTest(t, 20)
	tank := roleRunner(realEngine{}, si)
	first, err := sweepFigures(tank, req)
	if err != nil {
		t.Fatalf("first sweep: %v", err)
	}
	second, err := sweepFigures(tank, req)
	if err != nil {
		t.Fatalf("second sweep: %v", err)
	}
	if !reflect.DeepEqual(first, second) {
		t.Errorf("two tank sweeps differ:\n%v\n%v", first, second)
	}
	if first["per_point"] <= 0 || first["stamina"] != 1 || first["stamina_error"] >= 1 {
		t.Errorf("stamina = %v ± %v at %v score per point, want the reference positive beyond its error", first["stamina"], first["stamina_error"], first["per_point"])
	}
	if _, err := sweepFigures(realEngine{}, req); !errors.Is(err, adapter.ErrNoWeights) {
		t.Errorf("the damage sweep of a tank = %v, want ErrNoWeights (stamina deals no damage)", err)
	}
}
