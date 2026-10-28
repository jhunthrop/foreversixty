package coach

import "testing"

func TestTone_Problems(t *testing.T) {
	for _, s := range []string{"Great job!", "That was bad.", "Just heal more.", "Just heal  more.", "Stop being lazy.",
		"You need to pump harder.", "This guild would love it.", "A world first kill.", "That parse was great.", "Your rank #1 placement.", "You were rank # 3 tonight.", "You are in the top 5% of warriors.", "Top 10 % on this boss.",
		"A top player would hold it.", "You are ranked highly.", "Their ranking dropped.", "Your parses are rising."} {
		if len(toneProblems(s)) == 0 {
			t.Errorf("%q passed", s)
		}
	}
	if got := toneProblems("Bloodthirst is the button the Fury rotation is built around."); len(got) != 0 {
		t.Errorf("plain advice flagged: %v", got)
	}
	// The pack's own spell-rank wording is not a ranking claim (F5 says "Rank 4").
	for _, s := range []string{"Rank 5 is trainable at your level.", "Greater Heal (Rank 4) is your most-cast heal.",
		"Cast rank 6 instead.", "Top 25: 97%.", "Top 25 Fury warriors on Onyxia (40): 31% active."} {
		if got := toneProblems(s); len(got) != 0 {
			t.Errorf("%q flagged: %v", s, got)
		}
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
