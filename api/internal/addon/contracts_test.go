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
	"strings"
	"testing"
)

// contractTypes are contract 3's eight message types, every one of which MessageEntry decodes.
var contractTypes = []string{
	messageBuild, messageUpgrade, messageWeights, messageSimResult,
	messageLastRaid, messageLootValues, messageGuild, messageAck,
}

// contractCharacter is the fixture's character key.
const contractCharacter = "us/normal/obnoxious"

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

// contractSorted is raw re-encoded with every object's keys sorted and its whitespace
// removed, so two encodings of one value compare equal whatever order a struct lists its
// fields in.
func contractSorted(t *testing.T, raw []byte) []byte {
	t.Helper()
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	var v any
	if err := dec.Decode(&v); err != nil {
		t.Fatalf("decode %s: %v", raw, err)
	}
	out, err := json.Marshal(v) // a map's keys marshal sorted
	if err != nil {
		t.Fatal(err)
	}
	return out
}

// contractRoundTrip marshals v and requires it to equal raw once both have their keys sorted.
func contractRoundTrip(t *testing.T, raw []byte, v any) {
	t.Helper()
	got, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if g, w := contractSorted(t, got), contractSorted(t, raw); !bytes.Equal(g, w) {
		t.Errorf("round trip changed the value\n got: %s\nwant: %s", g, w)
	}
}

func TestInboxV2FixtureGolden(t *testing.T) {
	raw := contractFixture(t, "inbox-v2.json")

	t.Run("the whole body round trips through InboxResponse", func(t *testing.T) {
		var body InboxResponse
		contractStrict(t, raw, &body)
		contractRoundTrip(t, raw, body)
		if body.Format != InboxFormat || body.GeneratedAt != "2026-11-04T21:03:00Z" {
			t.Errorf("envelope = format %d, generated_at %q", body.Format, body.GeneratedAt)
		}
		if len(body.Builds) != 1 || body.Builds[0].Character != contractCharacter {
			t.Errorf("builds = %+v", body.Builds)
		}
	})

	var envelope struct {
		Format      int               `json:"format"`
		GeneratedAt string            `json:"generated_at"`
		Builds      []json.RawMessage `json:"builds"`
		Messages    []json.RawMessage `json:"messages"`
	}
	contractStrict(t, raw, &envelope)

	t.Run("every message strictly decodes into MessageEntry and round trips", func(t *testing.T) {
		seen := map[string]int{}
		for _, msg := range envelope.Messages {
			var m MessageEntry
			contractStrict(t, msg, &m)
			contractRoundTrip(t, msg, m)
			if m.Character != contractCharacter {
				t.Errorf("%s is addressed to %q, want the character key %q", m.Type, m.Character, contractCharacter)
			}
			seen[m.Type]++
		}
		for _, typ := range contractTypes {
			if seen[typ] != 1 {
				t.Errorf("fixture has %d %q messages, want 1", seen[typ], typ)
			}
		}
		if len(envelope.Messages) != len(contractTypes) {
			t.Errorf("the fixture has %d messages and contract 3 names %d types", len(envelope.Messages), len(contractTypes))
		}
	})

	byType := map[string]MessageEntry{}
	for _, msg := range envelope.Messages {
		var m MessageEntry
		contractStrict(t, msg, &m)
		byType[m.Type] = m
	}

	t.Run("the v2 fields read as contract 3 says", func(t *testing.T) {
		if u := byType[messageUpgrade]; u.Unit != "dps" {
			t.Errorf("upgrade unit = %q, want dps", u.Unit)
		}
		if w := byType[messageWeights]; w.DPSPerPoint == nil || *w.DPSPerPoint != 0.42 {
			t.Errorf("weights dps_per_point = %v, want 0.42", w.DPSPerPoint)
		}
		// The weights say when the run that made them was queued and which export it ran on:
		// the same export_seen_at the ack carries for that export.
		if w := byType[messageWeights]; w.GeneratedAt != "2026-11-04T14:11:00Z" || w.ExportSeenAt != byType[messageAck].ExportSeenAt {
			t.Errorf("weights generated_at %q, export_seen_at %q", w.GeneratedAt, w.ExportSeenAt)
		}
		if s := byType[messageSimResult]; s.DPS == nil || *s.DPS != 412.6 || s.Error == nil || *s.Error != 3.1 || s.Slots == nil ||
			len(*s.Slots) != 1 || (*s.Slots)[0] != (SimSlot{Slot: "chest", ItemID: 6627, Delta: 4.2}) {
			t.Errorf("sim_result = %+v", s)
		}
		lr := byType[messageLastRaid]
		if lr.ReportID != "r_8f2a1c" || lr.Date != "2026-12-09" || len(lr.Bosses) != 1 {
			t.Fatalf("last_raid = %+v", lr)
		}
		if b := lr.Bosses[0]; b.Name != "Lucifron" || b.Percentile != 71 || b.BiggestLoss != "downtime" ||
			b.LossSeconds != 18 || len(b.ConsumablesMissed) != 1 || b.Adherence == nil || *b.Adherence != 0.82 {
			t.Errorf("last_raid boss = %+v", b)
		}
		lv := byType[messageLootValues]
		if lv.Raid != "Molten Core" || lv.Items[17076] != (LootItemValue{Value: 6.3, Unit: "dps", Planned: true}) {
			t.Errorf("loot_values = %+v", lv)
		}
		if a := byType[messageAck]; a.ExportSeenAt != "2026-11-04T14:03:00Z" || a.SimRequestState != "done" ||
			a.RequestedAt != "2026-11-04T14:11:00Z" {
			t.Errorf("ack = %+v", a)
		}
		if b := byType[messageBuild]; b.Name != "Arms raid" || !strings.HasPrefix(b.Code, "FSB1:") {
			t.Errorf("build = %+v", b)
		}
	})

	// adherence is optional (contract 3): absent until the rotation task lands, so a
	// last_raid without it decodes strictly and stays without it.
	t.Run("a last_raid boss without adherence decodes and round trips", func(t *testing.T) {
		lr := byType[messageLastRaid]
		lr.Bosses = append([]LastRaidBoss(nil), lr.Bosses...)
		lr.Bosses[0].Adherence = nil
		b, err := json.Marshal(lr)
		if err != nil {
			t.Fatal(err)
		}
		if bytes.Contains(b, []byte("adherence")) {
			t.Fatalf("adherence is on the wire: %s", b)
		}
		var again MessageEntry
		contractStrict(t, b, &again)
		contractRoundTrip(t, b, again)
	})
}
