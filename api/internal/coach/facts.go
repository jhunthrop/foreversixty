package coach

import "strings"

// Fact problem kinds (§10.2–§10.4). A percentage is reported as KindPercent, a time as
// KindTime and a time range as KindRange; the rest are below.
const (
	ProblemName       = "name"
	ProblemNumberWord = "number_word"
)

// nearWindow is how many characters either side of a problem the retry message quotes.
const nearWindow = 18

// FactProblem is one number, time, range, number word or name the pack does not allow.
// Value is the exact text found; Near is the words around it for the retry message.
type FactProblem struct {
	Kind, Value, Near string
}

// FactChecker checks prose against the pack's allowed set. DrillText is the chosen
// drill's own text: what it states, the review may repeat.
type FactChecker struct {
	Allowed   Allowed
	Class     string
	Spec      string
	DrillText string
	Names     *NameDict
}

// Check returns every problem in the prose fields (opening, why, drill text, text), in
// field and reading order. None means every number, time, range and name is the pack's.
func (c FactChecker) Check(prose ...string) []FactProblem {
	numbers := c.allowedNumbers()
	names := c.allowedNames()
	drillWords := stringSet(NumberWords(c.DrillText))
	var out []FactProblem
	for _, s := range prose {
		out = append(out, numberProblems(s, numbers)...)
		for _, w := range NumberWords(s) {
			if !drillWords[w] {
				out = append(out, FactProblem{Kind: ProblemNumberWord, Value: w, Near: nearText(s, w)})
			}
		}
		out = append(out, c.nameProblems(s, names)...)
	}
	return out
}

func numberProblems(s string, allowed map[string]bool) []FactProblem {
	var out []FactProblem
	for _, t := range ExtractNumbers(s) {
		if !allowed[t.Norm] {
			out = append(out, FactProblem{Kind: t.Kind, Value: t.Raw, Near: nearText(s, t.Raw)})
		}
	}
	return out
}

func (c FactChecker) nameProblems(s string, allowed map[string]bool) []FactProblem {
	if c.Names == nil {
		return nil
	}
	var out []FactProblem
	for _, n := range c.Names.Find(s) {
		if !allowed[strings.ToLower(n)] {
			out = append(out, FactProblem{Kind: ProblemName, Value: n, Near: nearText(s, n)})
		}
	}
	return out
}

// allowedNumbers is the pack's numbers, times and ranges plus the drill's own, as
// normalised tokens; a range's two times are allowed on their own as well.
func (c FactChecker) allowedNumbers() map[string]bool {
	set := map[string]bool{}
	for _, entry := range append(append([]string{}, c.Allowed.Numbers...), c.DrillText) {
		for _, t := range ExtractNumbers(entry) {
			set[t.Norm] = true
			if t.Kind == KindRange {
				from, to, _ := strings.Cut(t.Norm, enDash)
				set[from], set[to] = true, true
			}
		}
	}
	return set
}

// allowedNames is the pack's names, the player's class and spec and the drill's names.
func (c FactChecker) allowedNames() map[string]bool {
	set := map[string]bool{}
	for _, n := range append(append([]string{}, c.Allowed.Names...), c.Class, c.Spec) {
		set[strings.ToLower(strings.TrimSpace(n))] = true
	}
	if c.Names != nil {
		for _, n := range c.Names.Find(c.DrillText) {
			set[strings.ToLower(n)] = true
		}
	}
	return set
}

func stringSet(items []string) map[string]bool {
	set := make(map[string]bool, len(items))
	for _, i := range items {
		set[i] = true
	}
	return set
}

// nearText is value with a few words of its surroundings, cut at word boundaries.
func nearText(s, value string) string {
	idx := strings.Index(strings.ToLower(s), strings.ToLower(value))
	if idx < 0 {
		return value
	}
	start, end := max(0, idx-nearWindow), min(len(s), idx+len(value)+nearWindow)
	window := s[start:end]
	if start > 0 {
		window = dropFirstWord(window)
	}
	if end < len(s) {
		window = dropLastWord(window)
	}
	return strings.TrimSpace(window)
}

func dropFirstWord(s string) string {
	if i := strings.Index(s, " "); i >= 0 {
		return s[i+1:]
	}
	return s
}

func dropLastWord(s string) string {
	if i := strings.LastIndex(s, " "); i >= 0 {
		return s[:i]
	}
	return s
}
