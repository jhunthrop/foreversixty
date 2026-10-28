package coach

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"slices"
	"strconv"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/llm"
)

func (f *t30) noJobNoCall(t *testing.T) {
	t.Helper()
	if len(f.jobs.Ran()) != 0 || len(f.llm.Calls) != 0 {
		t.Fatalf("jobs %v calls %d", f.jobs.Ran(), len(f.llm.Calls))
	}
}

// key is the review key the route uses for the Fury warrior's pack.
func (f *t30) key(t *testing.T) ReviewKey {
	t.Helper()
	p, err := f.svc.PackFor(context.Background(), f.rep(t), t30Fight, f.furyKey)
	if err != nil {
		t.Fatal(err)
	}
	return NewReviewKey(t30Report, t30Fight, f.furyKey, p, "m", fixturePersona(t))
}

func (f *t30) store(t *testing.T, source string, retryable bool) ReviewKey {
	t.Helper()
	k := f.key(t)
	doc := exampleReview()
	doc.Opening = "stored " + source
	if err := f.svc.Store.SaveReview(context.Background(), StoredReview{Key: k, Doc: doc, Source: source,
		Retryable: retryable, GeneratedAt: capsNow}); err != nil {
		t.Fatal(err)
	}
	return k
}

func wantNotFound(t *testing.T, code int, body map[string]any, want string) {
	t.Helper()
	if code != http.StatusNotFound || body["code"] != want {
		t.Fatalf("%d %v, want 404 %s", code, body, want)
	}
}

func TestFightCoach_StrangerGets404(t *testing.T) {
	f := newT30(t)
	code, _, body := f.get(t, user(f.stranger), "")
	wantNotFound(t, code, body, "not_found")
	if body["message"] != msgNotFound {
		t.Fatalf("the same body as an unknown report: %v", body)
	}
	f.noJobNoCall(t)
}

func TestFightCoach_UnknownReportIsTheSame404(t *testing.T) {
	f := newT30(t)
	f.svc.Reports.Pool.Exec(context.Background(), `delete from reports where id = $1`, t30Report)
	code, _, body := f.get(t, user(f.premium), "")
	wantNotFound(t, code, body, "not_found")
	if body["message"] != msgNotFound {
		t.Fatalf("%v", body)
	}
}

func TestFightCoach_PrivateReportNotVisible(t *testing.T) {
	f := newT30(t)
	f.pool.Exec(context.Background(), `update reports set visibility = 'private' where id = $1`, t30Report)
	code, _, body := f.get(t, user(f.stranger), "")
	wantNotFound(t, code, body, "not_found")
}

func TestFightCoach_FlaggedAnswersFlagged(t *testing.T) {
	f := newT30(t)
	f.pool.Exec(context.Background(), `update reports set flagged = 'tampered' where id = $1`, t30Report)
	code, _, body := f.get(t, user(f.premium), "")
	wantNotFound(t, code, body, "flagged")
	f.noJobNoCall(t)
}

func TestFightCoach_FreeOwnerGetsOnlyTheTeaser(t *testing.T) {
	f := newT30(t)
	code, _, body := f.get(t, user(f.freeOwner), "")
	if code != http.StatusOK || body["access"] != "teaser" || body["finding_count"] == nil || body["spec"] != "Fury" || len(body) != 5 {
		t.Fatalf("%d %v", code, body)
	}
	if _, leaks := body["review"]; leaks {
		t.Fatal("the gate is on the server")
	}
	first, ok := body["fix_first"].(map[string]any)
	if !ok || first["title"] == "" || first["code"] == "" || len(first) != 3 {
		t.Fatalf("fix_first carries title, code and spell_id: %v", body["fix_first"])
	}
	if _, has := first["spell_id"]; !has {
		t.Fatalf("spell_id is sent, null when there is none: %v", first)
	}
	f.noJobNoCall(t)
}

func TestFightCoach_PremiumMissQueuesAndAnswersWriting(t *testing.T) {
	f := newT30(t)
	code, _, body := f.get(t, user(f.premium), "")
	if code != http.StatusOK || body["status"] != "writing" || body["source"] != "rules" || body["access"] != "full" || body["busy"] != false {
		t.Fatalf("%d %v", code, body)
	}
	ran := f.jobs.Ran()
	want := []string{ReviewJobCommand, t30Report, "1", f.furyKey, strconv.FormatInt(f.premium, 10)}
	if len(ran) != 1 || len(ran[0]) != 6 || !slices.Equal(ran[0][:5], want) || ran[0][5] == "" {
		t.Fatalf("one coach-review job with the key, the account and the lease holder: %v", ran)
	}
	f.get(t, user(f.premium), "")
	if len(f.jobs.Ran()) != 1 {
		t.Fatalf("the lease holds: no second job %v", f.jobs.Ran())
	}
	if persona := body["persona"].(map[string]any); persona["label"] != "Your Fury review" {
		t.Fatalf("the rule-only rendering claims no coach: %v", persona)
	}
}

func TestFightCoach_LineFormNeverQueues(t *testing.T) {
	f := newT30(t)
	code, _, body := f.get(t, user(f.premium), "&line=1")
	if code != http.StatusOK || body["status"] != "ready" || body["source"] != "rules" {
		t.Fatalf("%d %v", code, body)
	}
	var leases int
	f.pool.QueryRow(context.Background(), `select count(*) from coach_review_leases where report_id = $1`, t30Report).Scan(&leases)
	if leases != 0 {
		t.Fatalf("the line form takes no lease: %d", leases)
	}
	f.noJobNoCall(t)
	f.store(t, SourceModel, false)
	_, _, body = f.get(t, user(f.premium), "&line=1")
	if body["source"] != "model" {
		t.Fatalf("the stored review's line: %v", body)
	}
}

// passingDoc is a model answer the validator accepts for the pack: the rule-only rendering.
func passingDoc(t *testing.T, f *t30) string {
	t.Helper()
	p, err := f.svc.PackFor(context.Background(), f.rep(t), t30Fight, f.furyKey)
	if err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(RenderRules(p, fixturePersona(t).Drills))
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func TestFightCoach_ServesTheStoredModelReview(t *testing.T) {
	f := newT30(t)
	f.llm.Responses = []llm.Response{{Content: passingDoc(t, f), Model: "m"}}
	args := []string{t30Report, "1", f.furyKey, strconv.FormatInt(f.premium, 10), "holder-1"}
	if err := RunReviewJob(context.Background(), f.svc, args); err != nil {
		t.Fatal(err)
	}
	stored, ok, err := f.svc.Store.Review(context.Background(), f.key(t))
	if err != nil || !ok || stored.Source != SourceModel {
		t.Fatalf("the job stored a model review: %+v %v %v", stored.Validation, ok, err)
	}
	_, _, body := f.get(t, user(f.premium), "")
	persona := body["persona"].(map[string]any)
	if body["status"] != "ready" || body["source"] != "model" || persona["label"] != "Your Fury coach" || persona["status"] != "full" {
		t.Fatalf("%v", body)
	}
	if body["reviewed_at"] != capsNow.Format(time.RFC3339) || len(f.jobs.Ran()) != 0 {
		t.Fatalf("reviewed_at %v jobs %v", body["reviewed_at"], f.jobs.Ran())
	}
}

func TestFightCoach_ModelReviewIsServedOverAnyRulesSave(t *testing.T) {
	f := newT30(t)
	k := f.store(t, SourceModel, false)
	f.svc.Store.SaveReview(context.Background(), StoredReview{Key: k, Doc: exampleReview(), Source: SourceRules,
		Retryable: true, GeneratedAt: capsNow})
	_, _, body := f.get(t, user(f.premium), "")
	if body["source"] != "model" || body["review"].(map[string]any)["opening"] != "stored model" || len(f.jobs.Ran()) != 0 {
		t.Fatalf("%v jobs %v", body, f.jobs.Ran())
	}
}

func TestFightCoach_SwitchOffServesRulesNotTheStoredModelReview(t *testing.T) {
	f := newT30(t)
	f.store(t, SourceModel, false)
	f.svc.Enabled = false
	_, _, body := f.get(t, user(f.premium), "")
	if body["source"] != "rules" || body["status"] != "ready" || body["busy"] != false ||
		body["persona"].(map[string]any)["label"] != "Your Fury review" {
		t.Fatalf("%v", body)
	}
	f.noJobNoCall(t)
}

func TestFightCoach_CapHitIsBusyAndQueuesNothing(t *testing.T) {
	f := newT30(t)
	if _, err := f.pool.Exec(context.Background(), `insert into llm_spend (day, usd) values ($1, 25)`, capsNow); err != nil {
		t.Fatal(err)
	}
	_, _, body := f.get(t, user(f.premium), "")
	if _, ok := body["cap"]; body["busy"] != true || body["source"] != "rules" || body["status"] != "ready" || ok {
		t.Fatalf("%v", body)
	}
	f.noJobNoCall(t)
}

func TestFightCoach_AccountCapIsBusy(t *testing.T) {
	f := newT30(t)
	f.svc.Caps.PerAccount = 1
	if v, err := f.svc.Caps.Admit(context.Background(), f.premium); err != nil || v.Refusal != RefusalNone {
		t.Fatal(v, err)
	}
	_, _, body := f.get(t, user(f.premium), "")
	want := map[string]any{"period": "day", "until": "2026-10-28"}
	if body["busy"] != true || len(f.jobs.Ran()) != 0 || !reflect.DeepEqual(body["cap"], want) {
		t.Fatalf("%v", body)
	}
}

func TestFightCoach_AccountMonthlyCapSaysWhenItLifts(t *testing.T) {
	f := newT30(t)
	if _, err := f.pool.Exec(context.Background(), `insert into coach_account_day (day, user_id, usd) values ('2026-10-03', $1, 1)`,
		f.premium); err != nil {
		t.Fatal(err)
	}
	_, _, body := f.get(t, user(f.premium), "")
	want := map[string]any{"period": "month", "until": "2026-11-01"}
	if body["busy"] != true || body["source"] != "rules" || !reflect.DeepEqual(body["cap"], want) {
		t.Fatalf("%v", body)
	}
	f.noJobNoCall(t)
}

func TestFightCoach_RetryableRulesRowRegeneratesWhenAllowed(t *testing.T) {
	f := newT30(t)
	f.store(t, SourceRules, true)
	_, _, body := f.get(t, user(f.premium), "")
	if body["status"] != "writing" || len(f.jobs.Ran()) != 1 {
		t.Fatalf("%v jobs %v", body, f.jobs.Ran())
	}
}

func TestFightCoach_ValidationFailureRowIsNotRetried(t *testing.T) {
	f := newT30(t)
	f.store(t, SourceRules, false)
	_, _, body := f.get(t, user(f.premium), "")
	if body["status"] != "ready" || body["source"] != "rules" || body["review"].(map[string]any)["opening"] != "stored rules" {
		t.Fatalf("%v", body)
	}
	f.noJobNoCall(t)
}

func TestFightCoach_OfficerOfTheGuildSeesFullMarkedOfficer(t *testing.T) {
	f := newT30(t)
	_, _, body := f.get(t, user(f.officerA), "")
	if body["access"] != "officer" || body["review"] == nil {
		t.Fatalf("%v", body)
	}
	ran := f.jobs.Ran()
	if len(ran) != 1 || ran[0][4] != strconv.FormatInt(f.officerA, 10) {
		t.Fatalf("the generation counts against the officer: %v", ran)
	}
}

func TestFightCoach_OfficerOfAnotherGuildGets404(t *testing.T) {
	f := newT30(t)
	code, _, body := f.get(t, user(f.officerB), "")
	wantNotFound(t, code, body, "not_found")
	if body["message"] != msgNotFound {
		t.Fatalf("%v", body)
	}
	f.noJobNoCall(t)
}

func TestFightCoach_NotAVerifiedMemberGets404ForOfficer(t *testing.T) {
	f := newT30(t)
	f.linkCharacter(t, f.freeOwner, false)
	code, _, body := f.get(t, user(f.officerA), "")
	wantNotFound(t, code, body, "not_found")
	f.noJobNoCall(t)
}

func TestFightCoach_ModeratorSeesFullExceptAnonymised(t *testing.T) {
	f := newT30(t)
	mod := auth.Actor{UserID: f.mod, Role: "moderator", Method: "session"}
	code, _, body := f.get(t, mod, "")
	if code != http.StatusOK || body["access"] != "officer" || body["review"] == nil {
		t.Fatalf("%d %v", code, body)
	}
	f.pool.Exec(context.Background(), `update users set anonymize = true where id = $1`, f.freeOwner)
	code, _, body = f.get(t, mod, "")
	wantNotFound(t, code, body, "not_found")
	if _, _, body = f.get(t, user(f.officerA), ""); body["access"] != "officer" {
		t.Fatalf("the guild's officers still see an anonymised member: %v", body)
	}
}

func TestFightCoach_NoStoreHeader(t *testing.T) {
	f := newT30(t)
	for _, actor := range []auth.Actor{user(f.premium), user(f.freeOwner), user(f.stranger)} {
		_, h, _ := f.get(t, actor, "")
		if h.Get("Cache-Control") != "private, no-store" {
			t.Fatalf("actor %d: %q", actor.UserID, h.Get("Cache-Control"))
		}
	}
}

func TestFightCoach_FullShapeRoundTrips(t *testing.T) {
	f := newT30(t)
	_, _, body := f.get(t, user(f.premium), "")
	for _, field := range []string{"report_id", "fight_index", "player_key", "spec", "role", "size", "kill", "status",
		"source", "busy", "access", "persona", "reviewed_at", "review", "anchors", "basis_labels"} {
		if _, ok := body[field]; !ok {
			t.Errorf("missing %s", field)
		}
	}
	review := body["review"].(map[string]any)
	fix, ok := review["fixFirst"].(map[string]any)
	if !ok {
		t.Fatalf("a fix first: %v", review)
	}
	anchors := body["anchors"].(map[string]any)
	for _, id := range fix["anchors"].([]any) {
		if _, ok := anchors[id.(string)]; !ok {
			t.Errorf("anchor %v is not in the anchors map", id)
		}
	}
	for _, l := range fix["links"].([]any) {
		if l.(map[string]any)["sentence"] == nil {
			t.Errorf("a rule-only link carries its sentence: %v", l)
		}
	}
}

func TestFightCoach_UnknownPlayerIs404(t *testing.T) {
	f := newT30(t)
	f.furyKey = "us/normal/nobody"
	f.owners[f.premium]["us/normal/nobody"] = true
	code, _, body := f.get(t, user(f.premium), "")
	wantNotFound(t, code, body, "not_found")
}

// TestFightCoach_PersonaLabelsAreTheWebCopy pins the labels one way to the experience spec's
// §8.2 web copy rows; they are not §8.1 pack copy.
func TestFightCoach_PersonaLabelsAreTheWebCopy(t *testing.T) {
	rows := experienceRows(t, "### 8.2 Web copy")
	for key, want := range map[string]string{
		"titleModel": PersonaLabelModel, "titleStarter": PersonaLabelStarter, "titleRules": PersonaLabelRules,
	} {
		if cell := rows[key]; len(cell) == 0 || cell[0] != want {
			t.Errorf("%s: §8.2 says %q, copy.go has %q", key, cell, want)
		}
	}
}

func TestServiceCapsPeek_DecidesWithoutCounting(t *testing.T) {
	f := newT30(t)
	caps := f.svc.Caps
	caps.PerAccount = 1
	ctx := context.Background()
	for i := 0; i < 2; i++ {
		if v, err := caps.Peek(ctx, f.premium); err != nil || v.Refusal != RefusalNone {
			t.Fatalf("peek %d: %+v %v", i, v, err)
		}
	}
	if _, err := caps.Admit(ctx, f.premium); err != nil {
		t.Fatal(err)
	}
	if v, _ := caps.Peek(ctx, f.premium); v.Refusal != RefusalAccountCap || v.Period != CapPeriodDay {
		t.Fatalf("%+v", v)
	}
	caps.Enabled = false
	if v, _ := caps.Peek(ctx, f.premium); v.Refusal != RefusalSwitchOff {
		t.Fatalf("%+v", v)
	}
}
