package adapter

import (
	"encoding/json"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
)

func spellAction(id int32) *proto.ActionID {
	return &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: id}}
}

func attackAction(tag int32) *proto.ActionID {
	return &proto.ActionID{RawId: &proto.ActionID_OtherId{OtherId: proto.OtherAction_OtherActionAttack}, Tag: tag}
}

func TestHistogramIsContiguousAndSumsToIterations(t *testing.T) {
	h := dpsHistogram(&proto.DistributionMetrics{Hist: map[int32]int32{900: 4, 930: 1, 910: 5}})
	if h.Start != 900 || h.BinDPS != 10 {
		t.Fatalf("start %d bin %d, want 900 and 10", h.Start, h.BinDPS)
	}
	want := []int{4, 5, 0, 1}
	if len(h.Counts) != len(want) {
		t.Fatalf("counts = %v, want %v", h.Counts, want)
	}
	for i := range want {
		if h.Counts[i] != want[i] {
			t.Errorf("counts = %v, want %v", h.Counts, want)
		}
	}
}

func TestHistogramWidensPastTheBinLimit(t *testing.T) {
	h := dpsHistogram(&proto.DistributionMetrics{Hist: map[int32]int32{0: 1, 5000: 2, 10000: 3}})
	if len(h.Counts) > maxHistogramBins {
		t.Fatalf("%d bins, the limit is %d", len(h.Counts), maxHistogramBins)
	}
	total := 0
	for _, n := range h.Counts {
		total += n
	}
	if total != 6 {
		t.Errorf("counts sum to %d, want 6", total)
	}
}

func TestNoHistogramWithoutADistribution(t *testing.T) {
	if h := dpsHistogram(&proto.DistributionMetrics{}); h != nil {
		t.Errorf("histogram = %+v, want none", h)
	}
}

// Ruling R5: a self-damage rage source is a gain, never a spend. Bloodrage pays
// in health; the engine books the rage as a gain and the health as a health
// row, which is not a bar the page reads.
func TestBloodrageIsAGainAndHealthIsNotAResource(t *testing.T) {
	u := &proto.UnitMetrics{Resources: []*proto.ResourceMetrics{
		{Id: spellAction(2687), Type: proto.ResourceType_ResourceTypeRage, Gain: 1000, ActualGain: 900},
		{Id: spellAction(2687), Type: proto.ResourceType_ResourceTypeHealth, Gain: -5000, ActualGain: -5000},
		{Id: spellAction(23894), Type: proto.ResourceType_ResourceTypeRage, Gain: -3000, ActualGain: -3000},
		{Id: attackAction(1), Type: proto.ResourceType_ResourceTypeRage, Gain: 8000, ActualGain: 7770},
		{Id: spellAction(2098), Type: proto.ResourceType_ResourceTypeComboPoints, Gain: 5, ActualGain: 5},
	}}
	got := resourceSources(u, 10)
	if len(got) != 1 || got[0].Resource != "rage" {
		t.Fatalf("sources = %+v, want rage alone", got)
	}
	wantGained := []api.ResourceSourceRow{
		{Action: "other:attack/1", Amount: 777, Wasted: 23},
		{Action: "spell:2687", Amount: 90, Wasted: 10},
	}
	for i, row := range got[0].Gained {
		if row != wantGained[i] {
			t.Errorf("gained[%d] = %+v, want %+v", i, row, wantGained[i])
		}
	}
	if len(got[0].Spent) != 1 || got[0].Spent[0] != (api.ResourceSourceRow{Action: "spell:23894", Amount: 300}) {
		t.Errorf("spent = %+v", got[0].Spent)
	}
}

// Every source is named by its action key. Windfury Totem granted extra
// swings until push 113887 made its attack a strike of its own; results
// saved before still carry its simbuffs id (web/src/lib/sim/extra-attacks.ts).
func TestExtraAttackSourcesAreNamedByTheirKey(t *testing.T) {
	u := &proto.UnitMetrics{ExtraAttacks: []*proto.ExtraAttackMetrics{
		{Id: spellAction(15600), Attacks: 50},
		{Id: spellAction(13964), Attacks: 212},
	}}
	got := extraAttackSources(u, 10)
	if len(got) != 2 || got[0] != (api.ExtraAttackSource{Source: "spell:13964", Count: 21.2}) || got[1].Source != "spell:15600" {
		t.Errorf("sources = %+v", got)
	}
}

// A grant that only speeds up the next swing is not a swing of its own: the
// sources share out what the summary's extra attack row reports.
func TestExtraAttackSourcesAgreeWithTheExtraSwingsRow(t *testing.T) {
	u := &proto.UnitMetrics{
		ExtraAttacks: []*proto.ExtraAttackMetrics{
			{Id: spellAction(13964), Attacks: 300},
			{Id: spellAction(15600), Attacks: 100},
		},
		Actions: []*proto.ActionMetrics{
			{Id: attackAction(1), Targets: []*proto.TargetedActionMetrics{{Casts: 999}}},
			{Id: attackAction(3), Targets: []*proto.TargetedActionMetrics{{Casts: 200}}},
		},
	}
	got := extraAttackSources(u, 10)
	if len(got) != 2 || got[0].Count != 15 || got[1].Count != 5 {
		t.Errorf("sources = %+v, want 15 and 5 of the 20 swings", got)
	}
}

func auraEvent(at int64, id *proto.ActionID, target string, active bool, stacks int32) *proto.SampleAuraEvent {
	return &proto.SampleAuraEvent{AtMs: at, AuraId: id, Target: target, Active: active, Stacks: stacks}
}

func TestSampleAurasFoldIntoWindows(t *testing.T) {
	sunder := spellAction(11597)
	got := sampleAuras(&proto.SampleIteration{DurationSeconds: 100, Auras: []*proto.SampleAuraEvent{
		auraEvent(-2000, spellAction(12328), "", true, 0),
		auraEvent(3000, spellAction(12328), "", false, 0),
		auraEvent(10000, spellAction(12328), "", true, 0),
		auraEvent(20000, spellAction(12328), "", false, 0),
		auraEvent(0, sunder, "Target 1", true, 0),
		auraEvent(0, sunder, "Target 1", true, 1),
		auraEvent(200, sunder, "Target 1", true, 2),
		auraEvent(100000, sunder, "Target 1", false, 0),
		// The engine's extra attack counter, and an aura never up in the fight.
		auraEvent(0, spellAction(21919), "", true, 0),
		auraEvent(100000, spellAction(21919), "", false, 0),
		auraEvent(-500, spellAction(7), "", true, 0),
		auraEvent(-100, spellAction(7), "", false, 0),
	}})
	if len(got) != 2 {
		t.Fatalf("auras = %+v, want Sweeping Strikes and Sunder only", got)
	}
	player := got[0]
	if player.Action != "spell:12328" || player.On != "player" ||
		len(player.Segments) != 2 || player.Segments[0] != [2]int64{0, 3000} || player.Segments[1] != [2]int64{10000, 20000} {
		t.Errorf("player aura = %+v", player)
	}
	target := got[1]
	if target.On != "target" || len(target.Stacks) != 2 || target.Stacks[0] != [2]int64{0, 1} || target.Stacks[1] != [2]int64{200, 2} {
		t.Errorf("target aura = %+v", target)
	}
}

func TestDebuffTracksComeFromTheTargetMetrics(t *testing.T) {
	targets := []*proto.UnitMetrics{{
		Name: "Target 1",
		Auras: []*proto.AuraMetrics{
			{Id: spellAction(11597), ProcsAvg: 5, UptimeSecondsAvg: 180.1},
			{Id: spellAction(11556), ProcsAvg: 0, UptimeSecondsAvg: 0},
		},
	}}
	sampled := []api.SampleAura{{Action: "spell:11597", On: "target", Stacks: [][2]int64{{0, 1}, {800, 5}}}}
	rows := debuffTracks(targets, sampled)
	if len(rows) != 1 {
		t.Fatalf("rows = %+v, want the one debuff that was up", rows)
	}
	row := rows[0]
	if row.Type != "DEBUFF" || row.TargetGUID != "target-1" || row.TargetName != "Target 1" ||
		row.UptimeMS != 180100 || row.MaxStacks != 5 || row.Applications != 5 {
		t.Errorf("row = %+v", row)
	}
	if row.Segments == nil || row.Appliers == nil {
		t.Error("a nil list marshals as null, which the page does not parse")
	}
}

func TestSheetKeepsZeroHasteAndHidesSpellStatsOnAWarrior(t *testing.T) {
	sheet := &proto.SampleSheet{AttackPower: 1775.2, MeleeHitPct: 9.5, MeleeCritPctVsTarget: 29.26, MainHandSkill: 305, TargetLevel: 63, TargetArmor: 3731, TargetArmorAfterDebuffs: 471}
	final, target := sheetStats(sheet, "warrior-fury")
	b, err := json.Marshal(final)
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]any
	if err := json.Unmarshal(b, &fields); err != nil {
		t.Fatal(err)
	}
	if fields["melee_haste_pct"] != 0.0 {
		t.Errorf("melee_haste_pct = %v, want an explicit 0 (ruling R3)", fields["melee_haste_pct"])
	}
	if _, ok := fields["spell_power"]; ok {
		t.Errorf("a warrior has no spell block: %v", fields)
	}
	if final.AttackPower != 1775 || final.MeleeCritPctVsTarget != 29.3 || final.WeaponSkill.MainHand != 305 {
		t.Errorf("final = %+v", final)
	}
	if target.Level != 63 || target.Armor != 3731 || target.ArmorAfterDebuffs != 471 {
		t.Errorf("target = %+v", target)
	}
}

func TestSheetUsesRangedPowerForARangedHunter(t *testing.T) {
	sheet := &proto.SampleSheet{AttackPower: 100, RangedAttackPower: 900, HasManaBar: true}
	if final, _ := sheetStats(sheet, "hunter-marksmanship"); final.AttackPower != 900 || final.SpellPower != nil {
		t.Errorf("marksmanship = %+v", final)
	}
	if final, _ := sheetStats(sheet, "hunter-survival"); final.AttackPower != 100 {
		t.Errorf("survival fights in melee: %+v", final)
	}
	if final, _ := sheetStats(sheet, "mage-frost"); final.SpellPower == nil {
		t.Errorf("a mage has the spell block: %+v", final)
	}
}

func TestSampleRowsCarryOutcomeAndAmount(t *testing.T) {
	cast := func(outcome proto.SampleOutcome, amount float64) *proto.SampleCast {
		return &proto.SampleCast{ActionId: spellAction(1), Outcome: outcome, Amount: amount}
	}
	rows := Sample(&proto.RaidSimResult{SampleIteration: &proto.SampleIteration{Casts: []*proto.SampleCast{
		cast(proto.SampleOutcome_SampleOutcomeCrit, 2274.6),
		cast(proto.SampleOutcome_SampleOutcomeTickCrit, 377),
		cast(proto.SampleOutcome_SampleOutcomeMiss, 0),
		cast(proto.SampleOutcome_SampleOutcomeHit, 0),
		cast(proto.SampleOutcome_SampleOutcomeNone, 0),
	}}})
	want := []struct {
		outcome string
		amount  int64
	}{{"crit", 2275}, {"tick_crit", 377}, {"miss", 0}, {"", 0}, {"", 0}}
	for i, w := range want {
		if rows[i].Outcome != w.outcome || rows[i].Amount != w.amount {
			t.Errorf("row %d = %+v, want %+v", i, rows[i], w)
		}
	}
}

func TestWithDetailLeavesItsArgumentAlone(t *testing.T) {
	res := &proto.RaidSimResult{
		IterationsDone: 2,
		RaidMetrics: &proto.RaidMetrics{
			Dps:     &proto.DistributionMetrics{Hist: map[int32]int32{10: 2}},
			Parties: []*proto.PartyMetrics{{Players: []*proto.UnitMetrics{{Name: "P"}}}},
		},
	}
	in := api.SimResult{Request: api.SimRequest{Spec: "warrior-fury"}}
	out, err := WithDetail(in, res)
	if err != nil {
		t.Fatal(err)
	}
	if in.DPSHistogram != nil {
		t.Error("the argument changed")
	}
	if out.DPSHistogram == nil {
		t.Error("no histogram")
	}
	if out.FinalStats != nil || out.TargetStats != nil {
		t.Errorf("no sample means no sheet, and the page hides the block: %+v %+v", out.FinalStats, out.TargetStats)
	}
}

func TestDamageTableLosesTheEnginesBookkeepingRows(t *testing.T) {
	in := []summary.Actor{{Abilities: []summary.Ability{
		{Name: "other:rage_gain"},
		{Name: "other:attack/3"},
		{Name: "spell:12867", Total: 10},
		{Name: "spell:7"},
	}}}
	got := withoutBookkeepingRows(in)[0].Abilities
	if len(got) != 3 || got[0].Name != "other:attack/3" {
		t.Errorf("rows = %+v", got)
	}
	if len(in[0].Abilities) != 4 {
		t.Error("the argument changed")
	}
}
