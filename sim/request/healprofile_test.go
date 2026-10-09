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
