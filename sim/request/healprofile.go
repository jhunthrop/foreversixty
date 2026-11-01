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
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"math"
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

// HealTuning is the multiplier the tank hit and the raid pulse are scaled by (the retune
// rule), the bounds it was tuned between, and the unscaled figures it multiplies.
type HealTuning struct {
	Scale             float64   `json:"scale"`
	Bounds            []float64 `json:"bounds"`
	BaseTankHitDamage float64   `json:"base_tank_hit_damage"`
	BasePulseDamage   float64   `json:"base_pulse_damage"`
	Reason            string    `json:"reason"`
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
	Tuning       HealTuning          `json:"tuning"`
	Tank         HealTank            `json:"tank"`
	Members      HealMembers         `json:"members"`
	Pulse        HealPulse           `json:"pulse"`
	Sources      []HealProfileSource `json:"sources"`
}

// healProfileJSON is data/curated/heal-profile.json, embedded the way
// tank-encounter.json is (`make heal-profile-sync`), so a request built in
// the browser's wasm can attach the same fake raid the ranker does.
// TestEmbeddedHealProfileMatchesCurated proves the copy is current.
//
//go:embed heal-profile.json
var healProfileJSON []byte

// LoadHealProfile reads and validates the curated profile at path.
func LoadHealProfile(path string) (HealProfile, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return HealProfile{}, fmt.Errorf("request: reading the heal profile: %w", err)
	}
	return parseHealProfile(raw, path)
}

// EmbeddedHealProfile is the curated profile compiled into this build.
func EmbeddedHealProfile() (HealProfile, error) {
	return parseHealProfile(healProfileJSON, "embedded heal-profile.json")
}

func parseHealProfile(raw []byte, name string) (HealProfile, error) {
	var profile HealProfile
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&profile); err != nil {
		return HealProfile{}, fmt.Errorf("request: decoding %s: %w", name, err)
	}
	if err := profile.validate(); err != nil {
		return HealProfile{}, fmt.Errorf("%s: %w", name, err)
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
	return p.Tuning.validate(p.Tank.HitDamage, p.Pulse.Damage)
}

// validate checks the bounds and that the tuned figures are the base figures times the scale.
func (t HealTuning) validate(hitDamage, pulseDamage float64) error {
	if len(t.Bounds) != 2 || t.Bounds[0] <= 0 || t.Bounds[0] > t.Bounds[1] {
		return fmt.Errorf("%w: tuning.bounds must be two positive scales, low first", ErrBadHealProfile)
	}
	if t.Scale < t.Bounds[0] || t.Scale > t.Bounds[1] {
		return fmt.Errorf("%w: tuning.scale %v is outside its bounds %v", ErrBadHealProfile, t.Scale, t.Bounds)
	}
	if t.Reason == "" {
		return fmt.Errorf("%w: tuning states its reason", ErrBadHealProfile)
	}
	if scaled(t.BaseTankHitDamage, t.Scale) != hitDamage || scaled(t.BasePulseDamage, t.Scale) != pulseDamage {
		return fmt.Errorf("%w: tank.hit_damage and pulse.damage must be the base figures times tuning.scale", ErrBadHealProfile)
	}
	return nil
}

// scaled is a base figure times a scale, to the whole point of damage the profile states.
func scaled(base, scale float64) float64 { return math.Round(base * scale) }

// AtScale is the profile with the tank hit and the raid pulse scaled by scale in place of
// the tuned one; everything else is kept. The receiver is not changed.
func (p HealProfile) AtScale(scale float64) HealProfile {
	out := p
	out.Tuning.Scale = scale
	out.Tank.HitDamage = scaled(p.Tuning.BaseTankHitDamage, scale)
	out.Pulse.Damage = scaled(p.Tuning.BasePulseDamage, scale)
	return out
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
