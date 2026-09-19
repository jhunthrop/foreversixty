package main

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
)

// The planted log sim/measure's own tests read. This command is a thin
// shell over that package, so what is tested here is the shell: the
// flags, the two output shapes and the exit codes a caller branches on.
var planted = filepath.Join("..", "..", "measure", "testdata", "planted-v22.log")

// The actor the planted log is written around.
const actor = "Testwarrior-Beta"

func runArgs(t *testing.T, args ...string) (int, string, string) {
	t.Helper()
	var stdout, stderr bytes.Buffer
	code := run(args, &stdout, &stderr)
	return code, stdout.String(), stderr.String()
}

func TestATableIsPrinted(t *testing.T) {
	code, stdout, stderr := runArgs(t, "-log", planted, "-actor", actor)
	if code != exitOK {
		t.Fatalf("exit %d, stderr %q", code, stderr)
	}
	// The actor's own rows, not an empty frame: a report for a name no
	// line mentions still prints headings.
	if !strings.Contains(stdout, actor) {
		t.Errorf("the table does not mention %q:\n%s", actor, stdout)
	}
}

// -json is what the nightly validation job reads, so it has to be a
// document rather than a table with brackets.
func TestJSONIsADocument(t *testing.T) {
	code, stdout, stderr := runArgs(t, "-log", planted, "-actor", actor, "-json")
	if code != exitOK {
		t.Fatalf("exit %d, stderr %q", code, stderr)
	}
	var doc map[string]any
	if err := json.Unmarshal([]byte(stdout), &doc); err != nil {
		t.Fatalf("-json did not produce JSON: %v\n%s", err, stdout)
	}
	if len(doc) == 0 {
		t.Error("the report is empty")
	}
}

// The two exit codes a caller branches on: 2 is "you called it wrong",
// 1 is "it could not do the job". A script that retried the first
// forever is what the distinction prevents.
func TestExitCodes(t *testing.T) {
	cases := []struct {
		name string
		args []string
		want int
	}{
		{"no log", []string{"-actor", actor}, exitBadFlags},
		{"no actor", []string{"-log", planted}, exitBadFlags},
		{"neither", nil, exitBadFlags},
		{"an unknown flag", []string{"-nope"}, exitBadFlags},
		{"a log that is not there", []string{"-log", filepath.Join(t.TempDir(), "nope.log"), "-actor", actor}, exitFailed},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			code, stdout, stderr := runArgs(t, tc.args...)
			if code != tc.want {
				t.Errorf("exit %d, want %d (stderr %q)", code, tc.want, stderr)
			}
			if stdout != "" {
				t.Errorf("a failed run wrote to stdout: %q", stdout)
			}
			if stderr == "" {
				t.Error("a failed run said nothing on stderr")
			}
		})
	}
}
