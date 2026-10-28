package coach

import (
	"slices"
	"strings"
	"testing"
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

func TestPreparation_OnePaladinGivesOneBlessing(t *testing.T) {
	in := rosterOf(furyPrep(t, 40, fullKit()), rosterClass("pal", "Paladin"))
	in = withRaidBuff(t, in, "Blessing of Kings")
	if o := preparation(in, CountedTime(in)); o.Finding != nil {
		t.Fatalf("one paladin gave Kings and cannot also give Might: %+v", o.Finding)
	}
	two := rosterOf(in, rosterClass("pal", "Paladin"), rosterClass("pal2", "Paladin"))
	f := preparation(two, CountedTime(two)).Finding
	if f == nil || !strings.Contains(f.Title, "Blessing of Might") {
		t.Fatalf("two paladins cover Kings and Might, so Might is owed: %+v", f)
	}
}

func TestPreparation_PotionsShortOfMaxUsesAreMinor(t *testing.T) {
	in := furyPrep(t, 40, fullKit())
	in.Fight.DurationMS = 300_000
	f := preparation(in, CountedTime(in)).Finding
	if f == nil || f.Severity != SeverityMinor || !strings.HasPrefix(f.Title, "Combat potions") {
		t.Fatalf("%+v", f)
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
