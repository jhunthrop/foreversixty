package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
	googleproto "google.golang.org/protobuf/proto"
)

func smallRequest(t *testing.T) []byte {
	t.Helper()
	req := api.SimRequest{
		EngineVersion: "test",
		Spec:          "warrior-fury",
		Character: api.CharacterSpec{
			Name: "CLI Test", Race: "orc", Class: "warrior", Level: 60,
			Talents: "30305001302-05050005525010051",
		},
		Encounter:  api.EncounterSpec{DurationSec: 60, Variation: 0, Targets: 1, ExecuteRatio: 0.25},
		Iterations: 500,
		RandomSeed: 1,
	}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func TestRunProducesASimResult(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "req.json")
	if err := os.WriteFile(in, smallRequest(t), 0o644); err != nil {
		t.Fatal(err)
	}
	out := filepath.Join(dir, "res.json")

	var progress bytes.Buffer
	if err := run(in, out, 0, &progress); err != nil {
		t.Fatalf("run: %v", err)
	}

	b, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	var res api.SimResult
	if err := json.Unmarshal(b, &res); err != nil {
		t.Fatalf("the output is not a SimResult: %v", err)
	}
	if res.Error != "" {
		t.Fatalf("the sim reported an error: %s", res.Error)
	}
	if res.Lane != api.LaneServer {
		t.Errorf("Lane = %q, want %q", res.Lane, api.LaneServer)
	}
	if res.IterationsRun != 500 {
		t.Errorf("IterationsRun = %d, want 500", res.IterationsRun)
	}
	if res.DPS.Mean <= 0 {
		t.Error("no DPS")
	}
	// The whole point of building the binary here rather than in the
	// engine: the result arrives with its summary already built.
	if len(res.Summary.DamageDone) == 0 {
		t.Error("the result carries no summary; sim/adapter did not run")
	}
	if res.Summary.EngineVersion == "" {
		t.Error("the summary carries no engine version")
	}
}

func TestProgressIsJSONLines(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "req.json")
	if err := os.WriteFile(in, smallRequest(t), 0o644); err != nil {
		t.Fatal(err)
	}
	var progress bytes.Buffer
	if err := run(in, filepath.Join(dir, "res.json"), 0, &progress); err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(progress.String()), "\n")
	if len(lines) == 0 || lines[0] == "" {
		t.Fatal("no progress was written")
	}
	for i, line := range lines {
		var got struct {
			Completed int     `json:"completed"`
			Total     int     `json:"total"`
			DPS       float64 `json:"dps"`
		}
		if err := json.Unmarshal([]byte(line), &got); err != nil {
			t.Fatalf("progress line %d is not JSON: %q (%v)", i, line, err)
		}
		if got.Total != 500 {
			t.Errorf("progress line %d has total %d, want 500", i, got.Total)
		}
	}
}

func TestIterationsOverride(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "req.json")
	if err := os.WriteFile(in, smallRequest(t), 0o644); err != nil {
		t.Fatal(err)
	}
	out := filepath.Join(dir, "res.json")
	if err := run(in, out, 3000, nil); err != nil {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(out)
	var res api.SimResult
	if err := json.Unmarshal(b, &res); err != nil {
		t.Fatal(err)
	}
	if res.IterationsRun != 3000 {
		t.Errorf("IterationsRun = %d, want the overridden 3000", res.IterationsRun)
	}
}

func TestBadInputIsRejected(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "junk.json")
	if err := os.WriteFile(in, []byte("{not json"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := run(in, filepath.Join(dir, "res.json"), 0, nil); err == nil {
		t.Fatal("junk input was accepted")
	}
	if err := run(filepath.Join(dir, "missing.json"), filepath.Join(dir, "res.json"), 0, nil); err == nil {
		t.Fatal("a missing input file was accepted")
	}
}

// -out-proto writes the engine's own RaidSimResult instead of ours. It
// is how sim/adapter's fixtures are regenerated, and a fixture that
// carried the cast log would be one iteration of text larger for nothing.
func TestOutProtoWritesAnEngineResult(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "req.json")
	if err := os.WriteFile(in, smallRequest(t), 0o644); err != nil {
		t.Fatal(err)
	}
	out := filepath.Join(dir, "res.pb")
	if err := runProto(in, out, 0, nil); err != nil {
		t.Fatalf("runProto: %v", err)
	}
	b, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	res := &proto.RaidSimResult{}
	if err := googleproto.Unmarshal(b, res); err != nil {
		t.Fatalf("the output is not a RaidSimResult: %v", err)
	}
	if res.IterationsDone != 500 {
		t.Errorf("IterationsDone = %d, want 500", res.IterationsDone)
	}
	if res.Logs != "" {
		t.Error("the fixture carries the cast log; it makes the file large and the adapter never reads it")
	}
	if _, err := adapter.Summarize(res, api.SimRequest{EngineVersion: "t", Spec: "warrior-fury"}); err != nil {
		t.Errorf("the written result does not summarize: %v", err)
	}
	if err := runProto(filepath.Join(dir, "missing.json"), out, 0, nil); err == nil {
		t.Error("runProto accepted a missing input file")
	}
}

// A request the engine cannot run is an engine failure, not bad input:
// the two exit differently, and a caller that retries a bad request
// forever is the thing the distinction prevents.
func TestAnUnrunnableRequestIsNotBadInput(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "req.json")
	var req api.SimRequest
	if err := json.Unmarshal(smallRequest(t), &req); err != nil {
		t.Fatal(err)
	}
	req.Character.Race = "murloc"
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(in, b, 0o644); err != nil {
		t.Fatal(err)
	}
	err = run(in, filepath.Join(dir, "res.json"), 0, nil)
	if err == nil {
		t.Fatal("an unknown race was accepted")
	}
	if !errors.Is(err, errBadInput) {
		t.Errorf("error %v is not errBadInput; a request the builder refuses is bad input", err)
	}
}
