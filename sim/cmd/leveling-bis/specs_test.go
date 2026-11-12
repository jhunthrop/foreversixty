package main

import (
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestParseSpecsFlag(t *testing.T) {
	if got := parseSpecsFlag(""); got != nil {
		t.Errorf("empty: got %v, want nil", got)
	}
	if got, want := parseSpecsFlag(" rogue-combat, hunter-marksmanship,,"), []string{"rogue-combat", "hunter-marksmanship"}; !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestKeepWrittenKeepsWrittenOrder(t *testing.T) {
	written := []string{"druid-balance", "priest-holy", "warrior-fury"}
	got, err := keepWritten(written, []string{"warrior-fury", "druid-balance"})
	if err != nil {
		t.Fatal(err)
	}
	if want := []string{"druid-balance", "warrior-fury"}; !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestKeepWrittenRefusesAnUnwrittenSpec(t *testing.T) {
	_, err := keepWritten([]string{"warrior-fury"}, []string{"warior-fury"})
	if err == nil || !strings.Contains(err.Error(), "warior-fury") {
		t.Errorf("err = %v, want it to name the unknown spec", err)
	}
}

// Without -specs, -all ranks every written spec (the fixture root).
func TestSpecsToRankWithoutSpecsIsEveryWrittenSpec(t *testing.T) {
	got, err := specsToRank(repoRootFixture, nil)
	if err != nil {
		t.Fatal(err)
	}
	want, _ := writtenSpecs(repoRootFixture)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

// -specs without -all is a mistake, not a quiet single-spec run.
func TestSpecsFlagNeedsAll(t *testing.T) {
	err := run("leveling-bis", []string{"-repo-root", repoRootFixture, "-specs", "hunter-marksmanship"})
	if err == nil || !strings.Contains(err.Error(), "-all") {
		t.Errorf("err = %v, want it to ask for -all", err)
	}
}

// bis.yml's matrix (tools/bis-matrix.sh) and -all read the same rule for a
// written spec; this keeps the script's legs, flattened, equal to -all's list
// over the real repository.
func TestBISMatrixScriptAgreesWithWrittenSpecs(t *testing.T) {
	if _, err := exec.LookPath("jq"); err != nil {
		t.Skip("jq is not installed")
	}
	script := filepath.Join(publishedRepoRoot, "tools", "bis-matrix.sh")
	out, err := exec.Command(script, publishedRepoRoot).Output()
	if err != nil {
		t.Fatalf("%s: %v", script, err)
	}
	jq := exec.Command("jq", "-r", `[.include[].specs] | join(",")`)
	jq.Stdin = strings.NewReader(string(out))
	flat, err := jq.Output()
	if err != nil {
		t.Fatal(err)
	}
	got := parseSpecsFlag(strings.TrimSpace(string(flat)))
	want, err := writtenSpecs(publishedRepoRoot)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("matrix legs flattened = %v\nwritten specs = %v", got, want)
	}
}
