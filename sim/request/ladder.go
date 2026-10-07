package request

// Small, test-independent helpers for the rotation ladder
// (ladder_test.go): talent truncation, gear picking, cast tallying, and
// golden rendering. Nothing here reaches into rotations_smoke_test.go's
// _test.go-only symbols (repoRoot, activeBuild, readBuildJSON, ...) -
// every function below takes the paths and build strings it needs as
// plain arguments, so this file compiles into the ordinary package as
// well as the test binary.

import (
	"encoding/json"
	"fmt"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/specs"
	"github.com/wowsims/classic/sim/core/proto"
)

// ladderLevels is Phase 1a's ladder: the seven character levels every
// written rotation is measured at.
var ladderLevels = []int{10, 20, 30, 38, 40, 50, 60}

// ladderIterations and ladderSeed are the ladder's fixed run: enough
// iterations that a rarely-reached line still gets its turn, one seed so
// a failure reproduces.
const (
	ladderIterations = 300
	ladderSeed       = 1
)

// dpsRegressionTolerance is harness rule 4 (this wave's brief): a level
// scoring up to this much lower than the ladder's previous rung is the
// 300-iteration run's own noise, not a real regression -
// shaman-elemental's level 40 (46.0) against its level 38 (46.1) is
// what set this.
const dpsRegressionTolerance = 0.01

// ---------------------------------------------------------------------
// Talents: truncating a guide's level-60 build to a lower level.
//
// Moved to sim/leveling (lane bis-all) so the leveling BiS pipeline
// (sim/cmd/leveling-bis) can share the exact same rule instead of
// approximating its own. talentNode/talentTree/talentsFile/
// loadTalentTrees/guideBuildTalents/guideTalentTargets/
// ladderTalentString now live there as leveling.TalentNode/
// leveling.TalentTree/leveling.TalentsFile/leveling.LoadTalentTrees/
// leveling.GuideBuildTalents/leveling.GuideTalentTargets/
// leveling.LadderTalentString, with identical behaviour; only the call
// sites below were renamed to match.
// ---------------------------------------------------------------------

// ---------------------------------------------------------------------
// Gear: a weapon set the engine's item database knows, chosen by level.
// ---------------------------------------------------------------------

// buildItem is the fields ladder gear picking needs from one row of
// data/builds/<build>/items/<class>.json.
type buildItem struct {
	ID            int     `json:"id"`
	Slot          string  `json:"slot"`
	RequiredLevel int     `json:"required_level"`
	ItemLevel     int     `json:"item_level"`
	TwoHand       bool    `json:"two_hand"`
	Speed         float64 `json:"speed"`
	DamageMax     int     `json:"damage_max"`
	// Icon is only read by pickWandItem (harness rule 3): a ranged-slot
	// row's icon is how this item table distinguishes a real wand from
	// a bow, gun or thrown weapon, since every wand row's own
	// damage_max is 0 in this build (see pickWandItem's own comment).
	Icon string `json:"icon"`
	// WeaponClass and WeaponSubclass are joined in from
	// data/builds/<build>/items.json (see loadItemWeaponTypes): the
	// class item file (items/<class>.json) carries no weapon-type field
	// of its own at all, only slot/level/speed/damage. WeaponClass is
	// -1 for a row this build's items.json does not carry (should not
	// happen for anything ladderGear ever sees, since every id here
	// also has to be in simitems.json/loot.json to be picked).
	WeaponClass    int `json:"-"`
	WeaponSubclass int `json:"-"`
	// InventoryType is the client's slot code (13 = a one-hander that
	// fits either hand, 21 = main hand only, 22 = off hand only), also
	// joined from items.json; items/<class>.json files every
	// either-hand weapon under slot "main_hand".
	InventoryType int `json:"-"`
	// Stats is the row's own stats map (spell_power, attack_power,
	// strength, agility, ...), read directly off
	// data/builds/<build>/items/<class>.json rather than recomputed -
	// weaponScore (harness rule 2, this wave's brief) sums whichever of
	// these keys the picking spec's own WeightStats names.
	Stats map[string]float64 `json:"stats"`
	// DPS is the row's own precomputed damage-per-second
	// (damage_avg/speed, already carried by this build's item table),
	// weaponScore's tiebreak for a melee or hunter spec, whose damage
	// scales with the weapon's raw output rather than a stat on it.
	DPS float64 `json:"dps"`
}

// Weapon and shield item/subclass ids, per the client's own item
// taxonomy (data/builds/<build>/items.json's class_id/subclass_id),
// confirmed against this build's own rows (warrior.json's axes, maces,
// swords, polearms, staves, fist weapons; rogue.json's daggers; hunter's
// bows/guns; shaman.json's shields) rather than assumed from upstream
// Classic knowledge alone. class_id 2 is every real weapon; class_id 4
// subclass 6 is a shield (armor, not a weapon - pickShieldItem, not
// pickGearItem, picks these).
const (
	itemClassWeapon = 2
	itemClassArmor  = 4

	weaponAxe1H    = 0
	weaponAxe2H    = 1
	weaponBow      = 2
	weaponGun      = 3
	weaponMace1H   = 4
	weaponMace2H   = 5
	weaponPolearm  = 6
	weaponSword1H  = 7
	weaponSword2H  = 8
	weaponStaff    = 10
	weaponFist     = 13
	weaponDagger   = 15
	weaponThrown   = 16
	weaponCrossbow = 18

	armorSubclassShield = 6
)

// twoHandWeaponSubclasses is every subclass a "staff or two-hander" rule
// (druid-balance/feral, warrior-arms, paladin-retribution) accepts:
// every two-hand-only subclass in this table, plus the staff subclass
// (a staff is two_hand:true on every row checked, so TwoHand handedness
// alone would already select it, but the type list names it explicitly
// for callers that constrain type without constraining hand).
var twoHandWeaponSubclasses = []int{weaponAxe2H, weaponMace2H, weaponPolearm, weaponSword2H, weaponStaff}

type classItemsFile struct {
	Items []buildItem `json:"items"`
}

// loadClassItems reads one class's item rows - already restricted to
// items that class can equip, which is the "class allow" the design
// asks for - and joins in each row's weapon class/subclass from
// data/builds/<build>/items.json (loadItemWeaponTypes), since the class
// item file itself carries no type field.
func loadClassItems(repoRoot, build, class string) ([]buildItem, error) {
	path := filepath.Join(repoRoot, "data", "builds", build, "items", class+".json")
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ladder: reading %s: %w", path, err)
	}
	var f classItemsFile
	if err := json.Unmarshal(b, &f); err != nil {
		return nil, fmt.Errorf("ladder: parsing %s: %w", path, err)
	}
	types, err := loadItemWeaponTypes(repoRoot, build)
	if err != nil {
		return nil, err
	}
	for i := range f.Items {
		t, ok := types[f.Items[i].ID]
		if !ok {
			f.Items[i].WeaponClass = -1
			f.Items[i].WeaponSubclass = -1
			continue
		}
		f.Items[i].WeaponClass = t.ClassID
		f.Items[i].WeaponSubclass = t.SubclassID
		f.Items[i].InventoryType = t.InventoryType
	}
	return f.Items, nil
}

// itemTypeRow is the two fields loadItemWeaponTypes needs from one row
// of data/builds/<build>/items.json (the build's full item table, every
// class combined - unlike items/<class>.json, this file names each
// item's class_id/subclass_id, which is how a weapon's TYPE - dagger,
// sword, mace, axe, staff, fist, bow, gun, or a shield - is read).
type itemTypeRow struct {
	ID            int `json:"id"`
	ClassID       int `json:"class_id"`
	SubclassID    int `json:"subclass_id"`
	InventoryType int `json:"inventory_type"`
}

// inventoryTypeEitherHand is the client's INVTYPE_WEAPON: a one-hander
// that equips in either hand, which the class item files list under
// slot "main_hand" only.
const inventoryTypeEitherHand = 13

// loadItemWeaponTypes reads items.json (a bare array, unlike the other
// build files this ladder reads) into an id -> (class, subclass) map.
func loadItemWeaponTypes(repoRoot, build string) (map[int]itemTypeRow, error) {
	path := filepath.Join(repoRoot, "data", "builds", build, "items.json")
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ladder: reading %s: %w", path, err)
	}
	var rows []itemTypeRow
	if err := json.Unmarshal(b, &rows); err != nil {
		return nil, fmt.Errorf("ladder: parsing %s: %w", path, err)
	}
	out := make(map[int]itemTypeRow, len(rows))
	for _, r := range rows {
		out[r.ID] = r
	}
	return out, nil
}

type simItemsFile struct {
	Items []int `json:"items"`
}

// lootSourcesFile is the part of loot.json the ladder reads: every item
// id any source names (whether as a flat list, an instance's trash, or
// a boss drop), each source's kind (2026-09-28 quest-levels lane: which
// items are "crafted", for the item-level-proxy fallback), and the
// per-item quests map (the quest's own min_level, the real level gate
// for the 848 of loot.json's 1,140 quest-reward items whose own
// required_level is 0 in the client).
type lootSourcesFile struct {
	Sources []struct {
		Kind   string `json:"kind"`
		Items  []int  `json:"items"`
		Trash  []int  `json:"trash"`
		Bosses []struct {
			Items []int `json:"items"`
		} `json:"bosses"`
	} `json:"sources"`
	Quests map[string][]struct {
		MinLevel int `json:"min_level"`
		Level    int `json:"level"`
	} `json:"quests"`
}

// loadSourcedItemIDs reads loot.json and returns the ids with at least
// one known source. The ladder equips only such items: the client's item
// table also carries placeholder weapons no player can obtain ("Bland
// Dagger", "90 Epic Rogue Dagger" - required_level 0, no source), and
// once the client's damage curves gave those rows a damage value they
// outranked every real weapon by item level.
func loadSourcedItemIDs(repoRoot, build string) (map[int]bool, error) {
	path := filepath.Join(repoRoot, "data", "builds", build, "loot.json")
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ladder: reading %s: %w", path, err)
	}
	var f lootSourcesFile
	if err := json.Unmarshal(b, &f); err != nil {
		return nil, fmt.Errorf("ladder: parsing %s: %w", path, err)
	}
	out := map[int]bool{}
	for _, source := range f.Sources {
		for _, id := range source.Items {
			out[id] = true
		}
		for _, id := range source.Trash {
			out[id] = true
		}
		for _, boss := range source.Bosses {
			for _, id := range boss.Items {
				out[id] = true
			}
		}
	}
	return out, nil
}

// obtainableItemIDs is the ladder's pick pool: items the pinned engine
// knows (simitems.json) that also have a loot source (loot.json).
func obtainableItemIDs(repoRoot, build string) (map[int]bool, error) {
	known, err := loadSimItemIDs(repoRoot, build)
	if err != nil {
		return nil, err
	}
	sourced, err := loadSourcedItemIDs(repoRoot, build)
	if err != nil {
		return nil, err
	}
	out := make(map[int]bool, len(sourced))
	for id := range sourced {
		if known[id] {
			out[id] = true
		}
	}
	return out, nil
}

// loadRequiredLevelFloors reads loot.json once for the per-item level
// floor pickGearItem/pickShieldItem/pickWandItem apply on top of a
// buildItem's own RequiredLevel (leveling.EffectiveRequiredLevel;
// 2026-09-28 quest-levels lane) -- the client states required_level 0
// for a quest reward or crafted item far more often than not, gated
// instead by the quest's own min_level or, for a crafted item with no
// recipe skill level in loot.json (true of every crafted item today,
// since the engine fork's own database carries no such field), the
// item's own item_level run through leveling.ItemLevelProxyRequiredLevel
// -- the very fix Deadhead Blade (item 274271, item level 58,
// required_level 0) needed: pickGearItem used to equip it on a
// level-10 rogue.
//
// items is the class's own buildItem rows (already loaded by the
// caller for loadRequiredLevelFloors' item_level lookups on a crafted
// id); this function does not re-read items/<class>.json itself.
func loadRequiredLevelFloors(repoRoot, build string, items []buildItem) (map[int]int, error) {
	path := filepath.Join(repoRoot, "data", "builds", build, "loot.json")
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ladder: reading %s: %w", path, err)
	}
	var f lootSourcesFile
	if err := json.Unmarshal(b, &f); err != nil {
		return nil, fmt.Errorf("ladder: parsing %s: %w", path, err)
	}
	questFloor := make(map[int]int, len(f.Quests))
	for idStr, entries := range f.Quests {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			continue
		}
		levels := make([]int, len(entries))
		for i, e := range entries {
			levels[i] = leveling.QuestFloor(e.MinLevel, e.Level)
		}
		questFloor[id] = leveling.LowestFloor(levels)
	}
	crafted := map[int]bool{}
	for _, source := range f.Sources {
		if source.Kind != "crafted" {
			continue
		}
		for _, id := range source.Items {
			crafted[id] = true
		}
	}
	floors := make(map[int]int, len(items))
	for _, it := range items {
		if qf, ok := questFloor[it.ID]; ok && qf > 0 {
			floors[it.ID] = qf
		} else if crafted[it.ID] || it.RequiredLevel == 0 {
			// A recipe's product, or a row with no required level at all
			// (Forever-new items ship required_level 0 -- a data gap, the
			// same gate sim/cmd/leveling-bis applies), floors at its item
			// level's proxy.
			floors[it.ID] = leveling.ItemLevelProxyRequiredLevel(it.ItemLevel)
		}
	}
	return floors, nil
}

// loadSimItemIDs reads simitems.json: the item ids the pinned engine's
// own item database carries. items/<class>.json is already a subset of
// this set on every build checked, but a gear pick still filters against
// it rather than trusting that silently, so a future build where a
// class item outruns the sim database fails to pick the item instead of
// handing the engine an id it panics on (see sim/adapter/golden_test.go).
func loadSimItemIDs(repoRoot, build string) (map[int]bool, error) {
	path := filepath.Join(repoRoot, "data", "builds", build, "simitems.json")
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ladder: reading %s: %w", path, err)
	}
	var f simItemsFile
	if err := json.Unmarshal(b, &f); err != nil {
		return nil, fmt.Errorf("ladder: parsing %s: %w", path, err)
	}
	out := make(map[int]bool, len(f.Items))
	for _, id := range f.Items {
		out[id] = true
	}
	return out, nil
}

// handedness constrains a weapon pick to one-handed, two-handed, or
// either.
type handedness int

const (
	handAny handedness = iota
	handOne
	handTwo
)

// gearProfile is which weapon slots a spec's ladder character fills, and
// what handedness main_hand takes. The zero value is a caster: main hand
// only, any handedness (a caster's main_hand may be a one-hand dagger or
// a two-hand staff and either is fine, since no off-hand is ever filled
// alongside it), no off hand, no ranged weapon - though see Wand below,
// which every caster spec DOES set.
type gearProfile struct {
	MainHand handedness
	OffHand  bool
	Ranged   bool
	// Wand is harness rule 3 (this wave's brief): priest, mage and
	// warlock fill the ranged slot with a wand the same way melee fill
	// main_hand/off_hand, via pickWandItem rather than pickGearItem (a
	// caster's ranged pick needs a DIFFERENT weapon-reality check - see
	// pickWandItem's own comment for why). Mutually exclusive with
	// Ranged in practice (hunter uses Ranged for its bow instead), but
	// nothing enforces that beyond no spec setting both.
	Wand bool
	// Skip is druid-feral: "none" in the design's own words. A feral
	// character fights shapeshifted and this ladder does not model
	// weapon-DPS-through-form-conversion, so it picks no weapon at all
	// rather than equip one that does nothing for the rotation measured.
	Skip bool
	// SkipFromLevel narrows Skip to apply only at or above this level;
	// below it, gear is picked normally. druid-feral's own Skip is real
	// only from Cat Form's own learn level on (spellconst: spell 768,
	// level 20) - below that the character has no Cat Form yet (and
	// Bear Form's damage kit is not modeled either, see
	// data/curated/apl/druid-feral.json's own notes on its pre-20
	// Moonfire/Wrath fallback), so it fights unshifted and a real
	// weapon matters for its auto-attack the same way it does for any
	// other caster. Zero means "Skip applies at every level."
	SkipFromLevel int
	// Shield fills off_hand with a shield (pickShieldItem: armor,
	// subclass 6) instead of a weapon - shaman's elemental/resto rule
	// (MH + shield). Mutually exclusive with OffHand in practice.
	Shield bool
	// MainHandTypes/OffHandTypes/RangedTypes are this wave's per-spec
	// weapon TYPE preference (harness rule 1): the subclass ids
	// (weaponDagger, weaponSword1H, ...) a slot's pick is restricted
	// to, beyond handedness. Nil means "any real weapon type this
	// handedness/known/level/loot-source filter already allows" - the
	// pre-existing rule, unchanged for every spec this wave's brief
	// does not name a type for.
	MainHandTypes []int
	OffHandTypes  []int
	RangedTypes   []int
}

// ladderGearProfiles is every written spec whose weapon rule is not the
// bare caster default: the two-handers (warrior-arms,
// paladin-retribution), the dual-wielders (warrior-fury, rogue's three
// specs, and hunter - hunters can dual-wield a melee stat-stick beside
// their bow in this build), hunter's ranged bow, feral's "none",
// shaman-elemental's and shaman-enhancement's shield, druid-balance's
// staff/two-hander, and every caster's wand (harness rule 3:
// priest-shadow, mage's three specs and warlock's three specs, so
// OtherActionShoot/wand lines have a real weapon to resolve against,
// the same way melee always has).
//
// shaman-enhancement is NOT a dual-wielder (owner rule, Justin,
// 2026-09-30, bis-ranker-integrity-13's brief: shamans cannot dual
// wield in Forever) - it shares shaman-elemental's Shield profile
// instead, one imbued main-hand weapon plus a shield, never a second
// weapon.
//
// Weapon TYPE (this wave's harness rule 1, from the class item table's
// weapon subclass, joined in by loadItemWeaponTypes): assassination and
// subtlety want a main-hand dagger (Ambush/Backstab require one, and
// Mutilate is the spec's namesake finisher even though the pinned
// engine's own Mutilate does not itself require one - see
// sim/rogue/mutilate.go's ExtraCastCondition comment, quoted in this
// lane's report); combat wants sword or mace main-hand (its own guide:
// "Hack and Slash... treats axes the same way it always treated
// swords" - Forever gives Combat axes too, but the brief this table
// follows names sword/mace, so axe is left out here even though the
// guide suggests it is also viable; a follow-up can widen this once
// that's confirmed); shaman ele/enhancement/resto want a shield rather
// than dual-wielding (elemental and enhancement are the two written
// specs today);
// druid-balance (and feral, if it ever stops skipping gear) wants a
// staff or a two-hand weapon, matching what an unconstrained "highest
// item_level" pick already happened to choose for balance in this
// build (a two-hand mace) - named explicitly here so that stays true
// on a future build where it might not. Hunter's ranged slot is
// restricted to bow/gun, per the brief's own wording, even though this
// build's ranged table also carries crossbows and thrown weapons a
// hunter could equip.
//
// Assassination's off hand is a dagger too (Mutilate needs one in each
// hand). The class item files list every either-hand one-hander under
// slot "main_hand" (the off_hand slot holds only off-hand-only weapons
// and held items), so fitsSlot lets the off-hand pick draw from those
// rows: 150 one-hand daggers this build's rogues can wear, none of them
// filed as off_hand. The two picks may not be the same item.
var ladderGearProfiles = map[string]gearProfile{
	"warrior-arms":        {MainHand: handTwo},
	"warrior-fury":        {MainHand: handOne, OffHand: true},
	"paladin-retribution": {MainHand: handTwo},
	// shaman-enhancement: owner rule (Justin, 2026-09-30,
	// bis-ranker-integrity-13's brief) - shamans cannot dual wield in
	// Forever, so enhancement's off hand is a shield, the same
	// pickShieldItem pool shaman-elemental already draws from, not a
	// second weapon.
	"shaman-enhancement": {MainHand: handOne, Shield: true},
	"shaman-elemental":   {MainHand: handOne, Shield: true},
	"druid-feral":        {Skip: true, SkipFromLevel: 20, MainHand: handAny},
	"druid-balance":      {MainHand: handAny, MainHandTypes: twoHandWeaponSubclasses},
	"rogue-assassination": {MainHand: handOne, OffHand: true,
		MainHandTypes: []int{weaponDagger}, OffHandTypes: []int{weaponDagger}},
	"rogue-combat": {MainHand: handOne, OffHand: true,
		MainHandTypes: []int{weaponSword1H, weaponMace1H}},
	"rogue-subtlety": {MainHand: handOne, OffHand: true,
		MainHandTypes: []int{weaponDagger}},
	"hunter-beast-mastery": {MainHand: handOne, OffHand: true, Ranged: true, RangedTypes: []int{weaponBow, weaponGun}},
	"hunter-marksmanship":  {MainHand: handOne, OffHand: true, Ranged: true, RangedTypes: []int{weaponBow, weaponGun}},
	"hunter-survival":      {MainHand: handOne, OffHand: true, Ranged: true, RangedTypes: []int{weaponBow, weaponGun}},
	"priest-shadow":        {Wand: true},
	"mage-arcane":          {Wand: true},
	"mage-fire":            {Wand: true},
	"mage-frost":           {Wand: true},
	"warlock-affliction":   {Wand: true},
	"warlock-demonology":   {Wand: true},
	"warlock-destruction":  {Wand: true},
}

// weaponRelevantStats restricts a spec's own WeightStats
// (data/curated/specs.json, via sim/specs) to the stat keys an actual
// weapon row's own stats map (data/builds/<build>/items/<class>.json)
// can carry. WeightStats also lists throughput stats - crit, hit,
// spell_haste, melee_haste, spell_penetration - that a weapon row in
// this build's item table essentially never carries, so scoring on them
// would only add noise to weaponScore below.
var weaponRelevantStats = map[string]bool{
	"spell_power": true, "healing_power": true,
	"attack_power": true, "ranged_attack_power": true, "feral_attack_power": true,
	"strength": true, "agility": true, "intellect": true,
}

// specWeaponPrimaryStats is spec's WeightStats filtered to
// weaponRelevantStats, in the order specs.json lists them - a caster's
// spell_power/healing_power sorts ahead of the strength/agility its own
// stats map can also carry, though weaponScore below sums the whole
// list rather than reading only the first entry.
func specWeaponPrimaryStats(spec string) []string {
	var out []string
	for _, s := range specs.ByKey[spec].WeightStats {
		if weaponRelevantStats[s] {
			out = append(out, s)
		}
	}
	return out
}

// specWeaponDPSMatters is harness rule 2 (this wave's brief): a melee or
// hunter spec's damage scales with the weapon's own raw DPS (Auto
// Attack, White damage windows, weapon-damage-effect abilities), where a
// caster or healer spec's does not - spellpower is what does there - so
// only a spec whose ReferenceStat is a physical one credits a
// candidate's DPS field in weaponScore.
func specWeaponDPSMatters(spec string) bool {
	switch specs.ByKey[spec].ReferenceStat {
	case "attack_power", "ranged_attack_power":
		return true
	default:
		return false
	}
}

// weaponScore is one candidate's rank for a spec's ladder gear pick. The
// order the two throughput numbers are compared in depends on
// weightDPS (specWeaponDPSMatters(spec)):
//
//   - A caster or healer spec (weightDPS false, dps always left at 0 by
//     scoreWeapon) ranks purely on primaryStat (the sum of every stat
//     point on specWeaponPrimaryStats' list) - a candidate with even one
//     point of the spec's own primary stat always beats a stat-less item
//     at a higher item level, the gap this lane's brief names
//     ("item-level only ... generally ignores the spec's stats").
//   - A melee or hunter spec (weightDPS true) ranks on dps FIRST: white
//     damage and most weapon-damage-coefficient abilities scale off the
//     weapon's own raw output, which a str/agi/AP stat point does not
//     convert 1:1 against (a 35-strength, 86 DPS axe is a clear upgrade
//     over a 62-attack-power, 54 DPS one - primaryStat alone would get
//     that backwards). primaryStat only breaks a tie in DPS.
//
// itemLevel and finally id are the last, fully deterministic tiebreaks,
// so two candidates that also tie on the throughput numbers above still
// resolve without depending on the source file's row order.
type weaponScore struct {
	primaryStat float64
	dps         float64
	itemLevel   int
	id          int
	weightDPS   bool
}

func scoreWeapon(it buildItem, primaryStats []string, weightDPS bool) weaponScore {
	var stat float64
	for _, s := range primaryStats {
		stat += it.Stats[s]
	}
	dps := 0.0
	if weightDPS {
		dps = it.DPS
	}
	return weaponScore{primaryStat: stat, dps: dps, itemLevel: it.ItemLevel, id: it.ID, weightDPS: weightDPS}
}

// betterThan reports whether a is weaponScore's preferred candidate over
// b: see weaponScore's own comment for the tiebreak order.
func (a weaponScore) betterThan(b weaponScore) bool {
	if a.weightDPS {
		if a.dps != b.dps {
			return a.dps > b.dps
		}
		if a.primaryStat != b.primaryStat {
			return a.primaryStat > b.primaryStat
		}
	} else if a.primaryStat != b.primaryStat {
		return a.primaryStat > b.primaryStat
	}
	if a.itemLevel != b.itemLevel {
		return a.itemLevel > b.itemLevel
	}
	return a.id < b.id
}

// pickGearItem is weaponScore's best candidate in slot with
// leveling.EffectiveRequiredLevel(RequiredLevel, floors[id]) <= level
// (2026-09-28 quest-levels lane: NOT RequiredLevel alone -- floors is
// loadRequiredLevelFloors' result, the quest/crafted level gate a
// candidate's own RequiredLevel very often does not state, e.g.
// Deadhead Blade, item 274271, item level 58, required_level 0, that
// used to be pickGearItem's level-10 pick), restricted to hand's
// handedness, to allowed weapon subclasses when the caller names any
// (harness rule 1's per-spec weapon TYPE table - nil means every type
// this filter set already allows), and to items this build's sim
// database knows (known). A candidate must be a real weapon: WeaponClass ==
// itemClassWeapon (joined in from items.json by loadClassItems - a
// held-in-off-hand item is class 4, not a weapon at all, even on a row
// that happens to carry nonzero speed/damage fields), speed > 0 (which
// is what keeps a shield out of a dual-wielder's off hand: an off_hand
// row with speed == 0 is armor, not a weapon, in this item table) and
// damage_max > 0. The last check exists because this build's item table
// carries a large number of quality-3 "rare" weapon rows with
// damage_min = damage_max = 0 and no required_level (Bland Dagger, item
// 24071, is one) - unfinished or placeholder rows, not a character's
// real choice, and equipping one of them is what made the engine hang
// mid-sim rather than simulate a zero-damage weapon (see the report).
// primaryStats/weightDPS are specWeaponPrimaryStats(spec)/
// specWeaponDPSMatters(spec) - passed in rather than a spec string so
// this function (and its existing fixture tests, which pass nil/false
// for "no stat preference") stays independent of the specs package.
// fitsSlot is whether a row can be equipped in slot: its own slot, or,
// for the off hand, an either-hand one-hander the class item file
// lists under main_hand (a rogue's off-hand dagger is one of those;
// the off_hand slot itself holds only off-hand-only weapons and held
// items).
func fitsSlot(it buildItem, slot string) bool {
	if it.Slot == slot {
		return true
	}
	return slot == "off_hand" && it.Slot == "main_hand" && it.InventoryType == inventoryTypeEitherHand && !it.TwoHand
}

// withoutItem is items minus every id already in gear.
func withoutItem(items []buildItem, gear []api.GearSlot) []buildItem {
	if len(gear) == 0 {
		return items
	}
	taken := make(map[int]bool, len(gear))
	for _, g := range gear {
		taken[g.ItemID] = true
	}
	out := make([]buildItem, 0, len(items))
	for _, it := range items {
		if !taken[it.ID] {
			out = append(out, it)
		}
	}
	return out
}

func pickGearItem(items []buildItem, known map[int]bool, floors map[int]int, slot string, level int, hand handedness, allowed []int, primaryStats []string, weightDPS bool) (buildItem, bool) {
	var best buildItem
	var bestScore weaponScore
	found := false
	for _, it := range items {
		if !fitsSlot(it, slot) || it.WeaponClass != itemClassWeapon || it.Speed <= 0 || it.DamageMax <= 0 || !known[it.ID] {
			continue
		}
		if leveling.EffectiveRequiredLevel(it.RequiredLevel, floors[it.ID]) > level {
			continue
		}
		if hand == handOne && it.TwoHand {
			continue
		}
		if hand == handTwo && !it.TwoHand {
			continue
		}
		if len(allowed) > 0 && !containsInt(allowed, it.WeaponSubclass) {
			continue
		}
		score := scoreWeapon(it, primaryStats, weightDPS)
		if !found || score.betterThan(bestScore) {
			best, bestScore, found = it, score, true
		}
	}
	return best, found
}

// containsInt reports whether needle is in haystack; the per-spec
// weapon TYPE lists this file builds (assassination's dagger MH+OH,
// combat's sword/mace MH, hunter's bow/gun ranged, ...) are always
// short enough that a linear scan is simpler than a set.
func containsInt(haystack []int, needle int) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

// pickShieldItem is ladderGear's shaman-elemental/resto counterpart to
// pickGearItem (harness rule 1's "MH + shield" rule): the highest
// item_level off_hand row that is a shield - itemClassArmor, subclass
// armorSubclassShield, in items.json's own taxonomy, since a shield is
// armor, not a weapon, and pickGearItem's speed>0/damage_max>0 "is a
// weapon" checks would always reject one.
func pickShieldItem(items []buildItem, known map[int]bool, floors map[int]int, level int) (buildItem, bool) {
	var best buildItem
	found := false
	for _, it := range items {
		if it.Slot != "off_hand" || it.WeaponClass != itemClassArmor || it.WeaponSubclass != armorSubclassShield || !known[it.ID] {
			continue
		}
		if leveling.EffectiveRequiredLevel(it.RequiredLevel, floors[it.ID]) > level {
			continue
		}
		if !found || it.ItemLevel > best.ItemLevel || (it.ItemLevel == best.ItemLevel && it.ID < best.ID) {
			best, found = it, true
		}
	}
	return best, found
}

// ladderGear is a level-appropriate weapon set for spec: main_hand,
// and off_hand/ranged where the spec's profile calls for them. Every
// other slot is bare, as the design's Phase 1a asks ("a bare character
// shows the rotation, not the raid"). floors is
// loadRequiredLevelFloors' result (2026-09-28 quest-levels lane): the
// quest/crafted level gate a candidate's own RequiredLevel does not
// state, so a quest-reward or crafted weapon is never equipped below
// the level it is actually obtainable at. primaryStats/weightDPS
// (harness rule 2, this wave's brief) are the same for every slot this
// spec fills - main_hand, off_hand and ranged all score against the
// spec's own stats, not a per-slot preference.
func ladderGear(items []buildItem, known map[int]bool, floors map[int]int, spec string, level int) []api.GearSlot {
	profile := ladderGearProfiles[spec]
	if profile.Skip && (profile.SkipFromLevel == 0 || level >= profile.SkipFromLevel) {
		return nil
	}
	primaryStats := specWeaponPrimaryStats(spec)
	weightDPS := specWeaponDPSMatters(spec)
	var gear []api.GearSlot
	if it, ok := pickGearItem(items, known, floors, "main_hand", level, profile.MainHand, profile.MainHandTypes, primaryStats, weightDPS); ok {
		gear = append(gear, api.GearSlot{Slot: "main_hand", ItemID: it.ID})
	}
	if profile.OffHand {
		// The main hand's item is excluded so a class with one standout
		// weapon does not hold two copies of it.
		if it, ok := pickGearItem(withoutItem(items, gear), known, floors, "off_hand", level, handOne, profile.OffHandTypes, primaryStats, weightDPS); ok {
			gear = append(gear, api.GearSlot{Slot: "off_hand", ItemID: it.ID})
		}
	}
	if profile.Shield {
		if it, ok := pickShieldItem(items, known, floors, level); ok {
			gear = append(gear, api.GearSlot{Slot: "off_hand", ItemID: it.ID})
		}
	}
	if profile.Ranged {
		if it, ok := pickGearItem(items, known, floors, "ranged", level, handAny, profile.RangedTypes, primaryStats, weightDPS); ok {
			gear = append(gear, api.GearSlot{Slot: "ranged", ItemID: it.ID})
		}
	}
	if profile.Wand {
		if it, ok := pickWandItem(items, known, floors, level); ok {
			gear = append(gear, api.GearSlot{Slot: "ranged", ItemID: it.ID})
		}
	}
	return gear
}

// formatGear renders a gear set for the golden: "bare" for none,
// otherwise each slot's item id, sorted so the line never reorders
// itself between regenerations.
func formatGear(gear []api.GearSlot) string {
	if len(gear) == 0 {
		return "bare"
	}
	parts := make([]string, len(gear))
	for i, g := range gear {
		parts[i] = fmt.Sprintf("%s:%d", g.Slot, g.ItemID)
	}
	sort.Strings(parts)
	return strings.Join(parts, " ")
}

// ---------------------------------------------------------------------
// Spell data: which abilities exist, when they are learned, and which
// of them deal damage.
// ---------------------------------------------------------------------

// spellRankEntry is one rank of one named ability, from
// data/builds/<build>/spellranks.json.
type spellRankEntry struct {
	ID    int `json:"id"`
	Level int `json:"level"`
	Rank  int `json:"rank"`
}

type spellRanksFile struct {
	Classes map[string]map[string][]spellRankEntry `json:"classes"`
}

// loadSpellRanks reads the whole spellranks.json once; callers index
// f.Classes[class] themselves so the (large) file is read only once per
// ladder run rather than once per spec.
func loadSpellRanks(repoRoot, build string) (spellRanksFile, error) {
	path := filepath.Join(repoRoot, "data", "builds", build, "spellranks.json")
	b, err := os.ReadFile(path)
	if err != nil {
		return spellRanksFile{}, fmt.Errorf("ladder: reading %s: %w", path, err)
	}
	var f spellRanksFile
	if err := json.Unmarshal(b, &f); err != nil {
		return spellRanksFile{}, fmt.Errorf("ladder: parsing %s: %w", path, err)
	}
	return f, nil
}

// junkAbilityName matches spellranks.json keys that are not a named
// ability at all: item-set-bonus rows key on a version-stamped label
// like "1.60.0 - Item - Tier 1 - Warrior - Arms/Fury 2P Bonus - Haste".
var junkAbilityName = regexp.MustCompile(`^\d`)

// rankTier is every id that shares one rank number of one ability, and
// the level that rank is learned at - the ladder's own mirror of
// sim/internal/spellranks' buildRankChain (unexported there, so this is
// a second, read-only copy of the same grouping rather than a call into
// it), used only for the informational "learned but unused" list
// (rule 3). A tier's level is the LEVEL OF THE FIRST ROW SEEN at that
// rank, matching that package exactly: a later duplicate at the same
// rank number and a different level (typically 0, an NPC/internal
// copy) widens the tier's id set without ever lowering its level.
type rankTier struct {
	Level int
	Rank  int
	IDs   []int
}

// classAbilities is one class's spellranks.json rows grouped by ability
// name into rank tiers. A name whose rows are ALL rank 0 has NO entry
// here at all - exactly like sim/internal/spellranks' table, which is
// how Bloodrage (two rank-0 rows, no real progression) and Judgement
// (three rank-0 rows) are represented: they are not "ranked spells"
// this ladder - or the engine's own rewrite - tracks at all.
type classAbilities struct {
	Tiers map[string][]rankTier
}

func buildClassAbilities(all spellRanksFile, class string) classAbilities {
	out := classAbilities{Tiers: map[string][]rankTier{}}
	for name, entries := range all.Classes[class] {
		if junkAbilityName.MatchString(name) {
			continue
		}
		levelByRank := map[int]int{}
		idsByRank := map[int][]int{}
		var ranks []int
		for _, e := range entries {
			if e.Rank <= 0 {
				continue
			}
			if _, seen := levelByRank[e.Rank]; !seen {
				levelByRank[e.Rank] = e.Level
				ranks = append(ranks, e.Rank)
			}
			idsByRank[e.Rank] = append(idsByRank[e.Rank], e.ID)
		}
		if len(ranks) == 0 {
			continue
		}
		sort.Ints(ranks)
		tiers := make([]rankTier, len(ranks))
		for i, r := range ranks {
			tiers[i] = rankTier{Level: levelByRank[r], Rank: r, IDs: idsByRank[r]}
		}
		out.Tiers[name] = tiers
	}
	return out
}

// learnedTierAtLevel is the highest-rank tier a character of level has
// reached, from an ability's own rank-ascending tier list (tiers are
// already sorted by rank, which in this table is also ascending by
// level in every case checked, so walking forward and keeping the last
// tier at or under level picks the same tier
// sim/internal/spellranks.highestLearnedIn would, walking from the top
// down).
func learnedTierAtLevel(tiers []rankTier, level int) (rankTier, bool) {
	best := -1
	for i, t := range tiers {
		if t.Level <= level {
			best = i
		}
	}
	if best == -1 {
		return rankTier{}, false
	}
	return tiers[best], true
}

// abilityNameByID is every tracked id's ability name, for labeling a
// zero_casts violation (rule 1 resolves the id sim/internal/spellranks'
// own HighestLearnedSpellID would cast, which carries no name). An id
// this table does not track - an unranked ability like Bloodrage, or
// one spellranks.json's generator dropped entirely - has no entry, and
// callers fall back to the bare id.
func abilityNameByID(abilities classAbilities) map[int]string {
	out := map[int]string{}
	for name, tiers := range abilities.Tiers {
		for _, tier := range tiers {
			for _, id := range tier.IDs {
				out[id] = name
			}
		}
	}
	return out
}

// spellEffectConst is one effect of one spell, from
// data/builds/<build>/spellconst/<class>.json.
type spellEffectConst struct {
	Effect   int     `json:"effect"`
	Aura     int     `json:"aura"`
	Amount   float64 `json:"amount"`
	PeriodMS int32   `json:"period_ms"`
}

type spellConstEntry struct {
	Name       string             `json:"name"`
	GCDMS      int32              `json:"gcd_ms"`
	CooldownMS int32              `json:"cooldown_ms"`
	DurationMS int32              `json:"duration_ms"`
	Effects    []spellEffectConst `json:"effects"`
}

type spellConstFile struct {
	Spells map[string]spellConstEntry `json:"spells"`
}

// loadSpellConst reads one class's spell constants, keyed by spell id.
func loadSpellConst(repoRoot, build, class string) (map[int]spellConstEntry, error) {
	path := filepath.Join(repoRoot, "data", "builds", build, "spellconst", class+".json")
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ladder: reading %s: %w", path, err)
	}
	var f spellConstFile
	if err := json.Unmarshal(b, &f); err != nil {
		return nil, fmt.Errorf("ladder: parsing %s: %w", path, err)
	}
	out := make(map[int]spellConstEntry, len(f.Spells))
	for idStr, entry := range f.Spells {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			continue
		}
		out[id] = entry
	}
	return out, nil
}

// isDamageEffect is the design's own rule for "this effect deals
// damage": type 2 (school damage), type 6 with aura 3 (periodic
// damage) or aura 226 (periodic trigger spell - the ground-effect
// pattern Consecration, Blizzard, Rain of Fire, Hurricane and Volley
// all use in this build's spellconst: an effect 6/aura 226 entry
// applies an aura whose own tick casts the real damage spell, so the
// container spell itself never carries an effect 2 or an aura-3 row of
// its own, and the old rule missed it entirely - a learned Consecration
// showed up as a false "learned but unused"), or type 121, 31 or 58
// (the three weapon-damage effect types). It is deliberately narrow -
// Heroic Strike's own bonus-weapon-damage effect is type 17 and does
// not match, which the report calls out as a gap in the rule rather
// than something this code quietly widens to cover.
func isDamageEffect(e spellEffectConst) bool {
	if e.Effect == 2 {
		return true
	}
	if e.Effect == 6 && (e.Aura == 3 || e.Aura == 226) {
		return true
	}
	switch e.Effect {
	case 121, 31, 58:
		return true
	}
	return false
}

// isDamageSpellID reports whether id, looked up in consts, deals damage
// by isDamageEffect. An id the table does not carry is not a damage
// spell as far as this ladder can tell.
func isDamageSpellID(consts map[int]spellConstEntry, id int) bool {
	entry, ok := consts[id]
	if !ok {
		return false
	}
	for _, e := range entry.Effects {
		if isDamageEffect(e) {
			return true
		}
	}
	return false
}

// ---------------------------------------------------------------------
// Cast tally.
// ---------------------------------------------------------------------

// CastSpellCounts sums every plain-spell action's casts across the run,
// keyed by the base spell id (a tag or rank variant of the same spell -
// see sim/adapter's ActionName - folds into the one id, since what this
// ladder asks is "was this ability cast at all", not which of its
// variants). Items, pets and "other" actions are not spells and are
// left out.
func CastSpellCounts(player *proto.UnitMetrics) map[int]int64 {
	out := map[int]int64{}
	for _, action := range player.GetActions() {
		id := action.GetId()
		if id == nil {
			continue
		}
		spellID, ok := id.RawId.(*proto.ActionID_SpellId)
		if !ok {
			continue
		}
		var casts int64
		for _, target := range action.GetTargets() {
			casts += int64(target.GetCasts())
		}
		if casts == 0 {
			continue
		}
		out[int(spellID.SpellId)] += casts
	}
	return out
}

// castTally is one action and how many times the player performed it
// per iteration - the ladder's own version of rotations_smoke_test.go's
// castCount, normalized against THIS run's own iteration count rather
// than that file's smokeIterations constant (200; the ladder runs 300 -
// reusing castSet here would silently divide by the wrong number).
type castTally struct {
	Name         string
	PerIteration float64
}

// ladderCastSet is every action the player actually performed, most
// first, normalized by iterations. The same source rotations_smoke_test.go's
// castSet reads (player.GetActions()), computed against this run's own
// iteration count.
func ladderCastSet(player *proto.UnitMetrics, iterations int) []castTally {
	var out []castTally
	for _, action := range player.GetActions() {
		var casts int32
		for _, target := range action.GetTargets() {
			casts += target.GetCasts()
		}
		if casts == 0 {
			continue
		}
		_, name := adapter.ActionName(action.GetId())
		out = append(out, castTally{Name: name, PerIteration: float64(casts) / float64(iterations)})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].PerIteration != out[j].PerIteration {
			return out[i].PerIteration > out[j].PerIteration
		}
		return out[i].Name < out[j].Name
	})
	return out
}

// ladderDistinctSpellCasts counts how many distinct "spell:"-named
// actions were cast at all - autoattacks and resource gains the engine
// records alongside them are not one, which is what makes this the
// design's rule-4 check ("no cast but auto-attack").
func ladderDistinctSpellCasts(tallies []castTally) int {
	n := 0
	for _, c := range tallies {
		if strings.HasPrefix(c.Name, "spell:") {
			n++
		}
	}
	return n
}

// ladderTopCasts renders the n most-performed actions for the golden.
func ladderTopCasts(tallies []castTally, n int) string {
	if n > len(tallies) {
		n = len(tallies)
	}
	parts := make([]string, n)
	for i, c := range tallies[:n] {
		parts[i] = fmt.Sprintf("%s=%.1f", c.Name, c.PerIteration)
	}
	return strings.Join(parts, ", ")
}

// extractCastSpellIDs walks a decoded APL JSON tree (any of
// map[string]any, []any, or a scalar) and collects the spell id of
// every castSpell action it finds, at any depth - a compound action
// (sequence, resetSequence, ...) can hold cast actions nested inside
// it, so this does not stop at the top level of prepullActions or
// priorityList.
func extractCastSpellIDs(node any, out map[int]bool) {
	switch v := node.(type) {
	case map[string]any:
		if cs, ok := v["castSpell"].(map[string]any); ok {
			if sid, ok := cs["spellId"].(map[string]any); ok {
				if raw, ok := sid["spellId"].(float64); ok {
					out[int(raw)] = true
				}
			}
		}
		for _, child := range v {
			extractCastSpellIDs(child, out)
		}
	case []any:
		for _, child := range v {
			extractCastSpellIDs(child, out)
		}
	}
}

// authoredCastSpellIDs is every spell id a curated rotation's own JSON
// (its "rotation" field, the same source apl/<spec>.apl.json embeds)
// names in a castSpell action - the lines the rotation lane wrote down,
// at the ranks it wrote them at, before any level rewrite.
func authoredCastSpellIDs(rotation json.RawMessage) (map[int]bool, error) {
	var tree any
	if err := json.Unmarshal(rotation, &tree); err != nil {
		return nil, fmt.Errorf("ladder: parsing the curated rotation: %w", err)
	}
	out := map[int]bool{}
	extractCastSpellIDs(tree, out)
	return out, nil
}

// ---------------------------------------------------------------------
// The curated file's ladder-relevant fields.
// ---------------------------------------------------------------------

// expectedIdleEntry is one line data/curated/apl/<spec>.json's
// expected_idle array may carry: an engine-resolvable spell id this
// ladder's bare run cannot cast, and why.
type expectedIdleEntry struct {
	ID     int    `json:"id"`
	Reason string `json:"reason"`
}

// ladderCurated is the half of data/curated/apl/<spec>.json the ladder
// reads: rotations_smoke_test.go's own curatedRotation reads spec/
// state/inert for its own purpose; this is a second, ladder-only
// reader carrying expected_idle beside inert (the new key this task
// adds, per the design) and the raw rotation JSON authoredCastSpellIDs
// needs.
type ladderCurated struct {
	Spec         string              `json:"spec"`
	State        string              `json:"state"`
	Inert        []int               `json:"inert"`
	ExpectedIdle []expectedIdleEntry `json:"expected_idle"`
	Rotation     json.RawMessage     `json:"rotation"`
}

func loadLadderCurated(repoRoot, spec string) (ladderCurated, error) {
	path := filepath.Join(repoRoot, "data", "curated", "apl", spec+".json")
	b, err := os.ReadFile(path)
	if err != nil {
		return ladderCurated{}, fmt.Errorf("ladder: reading %s: %w", path, err)
	}
	var doc ladderCurated
	if err := json.Unmarshal(b, &doc); err != nil {
		return ladderCurated{}, fmt.Errorf("ladder: parsing %s: %w", path, err)
	}
	if doc.Spec != spec {
		return ladderCurated{}, fmt.Errorf("ladder: %s declares spec %q", path, doc.Spec)
	}
	return doc, nil
}

// ---------------------------------------------------------------------
// Golden rendering.
// ---------------------------------------------------------------------

// ladderRow is one level's line in a spec's golden table.
type ladderRow struct {
	Level         int
	Talents       string
	Gear          string
	DPS           float64
	DistinctCasts int
	TopCasts      string
	Unresolved    []string
}

// unusedEntry is one (level, ability) the class had learned and could
// deal damage with, that this level's cast set never touched.
// Informational: the design's rule 3, not a failure.
type unusedEntry struct {
	Level int
	Name  string
	ID    int
}

// ladderKitConsumes is the class kit the ladder character carries
// (sim/leveling.KitConsumes: a rogue's poisons from 20, shaman-
// enhancement's weapon imbues throughout), shared with the leveling
// BiS ranker so both measure the same character. It takes the full
// spec slug, not the bare class, because the kit is a spec property.
func ladderKitConsumes(spec string, level int) []string { return leveling.KitConsumes(spec, level) }

// ladderRulesHeader documents, once, the rules every golden in this
// directory was generated under. It is repeated verbatim into each
// file rather than referenced, so a reader of one golden never has to
// go find the rule that produced it.
const ladderRulesHeader = `Rules this ladder runs under (Phase 1a,
docs/superpowers/specs/2026-09-28-rotation-accuracy-program-design.md):

- Levels: 10, 20, 30, 38, 40, 50, 60. 300 iterations, seed 1, the
  default encounter (a stationary target three levels above the
  character).
- Talents: the guide's level-60 FS1 build code
  (web/src/content/guides/<class>/<spec>.md), truncated to level-9
  points (0 below level 10). Each talent's guide-assigned rank is read
  by the talent's own stable id against the client build the guide
  names, then re-resolved to that talent's row in the ACTIVE build
  (data/builds/<active>/talents/<class>.json) - a talent the active
  build no longer carries is dropped rather than misaligning every
  digit after it. Points are spent walking the active build's trees
  top row down, the spec's own tree first, then the other two in the
  build code's own order (0, 1, 2, skipping the spec's own). A row the
  budget runs out before reaching is simply left at 0. If the guide
  build itself spends fewer than level-9 points, the remainder is left
  unspent.
- Gear: main_hand always, off_hand for the classes that dual-wield in
  this build (rogue, warrior-fury, shaman-enhancement, hunter),
  ranged for hunter only. Each slot picks the highest item_level
  weapon (speed > 0, so a shield never fills an off hand; damage_max >
  0, so an unfinished/placeholder weapon row this build's item table
  still carries - equipping one hangs the engine mid-sim rather than
  simulating a zero-damage weapon, see the report - is never picked)
  with required_level <= the character's level, from
  data/builds/<active>/items/<class>.json filtered to
  data/builds/<active>/simitems.json's known ids, honoring the spec's
  handedness (warrior-arms and paladin-retribution two-hand only;
  warrior-fury, rogue and shaman-enhancement one-hand only for both
  hands). druid-feral picks no weapon at all. Every caster spec
  (priest-shadow, mage's three specs, warlock's three specs) fills
  ranged with a wand instead: the same item table's ranged rows whose
  icon names them a real wand (every wand row's own damage_max is 0 in
  this build, unlike a bow or gun, so the melee pick's damage_max > 0
  check is replaced by that icon check rather than dropped), so
  OtherActionShoot/wand lines have something to resolve against.
  Every other slot is bare. Consumables: none (see the potion rule
  below).
- DPS regression: each level's DPS is compared against the ladder's own
  PREVIOUS rung (not literally level-10, since the ladder's own gaps
  are uneven - 30 to 38 is 8 levels, 38 to 40 is 2), tolerating up to a
  1% drop as the 300-iteration run's own noise (shaman-elemental's
  level 40, 46.0 vs a level-38 46.1, is exactly this). A level scoring
  more than 1% lower than the rung before it is a violation.
- Unresolved: an id the engine's ComputeStats warns it cannot resolve,
  with three standing exceptions before anything counts as a
  violation: (1) data/curated/apl/<spec>.json's own inert array names
  it; (2) it is the potion action ({OtherID: 13}) - the ladder
  character carries no consumes (a rogue's poisons, class kit from
  level 20, are the one exception: ladderKitConsumes), so this can
  never resolve, at any level, any spec; (3) it is a talent-granted spell
  (data/builds/<build>/talents/<class>.json's own "ranks[].spell_id")
  and the ladder's own truncated build (ladderTalentString's budget
  walk) has spent zero points on that talent at this level - expected
  right up until the level this ladder's approximation of the guide's
  build actually reaches that talent's row, a violation only once the
  build HAS spent points on it and the id still will not resolve.
  Anything else is a violation.
- Zero casts: one of the curated rotation's own castSpell lines, resolved
  to the id sim/internal/spellranks.HighestLearnedSpellID says the
  engine's OWN rank rewrite actually casts at this level (not a second,
  approximate copy of that resolution - this ladder calls the same
  function sim/request's rewriteRotationRanks calls), that never fires
  in the run - unless the engine could not resolve that id at all
  (already counted as unresolved) or data/curated/apl/<spec>.json's
  expected_idle array names the LINE'S AUTHORED id with a reason.
  HighestLearnedSpellID returning "not learned" means the engine's own
  rewrite already dropped the line before the request was built, which
  is not a violation to report twice.
- Learned but unused (informational, not a violation): every damage
  ability (spellconst effect 2, 6 with aura 3, or 121/31/58 - see
  ladder.go's isDamageEffect) the class has learned by this level that
  this level's cast set never touched, regardless of whether the
  rotation names it at all. An ability whose spellranks.json rows are
  ALL rank 0 (Bloodrage, Judgement - a single always-known ability, not
  a rank progression) is not tracked by this rule at all, the same way
  it is invisible to the engine's own rank rewrite. Heroic
  Strike-shaped bonus-weapon-damage effects (type 17) are NOT covered
  by the damage-effect rule and so never appear here even when
  genuinely unused - see the report for specs where that matters.
- Strict failures (FOREVER_LADDER_STRICT=1) are the four rules above;
  see the run's own "Violations" section below for what this file's own
  run found.
`

// renderLadderGolden is the markdown for one spec's golden file.
func renderLadderGolden(spec string, rows []ladderRow, unused []unusedEntry, violations []string) []byte {
	var b strings.Builder
	fmt.Fprintf(&b, "# %s rotation ladder\n\n", spec)
	b.WriteString(ladderRulesHeader)
	b.WriteString("\n## Ladder\n\n")
	b.WriteString("| Level | Talents | Gear | DPS | Distinct casts | Top casts | Unresolved |\n")
	b.WriteString("|---|---|---|---|---|---|---|\n")
	for _, r := range rows {
		unresolved := "-"
		if len(r.Unresolved) > 0 {
			unresolved = strings.Join(r.Unresolved, ", ")
		}
		topCasts := r.TopCasts
		if topCasts == "" {
			topCasts = "-"
		}
		fmt.Fprintf(&b, "| %d | %s | %s | %.1f | %d | %s | %s |\n",
			r.Level, r.Talents, r.Gear, r.DPS, r.DistinctCasts, topCasts, unresolved)
	}

	b.WriteString("\n## Learned but unused (informational)\n\n")
	if len(unused) == 0 {
		b.WriteString("None at any ladder level.\n")
	} else {
		sorted := append([]unusedEntry(nil), unused...)
		sort.Slice(sorted, func(i, j int) bool {
			if sorted[i].Level != sorted[j].Level {
				return sorted[i].Level < sorted[j].Level
			}
			if sorted[i].Name != sorted[j].Name {
				return sorted[i].Name < sorted[j].Name
			}
			return sorted[i].ID < sorted[j].ID
		})
		lastLevel := -1
		for _, u := range sorted {
			if u.Level != lastLevel {
				fmt.Fprintf(&b, "\n### Level %d\n\n", u.Level)
				lastLevel = u.Level
			}
			fmt.Fprintf(&b, "- %s (spell %d)\n", u.Name, u.ID)
		}
	}

	b.WriteString("\n## Violations found in this run\n\n")
	if len(violations) == 0 {
		b.WriteString("None.\n")
	} else {
		sorted := append([]string(nil), violations...)
		sort.Strings(sorted)
		for _, v := range sorted {
			fmt.Fprintf(&b, "- %s\n", v)
		}
	}

	return []byte(b.String())
}

// ---------------------------------------------------------------------
// Harness rule 1 (this wave's brief): a castSpell id the engine warns
// it cannot resolve is EXPECTED, not a violation, when either (a) it is
// a talent-granted spell and the ladder's own truncated build has zero
// points in that talent at this level, or (b) it is a spellranks.json
// rank learned above this level's band (the rank-rewrite
// - sim/internal/spellranks.HighestLearnedSpellID, the very function
// rewriteRotationRanks calls before a request ever reaches the engine -
// already drops a CAST of such an id; this id reaching the engine as an
// unresolved reference at all means it arrived some other way, a
// condition value rewriteRankedSpellIDs's castKeys deliberately leaves
// as-authored, per rotation_ranks.go's own comment).
//
// Both readers below are deliberately separate, read-only copies of
// data this file's talent-truncation code (talentNode, talentTree,
// loadTalentTrees, ladderTalentString - lines ~40-230, owned by lane
// bis-all's leveling-package migration) already reads, in the same
// spirit ladderCurated is "a second, ladder-only reader" beside
// rotations_smoke_test.go's curatedRotation: neither touches a
// protected line, and both use the exported talentTree/talentNode
// shapes that section already defines.
// ---------------------------------------------------------------------

// talentRankSpellID is one rank of one talent's own spell id, from
// data/builds/<build>/talents/<class>.json's "ranks" array - the field
// loadTalentTrees's talentNode does not carry (it only reads id, tier,
// column and max_rank, everything ladderTalentString's truncation walk
// needs and nothing more).
type talentRankSpellID struct {
	SpellID int `json:"spell_id"`
}

type talentSpellNode struct {
	ID    int                 `json:"id"`
	Ranks []talentRankSpellID `json:"ranks"`
}

type talentSpellTree struct {
	Talents []talentSpellNode `json:"talents"`
}

type talentSpellFile struct {
	Trees []talentSpellTree `json:"trees"`
}

// loadTalentSpellIDs reads build's talent file for class a second time,
// keyed by every rank's OWN spell id rather than by (tier, column): the
// map from a spell id any rank of any talent grants to that talent's
// own stable node id, which is what a warned action's id and a truncated
// build's per-talent point count (ladderTalentPoints) share as a common
// key.
func loadTalentSpellIDs(repoRoot, build, class string) (map[int]int, error) {
	path := filepath.Join(repoRoot, "data", "builds", build, "talents", class+".json")
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ladder: reading %s: %w", path, err)
	}
	var f talentSpellFile
	if err := json.Unmarshal(b, &f); err != nil {
		return nil, fmt.Errorf("ladder: parsing %s: %w", path, err)
	}
	out := map[int]int{}
	for _, tree := range f.Trees {
		for _, node := range tree.Talents {
			for _, r := range node.Ranks {
				if r.SpellID != 0 {
					out[r.SpellID] = node.ID
				}
			}
		}
	}
	return out, nil
}

// expandTalentSpellIDs widens loadTalentSpellIDs' map (every id
// talents.json's own ranks[].spell_id lists) to every id spellranks.json's
// ability-name grouping (buildClassAbilities' Tiers, keyed by name) shares
// with one of those literal ids. talents.json names only ONE id per talent
// node - conventionally the ability's rank 1 - even when the talent gates a
// whole multi-rank spellranks.json chain: Arcane Blast is talents.json node
// with ranks[].spell_id 400574 only, but spellranks.json's own "Arcane
// Blast" grouping carries five ranks (400574, 1239696, 1239697, 1239699,
// 1239700), and sim/mage/arcane_blast.go's registerArcaneBlastSpell gates
// every one of those five ranks on the SAME Talents.ArcaneBlast bool, not
// just rank 1. Without this widening, harness rule 1's "zero points on
// this talent" exemption only ever recognizes rank 1's id: a ladder level
// where the truncated build has NOT yet spent the point but IS already
// past a higher rank's own RequiredLevel (Arcane Blast rank 2, 1239696,
// RequiredLevel 30) reports a false unresolved_id instead of being
// silently absorbed the way rank 1 already is. Checked against a second
// case with the same shape (Shadowburn: talents.json names only 17877,
// spellranks.json's "Shadowburn" chain carries six ranks 17877..18871, and
// sim/warlock/shadowburn.go's registerShadowBurnSpell gates every rank on
// Talents.Shadowburn) before generalizing rather than one-offing Arcane
// Blast alone.
func expandTalentSpellIDs(talentSpellIDs map[int]int, abilities classAbilities) map[int]int {
	out := make(map[int]int, len(talentSpellIDs))
	for id, node := range talentSpellIDs {
		out[id] = node
	}
	for _, tiers := range abilities.Tiers {
		node, found := 0, false
		for _, tier := range tiers {
			for _, id := range tier.IDs {
				if n, ok := talentSpellIDs[id]; ok {
					node, found = n, true
				}
			}
		}
		if !found {
			continue
		}
		for _, tier := range tiers {
			for _, id := range tier.IDs {
				if _, already := out[id]; !already {
					out[id] = node
				}
			}
		}
	}
	return out
}

// ladderTalentPoints is ladderTalentString's own budget-spending walk
// (lines ~159-213), recomputed here to return points-per-talent-node
// rather than a rendered digit string - the shape rule 1 needs and the
// string does not carry. Any future change to the truncation rule
// belongs in ladderTalentString; this copy exists only because that
// function is protected for lane bis-all's migration and returns the
// wrong shape for this reader anyway.
func ladderTalentPoints(activeTrees []leveling.TalentTree, targets map[int]int, ownTreeIndex, level int) map[int]int {
	budget := level - 9
	if budget < 0 {
		budget = 0
	}
	order := make([]int, 0, len(activeTrees))
	order = append(order, ownTreeIndex)
	for i := range activeTrees {
		if i != ownTreeIndex {
			order = append(order, i)
		}
	}

	points := map[int]int{}
	for _, ti := range order {
		if ti < 0 || ti >= len(activeTrees) {
			continue
		}
		for _, node := range activeTrees[ti].Talents {
			if budget <= 0 {
				break
			}
			target := targets[node.ID]
			if target > node.MaxRank {
				target = node.MaxRank
			}
			give := target
			if give > budget {
				give = budget
			}
			points[node.ID] = give
			budget -= give
		}
	}
	return points
}

// idLearnLevel is every id spellranks.json tracks, for ANY class, mapped
// to the level its own row names - rule 1's "a rank whose
// spellranks.json level is above the band's level" half. Unlike
// buildClassAbilities (which drops rank-0 rows and groups by ability
// name for the "learned but unused" report), this keeps every row so a
// bare warned id - which carries no ability name - can still be looked
// up directly.
func idLearnLevel(ranks spellRanksFile, class string) map[int]int {
	out := map[int]int{}
	for name, entries := range ranks.Classes[class] {
		if junkAbilityName.MatchString(name) {
			continue
		}
		for _, e := range entries {
			// A rank <= 0 row is the same "not a rank progression"
			// signal buildClassAbilities filters on (Bloodrage,
			// Judgement - a single always-known ability, not a chain
			// with a learn level): sim/internal/spellranks' own
			// HighestLearnedSpellID treats such an id as untracked and
			// always learned, so keeping it here would disagree with
			// that package about an id that was never a real
			// above-band rank at all - EXCEPT for the small,
			// engine-verified allowlist idLearnLevelOverrides mirrors
			// below, which sim/internal/spellranks' own
			// singleTierLevelOverrides (same allowlist, same reasoning
			// - see that package's comment) also carries, precisely so
			// the two stay in agreement.
			if e.Rank <= 0 {
				continue
			}
			if prev, ok := out[e.ID]; !ok || e.Level < prev {
				out[e.ID] = e.Level
			}
		}
	}
	for id, level := range idLearnLevelOverrides[class] {
		out[id] = level
	}
	return out
}

// idLearnLevelOverrides is this file's own read-only mirror of
// sim/internal/spellranks' singleTierLevelOverrides (that package's own
// comment has the full reasoning and the per-id engine citations): a
// handful of ids whose spellranks.json rows are all rank <= 0 but that
// the pinned engine still gates behind a real character level. Kept in
// sync by hand rather than imported, the same "separate, read-only
// copy" shape this file's own comment on loadTalentSpellIDs and
// idLearnLevel already documents for the rest of this exemption logic -
// a second copy of a grouping already computed elsewhere, not a live
// dependency on it.
var idLearnLevelOverrides = map[string]map[int]int{
	"druid":  {5217: 24},
	"hunter": {2643: 18, 3045: 26},
}

// warnedSpellID pulls the bare spell id out of an engine warning
// action's rendering (spellAction's own "{SpellID: %d}", with no ", Tag:
// N" suffix - a tagged id is a rank-variant reference the talent and
// rank tables below do not key on, so it is left for the ordinary
// unresolved_id violation rather than guessed at).
var warnedSpellIDPattern = regexp.MustCompile(`^\{SpellID: (\d+)\}$`)

func warnedSpellID(action string) (int, bool) {
	m := warnedSpellIDPattern.FindStringSubmatch(action)
	if m == nil {
		return 0, false
	}
	id, err := strconv.Atoi(m[1])
	if err != nil {
		return 0, false
	}
	return id, true
}

// potionUnresolvedAction is harness rule 2 (this wave's brief): the
// ladder character carries no consumes, so the engine's own potion
// action can never resolve. Expected at every level, every spec.
const potionUnresolvedAction = "{OtherID: 13}"

// pickWandItem is ladderGear's caster counterpart to pickGearItem
// (harness rule 3): the highest item_level ranged-slot item whose icon
// marks it a real wand (data/builds/<build>/items/<class>.json's own
// "inv_wand*" icon naming - every wand row checked in this build has
// damage_max == 0, unlike a bow or gun, so pickGearItem's damage_max > 0
// weapon check would reject every candidate; this reader keeps that
// check for the OTHER reason it exists - the Bland Dagger-shaped
// placeholder rows this table also carries - by restricting the
// candidate pool to wand-icon rows first, rather than dropping the
// check) with required_level <= level, from a build's own item table
// filtered to its simitems.json-known ids, the same two filters
// pickGearItem applies.
func pickWandItem(items []buildItem, known map[int]bool, floors map[int]int, level int) (buildItem, bool) {
	var best buildItem
	found := false
	for _, it := range items {
		if it.Slot != "ranged" || it.DamageMax <= 0 || !known[it.ID] {
			continue
		}
		if !strings.HasPrefix(it.Icon, "inv_wand") {
			continue
		}
		if leveling.EffectiveRequiredLevel(it.RequiredLevel, floors[it.ID]) > level {
			continue
		}
		if !found || it.ItemLevel > best.ItemLevel || (it.ItemLevel == best.ItemLevel && it.ID < best.ID) {
			best, found = it, true
		}
	}
	return best, found
}
