package request

import (
	"math"
	"os"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
)

const curatedTankEncounter = "../../data/curated/tank-encounter.json"

func TestTankEncounterCopyMatchesCurated(t *testing.T) {
	curated, err := os.ReadFile(filepath.FromSlash(curatedTankEncounter))
	if err != nil {
		t.Fatalf("reading the curated tank encounter: %v", err)
	}
	if string(curated) != string(tankEncounterJSON) {
		t.Fatal("sim/request/tank-encounter.json differs from data/curated/tank-encounter.json; run `make tank-encounter-sync`")
	}
}

func TestEveryTankEncounterValueStatesItsReason(t *testing.T) {
	f, err := loadTankEncounter()
	if err != nil {
		t.Fatal(err)
	}
	reasons := map[string]string{
		"level_scaling.exponent":        f.LevelScaling.Exponent.Reason,
		"boss.swing_speed_sec":          f.Boss.SwingSpeedSec.Reason,
		"boss.min_base_damage":          f.Boss.MinBaseDamage.Reason,
		"boss.damage_spread":            f.Boss.DamageSpread.Reason,
		"boss.parry_haste":              f.Boss.ParryHaste.Reason,
		"boss.dual_wield":               f.Boss.DualWield.Reason,
		"healers.hps":                   f.Healers.HPS.Reason,
		"healers.cadence_sec":           f.Healers.CadenceSec.Reason,
		"healers.cadence_variation_sec": f.Healers.CadenceVariationSec.Reason,
		"healers.burst_window_sec":      f.Healers.BurstWindowSec.Reason,
	}
	for name, reason := range reasons {
		if reason == "" {
			t.Errorf("%s states no reason", name)
		}
	}
}

func TestTankProfileScalesDamageAndHealingWithLevel(t *testing.T) {
	full, err := TankProfileForLevel(60)
	if err != nil {
		t.Fatal(err)
	}
	half, err := TankProfileForLevel(30)
	if err != nil {
		t.Fatal(err)
	}
	f, _ := loadTankEncounter()
	ratio := math.Pow(0.5, f.LevelScaling.Exponent.Value)
	healingRatio := math.Pow(0.5, f.LevelScaling.HealingExponent.Value)
	if got, want := half.Boss.MinBaseDamage, full.Boss.MinBaseDamage*ratio; math.Abs(got-want) > 1e-9 {
		t.Errorf("level 30 boss damage = %v, want level 60's scaled by %v (%v)", got, ratio, want)
	}
	if got, want := half.Healers.HPS, full.Healers.HPS*healingRatio; math.Abs(got-want) > 1e-9 {
		t.Errorf("level 30 healing = %v, want level 60's scaled by %v (%v)", got, healingRatio, want)
	}
	if half.Boss.SwingSpeedSec != full.Boss.SwingSpeedSec || half.Boss.DamageSpread != full.Boss.DamageSpread {
		t.Error("level scaling must change damage and healing only, not the swing timer or the roll")
	}
}

func TestRawDPSIsTheMeanSwingOverItsSpeed(t *testing.T) {
	boss := TankBoss{SwingSpeedSec: 2, MinBaseDamage: 2000, DamageSpread: 0.5}
	if got, want := boss.RawDPS(), 1250.0; got != want {
		t.Errorf("RawDPS = %v, want %v", got, want)
	}
}

func TestIsTankSpecFollowsTheCanonicalRole(t *testing.T) {
	for spec, want := range map[string]bool{
		"warrior-protection": true,
		"paladin-protection": true,
		"druid-feral-bear":   true,
		"warrior-fury":       false,
		"druid-feral":        false,
		"no-such-spec":       false,
	} {
		if got := IsTankSpec(spec); got != want {
			t.Errorf("IsTankSpec(%q) = %v, want %v", spec, got, want)
		}
	}
}

func tankRequest(spec, class string, level int) api.SimRequest {
	return api.SimRequest{
		Spec:      spec,
		Source:    api.CharacterSource{Kind: api.SourceBuild},
		Character: api.CharacterSpec{Name: "t", Race: "dwarf", Class: class, Level: level},
		Encounter: api.DefaultEncounter(),
	}
}

func TestBuildPutsATankInFrontOfAHealedBossThatHitsIt(t *testing.T) {
	req := tankRequest("warrior-protection", "warrior", 60)
	player := &proto.Player{}
	raid := &proto.Raid{}
	enc := encounter(req.Encounter, 60)
	if err := applyTankFight(player, raid, enc, req); err != nil {
		t.Fatal(err)
	}
	profile, err := TankProfileForLevel(60)
	if err != nil {
		t.Fatal(err)
	}
	if !player.InFrontOfTarget {
		t.Error("a tank stands in front of its boss")
	}
	if player.HealingModel == nil || player.HealingModel.Hps != profile.Healers.HPS || player.HealingModel.BurstWindow != profile.Healers.BurstWindowSec {
		t.Errorf("healing model = %+v, want the curated healers", player.HealingModel)
	}
	if len(raid.Tanks) != 1 || raid.Tanks[0].Index != tankIndex {
		t.Errorf("raid tanks = %v, want the character", raid.Tanks)
	}
	boss := enc.Targets[0]
	if boss.TankIndex != tankIndex || boss.SwingSpeed != profile.Boss.SwingSpeedSec || boss.MinBaseDamage != profile.Boss.MinBaseDamage {
		t.Errorf("boss = %+v, want it swinging at the tank with the curated melee", boss)
	}
}

func TestATrainingDummyKeepsItsPeacefulTarget(t *testing.T) {
	req := tankRequest("warrior-protection", "warrior", 60)
	req.Encounter.Dummy = true
	player := &proto.Player{}
	raid := &proto.Raid{}
	enc := encounter(req.Encounter, 60)
	if err := applyTankFight(player, raid, enc, req); err != nil {
		t.Fatal(err)
	}
	if enc.Targets[0].SwingSpeed != 0 || len(raid.Tanks) != 0 || player.HealingModel != nil {
		t.Error("a dummy does not fight back, so nothing of the tank fight applies")
	}
}
