package main

import (
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

const (
	healTestIterations = 2.0
	healTestSeconds    = 10.0
)

func manaResource(id *proto.ActionID, gain, actual float64, events int32) *proto.ResourceMetrics {
	return &proto.ResourceMetrics{Id: id, Type: proto.ResourceType_ResourceTypeMana, Gain: gain, ActualGain: actual, Events: events}
}

// A healer's mana resources split into what came in (actual gain) and what
// the spells paid (negative gain), each per iteration, with the healing the
// spell landed beside its spend.
func TestSplitManaFlowsSeparatesIncomeFromSpend(t *testing.T) {
	player := &proto.UnitMetrics{
		Actions: []*proto.ActionMetrics{{
			Id:      spellID(2054, 0),
			Targets: []*proto.TargetedActionMetrics{{Casts: 6, EffectiveHealing: 1200}},
		}},
		Resources: []*proto.ResourceMetrics{
			manaResource(spellID(2054, 0), -300, -300, 6),
			manaResource(spellID(1317006, 0), 80, 60, 4),
			{Id: spellID(99, 0), Type: proto.ResourceType_ResourceTypeRage, Gain: 50, ActualGain: 50},
		},
	}
	names := map[int32]string{2054: "Heal", 1317006: "Litany of Light"}

	income, spend := splitManaFlows(player, names, healTestIterations)

	if len(income) != 1 || income[0].Name != "Litany of Light (1317006)" || income[0].Mana != 30 || income[0].Casts != 2 {
		t.Errorf("income = %+v, want one Litany of Light row of 30 mana in 2 events", income)
	}
	if len(spend) != 1 || spend[0].Name != "Heal (2054)" {
		t.Fatalf("spend = %+v, want one Heal row", spend)
	}
	if spend[0].Mana != 150 || spend[0].Casts != 3 || spend[0].Healed != 600 {
		t.Errorf("Heal row = %+v, want 150 mana, 3 casts and 600 healed a fight", spend[0])
	}
}

func TestWriteManaTableTotalsTheRowsAndPricesTheHealing(t *testing.T) {
	var b strings.Builder
	rows := []manaRow{{Name: "Heal (2054)", Casts: 3, Mana: 150, Healed: 600}, {Name: "Renew (139)", Casts: 1, Mana: 50, Healed: 100}}

	writeManaTable(&b, "Mana spend", "per fight", "per second", rows, healTestSeconds, true)

	out := b.String()
	for _, want := range []string{"| Heal (2054) | 3.0 | 150 | 15.00 | 600 | 4.00 |", "| total | | 200 | 20.00 |"} {
		if !strings.Contains(out, want) {
			t.Errorf("table lacks %q:\n%s", want, out)
		}
	}
}

func TestWriteFreeCastsListsOnlyActionsWithoutAManaRow(t *testing.T) {
	player := &proto.UnitMetrics{Actions: []*proto.ActionMetrics{
		{Id: spellID(2054, 0), Targets: []*proto.TargetedActionMetrics{{Casts: 6}}},
		{Id: spellID(14751, 0), Targets: []*proto.TargetedActionMetrics{{Casts: 4}}},
		{Id: spellID(1, 0), Targets: []*proto.TargetedActionMetrics{{Casts: 0}}},
	}}
	names := map[int32]string{2054: "Heal", 14751: "Inner Focus", 1: "Unused"}
	var b strings.Builder

	writeFreeCasts(&b, player, names, healTestIterations, []manaRow{{Name: "Heal (2054)"}})

	out := b.String()
	if !strings.Contains(out, "| Inner Focus (14751) | 2.0 |") {
		t.Errorf("the free cast is missing:\n%s", out)
	}
	if strings.Contains(out, "Heal (2054)") || strings.Contains(out, "Unused") {
		t.Errorf("a paid or uncast action is listed:\n%s", out)
	}
}

func TestSplitManaFlowsOfAnEmptyRunIsEmpty(t *testing.T) {
	income, spend := splitManaFlows(&proto.UnitMetrics{}, nil, healTestIterations)
	if len(income) != 0 || len(spend) != 0 {
		t.Errorf("income %v and spend %v, want both empty", income, spend)
	}
}
