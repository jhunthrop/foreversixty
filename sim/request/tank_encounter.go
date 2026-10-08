package request

// The tank fight.
//
// A tank spec is measured against one curated statement of the fight,
// data/curated/tank-encounter.json: a boss that swings at the character,
// and the healers the raid assigns to it. This file is the engine-side
// half of that statement. The curated file is the single source and
// carries the reasons; sim/request/tank-encounter.json is its embedded
// copy (`make tank-encounter-sync`), the way sim/request/apl/ carries
// the rotations, and TestTankEncounterCopyMatchesCurated proves the two
// agree. Nothing here invents a number: a value changes by editing the
// curated file.

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math"
	"sync"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/specs"
	"github.com/wowsims/classic/sim/core/proto"
)

// RoleTank is the role data/curated/specs.json gives a tank spec.
const RoleTank = "tank"

// tankIndex is the character's slot in the raid's tank list: the raid
// has one player, and the boss hits it.
const tankIndex = 0

//go:embed tank-encounter.json
var tankEncounterJSON []byte

// sourced is one curated value and the reason it is what it is.
type sourced[T any] struct {
	Value  T      `json:"value"`
	Reason string `json:"reason"`
}

type tankEncounterFile struct {
	LevelScaling struct {
		ReferenceLevel  int              `json:"reference_level"`
		Exponent        sourced[float64] `json:"exponent"`
		HealingExponent sourced[float64] `json:"healing_exponent"`
	} `json:"level_scaling"`
	Boss struct {
		SwingSpeedSec sourced[float64] `json:"swing_speed_sec"`
		MinBaseDamage sourced[float64] `json:"min_base_damage"`
		DamageSpread  sourced[float64] `json:"damage_spread"`
		ParryHaste    sourced[bool]    `json:"parry_haste"`
		DualWield     sourced[bool]    `json:"dual_wield"`
	} `json:"boss"`
	Healers struct {
		HPS                 sourced[float64] `json:"hps"`
		CadenceSec          sourced[float64] `json:"cadence_sec"`
		CadenceVariationSec sourced[float64] `json:"cadence_variation_sec"`
		BurstWindowSec      sourced[int32]   `json:"burst_window_sec"`
	} `json:"healers"`
}

// TankBoss is the boss's melee, resolved for one character level.
type TankBoss struct {
	SwingSpeedSec float64
	MinBaseDamage float64
	DamageSpread  float64
	ParryHaste    bool
	DualWield     bool
}

// RawDPS is the damage per second the boss would deal with every swing
// landing as a plain hit and no mitigation of any kind: the mean of a
// swing's roll over its speed. Effective health is measured against it.
func (b TankBoss) RawDPS() float64 {
	return b.MinBaseDamage * (1 + b.DamageSpread/2) / b.SwingSpeedSec
}

// TankHealers is the healing the tank is assumed to receive.
type TankHealers struct {
	HPS                 float64
	CadenceSec          float64
	CadenceVariationSec float64
	BurstWindowSec      int32
}

// TankProfile is the whole fight for one character level.
type TankProfile struct {
	Boss    TankBoss
	Healers TankHealers
}

var (
	tankEncounterOnce sync.Once
	tankEncounter     tankEncounterFile
	tankEncounterErr  error
)

func loadTankEncounter() (tankEncounterFile, error) {
	tankEncounterOnce.Do(func() {
		if err := json.Unmarshal(tankEncounterJSON, &tankEncounter); err != nil {
			tankEncounterErr = fmt.Errorf("request: tank-encounter.json is corrupt: %w", err)
			return
		}
		tankEncounterErr = tankEncounter.validate()
	})
	return tankEncounter, tankEncounterErr
}

func (f tankEncounterFile) validate() error {
	positive := map[string]float64{
		"level_scaling.reference_level":  float64(f.LevelScaling.ReferenceLevel),
		"level_scaling.exponent":         f.LevelScaling.Exponent.Value,
		"level_scaling.healing_exponent": f.LevelScaling.HealingExponent.Value,
		"boss.swing_speed_sec":           f.Boss.SwingSpeedSec.Value,
		"boss.min_base_damage":           f.Boss.MinBaseDamage.Value,
		"healers.hps":                    f.Healers.HPS.Value,
		"healers.burst_window_sec":       float64(f.Healers.BurstWindowSec.Value),
	}
	for name, v := range positive {
		if v <= 0 {
			return fmt.Errorf("request: tank-encounter.json: %s must be positive, got %v", name, v)
		}
	}
	if f.Boss.DamageSpread.Value < 0 {
		return fmt.Errorf("request: tank-encounter.json: boss.damage_spread must not be negative")
	}
	return nil
}

// TankProfileForLevel is the tank fight for a character of characterLevel:
// the curated numbers, with damage and healing scaled by the character's
// level over the reference level (data/curated/tank-encounter.json's
// level_scaling).
func TankProfileForLevel(characterLevel int) (TankProfile, error) {
	f, err := loadTankEncounter()
	if err != nil {
		return TankProfile{}, err
	}
	levelRatio := float64(characterLevel) / float64(f.LevelScaling.ReferenceLevel)
	damageScale := math.Pow(levelRatio, f.LevelScaling.Exponent.Value)
	healingScale := math.Pow(levelRatio, f.LevelScaling.HealingExponent.Value)
	return TankProfile{
		Boss: TankBoss{
			SwingSpeedSec: f.Boss.SwingSpeedSec.Value,
			MinBaseDamage: f.Boss.MinBaseDamage.Value * damageScale,
			DamageSpread:  f.Boss.DamageSpread.Value,
			ParryHaste:    f.Boss.ParryHaste.Value,
			DualWield:     f.Boss.DualWield.Value,
		},
		Healers: TankHealers{
			HPS:                 f.Healers.HPS.Value * healingScale,
			CadenceSec:          f.Healers.CadenceSec.Value,
			CadenceVariationSec: f.Healers.CadenceVariationSec.Value,
			BurstWindowSec:      f.Healers.BurstWindowSec.Value,
		},
	}, nil
}

// IsTankSpec reports whether spec is a tank on sim/specs' canonical list.
func IsTankSpec(spec string) bool {
	known, ok := specs.ByKey[spec]
	return ok && known.Role == RoleTank
}

// applyTankFight turns an ordinary single-target request into the tank
// fight: the boss swings at the character, the character stands in front
// of it, and the healers heal. A training dummy has no melee and keeps
// none.
func applyTankFight(player *proto.Player, raid *proto.Raid, enc *proto.Encounter, req api.SimRequest) error {
	if req.Encounter.Dummy {
		return nil
	}
	profile, err := TankProfileForLevel(req.Character.Level)
	if err != nil {
		return err
	}
	player.InFrontOfTarget = true
	player.HealingModel = &proto.HealingModel{
		Hps:              profile.Healers.HPS,
		CadenceSeconds:   profile.Healers.CadenceSec,
		CadenceVariation: profile.Healers.CadenceVariationSec,
		BurstWindow:      profile.Healers.BurstWindowSec,
	}
	raid.Tanks = []*proto.UnitReference{{Type: proto.UnitReference_Player, Index: tankIndex}}
	for _, target := range enc.Targets {
		target.TankIndex = tankIndex
		target.SwingSpeed = profile.Boss.SwingSpeedSec
		target.MinBaseDamage = profile.Boss.MinBaseDamage
		target.DamageSpread = profile.Boss.DamageSpread
		target.ParryHaste = profile.Boss.ParryHaste
		target.DualWield = profile.Boss.DualWield
	}
	return nil
}
