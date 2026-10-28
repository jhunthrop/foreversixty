package coach

import (
	"bytes"
	"context"
	"log/slog"
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/llm"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// privacyPack is examplePack with title in f1's title: a killer named from a player GUID.
func privacyPack(title string) Pack {
	p := examplePack()
	p.Findings = slices.Clone(p.Findings)
	p.Findings[0].Title = title
	return p
}

func TestAttempt_ARosterNameOrRealmInThePackNeverLeaves(t *testing.T) {
	for name, title := range map[string]string{
		"name":           "Killed by Zoë's melee",
		"name any case":  "Killed by ZOË's melee",
		"realm":          "Killed by a player from Mirage",
		"name and realm": "Killed by Zoë-Mirage's melee",
	} {
		var logged bytes.Buffer
		fake := &llm.Fake{Responses: []llm.Response{{Content: mustJSON(t, exampleReview())}}}
		g := dbFree(t, fake)
		g.Validator.Roster = []string{"Zoë", "Tyrim"}
		g.Realms = []string{"Mirage"}
		g.Log = slog.New(slog.NewJSONHandler(&logged, nil))
		p := privacyPack(title)
		a, err := g.Attempt(context.Background(), p, fixturePersona(t))
		if err != nil || len(fake.Calls) != 0 || a.Source != SourceRules || a.Reason != reasonPrivacy || a.Retryable ||
			!reflect.DeepEqual(a.Doc, RenderRules(p, fixturePersona(t).Drills)) {
			t.Fatalf("%s: no call, the rule-only review: %+v %v calls %d", name, a, err, len(fake.Calls))
		}
		out := logged.String()
		if !strings.Contains(out, string(CodeCoreCasts)) || strings.Contains(strings.ToLower(out), "zoë") ||
			strings.Contains(out, "Mirage") {
			t.Fatalf("%s: the warning names the finding code only: %s", name, out)
		}
	}
}

func TestAttempt_ARosterNameInsideAWordIsNotAHit(t *testing.T) {
	fake := &llm.Fake{Responses: []llm.Response{{Content: mustJSON(t, exampleReview())}}}
	g := dbFree(t, fake)
	g.Validator.Roster = []string{"Blood", "Ony"} // only inside "Bloodthirst" and "Onyxia"
	a, err := g.Attempt(context.Background(), examplePack(), fixturePersona(t))
	if err != nil || a.Source != SourceModel || len(fake.Calls) != 1 {
		t.Fatalf("%+v %v", a, err)
	}
}

func TestFightRosterOf_NamesAndDistinctRealms(t *testing.T) {
	sum := summary.Summary{Roster: []summary.RosterRow{{Name: "Zoë-Mirage"}, {Name: "Tyrim-Mirage"}, {Name: "Kraggor-Stalagg"}, {Name: "Mia"}}}
	got := fightRosterOf(sum)
	want := fightRoster{Names: []string{"Zoë", "Tyrim", "Kraggor", "Mia"}, Realms: []string{"Mirage", "Stalagg"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v want %+v", got, want)
	}
}

func TestGenerate_PrivacyHitTakesNoGenerationSlot(t *testing.T) {
	f := newGenFixture(t)
	fake := &llm.Fake{}
	r, err := f.gen(t, fake).Generate(context.Background(), t28Key(), privacyPack("Killed by Tyrim's melee"), fixturePersona(t), f.user)
	if err != nil || r.Source != SourceRules || r.Reason != reasonPrivacy || r.Retryable || len(fake.Calls) != 0 {
		t.Fatalf("%+v %v", r, err)
	}
	var n int
	if err := f.pool.QueryRow(context.Background(), `select count(*) from coach_account_day`).Scan(&n); err != nil || n != 0 {
		t.Fatalf("no generation slot is taken: %d %v", n, err)
	}
}
