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
// sourced tank hit, from its dungeons' last bosses and trash pulls). A band's boss and its
// healers now scale by that table: the band's tank hit over the reference band's, the same
// factor for both, so the share of the damage the healers cover is the level-60 fight's at
// every band. Below the first band a level scales down from band 20 by the exponents.
func TestTankProfileScalesDamageAndHealingByTheHealProfileBandTable(t *testing.T) {
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
		if g, w := got.Healers.HPS, full.Healers.HPS*scale; math.Abs(g-w) > 1e-9 {
			t.Errorf("level %d healing = %v, want level 60's times band %d's tank hit ratio %v (%v)", c.level, g, c.band, scale, w)
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
	f, _ := loadTankEncounter()
	ten, err := TankProfileForLevel(10)
	if err != nil {
		t.Fatal(err)
	}
	if g, w := ten.Boss.MinBaseDamage, full.Boss.MinBaseDamage*want[20]*math.Pow(0.5, f.LevelScaling.Exponent.Value); math.Abs(g-w) > 1e-9 {
		t.Errorf("level 10 boss damage = %v, want band 20's scaled by (10/20)^exponent (%v)", g, w)
	}
	if g, w := ten.Healers.HPS, full.Healers.HPS*want[20]*math.Pow(0.5, f.LevelScaling.HealingExponent.Value); math.Abs(g-w) > 1e-9 {
		t.Errorf("level 10 healing = %v, want band 20's scaled by (10/20)^healing_exponent (%v)", g, w)
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
