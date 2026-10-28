package coach

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/llm"
)

// genSecret stands in for the gateway key: nothing the generator logs may carry it.
const genSecret = "sk-t28-secret"

type genFixture struct {
	pool  *pgxpool.Pool
	store *Store
	user  int64
	log   *bytes.Buffer
}

func newGenFixture(t *testing.T) genFixture {
	t.Helper()
	pool, store := reviewPool(t)
	return genFixture{pool: pool, store: store, user: capsUser(t, pool, "t28-gen@example.com"), log: &bytes.Buffer{}}
}

func (f genFixture) gen(t *testing.T, client llm.Client) *Generator {
	t.Helper()
	return &Generator{
		LLM: client, Caps: newCaps(f.pool), Store: f.store, Model: "m",
		Validator: Validator{Names: buildNames(t), Roster: []string{"Tyrim"}},
		Log:       slog.New(slog.NewJSONHandler(f.log, nil)), Now: func() time.Time { return capsNow },
	}
}

func mustJSON(t *testing.T, d ReviewDoc) string {
	t.Helper()
	b, err := json.Marshal(d)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func badReview() ReviewDoc {
	bad := exampleReview()
	bad.FixFirst.Why = "12 casts went unused."
	return bad
}

func TestGenerate_PassStoresModelReview(t *testing.T) {
	f := newGenFixture(t)
	ctx := context.Background()
	fake := &llm.Fake{Responses: []llm.Response{{Content: mustJSON(t, exampleReview()),
		Usage: llm.Usage{PromptTokens: 7000, CompletionTokens: 800}, CostUSD: 0.002, RequestID: "g1", Model: "m"}}}
	r, err := f.gen(t, fake).Generate(ctx, t28Key(), examplePack(), fixturePersona(t), f.user)
	if err != nil || r.Source != SourceModel || r.TokensIn != 7000 || r.TokensOut != 800 || r.CostUSD != 0.002 ||
		r.RequestID != "g1" || r.ModelSlug != "m" || len(r.Validation) != 0 || len(fake.Calls) != 1 {
		t.Fatalf("%+v %v", r, err)
	}
	req := fake.Calls[0]
	if req.MaxTokens != reviewMaxTokens || len(req.JSONSchema) == 0 || !reflect.DeepEqual(req.Messages, BuildPrompt(examplePack(), fixturePersona(t))) {
		t.Fatalf("request %+v", req)
	}
	stored, ok, err := f.store.Review(ctx, t28Key())
	if err != nil || !ok || stored.Source != SourceModel || !reflect.DeepEqual(stored.Doc, exampleReview()) || stored.AccountID != f.user {
		t.Fatalf("the row is stored: %+v %v %v", stored, ok, err)
	}
}

func TestGenerate_RetryOnceWithTheProblemsListed(t *testing.T) {
	f := newGenFixture(t)
	fake := &llm.Fake{Responses: []llm.Response{
		{Content: mustJSON(t, badReview()), Usage: llm.Usage{PromptTokens: 7000, CompletionTokens: 800}, CostUSD: 0.002, RequestID: "g1"},
		{Content: mustJSON(t, exampleReview()), Usage: llm.Usage{PromptTokens: 8000, CompletionTokens: 700}, CostUSD: 0.003, RequestID: "g2"},
	}}
	r, err := f.gen(t, fake).Generate(context.Background(), t28Key(), examplePack(), fixturePersona(t), f.user)
	if err != nil || len(fake.Calls) != 2 {
		t.Fatalf("%v calls %d", err, len(fake.Calls))
	}
	msgs := fake.Calls[1].Messages
	last := msgs[len(msgs)-1].Content
	if r.Source != SourceModel || !strings.HasPrefix(last, "These were not in the evidence pack: ") ||
		!strings.HasSuffix(last, " Use only the pack.") || !strings.Contains(last, `"12"`) {
		t.Fatalf("%+v %q", r, last)
	}
	first := fake.Calls[0].Messages
	if len(msgs) != 4 || !reflect.DeepEqual(msgs[:2], first) || msgs[2].Role != "assistant" || msgs[2].Content != mustJSON(t, badReview()) {
		t.Fatalf("the retry carries the first two messages and the model's answer: %+v", msgs)
	}
	if r.TokensIn != 15000 || r.TokensOut != 1500 || r.CostUSD != 0.005 || r.RequestID != "g2" || len(r.Validation) != 0 {
		t.Fatalf("sums over both attempts, the last request id: %+v", r)
	}
}

func TestGenerate_SecondFailureFallsBackToRules(t *testing.T) {
	f := newGenFixture(t)
	ctx := context.Background()
	fake := &llm.Fake{Responses: []llm.Response{{Content: mustJSON(t, badReview())}, {Content: mustJSON(t, badReview())}}}
	r, err := f.gen(t, fake).Generate(ctx, t28Key(), examplePack(), fixturePersona(t), f.user)
	want := RenderRules(examplePack(), fixturePersona(t).Drills)
	if err != nil || r.Source != SourceRules || r.Reason != reasonValidation || r.Retryable || !reflect.DeepEqual(r.Doc, want) ||
		len(r.Validation) == 0 || len(fake.Calls) != 2 {
		t.Fatalf("%+v %v", r, err)
	}
	stored, ok, _ := f.store.Review(ctx, t28Key())
	if !ok || stored.Source != SourceRules || stored.Retryable || !reflect.DeepEqual(stored.Validation, r.Validation) {
		t.Fatalf("stored %+v", stored)
	}
}

func TestGenerate_UnparseableOutputIsAProblem(t *testing.T) {
	f := newGenFixture(t)
	fake := &llm.Fake{Responses: []llm.Response{{Content: "not json"}, {Content: mustJSON(t, exampleReview())}}}
	r, err := f.gen(t, fake).Generate(context.Background(), t28Key(), examplePack(), fixturePersona(t), f.user)
	last := fake.Calls[1].Messages[3].Content
	if err != nil || r.Source != SourceModel || !strings.HasPrefix(last, "Fix these: ") || !strings.HasSuffix(last, " Use only the pack.") {
		t.Fatalf("%+v %v %q", r, err, last)
	}
}

func TestGenerate_CapRefusalNeverCallsTheModel(t *testing.T) {
	f := newGenFixture(t)
	fake := &llm.Fake{}
	g := f.gen(t, fake)
	g.Caps.Enabled = false
	r, err := g.Generate(context.Background(), t28Key(), examplePack(), fixturePersona(t), f.user)
	if err != nil || r.Source != SourceRules || r.Reason != string(RefusalSwitchOff) || !r.Retryable ||
		len(fake.Calls) != 0 || len(r.Validation) != 0 || r.ModelSlug != "" {
		t.Fatalf("%+v %v", r, err)
	}
	if _, ok, _ := f.store.Review(context.Background(), t28Key()); !ok {
		t.Fatal("a refusal still stores the rule-only row")
	}
}

func TestGenerate_SpendCapRefusesBeforeTheCall(t *testing.T) {
	f := newGenFixture(t)
	fake := &llm.Fake{}
	g := f.gen(t, fake)
	g.Caps.SpendCapUSD = 0
	r, _ := g.Generate(context.Background(), t28Key(), examplePack(), fixturePersona(t), f.user)
	if r.Reason != string(RefusalSpendCap) || !r.Retryable || len(fake.Calls) != 0 {
		t.Fatalf("%+v", r)
	}
}

func TestGenerate_CapsErrorFailsClosed(t *testing.T) {
	f := newGenFixture(t)
	fake := &llm.Fake{}
	g := f.gen(t, fake)
	g.Caps.PerAccount = 30
	r, _ := g.Generate(context.Background(), t28Key(), examplePack(), fixturePersona(t), -1) // no such user: the reserve fails
	if r.Source != SourceRules || r.Reason != string(RefusalError) || !r.Retryable || len(fake.Calls) != 0 {
		t.Fatalf("%+v", r)
	}
}

func TestGenerate_ModelErrorIsRetryableRules(t *testing.T) {
	f := newGenFixture(t)
	fake := &llm.Fake{Errs: []error{errors.New("gateway 503")}}
	r, err := f.gen(t, fake).Generate(context.Background(), t28Key(), examplePack(), fixturePersona(t), f.user)
	if err != nil || r.Source != SourceRules || r.Reason != reasonModelError || !r.Retryable || len(fake.Calls) != 1 {
		t.Fatalf("%+v %v", r, err)
	}
}

func TestGenerate_ModelErrorOnTheRetryKeepsTheFirstAttempt(t *testing.T) {
	f := newGenFixture(t)
	fake := &llm.Fake{Responses: []llm.Response{{Content: mustJSON(t, badReview()), Usage: llm.Usage{PromptTokens: 7000}, RequestID: "g1"}},
		Errs: []error{nil, errors.New("gateway 503")}}
	r, _ := f.gen(t, fake).Generate(context.Background(), t28Key(), examplePack(), fixturePersona(t), f.user)
	if r.Reason != reasonModelError || !r.Retryable || r.TokensIn != 7000 || r.RequestID != "g1" || len(r.Validation) == 0 {
		t.Fatalf("%+v", r)
	}
}

func TestGenerate_RecordsSpendForEveryAttempt(t *testing.T) {
	f := newGenFixture(t)
	fake := &llm.Fake{Responses: []llm.Response{{Content: mustJSON(t, badReview()), CostUSD: 0.002}, {Content: mustJSON(t, badReview()), CostUSD: 0.002}}}
	if _, err := f.gen(t, fake).Generate(context.Background(), t28Key(), examplePack(), fixturePersona(t), f.user); err != nil {
		t.Fatal(err)
	}
	var usd float64
	var calls int
	if err := f.pool.QueryRow(context.Background(), `select usd::float8, calls from llm_spend where day = $1`, capsNow).Scan(&usd, &calls); err != nil {
		t.Fatal(err)
	}
	if usd != 0.004 || calls != 2 {
		t.Fatalf("usd %v calls %d", usd, calls)
	}
}

func TestGenerate_FailureLogCarriesNoKeyOrName(t *testing.T) {
	f := newGenFixture(t)
	bad := badReview()
	bad.Opening = "Tyrim, solid kill."
	second := badReview()
	second.Opening = "TYRIM held on."
	fake := &llm.Fake{Responses: []llm.Response{{Content: mustJSON(t, bad)}, {Content: mustJSON(t, second)}}}
	if _, err := f.gen(t, fake).Generate(context.Background(), t28Key(), examplePack(), fixturePersona(t), f.user); err != nil {
		t.Fatal(err)
	}
	logged := f.log.String()
	if strings.Count(logged, "\n") != 1 || !strings.Contains(logged, "solid kill.") || !strings.Contains(logged, "held on.") ||
		!strings.Contains(logged, "Bloodthirst: 9 casts lost") || !strings.Contains(logged, "12 casts") {
		t.Fatalf("one entry with the pack, both outputs and the problems: %s", logged)
	}
	for _, leak := range []string{genSecret, "Tyrim", "TYRIM", t28Key().PlayerKey, t28Report} {
		if strings.Contains(logged, leak) {
			t.Fatalf("the log carries %q: %s", leak, logged)
		}
	}
}

func TestGenerate_NeedsItsDependencies(t *testing.T) {
	if _, err := (&Generator{}).Generate(context.Background(), t28Key(), examplePack(), Persona{}, 1); err == nil {
		t.Fatal("a generator without a client, caps or store is an error")
	}
}

func TestRetryMessage_ListsFactsThenTheRest(t *testing.T) {
	got := RetryMessage([]Problem{{Kind: ProblemName, Value: "Whirlwind"}, {Kind: "number", Value: "12 casts"},
		{Kind: ProblemName, Value: "Whirlwind"}, {Kind: ProblemShape, Value: "fixFirst.title must equal the pack's"}})
	want := `These were not in the evidence pack: "Whirlwind", "12 casts". Fix these: fixFirst.title must equal the pack's. Use only the pack.`
	if got != want {
		t.Fatalf("got  %q\nwant %q", got, want)
	}
}
