package request

import "sort"

// LearnedAbility is everything cmd/rotation-search's unused-spell probe
// needs about one class ability by a given level: enough to pick a
// sensible default APL condition (dot, persistent toggle, or neither)
// and to write a legal castSpell action at the right rank. It is read
// through the exact data and grouping sim/request/ladder.go's own
// "learned but unused" check (rule 3) already uses - buildClassAbilities
// and learnedTierAtLevel - rather than a second copy of that walk.
type LearnedAbility struct {
	// Name is the spellranks.json ability name ("Lava Burst").
	Name string
	// IDs are every id sharing this ability's learned rank tier (a
	// ranked spell's rank and its talent-rank variants resolve to the
	// same ability at one level; see buildClassAbilities).
	IDs []int
	// Rank is the rank number to write into a castSpell action's
	// ActionID (action/{"spellId":{"spellId": IDs[0], "rank": Rank}}).
	// Zero means this ability carries no rank at all (spellranks.json
	// tracks it as a single rank-0 row) and the "rank" key is left off
	// the action entirely, the same way the curated rotations write a
	// rankless spell such as Tiger's Fury.
	Rank int
	// LearnedLevel is the level this rank tier is first learned at.
	LearnedLevel int
	// IsDamage is isDamageSpellID's verdict for this tier: a cast the
	// engine credits with dealing damage.
	IsDamage bool
	// IsDOT is a periodic damage effect (aura 3 or 226) with a finite
	// duration - a line a default condition should gate on dotIsActive
	// rather than cast unconditionally.
	IsDOT bool
	// RaisesDamageTaken is a timed debuff that makes the target take more
	// damage (a positive aura-87 effect over a finite duration): Curse of the
	// Elements. It deals none itself, so IsDamage misses it.
	RaisesDamageTaken bool
	// TickSeconds is the periodic effect's own period, in seconds, 0
	// if the ability carries none (set whenever IsDOT is true).
	TickSeconds float64
	// DurationMS is the ability's own duration_ms. -1 is the engine's
	// convention for a toggle that persists until removed or replaced
	// by another of the same kind (a form, a seal, a stance) rather
	// than a timed buff.
	DurationMS int32
	// GCDMS is the ability's own gcd_ms. Zero marks a "spell" that is
	// not a real player-chosen action at all - this build's internal
	// representation of the white-damage melee swing ("Attack") is
	// one, carrying a damage effect and no cooldown of its own, which
	// would make isDamageSpellID call it a damage spell and the
	// engine loop forever if an APL cast it explicitly with no
	// pacing mechanism to stop it recasting itself every tick. A
	// caller inserting a candidate into a priority list unconditionally
	// should require GCDMS > 0 (or a real CooldownMS, carried
	// separately if a future caller needs it) before trusting it is
	// safe to cast at all.
	GCDMS int32
}

// auraModDamageTakenPercent is the client's aura code for "increases damage
// taken by N percent".
const auraModDamageTakenPercent = 87

// LearnedAbilities is every damage-relevant ability a class has
// learned by level, one entry per highest-reached rank tier (ties the
// grouping buildClassAbilities already performs for the ladder's own
// "learned but unused" report to this package's public surface, so a
// caller outside sim/request never re-derives the same groups from
// spellranks.json by hand).
func LearnedAbilities(repoRoot, build, class string, level int) ([]LearnedAbility, error) {
	ranks, err := loadSpellRanks(repoRoot, build)
	if err != nil {
		return nil, err
	}
	consts, err := loadSpellConst(repoRoot, build, class)
	if err != nil {
		return nil, err
	}
	abilities := buildClassAbilities(ranks, class)
	out := make([]LearnedAbility, 0, len(abilities.Tiers))
	for name, tiers := range abilities.Tiers {
		tier, learned := learnedTierAtLevel(tiers, level)
		if !learned {
			continue
		}
		out = append(out, learnedAbilityFrom(consts, name, tier))
	}
	for name, tier := range unrankedSingleCastAbilities(ranks, class) {
		if _, ranked := abilities.Tiers[name]; ranked || tier.Level > level {
			continue
		}
		out = append(out, learnedAbilityFrom(consts, name, tier))
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out, nil
}

// unrankedSingleCastAbilities picks up class abilities spellranks.json
// carries at rank 0 only - this build's own convention for a single-
// rank ability (a talent proc, a stance, a seal, a form) rather than a
// multi-rank spell - that buildClassAbilities skips by design (see its
// own doc comment: Bloodrage, Judgement). Several different rank-0
// rows can share one ability NAME for an unrelated reason (three
// Judgement-of-* casts, an NPC/internal duplicate row) with no rank
// number to disambiguate them the way a real rank tier would, so this
// only admits a name whose rows resolve to exactly one id - Moonkin
// Form, a standalone single-rank strike - and leaves a name with more
// than one distinct id out, the same conservative line
// buildClassAbilities already draws for the ranked case.
func unrankedSingleCastAbilities(all spellRanksFile, class string) map[string]rankTier {
	out := map[string]rankTier{}
	for name, entries := range all.Classes[class] {
		if junkAbilityName.MatchString(name) {
			continue
		}
		ids := map[int]bool{}
		level, allZero := -1, true
		for _, e := range entries {
			if e.Rank != 0 {
				allZero = false
				break
			}
			ids[e.ID] = true
			if e.Level > 0 && (level == -1 || e.Level < level) {
				level = e.Level
			}
		}
		if !allZero || len(ids) != 1 || level == -1 {
			continue
		}
		var id int
		for k := range ids {
			id = k
		}
		out[name] = rankTier{Level: level, IDs: []int{id}}
	}
	return out
}

// learnedAbilityFrom classifies one rank tier: damage (isDamageSpellID),
// dot (a finite-duration periodic damage effect, with its own tick
// length), and the tier's duration_ms, read off its first id's
// spellconst entry - every id sharing a rank tier is the same ability,
// so its constants do not vary id to id.
func learnedAbilityFrom(consts map[int]spellConstEntry, name string, tier rankTier) LearnedAbility {
	out := LearnedAbility{Name: name, IDs: append([]int(nil), tier.IDs...), Rank: tier.Rank, LearnedLevel: tier.Level}
	for _, id := range tier.IDs {
		entry, ok := consts[id]
		if !ok {
			continue
		}
		out.DurationMS = entry.DurationMS
		out.GCDMS = entry.GCDMS
		for _, e := range entry.Effects {
			if isDamageEffect(e) {
				out.IsDamage = true
			}
			if e.Effect == 6 && e.Aura == auraModDamageTakenPercent && e.Amount > 0 && entry.DurationMS > 0 {
				out.RaisesDamageTaken = true
			}
			if e.Effect == 6 && e.Aura == 3 && entry.DurationMS > 0 {
				out.IsDOT = true
				if e.PeriodMS > 0 {
					out.TickSeconds = float64(e.PeriodMS) / 1000
				}
			}
		}
		break // every id in a tier shares one ability's constants.
	}
	return out
}
