// companion/internal/addon/contracts_test.go
// The companion's side of the v2 golden contract tests (docs/contracts/addon-v2.md, contracts 3
// and 4). It decodes the shared fixture docs/contracts/fixtures/inbox-v2.json with this
// package's Message and Build, strictly and byte for byte, renders it through RenderInbox, and
// loads the rendered chunk under Lua 5.1, the interpreter the game runs, asserting on the table
// it builds. It reads docs/contracts/fixtures/outbox-v1.json back through ScanOutbox from a
// saved-variables file written the way the game writes one. The addon
// (addon/tests/contracts_spec.lua) and the API (api/internal/addon/contracts_test.go) pin the
// same files from their side, so a change to either fails all three until each is updated.
package addon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"sort"
	"strings"
	"testing"
)

// contractTypes are contract 3's eight message types, in the fixture's order.
var contractTypes = []string{
	MessageBuild, MessageUpgrade, MessageWeights, MessageSimResult,
	MessageLastRaid, MessageLootValues, MessageGuild, MessageAck,
}

func contractFixture(t *testing.T, name string) []byte {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("..", "..", "..", "docs", "contracts", "fixtures", name))
	if err != nil {
		t.Fatalf("fixture %s: %v", name, err)
	}
	return b
}

func contractStrict(t *testing.T, raw []byte, v any) {
	t.Helper()
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.DisallowUnknownFields()
	if err := dec.Decode(v); err != nil {
		t.Fatalf("strict decode of %s: %v", raw, err)
	}
}

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

// contractCharacter is the fixture's character key.
const contractCharacter = "us/normal/obnoxious"

func TestInboxV2FixtureGolden(t *testing.T) {
	raw := contractFixture(t, "inbox-v2.json")
	var inbox struct {
		Format      int               `json:"format"`
		GeneratedAt string            `json:"generated_at"`
		Builds      []json.RawMessage `json:"builds"`
		Messages    []json.RawMessage `json:"messages"`
	}
	contractStrict(t, raw, &inbox)
	// The body GET /v1/addon/inbox answers: contract 3's envelope, every message addressed by
	// character key (region/ruleset/name-slug), the form Build.Character has always carried.
	if inbox.Format != 2 || inbox.GeneratedAt != "2026-11-04T21:03:00Z" {
		t.Fatalf("envelope = format %d, generated_at %q", inbox.Format, inbox.GeneratedAt)
	}

	// strict is the fixture decoded field by field: every key of every message must have a
	// home in Message, or DisallowUnknownFields fails the decode.
	var strict Inbox
	strict.Format, strict.GeneratedAt = inbox.Format, inbox.GeneratedAt
	for _, raw := range inbox.Builds {
		var b Build
		contractStrict(t, raw, &b)
		strict.Builds = append(strict.Builds, b)
	}
	for _, raw := range inbox.Messages {
		var m Message
		contractStrict(t, raw, &m)
		strict.Messages = append(strict.Messages, m)
	}

	t.Run("builds round trip through Build", func(t *testing.T) {
		if len(inbox.Builds) != 1 {
			t.Fatalf("builds = %d, want 1", len(inbox.Builds))
		}
		contractRoundTrip(t, inbox.Builds[0], strict.Builds[0])
	})

	t.Run("every message type round trips through Message", func(t *testing.T) {
		var types []string
		for i, raw := range inbox.Messages {
			m := strict.Messages[i]
			contractRoundTrip(t, raw, m)
			if m.Character != contractCharacter {
				t.Errorf("%s is addressed to %q", m.Type, m.Character)
			}
			types = append(types, m.Type)
		}
		if !slices.Equal(types, contractTypes) {
			t.Errorf("fixture types = %v, want contract 3's %v", types, contractTypes)
		}
	})

	t.Run("the sync's decode reads the same as the strict one", func(t *testing.T) {
		var got Inbox
		if err := json.Unmarshal(raw, &got); err != nil {
			t.Fatal(err)
		}
		if len(got.Skipped) != 0 {
			t.Errorf("the sync skipped fixture messages: %+v", got.Skipped)
		}
		if !reflect.DeepEqual(got, strict) {
			t.Errorf("the sync's decode differs from the strict one\n got: %+v\nwant: %+v", got, strict)
		}
	})

	t.Run("renders Lua the package reads back", func(t *testing.T) {
		root := parseInbox(t, string(RenderInbox(strict, t0)))
		if keyOf(root.Fields["format"]) != "2" || root.Get("generated_at") != "2026-12-09T20:00:00Z" {
			t.Errorf("envelope = %+v", root.Fields)
		}
		builds := root.Fields["builds"].(*Table)
		if len(builds.Items) != 1 {
			t.Fatalf("builds = %+v", builds)
		}
		build := builds.Items[0].(*Table)
		if build.Get("name") != "Arms raid" || build.Get("character") != contractCharacter ||
			!strings.HasPrefix(build.Get("code"), "FSB1:") {
			t.Errorf("build = %+v", build.Fields)
		}
		messages := root.Fields["messages"].(*Table)
		// Exactly the API's messages: the companion synthesises no build message of its own
		// from the builds list.
		if len(messages.Items) != len(contractTypes) {
			t.Fatalf("rendered %d messages, fixture has %d", len(messages.Items), len(contractTypes))
		}
		byType := map[string]*Table{}
		for i, item := range messages.Items {
			m := item.(*Table)
			if m.Get("type") != contractTypes[i] || m.Get("character") != contractCharacter {
				t.Errorf("message %d = %+v, want type %s", i, m.Fields, contractTypes[i])
			}
			byType[m.Get("type")] = m
		}
		if b := byType[MessageBuild]; b.Get("name") != "Arms raid" || !strings.HasPrefix(b.Get("code"), "FSB1:") {
			t.Errorf("build message = %+v", b.Fields)
		}
		upgrade := byType[MessageUpgrade]
		if upgrade.Get("slot") != "chest" || keyOf(upgrade.Fields["item_id"]) != "6627" ||
			upgrade.Get("item_name") != "Mutant Scale Breastplate" ||
			upgrade.Get("source") != "Blackrock Depths" || keyOf(upgrade.Fields["delta"]) != "4.2" ||
			upgrade.Get("unit") != "dps" {
			t.Errorf("upgrade = %+v", upgrade.Fields)
		}
		weights := byType[MessageWeights]
		if weights.Get("spec") != "warrior-arms" || len(weights.Fields["weights"].(*Table).Items) != 3 ||
			keyOf(weights.Fields["caps"].(*Table).Items[0]) != "hit" ||
			keyOf(weights.Fields["dps_per_point"]) != "0.42" {
			t.Errorf("weights = %+v", weights.Fields)
		}
		sim := byType[MessageSimResult]
		slot := sim.Fields["slots"].(*Table).Items[0].(*Table)
		if keyOf(sim.Fields["dps"]) != "412.6" || keyOf(sim.Fields["error"]) != "3.1" ||
			sim.Get("generated_at") != "2026-11-04T20:55:00Z" || slot.Get("slot") != "chest" ||
			keyOf(slot.Fields["item_id"]) != "6627" || keyOf(slot.Fields["delta"]) != "4.2" {
			t.Errorf("sim_result = %+v, slot %+v", sim.Fields, slot.Fields)
		}
		raid := byType[MessageLastRaid]
		boss := raid.Fields["bosses"].(*Table).Items[0].(*Table)
		if raid.Get("report_id") != "r_8f2a1c" || raid.Get("date") != "2026-12-09" ||
			boss.Get("name") != "Lucifron" || keyOf(boss.Fields["percentile"]) != "71" ||
			boss.Get("biggest_loss") != "downtime" || keyOf(boss.Fields["loss_seconds"]) != "18" ||
			keyOf(boss.Fields["deaths"]) != "0" || keyOf(boss.Fields["adherence"]) != "0.82" ||
			keyOf(boss.Fields["consumables_missed"].(*Table).Items[0]) != "flask" {
			t.Errorf("last_raid = %+v, boss %+v", raid.Fields, boss.Fields)
		}
		loot := byType[MessageLootValues]
		item := loot.Fields["items"].(*Table).Fields["17076"].(*Table)
		if loot.Get("raid") != "Molten Core" || keyOf(item.Fields["value"]) != "6.3" ||
			item.Get("unit") != "dps" || item.Fields["planned"] != true {
			t.Errorf("loot_values = %+v, item %+v", loot.Fields, item.Fields)
		}
		guild := byType[MessageGuild]
		if guild.Get("guild_name") != "OLYMPUS XXVII" || guild.Get("claim_state") != "claimed" ||
			keyOf(guild.Fields["pending_approvals"]) != "2" || guild.Get("rank") != "officer" {
			t.Errorf("guild = %+v", guild.Fields)
		}
		ack := byType[MessageAck]
		if ack.Get("export_seen_at") != "2026-11-04T14:03:00Z" || ack.Get("sim_request_state") != "done" ||
			ack.Get("requested_at") != "2026-11-04T14:11:00Z" {
			t.Errorf("ack = %+v", ack.Fields)
		}
	})

	t.Run("loads under Lua 5.1 with every nested shape", func(t *testing.T) {
		body := RenderInbox(strict, t0)
		ran, ok := runLua51(t, body, `
local I = assert(ForeverSixtyInbox, "no ForeverSixtyInbox")
assert(I.format == 2, "format")
assert(I.generated_at == "2026-12-09T20:00:00Z", "generated_at")
assert(#I.builds == 1 and I.builds[1].name == "Arms raid", "builds")
local want = { "build", "upgrade", "weights", "sim_result", "last_raid", "loot_values", "guild", "ack" }
assert(#I.messages == #want, "#messages = " .. #I.messages)
for i, t in ipairs(want) do
	assert(I.messages[i].type == t, "messages[" .. i .. "].type = " .. tostring(I.messages[i].type))
	assert(I.messages[i].character == "us/normal/obnoxious", "character " .. i)
end
local m = I.messages
assert(m[1].name == "Arms raid" and m[1].code:sub(1, 5) == "FSB1:", "build")
assert(m[2].item_id == 6627 and m[2].delta == 4.2 and m[2].unit == "dps", "upgrade")
assert(#m[3].weights == 3 and m[3].weights[2].stat == "strength" and m[3].weights[2].weight == 2.2, "weights")
assert(m[3].caps[1] == "hit" and m[3].dps_per_point == 0.42, "caps, dps_per_point")
assert(m[4].dps == 412.6 and m[4].error == 3.1 and m[4].generated_at == "2026-11-04T20:55:00Z", "sim_result")
assert(#m[4].slots == 1 and m[4].slots[1].slot == "chest" and m[4].slots[1].item_id == 6627 and m[4].slots[1].delta == 4.2, "slots")
local boss = m[5].bosses[1]
assert(m[5].report_id == "r_8f2a1c" and m[5].date == "2026-12-09" and #m[5].bosses == 1, "last_raid")
assert(boss.name == "Lucifron" and boss.percentile == 71 and boss.biggest_loss == "downtime", "boss")
assert(boss.loss_seconds == 18 and boss.deaths == 0 and boss.adherence == 0.82, "boss numbers")
assert(type(boss.consumables_missed) == "table" and boss.consumables_missed[1] == "flask", "consumables_missed")
assert(m[6].raid == "Molten Core", "raid")
assert(type(m[6].items[17076].value) == "number" and m[6].items[17076].value == 6.3, "loot value")
assert(m[6].items[17076].unit == "dps" and m[6].items[17076].planned == true, "loot unit, planned")
assert(m[7].guild_name == "OLYMPUS XXVII" and m[7].pending_approvals == 2 and m[7].rank == "officer", "guild")
assert(m[8].export_seen_at == "2026-11-04T14:03:00Z" and m[8].sim_request_state == "done" and m[8].requested_at == "2026-11-04T14:11:00Z", "ack")
`)
		if ok {
			t.Logf("Lua 5.1 check ran through %q", ran)
		} else {
			t.Log("no Lua 5.1 interpreter on PATH or in WSL: only the pure-Go ParseLua check ran")
		}
	})

	t.Run("optional fields stay absent when the API leaves them out", func(t *testing.T) {
		var bare []Message
		for _, m := range strict.Messages {
			switch m.Type {
			case MessageLastRaid:
				boss := m.Bosses[0]
				boss.Adherence, boss.BiggestLoss, boss.ConsumablesMissed = nil, "", []string{}
				m.Bosses = []LastRaidBoss{boss}
			case MessageAck:
				m.SimRequestState, m.RequestedAt = "", ""
			case MessageWeights:
				m.DPSPerPoint = nil
			case MessageSimResult:
				m.Slots = &[]SimSlot{}
			case MessageUpgrade:
				m.Unit = ""
			}
			bare = append(bare, m)
		}
		body := RenderInbox(Inbox{Messages: bare}, t0)
		for _, absent := range []string{`["adherence"]`, `["biggest_loss"]`, `["sim_request_state"]`, `["requested_at"]`, `["dps_per_point"]`} {
			if strings.Contains(string(body), absent) {
				t.Errorf("the render carries %s for a message without it:\n%s", absent, body)
			}
		}
		if _, ok := runLua51(t, body, `
local m = ForeverSixtyInbox.messages
assert(m[5].bosses[1].adherence == nil and m[5].bosses[1].biggest_loss == nil, "optional boss fields")
assert(type(m[5].bosses[1].consumables_missed) == "table" and #m[5].bosses[1].consumables_missed == 0, "empty list")
assert(type(m[4].slots) == "table" and #m[4].slots == 0, "empty slots")
assert(m[8].sim_request_state == nil and m[8].requested_at == nil and m[3].dps_per_point == nil and m[2].unit == nil, "absent fields")
`); !ok {
			t.Log("no Lua 5.1 interpreter: the absent-field check ran through the render text only")
		}
	})
}

// contractLuaValue writes a decoded JSON value as the game writes a saved-variables value:
// bracketed string keys, positional list entries.
func contractLuaValue(b *strings.Builder, v any, indent string) {
	switch x := v.(type) {
	case map[string]any:
		keys := make([]string, 0, len(x))
		for k := range x {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		b.WriteString("{\n")
		for _, k := range keys {
			if x[k] == nil {
				continue
			}
			fmt.Fprintf(b, "%s\t[%s] = ", indent, quote(k))
			contractLuaValue(b, x[k], indent+"\t")
			b.WriteString(",\n")
		}
		b.WriteString(indent + "}")
	case []any:
		b.WriteString("{\n")
		for i, e := range x {
			b.WriteString(indent + "\t")
			contractLuaValue(b, e, indent+"\t")
			fmt.Fprintf(b, ", -- [%d]\n", i+1)
		}
		b.WriteString(indent + "}")
	case string:
		b.WriteString(quote(x))
	case json.Number:
		b.WriteString(x.String())
	case bool:
		fmt.Fprint(b, x)
	}
}

// TestOutboxV1FixtureGolden pins contract 4's shape from the reading side: the fixture strictly
// decodes into the fields the companion reads, and written into ForeverSixtyDB.outbox the way
// the game writes saved variables, ScanOutbox reads back the whole of it unchanged.
func TestOutboxV1FixtureGolden(t *testing.T) {
	raw := contractFixture(t, "outbox-v1.json")
	var outbox struct {
		Format int `json:"format"`
		Export struct {
			Character string `json:"character"`
			Code      string `json:"code"`
			SavedAt   string `json:"saved_at"`
		} `json:"export"`
		SimRequest *struct {
			Character   string `json:"character"`
			Spec        string `json:"spec"`
			RequestedAt string `json:"requested_at"`
			GearCode    string `json:"gear_code"`
		} `json:"sim_request"`
		RosterSnapshot *struct {
			Guild   string `json:"guild"`
			Members []struct {
				Name    string `json:"name"`
				Level   int    `json:"level"`
				Spec    string `json:"spec"`
				GearGap int    `json:"gear_gap"`
				Build   string `json:"build"`
				SeenAt  string `json:"seen_at"`
			} `json:"members"`
		} `json:"roster_snapshot"`
	}
	contractStrict(t, raw, &outbox)

	if outbox.Format != OutboxFormat {
		t.Errorf("format = %d, want %d", outbox.Format, OutboxFormat)
	}
	if outbox.Export.Character != "Obnoxious-Yell" || !strings.HasPrefix(outbox.Export.Code, "FS1:") ||
		outbox.Export.SavedAt != "2026-11-04 14:02" {
		t.Errorf("export = %+v", outbox.Export)
	}
	if outbox.SimRequest == nil || outbox.SimRequest.Spec != "warrior-arms" ||
		!strings.HasPrefix(outbox.SimRequest.GearCode, "FS1:") {
		t.Errorf("sim_request = %+v", outbox.SimRequest)
	}
	if outbox.RosterSnapshot == nil || outbox.RosterSnapshot.Guild != "OLYMPUS XXVII" ||
		len(outbox.RosterSnapshot.Members) != 1 || outbox.RosterSnapshot.Members[0].GearGap != 14 {
		t.Errorf("roster_snapshot = %+v", outbox.RosterSnapshot)
	}

	t.Run("ScanOutbox reads the whole outbox back from the saved variables", func(t *testing.T) {
		dec := json.NewDecoder(bytes.NewReader(raw))
		dec.UseNumber()
		var value any
		if err := dec.Decode(&value); err != nil {
			t.Fatal(err)
		}
		var lua strings.Builder
		lua.WriteString("ForeverSixtyDB = {\n\t[\"outbox\"] = ")
		contractLuaValue(&lua, value, "\t")
		lua.WriteString(",\n}\n")
		path := filepath.Join(t.TempDir(), SavedVariablesName+".lua")
		if err := os.WriteFile(path, []byte(lua.String()), 0o600); err != nil {
			t.Fatal(err)
		}
		ob, err := ScanOutbox(path)
		if err != nil || ob == nil {
			t.Fatalf("ScanOutbox = %v, %v\n%s", ob, err, lua.String())
		}
		var got, want any
		if err := json.Unmarshal(ob.Body, &got); err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(raw, &want); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("the outbox changed on the way through\n got: %s\nwant: %s", ob.Body, raw)
		}
		if ob.SavedAt != outbox.Export.SavedAt || ob.RequestedAt != outbox.SimRequest.RequestedAt {
			t.Errorf("stamps = %q, %q", ob.SavedAt, ob.RequestedAt)
		}
	})
}
