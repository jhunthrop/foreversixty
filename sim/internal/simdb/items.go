package simdb

// The item rows the planner reads.
//
// sim/bulk decides which sims run, and to do that it needs to know what
// an item is: where it goes, who can wear it, whether two of it may be
// worn at once. Those facts are in the same embedded database the
// engine loads, and this file is the view of them that is NOT a
// protobuf.
//
// That matters because sim/request and sim/adapter are the only two
// packages in the product that touch one, and a third would be a third
// place the engine's schema leaks into our own vocabulary. So every
// enum here is a string in OUR words - the planner's slot names, class
// slugs, "two_hand" rather than HandTypeTwoHand - and the mapping
// happens once, here.

import (
	"strings"
	"sync"

	"github.com/jhunthrop/foreversixty/sim/internal/strcase"
	"github.com/wowsims/classic/sim/core/proto"
)

// Hand types, in our words.
const (
	HandNone = ""
	HandOne  = "one_hand"
	HandTwo  = "two_hand"
	HandMain = "main_hand"
	HandOff  = "off_hand"
)

// Faction restrictions, in our words. The empty string is "either".
const (
	FactionAny      = ""
	FactionAlliance = "alliance"
	FactionHorde    = "horde"
)

// Item is one row, in the planner's vocabulary.
type Item struct {
	ID   int
	Name string
	// Slots is every gear slot this item can be equipped in, in the
	// order the planner tries them: a ring is finger1 then finger2, a
	// one-hander is main_hand then off_hand. Empty means nothing can
	// equip it, which is most of a client's item table.
	Slots []string
	// ArmorType is "cloth", "leather", "mail", "plate" or "".
	ArmorType string
	// WeaponType is "sword", "axe", "shield" and the rest, or "".
	WeaponType string
	// HandType is one of the Hand constants.
	HandType string
	// Classes are the class slugs allowed to use it. Empty means every
	// class.
	Classes []string
	// RequiredLevel is 0 when the item has none.
	RequiredLevel int
	// Unique says only one may be equipped at a time.
	Unique bool
	// Faction is one of the Faction constants.
	Faction string
	// SuffixOptions are the random suffixes this item can roll.
	SuffixOptions []int
}

// slotsByType is the engine's own eligibility table (core.database.go's
// itemTypeToSlotsMap), in our slot names. Weapons are absent because a
// weapon's slot cannot be decided from its type alone.
var slotsByType = map[proto.ItemType][]string{
	proto.ItemType_ItemTypeHead:     {"head"},
	proto.ItemType_ItemTypeNeck:     {"neck"},
	proto.ItemType_ItemTypeShoulder: {"shoulder"},
	proto.ItemType_ItemTypeBack:     {"back"},
	proto.ItemType_ItemTypeChest:    {"chest"},
	proto.ItemType_ItemTypeWrist:    {"wrist"},
	proto.ItemType_ItemTypeHands:    {"hands"},
	proto.ItemType_ItemTypeWaist:    {"waist"},
	proto.ItemType_ItemTypeLegs:     {"legs"},
	proto.ItemType_ItemTypeFeet:     {"feet"},
	proto.ItemType_ItemTypeFinger:   {"finger1", "finger2"},
	proto.ItemType_ItemTypeTrinket:  {"trinket1", "trinket2"},
	proto.ItemType_ItemTypeRanged:   {"ranged"},
}

// slotsByHand is the weapon half of the same table.
var slotsByHand = map[proto.HandType][]string{
	proto.HandType_HandTypeTwoHand:  {"main_hand"},
	proto.HandType_HandTypeMainHand: {"main_hand"},
	proto.HandType_HandTypeOffHand:  {"off_hand"},
	proto.HandType_HandTypeOneHand:  {"main_hand", "off_hand"},
}

var handNames = map[proto.HandType]string{
	proto.HandType_HandTypeUnknown:  HandNone,
	proto.HandType_HandTypeOneHand:  HandOne,
	proto.HandType_HandTypeTwoHand:  HandTwo,
	proto.HandType_HandTypeMainHand: HandMain,
	proto.HandType_HandTypeOffHand:  HandOff,
}

// factionNames maps proto.SimItem's own nested FactionRestriction enum
// (common.proto's SimItem.FactionRestriction, redeclared there rather
// than imported from UIItem because ui.proto imports common.proto and
// common.proto cannot import back) into our words.
var factionNames = map[proto.SimItem_FactionRestriction]string{
	proto.SimItem_FACTION_RESTRICTION_UNSPECIFIED:   FactionAny,
	proto.SimItem_FACTION_RESTRICTION_ALLIANCE_ONLY: FactionAlliance,
	proto.SimItem_FACTION_RESTRICTION_HORDE_ONLY:    FactionHorde,
}

// trimmed is an enum value name in our words: the value name without
// its enum's prefix, lower snake case. ArmorTypeMail is "mail",
// WeaponTypeTwoHandedSword is "two_handed_sword", ClassWarrior is
// "warrior". The unknown value of every one of these enums is 0 and
// reads as "", which is what "this item has no armor type" means.
func trimmed(name, prefix string) string {
	name = strings.TrimPrefix(name, prefix)
	if name == "Unknown" || name == "" {
		return ""
	}
	return strcase.Snake(name)
}

// itemFrom maps one engine row into our vocabulary.
func itemFrom(row *proto.SimItem) Item {
	slots := slotsByType[row.Type]
	if row.Type == proto.ItemType_ItemTypeWeapon {
		slots = slotsByHand[row.HandType]
	}
	classes := make([]string, 0, len(row.ClassAllowlist))
	for _, c := range row.ClassAllowlist {
		classes = append(classes, trimmed(c.String(), "Class"))
	}
	suffixes := make([]int, 0, len(row.RandomSuffixOptions))
	for _, s := range row.RandomSuffixOptions {
		suffixes = append(suffixes, int(s))
	}
	out := Item{
		ID:            int(row.Id),
		Name:          row.Name,
		Slots:         slots,
		ArmorType:     trimmed(row.ArmorType.String(), "ArmorType"),
		WeaponType:    trimmed(row.WeaponType.String(), "WeaponType"),
		HandType:      handNames[row.HandType],
		RequiredLevel: int(row.RequiredLevel),
		Unique:        row.Unique,
		Faction:       factionNames[row.FactionRestriction],
	}
	if len(classes) > 0 {
		out.Classes = classes
	}
	if len(suffixes) > 0 {
		out.SuffixOptions = suffixes
	}
	return out
}

// items is the whole table, indexed, built once. The map is never
// handed out: Lookup returns a copy, whose only slices are the shared
// read-only ones built above, so a caller cannot retune the build's
// item table for the rest of the process.
var items = sync.OnceValues(func() (map[int]Item, error) {
	db, err := load()
	if err != nil {
		return nil, err
	}
	out := make(map[int]Item, len(db.Items))
	for _, row := range db.Items {
		out[int(row.Id)] = itemFrom(row)
	}
	return out, nil
})

// Lookup is one item of the active build. A miss is an item the build
// does not carry, which sim/bulk refuses rather than sims.
func Lookup(id int) (Item, bool) {
	table, err := items()
	if err != nil {
		return Item{}, false
	}
	it, ok := table[id]
	return it, ok
}

// Len is how many items the build carries. It is what a caller uses to
// prove the table loaded before it starts asking about ids.
func Len() (int, error) {
	table, err := items()
	if err != nil {
		return 0, err
	}
	return len(table), nil
}
