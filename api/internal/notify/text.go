package notify

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// markdownSpecials are escaped wherever they appear in member text.
const markdownSpecials = "\\*_~`|>#[]"

// lineStartSpecials are escaped only at the start of a line.
const lineStartSpecials = "-"

// escapeMarkdown makes member text render literally in Discord.
func escapeMarkdown(s string) string {
	var b strings.Builder
	lineStart := true
	for _, r := range s {
		if strings.ContainsRune(markdownSpecials, r) || (lineStart && strings.ContainsRune(lineStartSpecials, r)) {
			b.WriteByte('\\')
		}
		b.WriteRune(r)
		lineStart = r == '\n'
	}
	return b.String()
}

// collapse folds every run of whitespace into one space, so member text is
// one line.
func collapse(s string) string { return strings.Join(strings.Fields(s), " ") }

// escapedWithin cuts raw to rawMax runes, escapes it, and shortens it until
// the escaped text fits outMax runes, so no escape is ever cut in half.
func escapedWithin(raw string, rawMax, outMax int) string {
	runes := []rune(raw)
	if len(runes) > rawMax {
		runes = runes[:rawMax]
	}
	for {
		escaped := escapeMarkdown(string(runes))
		if len([]rune(escaped)) <= outMax || len(runes) == 0 {
			return escaped
		}
		runes = runes[:len(runes)-1]
	}
}

// thousands renders a rounded figure with comma separators.
func thousands(v float64) string {
	n := int64(math.Round(v))
	digits := fmt.Sprintf("%d", n)
	sign := ""
	if n < 0 {
		sign, digits = "-", digits[1:]
	}
	for i := len(digits) - 3; i > 0; i -= 3 {
		digits = digits[:i] + "," + digits[i:]
	}
	return sign + digits
}

// duration renders "48 s", "3 min 12 s" or "1 h 5 min".
func duration(d time.Duration) string {
	total := int(math.Round(d.Seconds()))
	if total < 0 {
		total = 0
	}
	hours, minutes, seconds := total/3600, total%3600/60, total%60
	switch {
	case hours > 0:
		return joinUnits(fmt.Sprintf("%d h", hours), minutes, "min")
	case minutes > 0:
		return joinUnits(fmt.Sprintf("%d min", minutes), seconds, "s")
	default:
		return fmt.Sprintf("%d s", seconds)
	}
}

func joinUnits(lead string, rest int, unit string) string {
	if rest == 0 {
		return lead
	}
	return fmt.Sprintf("%s %d %s", lead, rest, unit)
}

// EscapeWithin is escapedWithin for other packages that render member text
// into Discord messages.
func EscapeWithin(raw string, rawMax, outMax int) string { return escapedWithin(raw, rawMax, outMax) }

// Collapse is collapse for other packages: member text as one line.
func Collapse(s string) string { return collapse(s) }
