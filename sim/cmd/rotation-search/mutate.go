package main

import "fmt"

// mutation is one candidate rotation the search round screens, with a
// human label describing what changed from its base.
type mutation struct {
	Label    string
	Rotation rotation
}

// learnedCandidate is one ability the probe or the search may insert:
// a learned spell the rotation does not already cast, with the
// default condition its own shape (dot, persistent toggle, or plain)
// suggests.
type learnedCandidate struct {
	Name           string
	ID, Rank       int
	Condition      action // nil for "no condition"
	ConditionLabel string
}

// ---------------------------------------------------------------------
// List-level mutators: swap, remove, insert. Each returns a brand new
// rotation; base is never modified.
// ---------------------------------------------------------------------

// swapAdjacentMutations is one mutation per adjacent pair in base's
// priority list, each swapping that pair and nothing else. Labels
// carry the 1-based position as well as the spell name: two different
// entries can legally cast the same spell under different conditions
// (an execute clip and a maintenance line both casting Moonfire, for
// one), and the position is what tells them apart in a report.
func swapAdjacentMutations(base rotation, names map[int]string) []mutation {
	list := base.PriorityList
	out := make([]mutation, 0, len(list))
	for i := 0; i+1 < len(list); i++ {
		next := cloneEntries(list)
		next[i], next[i+1] = next[i+1], next[i]
		out = append(out, mutation{
			Label:    fmt.Sprintf("swap #%d (%s) and #%d (%s)", i+1, entryActionLabel(list[i].Action, names), i+2, entryActionLabel(list[i+1].Action, names)),
			Rotation: base.withPriorityList(next),
		})
	}
	return out
}

// removeActionMutations is one mutation per priority-list entry, each
// dropping that entry and nothing else.
func removeActionMutations(base rotation, names map[int]string) []mutation {
	list := base.PriorityList
	out := make([]mutation, 0, len(list))
	for i := range list {
		next := make([]entry, 0, len(list)-1)
		for j, e := range list {
			if j != i {
				next = append(next, cloneEntries([]entry{e})[0])
			}
		}
		out = append(out, mutation{
			Label:    fmt.Sprintf("remove #%d (%s)", i+1, entryActionLabel(list[i].Action, names)),
			Rotation: base.withPriorityList(next),
		})
	}
	return out
}

// insertCandidateMutations is one mutation per (candidate, position)
// pair: candidate inserted at position i of base's priority list
// (0..len, inclusive - len means appended at the end), cast with its
// own default condition.
func insertCandidateMutations(base rotation, candidates []learnedCandidate) []mutation {
	list := base.PriorityList
	var out []mutation
	for _, c := range candidates {
		newEntry := buildCastEntry(fmt.Sprintf("rotation-search: inserted %s, default condition (%s)", c.Name, c.ConditionLabel), c.ID, c.Rank, c.Condition)
		for i := 0; i <= len(list); i++ {
			next := make([]entry, 0, len(list)+1)
			next = append(next, cloneEntries(list[:i])...)
			next = append(next, cloneEntries([]entry{newEntry})[0])
			next = append(next, cloneEntries(list[i:])...)
			out = append(out, mutation{
				Label:    fmt.Sprintf("insert %s at position %d", c.Name, i+1),
				Rotation: base.withPriorityList(next),
			})
		}
	}
	return out
}

// ---------------------------------------------------------------------
// Refresh-condition variants for a maintenance line (one already
// written as not(dotIsActive(...)) or not(auraIsActive(...))).
// ---------------------------------------------------------------------

// refreshVariant is one of the three refresh-condition shapes the
// search tries for a dot/buff maintenance line.
type refreshVariant string

const (
	refreshNotActive    refreshVariant = "not active"
	refreshBelowOneTick refreshVariant = "remaining below one tick"
	refreshBelowThreeS  refreshVariant = "remaining below 3s"
)

// refreshConditionMutations is one mutation per maintenance entry in
// base's priority list, per refresh variant other than the one it
// already uses - identified by whichever entry's condition contains
// the literal not(dotIsActive(id)) or not(auraIsActive(id)) pattern.
// tickSeconds supplies the dot's own tick length for the "below one
// tick" variant (0 when unknown, which skips that variant rather than
// guessing a length).
func refreshConditionMutations(base rotation, tickSeconds map[int]float64, names map[int]string) []mutation {
	var out []mutation
	for i, e := range base.PriorityList {
		kind, id, found := findNotActive(e.Action)
		if !found {
			continue
		}
		variants := refreshVariants(kind, id, tickSeconds[id.SpellID])
		for _, v := range variants {
			if v.variant == refreshNotActive {
				continue // already the base's own condition.
			}
			newCond, replaced := replaceNotActive(e.Action["condition"], kind, id, v.condition)
			if !replaced {
				continue
			}
			newEntry := setCondition(e, newCond.(map[string]any))
			next := cloneEntries(base.PriorityList)
			next[i] = newEntry
			out = append(out, mutation{
				Label:    fmt.Sprintf("#%d (%s): refresh condition -> %s", i+1, entryActionLabel(e.Action, names), v.variant),
				Rotation: base.withPriorityList(next),
			})
		}
	}
	return out
}

type refreshOption struct {
	variant   refreshVariant
	condition action
}

func refreshVariants(kind string, id actionID, tickSeconds float64) []refreshOption {
	out := []refreshOption{{variant: refreshNotActive, condition: notCondition(remainActiveCondition(kind, id))}}
	remaining := remainingTimeValue(kind, id)
	out = append(out, refreshOption{
		variant:   refreshBelowThreeS,
		condition: cmpCondition("OpLe", remaining, "3s"),
	})
	if tickSeconds > 0 {
		out = append(out, refreshOption{
			variant:   refreshBelowOneTick,
			condition: cmpCondition("OpLe", remaining, fmt.Sprintf("%gs", tickSeconds)),
		})
	}
	return out
}

func remainActiveCondition(kind string, id actionID) action {
	if kind == "auraIsActive" {
		return auraIsActiveCondition(id.SpellID, id.Rank)
	}
	return dotIsActiveCondition(id.SpellID, id.Rank)
}

func remainingTimeValue(kind string, id actionID) action {
	if kind == "auraIsActive" {
		return action{"auraRemainingTime": map[string]any{"auraId": buildActionID(id.SpellID, id.Rank)}}
	}
	return dotRemainingTimeValue(id.SpellID, id.Rank)
}

// findNotActive returns the kind ("dotIsActive" or "auraIsActive") and
// ActionID of the first not(dotIsActive(...))/not(auraIsActive(...))
// pattern found anywhere in node - the shape a maintenance line's own
// "keep this up" condition uses.
func findNotActive(node any) (kind string, id actionID, found bool) {
	switch t := node.(type) {
	case map[string]any:
		if notV, ok := t["not"].(map[string]any); ok {
			if val, ok := notV["val"].(map[string]any); ok {
				for _, k := range []string{"dotIsActive", "auraIsActive"} {
					if inner, ok := val[k].(map[string]any); ok {
						idKey := "spellId"
						if k == "auraIsActive" {
							idKey = "auraId"
						}
						if parsedID, ok := readActionID(inner[idKey]); ok {
							return k, parsedID, true
						}
					}
				}
			}
		}
		for _, v := range t {
			if k, i, f := findNotActive(v); f {
				return k, i, f
			}
		}
	case []any:
		for _, v := range t {
			if k, i, f := findNotActive(v); f {
				return k, i, f
			}
		}
	}
	return "", actionID{}, false
}

// replaceNotActive rewrites the first not(dotIsActive(targetID))/
// not(auraIsActive(targetID)) node found in node to replacement,
// leaving everything else untouched and the original node's maps
// unshared with the result (every map on the changed path is rebuilt;
// unchanged subtrees are shared, which is safe because nothing ever
// writes into a shared map after it is built).
func replaceNotActive(node any, targetKind string, targetID actionID, replacement action) (any, bool) {
	switch t := node.(type) {
	case map[string]any:
		if notV, ok := t["not"].(map[string]any); ok {
			if val, ok := notV["val"].(map[string]any); ok {
				idKey := "spellId"
				if targetKind == "auraIsActive" {
					idKey = "auraId"
				}
				if inner, ok := val[targetKind].(map[string]any); ok {
					if parsedID, ok := readActionID(inner[idKey]); ok && parsedID == targetID {
						return map[string]any(replacement), true
					}
				}
			}
		}
		out := make(map[string]any, len(t))
		changed := false
		for k, v := range t {
			nv, ch := replaceNotActive(v, targetKind, targetID, replacement)
			out[k] = nv
			changed = changed || ch
		}
		if !changed {
			return node, false
		}
		return out, true
	case []any:
		out := make([]any, len(t))
		changed := false
		for i, v := range t {
			nv, ch := replaceNotActive(v, targetKind, targetID, replacement)
			out[i] = nv
			changed = changed || ch
		}
		if !changed {
			return node, false
		}
		return out, true
	default:
		return node, false
	}
}

func setCondition(e entry, cond action) entry {
	newAction := cloneAction(e.Action)
	if cond == nil {
		delete(newAction, "condition")
	} else {
		newAction["condition"] = cond
	}
	return entry{Notes: e.Notes, Action: newAction, DoAtValue: cloneAction(e.DoAtValue)}
}

// ---------------------------------------------------------------------
// Resource-gate threshold variants.
// ---------------------------------------------------------------------

// resourcePercentKeys are the APLValue oneof fields a resource-percent
// gate's left-hand side can use. Only currentManaPercent carries data
// in this engine's written rotations today - the proto defines no
// rage or energy *percent* value, so a rage/energy gate there compares
// a raw amount instead, out of scope for the three mana-based specs
// this search's first run covers; the mutator itself is generic and
// will pick up a rage/energy percent key the moment one exists.
var resourcePercentKeys = []string{"currentManaPercent", "currentRagePercent", "currentEnergyPercent"}

// gateThresholds are the percent values the search tries for a
// resource gate, per the task's own list.
var gateThresholds = []string{"0%", "20%", "30%", "40%", "50%"}

type resourceGate struct {
	LHSKey string
	Op     string
	RHS    string
}

func findResourceGate(node any) (resourceGate, bool) {
	switch t := node.(type) {
	case map[string]any:
		if cmp, ok := t["cmp"].(map[string]any); ok {
			if lhs, ok := cmp["lhs"].(map[string]any); ok {
				for _, key := range resourcePercentKeys {
					if _, ok := lhs[key]; ok {
						op, _ := cmp["op"].(string)
						return resourceGate{LHSKey: key, Op: op, RHS: constVal(cmp["rhs"])}, true
					}
				}
			}
		}
		for _, v := range t {
			if g, f := findResourceGate(v); f {
				return g, f
			}
		}
	case []any:
		for _, v := range t {
			if g, f := findResourceGate(v); f {
				return g, f
			}
		}
	}
	return resourceGate{}, false
}

func constVal(v any) string {
	m, ok := v.(map[string]any)
	if !ok {
		return ""
	}
	c, ok := m["const"].(map[string]any)
	if !ok {
		return ""
	}
	s, _ := c["val"].(string)
	return s
}

func replaceResourceGateValue(node any, g resourceGate, newRHS string) (any, bool) {
	switch t := node.(type) {
	case map[string]any:
		if cmp, ok := t["cmp"].(map[string]any); ok {
			if lhs, ok := cmp["lhs"].(map[string]any); ok {
				if _, ok := lhs[g.LHSKey]; ok {
					curOp, _ := cmp["op"].(string)
					if curOp == g.Op && constVal(cmp["rhs"]) == g.RHS {
						newCmp := map[string]any{"op": g.Op, "lhs": cloneAny(lhs), "rhs": map[string]any{"const": map[string]any{"val": newRHS}}}
						return map[string]any{"cmp": newCmp}, true
					}
				}
			}
		}
		out := make(map[string]any, len(t))
		changed := false
		for k, v := range t {
			nv, ch := replaceResourceGateValue(v, g, newRHS)
			out[k] = nv
			changed = changed || ch
		}
		if !changed {
			return node, false
		}
		return out, true
	case []any:
		out := make([]any, len(t))
		changed := false
		for i, v := range t {
			nv, ch := replaceResourceGateValue(v, g, newRHS)
			out[i] = nv
			changed = changed || ch
		}
		if !changed {
			return node, false
		}
		return out, true
	default:
		return node, false
	}
}

// resourceGateMutations is one mutation per resource-gated entry, per
// threshold in gateThresholds other than the one it already uses.
func resourceGateMutations(base rotation, names map[int]string) []mutation {
	var out []mutation
	for i, e := range base.PriorityList {
		cond, ok := e.Action["condition"]
		if !ok {
			continue
		}
		g, found := findResourceGate(cond)
		if !found {
			continue
		}
		for _, threshold := range gateThresholds {
			if threshold == g.RHS {
				continue
			}
			newCond, replaced := replaceResourceGateValue(cond, g, threshold)
			if !replaced {
				continue
			}
			newEntry := setCondition(e, newCond.(map[string]any))
			next := cloneEntries(base.PriorityList)
			next[i] = newEntry
			out = append(out, mutation{
				Label:    fmt.Sprintf("#%d (%s): %s gate -> %s", i+1, entryActionLabel(e.Action, names), g.LHSKey, threshold),
				Rotation: base.withPriorityList(next),
			})
		}
	}
	return out
}

// ---------------------------------------------------------------------
// Execute-phase gate toggle.
// ---------------------------------------------------------------------

// toggleExecuteGateMutations is one mutation per priority-list entry
// that already names isExecutePhase somewhere in its condition - an
// execute-phase nuke's own gate, or one the insertion step added with
// the execute-only default condition - removing a bare gate or, for
// one already stripped down to just that bare gate, nothing else to
// toggle. This never ADDS isExecutePhase to a line that never had
// one: this search has no generic way to tell "this spell is an
// execute-phase ability" from spellconst data alone, and guessing
// would mean testing the gate on every plain-damage line in the list.
func toggleExecuteGateMutations(base rotation, names map[int]string) []mutation {
	var out []mutation
	for i, e := range base.PriorityList {
		if _, ok := castSpellID(e.Action); !ok {
			continue
		}
		cond, hasCond := e.Action["condition"]
		if !hasCond || !hasExecutePhase(cond) {
			continue
		}
		newCond, label, ok := toggleExecuteGate(cond)
		if !ok {
			continue
		}
		newEntry := setCondition(e, newCond)
		next := cloneEntries(base.PriorityList)
		next[i] = newEntry
		out = append(out, mutation{
			Label:    fmt.Sprintf("#%d (%s): execute gate -> %s", i+1, entryActionLabel(e.Action, names), label),
			Rotation: base.withPriorityList(next),
		})
	}
	return out
}

// hasExecutePhase reports whether isExecutePhase appears anywhere in
// node, used only to decide whether a line is eligible for the
// execute-gate toggle at all (see toggleExecuteGateMutations).
func hasExecutePhase(node any) bool {
	switch t := node.(type) {
	case map[string]any:
		if _, ok := t["isExecutePhase"]; ok {
			return true
		}
		for _, v := range t {
			if hasExecutePhase(v) {
				return true
			}
		}
	case []any:
		for _, v := range t {
			if hasExecutePhase(v) {
				return true
			}
		}
	}
	return false
}

func isBareExecutePhase(cond any) bool {
	m, ok := cond.(map[string]any)
	if !ok {
		return false
	}
	_, ok = m["isExecutePhase"]
	return ok && len(m) == 1
}

// toggleExecuteGate flips one condition's isExecutePhase gate. Three
// shapes are handled: no condition (add a bare gate), a bare
// isExecutePhase condition (remove it), and an AND of exactly two
// clauses where one is a bare isExecutePhase (drop it, keeping the
// other clause alone). Anything else reports ok=false rather than
// guessing which clause is the gate.
func toggleExecuteGate(cond any) (action, string, bool) {
	if cond == nil {
		return isExecutePhaseCondition(), "add", true
	}
	if isBareExecutePhase(cond) {
		return nil, "remove", true
	}
	m, ok := cond.(map[string]any)
	if !ok {
		return nil, "", false
	}
	and, ok := m["and"].(map[string]any)
	if !ok {
		return nil, "", false
	}
	vals, ok := and["vals"].([]any)
	if !ok || len(vals) != 2 {
		return nil, "", false
	}
	for i, v := range vals {
		if isBareExecutePhase(v) {
			other := vals[1-i]
			if am, ok := other.(map[string]any); ok {
				return cloneAny(am).(map[string]any), "remove", true
			}
		}
	}
	return action{"and": map[string]any{"vals": []any{cloneAny(m), map[string]any(isExecutePhaseCondition())}}}, "add", true
}
