// logs/engine/layout/infer.go
package layout

import (
	"sort"
	"strings"

	"github.com/jhunthrop/foreversixty/logs/engine/lexer"
)

// Infer builds a row from the log itself when no registered row matches.
// It counts the fields of every event it sees and derives the advanced-block
// size from the events that must carry one, so an unknown dialect still
// decodes rather than falling back to raw text for the whole file.
//
// The result is always Verified: false and Name "inferred". Where the sample
// does not settle a question the field is left at its zero value; nothing is
// guessed.
func Infer(lines []lexer.Line) Layout {
	l := Layout{
		Name:      "inferred",
		Verified:  false,
		Prefixes:  map[string]int{"SWING": 0, "RANGE": 3, "SPELL_PERIODIC": 3, "SPELL_BUILDING": 3, "SPELL": 3},
		Suffixes:  map[string]Suffix{},
		Specials:  map[string]Special{},
		Combatant: Combatant{},
	}

	widths := map[string]map[int]int{} // event -> width -> count
	var stamp string
	for _, ln := range lines {
		if len(ln.Params) == 0 || ln.Params[0] == "" {
			continue
		}
		if stamp == "" && ln.Stamp != "" {
			stamp = ln.Stamp
		}
		if h, ok := ParseHeader(ln); ok {
			l.Version = h.Version
			l.ProjectID = h.ProjectID
			continue
		}
		ev := ln.Params[0]
		if widths[ev] == nil {
			widths[ev] = map[int]int{}
		}
		widths[ev][len(ln.Params)]++
	}
	l.StampYear, l.StampZone = inferStamp(stamp)

	// The advanced block is whatever is left on a SWING_DAMAGE line after
	// the common header and a damage suffix of 9 or 10 fields. Both
	// candidate suffix sizes are tried and the one that lands on a
	// plausible block size wins; SPELL_CAST_SUCCESS, which has no suffix at
	// all, breaks the tie.
	if w, ok := dominant(widths["SPELL_CAST_SUCCESS"]); ok && w > BaseParams+3 {
		l.Advanced = w - BaseParams - 3
	} else if w, ok := dominant(widths["SWING_DAMAGE"]); ok {
		for _, suffix := range []int{10, 9} {
			if adv := w - BaseParams - suffix; adv == 17 || adv == 19 {
				l.Advanced = adv
				break
			}
		}
	}

	// Iterate the events in sorted order. The reduction below is total on
	// Suffix.Params but not on the whole Suffix struct: two events sharing
	// a suffix can agree on Params and disagree on Advanced, and the write
	// at the end of the loop would then let Go's randomised map order pick
	// the winner. Sorting settles it, and keeps settling it if a later
	// field is added to Suffix.
	for _, ev := range sortedKeys(widths) {
		ws := widths[ev]
		w, ok := dominant(ws)
		if !ok {
			continue
		}
		prefix, suffix, known := splitAny(l.Prefixes, ev)
		if !known {
			all := make([]int, 0, len(ws))
			for width := range ws {
				all = append(all, width)
			}
			sort.Ints(all)
			l.Specials[ev] = Special{Widths: all}
			continue
		}
		adv := advancedForSuffix(suffix)
		params := w - BaseParams - l.Prefixes[prefix]
		if adv {
			params -= l.Advanced
		}
		if params < 0 {
			adv = false
			params = w - BaseParams - l.Prefixes[prefix]
		}
		cur, seen := l.Suffixes[suffix]
		if seen && cur.Params != params {
			// Two events sharing a suffix disagree: keep the smaller count
			// and let the decoder report the mismatch rather than pick.
			if cur.Params < params {
				continue
			}
		}
		l.Suffixes[suffix] = Suffix{Params: params, Advanced: adv}
	}
	return l
}

// advancedForSuffix lists the suffixes that carry the advanced block in
// every dialect documented so far.
func advancedForSuffix(suffix string) bool {
	switch suffix {
	case "_DAMAGE", "_DAMAGE_LANDED", "_HEAL", "_ENERGIZE", "_DRAIN", "_LEECH", "_CAST_SUCCESS":
		return true
	}
	return false
}

func splitAny(prefixes map[string]int, event string) (prefix, suffix string, ok bool) {
	best := ""
	for p := range prefixes {
		if strings.HasPrefix(event, p) && len(p) > len(best) {
			best = p
		}
	}
	if best == "" || len(event) == len(best) {
		return "", "", false
	}
	rest := event[len(best):]
	if !strings.HasPrefix(rest, "_") {
		return "", "", false
	}
	return best, rest, true
}

// sortedKeys returns m's keys in a fixed order, so a loop that reduces over
// them lands on the same answer whatever order the runtime hands them back.
func sortedKeys[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

// dominant returns the most common width, breaking a tie on the smaller
// width so the answer does not depend on map order.
func dominant(ws map[int]int) (int, bool) {
	best, bestN := 0, 0
	for w, n := range ws {
		if n > bestN || (n == bestN && w < best) {
			best, bestN = w, n
		}
	}
	return best, bestN > 0
}

func inferStamp(stamp string) (year, zone bool) {
	if stamp == "" {
		return false, false
	}
	date, clock, ok := strings.Cut(stamp, " ")
	if !ok {
		return false, false
	}
	year = strings.Count(date, "/") == 2
	zone = strings.ContainsAny(clock[1:], "+-")
	return year, zone
}
