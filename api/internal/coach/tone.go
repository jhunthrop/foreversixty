package coach

import (
	"regexp"
	"strings"
	"sync"
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

// claimRE is a claim about a guild, a ranking or a first: the coach never makes one, so
// no real player or group can be named by implication (§8.2, §9.3).
var claimRE = regexp.MustCompile(`(?i)\b(guilds?|ranked|ranking|rank\s+\d+|world\s+first|server\s+first|parses?|top\s+(?:player|parser|guild)s?)\b`)

const exclamation = "!"

// toneProblems returns each banned word or phrase, each exclamation mark and each claim
// about a guild or ranking found in s, lower-cased, banned words first.
func toneProblems(s string) []string {
	var found []string
	for _, m := range toneRE.FindAllString(s, -1) {
		found = append(found, strings.ToLower(m))
	}
	for range strings.Count(s, exclamation) {
		found = append(found, exclamation)
	}
	for _, m := range claimRE.FindAllString(s, -1) {
		found = append(found, strings.ToLower(m))
	}
	return found
}

// rosterNames returns each roster name found in s as a whole word, case-insensitively.
func rosterNames(s string, roster []string) []string {
	var found []string
	for _, name := range roster {
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}
		if rosterRE(name).MatchString(s) {
			found = append(found, name)
		}
	}
	return found
}

var rosterREs sync.Map // lower-cased name -> *regexp.Regexp

// rosterRE is the whole-word, case-insensitive matcher for one roster name, compiled once.
func rosterRE(name string) *regexp.Regexp {
	key := strings.ToLower(name)
	if re, ok := rosterREs.Load(key); ok {
		return re.(*regexp.Regexp)
	}
	re := regexp.MustCompile(`(?i)\b` + regexp.QuoteMeta(name) + `\b`)
	rosterREs.Store(key, re)
	return re
}
