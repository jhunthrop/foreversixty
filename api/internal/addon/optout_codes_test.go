// api/internal/addon/optout_codes_test.go
package addon

import (
	"encoding/json"
	"testing"
)

// The snapshot's opt-out codes keep only well-formed codes, each under the member's plain name.
func TestReadRosterKeepsWellFormedOptOutCodes(t *testing.T) {
	raw := json.RawMessage(`{"guild":"Ember Wardens","members":[],"optout_codes":[
		{"name":"Vela Thornbrook-PvPUS","code":"FS-7K2QXM"},
		{"name":"Orrin Ashmantle","code":"FS-0O1I00"},
		{"name":"<b>bad</b>","code":"FS-7K2QXM"},
		{"name":"Sel Duskwater","code":"remove me please"}]}`)
	snap, _, invalid := readRoster(raw)
	if invalid || snap == nil {
		t.Fatalf("snapshot invalid")
	}
	if len(snap.OptOutCodes) != 1 || snap.OptOutCodes[0].Name != "Vela Thornbrook" || snap.OptOutCodes[0].Code != "FS-7K2QXM" {
		t.Fatalf("codes = %+v", snap.OptOutCodes)
	}
}
