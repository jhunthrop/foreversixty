// companion/internal/addon/contracts_test.go
// The companion's side of the v2 golden contract test (docs/contracts/addon-v2.md,
// contract 3). It decodes the shared fixture docs/contracts/fixtures/inbox-v2.json with this
// package's Message and Build, renders it through RenderInbox, and checks the rendered chunk
// loads under Lua 5.1, the interpreter the game runs. The addon
// (addon/tests/contracts_spec.lua) and the API (api/internal/addon/contracts_test.go) pin the
// same file from their side, so a change to it fails all three until each is updated.
package addon

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// contractMessageV2 is Message plus the two fields contract 3 adds to messages the struct
// already carries: "unit" on upgrade and "dps_per_point" on weights. Lane L7 moves them into
// Message itself (and renders them) and deletes this shim.
type contractMessageV2 struct {
	Message
	Unit        string  `json:"unit,omitempty"`
	DPSPerPoint float64 `json:"dps_per_point,omitempty"`
}

// contractSupported are the message types Message can decode today; contractPending are the
// rest of contract 3's eight types, which Message has none of the fields for. Lane L7 moves
// them across.
var (
	contractSupported = []string{MessageUpgrade, MessageWeights, MessageGuild}
	contractPending   = []string{"build", "sim_result", "last_raid", "loot_values", "ack"}
)

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

func contractContains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

// contractLua51 checks that body loads (parses, does not run) under a Lua 5.1 interpreter. It
// tries lua5.1 on PATH, then lua5.1 inside WSL for a Windows developer, and reports which ran;
// with neither it reports false and the caller relies on the package's own pure-Go reader.
func contractLua51(t *testing.T, body []byte) (ran string, ok bool) {
	t.Helper()
	const script = "assert(loadstring(io.read('*a')))"
	candidates := [][]string{{"lua5.1"}, {"wsl", "-e", "lua5.1"}}
	for _, argv := range candidates {
		if _, err := exec.LookPath(argv[0]); err != nil {
			continue
		}
		probe := exec.Command(argv[0], append(append([]string{}, argv[1:]...), "-v")...)
		if err := probe.Run(); err != nil {
			continue // wsl without a distro or without lua5.1: try the next route
		}
		cmd := exec.Command(argv[0], append(append([]string{}, argv[1:]...), "-e", script)...)
		cmd.Stdin = bytes.NewReader(body)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("the rendered inbox does not load under Lua 5.1 (%s): %v\n%s\n%s",
				strings.Join(argv, " "), err, out, body)
		}
		return strings.Join(argv, " "), true
	}
	return "", false
}

func TestInboxV2FixtureGolden(t *testing.T) {
	var inbox struct {
		Builds   []json.RawMessage `json:"builds"`
		Messages []json.RawMessage `json:"messages"`
	}
	contractStrict(t, contractFixture(t, "inbox-v2.json"), &inbox)

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

	// decodeAll is what the sync would hold after fetching the fixture: Builds, and a Message for
	// every entry. A type Message cannot carry yet keeps only its type and character, which is
	// what RenderInbox writes for a kind it does not know (renderMessage).
	decodeAll := func(t *testing.T) Inbox {
		t.Helper()
		var in Inbox
		for _, raw := range inbox.Builds {
			var b Build
			contractStrict(t, raw, &b)
			in.Builds = append(in.Builds, b)
		}
		for i, raw := range inbox.Messages {
			if contractContains(contractSupported, types[i]) {
				var m contractMessageV2
				contractStrict(t, raw, &m)
				in.Messages = append(in.Messages, m.Message)
				continue
			}
			var head struct {
				Character string `json:"character"`
			}
			if err := json.Unmarshal(raw, &head); err != nil {
				t.Fatalf("message %s: %v", raw, err)
			}
			in.Messages = append(in.Messages, Message{Type: types[i], Character: head.Character})
		}
		return in
	}

	t.Run("builds round trip through Build", func(t *testing.T) {
		if len(inbox.Builds) != 1 {
			t.Fatalf("builds = %d, want 1", len(inbox.Builds))
		}
		for _, raw := range inbox.Builds {
			var b Build
			contractStrict(t, raw, &b)
			contractRoundTrip(t, raw, b)
		}
	})

	t.Run("supported messages round trip through Message", func(t *testing.T) {
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

	t.Run("v2 fields on supported types pending L7", func(t *testing.T) {
		for i, raw := range inbox.Messages {
			var m contractMessageV2
			switch types[i] {
			case MessageUpgrade:
				contractStrict(t, raw, &m)
				if m.Unit != "dps" {
					t.Errorf("upgrade unit = %q, want dps", m.Unit)
				}
			case MessageWeights:
				contractStrict(t, raw, &m)
				if m.DPSPerPoint != 0.42 {
					t.Errorf("weights dps_per_point = %v, want 0.42", m.DPSPerPoint)
				}
			}
		}
	})

	t.Run("v2 types pending L7", func(t *testing.T) {
		pending := 0
		for _, typ := range contractPending {
			if seen[typ] != 1 {
				t.Errorf("fixture has %d %q messages, want 1", seen[typ], typ)
			}
			pending += seen[typ]
		}
		t.Logf("%d v2-only message types are in the fixture and not yet decodable by Message: %v",
			pending, contractPending)
		if got := len(contractSupported) + len(contractPending); got != len(inbox.Messages) {
			t.Errorf("the fixture has %d messages and the type lists name %d", len(inbox.Messages), got)
		}
	})

	t.Run("renders Lua the package reads back", func(t *testing.T) {
		body := RenderInbox(decodeAll(t), t0)
		top, err := ParseLua(string(body))
		if err != nil {
			t.Fatalf("the rendered inbox does not parse: %v\n%s", err, body)
		}
		root, ok := top[InboxGlobal].(*Table)
		if !ok {
			t.Fatalf("the inbox does not assign %s:\n%s", InboxGlobal, body)
		}

		builds := root.Fields["builds"].(*Table)
		if len(builds.Items) != 1 {
			t.Fatalf("builds = %+v", builds)
		}
		build := builds.Items[0].(*Table)
		if build.Get("name") != "Arms raid" || build.Get("character") != "Obnoxious-Yell" ||
			!strings.HasPrefix(build.Get("code"), "FSB1:") {
			t.Errorf("build = %+v", build.Fields)
		}

		messages := root.Fields["messages"].(*Table)
		if len(messages.Items) != len(types) {
			t.Fatalf("rendered %d messages, fixture has %d", len(messages.Items), len(types))
		}
		for i, item := range messages.Items {
			m := item.(*Table)
			if m.Get("type") != types[i] || m.Get("character") != "Obnoxious-Yell" {
				t.Errorf("message %d = %+v, want type %s", i, m.Fields, types[i])
			}
		}
		byType := map[string]*Table{}
		for _, item := range messages.Items {
			m := item.(*Table)
			byType[m.Get("type")] = m
		}
		upgrade := byType[MessageUpgrade]
		if upgrade.Get("slot") != "chest" || keyOf(upgrade.Fields["item_id"]) != "6627" ||
			upgrade.Get("item_name") != "Mutant Scale Breastplate" ||
			upgrade.Get("source") != "Blackrock Depths" || keyOf(upgrade.Fields["delta"]) != "4.2" {
			t.Errorf("upgrade = %+v", upgrade.Fields)
		}
		weights := byType[MessageWeights]
		if weights.Get("spec") != "warrior-arms" || len(weights.Fields["weights"].(*Table).Items) != 3 ||
			keyOf(weights.Fields["caps"].(*Table).Items[0]) != "hit" {
			t.Errorf("weights = %+v", weights.Fields)
		}
		guild := byType[MessageGuild]
		if guild.Get("guild_name") != "OLYMPUS XXVII" || guild.Get("claim_state") != "claimed" ||
			keyOf(guild.Fields["pending_approvals"]) != "2" || guild.Get("rank") != "officer" {
			t.Errorf("guild = %+v", guild.Fields)
		}
		// A v2-only type renders as its envelope alone until Message carries its fields.
		if len(byType["sim_result"].Fields) != 2 {
			t.Errorf("sim_result = %+v, want type and character only", byType["sim_result"].Fields)
		}
	})

	t.Run("rendered chunk loads under Lua 5.1", func(t *testing.T) {
		body := RenderInbox(decodeAll(t), t0)
		// The package's own reader is the check that always runs.
		if _, err := ParseLua(string(body)); err != nil {
			t.Fatalf("the rendered inbox does not parse: %v", err)
		}
		if ran, ok := contractLua51(t, body); ok {
			t.Logf("Lua 5.1 load check ran through %q", ran)
		} else {
			t.Log("no Lua 5.1 interpreter on PATH or in WSL: only the pure-Go ParseLua check ran")
		}
	})
}

// TestOutboxV1FixtureGolden pins contract 4's shape from the reading side. The companion has
// no outbox type yet (Lane L7 adds one and reads it from ForeverSixty.lua), so the structs
// here are local, strict, and the fields L7 must read.
func TestOutboxV1FixtureGolden(t *testing.T) {
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
	contractStrict(t, contractFixture(t, "outbox-v1.json"), &outbox)

	if outbox.Format != 1 {
		t.Errorf("format = %d, want 1", outbox.Format)
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
}
