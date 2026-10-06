package main

// Read-only validation over the PUBLISHED, committed bis/*.json files -
// this lane's brief (ranker-weights-anchor), item 5: walk every
// data/builds/<build>/bis/<spec>.json band and list the ones whose
// primary-stat anchor row is insignificant, or whose own
// scale_reference_stat is not the spec's primary stat at all. It never
// writes or regenerates anything (the nightly workflow, bis.yml, owns
// that - this lane never commits regenerated data, per its own brief).
//
// TestCheckWeightsAnchorReportsTodaysCommittedFiles (below) is
// deliberately non-failing: it logs what it finds via t.Logf rather
// than t.Errorf, because the files committed before this lane's own
// fix landed are EXPECTED to show the old, pre-primary-stat anchoring
// convention - that "before" count is exactly what this lane's own
// report (design/reviews/2026-10-05-weights-anchor-report.md) needs,
// not a CI failure that would block on the nightly regenerating them.

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

// weightsAnchorIssue is one band+faction this check flagged, and why.
type weightsAnchorIssue struct {
	Spec    string
	Band    int
	Faction string
	Reason  string
}

// checkWeightsAnchorCoverage walks every data/builds/*/bis/*.json file
// under repoRoot and flags a band+faction when:
//   - this spec's own primary-stat anchor row (primaryAnchorStat,
//     primary_stat.go) is published Insignificant (should never happen
//     once ranker-weights-anchor's own forceAnchorRowSignificant,
//     weights.go, is in the pipeline that generated the file - a file
//     flagging this predates that fix), or
//   - the band's own ScaleReferenceStat (the anchor normalizeScaleFactors
//     actually used) is not the spec's primary-stat row at all - either
//     "" (no anchor found) or some other stat entirely (the pre-fix
//     "largest significant weight" rule).
//
// A band whose WeightsReason is already set (the whole sweep was
// untrustworthy - weightRow's own doc) is skipped: normalizeScaleFactors
// deliberately publishes no anchor there regardless of this lane's own
// fix, so it is not a "before" defect this check should count.
func checkWeightsAnchorCoverage(repoRoot string) ([]weightsAnchorIssue, error) {
	paths, err := filepath.Glob(filepath.Join(repoRoot, "data", "builds", "*", "bis", "*.json"))
	if err != nil {
		return nil, fmt.Errorf("globbing data/builds/*/bis/*.json: %w", err)
	}
	var issues []weightsAnchorIssue
	for _, path := range paths {
		b, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("reading %s: %w", path, err)
		}
		var report specReport
		if err := json.Unmarshal(b, &report); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", path, err)
		}
		spec, err := loadSpec(repoRoot, report.Spec)
		if err != nil {
			return nil, fmt.Errorf("%s: loading spec %s: %w", path, report.Spec, err)
		}
		primary, ok := primaryAnchorStat(spec)
		if !ok {
			issues = append(issues, weightsAnchorIssue{
				Spec:   report.Spec,
				Reason: "no primary_stat entry in primaryStatBySpec",
			})
			continue
		}
		for _, band := range report.Bands {
			if band.WeightsReason != "" {
				continue
			}
			if band.ScaleReferenceStat != primary {
				reason := fmt.Sprintf("scale_reference_stat=%q, want primary stat %q", band.ScaleReferenceStat, primary)
				issues = append(issues, weightsAnchorIssue{Spec: report.Spec, Band: band.Band, Faction: band.Faction, Reason: reason})
			}
			for _, row := range band.Weights {
				if row.Stat == primary && row.Insignificant {
					issues = append(issues, weightsAnchorIssue{
						Spec: report.Spec, Band: band.Band, Faction: band.Faction,
						Reason: fmt.Sprintf("primary stat %q published insignificant", primary),
					})
				}
			}
		}
	}
	return issues, nil
}

// TestCheckWeightsAnchorReportsTodaysCommittedFiles runs
// checkWeightsAnchorCoverage over this repository's own real,
// committed data/builds/*/bis/*.json files and logs a per-spec issue
// count - this lane's brief, item 5's own "before" measurement. It
// never fails the build: a build predating ranker-weights-anchor's own
// fix (every file committed before this lane) is EXPECTED to show
// issues, and the nightly workflow (bis.yml), not this test, is what
// regenerates them. Skipped the same way TestPublishedBISFilesPass
// SanityChecks's own doc describes when no bis directory exists at
// all under publishedRepoRoot (a repo clone with no build ever ranked).
func TestCheckWeightsAnchorReportsTodaysCommittedFiles(t *testing.T) {
	issues, err := checkWeightsAnchorCoverage(publishedRepoRoot)
	if err != nil {
		t.Fatalf("checkWeightsAnchorCoverage: %v", err)
	}
	if len(issues) == 0 {
		t.Log("checkWeightsAnchorCoverage: no issues found (or no bis/*.json files are committed yet)")
		return
	}
	counts := make(map[string]int, len(issues))
	for _, issue := range issues {
		counts[issue.Spec]++
	}
	specs := make([]string, 0, len(counts))
	for spec := range counts {
		specs = append(specs, spec)
	}
	sort.Strings(specs)
	t.Logf("checkWeightsAnchorCoverage: %d total issue(s) across %d spec(s):", len(issues), len(specs))
	for _, spec := range specs {
		t.Logf("  %s: %d band+faction row(s)", spec, counts[spec])
	}
}
