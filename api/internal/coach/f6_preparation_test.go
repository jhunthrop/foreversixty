package coach

import (
	"maps"
	"slices"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

func fullKit() seeds {
	return seededAt(flask("Flask of the Titans"), food("Grilled Squid"))
}

func TestPreparation_NoFlaskAt40IsMajor(t *testing.T) {
	in := furyPrep(t, 40, seededAt(food("Grilled Squid")))
	f := preparation(in, CountedTime(in)).Finding
	if f == nil || f.Title != "No flask at the pull" || f.Severity != SeverityMajor || len(f.Anchors) == 0 {
		t.Fatalf("%+v", f)
	}
	if !strings.Contains(f.Detail, "Grilled Squid") || !strings.HasSuffix(f.Detail, "and no flask.") {
		t.Fatal(f.Detail)
	}
}

func TestPreparation_SeverityBySize(t *testing.T) {
	for _, size := range []int{10, 20} {
		in := furyPrep(t, size, seededAt(food("Grilled Squid")))
		if f := preparation(in, CountedTime(in)).Finding; f == nil || f.Severity != SeverityMinor {
			t.Fatalf("size %d: %+v", size, f)
		}
	}
}

func TestPreparation_MissingWeaponEnchantMajorAtEverySize(t *testing.T) {
	for _, size := range []int{10, 20, 40} {
		in := withUnenchantedMainHand(furyPrep(t, size, fullKit()))
		f := preparation(in, CountedTime(in)).Finding
		if f == nil || f.Severity != SeverityMajor || len(f.Anchors) == 0 {
			t.Fatalf("size %d: %+v", size, f)
		}
	}
}

func TestPreparation_NoCharacterSheetGoesToCouldntCheck(t *testing.T) {
	in := withCombatantInput(furyPrep(t, 40, nil), furyGUID, false)
	o := preparation(in, CountedTime(in))
	if o.Finding != nil {
		t.Fatalf("no sheet, no claim: %+v", o.Finding)
	}
	if !slices.Contains(o.CouldNotCheck, "Consumables and enchants: your log has no character sheet for this pull.") {
		t.Fatal(o.CouldNotCheck)
	}
}

func TestPreparation_FoodListIncompleteIsNeverMissing(t *testing.T) {
	in := furyPrep(t, 40, seededAt(flask("Flask of the Titans")))
	o := preparation(in, CountedTime(in))
	if o.Finding != nil && strings.Contains(strings.ToLower(o.Finding.Title+o.Finding.Detail), "food") {
		t.Fatalf("an incomplete list never says food is missing: %+v", o.Finding)
	}
	if !slices.Contains(o.CouldNotCheck, "Food: the coach's food list isn't complete yet.") {
		t.Fatal(o.CouldNotCheck)
	}
}

func TestPreparation_RaidBuffNeedsProviderInRoster(t *testing.T) {
	in := tenManNoPriest(t)
	o := preparation(in, CountedTime(in))
	all := textOf(o)
	if strings.Contains(all, "Power Word: Fortitude") {
		t.Fatal("no priest was there")
	}
	if o.Finding == nil || !strings.Contains(all, "Arcane Intellect") || o.Finding.Severity != SeverityMinor {
		t.Fatalf("a mage was there: %+v", o.Finding)
	}
}

func TestPreparation_WisdomIsNothingToARageUser(t *testing.T) {
	in := rosterOf(furyPrep(t, 40, fullKit()), rosterClass("pal", "Paladin"), rosterClass("mage", "Mage"))
	in = withRaidBuff(t, in, "Blessing of Kings")
	if o := preparation(in, CountedTime(in)); o.Finding != nil {
		t.Fatalf("a warrior needs neither Wisdom nor Intellect: %+v", o.Finding)
	}
}

func TestPreparation_OnePaladinCoversOneBlessingWhicheverItWas(t *testing.T) {
	for _, held := range []string{"Blessing of Kings", "Blessing of Might"} {
		in := withRaidBuff(t, rosterOf(furyPrep(t, 40, fullKit()), rosterClass("pal", "Paladin")), held)
		if o := preparation(in, CountedTime(in)); o.Finding != nil {
			t.Fatalf("one paladin gave %s; the log cannot say Kings was owed: %+v", held, o.Finding)
		}
	}
}

func TestPreparation_NoBlessingAtAllWithAPaladinNamesNone(t *testing.T) {
	in := rosterOf(furyPrep(t, 40, fullKit()), rosterClass("pal", "Paladin"))
	f := preparation(in, CountedTime(in)).Finding
	if f == nil || f.Title != "No blessing at the pull" {
		t.Fatalf("%+v", f)
	}
}

func TestPreparation_MorePaladinsThanBlessingsHeldNamesTheOneMissing(t *testing.T) {
	in := withRaidBuff(t, rosterOf(furyPrep(t, 40, fullKit()), rosterClass("pal", "Paladin"), rosterClass("pal2", "Paladin")), "Blessing of Might")
	f := preparation(in, CountedTime(in)).Finding
	if f == nil || !strings.Contains(f.Title, "Blessing of Kings") {
		t.Fatalf("two paladins, one blessing held, one absent: %+v", f)
	}
}

func TestPreparation_ADpsWithoutTheListedPotionIsNeverTold0Used(t *testing.T) {
	in := furyPrep(t, 40, fullKit())
	o := preparation(in, CountedTime(in))
	if o.Finding != nil || strings.Contains(textOf(o), "0 used") {
		t.Fatalf("%+v", o.Finding)
	}
	if !slices.Contains(o.CouldNotCheck, "Combat potions: the coach's potion list isn't complete yet.") {
		t.Fatal(o.CouldNotCheck)
	}
}

// withCompletePotions is in as if the catalogue's potion list were complete.
func withCompletePotions(in Input) Input {
	done := maps.Clone(in.Tables.Consumables.Complete)
	done["combat_potion"] = true
	in.Tables.Consumables.Complete = done
	return in
}

func TestPreparation_PotionsShortOfWhatTheFightHoldsAreMinorAndQuoteBothNumbers(t *testing.T) {
	in := withCompletePotions(furyPrep(t, 40, fullKit()))
	in.Fight.DurationMS = 300_000
	in.Fight = seedPotion(in.Fight, in.Row.GUID)
	o := preparation(in, CountedTime(in))
	f := o.Finding
	if f == nil || f.Severity != SeverityMinor || f.Title != "Combat potions: 1 unused" || f.Detail != "You used 1 combat potion. The fight had room for 2." {
		t.Fatalf("%+v", f)
	}
	if !o.UnusedPotions || o.MissingFlask {
		t.Errorf("typed facts for the links: unused potions %v, missing flask %v", o.UnusedPotions, o.MissingFlask)
	}
	if len(f.Facts.Numbers) != 3 || f.Facts.Numbers[1].Text != "2" || f.Facts.Numbers[2].Text != "1" {
		t.Fatalf("the 'of 2' is a fact too: %+v", f.Facts.Numbers)
	}
}

func TestPreparation_AShortFightIsNotFaultedForTheSecondPotion(t *testing.T) {
	in := withCompletePotions(furyPrep(t, 40, fullKit()))
	in.Fight = seedPotion(in.Fight, in.Row.GUID) // 60 s: one potion is all it could hold
	if o := preparation(in, CountedTime(in)); o.Finding != nil {
		t.Fatalf("%+v", o.Finding)
	}
}

func TestPreparation_WeaponOilIsCouldntCheckNeverMissing(t *testing.T) {
	in := furyPrep(t, 40, fullKit())
	o := preparation(in, CountedTime(in))
	if o.Finding != nil || !slices.Contains(o.CouldNotCheck, "Weapon oils and stones: the coach can't read them from your log yet.") {
		t.Fatalf("%+v %v", o.Finding, o.CouldNotCheck)
	}
}

func TestPreparation_AFlaskAppliedWithinTheGraceCountsAndOneAfterItDoesNot(t *testing.T) {
	fl := flask("Flask of the Titans")
	for _, c := range []struct {
		startMS int64
		missing bool
	}{{9_000, false}, {12_000, true}} {
		in := furyPrep(t, 40, seededAt(food("Grilled Squid")))
		in.Fight = withAura(in.Fight, in.Row.GUID, fl.SpellID, fl.Name, []string{in.Row.GUID}, span(c.startMS, in.Fight.DurationMS))
		f := preparation(in, CountedTime(in)).Finding
		if got := f != nil && f.Title == "No flask at the pull"; got != c.missing {
			t.Fatalf("applied at %d ms: %+v", c.startMS, f)
		}
	}
}

func TestPreparation_EveryEnchantSlotHasALabel(t *testing.T) {
	for _, slot := range rating.DisciplineEnchantSlots {
		if enchantSlotNames[slot] == "" {
			t.Fatalf("slot %d has no label", slot)
		}
	}
	if !slices.Contains(rating.DisciplineEnchantSlots, mainHandSlot) {
		t.Fatal("the main hand is no longer an enchant slot")
	}
}

func TestPreparation_WorldBuffsAreFactsNotFindings(t *testing.T) {
	in := furyPrep(t, 40, fullKit())
	if o := preparation(in, CountedTime(in)); o.Finding != nil {
		t.Fatalf("no world buff is no finding: %+v", o.Finding)
	}
	wb := catalogueWorldBuff(t)
	in.Fight = withAura(in.Fight, in.Row.GUID, wb.SpellID, wb.Name, []string{in.Row.GUID}, span(0, in.Fight.DurationMS))
	o := preparation(in, CountedTime(in))
	if o.Finding != nil || len(o.Strengths) != 1 || !slices.Contains(o.Strengths[0].Facts.Items, wb.Name) {
		t.Fatalf("a world buff is a fact: %+v", o)
	}
}

func TestPreparation_EverythingPresentIsAStrength(t *testing.T) {
	in := furyPrep(t, 40, fullKit())
	o := preparation(in, CountedTime(in))
	if o.Finding != nil || len(o.Strengths) != 1 || o.Strengths[0].Sentence != "Flask of the Titans at the pull." {
		t.Fatalf("%+v", o)
	}
	if len(o.Strengths[0].Anchors) == 0 || len(o.Anchors) != 1 || o.Anchors[0].Kind != AnchorBuffs {
		t.Fatal("a strength anchors to the Buffs tab")
	}
}

func TestPreparation_KnownGoodFlaskOnTheRealPull(t *testing.T) {
	t.Skip("known-good boss pull is owner-provided; none on disk yet")
}

// §8.1: one slot names the slot; two or more count them and list the slots.
func TestPreparation_EnchantCopy(t *testing.T) {
	in := withUnenchantedMainHand(furyPrep(t, 40, fullKit()))
	f := preparation(in, CountedTime(in)).Finding
	if f == nil || f.Title != "Main Hand has no enchant" || f.Detail != "Main Hand had no permanent enchant at the pull." || f.Basis.Label != BasisGear {
		t.Fatalf("%+v", f)
	}
	gear := preparedGear(false)
	gear[slotChest].Enchants = nil
	in = withGear(furyPrep(t, 40, fullKit()), gear)
	f = preparation(in, CountedTime(in)).Finding
	if f == nil || f.Title != "2 items have no enchant" || f.Detail != "No permanent enchant on Main Hand and Chest." {
		t.Fatalf("%+v", f)
	}
	assertNumbers(t, f.Facts, "2")
}

// §8.1: the lead names the title; everything else missing follows in one sentence.
func TestPreparation_AlsoMissingFollowsTheLead(t *testing.T) {
	in := withUnenchantedMainHand(furyPrep(t, 40, seededAt(food("Grilled Squid"))))
	f := preparation(in, CountedTime(in)).Finding
	want := "You pulled with Grilled Squid, and no flask. Also missing: a Main Hand enchant."
	if f == nil || f.Title != "No flask at the pull" || f.Detail != want {
		t.Fatalf("got %+v, want detail %q", f, want)
	}
	bare := furyPrep(t, 40, seededAt())
	if f := preparation(bare, CountedTime(bare)).Finding; f == nil || !strings.HasPrefix(f.Detail, "You pulled with no flask.") {
		t.Fatalf("%+v", f)
	}
}

func TestPreparation_RaidBuffCopy(t *testing.T) {
	in := tenManNoPriest(t)
	f := preparation(in, CountedTime(in)).Finding
	want := "Arcane Intellect was not on you at the pull. A mage in the raid provides it."
	if f == nil || f.Title != "No Arcane Intellect at the pull" || f.Detail != want || f.Basis.Label != BasisRaidBuff {
		t.Fatalf("got %+v, want detail %q", f, want)
	}
	blessing := rosterOf(furyPrep(t, 40, fullKit()), rosterClass("pal", "Paladin"))
	f = preparation(blessing, CountedTime(blessing)).Finding
	if f == nil || f.Detail != "A blessing was not on you at the pull. A paladin in the raid provides it." {
		t.Fatalf("%+v", f)
	}
}

// A buff whose only provider in the roster is the player is F2's uptime, never F6 (A11):
// a lone warrior without Battle Shout, a lone mage without Arcane Intellect.
func TestPreparation_ThePlayerIsNeverTheirOwnRaidBuffProvider(t *testing.T) {
	in := furyPrep(t, 40, fullKit())
	in.Fight = withRoster(in.Fight, in.Row, rosterClass("rogue", "Rogue"))
	in.Fight.Auras = slices.DeleteFunc(slices.Clone(in.Fight.Auras), func(a summary.AuraTrack) bool {
		return slices.Contains(raidBuffByName(t, in, "Battle Shout").SpellIDs, a.SpellID)
	})
	if o := preparation(in, CountedTime(in)); strings.Contains(textOf(o), "Battle Shout") {
		t.Fatalf("the only warrior is the player: %+v", o.Finding)
	}
	in = rosterOf(in, rosterClass("war2", "Warrior"))
	if o := preparation(in, CountedTime(in)); !strings.Contains(textOf(o), "Battle Shout") {
		t.Fatalf("another warrior was there: %+v", o.Finding)
	}
}

func TestPreparation_ALonePaladinOwesNoBlessing(t *testing.T) {
	in := furyPrep(t, 40, fullKit())
	in.Row = withRole(in.Fight, in.Row.GUID, "Paladin", "Retribution", RoleDPS).Roster[0]
	in.Fight = withRoster(in.Fight, in.Row)
	if o := preparation(in, CountedTime(in)); strings.Contains(strings.ToLower(textOf(o)), "blessing") {
		t.Fatalf("the only paladin is the player: %+v", o.Finding)
	}
}
