package coach

import (
	"slices"
	"sort"
	"strconv"
	"strings"

	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
)

// needsRuleFile are the checks a spec's rule file drives (core abilities, owned auras,
// and the rotation check that ships with them).
var needsRuleFile = []Code{CodeCoreCasts, CodeUptime, CodeAdherence}

// reviewCheck is one of the eight checks and the code its finding carries.
type reviewCheck struct {
	code Code
	run  check
}

// reviewChecks is F1–F8 in check order: the order the couldn't-check lines follow.
var reviewChecks = []reviewCheck{
	{CodeCoreCasts, coreCasts}, {CodeUptime, uptime}, {CodeIdle, idle}, {CodeResource, resource},
	{CodeRanks, ranks}, {CodePreparation, preparation}, {CodeDeath, death}, {CodeAdherence, adherence},
}

// checked is one check's outcome with its code.
type checked struct {
	code Code
	out  outcome
}

// Review builds one player's evidence pack for one fight (§7). Pure. Version is left
// for the caller, which holds the rule file and catalogue bytes PackVersion hashes.
func Review(in Input) Pack {
	p := Pack{Status: StatusReady, Context: contextOf(in), Findings: []Finding{}, Strengths: []Strength{},
		Links: []Link{}, CouldNotCheck: []string{}, Anchors: []Anchor{}, BasisLabels: []string{}}
	if reason := notReviewable(in); reason != "" {
		p.Status, p.NotReviewable = StatusNotReviewable, reason
		return p
	}
	if in.Refs.SimPreparing && in.Refs.Pool == nil {
		p.Status = StatusPreparing
		return p
	}
	results := runChecks(in)
	book := &anchorBook{results: results}
	findings, from := rankFindings(in.Row.Role, results)
	p.Findings = numberFindings(findings, from, book)
	p.Strengths = collectStrengths(in.Row.Role, results, book)
	p.Anchors = book.anchors
	p.Links = numberLinks(append(computeLinks(p.Findings, findingWindows(p.Findings, from, results)), overhealLink(in, p.Findings)...))
	p.CouldNotCheck = couldNotCheck(results)
	p.BasisLabels = basisLabels(p.Findings)
	p.Allowed = allowedOf(p)
	return p
}

// contextOf is the fight as the model may see it: never a name, account, guild or realm.
func contextOf(in Input) Context {
	return Context{
		SpecLabel: in.Row.Spec, Class: in.Row.Class, Role: in.Row.Role, Boss: in.BossName, Size: in.Size,
		Kill: in.Fight.Kill, Duration: Clock(in.Fight.DurationMS), Level: int(in.Row.Level),
	}
}

// notReviewable is the reason the fight is not reviewed (§6.1), "" when it is.
func notReviewable(in Input) string {
	switch {
	case in.Row.Spec == "":
		return NotReviewableSpec
	case in.Fight.EncounterID == 0:
		return NotReviewableTrash
	case in.Fight.DurationMS < MinFightMS:
		return NotReviewableShort
	}
	return ""
}

// runChecks runs the checks the role gets (Precedence[role]), in check order. Without a
// rule file for the spec, the checks that read it are skipped and say so, never "checked,
// nothing wrong".
func runChecks(in Input) []checked {
	c := CountedTime(in)
	noRules := in.Tables.Rules.Spec == ""
	var out []checked
	for _, rc := range reviewChecks {
		switch {
		case !slices.Contains(Precedence[in.Row.Role], rc.code):
		case noRules && slices.Contains(needsRuleFile, rc.code):
			out = append(out, checked{code: rc.code, out: outcome{CouldNotCheck: []string{
				Fill(CannotSpecRules, map[string]string{"Spec": in.Row.Spec})}}})
		default:
			out = append(out, checked{code: rc.code, out: rc.run(in, c)})
		}
	}
	return out
}

// rankFindings is the fired findings, ranked: a raid-buff-only preparation finding last,
// then major before minor, then the role's precedence; capped at MaxFindings. from[i] is
// the result index findings[i] came from.
func rankFindings(role string, results []checked) (findings []Finding, from []int) {
	for i, r := range results {
		if r.out.Finding != nil {
			findings, from = append(findings, *r.out.Finding), append(from, i)
		}
	}
	order := make([]int, len(findings))
	for i := range order {
		order[i] = i
	}
	sort.SliceStable(order, func(a, b int) bool { return rankKey(role, findings[order[a]]) < rankKey(role, findings[order[b]]) })
	n := min(len(order), MaxFindings)
	sortedF, sortedFrom := make([]Finding, n), make([]int, n)
	for i, idx := range order[:n] {
		sortedF[i], sortedFrom[i] = findings[idx], from[idx]
	}
	return sortedF, sortedFrom
}

// rankKey orders findings: raid buffs alone last, then severity, then precedence.
func rankKey(role string, f Finding) int {
	key := slices.Index(Precedence[role], f.Code)
	if f.Severity != SeverityMajor {
		key += len(reviewChecks)
	}
	if raidBuffAlone(f) {
		key += 2 * len(reviewChecks)
	}
	return key
}

// raidBuffAlone reports a preparation finding about raid buffs alone. F6 words its finding
// from its worst issue, raid buffs (BasisRaidBuff) only when nothing ranks above them,
// and orders them after every other issue but the potions, which add numbers: so a
// raid-buff basis with no number is raid buffs and nothing else.
func raidBuffAlone(f Finding) bool {
	return f.Code == CodePreparation && f.Basis.Label == BasisRaidBuff && len(f.Facts.Numbers) == 0
}

// numberFindings gives ids f1… and ranks in rank order, and global anchor ids. Raid
// buffs are never rank 1: when one leads (it is then alone), ranks start at 2.
func numberFindings(fs []Finding, from []int, book *anchorBook) []Finding {
	first := 1
	if len(fs) > 0 && raidBuffAlone(fs[0]) {
		first = 2
	}
	out := make([]Finding, len(fs))
	for i, f := range fs {
		f.ID, f.Rank = "f"+strconv.Itoa(i+1), first+i
		f.Anchors = book.refs(from[i], f.Anchors)
		out[i] = f
	}
	return out
}

// collectStrengths is every check's strengths, ordered by the role's precedence of the
// check that measured them, with ids s1… and global anchor ids.
func collectStrengths(role string, results []checked, book *anchorBook) []Strength {
	order := make([]int, len(results))
	for i := range order {
		order[i] = i
	}
	sort.SliceStable(order, func(a, b int) bool {
		return slices.Index(Precedence[role], results[order[a]].code) < slices.Index(Precedence[role], results[order[b]].code)
	})
	out := []Strength{}
	for _, idx := range order {
		for _, s := range results[idx].out.Strengths {
			s.ID = "s" + strconv.Itoa(len(out)+1)
			s.Anchors = book.refs(idx, s.Anchors)
			out = append(out, s)
		}
	}
	return out
}

// findingWindows is each ranked finding's windows, by finding id, for the links.
func findingWindows(fs []Finding, from []int, results []checked) map[string][]ratingengine.Window {
	out := make(map[string][]ratingengine.Window, len(fs))
	for i, f := range fs {
		out[f.ID] = results[from[i]].out.Windows
	}
	return out
}

// couldNotCheck is the talents and range lines, then each check's lines in check order,
// one copy of each.
func couldNotCheck(results []checked) []string {
	out := []string{CannotTalents, CannotRange}
	for _, r := range results {
		for _, line := range r.out.CouldNotCheck {
			if !slices.Contains(out, line) {
				out = append(out, line)
			}
		}
	}
	return out
}

// basisLabels is every basis and comparison label the findings use, one copy each.
func basisLabels(fs []Finding) []string {
	out := []string{}
	add := func(label string) {
		if label != "" && !slices.Contains(out, label) {
			out = append(out, label)
		}
	}
	for _, f := range fs {
		add(f.Basis.Label)
		if f.Comparison != nil {
			add(f.Comparison.Label)
		}
	}
	return out
}

// anchorBook numbers the outcomes' local anchors a1… in the order they are first cited.
type anchorBook struct {
	results []checked
	anchors []Anchor
	ids     map[[2]int]string
}

// refs maps local "#n" references of result idx to global ids; a reference to no anchor
// is dropped.
func (b *anchorBook) refs(idx int, local []string) []string {
	out := []string{}
	for _, ref := range local {
		if id, ok := b.ref(idx, ref); ok {
			out = append(out, id)
		}
	}
	return out
}

// ref is the global id of result idx's local anchor ref, numbered on first citation.
func (b *anchorBook) ref(idx int, ref string) (string, bool) {
	n, err := strconv.Atoi(strings.TrimPrefix(ref, "#"))
	local := b.results[idx].out.Anchors
	if err != nil || !strings.HasPrefix(ref, "#") || n < 0 || n >= len(local) {
		return "", false
	}
	key := [2]int{idx, n}
	if id, ok := b.ids[key]; ok {
		return id, true
	}
	a := local[n]
	a.ID = "a" + strconv.Itoa(len(b.anchors)+1)
	// F7's death anchor string embeds the player's GUID (realm and character); the pack
	// carries none (§7). The page rebuilds it from the player it shows and at_ms.
	a.Death = ""
	if b.ids == nil {
		b.ids = map[[2]int]string{}
	}
	b.ids[key] = a.ID
	b.anchors = append(b.anchors, a)
	return a.ID, true
}
