package coach

import (
	"cmp"
	"regexp"
	"slices"
	"strings"
)

// NightFight is one boss fight of a character's night: its index in the report and its pack.
type NightFight struct {
	FightIndex int
	Pack       Pack
}

// NightLine is the night rule of §5.5, from packs only: the rank-1 findings across one
// character's fights, chosen by severity, then the finding type on the most bosses, then the
// earliest boss. On of Of bosses had the same code and Stem as their rank-1 finding; Of counts
// the ready packs. FindingCount sums every ready pack's findings.
type NightLine struct {
	Code         Code
	Stem         string
	On, Of       int
	FightIndex   int
	FindingCount int
}

// nightGroup is one (code, stem) across the night.
type nightGroup struct {
	code     Code
	stem     string
	severity Severity
	on       int
	earliest int
}

// Night is the night line of fights; false when no ready pack has a finding.
func Night(fights []NightFight) (NightLine, bool) {
	var groups []*nightGroup
	of, count := 0, 0
	for _, f := range fights {
		if f.Pack.Status != StatusReady {
			continue
		}
		of++
		count += len(f.Pack.Findings)
		if first, ok := nightRankOne(f.Pack); ok {
			groups = addToGroup(groups, first, f.FightIndex)
		}
	}
	if len(groups) == 0 {
		return NightLine{}, false
	}
	slices.SortStableFunc(groups, compareGroups)
	g := groups[0]
	return NightLine{Code: g.code, Stem: g.stem, On: g.on, Of: of, FightIndex: g.earliest, FindingCount: count}, true
}

// nightRankOne is the pack's rank-1 finding.
func nightRankOne(p Pack) (Finding, bool) {
	for _, f := range p.Findings {
		if f.Rank == 1 {
			return f, true
		}
	}
	return Finding{}, false
}

func addToGroup(groups []*nightGroup, f Finding, fightIndex int) []*nightGroup {
	stem := Stem(f)
	for _, g := range groups {
		if g.code == f.Code && g.stem == stem {
			g.on++
			g.earliest = min(g.earliest, fightIndex)
			if severityOrder(f.Severity) < severityOrder(g.severity) {
				g.severity = f.Severity
			}
			return groups
		}
	}
	return append(groups, &nightGroup{code: f.Code, stem: stem, severity: f.Severity, on: 1, earliest: fightIndex})
}

// compareGroups orders by severity, then the most bosses, then the earliest boss.
func compareGroups(a, b *nightGroup) int {
	return cmp.Or(
		cmp.Compare(severityOrder(a.severity), severityOrder(b.severity)),
		cmp.Compare(b.on, a.on),
		cmp.Compare(a.earliest, b.earliest),
	)
}

func severityOrder(s Severity) int {
	if s == SeverityMajor {
		return 0
	}
	return 1
}

// stemRule turns a finding's title into its night stem: the title template's fields fill the
// stem template. An empty title matches any title of the code.
type stemRule struct {
	code  Code
	title string
	stem  string
}

// stemRules are the §8.1 night stems by finding code and title form; the first match wins.
var stemRules = []stemRule{
	{CodeCoreCasts, TitleCoreCasts, NightCoreCasts},
	{CodeUptime, TitleUptime, StemUptime},
	{CodeIdle, "", StemIdle},
	{CodeResource, TitleOutOfMana, StemOutOfMana},
	{CodeResource, TitleCapped, StemCapped},
	{CodeRanks, TitleHealRank, StemHealRank},
	{CodeRanks, TitleLowRank, StemLowRank},
	{CodePreparation, TitleNoFlask, StemNoFlask},
	{CodePreparation, TitleNoFood, StemNoFood},
	{CodePreparation, TitleNoWeaponOil, StemEnchants},
	{CodePreparation, TitleNoWeaponEnchant, StemEnchants},
	{CodePreparation, TitleNoEnchants, StemEnchants},
	{CodePreparation, TitlePotions, StemPotions},
	{CodePreparation, TitleRaidBuffs, StemRaidBuffs},
	{CodePreparation, TitleRaidBuff, StemRaidBuffs},
	{CodeDeath, "", StemDeath},
	{CodeAdherence, "", StemAdherence},
}

// titleField is one {Field} of a title template.
var titleField = regexp.MustCompile(`\{([A-Za-z]+)\}`)

// stemPatterns is each rule's title template as a pattern, compiled once.
var stemPatterns = func() map[string]*regexp.Regexp {
	out := map[string]*regexp.Regexp{}
	for _, r := range stemRules {
		if r.title != "" {
			out[r.title] = titlePattern(r.title)
		}
	}
	return out
}()

// Stem is the title stem without the per-fight number ("Bloodthirst casts lost"), so the
// same stem on two bosses means the same words. A title no rule knows is its own stem.
func Stem(f Finding) string {
	for _, r := range stemRules {
		if r.code != f.Code {
			continue
		}
		if r.title == "" {
			return r.stem
		}
		if fields, ok := matchTitle(stemPatterns[r.title], f.Title); ok {
			return Fill(r.stem, fields)
		}
	}
	return f.Title
}

// matchTitle reads title's fields against re; false when title is not of its form.
func matchTitle(re *regexp.Regexp, title string) (map[string]string, bool) {
	m := re.FindStringSubmatch(title)
	if m == nil {
		return nil, false
	}
	fields := map[string]string{}
	for i, name := range re.SubexpNames() {
		if name != "" {
			fields[name] = m[i]
		}
	}
	return fields, true
}

// titlePattern is template as an anchored pattern, each {Field} a named lazy group.
func titlePattern(template string) *regexp.Regexp {
	var b strings.Builder
	b.WriteString("^")
	last := 0
	for _, loc := range titleField.FindAllStringSubmatchIndex(template, -1) {
		b.WriteString(regexp.QuoteMeta(template[last:loc[0]]))
		b.WriteString("(?P<" + template[loc[2]:loc[3]] + ">.+?)")
		last = loc[1]
	}
	b.WriteString(regexp.QuoteMeta(template[last:]) + "$")
	return regexp.MustCompile(b.String())
}
