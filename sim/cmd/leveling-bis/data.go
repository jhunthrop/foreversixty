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
	// states Alliance first, Horde second, and bulk/expand.go's
	// hordeRaces table is what this command checks that against (see
	// main.go's loadGuide caller).
	AllianceRace string
	HordeRace    string
}

var racesLineRE = regexp.MustCompile(`^recommendedRaces:\s*\[([^\]]*)\]\s*$`)

// loadGuideRaces reads a spec guide's frontmatter for its recommended
// races. It is a small, deliberately line-oriented reader rather than a
// YAML parser: the frontmatter is one line this command needs out of a
// much larger file, and a full YAML dependency for one regex is not
// worth adding to the sim module.
func loadGuideRaces(repoRoot, classSlug, specSlug string) (guideRaces, error) {
	path := filepath.Join(repoRoot, "web", "src", "content", "guides", classSlug, specSlug+".md")
	b, err := os.ReadFile(path)
	if err != nil {
		return guideRaces{}, fmt.Errorf("reading %s: %w", path, err)
	}
	var out guideRaces
	for _, line := range strings.Split(string(b), "\n") {
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
	if out.AllianceRace == "" || out.HordeRace == "" {
		return guideRaces{}, fmt.Errorf("%s: recommendedRaces: frontmatter did not parse into 2 races, got %q/%q", path, out.AllianceRace, out.HordeRace)
	}
	return out, nil
}
