package sims

import (
	"math"
	"slices"
	"strconv"
	"strings"
	"unicode"

	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

// maxHeadline bounds the stored line, in runes. A set or loadout name is
// a member's own string, so the line it lands in needs a bound.
const maxHeadline = 120

// headlineWeights is how many stat weights the weights headline names:
// the reference and its nearest rival, which is what makes the line mean
// anything at a glance.
const headlineWeights = 2

// weightSeparator joins them. It is a middle dot with spaces, as the
// contract's example spells it.
const weightSeparator = " · "

// Headline is the one line a history row shows for a finished result.
// Contract section 8 fixes one form per kind and 10.6 adds the "… and N
// more" and empty-result forms; this function is the only place any of
// them is composed, and the only place a figure on that list is
// formatted.
//
// It reads the stored result and nothing else — no item table, no loot
// table — which is why sim/bulk fills Substitution.Name for items and
// copies SourceName onto the substitution (contract A6).
func Headline(res simapi.SimResult) string {
	var line string
	switch kind := res.Request.Kind(); kind {
	case simapi.KindWeights:
		line = weightsHeadline(res.Weights)
	case simapi.KindDrops:
		line = dropsHeadline(res.Combos)
	case simapi.KindGear, simapi.KindTalents:
		line = comboHeadline(kind, res.Combos)
	default:
		line = withThousands(roundDPS(res.DPS.Mean)) + " DPS"
	}
	return trimRunes(line, maxHeadline)
}

// comboHeadline is Top Gear's and talent compare's line: what the best
// combination gained, and what it swapped in to gain it.
func comboHeadline(kind string, combos []simapi.Combo) string {
	if len(combos) == 0 {
		return "no combinations"
	}
	best := combos[0]
	delta := signedDPS(best.Delta.Mean)
	name := substitutionName(best.Substitutions)
	switch {
	case name == "":
		return delta
	case kind == simapi.KindTalents:
		return delta + " with '" + name + "'"
	default:
		return delta + " from " + name
	}
}

// substitutionName names a combination: the first substitution, and how
// many others rode with it (contract 10.6). An item with no name falls
// back to its id rather than vanishing — an unnamed item is a data gap
// worth seeing on the page.
func substitutionName(subs []simapi.Substitution) string {
	if len(subs) == 0 {
		return ""
	}
	name := subs[0].Name
	if name == "" && subs[0].ItemID != 0 {
		name = "item #" + strconv.Itoa(subs[0].ItemID)
	}
	if name == "" {
		return ""
	}
	if rest := len(subs) - 1; rest > 0 {
		name += " and " + strconv.Itoa(rest) + " more"
	}
	return name
}

// dropsHeadline is Droptimizer's line: how many of a source's drops beat
// what is equipped, and which source they came from.
func dropsHeadline(combos []simapi.Combo) string {
	upgrades := 0
	for _, c := range combos {
		if c.Delta.Mean > 0 {
			upgrades++
		}
	}
	line := strconv.Itoa(upgrades) + " upgrades"
	switch upgrades {
	case 0:
		line = "no upgrades"
	case 1:
		line = "1 upgrade"
	}
	if source := sharedSource(combos); source != "" {
		return line + " on " + source
	}
	return line
}

// sharedSource is the one loot source every substitution came from, or
// "" when they came from more than one, or from one the page could not
// name. A Droptimizer run is normally one boss; a request that mixed two
// is counted without being named rather than named wrongly.
//
// The name is Substitution.SourceName (contract A6), which the page
// filled from loot.json when it built the request. The API has no loot
// table of its own and never invents one from an origin id.
func sharedSource(combos []simapi.Combo) string {
	source := ""
	for _, c := range combos {
		for _, s := range c.Substitutions {
			if s.SourceName == "" {
				return ""
			}
			if source == "" {
				source = s.SourceName
			} else if source != s.SourceName {
				return ""
			}
		}
	}
	return source
}

// weightsHeadline is the stat weights line: the two largest weights, each
// to two decimals. The reference stat is exactly 1 (contract 2), so it
// leads unless something beat it, which is itself worth seeing.
func weightsHeadline(weights []simapi.StatWeight) string {
	if len(weights) == 0 {
		return "no weights"
	}
	sorted := slices.Clone(weights)
	slices.SortFunc(sorted, func(a, b simapi.StatWeight) int {
		switch {
		case a.Weight > b.Weight:
			return -1
		case a.Weight < b.Weight:
			return 1
		default:
			// A stable order, so two equal weights do not swap between
			// reads and make the same result read differently twice.
			return strings.Compare(a.Stat, b.Stat)
		}
	})
	if len(sorted) > headlineWeights {
		sorted = sorted[:headlineWeights]
	}
	parts := make([]string, 0, len(sorted))
	for _, w := range sorted {
		parts = append(parts, statLabel(w.Stat)+" "+strconv.FormatFloat(w.Weight, 'f', 2, 64))
	}
	return strings.Join(parts, weightSeparator)
}

// statLabel turns a stat id into its display form: "attack_power" becomes
// "Attack Power", "crit" becomes "Crit". The vocabulary is the fork's
// proto.Stat enum names in snake case (contract 10.8 — there is no
// melee_crit or spell_crit, only one crit), so there is no table to keep
// in step: a stat the engine gains is labelled the day it appears.
func statLabel(stat string) string { return titleWords(strings.ReplaceAll(stat, "_", " ")) }

// titleWords upper-cases the first rune of each word. It works by rune,
// so a name in any script keeps its characters whole.
func titleWords(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		r := []rune(w)
		words[i] = string(unicode.ToUpper(r[0])) + string(r[1:])
	}
	return strings.Join(words, " ")
}

// signedDPS is a delta: always signed, so "+0 DPS" reads as a comparison
// that found nothing rather than as an absolute figure.
func signedDPS(mean float64) string {
	n := roundDPS(mean)
	if n < 0 {
		return "-" + withThousands(-n) + " DPS"
	}
	return "+" + withThousands(n) + " DPS"
}

// roundDPS is how every DPS figure on the history list is rounded: to the
// nearest whole, halves away from zero, which is math.Round's own rule.
func roundDPS(mean float64) int64 { return int64(math.Round(mean)) }

// withThousands groups n into three-digit blocks with commas. There is no
// dependency for this: golang.org/x/text's printer is a locale package
// for one string, and the site's numbers are English-grouped everywhere
// else too.
func withThousands(n int64) string {
	sign := ""
	if n < 0 {
		sign, n = "-", -n
	}
	digits := strconv.FormatInt(n, 10)
	var b strings.Builder
	for i := range digits {
		if i > 0 && (len(digits)-i)%3 == 0 {
			b.WriteByte(',')
		}
		b.WriteByte(digits[i])
	}
	return sign + b.String()
}
