package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// fakeLevelingBisScript writes a tiny shell script standing in for the
// real leveling-bis binary that runAllSpecsIsolated/runSpecSubprocess
// would otherwise re-exec (os.Args[0]) - never the real thing: a real
// -all run takes 13 minutes and a lot of memory, exactly what this
// lane's brief says never to run from a test. The script ignores its
// flag arguments (real leveling-bis flags, which a plain shell script
// need not parse) and behaves per env var:
//
//	LBIS_FAKE_BEHAVIOR=fail  -> prints to stderr, exits 1
//	LBIS_FAKE_BEHAVIOR=hang  -> sleeps past any sane test timeout
//	anything else            -> exits 0 immediately
func fakeLevelingBisScript(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "fake-leveling-bis.sh")
	script := `#!/bin/sh
case "$LBIS_FAKE_BEHAVIOR" in
  fail)
    echo "fake leveling-bis: forced failure" >&2
    exit 1
    ;;
  hang)
    sleep 300
    ;;
  *)
    exit 0
    ;;
esac
`
	if err := os.WriteFile(path, []byte(script), 0o755); err != nil {
		t.Fatalf("writing fake script: %v", err)
	}
	return path
}

func TestRunSpecSubprocessSuccess(t *testing.T) {
	t.Setenv("LBIS_FAKE_BEHAVIOR", "")
	script := fakeLevelingBisScript(t)
	err := runSpecSubprocess(script, 5*time.Second, "repo", "build", "out", "hunter-marksmanship", "20,30", 5, "")
	if err != nil {
		t.Fatalf("runSpecSubprocess: %v", err)
	}
}

func TestRunSpecSubprocessPropagatesFailure(t *testing.T) {
	t.Setenv("LBIS_FAKE_BEHAVIOR", "fail")
	script := fakeLevelingBisScript(t)
	err := runSpecSubprocess(script, 5*time.Second, "repo", "build", "out", "hunter-marksmanship", "20,30", 5, "")
	if err == nil {
		t.Fatal("runSpecSubprocess with a failing subprocess: want an error, got nil")
	}
}

func TestRunSpecSubprocessTimesOut(t *testing.T) {
	t.Setenv("LBIS_FAKE_BEHAVIOR", "hang")
	script := fakeLevelingBisScript(t)
	err := runSpecSubprocess(script, 100*time.Millisecond, "repo", "build", "out", "hunter-marksmanship", "20,30", 5, "")
	if err == nil {
		t.Fatal("runSpecSubprocess with a hung subprocess: want a timeout error, got nil")
	}
	if !strings.Contains(err.Error(), "did not finish within") {
		t.Fatalf("err = %v, want the timeout message naming the deadline", err)
	}
}

func TestRunSpecSubprocessMemProfileSuffixesPerSpec(t *testing.T) {
	// This test's real assertion is indirect: runSpecSubprocess must not
	// error just because -memprofile was set (the flag is appended, not
	// validated by this function - the fake script ignores it entirely).
	t.Setenv("LBIS_FAKE_BEHAVIOR", "")
	script := fakeLevelingBisScript(t)
	err := runSpecSubprocess(script, 5*time.Second, "repo", "build", "out", "hunter-marksmanship", "20", 5, filepath.Join(t.TempDir(), "heap.prof"))
	if err != nil {
		t.Fatalf("runSpecSubprocess with -memprofile set: %v", err)
	}
}

func TestRunAllSpecsIsolatedRunsEveryWrittenSpec(t *testing.T) {
	t.Setenv("LBIS_FAKE_BEHAVIOR", "")
	script := fakeLevelingBisScript(t)
	// repoRootFixture's writtenSpecs is exactly [hunter-marksmanship]
	// (mage-fire is "draft", priest-shadow has no apl file) - see
	// data_test.go's TestWrittenSpecs for the same fixture asserted
	// directly.
	err := runAllSpecsIsolated(script, 5*time.Second, repoRootFixture, "testbuild", t.TempDir(), defaultBandsFlag, 5, "")
	if err != nil {
		t.Fatalf("runAllSpecsIsolated: %v", err)
	}
}

func TestRunAllSpecsIsolatedOneFailureFailsTheRunButNotBeforeTryingEveryOne(t *testing.T) {
	dir := t.TempDir()
	// Build a two-spec fixture repoRoot so this test can prove BOTH
	// specs are attempted even though the fake subprocess fails every
	// time (main.go's own doc: "should not cost every OTHER spec its
	// BiS list too" - log it, keep going, fail the whole run at the
	// end).
	if err := writeFile(t, filepath.Join(dir, "data", "curated", "specs.json"), `[
		{"spec":"spec-a","class_slug":"hunter","spec_slug":"a","name":"A","tree_index":0,"reference_stat":"agility","weight_stats":["agility"]},
		{"spec":"spec-b","class_slug":"hunter","spec_slug":"b","name":"B","tree_index":0,"reference_stat":"agility","weight_stats":["agility"]}
	]`); err != nil {
		t.Fatal(err)
	}
	if err := writeFile(t, filepath.Join(dir, "data", "curated", "apl", "spec-a.json"), `{"state":"written"}`); err != nil {
		t.Fatal(err)
	}
	if err := writeFile(t, filepath.Join(dir, "data", "curated", "apl", "spec-b.json"), `{"state":"written"}`); err != nil {
		t.Fatal(err)
	}

	t.Setenv("LBIS_FAKE_BEHAVIOR", "fail")
	script := fakeLevelingBisScript(t)
	err := runAllSpecsIsolated(script, 5*time.Second, dir, "testbuild", t.TempDir(), defaultBandsFlag, 5, "")
	if err == nil {
		t.Fatal("runAllSpecsIsolated with every subprocess failing: want an error, got nil")
	}
	if !strings.Contains(err.Error(), "spec-a") || !strings.Contains(err.Error(), "spec-b") {
		t.Fatalf("err = %v, want both spec-a and spec-b named as failed", err)
	}
}

func TestRunUnknownRepoRootErrors(t *testing.T) {
	err := run("leveling-bis", []string{"-repo-root", t.TempDir()})
	if err == nil {
		t.Fatal("run with a repo-root that has no data/curated/specs.json: want an error, got nil")
	}
	if !strings.Contains(err.Error(), "does not look like the site repository") {
		t.Fatalf("err = %v, want the repo-root sanity-check message", err)
	}
}

func TestRunInvalidBandsErrors(t *testing.T) {
	err := run("leveling-bis", []string{"-repo-root", repoRootFixture, "-bands", "not-a-number"})
	if err == nil {
		t.Fatal("run with -bands not-a-number: want an error, got nil")
	}
}

func TestRunMissingActiveBuildFileErrors(t *testing.T) {
	// -build is left empty and repoRootFixture's own
	// web/src/data/active-build.json DOES exist (leveling.ReadActiveBuild
	// succeeds) - this test instead points -repo-root at a fixture
	// with specs.json but no active-build.json, to reach
	// leveling.ReadActiveBuild's own error path through run().
	dir := t.TempDir()
	if err := writeFile(t, filepath.Join(dir, "data", "curated", "specs.json"), `[]`); err != nil {
		t.Fatal(err)
	}
	err := run("leveling-bis", []string{"-repo-root", dir})
	if err == nil {
		t.Fatal("run with no active-build.json and no -build flag: want an error, got nil")
	}
}

func TestRunFlagParseErrorIsReturnedNotFatal(t *testing.T) {
	err := run("leveling-bis", []string{"-not-a-real-flag"})
	if err == nil {
		t.Fatal("run with an unknown flag: want an error, got nil")
	}
}
