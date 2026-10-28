package coach

import "testing"

func packRank1(fight int, code Code, sev Severity, title string) NightFight {
	return NightFight{FightIndex: fight, Pack: Pack{Status: StatusReady, Findings: []Finding{
		{ID: "f1", Code: code, Severity: sev, Rank: 1, Title: title},
	}}}
}

func TestNight_SeverityThenMostBossesThenEarliest(t *testing.T) {
	packs := []NightFight{
		packRank1(1, CodeCoreCasts, SeverityMajor, "Bloodthirst: 9 casts lost"),
		packRank1(2, CodePreparation, SeverityMajor, "No flask at the pull"),
		packRank1(3, CodeCoreCasts, SeverityMajor, "Bloodthirst: 4 casts lost"),
		packRank1(4, CodeResource, SeverityMinor, "Rage full for 14 seconds"),
	}
	l, ok := Night(packs)
	if !ok || l.Code != CodeCoreCasts || l.Stem != "Bloodthirst casts lost" || l.On != 2 || l.Of != 4 {
		t.Fatalf("%+v", l)
	}
	if l.FightIndex != 1 || l.FindingCount != 4 {
		t.Fatalf("earliest boss 1, four findings: %+v", l)
	}
}

func TestNight_TieGoesToTheEarliestBoss(t *testing.T) {
	packs := []NightFight{
		packRank1(1, CodeResource, SeverityMajor, "Rage full for 14 seconds"),
		packRank1(2, CodeCoreCasts, SeverityMajor, "Bloodthirst: 4 casts lost"),
	}
	if l, ok := Night(packs); !ok || l.Code != CodeResource || l.FightIndex != 1 || l.On != 1 || l.Of != 2 {
		t.Fatalf("same severity, one boss each: the earliest boss wins, got %+v", l)
	}
}

func TestNight_MajorBeatsMoreMinorBosses(t *testing.T) {
	packs := []NightFight{
		packRank1(2, CodeIdle, SeverityMinor, "Idle for 9 seconds"),
		packRank1(5, CodeIdle, SeverityMinor, "Idle for 4 seconds"),
		packRank1(7, CodeDeath, SeverityMajor, "Died at 1:02 to Cleave"),
	}
	if l, ok := Night(packs); !ok || l.Code != CodeDeath || l.FightIndex != 7 || l.Stem != StemDeath {
		t.Fatalf("severity first: %+v", l)
	}
}

func TestNight_NoFindingsNoLine(t *testing.T) {
	if l, ok := Night([]NightFight{{Pack: Pack{Status: StatusReady}}, {Pack: Pack{Status: StatusReady}}}); ok {
		t.Fatalf("no findings, no line: %+v", l)
	}
}

func TestNight_OnlyReadyPacksCount(t *testing.T) {
	packs := []NightFight{
		packRank1(1, CodeIdle, SeverityMinor, "Idle for 9 seconds"),
		{FightIndex: 2, Pack: Pack{Status: StatusNotReviewable}},
	}
	if l, _ := Night(packs); l.Of != 1 {
		t.Fatalf("a pack that is not ready is no boss of the night: %+v", l)
	}
}

func TestStem_EveryTitleForm(t *testing.T) {
	cases := []struct {
		code        Code
		title, want string
	}{
		{CodeCoreCasts, "Bloodthirst: 9 casts lost", "Bloodthirst casts lost"},
		{CodeUptime, "Sunder Armor up 41% of the fight", "Sunder Armor uptime"},
		{CodeIdle, "Idle for 12 seconds", StemIdle},
		{CodeResource, "Rage full for 14 seconds", "Rage at the cap"},
		{CodeResource, "Out of mana from 2:10", StemOutOfMana},
		{CodeRanks, "Heal (Rank 2): 61% overheal", "Heal overheal"},
		{CodeRanks, "Frostbolt: 14 casts of Rank 3", "Frostbolt below your best rank"},
		{CodePreparation, TitleNoFlask, StemNoFlask},
		{CodePreparation, TitleNoFood, StemNoFood},
		{CodePreparation, "Main Hand has no enchant", StemEnchants},
		{CodePreparation, "3 items have no enchant", StemEnchants},
		{CodePreparation, TitleNoWeaponOil, StemEnchants},
		{CodePreparation, "Combat potions: 2 unused", StemPotions},
		{CodePreparation, "No Arcane Intellect at the pull", StemRaidBuffs},
		{CodePreparation, "2 raid buffs missing", StemRaidBuffs},
		{CodeDeath, "Died at 1:02 with Shield Wall ready", StemDeath},
		{CodeAdherence, "71% of casts on the rotation", StemAdherence},
		{CodeAdherence + "x", "Something new", "Something new"},
	}
	for _, c := range cases {
		if got := Stem(Finding{Code: c.code, Title: c.title}); got != c.want {
			t.Errorf("Stem(%s, %q) = %q, want %q", c.code, c.title, got, c.want)
		}
	}
}
