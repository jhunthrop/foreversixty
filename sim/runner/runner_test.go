package runner

import (
	"bufio"
	"context"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// stubBinary writes an executable shell script standing in for
// forever-sim and returns its path.
func stubBinary(t *testing.T, script string) string {
	t.Helper()
	if runtime.GOOS == "windows" {
		t.Skip("the stub binary is a shell script")
	}
	bin := filepath.Join(t.TempDir(), "forever-sim")
	if err := os.WriteFile(bin, []byte("#!/bin/sh\n"+script), 0o755); err != nil {
		t.Fatal(err)
	}
	return bin
}

func aRequest() api.SimRequest {
	return api.SimRequest{
		EngineVersion: "pinned", Spec: "warrior-fury", Iterations: 3000,
		Source:    api.CharacterSource{Kind: api.SourceManual},
		Character: api.CharacterSpec{Name: "T", Race: "orc", Class: "warrior", Level: 60},
		Encounter: api.DefaultEncounter(),
	}
}

func TestNativeStreamsProgressOffStderrAndReadsTheResultOffStdout(t *testing.T) {
	// The real binary reads the request from stdin, writes progress
	// lines to stderr and one SimResult to stdout.
	bin := stubBinary(t, `
cat > /dev/null
echo '{"completed":1000,"total":3000,"dps":990}' >&2
echo '{"completed":2000,"total":3000,"dps":1020}' >&2
printf '%s' '{"engine_version":"pinned","iterations_run":3000,"duration_ms":2470,"dps":{"mean":1042.5,"error":1.6}}'
`)

	var ticks []int
	res, err := (&Native{Binary: bin}).Run(context.Background(), aRequest(),
		func(done int, _ float64) { ticks = append(ticks, done) })
	if err != nil {
		t.Fatal(err)
	}
	if len(ticks) != 2 || ticks[0] != 1000 || ticks[1] != 2000 {
		t.Fatalf("ticks %v", ticks)
	}
	if res.DPS.Mean != 1042.5 || res.IterationsRun != 3000 {
		t.Fatalf("result %+v", res)
	}
	// The runner owns the lane and puts the request back on the
	// result, so a caller never has to.
	if res.Lane != api.LaneServer || res.Request.Spec != "warrior-fury" {
		t.Fatalf("lane %q request %+v", res.Lane, res.Request)
	}
}

func TestNativeSendsTheRequestOnStdin(t *testing.T) {
	// The script echoes back a result whose error field carries what
	// it was given, so the test can see the request really arrived.
	bin := stubBinary(t, `
req=$(cat)
case "$req" in
  *warrior-fury*) printf '%s' '{"dps":{"mean":1},"iterations_run":1}' ;;
  *) printf '%s' '{"dps":{"mean":0},"iterations_run":0,"error":"no spec on stdin"}' ;;
esac
`)
	res, err := (&Native{Binary: bin}).Run(context.Background(), aRequest(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != "" {
		t.Fatalf("the binary did not receive the request: %q", res.Error)
	}
}

func TestNativeIgnoresAChattyStderrLine(t *testing.T) {
	bin := stubBinary(t, `
cat > /dev/null
echo 'loading the database...' >&2
printf '%s' '{"dps":{"mean":7},"iterations_run":10}'
`)
	res, err := (&Native{Binary: bin}).Run(context.Background(), aRequest(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.DPS.Mean != 7 {
		t.Fatalf("mean %v", res.DPS.Mean)
	}
}

func TestNativeFailsWhenTheBinaryWritesNoResult(t *testing.T) {
	bin := stubBinary(t, "cat > /dev/null\necho 'the database is missing' >&2\n")
	_, err := (&Native{Binary: bin}).Run(context.Background(), aRequest(), nil)
	if err == nil {
		t.Fatal("expected an error when nothing came back")
	}
	if !strings.Contains(err.Error(), "the database is missing") {
		t.Fatalf("the error should carry what the binary said: %v", err)
	}
}

func TestNativeSaysWhichFailureItWas(t *testing.T) {
	engineFailure := stubBinary(t, "cat > /dev/null\necho boom >&2\nexit 1\n")
	if _, err := (&Native{Binary: engineFailure}).Run(context.Background(), aRequest(), nil); err == nil {
		t.Fatal("expected an error on exit 1")
	} else if errorsIsBadInput(err) {
		t.Fatalf("exit 1 is an engine failure, not bad input: %v", err)
	}

	badInput := stubBinary(t, "cat > /dev/null\necho 'not a SimRequest' >&2\nexit 2\n")
	if _, err := (&Native{Binary: badInput}).Run(context.Background(), aRequest(), nil); err == nil {
		t.Fatal("expected an error on exit 2")
	} else if !errorsIsBadInput(err) {
		t.Fatalf("exit 2 is bad input: %v", err)
	}
}

func TestNativeReturnsTheContextErrorWhenCancelled(t *testing.T) {
	// The stub outlives the deadline, so the context kills it. The
	// caller must be able to tell its own deadline apart from an
	// engine crash, not just see "signal: killed".
	bin := stubBinary(t, "cat > /dev/null\nsleep 5\n")
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	_, err := (&Native{Binary: bin}).Run(ctx, aRequest(), nil)
	if err == nil {
		t.Fatal("expected an error when the context is cancelled")
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected the error to wrap context.DeadlineExceeded: %v", err)
	}
}

func TestNativeReturnsThePartialResultAndErrAbortedOnExit130(t *testing.T) {
	// forever-sim writes a complete-shaped SimResult with aborted set
	// and then exits 130. The caller must get both: the partial result
	// (so what completed is not lost) and a typed, errors.Is-able
	// error (so the job never stores it as a finished run).
	bin := stubBinary(t, `
cat > /dev/null
echo '{"completed":1500,"total":3000,"dps":700}' >&2
printf '%s' '{"engine_version":"pinned","iterations_run":1500,"duration_ms":900,"aborted":true,"dps":{"mean":700}}'
exit 130
`)
	var ticks []int
	res, err := (&Native{Binary: bin}).Run(context.Background(), aRequest(),
		func(done int, _ float64) { ticks = append(ticks, done) })
	if !errors.Is(err, ErrAborted) {
		t.Fatalf("err = %v, want ErrAborted", err)
	}
	if !res.Aborted || res.IterationsRun != 1500 {
		t.Fatalf("result %+v, want the partial result preserved with Aborted set", res)
	}
	// The runner still owns the lane and the request, the same as a
	// finished run, so a caller need not special-case an abort to
	// store it.
	if res.Lane != api.LaneServer || res.Request.Spec != "warrior-fury" {
		t.Fatalf("lane %q request %+v", res.Lane, res.Request)
	}
	if len(ticks) != 1 || ticks[0] != 1500 {
		t.Fatalf("ticks %v", ticks)
	}
}

func TestNativeFallsBackToTheGenericErrorWhenAbortWroteNoResult(t *testing.T) {
	// A crash before anything reached stdout still exits 130 (the
	// signal handler ran) but there is no partial result to decode;
	// this must not claim ErrAborted over nothing.
	bin := stubBinary(t, "cat > /dev/null\necho 'panic: nil pointer' >&2\nexit 130\n")
	res, err := (&Native{Binary: bin}).Run(context.Background(), aRequest(), nil)
	if err == nil {
		t.Fatal("expected an error")
	}
	if errors.Is(err, ErrAborted) {
		t.Fatalf("no result was decoded, so this must not be ErrAborted: %v", err)
	}
	if res.Aborted {
		t.Fatalf("expected the zero result: %+v", res)
	}
}

func TestNativeSurvivesAStderrLineOver64KBAndKeepsReadingProgress(t *testing.T) {
	// forever-sim wraps the engine's own log output as one
	// {"log":"..."} line on the same stream as progress ticks. Before
	// the buffer was enlarged, one line past bufio.Scanner's 64KB
	// default silently killed the goroutine, and the progress bar
	// froze with nothing to say why. A 100KB line must not stop later
	// ticks from arriving.
	bin := stubBinary(t, `
cat > /dev/null
long=$(head -c 100000 /dev/zero | tr '\0' 'x')
printf '{"log":"%s"}\n' "$long" >&2
echo '{"completed":2200,"total":3000,"dps":1010}' >&2
printf '%s' '{"dps":{"mean":42},"iterations_run":5}'
`)
	var ticks []int
	res, err := (&Native{Binary: bin}).Run(context.Background(), aRequest(),
		func(done int, _ float64) { ticks = append(ticks, done) })
	if err != nil {
		t.Fatal(err)
	}
	if len(ticks) != 1 || ticks[0] != 2200 {
		t.Fatalf("ticks %v: the progress tick after a >64KB stderr line was dropped", ticks)
	}
	if res.DPS.Mean != 42 {
		t.Fatalf("result %+v", res)
	}
}

func TestNativeSurfacesAScanErrorWhenTheProcessAlsoFails(t *testing.T) {
	// A line past even the enlarged 1MB cap is still possible; when it
	// coincides with the process itself failing, the caller should be
	// told reading stderr also broke, not just left silently short.
	bin := stubBinary(t, `
cat > /dev/null
head -c 1100000 /dev/zero | tr '\0' 'x' >&2
echo >&2
exit 1
`)
	_, err := (&Native{Binary: bin}).Run(context.Background(), aRequest(), nil)
	if err == nil {
		t.Fatal("expected an error")
	}
	if !errors.Is(err, bufio.ErrTooLong) {
		t.Fatalf("expected the scan error to be surfaced: %v", err)
	}
}

func TestNativeFailsWhenThereIsNoBinary(t *testing.T) {
	if _, err := (&Native{Binary: "/nowhere/forever-sim"}).
		Run(context.Background(), aRequest(), nil); err == nil {
		t.Fatal("expected an error with no binary")
	}
}

func TestTheFixtureAnswersWithARealSummary(t *testing.T) {
	f := &Fixture{}
	var last int
	req := aRequest()
	res, err := f.Run(context.Background(), req, func(done int, _ float64) { last = done })
	if err != nil {
		t.Fatal(err)
	}
	if last != req.Iterations {
		t.Fatalf("the last tick was %d, want the whole run", last)
	}
	if res.DPS.Mean == 0 || res.DPS.Error == 0 {
		t.Fatalf("the fixture has no distribution: %+v", res.DPS)
	}
	if len(res.Summary.DamageDone) != 1 {
		t.Fatalf("%d damage rows", len(res.Summary.DamageDone))
	}
	abilities := res.Summary.DamageDone[0].Abilities
	if len(abilities) < 2 || abilities[1].Name != "Bloodthirst" {
		t.Fatalf("abilities %+v", abilities)
	}
	if len(res.Summary.Auras) == 0 || res.Summary.Auras[0].Name != "Flurry" {
		t.Fatalf("auras %+v", res.Summary.Auras)
	}
	if len(res.Summary.Casts) == 0 || len(res.Summary.Roster) != 1 {
		t.Fatalf("casts %d roster %d", len(res.Summary.Casts), len(res.Summary.Roster))
	}
	if res.Lane != api.LaneServer || res.IterationsRun != req.Iterations {
		t.Fatalf("lane %q iterations %d", res.Lane, res.IterationsRun)
	}
	if res.EngineVersion != req.EngineVersion {
		t.Fatalf("engine version %q", res.EngineVersion)
	}
	if len(f.Asked()) != 1 || f.Asked()[0].Spec != "warrior-fury" {
		t.Fatalf("the request was not recorded: %+v", f.Asked())
	}
}

func TestTheFixtureCanBePutOnAKnownNumberOrMadeToFail(t *testing.T) {
	f := &Fixture{Mean: 1200}
	res, err := f.Run(context.Background(), aRequest(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.DPS.Mean != 1200 {
		t.Fatalf("mean %v, want the override", res.DPS.Mean)
	}
	// Two runs do not share a summary: a caller that edits one must
	// not change the next.
	res.Summary.DamageDone[0].Total = 1
	again, err := f.Run(context.Background(), aRequest(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if again.Summary.DamageDone[0].Total == 1 {
		t.Fatal("the fixture handed out a shared summary")
	}

	f = &Fixture{Err: context.DeadlineExceeded}
	if _, err := f.Run(context.Background(), aRequest(), nil); err == nil {
		t.Fatal("expected the configured error")
	}
}

func TestTheFixtureCanSimulateAnAbort(t *testing.T) {
	f := &Fixture{Aborted: true}
	req := aRequest()
	var last int
	res, err := f.Run(context.Background(), req, func(done int, _ float64) { last = done })
	if !errors.Is(err, ErrAborted) {
		t.Fatalf("err = %v, want ErrAborted", err)
	}
	if !res.Aborted || res.IterationsRun != req.Iterations/2 {
		t.Fatalf("result %+v", res)
	}
	if last != req.Iterations/2 {
		t.Fatalf("last tick %d, want %d", last, req.Iterations/2)
	}
}

// errorsIsBadInput is the test's own reader of the sentinel, so the
// assertion reads as one line above.
func errorsIsBadInput(err error) bool { return err != nil && isBadInput(err) }
