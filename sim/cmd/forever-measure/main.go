// Command forever-measure recovers combat constants from a real combat
// log: the numbers the client tables do not carry.
//
// It prints a table. The measurement functions live in sim/measure and
// return values, because the nightly validation job consumes the same
// code and wants the numbers rather than a table on stdout.
//
//	forever-measure -log dummy.txt -actor "Yourname-Forever" \
//	  -spell-power 500 -attack-power 1200
//
// See sim/README.md for the dummy-target procedure that produces a log
// this can read.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/jhunthrop/foreversixty/sim/measure"
)

const (
	exitOK       = 0
	exitFailed   = 1
	exitBadFlags = 2
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

// run is main's body with its arguments and its two streams as
// parameters, and an exit code instead of an os.Exit, so the flag
// handling and the failure paths are testable rather than only the
// measurement behind them.
func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("forever-measure", flag.ContinueOnError)
	fs.SetOutput(stderr)
	logPath := fs.String("log", "", "the combat log to read")
	actor := fs.String("actor", "", "the player to measure, by name")
	spellPower := fs.Float64("spell-power", 0, "override the spell power the log reports for the actor")
	attackPower := fs.Float64("attack-power", 0, "override the attack power the log reports for the actor")
	minSamples := fs.Int("min-samples", measure.DefaultMinSamples, "figures under this many samples print as insufficient data")
	asJSON := fs.Bool("json", false, "print the report as JSON instead of a table")
	if err := fs.Parse(args); err != nil {
		return exitBadFlags
	}

	if *logPath == "" || *actor == "" {
		fmt.Fprintln(stderr, "forever-measure: -log and -actor are both required")
		fs.Usage()
		return exitBadFlags
	}

	// A log's timestamps carry no year in every dialect, so the clock is
	// seeded from the file's modification time, which is what a batch
	// parse does. Forever's own stamps carry the year and ignore it.
	base := time.Now()
	if fi, err := os.Stat(*logPath); err == nil {
		base = fi.ModTime()
	}

	events, err := measure.Load(*logPath, base)
	if err != nil {
		fmt.Fprintln(stderr, "forever-measure:", err)
		return exitFailed
	}
	rep, err := measure.Run(measure.Input{
		Events:      events,
		Actor:       *actor,
		SpellPower:  *spellPower,
		AttackPower: *attackPower,
		MinSamples:  *minSamples,
	})
	if err != nil {
		fmt.Fprintln(stderr, "forever-measure:", err)
		return exitFailed
	}

	if *asJSON {
		enc := json.NewEncoder(stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(rep); err != nil {
			fmt.Fprintln(stderr, "forever-measure:", err)
			return exitFailed
		}
		return exitOK
	}
	fmt.Fprint(stdout, rep.Table())
	return exitOK
}
