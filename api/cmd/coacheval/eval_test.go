package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/coach"
	"github.com/jhunthrop/foreversixty/api/internal/llm"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const (
	repoBuilds    = "../../../data/builds"
	repoCoach     = "../../../data/curated/coach"
	furyName      = "Zorgrimthak"
	otherName     = "Quillfeather"
	furyGUID      = "Player-4711-0ABCDEF1"
	testEncounter = 1084
	testDuration  = 180000
)

func testSummary(index int) summary.Summary {
	return summary.Summary{
		FightIndex: index, EncounterID: testEncounter, Difficulty: 1, Kill: true, DurationMS: testDuration,
		Roster: []summary.RosterRow{
			{GUID: furyGUID, Name: furyName, Class: "Warrior", Spec: "Fury", Role: "dps", Level: 60, ActiveMS: testDuration},
			{GUID: "Player-4711-0ABCDEF2", Name: otherName, Class: "Priest", Spec: "Holy", Role: "healer", Level: 60, ActiveMS: 1000},
		},
	}
}

func corpusDir(t *testing.T, n int) string {
	t.Helper()
	dir := t.TempDir()
	for i := 1; i <= n; i++ {
		data, err := json.Marshal(testSummary(i))
		if err != nil {
			t.Fatal(err)
		}
		name := filepath.Join(dir, "REPORT-secret-"+string(rune('a'+i))+".json")
		if err := os.WriteFile(name, data, 0o600); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}

var (
	namesOnce sync.Once
	namesDict *coach.NameDict
	namesErr  error
)

func sharedNames(t *testing.T) *coach.NameDict {
	t.Helper()
	namesOnce.Do(func() {
		ev := Evaluation{Tables: &coach.FileTables{BuildsDir: repoBuilds, CoachDir: repoCoach}, Now: time.Now}
		namesDict, namesErr = ev.loadNames([]Fight{{Summary: testSummary(1)}})
	})
	if namesErr != nil {
		t.Fatal(namesErr)
	}
	return namesDict
}

func testEvaluation(t *testing.T, n int, factory ClientFactory) Evaluation {
	t.Helper()
	now := time.Date(2026, 10, 28, 12, 0, 0, 0, time.UTC)
	return Evaluation{
		Source:   dirSource{dir: corpusDir(t, n), now: func() time.Time { return now }},
		Tables:   &coach.FileTables{BuildsDir: repoBuilds, CoachDir: repoCoach},
		CoachDir: repoCoach, N: n, Now: func() time.Time { return now },
		Names:     sharedNames(t),
		Models:    []ModelSpec{{Slug: "model-a", Temperature: 0.3}, {Slug: "model-b", Temperature: 0.3}},
		ClientFor: factory, Log: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
}

func TestEvaluate_FakeModelsPassAndNoDatabaseIsTouched(t *testing.T) {
	ev := testEvaluation(t, 2, fakeFactory) // Pool is nil: nothing here can reach a database.
	targets, results, err := ev.Evaluate(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(targets) != 2 || len(results) != 4 {
		t.Fatalf("targets %d results %d", len(targets), len(results))
	}
	for _, r := range results {
		if r.Outcome == OutcomeSkipped || r.Calls == 0 || r.TokensOut != r.Calls*fakeTokensOut || !r.CostEstimated {
			t.Fatalf("%s %s %s %s", r.Label, r.Model, r.Outcome, r.Note)
		}
	}
}

const scriptedIn = 10

func fakeOf(contents ...string) *llm.Fake {
	f := &llm.Fake{}
	for _, c := range contents {
		f.Responses = append(f.Responses, llm.Response{Content: c, Usage: llm.Usage{PromptTokens: scriptedIn, CompletionTokens: 5}, CostUSD: 0.001})
	}
	return f
}

func scripted(contents ...string) ClientFactory {
	return func(ModelSpec, Built) (llm.Client, error) { return fakeOf(contents...), nil }
}

// cleanBuilt is a ready pack with no finding, so a review of it can validate.
func cleanBuilt(t *testing.T) Built {
	t.Helper()
	b := builtOf(t)
	b.Pack.Findings, b.Pack.Strengths, b.Pack.Links, b.Pack.Anchors = []coach.Finding{}, []coach.Strength{}, []coach.Link{}, []coach.Anchor{}
	return b
}

func goodAnswer(t *testing.T, b Built) string {
	t.Helper()
	out, err := json.Marshal(coach.RenderRules(b.Pack, b.Persona.Drills))
	if err != nil {
		t.Fatal(err)
	}
	return string(out)
}

func builtOf(t *testing.T) Built {
	t.Helper()
	ev := testEvaluation(t, 1, fakeFactory)
	fights, err := ev.Source.Fights(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
	target := planTargets(fights, fightLabel)[0]
	b, err := Builder{Tables: ev.Tables, Personas: &coach.Personas{Dir: repoCoach}}.Build(context.Background(), target)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func TestRun_OutcomesFirstRetryFallbackAndError(t *testing.T) {
	built := cleanBuilt(t)
	good := goodAnswer(t, built)
	cases := []struct {
		name   string
		client llm.Client
		want   Outcome
		calls  int
	}{
		{"first", fakeOf(good), OutcomePassFirst, 1},
		{"retry", fakeOf("not json", good), OutcomePassRetry, 2},
		{"fallback", fakeOf("not json", "still not json"), OutcomeFallback, 2},
		{"error", &llm.Fake{Errs: []error{errors.New("gateway down")}}, OutcomeModelError, 0},
	}
	runner := Runner{Names: sharedNames(t), Now: time.Now}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := runner.Run(context.Background(), built, Model{Slug: "m", Client: c.client})
			if r.Outcome != c.want || r.Calls != c.calls {
				t.Fatalf("outcome %s calls %d note %q problems %v", r.Outcome, r.Calls, r.Note, r.Problems)
			}
			if c.want == OutcomeFallback && (len(r.Problems) == 0) {
				t.Fatal("problems not named")
			}
			if c.calls == 2 && r.TokensIn != 2*scriptedIn {
				t.Fatalf("tokens not summed: %d", r.TokensIn)
			}
		})
	}
}

func TestRun_SkipsAPackThatIsNotReady(t *testing.T) {
	built := cleanBuilt(t)
	built.Pack.Status = coach.StatusNotReviewable
	f := &llm.Fake{}
	r := Runner{Names: sharedNames(t), Now: time.Now}.Run(context.Background(), built, Model{Slug: "m", Client: f})
	if r.Outcome != OutcomeSkipped || len(f.Calls) != 0 {
		t.Fatalf("%s, %d calls", r.Outcome, len(f.Calls))
	}
}

func TestPackCarriesNoIdentifier(t *testing.T) {
	b := builtOf(t)
	js := packJSON(b.Pack)
	for _, id := range []string{furyName, otherName, furyGUID, "REPORT-secret"} {
		if strings.Contains(js, id) {
			t.Fatalf("pack carries %q", id)
		}
	}
}

func TestWriteOutputs_RedactsIdentifiersAndMapsFightsPrivately(t *testing.T) {
	leak := "RAWSENTINEL Zorgrimthak should use REPORT-secret-b more, said " + otherName
	ev := testEvaluation(t, 1, scripted(leak, leak))
	targets, results, err := ev.Evaluate(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	out := t.TempDir()
	if err := WriteOutputs(out, targets, results, RunMeta{Source: "corpus directory", Planned: len(targets)}); err != nil {
		t.Fatal(err)
	}
	for _, f := range []string{resultsFile, reportFile} {
		data, err := os.ReadFile(filepath.Join(out, f))
		if err != nil {
			t.Fatal(err)
		}
		for _, id := range []string{furyName, otherName, furyGUID, "REPORT-secret", "RAWSENTINEL"} {
			if strings.Contains(string(data), id) {
				t.Fatalf("%s carries %q", f, id)
			}
		}
	}
	private, err := os.ReadFile(filepath.Join(out, privateFile))
	if err != nil || !strings.Contains(string(private), "REPORT-secret") {
		t.Fatalf("private map missing the report id: %v", err)
	}
}

func TestReport_BuildNotesRowsAndFakeBanner(t *testing.T) {
	results := []Result{
		{Label: "f01", Model: "a", Outcome: OutcomePassFirst, TokensIn: 2000, TokensOut: 600, CostUSD: 0.01, CostEstimated: true},
		{Label: "f02", Model: "a", Outcome: OutcomeFallback, TokensIn: 4000, TokensOut: 800, Problems: []Problem{{Kind: "fact", Value: "12 casts"}}},
		{Label: "f03", Model: "a", Outcome: OutcomeSkipped},
	}
	md := RenderReport(results, RunMeta{Fake: true, Planned: 3})
	for _, want := range []string{"| Evaluation: a pass rate, mean tokens in/out | 50% (1/2), 3000 in / 700 out |", "Fake run", "~$0.0100", "fact 1"} {
		if !strings.Contains(md, want) {
			t.Fatalf("missing %q in\n%s", want, md)
		}
	}
}

func TestPlanTargets_SpreadsSpecs(t *testing.T) {
	mk := func(spec string) Fight {
		return Fight{Summary: summary.Summary{Roster: []summary.RosterRow{
			{Name: "a", Class: "Warrior", Spec: "Fury", ActiveMS: 9}, {Name: "b", Class: "Mage", Spec: spec, ActiveMS: 1}}}}
	}
	got := planTargets([]Fight{mk("Fire"), mk("Fire")}, fightLabel)
	if got[0].Row.Spec != "Fire" && got[0].Row.Spec != "Fury" || got[0].Row.Spec == got[1].Row.Spec {
		t.Fatalf("specs %s %s", got[0].Row.Spec, got[1].Row.Spec)
	}
}

func TestParseOptions(t *testing.T) {
	env := func(string) string { return "x" }
	base := []string{"-corpus", "c", "-out", "o", "-model", "m"}
	if _, err := ParseOptions(base, env, io.Discard); err != nil {
		t.Fatal(err)
	}
	for name, args := range map[string][]string{
		"spend":     append([]string{"-dry-spend=false"}, base...),
		"no source": {"-out", "o", "-model", "m"},
		"both":      append([]string{"-db", "u"}, base...),
		"no model":  {"-corpus", "c", "-out", "o"},
		"bad price": append([]string{"-price", "m=1"}, base...),
	} {
		if _, err := ParseOptions(args, env, io.Discard); err == nil {
			t.Fatalf("%s: accepted", name)
		}
	}
}

func TestGatewayClients_MissingKeyNamesTheVariableOnly(t *testing.T) {
	o := Options{Models: []string{"m"}, Prices: map[string]price{"m": {In: 1, Out: 2}}}
	_, err := gatewayClients(o, func(string) string { return "" })
	if err == nil || !strings.Contains(err.Error(), envKey) {
		t.Fatalf("err %v", err)
	}
	secret := "sk-very-secret"
	env := func(k string) string {
		if k == envKey {
			return secret
		}
		return ""
	}
	if _, err := gatewayClients(Options{Models: []string{"m"}}, env); err == nil || strings.Contains(err.Error(), secret) {
		t.Fatalf("price error %v", err)
	}
	if _, err := gatewayClients(o, env); err != nil {
		t.Fatal(err)
	}
}

func TestRun_FakeSmokeWritesTheOutputs(t *testing.T) {
	out := t.TempDir()
	args := []string{"-corpus", corpusDir(t, 2), "-out", out, "-fake", "-model", "m1", "-model", "m2",
		"-coach-dir", repoCoach, "-builds-dir", repoBuilds, "-n", "2"}
	if err := run(context.Background(), args, func(string) string { return "" }, io.Discard, io.Discard); err != nil {
		t.Fatal(err)
	}
	for _, f := range []string{resultsFile, reportFile, privateFile} {
		if _, err := os.Stat(filepath.Join(out, f)); err != nil {
			t.Fatal(err)
		}
	}
}
