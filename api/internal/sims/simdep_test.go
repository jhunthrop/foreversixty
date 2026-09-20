package sims

import (
	"reflect"
	"slices"
	"testing"
	"time"

	simapi "github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

func TestTheEncounterDefaultsAreTheDesignsSettingsBar(t *testing.T) {
	got := withEncounterDefaults(simapi.EncounterSpec{})
	want := simapi.DefaultEncounter()
	// EncounterSpec now carries TargetsOverTime ([]TargetCount), so it is
	// no longer comparable with ==; DeepEqual is the correct successor.
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

func TestTheEstimateIsIterationsOverTheJobsMeasuredRate(t *testing.T) {
	// Stated in multiples of the rate, not in literals: the rate is
	// sim/measure's benchmark figure and moves when the benchmark does.
	for _, c := range []struct {
		iterations int
		want       int
	}{
		{0, 0},
		{1, 1}, // rounded up: a run is never estimated at no time at all
		{nativeRate, 1},
		{nativeRate * 60, 60},
		{nativeRate*4000 + 1, 4001},
	} {
		if got := estimateSec(c.iterations); got != c.want {
			t.Errorf("estimateSec(%d) = %d, want %d", c.iterations, got, c.want)
		}
	}
}

func TestTheBulkBudgetFitsInsideTheCloudRunTaskTimeout(t *testing.T) {
	// api/README.md creates sim-run with --task-timeout 15m, and
	// contract A2 fixes the budget at 840 seconds. The in-process bound
	// has to leave room for start-up and the two writes at the end, or
	// the platform kills the job mid-write and the row never leaves
	// "running".
	const taskTimeout = 15 * time.Minute
	if BulkBudget != 840*time.Second {
		t.Fatalf("BulkBudget %s, want the contract's 840s", BulkBudget)
	}
	if BulkBudget >= taskTimeout {
		t.Fatalf("BulkBudget %s, want less than the job's %s", BulkBudget, taskTimeout)
	}
	if RunTimeout > BulkBudget {
		t.Fatalf("a plain run (%s) may not outlast a bulk one (%s)", RunTimeout, BulkBudget)
	}
}

func TestTheServerCapAndTheBudgetAgree(t *testing.T) {
	// Contract A2 lowered the server cap to 5,000 so that a full-cap fast
	// run can actually finish. LadderIterations is the module's own
	// costing of a ladder (sim/api/bulk.go), not a hand-rolled
	// approximation of it - it accounts for the +1 equipped baseline
	// that runs alongside every delta in every stage, which a bare
	// cap*100 + (cap/4)*1000 + 11*3000 leaves out.
	cap := simapi.Caps[simapi.LaneServer]
	budget := int(BulkBudget.Seconds())

	fast := simapi.LadderIterations(simapi.Ladders[simapi.PrecisionFast], cap)
	if est := estimateSec(fast); est > budget {
		t.Fatalf("a full-cap fast run estimates %ds against a %ds budget: "+
			"the cap and the budget disagree", est, budget)
	}

	// Normal and high at full cap deliberately do NOT fit the budget.
	// This is the intended division of labour: Caps[LaneServer] bounds
	// memory and UX (how big a request the server will even expand),
	// while the estimate-based too_large refusal bounds engine time.
	// The cap does not promise every precision finishes at full cap;
	// only the estimate does that, per request. If someone later moves
	// the cap, a ladder's stage sizes, or the measured rate, one of
	// these should flip and say which side of the trade moved.
	for _, precision := range []string{simapi.PrecisionNormal, simapi.PrecisionHigh} {
		total := simapi.LadderIterations(simapi.Ladders[precision], cap)
		est := estimateSec(total)
		if est <= budget {
			t.Fatalf("a full-cap %s run estimates %ds, within the %ds budget: "+
				"expected it to exceed the budget (too_large is what bounds it, not the cap)",
				precision, est, budget)
		}
	}
}
