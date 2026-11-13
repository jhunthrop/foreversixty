package request

// The raid-shaped healing fight (data/curated/heal-raid-profile.json).
//
// heal-profile.json is the fight the BiS ranking picks healer gear on: scaled up until
// gear separates, so every healer casts flat out until its mana runs dry. This profile is
// the second fight, the one the /tiers Healers tab ranks the five healers on (owner
// ruling, 2026-11-13): one healer's share of a logged level 60 raid kill, fitted to a
// stated log sample, so that idle time and overheal exist. It is a level 60 raid fight
// only, so it has no bands and no tuning; its tank hit, pulse and length are derived from
// its log sample, and loading refuses a file whose figures are not that derivation.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"os"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// RaidHealProfileFile is the raid-shaped profile under data/curated.
const RaidHealProfileFile = "heal-raid-profile.json"

// raidFightRoundSeconds is what the logged 90th percentile kill is rounded up to.
const raidFightRoundSeconds = 10

// RaidLogSample is the log sample a raid-shaped profile is fitted to.
type RaidLogSample struct {
	Source                string             `json:"source"`
	Collector             string             `json:"collector"`
	Kills                 int                `json:"kills"`
	HealerKills           int                `json:"healer_kills"`
	KillSecondsMedian     float64            `json:"kill_seconds_median"`
	KillSecondsP90        float64            `json:"kill_seconds_p90"`
	HealerDamagePerSecond float64            `json:"healer_damage_per_second"`
	TankShare             float64            `json:"tank_share"`
	TankSpike5            float64            `json:"tank_spike5"`
	OverhealMedianBySpec  map[string]float64 `json:"overheal_median_by_spec"`
	Reason                string             `json:"reason"`
}

// RaidHealProfile is data/curated/heal-raid-profile.json.
type RaidHealProfile struct {
	ID           string              `json:"id"`
	Label        string              `json:"label"`
	Summary      string              `json:"summary"`
	Notes        string              `json:"notes"`
	LogSample    RaidLogSample       `json:"log_sample"`
	DurationSec  int                 `json:"duration_sec"`
	DamageSpread float64             `json:"damage_spread"`
	Tank         HealTank            `json:"tank"`
	Members      HealMembers         `json:"members"`
	Pulse        HealPulse           `json:"pulse"`
	Sources      []HealProfileSource `json:"sources"`
}

// HealFight is a fight a healer can be run in: the fake raid and its damage laid onto a
// built engine request. HealProfile (at the healer's band) and RaidHealProfile both are.
type HealFight interface {
	Attach(req *proto.RaidSimRequest)
}

// LoadRaidHealProfile reads and validates the raid-shaped profile at path.
func LoadRaidHealProfile(path string) (RaidHealProfile, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return RaidHealProfile{}, fmt.Errorf("request: reading the raid heal profile: %w", err)
	}
	var profile RaidHealProfile
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&profile); err != nil {
		return RaidHealProfile{}, fmt.Errorf("request: decoding %s: %w", path, err)
	}
	if err := profile.validate(); err != nil {
		return RaidHealProfile{}, fmt.Errorf("%s: %w", path, err)
	}
	return profile, nil
}

func (p RaidHealProfile) validate() error {
	if p.ID == "" || p.Label == "" || p.Notes == "" {
		return fmt.Errorf("%w: id, label and notes are required", ErrBadHealProfile)
	}
	if err := p.LogSample.validate(); err != nil {
		return err
	}
	positive := map[string]float64{
		"tank.health":            p.Tank.Health,
		"tank.swing_seconds":     p.Tank.SwingSeconds,
		"members.health":         p.Members.Health,
		"pulse.interval_seconds": p.Pulse.IntervalSeconds,
		"pulse.members":          float64(p.Pulse.Members),
	}
	for name, value := range positive {
		if value <= 0 {
			return fmt.Errorf("%w: %s must be above zero", ErrBadHealProfile, name)
		}
	}
	if p.DamageSpread < 0 || p.DamageSpread >= 1 {
		return fmt.Errorf("%w: damage_spread must be in [0, 1)", ErrBadHealProfile)
	}
	if p.Tank.Reason == "" || p.Members.Reason == "" || p.Pulse.Reason == "" {
		return fmt.Errorf("%w: every figure group states its reason", ErrBadHealProfile)
	}
	return p.validateDerivation()
}

func (s RaidLogSample) validate() error {
	if s.Source == "" || s.Collector == "" || s.Reason == "" {
		return fmt.Errorf("%w: log_sample states its source, its collector and its reason", ErrBadHealProfile)
	}
	if s.Kills <= 0 || s.HealerKills <= 0 || s.KillSecondsP90 <= 0 || s.HealerDamagePerSecond <= 0 || s.TankSpike5 <= 0 {
		return fmt.Errorf("%w: log_sample's counts and figures must be above zero", ErrBadHealProfile)
	}
	if s.TankShare <= 0 || s.TankShare >= 1 {
		return fmt.Errorf("%w: log_sample.tank_share must be in (0, 1)", ErrBadHealProfile)
	}
	return nil
}

// validateDerivation refuses figures that are not the log sample's: the fight is the 90th
// percentile kill rounded up to ten seconds, and the healer's damage a second splits into
// the tank hit and the pulse by the tank share, at the stated cadences.
func (p RaidHealProfile) validateDerivation() error {
	want := p.derived()
	if p.DurationSec != want.durationSec {
		return fmt.Errorf("%w: duration_sec %d is not the logged 90th percentile kill rounded up (%d)", ErrBadHealProfile, p.DurationSec, want.durationSec)
	}
	if p.Tank.HitDamage != want.tankHit || p.Pulse.Damage != want.pulse {
		return fmt.Errorf("%w: tank.hit_damage %v and pulse.damage %v are not the log sample's (%v and %v)", ErrBadHealProfile, p.Tank.HitDamage, p.Pulse.Damage, want.tankHit, want.pulse)
	}
	return nil
}

type raidDerivation struct {
	durationSec    int
	tankHit, pulse float64
}

func (p RaidHealProfile) derived() raidDerivation {
	s := p.LogSample
	return raidDerivation{
		durationSec: int(math.Ceil(s.KillSecondsP90/raidFightRoundSeconds)) * raidFightRoundSeconds,
		tankHit:     math.Round(s.HealerDamagePerSecond * s.TankShare * p.Tank.SwingSeconds),
		pulse:       math.Round(s.HealerDamagePerSecond * (1 - s.TankShare) * p.Pulse.IntervalSeconds / float64(p.Pulse.Members)),
	}
}

// Model is the engine's raid damage model for the profile.
func (p RaidHealProfile) Model() *proto.RaidDamageModel {
	return &proto.RaidDamageModel{
		Profile:              p.ID,
		TankHealth:           p.Tank.Health,
		MemberHealth:         p.Members.Health,
		TankHitDamage:        p.Tank.HitDamage,
		TankSwingSeconds:     p.Tank.SwingSeconds,
		DamageSpread:         p.DamageSpread,
		PulseDamage:          p.Pulse.Damage,
		PulseIntervalSeconds: p.Pulse.IntervalSeconds,
		PulseMembers:         int32(p.Pulse.Members),
		PulsesReachHealer:    p.Pulse.ReachesHealer,
	}
}

// Attach lays the fight onto a built engine request: the fake raid, the damage it takes and
// the fight length with no variation. It is a level 60 raid fight at every level it is run.
func (p RaidHealProfile) Attach(req *proto.RaidSimRequest) {
	core.AddHealingFakeRaid(req.Raid, p.Model())
	req.Encounter.Duration = float64(p.DurationSec)
	req.Encounter.DurationVariation = 0
}
