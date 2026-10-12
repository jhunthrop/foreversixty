package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// Task L1.6: an item whose record is client_unconfirmed (its stats are
// 1.12's; the client may never resolve the id) is never a candidate. The
// filter is loadCandidates' own, so every slot ranks, verifies and lists
// runners-up from the confirmed set only.

// unconfirmedHeadsRepo is the shared reporoot fixture with four head
// items added: 1021 (client_unconfirmed, the best on paper) above three
// confirmed heads 1022 > 1023 > 1024, every one dropped by a dungeon boss.
func unconfirmedHeadsRepo(t *testing.T) (repoRoot, buildDir string) {
	t.Helper()
	repoRoot = t.TempDir()
	if err := os.CopyFS(repoRoot, os.DirFS(repoRootFixture)); err != nil {
		t.Fatalf("copying %s into %s: %v", repoRootFixture, repoRoot, err)
	}
	buildDir = filepath.Join(repoRoot, "data", "builds", "testbuild")
	if err := writeFile(t, filepath.Join(buildDir, "items.json"), `[
		{"id": 1001, "name": "Test Helm", "quality": 3, "item_level": 20, "required_level": 10, "class_id": 4, "subclass_id": 2, "inventory_type": 1, "suffixes": [], "faction_restriction": ""},
		{"id": 1002, "name": "Test Bow", "quality": 3, "item_level": 25, "required_level": 15, "class_id": 2, "subclass_id": 2, "inventory_type": 15, "suffixes": [], "faction_restriction": "horde"},
		{"id": 1021, "name": "Unconfirmed Helm", "quality": 3, "item_level": 20, "required_level": 10, "class_id": 4, "subclass_id": 2, "inventory_type": 1, "suffixes": [], "faction_restriction": "", "client_unconfirmed": true},
		{"id": 1022, "name": "Confirmed Helm A", "quality": 3, "item_level": 20, "required_level": 10, "class_id": 4, "subclass_id": 2, "inventory_type": 1, "suffixes": [], "faction_restriction": ""},
		{"id": 1023, "name": "Confirmed Helm B", "quality": 3, "item_level": 20, "required_level": 10, "class_id": 4, "subclass_id": 2, "inventory_type": 1, "suffixes": [], "faction_restriction": ""},
		{"id": 1024, "name": "Confirmed Helm C", "quality": 3, "item_level": 20, "required_level": 10, "class_id": 4, "subclass_id": 2, "inventory_type": 1, "suffixes": [], "faction_restriction": ""}
	]`); err != nil {
		t.Fatal(err)
	}
	if err := writeFile(t, filepath.Join(buildDir, "items", "hunter.json"), `{
		"build": "testbuild", "class_slug": "hunter",
		"items": [
			{"id": 1001, "name": "Test Helm", "slot": "head", "quality": 3, "required_level": 10, "item_level": 20, "armor": 40, "stats": {"agility": 5}, "damage_min": 0, "damage_max": 0, "speed": 0, "dps": 0, "two_hand": false, "effect_text": "", "set_id": null, "unique": false},
			{"id": 1002, "name": "Test Bow", "slot": "ranged", "quality": 3, "required_level": 15, "item_level": 25, "armor": 0, "stats": {"agility": 2}, "damage_min": 10, "damage_max": 20, "speed": 2.8, "dps": 5.3, "two_hand": false, "effect_text": "", "set_id": null, "unique": false, "weapon_type": "bow"},
			{"id": 1021, "name": "Unconfirmed Helm", "slot": "head", "quality": 3, "required_level": 10, "item_level": 20, "armor": 40, "stats": {"agility": 20}, "damage_min": 0, "damage_max": 0, "speed": 0, "dps": 0, "two_hand": false, "effect_text": "", "set_id": null, "unique": false, "client_unconfirmed": true},
			{"id": 1022, "name": "Confirmed Helm A", "slot": "head", "quality": 3, "required_level": 10, "item_level": 20, "armor": 40, "stats": {"agility": 12}, "damage_min": 0, "damage_max": 0, "speed": 0, "dps": 0, "two_hand": false, "effect_text": "", "set_id": null, "unique": false},
			{"id": 1023, "name": "Confirmed Helm B", "slot": "head", "quality": 3, "required_level": 10, "item_level": 20, "armor": 40, "stats": {"agility": 10}, "damage_min": 0, "damage_max": 0, "speed": 0, "dps": 0, "two_hand": false, "effect_text": "", "set_id": null, "unique": false},
			{"id": 1024, "name": "Confirmed Helm C", "slot": "head", "quality": 3, "required_level": 10, "item_level": 20, "armor": 40, "stats": {"agility": 8}, "damage_min": 0, "damage_max": 0, "speed": 0, "dps": 0, "two_hand": false, "effect_text": "", "set_id": null, "unique": false}
		]
	}`); err != nil {
		t.Fatal(err)
	}
	if err := writeFile(t, filepath.Join(buildDir, "loot.json"), `{
		"sources": [
			{"id": "quest-1", "kind": "quest", "name": "A Test Quest", "items": [1001]},
			{"id": "dungeon-1", "kind": "dungeon", "name": "A Test Dungeon", "bosses": [{"name": "Test Boss", "items": [1002, 1021, 1022, 1023, 1024]}]}
		],
		"quests": {
			"1001": [{"quest_id": 1, "name": "A Test Quest", "faction": "both", "min_level": 5, "level": 7, "level_source": "wowhead"}]
		}
	}`); err != nil {
		t.Fatal(err)
	}
	return repoRoot, buildDir
}

func TestLoadCandidatesDropsClientUnconfirmedItemsAndCountsThem(t *testing.T) {
	_, buildDir := unconfirmedHeadsRepo(t)
	items, unconfirmed, _, err := loadCandidates(buildDir, "hunter")
	if err != nil {
		t.Fatalf("loadCandidates: %v", err)
	}
	for _, c := range items {
		if c.ID == 1021 || c.ClientUnconfirmed {
			t.Fatalf("loadCandidates kept the client_unconfirmed item %d (%s)", c.ID, c.Name)
		}
	}
	if unconfirmed != 1 {
		t.Fatalf("unconfirmed = %d, want 1 (item 1021)", unconfirmed)
	}
}

func TestTheLadderPicksTheConfirmedHeadAndListsTwoConfirmedRunnersUp(t *testing.T) {
	repoRoot, buildDir := unconfirmedHeadsRepo(t)
	fake := &fakeEngine{
		DefaultDPS: 500,
		WeightsResult: map[string]api.StatWeight{
			"ranged_attack_power": {Stat: "ranged_attack_power", Weight: 1.0},
			"agility":             {Stat: "agility", Weight: 1.8},
		},
		ReferenceDPSPerPoint: 0.5,
	}
	outDir := t.TempDir()
	if err := runSpec(fake, repoRoot, buildDir, "testbuild", outDir, "hunter-marksmanship", []int{20}, 5, identityTalentLayout); err != nil {
		t.Fatalf("runSpec: %v", err)
	}
	path := filepath.Join(outDir, "hunter-marksmanship.json")
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	// Match the id as a JSON field, not a raw digit run: the report carries
	// wall-clock and DPS floats that may contain "1021".
	if regexp.MustCompile(`"item_id":\s*1021[,}\s]`).Match(raw) || strings.Contains(string(raw), "Unconfirmed Helm") {
		t.Fatalf("the published report names the client_unconfirmed item 1021 somewhere")
	}
	report := readSpecReportForTest(t, path)
	heads := 0
	for _, b := range report.Bands {
		for _, row := range b.Slots {
			if row.Slot != "head" {
				continue
			}
			heads++
			if row.ItemID != 1022 {
				t.Errorf("band %d %s head = %d, want 1022 (the best confirmed head)", b.Band, b.Faction, row.ItemID)
			}
			var got []int
			for _, ru := range row.RunnersUp {
				got = append(got, ru.ItemID)
			}
			if len(got) != 2 || got[0] != 1023 || got[1] != 1024 {
				t.Errorf("band %d %s head runners_up = %v, want [1023 1024]", b.Band, b.Faction, got)
			}
		}
	}
	if heads == 0 {
		t.Fatal("no head row published")
	}
}

// Surfaced by the L1.6 regeneration: with Headmaster's Charge (a
// client_unconfirmed staff) gone, druid-feral-bear's band-60 raid main hand
// became a one-hander and the off hand took White Bone Shredder, an
// off-hand fist weapon a druid cannot dual-wield. A spec outside
// DualWieldSpecs holds only a shield or a held item there.
func TestANonDualWielderNeverTakesAnOffHandWeapon(t *testing.T) {
	pool := []scored{
		{candidate: candidate{ID: 21522, Name: "Shadowsong's Sorrow", Slots: []string{"main_hand"}, ClassID: itemClassWeapon}, Score: 30},
		{candidate: candidate{ID: 11863, Name: "White Bone Shredder", Slots: []string{"off_hand"}, ClassID: itemClassWeapon}, Score: 20},
		{candidate: candidate{ID: 272087, Name: "Tome of the Darkspear Prophecy", Slots: []string{"off_hand"}, ClassID: 4}, Score: 10},
	}
	result := pick("druid-feral-bear", candidatesBySlot(pool))
	if oh := result["off_hand"].Item; oh == nil || oh.ID != 272087 {
		t.Fatalf("off_hand = %+v, want the held item 272087, never the off-hand weapon", oh)
	}
	if ru := result["off_hand"].RunnerUp; ru != nil && ru.ClassID == itemClassWeapon {
		t.Fatalf("off_hand runner-up = %+v, a weapon", ru)
	}
}

// Fix round 1, item 3: the published-file check's active-build branch.
// A bis file of the active build may not carry a client_unconfirmed item;
// one of an older, never re-ranked build may.
func TestThePublishedCheckRefusesAnUnconfirmedItemOnlyInTheActiveBuild(t *testing.T) {
	items := []candidate{
		{ID: 1021, Name: "Unconfirmed Helm", ClientUnconfirmed: true},
		{ID: 1022, Name: "Confirmed Helm A"},
	}
	active := publishedCandidates(items, "1.60.1.70291", "1.60.1.70291")
	if msg := unconfirmedError("item", 1021, "Unconfirmed Helm", active[1021]); !strings.Contains(msg, "client_unconfirmed") {
		t.Fatalf("active build: unconfirmedError = %q, want a refusal", msg)
	}
	if msg := unconfirmedError("item", 1022, "Confirmed Helm A", active[1022]); msg != "" {
		t.Fatalf("active build, confirmed item: unconfirmedError = %q, want none", msg)
	}
	older := publishedCandidates(items, "1.60.1.70009", "1.60.1.70291")
	if msg := unconfirmedError("alternative", 1021, "Unconfirmed Helm", older[1021]); msg != "" {
		t.Fatalf("older build: unconfirmedError = %q, want it tolerated", msg)
	}
}

// Fix round 1, item 1: one off-hand pool per spec, so the pick, the
// alternatives, the swaps and set completion all read the same shape.
func TestTheOffHandPoolIsWeaponsForADualWielderAndNoWeaponsForAnyoneElse(t *testing.T) {
	shredder := scored{candidate: candidate{ID: 11863, Name: "White Bone Shredder", ClassID: itemClassWeapon}}
	tome := scored{candidate: candidate{ID: 272087, Name: "Tome of the Darkspear Prophecy", ClassID: 4}}
	rod := scored{candidate: candidate{ID: 15962, Name: "Satyr's Rod", ClassID: 4}}
	if got := offHandPool("druid-feral-bear", []scored{shredder, tome}); len(got) != 1 || got[0].ID != 272087 {
		t.Fatalf("druid-feral-bear off-hand pool = %+v, want the held item only", got)
	}
	if got := offHandPool("hunter-beast-mastery", []scored{rod, shredder}); len(got) != 1 || got[0].ID != 11863 {
		t.Fatalf("hunter-beast-mastery off-hand pool = %+v, want the weapon only", got)
	}
}

func TestThePublishedCheckRefusesAnOffHandAlternativeOfTheWrongShape(t *testing.T) {
	byID := map[int]candidate{
		11863:  {ID: 11863, ClassID: itemClassWeapon},
		272087: {ID: 272087, ClassID: 4},
	}
	row := func(alts ...int) map[string]slotRow {
		r := slotRow{Slot: "off_hand", ItemID: 1}
		for _, id := range alts {
			r.Alternatives = append(r.Alternatives, alternativeRow{ItemID: id})
		}
		return map[string]slotRow{"off_hand": r}
	}
	label := func(s string) string { return s }
	for _, tc := range []struct {
		spec  string
		alt   int
		fails bool
	}{
		{"druid-feral-bear", 11863, true},
		{"druid-feral-bear", 272087, false},
		{"hunter-beast-mastery", 272087, true},
		{"hunter-beast-mastery", 11863, false},
	} {
		tb := &fakeTB{}
		checkOffHandAlternatives(tb, label, tc.spec, byID, row(tc.alt))
		if got := len(tb.errors) > 0; got != tc.fails {
			t.Errorf("%s off-hand alternative %d: refused = %v, want %v (%v)", tc.spec, tc.alt, got, tc.fails, tb.errors)
		}
	}
}
