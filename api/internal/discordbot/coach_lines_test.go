package discordbot

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"
)

const coachURL = "https://forever.test/coach/abc"

type fakeCoach struct {
	lines  []CoachLine
	err    error
	users  []int64
	budget time.Duration
}

func (f *fakeCoach) NightLinesFor(_ context.Context, userID int64, _ string, budget time.Duration) ([]CoachLine, error) {
	f.users = append(f.users, userID)
	f.budget = budget
	return f.lines, f.err
}

func coachWorld(t *testing.T, lines ...CoachLine) (*officerWorld, *fakeCoach) {
	t.Helper()
	w := newOfficerWorld(t)
	fc := &fakeCoach{lines: lines}
	w.svc.d.Coach = fc
	w.addChar(t, w.gid, w.uid, "us/normal/zulmara", "Zulmara", true)
	w.og.report = onyxiaReport(false, false)
	return w, fc
}

func line(name, stem string, on, of int) CoachLine {
	return CoachLine{Character: name, Stem: stem, On: on, Of: of, URL: coachURL}
}

func TestCoachLineTextVariants(t *testing.T) {
	open := ". [Open your coach's review](" + coachURL + ")"
	for name, tc := range map[string]struct {
		in   CoachLine
		want string
	}{
		"several": {line("A", "Bloodthirst casts lost", 5, 8), "Fix first: Bloodthirst casts lost — on 5 of 8 bosses" + open},
		"one":     {line("A", "Bloodthirst casts lost", 1, 1), "Fix first: Bloodthirst casts lost" + open},
		"all":     {line("A", "Bloodthirst casts lost", 8, 8), "Fix first: Bloodthirst casts lost — on all 8 bosses" + open},
		"pending": {CoachLine{Character: "A", Pending: true, URL: coachURL}, "[Your coach's review is on the site.](" + coachURL + ")"},
		"nocount": {line("A", "Stem", 0, 0), "Fix first: Stem" + open},
	} {
		if got := coachLineText(tc.in); got != tc.want {
			t.Errorf("%s = %q, want %q", name, got, tc.want)
		}
	}
}

func TestReport_PremiumOwnerGetsOneLinePerCharacter(t *testing.T) {
	w, fc := coachWorld(t, line("Zulmara", "Bloodthirst casts lost", 5, 8), line("Vexlash", "Early pulls", 8, 8))
	rep := w.run(t, pressDiscord, "report")
	got := headFields(rep)[coachCopy.FieldName]
	want := "**Zulmara** · Fix first: Bloodthirst casts lost — on 5 of 8 bosses. [Open your coach's review](" + coachURL + ")\n" +
		"**Vexlash** · Fix first: Early pulls — on all 8 bosses. [Open your coach's review](" + coachURL + ")"
	if got != want {
		t.Fatalf("field = %q", got)
	}
	if len(fc.users) != 1 || fc.users[0] != w.uid || fc.budget <= 0 || fc.budget > coachLinesBudget {
		t.Fatalf("coach asked %v with %v", fc.users, fc.budget)
	}
}

func TestReport_OneCharacterIsNotNamed(t *testing.T) {
	w, _ := coachWorld(t, line("Zulmara", "Stem", 3, 8))
	got := headFields(w.run(t, pressDiscord, "report"))[coachCopy.FieldName]
	if strings.Contains(got, "Zulmara") || !strings.HasPrefix(got, "Fix first: Stem") {
		t.Fatalf("field = %q", got)
	}
}

func TestReport_AtMostThreeLines(t *testing.T) {
	var lines []CoachLine
	for _, n := range []string{"A", "B", "C", "D", "E"} {
		lines = append(lines, line(n, "Stem", 2, 8))
	}
	w, _ := coachWorld(t, lines...)
	got := headFields(w.run(t, pressDiscord, "report"))[coachCopy.FieldName]
	if n := strings.Count(got, "\n") + 1; n != 3 {
		t.Fatalf("%d lines: %q", n, got)
	}
}

func TestReport_PendingLineReadsTheSiteLine(t *testing.T) {
	w, _ := coachWorld(t, CoachLine{Character: "Zulmara", URL: coachURL, Pending: true})
	got := headFields(w.run(t, pressDiscord, "report"))[coachCopy.FieldName]
	if got != "[Your coach's review is on the site.]("+coachURL+")" {
		t.Fatalf("field = %q", got)
	}
}

func TestReport_NoCoachDepsNoLine(t *testing.T) {
	w, _ := coachWorld(t, line("Zulmara", "Stem", 2, 8))
	with := w.run(t, pressDiscord, "report")
	w.svc.d.Coach = nil
	without := w.run(t, pressDiscord, "report")
	if _, ok := headFields(without)[coachCopy.FieldName]; ok {
		t.Fatal("a coach field without Deps.Coach")
	}
	if len(headFields(with)) != len(headFields(without))+1 {
		t.Fatalf("only the coach field may differ: %v vs %v", headFields(with), headFields(without))
	}
}

func TestReport_NoLinesOrAFailedReadLeaveTheAnswerAlone(t *testing.T) {
	w, fc := coachWorld(t)
	if _, ok := headFields(w.run(t, pressDiscord, "report"))[coachCopy.FieldName]; ok {
		t.Fatal("a coach field with no lines")
	}
	fc.err = errors.New("pack store down")
	rep := w.run(t, pressDiscord, "report")
	if _, ok := headFields(rep)[coachCopy.FieldName]; ok || len(rep.Message.Embeds) == 0 {
		t.Fatal("a failed coach read changed the answer")
	}
}

func TestReport_ShareInChannelRendersWithoutTheLine(t *testing.T) {
	w, _ := coachWorld(t, line("Zulmara", "Stem", 2, 8))
	w.og.report = onyxiaReport(true, true)
	answer := w.run(t, pressDiscord, "report").Message
	if _, ok := headFields(Reply{Message: answer})[coachCopy.FieldName]; !ok {
		t.Fatal("the ephemeral answer has no coach field")
	}
	in := interactionOf(typeComponent, pressDiscord, map[string]any{"custom_id": answer.Components[0][0].CustomID})
	in.Token = "tok-coach"
	rep, err := w.svc.share(context.Background(), in)
	if err != nil || rep.After == nil {
		t.Fatalf("press = %+v %v", rep, err)
	}
	rep.After(context.Background())
	for _, c := range w.rest.Calls() {
		if c.Method != "FollowUp" {
			continue
		}
		for _, e := range c.Message.Embeds {
			for _, f := range e.Fields {
				if f.Name == coachCopy.FieldName {
					t.Fatal("the shared post carries the coach field")
				}
			}
		}
	}
}

func TestReport_EveryLimitHolds(t *testing.T) {
	stem := strings.Repeat("s", 60)
	w, _ := coachWorld(t, line("Abcdefghijkl", stem, 2, 8), line("Mnopqrstuvwx", stem, 3, 8), line("Yzabcdefghij", stem, 8, 8))
	rep := w.run(t, pressDiscord, "report")
	assertMessageLimits(t, rep.Message)
	if v := headFields(rep)[coachCopy.FieldName]; len(v) > fieldValueMax || strings.Count(v, "\n") != 2 {
		t.Fatalf("field = %d bytes: %q", len(v), v)
	}
	raw, err := rep.Message.MarshalJSON()
	if err != nil || !strings.Contains(string(raw), `"allowed_mentions":{"parse":[]}`) {
		t.Fatalf("mentions: %s %v", raw, err)
	}
}

func TestCoachBudgetIsClippedByTheContext(t *testing.T) {
	if got := coachBudget(context.Background()); got != coachLinesBudget {
		t.Fatalf("budget = %v", got)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	if got := coachBudget(ctx); got <= 0 || got > 300*time.Millisecond {
		t.Fatalf("clipped budget = %v", got)
	}
}
