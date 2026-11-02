package request

import (
	"errors"
	"math"
	"os"
	"regexp"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

// Owner ruling 1, sim-accuracy round 10: below 60 the heal profile is scaled by band.
// Each band's damage comes from that band's content, by one rule stated in the curated
// file (bands.rule) and checked here.

// The curated bands are the leveling bands, the reference (60) among them.
func TestCuratedHealBandsAreTheLevelingBands(t *testing.T) {
	profile := mustCuratedHealProfile(t)
	var levels []int
	for _, row := range profile.Bands.Rows {
		levels = append(levels, row.Level)
	}
	want := []int{20, 30, 40, 50, 60}
	if len(levels) != len(want) {
		t.Fatalf("band levels = %v, want %v", levels, want)
	}
	for i := range want {
		if levels[i] != want[i] {
			t.Fatalf("band levels = %v, want %v", levels, want)
		}
	}
	if profile.Bands.ReferenceLevel != 60 {
		t.Errorf("reference level = %d, want 60", profile.Bands.ReferenceLevel)
	}
}

// The rule's first half: a band's content is the open dungeons (at 60, the open raid)
// whose last boss stands from the band's level to five above it.
func TestCuratedHealBandsFollowTheirContentRule(t *testing.T) {
	profile := mustCuratedHealProfile(t)
	for _, row := range profile.Bands.Rows {
		if len(row.Content) == 0 {
			t.Errorf("band %d names no content", row.Level)
		}
		for _, c := range row.Content {
			if c.BossLevel < row.Level || c.BossLevel > row.Level+profile.Bands.WindowLevels {
				t.Errorf("band %d: %s's last boss %s is level %d, outside %d to %d", row.Level, c.Name, c.LastBoss, c.BossLevel, row.Level, row.Level+profile.Bands.WindowLevels)
			}
		}
	}
}

// A last boss's melee damage a second is its creature_template hit over its attack time:
// Edwin VanCleef (74 attack power, 23 to 30 damage, x2.45, 1.5 s) lands about 70 a hit.
func TestHealBandContentMeleeDPS(t *testing.T) {
	vancleef := HealBandContent{MinMeleeDmg: 23, MaxMeleeDmg: 30, DamageMultiplier: 2.45, MeleeAttackPower: 74, AttackTimeMS: 1500}
	if got, want := vancleef.meleeDPS(), (74.0/14+26.5*2.45)/1.5; math.Abs(got-want) > 1e-9 {
		t.Errorf("VanCleef melee dps = %v, want %v", got, want)
	}
	band := HealBand{ContentDamageMod: 1.2, Content: []HealBandContent{
		{MinMeleeDmg: 10, MaxMeleeDmg: 10, DamageMultiplier: 1, AttackTimeMS: 1000},
		{MinMeleeDmg: 30, MaxMeleeDmg: 30, DamageMultiplier: 1, AttackTimeMS: 1000},
	}}
	if got := band.damage(); math.Abs(got-24) > 1e-9 {
		t.Errorf("band damage = %v, want the median 20 times 1.2", got)
	}
}

// The reference band is the profile as tuned: its figures are the top-level ones.
func TestHealProfileAtTheReferenceBandIsUnchanged(t *testing.T) {
	profile := mustCuratedHealProfile(t)
	for _, level := range []int{60, 0} {
		got := profile.ForLevel(level).Model()
		want := &proto.RaidDamageModel{
			Profile: profile.ID, TankHealth: 9500, MemberHealth: 5000, TankHitDamage: 2128,
			TankSwingSeconds: 3.3, DamageSpread: 0.25, PulseDamage: 555, PulseIntervalSeconds: 5,
			PulseMembers: 3, PulsesReachHealer: true,
		}
		if got.TankHealth != want.TankHealth || got.MemberHealth != want.MemberHealth || got.TankHitDamage != want.TankHitDamage ||
			got.PulseDamage != want.PulseDamage || got.TankSwingSeconds != want.TankSwingSeconds || got.PulseIntervalSeconds != want.PulseIntervalSeconds ||
			got.PulseMembers != want.PulseMembers || got.DamageSpread != want.DamageSpread || !got.PulsesReachHealer {
			t.Errorf("level %d: model %+v, want the level-60 profile %+v", level, got, want)
		}
	}
}

// The rule's second half: the tank hit and the pulse scale with the content's damage
// (its last bosses' median melee damage a second, times the client's dungeon modifier),
// and health with the client's expected player health at the band's level, each against
// the reference band.
func TestHealProfileScalesEachBandByItsContent(t *testing.T) {
	profile := mustCuratedHealProfile(t)
	cases := []struct {
		level                                int
		tankHealth, memberHealth, hit, pulse float64
	}{
		{20, 1836, 966, 124, 32},
		{29, 1836, 966, 124, 32},
		{30, 3204, 1687, 124, 32},
		{40, 4938, 2599, 294, 77},
		{50, 7036, 3703, 677, 177},
		{59, 7036, 3703, 677, 177},
		{10, 1836, 966, 124, 32},
	}
	for _, c := range cases {
		m := profile.ForLevel(c.level).Model()
		if m.TankHealth != c.tankHealth || m.MemberHealth != c.memberHealth || m.TankHitDamage != c.hit || m.PulseDamage != c.pulse {
			t.Errorf("level %d: tank %v hit %v, members %v pulse %v; want %v / %v, %v / %v",
				c.level, m.TankHealth, m.TankHitDamage, m.MemberHealth, m.PulseDamage, c.tankHealth, c.hit, c.memberHealth, c.pulse)
		}
		if m.TankSwingSeconds != profile.Tank.SwingSeconds || m.PulseIntervalSeconds != profile.Pulse.IntervalSeconds || m.PulseMembers != int32(profile.Pulse.Members) {
			t.Errorf("level %d: the cadence moved: %+v", c.level, m)
		}
	}
}

// A tuning bound rescales a band the same way it rescales the reference.
func TestHealProfileAtScaleKeepsTheBand(t *testing.T) {
	profile := mustCuratedHealProfile(t)
	low := profile.AtScale(1.65)
	if m := low.ForLevel(50).Model(); m.TankHitDamage != 604 || m.PulseDamage != 158 {
		t.Errorf("band 50 at 1.65: hit %v pulse %v, want 604 / 158", m.TankHitDamage, m.PulseDamage)
	}
	if m := profile.ForLevel(50).AtScale(1.65).Model(); m.TankHitDamage != 604 || m.PulseDamage != 158 {
		t.Errorf("band 50 then 1.65: hit %v pulse %v, want 604 / 158", m.TankHitDamage, m.PulseDamage)
	}
	if m := low.Model(); m.TankHitDamage != 1898 || m.PulseDamage != 495 {
		t.Errorf("reference at 1.65: hit %v pulse %v, want 1898 / 495", m.TankHitDamage, m.PulseDamage)
	}
}

// ForLevel re-resolves from the curated figures, so resolving twice is resolving once.
func TestHealProfileForLevelIsIdempotent(t *testing.T) {
	profile := mustCuratedHealProfile(t)
	once := profile.ForLevel(30).Model()
	twice := profile.ForLevel(20).ForLevel(30).Model()
	if once.TankHitDamage != twice.TankHitDamage || once.TankHealth != twice.TankHealth || once.PulseDamage != twice.PulseDamage {
		t.Errorf("ForLevel(20).ForLevel(30) = %+v, want ForLevel(30) = %+v", twice, once)
	}
}

// Attach and AttachWeights read the healer's own level from the request, so every tool
// that runs a healer (/sim, the API, the ranker, spec-breakdown, heal-builds,
// dps-per-point) fights its band's profile without being told.
func TestHealProfileAttachUsesTheHealersBand(t *testing.T) {
	profile := mustCuratedHealProfile(t)
	req := &proto.RaidSimRequest{
		Raid:      &proto.Raid{Parties: []*proto.Party{{Players: []*proto.Player{{Level: 40}}}}},
		Encounter: &proto.Encounter{Duration: 180},
	}
	profile.Attach(req)
	if m := req.Raid.RaidDamageModel; m.TankHitDamage != 294 || m.TankHealth != 4938 {
		t.Errorf("a level-40 healer's raid: hit %v on %v health, want 294 on 4938", m.TankHitDamage, m.TankHealth)
	}
	weights := &proto.StatWeightsRequest{Player: &proto.Player{Level: 20}, Encounter: &proto.Encounter{Duration: 180}}
	profile.AttachWeights(weights)
	if m := weights.RaidDamageModel; m.TankHitDamage != 124 || m.PulseDamage != 32 {
		t.Errorf("a level-20 weights request: hit %v pulse %v, want 124 / 32", m.TankHitDamage, m.PulseDamage)
	}
}

// Each band row states its derived figures, and loading refuses a row whose figures are
// not the rule's (the way tuning.scale is held to tank.hit_damage).
func TestHealProfileRefusesABandFigureThatIsNotTheRules(t *testing.T) {
	raw, err := os.ReadFile(curatedHealProfile)
	if err != nil {
		t.Fatal(err)
	}
	field := regexp.MustCompile(`"tank_hit_damage": 294`)
	if !field.Match(raw) {
		t.Fatal("the curated band 40 no longer states tank_hit_damage 294; update this test")
	}
	if _, err := parseHealProfile(field.ReplaceAll(raw, []byte(`"tank_hit_damage": 295`)), "broken"); !errors.Is(err, ErrBadHealProfile) {
		t.Errorf("err = %v, want ErrBadHealProfile", err)
	}
}

func mustCuratedHealProfile(t *testing.T) HealProfile {
	t.Helper()
	profile, err := LoadHealProfile(curatedHealProfile)
	if err != nil {
		t.Fatalf("the curated profile does not load: %v", err)
	}
	return profile
}
