package main

import (
	"sort"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/request"
)

// conditionKind names which of the task's four default-condition
// shapes a candidate's insertion used, for the report.
type conditionKind string

const (
	condPlain      conditionKind = "no condition (cooldown/plain damage)"
	condDOT        conditionKind = "dot: not active"
	condExecute    conditionKind = "execute-only: isExecutePhase"
	condBuffOrForm conditionKind = "buff/form/seal: not active"
	condDebuff     conditionKind = "target debuff: not active on the target"
)

// defaultCondition picks the sensible default condition the task
// calls for, from the shape request.LearnedAbilities already read off
// spellconst: a dot gets "not active" (dotIsActive false); an
// execute-phase nuke (name-matched - spellconst carries no explicit
// execute-phase flag) gets isExecutePhase; a persistent, non-damage
// toggle (duration_ms == -1: a form, a seal, a stance) gets "not
// active" (auraIsActive false); everything else - a cooldown or a
// plain damage spell - gets no condition at all, cast whenever its
// own cooldown and resources allow.
func defaultCondition(a request.LearnedAbility) (action, conditionKind) {
	id, rank := a.IDs[0], a.Rank
	switch {
	case a.IsDOT:
		return notCondition(dotIsActiveCondition(id, rank)), condDOT
	case isExecuteName(a.Name):
		return isExecutePhaseCondition(), condExecute
	case a.RaisesDamageTaken:
		return notCondition(targetAuraIsActiveCondition(id, rank)), condDebuff
	case a.DurationMS == -1 && !a.IsDamage:
		return notCondition(auraIsActiveCondition(id, rank)), condBuffOrForm
	default:
		return nil, condPlain
	}
}

func isExecuteName(name string) bool {
	return strings.Contains(strings.ToLower(name), "execute")
}

// isProbeWorthy is which learned abilities the probe and the search's
// insertion mutator bother trying: a damage-dealing cast
// (a.IsDamage), an execute-phase nuke by name, a persistent
// non-damage toggle (a form, a seal, a stance) - the task's "learned
// damage or DPS-cooldown spell" - or a debuff that raises the damage the
// target takes (Curse of the Elements), and not already in the rotation.
// GCDMS > 0 is required regardless of category: an ability with no
// GCD of its own is not a player-chosen rotation action at all - this
// build's internal "Attack" entry (the white-damage melee swing,
// carrying a damage effect and gcd_ms 0) is exactly this trap, and
// casting it explicitly with no pacing mechanism loops the engine
// forever.
func isProbeWorthy(a request.LearnedAbility) bool {
	if a.GCDMS <= 0 {
		return false
	}
	return a.IsDamage || isExecuteName(a.Name) || a.RaisesDamageTaken || (a.DurationMS == -1 && !a.IsDamage)
}

// learnedCandidates is every learned, probe-worthy ability the
// rotation does not already cast, each with the default condition
// defaultCondition picked for it - the pool both the action probe
// (step 2) and the mutation search's insertion operator (step 3) draw
// from.
func learnedCandidates(repoRoot, build, class string, level int, authored map[int]bool) ([]learnedCandidate, error) {
	abilities, err := request.LearnedAbilities(repoRoot, build, class, level)
	if err != nil {
		return nil, err
	}
	var out []learnedCandidate
	for _, a := range abilities {
		if !isProbeWorthy(a) || anyAuthored(a.IDs, authored) {
			continue
		}
		cond, kind := defaultCondition(a)
		out = append(out, learnedCandidate{
			Name: a.Name, ID: a.IDs[0], Rank: a.Rank,
			Condition: cond, ConditionLabel: string(kind),
		})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out, nil
}

func anyAuthored(ids []int, authored map[int]bool) bool {
	for _, id := range ids {
		if authored[id] {
			return true
		}
	}
	return false
}

// authoredCastIDs is every spell id base's prepull or priority
// actions cast, at any nesting depth - sim/request/ladder.go's own
// extractCastSpellIDs, reapplied to this package's typed entry/action
// tree instead of a bare decoded-JSON one.
func authoredCastIDs(base rotation) map[int]bool {
	out := map[int]bool{}
	collectCastIDs(base.PrepullActions, out)
	collectCastIDs(base.PriorityList, out)
	return out
}

func collectCastIDs(entries []entry, out map[int]bool) {
	for _, e := range entries {
		walkCastIDs(map[string]any(e.Action), out)
		if e.DoAtValue != nil {
			walkCastIDs(map[string]any(e.DoAtValue), out)
		}
	}
}

func walkCastIDs(node any, out map[int]bool) {
	switch t := node.(type) {
	case map[string]any:
		if cs, ok := t["castSpell"].(map[string]any); ok {
			if id, ok := readActionID(cs["spellId"]); ok {
				out[id.SpellID] = true
			}
		}
		for _, v := range t {
			walkCastIDs(v, out)
		}
	case []any:
		for _, v := range t {
			walkCastIDs(v, out)
		}
	}
}

// tickSecondsByID is every probe-worthy dot's own tick length, keyed
// by its representative id - refreshConditionMutations' "below one
// tick" variant needs it for a dot the rotation ALREADY casts, which
// learnedCandidates does not cover (that list is only the unused
// ones), so this reads the same request.LearnedAbilities data again,
// unfiltered by whether the rotation already casts it.
func tickSecondsByID(repoRoot, build, class string, level int) (map[int]float64, error) {
	abilities, err := request.LearnedAbilities(repoRoot, build, class, level)
	if err != nil {
		return nil, err
	}
	out := map[int]float64{}
	for _, a := range abilities {
		if !a.IsDOT || a.TickSeconds <= 0 {
			continue
		}
		for _, id := range a.IDs {
			out[id] = a.TickSeconds
		}
	}
	return out, nil
}
