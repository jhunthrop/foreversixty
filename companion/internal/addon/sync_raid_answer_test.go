// companion/internal/addon/sync_raid_answer_test.go
// The raid answer loop through the real Sync against the fakeapi server: the answers ride up in
// the outbox, and once the folder is quiet the ones the site accepted or refused come out of
// the game's file while everything else in it stays; a deferred one stays to go up again.
package addon_test

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/addon"
)

// answersSaved is outboxSaved with one raid answer in its outbox.
func answersSaved() string {
	return strings.Replace(outboxSaved("2026-11-04 14:10", 1), `		["roster_snapshot"] = {`, `		["raid_answer"] = {
			{
				["character"] = "Obnoxious-Yell",
				["raid_id"] = 42,
				["answer"] = "in",
				["answered_at"] = 1796753040,
			}, -- [1]
		},
		["roster_snapshot"] = {`, 1)
}

func raidAnswersIn(t *testing.T, path string) []any {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	top, err := addon.ParseLua(string(b))
	if err != nil {
		t.Fatal(err)
	}
	ob := top["ForeverSixtyDB"].(*addon.Table).Fields["outbox"].(*addon.Table)
	list, ok := ob.Fields["raid_answer"].(*addon.Table)
	if !ok {
		t.Fatalf("no raid_answer in the file:\n%s", b)
	}
	if ob.Get("export") != "" || ob.Fields["roster_snapshot"] == nil || ob.Fields["sim_request"] == nil {
		t.Fatalf("the rest of the outbox changed:\n%s", b)
	}
	return list.Items
}

func TestAnAnsweredRaidAnswerComesOutOfTheFile(t *testing.T) {
	r := newRig(t, answersSaved())
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	posts := r.posts()
	if len(posts) != 1 {
		t.Fatalf("posts = %d", len(posts))
	}
	var ob struct {
		RaidAnswer []map[string]any `json:"raid_answer"`
	}
	if err := json.Unmarshal(posts[0].Outbox, &ob); err != nil || len(ob.RaidAnswer) != 1 || ob.RaidAnswer[0]["raid_id"] != float64(42) {
		t.Fatalf("the answer did not go up: %s", posts[0].Outbox)
	}
	if !strings.Contains(r.logs.String(), "Raid answers: 1 taken, 0 refused, 0 to send again") {
		t.Errorf("the log does not say what became of the answer:\n%s", r.logs.String())
	}

	// Once the folder is quiet the answer comes out, and the companion's own edit is not
	// uploaded again.
	r.at = r.at.Add(10 * time.Second)
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	if left := raidAnswersIn(t, r.path); len(left) != 0 {
		t.Fatalf("left %+v", left)
	}
	r.at = r.at.Add(10 * time.Second)
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	if n := len(r.posts()); n != 1 {
		t.Fatalf("the companion's own edit went up again: %d posts", n)
	}
}

func TestADeferredRaidAnswerStays(t *testing.T) {
	r := newRig(t, answersSaved())
	r.srv.SetRaidAnswerStatus(addon.RaidAnswerDeferred)
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	r.at = r.at.Add(10 * time.Second)
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	if left := raidAnswersIn(t, r.path); len(left) != 1 {
		t.Fatalf("a deferred answer was removed: %+v", left)
	}
}
