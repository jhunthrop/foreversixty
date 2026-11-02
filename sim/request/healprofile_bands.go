package request

// The heal profile below 60 (owner ruling 1, sim-accuracy round 10).
//
// The curated profile is the level-60 raid fight. Each leveling band fights the same fight
// at its own level: bands.rule in data/curated/heal-profile.json states how a band's
// content and its figures are derived (its last bosses' melee damage, the client's
// dungeon damage modifier and expected player health), and each band row states the
// figures it comes to. This file is the one place that turns a band row into
// the engine's raid damage model; Attach and AttachWeights resolve the band from the
// healer's own level, so every tool that runs a healer fights the same band.

import (
	"fmt"
	"sort"

	"github.com/wowsims/classic/sim/core/proto"
)

// HealBandContent is one dungeon (or raid) a band's content is made of, by its last boss
// and that boss's creature_template melee figures.
type HealBandContent struct {
	Name             string  `json:"name"`
	LastBoss         string  `json:"last_boss"`
	NPCID            int     `json:"npc_id"`
	BossLevel        int     `json:"boss_level"`
	MinMeleeDmg      float64 `json:"min_melee_dmg"`
	MaxMeleeDmg      float64 `json:"max_melee_dmg"`
	DamageMultiplier float64 `json:"damage_multiplier"`
	MeleeAttackPower float64 `json:"melee_attack_power"`
	AttackTimeMS     float64 `json:"melee_base_attack_time_ms"`
}

// meleeDPS is the boss's melee damage a second: its hit as cmangos'
// Creature::UpdateDamagePhysical makes it (attack power over 14, plus the mean weapon
// damage times the damage multiplier) over its attack time.
func (c HealBandContent) meleeDPS() float64 {
	hit := c.MeleeAttackPower/14 + (c.MinMeleeDmg+c.MaxMeleeDmg)/2*c.DamageMultiplier
	return hit / (c.AttackTimeMS / 1000)
}

// HealBand is one band's content, the client figures for it, and the profile figures the
// rule derives from them at the tuned scale.
type HealBand struct {
	Level            int               `json:"level"`
	Content          []HealBandContent `json:"content"`
	ContentDamageMod float64           `json:"content_damage_mod"`
	PlayerHealth     float64           `json:"player_health"`
	TankHealth       float64           `json:"tank_health"`
	MemberHealth     float64           `json:"member_health"`
	TankHitDamage    float64           `json:"tank_hit_damage"`
	PulseDamage      float64           `json:"pulse_damage"`
}

// damage is the band's content damage: the median of its last bosses' melee damage a
// second, a half-way median between the middle two, times the content's damage modifier.
func (b HealBand) damage() float64 {
	dps := make([]float64, len(b.Content))
	for i, c := range b.Content {
		dps[i] = c.meleeDPS()
	}
	sort.Float64s(dps)
	n := len(dps)
	median := dps[n/2]
	if n%2 == 0 {
		median = (dps[n/2-1] + dps[n/2]) / 2
	}
	return median * b.ContentDamageMod
}

// HealBands is the profile's band table: the rule, the content window it names, the
// reference band (the curated top-level figures) and one row per band, lowest first.
type HealBands struct {
	Rule           string              `json:"rule"`
	WindowLevels   int                 `json:"window_levels"`
	ReferenceLevel int                 `json:"reference_level"`
	Notes          string              `json:"notes"`
	Rows           []HealBand          `json:"rows"`
	Sources        []HealProfileSource `json:"sources"`
}

// healBandFigures are the four figures a band changes.
type healBandFigures struct {
	TankHealth, MemberHealth, TankHitDamage, PulseDamage float64
}

// ForLevel is the profile resolved to the band a healer of level fights: the band at or
// below level, band 20 below 20, and the reference band at or above it (or for an unknown
// level, 0). Resolving always starts from the curated figures, so it can be repeated. The
// receiver is not changed.
func (p HealProfile) ForLevel(level int) HealProfile {
	out := p
	out.Band = p.Bands.bandFor(level).Level
	return out
}

func (b HealBands) bandFor(level int) HealBand {
	if level <= 0 || level >= b.ReferenceLevel {
		return b.reference()
	}
	chosen := b.Rows[0]
	for _, row := range b.Rows {
		if row.Level <= level {
			chosen = row
		}
	}
	return chosen
}

func (b HealBands) reference() HealBand {
	for _, row := range b.Rows {
		if row.Level == b.ReferenceLevel {
			return row
		}
	}
	panic("request: the heal profile has no reference band; validateBands refuses such a file")
}

// bandFigures is the rule: against the reference band, the tank hit and the pulse both
// scale with the content's damage, and both healths with the band's player health. The
// tuning scale multiplies the base hit and pulse exactly as at the reference.
func (p HealProfile) bandFigures() healBandFigures {
	return p.figuresFor(p.Bands.bandFor(p.Band), p.Tuning.Scale)
}

func (p HealProfile) figuresFor(row HealBand, scale float64) healBandFigures {
	ref := p.Bands.reference()
	damage := row.damage() / ref.damage()
	health := row.PlayerHealth / ref.PlayerHealth
	return healBandFigures{
		TankHealth:    scaled(p.Tank.Health, health),
		MemberHealth:  scaled(p.Members.Health, health),
		TankHitDamage: scaled(p.Tuning.BaseTankHitDamage*damage, scale),
		PulseDamage:   scaled(p.Tuning.BasePulseDamage*damage, scale),
	}
}

// validateBands refuses a band table the rule cannot run on, and a row whose stated
// figures are not what the rule derives at the tuned scale.
func (p HealProfile) validateBands() error {
	b := p.Bands
	if b.Rule == "" || b.WindowLevels <= 0 || len(b.Rows) == 0 {
		return fmt.Errorf("%w: bands states its rule, its window and its rows", ErrBadHealProfile)
	}
	hasReference := false
	for i, row := range b.Rows {
		if i > 0 && row.Level <= b.Rows[i-1].Level {
			return fmt.Errorf("%w: band rows run lowest first, one per level", ErrBadHealProfile)
		}
		if len(row.Content) == 0 || row.ContentDamageMod <= 0 || row.PlayerHealth <= 0 {
			return fmt.Errorf("%w: band %d states its content and positive client figures", ErrBadHealProfile, row.Level)
		}
		for _, c := range row.Content {
			if c.AttackTimeMS <= 0 || c.meleeDPS() <= 0 {
				return fmt.Errorf("%w: band %d: %s states a positive melee hit and attack time", ErrBadHealProfile, row.Level, c.LastBoss)
			}
		}
		hasReference = hasReference || row.Level == b.ReferenceLevel
	}
	if !hasReference {
		return fmt.Errorf("%w: bands has no row for the reference level %d", ErrBadHealProfile, b.ReferenceLevel)
	}
	if ref := b.reference(); ref.TankHealth != p.Tank.Health || ref.MemberHealth != p.Members.Health || ref.TankHitDamage != p.Tank.HitDamage || ref.PulseDamage != p.Pulse.Damage {
		return fmt.Errorf("%w: the reference band's figures must be the profile's own", ErrBadHealProfile)
	}
	for _, row := range b.Rows {
		want := p.figuresFor(row, p.Tuning.Scale)
		got := healBandFigures{TankHealth: row.TankHealth, MemberHealth: row.MemberHealth, TankHitDamage: row.TankHitDamage, PulseDamage: row.PulseDamage}
		if got != want {
			return fmt.Errorf("%w: band %d states %+v, the rule derives %+v", ErrBadHealProfile, row.Level, got, want)
		}
	}
	return nil
}

// healerLevel is the level of the raid's first player, the healer a profile is attached
// for (fake members are added after it); 0 when the request names none.
func healerLevel(raid *proto.Raid) int {
	for _, party := range raid.GetParties() {
		for _, player := range party.GetPlayers() {
			if player != nil {
				return int(player.GetLevel())
			}
		}
	}
	return 0
}
