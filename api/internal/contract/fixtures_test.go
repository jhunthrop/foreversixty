package contract

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Every response fixture is a well-formed envelope: {ok, data, error, request_id}, where a
// success carries data and a null error, and a failure carries a null data and a code. The
// fixtures of endpoints not built yet (app.*) are held to the same rule, so an app can rely on
// them before the handler lands. The static-file fixtures (cdn.*) are not envelopes.
func TestEveryResponseFixtureIsAnEnvelope(t *testing.T) {
	files, err := filepath.Glob(filepath.Join(fixtureDir(), "*.json"))
	if err != nil || len(files) == 0 {
		t.Fatalf("no fixtures found: %v", err)
	}
	for _, file := range files {
		name := strings.TrimSuffix(filepath.Base(file), ".json")
		if strings.HasPrefix(name, "cdn.") {
			continue
		}
		t.Run(name, func(t *testing.T) { assertEnvelope(t, Load(t, name), strings.Contains(name, "error")) })
	}
}

func assertEnvelope(t *testing.T, doc any, wantFailure bool) {
	t.Helper()
	env, ok := doc.(map[string]any)
	if !ok {
		t.Fatalf("a fixture is a JSON object, got %T", doc)
	}
	for _, key := range []string{"ok", "data", "error", "request_id"} {
		if _, present := env[key]; !present {
			t.Fatalf("the envelope has no %q", key)
		}
	}
	success, _ := env["ok"].(bool)
	if success == wantFailure {
		t.Fatalf("ok = %v, but the fixture's name says it is a %s", success, map[bool]string{true: "failure", false: "success"}[wantFailure])
	}
	if success && env["error"] != nil {
		t.Errorf("a success has an error: %v", env["error"])
	}
	if !success {
		failure, _ := env["error"].(map[string]any)
		if env["data"] != nil || failure == nil || failure["code"] == "" || failure["message"] == "" {
			t.Errorf("a failure has a null data, and an error with a code and a message: %v", env)
		}
	}
}

// The fixtures are valid JSON files that end in a newline, formatted as the generator writes
// them, so a diff of one shows only what changed.
func TestFixturesAreFormattedJSON(t *testing.T) {
	files, _ := filepath.Glob(filepath.Join(fixtureDir(), "*.json"))
	for _, file := range files {
		raw, err := os.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		if len(raw) == 0 || raw[len(raw)-1] != '\n' {
			t.Errorf("%s does not end in a newline", filepath.Base(file))
		}
	}
}
