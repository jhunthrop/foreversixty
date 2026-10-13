package main

import (
	"path/filepath"
	"strings"
	"testing"
)

// Task L1.7 (the persona passes' pick audit): loot.json marks a rare
// spawn's own rows `rare_spawn` (pipeline.loot.npc_rank: a rare-ranked
// creature with a pooled or open-world spawn, from the pinned classic-db
// dump). The ranker names such a source "(rare)", prefers any other source
// the item has, and never makes a rare spawn's drop the pick of a slot that
// has a candidate from another source at the band. A trash drop several
// creatures of one instance share, none with a rate, is named by its
// instance alone ("Gnomeregan"), never by whichever creature loot.json
// happens to list first.

func loadTestLootIndex(t *testing.T, lootJSON string) lootIndex {
	t.Helper()
	dir := t.TempDir()
	if err := writeFile(t, filepath.Join(dir, "loot.json"), lootJSON); err != nil {
		t.Fatal(err)
	}
	idx, _, err := loadLootIndex(dir, nil)
	if err != nil {
		t.Fatalf("loadLootIndex: %v", err)
	}
	return idx
}

func TestLoadLootIndexNamesARareSpawnsSourceRare(t *testing.T) {
	idx := loadTestLootIndex(t, `{
		"sources": [
			{"id": "dungeon:scarlet-monastery", "kind": "dungeon", "name": "Scarlet Monastery", "bosses": [
				{"name": "Ironspine", "npc_id": 6489, "rank": 2, "rare_spawn": true, "items": [7686]},
				{"name": "Herod", "npc_id": 3975, "rank": 1, "items": [7718]}
			]},
			{"id": "world:swiftmane", "kind": "world", "name": "Swiftmane", "rare_spawn": true, "items": [285330]},
			{"id": "world:bluegill-muckdweller", "kind": "world", "name": "Bluegill Muckdweller", "items": [3022]}
		],
		"quests": {}
	}`)
	cases := []struct {
		id    int
		label string
		rare  bool
	}{
		{7686, "Scarlet Monastery: Ironspine (rare)", true},
		{7718, "Scarlet Monastery: Herod", false},
		{285330, "Swiftmane (rare)", true},
		{3022, "Bluegill Muckdweller", false},
	}
	for _, c := range cases {
		srcs := idx[c.id]
		if len(srcs) != 1 || srcs[0].Label != c.label || srcs[0].RareSpawn != c.rare {
			t.Errorf("idx[%d] = %+v, want one source %q with RareSpawn %v", c.id, srcs, c.label, c.rare)
		}
	}
}

func TestLoadLootIndexNamesAMultiCreatureTrashDropWithNoRateByItsInstance(t *testing.T) {
	idx := loadTestLootIndex(t, `{
		"sources": [
			{"id": "dungeon:gnomeregan", "kind": "dungeon", "name": "Gnomeregan", "trash": [9509, 9510, 9511], "bosses": [
				{"name": "Caverndeep Burrower", "npc_id": 6206, "items": [9509, 9510, 9511], "item_chances": {"9510": 1.5}},
				{"name": "Caverndeep Ambusher", "npc_id": 6207, "items": [9509, 9510]},
				{"name": "Mekgineer Thermaplugg", "npc_id": 7800, "items": [9509, 9999]}
			]},
			{"id": "dungeon:wailing-caverns", "kind": "dungeon", "name": "Wailing Caverns", "trash": [10413], "bosses": [
				{"name": "Druid of the Fang", "npc_id": 3840, "items": [10413], "item_chances": {"10413": 2.34}}
			]}
		],
		"quests": {}
	}`)
	// Petrolspill Leggings: trash, three creatures, no rate anywhere. The
	// source keeps the creatures it stands for (the old build's label check).
	if srcs := idx[9509]; len(srcs) != 1 || srcs[0].Label != "Gnomeregan" || srcs[0].Kind != "dungeon" || srcs[0].Chance != 0 ||
		strings.Join(srcs[0].SharedTrashCreatures, ",") != "Caverndeep Burrower,Caverndeep Ambusher,Mekgineer Thermaplugg" {
		t.Errorf("idx[9509] = %+v, want the one instance-wide source \"Gnomeregan\" standing for its three creatures", srcs)
	}
	// A trash drop one creature states a rate for keeps its creatures.
	if src, ok := sourceFor(9510, 30, "alliance", "", idx); !ok || src.Label != "Gnomeregan: Caverndeep Burrower" || src.Chance != 1.5 {
		t.Errorf("sourceFor(9510) = %+v, %v, want the Burrower at 1.5%%", src, ok)
	}
	// A trash drop only one creature names keeps that creature.
	if srcs := idx[9511]; len(srcs) != 1 || srcs[0].Label != "Gnomeregan: Caverndeep Burrower" {
		t.Errorf("idx[9511] = %+v, want the one creature that drops it", srcs)
	}
	// A boss's own item that is not trash keeps its boss.
	if srcs := idx[9999]; len(srcs) != 1 || srcs[0].Label != "Gnomeregan: Mekgineer Thermaplugg" {
		t.Errorf("idx[9999] = %+v, want Thermaplugg", srcs)
	}
	// Gloves of the Fang: trash, one creature, with its rate.
	if srcs := idx[10413]; len(srcs) != 1 || srcs[0].Label != "Wailing Caverns: Druid of the Fang" || srcs[0].Chance != 2.34 {
		t.Errorf("idx[10413] = %+v, want the Druid of the Fang at 2.34%%", srcs)
	}
}

func TestSourceForPrefersAnySourceThatIsNotARareSpawn(t *testing.T) {
	idx := lootIndex{
		1: {
			{Kind: "dungeon", Label: "Scarlet Monastery: Ironspine (rare)", Chance: 20, RareSpawn: true},
			{Kind: "crafted", Label: "Leatherworking"},
		},
		2: {{Kind: "dungeon", Label: "Scarlet Monastery: Ironspine (rare)", RareSpawn: true}},
	}
	if src, ok := sourceFor(1, 30, "alliance", "", idx); !ok || src.Label != "Leatherworking" {
		t.Errorf("sourceFor(1) = %+v, %v, want the crafted source over the rare spawn", src, ok)
	}
	if src, ok := sourceFor(2, 30, "alliance", "", idx); !ok || !src.RareSpawn {
		t.Errorf("sourceFor(2) = %+v, %v, want the rare spawn, the item's only source", src, ok)
	}
}

func TestBuildBandPoolKeepsARareSpawnsDropOnlyWhereNothingElseFillsTheSlot(t *testing.T) {
	finger := []string{"finger1", "finger2"}
	items := []candidate{
		{ID: 7686, Name: "Ironspine's Eye", RequiredLevel: 30, EffectiveRequiredLevel: 30, Stats: map[string]float64{"agility": 9}, Slots: finger},
		{ID: 1491, Name: "Ring of Precision", RequiredLevel: 25, EffectiveRequiredLevel: 25, Stats: map[string]float64{"agility": 4}, Slots: finger},
		{ID: 6748, Name: "Monkey Ring", RequiredLevel: 25, EffectiveRequiredLevel: 25, Stats: map[string]float64{"agility": 3}, Slots: finger},
		{ID: 9001, Name: "Rare Helm", RequiredLevel: 25, EffectiveRequiredLevel: 25, Stats: map[string]float64{"agility": 5}, Slots: []string{"head"}},
	}
	idx := lootIndex{
		7686: {{Kind: "dungeon", Label: "Scarlet Monastery: Ironspine (rare)", RareSpawn: true}},
		1491: {{Kind: "dungeon", Label: "Blackfathom Deeps: Fallenroot Shadowstalker"}},
		6748: {{Kind: "vendor", Label: "Willix the Importer"}},
		9001: {{Kind: "world", Label: "Lupos (rare)", RareSpawn: true}},
	}
	pool := buildBandPool(items, idx, "hunter", 30, "alliance", map[string]float64{"agility": 1}, 0, false, true)
	got := map[int]bool{}
	for _, s := range pool.Scored {
		got[s.ID] = true
	}
	if got[7686] || !got[1491] || !got[6748] || !got[9001] || len(pool.Scored) != 3 {
		t.Errorf("pool.Scored ids = %v, want 1491, 6748 and 9001: the rare ring gives way to the two other rings, the rare helm has no rival", got)
	}
	if pool.RareSpawnSkipped != 1 {
		t.Errorf("pool.RareSpawnSkipped = %d, want 1", pool.RareSpawnSkipped)
	}
}

// Fix round 1, D1: a source sourceFor can never pick (a `zone` source is
// in neither sourceKindPriority nor bossKinds) does not count as the
// "other" source a rare spawn's drop gives way to - the item would be left
// with no source at all. Signet of the Zhevra: Swiftmane plus zone:17.
func TestSourceForKeepsARareSpawnWhoseOtherSourcesCannotBePicked(t *testing.T) {
	idx := lootIndex{
		285330: {
			{Kind: "world", Label: "Swiftmane (rare)", RareSpawn: true},
			{Kind: "zone", Label: "The Barrens"},
		},
		286535: {
			{Kind: "zone", Label: "Wetlands"},
			{Kind: "dungeon", Label: "Wailing Caverns: Deviate Faerie Dragon (rare)", Chance: 4, RareSpawn: true},
		},
	}
	if src, ok := sourceFor(285330, 20, "horde", "", idx); !ok || src.Label != "Swiftmane (rare)" || !src.RareSpawn {
		t.Errorf("sourceFor(285330) = %+v, %v, want the rare spawn labelled \"Swiftmane (rare)\": the zone source is not one sourceFor picks", src, ok)
	}
	if src, ok := sourceFor(286535, 20, "horde", "", idx); !ok || src.Label != "Wailing Caverns: Deviate Faerie Dragon (rare)" || src.Chance != 4 {
		t.Errorf("sourceFor(286535) = %+v, %v, want the rare boss with its chance", src, ok)
	}
}

func idsOf(list []scored) map[int]bool {
	out := make(map[int]bool, len(list))
	for _, s := range list {
		out[s.ID] = true
	}
	return out
}

func pickID(sp slotPick) int {
	if sp.Item == nil {
		return 0
	}
	return sp.Item.ID
}

// Fix round 1, D2: the slot rule counts only candidates that can fill the
// slot (a positive score, or an effect a sim ranks - the trinket and effect
// passes), and a paired slot needs two of them, distinct by id and name
// (pick's excludePaired), before the rare drop leaves its second slot. A
// rare item is excluded slot by slot: it never surfaces, as the pick or as
// an alternative, in a slot something else covers. warlock-affliction band
// 20 horde finger2 went empty this way (Sludge-Stained Band).
func TestAPairedSlotKeepsARareDropForTheSlotItsRivalsCannotFill(t *testing.T) {
	finger := []string{"finger1", "finger2"}
	rare := candidate{ID: 286535, Name: "Sludge-Stained Band", EffectiveRequiredLevel: 20, Stats: map[string]float64{"intellect": 6}, Slots: finger}
	ring := func(id int, name string, stats map[string]float64) candidate {
		return candidate{ID: id, Name: name, EffectiveRequiredLevel: 15, Stats: stats, Slots: finger}
	}
	intellect := func(n float64) map[string]float64 { return map[string]float64{"intellect": n} }
	cases := []struct {
		name         string
		rivals       []candidate
		inFinger1    bool // the rare ring is offered in finger1
		inFinger2    bool // ... and in finger2
		pick1, pick2 int
		skipped      int
	}{
		{"one rival", []candidate{ring(1, "Legionnaire's Band", intellect(4))}, false, true, 1, 286535, 0},
		{"one rival and one that scores 0", []candidate{ring(1, "Legionnaire's Band", intellect(4)), ring(2, "Strength Ring", map[string]float64{"strength": 5})}, false, true, 1, 286535, 0},
		{"two rivals sharing a name", []candidate{ring(1, "Legionnaire's Band", intellect(4)), ring(3, "Legionnaire's Band", intellect(3))}, false, true, 1, 286535, 0},
		{"no rival", nil, true, true, 286535, 0, 0},
		{"two rivals", []candidate{ring(1, "Legionnaire's Band", intellect(4)), ring(4, "Bounty Hunter's Ring", intellect(2))}, false, false, 1, 4, 1},
	}
	for _, c := range cases {
		items := append([]candidate{rare}, c.rivals...)
		idx := lootIndex{286535: {{Kind: "world", Label: "Sludgeguard (rare)", RareSpawn: true}}}
		for _, r := range c.rivals {
			idx[r.ID] = []itemSource{{Kind: "vendor", Label: "A Vendor"}}
		}
		pool := buildBandPool(items, idx, "warlock", 20, "horde", map[string]float64{"intellect": 1}, 0, false, false)
		bySlot := candidatesBySlot(pool.Scored)
		if got := idsOf(bySlot["finger1"])[286535]; got != c.inFinger1 {
			t.Errorf("%s: rare ring among finger1's candidates = %v, want %v", c.name, got, c.inFinger1)
		}
		if got := idsOf(bySlot["finger2"])[286535]; got != c.inFinger2 {
			t.Errorf("%s: rare ring among finger2's candidates = %v, want %v", c.name, got, c.inFinger2)
		}
		picks := pick("warlock-affliction", bySlot)
		if got := pickID(picks["finger1"]); got != c.pick1 {
			t.Errorf("%s: finger1 pick = %d, want %d", c.name, got, c.pick1)
		}
		if got := pickID(picks["finger2"]); got != c.pick2 {
			t.Errorf("%s: finger2 pick = %d, want %d", c.name, got, c.pick2)
		}
		if pool.RareSpawnSkipped != c.skipped {
			t.Errorf("%s: RareSpawnSkipped = %d, want %d", c.name, pool.RareSpawnSkipped, c.skipped)
		}
	}
}

// The coverage block publishes, per slot, how many rare-spawn candidates
// the slot rule kept out of that slot.
func TestTheCoverageBlockCountsTheRareDropsEachSlotKeptOut(t *testing.T) {
	finger := []string{"finger1", "finger2"}
	items := []candidate{
		{ID: 7686, Name: "Ironspine's Eye", EffectiveRequiredLevel: 30, Stats: map[string]float64{"agility": 9}, Slots: finger},
		{ID: 1491, Name: "Ring of Precision", EffectiveRequiredLevel: 25, Stats: map[string]float64{"agility": 4}, Slots: finger},
		{ID: 9001, Name: "Rare Cloak", EffectiveRequiredLevel: 25, Stats: map[string]float64{"agility": 5}, Slots: []string{"back"}},
		{ID: 9002, Name: "Plain Cloak", EffectiveRequiredLevel: 25, Stats: map[string]float64{"agility": 1}, Slots: []string{"back"}},
	}
	idx := lootIndex{
		7686: {{Kind: "dungeon", Label: "Scarlet Monastery: Ironspine (rare)", RareSpawn: true}},
		1491: {{Kind: "dungeon", Label: "Blackfathom Deeps"}},
		9001: {{Kind: "world", Label: "Rohh the Silent (rare)", RareSpawn: true}},
		9002: {{Kind: "vendor", Label: "A Vendor"}},
	}
	pool := buildBandPool(items, idx, "hunter", 30, "alliance", map[string]float64{"agility": 1}, 0, false, true)
	want := map[string]coverageRow{
		"finger1": {Eligible: 2, Sourced: 2, RareSpawnSkipped: 1},
		"finger2": {Eligible: 2, Sourced: 2},
		"back":    {Eligible: 2, Sourced: 2, RareSpawnSkipped: 1},
	}
	for slot, row := range want {
		if pool.Coverage[slot] != row {
			t.Errorf("Coverage[%s] = %+v, want %+v", slot, pool.Coverage[slot], row)
		}
	}
	if pool.RareSpawnSkipped != 1 {
		t.Errorf("RareSpawnSkipped = %d, want 1 (the cloak; the ring keeps finger2)", pool.RareSpawnSkipped)
	}
}

// The weapon hands are a pair for a dual-wielder: a rare one-hander that
// one rival covers in the main hand moves to the off hand; two rivals take
// it out. A spec that does not dual-wield has one weapon hand, so one rival
// covers it. A trinket's rival counts by the effect the trinket pass sims,
// whatever its score.
func TestTheWeaponHandsAndTrinketsArePairedForTheRareRule(t *testing.T) {
	sword := func(id int, name string, dps float64) candidate {
		return candidate{ID: id, Name: name, EffectiveRequiredLevel: 20, ClassID: itemClassWeapon, DPS: dps, Stats: map[string]float64{"agility": 2}, Slots: []string{"main_hand"}}
	}
	rareSword := sword(9100, "Rare Sword", 20)
	idx := lootIndex{
		9100: {{Kind: "world", Label: "Lupos (rare)", RareSpawn: true}},
		9101: {{Kind: "vendor", Label: "A Vendor"}},
		9102: {{Kind: "vendor", Label: "A Vendor"}},
	}
	weights := map[string]float64{"agility": 1}

	pool := buildBandPool([]candidate{rareSword, sword(9101, "Plain Sword", 10)}, idx, "rogue", 20, "horde", weights, 0, false, true)
	bySlot := candidatesBySlot(pool.Scored)
	if idsOf(bySlot["main_hand"])[9100] || !idsOf(bySlot["off_hand"])[9100] {
		t.Errorf("dual-wield, one rival: main_hand %v, off_hand %v, want the rare sword in the off hand only", idsOf(bySlot["main_hand"]), idsOf(bySlot["off_hand"]))
	}
	picks := pick("rogue-combat", bySlot)
	if pickID(picks["main_hand"]) != 9101 || pickID(picks["off_hand"]) != 9100 {
		t.Errorf("dual-wield, one rival: picks main %d off %d, want 9101 and 9100", pickID(picks["main_hand"]), pickID(picks["off_hand"]))
	}

	pool = buildBandPool([]candidate{rareSword, sword(9101, "Plain Sword", 10), sword(9102, "Other Sword", 9)}, idx, "rogue", 20, "horde", weights, 0, false, true)
	if idsOf(pool.Scored)[9100] {
		t.Errorf("dual-wield, two rivals: the rare sword stayed in the pool")
	}

	pool = buildBandPool([]candidate{rareSword, sword(9101, "Plain Sword", 10)}, idx, "paladin", 20, "horde", weights, 0, false, false)
	if idsOf(pool.Scored)[9100] {
		t.Errorf("one weapon hand, one rival: the rare sword stayed in the pool")
	}

	trinket := []string{"trinket1", "trinket2"}
	items := []candidate{
		{ID: 9200, Name: "Rare Charm", EffectiveRequiredLevel: 20, Stats: map[string]float64{"agility": 3}, Slots: trinket},
		{ID: 9201, Name: "Use Trinket", EffectiveRequiredLevel: 20, EffectText: "Use: Increases agility by 20 for 15 sec.", Slots: trinket},
		{ID: 9202, Name: "Dead Trinket", EffectiveRequiredLevel: 20, Slots: trinket},
	}
	idx = lootIndex{
		9200: {{Kind: "world", Label: "Mother Fang (rare)", RareSpawn: true}},
		9201: {{Kind: "quest", Label: "A Quest"}},
		9202: {{Kind: "quest", Label: "A Quest"}},
	}
	bySlot = candidatesBySlot(buildBandPool(items, idx, "rogue", 20, "horde", weights, 0, false, true).Scored)
	if idsOf(bySlot["trinket1"])[9200] || !idsOf(bySlot["trinket2"])[9200] {
		t.Errorf("trinkets, one rival with a use effect: trinket1 %v, trinket2 %v, want the rare charm in trinket2 only", idsOf(bySlot["trinket1"]), idsOf(bySlot["trinket2"]))
	}
}

// The nightly's published-file check (published_test.go) compares every
// published source with sourceFor's own. 1.60.1.70009's files were ranked
// before task L1.7 and are never re-ranked, so for that build alone the
// check accepts the one relabel L1.7 made there - a shared trash drop now
// named by its instance alone - and only where it applies: sourceFor's
// source is that build's shared-trash instance label and the published
// label names one of the creatures loot.json lists for the item (fix round
// 1, D3; the old clause accepted any "X: anything" against "X", and a
// "(rare)" clause no old file can meet). The active build must match
// sourceFor exactly.
func TestThePublishedSourceCheckAcceptsOnlyTaskL17sRelabelsInAnOlderBuild(t *testing.T) {
	gnomeregan := itemSource{Kind: "dungeon", Label: "Gnomeregan", SharedTrashCreatures: []string{"Caverndeep Burrower", "Caverndeep Ambusher"}}
	ironspine := itemSource{Kind: "dungeon", Label: "Scarlet Monastery: Ironspine (rare)", RareSpawn: true}
	instanceOnly := itemSource{Kind: "dungeon", Label: "Gnomeregan"} // a source's own item, not shared trash
	crafted := itemSource{Kind: "crafted", Label: "Tailoring"}
	cases := []struct {
		published string
		want      itemSource
		active    bool
		ok        bool
	}{
		{"Gnomeregan", gnomeregan, true, true},
		{"Gnomeregan: Caverndeep Burrower", gnomeregan, true, false},
		{"Gnomeregan: Caverndeep Burrower", gnomeregan, false, true},
		{"Gnomeregan: Caverndeep Ambusher", gnomeregan, false, true},
		{"Gnomeregan: Mekgineer Thermaplugg", gnomeregan, false, false},
		{"Gnomeregan: Caverndeep Burrower", instanceOnly, false, false},
		{"Scarlet Monastery: Ironspine", ironspine, false, false},
		{"Uldaman", gnomeregan, false, false},
		{"Leatherworking", crafted, false, false},
	}
	for _, c := range cases {
		msg := sourceLabelError("item", 9509, "Thing", c.published, c.want, c.active)
		if (msg == "") != c.ok {
			t.Errorf("sourceLabelError(%q, %+v, active=%v) = %q, want ok=%v", c.published, c.want, c.active, msg, c.ok)
		}
	}
}

// The ranker's per-band summary line names the slots that kept a rare drop
// out, in slot order (fix round 1, D6).
func TestTheRareSpawnSlotSummaryNamesEachSlotInOrder(t *testing.T) {
	cov := map[string]coverageRow{
		"back":    {Eligible: 3, Sourced: 3, RareSpawnSkipped: 1},
		"finger1": {Eligible: 4, Sourced: 4, RareSpawnSkipped: 2},
		"head":    {Eligible: 2, Sourced: 2},
	}
	if got, want := rareSpawnSlotSummary(cov), "back 1, finger1 2"; got != want {
		t.Errorf("rareSpawnSlotSummary = %q, want %q", got, want)
	}
	if got := rareSpawnSlotSummary(map[string]coverageRow{"head": {Eligible: 1}}); got != "no slot" {
		t.Errorf("rareSpawnSlotSummary(none) = %q, want \"no slot\"", got)
	}
}
