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
	"os"
	"time"

	"github.com/jhunthrop/foreversixty/sim/measure"
)

func main() {
	logPath := flag.String("log", "", "the combat log to read")
	actor := flag.String("actor", "", "the player to measure, by name")
	spellPower := flag.Float64("spell-power", 0, "override the spell power the log reports for the actor")
	attackPower := flag.Float64("attack-power", 0, "override the attack power the log reports for the actor")
	minSamples := flag.Int("min-samples", measure.DefaultMinSamples, "figures under this many samples print as insufficient data")
	asJSON := flag.Bool("json", false, "print the report as JSON instead of a table")
	flag.Parse()

	if *logPath == "" || *actor == "" {
		fmt.Fprintln(os.Stderr, "forever-measure: -log and -actor are both required")
		flag.Usage()
		os.Exit(2)
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
		fmt.Fprintln(os.Stderr, "forever-measure:", err)
		os.Exit(1)
	}
	rep, err := measure.Run(measure.Input{
		Events:      events,
		Actor:       *actor,
		SpellPower:  *spellPower,
		AttackPower: *attackPower,
		MinSamples:  *minSamples,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "forever-measure:", err)
		os.Exit(1)
	}

	if *asJSON {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(rep); err != nil {
			fmt.Fprintln(os.Stderr, "forever-measure:", err)
			os.Exit(1)
		}
		return
	}
	fmt.Print(rep.Table())
}
