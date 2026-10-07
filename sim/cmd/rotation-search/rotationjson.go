package main

import (
	"encoding/json"
	"fmt"
)

// action is one APL action or value, kept exactly as the engine's
// protojson encoding would decode or encode it: a generic JSON object
// whose keys are the proto oneof's field names ("condition",
// "castSpell", "autocastOtherCooldowns", "and", "cmp", "dotIsActive",
// ...). This package never needs a typed Go mirror of every
// APLAction/APLValue message - the same convention
// sim/request/rotation_ranks.go's walker and sim/request/ladder.go's
// extractCastSpellIDs already use for reading an APL's JSON - because
// every mutation this search makes either replaces a whole entry's
// action wholesale (a new cast, built by buildCastEntry) or edits one
// named sub-key it already knows how to find (condition, castSpell).
//
// It is a type ALIAS (= map[string]any), not a defined type: a value
// built as `action{...}` and one built as a plain `map[string]any{...}`
// are then the exact same type, so every nested builder and every
// recursive walker below can mix the two freely with no dynamic-type
// mismatches at a type switch or a type assertion.
type action = map[string]any

// entry is one list item of prepullActions or priorityList.
type entry struct {
	Notes     string `json:"notes,omitempty"`
	Action    action `json:"action"`
	DoAtValue action `json:"doAtValue,omitempty"`
}

// rotation is the engine APL JSON shape data/curated/apl/<spec>.json's
// "rotation" field, and sim/request/apl/<spec>.apl.json, both carry -
// only priorityList is mutated by this search; prepullActions is
// carried through unchanged.
type rotation struct {
	Type           string  `json:"type"`
	PrepullActions []entry `json:"prepullActions,omitempty"`
	PriorityList   []entry `json:"priorityList"`
}

// parseRotation decodes one curated or embedded rotation's raw JSON.
func parseRotation(raw json.RawMessage) (rotation, error) {
	var r rotation
	if err := json.Unmarshal(raw, &r); err != nil {
		return rotation{}, fmt.Errorf("parsing the rotation: %w", err)
	}
	return r, nil
}

// marshalIndent renders r as the same pretty JSON shape the curated
// files carry, for the report's pasteable fenced block.
func (r rotation) marshalIndent() ([]byte, error) {
	return json.MarshalIndent(r, "", "    ")
}

// clone is a deep copy of r: every mutator in this package returns a
// new rotation rather than editing one in place, so two candidates
// built from the same base never alias each other's maps.
func (r rotation) clone() rotation {
	return rotation{
		Type:           r.Type,
		PrepullActions: cloneEntries(r.PrepullActions),
		PriorityList:   cloneEntries(r.PriorityList),
	}
}

// withPriorityList is a copy of r with its priority list replaced;
// prepullActions is carried over unchanged, deep-copied the same way
// clone does.
func (r rotation) withPriorityList(list []entry) rotation {
	return rotation{
		Type:           r.Type,
		PrepullActions: cloneEntries(r.PrepullActions),
		PriorityList:   list,
	}
}

func cloneEntries(in []entry) []entry {
	if in == nil {
		return nil
	}
	out := make([]entry, len(in))
	for i, e := range in {
		out[i] = entry{Notes: e.Notes, Action: cloneAction(e.Action), DoAtValue: cloneAction(e.DoAtValue)}
	}
	return out
}

// cloneAction deep-copies an action map (action is an alias for
// map[string]any, so cloneAny's own type switch matches it directly).
func cloneAction(a action) action {
	if a == nil {
		return nil
	}
	return cloneAny(a).(map[string]any)
}

// cloneAny deep-copies a decoded-JSON value: map[string]any, []any, or
// a scalar (string/float64/bool/nil), the only shapes json.Unmarshal
// ever produces into an `any`.
func cloneAny(v any) any {
	switch t := v.(type) {
	case map[string]any:
		out := make(map[string]any, len(t))
		for k, val := range t {
			out[k] = cloneAny(val)
		}
		return out
	case []any:
		out := make([]any, len(t))
		for i, val := range t {
			out[i] = cloneAny(val)
		}
		return out
	default:
		return v
	}
}

// actionID is one castSpell's ActionID, read generically off an
// action map.
type actionID struct {
	SpellID int
	Rank    int
}

// castSpellID reads the top-level castSpell's spell id and rank off
// an action, if it has one (an action with no castSpell - a bare
// autocastOtherCooldowns, for instance - returns ok=false).
func castSpellID(a action) (actionID, bool) {
	cs, ok := a["castSpell"].(map[string]any)
	if !ok {
		return actionID{}, false
	}
	return readActionID(cs["spellId"])
}

// readActionID reads one {"spellId": <id>, "rank": <rank>} ActionID
// node - the same shape castSpell.spellId, dotIsActive.spellId, and
// auraIsActive.auraId all carry.
func readActionID(v any) (actionID, bool) {
	m, ok := v.(map[string]any)
	if !ok {
		return actionID{}, false
	}
	idf, ok := m["spellId"].(float64)
	if !ok {
		return actionID{}, false
	}
	rank := 0
	if rf, ok := m["rank"].(float64); ok {
		rank = int(rf)
	}
	return actionID{SpellID: int(idf), Rank: rank}, true
}

// buildActionID is the JSON an ActionID node carries: {"spellId": id}
// or, ranked, {"spellId": id, "rank": rank}.
func buildActionID(id, rank int) map[string]any {
	m := map[string]any{"spellId": float64(id)}
	if rank > 0 {
		m["rank"] = float64(rank)
	}
	return m
}

// buildCastEntry is one new priority-list entry casting id/rank,
// gated on cond (nil for no condition - cast whenever its own
// cooldown and resources allow).
func buildCastEntry(notes string, id, rank int, cond action) entry {
	act := action{"castSpell": map[string]any{"spellId": buildActionID(id, rank)}}
	if cond != nil {
		act["condition"] = cond
	}
	return entry{Notes: notes, Action: act}
}

// dotIsActiveCondition / auraIsActiveCondition / isExecutePhaseCondition
// are the three named condition shapes the task's default-condition
// vocabulary calls for, built generically off an ActionID.
func dotIsActiveCondition(id, rank int) action {
	return action{"dotIsActive": map[string]any{"spellId": buildActionID(id, rank)}}
}

func auraIsActiveCondition(id, rank int) action {
	return action{"auraIsActive": map[string]any{"auraId": buildActionID(id, rank)}}
}

// targetAuraIsActiveCondition is auraIsActive on the current target's own aura:
// a debuff the caster lands, which the caster's own aura list does not carry.
func targetAuraIsActiveCondition(id, rank int) action {
	return action{"auraIsActive": map[string]any{
		"sourceUnit": map[string]any{"type": "CurrentTarget"},
		"auraId":     buildActionID(id, rank),
	}}
}

func notCondition(inner action) action {
	return action{"not": map[string]any{"val": inner}}
}

func isExecutePhaseCondition() action {
	return action{"isExecutePhase": map[string]any{}}
}

// cmpCondition is {"cmp": {"op": op, "lhs": lhs, "rhs": {"const": {"val": rhsConst}}}}.
func cmpCondition(op string, lhs action, rhsConst string) action {
	return action{"cmp": map[string]any{
		"op":  op,
		"lhs": lhs,
		"rhs": map[string]any{"const": map[string]any{"val": rhsConst}},
	}}
}

func dotRemainingTimeValue(id, rank int) action {
	return action{"dotRemainingTime": map[string]any{"spellId": buildActionID(id, rank)}}
}

// findBareActiveGates collects every dotIsActive/auraIsActive id node
// requires to be PRESENT (a damage- or combo-dependent gate, not a
// not(...)-wrapped maintenance trigger) anywhere in node. It does not
// descend past a "not" key at all: a maintenance line's own
// not(dotIsActive(X)) is a different thing (a refresh trigger, true
// exactly when X is ABSENT) and must never be reported as a presence
// requirement. This is report.go's own tool for flagging a line whose
// gate nothing in the rest of the rotation can ever satisfy - see
// danglingGates.
func findBareActiveGates(node any, out *[]actionID) {
	switch t := node.(type) {
	case map[string]any:
		if _, ok := t["not"]; ok {
			return
		}
		for _, key := range []string{"dotIsActive", "auraIsActive"} {
			inner, ok := t[key].(map[string]any)
			if !ok {
				continue
			}
			idKey := "spellId"
			if key == "auraIsActive" {
				idKey = "auraId"
			}
			if id, ok := readActionID(inner[idKey]); ok {
				*out = append(*out, id)
			}
		}
		for _, v := range t {
			findBareActiveGates(v, out)
		}
	case []any:
		for _, v := range t {
			findBareActiveGates(v, out)
		}
	}
}
