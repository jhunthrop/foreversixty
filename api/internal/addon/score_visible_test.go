package addon

import (
	"fmt"
	"testing"
)

// scoreVisible reads characters.score_visible for the contract's character.
func (h *outboxHarness) scoreVisible() bool {
	h.t.Helper()
	var visible bool
	if err := h.pool.QueryRow(h.t.Context(), `select score_visible from characters where key = $1`, obnoxiousKey).Scan(&visible); err != nil {
		h.t.Fatal(err)
	}
	return visible
}

const scoreOutbox = `{"format":1,"export":{"character":"Obnoxious-Yell","code":"FS1:1.60.1.70291:warrior:human:0/0/0:","saved_at":"2026-11-08 20:00"%s}}`

// Contract 4's export.score_visible: false hides the character's score, absent shows it again.
func TestTheOutboxCarriesTheScoreChoice(t *testing.T) {
	h := newOutboxHarness(t, true)
	h.post(withOutbox(fmt.Sprintf(scoreOutbox, "")))
	if !h.scoreVisible() {
		t.Fatal("a score is visible by default")
	}
	if out := h.post(withOutbox(fmt.Sprintf(scoreOutbox, `,"score_visible":false`))); out.Outbox == nil || out.Outbox.Export != OutboxStored {
		t.Fatalf("outbox = %+v", out.Outbox)
	}
	if h.scoreVisible() {
		t.Fatal("score_visible false did not hide the score")
	}
	h.post(withOutbox(fmt.Sprintf(scoreOutbox, "")))
	if !h.scoreVisible() {
		t.Fatal("an outbox without score_visible did not show the score again")
	}
}
