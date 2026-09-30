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
	ID                 int
	Name               string
	Quality            int
	RequiredLevel      int
	ItemLevel          int
	ClassID            int
	SubclassID         int
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
}

// loadCandidates merges items.json and items/<class>.json for one
// class into the candidate pool. An id the class file carries but the
// flat file does not (should not happen; every build ships both from
// the same pipeline run) is skipped with a note, because eligibility
// cannot be decided without required_level/faction/armor-type.
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
	// Faction is "alliance" or "horde" for a "pvp" kind source only
	// (pipeline.loot.pvp_faction's own doc, pvp-faction lane 2026-09-29:
	// loot.json's own pvp:rank-N sources split into pvp:rank-N:alliance/
	// pvp:rank-N:horde, one per side, each carrying its own Faction) -
	// every other kind's own faction fact already reaches this struct a
	// different way (FactionID for rep, factionExclusiveDungeons for a
	// dungeon), so this field is empty for them.
	Faction     string             `json:"faction"`
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
	// reputation, crafted, PvP or open-world source, every one of which
	// this lane's brief (item 3) confirms IS open at launch. sourceObtainable
	// (band.go) reads this to gate a leveling list on it directly, so a
	// leveling character's OWN band-60 list stops naming raid gear no
	// launch-day 60 could possibly have - see this lane's report for
	// which band-60 picks moved once this landed.
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
				// pvp-faction lane, 2026-09-29: loot.json's own
				// pvp:rank-N:alliance/pvp:rank-N:horde sources each
				// carry their own Faction now (the module doc above) -
				// this is what makes sourceObtainable's Side check
				// actually gate a rank reward to its own side, the
				// third wow-player sweep's own defect (Alliance-titled
				// rewards reaching a Horde character's list and back).
				is.Side = src.Faction
			}
			if src.Kind == "rep" {
				factionID := src.FactionID
				if correctedID, correctedLabel, swapped := correctedRepSource(itemFactionRestriction[id], factionID, factionNames); swapped {
					factionID, is.Label = correctedID, correctedLabel
				}
				is.Side = repSide[factionID]
				is.Standing = src.Standing
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
			})
		}
		questFloors[id] = leveling.LowestFloor(levels)
	}
	return vendorInheritsPvpRankGate(vendorInheritsRepStandingGate(idx)), questFloors, nil
}

// vendorInheritsPvpRankGate returns a copy of idx where a "vendor"
// source sharing an item id with a "pvp" source inherits that pvp
// source's own Rank - the exact same shape vendorInheritsRepStandingGate
// (below) already fixes for reputation, found dogfooding this lane's
// brief item 3 (pvpRankCap): Captain O'Neal (Alliance's Grand Marshal
// rank-reward quartermaster, loot.json's own vendor:12782) lists the
// SAME items (Grand Marshal's Stave 18873, Grand Marshal's Sunderer
// 18830, Grand Marshal's Demolisher 23455, ...) loot.json's own
// pvp:rank-18 source already lists with Rank 18 - but "vendor" outranks
// "pvp" in sourceKindPriority (band.go), so sourceFor picked the vendor
// row for every one of them and band.go's pvpRankExceedsCap never saw a
// Rank at all (the vendor kind carries none on its own), silently
// defeating the whole cap: every caster/melee/hybrid spec's band-60
// main_hand this lane regenerated to check item 1 picked a Grand
// Marshal/High Warlord weapon from ITS OWN vendor before this fix, with
// the cap doing nothing. An item's vendor source with no matching pvp
// source (an ordinary gold vendor, or a rep-gated quartermaster with no
// rank reward) is returned unchanged.
func vendorInheritsPvpRankGate(idx lootIndex) lootIndex {
	out := make(lootIndex, len(idx))
	for id, srcs := range idx {
		var pvpSrc *itemSource
		for i := range srcs {
			if srcs[i].Kind == "pvp" {
				pvpSrc = &srcs[i]
				break
			}
		}
		if pvpSrc == nil {
			out[id] = srcs
			continue
		}
		gated := make([]itemSource, len(srcs))
		for i, s := range srcs {
			if s.Kind == "vendor" {
				s.Rank = pvpSrc.Rank
			}
			gated[i] = s
		}
		out[id] = gated
	}
	return out
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
