package coach

import (
	"strings"
	"testing"
)

func TestValidate_AcceptsTheWorkedExample(t *testing.T) {
	probs := Validator{Names: buildNames(t)}.Validate(exampleReview(), examplePack(), fixturePersona(t))
	if len(probs) != 0 {
		t.Fatalf("%+v", probs)
	}
}

func TestValidate_ShapeProblems(t *testing.T) {
	cases := map[string]func(*ReviewDoc){
		"fixFirst not rank 1":            func(d *ReviewDoc) { d.FixFirst.Finding = "f2" },
		"title differs from pack":        func(d *ReviewDoc) { d.FixFirst.Title = "Bloodthirst casts lost" },
		"four thenFix":                   func(d *ReviewDoc) { d.ThenFix = append(d.ThenFix, d.ThenFix...) },
		"thenFix repeats rank 1":         func(d *ReviewDoc) { d.ThenFix[0].Finding = "f1" },
		"couldntCheck reworded":          func(d *ReviewDoc) { d.CouldntCheck[0] = "Talents and race are not in the log." },
		"couldntCheck reordered":         func(d *ReviewDoc) { d.CouldntCheck[0], d.CouldntCheck[1] = d.CouldntCheck[1], d.CouldntCheck[0] },
		"unknown anchor":                 func(d *ReviewDoc) { d.FixFirst.Anchors = []string{"a9"} },
		"unknown link":                   func(d *ReviewDoc) { d.FixFirst.Links = []string{"l9"} },
		"unknown drill":                  func(d *ReviewDoc) { d.FixFirst.Drill = "drill.flying" },
		"unknown strength":               func(d *ReviewDoc) { d.DoingWell[0].Strength = "s9" },
		"opening three sentences":        func(d *ReviewDoc) { d.Opening = "One. Two. Three." },
		"opening over 280":               func(d *ReviewDoc) { d.Opening = strings.Repeat("a", 281) },
		"why over 320":                   func(d *ReviewDoc) { d.FixFirst.Why = strings.Repeat("a", 321) },
		"then why over 200":              func(d *ReviewDoc) { d.ThenFix[0].Why = strings.Repeat("a", 201) },
		"drill text over 240":            func(d *ReviewDoc) { d.FixFirst.DrillText = strings.Repeat("a", 241) },
		"well text over 160":             func(d *ReviewDoc) { d.DoingWell[0].Text = strings.Repeat("a", 161) },
		"three doingWell":                func(d *ReviewDoc) { d.DoingWell = append(d.DoingWell, d.DoingWell[0]) },
		"doingWell empty with strengths": func(d *ReviewDoc) { d.DoingWell = nil },
		"fixFirst null with a rank 1":    func(d *ReviewDoc) { d.FixFirst = nil },
		"relation without a link":        func(d *ReviewDoc) { d.FixFirst.Links = nil },
	}
	for name, mutate := range cases {
		if len(validateEdited(t, mutate)) == 0 {
			t.Errorf("%s: accepted", name)
		}
	}
}

func TestValidate_NamesTheProblem(t *testing.T) {
	probs := validateEdited(t, func(d *ReviewDoc) { d.FixFirst.Finding = "f2" })
	if !hasProblem(probs, ProblemShape, `fixFirst.finding is "f2"; it must be the pack's rank 1, "f1"`) {
		t.Fatalf("%+v", probs)
	}
	probs = validateEdited(t, func(d *ReviewDoc) { d.FixFirst.Links = nil })
	if len(probs) != 1 || probs[0].Kind != ProblemLink || !strings.Contains(probs[0].Value, "Rage full for 14 seconds") {
		t.Fatalf("%+v", probs)
	}
}

func TestValidate_NoRankOneMeansNullFixFirst(t *testing.T) {
	pack := examplePack()
	pack.Findings = pack.Findings[1:]
	pack.Findings[0].Rank, pack.Findings[1].Rank = 2, 3
	pack.Links = nil
	d := exampleReview()
	d.FixFirst, d.ThenFix = nil, d.ThenFix[:2]
	v := Validator{Names: buildNames(t)}
	if probs := v.Validate(d, pack, fixturePersona(t)); len(probs) != 0 {
		t.Fatalf("%+v", probs)
	}
	d.FixFirst = exampleReview().FixFirst
	if len(v.Validate(d, pack, fixturePersona(t))) == 0 {
		t.Error("a fixFirst with no rank 1 was accepted")
	}
}

func TestValidate_ALinkThatDoesNotJoinTheTwoFindingsDoesNotCount(t *testing.T) {
	probs := validateEdited(t, func(d *ReviewDoc) {
		d.ThenFix[0].Why = "Bring a flask; the rage cap is the other half of the story."
		d.ThenFix[0].Links = []string{"l1"}
	})
	if len(probs) != 1 || probs[0].Kind != ProblemLink {
		t.Fatalf("%+v", probs)
	}
}

func TestValidate_AnotherPlayersNameIsAProblem(t *testing.T) {
	v := Validator{Names: buildNames(t), Roster: []string{"Kraggor", "Mia"}}
	d := exampleReview()
	d.FixFirst.Why = "Kraggor opened with Bloodthirst on every cooldown."
	probs := v.Validate(d, examplePack(), fixturePersona(t))
	if !hasProblem(probs, ProblemTone, `another player, "Kraggor", in fixFirst.why`) {
		t.Fatalf("%+v", probs)
	}
}
