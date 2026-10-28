package coach

import (
	"strconv"

	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
)

// computeLinks is the §7 connections between findings that meet in this fight: a cast gap
// overlapping a resource cap, out of mana at a death, a death inside a cast gap, and a
// missing flask on running out of mana, and unused combat potions on a death. ev is each
// finding's outcome by finding id: its windows and typed facts (outcome.Resource,
// outcome.MissingFlask, outcome.UnusedPotions), never its copy. The
// overheal link needs the fight; overhealLink makes it. Ids are left to numberLinks.
func computeLinks(fs []Finding, ev map[string]outcome) []Link {
	core, hasCore := findingOf(fs, CodeCoreCasts)
	res, hasRes := findingOf(fs, CodeResource)
	dead, hasDeath := findingOf(fs, CodeDeath)
	prep, hasPrep := findingOf(fs, CodePreparation)
	oom := hasRes && ev[res.ID].Resource.Kind == resourceOutOfMana
	var out []Link
	if hasCore && hasRes && ev[res.ID].Resource.Kind == resourceCapped {
		out = append(out, gapCapLinks(core, res, ev)...)
	}
	if oom && hasDeath && deathMeets(ev[dead.ID].Windows, ev[res.ID].Windows) {
		out = append(out, Link{Findings: [2]string{res.ID, dead.ID}, Kind: LinkOutOfManaAtDeath, Sentence: LinkOutOfManaDeath})
	}
	if hasCore && hasDeath {
		out = append(out, deathInGapLinks(dead, core, ev)...)
	}
	if oom && hasPrep && ev[prep.ID].MissingFlask && len(ev[res.ID].Windows) > 0 {
		out = append(out, Link{Findings: [2]string{prep.ID, res.ID}, Kind: LinkConsumableCovers,
			Sentence: Fill(LinkNoFlaskMana, map[string]string{"t": Clock(ev[res.ID].Windows[0].FromMS)})})
	}
	if hasDeath && hasPrep && ev[prep.ID].UnusedPotions && len(ev[dead.ID].Windows) > 0 {
		out = append(out, Link{Findings: [2]string{prep.ID, dead.ID}, Kind: LinkConsumableCovers,
			Sentence: Fill(LinkPotionDeath, map[string]string{"t": Clock(ev[dead.ID].Windows[0].FromMS)})})
	}
	return out
}

// withWaste is o with its F4 finding's kind and bar name set.
func withWaste(o outcome, kind resourceKind, name string) outcome {
	o.Resource = resourceWaste{Kind: kind, Name: name}
	return o
}

func findingOf(fs []Finding, code Code) (Finding, bool) {
	for _, f := range fs {
		if f.Code == code {
			return f, true
		}
	}
	return Finding{}, false
}

// gapCapLinks links the first gap that overlaps a cap window.
func gapCapLinks(core, res Finding, ev map[string]outcome) []Link {
	for _, gap := range ev[core.ID].Windows {
		for _, capped := range ev[res.ID].Windows {
			if gap.FromMS < capped.ToMS && capped.FromMS < gap.ToMS {
				return []Link{{Findings: [2]string{core.ID, res.ID}, Kind: LinkGapOverlapsCap,
					Sentence: Fill(LinkGapCap, map[string]string{"Spell": coreSpell(core), "gap": Span(gap.FromMS, gap.ToMS),
						"resource": ev[res.ID].Resource.Name, "cap": Span(capped.FromMS, capped.ToMS)})}}
			}
		}
	}
	return nil
}

// deathInGapLinks links a death that falls inside a cast gap.
func deathInGapLinks(dead, core Finding, ev map[string]outcome) []Link {
	death := ev[dead.ID].Windows
	for _, gap := range ev[core.ID].Windows {
		if deathMeets(death, []ratingengine.Window{gap}) {
			return []Link{{Findings: [2]string{dead.ID, core.ID}, Kind: LinkDeathInGap,
				Sentence: Fill(LinkDeathInGapSentence, map[string]string{
					"Spell": coreSpell(core), "gap": Span(gap.FromMS, gap.ToMS), "t": Clock(death[0].FromMS),
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

// coreSpell is the F1 finding's worst ability: coreFacts puts it first.
func coreSpell(f Finding) string {
	if len(f.Facts.Spells) == 0 {
		return ""
	}
	return f.Facts.Spells[0].Name
}

// overhealLink links a healer's heal-rank overheal (F5) to running out of mana (F4) when
// that heal is the player's most-cast heal in this fight.
func overhealLink(in Input, fs []Finding, ev map[string]outcome) []Link {
	rank, hasRank := findingOf(fs, CodeRanks)
	res, hasRes := findingOf(fs, CodeResource)
	if in.Row.Role != RoleHealer || !hasRank || !hasRes || ev[res.ID].Resource.Kind != resourceOutOfMana || len(rank.Facts.Spells) == 0 {
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
