package coach

import "testing"

func factsPack() Allowed {
	return Allowed{
		Numbers: []string{"31", "40", "2,140", "44%", "3:05", "1:12", "1:12–1:25", "9.1"},
		Names:   []string{"Bloodthirst", "Onyxia", "Fury", "Warrior"},
	}
}

func factChecker(t *testing.T, drill string) FactChecker {
	t.Helper()
	return FactChecker{Allowed: factsPack(), Class: "Warrior", Spec: "Fury", DrillText: drill, Names: buildNames(t)}
}

func hasFactProblem(probs []FactProblem, kind, value string) bool {
	for _, p := range probs {
		if p.Kind == kind && p.Value == value {
			return true
		}
	}
	return false
}

const factsWhy = "Bloodthirst is the button the Fury rotation is built around"

func TestValidate_RejectsNumberNotInPack(t *testing.T) {
	probs := factChecker(t, "").Check(factsWhy + ", and 12 casts went unused.")
	if !hasFactProblem(probs, KindNumber, "12") || len(probs) != 1 {
		t.Fatalf("%+v", probs)
	}
	if probs[0].Near == "" || probs[0].Near == "12" {
		t.Errorf("no surrounding words: %+v", probs[0])
	}
}

func TestValidate_AcceptsPackNumbers(t *testing.T) {
	if probs := factChecker(t, "").Check(factsWhy + ": 31 of 40, 9.10 a minute, 44 percent, 2140 at 3:05."); len(probs) != 0 {
		t.Fatalf("%+v", probs)
	}
}

func TestValidate_RejectsSpellNotInPack(t *testing.T) {
	probs := factChecker(t, "").Check("Whirlwind waits while the rage builds.")
	if !hasFactProblem(probs, ProblemName, "Whirlwind") {
		t.Fatalf("%+v", probs)
	}
}

func TestValidate_AllowsTheChosenDrillsNumbersAndNames(t *testing.T) {
	drill := "Hold Bloodthirst first; Whirlwind and Heroic Strike wait their turn, 2 GCDs at a time."
	probs := factChecker(t, drill).Check("Whirlwind and Heroic Strike wait; hold it for 2 GCDs.")
	if len(probs) != 0 {
		t.Fatalf("%+v", probs)
	}
}

func TestValidate_PlayersClassAndSpecAreAllowed(t *testing.T) {
	if probs := factChecker(t, "").Check("A Fury Warrior holds Bloodthirst."); len(probs) != 0 {
		t.Fatalf("%+v", probs)
	}
}

func TestValidate_RangeMustMatchAPackRange(t *testing.T) {
	c := factChecker(t, "")
	if probs := c.Check("The gap ran 1:12–1:30."); !hasFactProblem(probs, KindRange, "1:12–1:30") {
		t.Errorf("wrong range accepted: %+v", probs)
	}
	if probs := c.Check("The gap ran 1:12–1:25."); len(probs) != 0 {
		t.Errorf("pack range rejected: %+v", probs)
	}
	if probs := c.Check("The gap ran 1:12 to 1:25."); len(probs) != 0 {
		t.Errorf("'to' range rejected: %+v", probs)
	}
}

func TestValidate_TimeMustBeInThePack(t *testing.T) {
	if probs := factChecker(t, "").Check("It happened at 4:44."); !hasFactProblem(probs, KindTime, "4:44") {
		t.Fatalf("%+v", probs)
	}
}

func TestValidate_NumberWordAboveThree(t *testing.T) {
	probs := factChecker(t, "").Check("Four casts went unused, but one, two and three are fine.")
	if !hasFactProblem(probs, ProblemNumberWord, "four") || len(probs) != 1 {
		t.Fatalf("%+v", probs)
	}
}

func TestValidate_ThousandsSeparatorsNormalise(t *testing.T) {
	c := factChecker(t, "")
	if probs := c.Check("It cost 2140 in total."); len(probs) != 0 {
		t.Errorf("plain digits rejected: %+v", probs)
	}
	if probs := c.Check("It cost 2,140 in total."); len(probs) != 0 {
		t.Errorf("separated digits rejected: %+v", probs)
	}
}
