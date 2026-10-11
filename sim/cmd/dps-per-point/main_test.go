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

// fakeMeasurer returns canned per-point figures by spec.
type fakeMeasurer map[string][2]float64

func (f fakeMeasurer) Measure(spec specs.Spec, _ api.SimRequest) (float64, float64, error) {
	v, ok := f[spec.Spec]
	if !ok {
		return 0, 0, errors.New("no figure")
	}
	return v[0], v[1], nil
}

func fakeCharacter(spec specs.Spec) (api.CharacterSpec, error) {
	return api.CharacterSpec{Name: "x", Race: "human", Class: spec.ClassSlug, Level: 60}, nil
}

func TestMeasureKeepsOnlyFiguresPositiveBeyondTheirError(t *testing.T) {
	m := fakeMeasurer{
		"warrior-arms": {0.4213, 0.01},
		"priest-holy":  {1.25, 0.05},
		"mage-fire":    {0.02, 0.05}, // inside its own error: noise, not a number
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
	if got["warrior-arms"] != (entry{Reference: "attack_power", DPS: 0.4213, Healing: false}) {
		t.Fatalf("warrior-arms = %+v", got["warrior-arms"])
	}
	if got["priest-holy"] != (entry{Reference: "healing_power", DPS: 1.25, Healing: true}) {
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
		`  "mage-fire": {"reference":"spell_power","dps":0.9,"healing":false},` + "\n" +
		`  "priest-holy": {"reference":"healing_power","dps":1.25,"healing":true},` + "\n" +
		`  "warrior-arms": {"reference":"attack_power","dps":0.4213,"healing":false}` + "\n" +
		"}\n"
	if string(b) != want {
		t.Fatalf("file =\n%s\nwant\n%s", b, want)
	}
}
