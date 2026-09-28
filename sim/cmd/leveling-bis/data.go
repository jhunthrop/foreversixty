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
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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
			ID:                 ci.ID,
			Name:               ci.Name,
			Quality:            ci.Quality,
			RequiredLevel:      ci.RequiredLevel,
			ItemLevel:          ci.ItemLevel,
			ClassID:            fi.ClassID,
			SubclassID:         fi.SubclassID,
			FactionRestriction: fi.FactionRestriction,
			Stats:              ci.Stats,
			DamageMin:          ci.DamageMin,
			DamageMax:          ci.DamageMax,
			Speed:              ci.Speed,
			DPS:                ci.DPS,
			TwoHand:            ci.TwoHand,
			Unique:             ci.Unique,
			Slots:              plannerSlots(ci.Slot),
			SetID:              ci.SetID,
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

// lootSource is one row of data/builds/<build>/loot.json's sources
// array. Bosses carries a raid/dungeon's per-boss item lists; Items
// carries every other kind's flat list (today: quest - one bucket for
// every quest reward, crafted, rep, pvp, world). See loot.json's own
// "kind" values; this struct is deliberately permissive about which
// fields are present so a kind this file has never carried yet
// decodes instead of failing the whole load.
type lootSource struct {
	ID         string `json:"id"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Profession string `json:"profession"`
	Standing   string `json:"standing"`
	Rank       int    `json:"rank"`
	Items      []int  `json:"items"`
	Bosses     []struct {
		Name  string `json:"name"`
		Items []int  `json:"items"`
	} `json:"bosses"`
}

type lootFile struct {
	Sources []lootSource `json:"sources"`
}

// itemSource is what loot.json says about one item: which kind of
// source named it, and a human label for the report/output badge.
type itemSource struct {
	Kind  string
	Label string
}

// lootIndex is item id -> every source that names it. An item can
// have more than one (a quest reward that is also a vendor item), so
// eligible() and the report pick the most specific: a boss kill over
// "Quests" in general.
type lootIndex map[int][]itemSource

// loadLootIndex builds the item -> source index from loot.json.
//
// TODO(bis-data): loot.json's "quest" kind is one flattened bucket
// with no per-item faction or quest name (see this lane's brief and
// the design doc's "The site's loot.json flattens quests into one
// bucket today; it grows"). Lane bis-data is adding quest id/name/
// faction, vendor and zone-drop kinds; once those land, this index
// gains "vendor" and "zone" cases below and the report's "lucky"
// (unsourced-but-a-known-zone-drop) distinction becomes possible.
// Today an item absent from every source here is reported as having
// no known source, full stop - see report.go's noSource accounting.
func loadLootIndex(buildDir string) (lootIndex, error) {
	b, err := os.ReadFile(filepath.Join(buildDir, "loot.json"))
	if err != nil {
		return nil, fmt.Errorf("reading loot.json: %w", err)
	}
	var f lootFile
	if err := json.Unmarshal(b, &f); err != nil {
		return nil, fmt.Errorf("decoding loot.json: %w", err)
	}
	idx := make(lootIndex)
	add := func(id int, kind, label string) {
		idx[id] = append(idx[id], itemSource{Kind: kind, Label: label})
	}
	for _, src := range f.Sources {
		for _, id := range src.Items {
			add(id, src.Kind, src.Name)
		}
		for _, boss := range src.Bosses {
			label := src.Name
			if boss.Name != "" {
				label = src.Name + ": " + boss.Name
			}
			for _, id := range boss.Items {
				add(id, src.Kind, label)
			}
		}
	}
	return idx, nil
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

func loadSpec(repoRoot, spec string) (specInfo, error) {
	b, err := os.ReadFile(filepath.Join(repoRoot, "data", "curated", "specs.json"))
	if err != nil {
		return specInfo{}, fmt.Errorf("reading curated/specs.json: %w", err)
	}
	var specs []specInfo
	if err := json.Unmarshal(b, &specs); err != nil {
		return specInfo{}, fmt.Errorf("decoding curated/specs.json: %w", err)
	}
	for _, s := range specs {
		if s.Spec == spec {
			return s, nil
		}
	}
	return specInfo{}, fmt.Errorf("no spec %q in curated/specs.json", spec)
}

// guideBuild is what this lane reads out of a spec guide's frontmatter:
// the level-60 build code and the two recommended races, one per
// faction (guides/*.md's own convention - see marksmanship.md's
// `recommendedRaces: [dwarf, troll]`, Alliance then Horde).
type guideBuild struct {
	// Trees is the build code's three tree strings, in tree position
	// order (0, 1, 2), digit-per-talent-slot, exactly as the engine's
	// own positional talent string spells one tree (sim/talents'
	// package doc) - the guide's `build:` frontmatter uses "/" where
	// the engine string uses "-".
	Trees []string
	// AllianceRace, HordeRace are recommendedRaces[0], [1]: the guide
	// states Alliance first, Horde second, and bulk/expand.go's
	// hordeRaces table is what this command checks that against (see
	// main.go's loadGuide caller).
	AllianceRace string
	HordeRace    string
}

var buildLineRE = regexp.MustCompile(`^build:\s*'FS1:[^:]*:[^:]*:[^:]*:([^:']*):?'\s*$`)
var racesLineRE = regexp.MustCompile(`^recommendedRaces:\s*\[([^\]]*)\]\s*$`)

// loadGuideBuild reads a spec guide's frontmatter for its level-60
// talent build and recommended races. It is a small, deliberately
// line-oriented reader rather than a YAML parser: the frontmatter is
// two lines this command needs out of a much larger file, and a full
// YAML dependency for two regexes is not worth adding to the sim
// module.
func loadGuideBuild(repoRoot, classSlug, specSlug string) (guideBuild, error) {
	path := filepath.Join(repoRoot, "web", "src", "content", "guides", classSlug, specSlug+".md")
	b, err := os.ReadFile(path)
	if err != nil {
		return guideBuild{}, fmt.Errorf("reading %s: %w", path, err)
	}
	var out guideBuild
	for _, line := range strings.Split(string(b), "\n") {
		if m := buildLineRE.FindStringSubmatch(line); m != nil {
			trees := strings.Split(m[1], "/")
			out.Trees = trees
		}
		if m := racesLineRE.FindStringSubmatch(line); m != nil {
			parts := strings.Split(m[1], ",")
			for i, p := range parts {
				p = strings.TrimSpace(p)
				p = strings.Trim(p, "'\"")
				if i == 0 {
					out.AllianceRace = p
				} else if i == 1 {
					out.HordeRace = p
				}
			}
		}
	}
	if len(out.Trees) != 3 {
		return guideBuild{}, fmt.Errorf("%s: build: frontmatter did not parse into 3 trees, got %v", path, out.Trees)
	}
	if out.AllianceRace == "" || out.HordeRace == "" {
		return guideBuild{}, fmt.Errorf("%s: recommendedRaces: frontmatter did not parse into 2 races, got %q/%q", path, out.AllianceRace, out.HordeRace)
	}
	return out, nil
}
