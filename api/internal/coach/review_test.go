package coach

import (
	"bytes"
	"encoding/json"
	"slices"
	"strings"
	"testing"
)

func TestReview_RanksMajorFirstThenPrecedence(t *testing.T) {
	// fury: a major F6 (no flask at 40), a major F1, a minor F4 -> f1 core_casts, f2 preparation, f3 resource
	p := Review(furyThreeFindings(t))
	if len(p.Findings) != 3 {
		t.Fatalf("%+v", p.Findings)
	}
	got := []Code{p.Findings[0].Code, p.Findings[1].Code, p.Findings[2].Code}
	if !slices.Equal(got, []Code{CodeCoreCasts, CodePreparation, CodeResource}) || p.Findings[0].ID != "f1" || p.Findings[0].Rank != 1 {
		t.Fatalf("%v", got)
	}
	for i, f := range p.Findings {
		if f.ID != "f"+string(rune('1'+i)) || f.Rank != i+1 {
			t.Errorf("finding %d is %s rank %d", i, f.ID, f.Rank)
		}
	}
}

func TestReview_RaidBuffNeverRankOne(t *testing.T) {
	p := Review(raidBuffAndMinorRanks(t)) // a raid-buff-only preparation finding and a minor F5
	if len(p.Findings) != 2 || p.Findings[0].Code != CodeRanks || p.Findings[0].Rank != 1 ||
		p.Findings[1].Code != CodePreparation || p.Findings[1].Rank != 2 {
		t.Fatalf("%+v", p.Findings)
	}
	alone := Review(raidBuffOnly(t))
	if len(alone.Findings) != 1 || alone.Findings[0].Rank != 2 || RenderRules(alone, nil).FixFirst != nil {
		t.Fatalf("a raid buff alone is rank 2 with no fix first: %+v", alone.Findings)
	}
}

func TestReview_FlaskAndFortitudeMissingFixesTheFlaskFirst(t *testing.T) {
	// flask missing + Fortitude missing -> fix first is the flask
	in := raidBuffOnly(t)
	in.Fight = withoutAura(in.Fight, furyGUID, flask("Flask of the Titans").SpellID)
	p := Review(in)
	doc := RenderRules(p, nil)
	if doc.FixFirst == nil || doc.FixFirst.Title != TitleNoFlask || p.Findings[0].Rank != 1 {
		t.Fatalf("%+v", p.Findings)
	}
	o := preparation(in, CountedTime(in))
	if o.RaidBuffOnly || !o.MissingFlask {
		t.Fatalf("typed flags: %+v", o)
	}
	if alone := preparation(raidBuffOnly(t), CountedTime(raidBuffOnly(t))); !alone.RaidBuffOnly || alone.MissingFlask {
		t.Fatalf("raid buff alone: %+v", alone)
	}
}

func TestReview_RaidBuffWithAPotionShortfallIsNotRankOne(t *testing.T) {
	// raid buffs order after the potions, so the potion shortfall names the finding and
	// a raid-buff title never leads a finding with something else in it
	// (the shipped catalogue's potion list is incomplete, so the issues are built directly)
	in := raidBuffOnly(t)
	buff := prepIssue{order: orderRaidBuff, severity: SeverityMinor, kind: kindRaidBuff, name: "Power Word: Fortitude",
		class: "priest", missing: "Power Word: Fortitude", item: "Power Word: Fortitude"}
	potion := prepIssue{order: orderPotions, severity: SeverityMinor, kind: kindPotions, missing: "combat potions",
		potions: &potionCount{uses: 0, allowed: 1}}
	issues := []prepIssue{buff, potion}
	f := prepFinding(in, issues, nil, nil)
	if strings.Contains(f.Title, "Fortitude") || allPrepKind(issues, kindRaidBuff) {
		t.Fatalf("a raid buff titles a finding with a potion shortfall: %q", f.Title)
	}
	if !allPrepKind([]prepIssue{buff}, kindRaidBuff) || allPrepKind(nil, kindRaidBuff) {
		t.Fatal("raid buffs alone is every issue a raid buff, and at least one")
	}
}

func TestReview_DeathAnchorIsItsTimeWithNoGUID(t *testing.T) {
	in := furyFightDeadAt(t, reviewShortMS, bloodthirstRank4ID, []int64{0, 30_000}, repeat(reviewRageHalf, 60), 40_000)
	p := Review(in)
	var death *Anchor
	for i, a := range p.Anchors {
		if a.Kind == AnchorDeath {
			death = &p.Anchors[i]
		}
	}
	if death == nil || death.AtMS == nil || *death.AtMS != 40_000 {
		t.Fatalf("%+v", p.Anchors)
	}
	b, err := json.Marshal(*death)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != `{"id":"`+death.ID+`","kind":"death","at_ms":40000}` {
		t.Fatalf("%s", b)
	}
	whole, _ := json.Marshal(p)
	if bytes.Contains(whole, []byte(furyGUID)) || bytes.Contains(whole, []byte("Player-")) {
		t.Fatal("the pack carries the GUID")
	}
}

func TestReview_CouldNotCheckStartsWithTalentsAndRange(t *testing.T) {
	p := Review(furyInput(t))
	if len(p.CouldNotCheck) < 2 || p.CouldNotCheck[0] != "Talents and race: Forever logs don't record them." ||
		p.CouldNotCheck[1] != "Range and movement: the log doesn't say where you stood." {
		t.Fatal(p.CouldNotCheck)
	}
}

func TestReview_CouldNotCheckDedupesInCheckOrder(t *testing.T) {
	thin := Fill(ThinPool, map[string]string{"n": "0", "Spec": "Fury", "Boss": "Onyxia", "size": "40"})
	lacks := Drafts["pool_lacks_measure"]
	results := []checked{
		{CodeCoreCasts, outcome{CouldNotCheck: []string{thin}}},
		{CodeUptime, outcome{CouldNotCheck: []string{lacks, thin}}},
		{CodeRanks, outcome{CouldNotCheck: []string{Drafts["cannot_level"], lacks}}},
	}
	want := []string{CannotTalents, CannotRange, thin, lacks, Drafts["cannot_level"]}
	if got := couldNotCheck(results); !slices.Equal(got, want) {
		t.Fatalf("%q", got)
	}
	p := Review(furyInput(t))
	seen := map[string]bool{}
	for _, line := range p.CouldNotCheck {
		if seen[line] {
			t.Errorf("twice: %q", line)
		}
		seen[line] = true
	}
}

func TestReview_NotReviewable(t *testing.T) {
	noSpec := furyInput(t)
	noSpec.Row.Spec = ""
	short := furyInput(t)
	short.Fight = withDuration(short.Fight, 25_000)
	trash := furyInput(t)
	trash.Fight.EncounterID = 0
	for name, in := range map[string]Input{"no spec": noSpec, "25 s": short, "trash": trash} {
		if p := Review(in); p.Status != StatusNotReviewable || len(p.Findings) != 0 || p.NotReviewable == "" {
			t.Errorf("%s: %s with %d findings", name, p.Status, len(p.Findings))
		}
	}
	if got := Review(noSpec).NotReviewable; got != "We couldn't tell this player's spec from the log." {
		t.Fatal(got)
	}
}

func TestReview_PreparingWhileTheSimIsTheOnlyBasis(t *testing.T) {
	in := furyThreeFindings(t)
	in.Refs = Refs{SimPreparing: true}
	if p := Review(in); p.Status != StatusPreparing || len(p.Findings) != 0 {
		t.Fatalf("%s with %d findings", p.Status, len(p.Findings))
	}
	in.Refs = poolCPM("Bloodthirst", 9, 25)
	in.Refs.SimPreparing = true
	if p := Review(in); p.Status != StatusReady {
		t.Fatalf("a pool answers; the sim isn't needed: %s", p.Status)
	}
}

func TestReview_PackCarriesNoNames(t *testing.T) {
	for name, in := range allFixtureInputs(t) {
		b, err := json.Marshal(Review(in))
		if err != nil {
			t.Fatal(err)
		}
		for _, s := range []string{"Tyrim", "Tichondrius", "Aegwynn", "Player-1136", "Player-11-", "Skiimaster"} {
			if bytes.Contains(b, []byte(s)) {
				t.Errorf("%s: the pack leaks %q", name, s)
			}
		}
	}
}

func TestReview_ContextIsTheFight(t *testing.T) {
	in := furyThreeFindings(t)
	want := Context{SpecLabel: "Fury", Class: "Warrior", Role: RoleDPS, Boss: onyxiaName, Size: 40, Kill: true, Duration: "2:00", Level: fixtureLevel}
	if got := Review(in).Context; got != want {
		t.Fatalf("%+v", got)
	}
}

func TestReview_EveryFindingHasAnAnchor(t *testing.T) {
	for name, in := range allFixtureInputs(t) {
		p := Review(in)
		for _, f := range p.Findings {
			if len(f.Anchors) == 0 {
				t.Errorf("%s: %s has no anchor", name, f.ID)
			}
			for _, a := range f.Anchors {
				if !anchorExists(p, a) {
					t.Errorf("%s: %s cites unknown anchor %s", name, f.ID, a)
				}
			}
		}
		for _, s := range p.Strengths {
			for _, a := range s.Anchors {
				if !anchorExists(p, a) {
					t.Errorf("%s: %s cites unknown anchor %s", name, s.ID, a)
				}
			}
		}
	}
}

func TestReview_AnchorsFollowRankOrder(t *testing.T) {
	p := Review(furyThreeFindings(t))
	var cited []string
	for _, f := range p.Findings {
		cited = append(cited, f.Anchors...)
	}
	for i, a := range p.Anchors {
		if a.ID != "a"+string(rune('1'+i)) || (i < len(cited) && cited[i] != a.ID) {
			t.Fatalf("anchors %+v, cited %v", p.Anchors, cited)
		}
	}
	if p.Anchors[0].Kind != AnchorCast || p.Anchors[len(cited)-1].Kind != AnchorResource {
		t.Fatalf("%+v", p.Anchors)
	}
}

func TestReview_StrengthsFollowPrecedence(t *testing.T) {
	in := withPool(furyThreeFindings(t), poolOf(25, Profile{Uptime: map[string]float64{SpellKey(6673): 90}}))
	p := Review(in)
	if len(p.Strengths) < 2 || p.Strengths[0].Sentence != StrengthAlive || !strings.HasPrefix(p.Strengths[1].Sentence, "Battle Shout up 100%") ||
		p.Strengths[0].ID != "s1" || p.Strengths[1].ID != "s2" {
		t.Fatalf("death precedes uptime: %+v", p.Strengths)
	}
}

func TestReview_LinksAndBasisLabels(t *testing.T) {
	p := Review(furyThreeFindings(t))
	if len(p.Links) != 1 || p.Links[0].ID != "l1" || p.Links[0].Findings != [2]string{"f1", "f3"} ||
		p.Links[0].Sentence != "The Bloodthirst gap at 1:36–2:00 overlaps the rage cap at 1:40–1:52." {
		t.Fatalf("%+v", p.Links)
	}
	want := []string{p.Findings[0].Basis.Label, p.Findings[1].Basis.Label, p.Findings[2].Basis.Label}
	if !slices.Equal(p.BasisLabels, want) {
		t.Fatalf("%q", p.BasisLabels)
	}
	if !p.Allowed.HasRange("1:36–2:00") || !p.Allowed.HasID("l1") {
		t.Fatal("the link's values are allowed")
	}
}

func TestReview_HealerHasNoIdleOrAdherence(t *testing.T) {
	// holy with long idle and low adherence -> neither code; no title, detail or strength contains "heal more"
	p := Review(holyIdleAndOffRotation(t))
	if p.Status != StatusReady {
		t.Fatal(p.Status)
	}
	for _, f := range p.Findings {
		if f.Code == CodeIdle || f.Code == CodeAdherence {
			t.Errorf("a healer got %s", f.Code)
		}
		if strings.Contains(strings.ToLower(f.Title+f.Detail), "heal more") {
			t.Errorf("%q", f.Title+f.Detail)
		}
	}
	for _, s := range p.Strengths {
		if strings.Contains(strings.ToLower(s.Sentence), "heal more") {
			t.Errorf("%q", s.Sentence)
		}
	}
	// the same idle on a Fury warrior fires: the healer's silence is the role, not the data
	if !slices.ContainsFunc(Review(furyIdle(t, reviewLongMS, span(10_000, 60_000))).Findings, func(f Finding) bool { return f.Code == CodeIdle }) {
		t.Fatal("the fury control has no idle finding")
	}
}

func TestReview_NoRuleFileSkipsItsChecksAndSaysSo(t *testing.T) {
	in := furyThreeFindings(t)
	in.Tables.Rules = SpecRules{}
	in.Tables.RotationOK = true
	p := Review(in)
	line := "Core casts, uptimes and the rotation: the coach doesn't know the Fury rotation yet."
	if n := countOf(p.CouldNotCheck, line); n != 1 || p.CouldNotCheck[2] != line {
		t.Fatalf("one copy, first after talents and range: %q", p.CouldNotCheck)
	}
	for _, f := range p.Findings {
		if slices.Contains(needsRuleFile, f.Code) {
			t.Errorf("%s ran without a rule file", f.Code)
		}
	}
	if !slices.ContainsFunc(p.Findings, func(f Finding) bool { return f.Code == CodePreparation }) {
		t.Fatal("checks that need no rule file still run")
	}
	if with := Review(furyThreeFindings(t)); countOf(with.CouldNotCheck, line) != 0 {
		t.Fatal("a spec with a rule file never says it has none")
	}
}

func countOf(lines []string, s string) int {
	n := 0
	for _, l := range lines {
		if l == s {
			n++
		}
	}
	return n
}

func TestReview_CapsAtMaxFindings(t *testing.T) {
	results := make([]checked, 0, MaxFindings+2)
	for i := range MaxFindings + 2 {
		results = append(results, checked{CodeCoreCasts, outcome{Finding: &Finding{Code: Precedence[RoleDPS][i%len(Precedence[RoleDPS])], Severity: SeverityMinor}}})
	}
	if fs, _ := rankFindings(RoleDPS, results); len(fs) != MaxFindings {
		t.Fatal(len(fs))
	}
}
