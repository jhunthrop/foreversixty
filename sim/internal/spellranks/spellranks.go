// Package spellranks carries the active build's spell rank table: which
// id is which rank of which spell, and the level each rank is learned
// at.
//
// It is embedded exactly the way sim/internal/simdb embeds simdb.bin
// and enchants.json - `make simdb` copies
// data/builds/<build>/spellranks.json here, git-ignored, committed
// only at that source path - but it lives in its own package rather
// than inside sim/internal/simdb: simdb_test.go (package simdb, since
// it reaches simdb's unexported load()) already imports sim/request to
// prove request.Build attaches no database of its own, and
// sim/request's rotation rank rewrite is the one caller of this
// table's lookup. Putting the table in package simdb too would make
// simdb import sim/request AND sim/request import simdb - an import
// cycle Go refuses at test-compile time, not a real circular
// dependency, so the fix is to give the table its own leaf package
// rather than to unpick simdb_test.go's existing, unrelated coverage.
//
// sim/request's rotation rank rewrite needs to know, at request time,
// what a character's own level has learned, and the browser lane has
// no round trip to ask with.
//
// The data lane's generator (design doc "Data") emits, per class, only
// spell NAMES with more than one id or a learn level, each as its
// entries sorted by rank ascending. A name can still carry rank-0
// entries beside real player ranks - NPC or internal copies of the
// same spell, e.g. Sinister Strike's five rank-0 ids at level 20
// alongside its eight player ranks - so this parser only ever chains
// entries with rank >= 1 into a spell's rank progression; it never
// treats a rank-0 duplicate as one of that progression's ranks. A name
// whose entries are ALL rank 0 has no ranked-spell entry here at all
// (per the generator's own "more than one id or a learn level" filter,
// a single always-known ability has nothing to rewrite), so an id that
// never appears in this table is, by construction, not a ranked spell
// this package tracks - HighestLearnedSpellID returns it unchanged.
package spellranks

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"sort"
	"sync"
)

//go:embed spellranks.json
var rawSpellRanks []byte

// spellRankRow is one entry of spellranks.json's per-spell array.
type spellRankRow struct {
	ID    int32 `json:"id"`
	Rank  int   `json:"rank"`
	Level int   `json:"level"`
}

// spellRanksFile is the whole embedded table.
type spellRanksFile struct {
	Build   string                               `json:"build"`
	Classes map[string]map[string][]spellRankRow `json:"classes"`
}

// rankTier is every id that shares one rank number of one spell, and
// the level that rank is learned at.
type rankTier struct {
	level int
	ids   []int32
}

// rankChain is one spell's rank progression, tiers sorted ascending by
// rank number (and therefore, in practice, by level).
type rankChain struct {
	tiers []rankTier
}

// parseSpellRanks decodes the embedded table into, per class, a lookup
// from any id that belongs to a ranked spell to that spell's whole
// chain - so a lookup works regardless of which rank's id happens to
// be the one a rotation references (the APL always references the
// highest, since it is authored for MaxLevel, but nothing here assumes
// that).
func parseSpellRanks(b []byte) (map[string]map[int32]*rankChain, error) {
	var file spellRanksFile
	if err := json.Unmarshal(b, &file); err != nil {
		return nil, fmt.Errorf("spellranks: the table is not readable: %w", err)
	}
	if len(file.Classes) == 0 {
		return nil, errors.New("spellranks: the table carries no classes")
	}
	out := make(map[string]map[int32]*rankChain, len(file.Classes))
	for class, spells := range file.Classes {
		byID := make(map[int32]*rankChain)
		for _, rows := range spells {
			chain := buildRankChain(rows)
			if chain == nil {
				// Every entry was rank 0 (or the name carried none): not
				// a ranked spell this table tracks. See the package doc
				// comment's reference caveat.
				continue
			}
			for _, tier := range chain.tiers {
				for _, id := range tier.ids {
					byID[id] = chain
				}
			}
		}
		out[class] = byID
	}
	return out, nil
}

// buildRankChain groups rows sharing a rank number (rank >= 1 only)
// into tiers, sorted ascending by rank, in the order rows are first
// seen for that rank - which is also the file's own rank-ascending
// order, so the FIRST row seen at a rank is the one whose level this
// tier trusts. A later duplicate at the same rank number but a
// different (typically zero) level - Serpent Sting's rank-9 NPC copies
// at level 0, alongside its real rank-9 ids at level 60 - widens the
// tier's id set without ever lowering its level, because only the
// first-seen row's level is recorded.
func buildRankChain(rows []spellRankRow) *rankChain {
	levelByRank := make(map[int]int)
	idsByRank := make(map[int][]int32)
	var ranks []int
	for _, row := range rows {
		if row.Rank <= 0 {
			continue
		}
		if _, seen := levelByRank[row.Rank]; !seen {
			levelByRank[row.Rank] = row.Level
			ranks = append(ranks, row.Rank)
		}
		idsByRank[row.Rank] = append(idsByRank[row.Rank], row.ID)
	}
	if len(ranks) == 0 {
		return nil
	}
	sort.Ints(ranks)
	tiers := make([]rankTier, len(ranks))
	for i, r := range ranks {
		tiers[i] = rankTier{level: levelByRank[r], ids: idsByRank[r]}
	}
	return &rankChain{tiers: tiers}
}

var spellRankTable = sync.OnceValues(func() (map[string]map[int32]*rankChain, error) {
	return parseSpellRanks(rawSpellRanks)
})

// HighestLearnedSpellID resolves id to the id of the highest rank of
// its spell that a character of class has learned by level.
//
// ok is false when id belongs to a KNOWN ranked spell of class but no
// rank of it is learned yet at level - the caller (sim/request's
// rotation rank rewrite) drops whatever cast or condition named it.
// ok is true and newID equals id unchanged when id is not part of any
// ranked spell this table tracks for class (an unranked ability, or
// one the embed's class has no entry for at all): the rotation leaves
// it exactly as authored.
//
// A tier below the character's level whose ids do not include the
// original id (the character out-ranks the tier the rotation was
// written against) resolves to that tier's first id; ranks that share
// a tier - e.g. a talent-glyph variant sharing a level and rank number
// with the plain cast - are otherwise indistinguishable from this
// table alone, and picking the first is deterministic rather than
// arbitrary per call.
func HighestLearnedSpellID(class string, id int32, level int) (newID int32, ok bool) {
	table, err := spellRankTable()
	if err != nil {
		// A build problem (the embed failed to parse), not a per-request
		// one. Leaving every id untouched is the same failure mode as
		// never having this table at all, which is safer than treating
		// every ranked action in every rotation as unlearned.
		return id, true
	}
	byID, ok := table[class]
	if !ok {
		return id, true
	}
	return highestLearnedIn(byID, id, level)
}

// highestLearnedIn is HighestLearnedSpellID's resolution over one
// class's already-parsed table, factored out so a test can exercise it
// against parseSpellRanks(excerpt) without going through the embed.
func highestLearnedIn(byID map[int32]*rankChain, id int32, level int) (newID int32, ok bool) {
	chain, ok := byID[id]
	if !ok {
		return id, true
	}
	for i := len(chain.tiers) - 1; i >= 0; i-- {
		tier := chain.tiers[i]
		if tier.level > level {
			continue
		}
		if slices.Contains(tier.ids, id) {
			return id, true
		}
		return sameBandOrFirst(tier.ids, id), true
	}
	return 0, false
}

// foreverIDFloor separates the classic client's spell ids (all below it)
// from the ids Forever added (all above it). When a tier still holds
// more than one castable id -- Earth Shock keeps a classic id and two
// Forever copies at every rank -- the rewrite stays in the band the
// rotation named, since that is the family the engine registers.
const foreverIDFloor = 100000

func sameBandOrFirst(ids []int32, original int32) int32 {
	forever := original >= foreverIDFloor
	for _, candidate := range ids {
		if (candidate >= foreverIDFloor) == forever {
			return candidate
		}
	}
	return ids[0]
}
