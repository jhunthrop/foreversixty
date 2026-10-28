package consumables

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// foreverGolden is the summary of the one real Forever log in the repo
// (logs/engine/event/testdata/forever-1.60.log: beta 1.60.1, a 17 s
// open-world trash fight on Zephras Isle, 18 players). It has no
// COMBATANT_INFO and no classes, which is what it can prove: with no
// character sheet, a real Forever pull reads every player's preparation as
// unknown, never as missing. No real Forever boss pull with a character
// sheet exists in the repo or in the owner's beta Logs folder yet, so the
// flask-present / flask-absent cases below run on hand-built summaries
// that exercise the reader, and are not known-good fixtures (§20.6).
const foreverGolden = "forever-1.60.summary.json.golden"

// foreverTank is Tester01-Beta in forever-1.60.log, a roster row with no
// combatant row.
const foreverTank = "Player-4619-005746D8"

const (
	flaskHolderGUID = "Player-1-FLASK"
	noFlaskGUID     = "Player-1-NOFLASK"
	flaskSpellID    = 17628 // Flask of Supreme Power
	potionSpellID   = 17528 // Mighty Rage Potion
)

func loadGolden(t *testing.T, name string) summary.Summary {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("..", "..", "summary", "testdata", name))
	if err != nil {
		t.Fatal(err)
	}
	var s summary.Summary
	if err := json.Unmarshal(data, &s); err != nil {
		t.Fatal(err)
	}
	return s
}

func seededTrack(guid string, spellID int64) summary.AuraTrack {
	return summary.AuraTrack{TargetGUID: guid, SpellID: spellID,
		Segments: []summary.Segment{{StartMS: 0, EndMS: 60_000}}}
}

// sheetFight is two players with gear, one whose flask was seeded at the
// pull and one whose was not.
func sheetFight() summary.Summary {
	gear := []event.Item{{ID: 1}}
	return summary.Summary{
		Combatants: []summary.CombatantRow{
			{GUID: flaskHolderGUID, Gear: gear},
			{GUID: noFlaskGUID, Gear: gear},
		},
		Auras: []summary.AuraTrack{seededTrack(flaskHolderGUID, flaskSpellID)},
	}
}

func TestAtPull_SeededFlaskIsPresent(t *testing.T) {
	p := AtPull(sheetFight(), flaskHolderGUID, Load().AllRoles())
	if !p.Known || !p.Has(CategoryFlask) {
		t.Fatalf("flask holder: %+v", p)
	}
	if got := p.Present[CategoryFlask]; len(got) != 1 || got[0].SpellID != flaskSpellID {
		t.Fatalf("flask entries = %+v, want only %d", got, flaskSpellID)
	}
}

func TestAtPull_NoFlaskIsAbsent(t *testing.T) {
	if p := AtPull(sheetFight(), noFlaskGUID, Load().AllRoles()); !p.Known || p.Has(CategoryFlask) {
		t.Fatalf("no-flask player: %+v", p)
	}
}

func TestAtPull_NoCharacterSheetIsUnknown(t *testing.T) {
	fight := sheetFight()
	fight.Combatants = nil
	if p := AtPull(fight, flaskHolderGUID, Load().AllRoles()); p.Known || len(p.Present) != 0 {
		t.Fatalf("no sheet must be unknown, got %+v", p)
	}
}

func TestAtPull_CombatantWithoutGearIsUnknown(t *testing.T) {
	fight := sheetFight()
	fight.Combatants[0].Gear = nil
	if p := AtPull(fight, flaskHolderGUID, Load().AllRoles()); p.Known || len(p.Present) != 0 {
		t.Fatalf("a combatant row with no gear is no sheet, got %+v", p)
	}
}

// The real Forever log: every roster player is unknown, none is missing.
func TestAtPull_RealForeverLogWithoutSheetsIsUnknown(t *testing.T) {
	fight := loadGolden(t, foreverGolden)
	if len(fight.Roster) != 18 || len(fight.Combatants) != 0 {
		t.Fatalf("forever-1.60 golden moved: %d roster rows, %d combatants (want 18, 0)",
			len(fight.Roster), len(fight.Combatants))
	}
	all := Load().AllRoles()
	for _, r := range fight.Roster {
		if p := AtPull(fight, r.GUID, all); p.Known || len(p.Present) != 0 || p.PotionUses != 0 {
			t.Errorf("%s (%s): %+v, want unknown", r.Name, r.GUID, p)
		}
	}
	for _, s := range RaidBuffsAtPull(fight, foreverTank, Load().RaidBuffs) {
		if s.Present || s.ProviderInRoster {
			t.Errorf("%s: %+v, want absent with no provider (the log names no class)", s.Buff.Name, s)
		}
	}
}

func TestAtPull_PotionUsesFoldPetsByOwner(t *testing.T) {
	fight := sheetFight()
	fight.Casts = []summary.CastRow{
		{GUID: noFlaskGUID, OwnerGUID: noFlaskGUID, SpellID: potionSpellID, Succeeded: 1},
		{GUID: "Pet-1", OwnerGUID: noFlaskGUID, SpellID: potionSpellID, Succeeded: 1},
		{GUID: flaskHolderGUID, OwnerGUID: flaskHolderGUID, SpellID: potionSpellID, Succeeded: 5},
		{GUID: noFlaskGUID, OwnerGUID: noFlaskGUID, SpellID: flaskSpellID, Succeeded: 3},
	}
	if p := AtPull(fight, noFlaskGUID, Load().AllRoles()); p.PotionUses != 2 {
		t.Fatalf("PotionUses = %d, want 2 (own + pet, not another player's, not a non-potion)", p.PotionUses)
	}
}

func TestSeeded_OnlyZeroStartSegments(t *testing.T) {
	// a flask track on the player whose first segment starts at 1200 ms is not seeded
	fight := summary.Summary{Auras: []summary.AuraTrack{{TargetGUID: "P", SpellID: flaskSpellID,
		Segments: []summary.Segment{{StartMS: 1200, EndMS: 9000}}}}}
	if Seeded(fight, "P", flaskSpellID) {
		t.Fatal("a mid-fight application is not at the pull")
	}
	fight.Auras = append(fight.Auras, seededTrack("Q", flaskSpellID))
	if Seeded(fight, "P", flaskSpellID) || !Seeded(fight, "Q", flaskSpellID) {
		t.Fatal("seeded must be per target")
	}
}

func TestRaidBuffsAtPull_ProviderInRoster(t *testing.T) {
	fight := summary.Summary{Roster: []summary.RosterRow{{GUID: "P", Class: "Warrior"}, {GUID: "M", Class: "Mage"}}}
	got := RaidBuffsAtPull(fight, "P", Load().RaidBuffs)
	byName := map[string]RaidBuffState{}
	for _, s := range got {
		byName[s.Buff.Name] = s
	}
	if !byName["Arcane Intellect"].ProviderInRoster || byName["Power Word: Fortitude"].ProviderInRoster {
		t.Fatalf("provider flags: %+v", byName)
	}
}

func TestRaidBuffsAtPull_AnyRankIsPresent(t *testing.T) {
	const fortitudeR6 = 10938
	fight := summary.Summary{Auras: []summary.AuraTrack{seededTrack("P", fortitudeR6)}}
	for _, s := range RaidBuffsAtPull(fight, "P", Load().RaidBuffs) {
		if want := s.Buff.Name == "Power Word: Fortitude"; s.Present != want {
			t.Errorf("%s present = %v, want %v", s.Buff.Name, s.Present, want)
		}
	}
}
