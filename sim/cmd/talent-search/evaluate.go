package main

import (
	"fmt"
	"sort"
)

// result is one build's final standing against the guide.
type result struct {
	candidate
	DPS         estimate
	Delta       float64
	DeltaErr    float64
	Significant bool // |Delta| > DeltaErr
}

// evaluation is a spec's whole sim outcome.
type evaluation struct {
	Guide    result
	Final    []result // best first, guide included
	Screened []result // every candidate at screening iterations, best first
}

// evaluate screens every candidate (and the guide) at screenIters,
// then re-sims the guide, the best topN and the best candidate clean
// admits (if the topN missed it) at finalIters.
func evaluate(guide build, cands []candidate, run dpsFunc, screenIters, finalIters, topN int, clean func(build) bool) (evaluation, error) {
	guideCand := candidate{Label: "guide", Build: guide}
	screened, err := simAll(append([]candidate{guideCand}, cands...), run, screenIters)
	if err != nil {
		return evaluation{}, err
	}
	finalists := []candidate{guideCand}
	cleanIn := false
	for _, r := range screened {
		if r.Label == guideCand.Label {
			continue
		}
		top := len(finalists) <= topN
		if top || (!cleanIn && clean(r.Build)) {
			finalists = append(finalists, r.candidate)
			cleanIn = cleanIn || clean(r.Build)
		}
		if !top && cleanIn {
			break
		}
	}
	final, err := simAll(finalists, run, finalIters)
	if err != nil {
		return evaluation{}, err
	}
	var guideRes result
	for _, r := range final {
		if r.Label == guideCand.Label {
			guideRes = r
		}
	}
	return evaluation{Guide: guideRes, Final: final, Screened: screened}, nil
}

// simAll sims every candidate at iterations and returns them best
// first, each with its delta against the first candidate (the guide).
func simAll(cands []candidate, run dpsFunc, iterations int) ([]result, error) {
	out := make([]result, 0, len(cands))
	for _, c := range cands {
		est, err := run(c.Build, iterations)
		if err != nil {
			return nil, fmt.Errorf("simming %q: %w", c.Label, err)
		}
		out = append(out, result{candidate: c, DPS: est})
	}
	base := out[0].DPS
	for i := range out {
		out[i].Delta = out[i].DPS.Mean - base.Mean
		out[i].DeltaErr = combinedErr(out[i].DPS, base)
		out[i].Significant = i > 0 && abs64(out[i].Delta) > out[i].DeltaErr
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].DPS.Mean > out[j].DPS.Mean })
	return out, nil
}

func abs64(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
