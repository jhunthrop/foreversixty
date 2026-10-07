// Command rotation-search sims rotation variants for one spec at
// level 60 and reports whether the curated APL is right for the
// spec's guide talent build, and whether a better rotation exists.
//
// Gear and talents are held fixed - the committed level-60 BiS band's
// gear and the guide's own FS1 build - exactly what sim/cmd/talent-search
// holds fixed on the OTHER axis (it varies talents over fixed gear and
// a fixed rotation). Step 1 sims the curated rotation as a baseline.
// Step 2 (the "is it right" answer) removes each priority-list action
// in turn and inserts each learned-but-uncast damage or DPS-cooldown
// spell in turn, each against the baseline, to find wrong and missing
// lines. Step 3 greedily hill-climbs over swap/remove/insert/
// refresh-condition/resource-gate/execute-gate mutations, screening
// every mutation each round and keeping the best if it beats the
// incumbent beyond combined error. A markdown report goes to
// design/reviews/rotation-search/<spec>.md. Nothing under data/ or
// web/ is ever written.
package main

import (
	"errors"
	"fmt"
	"github.com/jhunthrop/foreversixty/sim/request"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	err := run(os.Args[1:])
	switch {
	case errors.Is(err, errSkipped):
		log.Print(err)
		os.Exit(exitSkipped)
	case err != nil:
		log.Fatal(err)
	}
}

func run(args []string) error {
	o, err := parseOptions(args)
	if err != nil {
		return err
	}
	start := time.Now()
	in, err := prepare(o)
	if err != nil {
		return err
	}
	runDPS, err := engineRun(in.setup)
	if err != nil {
		return err
	}

	names, err := buildSpellNames(o.repoRoot, in.clientBuild, in.setup.spec.ClassSlug, o.level)
	if err != nil {
		return err
	}

	baseline, err := runDPS(in.base, o.confirmIterations)
	if err != nil {
		return fmt.Errorf("simming the baseline: %w", err)
	}
	log.Printf("rotation-search: %s: baseline %.1f ± %.1f DPS (%d candidates to probe)", o.spec, baseline.Mean, baseline.Err, len(in.candidates))

	probe := runProbe(in.base, in.candidates, baseline, runDPS, o.confirmIterations, names)

	rep := report{opts: o, in: in, baseline: baseline, probe: probe, probeOnly: o.probeOnly, names: names}
	if !o.probeOnly {
		best, bestEst, accepted := huntBest(in.base, in.candidates, in.tickSeconds, runDPS, o.iterations, o.rounds, baseline, names)
		// Re-confirm the final winner at the confirm iteration count,
		// per the design (screening and confirming use different
		// iteration counts, so the reported winner's own number is not
		// just carried over from a cheaper screening run).
		if bestEst, err = runDPS(best, o.confirmIterations); err != nil {
			return err
		}
		rep.best, rep.bestEst, rep.accepted = best, bestEst, accepted
		if rep.finishers, err = tallyFinishers(in, in.base, best, o.confirmIterations, names); err != nil {
			return err
		}
	}
	rep.elapsed = time.Since(start)

	if err := os.MkdirAll(o.out, 0o755); err != nil {
		return err
	}
	path := filepath.Join(o.out, reportName(o))
	if err := os.WriteFile(path, []byte(rep.markdown()), 0o644); err != nil {
		return err
	}
	log.Printf("rotation-search: %s: wrote %s in %s", o.spec, path, time.Since(start).Round(time.Second))
	return nil
}

// reportName is the report file for one search: the bare search keeps its
// historical <spec>.md, any other preset gets its own file beside it.
func reportName(o options) string {
	if o.preset == request.BarePreset {
		return o.spec + ".md"
	}
	return o.spec + "-" + o.preset + ".md"
}
