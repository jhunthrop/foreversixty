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
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
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
	progress := flag.Bool("progress", false, "write JSON-lines progress to stderr; the engine's own log output is wrapped as {\"log\":...} so every line of that stream parses")
	version := flag.Bool("version", false, "print the engine version and exit")
	flag.Parse()

	if *version {
		fmt.Println(enginever.Version)
		os.Exit(exitOK)
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
	if err := run(*in, outPath, *iterations, sink); err != nil {
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

// run is main's body, with its files and its progress sink as parameters
// so it is testable.
func run(inPath, outPath string, iterations int, progress io.Writer) error {
	req, err := load(inPath, iterations)
	if err != nil {
		return err
	}
	res, err := Execute(req, progress)
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

// runProto is run with the engine's own result as the output instead of
// ours. It exists for one caller: sim/adapter's fixtures are checked-in
// RaidSimResult protobufs, so regenerating them needs the thing before
// the adapter rather than after it. Everything else wants run.
func runProto(inPath, outProtoPath string, iterations int, progress io.Writer) error {
	req, err := load(inPath, iterations)
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

// load reads a SimRequest and applies the iteration override.
func load(inPath string, iterations int) (api.SimRequest, error) {
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
	if iterations > 0 {
		req.Iterations = iterations
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

// drainToResult reads reporter until it has the engine's final result,
// reporting every intermediate tick to enc along the way (enc may be
// nil, when the caller asked for no progress output).
//
// It does NOT simply break on the first FinalRaidResult: run() sends
// that message from INSIDE the producer goroutine, then keeps running
// - for a sample request it still has to return up to runSim, which
// replays the median iteration and only THEN sets
// FinalRaidResult.SampleIteration on that same pointer, before closing
// the channel. Breaking on sight of the message used to read the
// struct while that replay was still in flight, which raced
// SampleIteration nil almost every time. Draining to the channel's
// close instead relies on Go's channel-close happens-before: whatever
// the producer did before close(progress), including that mutation,
// is guaranteed visible once range observes the close.
//
// That fix only holds for a SUCCESSFUL run, because the engine's sim
// body is the only path that closes progress on its way out
// (core/sim.go). Two other engine paths send a FinalRaidResult and
// then return WITHOUT ever closing the channel: a failed
// simsignals.RegisterWithId - an empty or duplicate request id -
// inside RunRaidSimAsync/RunRaidSimConcurrentAsync (core/api.go), and
// SimOptions.IsTest, which registers no closing defer at all
// (core/sim.go) - this package's own requests never set IsTest, but
// nothing stops a future caller from being the first to. A blanket
// drain hangs forever on either. Neither ever carries a sample
// (core/sim.go's replay is itself guarded on result.Error == nil), so
// an error result has nothing left worth waiting for: take it and
// stop rather than block on a close that may never come.
func drainToResult(reporter chan *proto.ProgressMetrics, enc *json.Encoder) *proto.RaidSimResult {
	var engineRes *proto.RaidSimResult
	for p := range reporter {
		if p.FinalRaidResult != nil {
			engineRes = p.FinalRaidResult
			if engineRes.Error != nil {
				break
			}
			continue
		}
		if enc != nil {
			_ = enc.Encode(struct {
				Completed int32   `json:"completed"`
				Total     int32   `json:"total"`
				DPS       float64 `json:"dps"`
			}{p.CompletedIterations, p.TotalIterations, p.Dps})
		}
	}
	return engineRes
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

	var enc *json.Encoder
	if progress != nil {
		enc = json.NewEncoder(progress)
	}

	engineRes := drainToResult(reporter, enc)
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
