package api

import (
	"strings"
	"testing"
)

func TestEncounterAdditionsValidation(t *testing.T) {
	cases := []struct {
		name string
		edit func(*SimRequest)
		want string
	}{
		{"the default encounter is unchanged and legal", func(*SimRequest) {}, ""},
		{"a style label rides along", func(r *SimRequest) { r.Encounter.Style = "heavy-movement" }, ""},
		{"movement away", func(r *SimRequest) {
			r.Encounter.Movement = &Movement{IntervalSec: 20, DurationSec: 5, Kind: MovementAway}
		}, ""},
		{"movement with an unknown kind", func(r *SimRequest) {
			r.Encounter.Movement = &Movement{IntervalSec: 20, DurationSec: 5, Kind: "teleport"}
		}, "movement.kind"},
		{"movement longer than its interval", func(r *SimRequest) {
			r.Encounter.Movement = &Movement{IntervalSec: 5, DurationSec: 20, Kind: MovementAway}
		}, "shorter than the interval"},
		{"movement with no interval", func(r *SimRequest) {
			r.Encounter.Movement = &Movement{DurationSec: 5, Kind: MovementAway}
		}, "movement.interval_sec"},
		{"a target-count timeline", func(r *SimRequest) {
			r.Encounter.TargetsOverTime = []TargetCount{{AtSec: 0, Count: 1}, {AtSec: 40, Count: 3}}
		}, ""},
		{"a timeline that does not start at zero", func(r *SimRequest) {
			r.Encounter.TargetsOverTime = []TargetCount{{AtSec: 40, Count: 3}}
		}, "starts at 0"},
		{"a timeline out of order", func(r *SimRequest) {
			r.Encounter.TargetsOverTime = []TargetCount{{AtSec: 0, Count: 1}, {AtSec: 40, Count: 3}, {AtSec: 20, Count: 5}}
		}, "in time order"},
		{"a timeline over the target cap", func(r *SimRequest) {
			r.Encounter.TargetsOverTime = []TargetCount{{AtSec: 0, Count: MaxTargets + 1}}
		}, "targets_over_time"},
		{"target level 60", func(r *SimRequest) { r.Encounter.TargetLevel = 60 }, ""},
		{"target level 64", func(r *SimRequest) { r.Encounter.TargetLevel = 64 }, "target_level"},
		{"an armor override", func(r *SimRequest) { r.Encounter.TargetArmor = ptr(2500) }, ""},
		// 2026-09-21 result-page review, Defect 3: a typed 0 is a real request (an
		// unarmoured target), not an error, and must validate exactly like any other
		// non-negative override.
		{"an explicit zero armor override", func(r *SimRequest) { r.Encounter.TargetArmor = ptr(0) }, ""},
		{"negative armor", func(r *SimRequest) { r.Encounter.TargetArmor = ptr(-1) }, "target_armor"},
		{"a target type", func(r *SimRequest) { r.Encounter.TargetType = "undead" }, ""},
		{"an unknown target type", func(r *SimRequest) { r.Encounter.TargetType = "murloc" }, "target_type"},
		{"the dummy", func(r *SimRequest) { r.Encounter.Dummy = true }, ""},
		{"a cooldown on cooldown", func(r *SimRequest) {
			r.Character.Cooldowns = []CooldownSpec{{ID: "spell:1719"}}
		}, ""},
		{"a cooldown at fixed times", func(r *SimRequest) {
			r.Character.Cooldowns = []CooldownSpec{{ID: "spell:1719", AtSec: []float64{0, 90}}}
		}, ""},
		{"a cooldown with no id", func(r *SimRequest) {
			r.Character.Cooldowns = []CooldownSpec{{AtSec: []float64{0}}}
		}, "cooldowns"},
		{"a cooldown used before the pull", func(r *SimRequest) {
			r.Character.Cooldowns = []CooldownSpec{{ID: "spell:1719", AtSec: []float64{-1}}}
		}, "at_sec"},
		{"a twenty-second fight", func(r *SimRequest) { r.Encounter.DurationSec = MinDurationSec }, ""},
		{"a ten-minute fight", func(r *SimRequest) { r.Encounter.DurationSec = MaxDurationSec }, ""},
		{"nineteen seconds", func(r *SimRequest) { r.Encounter.DurationSec = 19 }, "duration_sec"},
		{"eleven minutes", func(r *SimRequest) { r.Encounter.DurationSec = 660 }, "duration_sec"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := runReq()
			c.edit(&req)
			err := req.Validate()
			switch {
			case c.want == "" && err != nil:
				t.Fatalf("a legal request was refused: %v", err)
			case c.want != "" && err == nil:
				t.Fatalf("an illegal request was accepted; the error should mention %q", c.want)
			case c.want != "" && !strings.Contains(err.Error(), c.want):
				t.Errorf("error %q does not mention %q", err, c.want)
			}
		})
	}
}

// encounter.target_level's own bounds move with the character: a
// level-38 character's target ranges 38..41, not 60..63, and 60 is out
// of range for it even though it was always in range for a MaxLevel
// character.
func TestEncounterTargetLevelBoundsFollowCharacterLevel(t *testing.T) {
	req := runReq()
	req.Character.Level = 38
	for _, level := range []int{38, 39, 40, 41} {
		req.Encounter.TargetLevel = level
		if err := req.Validate(); err != nil {
			t.Errorf("target_level %d for a level-38 character was refused: %v", level, err)
		}
	}
	for _, level := range []int{37, 42, 60} {
		req.Encounter.TargetLevel = level
		if err := req.Validate(); err == nil || !strings.Contains(err.Error(), "target_level") {
			t.Errorf("target_level %d for a level-38 character should have been refused, got %v", level, err)
		}
	}
}

// ptr is a *int literal for a table test: Go has no address-of operator on
// a literal, and TargetArmor's whole point (Defect 3 below) is that nil
// and a pointer at 0 must be two different requests.
func ptr(v int) *int { return &v }

// A nil override means the level's preset. An override pointing at 0 means
// the target has NO armor - a real request a player can make (2026-09-21
// result-page review, Defect 3: before TargetArmor became a pointer, 0 was
// the only spelling of "unset" a plain int had, so a typed 0 and an absent
// field were the same request and ran identically).
func TestTargetArmorFor(t *testing.T) {
	cases := []struct {
		name     string
		level    int
		override *int
		want     int
	}{
		{"nil override, boss level", BossLevel, nil, TargetArmorByLevel[BossLevel]},
		{"nil override, level 60", 60, nil, TargetArmorByLevel[60]},
		{"nil override, unset level", 0, nil, TargetArmorByLevel[BossLevel]}, // the default boss
		{"nil override, a level with no preset above the cap", 99, nil, TargetArmorByLevel[BossLevel]},
		// Below 60 there is no exact preset row, so TargetArmorFor scales
		// TargetArmorByLevel[MaxLevel] (3,300) proportionally to level:
		// a level-41 target's default is round(3300*41/60) = 2,255, and
		// a level-38 hunter's default target (level 41, DefaultTargetLevel)
		// resolves the same way.
		{"nil override, a level below the cap", 41, nil, 2255},
		{"nil override, level 1", 1, nil, 55},
		{"a positive override", BossLevel, ptr(2500), 2500},
		{"an explicit zero override is zero, not the preset", BossLevel, ptr(0), 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := TargetArmorFor(c.level, c.override); got != c.want {
				t.Errorf("TargetArmorFor(%d, %v) = %d, want %d", c.level, c.override, got, c.want)
			}
		})
	}
	for level := MaxLevel; level <= BossLevel; level++ {
		if TargetArmorByLevel[level] <= 0 {
			t.Errorf("no armor preset for target level %d", level)
		}
	}
}
