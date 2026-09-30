package main

import (
	"encoding/json"
	"path/filepath"
	"testing"
)

func TestCollectSpellIDsFindsBareAndRankedShapes(t *testing.T) {
	var doc any = map[string]any{
		"prepullActions": []any{
			map[string]any{"action": map[string]any{"castSpell": map[string]any{
				"spellId": map[string]any{"spellId": float64(11269), "rank": float64(6)},
			}}},
		},
		"priorityList": []any{
			map[string]any{"action": map[string]any{"castSpell": map[string]any{
				"spellId": map[string]any{"spellId": float64(14278)},
			}}},
			map[string]any{"condition": map[string]any{"cmp": map[string]any{
				"lhs": map[string]any{"currentComboPoints": map[string]any{}},
			}}},
		},
	}
	got := map[int]bool{}
	collectSpellIDs(doc, got)
	want := map[int]bool{11269: true, 14278: true}
	if len(got) != len(want) {
		t.Fatalf("collectSpellIDs = %v, want %v", got, want)
	}
	for id := range want {
		if !got[id] {
			t.Errorf("collectSpellIDs missing id %d: %v", id, got)
		}
	}
}

func writeAPLFixture(t *testing.T, dir, spec string, spellIDs ...int) {
	t.Helper()
	actions := make([]any, len(spellIDs))
	for i, id := range spellIDs {
		actions[i] = map[string]any{"action": map[string]any{"castSpell": map[string]any{"spellId": map[string]any{"spellId": id}}}}
	}
	doc := map[string]any{
		"spec":  spec,
		"state": "written",
		"rotation": map[string]any{
			"type":         "TypeAPL",
			"priorityList": actions,
		},
	}
	b, err := json.Marshal(doc)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, "data", "curated", "apl", spec+".json")
	if err := writeFile(t, path, string(b)); err != nil {
		t.Fatal(err)
	}
}

// This lane's brief, item 5: rogue-subtlety's own rotation casts Ambush
// (spell id 11269 at rank 6, matching data/curated/apl/rogue-
// subtlety.json's real prepull action) - a dagger-only Classic ability
// - so its rotation must be found to require a dagger.
func TestAPLRotationRequiresDaggerTrueForAmbush(t *testing.T) {
	dir := t.TempDir()
	writeAPLFixture(t, dir, "rogue-subtlety", 11269, 14278)
	got, err := aplRotationRequiresDagger(dir, "rogue-subtlety")
	if err != nil {
		t.Fatalf("aplRotationRequiresDagger: %v", err)
	}
	if !got {
		t.Error("aplRotationRequiresDagger = false, want true: this rotation casts Ambush (id 11269)")
	}
}

// A rotation that never names Backstab or Ambush at any rank (Ghostly
// Strike, id 14278, is neither) must not be flagged - this is the
// negative control every other rogue spec, and every other class
// entirely, relies on to stay unrestricted.
func TestAPLRotationRequiresDaggerFalseWithNeitherAbility(t *testing.T) {
	dir := t.TempDir()
	writeAPLFixture(t, dir, "rogue-combat", 14278, 1752)
	got, err := aplRotationRequiresDagger(dir, "rogue-combat")
	if err != nil {
		t.Fatalf("aplRotationRequiresDagger: %v", err)
	}
	if got {
		t.Error("aplRotationRequiresDagger = true, want false: this rotation names neither Backstab nor Ambush")
	}
}

// A spec with no apl file yet (not curated - writtenSpecs' own "skipped,
// not an error" reading) must not fail runSpec over a requirement that
// does not apply to it yet.
func TestAPLRotationRequiresDaggerMissingFileIsFalseNotError(t *testing.T) {
	dir := t.TempDir()
	got, err := aplRotationRequiresDagger(dir, "nonexistent-spec")
	if err != nil {
		t.Fatalf("aplRotationRequiresDagger for a missing apl file: want no error, got %v", err)
	}
	if got {
		t.Error("aplRotationRequiresDagger for a missing apl file: want false")
	}
}

// restrictToDaggers's own contract, this lane's brief item 5: Ardent
// Custodian (a real Mace1H, subclass 4) is dropped; a real dagger and a
// non-weapon (should never appear in main_hand/off_hand, but left alone
// rather than assumed impossible) both pass through.
func TestRestrictToDaggersDropsNonDaggerWeaponsOnly(t *testing.T) {
	mace := scored{candidate: candidate{ID: 868, Name: "Ardent Custodian", ClassID: itemClassWeapon, SubclassID: 4}}
	dagger := scored{candidate: candidate{ID: 1, Name: "A Dagger", ClassID: itemClassWeapon, SubclassID: daggerSubclassID}}
	nonWeapon := scored{candidate: candidate{ID: 2, Name: "Not A Weapon", ClassID: armorClassID}}
	got := restrictToDaggers([]scored{mace, dagger, nonWeapon})
	if len(got) != 2 {
		t.Fatalf("restrictToDaggers = %+v, want 2 entries (the mace dropped)", got)
	}
	for _, c := range got {
		if c.ID == 868 {
			t.Fatalf("restrictToDaggers kept the mace (id 868): %+v", got)
		}
	}
}

// This lane's brief (bis-ranker-integrity-6), item 9: a caster's ranged
// slot is only ever a real wand - score()'s own Shoot fallback
// (score.go) had nothing checking that before this fix, so a thrown
// weapon or a bow in the ranged slot scored identically to a real
// wand. (Not Torch of Light/Cold Snap: the controller's own direct
// note corrected this lane's first pass - both are real wands, not
// thrown weapons, so fictional ids are used here instead of asserting
// a false fact about a real item.)
func containsID(list []scored, id int) bool {
	for _, c := range list {
		if c.ID == id {
			return true
		}
	}
	return false
}

func TestRestrictRangedByProficiencyCasterKeepsOnlyWands(t *testing.T) {
	wand := scored{candidate: candidate{ID: 1, Name: "A Wand", ClassID: itemClassWeapon, WeaponType: "wand"}}
	thrown := scored{candidate: candidate{ID: 2, Name: "A Thrown Weapon", ClassID: itemClassWeapon, WeaponType: "thrown"}}
	bow := scored{candidate: candidate{ID: 3, Name: "A Bow", ClassID: itemClassWeapon, WeaponType: "bow"}}
	unknown := scored{candidate: candidate{ID: 4, Name: "A Ranged Weapon With No weapon_type Yet", ClassID: itemClassWeapon, WeaponType: ""}}
	for _, casterClass := range []string{"mage", "priest", "warlock"} {
		got := restrictRangedByProficiency([]scored{wand, thrown, bow, unknown}, casterClass)
		if len(got) != 1 || got[0].ID != 1 {
			t.Fatalf("%s: restrictRangedByProficiency = %+v, want only the wand (id 1)", casterClass, got)
		}
	}
}

// A relic (paladin/shaman/druid's own ranged slot, ClassID armorClassID)
// is never this gate's concern, for any class.
func TestRestrictRangedByProficiencyNeverTouchesRelics(t *testing.T) {
	relic := scored{candidate: candidate{ID: 3, Name: "A Libram", ClassID: armorClassID, SubclassID: armorSublibramID}}
	for _, classSlug := range []string{"mage", "hunter", "paladin", "shaman", "druid"} {
		got := restrictRangedByProficiency([]scored{relic}, classSlug)
		if !containsID(got, 3) {
			t.Fatalf("%s: restrictRangedByProficiency dropped a relic: %+v", classSlug, got)
		}
	}
}

// Hunter/warrior/rogue keep every recognised ranged weapon type per
// this lane's brief's own list (bows/guns/crossbows/thrown), and lose
// a wand - the one type they have no proficiency for.
func TestRestrictRangedByProficiencySkillClassesKeepEveryRecognisedType(t *testing.T) {
	bow := scored{candidate: candidate{ID: 1, Name: "A Bow", ClassID: itemClassWeapon, WeaponType: "bow"}}
	gun := scored{candidate: candidate{ID: 2, Name: "A Gun", ClassID: itemClassWeapon, WeaponType: "gun"}}
	crossbow := scored{candidate: candidate{ID: 3, Name: "A Crossbow", ClassID: itemClassWeapon, WeaponType: "crossbow"}}
	thrown := scored{candidate: candidate{ID: 4, Name: "A Thrown Weapon", ClassID: itemClassWeapon, WeaponType: "thrown"}}
	wand := scored{candidate: candidate{ID: 5, Name: "A Wand", ClassID: itemClassWeapon, WeaponType: "wand"}}
	for _, classSlug := range []string{"hunter", "warrior", "rogue"} {
		got := restrictRangedByProficiency([]scored{bow, gun, crossbow, thrown, wand}, classSlug)
		if len(got) != 4 {
			t.Fatalf("%s: restrictRangedByProficiency = %+v, want bow/gun/crossbow/thrown kept, wand dropped", classSlug, got)
		}
		if containsID(got, 5) {
			t.Fatalf("%s: restrictRangedByProficiency kept a wand (id 5), no proficiency for one: %+v", classSlug, got)
		}
	}
}

// This lane's brief's own scoping: "until the data carries it... exclude
// it from CASTER ranged slots" - a hunter/warrior/rogue's own real bow/
// gun/crossbow/thrown item with an empty WeaponType (the data gap this
// lane's brief describes as almost every ranged row's current state)
// must NOT be excluded, or every one of those classes' ranged slots
// would empty over the identical gap the caster fix targets.
func TestRestrictRangedByProficiencySkillClassesGrandfatherUnknownType(t *testing.T) {
	unknownType := scored{candidate: candidate{ID: 6, Name: "A Real Bow With No weapon_type Yet", ClassID: itemClassWeapon, WeaponType: ""}}
	for _, classSlug := range []string{"hunter", "warrior", "rogue"} {
		got := restrictRangedByProficiency([]scored{unknownType}, classSlug)
		if !containsID(got, 6) {
			t.Fatalf("%s: restrictRangedByProficiency dropped an unknown-type item, want it grandfathered in: %+v", classSlug, got)
		}
	}
}

// A weapon-class ranged candidate for a relic-only class (paladin/
// shaman/druid) is a data anomaly this gate excludes outright - their
// own per-class item file was never going to hand it one for real
// (rangedWeaponSkillClasses' own doc).
func TestRestrictRangedByProficiencyExcludesWeaponClassForRelicOnlyClasses(t *testing.T) {
	bow := scored{candidate: candidate{ID: 7, Name: "A Bow", ClassID: itemClassWeapon, WeaponType: "bow"}}
	for _, classSlug := range []string{"paladin", "shaman", "druid"} {
		got := restrictRangedByProficiency([]scored{bow}, classSlug)
		if len(got) != 0 {
			t.Fatalf("%s: restrictRangedByProficiency = %+v, want the weapon-class candidate excluded", classSlug, got)
		}
	}
}
