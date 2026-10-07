package main

import (
	"fmt"
	"log"
	"math"
	"runtime"
	"sort"
	"sync"
)

// verdict is helps / hurts / no effect against the combined error.
type verdict string

const (
	verdictHelps    verdict = "helps"
	verdictHurts    verdict = "hurts"
	verdictNoEffect verdict = "no effect"
	// verdictError is a mutation or probe candidate this search could
	// not sim at all (the engine returned an error, or - rarer -
	// panicked, which runParallel recovers from rather than letting
	// take the whole run down with it). It is excluded from every
	// wrong/missing/best conclusion rather than silently scored as
	// "no effect".
	verdictError verdict = "sim error"
)

func classifyVerdict(delta, err float64) verdict {
	switch {
	case delta > err:
		return verdictHelps
	case delta < -err:
		return verdictHurts
	default:
		return verdictNoEffect
	}
}

// probeRow is one action-probe line: either "remove: <action>" (does
// the rotation need this line?) or "insert: <spell>" (is a learned
// spell missing?). Delta is the ACTION's own measured contribution:
// for a removal, baseline-minus-removed (positive means the action
// helps); for an insertion, inserted-minus-baseline (positive means
// the missing spell helps).
type probeRow struct {
	Kind    string // "remove" or "insert"
	Label   string
	Delta   float64
	Err     float64
	Verdict verdict
	Error   string // set only when Verdict is verdictError
}

// runProbe is step 2, the "is it right" answer: every priority-list
// action removed, and every probe-worthy learned spell the rotation
// does not cast inserted at the top with its default condition - each
// sim'd once at o.confirmIterations against the baseline estimate. A
// candidate the engine cannot sim (a malformed mutation, an engine
// error) is reported as a sim-error row rather than failing the whole
// probe - one bad candidate should never hide every other row's
// findings.
func runProbe(base rotation, candidates []learnedCandidate, baseline estimate, run dpsFunc, iterations int, names map[int]string) []probeRow {
	type job struct {
		kind, label string
		rot         rotation
	}
	var jobs []job
	for i, e := range base.PriorityList {
		next := removeEntry(base.PriorityList, e)
		label := fmt.Sprintf("#%d (%s)", i+1, entryActionLabel(e.Action, names))
		jobs = append(jobs, job{kind: "remove", label: label, rot: base.withPriorityList(next)})
	}
	for _, c := range candidates {
		newEntry := buildCastEntry(fmt.Sprintf("rotation-search probe: %s (%s)", c.Name, c.ConditionLabel), c.ID, c.Rank, c.Condition)
		next := append(cloneEntries([]entry{newEntry}), cloneEntries(base.PriorityList)...)
		jobs = append(jobs, job{kind: "insert", label: c.Name, rot: base.withPriorityList(next)})
	}

	rows := make([]probeRow, len(jobs))
	for i, j := range jobs {
		// Pre-filled so a recovered panic (see runParallel's own doc)
		// leaves a labeled sim-error row behind rather than a blank
		// zero-value one that would misread as "no effect".
		rows[i] = probeRow{Kind: j.kind, Label: j.label, Verdict: verdictError, Error: "did not finish"}
	}
	runParallel(len(jobs), func(i int) error {
		j := jobs[i]
		est, err := run(j.rot, iterations)
		if err != nil {
			rows[i] = probeRow{Kind: j.kind, Label: j.label, Verdict: verdictError, Error: err.Error()}
			return nil
		}
		var delta float64
		if j.kind == "remove" {
			delta = baseline.Mean - est.Mean
		} else {
			delta = est.Mean - baseline.Mean
		}
		e := combinedErr(est, baseline)
		rows[i] = probeRow{Kind: j.kind, Label: j.label, Delta: delta, Err: e, Verdict: classifyVerdict(delta, e)}
		return nil
	})
	sort.SliceStable(rows, func(i, j int) bool {
		if rows[i].Kind != rows[j].Kind {
			return rows[i].Kind == "remove"
		}
		return rows[i].Label < rows[j].Label
	})
	return rows
}

func removeEntry(list []entry, target entry) []entry {
	out := make([]entry, 0, len(list))
	removed := false
	for _, e := range list {
		if !removed && sameEntry(e, target) {
			removed = true
			continue
		}
		out = append(out, cloneEntries([]entry{e})[0])
	}
	return out
}

// sameEntry compares two entries by their rendered JSON - good enough
// identity for the one caller (removeEntry), which always removes an
// entry drawn from the same list it is searching.
func sameEntry(a, b entry) bool {
	return fmt.Sprintf("%v|%v|%v", a.Notes, a.Action, a.DoAtValue) == fmt.Sprintf("%v|%v|%v", b.Notes, b.Action, b.DoAtValue)
}

// acceptedMutation is one round's winning move.
type acceptedMutation struct {
	Round         int
	Label         string
	Gain, GainErr float64
}

// huntBest is step 3's greedy hill-climb: each round generates every
// mutation of the incumbent, screens them all at screenIters, and
// keeps the best if it beats the incumbent beyond combined error;
// stops when no mutation wins or after rounds rounds. A mutation the
// engine cannot sim is scored at -Inf (never chosen, never fatal) and
// logged.
func huntBest(base rotation, candidates []learnedCandidate, ticks map[int]float64, run dpsFunc, screenIters, rounds int, baseline estimate, names map[int]string) (rotation, estimate, []acceptedMutation) {
	current, currentEst := base, baseline
	var accepted []acceptedMutation
	for round := 1; round <= rounds; round++ {
		muts := allMutations(current, candidates, ticks, names)
		if len(muts) == 0 {
			break
		}
		ests := make([]estimate, len(muts))
		for i := range ests {
			ests[i] = estimate{Mean: math.Inf(-1)} // see runParallel's doc: a recovered panic leaves this sentinel in place.
		}
		runParallel(len(muts), func(i int) error {
			est, err := run(muts[i].Rotation, screenIters)
			if err != nil {
				log.Printf("rotation-search: round %d: %q could not be simmed, skipping: %v", round, muts[i].Label, err)
				ests[i] = estimate{Mean: math.Inf(-1)}
				return nil
			}
			ests[i] = est
			return nil
		})
		bestIdx := 0
		for i := 1; i < len(ests); i++ {
			if ests[i].Mean > ests[bestIdx].Mean {
				bestIdx = i
			}
		}
		best := muts[bestIdx]
		bestEst := ests[bestIdx]
		if math.IsInf(bestEst.Mean, -1) {
			break // every mutation this round failed to sim.
		}
		gain := bestEst.Mean - currentEst.Mean
		gainErr := combinedErr(bestEst, currentEst)
		if gain <= gainErr {
			break
		}
		current, currentEst = best.Rotation, bestEst
		accepted = append(accepted, acceptedMutation{Round: round, Label: best.Label, Gain: gain, GainErr: gainErr})
	}
	return current, currentEst, accepted
}

// allMutations is one round's whole mutation pool: swap, remove,
// insert, refresh-condition variants, resource-gate thresholds, and
// the execute-gate toggle, all generated fresh against current so
// every mutation always starts from the search's own incumbent.
func allMutations(current rotation, candidates []learnedCandidate, ticks map[int]float64, names map[int]string) []mutation {
	var out []mutation
	out = append(out, swapAdjacentMutations(current, names)...)
	out = append(out, removeActionMutations(current, names)...)
	out = append(out, insertCandidateMutations(current, candidates)...)
	out = append(out, replaceMaintenanceMutations(current, candidates, names)...)
	out = append(out, refreshConditionMutations(current, ticks, names)...)
	out = append(out, resourceGateMutations(current, names)...)
	out = append(out, toggleExecuteGateMutations(current, names)...)
	return out
}

// runParallel calls work(i) for every i in [0, n) across a bounded
// pool of goroutines - the inproc path is safe for concurrent calls
// (see run.go's engineRun doc), and screening dozens to hundreds of
// mutations a round is the one place this search's wall-clock time
// actually lives. A panic inside work (an engine bug a pathological
// mutation triggers, say) is recovered per job and turned into the
// error work itself would have returned, so one bad mutation cannot
// take the rest of a round - or the process - down with it.
func runParallel(n int, work func(i int) error) {
	if n == 0 {
		return
	}
	workers := runtime.GOMAXPROCS(0)
	if workers > n {
		workers = n
	}
	if workers < 1 {
		workers = 1
	}
	var wg sync.WaitGroup
	next := make(chan int)
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range next {
				safeWork(work, i)
			}
		}()
	}
	for i := 0; i < n; i++ {
		next <- i
	}
	close(next)
	wg.Wait()
}

// safeWork runs work(i), recovering a panic into a logged warning so
// the caller's own error handling (each work closure records its
// result at index i before returning) is never skipped.
func safeWork(work func(i int) error, i int) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("rotation-search: job %d panicked, treating as a sim error: %v", i, r)
		}
	}()
	if err := work(i); err != nil {
		log.Printf("rotation-search: job %d: %v", i, err)
	}
}
