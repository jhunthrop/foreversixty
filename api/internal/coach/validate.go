package coach

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"
)

// Problem kinds. Fact kinds are FactChecker's (number, percent, time, range, name,
// number_word).
const (
	ProblemShape = "shape"
	ProblemLink  = "link"
	ProblemTone  = "tone"
)

// Problem is one validation failure, worded so it can be listed back to the model. Value
// is the exact text or a plain description; Near is the words around a fact problem.
type Problem struct {
	Kind  string
	Value string
	Near  string
}

// Validator checks a review against its pack. Names is Task 27's dictionary; Roster is the
// fight's player names, used only to catch a second player.
type Validator struct {
	Names  *NameDict
	Roster []string
}

// Validate returns every problem; none means the review may be shown.
func (v Validator) Validate(doc ReviewDoc, p Pack, persona Persona) []Problem {
	var out []Problem
	out = append(out, shapeProblems(doc, p, persona)...)
	out = append(out, linkProblems(doc, p)...)
	out = append(out, v.factProblems(doc, p, persona)...)
	out = append(out, v.toneProblems(doc)...)
	return out
}

func shape(format string, args ...any) Problem {
	return Problem{Kind: ProblemShape, Value: fmt.Sprintf(format, args...)}
}

func chars(s string) int { return utf8.RuneCountInString(s) }

func capProblem(field, s string, limit int) []Problem {
	if chars(s) > limit {
		return []Problem{shape("%s is %d characters; the limit is %d", field, chars(s), limit)}
	}
	return nil
}

func requiredProblem(field, s string) []Problem {
	if strings.TrimSpace(s) == "" {
		return []Problem{shape("%s is empty", field)}
	}
	return nil
}

func shapeProblems(doc ReviewDoc, p Pack, persona Persona) []Problem {
	var out []Problem
	out = append(out, openingProblems(doc.Opening)...)
	out = append(out, fixFirstProblems(doc.FixFirst, p, persona)...)
	out = append(out, thenFixProblems(doc, p)...)
	out = append(out, doingWellProblems(doc.DoingWell, p)...)
	if !slices.Equal(doc.CouldntCheck, p.CouldNotCheck) && (len(doc.CouldntCheck)+len(p.CouldNotCheck) > 0) {
		out = append(out, shape("couldntCheck must be the pack's lines word for word, in order"))
	}
	return out
}

func openingProblems(opening string) []Problem {
	out := append(requiredProblem("opening", opening), capProblem("opening", opening, CapOpeningChars)...)
	if n := len(sentences(opening)); n > CapOpeningSentences {
		out = append(out, shape("opening is %d sentences; the limit is %d", n, CapOpeningSentences))
	}
	return out
}

func packFinding(p Pack, id string) (Finding, bool) {
	for _, f := range p.Findings {
		if f.ID == id {
			return f, true
		}
	}
	return Finding{}, false
}

func rankOne(p Pack) (Finding, bool) {
	for _, f := range p.Findings {
		if f.Rank == 1 {
			return f, true
		}
	}
	return Finding{}, false
}

func fixFirstProblems(item *FixItem, p Pack, persona Persona) []Problem {
	top, hasTop := rankOne(p)
	switch {
	case item == nil && hasTop:
		return []Problem{shape("fixFirst is null but the pack's rank 1 is %s", top.ID)}
	case item == nil:
		return nil
	case !hasTop:
		return []Problem{shape("fixFirst must be null: the pack has no rank-1 finding")}
	}
	var out []Problem
	if item.Finding != top.ID {
		out = append(out, shape("fixFirst.finding is %q; it must be the pack's rank 1, %q", item.Finding, top.ID))
	}
	out = append(out, titleProblems("fixFirst", item.Finding, item.Title, p)...)
	out = append(out, requiredProblem("fixFirst.why", item.Why)...)
	out = append(out, capProblem("fixFirst.why", item.Why, CapFixWhyChars)...)
	out = append(out, requiredProblem("fixFirst.drillText", item.DrillText)...)
	out = append(out, capProblem("fixFirst.drillText", item.DrillText, CapDrillChars)...)
	if _, ok := persona.Drills[item.Drill]; !ok {
		out = append(out, shape("fixFirst.drill %q is not one of the persona's drills", item.Drill))
	}
	return append(out, refProblems("fixFirst", item.Anchors, item.Links, p)...)
}

func titleProblems(field, id, title string, p Pack) []Problem {
	f, ok := packFinding(p, id)
	switch {
	case !ok:
		return []Problem{shape("%s.finding %q is not in the pack", field, id)}
	case title != f.Title:
		return []Problem{shape("%s.title %q must equal the pack's title %q", field, title, f.Title)}
	}
	return nil
}

func refProblems(field string, anchors, links []string, p Pack) []Problem {
	var out []Problem
	for _, a := range anchors {
		if !slices.ContainsFunc(p.Anchors, func(x Anchor) bool { return x.ID == a }) {
			out = append(out, shape("%s cites anchor %q, which is not in the pack", field, a))
		}
	}
	for _, l := range links {
		if !slices.ContainsFunc(p.Links, func(x Link) bool { return x.ID == l }) {
			out = append(out, Problem{Kind: ProblemLink, Value: fmt.Sprintf("%s cites link %q, which is not one of the pack's links", field, l)})
		}
	}
	return out
}

func thenFixProblems(doc ReviewDoc, p Pack) []Problem {
	var out []Problem
	if len(doc.ThenFix) > CapThenItems {
		out = append(out, shape("thenFix has %d items; the limit is %d", len(doc.ThenFix), CapThenItems))
	}
	top, _ := rankOne(p)
	seen := map[string]bool{}
	for i, it := range doc.ThenFix {
		field := fmt.Sprintf("thenFix[%d]", i)
		switch {
		case it.Finding == top.ID && top.ID != "":
			out = append(out, shape("%s repeats the fix-first finding %q", field, it.Finding))
		case seen[it.Finding]:
			out = append(out, shape("%s repeats finding %q", field, it.Finding))
		}
		seen[it.Finding] = true
		out = append(out, titleProblems(field, it.Finding, it.Title, p)...)
		out = append(out, requiredProblem(field+".why", it.Why)...)
		out = append(out, capProblem(field+".why", it.Why, CapThenWhyChars)...)
		out = append(out, refProblems(field, it.Anchors, it.Links, p)...)
	}
	return out
}

func doingWellProblems(items []WellItem, p Pack) []Problem {
	var out []Problem
	switch {
	case len(p.Strengths) == 0 && len(items) > 0:
		out = append(out, shape("doingWell must be empty: the pack has no strengths"))
	case len(p.Strengths) > 0 && (len(items) == 0 || len(items) > CapWellItems):
		out = append(out, shape("doingWell has %d items; it needs one or two", len(items)))
	}
	seen := map[string]bool{}
	for i, it := range items {
		field := fmt.Sprintf("doingWell[%d]", i)
		if !slices.ContainsFunc(p.Strengths, func(s Strength) bool { return s.ID == it.Strength }) {
			out = append(out, shape("%s.strength %q is not in the pack", field, it.Strength))
		}
		if seen[it.Strength] {
			out = append(out, shape("%s repeats strength %q", field, it.Strength))
		}
		seen[it.Strength] = true
		out = append(out, requiredProblem(field+".text", it.Text)...)
		out = append(out, capProblem(field+".text", it.Text, CapWellChars)...)
		out = append(out, refProblems(field, it.Anchors, nil, p)...)
	}
	return out
}

// sentences splits s at ., ! or ? followed by white space or the end.
func sentences(s string) []string {
	var out []string
	start := 0
	runes := []rune(s)
	for i, r := range runes {
		if (r == '.' || r == '!' || r == '?') && (i+1 == len(runes) || runes[i+1] == ' ' || runes[i+1] == '\n') {
			if part := strings.TrimSpace(string(runes[start : i+1])); part != "" {
				out = append(out, part)
			}
			start = i + 1
		}
	}
	if tail := strings.TrimSpace(string(runes[start:])); tail != "" {
		out = append(out, tail)
	}
	return out
}

// proseFields is every prose string of the review, named for the retry message.
type proseField struct{ name, text string }

func proseFields(doc ReviewDoc) []proseField {
	out := []proseField{{"opening", doc.Opening}}
	if f := doc.FixFirst; f != nil {
		out = append(out, proseField{"fixFirst.why", f.Why}, proseField{"fixFirst.drillText", f.DrillText})
	}
	for i, it := range doc.ThenFix {
		out = append(out, proseField{fmt.Sprintf("thenFix[%d].why", i), it.Why})
	}
	for i, it := range doc.DoingWell {
		out = append(out, proseField{fmt.Sprintf("doingWell[%d].text", i), it.Text})
	}
	return out
}

func (v Validator) factProblems(doc ReviewDoc, p Pack, persona Persona) []Problem {
	checker := FactChecker{Allowed: p.Allowed, Class: p.Context.Class, Spec: p.Context.SpecLabel, Names: v.Names}
	if doc.FixFirst != nil {
		checker.DrillText = persona.Drills[doc.FixFirst.Drill]
	}
	var out []Problem
	for _, f := range proseFields(doc) {
		for _, fp := range checker.Check(f.text) {
			out = append(out, Problem{Kind: fp.Kind, Value: fp.Value, Near: fp.Near})
		}
	}
	return out
}

func (v Validator) toneProblems(doc ReviewDoc) []Problem {
	var out []Problem
	for _, f := range proseFields(doc) {
		for _, bad := range toneProblems(f.text) {
			out = append(out, Problem{Kind: ProblemTone, Value: fmt.Sprintf("%q in %s", bad, f.name)})
		}
		for _, name := range rosterNames(f.text, v.Roster) {
			out = append(out, Problem{Kind: ProblemTone, Value: fmt.Sprintf("another player, %q, in %s", name, f.name)})
		}
	}
	return out
}

// Link check (§9.3): a sentence that names two findings cites a link joining them.

// genericTitleWords are title words too common to identify a finding.
var genericTitleWords = stringSet(strings.Fields(
	"casts cast lost with from that this what full ready seconds fight pull rank died time down left over while your"))

var titleWordRE = regexp.MustCompile(`[A-Za-z']+`)

// subjectTerms is the lower-case words and phrases that identify a finding in prose: its
// title's distinctive words, its spells and its items.
func subjectTerms(f Finding) []string {
	var terms []string
	for _, w := range titleWordRE.FindAllString(strings.ToLower(f.Title), -1) {
		if len(w) >= 4 && !genericTitleWords[w] {
			terms = append(terms, w)
		}
	}
	for _, s := range f.Facts.Spells {
		terms = append(terms, strings.ToLower(s.Name))
	}
	for _, it := range f.Facts.Items {
		terms = append(terms, strings.ToLower(it))
	}
	return terms
}

func mentions(sentence string, terms []string) bool {
	low := strings.ToLower(sentence)
	return slices.ContainsFunc(terms, func(t string) bool {
		return regexp.MustCompile(`\b` + regexp.QuoteMeta(t) + `\b`).MatchString(low)
	})
}

// linkedPairs is the set of finding pairs the cited links join.
func linkedPairs(p Pack, cited []string) map[[2]string]bool {
	pairs := map[[2]string]bool{}
	for _, l := range p.Links {
		if slices.Contains(cited, l.ID) {
			pairs[[2]string{l.Findings[0], l.Findings[1]}] = true
			pairs[[2]string{l.Findings[1], l.Findings[0]}] = true
		}
	}
	return pairs
}

// itemLinkProblems finds each other finding the item's sentences name without a link.
func itemLinkProblems(field, own, why string, cited []string, p Pack) []Problem {
	ownFinding, ok := packFinding(p, own)
	if !ok {
		return nil
	}
	ownTerms := subjectTerms(ownFinding)
	pairs := linkedPairs(p, cited)
	var out []Problem
	reported := map[string]bool{}
	for _, s := range sentences(why) {
		for _, g := range p.Findings {
			if g.ID == own || reported[g.ID] || pairs[[2]string{own, g.ID}] {
				continue
			}
			foreign := slices.DeleteFunc(subjectTerms(g), func(t string) bool { return slices.Contains(ownTerms, t) })
			if mentions(s, foreign) {
				reported[g.ID] = true
				out = append(out, Problem{Kind: ProblemLink, Value: fmt.Sprintf(
					"%s relates %q to %q (%s) without citing a link that joins them", field, ownFinding.Title, g.Title, g.ID)})
			}
		}
	}
	return out
}

func linkProblems(doc ReviewDoc, p Pack) []Problem {
	var out []Problem
	if f := doc.FixFirst; f != nil {
		out = append(out, itemLinkProblems("fixFirst", f.Finding, f.Why, f.Links, p)...)
	}
	for i, it := range doc.ThenFix {
		out = append(out, itemLinkProblems(fmt.Sprintf("thenFix[%d]", i), it.Finding, it.Why, it.Links, p)...)
	}
	return out
}
