package api

import (
	"encoding/json"
	"slices"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// The JSON field names are the contract, shared verbatim with
// web/src/lib/sim/types.ts. A rename here is a break there, so the test
// pins the wire form rather than the Go field names.
func TestSimRequestJSONFieldNames(t *testing.T) {
	req := SimRequest{
		EngineVersion: enginever.Version,
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
		EngineVersion: enginever.Version, Spec: "mage-frost", Iterations: 3000,
		Source:    CharacterSource{Kind: SourceManual},
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
		EngineVersion: enginever.Version,
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

// A worker's share of a split run is 750 iterations, or 250, or whatever
// the division produced. Validate must refuse that and ValidatePart must
// accept it, or the browser's worker pool cannot build the parts
// combine.Split just handed it.
func TestValidatePartAcceptsASplitShareAndNothingElse(t *testing.T) {
	part := SimRequest{
		EngineVersion: enginever.Version, Spec: "mage-frost", Iterations: 750,
		Source:    CharacterSource{Kind: SourceManual},
		Encounter: DefaultEncounter(),
		Character: CharacterSpec{Name: "Jaina", Race: "gnome", Class: "mage", Level: 60},
	}
	if err := part.Validate(); err == nil {
		t.Error("Validate accepted 750 iterations; the settings bar offers no such run")
	}
	if err := part.ValidatePart(); err != nil {
		t.Errorf("ValidatePart rejected a split share: %v", err)
	}

	// Everything else Validate checks still applies to a part.
	noRace := part
	noRace.Character.Race = ""
	if err := noRace.ValidatePart(); err == nil {
		t.Error("ValidatePart accepted a part with no race")
	}

	// A part is a share of a whole request, so it is bounded by the
	// largest whole request. Nothing may smuggle a million-iteration run
	// past the closed set by calling itself a part.
	for _, n := range []int{0, -1, MaxIterations + 1} {
		bad := part
		bad.Iterations = n
		if err := bad.ValidatePart(); err == nil {
			t.Errorf("ValidatePart accepted %d iterations", n)
		}
	}
	if err := func() error { p := part; p.Iterations = MaxIterations; return p.ValidatePart() }(); err != nil {
		t.Errorf("ValidatePart rejected the largest whole run: %v", err)
	}
}

// MaxIterations is the top of ValidIterations rather than a second copy
// of the number; adding a larger run to the closed set must move it.
func TestMaxIterationsTracksTheClosedSet(t *testing.T) {
	for _, n := range ValidIterations {
		if n > MaxIterations {
			t.Errorf("ValidIterations has %d, above MaxIterations %d", n, MaxIterations)
		}
	}
	if !slices.Contains(ValidIterations, MaxIterations) {
		t.Errorf("MaxIterations %d is not in ValidIterations %v", MaxIterations, ValidIterations)
	}
}

// The progress payload is a contract with the web: the sim island reads
// a partial result with the same accessors it reads a finished one with,
// so the two keys must be SimResult's own.
func TestProgressIsAPickOfSimResult(t *testing.T) {
	b, err := json.Marshal(Progress{IterationsRun: 250, DPS: Estimate{Mean: 1427.4}})
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatal(err)
	}
	if len(m) != 2 {
		t.Errorf("Progress has %d keys, want exactly iterations_run and dps: %v", len(m), m)
	}
	// The same two keys a SimResult carries, spelled the same way.
	var res map[string]any
	rb, err := json.Marshal(SimResult{})
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(rb, &res); err != nil {
		t.Fatal(err)
	}
	for _, k := range []string{"iterations_run", "dps"} {
		if _, ok := m[k]; !ok {
			t.Errorf("Progress is missing JSON key %q", k)
		}
		if _, ok := res[k]; !ok {
			t.Errorf("SimResult has no key %q, so Progress is not a Pick of it", k)
		}
	}
	dps, ok := m["dps"].(map[string]any)
	if !ok {
		t.Fatalf("dps is not an Estimate object: %T", m["dps"])
	}
	if dps["mean"] != 1427.4 {
		t.Errorf("dps.mean = %v, want the running mean 1427.4", dps["mean"])
	}
	for _, k := range []string{"stddev", "error", "min", "max"} {
		if dps[k] != 0.0 {
			t.Errorf("dps.%s = %v mid-run, want 0 until the run completes", k, dps[k])
		}
	}
}

// One string identifies the engine build everywhere. A request naming
// another cannot be answered by this binary: its talents, its item rows
// and its spell constants all belong to the pinned sha, so running it
// would produce a row stamped with a version that did not produce it -
// and SimResult.Stale, which exists to catch exactly that, could never
// fire because the stamp came from the claim.
func TestValidateRefusesARequestForAnotherEngine(t *testing.T) {
	req := SimRequest{
		EngineVersion: "deadbee", Spec: "mage-frost", Iterations: 3000,
		Source:    CharacterSource{Kind: SourceManual},
		Encounter: DefaultEncounter(),
		Character: CharacterSpec{Name: "Jaina", Race: "gnome", Class: "mage", Level: 60},
	}
	err := req.Validate()
	if err == nil {
		t.Fatal("a request naming another engine was accepted")
	}
	for _, want := range []string{"deadbee", enginever.Version} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("the error does not name %q: %v", want, err)
		}
	}
	// A part of a split run is held to it too: a worker cannot smuggle
	// a foreign engine version past the whole-request check.
	if err := req.ValidatePart(); err == nil {
		t.Error("ValidatePart accepted a request for another engine")
	}
	req.EngineVersion = enginever.Version
	if err := req.Validate(); err != nil {
		t.Errorf("the pinned engine's own version was rejected: %v", err)
	}
}

// Source.Kind had five constants declared beside it and nothing ever
// compared to them, so a typo rode through as a stored row nothing
// could join on.
func TestValidateChecksTheSourceKind(t *testing.T) {
	req := SimRequest{
		EngineVersion: enginever.Version, Spec: "mage-frost", Iterations: 3000,
		Encounter: DefaultEncounter(),
		Character: CharacterSpec{Name: "Jaina", Race: "gnome", Class: "mage", Level: 60},
	}
	for _, kind := range Sources {
		r := req
		r.Source.Kind = kind
		if err := r.Validate(); err != nil {
			t.Errorf("source kind %q was rejected: %v", kind, err)
		}
	}
	for _, kind := range []string{"", "Armory", "wowhead", "addon "} {
		r := req
		r.Source.Kind = kind
		if err := r.Validate(); err == nil {
			t.Errorf("source kind %q was accepted", kind)
		}
	}
}

// The profile was accepted and then thrown away: biomeFor ignored its
// argument, so "encounter:onyxia" produced a sim byte-identical to a
// blank one and said nothing. The vocabulary is refused at the boundary
// instead, and the one profile that IS what the sim builds is accepted.
func TestValidateChecksTheEncounterProfile(t *testing.T) {
	req := SimRequest{
		EngineVersion: enginever.Version, Spec: "mage-frost", Iterations: 3000,
		Source:    CharacterSource{Kind: SourceManual},
		Encounter: DefaultEncounter(),
		Character: CharacterSpec{Name: "Jaina", Race: "gnome", Class: "mage", Level: 60},
	}
	for _, profile := range []string{"", ProfilePatchwerk} {
		r := req
		r.Encounter.Profile = profile
		if err := r.Validate(); err != nil {
			t.Errorf("profile %q was rejected, and it is exactly the fight the sim builds: %v", profile, err)
		}
	}
	for _, tc := range []struct{ profile, want string }{
		{"encounter:onyxia", "not simulated yet"},
		{"encounter:", "names no encounter"},
		{"patchwork", "encounter.profile must be"},
		{"Patchwerk", "encounter.profile must be"},
	} {
		r := req
		r.Encounter.Profile = tc.profile
		err := r.Validate()
		if err == nil {
			t.Errorf("profile %q was accepted and would have been ignored", tc.profile)
			continue
		}
		if !strings.Contains(err.Error(), tc.want) {
			t.Errorf("profile %q: error %v does not say %q", tc.profile, err, tc.want)
		}
	}
}

// An abort is a state of its own on the result, not an error string: a
// run the user stopped is not a run that went wrong, and the two are
// rendered differently.
func TestAbortedIsOmittedWhenFalse(t *testing.T) {
	b, err := json.Marshal(SimResult{EngineVersion: enginever.Version})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(b), "aborted") {
		t.Errorf("a finished result carries an aborted key: %s", b)
	}
	b, err = json.Marshal(SimResult{EngineVersion: enginever.Version, Aborted: true})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(b), `"aborted":true`) {
		t.Errorf("a stopped result does not say so: %s", b)
	}
}

// BossLevel is one number two lanes read: sim/request builds the
// encounter at it and sim/measure refuses a log that never saw it. If
// they drifted the validation job would compare a sim against a log of
// a different target tier and call the gap a modelling error.
func TestBossLevelIsThreeAboveThePlayer(t *testing.T) {
	if BossLevel != SimLevel+3 {
		t.Errorf("BossLevel = %d, want %d", BossLevel, SimLevel+3)
	}
}
