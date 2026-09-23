// api/internal/bnetbuild/encode.go

// Package bnetbuild turns a Blizzard character profile, equipment and specializations
// response into the FS1 v1 export string the simulator and planner already read (spec
// docs/superpowers/specs/2026-09-22-battlenet-first-design.md §2.2). Every function here is
// pure: no I/O, no database, no HTTP — bnetimport does all three and calls Encode with what
// it captured.
//
// Note on per-rank spell ids: today's data pipeline writes the same spell_id into every rank
// of a multi-rank talent (talents.go's blizzardTalentEntry-matching doc explains why this
// matters) rather than each rank's own distinct spell id, so TalentBySpellID's fallback
// rarely fires in practice — Blizzard's classic1x API reports the CURRENT rank's spell id,
// which this site's flattened data can only coincidentally hold. Fixing that is a data-lane
// change (a per-rank spell_id column), not something this package can correct on its own.
package bnetbuild

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
)

// Inputs is everything Encode needs, exactly the spec §2.2 signature.
type Inputs struct {
	Build     string
	Profile   bnetapi.CharacterProfile
	Equipment json.RawMessage
	Talents   json.RawMessage
	Talent    TalentTable
	Enchants  EnchantTable
	Suffixes  SuffixTable
	Races     RaceTable
}

// Report is what Encode could not map cleanly, plus how many talents each matching strategy
// resolved (MatchedByName/MatchedByID/MatchedBySpell) — logged at INFO with the character's
// key by the caller (bnetimport), never inside this pure package (spec §2.2).
type Report struct {
	UnmatchedTalents []string
	Clamped          []string
	NoSuffix         []string
	SkippedSlots     []string
	MatchedByName    int
	MatchedByID      int
	MatchedBySpell   int
}

// Encode builds one FS1 v1 string: "FS1:<build>:<class>:<race>:<t1>/<t2>/<t3>:<gear>" (spec
// §2.2, grammar in docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md §7,
// §10.5). An unmapped race is the one hard failure — every other mismatch (an unmatched
// talent, a clamped rank, an unresolved suffix, a skipped gear slot) is dropped and named in
// Report instead, per spec §1.3's honesty rule: nothing here is inferred or faked, but a
// partial build is still worth serving.
func Encode(in Inputs) (code string, report Report, err error) {
	raceSlug, ok := in.Races[in.Profile.RaceName]
	if !ok {
		return "", Report{}, fmt.Errorf("bnetbuild: unknown race %q", in.Profile.RaceName)
	}

	talentResult, err := encodeTalents(in.Talents, in.Talent)
	if err != nil {
		return "", Report{}, err
	}
	gear, skippedSlots, noSuffix, err := encodeGear(in.Equipment, in.Suffixes)
	if err != nil {
		return "", Report{}, err
	}

	treeFields := make([]string, 3)
	for i := 0; i < 3; i++ {
		treeFields[i] = encodeTree(talentResult.TreeRanks[i])
	}

	code = strings.Join([]string{
		"FS1", in.Build, in.Profile.ClassSlug, raceSlug, strings.Join(treeFields, "/"), gear,
	}, ":")
	return code, Report{
		UnmatchedTalents: talentResult.Unmatched, Clamped: talentResult.Clamped,
		NoSuffix: noSuffix, SkippedSlots: skippedSlots,
		MatchedByName: talentResult.MatchedByName, MatchedByID: talentResult.MatchedByID,
		MatchedBySpell: talentResult.MatchedBySpell,
	}, nil
}

// base36Digits is every legal single base-36 digit, lowercase, indexed by value — a talent
// rank never exceeds this game's single-digit max_rank, so encodeTree looks up a digit
// rather than reaching for strconv/fmt for one character at a time.
const base36Digits = "0123456789abcdefghijklmnopqrstuvwxyz"

// encodeTree matches web/src/lib/planner/fs1.ts's encodeTree exactly: one base-36 digit per
// rank in tab order, trailing zeros trimmed, "0" for an empty tree.
func encodeTree(ranks []int) string {
	var b strings.Builder
	for _, r := range ranks {
		if r < 0 {
			r = 0
		}
		if r >= len(base36Digits) {
			r = len(base36Digits) - 1
		}
		b.WriteByte(base36Digits[r])
	}
	trimmed := strings.TrimRight(b.String(), "0")
	if trimmed == "" {
		return "0"
	}
	return trimmed
}
