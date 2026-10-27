package guilds

import "testing"

// The WP 9 ruling (spec §3.5): first seen past the grace is late only for a raider signed In; a
// raider signed Late who arrives late did what they said and is present.
func TestLateIsOnlyForRaidersSignedIn(t *testing.T) {
	past := lateGraceSeconds + 60
	in := attendanceInput{
		signups: []Signup{
			{CharacterKey: "us/pvp/a", Name: "A", State: SignupIn, Role: "dps"},
			{CharacterKey: "us/pvp/b", Name: "B", State: SignupLate, Role: "dps"},
		},
		present: map[string]logPresence{
			"us/pvp/a": {Name: "A", FirstSeenOffsetS: &past, PullsPresent: 3},
			"us/pvp/b": {Name: "B", FirstSeenOffsetS: &past, PullsPresent: 3},
		},
		pullsTotal: 4,
	}
	rows := computeAttendance(in)
	if len(rows) != 2 || rows[0].Outcome != OutcomeLate || rows[1].Outcome != OutcomePresent {
		t.Fatalf("rows = %+v", rows)
	}
	if rows[1].FirstSeenOffsetS == nil || *rows[1].FirstSeenOffsetS != past {
		t.Fatalf("the signed-Late row keeps its first seen: %+v", rows[1])
	}
}

func TestLootItemFactsCarriesTheLootTabIcon(t *testing.T) {
	if name, quality, icon := lootItemFacts(16914); name != "Netherwind Crown" || quality != 4 || icon != "inv_helmet_70" {
		t.Fatalf("16914 = %q %d %q", name, quality, icon)
	}
	// A quest item's icon comes from the client's Item.IconFileDataID, not the class item files.
	if _, _, icon := lootItemFacts(18423); icon != "inv_misc_head_dragon_01" {
		t.Fatalf("Head of Onyxia = %q", icon)
	}
	if name, _, icon := lootItemFacts(1); name != "Item 1" || icon != "" {
		t.Fatalf("unknown = %q %q", name, icon)
	}
}
