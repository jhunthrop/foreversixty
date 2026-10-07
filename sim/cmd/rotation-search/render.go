package main

import (
	"fmt"

	"github.com/jhunthrop/foreversixty/sim/request"
)

// spellName is id -> ability name, best-effort (request.LearnedAbilities
// at this spec's level, which is every learned ability spellranks.json
// tracks). An id outside that table - a pseudo-action with no spell
// id, or a real spell spellranks.json does not carry - renders as its
// bare number.
func spellLabel(names map[int]string, id, rank int) string {
	name, ok := names[id]
	if !ok {
		name = fmt.Sprintf("spell %d", id)
	}
	if rank > 0 {
		return fmt.Sprintf("%s (rank %d)", name, rank)
	}
	return name
}

// renderEntry is one priority-list entry as a reader would say it:
// the spell it casts (or the bare action name, for a non-cast action
// such as autocastOtherCooldowns) and its condition in words.
func renderEntry(e entry, names map[int]string) string {
	label := entryActionLabel(e.Action, names)
	cond, ok := e.Action["condition"]
	if !ok {
		return fmt.Sprintf("%s -- no condition", label)
	}
	return fmt.Sprintf("%s -- condition: %s", label, renderCondition(cond, names))
}

func entryActionLabel(a action, names map[int]string) string {
	if id, ok := castSpellID(a); ok {
		return spellLabel(names, id.SpellID, id.Rank)
	}
	for k := range a {
		return fmt.Sprintf("(%s)", k)
	}
	return "(empty action)"
}

// renderCondition renders a condition/value subtree into words. It
// covers every shape this program's three target specs' curated
// rotations use (and, not, cmp, dotIsActive, currentManaPercent,
// numberTargets, const) plus the shapes this search's own mutators
// introduce (auraIsActive, isExecutePhase, dotRemainingTime,
// auraRemainingTime). An unrecognised shape renders as its bare key
// name rather than failing the report.
func renderCondition(node any, names map[int]string) string {
	m, ok := node.(map[string]any)
	if !ok {
		return fmt.Sprintf("%v", node)
	}
	if and, ok := m["and"].(map[string]any); ok {
		return joinVals(and["vals"], " and ", names)
	}
	if or, ok := m["or"].(map[string]any); ok {
		return joinVals(or["vals"], " or ", names)
	}
	if not, ok := m["not"].(map[string]any); ok {
		return fmt.Sprintf("not (%s)", renderCondition(not["val"], names))
	}
	if cmp, ok := m["cmp"].(map[string]any); ok {
		op, _ := cmp["op"].(string)
		return fmt.Sprintf("%s %s %s", renderCondition(cmp["lhs"], names), opWords(op), renderCondition(cmp["rhs"], names))
	}
	if c, ok := m["const"].(map[string]any); ok {
		return fmt.Sprintf("%v", c["val"])
	}
	if dot, ok := m["dotIsActive"].(map[string]any); ok {
		return fmt.Sprintf("%s dot is active", idLabel(dot["spellId"], names))
	}
	if dot, ok := m["dotRemainingTime"].(map[string]any); ok {
		return fmt.Sprintf("%s dot remaining time", idLabel(dot["spellId"], names))
	}
	if aura, ok := m["auraIsActive"].(map[string]any); ok {
		return fmt.Sprintf("%s aura is active", idLabel(aura["auraId"], names))
	}
	if aura, ok := m["auraRemainingTime"].(map[string]any); ok {
		return fmt.Sprintf("%s aura remaining time", idLabel(aura["auraId"], names))
	}
	if _, ok := m["isExecutePhase"]; ok {
		return "execute phase"
	}
	if _, ok := m["numberTargets"]; ok {
		return "number of targets"
	}
	for _, key := range []string{"currentManaPercent", "currentRagePercent", "currentEnergyPercent"} {
		if _, ok := m[key]; ok {
			return key
		}
	}
	for k := range m {
		return k
	}
	return "(empty)"
}

func idLabel(v any, names map[int]string) string {
	id, ok := readActionID(v)
	if !ok {
		return "?"
	}
	return spellLabel(names, id.SpellID, id.Rank)
}

func joinVals(v any, sep string, names map[int]string) string {
	lst, ok := v.([]any)
	if !ok || len(lst) == 0 {
		return "(empty)"
	}
	parts := make([]string, len(lst))
	for i, c := range lst {
		parts[i] = "(" + renderCondition(c, names) + ")"
	}
	out := parts[0]
	for _, p := range parts[1:] {
		out += sep + p
	}
	return out
}

func opWords(op string) string {
	switch op {
	case "OpEq":
		return "=="
	case "OpNe":
		return "!="
	case "OpLt":
		return "<"
	case "OpLe":
		return "<="
	case "OpGt":
		return ">"
	case "OpGe":
		return ">="
	default:
		return op
	}
}

// buildSpellNames is id -> ability name for every ability
// spellranks.json tracks by level - a superset of the probe-worthy
// candidates, so a baseline action's existing spell (Lightning Bolt,
// say) still renders by name rather than by bare id.
func buildSpellNames(repoRoot, build, class string, level int) (map[int]string, error) {
	abilities, err := request.LearnedAbilities(repoRoot, build, class, level)
	if err != nil {
		return nil, err
	}
	out := map[int]string{}
	for _, a := range abilities {
		for _, id := range a.IDs {
			out[id] = a.Name
		}
	}
	return out, nil
}
