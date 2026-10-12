package main

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/specs"
	"github.com/wowsims/classic/sim/core/proto"
)

func TestTanksAreNotMeasuredAndHealersReportHealing(t *testing.T) {
	cases := map[string]struct {
		measured, healing bool
	}{
		"warrior-arms":        {true, false},
		"hunter-marksmanship": {true, false},
		"priest-holy":         {true, true},
		"warrior-protection":  {false, false},
		"druid-feral-bear":    {false, false},
	}
	for key, want := range cases {
		spec := specs.ByKey[key]
		if got := measured(spec); got != want.measured {
			t.Errorf("measured(%s) = %v, want %v", key, got, want.measured)
		}
		if got := healing(spec); got != want.healing {
			t.Errorf("healing(%s) = %v, want %v", key, got, want.healing)
		}
	}
}

func TestTheWeightsRequestWeighsOnlyTheReferenceStatAtSixty(t *testing.T) {
	spec := specs.ByKey["warrior-arms"]
	ch := api.CharacterSpec{Name: "x", Race: "human", Class: "warrior", Level: 60, Talents: "t"}
	req := weightsRequest(spec, ch, nil, 250, 7)
	if req.Weights == nil || !reflect.DeepEqual(req.Weights.Stats, []string{"attack_power"}) || req.Weights.Reference != "attack_power" {
		t.Fatalf("weights = %+v, want only attack_power as its own reference", req.Weights)
	}
	if req.Character.Level != 60 || req.Iterations != 250 || req.RandomSeed != 7 || req.Spec != "warrior-arms" {
		t.Fatalf("request = %+v", req)
	}
	if req.Encounter.TargetType != api.TargetTypeUnknown {
		t.Fatalf("target type %q: a raid boss grants no slaying bonus", req.Encounter.TargetType)
	}
	holy := weightsRequest(specs.ByKey["priest-holy"], api.CharacterSpec{Class: "priest", Level: 60}, nil, 250, 7)
	if holy.Iterations != 250*healIterationScale {
		t.Fatalf("a healer sweeps %d iterations, want %d: its healing figure is far noisier than a damage one", holy.Iterations, 250*healIterationScale)
	}
	mage := weightsRequest(specs.ByKey["mage-fire"], api.CharacterSpec{Class: "mage", Level: 60}, nil, 250, 7)
	if len(mage.Character.Buffs) == 0 {
		t.Fatal("the mage's class kit (Arcane Brilliance) is missing")
	}
	if mage.Character.DistanceFromTarget == 0 {
		t.Fatal("a caster stands out of melee range")
	}
}

// fakeMeasurer returns canned figures by spec: the reference stat's per
// point and its error, then the weapon term and its error.
type fakeMeasurer map[string][4]float64

func (f fakeMeasurer) Measure(spec specs.Spec, _ api.SimRequest) (measurement, error) {
	v, ok := f[spec.Spec]
	if !ok {
		return measurement{}, errors.New("no figure")
	}
	return measurement{PerPoint: v[0], StdErr: v[1], Weapon: v[2], WeaponErr: v[3]}, nil
}

func fakeCharacter(spec specs.Spec) (api.CharacterSpec, error) {
	return api.CharacterSpec{Name: "x", Race: "human", Class: spec.ClassSlug, Level: 60}, nil
}

func TestMeasureKeepsOnlyFiguresPositiveBeyondTheirError(t *testing.T) {
	m := fakeMeasurer{
		"warrior-arms": {0.4213, 0.01, 5.9, 0.1},
		"priest-holy":  {1.25, 0.05, 0, 0},
		"mage-fire":    {0.02, 0.05, 0, 0}, // inside its own error: noise, not a number
	}
	want := []string{"warrior-arms", "priest-holy", "mage-fire", "warrior-protection"}
	var picked []specs.Spec
	for _, key := range want {
		picked = append(picked, specs.ByKey[key])
	}
	got, skipped := measureAll(picked, fakeCharacter, nil, m, 100, 7)
	if len(got) != 2 {
		t.Fatalf("got %v, want warrior-arms and priest-holy only", got)
	}
	if !reflect.DeepEqual(got["warrior-arms"], entry{Reference: "attack_power", DPS: 0.4213, Healing: false, WeaponDPS: 5.9}) {
		t.Fatalf("warrior-arms = %+v", got["warrior-arms"])
	}
	if !reflect.DeepEqual(got["priest-holy"], entry{Reference: "healing_power", DPS: 1.25, Healing: true}) {
		t.Fatalf("priest-holy = %+v", got["priest-holy"])
	}
	if len(skipped) != 2 || !strings.Contains(strings.Join(skipped, "\n"), "mage-fire") || !strings.Contains(strings.Join(skipped, "\n"), "warrior-protection") {
		t.Fatalf("skipped = %v", skipped)
	}
}

func TestWriteMergesIntoTheExistingFileInSpecOrder(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "dps-per-point.json")
	if err := os.WriteFile(path, []byte(`{"warrior-arms": {"reference": "attack_power", "dps": 0.1, "healing": false}, "mage-fire": {"reference": "spell_power", "dps": 0.9, "healing": false}}`), 0o644); err != nil {
		t.Fatal(err)
	}
	err := writeMerged(path, map[string]entry{
		"warrior-arms": {Reference: "attack_power", DPS: 0.42134567, Healing: false},
		"priest-holy":  {Reference: "healing_power", DPS: 1.25, Healing: true},
	})
	if err != nil {
		t.Fatal(err)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	want := "{\n" +
		`  "mage-fire": {"reference":"spell_power","dps":0.9,"healing":false,"weapon_dps":0},` + "\n" +
		`  "priest-holy": {"reference":"healing_power","dps":1.25,"healing":true,"weapon_dps":0},` + "\n" +
		`  "warrior-arms": {"reference":"attack_power","dps":0.4213,"healing":false,"weapon_dps":0}` + "\n" +
		"}\n"
	if string(b) != want {
		t.Fatalf("file =\n%s\nwant\n%s", b, want)
	}
}

// --- Task L1.5: per band, with a measured weapon term ---

func TestABandIsMeasuredInTheMiddleOfTheLevelsTheAddonReadsItAt(t *testing.T) {
	// Gear.bandFor rounds a level down to its band: band 20 serves 20-29.
	for band, want := range map[int]int{20: 25, 30: 35, 40: 45, 50: 55, 60: 60} {
		if got := bandLevel(band); got != want {
			t.Errorf("bandLevel(%d) = %d, want %d", band, got, want)
		}
	}
}

func TestTheWeaponTermIsTheWeaponTheReferenceStatSwings(t *testing.T) {
	cases := map[string]proto.PseudoStat{
		"warrior-arms":        proto.PseudoStat_PseudoStatMainHandDps,
		"rogue-combat":        proto.PseudoStat_PseudoStatMainHandDps,
		"hunter-marksmanship": proto.PseudoStat_PseudoStatRangedDps,
		"hunter-survival":     proto.PseudoStat_PseudoStatMainHandDps,
		"mage-frost":          proto.PseudoStat_PseudoStatMainHandDps,
	}
	for key, want := range cases {
		if got := weaponPseudoStat(specs.ByKey[key]); got != want {
			t.Errorf("weaponPseudoStat(%s) = %v, want %v", key, got, want)
		}
	}
}

func TestMeasureBandsKeepsBothTermsAndZeroesAWeaponTermInsideItsError(t *testing.T) {
	m := fakeMeasurer{
		"warrior-arms": {0.05, 0.001, 0.7, 0.02},
		"mage-frost":   {0.11, 0.002, 0, 0},
		"rogue-combat": {0.06, 0.001, 0.01, 0.05}, // a weapon term inside its error is noise
	}
	var picked []specs.Spec
	for _, key := range []string{"warrior-arms", "mage-frost", "rogue-combat", "warrior-protection"} {
		picked = append(picked, specs.ByKey[key])
	}
	var levels []int
	characterFor := func(spec specs.Spec, band int) (api.CharacterSpec, error) {
		levels = append(levels, bandLevel(band))
		return api.CharacterSpec{Name: "x", Race: "human", Class: spec.ClassSlug, Level: bandLevel(band)}, nil
	}
	got, skipped := measureBands(picked, []int{20, 60}, characterFor, m, 100, 7)
	want := map[string]map[string]bandEntry{
		"warrior-arms": {
			"20": {Level: 25, DPS: 0.05, DPSError: 0.001, WeaponDPS: 0.7, WeaponError: 0.02},
			"60": {Level: 60, DPS: 0.05, DPSError: 0.001, WeaponDPS: 0.7, WeaponError: 0.02},
		},
		"mage-frost": {
			"20": {Level: 25, DPS: 0.11, DPSError: 0.002},
			"60": {Level: 60, DPS: 0.11, DPSError: 0.002},
		},
		"rogue-combat": {
			"20": {Level: 25, DPS: 0.06, DPSError: 0.001, WeaponError: 0.05},
			"60": {Level: 60, DPS: 0.06, DPSError: 0.001, WeaponError: 0.05},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("bands =\n%+v\nwant\n%+v", got, want)
	}
	if !reflect.DeepEqual(levels, []int{25, 60, 25, 60, 25, 60}) {
		t.Fatalf("characters built at levels %v", levels)
	}
	if len(skipped) != 1 || !strings.Contains(skipped[0], "warrior-protection") {
		t.Fatalf("skipped = %v", skipped)
	}
}

func TestWriteKeepsTheWeaponTermAndTheBands(t *testing.T) {
	path := filepath.Join(t.TempDir(), "dps-per-point.json")
	err := writeMerged(path, map[string]entry{
		"warrior-arms": {
			Reference: "attack_power", DPS: 0.2342, WeaponDPS: 3.21456,
			Bands: map[string]bandEntry{"20": {Level: 25, DPS: 0.051234, DPSError: 0.00123, WeaponDPS: 0.712345, WeaponError: 0.0201}},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	want := "{\n" +
		`  "warrior-arms": {"reference":"attack_power","dps":0.2342,"healing":false,"weapon_dps":3.2146,` +
		`"bands":{"20":{"level":25,"dps":0.0512,"dps_error":0.0012,"weapon_dps":0.7123,"weapon_dps_error":0.0201}}}` + "\n" +
		"}\n"
	if string(b) != want {
		t.Fatalf("file =\n%s\nwant\n%s", b, want)
	}
}
