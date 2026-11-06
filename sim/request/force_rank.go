package request

import (
	"encoding/json"
	"fmt"

	"github.com/jhunthrop/foreversixty/sim/internal/spellranks"
)

// ForceRotationRank is the rank-swap probe's rewrite: every reference to
// authored in a rotation's JSON (casts, conditions such as dotIsActive and
// spellTimeToReady, every line and every pre-pull step) becomes rank rank of
// that spell, as a character of class learns it by level. Round 18 priest X4:
// a probe that swapped only the cast left a DoT condition watching the
// authored rank, so the swapped line recast on every global and Shadow Word:
// Pain rank 4 read -29%. It errs when the rank is not learned by level.
// Run RewriteForcedRotationRanks after it with the forced authored ids: it
// drops the rank rows on them, so the forced rank is what the line casts,
// the authored rank itself included.
func ForceRotationRank(raw []byte, class string, authored int32, rank, level int) ([]byte, error) {
	forced, ok := spellranks.SpellIDOfRank(class, authored, rank, level)
	if !ok {
		return nil, fmt.Errorf("spell %d has no rank %d a %s learns by level %d", authored, rank, class, level)
	}
	var root any
	if err := json.Unmarshal(raw, &root); err != nil {
		return nil, fmt.Errorf("the rotation is not valid JSON: %w", err)
	}
	replaceActionID(root, authored, forced)
	out, err := json.Marshal(root)
	if err != nil {
		return nil, fmt.Errorf("re-encoding the forced rotation: %w", err)
	}
	return out, nil
}

// replaceActionID rewrites, in place, every ActionID (a map whose own
// "spellId" is a JSON number; see rewriteRankedSpellIDs) naming from to name
// to instead, and drops its cosmetic "rank" annotation.
func replaceActionID(node any, from, to int32) {
	switch v := node.(type) {
	case map[string]any:
		if num, ok := v["spellId"].(float64); ok {
			if int32(num) == from {
				v["spellId"] = float64(to)
				delete(v, "rank")
			}
			return
		}
		for _, val := range v {
			replaceActionID(val, from, to)
		}
	case []any:
		for _, item := range v {
			replaceActionID(item, from, to)
		}
	}
}
