package addon

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// The inbox tells the addon which outbox entries this companion uploads (contract 3, accepts),
// so the panel offers its Tonight answer only when it will be sent.
func TestTheInboxSaysTheCompanionAcceptsRaidAnswers(t *testing.T) {
	top := parseInbox(t, string(RenderInbox(Inbox{}, t0)))
	accepts, ok := top.Fields["accepts"].(*Table)
	if !ok || len(accepts.Items) != 1 || accepts.Items[0] != OutboxRaidAnswer {
		t.Fatalf("accepts = %+v", top.Fields["accepts"])
	}
}

// The API's raid_signup message (docs/contracts/fixtures/inbox-raid-signup.json) is flat, so it
// is carried through and rendered whole.
func TestTheRaidSignupFixtureIsCarriedThrough(t *testing.T) {
	raw := contractFixture(t, "inbox-raid-signup.json")
	var in Inbox
	if err := json.Unmarshal([]byte(`{"format":2,"generated_at":"2026-11-04T21:03:00Z","builds":[],"messages":[`+string(raw)+`]}`), &in); err != nil {
		t.Fatal(err)
	}
	if len(in.Messages) != 1 || len(in.Skipped) != 0 {
		t.Fatalf("messages = %+v, skipped = %+v", in.Messages, in.Skipped)
	}
	row := parseInbox(t, string(RenderInbox(in, t0))).Fields["messages"].(*Table).Items[0].(*Table)
	if row.Get("type") != "raid_signup" || row.Get("answer") != "late" || row.Get("role") != "dps" ||
		keyOf(row.Fields["raid_id"]) != "42" || keyOf(row.Fields["late_minutes"]) != "15" ||
		row.Get("answered_at") != "2026-12-08T19:04:00Z" {
		t.Errorf("rendered = %+v", row.Fields)
	}
}

// savedWithAnswers is a SavedVariables file whose outbox holds three answers, beside keys the
// companion must never touch.
const savedWithAnswers = `
ForeverSixtyDB = {
	["characters"] = {
		["Obnoxious-Normal"] = {
			["name"] = "Obnoxious",
			["export"] = "FS1:1.60.1.70291:warrior:human:371000000/0/0:chest=6627",
		},
	},
	["outbox"] = {
		["format"] = 1,
		["export"] = {
			["character"] = "Obnoxious-Yell",
			["code"] = "FS1:1.60.1.70291:warrior:human:371000000/0/0:chest=6627",
			["saved_at"] = "2026-11-04 14:02",
		},
		["raid_answer"] = {
			{
				["character"] = "Obnoxious-Yell",
				["raid_id"] = 42,
				["answer"] = "in",
				["answered_at"] = 1796753040,
			}, -- [1]
			{
				["character"] = "Obnoxious-Yell",
				["raid_id"] = 43,
				["answer"] = "late",
				["late_minutes"] = 15,
				["answered_at"] = 1796753041,
			}, -- [2]
			{
				["character"] = "Obnoxious-Yell",
				["raid_id"] = 44,
				["answer"] = "out",
				["answered_at"] = 1796753042,
			}, -- [3]
		},
	},
	["settings"] = { ["panel"] = true },
}
ForeverSixtyCharDB = { ["seen"] = 3 }
`

func writeAddonSaved(t *testing.T, body string) string {
	t.Helper()
	p := filepath.Join(t.TempDir(), SavedVariablesName+".lua")
	if err := os.WriteFile(p, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
	return p
}

// withoutRaidAnswers is a parsed file with ForeverSixtyDB.outbox.raid_answer taken out.
func withoutRaidAnswers(t *testing.T, src string) map[string]any {
	t.Helper()
	top, err := ParseLua(src)
	if err != nil {
		t.Fatal(err)
	}
	delete(top["ForeverSixtyDB"].(*Table).Fields["outbox"].(*Table).Fields, OutboxRaidAnswer)
	return top
}

// Only the answers the site took are removed, and nothing else in the file changes.
func TestClearRaidAnswersRemovesOnlyTheAnsweredEntries(t *testing.T) {
	p := writeAddonSaved(t, savedWithAnswers)
	done := []RaidAnswerKey{{Character: "Obnoxious-Yell", RaidID: 42, AnsweredAt: 1796753040},
		{Character: "Obnoxious-Yell", RaidID: 44, AnsweredAt: 1796753042},
		// An answer the file no longer holds is skipped.
		{Character: "Obnoxious-Yell", RaidID: 45, AnsweredAt: 1796753043}}
	n, err := ClearRaidAnswers(p, done)
	if err != nil || n != 2 {
		t.Fatalf("cleared %d: %v", n, err)
	}
	b, _ := os.ReadFile(p)
	top, err := ParseLua(string(b))
	if err != nil {
		t.Fatalf("the file no longer parses: %v\n%s", err, b)
	}
	left := top["ForeverSixtyDB"].(*Table).Fields["outbox"].(*Table).Fields[OutboxRaidAnswer].(*Table)
	if len(left.Items) != 1 || keyOf(left.Items[0].(*Table).Fields["raid_id"]) != "43" ||
		keyOf(left.Items[0].(*Table).Fields["late_minutes"]) != "15" {
		t.Fatalf("left = %+v", left.Items)
	}
	if !reflect.DeepEqual(withoutRaidAnswers(t, string(b)), withoutRaidAnswers(t, savedWithAnswers)) {
		t.Fatalf("something else in the file changed:\n%s", b)
	}
	if !strings.Contains(string(b), "ForeverSixtyCharDB = { [\"seen\"] = 3 }") {
		t.Fatalf("the rest of the file was rewritten:\n%s", b)
	}

	// Nothing left to clear writes nothing.
	before, _ := os.Stat(p)
	if n, err := ClearRaidAnswers(p, done); err != nil || n != 0 {
		t.Fatalf("second clear %d %v", n, err)
	}
	after, _ := os.Stat(p)
	if !after.ModTime().Equal(before.ModTime()) {
		t.Fatal("a clear with nothing to remove wrote the file")
	}
}

// Every answer gone leaves an empty list; a file with no answers is left alone; only the addon's
// own SavedVariables file is ever written.
func TestClearRaidAnswersEdges(t *testing.T) {
	p := writeAddonSaved(t, savedWithAnswers)
	all := []RaidAnswerKey{{"Obnoxious-Yell", 42, 1796753040}, {"Obnoxious-Yell", 43, 1796753041}, {"Obnoxious-Yell", 44, 1796753042}}
	if n, err := ClearRaidAnswers(p, all); err != nil || n != 3 {
		t.Fatalf("cleared %d: %v", n, err)
	}
	b, _ := os.ReadFile(p)
	top, err := ParseLua(string(b))
	if err != nil {
		t.Fatal(err)
	}
	if left := top["ForeverSixtyDB"].(*Table).Fields["outbox"].(*Table).Fields[OutboxRaidAnswer].(*Table); len(left.Items) != 0 || len(left.Fields) != 0 {
		t.Fatalf("left = %+v", left)
	}

	plain := writeAddonSaved(t, `ForeverSixtyDB = { ["outbox"] = { ["format"] = 1 } }`)
	if n, err := ClearRaidAnswers(plain, all); err != nil || n != 0 {
		t.Fatalf("no answers: %d %v", n, err)
	}
	other := filepath.Join(t.TempDir(), "Other.lua")
	if err := os.WriteFile(other, []byte(savedWithAnswers), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := ClearRaidAnswers(other, all); err == nil {
		t.Fatal("wrote a file that is not the addon's")
	}
}

// The outbox's key moves with its answers, so a new answer is sent even when nothing else in
// the outbox changed.
func TestTheOutboxKeyFollowsItsAnswers(t *testing.T) {
	a, err := ParseLua(savedWithAnswers)
	if err != nil {
		t.Fatal(err)
	}
	b, err := ParseLua(strings.Replace(savedWithAnswers, "1796753042", "1796753099", 1))
	if err != nil {
		t.Fatal(err)
	}
	oa, _ := outboxOf(a)
	ob, _ := outboxOf(b)
	if oa == nil || ob == nil || oa.key() == ob.key() || oa.RaidAnswerHash == "" {
		t.Fatalf("keys %q %q", oa.key(), ob.key())
	}
}
