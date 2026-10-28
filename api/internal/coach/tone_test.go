package coach

import "testing"

func TestTone_Problems(t *testing.T) {
	for _, s := range []string{"Great job!", "That was bad.", "Just heal more.", "Just heal  more.", "Stop being lazy.",
		"You need to pump harder.", "This guild would love it.", "A world first kill."} {
		if len(toneProblems(s)) == 0 {
			t.Errorf("%q passed", s)
		}
	}
	if got := toneProblems("Bloodthirst is the button the Fury rotation is built around."); len(got) != 0 {
		t.Errorf("plain advice flagged: %v", got)
	}
	if got := toneProblems("A failover is not a fail."); len(got) != 1 {
		t.Errorf("whole words only: %v", got)
	}
}

func TestTone_RosterNamesAreWholeWordsCaseInsensitive(t *testing.T) {
	roster := []string{"Mia", "Kraggor", " "}
	if got := rosterNames("kraggor stood in the fire", roster); len(got) != 1 || got[0] != "Kraggor" {
		t.Errorf("%v", got)
	}
	if got := rosterNames("A miasma of rage", roster); len(got) != 0 {
		t.Errorf("substring matched: %v", got)
	}
}

func TestValidate_ToneIsNamedInTheProse(t *testing.T) {
	probs := validateEdited(t, func(d *ReviewDoc) { d.Opening = "That was bad! Almost everything sits in one button." })
	if !hasProblem(probs, ProblemTone, `"bad" in opening`) || !hasProblem(probs, ProblemTone, `"!" in opening`) {
		t.Fatalf("%+v", probs)
	}
}
