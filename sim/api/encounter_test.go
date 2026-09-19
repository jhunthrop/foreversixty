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
		{"an armor override", func(r *SimRequest) { r.Encounter.TargetArmor = 2500 }, ""},
		{"negative armor", func(r *SimRequest) { r.Encounter.TargetArmor = -1 }, "target_armor"},
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

// Zero armor means the level's preset, not a naked target. Getting that
// backwards would report every sim against an unarmoured boss.
func TestTargetArmorFor(t *testing.T) {
	cases := []struct {
		level, override, want int
	}{
		{BossLevel, 0, TargetArmorByLevel[BossLevel]},
		{60, 0, TargetArmorByLevel[60]},
		{0, 0, TargetArmorByLevel[BossLevel]}, // an unset level is the default boss
		{BossLevel, 2500, 2500},
		{99, 0, TargetArmorByLevel[BossLevel]}, // a level with no preset falls back
	}
	for _, c := range cases {
		if got := TargetArmorFor(c.level, c.override); got != c.want {
			t.Errorf("TargetArmorFor(%d, %d) = %d, want %d", c.level, c.override, got, c.want)
		}
	}
	for level := MinTargetLevel; level <= MaxTargetLevel; level++ {
		if TargetArmorByLevel[level] <= 0 {
			t.Errorf("no armor preset for target level %d", level)
		}
	}
}
