package sims

import (
	"reflect"
	"slices"
	"testing"

	simapi "github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

func TestTheEncounterDefaultsAreTheDesignsSettingsBar(t *testing.T) {
	got := withEncounterDefaults(simapi.EncounterSpec{})
	want := simapi.DefaultEncounter()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("defaults: got %+v want %+v", got, want)
	}
	if want.DurationSec != 180 || want.Targets != 1 {
		t.Fatalf("the envelope's defaults moved: %+v", want)
	}
	// A value the caller did set survives.
	set := withEncounterDefaults(simapi.EncounterSpec{DurationSec: 300, Targets: 4})
	if set.DurationSec != 300 || set.Targets != 4 || set.Variation != want.Variation {
		t.Fatalf("explicit values overwritten: %+v", set)
	}
}

func TestADefaultedRequestPassesTheEnvelopesOwnValidation(t *testing.T) {
	req := simapi.SimRequest{
		EngineVersion: enginever.Version, Spec: "warrior-fury",
		Source:     simapi.CharacterSource{Kind: simapi.SourceAddon, Ref: "us/normal/baelgrim"},
		Character:  simapi.CharacterSpec{Name: "Baelgrim", Race: "orc", Class: "warrior", Level: 60},
		Encounter:  withEncounterDefaults(simapi.EncounterSpec{}),
		Iterations: defaultIterations,
	}
	if err := req.Validate(); err != nil {
		t.Fatalf("a good request was rejected: %v", err)
	}
	req.Iterations = 7
	if err := req.Validate(); err == nil {
		t.Fatal("an iteration count outside the closed set was accepted")
	}
	// One level runs, and the envelope is the only thing that says so.
	req.Iterations = defaultIterations
	req.Character.Level = simapi.SimLevel - 1
	if err := req.Validate(); err == nil {
		t.Fatalf("level %d was accepted; only %d runs", req.Character.Level, simapi.SimLevel)
	}
}

func TestTheIterationCountsAreTheOnesTheEnvelopeAccepts(t *testing.T) {
	want := []int{liveIterations, defaultIterations, preciseIterations}
	if !slices.Equal(simapi.ValidIterations, want) {
		t.Fatalf("ValidIterations = %v, want %v", simapi.ValidIterations, want)
	}
}

func TestTheSpecListIsTheDataLanesAndCarriesDPSSpecs(t *testing.T) {
	if len(specs.All) == 0 {
		t.Fatal("the generated spec list is empty")
	}
	dps := DPSSpecs()
	if len(dps) == 0 || len(dps) >= len(specs.All) {
		t.Fatalf("%d dps specs out of %d; the role filter is wrong", len(dps), len(specs.All))
	}
	if !slices.IsSorted(dps) {
		t.Fatalf("DPSSpecs is not sorted: %v", dps)
	}
	if !slices.Contains(dps, "warrior-fury") || !slices.Contains(dps, "mage-frost") {
		t.Fatalf("the two launch specs are missing: %v", dps)
	}
	for _, s := range specs.All {
		if s.Spec != s.ClassSlug+"-"+s.SpecSlug {
			t.Errorf("%q is not <class>-<spec>", s.Spec)
		}
	}
}

func TestTheEnginePinIsASha(t *testing.T) {
	if len(enginever.Version) < 7 {
		t.Fatalf("engine version %q does not look like a short sha", enginever.Version)
	}
}
