// Command forever-sim runs one SimRequest natively and writes the
// SimResult, both as JSON. It is the server lane's binary: the premium
// Cloud Run job, the nightly validation job and the execution scorer all
// invoke it.
//
// It reads and writes our envelope, not the engine's protobuf, because
// sim/request and sim/adapter are linked in here and the boundary is
// this binary's own. The active build's item database is linked in too,
// through sim/internal/simdb; see that package for why not
// --tags=with_db. That is the same arrangement the browser gets from
// sim/cmd/wasm, which is the point: one mapping, one language, two
// lanes.
//
//	forever-sim -in request.json -out result.json -progress
//	forever-sim -in - -out - < request.json > result.json
//	forever-sim -plan -in request.json -out -
//
// The binary detects the request's kind from the request itself
// (api.SimRequest.Kind) rather than from a flag: a gear, talents or
// drops request runs the plan-rank loop over sim/bulk's stages
// instead of one plain sim, and -plan answers "how many combinations,
// and what would the first stage be" without running anything.
//
// SIGINT and SIGTERM stop the run the way the browser's Stop button
// does, through the engine's own abort signal, and the partial run is
// written out as a SimResult with aborted set rather than as a failure
// or a half-written file.
//
// With -progress every line of stderr is JSON: our own
// {"completed","total","dps"} ticks, and the engine's log output wrapped
// as {"log": "..."} rather than interleaved raw.
//
// Concurrency is automatic: core.RunRaidSimConcurrentAsync splits across
// runtime.NumCPU() and recombines the distribution metrics, offsetting
// each split's seed so the stream matches a serial run. The one
// exception is a request that asks for the sample iteration: see
// entryPointFor, which keeps that run on the engine's single-threaded
// entry point because the concurrent path's sample is only an
// approximation.
//
// A plain run asks for the sample, because that is the product's
// behaviour and the page renders it. -no-sample is the opt-out, and it
// is for the batch callers above: the nightly validation job and the
// execution scorer run thousands of sims whose cast log nobody reads,
// and each of those would otherwise give up the concurrent entry point
// and replay one whole extra iteration to record it.
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/bulk"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/internal/simdrain"
	"github.com/jhunthrop/foreversixty/sim/request"
	engine "github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	googleproto "google.golang.org/protobuf/proto"
)

const (
	exitOK      = 0
	exitSimFail = 1
	exitBadArgs = 2
	exitAborted = 130 // the shell's convention for "killed by SIGINT"
)

var errBadInput = errors.New("bad input")

func main() {
	in := flag.String("in", "-", "SimRequest JSON; - for stdin")
	out := flag.String("out", "-", "SimResult JSON; - for stdout")
	outProto := flag.String("out-proto", "", "write the engine's raw RaidSimResult protobuf here instead; only sim/adapter's fixture refresh wants this")
	iterations := flag.Int("iterations", 0, "override the request's iteration count")
	noSample := flag.Bool("no-sample", false, "skip the sample iteration: the run keeps the engine's concurrent entry point instead of replaying one fight single-threaded. For a batch caller - the nightly validation job, the execution scorer - whose cast log is never read")
	plan := flag.Bool("plan", false, "print the combination count and the first stage as JSON, and run nothing; the API counts a bulk request this way")
	progress := flag.Bool("progress", false, "write JSON-lines progress to stderr; the engine's own log output is wrapped as {\"log\":...} so every line of that stream parses")
	version := flag.Bool("version", false, "print the engine version and exit")
	flag.Parse()

	if *version {
		fmt.Println(enginever.Version)
		os.Exit(exitOK)
	}
	// -out-proto ignores planOnly by design (see runProto's own doc): it
	// exists for one caller regenerating a plain-run fixture. Accepting
	// -plan alongside it silently, and then running a full sim anyway,
	// is a trap; refuse the combination instead.
	if *plan && *outProto != "" {
		fmt.Fprintln(os.Stderr, "forever-sim:", fmt.Errorf("%w: -plan and -out-proto cannot be combined", errBadInput))
		os.Exit(exitBadArgs)
	}

	// Once, here, rather than once per run inside execute: the engine
	// guards RegisterAll with an unsynchronised package bool, so two
	// concurrent runs in one process are a data race and a possible
	// double registration. The Cloud Run job is that concurrent caller.
	engine.RegisterAll()

	var sink io.Writer
	if *progress {
		// One writer for the whole stream, and the engine's standard
		// logger routed through it: see progress.go. Without this the
		// engine's "Running N iterations" lands between two ticks and
		// stderr is only mostly JSON.
		stream := newProgressStderr(os.Stderr)
		log.SetOutput(stream.LogOutput())
		sink = stream
	}
	// -out-proto replaces -out rather than joining it: the one caller
	// that wants the engine's own result wants nothing else.
	run := run
	outPath := *out
	if *outProto != "" {
		run, outPath = runProto, *outProto
	}
	if err := run(*in, outPath, overrides{iterations: *iterations, noSample: *noSample}, *plan, sink); err != nil {
		fmt.Fprintln(os.Stderr, "forever-sim:", err)
		switch {
		case errors.Is(err, errBadInput):
			os.Exit(exitBadArgs)
		case errors.Is(err, adapter.ErrAborted):
			// Stopping a run is something the operator asked for, and
			// run has already written what completed.
			os.Exit(exitAborted)
		}
		os.Exit(exitSimFail)
	}
}

// overrides are the flags that change a request after it is read. They
// travel as one value so adding the next one is not another positional
// bool at every call site.
type overrides struct {
	// iterations replaces the request's count when above zero.
	iterations int
	// noSample sets api.SimRequest.NoSample, whatever the file said.
	// It only ever turns the sample OFF: a request that wants one says
	// so by being a plain run, which is the default.
	noSample bool
}

// run is main's body, with its files and its progress sink as parameters
// so it is testable.
func run(inPath, outPath string, over overrides, planOnly bool, progress io.Writer) error {
	req, err := load(inPath, over)
	if err != nil {
		return err
	}
	if planOnly {
		return writePlan(req, outPath)
	}
	res, err := dispatch(req, bulk.Options{}, progress)
	if err != nil && !errors.Is(err, adapter.ErrAborted) {
		return err
	}
	// An abort still writes: the caller asked for the run to stop, and
	// the partial result says how far it got. The error is returned
	// afterwards so main can exit on it.
	b, marshalErr := json.Marshal(res)
	if marshalErr != nil {
		return fmt.Errorf("marshalling the result: %w", marshalErr)
	}
	if writeErr := write(outPath, b); writeErr != nil {
		return writeErr
	}
	return err
}

// dispatch picks the pipeline the request's kind asks for. The kind is
// derived from the request, never passed as a flag: an operator who
// could say "run this as a bulk" could say it of a request with no
// bulk block.
func dispatch(req api.SimRequest, opt bulk.Options, progress io.Writer) (api.SimResult, error) {
	switch req.Kind() {
	case api.KindGear, api.KindTalents, api.KindDrops:
		return executeBulk(req, opt, progress)
	case api.KindWeights:
		// Task 24 adds the weights run; until then the binary refuses
		// rather than silently running a plain sim of a request that
		// asked for something else.
		return api.SimResult{}, fmt.Errorf("%w: weights requests are not yet supported by this binary", errBadInput)
	default:
		return Execute(req, progress)
	}
}

// writePlan answers -plan: how many combinations, and what the first
// stage would be, without running a sim.
//
// The API calls it to count a request at submit time. It could not
// call bulk.Count directly: sim/bulk reaches sim/internal/simdb for
// the item rows, and an internal package is not importable from
// another module. So the binary the job already ships is the door.
//
// A cap breach comes back as the SAME structured error the wasm
// returns, so the API answers 400 cap_exceeded from either lane
// without parsing a sentence.
func writePlan(req api.SimRequest, outPath string) error {
	if req.Bulk == nil {
		return fmt.Errorf("%w: -plan takes a bulk request; this one has none", errBadInput)
	}
	stage, err := bulk.PlanWith(req, bulk.Options{})
	if err != nil {
		var capped api.ErrCapExceeded
		if errors.As(err, &capped) {
			b, marshalErr := json.Marshal(struct {
				Error        string `json:"error"`
				Cap          int    `json:"cap"`
				Combinations int    `json:"combinations"`
			}{"cap_exceeded", capped.Cap, capped.Combinations})
			if marshalErr != nil {
				return marshalErr
			}
			if writeErr := write(outPath, b); writeErr != nil {
				return writeErr
			}
			return fmt.Errorf("%w: %v", errBadInput, err)
		}
		return fmt.Errorf("%w: %v", errBadInput, err)
	}
	b, err := json.Marshal(struct {
		Combinations int                `json:"combinations"`
		Stage        bulk.StageRequests `json:"stage"`
	}{len(stage.Combos), stage})
	if err != nil {
		return err
	}
	return write(outPath, b)
}

// runProto is run with the engine's own result as the output instead of
// ours. It exists for one caller: sim/adapter's fixtures are checked-in
// RaidSimResult protobufs, so regenerating them needs the thing before
// the adapter rather than after it. Everything else wants run.
//
// planOnly is accepted only so runProto's signature matches run's: main
// assigns whichever of the two -out-proto picks to one local variable
// and calls it uniformly. runProto's one caller regenerates fixtures
// from a plain request, so the flag means nothing here.
func runProto(inPath, outProtoPath string, over overrides, _ bool, progress io.Writer) error {
	req, err := load(inPath, over)
	if err != nil {
		return err
	}
	engineRes, err := execute(req, progress)
	if err != nil {
		return err
	}
	// The cast log is one iteration of text; it makes the fixture large
	// without changing anything the adapter reads.
	engineRes.Logs = ""
	b, err := googleproto.Marshal(engineRes)
	if err != nil {
		return fmt.Errorf("marshalling the engine result: %w", err)
	}
	return write(outProtoPath, b)
}

// load reads a SimRequest and applies the command line's overrides.
func load(inPath string, over overrides) (api.SimRequest, error) {
	var raw []byte
	var err error
	if inPath == "-" {
		raw, err = io.ReadAll(os.Stdin)
	} else {
		raw, err = os.ReadFile(inPath)
	}
	if err != nil {
		return api.SimRequest{}, fmt.Errorf("%w: reading the request: %v", errBadInput, err)
	}

	// Strictly: a field the envelope does not carry is a caller sending
	// something this build cannot honour, and accepting it would run a
	// sim that quietly ignored part of the request.
	var req api.SimRequest
	dec := json.NewDecoder(strings.NewReader(string(raw)))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		return api.SimRequest{}, fmt.Errorf("%w: the input is not a SimRequest: %v", errBadInput, err)
	}
	if over.iterations > 0 {
		req.Iterations = over.iterations
	}
	// -no-sample can only take the sample away, never add one: a
	// request that already opted out stays opted out.
	if over.noSample {
		req.NoSample = true
	}
	// An operator who wrote no engine version means this binary's.
	// Naming a different one is refused by api.SimRequest.Validate.
	if req.EngineVersion == "" {
		req.EngineVersion = enginever.Version
	}
	return req, nil
}

// write puts bytes where the caller asked, with - meaning stdout.
func write(path string, b []byte) error {
	if path == "-" {
		_, err := os.Stdout.Write(b)
		return err
	}
	return os.WriteFile(path, b, 0o644)
}

// Execute is the whole pipeline: our envelope in, the engine in the
// middle, our envelope out. sim/cmd/wasm calls the same three steps.
func Execute(req api.SimRequest, progress io.Writer) (api.SimResult, error) {
	start := time.Now()
	// EngineVersion is this binary's own, never the request's claim: a
	// row's provenance is a fact about what produced it, and
	// api.SimResult.Stale can only ever fire if it is one.
	base := api.SimResult{EngineVersion: enginever.Version, Request: req, Lane: api.LaneServer}
	engineRes, err := execute(req, progress)
	if errors.Is(err, adapter.ErrAborted) {
		base.Aborted = true
		base.DurationMS = time.Since(start).Milliseconds()
		if engineRes != nil {
			base.IterationsRun = int(engineRes.IterationsDone)
		}
		// An abort folded no fight, but it still carries a summary of
		// the same SHAPE as a finished one: a zero summary.Summary
		// marshals its sixteen lists as null, and the page that renders
		// a stopped run would need a null check per key.
		base.Summary = adapter.EmptySummary()
		return base, err
	}
	if err != nil {
		return api.SimResult{}, err
	}
	sum, err := adapter.Summarize(engineRes, req)
	if err != nil {
		return api.SimResult{}, fmt.Errorf("adapting the result: %w", err)
	}
	base.DPS = adapter.DPS(engineRes)
	base.IterationsRun = int(engineRes.IterationsDone)
	base.DurationMS = time.Since(start).Milliseconds()
	base.Summary = sum
	base.Sample = adapter.Sample(engineRes)
	return base, nil
}

// runs counts the sims this process has started, so each gets a signal
// id of its own.
var runs atomic.Int64

func runID() string {
	return fmt.Sprintf("forever-sim-%d-%d", os.Getpid(), runs.Add(1))
}

// simEntryPoint is the shared signature of the engine's two ways to run
// a request: core.RunRaidSimAsync and core.RunRaidSimConcurrentAsync.
type simEntryPoint func(*proto.RaidSimRequest, chan *proto.ProgressMetrics, string)

// entryPointFor picks which of the engine's two entry points runs one
// request.
//
// The concurrent path (core.RunRaidSimConcurrentAsync) splits the
// iteration count across runtime.NumCPU() shards and, per the engine
// fork's own pickSampleIteration (sim/core/sim_concurrent.go), keeps
// whichever shard's LOCAL median lands closest to the combined mean -
// its own comment calls this "a KNOWN APPROXIMATION, not the genuine
// global median". The single-threaded path (core.RunRaidSimAsync, the
// same one sim/cmd/wasm uses because wasm has no threads to split
// across) computes the exact median over the whole run and never goes
// through that approximation.
//
// A sample iteration exists to show the player one real fight, so a
// request that asked for one gets the exact entry point rather than the
// fast, approximate one; every other request keeps the concurrent
// split. DO NOT "optimise" this back to always using the concurrent
// call - that would swap an exact sample for a plausible-looking one.
func entryPointFor(sampleIteration bool) simEntryPoint {
	if sampleIteration {
		return core.RunRaidSimAsync
	}
	return core.RunRaidSimConcurrentAsync
}

// execute is the engine half: our request in, the engine's own result
// out, with progress reported as JSON lines along the way.
func execute(req api.SimRequest, progress io.Writer) (*proto.RaidSimResult, error) {
	// A test binary and any caller that is not main() still needs the
	// registry, and sync.Once makes the second call free and safe.
	registerOnce.Do(engine.RegisterAll)

	// OpenIterations: -iterations is an operator override, so the count
	// is bounded rather than held to the closed set the settings bar
	// offers. Without it `-iterations 100` died on "iterations must be
	// one of [500 3000 10000]", which the flag's own help denies.
	engineReq, err := request.BuildWith(req, request.Options{OpenIterations: true, NoSampleIteration: req.NoSample})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errBadInput, err)
	}
	// Forever's own item rows, from the active build. Without them an
	// item id resolves to nothing and the engine dies mid-run.
	if err := simdb.Attach(engineReq); err != nil {
		return nil, err
	}

	reporter := make(chan *proto.ProgressMetrics, 32)
	// A fresh id per run: simsignals.RegisterWithId refuses a duplicate,
	// so a constant would collide the moment one process ran two sims -
	// which the tests in this package already do.
	id := runID()
	defer onInterrupt(id)()
	entryPointFor(engineReq.SimOptions.GetSampleIteration())(engineReq, reporter, id)

	// A nil tick is the no-progress case; ToResult takes one.
	var tick func(*proto.ProgressMetrics)
	if progress != nil {
		enc := json.NewEncoder(progress)
		tick = func(p *proto.ProgressMetrics) {
			_ = enc.Encode(struct {
				Completed int32   `json:"completed"`
				Total     int32   `json:"total"`
				DPS       float64 `json:"dps"`
			}{p.CompletedIterations, p.TotalIterations, p.Dps})
		}
	}

	// The drain loop lives in sim/internal/simdrain because sim/cmd/wasm
	// runs the same one; its invariant (why this does not break on the
	// first FinalRaidResult) is documented on ToResult.
	engineRes := simdrain.ToResult(reporter, tick)
	if engineRes == nil {
		return nil, errors.New("the engine produced no result")
	}
	// Switching on the type, not on the message: an abort's
	// ErrorOutcome carries none, so a message test reported Ctrl-C as
	// a corrupt result.
	if err := adapter.ResultError(engineRes); err != nil {
		if errors.Is(err, adapter.ErrAborted) {
			return engineRes, err
		}
		return nil, fmt.Errorf("the sim failed: %w", err)
	}
	return engineRes, nil
}

// registerOnce guards the engine's spell registry. The engine's own
// guard is an unsynchronised package bool.
var registerOnce sync.Once

// onInterrupt makes SIGINT and SIGTERM abort the run with the id given,
// and returns the function that takes the handler back down. The engine
// answers an abort with a result rather than a panic, so the process
// gets to write what completed instead of dying mid-file.
func onInterrupt(id string) func() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	done := make(chan struct{})
	go func() {
		select {
		case <-ch:
			simsignals.AbortById(id)
		case <-done:
		}
	}()
	return func() {
		signal.Stop(ch)
		close(done)
	}
}
