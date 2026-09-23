// api/internal/bnetbuild/talents.go
package bnetbuild

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

type blizzardSpecGroups struct {
	SpecializationGroups []blizzardSpecGroup `json:"specialization_groups"`
}

type blizzardSpecGroup struct {
	IsActive        bool                `json:"is_active"`
	Specializations []blizzardSpecEntry `json:"specializations"`
}

type blizzardSpecEntry struct {
	Talents []blizzardTalentEntry `json:"talents"`
}

type blizzardTalentEntry struct {
	Talent struct {
		ID int `json:"id"`
	} `json:"talent"`
	SpellTooltip struct {
		Spell struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"spell"`
	} `json:"spell_tooltip"`
	TalentRank int `json:"talent_rank"`
}

// talentEncodeResult is encodeTalents' own return shape: the per-tree ranks, what could not
// be placed or had to be clamped, and how many talents matched by each key — a ruling-round
// addition (spec docs/superpowers/specs/2026-09-22-battlenet-first-design.md §2.2's Report
// carries the counts too, not just the names) so the caller can see at a glance which
// matching strategy is actually doing the work on a given fixture, rather than only what
// failed.
type talentEncodeResult struct {
	TreeRanks      [3][]int
	Unmatched      []string
	Clamped        []string
	MatchedByName  int
	MatchedByID    int
	MatchedBySpell int
}

// encodeTalents reads the Blizzard specializations body and returns one rank slice per
// tree, in the class's tree position order (spec
// docs/superpowers/specs/2026-09-22-battlenet-first-design.md §2.2: "the active
// specialization group is used; when none is active, the first").
//
// A Blizzard talent is matched, in order: (1) by name — Blizzard's spell_tooltip.spell.name
// against our talent's own name, trimmed and case-folded, within the same class; (2) by
// talent.id against our talent id; (3) by the spell id against our spell_id/ranks[].spell_id.
// Name matching goes first because it is, empirically, the only key that actually agrees
// between Blizzard's classic1x profile API and this site's 1.60.1.69893 data on this
// fixture: Blizzard's Era talent.id is the old Talent.db2 id (Cruelty is 157 there, 105939
// in our data — a different DBC generation), and Blizzard reports the CURRENT rank's spell
// id (Cruelty rank 5 is spell 12856) while our data holds one spell id per talent regardless
// of rank (Cruelty is 12320 in every rank — see the package doc for why). Both id keys miss
// on nearly everything; the name still agrees because Blizzard and this site both render the
// talent's own display name. A talent this site renamed or removed for Forever has no name
// to match and falls through to the id keys (which also miss, on Era data), landing in
// Unmatched — the correct outcome, not a bug: Forever's own talent no longer has that name.
func encodeTalents(raw json.RawMessage, table TalentTable) (talentEncodeResult, error) {
	var body blizzardSpecGroups
	if err := json.Unmarshal(raw, &body); err != nil {
		return talentEncodeResult{}, fmt.Errorf("bnetbuild: decode specializations: %w", err)
	}
	group := firstActiveOrFirst(body.SpecializationGroups)

	classTrees := table.Build.Trees(table.ClassID)
	byName := talentsByName(classTrees)

	var result talentEncodeResult
	for i, tree := range classTrees {
		if i >= 3 {
			break
		}
		result.TreeRanks[i] = make([]int, len(tree.Talents))
	}

	for _, spec := range group.Specializations {
		for _, bt := range spec.Talents {
			ref, matchedKey, ok := matchTalent(bt, table, byName)
			if !ok {
				result.Unmatched = append(result.Unmatched, displayName(bt))
				continue
			}
			switch matchedKey {
			case matchByName:
				result.MatchedByName++
			case matchByID:
				result.MatchedByID++
			case matchBySpell:
				result.MatchedBySpell++
			}
			rank := bt.TalentRank
			if rank > ref.MaxRank {
				result.Clamped = append(result.Clamped, displayName(bt))
				rank = ref.MaxRank
			}
			treeIndex, talentIndex, found := positionOf(classTrees, ref.TreeID, ref.ID)
			if !found || treeIndex >= 3 {
				continue
			}
			result.TreeRanks[treeIndex][talentIndex] = rank
		}
	}
	return result, nil
}

// matchKey names which of the three strategies matched a talent, for talentEncodeResult's
// per-key counts.
type matchKey int

const (
	matchByName matchKey = iota
	matchByID
	matchBySpell
)

// matchTalent tries name, then talent.id, then spell id, in that order (see encodeTalents'
// doc comment for why this order and not the reverse).
func matchTalent(bt blizzardTalentEntry, table TalentTable, byName map[string]trees.TalentRef) (trees.TalentRef, matchKey, bool) {
	if ref, ok := byName[normalizeName(bt.SpellTooltip.Spell.Name)]; ok {
		return ref, matchByName, true
	}
	if ref, ok := table.Build.Talent(table.ClassID, bt.Talent.ID); ok {
		return ref, matchByID, true
	}
	if ref, ok := table.Build.TalentBySpellID(table.ClassID, bt.SpellTooltip.Spell.ID); ok {
		return ref, matchBySpell, true
	}
	return trees.TalentRef{}, 0, false
}

// normalizeName trims and lowercases a talent name for name-key comparison. Both sides of a
// name match go through this, so "Cruelty" and " cruelty " compare equal.
func normalizeName(name string) string {
	return strings.ToLower(strings.TrimSpace(name))
}

// talentsByName indexes every talent in classTrees by its normalized name. A name that
// collides within one class (none do, in this game's data, but nothing enforces it) keeps
// the first talent seen — the same "first wins" rule trees.Load's own duplicate-id refusal
// would apply if this were load-time validation instead of a lookup.
func talentsByName(classTrees []trees.Tree) map[string]trees.TalentRef {
	out := map[string]trees.TalentRef{}
	for _, tree := range classTrees {
		for _, t := range tree.Talents {
			key := normalizeName(t.Name)
			if _, exists := out[key]; !exists {
				out[key] = trees.TalentRef{Talent: t, TreeID: tree.ID, TreeName: tree.Name, TreePosition: tree.Position}
			}
		}
	}
	return out
}

// firstActiveOrFirst is spec §2.2's own words: "the active specialization group is used;
// when none is active, the first."
func firstActiveOrFirst(groups []blizzardSpecGroup) blizzardSpecGroup {
	for _, g := range groups {
		if g.IsActive {
			return g
		}
	}
	if len(groups) == 0 {
		return blizzardSpecGroup{}
	}
	return groups[0]
}

// displayName prefers the spell's own name (what a player recognises) over the bare talent
// id, falling back to the id when Blizzard sent no tooltip name.
func displayName(bt blizzardTalentEntry) string {
	if bt.SpellTooltip.Spell.Name != "" {
		return bt.SpellTooltip.Spell.Name
	}
	return fmt.Sprintf("talent %d", bt.Talent.ID)
}

// positionOf finds a talent's tree index (by tree position, 0-based) and its index within
// that tree's own Talents slice — the two coordinates encodeTree (encode.go) needs,
// matching web/src/lib/planner/fs1.ts's own "array index is tab order" contract.
func positionOf(classTrees []trees.Tree, treeID, talentID int) (treeIndex, talentIndex int, found bool) {
	for i, tree := range classTrees {
		if tree.ID != treeID {
			continue
		}
		for j, t := range tree.Talents {
			if t.ID == talentID {
				return i, j, true
			}
		}
	}
	return 0, 0, false
}
