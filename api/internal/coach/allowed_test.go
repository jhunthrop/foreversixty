package coach

import "testing"

// withHolyValues is p with two of §5.3's Holy priest findings added, for the thousands
// separator and the percentage the Fury example lacks.
func withHolyValues(p Pack) Pack {
	p.Findings = append(p.Findings,
		Finding{ID: "f7", Code: CodeDeath, Title: "Died at 3:20 to Flame Breath",
			Detail: "Flame Breath hit for 2,140 at 3:20. No healing potion or Healthstone used in the 10 s before.",
			Facts: Facts{Numbers: []NumberFact{{V: 2140, Text: "2,140", Unit: "damage", Of: "killing blow damage"}},
				Spells: []SpellFact{{ID: 18435, Name: "Flame Breath"}}, Times: []string{"3:20"}}},
		Finding{ID: "f5", Code: CodeRanks, Title: "Greater Heal (Rank 4): 44% overheal",
			Facts: Facts{Numbers: []NumberFact{{V: 44, Text: "44", Unit: "%", Of: "Greater Heal overheal"}}}},
	)
	return p
}

func TestAllowed_NumbersTimesRangesNamesIDs(t *testing.T) {
	a := allowedOf(withHolyValues(examplePack()))
	for _, ok := range []bool{a.HasNumber("31"), a.HasNumber("2140"), a.HasNumber("44%"), a.HasTime("1:12"),
		a.HasRange("1:12–1:25"), a.HasName("Bloodthirst"), a.HasName("Onyxia"), a.HasID("a3")} {
		if !ok {
			t.Fatal("a pack value is missing from the allowed set")
		}
	}
	if a.HasNumber("12") || a.HasName("Whirlwind") || a.HasRange("1:12–1:30") {
		t.Fatal("the allowed set holds a value the pack does not")
	}
}

func TestAllowed_ContextLinksAndAnchors(t *testing.T) {
	p := examplePack()
	p.Context.Level = fixtureLevel
	a := allowedOf(p)
	checks := map[string]bool{
		"size 40":              a.HasNumber("40"),
		"level 60":             a.HasNumber("60"),
		"duration 3:52":        a.HasTime("3:52"),
		"basis 6 s cooldown":   a.HasNumber("6"),
		"comparison 9.1":       a.HasNumber("9.1"),
		"link's cap range":     a.HasRange("2:41–2:50"),
		"strength 97%":         a.HasNumber("97%"),
		"uptime 100 bare":      a.HasNumber("100"),
		"spec name":            a.HasName("fury"),
		"class name":           a.HasName("Warrior"),
		"item":                 a.HasName("Grilled Squid"),
		"strength spell":       a.HasName("Battle Shout"),
		"link id":              a.HasID("l1"),
		"strength id":          a.HasID("s2"),
		"anchor a5's start":    a.HasTime("2:41"),
		"not a time as number": !a.HasNumber("1:12"),
	}
	for name, ok := range checks {
		if !ok {
			t.Errorf("%s", name)
		}
	}
}

func TestAllowed_MeleeKillerByName(t *testing.T) {
	p := examplePack()
	p.Findings = append(p.Findings, Finding{ID: "f7", Code: CodeDeath,
		Facts: Facts{Spells: []SpellFact{{ID: 0, Name: "Onyxia's melee"}}}})
	if !allowedOf(p).HasName("Onyxia's melee") {
		t.Fatal("a melee killer (spell id 0) is allowed by name")
	}
}

func TestAllowed_IsWhatTheFactCheckerConsumes(t *testing.T) {
	p := withHolyValues(examplePack())
	c := FactChecker{Allowed: allowedOf(p), Class: "Warrior", Spec: "Fury"}
	if probs := c.Check("It hit for 2,140 at 3:20, 44% overheal, at 2:40 to 2:51."); len(probs) != 0 {
		t.Fatalf("%+v", probs)
	}
	if probs := c.Check("12 casts at 4:44."); len(probs) != 2 {
		t.Fatalf("%+v", probs)
	}
}
