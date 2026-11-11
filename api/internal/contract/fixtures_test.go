package contract

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/push"
)

// Every response fixture is a well-formed envelope: {ok, data, error, request_id}, where a
// success carries data and a null error, and a failure carries a null data and a code. The
// fixtures of endpoints not built yet (app.*) are held to the same rule, so an app can rely on
// them before the handler lands. The static-file fixtures (cdn.*) and the push payloads (push.*,
// and the live-surface pushes surface.*, pinned to their builders in api/internal/livesurface)
// are not envelopes.
func TestEveryResponseFixtureIsAnEnvelope(t *testing.T) {
	files, err := filepath.Glob(filepath.Join(fixtureDir(), "*.json"))
	if err != nil || len(files) == 0 {
		t.Fatalf("no fixtures found: %v", err)
	}
	for _, file := range files {
		name := strings.TrimSuffix(filepath.Base(file), ".json")
		if strings.HasPrefix(name, "cdn.") || strings.HasPrefix(name, "push.") || strings.HasPrefix(name, "surface.") {
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

// The push payloads (lane A2, api/internal/push) are what APNs and FCM hand the app. Until A2
// merges and pins them to its own builders, they are held to the fields app-v1.md promises.
func TestPushFixturesCarryThePayloadFields(t *testing.T) {
	files, _ := filepath.Glob(filepath.Join(fixtureDir(), "push.*.json"))
	if len(files) == 0 {
		t.Fatal("no push fixtures found")
	}
	for _, file := range files {
		name := strings.TrimSuffix(filepath.Base(file), ".json")
		t.Run(name, func(t *testing.T) {
			doc, _ := Load(t, name).(map[string]any)
			switch {
			case strings.HasSuffix(name, ".apns"):
				aps, _ := doc["aps"].(map[string]any)
				alert, _ := aps["alert"].(map[string]any)
				requireKeys(t, alert, "title", "body")
				requireKeys(t, aps, "thread-id")
				requireKeys(t, doc, "kind", "link", "data")
				if _, hasActions := doc["actions"]; hasActions && !pushCategories[fmt.Sprint(aps["category"])] {
					t.Errorf("a payload with actions names a button set the apps register, got %v", aps["category"])
				}
			case strings.HasSuffix(name, ".fcm"):
				assertDataOnlyFCM(t, doc)
			default:
				t.Fatalf("a push fixture ends in .apns or .fcm: %s", name)
			}
		})
	}
}

// pushCategories are the button sets a payload with actions can name (push.MessageOf).
var pushCategories = map[string]bool{push.RaidSignupCategory: true, push.GuildApplicantCategory: true, push.RaidShortCategory: true}

// assertDataOnlyFCM holds an FCM fixture to the data-only message Android is sent (since
// a8ad9c882): no notification block, so the app's own service draws every push with its group;
// everything the app shows or does is a string in data, the actions a JSON list beside the
// button set (category) the apps register.
func assertDataOnlyFCM(t *testing.T, doc map[string]any) {
	t.Helper()
	msg, _ := doc["message"].(map[string]any)
	requireKeys(t, msg, "token", "data", "android")
	if _, ok := msg["notification"]; ok {
		t.Error("an FCM message is data-only: it carries no notification block")
	}
	data, _ := msg["data"].(map[string]any)
	requireKeys(t, data, "kind", "group", "title", "body", "link", "thread", "collapse")
	for key, value := range data {
		if _, isString := value.(string); !isString {
			t.Errorf("FCM data values are strings; %q is %T", key, value)
		}
	}
	raw, hasActions := data["actions"].(string)
	if !hasActions {
		return
	}
	if !pushCategories[fmt.Sprint(data["category"])] {
		t.Errorf("a payload with actions names a button set the apps register, got %v", data["category"])
	}
	var actions []map[string]any
	if err := json.Unmarshal([]byte(raw), &actions); err != nil || len(actions) == 0 {
		t.Errorf("data.actions is a JSON list of actions: %v (%s)", err, raw)
	}
	for _, a := range actions {
		requireKeys(t, a, "id", "title", "method", "path")
	}
}

func requireKeys(t *testing.T, doc map[string]any, keys ...string) {
	t.Helper()
	for _, key := range keys {
		if _, ok := doc[key]; !ok {
			t.Errorf("missing %q in %v", key, doc)
		}
	}
}
