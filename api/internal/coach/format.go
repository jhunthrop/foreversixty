package coach

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Clock is a fight time as m:ss from the pull (185000 -> "3:05"). Negative times read as
// the pull.
func Clock(ms int64) string {
	secs := max(ms, 0) / 1000
	return fmt.Sprintf("%d:%02d", secs/60, secs%60)
}

// Span is a fight window as m:ss–m:ss, with an en dash.
func Span(from, to int64) string { return Clock(from) + "–" + Clock(to) }

// Thousands is a whole number with comma separators (2140 -> "2,140").
func Thousands(n int64) string {
	digits := strconv.FormatInt(n, 10)
	sign := ""
	if n < 0 {
		sign, digits = "-", digits[1:]
	}
	var b strings.Builder
	for i, d := range digits {
		if i > 0 && (len(digits)-i)%3 == 0 {
			b.WriteByte(',')
		}
		b.WriteRune(d)
	}
	return sign + b.String()
}

// Pct is a whole percent (44.4 -> "44%").
func Pct(v float64) string { return PctNumber(v) + "%" }

// PctNumber is Pct without the sign, for templates that carry their own "%".
func PctNumber(v float64) string { return strconv.FormatFloat(math.Round(v), 'f', 0, 64) }

// PerMinute is a rate to one decimal (9.06 -> "9.1").
func PerMinute(v float64) string { return strconv.FormatFloat(v, 'f', 1, 64) }

// ClassPlural is a class name as the copy counts it ("Warrior" -> "warriors").
func ClassPlural(class string) string { return strings.ToLower(class) + "s" }
