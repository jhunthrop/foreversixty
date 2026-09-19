package api

import (
	"encoding/json"
	"strings"
	"testing"
)

// The JSON field names are the contract, shared verbatim with
// web/src/lib/sim/types.ts. A rename here is a break there, so the test
// pins the wire form rather than the Go field names.
func TestSimRequestJSONFieldNames(t *testing.T) {
	req := SimRequest{
		EngineVersion: "7779ebb",
		Spec:          "warrior-fury",
		Source:        CharacterSource{Kind: SourceArmory, Ref: "us/normal/thrall", CapturedAt: "2026-09-14T00:00:00Z"},
		Encounter:     DefaultEncounter(),
		Character: CharacterSpec{
			Name: "Thrall", Race: "orc", Class: "warrior", Level: 60,
			Talents: "30305001302-05050005525010051",
			Gear:    []GearSlot{{Slot: "main_hand", ItemID: 19352, Enchant: 2568}},
			Buffs:   []string{"battle_shout"},
		},
		Iterations: 3000,
		RandomSeed: 0,
	}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatal(err)
	}
	for _, k := range []string{"engine_version", "spec", "source", "character", "encounter", "iterations", "random_seed"} {
		if _, ok := m[k]; !ok {
			t.Errorf("SimRequest is missing JSON key %q", k)
		}
	}
	src, ok := m["source"].(map[string]any)
	if !ok {
		t.Fatalf("source is not an object: %T", m["source"])
	}
	for _, k := range []string{"kind", "ref", "captured_at"} {
		if _, ok := src[k]; !ok {
			t.Errorf("CharacterSource is missing JSON key %q", k)
		}
	}
	ch, ok := m["character"].(map[string]any)
	if !ok {
		t.Fatalf("character is not an object: %T", m["character"])
	}
	for _, k := range []string{"name", "race", "class", "level", "talents", "gear", "buffs", "consumes"} {
		if _, ok := ch[k]; !ok {
			t.Errorf("CharacterSpec is missing JSON key %q", k)
		}
	}
	gear, ok := ch["gear"].([]any)
	if !ok || len(gear) == 0 {
		t.Fatalf("gear is not a non-empty array: %v", ch["gear"])
	}
	for _, k := range []string{"slot", "item_id"} {
		if _, ok := gear[0].(map[string]any)[k]; !ok {
			t.Errorf("GearSlot is missing JSON key %q", k)
		}
	}

	enc, ok := m["encounter"].(map[string]any)
	if !ok {
		t.Fatalf("encounter is not an object: %T", m["encounter"])
	}
	for _, k := range []string{"duration_sec", "variation", "targets", "execute_ratio", "profile"} {
		if _, ok := enc[k]; !ok {
			t.Errorf("EncounterSpec is missing JSON key %q", k)
		}
	}
}

func TestDefaultEncounterMatchesTheContract(t *testing.T) {
	got := DefaultEncounter()
	want := EncounterSpec{DurationSec: 180, Variation: 0.2, Targets: 1, ExecuteRatio: 0.25, Profile: ""}
	if got != want {
		t.Errorf("DefaultEncounter() = %+v, want %+v", got, want)
	}
}

func TestValidateRejectsBadRequests(t *testing.T) {
	good := SimRequest{
		EngineVersion: "7779ebb", Spec: "mage-frost", Iterations: 3000,
		Encounter: DefaultEncounter(),
		Character: CharacterSpec{Name: "Jaina", Race: "gnome", Class: "mage", Level: 60},
	}
	if err := good.Validate(); err != nil {
		t.Fatalf("a good request was rejected: %v", err)
	}
	cases := []struct {
		name string
		mut  func(*SimRequest)
		want string
	}{
		{"no engine version", func(r *SimRequest) { r.EngineVersion = "" }, "engine_version"},
		{"no spec", func(r *SimRequest) { r.Spec = "" }, "spec"},
		{"odd iteration count", func(r *SimRequest) { r.Iterations = 1234 }, "iterations"},
		{"no duration", func(r *SimRequest) { r.Encounter.DurationSec = 0 }, "duration_sec"},
		{"too many targets", func(r *SimRequest) { r.Encounter.Targets = 11 }, "targets"},
		{"no class", func(r *SimRequest) { r.Character.Class = "" }, "character.class"},
		{"no race", func(r *SimRequest) { r.Character.Race = "" }, "character.race"},
		{"no level", func(r *SimRequest) { r.Character.Level = 0 }, "character.level"},
		// The engine has one level. A request for any other cannot be
		// run, so it is refused here rather than queued and failed at
		// the worker.
		{"a level the engine cannot sim", func(r *SimRequest) { r.Character.Level = 40 }, "character.level"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := good
			tc.mut(&req)
			err := req.Validate()
			if err == nil {
				t.Fatalf("expected an error mentioning %q, got nil", tc.want)
			}
			if !strings.Contains(err.Error(), tc.want) {
				t.Errorf("error %q does not mention %q", err, tc.want)
			}
		})
	}
}

// A stored sim row carries its whole request, and the request is JSON all
// the way down: no protobuf crosses a lane boundary, so a stored row can
// be re-run by handing it straight back to sim/request.
func TestSimResultRoundTripsItsRequest(t *testing.T) {
	res := SimResult{
		EngineVersion: "7779ebb",
		Request: SimRequest{
			Spec:      "warrior-fury",
			Character: CharacterSpec{Name: "Thrall", Race: "orc", Class: "warrior", Level: 60},
		},
		Lane: LaneBrowser,
		DPS:  Estimate{Mean: 1791.1, StdDev: 120, Error: 2.2, Min: 1400, Max: 2100},
	}
	b, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(b), `"raw"`) {
		t.Errorf("the envelope still carries a raw protobuf field: %s", b)
	}
	var back SimResult
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatal(err)
	}
	again, err := json.Marshal(back)
	if err != nil {
		t.Fatal(err)
	}
	if string(again) != string(b) {
		t.Errorf("the request did not survive a round trip:\n got %s\nwant %s", again, b)
	}
}

func TestStaleComparesEngineVersions(t *testing.T) {
	res := SimResult{EngineVersion: "aaaaaaa"}
	if res.Stale("aaaaaaa") {
		t.Error("a result from the current engine reported stale")
	}
	if !res.Stale("bbbbbbb") {
		t.Error("a result from another engine build did not report stale")
	}
}
