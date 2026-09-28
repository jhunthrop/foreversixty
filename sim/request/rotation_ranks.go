package request

import (
	"encoding/json"
	"fmt"

	"github.com/jhunthrop/foreversixty/sim/internal/spellranks"
)

// rewriteRotationRanks rewrites every ranked spell's id in an embedded
// APL's JSON to the highest rank the character's class has learned by
// level, and drops any top-level action that CASTS a ranked spell with
// no rank learned yet (a condition naming one is the engine's to nil
// out: see castKeys).
//
// It operates on the APL's JSON representation rather than the parsed
// proto.APLRotation: an ActionID always serializes as a message field
// whose OWN "spellId" key holds a JSON number (protojson int32
// encoding), while the FIELD that carries an ActionID is a JSON object
// under a key also named "spellId" - {"castSpell": {"spellId":
// {"spellId": 9835, "rank": 10}}} - so a plain recursive walk over
// map[string]any/[]any already distinguishes "this map IS an ActionID"
// (its "spellId" value is a number) from "this map HOLDS one" (its
// "spellId" value is an object) with no proto reflection and no hand
// list of the wrapper messages (APLActionCastSpell, APLValueSpellCanCast,
// APLValueDotIsActive, ...) that carry one. It also sidesteps ever
// producing a *proto.APLValue with an empty oneof or an *proto.APLAction
// the engine has to special-case: every rewrite lands in the same JSON
// the embed already ships, and protojson.Unmarshal (rotation, in
// request.go) parses the result exactly like the original file.
//
// Root actions live in two top-level arrays, "prepullActions" and
// "priorityList", each an array of {notes?, action: {...}, doAtValue?}.
// A nested action list - today only priest-shadow's strictSequence -
// serializes its own steps as a bare "actions" array of APLAction
// objects with no {action: ...} wrapper. Both shapes are walked
// identically: rewriteRankedSpellIDs recurses into every map value and
// every list element regardless of the key it sits under, so a spellId
// nested inside a strictSequence's actions is found the same way one
// directly on a priorityList entry is.
//
// Dropping is all-or-nothing per TOP-LEVEL entry, never per spell
// inside a compound action: if any CAST ActionID anywhere under a
// prepullActions/priorityList entry - the action's own spell, or a
// nested sequence's step - belongs to a ranked spell with no learned
// rank, the WHOLE entry is dropped, including every step of an
// enclosing sequence. A sequence is a fixed order of casts; the engine
// has no notion of "skip step 3 of this strict sequence," so trying to
// keep the rest of one whose spell dropped out from under it would be
// inventing a fourth thing for the engine to tolerate rather than
// reusing the two it already does (an action list a rewrite happened
// to leave shorter, and - not used here, see below - an APLValue with
// no oneof set). rot.newValueAnd's own nil-filtering (sim/core,
// wowsims-forever fork) proves the engine already drops a nil
// condition value from an AND/OR list; dropping the whole action here
// instead is the simpler of the two options the level-aware sim design
// doc offers, chosen because it needs no case-by-case proof that every
// OTHER value context (Cmp, Math, Min/Max, a bare condition) tolerates
// a nilled-out operand the same way AND/OR does - dropping the action
// is safe wherever a spellId can appear, not just inside AND/OR.
func rewriteRotationRanks(raw []byte, class string, level int) ([]byte, error) {
	var root map[string]any
	if err := json.Unmarshal(raw, &root); err != nil {
		return nil, fmt.Errorf("the embedded APL is not valid JSON: %w", err)
	}
	resolve := func(id int32) (int32, bool) {
		return spellranks.HighestLearnedSpellID(class, id, level)
	}
	for _, key := range []string{"prepullActions", "priorityList"} {
		if v, ok := root[key]; ok {
			root[key] = filterRankedActions(v, resolve)
		}
	}
	out, err := json.Marshal(root)
	if err != nil {
		return nil, fmt.Errorf("re-encoding the rewritten APL: %w", err)
	}
	return out, nil
}

// filterRankedActions rewrites and filters one top-level action array
// (prepullActions or priorityList). Anything that is not an array
// (there is none today, but a malformed or future field should not
// panic) passes through unchanged.
func filterRankedActions(v any, resolve func(int32) (int32, bool)) any {
	arr, ok := v.([]any)
	if !ok {
		return v
	}
	kept := make([]any, 0, len(arr))
	for _, item := range arr {
		drop := false
		rewriteRankedSpellIDs(item, resolve, &drop, false)
		if !drop {
			kept = append(kept, item)
		}
	}
	return kept
}

// castKeys are the APLAction fields whose ActionID IS the thing the
// action casts. An unlearned spell under one of these drops the whole
// top-level entry; anywhere else - a condition's spellTimeToReady, a
// dotRemainingTime, a spellIsReady - the id is left as written, and the
// engine's own rule takes over: rot.GetAPLSpell finds no such spell,
// the value builds to nil, and newValueAnd/newValueCompare drop a nil
// operand (sim/core/apl_values_operators.go, wowsims-forever fork). A
// level-18 hunter's Serpent Sting, conditioned on Aimed Shot's cooldown
// (learned at 20), used to lose the whole action to that condition and
// sim on auto shots and Multi-Shot alone (found 2026-09-28).
var castKeys = map[string]bool{
	"castSpell":         true,
	"channelSpell":      true,
	"castFriendlySpell": true,
	"multidot":          true,
	"multishield":       true,
}

// rewriteRankedSpellIDs walks node (a JSON object, array, or scalar)
// looking for ActionID objects - {"spellId": <number>, "rank":
// <number>?} - generically: any map whose OWN "spellId" entry is a
// JSON number, rather than a hand list of the messages that embed one.
// A ranked spell resolves to the character's learned rank in place. One
// belonging to a ranked spell with nothing learned sets *drop when it
// is the action's own cast (underCast: the node sits under one of
// castKeys) and is otherwise left as written for the engine to nil out
// (see castKeys); either way its "rank" annotation is cosmetic and not
// worth rewriting on the way out.
func rewriteRankedSpellIDs(node any, resolve func(int32) (int32, bool), drop *bool, underCast bool) {
	switch v := node.(type) {
	case map[string]any:
		if raw, ok := v["spellId"]; ok {
			if num, isNumber := raw.(float64); isNumber {
				id := int32(num)
				newID, learned := resolve(id)
				if !learned {
					if underCast {
						*drop = true
					}
					return
				}
				if newID != id {
					v["spellId"] = float64(newID)
				}
				return
			}
		}
		for key, val := range v {
			rewriteRankedSpellIDs(val, resolve, drop, underCast || castKeys[key])
		}
	case []any:
		for _, item := range v {
			rewriteRankedSpellIDs(item, resolve, drop, underCast)
		}
	}
}
