package coach

import "testing"

// validateEdited runs the whole Validate on the worked example after mutate.
func validateEdited(t *testing.T, mutate func(*ReviewDoc)) []Problem {
	t.Helper()
	d := exampleReview()
	mutate(&d)
	return Validator{Names: buildNames(t)}.Validate(d, examplePack(), fixturePersona(t))
}

func hasProblem(probs []Problem, kind, value string) bool {
	for _, p := range probs {
		if p.Kind == kind && p.Value == value {
			return true
		}
	}
	return false
}

func withWhy(why string) func(*ReviewDoc) {
	return func(d *ReviewDoc) { d.FixFirst.Why = why }
}

const factsWhy = "Bloodthirst is the button the Fury rotation is built around"

func TestValidate_RejectsNumberNotInPack(t *testing.T) {
	probs := validateEdited(t, withWhy(factsWhy+", and 12 casts went unused."))
	if !hasProblem(probs, KindNumber, "12") || len(probs) != 1 {
		t.Fatalf("%+v", probs)
	}
	if probs[0].Near == "" || probs[0].Near == "12" {
		t.Errorf("no surrounding words: %+v", probs[0])
	}
}

func TestValidate_AcceptsPackNumbers(t *testing.T) {
	probs := validateEdited(t, withWhy(factsWhy+": 31 of 40, 9.10 a minute, 100 percent, 1:12 to 1:25."))
	if len(probs) != 0 {
		t.Fatalf("%+v", probs)
	}
}

func TestValidate_RejectsSpellNotInPack(t *testing.T) {
	probs := validateEdited(t, func(d *ReviewDoc) {
		d.FixFirst.Drill, d.FixFirst.DrillText = "drill.uptime", "Put Battle Shout where you can see its timer."
		d.FixFirst.Why = "Whirlwind waits while the rage builds."
	})
	if !hasProblem(probs, ProblemName, "Whirlwind") {
		t.Fatalf("%+v", probs)
	}
}

func TestValidate_AllowsTheChosenDrillsNumbersAndNames(t *testing.T) {
	probs := validateEdited(t, withWhy("Whirlwind and Heroic Strike wait their turn."))
	if hasProblem(probs, ProblemName, "Whirlwind") {
		t.Fatalf("drill facts rejected: %+v", probs)
	}
}

func TestValidate_PlayersClassAndSpecAreAllowed(t *testing.T) {
	if probs := validateEdited(t, withWhy("A Fury Warrior holds Bloodthirst.")); len(probs) != 0 {
		t.Fatalf("%+v", probs)
	}
}

func TestValidate_RangeMustMatchAPackRange(t *testing.T) {
	if probs := validateEdited(t, withWhy("The gap ran 1:12–1:30.")); !hasProblem(probs, KindRange, "1:12–1:30") {
		t.Errorf("wrong range accepted: %+v", probs)
	}
	if probs := validateEdited(t, withWhy("The gap ran 1:12–1:25.")); len(probs) != 0 {
		t.Errorf("pack range rejected: %+v", probs)
	}
	if probs := validateEdited(t, withWhy("The gap ran 1:12 to 1:25.")); len(probs) != 0 {
		t.Errorf("'to' range rejected: %+v", probs)
	}
}

func TestValidate_TimeMustBeInThePack(t *testing.T) {
	if probs := validateEdited(t, withWhy("It happened at 4:44.")); !hasProblem(probs, KindTime, "4:44") {
		t.Fatalf("%+v", probs)
	}
}

func TestValidate_NumberWordAboveThree(t *testing.T) {
	probs := validateEdited(t, withWhy("Four casts went unused, but one, two and three are fine."))
	if !hasProblem(probs, ProblemNumberWord, "four") || len(probs) != 1 {
		t.Fatalf("%+v", probs)
	}
}

func TestValidate_ThousandsSeparatorsNormalise(t *testing.T) {
	pack := examplePack()
	pack.Allowed.Numbers = append(pack.Allowed.Numbers, "2,140")
	for _, prose := range []string{"It cost 2140 in total.", "It cost 2,140 in total."} {
		d := exampleReview()
		d.FixFirst.Why = prose
		probs := Validator{Names: buildNames(t)}.Validate(d, pack, fixturePersona(t))
		if len(probs) != 0 {
			t.Errorf("%q: %+v", prose, probs)
		}
	}
}
