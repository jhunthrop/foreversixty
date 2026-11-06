package guilds

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const (
	digestFlaskID = 17628 // Flask of Supreme Power, in the catalogue and in 1.60.1.70790's spells.json
	tankGUID      = "Player-1-TANK"
	mageGUID      = "Player-1-MAGE"
)

func hit(at int64, spell int64, name, source string, amount, overkill int64) summary.DamageRef {
	return summary.DamageRef{AtMS: at, SpellID: spell, SpellName: name, SourceName: source, Amount: amount, Overkill: overkill}
}

// deathSummary is one pull: the mage died to Deep Breath after five hits, the tank has a sheet
// and a flask up, the mage a sheet and no flask.
func deathSummary(withSheets bool) summary.Summary {
	sum := summary.Summary{
		Roster: []summary.RosterRow{{GUID: tankGUID, Name: "Thornhide", Role: "tank"}, {GUID: mageGUID, Name: "Sunderfel", Role: "dps"}},
		Deaths: []summary.Death{{GUID: mageGUID, Name: "Sunderfel", Class: "Mage", AtMS: 61_000,
			KillingBlow: ptrRef(hit(60_900, 18435, "Deep Breath", "Onyxia", 4000, 2100)),
			Last: []summary.DamageRef{hit(55_000, 1, "Melee", "Onyxian Whelp", 100, 0), hit(56_000, 1, "Melee", "Onyxian Whelp", 100, 0),
				hit(58_000, 18392, "Fireball", "Onyxia", 900, 0), hit(59_000, 1, "Melee", "Onyxian Whelp", 120, 0),
				hit(60_900, 18435, "Deep Breath", "Onyxia", 4000, 2100)}}},
		Auras: []summary.AuraTrack{{TargetGUID: tankGUID, SpellID: digestFlaskID, Seeded: true,
			Segments: []summary.Segment{{StartMS: 0, EndMS: 60_000}}}},
	}
	if withSheets {
		gear := []event.Item{{ID: 1}}
		sum.Combatants = []summary.CombatantRow{{GUID: tankGUID, Gear: gear}, {GUID: mageGUID, Gear: gear}}
	}
	return sum
}

func ptrRef(r summary.DamageRef) *summary.DamageRef { return &r }

func TestDigestKeepsTheKillingBlowAndTheLastThreeHits(t *testing.T) {
	d := digestOf(deathSummary(true))
	if len(d.Deaths) != 1 {
		t.Fatalf("deaths = %+v", d.Deaths)
	}
	death := d.Deaths[0]
	if death.KillingBlow == nil || death.KillingBlow.Ability != "Deep Breath" || death.KillingBlow.SpellID != 18435 ||
		death.KillingBlow.Source != "Onyxia" || death.KillingBlow.Overkill != 2100 || death.AtMS != 61_000 {
		t.Fatalf("death = %+v", death)
	}
	if len(death.Last) != deathLastHits || death.Last[0].Ability != "Fireball" || death.Last[2].Ability != "Deep Breath" {
		t.Fatalf("last hits = %+v, want the final three, oldest first", death.Last)
	}
}

func TestDigestReadsConsumablesOnlyFromACharacterSheet(t *testing.T) {
	if d := digestOf(deathSummary(false)); d.Consumables != nil {
		t.Fatalf("no sheet: consumables = %+v, want nil (unknown, never missing)", d.Consumables)
	}
	d := digestOf(deathSummary(true))
	by := map[string]digestConsumable{}
	for _, c := range d.Consumables {
		by[c.Name] = c
	}
	tank, mage := by["Thornhide"], by["Sunderfel"]
	if !tank.Known || tank.Flask == nil || !*tank.Flask || len(tank.Brought) != 1 {
		t.Fatalf("tank = %+v", tank)
	}
	if !mage.Known || mage.Flask == nil || *mage.Flask {
		t.Fatalf("mage flask = %+v, want proven missing (the flask list is complete)", mage)
	}
	if mage.Food != nil {
		t.Fatalf("mage food = %v, want null: the food list can prove presence, never absence", *mage.Food)
	}
}

func TestADeathWithNoHitHasNoKillingBlow(t *testing.T) {
	d := digestOf(summary.Summary{Deaths: []summary.Death{{Name: "Ghost", AtMS: 5}}})
	if d.Deaths[0].KillingBlow != nil || d.Deaths[0].Last == nil || len(d.Deaths[0].Last) != 0 {
		t.Fatalf("death = %+v", d.Deaths[0])
	}
}
