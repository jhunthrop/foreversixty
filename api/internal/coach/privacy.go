package coach

import (
	"encoding/json"
	"log/slog"
	"regexp"
	"slices"
	"strings"
)

// Unicode-aware word edges: Go's \b is ASCII-only, so "Zoë" or "Ælin" would never match.
const (
	wordStart = `(?:^|[^\p{L}\p{N}])`
	wordEnd   = `(?:$|[^\p{L}\p{N}])`
)

// withheldPattern matches any of words as a whole word in any case; nil when there are none.
func withheldPattern(words []string) *regexp.Regexp {
	var quoted []string
	for _, w := range words {
		if w = strings.TrimSpace(w); w != "" {
			quoted = append(quoted, regexp.QuoteMeta(w))
		}
	}
	if len(quoted) == 0 {
		return nil
	}
	return regexp.MustCompile(`(?i)` + wordStart + `(?:` + strings.Join(quoted, "|") + `)` + wordEnd)
}

// withheld reports whether the prompt pack carries a roster name or realm (§19: no character
// name or realm leaves for the gateway). On a hit it logs the findings that carry one, by
// code only: never the name, the realm or the text around it.
func (g *Generator) withheld(p Pack) bool {
	re := withheldPattern(slices.Concat(g.Validator.Roster, g.Realms))
	if re == nil {
		return false
	}
	// A promptPack holds only strings, numbers and slices of them: Marshal cannot fail.
	whole, _ := json.Marshal(promptPackOf(p))
	if !re.Match(whole) {
		return false
	}
	g.logger().Warn("coach: the prompt pack carries a roster name or realm; serving the rule-only review",
		slog.String("pack_version", p.Version),
		slog.Any("findings", findingCodesMatching(re, p.Findings)))
	return true
}

// findingCodesMatching is the codes of the findings whose JSON re matches, in pack order.
func findingCodesMatching(re *regexp.Regexp, findings []Finding) []Code {
	codes := []Code{}
	for _, f := range findings {
		// A Finding holds only strings, numbers and slices of them: Marshal cannot fail.
		b, _ := json.Marshal(f)
		if re.Match(b) && !slices.Contains(codes, f.Code) {
			codes = append(codes, f.Code)
		}
	}
	return codes
}
