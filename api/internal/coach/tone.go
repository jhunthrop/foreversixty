package coach

import (
	"regexp"
	"strings"
)

// bannedTone is the words and phrases the coach never uses (§9.3). It applies in full to
// model output; in personas it applies everywhere except the Voice section's rule lines,
// which may name a rule in its own words. Later specs may add to this list, never remove.
var bannedTone = []string{"fail", "bad", "terrible", "noob", "heal more", "pump", "lazy"}

var toneRE = regexp.MustCompile(`(?i)\b(` + strings.Join(quoteAll(bannedTone), "|") + `)\b`)

// quoteAll escapes each phrase; a space matches any run of whitespace.
func quoteAll(words []string) []string {
	out := make([]string, len(words))
	for i, w := range words {
		out[i] = strings.ReplaceAll(regexp.QuoteMeta(w), " ", `\s+`)
	}
	return out
}

// toneProblems returns each banned word or phrase found in s, lower-cased, in order.
func toneProblems(s string) []string {
	var found []string
	for _, m := range toneRE.FindAllString(s, -1) {
		found = append(found, strings.ToLower(m))
	}
	return found
}
