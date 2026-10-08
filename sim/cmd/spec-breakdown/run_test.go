package main

import (
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

// realRoot is the repository the committed BiS files and presets live in.
const realRoot = "../../.."

// captureStdout runs fn and returns what it printed.
func captureStdout(t *testing.T, fn func() error) (string, error) {
	t.Helper()
	orig := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w
	runErr := fn()
	os.Stdout = orig
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	return string(out), runErr
}

func realOptions(spec string) options {
	return options{repoRoot: realRoot, spec: spec, preset: "raid", band: 60, faction: "alliance", iterations: 3, seed: 1}
}

func TestRunPrintsTheMarkdownBreakdownOfARaidEntry(t *testing.T) {
	out, err := captureStdout(t, func() error { return run(realOptions(testSpec)) })
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{testSpec + "  DPS ", "consumes:", "| Action | DPS |", "| Aura | Uptime % |"} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q:\n%s", want, out)
		}
	}
}

func TestRunWithJSONPrintsDecodableTotals(t *testing.T) {
	o := realOptions(testSpec)
	o.jsonOut = true
	out, err := captureStdout(t, func() error { return run(o) })
	if err != nil {
		t.Fatal(err)
	}
	var got report
	if err := json.Unmarshal([]byte(out), &got); err != nil {
		t.Fatalf("not JSON: %v\n%s", err, out)
	}
	if got.Spec != testSpec || got.DPS <= 0 || len(got.Actions) == 0 {
		t.Fatalf("report = %+v", got)
	}
}

func TestGearOverrideReachesTheRequestAndABadPairFailsTheRun(t *testing.T) {
	o := realOptions(testSpec)
	o.gear = "head:not-a-number"
	if err := run(o); err == nil {
		t.Fatal("a malformed -gear pair ran")
	}
	o.gear = "head:1"
	req, err := buildRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, s := range req.Character.Gear {
		if s.Slot == "head" {
			found = true
			if s.ItemID != 1 {
				t.Fatalf("head = %d, want the override 1", s.ItemID)
			}
		}
	}
	if !found {
		t.Fatal("the overridden slot vanished")
	}
}

func TestRunRejectsAMissingRotationFile(t *testing.T) {
	o := realOptions(testSpec)
	o.rotation = t.TempDir() + "/missing.json"
	if err := run(o); err == nil {
		t.Fatal("a missing -rotation file ran")
	}
}

func TestRunTankPrintsTheDefensiveReport(t *testing.T) {
	o := realOptions("warrior-protection")
	o.tank = true
	out, err := captureStdout(t, func() error { return run(o) })
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"warrior-protection tank fight, 3 iterations", "| health |", "| dodge % |", "| chance of death |", "| Boss swings per fight |", "| Tank action |", "| Aura |"} {
		if !strings.Contains(out, want) {
			t.Errorf("tank report missing %q:\n%s", want, out)
		}
	}
}

func TestRunTankRefusesADPSSpec(t *testing.T) {
	o := realOptions(testSpec)
	o.tank = true
	if _, err := runTankReport(o); err == nil || !strings.Contains(err.Error(), "-tank needs a tank spec") {
		t.Fatalf("err = %v", err)
	}
	o.spec = "no-such-spec"
	if _, err := runTankReport(o); err == nil {
		t.Fatal("unknown spec ran")
	}
}

func finalStatsResult(values map[stats.Stat]float64) *proto.ComputeStatsResult {
	final := make([]float64, stats.Len)
	for s, v := range values {
		final[s] = v
	}
	return &proto.ComputeStatsResult{RaidStats: &proto.RaidStats{Parties: []*proto.PartyStats{{
		Players: []*proto.PlayerStats{{FinalStats: &proto.UnitStats{Stats: final, PseudoStats: make([]float64, len(proto.PseudoStat_name))}}},
	}}}}
}

func TestFirstFinalStatsReadsTheFirstPlayerOnly(t *testing.T) {
	if firstFinalStats(&proto.ComputeStatsResult{}) != nil {
		t.Fatal("an empty raid has no final stats")
	}
	got := firstFinalStats(finalStatsResult(map[stats.Stat]float64{stats.Health: 4321}))
	if got.GetStats()[stats.Health] != 4321 {
		t.Fatalf("health = %v", got.GetStats()[stats.Health])
	}
}

func TestTankMarkdownDividesTotalsByIterationsAndSkipsIdleRows(t *testing.T) {
	req := api.SimRequest{Spec: "warrior-protection", Encounter: api.DefaultEncounter()}
	req.Encounter.DurationSec = 10
	res := &proto.RaidSimResult{IterationsDone: 2}
	player := &proto.UnitMetrics{
		Dtps:          &proto.DistributionMetrics{Avg: 321.5},
		ChanceOfDeath: 0.25,
		Actions: []*proto.ActionMetrics{
			{Id: spellID(23922, 0), Targets: []*proto.TargetedActionMetrics{{Casts: 8, Shielding: 400, Threat: 600}}},
			{Id: spellID(99, 0), Targets: []*proto.TargetedActionMetrics{{}}}, // never cast: skipped
		},
		Auras: []*proto.AuraMetrics{
			{Id: spellID(2565, 0), UptimeSecondsAvg: 5, ProcsAvg: 1.5},
			{Id: spellID(98, 0)}, // never up: skipped
		},
	}
	out := tankMarkdown(req, finalStatsResult(map[stats.Stat]float64{stats.Health: 3000, stats.Dodge: 5.5}), res, player,
		map[int32]string{23922: "Shield Slam", 2565: "Shield Block"})
	for _, want := range []string{
		"warrior-protection tank fight, 2 iterations",
		"| health | 3000.00 |",
		"| dodge % | 5.50 |",
		"| damage taken /s | 321.5 |",
		"| chance of death | 25.0% |",
		"| Shield Slam (23922) | 4.00 | 20.0 | 30.0 |", // 8 casts / 2 iterations; 400 / 2 / 10 s; 600 / 2 / 10 s
		"| Shield Block (2565) | 50.0 | 1.50 |",        // 5 s of 10 s
	} {
		if !strings.Contains(out, want) {
			t.Errorf("missing %q in:\n%s", want, out)
		}
	}
	if strings.Contains(out, "spell (99)") || strings.Contains(out, "spell (98)") {
		t.Errorf("idle rows must be skipped:\n%s", out)
	}
	if bare := tankMarkdown(req, &proto.ComputeStatsResult{}, res, player, nil); strings.Contains(bare, "| health |") {
		t.Error("no final stats, no stat table")
	}
}

func TestBossSwingsWithNoSwingsIsZero(t *testing.T) {
	if got := bossSwings(&proto.RaidSimResult{}, 1); got != (bossSwingRow{}) {
		t.Fatalf("got %+v", got)
	}
}

func TestBossSwingsBlockedShareOfDamage(t *testing.T) {
	res := &proto.RaidSimResult{EncounterMetrics: &proto.EncounterMetrics{Targets: []*proto.UnitMetrics{{
		Actions: []*proto.ActionMetrics{{Targets: []*proto.TargetedActionMetrics{{Hits: 4, Blocks: 4, Damage: 600, BlockDamage: 200}}}},
	}}}}
	got := bossSwings(res, 2)
	if !near(got.Swings, 4) || !near(got.BlockedSwingsDamageShare, 25) {
		t.Fatalf("got %+v", got)
	}
}
