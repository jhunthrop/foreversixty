// api/internal/bnetbuild/talents.go
package bnetbuild

import (
	"encoding/json"
	"fmt"

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

// encodeTalents reads the Blizzard specializations body and returns one rank slice per
// tree, in the class's tree position order (spec
// docs/superpowers/specs/2026-09-22-battlenet-first-design.md §2.2: "the active
// specialization group is used; when none is active, the first"). unmatched and clamped
// name every talent that could not be placed or whose rank exceeded max_rank, by spell
// name (Report.UnmatchedTalents, Report.Clamped).
func encodeTalents(raw json.RawMessage, table TalentTable) (treeRanks [3][]int, unmatched, clamped []string, err error) {
	var body blizzardSpecGroups
	if err := json.Unmarshal(raw, &body); err != nil {
		return treeRanks, nil, nil, fmt.Errorf("bnetbuild: decode specializations: %w", err)
	}
	group := firstActiveOrFirst(body.SpecializationGroups)

	classTrees := table.Build.Trees(table.ClassID)
	for i, tree := range classTrees {
		if i >= 3 {
			break
		}
		treeRanks[i] = make([]int, len(tree.Talents))
	}

	for _, spec := range group.Specializations {
		for _, bt := range spec.Talents {
			ref, ok := table.Build.Talent(table.ClassID, bt.Talent.ID)
			if !ok {
				ref, ok = table.Build.TalentBySpellID(table.ClassID, bt.SpellTooltip.Spell.ID)
			}
			if !ok {
				unmatched = append(unmatched, displayName(bt))
				continue
			}
			rank := bt.TalentRank
			if rank > ref.MaxRank {
				clamped = append(clamped, displayName(bt))
				rank = ref.MaxRank
			}
			treeIndex, talentIndex, found := positionOf(classTrees, ref.TreeID, ref.ID)
			if !found || treeIndex >= 3 {
				continue
			}
			treeRanks[treeIndex][talentIndex] = rank
		}
	}
	return treeRanks, unmatched, clamped, nil
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
