package coach

import (
	"slices"
	"testing"
)

func TestThresholdsAreTheSpecValues(t *testing.T) {
	cases := []struct {
		name      string
		got, want float64
	}{
		{"MinFightMS", MinFightMS, 30000}, {"WipeTailMS", WipeTailMS, 15000}, {"MaxFindings", MaxFindings, 8},
		{"MoreAbilitiesInDetail", MoreAbilitiesInDetail, 2},
		{"CoreCooldownMinMS", CoreCooldownMinMS, 4000}, {"CoreCooldownMaxMS", CoreCooldownMaxMS, 600000},
		{"CoreGapGraceMS", CoreGapGraceMS, 1500}, {"CoreFiresLost", CoreFiresLost, 2},
		{"CoreFiresPct", CoreFiresPct, 10}, {"CoreMajorPct", CoreMajorPct, 20},
		{"UptimeFiresPoints", UptimeFiresPoints, 10}, {"UptimeFiresUnderPct", UptimeFiresUnderPct, 90},
		{"UptimeMajorPoints", UptimeMajorPoints, 25},
		{"IdleFiresMS", IdleFiresMS, 10000}, {"IdleFiresPct", IdleFiresPct, 5}, {"IdleMajorPct", IdleMajorPct, 10},
		{"IdleWindowsShown", IdleWindowsShown, 3},
		{"CapFiresMS", CapFiresMS, 10000}, {"CapFiresPct", CapFiresPct, 5}, {"CapMajorPct", CapMajorPct, 15},
		{"ManaLowPct", ManaLowPct, 5}, {"ManaFiresMS", ManaFiresMS, 10000},
		{"ManaFiresBeforePct", ManaFiresBeforePct, 90}, {"ManaMajorBeforePct", ManaMajorBeforePct, 75},
		{"RanksFiresCasts", RanksFiresCasts, 3}, {"RanksFiresPct", RanksFiresPct, 10},
		{"HealRankMinCasts", HealRankMinCasts, 10}, {"HealRankFiresPoints", HealRankFiresPoints, 15},
		{"HealRankMajorPoints", HealRankMajorPoints, 30},
		{"DeathFiresBeforePct", DeathFiresBeforePct, 90}, {"DeathMajorBeforePct", DeathMajorBeforePct, 75},
		{"DeathPotionWindowMS", DeathPotionWindowMS, 10000},
		{"AdherenceFiresPoints", AdherenceFiresPoints, 10}, {"AdherenceFiresUnderPct", AdherenceFiresUnderPct, 85},
		{"AdherenceMajorPoints", AdherenceMajorPoints, 25}, {"OffRotationShown", OffRotationShown, 3},
		{"PoolMinPlayers", PoolMinPlayers, 20}, {"PoolTopN", PoolTopN, 25},
	}
	for _, c := range cases {
		if c.got != c.want {
			t.Errorf("%s = %v, want %v (spec §6)", c.name, c.got, c.want)
		}
	}
}

// §6.3: DPS F7, F1, F3, F6, F2, F4, F8, F5. Healer F7, F4, F5, F1, F6, F2. Tank F7, F2,
// F1, F3, F6, F4, F8, F5.
func TestPrecedenceIsTheSpecOrder(t *testing.T) {
	want := map[string][]Code{
		RoleDPS:    {CodeDeath, CodeCoreCasts, CodeIdle, CodePreparation, CodeUptime, CodeResource, CodeAdherence, CodeRanks},
		RoleHealer: {CodeDeath, CodeResource, CodeRanks, CodeCoreCasts, CodePreparation, CodeUptime},
		RoleTank:   {CodeDeath, CodeUptime, CodeCoreCasts, CodeIdle, CodePreparation, CodeResource, CodeAdherence, CodeRanks},
	}
	if len(Precedence) != len(want) {
		t.Fatalf("Precedence has roles %v, want dps, healer and tank", Precedence)
	}
	for role, order := range want {
		if !slices.Equal(Precedence[role], order) {
			t.Errorf("%s precedence = %v, want %v", role, Precedence[role], order)
		}
	}
	for _, off := range []Code{CodeIdle, CodeAdherence} {
		if slices.Contains(Precedence[RoleHealer], off) {
			t.Errorf("healers have no %s finding (§6.3)", off)
		}
	}
}
