package coach

import (
	"regexp"
	"strconv"
	"strings"
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
)

// NumberToken is one number as the prose wrote it (Raw) and as the pack may state it (Norm).
type NumberToken struct {
	Raw, Norm string
	Kind      string // KindNumber, KindPercent, KindTime or KindRange
}

const (
	timePattern     = `\d{1,2}:\d{2}`
	rangeJoin       = `(?:\s*[–—-]\s*|\s+to\s+)`
	plainNumber     = `\d{1,3}(?:,\d{3})+(?:\.\d+)?|\d+(?:\.\d+)?`
	numberSuffix    = `\s*%|\s+per\s?cent\b|[kK]\b`
	numberPatternRE = `(` + timePattern + `)` + rangeJoin + `(` + timePattern + `)|(` + timePattern + `)|(` + plainNumber + `)(` + numberSuffix + `)?`
)

var numberRE = regexp.MustCompile(numberPatternRE)

// Submatch groups of numberRE.
const (
	groupRangeFrom = 1
	groupRangeTo   = 2
	groupTime      = 3
	groupNumber    = 4
	groupSuffix    = 5
)

// ExtractNumbers finds every digit sequence in prose: integers, decimals, percentages
// ("44%", "44 %", "44 percent"), thousands separators ("2,140"), "2.1k", m:ss times and
// m:ss ranges joined by a hyphen, dash or "to". Norm is the form the pack states: commas
// removed, a trailing zero dropped, a range joined by an en dash.
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
		case group(groupRangeFrom) != "":
			norm := normTime(group(groupRangeFrom)) + enDash + normTime(group(groupRangeTo))
			out = append(out, NumberToken{Raw: raw, Norm: norm, Kind: KindRange})
		case group(groupTime) != "":
			out = append(out, NumberToken{Raw: raw, Norm: normTime(group(groupTime)), Kind: KindTime})
		default:
			out = append(out, numberToken(raw, group(groupNumber), strings.TrimSpace(group(groupSuffix))))
		}
	}
	return out
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
	v, err := strconv.ParseFloat(strings.ReplaceAll(digits, ",", ""), 64)
	if err != nil {
		return 0
	}
	return v
}

func formatNumber(v float64) string { return strconv.FormatFloat(v, 'f', -1, 64) }

// normTime drops a leading zero from the minutes: "03:05" is "3:05".
func normTime(t string) string {
	minutes, seconds, _ := strings.Cut(t, ":")
	return strconv.Itoa(int(parseNumber(minutes))) + ":" + seconds
}

// numberWordPattern lists the number words above three a model might write in place of
// a digit, including teens, tens, magnitudes and ordinals.
const numberWordPattern = `(?i)\b(?:four|five|six|seven|eight|nine|ten|eleven|twelve|thirteen|fourteen|fifteen|` +
	`sixteen|seventeen|eighteen|nineteen|twenty|thirty|forty|fifty|sixty|seventy|eighty|ninety|` +
	`hundred|thousand|million|dozen|fourth|fifth|sixth|seventh|eighth|ninth|tenth)s?\b`

var numberWordRE = regexp.MustCompile(numberWordPattern)

// NumberWords finds number words above "three", lowercased, in the order written.
func NumberWords(s string) []string {
	found := numberWordRE.FindAllString(s, -1)
	for i, w := range found {
		found[i] = strings.ToLower(w)
	}
	return found
}
