package request

// The heal profile below 60 (owner ruling 1, sim-accuracy round 10).
//
// The curated profile is the level-60 raid fight. Each leveling band fights the same fight
// at its own level: bands.rule in data/curated/heal-profile.json states how a band's
// content and its figures are derived (its last bosses' melee and spell damage, the client's
// dungeon damage modifier and expected player health), and each band row states the
// figures it comes to. This file is the one place that turns a band row into
// the engine's raid damage model; Attach and AttachWeights resolve the band from the
// healer's own level, so every tool that runs a healer fights the same band.

import (
	"fmt"
	"math"
	"sort"

	"github.com/wowsims/classic/sim/core/proto"
)

// HealBandContent is one dungeon (or raid) a band's content is made of, by its last boss
// and that boss's creature_template melee figures.
type HealBandContent struct {
	Name             string          `json:"name"`
	LastBoss         string          `json:"last_boss"`
	NPCID            int             `json:"npc_id"`
	BossLevel        int             `json:"boss_level"`
	MinMeleeDmg      float64         `json:"min_melee_dmg"`
	MaxMeleeDmg      float64         `json:"max_melee_dmg"`
	DamageMultiplier float64         `json:"damage_multiplier"`
	MeleeAttackPower float64         `json:"melee_attack_power"`
	AttackTimeMS     float64         `json:"melee_base_attack_time_ms"`
	Spells           []HealBossSpell `json:"spells"`
}

// HealBossSpell is one damaging spell a last boss casts in combat: its mean damage a cast
// at the boss's level, the mean of its AI's repeat timer, and how long the boss spends
// casting (or channelling) it. Each row names its source and how the damage was derived.
type HealBossSpell struct {
	SpellID         int     `json:"spell_id"`
	Name            string  `json:"name"`
	Source          string  `json:"source"`
	DamagePerCast   float64 `json:"damage_per_cast"`
	Derivation      string  `json:"derivation"`
	IntervalSeconds float64 `json:"interval_seconds"`
	CastSeconds     float64 `json:"cast_seconds"`
}

// meleeDPS is the boss's melee damage a second: its hit as cmangos'
// Creature::UpdateDamagePhysical makes it (attack power over 14, plus the mean weapon
// damage times the damage multiplier) over its attack time.
func (c HealBandContent) meleeDPS() float64 {
	hit := c.MeleeAttackPower/14 + (c.MinMeleeDmg+c.MaxMeleeDmg)/2*c.DamageMultiplier
	return hit / (c.AttackTimeMS / 1000)
}

// damagePerSecond is the boss's melee, less the share of the fight it spends casting, plus
// each spell's damage a cast over its interval.
func (c HealBandContent) damagePerSecond() float64 {
	casting, spells := 0.0, 0.0
	for _, s := range c.Spells {
		casting += s.CastSeconds / s.IntervalSeconds
		spells += s.DamagePerCast / s.IntervalSeconds
	}
	return c.meleeDPS()*max(0, 1-casting) + spells
}

// HealPullMob is one kind of trash mob in a band's stated pull: how many of it the pull
// holds, its creature_template melee figures and the damage a second it deals to the tank
// (its melee and the spells it aims at its victim) and to the group (spells at a random
// target or an area), derived offline the way a last boss's are.
type HealPullMob struct {
	NPCID                int                `json:"npc_id"`
	Name                 string             `json:"name"`
	Count                int                `json:"count"`
	Level                int                `json:"level"`
	Rank                 int                `json:"rank"`
	MinMeleeDmg          float64            `json:"min_melee_dmg"`
	MaxMeleeDmg          float64            `json:"max_melee_dmg"`
	DamageMultiplier     float64            `json:"damage_multiplier"`
	MeleeAttackPower     float64            `json:"melee_attack_power"`
	AttackTimeMS         float64            `json:"melee_base_attack_time_ms"`
	TankDamagePerSecond  float64            `json:"tank_damage_per_second"`
	GroupDamagePerSecond float64            `json:"group_damage_per_second"`
	Spells               []HealPullMobSpell `json:"spells"`
}

// HealPullMobSpell is one damaging spell a trash mob casts in combat, with its source row.
type HealPullMobSpell struct {
	SpellID         int     `json:"spell_id"`
	Name            string  `json:"name"`
	DamagePerCast   float64 `json:"damage_per_cast"`
	IntervalSeconds float64 `json:"interval_seconds"`
	Hits            string  `json:"hits"`
	Source          string  `json:"source"`
}

// HealPull is the pull at the band's stated percentile: where it stands and what it holds.
type HealPull struct {
	Dungeon       string        `json:"dungeon"`
	CreatureGUIDs []int         `json:"creature_guids"`
	Mobs          []HealPullMob `json:"mobs"`
}

// damagePerSecond is the pull's damage a second: every mob's, tank and group, at its count.
func (p HealPull) damagePerSecond() float64 {
	total := 0.0
	for _, m := range p.Mobs {
		total += float64(m.Count) * (m.TankDamagePerSecond + m.GroupDamagePerSecond)
	}
	return total
}

// HealPullDungeon is one of the band's dungeons in the pull count.
type HealPullDungeon struct {
	Name                      string `json:"name"`
	MapID                     int    `json:"map_id"`
	TrashMobs                 int    `json:"trash_mobs"`
	TrashPulls                int    `json:"trash_pulls"`
	MobWeightedMedianPullSize int    `json:"mob_weighted_median_pull_size"`
}

// HealBandPulls is the band's trash, pulled under the pull rule (bands.pull_rule): how many
// mobs and pulls its dungeons hold, the median and the stated percentile's pull (both
// weighted by mobs, so a pull counts once for each mob killed in it), and the share of the
// trash's damage aimed at the tank.
type HealBandPulls struct {
	TrashMobs                 int               `json:"trash_mobs"`
	TrashPulls                int               `json:"trash_pulls"`
	MobWeightedMedianPullSize int               `json:"mob_weighted_median_pull_size"`
	PullSize                  int               `json:"pull_size"`
	MedianPullDamagePerSecond float64           `json:"median_pull_damage_per_second"`
	PullDamagePerSecond       float64           `json:"pull_damage_per_second"`
	TankShare                 float64           `json:"tank_share"`
	Dungeons                  []HealPullDungeon `json:"dungeons"`
	Pull                      HealPull          `json:"pull"`
}

// HealBand is one band's content, the client figures for it, and the profile figures the
// rule derives from them at the tuned scale.
type HealBand struct {
	Level            int               `json:"level"`
	Content          []HealBandContent `json:"content"`
	Pulls            *HealBandPulls    `json:"pulls,omitempty"`
	PullMultiplier   float64           `json:"pull_multiplier"`
	ContentDamageMod float64           `json:"content_damage_mod"`
	PlayerHealth     float64           `json:"player_health"`
	TankHealth       float64           `json:"tank_health"`
	MemberHealth     float64           `json:"member_health"`
	TankHitDamage    float64           `json:"tank_hit_damage"`
	PulseDamage      float64           `json:"pulse_damage"`
}

// damage is the band's content damage: the median of its last bosses' damage a second (a
// half-way median between the middle two), times the pull multiplier, times the content's
// damage modifier.
func (b HealBand) damage() float64 {
	return b.bossDamage() * b.pullMultiplier() * b.ContentDamageMod
}

// pullMultiplier is the band's stated trash pull over its median last boss, never below 1:
// the band fights the harder of the two. A band with no pulls (the raid) is 1.
func (b HealBand) pullMultiplier() float64 {
	if b.Pulls == nil {
		return 1
	}
	return max(1, b.Pulls.PullDamagePerSecond/b.bossDamage())
}

// bossDamage is the median of the last bosses' damage a second.
func (b HealBand) bossDamage() float64 {
	dps := make([]float64, len(b.Content))
	for i, c := range b.Content {
		dps[i] = c.damagePerSecond()
	}
	sort.Float64s(dps)
	n := len(dps)
	median := dps[n/2]
	if n%2 == 0 {
		median = (dps[n/2-1] + dps[n/2]) / 2
	}
	return median
}

// HealBands is the profile's band table: the rule, the content window it names, the
// reference band (the curated top-level figures) and one row per band, lowest first.
type HealBands struct {
	Rule           string              `json:"rule"`
	PullRule       HealPullRule        `json:"pull_rule"`
	WindowLevels   int                 `json:"window_levels"`
	ReferenceLevel int                 `json:"reference_level"`
	Notes          string              `json:"notes"`
	Rows           []HealBand          `json:"rows"`
	Sources        []HealProfileSource `json:"sources"`
}

// HealPullRule is how a band's trash is cut into pulls and which pull the band fights:
// creatures within the assistance radius of each other (and spawn groups and links that
// aggro together) are one pull, and the band fights the pull at the stated percentile,
// weighted by mobs. The rule text is in bands.rule.
type HealPullRule struct {
	AssistRadiusYards float64 `json:"assist_radius_yards"`
	Percentile        float64 `json:"percentile"`
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
	if b.PullRule.AssistRadiusYards <= 0 || b.PullRule.Percentile <= 0 || b.PullRule.Percentile >= 1 {
		return fmt.Errorf("%w: bands.pull_rule states a positive assistance radius and a percentile inside (0, 1)", ErrBadHealProfile)
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
			for _, sp := range c.Spells {
				if sp.IntervalSeconds <= 0 || sp.DamagePerCast < 0 || sp.CastSeconds < 0 || sp.Source == "" {
					return fmt.Errorf("%w: band %d: %s's %s states its source, a positive interval and no negative damage", ErrBadHealProfile, row.Level, c.LastBoss, sp.Name)
				}
			}
		}
		if err := b.validatePulls(row); err != nil {
			return err
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

// pullFigureTolerance is the rounding the offline derivation states pull figures to: half a
// tenth of a point a second for each mob in the pull.
const pullFigureTolerance = 0.05

// validatePulls refuses a band below the reference with no pulls, a raid band with pulls,
// a stated pull whose mobs do not sum to its stated damage, and a stated multiplier that is
// not the rule's (to the hundredth it is stated to).
func (b HealBands) validatePulls(row HealBand) error {
	if row.Level == b.ReferenceLevel {
		if row.Pulls != nil || row.PullMultiplier != 1 {
			return fmt.Errorf("%w: the reference band (a raid) states no pulls and a pull multiplier of 1", ErrBadHealProfile)
		}
		return nil
	}
	pulls := row.Pulls
	if pulls == nil || pulls.PullSize <= 0 || len(pulls.Pull.Mobs) == 0 || len(pulls.Dungeons) == 0 || len(pulls.Pull.CreatureGUIDs) != pulls.PullSize {
		return fmt.Errorf("%w: band %d states its dungeons' pulls and the pull it fights, spawn by spawn", ErrBadHealProfile, row.Level)
	}
	count := 0
	for _, m := range pulls.Pull.Mobs {
		count += m.Count
		for _, sp := range m.Spells {
			if sp.Source == "" || sp.IntervalSeconds <= 0 {
				return fmt.Errorf("%w: band %d: %s's %s states its source and a positive interval", ErrBadHealProfile, row.Level, m.Name, sp.Name)
			}
		}
	}
	if count != pulls.PullSize {
		return fmt.Errorf("%w: band %d's pull holds %d mobs, it states %d", ErrBadHealProfile, row.Level, count, pulls.PullSize)
	}
	if got := pulls.Pull.damagePerSecond(); math.Abs(got-pulls.PullDamagePerSecond) > pullFigureTolerance*float64(pulls.PullSize) {
		return fmt.Errorf("%w: band %d's pull mobs deal %.1f a second, it states %.1f", ErrBadHealProfile, row.Level, got, pulls.PullDamagePerSecond)
	}
	if want := math.Round(row.pullMultiplier()*100) / 100; row.PullMultiplier != want {
		return fmt.Errorf("%w: band %d states pull multiplier %v, the rule derives %v", ErrBadHealProfile, row.Level, row.PullMultiplier, want)
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
