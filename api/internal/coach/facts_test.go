package coach

import "testing"

func TestFacts_SpecAndClassForms(t *testing.T) {
	names := buildNames(t)
	for _, spec := range []string{"Beast Mastery", "beast-mastery", "hunter-beast-mastery"} {
		c := FactChecker{Class: "hunter", Spec: spec, Names: names}
		if probs := c.Check("A Beast Mastery Hunter holds the pet."); len(probs) != 0 {
			t.Errorf("spec %q: %+v", spec, probs)
		}
	}
}

func TestFacts_DrillWordsAndFractionsAreExemptOnlyForTheDrill(t *testing.T) {
	c := FactChecker{DrillText: "Hold it for half a second.", Names: buildNames(t)}
	if probs := c.Check("Hold it for half a second."); len(probs) != 0 {
		t.Errorf("drill word rejected: %+v", probs)
	}
	if probs := c.Check("Triple the casts."); len(probs) != 1 {
		t.Errorf("%+v", probs)
	}
}

// The spec's §5.3 sentences pass the fact checks word for word.
func TestFacts_TheWorkedExampleSentencesPass(t *testing.T) {
	c := FactChecker{
		Allowed: Allowed{Numbers: []string{"9", "40", "1:12–1:25", "2:40–2:51"}, Names: []string{"Bloodthirst"}},
		Class:   "Warrior", Spec: "Fury", Names: buildNames(t),
	}
	why := "Bloodthirst is the button the Fury rotation is built around, and 9 of a possible 40 went unused while you had the rage, " +
		"longest at 1:12–1:25. At 2:40–2:51 you sat on a full bar instead, so that one gap cost you twice."
	if probs := c.Check(why); len(probs) != 0 {
		t.Fatalf("%+v", probs)
	}
}
