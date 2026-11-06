package addon

import (
	"encoding/json"
	"testing"
)

// The API's raid_assignment message (docs/contracts/fixtures/inbox-raid-assignment.json) is a
// type this build does not know: it is flat, so it is carried through and rendered whole.
func TestTheRaidAssignmentFixtureIsCarriedThrough(t *testing.T) {
	raw := contractFixture(t, "inbox-raid-assignment.json")
	var in Inbox
	if err := json.Unmarshal([]byte(`{"format":2,"generated_at":"2026-11-04T21:03:00Z","builds":[],"messages":[`+string(raw)+`]}`), &in); err != nil {
		t.Fatal(err)
	}
	if len(in.Messages) != 1 || len(in.Skipped) != 0 {
		t.Fatalf("messages = %+v, skipped = %+v", in.Messages, in.Skipped)
	}
	row := parseInbox(t, string(RenderInbox(in, t0))).Fields["messages"].(*Table).Items[0].(*Table)
	if row.Get("type") != "raid_assignment" || row.Get("character") != "us/normal/obnoxious" || row.Get("slot") != "off_tank" ||
		row.Get("warning") != "OT" || row.Get("with") != "Thornhide" || keyOf(row.Fields["raid_id"]) != "42" ||
		row.Get("starts_at") != "2026-12-10T02:00:00Z" {
		t.Errorf("rendered = %+v", row.Fields)
	}
}
