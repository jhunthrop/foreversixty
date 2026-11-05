// api/internal/guilds/loot_candidates.go
//
// Who the Loot tab lists for a drop, and in what order (guild loot spec 2026-11-05 §3.2, §3.3). A
// character is a candidate only when every test passes, before any ranking: verified with a known
// class, at the drop's required level, and able to equip it (equip_rules.go). Then two tiers: tier
// 0, the character's BiS band names the item; tier 1, the spec wants the item's stats and it is an
// item-level upgrade over what they wear, or the slot is empty. A slotless drop never has
// candidates: there is no slot to compare against (loot_quests.go lists those). A character whose
// gear is unknown (no gear list on file: the roster reads the export only at gear consent, so
// this is gearOnFile) is no candidate at all: its slots are unknown, not empty.
package guilds

import (
	"math"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/spec"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// lootLevelCap is the level a quest drop's or the bag's award picker asks for, and what the page
// counts as "at 60". Every Onyxia drop with a slot requires it through its own required level.
const lootLevelCap = 60

// Gain units a BiS band's figures count (LootCandidate.GainUnit).
const (
	gainUnitDPS = "dps"
	gainUnitHPS = "hps"
	gainUnitTPS = "tps"
)

// rosterLevel is the character's level: the export's own, else the guild roster snapshot's; false
// when neither says.
func rosterLevel(row RosterRow) (int, bool) {
	if row.Level != nil && *row.Level > 0 {
		return *row.Level, true
	}
	if row.Addon != nil && row.Addon.Level != nil && *row.Addon.Level > 0 {
		return *row.Addon.Level, true
	}
	return 0, false
}

// awardKeyOf is the guild_characters key an award to row names.
func awardKeyOf(row RosterRow) string {
	if row.awardKey != "" {
		return row.awardKey
	}
	return row.CharacterKey
}

// countAtLevelCap counts the verified roster characters whose level on file is the cap.
func countAtLevelCap(roster []RosterRow) int {
	n := 0
	for _, row := range roster {
		if level, ok := rosterLevel(row); row.Verified && ok && level >= lootLevelCap {
			n++
		}
	}
	return n
}

// lootCandidates ranks every roster character who can take item, in two tiers (see the file
// comment); the rest are not listed at all.
func (s *Store) lootCandidates(roster []RosterRow, item lootItem) []LootCandidate {
	out := []LootCandidate{}
	if item.Slot == "" {
		return out
	}
	stats, isWeapon := s.dropStats(item), item.ItemClass == itemClassWeapon
	forTanks := hasAny(stats, tankStats)
	for _, row := range roster {
		level, ok := rosterLevel(row)
		if !row.Verified || row.className == "" || len(row.gear) == 0 || !ok ||
			level < item.RequiredLevel || !canEquip(row.className, level, item) {
			continue
		}
		slots := characterSlots(row, item, level)
		role := roleOf(row)
		cand := LootCandidate{CharacterKey: awardKeyOf(row), Name: row.Name, Attendance: row.Attendance,
			tankFirst: forTanks && role == RoleTank}
		if row.Class != nil {
			cand.Class = *row.Class
		}
		if row.Spec != nil {
			cand.Spec = *row.Spec
		}
		if match, ok := s.bisMatch(row, item, slots); ok {
			cand.Tier = 0
			cand.GainDps = match.verdict.GainDps
			if match.verdict.GainDps != nil {
				cand.GainUnit = &match.unit
			}
			cand.NotSimChecked = match.verdict.NotSimChecked
			cand.AlreadyEquivalent = !match.verdict.Upgrade
			_, worn := row.gear[match.slot]
			cand.FillsEmptySlot = !worn
			s.fillReplaces(&cand, row, match.slot)
			out = append(out, cand)
			continue
		}
		families := familiesFor(row.className, row.specStr, role)
		if !suitsFamily(stats, isWeapon, families, fightsInMeleeWith(row.className, row.specStr, role)) {
			continue
		}
		if delta, slot, ok := s.fallbackMatch(row, item, slots); ok {
			cand.Tier = 1
			cand.NotSimChecked = true
			cand.IlvlDelta = &delta
			_, worn := row.gear[slot]
			cand.FillsEmptySlot = !worn
			s.fillReplaces(&cand, row, slot)
			out = append(out, cand)
		}
	}
	sortLootCandidates(out)
	return out
}

// roleOf is the row's logged role, "" when none.
func roleOf(row RosterRow) string {
	if row.Role == nil {
		return ""
	}
	return *row.Role
}

// fillReplaces names the item row wears in slot, if any.
func (s *Store) fillReplaces(cand *LootCandidate, row RosterRow, slot string) {
	wornID, ok := row.gear[slot]
	if !ok {
		return
	}
	it, found := s.buildItemFor(row.className, wornID)
	if !found {
		return
	}
	cand.Replaces = it.Name
	cand.ReplacesItem = &LootWornItem{ItemID: wornID, Name: it.Name, Icon: it.Icon, Quality: it.Quality}
}

// characterSlots is the slots item can go in on this character: both rings, both trinkets, the
// off hand too for a one-hander when the character has Dual Wield. A unique item the character
// already holds can only go where it already is.
func characterSlots(row RosterRow, item lootItem, level int) []string {
	var slots []string
	switch {
	case item.Slot == "finger":
		slots = []string{"finger1", "finger2"}
	case item.Slot == "trinket":
		slots = []string{"trinket1", "trinket2"}
	case item.InventoryType == invOneHand && canDualWield(row.className, level):
		slots = []string{"main_hand", "off_hand"}
	default:
		slots = []string{item.Slot}
	}
	if item.Unique {
		for _, slot := range slots {
			if row.gear[slot] == item.ItemID {
				return []string{slot}
			}
		}
	}
	return slots
}

// bisSlotMatch is one candidate's tier-0 verdict: the character slot it is for, and the unit the
// band's figures count.
type bisSlotMatch struct {
	slot    string
	verdict bis.SlotVerdict
	unit    string
}

// bisMatch is the tier-0 test: row's spec band names item in one of slots (the pick or a listed
// alternative), scored the normal gear-gap way. With two slots named (a ring, a trinket) it is the
// verdict over the weaker worn item, the one the drop pushes out: an empty slot, else the larger
// gain, else the lower item level.
func (s *Store) bisMatch(row RosterRow, item lootItem, slots []string) (bisSlotMatch, bool) {
	if row.specStr == "" {
		return bisSlotMatch{}, false
	}
	band, hasBand := s.loadBandFor(row.className, row.specStr, row.faction)
	if !hasBand {
		return bisSlotMatch{}, false
	}
	bySlot := band.BySlot()
	var best bisSlotMatch
	found := false
	for _, slot := range slots {
		pick, named := bySlot[slot]
		if !named || pick.ItemID <= 0 || !namesItem(pick, item.ItemID) {
			continue
		}
		wornID, hasWorn := row.gear[slot]
		match := bisSlotMatch{slot: slot, verdict: bis.VerdictFor(pick, wornID, hasWorn), unit: gainUnitFor(band)}
		if !found || s.pushesOut(row, match, best) {
			best, found = match, true
		}
	}
	return best, found
}

// pushesOut reports whether a's worn item is weaker than b's, so a is the slot the drop takes.
func (s *Store) pushesOut(row RosterRow, a, b bisSlotMatch) bool {
	_, aWorn := row.gear[a.slot]
	_, bWorn := row.gear[b.slot]
	if aWorn != bWorn {
		return !aWorn
	}
	if ag, bg := gainOrLowest(a.verdict.GainDps), gainOrLowest(b.verdict.GainDps); ag != bg {
		return ag > bg
	}
	return s.wornItemLevel(row, a.slot) < s.wornItemLevel(row, b.slot)
}

// wornItemLevel is the item level of what row wears in slot; 0 when empty or unknown.
func (s *Store) wornItemLevel(row RosterRow, slot string) int {
	wornID, ok := row.gear[slot]
	if !ok {
		return 0
	}
	it, found := s.buildItemFor(row.className, wornID)
	if !found {
		return 0
	}
	return it.ItemLevel
}

// gainUnitFor is what a band's gains count: a tank band's Forever Score reads as threat, a healer
// band's as healing, anything else as damage.
func gainUnitFor(band bis.Band) string {
	switch {
	case band.ScoreUnit == bis.TankScoreUnit || band.Role == "tank":
		return gainUnitTPS
	case band.Role == "healer":
		return gainUnitHPS
	default:
		return gainUnitDPS
	}
}

// fallbackMatch is the tier-1 test: item's level outranks what row wears in the weakest of slots,
// or one of them is empty. A worn item whose level cannot be resolved (an id this class's item
// table does not carry) rules its slot out rather than guessing. Returns the delta and the slot.
func (s *Store) fallbackMatch(row RosterRow, item lootItem, slots []string) (int, string, bool) {
	bestDelta, bestSlot, bestEmpty, found := 0, "", false, false
	for _, slot := range slots {
		var worn *int
		wornID, hasWorn := row.gear[slot]
		if hasWorn {
			it, ok := s.buildItemFor(row.className, wornID)
			if !ok {
				continue
			}
			level := it.ItemLevel
			worn = &level
		}
		delta, ok := fallbackVerdict(item.ItemLevel, worn)
		if !ok {
			continue
		}
		empty := !hasWorn
		if !found || (empty && !bestEmpty) || (empty == bestEmpty && delta > bestDelta) {
			bestDelta, bestSlot, bestEmpty, found = delta, slot, empty, true
		}
	}
	return bestDelta, bestSlot, found
}

// fallbackVerdict is pure: given the drop's own item level and the worn item's level (nil
// for an empty slot), decides the fallback tier's own item-level-upgrade test.
func fallbackVerdict(dropItemLevel int, wornItemLevel *int) (delta int, qualifies bool) {
	if wornItemLevel == nil {
		return dropItemLevel, true
	}
	d := dropItemLevel - *wornItemLevel
	if d <= 0 {
		return 0, false
	}
	return d, true
}

// dropStats is the drop's stats from the first class item table that carries it (a class file's
// stats are the item's own, not the class's); nil for an item no class table carries, which reads
// as carrying no family stat.
func (s *Store) dropStats(item lootItem) map[string]int {
	for class := range classMaskBit {
		if it, ok := s.buildItemFor(class, item.ItemID); ok {
			return it.Stats
		}
	}
	return nil
}

// buildItemFor resolves itemID through classSlug's own per-class item table
// (data/builds/<build>/items/<classSlug>.json via api/internal/trees).
func (s *Store) buildItemFor(classSlug string, itemID int) (trees.Item, bool) {
	return s.ItemIn(s.BisBuild(), classSlug, itemID)
}

// ItemIn is itemID in version's per-class item table, the one the loot tab and the BiS page
// read (name, quality, item level, icon); false for an unknown build, class or item.
func (s *Store) ItemIn(version, classSlug string, itemID int) (trees.Item, bool) {
	if s.Trees == nil {
		return trees.Item{}, false
	}
	build, ok := s.Trees.Build(version)
	if !ok {
		return trees.Item{}, false
	}
	classID, ok := spec.New(build).ClassID(classSlug)
	if !ok {
		return trees.Item{}, false
	}
	return build.Item(classID, itemID)
}

// namesItem reports whether itemID is pick's own top choice or one of its listed
// alternatives - the membership test a loot drop must pass to be "this character's own
// candidate" at all.
func namesItem(pick bis.Slot, itemID int) bool {
	if pick.ItemID == itemID {
		return true
	}
	for _, alt := range pick.Alternatives {
		if alt.ItemID == itemID {
			return true
		}
	}
	return false
}

func sortLootCandidates(cands []LootCandidate) {
	for i := 1; i < len(cands); i++ {
		for j := i; j > 0 && lootCandidateLess(cands[j], cands[j-1]); j-- {
			cands[j], cands[j-1] = cands[j-1], cands[j]
		}
	}
}

// lootCandidateLess reports whether a ranks ahead of b: tier 0 (BiS-matched) always ahead
// of tier 1 (fallback); within a tier, a tank ahead on a drop with Defense, then a real upgrade
// over a worn item ahead of filling an empty slot; then within tier 0, gain_dps desc (a nil gain
// sorts last within the tier, never ahead of a measured one); within tier 1, ilvl_delta desc;
// either tier, attendance (present over the candidate's own nights, then present) desc breaks the
// remaining tie.
func lootCandidateLess(a, b LootCandidate) bool {
	if a.Tier != b.Tier {
		return a.Tier < b.Tier
	}
	if a.tankFirst != b.tankFirst {
		return a.tankFirst
	}
	if a.FillsEmptySlot != b.FillsEmptySlot {
		return !a.FillsEmptySlot
	}
	if a.Tier == 0 {
		ag, bg := gainOrLowest(a.GainDps), gainOrLowest(b.GainDps)
		if ag != bg {
			return ag > bg
		}
	} else {
		ad, bd := deltaOrZero(a.IlvlDelta), deltaOrZero(b.IlvlDelta)
		if ad != bd {
			return ad > bd
		}
	}
	if ar, br := attendanceRatio(a.Attendance), attendanceRatio(b.Attendance); ar != br {
		return ar > br
	}
	return a.Attendance.Present > b.Attendance.Present
}

func gainOrLowest(g *float64) float64 {
	if g == nil {
		return math.Inf(-1)
	}
	return *g
}

func deltaOrZero(d *int) int {
	if d == nil {
		return 0
	}
	return *d
}
