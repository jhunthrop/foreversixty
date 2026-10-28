package coach

import (
	"slices"
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
	if !hasProblem(probs, ProblemLink, `fixFirst states a relation ("instead") without citing one of the pack's links`) {
		t.Fatalf("%+v", probs)
	}
	for _, p := range probs {
		if p.Kind == ProblemLink && p.Near == "" {
			t.Errorf("no surrounding words: %+v", p)
		}
	}
}

func TestValidate_NoRankOneMeansNullFixFirst(t *testing.T) {
	pack := examplePack()
	pack.Findings = pack.Findings[1:]
	pack.Findings[0].Rank, pack.Findings[1].Rank = 2, 3
	pack.Links = nil
	d := exampleReview()
	d.FixFirst, d.ThenFix = nil, d.ThenFix[:2]
	d.ThenFix[1].Links = nil
	d.ThenFix[1].Why = "Heroic Strike turns a full bar into damage."
	v := Validator{Names: buildNames(t)}
	if probs := v.Validate(d, pack, fixturePersona(t)); len(probs) != 0 {
		t.Fatalf("%+v", probs)
	}
	d.FixFirst = exampleReview().FixFirst
	if len(v.Validate(d, pack, fixturePersona(t))) == 0 {
		t.Error("a fixFirst with no rank 1 was accepted")
	}
}

func TestValidate_ALinkThatDoesNotInvolveTheItemDoesNotCount(t *testing.T) {
	probs := validateEdited(t, func(d *ReviewDoc) {
		d.ThenFix[0].Why = "Bring a flask, because the rage cap is the other half of the story."
		d.ThenFix[0].Links = []string{"l1"}
	})
	if len(probs) != 1 || probs[0].Kind != ProblemLink {
		t.Fatalf("%+v", probs)
	}
}

// relationCases are sentences that state a relation; each must be refused with no link.
var relationCases = map[string]string{
	"the worked sentence":   "At 2:40–2:51 you sat on a full bar instead, so that one gap cost you twice.",
	"plural flasks":         "Bring flasks, because they are the cheapest damage you will buy.",
	"died paraphrase":       "You died early, which is why the raid lost its healer.",
	"idle paraphrase":       "You stood idle at the door, and that's why the boss reset.",
	"curly apostrophe":      "That’s why the rage was capped.",
	"capped paraphrase":     "Rage capped, which meant Heroic Strike went unused.",
	"led to":                "The gap led to a wasted bar.",
	"as a result":           "The bar filled. As a result, damage fell.",
	"hence":                 "Hence the lost casts.",
	"since":                 "Since the bar was full, Heroic Strike was free.",
	"caused":                "Standing still caused the gap.",
	"costs you":             "Skipping it costs you damage.",
	"upper case connective": "BECAUSE the bar was full.",
}

func TestValidate_RelationWithoutALinkIsRefusedInEverySlot(t *testing.T) {
	slots := map[string]func(*ReviewDoc, string){
		"opening":   func(d *ReviewDoc, s string) { d.Opening = s },
		"fixFirst":  func(d *ReviewDoc, s string) { d.FixFirst.Why, d.FixFirst.Links = s, nil },
		"thenFix":   func(d *ReviewDoc, s string) { d.ThenFix[1].Why, d.ThenFix[1].Links = s, nil },
		"doingWell": func(d *ReviewDoc, s string) { d.DoingWell[0].Text = s },
	}
	for slot, set := range slots {
		for name, sentence := range relationCases {
			probs := validateEdited(t, func(d *ReviewDoc) { set(d, sentence) })
			if !slices.ContainsFunc(probs, func(p Problem) bool { return p.Kind == ProblemLink && p.Near != "" }) {
				t.Errorf("%s / %s: accepted: %+v", slot, name, probs)
			}
		}
	}
}

func TestValidate_RelationWithAPackLinkIsAccepted(t *testing.T) {
	probs := validateEdited(t, func(d *ReviewDoc) {
		d.Opening, d.OpeningLinks = "Solid kill, because the gap met the rage cap.", []string{"l1"}
		d.DoingWell[0].Text, d.DoingWell[0].Links = "Alive, which is why the gap did not cost more.", []string{"l1"}
	})
	if len(probs) != 0 {
		t.Fatalf("%+v", probs)
	}
}

func TestValidate_PlainDescriptionNeedsNoLink(t *testing.T) {
	probs := validateEdited(t, func(d *ReviewDoc) {
		d.Opening = "You were alive the whole fight. The bar sat full at 2:41–2:50."
		d.ThenFix[1].Why = "Heroic Strike turns a full bar into damage."
	})
	if len(probs) != 0 {
		t.Fatalf("%+v", probs)
	}
}

func TestValidate_LinkIdsNotInThePackAreProblemsInEverySlot(t *testing.T) {
	probs := validateEdited(t, func(d *ReviewDoc) {
		d.OpeningLinks = []string{"l7"}
		d.ThenFix[0].Links = []string{"l8"}
		d.DoingWell[0].Links = []string{"l9"}
	})
	for _, f := range []string{"opening", "thenFix[0]", "doingWell[0]"} {
		if !slices.ContainsFunc(probs, func(p Problem) bool { return p.Kind == ProblemLink && strings.HasPrefix(p.Value, f+" cites link") }) {
			t.Errorf("%s: %+v", f, probs)
		}
	}
}

func TestValidate_AdmitsTheMeleeKillerByName(t *testing.T) {
	pack := examplePack()
	pack.Findings[0].Facts.Spells = []SpellFact{{ID: 0, Name: "Onyxia's melee"}}
	d := exampleReview()
	d.FixFirst.Why = "Onyxia's melee ended it; Bloodthirst is the button the Fury rotation is built around."
	if probs := (Validator{Names: buildNames(t)}).Validate(d, pack, fixturePersona(t)); len(probs) != 0 {
		t.Fatalf("%+v", probs)
	}
}

func TestValidate_BossNameIsAllowedWithoutBeingListed(t *testing.T) {
	pack := examplePack()
	pack.Allowed.Names = []string{"Bloodthirst", "Heroic Strike", "Battle Shout", "Grilled Squid", "Elemental Sharpening Stone"}
	if probs := (Validator{Names: buildNames(t)}).Validate(exampleReview(), pack, fixturePersona(t)); len(probs) != 0 {
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
