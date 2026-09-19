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
// With -progress every line of stderr is JSON: our own
// {"completed","total","dps"} ticks, and the engine's log output wrapped
// as {"log": "..."} rather than interleaved raw.
//
// Concurrency is automatic: core.RunRaidSimConcurrentAsync splits across
// runtime.NumCPU() and recombines the distribution metrics, offsetting
// each split's seed so the stream matches a serial run.
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync/atomic"
	"time"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/request"
	engine "github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	googleproto "google.golang.org/protobuf/proto"
)

// Version is set by the makefile to the short sha of the engine the
// binary was built against: the same string as enginever.Version.
var Version = "dev"

const (
	exitOK      = 0
	exitSimFail = 1
	exitBadArgs = 2
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
		fmt.Println(Version)
		os.Exit(exitOK)
	}

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
		if errors.Is(err, errBadInput) {
			os.Exit(exitBadArgs)
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
	if err != nil {
		return err
	}
	b, err := json.Marshal(res)
	if err != nil {
		return fmt.Errorf("marshalling the result: %w", err)
	}
	return write(outPath, b)
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

	var req api.SimRequest
	if err := json.Unmarshal(raw, &req); err != nil {
		return api.SimRequest{}, fmt.Errorf("%w: the input is not a SimRequest: %v", errBadInput, err)
	}
	if iterations > 0 {
		req.Iterations = iterations
	}
	if req.EngineVersion == "" {
		req.EngineVersion = Version
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
	engineRes, err := execute(req, progress)
	if err != nil {
		return api.SimResult{}, err
	}
	sum, err := adapter.Summarize(engineRes, req)
	if err != nil {
		return api.SimResult{}, fmt.Errorf("adapting the result: %w", err)
	}
	return api.SimResult{
		EngineVersion: req.EngineVersion,
		Request:       req,
		Lane:          api.LaneServer,
		DPS:           adapter.DPS(engineRes),
		IterationsRun: int(engineRes.IterationsDone),
		DurationMS:    time.Since(start).Milliseconds(),
		Summary:       sum,
	}, nil
}

// runs counts the sims this process has started, so each gets a signal
// id of its own.
var runs atomic.Int64

func runID() string {
	return fmt.Sprintf("forever-sim-%d-%d", os.Getpid(), runs.Add(1))
}

// execute is the engine half: our request in, the engine's own result
// out, with progress reported as JSON lines along the way.
func execute(req api.SimRequest, progress io.Writer) (*proto.RaidSimResult, error) {
	engine.RegisterAll()

	engineReq, err := request.Build(req)
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
	core.RunRaidSimConcurrentAsync(engineReq, reporter, runID())

	var enc *json.Encoder
	if progress != nil {
		enc = json.NewEncoder(progress)
	}

	var engineRes *proto.RaidSimResult
	for p := range reporter {
		if p.FinalRaidResult != nil {
			engineRes = p.FinalRaidResult
			break
		}
		if enc != nil {
			_ = enc.Encode(struct {
				Completed int32   `json:"completed"`
				Total     int32   `json:"total"`
				DPS       float64 `json:"dps"`
			}{p.CompletedIterations, p.TotalIterations, p.Dps})
		}
	}
	if engineRes == nil {
		return nil, errors.New("the engine produced no result")
	}
	if engineRes.Error != nil && engineRes.Error.Message != "" {
		return nil, fmt.Errorf("the sim failed: %s", engineRes.Error.Message)
	}
	return engineRes, nil
}
