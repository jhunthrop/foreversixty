package coach

import (
	"strconv"
	"strings"
)

// romanPattern is an ASCII roman numeral from II to XXXIX, or a lone V or X ("Rank IV",
// "Rank V"), in capitals. A lone I is the pronoun and never a number.
const romanPattern = `\b(?:X{1,3}(?:IX|IV|VI{0,3}|I{0,3})|IX|IV|VI{0,3}|I{2,3})\b`

// numberFormPattern is a run of number forms outside \p{Nd} (superscripts, subscripts, roman
// numeral signs, circled numbers, ...). The vulgar fraction signs are number words instead.
const numberFormPattern = `[\p{No}\p{Nl}]+`

// vulgarFractions are the fraction signs NumberWords reports.
const vulgarFractions = "¼½¾⅐⅑⅒⅓⅔⅕⅖⅗⅘⅙⅚⅛⅜⅝⅞"

var romanValues = map[byte]int{'I': 1, 'V': 5, 'X': 10}

// romanValue is the value of a numeral romanPattern matched.
func romanValue(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {
		v := romanValues[s[i]]
		if i+1 < len(s) && v < romanValues[s[i+1]] {
			total -= v
		} else {
			total += v
		}
	}
	return total
}

// numberFormRange is a block of number forms whose values count up from base.
type numberFormRange struct {
	lo, hi rune
	base   int
}

var numberFormRanges = []numberFormRange{
	{'⁰', '⁰', 0}, {'¹', '¹', 1}, {'²', '³', 2}, {'⁴', '⁹', 4}, // superscripts
	{'₀', '₉', 0}, // subscripts
	{'Ⅰ', 'Ⅻ', 1}, // Ⅰ-Ⅻ
	{'ⅰ', 'ⅻ', 1}, // ⅰ-ⅻ
	{'①', '⑳', 1}, // ①-⑳
	{'⑴', '⒇', 1}, // ⑴-⒇
	{'⒈', '⒛', 1}, // ⒈-⒛
	{'❶', '❿', 1}, // ❶-❿
}

// numberFormNorm is the value a run of number forms writes: each sign's value in turn
// ("²" is "2", "Ⅻ" is "12"). A sign with no known value keeps the run as written, so it
// passes only if the pack wrote the very same signs.
func numberFormNorm(run string) string {
	var b strings.Builder
	for _, r := range run {
		v, ok := numberFormValue(r)
		if !ok {
			return run
		}
		b.WriteString(strconv.Itoa(v))
	}
	return b.String()
}

func numberFormValue(r rune) (int, bool) {
	for _, rg := range numberFormRanges {
		if rg.lo <= r && r <= rg.hi {
			return rg.base + int(r-rg.lo), true
		}
	}
	return 0, false
}

// onlyVulgarFractions reports a run made of fraction signs only.
func onlyVulgarFractions(run string) bool {
	return strings.Trim(run, vulgarFractions) == ""
}
