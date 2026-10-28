package coach

import (
	"slices"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// The names a healing potion's or a Healthstone's use is cast under, across ranks and
// variants ("Healing Potion", "Superior Healing Potion", "Major Rejuvenation Potion",
// "Major Healthstone"). Conjuring a stone ("Create Healthstone") and the warlock talent
// are not uses.
const (
	deathPotionSuffix      = "Healing Potion"
	deathRejuvSuffix       = "Rejuvenation Potion"
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
	Name    string // a melee swing reads as KillerMelee
	Amount  int64
	AtMS    int64
}

// deathReady is a defensive cooldown the player could have pressed before the death:
// off cooldown since ReadyAtMS, or (NotUsed) never cast this pull with a cooldown longer
// than the time to the death, so a cast on an earlier pull may still have been cooling
// down and only "not used this pull" is proven.
type deathReady struct {
	Name      string
	SpellID   int64
	NotUsed   bool
	ReadyAtMS int64 // when its cooldown had certainly come back; unset when NotUsed
}

// deathDefensives is what the rule file's defensives say at a death: the ones proven
// ready, in list order; the first never-cast one not proven ready; and whether a
// baseline defensive went unjudged for want of the level.
type deathDefensives struct {
	Ready        []deathReady
	NotUsed      *deathReady
	LevelSkipped bool
}

// pick is the defensive the finding names: the first proven ready, else the first not
// used this pull.
func (d deathDefensives) pick() (deathReady, bool) {
	switch {
	case len(d.Ready) > 0:
		return d.Ready[0], true
	case d.NotUsed != nil:
		return *d.NotUsed, true
	}
	return deathReady{}, false
}

// death is F7 (§6.2): the player's first reviewable death, what killed them, whether it
// is on the encounter's avoidable list, a defensive that was ready, and whether a healing
// potion or Healthstone was cast in the DeathPotionWindowMS before. No death at all is a
// strength. A death the log cannot name is not reviewed, and says so; so does a narrowed
// defensive check (no defensives list, or a level the log doesn't record).
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
		return outcome{CouldNotCheck: []string{Fill(CannotDeathKiller, map[string]string{"t": Clock(d.AtMS)})}}
	}
	defs := defensivesAt(in, d.AtMS)
	ready, hasReady := defs.pick()
	anchors := []Anchor{deathAnchor(d.AtMS)}
	if hasReady && !ready.NotUsed { // a defensive never cast has no Casts row to open
		anchors = append(anchors, Anchor{Kind: AnchorCast, SpellID: ready.SpellID})
	}
	return outcome{
		Finding:       deathFinding(in, d.AtMS, hit, defs, len(anchors)),
		CouldNotCheck: defensivesNarrowed(in, hasReady, defs.LevelSkipped),
		Anchors:       anchors,
		Windows:       []ratingengine.Window{{FromMS: d.AtMS, ToMS: d.AtMS}},
	}
}

// defensivesNarrowed is the couldn't-check line for a defensive check the data narrowed:
// no defensives list at all, or a baseline defensive skipped for want of the level when
// nothing else was ready.
func defensivesNarrowed(in Input, hasReady, levelSkipped bool) []string {
	switch {
	case len(in.Tables.Rules.Defensives) == 0:
		return []string{CannotDefensives}
	case levelSkipped && !hasReady:
		return []string{CannotDefensiveLevel}
	}
	return nil
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
// is named by its source (the boss when the log names none), never by a player's or a
// pet's name: a mind-controlled player is "another player's melee", a pet "a pet" (§19:
// no character name leaves for the model). false: nothing names it.
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
		name = meleeKiller(in, *ref)
	}
	return deathHit{SpellID: ref.SpellID, Name: name, Amount: ref.Amount, AtMS: ref.AtMS}, name != ""
}

// meleeKiller names a melee swing by its source: another player or a pet by what it is,
// an NPC by its name, the boss when the log names no NPC; "" when nothing names it.
func meleeKiller(in Input, ref summary.DamageRef) string {
	source := ""
	switch units.Parse(ref.SourceGUID).Kind {
	case units.KindPlayer:
		return KillerPlayerMelee
	case units.KindPet:
		return KillerPet
	case units.KindCreature, units.KindVehicle, units.KindGameObject:
		source = ref.SourceName
	}
	if source == "" || source == ref.SourceGUID {
		source = in.BossName
	}
	if source == "" {
		return ""
	}
	return Fill(KillerMelee, map[string]string{"Mob": source})
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

// deathAnchor is the death at its time; the page adds the player's GUID (no GUID in the pack).
func deathAnchor(atMS int64) Anchor {
	at := atMS
	return Anchor{Kind: AnchorDeath, AtMS: &at}
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
	if strings.HasSuffix(name, deathPotionSuffix) || strings.HasSuffix(name, deathRejuvSuffix) {
		return true
	}
	return strings.HasSuffix(name, deathStoneSuffix) && !strings.HasPrefix(name, deathStoneNotUsePrefix) && name != deathStoneTalent
}

// defensivesAt reads the rule file's defensives at atMS. A talent counts only if cast that
// fight. A cast whose cooldown the data does not state is never claimed ready; one back
// off cooldown is ready from then. A baseline ability never cast counts from its
// trainable level, and only when the level is known: it is proven ready from its
// cooldown's length into the pull (a cast at the end of an earlier pull is back by then),
// and otherwise is only not used this pull (§20.4: fires only when the data proves it).
func defensivesAt(in Input, atMS int64) deathDefensives {
	var out deathDefensives
	for _, e := range in.Tables.Rules.Defensives {
		lastID, lastMS, cast := lastCastBefore(in.Fight, in.Row.GUID, e.SpellIDs, atMS)
		if cast {
			sc, known := in.Tables.Spells[lastID]
			if known && sc.CooldownMS > 0 && lastMS+sc.CooldownMS <= atMS {
				out.Ready = append(out.Ready, deathReady{Name: e.Family, SpellID: lastID, ReadyAtMS: lastMS + sc.CooldownMS})
			}
			continue
		}
		if e.Talent {
			continue
		}
		if in.Row.Level <= 0 {
			out.LevelSkipped = true
			continue
		}
		id, trainable := trainableAtLevel(in, e)
		if !trainable {
			continue
		}
		sc, known := in.Tables.Spells[id]
		switch {
		case known && sc.CooldownMS > 0 && sc.CooldownMS <= atMS:
			out.Ready = append(out.Ready, deathReady{Name: e.Family, SpellID: id, ReadyAtMS: sc.CooldownMS})
		case out.NotUsed == nil:
			out.NotUsed = &deathReady{Name: e.Family, SpellID: id, NotUsed: true}
		}
	}
	return out
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
func deathFinding(in Input, atMS int64, hit deathHit, defs deathDefensives, nAnchors int) *Finding {
	t := Clock(atMS)
	severity := SeverityMinor
	if deathAtPct(in.Fight, atMS) < DeathMajorBeforePct {
		severity = SeverityMajor
	}
	f := &Finding{Code: CodeDeath, Severity: severity, Facts: deathFacts(hit, t)}
	for i := range nAnchors {
		f.Anchors = append(f.Anchors, "#"+strconv.Itoa(i))
	}
	ready, hasReady := defs.pick()
	if !hasReady {
		f.Title = Fill(TitleDiedTo, map[string]string{"t": t, "Spell": hit.Name})
		f.Basis = Basis{Label: BasisFromLog, Kind: BasisOwnLog}
		if deathAvoidable(in, hit) {
			f.Basis.Label = Fill(BasisAvoidable, map[string]string{"Spell": hit.Name})
		}
		f.Detail = deathDetail(in, hit, atMS)
		return f
	}
	f.Facts.Spells = append(f.Facts.Spells, readyFact(in, ready))
	f.Detail = readyDetail(in, hit, atMS, ready)
	if ready.NotUsed {
		f.Title = Fill(TitleDiedNotUsed, map[string]string{"t": t, "Spell": ready.Name})
		f.Basis = Basis{Label: BasisFromLog, Kind: BasisOwnLog}
		return f
	}
	f.Title = Fill(TitleDiedWithReady, map[string]string{"t": t, "Spell": ready.Name})
	f.Basis = Basis{Label: Fill(BasisDeathCD, map[string]string{"Spell": ready.Name}), Kind: BasisSpellData}
	f.Facts.Times = append(f.Facts.Times, Clock(ready.ReadyAtMS))
	if also := defs.Ready[1:]; len(also) > 0 {
		names := make([]string, len(also))
		for i, r := range also {
			names[i] = r.Name
			f.Facts.Spells = append(f.Facts.Spells, readyFact(in, r))
		}
		f.Detail += " " + Fill(DetailAlsoReady, map[string]string{"spells": joinAnd(names)})
	}
	return f
}

func readyFact(in Input, r deathReady) SpellFact {
	return SpellFact{ID: r.SpellID, Name: r.Name, Rank: in.Tables.Spells[r.SpellID].Rank}
}

// readyDetail: the killer and the defensive, ready since its cooldown had certainly come
// back, or not used this pull. An avoidable killer is said too, in §6.2's order: killing
// blow, avoidable, defensive. The killer leads the sentence, so it starts upper case.
func readyDetail(in Input, hit deathHit, atMS int64, ready deathReady) string {
	fields := map[string]string{"Killer": sentenceStart(hit.Name), "t": Clock(atMS), "Spell": ready.Name}
	avoidable := deathAvoidable(in, hit)
	switch {
	case ready.NotUsed && avoidable:
		return Fill(DetailDiedNotUsedAvoidable, fields)
	case ready.NotUsed:
		return Fill(DetailDiedNotUsed, fields)
	}
	fields["since"] = Clock(ready.ReadyAtMS)
	if avoidable {
		return Fill(DetailDiedWithReadyAvoidable, fields)
	}
	return Fill(DetailDiedWithReadySince, fields)
}

// sentenceStart is s with its first letter upper case ("another player's melee" leading
// a sentence).
func sentenceStart(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	if size == 0 {
		return s
	}
	return string(unicode.ToUpper(r)) + s[size:]
}

// deathDetail, with no defensive ready: the hit (when the log kept its amount) and, when
// no healing potion or Healthstone was cast in the window, that line. The basis says an
// avoidable killer.
func deathDetail(in Input, hit deathHit, atMS int64) string {
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
	// A melee killer is SpellFact{ID: 0, Name: "Onyxia's melee"} (KillerMelee): id 0 is
	// the melee sentinel, not a spell. Task 26's validator must admit it by name into the
	// allowed set, never refuse it as a non-positive spell id.
	facts := Facts{Spells: []SpellFact{{ID: hit.SpellID, Name: hit.Name}}, Times: []string{t}}
	if hit.Amount > 0 {
		facts.Numbers = []NumberFact{{V: float64(hit.Amount), Text: Thousands(hit.Amount), Unit: deathFactUnit, Of: deathFactAmount}}
		if ht := Clock(hit.AtMS); ht != t {
			facts.Times = append(facts.Times, ht)
		}
	}
	return facts
}
