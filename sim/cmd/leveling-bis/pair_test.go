package main

import "testing"

// Round 9: the battleground and Darkspear trinkets are "Unique-Equipped"
// by family. The client puts Rune of Perfection and Rune of Duty in
// ItemLimitCategory 714 "Rune of Battle", Darkspear Voodoo Seal and
// Relentless Raider's Seal in 715 "Seal of Battle", Talisman of Arathor and
// Defiler's Talisman in 716 and Arena Master and Arena Grand Master in 718,
// each with Quantity 1, so one character wears one of each family. The
// ranker paired them by id and name only and, with the gates open, put both
// Runes on a band-20 bear and both Seals on a band-50 bear.
func TestOneLimitCategoryFillsOnlyOneOfAPair(t *testing.T) {
	perfection := scored{candidate: candidate{ID: 21565, Name: "Rune of Perfection", LimitCategory: 714, Slots: []string{"trinket1", "trinket2"}}, Score: 9}
	duty := scored{candidate: candidate{ID: 21567, Name: "Rune of Duty", LimitCategory: 714, Slots: []string{"trinket1", "trinket2"}}, Score: 8}
	other := scored{candidate: candidate{ID: 1, Name: "Other", Stats: map[string]float64{"stamina": 5}, Slots: []string{"trinket1", "trinket2"}}, Score: 1}
	if got := excludePaired([]scored{duty, other}, mateOf(&perfection)); len(got) != 1 || got[0].ID != 1 {
		t.Errorf("excludePaired kept %v, want only the item outside Rune of Battle", got)
	}
	picks := map[string]slotPick{"trinket1": {Item: &perfection}}
	if !wouldDuplicatePairMate(picks, "trinket2", duty.candidate) {
		t.Error("Rune of Duty beside Rune of Perfection is allowed; want refused (one Rune of Battle)")
	}
	if wouldDuplicatePairMate(picks, "trinket2", other.candidate) {
		t.Error("an unrelated trinket is refused")
	}
	if got := trinketShortlist([]scored{duty, other}, mateOf(&perfection), map[string]float64{"stamina": 1}); len(got) != 1 || got[0].ID != 1 {
		t.Errorf("trinketShortlist = %v, want only the item outside Rune of Battle", got)
	}
	// No category is no conflict: two unrelated items with LimitCategory 0.
	if mateOf(&other).excludes(candidate{ID: 2, Name: "Another"}) {
		t.Error("two items with no limit category conflict")
	}
}

func TestTheClassFileCarriesTheLimitCategory(t *testing.T) {
	dir := t.TempDir()
	if err := writeFile(t, dir+"/items.json", `[{"id": 21565, "name": "Rune of Perfection", "inventory_type": 12}]`); err != nil {
		t.Fatal(err)
	}
	if err := writeFile(t, dir+"/items/druid.json", `{"build": "b", "class_slug": "druid", "items": [{"id": 21565, "name": "Rune of Perfection", "slot": "trinket", "quality": 3, "limit_category": 714}]}`); err != nil {
		t.Fatal(err)
	}
	got, _, err := loadClassCandidates(dir, "druid")
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || got[0].LimitCategory != 714 {
		t.Fatalf("candidates = %+v, want limit category 714", got)
	}
}
