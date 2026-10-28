package coach

import (
	"slices"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// deathAnchorPrefix is the rating's death Moment anchor, "death-<guid>-<at_ms>", which
// DeathsTab.svelte renders as the death row's id.
const deathAnchorPrefix = "death-"

// The names a healing potion's or a Healthstone's use is cast under, across ranks and
// variants ("Healing Potion", "Superior Healing Potion", "Major Healthstone"). Conjuring
// a stone ("Create Healthstone") and the warlock talent are not uses.
const (
	deathPotionSuffix      = "Healing Potion"
	deathStoneSuffix       = "Healthstone"
	deathStoneNotUsePrefix = "Create "
	deathStoneTalent       = "Improved Healthstone"
)

const (
	deathFactAmount = "killing blow damage"
	deathFactUnit   = "damage"
)

// deathHit is what killed the player: the killing blow, else the last hit the log kept.
type deathHit struct {
	SpellID int64
	Name    string // a melee swing reads as Drafts["melee"]
	Amount  int64
	AtMS    int64
}

// deathReady is a defensive cooldown off cooldown at the death.
type deathReady struct {
	Name      string
	SpellID   int64
	SincePull bool // never cast before the death: ready from the pull
}

// death is F7 (§6.2): the player's first reviewable death, what killed them, whether it
// is on the encounter's avoidable list, a defensive that was ready, and whether a healing
// potion or Healthstone was cast in the DeathPotionWindowMS before. No death at all is a
// strength. A death the log cannot name is not reviewed.
func death(in Input, _ Counted) outcome {
	d, died, reviewable := firstReviewableDeath(in)
	if !died {
		return outcome{Strengths: []Strength{{Sentence: StrengthAlive}}}
	}
	if !reviewable || !deathFires(in.Fight, d.AtMS) {
		return outcome{}
	}
	hit, ok := deathHitOf(in, d)
	if !ok {
		return outcome{}
	}
	ready, hasReady := readyDefensive(in, d.AtMS)
	anchors := []Anchor{deathAnchor(in.Row.GUID, d.AtMS)}
	if hasReady {
		anchors = append(anchors, Anchor{Kind: AnchorCast, SpellID: ready.SpellID})
	}
	return outcome{
		Finding: deathFinding(in, d.AtMS, hit, ready, hasReady, len(anchors)),
		Anchors: anchors,
		Windows: []ratingengine.Window{{FromMS: d.AtMS, ToMS: d.AtMS}},
	}
}

// firstReviewableDeath is the player's earliest reviewable death. died is false when the
// player never died; reviewable is false when every death is part of the wipe.
func firstReviewableDeath(in Input) (first summary.Death, died, reviewable bool) {
	for _, d := range in.Fight.Deaths {
		if d.GUID != in.Row.GUID {
			continue
		}
		died = true
		if ReviewableDeath(in, d) && (!reviewable || d.AtMS < first.AtMS) {
			first, reviewable = d, true
		}
	}
	return first, died, reviewable
}

// deathFires: on a kill, a death before the last (100 − DeathFiresBeforePct)% of the
// fight; on a wipe, any reviewable death.
func deathFires(fight summary.Summary, atMS int64) bool {
	return !fight.Kill || deathAtPct(fight, atMS) < DeathFiresBeforePct
}

func deathAtPct(fight summary.Summary, atMS int64) float64 {
	if fight.DurationMS <= 0 {
		return percentOf
	}
	return float64(atMS) / float64(fight.DurationMS) * percentOf
}

// deathHitOf is the killing blow, else the latest hit in the death's recap. A melee swing
// is named by its source (the boss when the log names none). false: nothing names it.
func deathHitOf(in Input, d summary.Death) (deathHit, bool) {
	ref := d.KillingBlow
	if ref == nil {
		ref = latestHit(d.Last)
	}
	if ref == nil {
		return deathHit{}, false
	}
	name := ref.SpellName
	if ref.SpellID == 0 {
		source := ref.SourceName
		if source == "" {
			source = in.BossName
		}
		name = ""
		if source != "" {
			name = Fill(Drafts["melee"], map[string]string{"Boss": source})
		}
	}
	return deathHit{SpellID: ref.SpellID, Name: name, Amount: ref.Amount, AtMS: ref.AtMS}, name != ""
}

func latestHit(hits []summary.DamageRef) *summary.DamageRef {
	var latest *summary.DamageRef
	for i := range hits {
		if latest == nil || hits[i].AtMS >= latest.AtMS {
			latest = &hits[i]
		}
	}
	return latest
}

func deathAnchor(guid string, atMS int64) Anchor {
	at := atMS
	return Anchor{Kind: AnchorDeath, AtMS: &at, Death: deathAnchorPrefix + guid + "-" + strconv.FormatInt(atMS, 10)}
}

// deathAvoidable: the hit is on the encounter's avoidable list, and not an ability the
// table says the player's own role is meant to take (Flame Breath is the tank's).
func deathAvoidable(in Input, hit deathHit) bool {
	if hit.SpellID == 0 || in.Tables.Mechanics == nil {
		return false
	}
	m, ok := in.Tables.Mechanics.Lookup(hit.SpellID)
	return ok && m.Kind == mechanics.Avoidable && (m.Role == "" || m.Role != in.Row.Role)
}

// potionBefore reports a healing potion or Healthstone cast by the player in the
// DeathPotionWindowMS up to the death: a fact about casts, never about bags.
func potionBefore(fight summary.Summary, guid string, atMS int64) bool {
	for _, c := range fight.Casts {
		if c.GUID != guid || !isHealingConsumable(c.SpellName) {
			continue
		}
		for _, at := range c.Sequence {
			if at >= atMS-DeathPotionWindowMS && at <= atMS {
				return true
			}
		}
	}
	return false
}

func isHealingConsumable(name string) bool {
	if strings.HasSuffix(name, deathPotionSuffix) {
		return true
	}
	return strings.HasSuffix(name, deathStoneSuffix) && !strings.HasPrefix(name, deathStoneNotUsePrefix) && name != deathStoneTalent
}

// readyDefensive is the first of the rule file's defensives off cooldown at atMS. A talent
// counts only if cast that fight; a baseline ability never cast counts from its trainable
// level, and only when the level is known. A cast whose cooldown the data does not state
// is never claimed ready.
func readyDefensive(in Input, atMS int64) (deathReady, bool) {
	for _, e := range in.Tables.Rules.Defensives {
		lastID, lastMS, cast := lastCastBefore(in.Fight, in.Row.GUID, e.SpellIDs, atMS)
		if !cast {
			if id, ok := trainableAtLevel(in, e); ok && !e.Talent {
				return deathReady{Name: e.Family, SpellID: id, SincePull: true}, true
			}
			continue
		}
		sc, ok := in.Tables.Spells[lastID]
		if ok && sc.CooldownMS > 0 && lastMS+sc.CooldownMS <= atMS {
			return deathReady{Name: e.Family, SpellID: lastID}, true
		}
	}
	return deathReady{}, false
}

// lastCastBefore is guid's latest successful cast of any of ids before atMS.
func lastCastBefore(fight summary.Summary, guid string, ids []int64, atMS int64) (id, lastMS int64, ok bool) {
	for _, c := range fight.Casts {
		if c.GUID != guid || !slices.Contains(ids, c.SpellID) {
			continue
		}
		for _, at := range c.Sequence {
			if at < atMS && (!ok || at > lastMS) {
				id, lastMS, ok = c.SpellID, at, true
			}
		}
	}
	return id, lastMS, ok
}

// trainableAtLevel is the highest rank of e the player's level can train, by id.
func trainableAtLevel(in Input, e RuleEntry) (int64, bool) {
	level := in.Row.Level
	if level <= 0 {
		return 0, false
	}
	var best TrainableRank
	found := false
	for _, r := range in.Tables.Trainable[e.Family] {
		if int64(r.Level) <= level && slices.Contains(e.SpellIDs, r.ID) && (!found || r.Level >= best.Level) {
			best, found = r, true
		}
	}
	return best.ID, found
}

// deathFinding is F7's finding; nAnchors local anchors, the death first.
func deathFinding(in Input, atMS int64, hit deathHit, ready deathReady, hasReady bool, nAnchors int) *Finding {
	t := Clock(atMS)
	severity := SeverityMinor
	if deathAtPct(in.Fight, atMS) < DeathMajorBeforePct {
		severity = SeverityMajor
	}
	f := &Finding{Code: CodeDeath, Severity: severity, Facts: deathFacts(hit, t)}
	for i := range nAnchors {
		f.Anchors = append(f.Anchors, "#"+strconv.Itoa(i))
	}
	if hasReady {
		f.Facts.Spells = append(f.Facts.Spells, SpellFact{ID: ready.SpellID, Name: ready.Name, Rank: in.Tables.Spells[ready.SpellID].Rank})
		f.Title = Fill(TitleDiedWithReady, map[string]string{"t": t, "Spell": ready.Name})
		f.Basis = Basis{Label: Fill(BasisDeathCD, map[string]string{"Spell": ready.Name}), Kind: BasisSpellData}
		f.Detail = deathDetail(in, hit, atMS, ready)
		return f
	}
	f.Title = Fill(TitleDiedTo, map[string]string{"t": t, "Spell": hit.Name})
	f.Basis = Basis{Label: FromYourLog + ".", Kind: BasisOwnLog}
	if deathAvoidable(in, hit) {
		f.Basis.Label = Fill(BasisAvoidable, map[string]string{"Spell": hit.Name})
	}
	f.Detail = deathDetail(in, hit, atMS, deathReady{})
	return f
}

// deathDetail: with a defensive ready since the pull, §5.3's sentence; otherwise the hit
// (when the log kept its amount) and, when no healing potion or Healthstone was cast in
// the window, that line.
func deathDetail(in Input, hit deathHit, atMS int64, ready deathReady) string {
	if ready.SincePull {
		return Fill(DetailDiedWithReady, map[string]string{"Killer": hit.Name, "t": Clock(atMS), "Spell": ready.Name})
	}
	var parts []string
	if hit.Amount > 0 {
		parts = append(parts, Fill(DetailDeathHit, map[string]string{
			"Spell": hit.Name, "amount": Thousands(hit.Amount), "t": Clock(hit.AtMS),
		}))
	}
	if !potionBefore(in.Fight, in.Row.GUID, atMS) {
		parts = append(parts, DeathNoPotion)
	}
	return strings.Join(parts, " ")
}

func deathFacts(hit deathHit, t string) Facts {
	facts := Facts{Spells: []SpellFact{{ID: hit.SpellID, Name: hit.Name}}, Times: []string{t}}
	if hit.Amount > 0 {
		facts.Numbers = []NumberFact{{V: float64(hit.Amount), Text: Thousands(hit.Amount), Unit: deathFactUnit, Of: deathFactAmount}}
		if ht := Clock(hit.AtMS); ht != t {
			facts.Times = append(facts.Times, ht)
		}
	}
	return facts
}
