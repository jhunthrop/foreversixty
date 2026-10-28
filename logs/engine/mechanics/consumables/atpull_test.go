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
	return summary.AuraTrack{TargetGUID: guid, SpellID: spellID, Seeded: true,
		Segments: []summary.Segment{{StartMS: 0, EndMS: 60_000}}}
}

// appliedTrack is an aura APPLIED during the fight at startMS, not in the
// snapshot.
func appliedTrack(guid string, spellID, startMS int64) summary.AuraTrack {
	return summary.AuraTrack{TargetGUID: guid, SpellID: spellID,
		Segments: []summary.Segment{{StartMS: startMS, EndMS: startMS + 60_000}}}
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
		if s.Known || s.Present || s.ProviderInRoster {
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

func TestSeeded_ReadsTheSnapshotFlagNotTheStart(t *testing.T) {
	// an APPLIED at 0 ms with no snapshot is not seeded; a snapshot track whose
	// segment starts at 1 ms (COMBATANT_INFO written after ENCOUNTER_START) is
	fight := summary.Summary{Auras: []summary.AuraTrack{appliedTrack("P", flaskSpellID, 0)}}
	if Seeded(fight, "P", flaskSpellID) {
		t.Fatal("an APPLIED line is not the snapshot")
	}
	late := seededTrack("Q", flaskSpellID)
	late.Segments[0].StartMS = 1
	fight.Auras = append(fight.Auras, late)
	if !Seeded(fight, "Q", flaskSpellID) {
		t.Fatal("a snapshot written 1 ms after the pull opened is still the pull")
	}
	if Seeded(fight, "R", flaskSpellID) {
		t.Fatal("seeded must be per target")
	}
}

func TestAtPull_LateSnapshotStillCountsTheFlask(t *testing.T) {
	fight := sheetFight()
	fight.Auras[0].Segments[0].StartMS = 1
	if p := AtPull(fight, flaskHolderGUID, Load().AllRoles()); !p.Has(CategoryFlask) {
		t.Fatalf("a snapshot 1 ms late must not read the flask as missing: %+v", p)
	}
}

func TestAtPull_GraceWindowBothSides(t *testing.T) {
	cases := []struct {
		startMS int64
		want    bool
	}{{AtPullGraceMS, true}, {AtPullGraceMS + 1, false}, {300, true}}
	for _, c := range cases {
		fight := sheetFight()
		fight.Auras = []summary.AuraTrack{appliedTrack(noFlaskGUID, flaskSpellID, c.startMS)}
		if got := AtPull(fight, noFlaskGUID, Load().AllRoles()).Has(CategoryFlask); got != c.want {
			t.Errorf("flask applied at %d ms: Has = %v, want %v", c.startMS, got, c.want)
		}
	}
}

func TestAtPull_PotionAppliedInGraceIsNotPresent(t *testing.T) {
	// a potion drunk on the pull is a use (F6's rule), not something brought in
	fight := sheetFight()
	fight.Auras = append(fight.Auras, appliedTrack(noFlaskGUID, potionSpellID, 500))
	if p := AtPull(fight, noFlaskGUID, Load().AllRoles()); p.Has(CategoryCombatPotion) {
		t.Fatalf("potion inside the grace window: %+v", p)
	}
}

func TestRaidBuffsAtPull_GraceWindowBothSides(t *testing.T) {
	const battleShoutR7 = 25289
	for _, c := range []struct {
		startMS int64
		want    bool
	}{{AtPullGraceMS, true}, {AtPullGraceMS + 1, false}} {
		fight := sheetFight()
		fight.Auras = []summary.AuraTrack{appliedTrack(noFlaskGUID, battleShoutR7, c.startMS)}
		for _, s := range RaidBuffsAtPull(fight, noFlaskGUID, Load().RaidBuffs) {
			if s.Buff.Name == "Battle Shout" && s.Present != c.want {
				t.Errorf("Battle Shout at %d ms: Present = %v, want %v", c.startMS, s.Present, c.want)
			}
		}
	}
}

func TestRaidBuffsAtPull_NoSheetIsUnknown(t *testing.T) {
	fight := summary.Summary{Roster: []summary.RosterRow{{GUID: "P", Class: "Priest"}},
		Auras: []summary.AuraTrack{seededTrack("P", 10938)}}
	for _, s := range RaidBuffsAtPull(fight, "P", Load().RaidBuffs) {
		if s.Known || s.Present {
			t.Errorf("%s with no sheet: %+v, want unknown and not present", s.Buff.Name, s)
		}
	}
}

// TestAtPull_KnownGoodForeverPull is the §13.5 / §20.6 known-good fixture: a
// real Forever boss pull with COMBATANT_INFO, a flask holder and a player
// without one, a Fury warrior and an avoidable death. When it is on disk, cut
// it to logs/engine/event/testdata/forever-boss-pull.log, add its
// goldenDialects row, and pin both flask answers here (this is also where
// the snapshot's timing against ENCOUNTER_START is first seen on Forever).
func TestAtPull_KnownGoodForeverPull(t *testing.T) {
	t.Skip("owner-provided Forever boss pull not yet on disk")
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
	fight := sheetFight()
	fight.Auras = []summary.AuraTrack{seededTrack(noFlaskGUID, fortitudeR6)}
	for _, s := range RaidBuffsAtPull(fight, noFlaskGUID, Load().RaidBuffs) {
		if !s.Known {
			t.Errorf("%s: a player with a sheet is known", s.Buff.Name)
		}
		if want := s.Buff.Name == "Power Word: Fortitude"; s.Present != want {
			t.Errorf("%s present = %v, want %v", s.Buff.Name, s.Present, want)
		}
	}
}
