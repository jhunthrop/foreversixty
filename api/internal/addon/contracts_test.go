// api/internal/addon/contracts_test.go
// The API's side of the v2 golden contract test (docs/contracts/addon-v2.md,
// contract 3). The shared fixture docs/contracts/fixtures/inbox-v2.json is the
// body GET /v1/addon/inbox answers with. The addon (addon/tests/contracts_spec.lua)
// and the companion (companion/internal/addon/contracts_test.go) pin the same file from
// their side, so a change to it fails all three until each is updated.
package addon

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// contractMessageV2 is MessageEntry plus the two fields contract 3 adds to messages the
// struct already carries: "unit" on upgrade and "dps_per_point" on weights. Lane L6.3 moves
// them into MessageEntry itself and deletes this shim. (Once they are there, Go's
// shallower-field-wins rule keeps this file compiling and decoding until it is deleted.)
type contractMessageV2 struct {
	MessageEntry
	Unit        string  `json:"unit,omitempty"`
	DPSPerPoint float64 `json:"dps_per_point,omitempty"`
}

// contractSupported are the message types MessageEntry can decode today. The rest of
// contract 3's eight types are listed in contractPending.
var contractSupported = []string{messageUpgrade, messageWeights, messageGuild}

// contractPending are the contract 3 types the API cannot emit yet: MessageEntry has none of
// their fields (name/code, dps/error/slots, report_id/date/bosses, raid/items,
// export_seen_at). Lane L6.3 adds them, and this test then moves them to contractSupported.
var contractPending = []string{"build", "sim_result", "last_raid", "loot_values", "ack"}

// contractPendingKey is, per pending type, a JSON key only that type's message uses. The
// pending test fails the moment MessageEntry gains a field with one of them: that lane must then
// move the type from contractPending to a strict decode and round trip in this file.
var contractPendingKey = map[string]string{
	"build": "code", "sim_result": "slots", "last_raid": "report_id",
	"loot_values": "items", "ack": "export_seen_at",
}

// contractHasJSONKey reports whether the struct v has a field whose JSON name is key.
func contractHasJSONKey(v any, key string) bool {
	typ := reflect.TypeOf(v)
	for i := 0; i < typ.NumField(); i++ {
		if name, _, _ := strings.Cut(typ.Field(i).Tag.Get("json"), ","); name == key {
			return true
		}
	}
	return false
}

func contractFixture(t *testing.T, name string) []byte {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("..", "..", "..", "docs", "contracts", "fixtures", name))
	if err != nil {
		t.Fatalf("fixture %s: %v", name, err)
	}
	return b
}

// contractStrict decodes raw into v and refuses any field v does not have.
func contractStrict(t *testing.T, raw []byte, v any) {
	t.Helper()
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.DisallowUnknownFields()
	if err := dec.Decode(v); err != nil {
		t.Fatalf("strict decode of %s: %v", raw, err)
	}
}

// contractRoundTrip marshals v and requires the bytes to equal raw with its whitespace removed.
func contractRoundTrip(t *testing.T, raw []byte, v any) {
	t.Helper()
	var want bytes.Buffer
	if err := json.Compact(&want, raw); err != nil {
		t.Fatalf("compact %s: %v", raw, err)
	}
	got, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if !bytes.Equal(got, want.Bytes()) {
		t.Errorf("round trip changed the message\n got: %s\nwant: %s", got, want.Bytes())
	}
}

func contractContains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

func TestInboxV2FixtureGolden(t *testing.T) {
	var inbox struct {
		Builds   []json.RawMessage `json:"builds"`
		Messages []json.RawMessage `json:"messages"`
	}
	contractStrict(t, contractFixture(t, "inbox-v2.json"), &inbox)

	t.Run("builds round trip through InboxEntry", func(t *testing.T) {
		if len(inbox.Builds) != 1 {
			t.Fatalf("builds = %d, want 1", len(inbox.Builds))
		}
		for _, raw := range inbox.Builds {
			var entry InboxEntry
			contractStrict(t, raw, &entry)
			contractRoundTrip(t, raw, entry)
		}
	})

	// The type of every message, found before any subtest runs so each can run alone.
	seen := map[string]int{}
	types := make([]string, len(inbox.Messages))
	for i, raw := range inbox.Messages {
		var head struct {
			Type string `json:"type"`
		}
		if err := json.Unmarshal(raw, &head); err != nil {
			t.Fatalf("message %s: %v", raw, err)
		}
		types[i] = head.Type
		seen[head.Type]++
	}

	t.Run("supported messages round trip through MessageEntry", func(t *testing.T) {
		for i, raw := range inbox.Messages {
			if !contractContains(contractSupported, types[i]) {
				continue
			}
			var m contractMessageV2
			contractStrict(t, raw, &m)
			contractRoundTrip(t, raw, m)
			if m.Type != types[i] || m.Character != "Obnoxious-Yell" {
				t.Errorf("%s decoded as type %q for %q", types[i], m.Type, m.Character)
			}
		}
		for _, typ := range contractSupported {
			if seen[typ] != 1 {
				t.Errorf("fixture has %d %q messages, want 1", seen[typ], typ)
			}
		}
	})

	t.Run("v2 fields on supported types pending L6.3", func(t *testing.T) {
		// "unit" rides on the upgrade and "dps_per_point" on the weights message in the
		// fixture; MessageEntry has neither field yet, so contractMessageV2 carries them.
		for i, raw := range inbox.Messages {
			var m contractMessageV2
			switch types[i] {
			case messageUpgrade:
				contractStrict(t, raw, &m)
				if m.Unit != "dps" {
					t.Errorf("upgrade unit = %q, want dps", m.Unit)
				}
			case messageWeights:
				contractStrict(t, raw, &m)
				if m.DPSPerPoint != 0.42 {
					t.Errorf("weights dps_per_point = %v, want 0.42", m.DPSPerPoint)
				}
			}
		}
	})

	t.Run("v2 types pending L6.3", func(t *testing.T) {
		for _, typ := range contractPending {
			if contractHasJSONKey(MessageEntry{}, contractPendingKey[typ]) {
				t.Errorf("MessageEntry now has the %q field, so %q is no longer pending: move it from "+
					"contractPending to a strict decode and byte-equal round trip in this file (%s)",
					contractPendingKey[typ], typ, "Lane L6.3")
			}
		}
		pending := 0
		for _, typ := range contractPending {
			if seen[typ] != 1 {
				t.Errorf("fixture has %d %q messages, want 1", seen[typ], typ)
			}
			pending += seen[typ]
		}
		t.Logf("%d v2-only message types are in the fixture and not yet decodable by MessageEntry: %v",
			pending, contractPending)
		if got := len(contractSupported) + len(contractPending); got != len(inbox.Messages) {
			t.Errorf("the fixture has %d messages and the type lists name %d", len(inbox.Messages), got)
		}
	})
}
