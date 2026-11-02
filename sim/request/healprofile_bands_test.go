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
	if got, want := vancleef.damagePerSecond(), (74.0/14+26.5*2.45)/1.5; math.Abs(got-want) > 1e-9 {
		t.Errorf("VanCleef melee dps = %v, want %v", got, want)
	}
	// A boss's spells add their damage per cast over their interval, and the time it
	// spends casting is time it does not swing: 2 s of every 4 s casting halves the melee.
	caster := HealBandContent{MinMeleeDmg: 10, MaxMeleeDmg: 10, DamageMultiplier: 1, AttackTimeMS: 1000,
		Spells: []HealBossSpell{{DamagePerCast: 100, IntervalSeconds: 4, CastSeconds: 2}}}
	if got := caster.damagePerSecond(); math.Abs(got-(10*0.5+25)) > 1e-9 {
		t.Errorf("caster damage a second = %v, want half its melee plus 25", got)
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
		{20, 1836, 966, 143, 37},
		{29, 1836, 966, 143, 37},
		{30, 3204, 1687, 263, 69},
		{40, 4938, 2599, 470, 123},
		{50, 7036, 3703, 1010, 263},
		{59, 7036, 3703, 1010, 263},
		{10, 1836, 966, 143, 37},
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
	if m := low.ForLevel(50).Model(); m.TankHitDamage != 900 || m.PulseDamage != 235 {
		t.Errorf("band 50 at 1.65: hit %v pulse %v, want 900 / 235", m.TankHitDamage, m.PulseDamage)
	}
	if m := profile.ForLevel(50).AtScale(1.65).Model(); m.TankHitDamage != 900 || m.PulseDamage != 235 {
		t.Errorf("band 50 then 1.65: hit %v pulse %v, want 900 / 235", m.TankHitDamage, m.PulseDamage)
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
	if m := req.Raid.RaidDamageModel; m.TankHitDamage != 470 || m.TankHealth != 4938 {
		t.Errorf("a level-40 healer's raid: hit %v on %v health, want 470 on 4938", m.TankHitDamage, m.TankHealth)
	}
	weights := &proto.StatWeightsRequest{Player: &proto.Player{Level: 20}, Encounter: &proto.Encounter{Duration: 180}}
	profile.AttachWeights(weights)
	if m := weights.RaidDamageModel; m.TankHitDamage != 143 || m.PulseDamage != 37 {
		t.Errorf("a level-20 weights request: hit %v pulse %v, want 143 / 37", m.TankHitDamage, m.PulseDamage)
	}
}

// Each band row states its derived figures, and loading refuses a row whose figures are
// not the rule's (the way tuning.scale is held to tank.hit_damage).
func TestHealProfileRefusesABandFigureThatIsNotTheRules(t *testing.T) {
	raw, err := os.ReadFile(curatedHealProfile)
	if err != nil {
		t.Fatal(err)
	}
	field := regexp.MustCompile(`"tank_hit_damage": 470`)
	if !field.Match(raw) {
		t.Fatal("the curated band 40 no longer states tank_hit_damage 470; update this test")
	}
	if _, err := parseHealProfile(field.ReplaceAll(raw, []byte(`"tank_hit_damage": 471`)), "broken"); !errors.Is(err, ErrBadHealProfile) {
		t.Errorf("err = %v, want ErrBadHealProfile", err)
	}
}

// Owner ruling 1, sim-accuracy round 11: a band fights the harder of its last boss and its
// typical hard trash pull. The pull multiplier is the pull's damage a second over the
// median last boss's, never below 1, so a band never fights lighter than its bosses.
func TestHealBandPullMultiplier(t *testing.T) {
	boss := []HealBandContent{{MinMeleeDmg: 20, MaxMeleeDmg: 20, DamageMultiplier: 1, AttackTimeMS: 1000}}
	heavy := HealBand{ContentDamageMod: 1.2, Content: boss, Pulls: &HealBandPulls{PullDamagePerSecond: 50}}
	if got := heavy.pullMultiplier(); math.Abs(got-2.5) > 1e-9 {
		t.Errorf("a 50 a second pull over a 20 a second boss: multiplier %v, want 2.5", got)
	}
	if got := heavy.damage(); math.Abs(got-60) > 1e-9 {
		t.Errorf("band damage = %v, want the pull's 50 times 1.2", got)
	}
	light := HealBand{ContentDamageMod: 1.2, Content: boss, Pulls: &HealBandPulls{PullDamagePerSecond: 10}}
	if got := light.pullMultiplier(); got != 1 {
		t.Errorf("a pull lighter than the boss: multiplier %v, want 1", got)
	}
	none := HealBand{ContentDamageMod: 1, Content: boss}
	if got := none.pullMultiplier(); got != 1 {
		t.Errorf("a band with no pulls (the raid): multiplier %v, want 1", got)
	}
}

// A pull's damage a second is the sum of its mobs': each mob's tank and group damage a
// second, times how many of it the pull holds.
func TestHealBandPullDamageIsItsMobs(t *testing.T) {
	pull := HealPull{Mobs: []HealPullMob{
		{Count: 2, TankDamagePerSecond: 28.5},
		{Count: 1, TankDamagePerSecond: 20, GroupDamagePerSecond: 5},
	}}
	if got := pull.damagePerSecond(); math.Abs(got-82) > 1e-9 {
		t.Errorf("pull damage = %v, want 2 x 28.5 + 25", got)
	}
}

// Every band below the reference states its pulls under the one pull rule (the stated
// assistance radius and percentile), a pull whose mobs sum to its stated damage, and the
// multiplier the rule derives; the reference band (a raid) states none.
func TestCuratedHealBandsStateTheirPulls(t *testing.T) {
	profile := mustCuratedHealProfile(t)
	if profile.Bands.PullRule.AssistRadiusYards <= 0 || profile.Bands.PullRule.Percentile <= 0 || profile.Bands.PullRule.Percentile >= 1 {
		t.Fatalf("the pull rule states radius %v and percentile %v", profile.Bands.PullRule.AssistRadiusYards, profile.Bands.PullRule.Percentile)
	}
	want := map[int]float64{20: 1.61, 30: 2.06, 40: 3.05, 50: 1.85, 60: 1}
	for _, row := range profile.Bands.Rows {
		if row.Level == profile.Bands.ReferenceLevel {
			if row.Pulls != nil {
				t.Errorf("the reference band states pulls")
			}
			continue
		}
		if row.Pulls == nil {
			t.Fatalf("band %d states no pulls", row.Level)
		}
		if got := math.Round(row.pullMultiplier()*100) / 100; got != want[row.Level] {
			t.Errorf("band %d pull multiplier = %v, want %v", row.Level, got, want[row.Level])
		}
		if len(row.Pulls.Pull.CreatureGUIDs) != row.Pulls.PullSize {
			t.Errorf("band %d: the stated pull names %d spawns, its size is %d", row.Level, len(row.Pulls.Pull.CreatureGUIDs), row.Pulls.PullSize)
		}
	}
}

// Loading refuses a band whose stated pull damage is not its mobs' sum, or whose stated
// multiplier is not the rule's.
func TestHealProfileRefusesAPullThatIsNotItsMobs(t *testing.T) {
	raw, err := os.ReadFile(curatedHealProfile)
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct{ from, to string }{
		{`"pull_damage_per_second": 270.1`, `"pull_damage_per_second": 280.1`},
		{`"pull_multiplier": 3.05`, `"pull_multiplier": 3.1`},
	} {
		field := regexp.MustCompile(regexp.QuoteMeta(c.from))
		if !field.Match(raw) {
			t.Fatalf("the curated band 40 no longer states %s; update this test", c.from)
		}
		if _, err := parseHealProfile(field.ReplaceAll(raw, []byte(c.to)), "broken"); !errors.Is(err, ErrBadHealProfile) {
			t.Errorf("%s: err = %v, want ErrBadHealProfile", c.to, err)
		}
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
