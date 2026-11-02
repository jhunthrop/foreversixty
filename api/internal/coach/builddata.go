package coach

// Task 29b's file tables load Tables.Spells and Tables.Trainable through ParseSpellConst
// (data/builds/<build>/spellconst/<class>.json) and ParseTrainables
// (data/builds/<build>/trainables/<class>.json, the SkillLineAbility "what a Forever player
// learns" source), so production reads the files the fixtures are proved against
// (TestFixtureSpellConstMatchesData, TestFixtureTrainablesMatchData). spellranks.json is
// read for one fact the trainables file lacks: which ranks are Ahn'Qiraj books
// (ParseBookRanks), which no site rotation casts.

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
)

// spellConstFile is the shape of data/builds/<build>/spellconst/<class>.json, the fields
// the coach reads.
type spellConstFile struct {
	Spells map[string]struct {
		Name               string  `json:"name"`
		Rank               int     `json:"rank"`
		GCDMS              int64   `json:"gcd_ms"`
		CooldownMS         int64   `json:"cooldown_ms"`
		CategoryCooldownMS int64   `json:"category_cooldown_ms"`
		Cost               *int64  `json:"cost"`
		CostType           int64   `json:"cost_type"`
		CostPct            float64 `json:"cost_pct"`
	} `json:"spells"`
}

// ParseSpellConst reads one class's spellconst file into Tables.Spells. The cooldown is
// the larger of cooldown_ms and category_cooldown_ms; the cost is known when the file
// states it as a flat amount (a cost given only as a share of the bar is not).
func ParseSpellConst(data []byte) (map[int64]SpellConst, error) {
	var f spellConstFile
	if err := json.Unmarshal(data, &f); err != nil {
		return nil, fmt.Errorf("coach spellconst: %w", err)
	}
	out := make(map[int64]SpellConst, len(f.Spells))
	for key, s := range f.Spells {
		id, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("coach spellconst: spell key %q: %w", key, err)
		}
		c := SpellConst{Name: s.Name, Rank: s.Rank, CooldownMS: max(s.CooldownMS, s.CategoryCooldownMS), GCDMS: s.GCDMS, CostType: s.CostType}
		if s.Cost != nil && (*s.Cost > 0 || s.CostPct == 0) {
			c.Cost, c.CostKnown = *s.Cost, true
		}
		out[id] = c
	}
	return out, nil
}

// trainablesFile is the shape of data/builds/<build>/trainables/<class>.json.
type trainablesFile struct {
	Trainables []struct {
		Name  string `json:"name"`
		Ranks []struct {
			ID    int64 `json:"id"`
			Rank  int   `json:"rank"`
			Level int   `json:"level"`
		} `json:"ranks"`
	} `json:"trainables"`
}

// ParseTrainables reads one class's trainables file into Tables.Trainable: family ->
// ranks with the level each is learned at, in the file's order.
func ParseTrainables(data []byte) (map[string][]TrainableRank, error) {
	var f trainablesFile
	if err := json.Unmarshal(data, &f); err != nil {
		return nil, fmt.Errorf("coach trainables: %w", err)
	}
	out := make(map[string][]TrainableRank, len(f.Trainables))
	for _, t := range f.Trainables {
		for _, r := range t.Ranks {
			out[t.Name] = append(out[t.Name], TrainableRank{ID: r.ID, Rank: r.Rank, Level: r.Level})
		}
	}
	return out, nil
}

// spellRanksFile is the shape of data/builds/<build>/spellranks.json, the fields the coach
// reads: each rank's id and its book flag.
type spellRanksFile struct {
	Classes map[string]map[string][]struct {
		ID   int64 `json:"id"`
		Book bool  `json:"book"`
	} `json:"classes"`
}

// ParseBookRanks is the ids of class's Ahn'Qiraj book ranks in a spellranks.json file.
func ParseBookRanks(data []byte, class string) (map[int64]bool, error) {
	var f spellRanksFile
	if err := json.Unmarshal(data, &f); err != nil {
		return nil, fmt.Errorf("coach spellranks: %w", err)
	}
	books := map[int64]bool{}
	for _, ranks := range f.Classes[class] {
		for _, r := range ranks {
			if r.Book {
				books[r.ID] = true
			}
		}
	}
	return books, nil
}

// withoutBookRanks is trainable with every book rank left out: the trainer's top rank is
// the one the site's rotations cast (sim/internal/spellranks never casts a book rank).
func withoutBookRanks(trainable map[string][]TrainableRank, books map[int64]bool) map[string][]TrainableRank {
	out := make(map[string][]TrainableRank, len(trainable))
	for family, ranks := range trainable {
		out[family] = slices.DeleteFunc(slices.Clone(ranks), func(r TrainableRank) bool { return books[r.ID] })
	}
	return out
}
