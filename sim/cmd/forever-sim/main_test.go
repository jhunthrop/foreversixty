package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	googleproto "google.golang.org/protobuf/proto"
)

func smallRequest(t *testing.T) []byte {
	t.Helper()
	req := api.SimRequest{
		EngineVersion: enginever.Version,
		Spec:          "warrior-fury",
		Source:        api.CharacterSource{Kind: api.SourceManual},
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

// noSampleRequest is smallRequest with NoSample set, which is the one
// thing that sends a plain-shaped request down the concurrent entry
// point (see entryPointFor): a sample request stays single-threaded and
// so it emits no "N concurrent sims" log line for a test that wants one.
func noSampleRequest(t *testing.T) []byte {
	t.Helper()
	var req api.SimRequest
	if err := json.Unmarshal(smallRequest(t), &req); err != nil {
		t.Fatal(err)
	}
	req.NoSample = true
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

// fixtureRequest loads a checked-in SimRequest fixture from
// sim/adapter/testdata, decoded strictly so a field this build cannot
// honour is caught by the loader rather than silently ignored.
func fixtureRequest(t *testing.T, spec string) api.SimRequest {
	t.Helper()
	path := filepath.Join("..", "..", "adapter", "testdata", spec+".request.json")
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var req api.SimRequest
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		t.Fatalf("%s is not a SimRequest: %v", path, err)
	}
	return req
}

func TestRunProducesASimResult(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "req.json")
	if err := os.WriteFile(in, smallRequest(t), 0o644); err != nil {
		t.Fatal(err)
	}
	out := filepath.Join(dir, "res.json")

	var progress bytes.Buffer
	if err := run(in, out, overrides{}, false, &progress); err != nil {
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

// buildBinary builds forever-sim the way `make artifacts` does, so a
// test can look at what the real process writes to the real stderr.
// Nothing in-process can: the engine talks to the standard log package,
// which captured os.Stderr when it was initialised, so reassigning
// os.Stderr here would redirect nothing.
func buildBinary(t *testing.T) string {
	t.Helper()
	bin := filepath.Join(t.TempDir(), "forever-sim")
	cmd := exec.Command("go", "build", "-o", bin, ".")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}
	return bin
}

// -progress promises JSON lines on stderr, and the engine's own log
// output goes to stderr too. Every line has to parse, or the reader on
// the other end - sim/runner.Native - is guessing which lines are its
// own. This reads the real stderr of the real binary, because that is
// the only place the two streams actually meet.
//
// It runs a NoSample request rather than smallRequest's plain shape:
// a sample request now stays on the engine's single-threaded entry
// point (entryPointFor), which - for a run this small - logs nothing
// at all, and a test that wants an engine log line to wrap needs the
// concurrent path's "N concurrent sims" line to prove the wrapper
// against.
func TestProgressIsJSONLines(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "req.json")
	if err := os.WriteFile(in, noSampleRequest(t), 0o644); err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command(buildBinary(t), "-in", in, "-out", filepath.Join(dir, "res.json"), "-progress")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("forever-sim: %v\n%s", err, stderr.String())
	}

	lines := strings.Split(strings.TrimSpace(stderr.String()), "\n")
	if len(lines) == 0 || lines[0] == "" {
		t.Fatal("no progress was written")
	}
	var ticks, logs int
	for i, line := range lines {
		var got struct {
			Completed *int     `json:"completed"`
			Total     *int     `json:"total"`
			DPS       *float64 `json:"dps"`
			Log       *string  `json:"log"`
		}
		if err := json.Unmarshal([]byte(line), &got); err != nil {
			t.Fatalf("stderr line %d is not JSON: %q (%v); the engine's log is interleaving with the ticks", i, line, err)
		}
		switch {
		case got.Total != nil:
			ticks++
			if *got.Total != 500 {
				t.Errorf("progress line %d has total %d, want 500", i, *got.Total)
			}
			if got.Completed == nil || got.DPS == nil {
				t.Errorf("progress line %d is missing completed or dps: %q", i, line)
			}
		case got.Log != nil:
			logs++
		default:
			t.Errorf("stderr line %d is neither a tick nor a log line: %q", i, line)
		}
	}
	if ticks == 0 {
		t.Error("no progress ticks were written")
	}
	// The engine says at least "Running 500 iterations…"; if that ever
	// stops being wrapped, this catches it before the reader does.
	if logs == 0 {
		t.Error("the engine's log output did not come through the wrapper")
	}
}

// The inverse of TestProgressIsJSONLines, and the test that actually
// pins entryPointFor's call site (main.go's own comment: "DO NOT
// optimise this back"). TestEntryPointForChoosesTheSerialPathOnlyForASample
// proves the mapping function in isolation, but nothing stopped someone
// from reverting the ONE call to it - entryPointFor would just become
// an unused function, which Go does not flag, and every other test
// still passes because the concurrent path also returns a
// SampleIteration (pickSampleIteration's approximation). Running the
// real binary and checking stderr NEVER logs "concurrent sims" for a
// plain request is what would actually catch that revert.
func TestASampleRequestNeverLogsTheConcurrentPath(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "req.json")
	if err := os.WriteFile(in, smallRequest(t), 0o644); err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command(buildBinary(t), "-in", in, "-out", filepath.Join(dir, "res.json"), "-progress")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("forever-sim: %v\n%s", err, stderr.String())
	}

	if strings.Contains(stderr.String(), "concurrent sims") {
		t.Error("a plain (sample) request logged the concurrent path; entryPointFor's call site was reverted")
	}

	// And the positive half: it still produced a result with a sample,
	// so this is not passing by accident (e.g. the run failing before
	// it reaches either entry point).
	out, err := os.ReadFile(filepath.Join(dir, "res.json"))
	if err != nil {
		t.Fatal(err)
	}
	var res api.SimResult
	if err := json.Unmarshal(out, &res); err != nil {
		t.Fatal(err)
	}
	if res.Error != "" {
		t.Fatalf("the sim reported an error: %s", res.Error)
	}
	if len(res.Sample) == 0 {
		t.Error("a plain request carried no sample")
	}
}

// The same stream, in process, for the payload's own shape: run takes
// the sink as a parameter precisely so this needs no subprocess.
func TestProgressPayloadFields(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "req.json")
	if err := os.WriteFile(in, smallRequest(t), 0o644); err != nil {
		t.Fatal(err)
	}
	var progress bytes.Buffer
	if err := run(in, filepath.Join(dir, "res.json"), overrides{}, false, &progress); err != nil {
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
	// 100 is deliberately NOT one of api.ValidIterations. The flag
	// advertises "override the request's iteration count", and it used
	// to die with "iterations must be one of [500 3000 10000]" for any
	// number outside the settings bar's closed set - an operator
	// reproducing something quickly could not use the flag the binary
	// offered them. An override goes through ValidatePart, which is
	// the shape it is.
	if err := run(in, out, overrides{iterations: 100}, false, nil); err != nil {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(out)
	var res api.SimResult
	if err := json.Unmarshal(b, &res); err != nil {
		t.Fatal(err)
	}
	if res.IterationsRun != 100 {
		t.Errorf("IterationsRun = %d, want the overridden 100", res.IterationsRun)
	}
	// Bounded, not unbounded: a part is a share of a whole run and can
	// never legitimately exceed the largest one.
	if err := run(in, out, overrides{iterations: api.MaxIterations + 1}, false, nil); err == nil {
		t.Error("an override larger than the largest whole run was accepted")
	}
}

func TestBadInputIsRejected(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "junk.json")
	if err := os.WriteFile(in, []byte("{not json"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := run(in, filepath.Join(dir, "res.json"), overrides{}, false, nil); err == nil {
		t.Fatal("junk input was accepted")
	}
	if err := run(filepath.Join(dir, "missing.json"), filepath.Join(dir, "res.json"), overrides{}, false, nil); err == nil {
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
	if err := runProto(in, out, overrides{}, false, nil); err != nil {
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
	if err := runProto(filepath.Join(dir, "missing.json"), out, overrides{}, false, nil); err == nil {
		t.Error("runProto accepted a missing input file")
	}
}

// A request the builder refuses IS bad input: it never reaches the
// engine, and the exit code says so, because a caller that retried a
// malformed request forever is what the distinction prevents. (The
// test used to be named for the opposite of what it asserts.)
func TestARequestTheBuilderRefusesIsBadInput(t *testing.T) {
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
	err = run(in, filepath.Join(dir, "res.json"), overrides{}, false, nil)
	if err == nil {
		t.Fatal("an unknown race was accepted")
	}
	if !errors.Is(err, errBadInput) {
		t.Errorf("error %v is not errBadInput; a request the builder refuses is bad input", err)
	}
}

// The row's engine version is a fact about the binary that produced it,
// not the request's claim. It used to be copied straight off the
// request, so a cached request pinned to an old sha came back stamped
// with the old sha and api.SimResult.Stale could never fire.
func TestTheResultIsStampedWithTheBinarysOwnEngine(t *testing.T) {
	var req api.SimRequest
	if err := json.Unmarshal(smallRequest(t), &req); err != nil {
		t.Fatal(err)
	}
	res, err := Execute(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.EngineVersion != enginever.Version {
		t.Errorf("EngineVersion = %q, want %q", res.EngineVersion, enginever.Version)
	}
	if want := "sim:" + enginever.Version; res.Summary.EngineVersion != want {
		t.Errorf("Summary.EngineVersion = %q, want %q", res.Summary.EngineVersion, want)
	}
	if res.Stale(enginever.Version) {
		t.Error("a result this binary just produced reads as stale")
	}
	if !res.Stale("deadbee") {
		t.Error("Stale never fires; the stamp is not a fact")
	}
}

// A sample iteration is exact only on the engine's single-threaded
// entry point (see entryPointFor's own comment): the concurrent split
// keeps whichever shard's local median lands closest to the combined
// mean, which the engine fork's own pickSampleIteration documents as
// "a KNOWN APPROXIMATION, not the genuine global median". A request
// that asked for a sample must not silently get that approximation,
// and a request that did not ask for one must keep the fast concurrent
// path - so this pins the mapping by function identity rather than by
// running two full sims and hoping a stray difference shows through.
func TestEntryPointForChoosesTheSerialPathOnlyForASample(t *testing.T) {
	if got, want := reflect.ValueOf(entryPointFor(true)).Pointer(), reflect.ValueOf(core.RunRaidSimAsync).Pointer(); got != want {
		t.Error("a sample request did not choose the engine's single-threaded entry point")
	}
	if got, want := reflect.ValueOf(entryPointFor(false)).Pointer(), reflect.ValueOf(core.RunRaidSimConcurrentAsync).Pointer(); got != want {
		t.Error("a plain request did not choose the engine's concurrent entry point")
	}
}

// A plain run - the shape every request has unless something says
// otherwise - always asks for the sample and gets one back: the
// report's sample card is not optional for a run nobody flagged as
// one of many.
func TestExecuteFillsTheSampleForAPlainRun(t *testing.T) {
	var req api.SimRequest
	if err := json.Unmarshal(smallRequest(t), &req); err != nil {
		t.Fatal(err)
	}
	res, err := Execute(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Sample) == 0 {
		t.Error("a plain run carried no sample; the report's sample card would render nothing")
	}
}

// api.SimRequest.NoSample is how a caller building dozens of stage
// requests - one per bulk combination - says its cast log would never
// be read. Setting it must turn the sample off end to end, not just at
// the protobuf boundary.
func TestExecuteOmitsTheSampleWhenNoSampleIsSet(t *testing.T) {
	var req api.SimRequest
	if err := json.Unmarshal(smallRequest(t), &req); err != nil {
		t.Fatal(err)
	}
	req.NoSample = true
	res, err := Execute(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Sample) != 0 {
		t.Errorf("a NoSample run carried %d sample rows, want 0", len(res.Sample))
	}
}

// -no-sample is the batch callers' opt-out. The binary's own header
// names the nightly validation job and the execution scorer, and each
// runs thousands of sims whose cast log nobody reads; without a flag
// their only way to decline was to write api.SimRequest.NoSample into
// the JSON, which that field's own doc used to reserve for sim/bulk.
// The override is one-way: it can take the sample off, never put one
// on.
func TestNoSampleOverrideOnlyEverTurnsTheSampleOff(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "req.json")
	if err := os.WriteFile(in, smallRequest(t), 0o644); err != nil {
		t.Fatal(err)
	}

	plain, err := load(in, overrides{})
	if err != nil {
		t.Fatal(err)
	}
	if plain.NoSample {
		t.Error("a request loaded without -no-sample opted out of the sample")
	}
	off, err := load(in, overrides{noSample: true})
	if err != nil {
		t.Fatal(err)
	}
	if !off.NoSample {
		t.Error("-no-sample did not reach api.SimRequest.NoSample")
	}

	// A request that already opted out keeps its own answer whether the
	// flag is given or not.
	var req api.SimRequest
	if err := json.Unmarshal(smallRequest(t), &req); err != nil {
		t.Fatal(err)
	}
	req.NoSample = true
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	optedOut := filepath.Join(dir, "no-sample.json")
	if err := os.WriteFile(optedOut, b, 0o644); err != nil {
		t.Fatal(err)
	}
	stillOff, err := load(optedOut, overrides{})
	if err != nil {
		t.Fatal(err)
	}
	if !stillOff.NoSample {
		t.Error("a request that opted out had the sample put back")
	}
}

// End to end through the real flag set: the binary must accept
// -no-sample and the SimResult it writes must carry no sample rows.
// This is what pins the flag to the override - the unit test above
// covers the override alone, and a flag nobody wired to it would still
// pass that.
func TestTheBinaryAcceptsNoSample(t *testing.T) {
	bin := filepath.Join(t.TempDir(), "forever-sim")
	build := exec.Command("go", "build", "-o", bin, ".")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}
	dir := t.TempDir()
	in := filepath.Join(dir, "req.json")
	if err := os.WriteFile(in, smallRequest(t), 0o644); err != nil {
		t.Fatal(err)
	}
	out := filepath.Join(dir, "res.json")
	cmd := exec.Command(bin, "-in", in, "-out", out, "-iterations", "50", "-no-sample")
	if b, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("forever-sim -no-sample: %v\n%s", err, b)
	}
	b, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	var res api.SimResult
	if err := json.Unmarshal(b, &res); err != nil {
		t.Fatal(err)
	}
	if res.Error != "" {
		t.Fatalf("the run failed: %s", res.Error)
	}
	if len(res.Sample) != 0 {
		t.Errorf("-no-sample still produced %d sample rows", len(res.Sample))
	}
	if !res.Request.NoSample {
		t.Error("the echoed request does not record the opt-out")
	}
}

// -version prints the pin the binary was compiled with. sim/enginever
// used to be imported by no Go file at all: both mains declared
// `var Version = "dev"` and relied on -ldflags, so a plain `go build`
// produced a binary that stamped "dev" on real rows.
func TestTheBinaryKnowsItsOwnEngineWithoutLdflags(t *testing.T) {
	bin := filepath.Join(t.TempDir(), "forever-sim")
	build := exec.Command("go", "build", "-o", bin, ".")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}
	out, err := exec.Command(bin, "-version").Output()
	if err != nil {
		t.Fatalf("-version: %v", err)
	}
	if got := strings.TrimSpace(string(out)); got != enginever.Version {
		t.Errorf("-version printed %q, want %q", got, enginever.Version)
	}
}

// A field the envelope does not carry is a caller sending something
// this build cannot honour, and running anyway would drop it silently.
func TestUnknownFieldsAreRefused(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "req.json")
	var raw map[string]any
	if err := json.Unmarshal(smallRequest(t), &raw); err != nil {
		t.Fatal(err)
	}
	raw["covenant"] = "kyrian"
	b, err := json.Marshal(raw)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(in, b, 0o644); err != nil {
		t.Fatal(err)
	}
	err = run(in, filepath.Join(dir, "res.json"), overrides{}, false, nil)
	if err == nil {
		t.Fatal("a request carrying an unknown field was accepted")
	}
	if !errors.Is(err, errBadInput) {
		t.Errorf("error %v is not errBadInput", err)
	}
}

// Stopping a run is something the caller asked for, not a corruption.
// The engine reports an abort as an ErrorOutcome with an EMPTY message,
// so the old guard - Error != nil && Message != "" - let it through and
// the adapter then said "iterations_done is 0".
func TestAnAbortedRunIsWrittenAsAnAbort(t *testing.T) {
	var req api.SimRequest
	if err := json.Unmarshal(smallRequest(t), &req); err != nil {
		t.Fatal(err)
	}
	req.Iterations = api.MaxIterations

	// Abort as soon as the run registers its id. runID is this
	// process's counter, so the next id is the one Execute is about to
	// take; the loop covers the race between registering and aborting.
	next := fmt.Sprintf("forever-sim-%d-%d", os.Getpid(), runs.Load()+1)
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := 0; i < 2000; i++ {
			if simsignals.AbortById(next) {
				return
			}
			time.Sleep(time.Millisecond)
		}
	}()

	dir := t.TempDir()
	out := filepath.Join(dir, "res.json")
	in := filepath.Join(dir, "req.json")
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(in, b, 0o644); err != nil {
		t.Fatal(err)
	}
	err = run(in, out, overrides{}, false, nil)
	<-done
	if !errors.Is(err, adapter.ErrAborted) {
		t.Fatalf("run returned %v, want adapter.ErrAborted", err)
	}
	if errors.Is(err, errBadInput) {
		t.Error("an abort reads as bad input")
	}

	// The partial result is still written: the caller asked for the
	// run to stop, and what it got is what it got.
	raw, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("an aborted run wrote no result: %v", err)
	}
	var res api.SimResult
	if err := json.Unmarshal(raw, &res); err != nil {
		t.Fatal(err)
	}
	if !res.Aborted {
		t.Error("the result does not say it was stopped")
	}
	if res.Error != "" {
		t.Errorf("an abort carries an error message: %q", res.Error)
	}
	if res.EngineVersion != enginever.Version {
		t.Errorf("EngineVersion = %q, want %q", res.EngineVersion, enginever.Version)
	}
	// And it is written in the SHAPE a finished result has. A zero
	// summary.Summary marshals its lists as null, so a page rendering a
	// stopped run would need a null check per key; the abort path
	// carries adapter.EmptySummary() instead. Asserted on the bytes
	// that were actually written, not on the struct, because it is the
	// marshalling that differs.
	var onDisk struct {
		Summary map[string]json.RawMessage `json:"summary"`
	}
	if err := json.Unmarshal(raw, &onDisk); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"damage_done", "auras", "casts", "resources", "roster", "phases"} {
		if got := string(onDisk.Summary[key]); got != "[]" {
			t.Errorf("an aborted run wrote %q as %s, want []", key, got)
		}
	}
}

// The signal path: SIGINT and SIGTERM stop the run the way the
// browser's Stop button does, through the engine's own abort signal,
// rather than killing the process mid-write. The handler is installed
// for the duration of a run, so the default "terminate" action is not
// in force while this test raises the signal at itself.
func TestASignalAbortsTheRun(t *testing.T) {
	for _, sig := range []syscall.Signal{syscall.SIGINT, syscall.SIGTERM} {
		t.Run(sig.String(), func(t *testing.T) {
			id := runID()
			signals, err := simsignals.RegisterWithId(id)
			if err != nil {
				t.Fatal(err)
			}
			defer simsignals.UnregisterId(id)

			stop := onInterrupt(id)
			defer stop()
			if err := syscall.Kill(os.Getpid(), sig); err != nil {
				t.Fatal(err)
			}
			deadline := time.Now().Add(5 * time.Second)
			for !signals.Abort.IsTriggered() {
				if time.Now().After(deadline) {
					t.Fatalf("%v did not abort the run", sig)
				}
				time.Sleep(time.Millisecond)
			}
		})
	}

	// Once the run is over the handler comes back down, so a later
	// signal is the shell's business again and not a stale abort.
	id := runID()
	if _, err := simsignals.RegisterWithId(id); err != nil {
		t.Fatal(err)
	}
	defer simsignals.UnregisterId(id)
	onInterrupt(id)()
}
