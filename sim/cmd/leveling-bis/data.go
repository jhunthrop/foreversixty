// Command leveling-bis ranks the best gear a leveling character can
// wear at a handful of level bands, for one spec, from the simulator's
// own stat weights and Top Gear verification. See
// docs/superpowers/specs/2026-09-28-leveling-bis-design.md for the
// design this prototypes, and
// /private/tmp/claude-501/.../scratchpad/accuracy/lane-bis-proto.md
// for this lane's exact brief.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// flatItem is one row of data/builds/<build>/items.json: the fields
// every item carries regardless of class, including the two the
// per-class files do not (faction_restriction, and the class/subclass
// ids armor proficiency needs).
type flatItem struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Quality            int    `json:"quality"`
	ItemLevel          int    `json:"item_level"`
	RequiredLevel      int    `json:"required_level"`
	ClassID            int    `json:"class_id"`
	SubclassID         int    `json:"subclass_id"`
	InventoryType      int    `json:"inventory_type"`
	Suffixes           []int  `json:"suffixes"`
	FactionRestriction string `json:"faction_restriction"`
	// SupersededBy is the data-followups-4 lane's own key (this lane's
	// brief, item 2): a legacy client row Forever kept in items.json
	// (the vanilla id, still shipped by the client) alongside its own
	// re-itemised copy under a different id (the quartermaster/vendor/
	// loot source actually hands out the NEW id) - the pinned example
	// being "Grand Marshal's Stave" 18873 (superseded_by 234571) and
	// "High Warlord's War Staff" 18874 (superseded_by 234549). Non-zero
	// means this exact row is never a candidate: loadCandidates below
	// excludes it, the same way a class-file/flat-file id mismatch
	// already gets excluded with a note, rather than silently ranking a
	// row nothing in the game actually awards any more. The row itself
	// stays in items.json (data-followups-4's own words: "for tooltips
	// of anything a player already owns"), so this field's presence,
	// not the row's absence, is what "never a candidate" reads.
	SupersededBy int `json:"superseded_by,omitempty"`
}

// classItem is one row of data/builds/<build>/items/<class>.json: the
// planner's resolved view, already filtered to items that class can
// equip, with per-item stats already resolved through the item-level
// curves (data/pipeline/normalize/item_curves.py; see data/README.md).
type classItem struct {
	ID            int                `json:"id"`
	Name          string             `json:"name"`
	Slot          string             `json:"slot"`
	Quality       int                `json:"quality"`
	RequiredLevel int                `json:"required_level"`
	ItemLevel     int                `json:"item_level"`
	Armor         int                `json:"armor"`
	Stats         map[string]float64 `json:"stats"`
	DamageMin     float64            `json:"damage_min"`
	DamageMax     float64            `json:"damage_max"`
	Speed         float64            `json:"speed"`
	DPS           float64            `json:"dps"`
	TwoHand       bool               `json:"two_hand"`
	EffectText    string             `json:"effect_text"`
	SetID         *int               `json:"set_id"`
	Unique        bool               `json:"unique"`
	// WeaponType is a weapon row's own bow/gun/crossbow/wand/thrown (or
	// melee axe/mace/polearm/sword/staff/fist/dagger) kind - this
	// lane's brief (bis-ranker-integrity-6), item 9, corrected per the
	// controller's own direct note: subclass_id (flatItem, items.json)
	// is unreliable for a ranged row, so the classicdb-fidelity lane
	// (merged origin/main@0ace46ba) adds this field instead, by this
	// exact JSON key (data/pipeline/normalize/gear.py's
	// weapon_type_for), to the PER-CLASS row (items/<class>.json,
	// classItem - to_gear_item's own construction), not the flat
	// items.json row - this field lives here, not on flatItem, for
	// that reason. Empty for a non-weapon row, and for any row the
	// data has not been rebuilt with this field yet -
	// restrictRangedByProficiency (weapon_requirements.go) is the one
	// place this is read, and treats empty as "unknown", never as a
	// wand.
	WeaponType string `json:"weapon_type"`
	// SupersededBy mirrors flatItem.SupersededBy's own doc: the
	// data-followups-4 lane marks the legacy row on "the flat and
	// per-class rows (models + JSON)" (this lane's brief, item 2), so
	// loadCandidates below reads it defensively off whichever row
	// actually carries it - a class-file regeneration that only
	// updates one of the two rows still excludes the item, rather than
	// only working when both happen to agree.
	SupersededBy int `json:"superseded_by,omitempty"`
}

type classItemFile struct {
	Build     string      `json:"build"`
	ClassSlug string      `json:"class_slug"`
	Items     []classItem `json:"items"`
}

// slotAliases mirrors web/src/lib/planner/rules.ts's SLOT_ALIASES: the
// planner's "finger"/"trinket" collapse to two numbered slots apiece.
var slotAliases = map[string][]string{
	"finger":  {"finger1", "finger2"},
	"trinket": {"trinket1", "trinket2"},
}

// plannerSlots is every slot planner name a class item's .Slot may
// expand to, in the order sim/api.GearSlots (and the engine's own
// equipment array) uses.
func plannerSlots(slot string) []string {
	if aliased, ok := slotAliases[slot]; ok {
		return aliased
	}
	return []string{slot}
}

// candidate is one item merged from items.json and items/<class>.json,
// the unit eligible/score/pick all work on.
type candidate struct {
	ID            int
	Name          string
	Quality       int
	RequiredLevel int
	ItemLevel     int
	ClassID       int
	SubclassID    int
	// WeaponType is classItem.WeaponType carried through unchanged -
	// see that field's own doc (data.go's classItem) for why this
	// exists alongside SubclassID rather than reusing it.
	WeaponType         string
	FactionRestriction string
	Stats              map[string]float64
	DamageMin          float64
	DamageMax          float64
	Speed              float64
	DPS                float64
	TwoHand            bool
	Unique             bool
	Slots              []string
	SetID              *int
	// EffectText is the item's own on-hit/on-use/proc description
	// (data/builds/<build>/items/<class>.json's own effect_text), empty
	// for an item with no such effect. Carried through from
	// classItem.EffectText, unused before this lane (2026-09-28
	// weights-effects): score() cannot see a proc at all (it only dots
	// Stats against a spec's weights), which is exactly the gap this
	// lane's brief calls out ("Do we factor in on-hit effects, set
	// bonuses, etc. in our sims?") -- effectcoverage.go and rank.go read
	// this field to decide whether a candidate needs an engine-verified
	// run instead of (or in addition to) score()'s stat total.
	EffectText string
	// EffectiveRequiredLevel is the level gate this candidate really
	// has: RequiredLevel unless a quest or crafted source floors it
	// higher (leveling.EffectiveRequiredLevel; see
	// applyEffectiveRequiredLevels, which sets this once per run,
	// right after loadCandidates+loadLootIndex, from loot.json's
	// quests map). eligible() gates on THIS field, not RequiredLevel
	// directly -- see this lane's brief and eligible.go's own doc.
	EffectiveRequiredLevel int
	// NotInSimDB is true when this build's embedded item database
	// (simdb.Known) does not carry a row for ID - set once per spec
	// run by markNotInSimDB, immediately after loadCandidates
	// (main.go's runSpec, the same point convertCandidateRatings/
	// applyEffectiveRequiredLevels already run at), this lane's brief,
	// item 1: simdb.Attach/AttachWeights' own UnequipUnknown silently
	// strips such an item from the character before ANY sim this
	// command ever builds, so a real sim never actually measured it,
	// however this candidate scored or which pass promoted it -
	// report.go's buildReport reads this field off the final pick to
	// keep the report honest (SimStatus, never Verified/SimDPS/
	// DPSDelta - see slotRow's own doc). A candidate a test builds
	// directly, bypassing loadCandidates/markNotInSimDB entirely
	// (every existing *_test.go fixture, all synthetic ids the real
	// simdb.bin has never heard of), defaults to false - "assume
	// known", this field's own zero value and the field's absence's
	// old behaviour alike - so the existing suite's fixtures are
	// unaffected; only a test that sets this field explicitly (this
	// lane's brief, item 2's contract test) exercises the new path.
	NotInSimDB bool
}

// markNotInSimDB sets NotInSimDB (above) on every item this build's
// embedded item database does not carry - this lane's brief, item 1.
// Called once per spec run, right after loadCandidates, so every
// candidate that ever reaches band.go's eligible()/buildBandPool
// already carries the flag before pick()/rankTrinketSlot/
// rankSlotWithEffects/trySetCompletion/verify.go ever see it.
func markNotInSimDB(items []candidate) []candidate {
	out := make([]candidate, len(items))
	for i, c := range items {
		c.NotInSimDB = !simdb.Known(int32(c.ID))
		out[i] = c
	}
	return out
}

// loadCandidates merges items.json and items/<class>.json for one
// class into the candidate pool. An id the class file carries but the
// flat file does not (should not happen; every build ships both from
// the same pipeline run) is skipped with a note, because eligibility
// cannot be decided without required_level/faction/armor-type. An id
// either row marks superseded_by (data-followups-4 lane, this lane's
// brief item 2) is skipped with a note too: it is never a candidate,
// whatever its own score would have been - see flatItem.SupersededBy's
// own doc for why the row still exists in items.json without ever
// reaching this pool.
func loadCandidates(buildDir, classSlug string) ([]candidate, []string, error) {
	flatItems, err := loadFlatItems(buildDir)
	if err != nil {
		return nil, nil, err
	}
	byID := make(map[int]flatItem, len(flatItems))
	for _, it := range flatItems {
		byID[it.ID] = it
	}
	classFile, err := loadClassItems(buildDir, classSlug)
	if err != nil {
		return nil, nil, err
	}
	var out []candidate
	var missing []string
	for _, ci := range classFile.Items {
		fi, ok := byID[ci.ID]
		if !ok {
			missing = append(missing, fmt.Sprintf("%d %s: in items/%s.json but not items.json", ci.ID, ci.Name, classSlug))
			continue
		}
		// Read defensively off whichever row carries it (flatItem.
		// SupersededBy's own doc): fi.SupersededBy wins when both are
		// set and happen to disagree, since items.json is the flat,
		// class-independent source of truth for this fact.
		supersededBy := fi.SupersededBy
		if supersededBy == 0 {
			supersededBy = ci.SupersededBy
		}
		if supersededBy != 0 {
			missing = append(missing, fmt.Sprintf("%d %s: superseded_by %d, excluded from candidates", ci.ID, ci.Name, supersededBy))
			continue
		}
		out = append(out, candidate{
			ID:            ci.ID,
			Name:          ci.Name,
			Quality:       ci.Quality,
			RequiredLevel: ci.RequiredLevel,
			// Defaulted to the item's own required_level; applyEffectiveRequiredLevels
			// (called once the loot index is loaded) raises this for a
			// quest/crafted candidate whose true gate is higher.
			EffectiveRequiredLevel: ci.RequiredLevel,
			ItemLevel:              ci.ItemLevel,
			ClassID:                fi.ClassID,
			SubclassID:             fi.SubclassID,
			WeaponType:             ci.WeaponType,
			FactionRestriction:     factionOfRestriction(fi.FactionRestriction),
			Stats:                  ci.Stats,
			DamageMin:              ci.DamageMin,
			DamageMax:              ci.DamageMax,
			Speed:                  ci.Speed,
			DPS:                    ci.DPS,
			TwoHand:                ci.TwoHand,
			Unique:                 ci.Unique,
			Slots:                  plannerSlots(ci.Slot),
			SetID:                  ci.SetID,
			EffectText:             ci.EffectText,
		})
	}
	return out, missing, nil
}

func loadFlatItems(buildDir string) ([]flatItem, error) {
	b, err := os.ReadFile(filepath.Join(buildDir, "items.json"))
	if err != nil {
		return nil, fmt.Errorf("reading items.json: %w", err)
	}
	var items []flatItem
	if err := json.Unmarshal(b, &items); err != nil {
		return nil, fmt.Errorf("decoding items.json: %w", err)
	}
	return items, nil
}

func loadClassItems(buildDir, classSlug string) (classItemFile, error) {
	path := filepath.Join(buildDir, "items", classSlug+".json")
	b, err := os.ReadFile(path)
	if err != nil {
		return classItemFile{}, fmt.Errorf("reading %s: %w", path, err)
	}
	var f classItemFile
	if err := json.Unmarshal(b, &f); err != nil {
		return classItemFile{}, fmt.Errorf("decoding %s: %w", path, err)
	}
	return f, nil
}

// ratingStatColumns mirrors data/pipeline/simdb/ratings.py's own
// RATING_STAT_COLUMNS: the engine-bound stat key (candidate.Stats' own
// key, matching sim/leveling and the engine's unified Stat naming) ->
// the gametables/combatratings.txt column(s) that state, at level 60,
// how many rating points buy 1% of that stat. hit and crit each list
// all three schools because the engine unifies them into one Stat
// (data/pipeline/simdb/statmap.py); loadRatingFactors requires every
// column for a key to agree, so a build whose schools actually
// diverge fails loudly instead of silently picking one - exactly
// ratings.py's own rule, mirrored so the ranker's units never drift
// from what pipeline.simdb.ratings already divided out of simdb.bin.
var ratingStatColumns = map[string][]string{
	"hit":     {"Hit - Melee", "Hit - Ranged", "Hit - Spell"},
	"crit":    {"Crit - Melee", "Crit - Ranged", "Crit - Spell"},
	"dodge":   {"Dodge"},
	"parry":   {"Parry"},
	"block":   {"Block"},
	"defense": {"Defense Skill"},
}

// ratingFactors is level-60 rating points per 1%, one entry per
// ratingStatColumns key - loadRatingFactors' own return and
// convertRatingStats' own divisor.
type ratingFactors map[string]float64

// loadRatingFactors reads buildDir's own
// gametables/combatratings.txt (the same file
// data/pipeline/simdb/ratings.py reads to build simdb.bin) and
// returns its level-60 row's rating-to-percent factors for every
// rating-family stat score() dots against a weight the weights sweep
// measured per sim unit (percent), never per rating point. Loaded
// once per spec run (runSpec, alongside every other per-run loader in
// this file) rather than per band or per item: the table does not
// vary by band, and re-reading/re-parsing it per candidate would be
// pure overhead for a value that never changes within a run.
//
// Fails loudly - never silently skips a stat, defaults a missing
// column, or averages disagreeing ones - if the file, its level-60
// row, or any named column is missing or non-numeric, or if a key's
// own columns disagree: this build's own gametables/combatratings.txt
// is the single source of truth simdb.bin was already divided by, and
// a build where that has changed (or where a school's factor
// genuinely diverged) needs to be caught here, not ranked against a
// stale or averaged number.
func loadRatingFactors(buildDir string) (ratingFactors, error) {
	path := filepath.Join(buildDir, "gametables", "combatratings.txt")
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	var rows [][]string
	for _, line := range strings.Split(strings.ReplaceAll(string(b), "\r\n", "\n"), "\n") {
		if line == "" {
			continue
		}
		rows = append(rows, strings.Split(line, "\t"))
	}
	if len(rows) < 2 {
		return nil, fmt.Errorf("%s has a header and no rows", path)
	}
	header := rows[0]
	var level60 []string
	for _, row := range rows[1:] {
		if len(row) > 0 && row[0] == "60" {
			level60 = row
			break
		}
	}
	if level60 == nil {
		return nil, fmt.Errorf("%s has no level 60 row", path)
	}
	if len(level60) != len(header) {
		return nil, fmt.Errorf("%s level 60 row has %d cells against a %d-column header", path, len(level60), len(header))
	}
	byColumn := make(map[string]string, len(header))
	for i, name := range header {
		byColumn[name] = level60[i]
	}
	factors := make(ratingFactors, len(ratingStatColumns))
	for key, columns := range ratingStatColumns {
		var factor float64
		for i, col := range columns {
			raw, ok := byColumn[col]
			if !ok {
				return nil, fmt.Errorf("%s level 60 has no %q column (needed for %q)", path, col, key)
			}
			v, err := strconv.ParseFloat(raw, 64)
			if err != nil {
				return nil, fmt.Errorf("%s level 60 %q = %q: %w", path, col, raw, err)
			}
			if i == 0 {
				factor = v
			} else if v != factor {
				return nil, fmt.Errorf(
					"%s level 60 %v disagree (%v vs %v) for %q; the engine's unified stat needs one factor, not per-column ones",
					path, columns, factor, v, key,
				)
			}
		}
		if factor <= 0 {
			return nil, fmt.Errorf("%s level 60 %q factor is %v, not positive", path, key, factor)
		}
		factors[key] = factor
	}
	return factors, nil
}

// convertRatingStats returns a NEW map (this package's immutability
// rule; the same map a classItem/candidate carries may still be read
// elsewhere for what the client's own tooltip states) with every
// ratingFactors key's amount divided down from a combat-rating number
// to the flat percentage data/pipeline/simdb/ratings.py's own
// convert_rating_stats already produced for simdb.bin - the unit
// score() (score.go) and every weight it dots against actually share.
// A stat not in factors (agility, spell power, and so on - anything
// the client never itemises through ItemModType 31/32/12-15) passes
// through unchanged.
func convertRatingStats(stats map[string]float64, factors ratingFactors) map[string]float64 {
	out := make(map[string]float64, len(stats))
	for stat, amount := range stats {
		if factor, ok := factors[stat]; ok {
			out[stat] = amount / factor
			continue
		}
		out[stat] = amount
	}
	return out
}

// convertCandidateRatings returns a NEW slice, one new candidate per
// entry in items, with each candidate's Stats run through
// convertRatingStats - see that function's own doc. Called once per
// spec run (runSpec), immediately after loadCandidates, so every
// downstream consumer that dots a candidate's stats against a weight
// (score() first among them; pick.go's promoteLowValueWeapon only
// checks which keys are PRESENT, never their magnitude, so it is
// unaffected either way) operates in sim units throughout.
func convertCandidateRatings(items []candidate, factors ratingFactors) []candidate {
	out := make([]candidate, len(items))
	for i, c := range items {
		c.Stats = convertRatingStats(c.Stats, factors)
		out[i] = c
	}
	return out
}

// lootBoss is one entry of a raid/dungeon lootSource's own Bosses list.
// ItemChances (src-classicdb lane, 2026-09-29; wowhead-world-drops lane,
// 2026-09-29 for the wowhead-scraped case) is item id (string key,
// matching loot.json's own encoding) -> percent drop chance (0-100),
// present only for an item classic-db or wowhead itself states a chance
// for -- band.go's sourceFor reads it to pick the highest-chance boss
// among several naming the same item, rather than whichever the array
// happens to list first.
type lootBoss struct {
	Name        string             `json:"name"`
	Items       []int              `json:"items"`
	ItemChances map[string]float64 `json:"item_chances"`
}

// lootSource is one row of data/builds/<build>/loot.json's sources
// array. Bosses carries a raid/dungeon's per-boss item lists; Items
// carries every other kind's flat list (today: quest - one bucket for
// every quest reward, crafted, rep, pvp, world). ItemChances is Items'
// own per-item chance map, same shape and source as lootBoss.ItemChances
// above, for a flat (non-bossed) source such as a `world` mob or a
// `world_drop` pool. See loot.json's own "kind" values; this struct is
// deliberately permissive about which fields are present so a kind this
// file has never carried yet decodes instead of failing the whole load.
type lootSource struct {
	ID         string `json:"id"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Profession string `json:"profession"`
	FactionID  int    `json:"faction_id"`
	Standing   string `json:"standing"`
	Rank       int    `json:"rank"`
	// Faction is "alliance" or "horde" for a "pvp" kind source
	// (pipeline.loot.pvp_faction's own doc, pvp-faction lane 2026-09-29:
	// loot.json's own pvp:rank-N sources split into pvp:rank-N:alliance/
	// pvp:rank-N:horde, one per side, each carrying its own Faction), OR
	// for a "vendor" kind source pipeline.loot.pvp_faction resolved to a
	// rank quartermaster's own npc (fourth wow-player sweep, item 1:
	// vendorInheritsPvpRankGate below reads this to tell which of
	// possibly several same-id pvp sources a vendor row's own side
	// matches). Every other kind's own faction fact already reaches this
	// struct a different way (FactionID for rep, factionExclusiveDungeons
	// for a dungeon), so this field is empty for them.
	Faction string `json:"faction"`
	// Title is the pvp source's own in-game rank title
	// (pipeline.loot.pvp_faction.rank_title), for a "pvp" kind source
	// only - fourth wow-player sweep, item 2: "PvP rank 9 · Master
	// Sergeant · Alliance", never the bare bucket name.
	Title       string             `json:"title"`
	Items       []int              `json:"items"`
	ItemChances map[string]float64 `json:"item_chances"`
	Bosses      []lootBoss         `json:"bosses"`
	// Opens is the content phase this source unlocks in
	// (data/curated/loot/forever-raid-phases.json's own curated fact,
	// "later" for every Era raid this build's generator emits and
	// "raids-1" for Onyxia's Lair - none of them launch-day open: that
	// file's own notes say plainly "Nothing raids at launch on 4
	// November; the first tier opens on 9 December"). Empty for every
	// source this build has NOT curated a phase for - a dungeon, quest,
	// crafted, PvP or open-world source, and almost every reputation
	// source, all of which a previous lane's own item 3 confirmed IS
	// open at launch. sourceObtainable (band.go) reads this to gate a
	// leveling list on it directly, so a leveling character's OWN
	// band-60 list stops naming raid gear no launch-day 60 could
	// possibly have - see that lane's report for which band-60 picks
	// moved once this landed.
	//
	// bis-ranker-integrity-5 lane, item 4: that "every reputation
	// source is open at launch" confirmation has one real exception
	// loot.json itself never states - a rep-kind source's own faction
	// can be locked behind the same raid-era content patch as a raid
	// itself (Cenarion Circle, faction 609: introduced in Patch 1.9,
	// Gates of Ahn'Qiraj, the same patch that opens the Ahn'Qiraj raid
	// this build already gates "later" above and starts the AQ War
	// Effort the reputation is earned through - Wowhead's own Cenarion
	// Circle faction page dates its earliest quest to that event
	// chain). loadLootIndex applies repFactionRaidPhaseOpens (below) on
	// top of this field for exactly that case, since forever-raid-
	// phases.json's own curated facts patch raid/dungeon sources by id,
	// never a reputation faction.
	Opens string `json:"opens,omitempty"`
}

// lootQuestEntry is one entry of loot.json's `quests` map: one quest
// that hands out an item, with the level fields
// data/pipeline/wowhead_quests.py resolves (2026-09-28 quest-levels
// lane) -- min_level is the quest's own minimum level, which gates the
// reward far more often than the item's own required_level does (see
// applyEffectiveRequiredLevels).
type lootQuestEntry struct {
	QuestID     int    `json:"quest_id"`
	Name        string `json:"name"`
	Faction     string `json:"faction"`
	MinLevel    int    `json:"min_level"`
	Level       int    `json:"level"`
	LevelSource string `json:"level_source"`
	// Opens is QuestSource.opens (data/pipeline/loot/sources.py's own
	// apply_quest_opens_gate, quest-gates lane 2026-09-29): set when this
	// quest's own classic-db turn-in item(s) -- SrcItemId/ReqItemId1-4,
	// resolved through the quest's PrevQuestId chain -- are themselves
	// only obtainable from an opens-gated source (a raid boss drop, or
	// another such quest, recursively). Empty for the overwhelming
	// majority of quests, which need no raid-exclusive item at all.
	Opens string `json:"opens,omitempty"`
	// Classes is which classes' own quest this is - a data-followups
	// lane running in parallel (bis-ranker-integrity-12's brief, item
	// 3) is adding this field for a genuinely class-restricted quest
	// (quest 8253 "Destroy Morphaz" is a MAGE class quest, which is why
	// Fire Ruby, its reward, can never be a hunter's or shaman's
	// trinket - hunter-beast-mastery/shaman-elemental band 50's own
	// repro, this lane's item 1). Read defensively: nil/absent (every
	// quest in this build as of this lane) means "no class restriction
	// at all", so loadLootIndex/classAllowed (band.go) already handle
	// it correctly before the data lane ever lands the field, and start
	// gating the moment it does, with no second code change needed.
	Classes []string `json:"classes,omitempty"`
}

type lootFile struct {
	Sources []lootSource `json:"sources"`
	// Quests is item id (as a string key, matching loot.json's own
	// encoding) -> every quest that hands it out.
	Quests map[string][]lootQuestEntry `json:"quests"`
}

// itemSource is what loot.json says about one item: which kind of
// source named it, and a human label for the report/output badge.
type itemSource struct {
	Kind  string
	Label string
	// Side is "alliance" or "horde" for a source only one side can use:
	// a reputation only that side earns (the battleground factions) or a
	// dungeon inside the other side's capital (factionExclusiveDungeons);
	// "" for every other source.
	Side string
	// Standing is a rep source's required standing ("friendly" ..
	// "exalted"), "" for every other kind.
	Standing string
	// Chance is the item's own percent drop chance (0-100) at this
	// source, from lootSource.ItemChances/lootBoss.ItemChances --
	// wowhead-world-drops lane, 2026-09-29. 0 when this source states
	// none (the fork database and wowhead's dropped-by scrape usually
	// carry none), which sourceFor's bestBoss/bestOf treat as "lowest,
	// never preferred over a source that does state one" rather than as
	// a real 0% chance.
	Chance float64
	// Opens is lootSource.Opens, carried through unchanged - this lane's
	// brief, item 3: sourceObtainable (band.go) refuses any source this
	// is non-empty for, at every band, not only below 60.
	Opens string
	// Rank is lootSource.Rank, carried through unchanged, for a Kind ==
	// "pvp" source only (the Go zero value for every other kind - a rep/
	// vendor/quest/etc. source has no rank at all, and 0 is not itself a
	// real PvP rank, so band.go's pvpRankExceedsCap gates on Kind == "pvp"
	// first rather than trusting a bare Rank > pvpRankCap check alone).
	// bis-ranker-integrity-2 lane, item 3: loot.json's own pvp sources
	// (pvp:rank-5 .. pvp:rank-18) already carry this; nothing upstream
	// read it before this lane.
	Rank int
	// Title is lootSource.Title, carried through unchanged, for a Kind
	// == "pvp" source only - fourth wow-player sweep, item 2:
	// pvpSourceLabel (band.go) reads Rank/Title/Side together to format
	// "PvP rank 9 · Master Sergeant · Alliance" instead of the bare
	// bucket name loot.json's own Label would otherwise be.
	Title string
	// Classes mirrors lootQuestEntry.Classes (this lane's brief, item
	// 3) - nil for every source but a class-restricted quest reward,
	// which is every source in this build today (the field does not
	// exist in loot.json yet). classAllowed (band.go) is the one place
	// this is read.
	Classes []string
}

// repSide names the reputations only one side can earn, by the client's
// faction id as this build's loot.json carries it: Warsong Outriders 890 /
// Silverwing Sentinels 889, The Defilers 510 / The League of Arathor 509,
// Frostwolf Clan 729 / Stormpike Guard 730. Outrunner's Bow (a Warsong
// Outriders reward) was once the ALLIANCE level-20 hunter's bow.
// factionExclusiveDungeons is every dungeon source id (lootSource.ID)
// inside the opposite faction's capital, which the other faction cannot
// enter no matter what the item itself allows: Ragefire Chasm sits in
// Orgrimmar. Found by the warrior audit: Subterranean Cape (14149) has
// Taragaman the Hungerer as its only source and was the ALLIANCE level-20
// back pick. Vanilla has no Alliance-side mirror (Deadmines, Wailing
// Caverns and Shadowfang Keep are open-world instances either side walks
// to), so this table is one entry until another is found.
var factionExclusiveDungeons = map[string]string{
	"dungeon:ragefire-chasm": "horde",
}

var repSide = map[int]string{
	890: "horde", 510: "horde", 729: "horde",
	889: "alliance", 509: "alliance", 730: "alliance",
}

// questFactionSide maps loot.json's `quests[item][].faction` ("alliance",
// "horde" or "both") to the Side band.go's sourceObtainable expects ("" -
// no restriction - for "both", the faction name itself otherwise).
// quest-faction lane, 2026-09-29: loadLootIndex builds one itemSource per
// quest from this, rather than trusting the flat "quest" LootSource's own
// Side (always "", since its id is literally "quest") - see loadLootIndex's
// own doc for why a per-quest record, not a per-item merge, is required.
var questFactionSide = map[string]string{"alliance": "alliance", "horde": "horde", "both": ""}

// raidLockedQuestOpens is quest_id -> the content phase it needs, for a
// quest reward whose own loot.json entry reads "kind": "quest" (so the
// general Opens-from-a-raid-source gate in sourceObtainable never sees
// it - loadLootIndex skips "quest"-kind rows from that loop entirely
// and builds this one's itemSource separately, below) but whose quest
// chain is actually gated behind raid boss kills or a raid-tier
// community event - this lane's brief, item 3's own named example was
// "Atiesh, Greatstaff of the Guardian" (the Naxxramas class-quest
// chain's reward, not obtainable without having already killed
// Kel'Thuzad) and "Rise, Thunderfury!" (needs 8 Bindings of the
// Windseeker, which drop randomly off EVERY Molten Core boss); both are
// legendaries (items.json quality 5) and are gated by
// band.go's legendaryGatedLater now (bis-ranker-integrity-2 lane, item
// 2: "any item of quality 5 is gated later regardless of source kind"
// replaces the two hand-picked quest-id rows this map used to carry for
// them - one rule instead of one row per legendary this build ever
// adds).
//
// What is LEFT here is not a legendary at all: 8756/8789/8790 are the
// Ahn'Qiraj War Effort's three reward-item quests ("The Qiraji
// Conqueror", "Imperial Qiraji Armaments", "Imperial Qiraji Regalia" -
// found re-checking mage-fire band 60 after the Atiesh fix landed:
// "Blessed Qiraji Acolyte Staff", quest 8790, immediately became the
// new band-60 main_hand pick, itself gated behind the same War Effort a
// direct Ahn'Qiraj raid drop already is via forever-raid-phases.json's
// `raid:ahnqiraj` curated fact) - their own items.json quality is well
// below 5, so legendaryGatedLater's rule has no way to catch them; their
// gate is the war-effort RAID EVENT itself, which loot.json's own
// quest-kind rows never state.
//
// A broader systematic audit of every OTHER quest-kind reward that
// turns out to be raid-gated (the same category of gap
// forever-raid-phases.json's own curated facts fill for direct raid
// sources) is a follow-up beyond this lane's scope - see this lane's
// report for exactly what was and was not checked.
var raidLockedQuestOpens = map[int]string{
	8756: "later", // The Qiraji Conqueror - Ahn'Qiraj War Effort reward
	8789: "later", // Imperial Qiraji Armaments - Ahn'Qiraj War Effort reward
	8790: "later", // Imperial Qiraji Regalia - Ahn'Qiraj War Effort reward
}

// repFactionRaidPhaseOpens is a reputation faction id -> the content
// phase its OWN reputation track requires, for a faction loot.json's
// "rep"-kind sources never carry an Opens value for at all
// (lootSource.Opens' own doc above: forever-raid-phases.json's curated
// facts patch a raid/dungeon source by id, never a reputation
// faction) but whose reputation is not actually earnable until the
// same content patch that opens a raid this build already gates
// "later".
//
// bis-ranker-integrity-5 lane, 2026-09-30, this lane's brief item 4:
// Earthstrike (21180, rep:cenarion-circle:exalted) and its two
// standing-mates published as launch-day gear with no gate at all -
// Cenarion Circle (faction 609) is Gates of Ahn'Qiraj (Patch 1.9)
// content, the same patch this build's own raid:ahnqiraj source is
// already gated "later" for, so its reputation cannot be earned any
// earlier than the raid itself opens. A rep-kind source's Opens is
// firstNonEmpty(src.Opens, this map[factionID]) in loadLootIndex - the
// pipeline-curated fact wins when one exists; this hand-maintained
// fallback only fires where loot.json is silent.
var repFactionRaidPhaseOpens = map[int]string{
	609: "later", // Cenarion Circle - Gates of Ahn'Qiraj (AQ War Effort)
}

// worldBossSources is a loot.json "world"-kind source id -> the content
// phase its OWN world boss opens in, for the six world bosses whose
// loot.json entry (`kind: "world"`, the SAME generic bucket every
// ordinary named-mob world drop uses) carries no `opens` value at all -
// this lane's brief, item 3: "World bosses are raid content." A fresh
// level-60 character does not solo Lord Kazzak, Azuregos, or the four
// Dragons of Nightmare (Emeriss, Lethon, Taerar, Ysondre) - each needs a
// raid-sized group, exactly the same barrier forever-raid-phases.json's
// own curated facts already gate every direct raid-instance source
// behind (raid:onyxias-lair's own "raids-1", the earliest phase any
// raid content opens) - so each one is hand-listed here at that same
// phase, the way data.go's own raidLockedQuestOpens (above) hand-lists
// a quest reward loot.json's own per-item shape cannot state a raid
// gate for either. loadLootIndex's own `add` closure reads this with
// firstNonEmpty(src.Opens, worldBossSources[src.ID]) - the pipeline's
// own computed Opens wins the moment a data lane teaches
// pipeline.loot.classicdb (or a future world-boss-specific curated
// fact) to state one, exactly like every other firstNonEmpty gate in
// this file. Found dogfooding this lane's brief: warlock-affliction/
// -demonology/-destruction band 60 main_hand (both factions) published
// Amberseal Keeper, sourced to world:lord-kazzak with no gate at all,
// 2.3 DPS ahead of Ironbark Staff (a real, obtainable alternative).
var worldBossSources = map[string]string{
	"world:lord-kazzak": "raids-1",
	"world:azuregos":    "raids-1",
	"world:emeriss":     "raids-1",
	"world:lethon":      "raids-1",
	"world:taerar":      "raids-1",
	"world:ysondre":     "raids-1",
}

// firstNonEmpty returns the first non-empty string, or "" when every one
// is - used wherever a computed fact (the pipeline's own
// QuestSource.opens) should win over a hand-maintained fallback
// (raidLockedQuestOpens) without a chain of if-empty checks at the call
// site.
func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

// repFactionSwap pairs the six battleground faction ids with their
// opposite number, for correctedRepSource's fix below.
var repFactionSwap = map[int]int{
	890: 889, 889: 890,
	510: 509, 509: 510,
	729: 730, 730: 729,
}

// correctedRepSource swaps a rep source to its actual battleground
// faction whenever the item's own client-stated faction_restriction
// (itemFactionRestriction, from items.json, "" for an unrestricted
// item) disagrees with repSide[factionID] - a general rule, not a
// per-item allowlist, because this data error is not confined to one
// or two ids: Scout's Medallion (20442, horde_only, sold by Kelm
// Hargunth in the Barrens per wowhead) is mined under Silverwing
// Sentinels/alliance; Sentinel's Medallion (20444, alliance_only, sold
// by Illiyana Moonblaze in Ashenvale) is mined under Warsong
// Outriders/horde; the honored-tier Sentinel's Medallion (19541,
// alliance_only) is ALSO mined only under Warsong Outriders/horde -
// three items, two different WSG "honored" batches, each the exact
// inverse of the item's own restriction. A restricted item can only
// ever be equipped by the one faction it names, so that hard client
// fact outranks a mined rep source's Side whenever they disagree,
// rather than leaving the item unobtainable by anyone (band.go's
// sourceObtainable already trusts a matching itemFactionRestriction
// over a mismatched Side, which fixes obtainability alone; this
// function additionally corrects the DISPLAYED source name and
// standing side for the same disagreement, since a right item under a
// wrong-sounding source name still fails tenets.md #2's "the source
// line is right"). factionNames is the FactionID -> the fork's own
// name for the CORRECTED id (loadLootIndex's own pre-pass), so the
// label never hand-spells a name that could drift from the fork's own
// spelling. An unrestricted item, or one whose restriction already
// agrees with the mined Side, is returned unchanged.
func correctedRepSource(itemFactionRestriction string, factionID int, factionNames map[int]string) (int, string, bool) {
	if itemFactionRestriction == "" || repSide[factionID] == itemFactionRestriction {
		return factionID, "", false
	}
	swapped, ok := repFactionSwap[factionID]
	if !ok || repSide[swapped] != itemFactionRestriction {
		return factionID, "", false
	}
	return swapped, factionNames[swapped], true
}

// lootIndex is item id -> every source that names it. An item can
// have more than one (a quest reward that is also a vendor item), so
// eligible() and the report pick the most specific: a boss kill over
// "Quests" in general.
type lootIndex map[int][]itemSource

// loadLootIndex builds the item -> source index from loot.json, plus
// the item -> quest-level-floor map (2026-09-28 quest-levels lane;
// applyEffectiveRequiredLevels uses it).
//
// TODO(bis-data): loot.json's "quest" kind is one flattened bucket
// with no per-item faction or quest name (see this lane's brief and
// the design doc's "The site's loot.json flattens quests into one
// bucket today; it grows"). loot.json's vendor and zone-drop kinds have
// since landed and this generic add() loop already carries them into
// idx with Kind "vendor"/"zone" - the 2026-09-28 night-bis-sources lane
// added "vendor" to band.go's sourceKindPriority so a vendor source is
// actually picked, but "zone" (a farm-anywhere-in-this-zone drop, more
// speculative than a named world boss) is still absent from that
// priority list on purpose; report.go's "lucky" (unsourced-but-a-known-
// zone-drop) distinction still does not exist, so a zone-only item is
// reported as having no known source, full stop - see report.go's
// noSource accounting.
func loadLootIndex(buildDir string, itemFactionRestriction map[int]string) (lootIndex, map[int]int, error) {
	b, err := os.ReadFile(filepath.Join(buildDir, "loot.json"))
	if err != nil {
		return nil, nil, fmt.Errorf("reading loot.json: %w", err)
	}
	var f lootFile
	if err := json.Unmarshal(b, &f); err != nil {
		return nil, nil, fmt.Errorf("decoding loot.json: %w", err)
	}
	// factionNames is faction id -> the fork's own name for it (e.g. 889
	// -> "Silverwing Sentinels"), read from every rep source's own Name
	// before any item is added, so correctedRepSource can name a
	// misattributed item's ACTUAL faction without hand-spelling a label
	// that could drift from the fork's own spelling.
	factionNames := make(map[int]string)
	for _, src := range f.Sources {
		if src.Kind == "rep" {
			factionNames[src.FactionID] = src.Name
		}
	}
	idx := make(lootIndex)
	for _, src := range f.Sources {
		// quest-faction lane, 2026-09-29: the flat "quest" LootSource
		// names every quest-sourced item with no per-quest faction at all
		// (its own Side is always "" - factionExclusiveDungeons has no
		// entry for id "quest"), which is exactly what let an item reach
		// an alliance character through a horde-only quest (Hammerbone,
		// item 270018, quest 914 Leaders of the Fang) as long as the item
		// itself carried no client-stated factionRestriction. Skipped
		// here in favour of one itemSource PER QUEST below, built from
		// f.Quests' own per-quest Faction, so band.go's sourceObtainable
		// can actually tell a horde-only quest reward from an
		// alliance-only or unrestricted one.
		if src.Kind == "quest" {
			continue
		}
		add := func(id int, label string, chance float64) {
			is := itemSource{
				Kind: src.Kind, Label: label, Side: factionExclusiveDungeons[src.ID], Chance: chance, Opens: src.Opens,
			}
			if src.Kind == "pvp" {
				is.Rank = src.Rank
				is.Title = src.Title
				// pvp-faction lane, 2026-09-29: loot.json's own
				// pvp:rank-N:alliance/pvp:rank-N:horde sources each
				// carry their own Faction now (the module doc above) -
				// this is what makes sourceObtainable's Side check
				// actually gate a rank reward to its own side, the
				// third wow-player sweep's own defect (Alliance-titled
				// rewards reaching a Horde character's list and back).
				is.Side = src.Faction
				// Fourth wow-player sweep, item 2: never the bare bucket
				// name loot.json's own Name carries ("Rank 9
				// (Alliance)") - pvpSourceLabel (band.go) names the
				// rank, its in-game title and the faction together.
				is.Label = pvpSourceLabel(is.Rank, is.Title, is.Side)
			}
			if src.Kind == "world" {
				// This lane's brief, item 3: worldBossSources' own doc
				// above - the six named world bosses are raid content
				// even though loot.json's own "world" kind (shared with
				// every ordinary named-mob world drop) never carries an
				// opens value at all.
				is.Opens = firstNonEmpty(src.Opens, worldBossSources[src.ID])
			}
			if src.Kind == "vendor" && src.Faction != "" {
				// Fourth wow-player sweep, item 1: pipeline.loot.
				// pvp_faction already resolved this rank quartermaster's
				// OWN side (its own doc) and wrote it onto the vendor
				// row - carried here as this source's own Side so
				// vendorInheritsPvpRankGate below can tell which of
				// (possibly several) same-id pvp sources is really this
				// vendor's, rather than guessing at the first one found.
				is.Side = src.Faction
			}
			if src.Kind == "rep" {
				factionID := src.FactionID
				if correctedID, correctedLabel, swapped := correctedRepSource(itemFactionRestriction[id], factionID, factionNames); swapped {
					factionID, is.Label = correctedID, correctedLabel
				}
				is.Side = repSide[factionID]
				is.Standing = src.Standing
				// This lane's brief, item 4: repFactionRaidPhaseOpens'
				// own doc above - a reputation faction can be raid-era
				// content even though loot.json's own rep sources never
				// carry an opens value at all.
				is.Opens = firstNonEmpty(src.Opens, repFactionRaidPhaseOpens[factionID])
			}
			idx[id] = append(idx[id], is)
		}
		for _, id := range src.Items {
			add(id, src.Name, src.ItemChances[strconv.Itoa(id)])
		}
		for _, boss := range src.Bosses {
			label := src.Name
			if boss.Name != "" {
				label = src.Name + ": " + boss.Name
			}
			for _, id := range boss.Items {
				add(id, label, boss.ItemChances[strconv.Itoa(id)])
			}
		}
	}
	questFloors := make(map[int]int, len(f.Quests))
	for idStr, entries := range f.Quests {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			continue
		}
		levels := make([]int, len(entries))
		for i, e := range entries {
			levels[i] = leveling.QuestFloor(e.MinLevel, e.Level)
			// One itemSource per quest, not one merged record: an item
			// two faction-mirrored quest chains both award (one
			// alliance-only, one horde-only) stays obtainable by BOTH
			// sides, each through its own record, rather than the two
			// quests' factions collapsing into a single wrong Side.
			idx[id] = append(idx[id], itemSource{
				Kind:  "quest",
				Label: e.Name,
				Side:  questFactionSide[e.Faction],
				// e.Opens (the pipeline's own computed gate, quest-gates
				// lane 2026-09-29) wins when set; raidLockedQuestOpens
				// stays as the fallback for the Ahn'Qiraj war-effort
				// three (8756/8789/8790), whose own gate is the war
				// EVENT itself, not a classic-db turn-in item the
				// generic computation can see (that map's own doc).
				Opens: firstNonEmpty(e.Opens, raidLockedQuestOpens[e.QuestID]),
				// e.Classes (this lane's brief, item 3): nil for every
				// quest in this build today, carried through unchanged
				// so classAllowed (band.go) gates on it the moment the
				// parallel data lane lands it.
				Classes: e.Classes,
			})
		}
		questFloors[id] = leveling.LowestFloor(levels)
	}
	return vendorInheritsPvpRankGate(vendorInheritsRepStandingGate(idx)), questFloors, nil
}

// vendorInheritsPvpRankGate returns a copy of idx where a "vendor"
// source sharing an item id with a "pvp" source inherits that pvp
// source's own Rank AND Faction (Side) - the exact same shape
// vendorInheritsRepStandingGate (below) already fixes for reputation,
// found dogfooding this lane's brief item 3 (pvpRankCap): Captain
// O'Neal (Alliance's Grand Marshal rank-reward quartermaster,
// loot.json's own vendor:12782) lists the SAME items (Grand Marshal's
// Stave 18873, Grand Marshal's Sunderer 18830, Grand Marshal's
// Demolisher 23455, ...) loot.json's own pvp:rank-18 source already
// lists with Rank 18 - but "vendor" outranks "pvp" in
// sourceKindPriority (band.go), so sourceFor picked the vendor row for
// every one of them and band.go's pvpRankExceedsCap never saw a Rank at
// all (the vendor kind carries none on its own), silently defeating the
// whole cap: every caster/melee/hybrid spec's band-60 main_hand this
// lane regenerated to check item 1 picked a Grand Marshal/High Warlord
// weapon from ITS OWN vendor before this fix, with the cap doing
// nothing.
//
// Fourth wow-player sweep, item 1: the pvp-faction lane (2026-09-29)
// renamed loot.json's own bucket from a single "pvp:rank-N" to
// "pvp:rank-N:alliance"/"pvp:rank-N:horde", so the ORIGINAL version of
// this function (matching on Kind == "pvp" alone, taking the first
// match) stopped copying Rank at all - worse, it never copied Side
// either even before that rename, so a vendor row that DID inherit a
// Rank was still obtainable by both factions at once, with only the
// rank cap (not the faction gate) doing any work. Fixed by: matching
// every "pvp" source for this item id (ordinarily exactly one), and
// when there IS exactly one, trusting it outright regardless of what
// (if anything) the vendor row's own Side already said - the real-data
// case this bug actually hit. The rarer, pathological case both
// factions' own pvp sources somehow name the SAME classic item id (a
// re-itemised Forever-new item copied onto a classic twin that landed
// in both splits, or a vendor row this build's pvp_faction lane could
// not otherwise disambiguate) is resolved by the vendor row's OWN known
// side instead (data.go's add() closure already copied
// pipeline.loot.pvp_faction's own npc-faction resolution onto it, via
// lootSource.Faction) - a vendor row whose own side is unknown AND
// faces more than one matching pvp source is left alone rather than
// guessed at (tenet 8: unverifiable is left out, never shown as fact).
//
// An item's vendor source with no matching pvp source (an ordinary gold
// vendor, or a rep-gated quartermaster with no rank reward) is returned
// unchanged.
func vendorInheritsPvpRankGate(idx lootIndex) lootIndex {
	out := make(lootIndex, len(idx))
	for id, srcs := range idx {
		var pvpSrcs []itemSource
		for _, s := range srcs {
			if s.Kind == "pvp" {
				pvpSrcs = append(pvpSrcs, s)
			}
		}
		if len(pvpSrcs) == 0 {
			out[id] = srcs
			continue
		}
		gated := make([]itemSource, len(srcs))
		for i, s := range srcs {
			if s.Kind != "vendor" {
				gated[i] = s
				continue
			}
			match, ok := pvpSourceForVendor(s, pvpSrcs)
			if ok {
				s.Rank, s.Side, s.Title = match.Rank, match.Side, match.Title
				s.Label = pvpSourceLabel(s.Rank, s.Title, s.Side)
			}
			gated[i] = s
		}
		out[id] = gated
	}
	return out
}

// pvpSourceForVendor is vendorInheritsPvpRankGate's own matching rule
// (its doc above): the single pvp source sharing this item id when
// there is only one, or - when more than one exists - whichever pvp
// source's own Side matches the vendor row's already-known Side
// (data.go's add() closure, from pipeline.loot.pvp_faction's own npc-
// faction resolution). ok is false when neither rule settles it (more
// than one pvp source and no known/matching vendor side), in which case
// the caller leaves the vendor row untouched.
func pvpSourceForVendor(vendor itemSource, pvpSrcs []itemSource) (itemSource, bool) {
	if len(pvpSrcs) == 1 {
		return pvpSrcs[0], true
	}
	if vendor.Side == "" {
		return itemSource{}, false
	}
	for _, p := range pvpSrcs {
		if p.Side == vendor.Side {
			return p, true
		}
	}
	return itemSource{}, false
}

// vendorInheritsRepStandingGate returns a copy of idx where a "vendor"
// source sharing an item id with a "rep" source inherits that rep
// source's own Side/Standing (this lane's brief, defect 2). A
// reputation quartermaster's own vendor row (Illiyana Moonblaze,
// Kelm Hargunth, ...) duplicates its faction's rep-reward item list
// verbatim - loot.json's own data confirms this for every rep tier
// this build carries - but the vendor kind itself names no
// FactionID/Standing (lootSource's Standing field is populated only
// on the `src.Kind == "rep"` branch above), so sourceObtainable saw a
// bare, level-agnostic "vendor" purchase for an item that actually
// needs Silverwing Sentinels revered: Outrunner's Bow (item 19562+)
// showed as an ordinary level-18 vendor item, and the honored-tier
// Sentinel's Medallion's own report line read "vendor" instead of
// "Silverwing Sentinels (honored)" - the reputation gate the item
// really has, bypassed by its own vendor row. An item's vendor source
// with no matching rep source (an ordinary gold vendor) is returned
// unchanged.
func vendorInheritsRepStandingGate(idx lootIndex) lootIndex {
	out := make(lootIndex, len(idx))
	for id, srcs := range idx {
		var repSrc *itemSource
		for i := range srcs {
			if srcs[i].Kind == "rep" {
				repSrc = &srcs[i]
				break
			}
		}
		if repSrc == nil {
			out[id] = srcs
			continue
		}
		gated := make([]itemSource, len(srcs))
		for i, s := range srcs {
			if s.Kind == "vendor" {
				s.Side = repSrc.Side
				s.Standing = repSrc.Standing
			}
			gated[i] = s
		}
		out[id] = gated
	}
	return out
}

// applyEffectiveRequiredLevels returns a copy of items with each
// candidate's EffectiveRequiredLevel set from leveling.
// EffectiveRequiredLevel (2026-09-28 quest-levels lane): a quest-reward
// candidate's floor is questFloors[c.ID] (loot.json's quests map, the
// LOWEST min_level among the quests that award it -- any one suffices);
// a crafted candidate (idx[c.ID] names a "crafted" source) with no
// quest floor uses leveling.ItemLevelProxyRequiredLevel(c.ItemLevel),
// since loot.json's crafted source carries no recipe skill level today.
// Called once per run, right after loadCandidates+loadLootIndex, before
// any band is built - see main.go's runSpec.
func applyEffectiveRequiredLevels(items []candidate, idx lootIndex, questFloors map[int]int) []candidate {
	out := make([]candidate, len(items))
	for i, c := range items {
		floor := questFloors[c.ID]
		if floor == 0 {
			for _, src := range idx[c.ID] {
				if src.Kind == "crafted" {
					floor = leveling.ItemLevelProxyRequiredLevel(c.ItemLevel)
					break
				}
			}
		}
		if floor == 0 && c.RequiredLevel == 0 {
			// A row with no required level at all (772 Forever-new items on
			// 1.60.1.70009 ship required_level 0 -- a data gap, not "usable at
			// 1": Swamp Ring 270052, item level 35, headed a level-20 list) is
			// gated by its item level the same way an unfloored recipe is.
			floor = leveling.ItemLevelProxyRequiredLevel(c.ItemLevel)
		}
		c.EffectiveRequiredLevel = leveling.EffectiveRequiredLevel(c.RequiredLevel, floor)
		out[i] = c
	}
	return out
}

// specInfo is the fields of one data/curated/specs.json row this
// command needs.
type specInfo struct {
	Spec          string   `json:"spec"`
	ClassSlug     string   `json:"class_slug"`
	SpecSlug      string   `json:"spec_slug"`
	Name          string   `json:"name"`
	TreeIndex     int      `json:"tree_index"`
	ReferenceStat string   `json:"reference_stat"`
	WeightStats   []string `json:"weight_stats"`
}

func loadAllSpecs(repoRoot string) ([]specInfo, error) {
	b, err := os.ReadFile(filepath.Join(repoRoot, "data", "curated", "specs.json"))
	if err != nil {
		return nil, fmt.Errorf("reading curated/specs.json: %w", err)
	}
	var specs []specInfo
	if err := json.Unmarshal(b, &specs); err != nil {
		return nil, fmt.Errorf("decoding curated/specs.json: %w", err)
	}
	return specs, nil
}

func loadSpec(repoRoot, spec string) (specInfo, error) {
	specs, err := loadAllSpecs(repoRoot)
	if err != nil {
		return specInfo{}, err
	}
	for _, s := range specs {
		if s.Spec == spec {
			return s, nil
		}
	}
	return specInfo{}, fmt.Errorf("no spec %q in curated/specs.json", spec)
}

// aplState is the one field this command needs out of
// data/curated/apl/<spec>.json - sim/request/ladder_test.go's own
// ladderCurated reads the same field the same way (loadLadderCurated),
// duplicated here rather than imported because that type lives in
// package request's _test.go file, not exported for another command to
// use.
type aplState struct {
	State string `json:"state"`
}

func loadAPLState(repoRoot, spec string) (string, error) {
	path := filepath.Join(repoRoot, "data", "curated", "apl", spec+".json")
	b, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("reading %s: %w", path, err)
	}
	var f aplState
	if err := json.Unmarshal(b, &f); err != nil {
		return "", fmt.Errorf("decoding %s: %w", path, err)
	}
	return f.State, nil
}

// writtenSpecs is -all's spec list: every data/curated/specs.json row
// whose own data/curated/apl/<spec>.json rotation is state == "written"
// - the same test sim/request/ladder_test.go's TestRotationLadder
// applies before it measures a spec, and the same "written" gate
// docs/superpowers/specs/2026-09-28-leveling-bis-design.md's own
// pipeline section names. A spec with no apl file at all (not yet
// curated) is skipped, not an error - this command ranks what is
// written today, not what will exist eventually.
func writtenSpecs(repoRoot string) ([]string, error) {
	specs, err := loadAllSpecs(repoRoot)
	if err != nil {
		return nil, err
	}
	var out []string
	for _, s := range specs {
		state, err := loadAPLState(repoRoot, s.Spec)
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				continue
			}
			return nil, err
		}
		if state == "written" {
			out = append(out, s.Spec)
		}
	}
	return out, nil
}

// guideRaces is what this lane reads out of a spec guide's frontmatter:
// the two recommended races, one per faction (guides/*.md's own
// convention - see marksmanship.md's `recommendedRaces: [dwarf,
// troll]`, Alliance then Horde). The guide's level-60 talent BUILD
// (the `build:` line this struct used to also carry as Trees) is no
// longer read here - sim/leveling.GuideBuildTalents reads it directly
// by stable talent id instead (see runSpec's talent truncation, and
// sim/leveling's own package doc for why: a positional read here would
// silently misalign against an active build whose tree lost a talent
// since the guide was authored).
type guideRaces struct {
	// AllianceRace, HordeRace are recommendedRaces[0], [1]: the guide
	// states Alliance first, Horde second, and loadGuideRaces itself now
	// checks each against data/builds/<build>/races.json's own
	// slug->faction fact before returning (this lane's brief, item 4) -
	// a caller holding a guideRaces value already knows both races are
	// real and on the faction the field name says.
	AllianceRace string
	HordeRace    string
}

var racesLineRE = regexp.MustCompile(`^recommendedRaces:\s*\[([^\]]*)\]\s*$`)

// raceFaction is race slug -> "alliance"/"horde", from
// data/builds/<build>/races.json - the client's own, single source of
// truth for which faction can play which race (this lane's brief,
// item 4: "fix it for every class/faction against races.json").
type raceFaction map[string]string

type raceFile struct {
	Slug    string `json:"slug"`
	Faction string `json:"faction"`
}

func loadRaceFactions(buildDir string) (raceFaction, error) {
	path := filepath.Join(buildDir, "races.json")
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	var rows []raceFile
	if err := json.Unmarshal(b, &rows); err != nil {
		return nil, fmt.Errorf("decoding %s: %w", path, err)
	}
	out := make(raceFaction, len(rows))
	for _, r := range rows {
		out[r.Slug] = r.Faction
	}
	return out, nil
}

// loadGuideRaces reads a spec guide's frontmatter for its recommended
// races. It is a small, deliberately line-oriented reader rather than a
// YAML parser: the frontmatter is one line this command needs out of a
// much larger file, and a full YAML dependency for one regex is not
// worth adding to the sim module.
//
// The guide's own convention (guideRaces' own doc) is exactly 2 races,
// Alliance then Horde - racesLineRE's own match is split on every comma
// with no bound on how many. Before this lane, a third entry (a copy-
// paste slip, not a typo: paladin/retribution.md's own frontmatter read
// "recommendedRaces: [human, dwarf, undead]" - three viable-looking
// races, but position [1] silently became HordeRace) was accepted
// without complaint, positions [0]/[1] taken and everything after
// dropped on the floor - reading Alliance-only Dwarf into HordeRace for
// every band of paladin-retribution's Horde list (Forever's actual new
// Horde Paladin race is Undead, research/01-official-facts.md; the
// guide's own prose already said so, just not its frontmatter). Fixed
// two ways: parts is now rejected outright unless it has exactly 2
// entries (fail fast on the shape this bug actually took), and each
// race is checked against raceFactions - the client's own
// faction_restriction fact, not another hand-maintained guess - so a
// future guide naming a Horde race for Alliance (or vice versa) fails
// the same load instead of silently mis-racing every band it publishes.
func loadGuideRaces(repoRoot, buildDir, classSlug, specSlug string) (guideRaces, error) {
	path := filepath.Join(repoRoot, "web", "src", "content", "guides", classSlug, specSlug+".md")
	b, err := os.ReadFile(path)
	if err != nil {
		return guideRaces{}, fmt.Errorf("reading %s: %w", path, err)
	}
	factions, err := loadRaceFactions(buildDir)
	if err != nil {
		return guideRaces{}, err
	}
	var out guideRaces
	var found bool
	for _, line := range strings.Split(string(b), "\n") {
		m := racesLineRE.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		found = true
		parts := strings.Split(m[1], ",")
		if len(parts) != 2 {
			return guideRaces{}, fmt.Errorf("%s: recommendedRaces: want exactly 2 races (Alliance, Horde), got %d: %q", path, len(parts), m[1])
		}
		out.AllianceRace = strings.Trim(strings.TrimSpace(parts[0]), "'\"")
		out.HordeRace = strings.Trim(strings.TrimSpace(parts[1]), "'\"")
	}
	if !found || out.AllianceRace == "" || out.HordeRace == "" {
		return guideRaces{}, fmt.Errorf("%s: recommendedRaces: frontmatter did not parse into 2 races, got %q/%q", path, out.AllianceRace, out.HordeRace)
	}
	if f := factions[out.AllianceRace]; f != "alliance" {
		return guideRaces{}, fmt.Errorf("%s: recommendedRaces: %q is not an Alliance race per races.json (faction %q)", path, out.AllianceRace, f)
	}
	if f := factions[out.HordeRace]; f != "horde" {
		return guideRaces{}, fmt.Errorf("%s: recommendedRaces: %q is not a Horde race per races.json (faction %q)", path, out.HordeRace, f)
	}
	return out, nil
}

// factionOfRestriction maps items.json's faction_restriction value
// ("alliance_only", "horde_only", "") onto the faction words the bands
// are ranked for ("alliance", "horde"), so eligible() compares like with
// like. Before this, every restricted item -- 870 in this build, Tunic of
// Westfall among them -- failed the comparison for BOTH factions and was
// silently left out of every list.
func factionOfRestriction(restriction string) string {
	return strings.TrimSuffix(restriction, "_only")
}
