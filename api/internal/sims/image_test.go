package sims

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/runner"
)

// TestTheImageCarriesTheEngineWhereTheRunnerLooksForIt is a guard on
// two files that have to agree and are edited months apart: the
// Dockerfile puts the binary somewhere, and runner.DefaultBinary is
// where the jobs look for it.
func TestTheImageCarriesTheEngineWhereTheRunnerLooksForIt(t *testing.T) {
	b, err := os.ReadFile(filepath.Join("..", "..", "Dockerfile"))
	if err != nil {
		t.Fatal(err)
	}
	docker := string(b)
	if !strings.Contains(docker, "COPY --from=engine /out/forever-sim "+runner.DefaultBinary) {
		t.Errorf("the Dockerfile does not copy the engine to %q, which is where the jobs look",
			runner.DefaultBinary)
	}
	// The stage builds from the pin, not from a floating branch: the
	// sha is read out of sim/enginever/version.go inside the build.
	if !strings.Contains(docker, "enginever/version.go") {
		t.Error("the engine build stage does not read the pin from sim/enginever/version.go")
	}
	if !strings.Contains(docker, "./cmd/forever-sim") {
		t.Error("the engine build stage does not build sim/cmd/forever-sim")
	}
}
