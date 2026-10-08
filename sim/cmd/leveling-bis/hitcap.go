package main

import (
	"encoding/json"
	"fmt"

	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/wowsims/classic/sim/core"
)

// hitToCap is how far a band's weights character is from the miss-table
// caps, in engine percent (one point is one percent of hit chance, the unit
// the published hit weight is measured in per percent). The site reads it to
// say "hit to cap first": hit is worth its full weight only up to the cap,
// and nothing past it.
type hitToCap struct {
	// Baseline is the weights character's own hit, from gear, talents and
	// buffs.
	Baseline float64 `json:"baseline"`
	// Specials is the hit still worth full value to special attacks:
	// the cap (base miss plus the level-gap suppression) less Baseline.
	Specials float64 `json:"specials"`
	// White is the hit still worth anything to white swings. It is set only
	// for a dual wielder, whose white swings miss 19 points more than a
	// special does; everyone else's white swings share the special cap.
	White *float64 `json:"white,omitempty"`
}

// hitToCapFromProfile reads the engine's hit profile into the published
// figure. A spec that does not swing a weapon (a caster) has no physical
// miss table to cap against, so it publishes none.
func hitToCapFromProfile(profile core.HitProfile) *hitToCap {
	if !profile.Physical {
		return nil
	}
	out := &hitToCap{Baseline: profile.Hit, Specials: profile.ToSpecialCap()}
	if profile.DualWielding {
		white := profile.ToWhiteCap()
		out.White = &white
	}
	return out
}

// spellHitKind marks the caster shape of hit_to_cap; the melee shape has no
// kind, so a reader tells them apart by it.
const spellHitKind = "spell"

// hitToCapFigure is either published shape of hit_to_cap: hitToCap for a
// weapon user, spellHitToCap for a caster.
type hitToCapFigure interface{ isHitToCapFigure() }

func (*hitToCap) isHitToCapFigure()      {}
func (*spellHitToCap) isHitToCapFigure() {}

// publishedHitToCap is the report's hit_to_cap key: either shape, told apart
// on the way back in by the caster shape's kind.
type publishedHitToCap struct{ hitToCapFigure }

func (p publishedHitToCap) MarshalJSON() ([]byte, error) { return json.Marshal(p.hitToCapFigure) }

func (p *publishedHitToCap) UnmarshalJSON(data []byte) error {
	var probe struct {
		Kind string `json:"kind"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return fmt.Errorf("reading hit_to_cap: %w", err)
	}
	var figure hitToCapFigure
	switch probe.Kind {
	case "":
		figure = &hitToCap{}
	case spellHitKind:
		figure = &spellHitToCap{}
	default:
		return fmt.Errorf("hit_to_cap kind %q is neither melee (absent) nor %q", probe.Kind, spellHitKind)
	}
	if err := json.Unmarshal(data, figure); err != nil {
		return fmt.Errorf("reading hit_to_cap: %w", err)
	}
	p.hitToCapFigure = figure
	return nil
}

// reportHitToCap wraps a figure for the report; nil stays absent.
func reportHitToCap(figure hitToCapFigure) *publishedHitToCap {
	if figure == nil {
		return nil
	}
	return &publishedHitToCap{figure}
}

// spellHitToCap is how far a caster's weights character is from the spell
// hit cap, in engine percent. A spell misses at the target's base spell
// miss less a 1% residual (16 points against a level-63 boss), whatever the
// weapon; talents that add hit to a school move that school's own cap in.
type spellHitToCap struct {
	// Kind is always spellHitKind.
	Kind string `json:"kind"`
	// Baseline is the weights character's own hit, from gear, buffs and the
	// stat-wide bonuses every damaging spell gets.
	Baseline float64 `json:"baseline"`
	// Spell is the hit still worth full value to a spell with only that
	// baseline: the 16-point cap less Baseline.
	Spell float64 `json:"spell"`
	// School is set when the character's talents add hit to some spells
	// (Elemental Precision, Arcane Focus, Nature's Reach, Suppression): the
	// distance for those spells, which is less than Spell by the bonus.
	School *schoolHitToCap `json:"school,omitempty"`
}

// schoolHitToCap is the cap distance of the spells that carry a school hit
// bonus.
type schoolHitToCap struct {
	// Names are the schools of those spells, sorted ("Fire", "Frost").
	Names []string `json:"names"`
	// Hit is those spells' total hit, in engine percent.
	Hit float64 `json:"hit"`
	// ToCap is the hit still worth full value to them.
	ToCap float64 `json:"to_cap"`
}

// spellHitToCapFromProfile reads the engine's spell hit profile into the
// published caster figure; nil when the character casts no damaging spell.
func spellHitToCapFromProfile(profile core.HitProfile) *spellHitToCap {
	if !profile.Spell {
		return nil
	}
	spell := profile.SpellProfile
	out := &spellHitToCap{Kind: spellHitKind, Baseline: spell.Hit, Spell: spell.ToSpellCap()}
	if spell.HasSchoolBonus() {
		out.School = &schoolHitToCap{Names: spell.SchoolNames, Hit: spell.SchoolHit, ToCap: spell.ToSchoolCap()}
	}
	return out
}

// hitToCapFor picks the figure a spec publishes: a caster
// (leveling.NoMeleeAutoAttackSpecs) never stands in melee, so its physical
// miss table is beside the point even when the engine's profile reports one
// for the weapon it holds, and it publishes its spell hit distance instead;
// everyone else publishes the melee figure. It returns nil, not a typed nil,
// when there is nothing to publish.
func hitToCapFor(spec string, profile core.HitProfile) hitToCapFigure {
	if leveling.NoMeleeAutoAttackSpecs[spec] {
		if figure := spellHitToCapFromProfile(profile); figure != nil {
			return figure
		}
		return nil
	}
	if figure := hitToCapFromProfile(profile); figure != nil {
		return figure
	}
	return nil
}
