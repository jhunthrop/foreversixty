package coach

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Token kinds ExtractNumbers reports.
const (
	KindNumber  = "number"
	KindPercent = "percent"
	KindTime    = "time"
	KindRange   = "range"
)

const (
	enDash        = "–"
	thousandsStep = 1000
	digitsPerRun  = 10
)

// NumberToken is one number as the prose wrote it (Raw) and as the pack may state it (Norm).
type NumberToken struct {
	Raw, Norm string
	Kind      string // KindNumber, KindPercent, KindTime or KindRange
}

const (
	timePattern     = `\p{Nd}{1,2}:\p{Nd}{2}`
	rangeJoin       = `(?:\s*[–—-]\s*|\s+to\s+)`
	plainNumber     = `\p{Nd}{1,3}(?:,\p{Nd}{3})+(?:\.\p{Nd}+)?|\p{Nd}+(?:\.\p{Nd}+)?`
	numberSuffix    = `\s*%|\s+per\s?cent\b|[kK]\b`
	betweenRange    = `(?i:between)\s+(` + timePattern + `)\s+(?i:and)\s+(` + timePattern + `)`
	numberPatternRE = betweenRange + `|(` + timePattern + `)` + rangeJoin + `(` + timePattern + `)|(` + timePattern + `)|(` + plainNumber + `)(` + numberSuffix + `)?`
)

var numberRE = regexp.MustCompile(numberPatternRE)

// Submatch groups of numberRE.
const (
	groupBetweenFrom = 1
	groupBetweenTo   = 2
	groupRangeFrom   = 3
	groupRangeTo     = 4
	groupTime        = 5
	groupNumber      = 6
	groupSuffix      = 7
)

// ExtractNumbers finds every digit sequence in prose: integers, decimals, percentages
// ("44%", "44 %", "44 percent"), thousands separators ("2,140"), "2.1k", m:ss times and
// m:ss ranges joined by a hyphen, dash, "to" or "between ... and ...". Digits of any script
// count. Norm is the form the pack states: ASCII digits, commas removed, a trailing zero
// dropped, a range joined by an en dash. A range is one token: it passes only when the
// pack holds that very pair, never because its two ends appear in different ranges.
func ExtractNumbers(s string) []NumberToken {
	var out []NumberToken
	for _, m := range numberRE.FindAllStringSubmatchIndex(s, -1) {
		group := func(i int) string {
			if m[2*i] < 0 {
				return ""
			}
			return s[m[2*i]:m[2*i+1]]
		}
		raw := strings.TrimSpace(s[m[0]:m[1]])
		switch {
		case group(groupBetweenFrom) != "":
			out = append(out, rangeToken(raw, group(groupBetweenFrom), group(groupBetweenTo)))
		case group(groupRangeFrom) != "":
			out = append(out, rangeToken(raw, group(groupRangeFrom), group(groupRangeTo)))
		case group(groupTime) != "":
			out = append(out, NumberToken{Raw: raw, Norm: normTime(group(groupTime)), Kind: KindTime})
		default:
			out = append(out, numberToken(raw, group(groupNumber), strings.TrimSpace(group(groupSuffix))))
		}
	}
	return out
}

func rangeToken(raw, from, to string) NumberToken {
	return NumberToken{Raw: raw, Norm: normTime(from) + enDash + normTime(to), Kind: KindRange}
}

func numberToken(raw, digits, suffix string) NumberToken {
	value := parseNumber(digits)
	switch {
	case strings.EqualFold(suffix, "k"):
		value *= thousandsStep
	case suffix != "":
		return NumberToken{Raw: raw, Norm: formatNumber(value) + "%", Kind: KindPercent}
	}
	return NumberToken{Raw: raw, Norm: formatNumber(value), Kind: KindNumber}
}

// parseNumber reads digits with thousands separators; ExtractNumbers' pattern guarantees
// they parse.
func parseNumber(digits string) float64 {
	v, err := strconv.ParseFloat(strings.ReplaceAll(asciiDigits(digits), ",", ""), 64)
	if err != nil {
		return 0
	}
	return v
}

// asciiDigits rewrites digits of any script (Arabic-Indic, full-width, ...) as ASCII. Unicode
// lays each script's digits out in runs of ten starting at zero.
func asciiDigits(s string) string {
	return strings.Map(func(r rune) rune {
		if r < utf8.RuneSelf || !unicode.Is(unicode.Nd, r) {
			return r
		}
		for _, rg := range unicode.Nd.R16 {
			if rune(rg.Lo) <= r && r <= rune(rg.Hi) {
				return '0' + (r-rune(rg.Lo))%digitsPerRun
			}
		}
		for _, rg := range unicode.Nd.R32 {
			if rune(rg.Lo) <= r && r <= rune(rg.Hi) {
				return '0' + (r-rune(rg.Lo))%digitsPerRun
			}
		}
		return r
	}, s)
}

func formatNumber(v float64) string { return strconv.FormatFloat(v, 'f', -1, 64) }

// normTime drops a leading zero from the minutes: "03:05" is "3:05".
func normTime(t string) string {
	minutes, seconds, _ := strings.Cut(t, ":")
	return strconv.Itoa(int(parseNumber(minutes))) + ":" + asciiDigits(seconds)
}

// numberWordPattern lists the number words above three a model might write in place of
// a digit: teens, tens, magnitudes, ordinals, plus the vulgar fraction signs.
const numberWordPattern = `(?i)\b(?:four|five|six|seven|eight|nine|ten|eleven|twelve|thirteen|fourteen|fifteen|` +
	`sixteen|seventeen|eighteen|nineteen|twenty|thirty|forty|fifty|sixty|seventy|eighty|ninety|` +
	`hundred|thousand|million|dozen|fourth|fifth|sixth|seventh|eighth|ninth|tenth)s?\b|[\x{00BC}-\x{00BE}\x{2150}-\x{215E}]`

// quantifiedWordPattern lists fractions and multiples ("half", "a third", "two-thirds",
// "a quarter", "twice", "double", "triple") only where they quantify something: followed by
// "of", "as many/much/often", a determiner or a digit. "Cost you twice." and "past half on a
// pull" are idiom and pass; "half the casts" and "twice as often" are numeric claims.
const quantifiedWordPattern = `(?i)\b(half|halves|thirds?|quarters?|twice|double|doubled|triple|tripled)\b` +
	`(?:\s+(?:of|as|the|an?|your|their|its|all|every|each|that|this|those|these|many|much|often)\b|\s+\p{Nd})`

// alwaysQuantifiedPattern lists the phrases that are a fraction whatever follows: "a third",
// "a quarter" ("It fell by a third", "a third were late") and "by half".
const alwaysQuantifiedPattern = `(?i)\b(?:a\s+(third|quarter)|by\s+(half))\b`

var (
	numberWordRE         = regexp.MustCompile(numberWordPattern)
	quantifiedWordRE     = regexp.MustCompile(quantifiedWordPattern)
	alwaysQuantifiedWord = regexp.MustCompile(alwaysQuantifiedPattern)
)

// figurativeHalf is the one quantified use of "half" that counts nothing: "the other half of the story".
const figurativeHalf = "other half"

// wordHit is a number word and where it starts.
type wordHit struct {
	at   int
	word string
}

// alwaysQuantifiedHits finds the phrases of alwaysQuantifiedPattern that known has not
// already reported at the same place.
func alwaysQuantifiedHits(s string, known []wordHit) []wordHit {
	seen := map[int]bool{}
	for _, h := range known {
		seen[h.at] = true
	}
	var out []wordHit
	for _, m := range alwaysQuantifiedWord.FindAllStringSubmatchIndex(s, -1) {
		start, end := m[2], m[3]
		if start < 0 {
			start, end = m[4], m[5]
		}
		if !seen[start] {
			out = append(out, wordHit{start, strings.ToLower(s[start:end])})
		}
	}
	return out
}

// NumberWords finds number words above "three", and fractions and multiples that quantify
// something, lowercased, in the order written.
func NumberWords(s string) []string {
	var hits []wordHit
	for _, loc := range numberWordRE.FindAllStringIndex(s, -1) {
		hits = append(hits, wordHit{loc[0], strings.ToLower(s[loc[0]:loc[1]])})
	}
	for _, m := range quantifiedWordRE.FindAllStringSubmatchIndex(s, -1) {
		word := strings.ToLower(s[m[2]:m[3]])
		if word == "half" && strings.HasSuffix(strings.ToLower(s[:m[3]]), figurativeHalf) {
			continue
		}
		hits = append(hits, wordHit{m[2], word})
	}
	hits = append(hits, alwaysQuantifiedHits(s, hits)...)
	sort.Slice(hits, func(i, j int) bool { return hits[i].at < hits[j].at })
	out := make([]string, len(hits))
	for i, h := range hits {
		out[i] = h.word
	}
	return out
}
