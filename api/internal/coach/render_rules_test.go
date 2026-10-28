package coach

import (
	"slices"
	"strings"
	"testing"
)

func TestRenderRules_MapsThePackMechanically(t *testing.T) {
	drill := "On a target dummy, press Bloodthirst the moment it comes off cooldown, before anything else; Whirlwind and Heroic Strike wait their turn."
	doc := RenderRules(examplePack(), map[string]string{"drill.core_casts": drill})
	if doc.Opening != "Here is what your log shows for Onyxia." || doc.FixFirst.Finding != "f1" ||
		doc.FixFirst.Why != examplePack().Findings[0].Detail || doc.FixFirst.DrillText != drill || len(doc.ThenFix) != 2 {
		t.Fatalf("%+v", doc)
	}
	if doc.FixFirst.Drill != "drill.core_casts" || doc.FixFirst.Title != "Bloodthirst: 9 casts lost" ||
		!slices.Equal(doc.FixFirst.Anchors, []string{"a1", "a2", "a3"}) || !slices.Equal(doc.FixFirst.Links, []string{"l1"}) {
		t.Fatalf("%+v", doc.FixFirst)
	}
	if doc.ThenFix[0].Finding != "f6" || doc.ThenFix[1].Finding != "f4" || doc.ThenFix[1].Why != examplePack().Findings[2].Detail {
		t.Fatalf("%+v", doc.ThenFix)
	}
	if len(doc.DoingWell) != 2 || doc.DoingWell[0].Strength != "s1" || doc.DoingWell[1].Text != "Battle Shout up 100% of the fight. Top 25: 97%." {
		t.Fatalf("%+v", doc.DoingWell)
	}
	if !slices.Equal(doc.CouldntCheck, examplePack().CouldNotCheck) {
		t.Fatal("couldn't check is copied as is")
	}
}

func TestRenderRules_NoFindings(t *testing.T) {
	// opening "Nothing to fix on this pull that the log can show."; FixFirst nil; doing well and couldn't check follow
	p := examplePack()
	p.Findings, p.Links = nil, nil
	doc := RenderRules(p, nil)
	if doc.Opening != "Nothing to fix on this pull that the log can show." || doc.FixFirst != nil || len(doc.ThenFix) != 0 {
		t.Fatalf("%+v", doc)
	}
	if len(doc.DoingWell) != 2 || !slices.Equal(doc.CouldntCheck, p.CouldNotCheck) {
		t.Fatalf("%+v", doc)
	}
}

func TestRenderRules_ThenFixIsRanksTwoToFour(t *testing.T) {
	p := examplePack()
	for i := range 3 {
		p.Findings = append(p.Findings, Finding{ID: "x" + string(rune('1'+i)), Rank: 4 + i, Title: "extra"})
	}
	p.Strengths = append(p.Strengths, Strength{ID: "s3", Sentence: "third"})
	doc := RenderRules(p, nil)
	if len(doc.ThenFix) != 3 || doc.ThenFix[2].Finding != "x1" || len(doc.DoingWell) != 2 {
		t.Fatalf("%+v", doc)
	}
}

func TestRenderRules_RaidBuffAloneStartsAtThen(t *testing.T) {
	doc := RenderRules(Review(raidBuffOnly(t)), nil)
	if doc.FixFirst != nil || len(doc.ThenFix) != 1 || !strings.HasPrefix(doc.ThenFix[0].Title, "No Power Word: Fortitude") {
		t.Fatalf("%+v", doc)
	}
}

func TestRenderRules_PassesTheValidatorsShapeChecks(t *testing.T) {
	p := Review(furyThreeFindings(t))
	doc := RenderRules(p, map[string]string{"drill.core_casts": "Press Bloodthirst on cooldown."})
	persona := Persona{Drills: map[string]string{"drill.core_casts": "Press Bloodthirst on cooldown."}}
	if probs := shapeProblems(doc, p, persona); len(probs) != 0 {
		t.Fatalf("%+v", probs)
	}
}

func TestPackVersion_MovesWithRuleFiles(t *testing.T) {
	rules, catalogue := []byte(`{"spec":"warrior-fury"}`), []byte(`{"flask":[]}`)
	v := PackVersion(rules, catalogue)
	if !strings.HasPrefix(v, "pack-v1.") || v != PackVersion(rules, catalogue) {
		t.Fatal(v)
	}
	for name, other := range map[string]string{
		"rule byte":      PackVersion([]byte(`{"spec":"warrior-furz"}`), catalogue),
		"catalogue byte": PackVersion(rules, []byte(`{"flask":[ ]}`)),
		"moved between":  PackVersion(append(rules, catalogue[0]), catalogue[1:]),
	} {
		if other == v {
			t.Errorf("%s did not move the version", name)
		}
	}
}
