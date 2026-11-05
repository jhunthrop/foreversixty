// api/internal/guilds/loot_quests.go
//
// Onyxia's drops nobody ranks (guild loot spec 2026-11-05 §3.5, §3.6): the quest items, the bag and
// the skinned scale. They collapse into one line at the bottom of the Loot tab, and an officer can
// still award the quest items and the bag to a character the quest is for.
package guilds

import (
	"sort"
	"strings"
)

// Kinds of loot row (LootItemView.Kind).
const (
	lootKindGear    = "gear"    // a slotted drop the council ranks
	lootKindQuest   = "quest"   // a drop a quest takes
	lootKindBag     = "bag"     // Onyxia Hide Backpack
	lootKindSkinned = "skinned" // Scale of Onyxia: skinning loot, never a pick
)

// itemClassContainer is Item.ClassID 1, a bag.
const itemClassContainer = 1

// scaleOfOnyxia is skinned from Onyxia's corpse, not dropped.
const scaleOfOnyxia = 15410

// lootQuest is one quest item: the quest that takes it, who that quest is for ("" for anyone), and
// the content-phases gate that keeps the quest shut in Forever today ("" when it is open now).
// Copied from data/curated/raids/onyxia-quest-items.json, which records each classic-db source;
// loot_rules_test.go pins the two, and the gates to data/curated/content-phases.json.
type lootQuest struct {
	ItemID  int
	Quest   string
	Class   string
	Faction string
	Gate    string
}

var onyxiaQuestItems = []lootQuest{
	{18422, "Victory for the Horde", "", "horde", ""},
	{18423, "Victory for the Alliance", "", "alliance", ""},
	{18705, "A Proper String", "hunter", "", ""},
	// The Ahn'Qiraj Scepter chain: shut until the war effort opens.
	{21108, "The Only Prescription", "", "", "aq-war-effort"},
}

// questFor is itemID's quest row.
func questFor(itemID int) (lootQuest, bool) {
	for _, q := range onyxiaQuestItems {
		if q.ItemID == itemID {
			return q, true
		}
	}
	return lootQuest{}, false
}

// lootKind is which of the four kinds item is.
func lootKind(item lootItem) string {
	switch {
	case item.ItemID == scaleOfOnyxia:
		return lootKindSkinned
	case item.Slot != "":
		return lootKindGear
	case item.ItemClass == itemClassContainer:
		return lootKindBag
	default:
		return lootKindQuest
	}
}

// questLabel is the collapsed line's words for the quest items listed under one name ("Quest:
// Victory for the Alliance / Victory for the Horde", "Quest: A Proper String (Hunter)"); "Quest
// item" when no sourced quest open in Forever today names any of them. A gated quest is never
// named: a player cannot take it yet.
func questLabel(itemIDs []int) string {
	var names []string
	class := ""
	for _, id := range itemIDs {
		if q, ok := questFor(id); ok && q.Gate == "" {
			names = append(names, q.Quest)
			class = q.Class
		}
	}
	if len(names) == 0 {
		return "Quest item"
	}
	sort.Strings(names)
	label := "Quest: " + strings.Join(names, " / ")
	if class != "" {
		label += " (" + strings.ToUpper(class[:1]) + class[1:] + ")"
	}
	return label
}

// LootEligible is one character an officer can hand a quest item or the bag to, and the copy of
// the item they would get (the Alliance or the Horde Head of Onyxia).
type LootEligible struct {
	CharacterKey string `json:"character_key"`
	Name         string `json:"name"`
	Class        string `json:"class"`
	ItemID       int    `json:"item_id"`
}

// questEligible is every verified level-60 roster character the drop is for: the quest's class,
// the quest's faction (a character whose faction is unknown stays in), anyone for the bag. The
// skinned scale is for no one.
func questEligible(roster []RosterRow, item lootItem, kind string) []LootEligible {
	out := []LootEligible{}
	if kind == lootKindSkinned {
		return out
	}
	quest, _ := questFor(item.ItemID)
	for _, row := range roster {
		level, ok := rosterLevel(row)
		if !row.Verified || row.className == "" || !ok || level < lootLevelCap {
			continue
		}
		if quest.Class != "" && row.className != quest.Class {
			continue
		}
		if quest.Faction != "" && row.faction != "" && row.faction != quest.Faction {
			continue
		}
		out = append(out, LootEligible{CharacterKey: awardKeyOf(row), Name: row.Name, Class: row.className, ItemID: item.ItemID})
	}
	return out
}

// questEligibleAll is questEligible over every copy listed under one name, each character once,
// on the first copy that is for them.
func questEligibleAll(roster []RosterRow, group []lootItem) []LootEligible {
	out := []LootEligible{}
	seen := map[string]bool{}
	for _, item := range group {
		for _, e := range questEligible(roster, item, lootKindQuest) {
			if !seen[e.CharacterKey] {
				seen[e.CharacterKey] = true
				out = append(out, e)
			}
		}
	}
	return out
}
