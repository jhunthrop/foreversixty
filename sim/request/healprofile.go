package request

// The incoming-damage profile a healer is ranked under.
//
// A healing sim has no boss to hit, so it needs a stated statement of what
// the healer is asked to heal: data/curated/heal-profile.json. This file
// reads it, refuses a profile that cannot be simulated, and lays it onto
// the engine's request as the fake raid (sim/core's RaidDamageModel). The
// profile's name travels with every number measured under it.

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// ErrBadHealProfile is returned for a profile file the engine cannot run.
var ErrBadHealProfile = errors.New("request: the heal profile cannot be simulated")

// HealTank is the fake raid's tank.
type HealTank struct {
	Health       float64 `json:"health"`
	HitDamage    float64 `json:"hit_damage"`
	SwingSeconds float64 `json:"swing_seconds"`
	Reason       string  `json:"reason"`
}

// HealMembers are the fake raid members that are not the tank.
type HealMembers struct {
	Health float64 `json:"health"`
	Reason string  `json:"reason"`
}

// HealPulse is the raid-wide damage the non-tank members take.
type HealPulse struct {
	Damage          float64 `json:"damage"`
	IntervalSeconds float64 `json:"interval_seconds"`
	Members         int     `json:"members"`
	Reason          string  `json:"reason"`
}

// HealProfileSource is where a profile figure comes from.
type HealProfileSource struct {
	Label string `json:"label"`
	URL   string `json:"url"`
	Kind  string `json:"kind"`
}

// HealProfile is data/curated/heal-profile.json.
type HealProfile struct {
	ID           string              `json:"id"`
	Label        string              `json:"label"`
	Summary      string              `json:"summary"`
	Notes        string              `json:"notes"`
	DurationSec  int                 `json:"duration_sec"`
	DamageSpread float64             `json:"damage_spread"`
	Tank         HealTank            `json:"tank"`
	Members      HealMembers         `json:"members"`
	Pulse        HealPulse           `json:"pulse"`
	Sources      []HealProfileSource `json:"sources"`
}

// LoadHealProfile reads and validates the curated profile at path.
func LoadHealProfile(path string) (HealProfile, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return HealProfile{}, fmt.Errorf("request: reading the heal profile: %w", err)
	}
	var profile HealProfile
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&profile); err != nil {
		return HealProfile{}, fmt.Errorf("request: decoding %s: %w", path, err)
	}
	if err := profile.validate(); err != nil {
		return HealProfile{}, fmt.Errorf("%s: %w", path, err)
	}
	return profile, nil
}

func (p HealProfile) validate() error {
	positive := map[string]float64{
		"duration_sec":           float64(p.DurationSec),
		"tank.health":            p.Tank.Health,
		"tank.hit_damage":        p.Tank.HitDamage,
		"tank.swing_seconds":     p.Tank.SwingSeconds,
		"members.health":         p.Members.Health,
		"pulse.damage":           p.Pulse.Damage,
		"pulse.interval_seconds": p.Pulse.IntervalSeconds,
		"pulse.members":          float64(p.Pulse.Members),
	}
	for name, value := range positive {
		if value <= 0 {
			return fmt.Errorf("%w: %s must be above zero", ErrBadHealProfile, name)
		}
	}
	if p.ID == "" || p.Label == "" {
		return fmt.Errorf("%w: id and label are required", ErrBadHealProfile)
	}
	if p.DamageSpread < 0 || p.DamageSpread >= 1 {
		return fmt.Errorf("%w: damage_spread must be in [0, 1)", ErrBadHealProfile)
	}
	if p.Tank.Reason == "" || p.Members.Reason == "" || p.Pulse.Reason == "" {
		return fmt.Errorf("%w: every figure group states its reason", ErrBadHealProfile)
	}
	return nil
}

// Model is the engine's raid damage model for the profile.
func (p HealProfile) Model() *proto.RaidDamageModel {
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
	}
}

// Duration is the fight length the profile is stated for.
func (p HealProfile) Duration() time.Duration {
	return time.Duration(p.DurationSec) * time.Second
}

// Attach lays the profile onto a built engine request: the fake raid, the
// damage it takes, and the fight length with no variation, so "lasts the
// fight" means the same fight every run.
func (p HealProfile) Attach(req *proto.RaidSimRequest) {
	core.AddHealingFakeRaid(req.Raid, p.Model())
	req.Encounter.Duration = float64(p.DurationSec)
	req.Encounter.DurationVariation = 0
}

// AttachWeights is Attach for a stat-weights request.
func (p HealProfile) AttachWeights(req *proto.StatWeightsRequest) {
	req.RaidDamageModel = p.Model()
	req.Encounter.Duration = float64(p.DurationSec)
	req.Encounter.DurationVariation = 0
}
