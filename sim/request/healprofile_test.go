package request

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

const curatedHealProfile = "../../data/curated/heal-profile.json"

func TestCuratedHealProfileLoads(t *testing.T) {
	profile, err := LoadHealProfile(curatedHealProfile)
	if err != nil {
		t.Fatalf("the curated profile does not load: %v", err)
	}
	model := profile.Model()
	if model.Profile != profile.ID {
		t.Errorf("the model names itself %q, want the profile's id %q", model.Profile, profile.ID)
	}
	if model.TankHitDamage != profile.Tank.HitDamage || model.PulseMembers != int32(profile.Pulse.Members) {
		t.Errorf("the model does not carry the profile's figures: %+v", model)
	}
}

func TestHealProfileAttachLaysOutTheFakeRaid(t *testing.T) {
	profile, err := LoadHealProfile(curatedHealProfile)
	if err != nil {
		t.Fatal(err)
	}
	req := &proto.RaidSimRequest{
		Raid:      &proto.Raid{Parties: []*proto.Party{{Players: []*proto.Player{{}}}}},
		Encounter: &proto.Encounter{Duration: 180, DurationVariation: 36},
	}
	profile.Attach(req)

	if got := len(req.Raid.Parties); got != 2 {
		t.Errorf("parties = %d, want 2 (the tank stands in the second)", got)
	}
	if req.Raid.TargetDummies != 5 {
		t.Errorf("fake raid members = %d, want 5", req.Raid.TargetDummies)
	}
	if req.Raid.RaidDamageModel == nil {
		t.Error("the raid carries no damage model")
	}
	if req.Encounter.Duration != float64(profile.DurationSec) || req.Encounter.DurationVariation != 0 {
		t.Errorf("fight = %v +/- %v, want the profile's %d seconds exactly", req.Encounter.Duration, req.Encounter.DurationVariation, profile.DurationSec)
	}
}

func TestHealProfileAttachWeightsCarriesTheModel(t *testing.T) {
	profile, err := LoadHealProfile(curatedHealProfile)
	if err != nil {
		t.Fatal(err)
	}
	req := &proto.StatWeightsRequest{Encounter: &proto.Encounter{Duration: 180}}
	profile.AttachWeights(req)
	if req.RaidDamageModel == nil || req.Encounter.Duration != float64(profile.DurationSec) {
		t.Errorf("weights request = %+v, want the damage model and the profile's fight length", req)
	}
}

func TestHealProfileRefusesWhatTheEngineCannotRun(t *testing.T) {
	cases := map[string]struct{ key, broken string }{
		"no tank swings":      {"swing_seconds", "0"},
		"no pulse interval":   {"interval_seconds", "0"},
		"spread of a hundred": {"damage_spread", "1"},
	}
	good, err := os.ReadFile(curatedHealProfile)
	if err != nil {
		t.Fatal(err)
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			field := regexp.MustCompile(`"` + c.key + `":\s*[0-9.]+`)
			if !field.Match(good) {
				t.Fatalf("the curated profile no longer states %q as a number; update this test", c.key)
			}
			path := filepath.Join(t.TempDir(), "profile.json")
			broken := field.ReplaceAll(good, []byte(`"`+c.key+`": `+c.broken))
			if err := os.WriteFile(path, broken, 0o644); err != nil {
				t.Fatal(err)
			}
			if _, err := LoadHealProfile(path); !errors.Is(err, ErrBadHealProfile) {
				t.Errorf("err = %v, want ErrBadHealProfile", err)
			}
		})
	}
}

func TestHealProfileRefusesAnUnknownField(t *testing.T) {
	path := filepath.Join(t.TempDir(), "profile.json")
	if err := os.WriteFile(path, []byte(`{"id": "x", "surprise": 1}`), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadHealProfile(path); err == nil {
		t.Error("an unknown field loaded without error")
	}
}

func TestEmbeddedHealProfileMatchesCurated(t *testing.T) {
	curated, err := os.ReadFile(filepath.FromSlash(curatedHealProfile))
	if err != nil {
		t.Fatalf("reading the curated heal profile: %v", err)
	}
	if string(curated) != string(healProfileJSON) {
		t.Fatal("sim/request/heal-profile.json differs from data/curated/heal-profile.json; run `make heal-profile-sync`")
	}
	embedded, err := EmbeddedHealProfile()
	if err != nil {
		t.Fatalf("the embedded profile does not load: %v", err)
	}
	loaded, err := LoadHealProfile(curatedHealProfile)
	if err != nil {
		t.Fatal(err)
	}
	if embedded.ID != loaded.ID || embedded.DurationSec != loaded.DurationSec {
		t.Errorf("embedded %q/%ds, curated %q/%ds", embedded.ID, embedded.DurationSec, loaded.ID, loaded.DurationSec)
	}
}

// The profile's tuning is configuration (owner ruling, round 5): the tuned scale and the
// bounds it was tuned between, with the unscaled figures the scale multiplies.
func TestCuratedHealProfileStatesItsTuning(t *testing.T) {
	profile, err := LoadHealProfile(curatedHealProfile)
	if err != nil {
		t.Fatal(err)
	}
	tuning := profile.Tuning
	if tuning.Scale != 1.65 || len(tuning.Bounds) != 2 || tuning.Bounds[0] != 1.5 || tuning.Bounds[1] != 1.65 {
		t.Errorf("tuning = %+v, want scale 1.65 between 1.5 and 1.65", tuning)
	}
	if profile.Tank.HitDamage != 1898 || profile.Pulse.Damage != 495 {
		t.Errorf("tuned figures = %v / %v, want 1898 / 495", profile.Tank.HitDamage, profile.Pulse.Damage)
	}
}

func TestHealProfileAtScaleRescalesTheHitAndThePulse(t *testing.T) {
	profile, err := LoadHealProfile(curatedHealProfile)
	if err != nil {
		t.Fatal(err)
	}
	low := profile.AtScale(1.5)
	if low.Tank.HitDamage != 1725 || low.Pulse.Damage != 450 || low.Tuning.Scale != 1.5 {
		t.Errorf("at 1.5: hit %v pulse %v scale %v, want 1725 / 450 / 1.5", low.Tank.HitDamage, low.Pulse.Damage, low.Tuning.Scale)
	}
	if profile.Tank.HitDamage != 1898 {
		t.Error("AtScale changed the profile it was called on")
	}
	same := profile.AtScale(profile.Tuning.Scale)
	if same.Tank.HitDamage != profile.Tank.HitDamage || same.Pulse.Damage != profile.Pulse.Damage {
		t.Errorf("at the tuned scale: %v / %v", same.Tank.HitDamage, same.Pulse.Damage)
	}
}

func TestHealProfileRefusesATuningThatDisagreesWithItsFigures(t *testing.T) {
	raw, err := os.ReadFile(curatedHealProfile)
	if err != nil {
		t.Fatal(err)
	}
	for name, bad := range map[string]*regexp.Regexp{
		"scale outside bounds": regexp.MustCompile(`"scale": 1\.65`),
		"hit not base x scale": regexp.MustCompile(`"hit_damage": 1898`),
	} {
		replacement := map[string]string{"scale outside bounds": `"scale": 1.8`, "hit not base x scale": `"hit_damage": 1900`}[name]
		if _, err := parseHealProfile(bad.ReplaceAll(raw, []byte(replacement)), name); !errors.Is(err, ErrBadHealProfile) {
			t.Errorf("%s: err = %v, want ErrBadHealProfile", name, err)
		}
	}
}

// Raid-wide pulses land on the healer standing in the group as well as on
// the other members (round 6 healers): the curated profile says so and the
// engine model carries it, or every party heal puts a fifth of its healing
// on a healer who is never hurt.
func TestCuratedHealProfilePulsesReachTheHealer(t *testing.T) {
	profile, err := LoadHealProfile(curatedHealProfile)
	if err != nil {
		t.Fatal(err)
	}
	if !profile.Pulse.ReachesHealer {
		t.Error("the curated pulse does not reach the healer")
	}
	if !profile.Model().PulsesReachHealer {
		t.Error("the engine model does not carry pulse.reaches_healer")
	}
}
