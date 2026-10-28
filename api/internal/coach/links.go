package coach

import (
	"regexp"
	"strconv"
	"strings"
	"sync"

	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
)

// computeLinks is the §7 connections between findings that meet in this fight: a cast gap
// overlapping a resource cap, out of mana at a death, a death inside a cast gap, and a
// missing flask on running out of mana. The overheal link needs the fight; overhealLink
// makes it. Ids are left to numberLinks.
func computeLinks(fs []Finding, windows map[string][]ratingengine.Window) []Link {
	core, hasCore := findingOf(fs, CodeCoreCasts)
	res, hasRes := findingOf(fs, CodeResource)
	dead, hasDeath := findingOf(fs, CodeDeath)
	prep, hasPrep := findingOf(fs, CodePreparation)
	var out []Link
	if hasCore && hasRes {
		out = append(out, gapCapLinks(core, res, windows)...)
	}
	if hasRes && hasDeath && outOfMana(res) && deathMeets(windows[dead.ID], windows[res.ID]) {
		out = append(out, Link{Findings: [2]string{res.ID, dead.ID}, Kind: LinkOutOfManaAtDeath, Sentence: LinkOutOfManaDeath})
	}
	if hasCore && hasDeath {
		out = append(out, deathInGapLinks(dead, core, windows)...)
	}
	if t, oom := outOfManaFrom(res); hasPrep && hasRes && prep.Title == TitleNoFlask && oom {
		out = append(out, Link{Findings: [2]string{prep.ID, res.ID}, Kind: LinkConsumableCovers,
			Sentence: Fill(LinkNoFlaskMana, map[string]string{"t": t})})
	}
	return out
}

func findingOf(fs []Finding, code Code) (Finding, bool) {
	for _, f := range fs {
		if f.Code == code {
			return f, true
		}
	}
	return Finding{}, false
}

// gapCapLinks links the first gap that overlaps a cap window, when the resource finding
// is a cap.
func gapCapLinks(core, res Finding, windows map[string][]ratingengine.Window) []Link {
	bar, ok := titleField(TitleCapped, "Resource", res.Title)
	if !ok {
		return nil
	}
	for _, gap := range windows[core.ID] {
		for _, capped := range windows[res.ID] {
			if gap.FromMS < capped.ToMS && capped.FromMS < gap.ToMS {
				return []Link{{Findings: [2]string{core.ID, res.ID}, Kind: LinkGapOverlapsCap,
					Sentence: Fill(LinkGapCap, map[string]string{"Spell": coreSpell(core), "gap": Span(gap.FromMS, gap.ToMS),
						"resource": strings.ToLower(bar), "cap": Span(capped.FromMS, capped.ToMS)})}}
			}
		}
	}
	return nil
}

// deathInGapLinks links a death that falls inside a cast gap.
func deathInGapLinks(dead, core Finding, windows map[string][]ratingengine.Window) []Link {
	for _, gap := range windows[core.ID] {
		if deathMeets(windows[dead.ID], []ratingengine.Window{gap}) {
			return []Link{{Findings: [2]string{dead.ID, core.ID}, Kind: LinkDeathInGap,
				Sentence: Fill(LinkDeathInGapSentence, map[string]string{
					"Spell": coreSpell(core), "gap": Span(gap.FromMS, gap.ToMS), "t": Clock(windows[dead.ID][0].FromMS),
				})}}
		}
	}
	return nil
}

// deathMeets reports the death instant (the death finding's one window) inside any of ws,
// ends included.
func deathMeets(death, ws []ratingengine.Window) bool {
	if len(death) == 0 {
		return false
	}
	at := death[0].FromMS
	for _, w := range ws {
		if w.FromMS <= at && at <= w.ToMS {
			return true
		}
	}
	return false
}

func outOfMana(f Finding) bool {
	_, ok := outOfManaFrom(f)
	return ok
}

// outOfManaFrom is the time an out-of-mana finding's title gives.
func outOfManaFrom(f Finding) (string, bool) {
	return titleField(TitleOutOfMana, "t", f.Title)
}

// coreSpell is the F1 finding's worst ability: its first spell fact, else its title's.
func coreSpell(f Finding) string {
	if len(f.Facts.Spells) > 0 {
		return f.Facts.Spells[0].Name
	}
	name, _ := titleField(TitleCoreCasts, "Spell", f.Title)
	return name
}

// overhealLink links a heal rank's overheal (F5) to running out of mana (F4) when that
// heal is the player's most-cast heal in this fight.
func overhealLink(in Input, fs []Finding) []Link {
	rank, hasRank := findingOf(fs, CodeRanks)
	res, hasRes := findingOf(fs, CodeResource)
	if !hasRank || !hasRes || !outOfMana(res) || len(rank.Facts.Spells) == 0 {
		return nil
	}
	if _, ok := titleField(TitleHealRank, "Spell", rank.Title); !ok {
		return nil
	}
	heal := rank.Facts.Spells[0]
	if mostCastHeal(in) != heal.ID {
		return nil
	}
	return []Link{{Findings: [2]string{rank.ID, res.ID}, Kind: LinkOverhealOnMostCast,
		Sentence: Fill(LinkOverhealMostCast, map[string]string{"Spell": heal.Name, "r": strconv.Itoa(heal.Rank)})}}
}

// mostCastHeal is the spell id of the player's own heal with the most successful casts.
func mostCastHeal(in Input) int64 {
	var best int64
	most := 0
	for _, a := range playerHealing(in) {
		if n := succeededCasts(in, a.SpellID); a.Via == "" && n > most {
			best, most = a.SpellID, n
		}
	}
	return best
}

// numberLinks gives the links ids l1… in order.
func numberLinks(ls []Link) []Link {
	out := make([]Link, len(ls))
	for i, l := range ls {
		l.ID = "l" + strconv.Itoa(i+1)
		out[i] = l
	}
	return out
}

var (
	titlePatternsMu sync.Mutex
	titlePatterns   = map[string]*regexp.Regexp{}
)

// titleField reads field out of title when title has format's shape (a copy template
// with {name} placeholders).
func titleField(format, field, title string) (string, bool) {
	names := placeholder.FindAllStringSubmatch(format, -1)
	m := titlePattern(format).FindStringSubmatch(title)
	if m == nil {
		return "", false
	}
	for i, n := range names {
		if n[1] == field {
			return m[i+1], true
		}
	}
	return "", false
}

func titlePattern(format string) *regexp.Regexp {
	titlePatternsMu.Lock()
	defer titlePatternsMu.Unlock()
	if re, ok := titlePatterns[format]; ok {
		return re
	}
	parts := placeholder.Split(format, -1)
	var b strings.Builder
	b.WriteString("^")
	for i, part := range parts {
		b.WriteString(regexp.QuoteMeta(part))
		if i < len(parts)-1 {
			b.WriteString("(.+?)")
		}
	}
	b.WriteString("$")
	re := regexp.MustCompile(b.String())
	titlePatterns[format] = re
	return re
}
