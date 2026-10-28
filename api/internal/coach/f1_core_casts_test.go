package coach

import (
	"slices"
	"strings"
	"testing"
)

func TestCoreCasts(t *testing.T) {
	const bt = fixtureBloodthirstID
	steady := repeat(500, 60) // 50 rage every second of a 60 s fight
	five := []int64{0, 6000, 12000, 18000, 24000}
	cases := []struct {
		name               string
		in                 Input
		wantTitle          string
		wantSev            Severity
		wantNums, wantTime []string
		wantDetail         string // substring
		none               bool
	}{
		{name: "dps gap with rage: 4 lost of 9", in: furyFight(t, 60_000, bt, five, steady),
			wantTitle: "Bloodthirst: 4 casts lost", wantSev: SeverityMajor, wantNums: []string{"5", "9", "4"}, wantTime: []string{"0:30–1:00"}},
		{name: "no rage in the gap: nothing lost", in: furyFight(t, 60_000, bt, five, append(repeat(500, 30), repeat(0, 30)...)), none: true},
		// 8 casts end at 42 s: 48-60 s is 12 s ready, less the 1.5 s grace = 10.5 s -> exactly 1 lost of 9 possible (11%), under CoreFiresLost
		{name: "one lost is under CoreFiresLost", in: furyFight(t, 60_000, bt, seqEvery(0, 6000, 8), steady), none: true},
		{name: "unknown cost counts gaps and says so", in: withUnknownCost(furyFight(t, 60_000, bt, five, nil), bt),
			wantTitle: "Bloodthirst: 4 casts lost", wantSev: SeverityMajor, wantDetail: "Couldn't check whether you had the rage."},
		{name: "dead time is never counted", in: furyFightDeadAt(t, 60_000, bt, five, steady, 30_000), none: true},
		{name: "talent ability never cast is not assessed", in: furyTalentUncast(t), none: true},
		{name: "healer: only efficiency cooldowns", in: holyWithBloodthirstRule(t), none: true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			o := coreCasts(tc.in, CountedTime(tc.in))
			if tc.none {
				if o.Finding != nil {
					t.Fatalf("unexpected finding %+v", o.Finding)
				}
				return
			}
			f := o.Finding
			if f == nil || f.Code != CodeCoreCasts || f.Title != tc.wantTitle || f.Severity != tc.wantSev {
				t.Fatalf("got %+v", f)
			}
			if f.Basis.Kind != BasisSpellData || f.Basis.Label != "Measured against Bloodthirst's 6-second cooldown and your own rage bar." {
				t.Errorf("basis %+v", f.Basis)
			}
			assertNumbers(t, f.Facts, tc.wantNums...)
			assertTimes(t, f.Facts, tc.wantTime...)
			if tc.wantDetail != "" && !strings.Contains(f.Detail, tc.wantDetail) {
				t.Errorf("detail %q", f.Detail)
			}
			if len(f.Anchors) == 0 {
				t.Error("a finding needs an anchor")
			}
		})
	}
}

func TestCoreCasts_PoolComparisonIsCastsPerMinute(t *testing.T) {
	in := withPool(furyFight(t, 60_000, 23881, []int64{0, 6000, 12000, 18000, 24000}, repeat(500, 60)),
		poolCPM("Bloodthirst", 9.1, 25))
	f := coreCasts(in, CountedTime(in)).Finding
	if f == nil || f.Comparison == nil || f.Comparison.Kind != BasisTopLogs || !strings.HasPrefix(f.Comparison.Label, "Top 25 Fury warriors on Onyxia (40): ") {
		t.Fatalf("comparison %+v", f)
	}
	assertNumbers(t, f.Facts, "9.1", "5.0")
	if !strings.HasSuffix(f.Comparison.Label, "9.1 a minute. You: 5.0.") {
		t.Errorf("label %q", f.Comparison.Label)
	}
}

func TestCoreCasts_WorstActionsSkipsTheSim(t *testing.T) {
	five := []int64{0, 6000, 12000, 18000, 24000}
	simmed := withPool(furyFight(t, 60_000, 23881, five, repeat(500, 60)), simCPM("Bloodthirst", 10.2, 60)) // no pool, a validated SimRef
	if f := coreCasts(simmed, CountedTime(simmed)).Finding; f == nil || f.Comparison == nil || f.Comparison.Kind != BasisSim {
		t.Fatalf("off WorstActions the sim compares: %+v", f)
	}
	worst := withWorstAction(simmed, "Bloodthirst")
	f := coreCasts(worst, CountedTime(worst)).Finding
	if f == nil || f.Comparison != nil || f.Basis.Kind != BasisSpellData {
		t.Fatalf("on WorstActions the sim never compares: %+v", f)
	}
}

func TestCoreCasts_NoComparisonSaysWhyInCouldNotCheck(t *testing.T) {
	in := furyFight(t, 60_000, 23881, []int64{0, 6000, 12000, 18000, 24000}, repeat(500, 60))
	in.Refs = noPool(7)
	o := coreCasts(in, CountedTime(in))
	want := "not enough logs yet — 7 Fury kills of Onyxia (40) on the site, the coach needs 20."
	if o.Finding == nil || o.Finding.Comparison != nil || len(o.CouldNotCheck) != 1 || o.CouldNotCheck[0] != want {
		t.Fatalf("finding %+v couldn't check %q", o.Finding, o.CouldNotCheck)
	}
}

func TestCoreCasts_AnchorsAreTheCastsRowThenTheLongestGaps(t *testing.T) {
	// Casts at 0, 6, 12, 30 and 50 s: ready 18-30 s (10.5 s after grace, 1 lost) and
	// 36-50 s (12.5 s, 2 lost); the 56-60 s tail is under the grace.
	in := furyFight(t, 60_000, 23881, []int64{0, 6000, 12000, 30000, 50000}, repeat(500, 60))
	o := coreCasts(in, CountedTime(in))
	f := o.Finding
	if f == nil || len(f.Anchors) != 3 || len(o.Anchors) != 3 {
		t.Fatalf("finding %+v anchors %+v", f, o.Anchors)
	}
	if o.Anchors[0].Kind != AnchorCast || o.Anchors[0].SpellID != 23881 {
		t.Errorf("first anchor %+v", o.Anchors[0])
	}
	if a := o.Anchors[1]; a.Kind != AnchorTime || *a.AtMS != 36_000 || *a.ToMS != 50_000 {
		t.Errorf("longest gap anchor %+v", a)
	}
	assertTimes(t, f.Facts, "0:36–0:50", "0:18–0:30")
	if !strings.Contains(f.Detail, "Longest gaps 0:36–0:50 and 0:18–0:30.") {
		t.Errorf("detail %q", f.Detail)
	}
	if len(o.Windows) != 2 || o.Windows[0].FromMS != 18_000 || o.Windows[1].ToMS != 50_000 {
		t.Errorf("windows %+v, want every counted gap in time order", o.Windows)
	}
}

func TestCoreCasts_StrengthWhenAtOrAboveTheComparison(t *testing.T) {
	steady := repeat(500, 60)
	full := withPool(furyFight(t, 60_000, 23881, seqEvery(0, 6000, 10), steady), poolCPM("Bloodthirst", 9.0, 25))
	o := coreCasts(full, CountedTime(full))
	if o.Finding != nil || len(o.Strengths) != 1 {
		t.Fatalf("finding %+v strengths %+v", o.Finding, o.Strengths)
	}
	s := o.Strengths[0]
	if s.Sentence != "Bloodthirst: 10.0 casts a minute. Top 25: 9.0." || len(s.Anchors) != 1 {
		t.Errorf("strength %+v", s)
	}
	assertNumbers(t, s.Facts, "10.0", "9.0")
	assertSpells(t, s.Facts, "Bloodthirst")
	below := withPool(furyFight(t, 60_000, 23881, seqEvery(0, 6000, 10), steady), poolCPM("Bloodthirst", 11.0, 25))
	if o := coreCasts(below, CountedTime(below)); len(o.Strengths) != 0 {
		t.Errorf("below the comparison is no strength: %+v", o.Strengths)
	}
}

func TestCoreCasts_NamesTheWorstAbilityAndTheOthersInTheDetail(t *testing.T) {
	// Protection: Revenge is never cast (58.5 s ready / 5 s = 11 lost); Shield Slam is
	// cast once at 0 (52.5 s ready / 6 s = 8 lost).
	in := protInput(t)
	fight := withResource(withoutCasts(withDuration(in.Fight, 60_000), furyGUID), furyGUID, fixtureRagePower, fixtureRageMax, repeat(500, 60)...)
	in.Fight = withCasts(fight, furyGUID, 23925, "Shield Slam", 0)
	f := coreCasts(in, CountedTime(in)).Finding
	if f == nil || f.Title != "Revenge: 11 casts lost" {
		t.Fatalf("finding %+v", f)
	}
	if !strings.Contains(f.Detail, "Shield Slam") {
		t.Errorf("detail should name the next ability: %q", f.Detail)
	}
	assertSpells(t, f.Facts, "Revenge", "Shield Slam")
	assertNumbers(t, f.Facts, "11", "8")
}

func TestCoreCasts_PetCastsFoldIntoTheOwner(t *testing.T) {
	in := furyFight(t, 60_000, 23881, []int64{0, 6000, 12000}, repeat(500, 60))
	pet := in.Fight.Casts[len(in.Fight.Casts)-1]
	pet.GUID, pet.Sequence = "Pet-0-1", []int64{18000, 24000}
	in.Fight.Casts = append(in.Fight.Casts, pet)
	f := coreCasts(in, CountedTime(in)).Finding
	if f == nil {
		t.Fatal("expected a finding")
	}
	assertNumbers(t, f.Facts, "5", "9")
}

// An omitted or empty core list (efficiency for a healer) is unknown, never "nothing to
// check": it says so, as F7 does for defensives (A7). The healer fixture lists Bloodthirst
// as core and no efficiency cooldown at all.
func TestCoreCasts_NoListIsCouldntCheck(t *testing.T) {
	holy := holyWithBloodthirstRule(t)
	holy.Tables.Rules.Efficiency = nil
	if o := coreCasts(holy, CountedTime(holy)); o.Finding != nil || !slices.Equal(o.CouldNotCheck, []string{CannotCoreList}) {
		t.Errorf("healer with no efficiency list: %+v %q", o.Finding, o.CouldNotCheck)
	}
	for name, core := range map[string][]RuleEntry{"omitted": nil, "empty": {}} {
		in := furyInput(t)
		in.Tables.Rules.Core = core
		if o := coreCasts(in, CountedTime(in)); o.Finding != nil || !slices.Equal(o.CouldNotCheck, []string{CannotCoreList}) {
			t.Errorf("%s: %+v %q", name, o.Finding, o.CouldNotCheck)
		}
	}
}
