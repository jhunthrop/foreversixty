package rating

// The per-fight facts the addon's last-raid region shows (docs/contracts/addon-v2.md
// contract 3, the last_raid message; docs/superpowers/specs/2026-10-11-addon-v2-design.md
// §4.7, §5.2). They are worked out once, here, when the fight is rated, and stored on the
// player's rating_scores row (migration 0036_last_raid): the inbox reads rows and never a
// fight summary, so a companion polling every ten minutes never touches the bucket.

import (
	"math"
	"strings"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// The biggest-loss phrases. Each is plain words the addon shows as given (contract 3's
// copy rule): never an id, never a code.
const (
	LossDowntime     = "downtime"
	LossDeadEarly    = "dead early"
	LossDiedTo       = "died to "
	LossADeath       = "a death"
	LossAvoidable    = "standing in fire"
	LossConsumables  = "missing consumables"
	LossMechanics    = "missed mechanics"
	LossOutput       = "low output"
	LossUtility      = "utility"
	consumableFlask  = "flask"
	consumableFood   = "food"
	unknownCauseName = "Unknown cause" // the engine's own name for a death with no killing blow
)

// LastRaidFacts is what the last_raid message says about one player's fight.
type LastRaidFacts struct {
	// Deaths is how many times the player died.
	Deaths int
	// BiggestLoss is the plain phrase for the component that cost the most points; empty
	// when no scored component cost anything.
	BiggestLoss string
	// LossSeconds is the downtime seconds when the loss is downtime, else the seconds the
	// player lay dead, else 0.
	LossSeconds int
	// ConsumablesMissed names the consumable categories the player pulled without; empty,
	// never nil, when none was missed or the log cannot say.
	ConsumablesMissed []string
	// Adherence is the rotation adherence (adherence.go), 0 to 1 to two decimals; nil when the
	// spec has no shipped rotation or the player cast fewer than AdherenceMinCasts abilities.
	// Set by the caller (Store.RateFight), which holds the database for the character's level.
	Adherence *float64
}

// lastRaidFacts works out row's facts from its scored card and the fight. idleMS is the
// Activity component's own idle time for the player (ratingengine.ActivityIdleMS, under
// the curated table and assignments the card was scored with).
func lastRaidFacts(fight summary.Summary, row summary.RosterRow, card ratingengine.Card, idleMS int64) LastRaidFacts {
	phrase, seconds := biggestLoss(fight, row, card, idleMS)
	return LastRaidFacts{
		Deaths:            deathsOf(fight, row),
		BiggestLoss:       phrase,
		LossSeconds:       seconds,
		ConsumablesMissed: consumablesMissed(fight, row),
	}
}

// biggestLoss is ruling 3 of the L6.3 brief, in one place: the card's scored component
// that cost the most points - its shortfall from 100 times the weight the card applied -
// named in plain words, and the seconds that go with it. An excluded component cost
// nothing (it was not scored); a tie goes to the component the card lists first.
//
//	activity    -> "downtime"; seconds: idleMS, the Activity component's own measure of
//	               the time alive, outside downtime and assignment windows, with nothing done
//	survival    -> after a death in the first half of the fight, "dead early"; after a
//	               later one, "died to <the engine's cause, lower-cased>", or "a death"
//	               when the log never named a cause; with no death (avoidable damage
//	               alone), "standing in fire"
//	preparation -> "missing consumables"
//	mechanics   -> "missed mechanics"
//	output      -> "low output"
//	utility     -> "utility"
//
// For anything but downtime, seconds is the time the player lay dead (0 alive).
func biggestLoss(fight summary.Summary, row summary.RosterRow, card ratingengine.Card, idleMS int64) (string, int) {
	var worst *ratingengine.Component
	worstCost := 0.0
	for i := range card.Components {
		c := &card.Components[i]
		if c.Excluded {
			continue
		}
		if cost := (100 - c.Score) * c.Weight; cost > worstCost {
			worst, worstCost = c, cost
		}
	}
	if worst == nil {
		return "", 0
	}
	deathAt, died := firstDeathMS(fight, row)
	deadSeconds := 0
	if died {
		deadSeconds = msToSeconds(fight.DurationMS - deathAt)
	}
	switch worst.Name {
	case ratingengine.ComponentNameActivity:
		return LossDowntime, msToSeconds(idleMS)
	case ratingengine.ComponentNameSurvival:
		if !died {
			return LossAvoidable, 0
		}
		if fight.DurationMS > 0 && deathAt*2 < fight.DurationMS {
			return LossDeadEarly, deadSeconds
		}
		return deathCause(worst.Moments), deadSeconds
	case ratingengine.ComponentNamePreparation:
		return LossConsumables, deadSeconds
	case ratingengine.ComponentNameMechanics:
		return LossMechanics, deadSeconds
	case ratingengine.ComponentNameOutput:
		return LossOutput, deadSeconds
	case ratingengine.ComponentNameUtility:
		return LossUtility, deadSeconds
	}
	return strings.ReplaceAll(worst.Name, "_", " "), deadSeconds
}

// deathCause is "died to <cause>" from the survival component's first death moment, the
// engine's own spell name lower-cased, or "a death" when the engine had no cause to name.
func deathCause(moments []ratingengine.Moment) string {
	for _, m := range moments {
		if m.Kind != "death" {
			continue
		}
		name := strings.TrimSpace(m.SpellName)
		if name == "" || name == unknownCauseName {
			return LossADeath
		}
		return LossDiedTo + strings.ToLower(name)
	}
	return LossADeath
}

// firstDeathMS is when row's player first died, and whether they did.
func firstDeathMS(fight summary.Summary, row summary.RosterRow) (int64, bool) {
	first, died := int64(0), false
	for _, d := range fight.Deaths {
		if d.GUID == row.GUID && (!died || d.AtMS < first) {
			first, died = d.AtMS, true
		}
	}
	return first, died
}

// deathsOf counts row's deaths in the fight's death list, falling back to the roster's own
// count when the list has none for the player.
func deathsOf(fight summary.Summary, row summary.RosterRow) int {
	n := 0
	for _, d := range fight.Deaths {
		if d.GUID == row.GUID {
			n++
		}
	}
	if n == 0 {
		return row.Deaths
	}
	return n
}

func msToSeconds(ms int64) int {
	if ms <= 0 {
		return 0
	}
	return int(math.Round(float64(ms) / 1000))
}

// consumablesMissed is the plain names of the categories the player pulled without, read
// through consumables.AtPull against the same catalogue the discipline fact uses
// (disciplineCatalogue: flask and food; the weapon enchant stays out until a beta log
// shows how stones and oils appear, see score_v2.go). A category is named only when the
// pull proves it absent (consumables.RoleCatalogue.MissedAtPull): with no character sheet,
// or with an incomplete list (food today), the log cannot say, and it is left out: unknown
// is never missing.
func consumablesMissed(fight summary.Summary, row summary.RosterRow) []string {
	missed := []string{}
	pull := consumables.AtPull(fight, row.GUID, disciplineCatalogue)
	if disciplineCatalogue.MissedAtPull(pull, consumables.CategoryFlask) {
		missed = append(missed, consumableFlask)
	}
	if disciplineCatalogue.MissedAtPull(pull, consumables.CategoryFood) {
		missed = append(missed, consumableFood)
	}
	return missed
}
