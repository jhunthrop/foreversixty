package sims

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/enginever"
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

// TestTheEngineBinaryReportsThePin is the runtime half of the guard
// above: the Dockerfile text can name the right files and still ship a
// binary built from something else, e.g. a stale build-cache layer
// that serves an old sim/ tree at the right paths. It only means
// anything inside the image - runner.DefaultBinary does not exist on a
// developer's machine or in `go test`'s own sandbox, including CI's
// own "api" job, which never builds the image - so it skips cleanly
// whenever the binary is absent and only asserts once there is one to
// check, which is what `.github/workflows/api.yml`'s deploy job does
// right after `docker build`.
func TestTheEngineBinaryReportsThePin(t *testing.T) {
	bin := runner.DefaultBinary
	if _, err := os.Stat(bin); err != nil {
		t.Skipf("no engine binary at %s (expected outside the built image): %v", bin, err)
	}
	out, err := exec.Command(bin, "-version").CombinedOutput()
	if err != nil {
		t.Fatalf("running %s -version: %v (%s)", bin, err, out)
	}
	got := strings.TrimSpace(string(out))
	if got != enginever.Version {
		t.Errorf("the image's engine binary reports version %q; sim/enginever.Version (the pin) is %q",
			got, enginever.Version)
	}
}
