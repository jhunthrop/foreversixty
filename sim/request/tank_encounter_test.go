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
		"boss.attack_power_by_level":    f.Boss.AttackPowerByLevel.Reason,
		"boss.damage_spread":            f.Boss.DamageSpread.Reason,
		"boss.parry_haste":              f.Boss.ParryHaste.Reason,
		"boss.dual_wield":               f.Boss.DualWield.Reason,
		"healers.hps":                   f.Healers.HPS.Reason,
		"healers.cadence_sec":           f.Healers.CadenceSec.Reason,
		"healers.cadence_variation_sec": f.Healers.CadenceVariationSec.Reason,
		"healers.burst_window_sec":      f.Healers.BurstWindowSec.Reason,
		"dungeon_fight.duration_sec":    f.DungeonFight.DurationSec.Reason,
	}
	for name, reason := range reasons {
		if reason == "" {
			t.Errorf("%s states no reason", name)
		}
	}
}

// Sim-accuracy round 30 (walks-heal T2): below 60 the boss was the level-60 one scaled by
// (band/60)^2 and the healing by (band/60)^1.5, which hit 1.4 to 2 times harder than the
// dungeon table the heal profile carries (data/curated/heal-profile.json bands: each band's
// sourced tank hit, from its dungeons' last bosses and trash pulls). A band's boss now scales
// by that table, the band's tank hit over the reference band's, and the healing keeps the old
// stand-in's ratio to the damage ((band/60)^(1.5-2)), so only the swings against the tank's
// health shrink. Below the first band a level scales down from band 20 by the exponent.
func TestTankProfileScalesDamageAndHealingByTheHealProfileBandTable(t *testing.T) {
	f, _ := loadTankEncounter()
	healingOverDamage := func(level int) float64 {
		return math.Pow(float64(level)/60, f.LevelScaling.HealingExponent.Value-f.LevelScaling.Exponent.Value)
	}
	full, err := TankProfileForLevel(60)
	if err != nil {
		t.Fatal(err)
	}
	heal, err := EmbeddedHealProfile()
	if err != nil {
		t.Fatal(err)
	}
	reference := heal.Bands.reference().TankHitDamage
	want := map[int]float64{}
	for _, row := range heal.Bands.Rows {
		want[row.Level] = row.TankHitDamage / reference
	}
	for _, c := range []struct{ level, band int }{{20, 20}, {29, 20}, {30, 30}, {35, 30}, {40, 40}, {50, 50}, {59, 50}, {60, 60}} {
		got, err := TankProfileForLevel(c.level)
		if err != nil {
			t.Fatal(err)
		}
		scale := want[c.band]
		if scale <= 0 || (c.band < 60 && scale >= 1) {
			t.Fatalf("band %d: heal profile tank hit ratio %v, want in (0, 1)", c.band, scale)
		}
		if g, w := got.Boss.MinBaseDamage, full.Boss.MinBaseDamage*scale; math.Abs(g-w) > 1e-9 {
			t.Errorf("level %d boss damage = %v, want level 60's times band %d's tank hit ratio %v (%v)", c.level, g, c.band, scale, w)
		}
		if g, w := got.Healers.HPS, full.Healers.HPS*scale*healingOverDamage(c.band); math.Abs(g-w) > 1e-9 {
			t.Errorf("level %d healing = %v, want the damage scale %v times the stand-in's healing-to-damage ratio at band %d (%v)", c.level, g, scale, c.band, w)
		}
		if got.Boss.SwingSpeedSec != full.Boss.SwingSpeedSec || got.Boss.DamageSpread != full.Boss.DamageSpread {
			t.Error("level scaling must change damage and healing only, not the swing timer or the roll")
		}
	}
	// The table at 30 and 40 is about half of the old (band/60)^2 stand-in: the gap T2 measured.
	for _, band := range []int{30, 40} {
		old := math.Pow(float64(band)/60, 2)
		if ratio := want[band] / old; ratio > 0.6 {
			t.Errorf("band %d tank hit ratio %v is %v of the old stand-in; T2 measured about half", band, want[band], ratio)
		}
	}
	ten, err := TankProfileForLevel(10)
	if err != nil {
		t.Fatal(err)
	}
	if g, w := ten.Boss.MinBaseDamage, full.Boss.MinBaseDamage*want[20]*math.Pow(0.5, f.LevelScaling.Exponent.Value); math.Abs(g-w) > 1e-9 {
		t.Errorf("level 10 boss damage = %v, want band 20's scaled by (10/20)^exponent (%v)", g, w)
	}
	if g, w := ten.Healers.HPS, full.Healers.HPS*want[20]*math.Pow(0.5, f.LevelScaling.Exponent.Value)*healingOverDamage(10); math.Abs(g-w) > 1e-9 {
		t.Errorf("level 10 healing = %v, want its damage scale times the healing-to-damage ratio at 10 (%v)", g, w)
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
	if boss.TankIndex != tankIndex || boss.SwingSpeed != profile.Boss.SwingSpeedSec || boss.MinBaseDamage != profile.Boss.weaponMinDamage(boss.Stats[proto.Stat_StatAttackPower]) {
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

// Sim-accuracy 2026-11-11 tanks below 45: below 60 the tank fight lasted the ranker's default
// 180 seconds, a raid boss's length, against a dungeon band's hardest pull. Every logged Forever
// dungeon fight on foreverlogs.gg is far shorter (bosses and trash pulls at 20 to 45 last about
// half a minute), so a tank whose damage taken sat a little above the band's healing bled out
// over three minutes it never fights: a bare bear at 30 to 35 died in 96% of fights. Below the
// reference level the fight is now a dungeon pull of the curated length; level 60 keeps the
// request's raid-length fight.
func TestTankFightBelowTheReferenceLevelIsADungeonPull(t *testing.T) {
	f, err := loadTankEncounter()
	if err != nil {
		t.Fatal(err)
	}
	pull := f.DungeonFight.DurationSec.Value
	if pull <= 0 || pull >= float64(api.DefaultEncounter().DurationSec) {
		t.Fatalf("dungeon_fight.duration_sec = %v, want a positive pull shorter than the %ds raid fight", pull, api.DefaultEncounter().DurationSec)
	}
	for _, level := range []int{10, 20, 29, 32, 45, 59} {
		profile, err := TankProfileForLevel(level)
		if err != nil {
			t.Fatal(err)
		}
		if profile.FightSec != pull {
			t.Errorf("level %d fight = %vs, want the dungeon pull %vs", level, profile.FightSec, pull)
		}
		req := tankRequest("druid-feral-bear", "druid", level)
		enc := encounter(req.Encounter, level)
		if err := applyTankFight(&proto.Player{}, &proto.Raid{}, enc, req); err != nil {
			t.Fatal(err)
		}
		wantVariation := pull * req.Encounter.Variation
		if enc.Duration != pull || math.Abs(enc.DurationVariation-wantVariation) > 1e-9 {
			t.Errorf("level %d encounter = %vs +/- %v, want the pull %vs +/- %v (the request's share)", level, enc.Duration, enc.DurationVariation, pull, wantVariation)
		}
	}
	raid, err := TankProfileForLevel(60)
	if err != nil {
		t.Fatal(err)
	}
	if raid.FightSec != 0 {
		t.Errorf("level 60 fight = %vs, want 0: the raid boss keeps the request's length", raid.FightSec)
	}
	req := tankRequest("druid-feral-bear", "druid", 60)
	enc := encounter(req.Encounter, 60)
	if err := applyTankFight(&proto.Player{}, &proto.Raid{}, enc, req); err != nil {
		t.Fatal(err)
	}
	if enc.Duration != float64(req.Encounter.DurationSec) {
		t.Errorf("level 60 encounter = %vs, want the request's %ds", enc.Duration, req.Encounter.DurationSec)
	}
}

func TestFightSecondsIsThePullForATankBelowTheReferenceLevel(t *testing.T) {
	f, _ := loadTankEncounter()
	for _, c := range []struct {
		spec, class string
		level       int
		want        float64
	}{
		{"druid-feral-bear", "druid", 32, f.DungeonFight.DurationSec.Value},
		{"druid-feral-bear", "druid", 60, float64(api.DefaultEncounter().DurationSec)},
		{"warrior-fury", "warrior", 32, float64(api.DefaultEncounter().DurationSec)},
	} {
		got, err := FightSeconds(tankRequest(c.spec, c.class, c.level))
		if err != nil {
			t.Fatal(err)
		}
		if got != c.want {
			t.Errorf("%s at %d fights %vs, want %vs", c.spec, c.level, got, c.want)
		}
	}
}
