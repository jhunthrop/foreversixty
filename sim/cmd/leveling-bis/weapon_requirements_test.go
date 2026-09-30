package main

import (
	"encoding/json"
	"path/filepath"
	"strings"
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

// TestRestrictToProficientWeaponsDropsIllegalSubclassOnly is this
// lane's brief (bis-ranker-integrity-10), item 3's own repro: a
// paladin's allow-list (classWeaponSubclasses["paladin"]) carries no
// axe subclass (0 one-hand, 1 two-hand) - Nightfall 19169, the exact
// item data-followups-5's own brief names, is a 2H axe and must be
// dropped; a legal mace (subclass 4) and a non-weapon candidate (a
// relic, which this gate never touches) both pass through.
func TestRestrictToProficientWeaponsDropsIllegalSubclassOnly(t *testing.T) {
	axe := scored{candidate: candidate{ID: 19169, Name: "Nightfall", ClassID: itemClassWeapon, SubclassID: 1}}
	mace := scored{candidate: candidate{ID: 1, Name: "A Mace", ClassID: itemClassWeapon, SubclassID: 4}}
	relic := scored{candidate: candidate{ID: 2, Name: "A Libram", ClassID: armorClassID, SubclassID: armorSublibramID}}
	got := restrictToProficientWeapons([]scored{axe, mace, relic}, classWeaponSubclasses["paladin"])
	if len(got) != 2 {
		t.Fatalf("restrictToProficientWeapons = %+v, want 2 entries (the axe dropped)", got)
	}
	if containsID(got, 19169) {
		t.Fatalf("restrictToProficientWeapons kept the axe (id 19169): %+v", got)
	}
	if !containsID(got, 1) || !containsID(got, 2) {
		t.Fatalf("restrictToProficientWeapons dropped the mace or the relic: %+v", got)
	}
}

// TestClassWeaponSubclassesMatchesForeverDocumentedFacts is this
// lane's brief, item 3: paladin has no axe, druid no polearm (both
// Classic 1.x facts this static fallback must get right, and the
// exact regression data-followups-5's own brief names), and shaman/
// rogue carry Forever's two documented deviations
// (.claude/agents/wow-player.md: "shamans train one- and two-handed
// axes and maces, rogues train maces") on top of the Classic 1.x
// baseline.
func TestClassWeaponSubclassesMatchesForeverDocumentedFacts(t *testing.T) {
	const (
		axe1H   = 0
		axe2H   = 1
		mace1H  = 4
		mace2H  = 5
		polearm = 6
	)
	if classWeaponSubclasses["paladin"][axe1H] || classWeaponSubclasses["paladin"][axe2H] {
		t.Error(`classWeaponSubclasses["paladin"] carries an axe subclass, want none (Classic 1.x: maces/polearm/swords only)`)
	}
	if classWeaponSubclasses["druid"][polearm] {
		t.Error(`classWeaponSubclasses["druid"] carries polearm, want none (Classic 1.x: daggers/fist/maces/staves only)`)
	}
	if !classWeaponSubclasses["shaman"][axe1H] || !classWeaponSubclasses["shaman"][axe2H] || !classWeaponSubclasses["shaman"][mace2H] {
		t.Error(`classWeaponSubclasses["shaman"] missing an axe or two-hand mace subclass, want Forever's documented "one- and two-handed axes and maces"`)
	}
	if !classWeaponSubclasses["rogue"][mace1H] {
		t.Error(`classWeaponSubclasses["rogue"] missing one-hand mace, want Forever's documented "rogues train maces"`)
	}
}

// TestLoadWeaponSubclassesFallsBackToStaticTableWhenNothingPublished
// is this lane's brief, item 3: "if nothing is published yet, read
// data/curated/ for a static table and say so" - as of this lane,
// neither data/builds/<build>/proficiency.json nor a data/curated/
// classes.json weapon-proficiency field exists, so a buildDir with no
// proficiency.json at all must fall back to classWeaponSubclasses and
// say so in its returned source string.
func TestLoadWeaponSubclassesFallsBackToStaticTableWhenNothingPublished(t *testing.T) {
	dir := t.TempDir()
	got, source, err := loadWeaponSubclasses(dir, "paladin")
	if err != nil {
		t.Fatalf("loadWeaponSubclasses: %v", err)
	}
	if got[0] {
		t.Error(`loadWeaponSubclasses("paladin") carries axe subclass 0, want the static fallback's own paladin table (no axes)`)
	}
	if !strings.Contains(source, "static fallback") {
		t.Errorf("source = %q, want it to say this is a static fallback", source)
	}
}

// TestLoadWeaponSubclassesPrefersAPublishedFile confirms the other
// side of the same preference order: a data/builds/<build>/
// proficiency.json this lane's own brief anticipates data-followups-5
// (or a later lane) will publish wins over the static fallback the
// moment it exists, without requiring a code change here to read it.
func TestLoadWeaponSubclassesPrefersAPublishedFile(t *testing.T) {
	dir := t.TempDir()
	if err := writeFile(t, filepath.Join(dir, "proficiency.json"), `{"paladin": [4, 5, 6, 7, 8, 99]}`); err != nil {
		t.Fatal(err)
	}
	got, source, err := loadWeaponSubclasses(dir, "paladin")
	if err != nil {
		t.Fatalf("loadWeaponSubclasses: %v", err)
	}
	// 99 is not a real subclass id; asserting it came through anyway
	// (rather than only checking the real ones) is what proves this
	// read the published file and not classWeaponSubclasses, which
	// carries no such entry.
	if !got[99] || len(got) != 6 {
		t.Fatalf("loadWeaponSubclasses = %v, want exactly the published file's own 6 entries including the sentinel 99", got)
	}
	if source != filepath.Join(dir, "proficiency.json") {
		t.Errorf("source = %q, want the published file's own path", source)
	}
}

// TestLoadWeaponSubclassesPublishedFileMissingClassEntryErrors: a
// published proficiency.json that exists but has nothing for this
// class is a data bug worth failing loudly over (loadWeaponSubclasses'
// own doc: "not on any decode error against a file that DOES exist"),
// not a silent fall-through to the static table.
func TestLoadWeaponSubclassesPublishedFileMissingClassEntryErrors(t *testing.T) {
	dir := t.TempDir()
	if err := writeFile(t, filepath.Join(dir, "proficiency.json"), `{"warrior": [0, 1]}`); err != nil {
		t.Fatal(err)
	}
	if _, _, err := loadWeaponSubclasses(dir, "paladin"); err == nil {
		t.Fatal("loadWeaponSubclasses: want an error, the published file exists but carries no paladin entry")
	}
}

// TestLoadWeaponSubclassesUnknownClassWithNoFallbackErrors: a class
// slug neither a published file nor classWeaponSubclasses has ever
// heard of must fail loudly, not silently return an empty (therefore
// "no weapon is ever legal") allow-list.
func TestLoadWeaponSubclassesUnknownClassWithNoFallbackErrors(t *testing.T) {
	dir := t.TempDir()
	if _, _, err := loadWeaponSubclasses(dir, "not-a-real-class"); err == nil {
		t.Fatal("loadWeaponSubclasses: want an error for an unknown class with no static fallback entry")
	}
}
