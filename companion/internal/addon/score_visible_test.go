package addon

import (
	"encoding/json"
	"strings"
	"testing"
)

// Contract 4's export.score_visible goes up as the addon wrote it: false when the player hid the
// Forever Score, absent otherwise.
func TestTheOutboxForwardsTheScoreChoice(t *testing.T) {
	read := func(src string) map[string]any {
		t.Helper()
		top, err := ParseLua(src)
		if err != nil {
			t.Fatal(err)
		}
		ob, err := outboxOf(top)
		if err != nil || ob == nil {
			t.Fatalf("outbox %v %v", ob, err)
		}
		var body struct {
			Export map[string]any `json:"export"`
		}
		if err := json.Unmarshal(ob.Body, &body); err != nil {
			t.Fatal(err)
		}
		return body.Export
	}
	hidden := strings.Replace(savedWithAnswers, `["saved_at"] = "2026-11-04 14:02",`,
		`["saved_at"] = "2026-11-04 14:02",
			["score_visible"] = false,`, 1)
	if v, ok := read(hidden)["score_visible"]; !ok || v != false {
		t.Fatalf("score_visible = %v (%v)", v, ok)
	}
	if _, ok := read(savedWithAnswers)["score_visible"]; ok {
		t.Fatal("an export that says nothing carried score_visible")
	}
}
