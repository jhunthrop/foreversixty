package main

import (
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/wowsims/classic/sim/core"
)

// expertiseToCap is how far a band's weights character is from the boss's
// dodge and parry chances, in engine percent (one point of expertise takes
// one percent off a dodge or parry chance). The site reads it to say
// "expertise to cap": expertise is worth its full weight only until the
// boss can no longer dodge, or for a tank parry, the character.
type expertiseToCap struct {
	// Baseline is the weights character's own expertise, from gear and
	// talents (the rogue's Weapon Expertise).
	Baseline float64 `json:"baseline"`
	// Dodge is the expertise still reducing the boss's dodge chance against
	// this attacker: that chance less Baseline, floored at 0.
	Dodge float64 `json:"dodge"`
	// Parry is the same figure against the boss's parry chance. It is set
	// only for a tank, who faces the boss; a damage dealer stands behind it
	// and is never parried.
	Parry *float64 `json:"parry,omitempty"`
}

// expertiseToCapFor picks the figure a spec publishes: nil for a spec that
// never swings a weapon in melee (leveling.NoMeleeAutoAttackSpecs, the rule
// hitToCapFor uses) or whose engine character does not auto-attack in melee,
// the dodge distance for a damage dealer, plus the parry distance for a tank.
func expertiseToCapFor(spec string, profile core.HitProfile) *expertiseToCap {
	if leveling.NoMeleeAutoAttackSpecs[spec] || !profile.Physical || !profile.Melee {
		return nil
	}
	out := &expertiseToCap{Baseline: profile.Expertise, Dodge: profile.ToDodgeCap()}
	if request.IsTankSpec(spec) {
		parry := profile.ToParryCap()
		out.Parry = &parry
	}
	return out
}
