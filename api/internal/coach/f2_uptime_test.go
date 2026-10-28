package coach

import (
	"slices"
	"strings"
	"testing"
)

func TestUptime(t *testing.T) {
	cases := []struct {
		name                     string
		in                       Input
		wantTitle                string
		wantSev                  Severity
		wantCannot, wantStrength string
		none                     bool
	}{
		{name: "tank Shield Block 18% vs 34%", in: protShieldBlock(t, 42_000, 232_000, poolDefensive(2565, 34, 20)),
			wantTitle: "Shield Block up 18% of the fight", wantSev: SeverityMinor},
		{name: "dps self-buff at 100% vs 97% is a strength", in: furyBattleShout(t, 1.0, poolUptime(6673, 97, 25)),
			none: true, wantStrength: "Battle Shout up 100% of the fight. Top 25: 97%."},
		{name: "shared debuff goes to couldn't check", in: furySunderShared(t, 3), none: true,
			wantCannot: "Your share of Sunder Armor: 3 warriors applied it, and the log merges them."},
		{name: "tank with no pool: top logs only", in: protShieldBlock(t, 42_000, 232_000, noPool(0)), none: true},
		{name: "25 points under is major", in: furyRendLow(t, poolUptime(f2RendID, 75, 25)),
			wantTitle: "Rend up 20% of the fight", wantSev: SeverityMajor},
		{name: "under 10 points is not a finding", in: furyBattleShout(t, 0.9, poolUptime(6673, 97, 25)), none: true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			o := uptime(tc.in, CountedTime(tc.in))
			if tc.none != (o.Finding == nil) {
				t.Fatalf("finding %+v", o.Finding)
			}
			if !tc.none && (o.Finding.Title != tc.wantTitle || o.Finding.Severity != tc.wantSev || len(o.Finding.Anchors) == 0) {
				t.Fatalf("got %+v", o.Finding)
			}
			if tc.wantCannot != "" && !slices.Contains(o.CouldNotCheck, tc.wantCannot) {
				t.Errorf("couldn't check %q", o.CouldNotCheck)
			}
			if tc.wantStrength != "" && (len(o.Strengths) == 0 || o.Strengths[0].Sentence != tc.wantStrength) {
				t.Errorf("strengths %+v", o.Strengths)
			}
		})
	}
}

func TestUptime_ShieldBlockFinding(t *testing.T) {
	in := protShieldBlock(t, 42_000, 232_000, poolDefensive(2565, 34, 20))
	f := uptime(in, CountedTime(in)).Finding
	if f == nil {
		t.Fatal("no finding")
	}
	if f.Code != CodeUptime || f.Detail != "Shield Block was up for 42 s of 3:52." {
		t.Errorf("code %q detail %q", f.Code, f.Detail)
	}
	if f.Basis.Kind != BasisTopLogs || f.Basis.N != 20 ||
		f.Basis.Label != "Top 20 Protection warriors on Onyxia (40): 34% up. You: 18%." {
		t.Errorf("basis %+v", f.Basis)
	}
	assertNumbers(t, f.Facts, "18", "34", "42")
	assertSpells(t, f.Facts, "Shield Block")
	assertTimes(t, f.Facts, "3:52")
}

func TestUptime_SimBasisForValidatedDPS(t *testing.T) {
	in := furyRendLow(t, noPool(3))
	in.Refs.Sim = &SimRef{Band: 60, Uptime: map[int64]float64{f2RendID: 95}}
	f := uptime(in, CountedTime(in)).Finding
	if f == nil || f.Basis.Kind != BasisSim || !strings.HasPrefix(f.Basis.Label, "The written Fury rotation, simulated at level 60: ") {
		t.Fatalf("finding %+v", f)
	}
}

func TestUptime_TankNeverUsesTheSim(t *testing.T) {
	in := protShieldBlock(t, 42_000, 232_000, noPool(3))
	in.Refs.Sim = &SimRef{Band: 60, Uptime: map[int64]float64{f2ShieldBlockID: 95}}
	o := uptime(in, CountedTime(in))
	if o.Finding != nil || len(o.CouldNotCheck) != 1 || !strings.HasPrefix(o.CouldNotCheck[0], "not enough logs yet — 3 Protection") {
		t.Fatalf("finding %+v, couldn't check %q", o.Finding, o.CouldNotCheck)
	}
}

func TestUptime_TimeDeadIsLeftOut(t *testing.T) {
	in := furyBattleShout(t, 0.5, poolUptime(6673, 97, 25))
	in.Fight = withDeath(in.Fight, furyGUID, 100_000, 0, "")
	in.Fight = withAura(in.Fight, furyGUID, f2BattleShoutID, f2BattleShoutNme, []string{furyGUID}, span(0, 100_000))
	o := uptime(in, CountedTime(in))
	if o.Finding != nil {
		t.Fatalf("up for all of the 100 s alive, got %+v", o.Finding)
	}
	if len(o.Strengths) != 1 || !strings.HasPrefix(o.Strengths[0].Sentence, "Battle Shout up 100%") {
		t.Fatalf("strengths %+v", o.Strengths)
	}
}

func TestUptime_HealerJudgesOnlySelfBuffs(t *testing.T) {
	in := holyInput(t)
	in.Tables.Rules.Auras = append(slices.Clone(in.Tables.Rules.Auras),
		AuraRule{SpellID: f2RendID, Name: f2RendName, Kind: AuraDot, Verified: fixtureVerified})
	in.Fight = withAura(in.Fight, priestGUID, 10952, "Inner Fire", []string{priestGUID}, span(0, in.Fight.DurationMS))
	in.Fight = withAura(in.Fight, f2BossGUID, f2RendID, f2RendName, []string{priestGUID}, span(0, 1_000))
	refs := poolUptime(10952, 90, 25)
	refs.Pool.Median.Uptime[SpellKey(f2RendID)] = 99
	o := uptime(withPool(in, refs), CountedTime(in))
	if o.Finding != nil || len(o.Strengths) != 1 || !strings.HasPrefix(o.Strengths[0].Sentence, "Inner Fire up 100%") {
		t.Fatalf("finding %+v strengths %+v", o.Finding, o.Strengths)
	}
}

func TestUptime_ThinPoolSaysTheRealCount(t *testing.T) {
	in := furyRendLow(t, noPool(19)) // 19 Fury kills of Onyxia (40) on the site, no validated sim
	o := uptime(in, CountedTime(in))
	if o.Finding != nil {
		t.Fatal("thin basis means no finding")
	}
	want := "not enough logs yet — 19 Fury kills of Onyxia (40) on the site, the coach needs 20."
	if !slices.Contains(o.CouldNotCheck, want) {
		t.Fatalf("couldn't check = %q", o.CouldNotCheck)
	}
}

func TestUptime_NoPoolNoSimIsCouldntCheck(t *testing.T) {
	in := furyRendLow(t, noPool(0)) // not validated: Refs.Sim nil, SimPreparing false
	o := uptime(in, CountedTime(in))
	want := "not enough logs yet — 0 Fury kills of Onyxia (40) on the site, the coach needs 20."
	if o.Finding != nil || !slices.Contains(o.CouldNotCheck, want) {
		t.Fatalf("finding %+v, couldn't check %q", o.Finding, o.CouldNotCheck)
	}
	if strings.Contains(strings.ToLower(textOf(o)), "preparing") {
		t.Fatal("no sim was asked for, so nothing is preparing")
	}
}

func TestUptime_PoolWithoutTheAuraGivesNoBasisAndNoThinLine(t *testing.T) {
	in := furyRendLow(t, poolUptime(6673, 97, 25))
	o := uptime(in, CountedTime(in))
	if o.Finding != nil || len(o.CouldNotCheck) != 0 {
		t.Fatalf("finding %+v, couldn't check %q", o.Finding, o.CouldNotCheck)
	}
}
