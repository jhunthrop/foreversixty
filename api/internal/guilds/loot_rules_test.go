// api/internal/guilds/loot_rules_test.go
package guilds

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/equipskills"
)

// onyxiaRow is the table's row for itemID, failing the test when it is not there.
func onyxiaRow(t *testing.T, itemID int) lootItem {
	t.Helper()
	item, ok := onyxiaItem(itemID)
	if !ok {
		t.Fatalf("no onyxiaLoot row for %d", itemID)
	}
	return item
}

// candidate is a verified roster row at level, of a class and spec, wearing nothing; no faction, so
// no BiS band applies and the row is ranked on the item-level tier.
func candidate(name, class, spec string, level int) RosterRow {
	return RosterRow{
		CharacterKey: "us/hardcore/" + name, Name: name, Verified: true,
		Class: &class, className: class, Spec: &spec, specStr: spec,
		Level: i(level), gear: map[string]int{},
	}
}

// everySpec is one level-60 character of every class and spec.
func everySpec() []RosterRow {
	out := []RosterRow{}
	for key := range specFamilies {
		class, spec := splitClassSpec(key)
		out = append(out, candidate(class+"-"+spec, class, spec, 60))
	}
	return out
}

func splitClassSpec(key string) (class, spec string) {
	class, spec, _ = strings.Cut(key, "/")
	return class, spec
}

func namesOf(cands []LootCandidate) map[string]bool {
	out := map[string]bool{}
	for _, c := range cands {
		out[c.Name] = true
	}
	return out
}

func classesOf(cands []LootCandidate) map[string]bool {
	out := map[string]bool{}
	for _, c := range cands {
		out[c.Class] = true
	}
	return out
}

func TestAncientCornerstoneGrimoireNeverListsAHunterWarriorOrRogue(t *testing.T) {
	s := realTreesStore(t)
	cands := s.lootCandidates(everySpec(), onyxiaRow(t, 17067))
	classes := classesOf(cands)
	for _, class := range []string{"hunter", "warrior", "rogue"} {
		if classes[class] {
			t.Errorf("Grimoire lists a %s: %+v", class, cands)
		}
	}
	names := namesOf(cands)
	for _, name := range []string{"mage-frost", "priest-holy", "druid-restoration", "paladin-holy", "shaman-elemental"} {
		if !names[name] {
			t.Errorf("Grimoire does not list %s: %v", name, names)
		}
	}
	for _, name := range []string{"paladin-retribution", "shaman-enhancement", "druid-feral"} {
		if names[name] {
			t.Errorf("Grimoire lists physical spec %s", name)
		}
	}
}

func TestDeathbringerListsPhysicalSpecsThatCanWieldIt(t *testing.T) {
	s := realTreesStore(t)
	names := namesOf(s.lootCandidates(everySpec(), onyxiaRow(t, 17068)))
	for _, name := range []string{"warrior-fury", "warrior-arms", "warrior-protection", "hunter-survival",
		"paladin-retribution", "paladin-protection", "shaman-enhancement"} {
		if !names[name] {
			t.Errorf("Deathbringer does not list %s: %v", name, names)
		}
	}
	for _, name := range []string{"paladin-holy", "priest-holy", "priest-discipline", "priest-shadow",
		"rogue-combat", "mage-fire", "shaman-restoration", "shaman-elemental", "druid-feral"} {
		if names[name] {
			t.Errorf("Deathbringer lists %s", name)
		}
	}
}

func TestViskagNeverListsACaster(t *testing.T) {
	s := realTreesStore(t)
	names := namesOf(s.lootCandidates(everySpec(), onyxiaRow(t, 17075)))
	for _, name := range []string{"rogue-combat", "warrior-fury", "hunter-beast mastery", "paladin-retribution"} {
		if !names[name] {
			t.Errorf("Vis'kag does not list %s: %v", name, names)
		}
	}
	for _, name := range []string{"mage-arcane", "mage-fire", "mage-frost", "warlock-affliction",
		"warlock-destruction", "paladin-holy"} {
		if names[name] {
			t.Errorf("Vis'kag lists %s", name)
		}
	}
}

func TestALevel58IsOnNoOnyxiaItem(t *testing.T) {
	s := realTreesStore(t)
	for _, item := range onyxiaLoot {
		roster := []RosterRow{candidate("Warrior", "warrior", "Fury", 58), candidate("Mage", "mage", "Frost", 58)}
		if cands := s.lootCandidates(roster, item); len(cands) != 0 {
			t.Errorf("%s lists a level 58: %+v", item.Name, cands)
		}
	}
}

func TestAnUnknownLevelIsNotACandidate(t *testing.T) {
	s := realTreesStore(t)
	row := candidate("Warrior", "warrior", "Fury", 0)
	row.Level = nil
	if cands := s.lootCandidates([]RosterRow{row}, onyxiaRow(t, 16963)); len(cands) != 0 {
		t.Fatalf("an unknown level is listed: %+v", cands)
	}
}

func TestTheRosterSnapshotLevelStandsInForAMissingExportLevel(t *testing.T) {
	s := realTreesStore(t)
	row := candidate("Warrior", "warrior", "Fury", 0)
	row.Level = nil
	level := 60
	row.Addon = &AddonMemberView{Level: &level}
	if cands := s.lootCandidates([]RosterRow{row}, onyxiaRow(t, 16963)); len(cands) != 1 {
		t.Fatalf("candidates = %+v, want the snapshot's level 60 to count", cands)
	}
}

func TestTierHelmsListOnlyTheirClass(t *testing.T) {
	s := realTreesStore(t)
	want := map[int]string{
		16900: "druid", 16908: "rogue", 16914: "mage", 16921: "priest", 16929: "warlock",
		16939: "hunter", 16947: "shaman", 16955: "paladin", 16963: "warrior",
	}
	for itemID, class := range want {
		classes := classesOf(s.lootCandidates(everySpec(), onyxiaRow(t, itemID)))
		if len(classes) != 1 || !classes[class] {
			t.Errorf("item %d lists %v, want only %s", itemID, classes, class)
		}
	}
}

func TestArmourNeedsItsLearnLevel(t *testing.T) {
	plate, mail := onyxiaRow(t, 16963), onyxiaRow(t, 16939)
	cases := []struct {
		class string
		level int
		item  lootItem
		want  bool
	}{
		{"warrior", 39, plate, false},
		{"warrior", 40, plate, true},
		{"hunter", 39, mail, false},
		{"hunter", 40, mail, true},
	}
	for _, c := range cases {
		if got := canEquip(c.class, c.level, c.item); got != c.want {
			t.Errorf("canEquip(%s, %d, %s) = %v, want %v", c.class, c.level, c.item.Name, got, c.want)
		}
	}
	// A class with no skill for the type never equips it.
	shamanPlate := lootItem{ItemClass: itemClassArmor, Subclass: 4, AllowableClass: -1}
	if canEquip("shaman", 60, shamanPlate) {
		t.Error("a shaman can equip plate")
	}
}

func TestDualWieldLevels(t *testing.T) {
	cases := []struct {
		class string
		level int
		want  bool
	}{
		{"shaman", 60, false}, {"rogue", 9, false}, {"rogue", 10, true},
		{"warrior", 19, false}, {"warrior", 20, true}, {"hunter", 20, true},
		{"paladin", 60, false}, {"mage", 60, false},
	}
	for _, c := range cases {
		if got := canDualWield(c.class, c.level); got != c.want {
			t.Errorf("canDualWield(%s, %d) = %v, want %v", c.class, c.level, got, c.want)
		}
	}
}

func TestAShamanIsNeverACandidateForAOneHanderInTheOffHand(t *testing.T) {
	deathbringer := onyxiaRow(t, 17068)
	shaman := candidate("Shaman", "shaman", "Enhancement", 60)
	if got := characterSlots(shaman, deathbringer, 60); !reflect.DeepEqual(got, []string{"main_hand"}) {
		t.Errorf("shaman slots = %v, want the main hand alone", got)
	}
	warrior := candidate("Warrior", "warrior", "Fury", 60)
	if got := characterSlots(warrior, deathbringer, 60); !reflect.DeepEqual(got, []string{"main_hand", "off_hand"}) {
		t.Errorf("warrior slots = %v, want both hands", got)
	}
}

func TestAOneHanderOnADualWielderReplacesTheWeakerHand(t *testing.T) {
	s := realTreesStore(t)
	// Only the level of what is worn matters to the rule: Vis'kag (74) in the main hand, and Red
	// Winter Hat (item level 1) stands in for a weak off hand.
	strong, weak := onyxiaRow(t, 17075).ItemID, 21524
	row := candidate("Warrior", "warrior", "Fury", 60)
	row.gear = map[string]int{"main_hand": strong, "off_hand": weak}
	delta, slot, ok := s.fallbackMatch(row, onyxiaRow(t, 17068), characterSlots(row, onyxiaRow(t, 17068), 60))
	if !ok || slot != "off_hand" || delta != 74 {
		t.Fatalf("fallbackMatch = %d/%s/%v, want off_hand at +74 (75 over item level 1)", delta, slot, ok)
	}
}

func TestARingRowNamesTheWeakerWornRing(t *testing.T) {
	s := realTreesStore(t)
	// Only the level of what is worn matters to the rule: Sapphiron Drape (72) stands in for a strong
	// ring and Red Winter Hat (item level 1) for a weak one.
	ring := onyxiaRow(t, 18813)
	strong := onyxiaRow(t, 17078).ItemID // Sapphiron Drape, item level 72
	row := candidate("Warrior", "warrior", "Protection", 60)
	row.gear = map[string]int{"finger1": strong, "finger2": 21524}
	cands := s.lootCandidates([]RosterRow{row}, ring)
	if len(cands) != 1 || cands[0].Replaces != "Red Winter Hat" {
		t.Fatalf("candidates = %+v, want one replacing the weaker worn ring", cands)
	}
	row.gear = map[string]int{"finger1": strong}
	cands = s.lootCandidates([]RosterRow{row}, ring)
	if len(cands) != 1 || cands[0].ReplacesItem != nil || cands[0].IlvlDelta == nil || *cands[0].IlvlDelta != ring.ItemLevel {
		t.Fatalf("candidates = %+v, want the empty ring slot first", cands)
	}
}

func TestAUniqueDropAlreadyHeldIsNoUpgrade(t *testing.T) {
	s := realTreesStore(t)
	viskag := onyxiaRow(t, 17075)
	row := candidate("Rogue", "rogue", "Combat", 60)
	row.gear = map[string]int{"main_hand": viskag.ItemID}
	if cands := s.lootCandidates([]RosterRow{row}, viskag); len(cands) != 0 {
		t.Fatalf("candidates = %+v, want a held unique drop not to fill the off hand", cands)
	}
}

func TestSuitsFamily(t *testing.T) {
	phys, caster, healer := []statFamily{familyPhysical}, []statFamily{familyCaster}, []statFamily{familyHealer}
	tank := []statFamily{familyTank}
	cases := []struct {
		name     string
		stats    map[string]int
		weapon   bool
		families []statFamily
		want     bool
	}{
		{"strength for physical", map[string]int{"strength": 5}, false, phys, true},
		{"strength not for a caster", map[string]int{"strength": 5}, false, caster, false},
		{"intellect for a healer", map[string]int{"intellect": 5}, false, healer, true},
		{"intellect not for physical", map[string]int{"intellect": 5}, false, phys, false},
		{"mp5 for a healer", map[string]int{"mp5": 5}, false, healer, true},
		{"mp5 not for a caster", map[string]int{"mp5": 5}, false, caster, false},
		{"hit for everyone", map[string]int{"hit": 1}, false, caster, true},
		{"crit for everyone", map[string]int{"crit": 1, "stamina": 5}, false, healer, true},
		{"stamina only for anyone", map[string]int{"stamina": 5, "frost_res": 5}, false, caster, true},
		{"no stats weapon is physical", nil, true, phys, true},
		{"no stats weapon is a tank's", nil, true, tank, true},
		{"no stats weapon not a caster's", nil, true, caster, false},
		{"defense is a tank's", map[string]int{"defense": 4, "fire_res": 10}, false, tank, true},
		{"defense not a caster's", map[string]int{"defense": 4, "fire_res": 10}, false, caster, false},
	}
	for _, c := range cases {
		if got := suitsFamily(c.stats, c.weapon, c.families); got != c.want {
			t.Errorf("%s: suitsFamily = %v, want %v", c.name, got, c.want)
		}
	}
}

func TestFamiliesForReadsTheSpecThenTheRole(t *testing.T) {
	if got := familiesFor("druid", "Feral", "tank"); !reflect.DeepEqual(got, []statFamily{familyTank}) {
		t.Errorf("bear = %v, want tank", got)
	}
	if got := familiesFor("druid", "Feral", "dps"); !reflect.DeepEqual(got, []statFamily{familyPhysical}) {
		t.Errorf("cat = %v, want physical", got)
	}
	if got := familiesFor("paladin", "Holy", "healer"); !reflect.DeepEqual(got, []statFamily{familyHealer}) {
		t.Errorf("holy paladin = %v, want healer", got)
	}
	if got := familiesFor("mage", "", ""); !reflect.DeepEqual(got, []statFamily{familyCaster}) {
		t.Errorf("mage with no spec = %v, want caster", got)
	}
	if got := familiesFor("paladin", "", ""); len(got) != 3 {
		t.Errorf("paladin with no spec = %v, want its three families", got)
	}
}

func TestLootKinds(t *testing.T) {
	cases := map[int]string{
		15410: lootKindSkinned, 16963: lootKindGear, 17966: lootKindBag,
		18422: lootKindQuest, 18705: lootKindQuest, 21108: lootKindQuest,
	}
	for itemID, want := range cases {
		if got := lootKind(onyxiaRow(t, itemID)); got != want {
			t.Errorf("lootKind(%d) = %s, want %s", itemID, got, want)
		}
	}
}

func TestQuestLabels(t *testing.T) {
	if got := questLabel([]int{18422, 18423}); got != "Quest: Victory for the Alliance / Victory for the Horde" {
		t.Errorf("Head of Onyxia label = %q", got)
	}
	if got := questLabel([]int{18705}); got != "Quest: A Proper String (Hunter)" {
		t.Errorf("Sinew label = %q", got)
	}
	if got := questLabel([]int{1}); got != "Quest item" {
		t.Errorf("unsourced label = %q", got)
	}
}

func TestMatureBlackDragonSinewIsForHuntersAtSixtyAlone(t *testing.T) {
	roster := []RosterRow{
		candidate("Hunter", "hunter", "Survival", 60),
		candidate("Young", "hunter", "Survival", 59),
		candidate("Warrior", "warrior", "Fury", 60),
	}
	got := questEligible(roster, onyxiaRow(t, 18705), lootKindQuest)
	if len(got) != 1 || got[0].Name != "Hunter" || got[0].ItemID != 18705 {
		t.Fatalf("eligible = %+v, want the level-60 hunter alone", got)
	}
}

func TestHeadOfOnyxiaListsTheGuildsFaction(t *testing.T) {
	s := realTreesStore(t)
	alliance, horde := candidate("Ally", "warrior", "Fury", 60), candidate("Hordie", "warrior", "Fury", 60)
	alliance.faction, horde.faction = "alliance", "horde"
	roster := []RosterRow{alliance, horde}
	var group []lootItem
	for _, g := range lootGroups(onyxiaLoot) {
		if g[0].Name == "Head of Onyxia" {
			group = g
		}
	}
	guild := "alliance"
	iv := s.lootItemView(roster, group, &guild)
	if iv.ItemID != 18423 || len(iv.Eligible) != 1 || iv.Eligible[0].Name != "Ally" || iv.Eligible[0].ItemID != 18423 {
		t.Errorf("alliance guild = %+v, want the Alliance head for Ally", iv)
	}
	if iv.Copies != 2 || iv.QuestLabel != "Quest: Victory for the Alliance / Victory for the Horde" {
		t.Errorf("row = %+v, want both copies under one label", iv)
	}
	// A guild with no faction yet offers each character their own side's head.
	iv = s.lootItemView(roster, group, nil)
	if len(iv.Eligible) != 2 {
		t.Fatalf("eligible = %+v, want one per side", iv.Eligible)
	}
	for _, e := range iv.Eligible {
		if (e.Name == "Ally") != (e.ItemID == 18423) {
			t.Errorf("%s gets %d", e.Name, e.ItemID)
		}
	}
}

func TestTheBagIsForAnyoneAtSixtyAndTheScaleForNoOne(t *testing.T) {
	s := realTreesStore(t)
	roster := []RosterRow{candidate("Mage", "mage", "Frost", 60), candidate("Young", "rogue", "Combat", 20)}
	bag := s.lootItemView(roster, []lootItem{onyxiaRow(t, 17966)}, nil)
	if bag.Kind != lootKindBag || len(bag.Eligible) != 1 || bag.Eligible[0].Name != "Mage" {
		t.Errorf("bag = %+v, want the level 60 alone", bag)
	}
	scale := s.lootItemView(roster, []lootItem{onyxiaRow(t, scaleOfOnyxia)}, nil)
	if scale.Kind != lootKindSkinned || len(scale.Eligible) != 0 || len(scale.Candidates) != 0 {
		t.Errorf("scale = %+v, want nobody", scale)
	}
}

func TestCountAtLevelCap(t *testing.T) {
	unverified := candidate("Unverified", "mage", "Frost", 60)
	unverified.Verified = false
	roster := []RosterRow{candidate("A", "mage", "Frost", 60), candidate("B", "mage", "Frost", 59), unverified}
	if got := countAtLevelCap(roster); got != 1 {
		t.Errorf("countAtLevelCap = %d, want 1", got)
	}
}

func TestGainUnitFor(t *testing.T) {
	cases := []struct {
		role, unit, want string
	}{
		{"", "reference_stat_points", gainUnitDPS},
		{"healer", "reference_stat_points", gainUnitHPS},
		{"tank", "tank_score", gainUnitTPS},
		{"", "tank_score", gainUnitTPS},
	}
	for _, c := range cases {
		if got := gainUnitFor(bis.Band{Role: c.role, ScoreUnit: c.unit}); got != c.want {
			t.Errorf("gainUnitFor(%s, %s) = %s, want %s", c.role, c.unit, got, c.want)
		}
	}
}

// buildFile reads a committed file of the build the tables were copied from.
func buildFile(t *testing.T, name string) []byte {
	t.Helper()
	body, err := os.ReadFile(filepath.Join(repoDataDirForTrees(t), equipSkillsBuild, name))
	if err != nil {
		t.Fatal(err)
	}
	return body
}

// TestOnyxiaLootMatchesTheBuild pins the table's client facts to items.json and classic-db.
func TestOnyxiaLootMatchesTheBuild(t *testing.T) {
	var items []struct {
		ID            int `json:"id"`
		ClassID       int `json:"class_id"`
		SubclassID    int `json:"subclass_id"`
		InventoryType int `json:"inventory_type"`
		RequiredLevel int `json:"required_level"`
		ItemLevel     int `json:"item_level"`
	}
	if err := json.Unmarshal(buildFile(t, "items.json"), &items); err != nil {
		t.Fatal(err)
	}
	var template struct {
		Items []struct {
			ID             int  `json:"id"`
			AllowableClass int  `json:"allowable_class"`
			Unique         bool `json:"unique"`
		} `json:"items"`
	}
	if err := json.Unmarshal(buildFile(t, "raw/classicdb/item_template.json"), &template); err != nil {
		t.Fatal(err)
	}
	seen := 0
	for _, row := range items {
		item, ok := onyxiaItem(row.ID)
		if !ok {
			continue
		}
		seen++
		if item.ItemClass != row.ClassID || item.Subclass != row.SubclassID || item.InventoryType != row.InventoryType ||
			item.RequiredLevel != row.RequiredLevel || item.ItemLevel != row.ItemLevel {
			t.Errorf("%s: table %+v, items.json %+v", item.Name, item, row)
		}
	}
	if seen != len(onyxiaLoot) {
		t.Errorf("items.json carries %d of the table's %d items", seen, len(onyxiaLoot))
	}
	for _, row := range template.Items {
		item, ok := onyxiaItem(row.ID)
		if !ok {
			continue
		}
		if item.AllowableClass != row.AllowableClass || item.Unique != row.Unique {
			t.Errorf("%s: table allowable %d unique %v, classic-db %d %v", item.Name, item.AllowableClass, item.Unique, row.AllowableClass, row.Unique)
		}
	}
}

// TestQuestTableMatchesItsCuratedFile pins onyxiaQuestItems to data/curated/raids/onyxia-quest-items.json.
func TestQuestTableMatchesItsCuratedFile(t *testing.T) {
	body, err := os.ReadFile(filepath.Join(repoDataDirForTrees(t), "..", "curated", "raids", "onyxia-quest-items.json"))
	if err != nil {
		t.Fatal(err)
	}
	var file struct {
		Items []struct {
			ItemID  int     `json:"item_id"`
			Quest   string  `json:"quest"`
			Class   *string `json:"class"`
			Faction *string `json:"faction"`
		} `json:"items"`
	}
	if err := json.Unmarshal(body, &file); err != nil {
		t.Fatal(err)
	}
	if len(file.Items) != len(onyxiaQuestItems) {
		t.Fatalf("file has %d rows, table %d", len(file.Items), len(onyxiaQuestItems))
	}
	for _, row := range file.Items {
		q, ok := questFor(row.ItemID)
		if !ok || q.Quest != row.Quest || q.Class != orEmpty(row.Class) || q.Faction != orEmpty(row.Faction) {
			t.Errorf("item %d: table %+v, file %+v", row.ItemID, q, row)
		}
	}
}

func orEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// TestEquipRulesTableMatchesItsGenerator pins equip_rules_table.go to what cmd/genequiprules writes
// from the build's equip-skills.json.
func TestEquipRulesTableMatchesItsGenerator(t *testing.T) {
	doc, err := equipskills.Parse(buildFile(t, "equip-skills.json"))
	if err != nil {
		t.Fatal(err)
	}
	want, err := equipskills.GoSource(doc, equipskills.SourcePath(equipSkillsBuild))
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile("equip_rules_table.go")
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(want) {
		t.Fatal("equip_rules_table.go is stale: run go generate ./internal/guilds/ from api/")
	}
}

// TestEquipRulesHoldTheClassicFacts sanity-checks the generated table against the Classic facts the
// loot spec lists (§3.2 rule 3). Shamans training fists is Forever's own change.
func TestEquipRulesHoldTheClassicFacts(t *testing.T) {
	const axe, twoHandAxe, mace, twoHandMace, polearm, sword, twoHandSword, staff, fist, dagger = 0, 1, 4, 5, 6, 7, 8, 10, 13, 15
	const mail, plate = 3, 4
	no := []struct {
		class string
		table func(equipClass) map[int]int
		sub   int
	}{
		{"rogue", weapons, axe}, {"rogue", weapons, twoHandAxe}, {"rogue", weapons, twoHandMace},
		{"rogue", weapons, twoHandSword}, {"rogue", weapons, polearm}, {"rogue", weapons, staff},
		{"priest", weapons, sword}, {"hunter", weapons, mace},
		{"paladin", weapons, dagger}, {"paladin", weapons, fist}, {"paladin", weapons, staff},
		{"rogue", armor, mail}, {"druid", armor, mail}, {"hunter", armor, plate}, {"shaman", armor, plate},
	}
	for _, c := range no {
		if _, ok := c.table(equipSkills[c.class])[c.sub]; ok {
			t.Errorf("%s has subclass %d", c.class, c.sub)
		}
	}
	for class, want := range map[string]map[int]int{
		"warrior": {plate: 40}, "paladin": {plate: 40}, "hunter": {mail: 40}, "shaman": {mail: 40},
	} {
		for sub, level := range want {
			if got := equipSkills[class].armor[sub]; got != level {
				t.Errorf("%s armour %d learned at %d, want %d", class, sub, got, level)
			}
		}
	}
	if _, ok := equipSkills["shaman"].weapons[fist]; !ok {
		t.Error("a Forever shaman trains fist weapons")
	}
}

func weapons(c equipClass) map[int]int { return c.weapons }
func armor(c equipClass) map[int]int   { return c.armor }
