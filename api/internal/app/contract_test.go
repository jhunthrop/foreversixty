// The API's side of the app contract (docs/contracts/app-v1.md, app-v1-home-config.md): each fixture under
// docs/contracts/fixtures/app is the body its route answers for the reference world, and the
// handler's real response must have the fixture's shape: the same keys at every level and the
// same JSON type for every value both sides state (null matches any type, since a nullable
// field's fixture shows one case). An array's elements compare with the fixture's element at
// the same index, the last one past its end. `go test -run Contract
// -update` rewrites the fixtures from the handlers.
package app

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/rankings"
)

var update = flag.Bool("update", false, "rewrite the app contract fixtures from the handlers")

// fixtureRequestID is the request_id every fixture carries in place of a live one.
const fixtureRequestID = "req_fixture"

func TestContractFixtures(t *testing.T) {
	everyPick := map[string]int{"head": 226887, "neck": 19491, "shoulder": 272105, "back": 13340, "chest": 252483,
		"wrist": 226885, "hands": 272099, "waist": 20216, "legs": 237819, "feet": 226881, "finger1": 234202,
		"finger2": 272408, "trinket1": 249469, "trinket2": 249470, "main_hand": 250619, "ranged": 279273}
	cases := []struct {
		fixture string
		path    string
		edit    func(*homeWorld)
	}{
		{"character.home.json", "/v1/characters/us/normal/velaren/home", nil},
		// Lane F11's Home for Lightningbolltz, written by f11_fixtures.py: held to the same shape.
		{"character.home.yours.json", "/v1/characters/us/normal/velaren/home", nil},
		{"character.home.empty.json", "/v1/characters/us/normal/velaren/home", func(w *homeWorld) {
			w.chars.cards["velaren"] = velaren(everyPick)
		}},
		{"character.home.no-gear.json", "/v1/characters/us/normal/velaren/home", func(w *homeWorld) {
			card := w.chars.cards["velaren"]
			card.Gear = []rankings.GearSlot{}
			w.chars.cards["velaren"] = card
			delete(w.exports.rows, homeKey)
		}},
		{"app.config.json", "/v1/app/config", func(w *homeWorld) {
			w.svc.Settings.Features = []string{"raids", "sign_in"}
		}},
	}
	for _, c := range cases {
		t.Run(c.fixture, func(t *testing.T) {
			w := newHomeWorld(t)
			if c.edit != nil {
				c.edit(w)
			}
			rec := w.get(t, c.path)
			if rec.Code != http.StatusOK {
				t.Fatalf("%s = %d %s", c.path, rec.Code, rec.Body)
			}
			got := withFixtureID(t, rec.Body.Bytes())
			path := filepath.Join(homeContracts, c.fixture)
			if *update {
				writeFixture(t, path, got)
			}
			want, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			if diffs := shapeDiff("", decodeAny(t, want), decodeAny(t, got)); len(diffs) > 0 {
				t.Fatalf("%s differs from its fixture's shape:\n%s", c.path, strings.Join(diffs, "\n"))
			}
		})
	}
}

// The shape check itself fails on a missing key, an extra key and a changed type.
func TestShapeDiffCatchesDrift(t *testing.T) {
	want := map[string]any{"a": 1.0, "b": []any{map[string]any{"c": "x"}}, "d": nil}
	for name, got := range map[string]any{
		"a missing key": map[string]any{"b": []any{map[string]any{"c": "x"}}, "d": nil},
		"an extra key":  map[string]any{"a": 1.0, "b": []any{map[string]any{"c": "x"}}, "d": nil, "e": true},
		"a type change": map[string]any{"a": "1", "b": []any{map[string]any{"c": "x"}}, "d": nil},
		"inside arrays": map[string]any{"a": 1.0, "b": []any{map[string]any{"c": 2.0}}, "d": nil},
	} {
		if len(shapeDiff("", want, got)) == 0 {
			t.Errorf("%s passed", name)
		}
	}
	if d := shapeDiff("", want, map[string]any{"a": 2.0, "b": []any{}, "d": "now set"}); len(d) != 0 {
		t.Errorf("values, empty arrays and a set nullable must pass: %v", d)
	}
}

func withFixtureID(t *testing.T, body []byte) []byte {
	t.Helper()
	var env map[string]any
	if err := json.Unmarshal(body, &env); err != nil {
		t.Fatal(err)
	}
	env["request_id"] = fixtureRequestID
	out, err := json.MarshalIndent(env, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	return append(out, '\n')
}

func writeFixture(t *testing.T, path string, body []byte) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, body, 0o644); err != nil {
		t.Fatal(err)
	}
}

func decodeAny(t *testing.T, raw []byte) any {
	t.Helper()
	var v any
	dec := json.NewDecoder(bytes.NewReader(raw))
	if err := dec.Decode(&v); err != nil {
		t.Fatalf("decode: %v", err)
	}
	return v
}

// optionalKeys are fixture keys the spec lets a response leave out ("absent or null").
var optionalKeys = map[string]bool{"drop_chance": true}

// shapeDiff lists where got's shape departs from want's, by JSON path.
func shapeDiff(path string, want, got any) []string {
	if want == nil || got == nil {
		return nil
	}
	switch w := want.(type) {
	case map[string]any:
		g, ok := got.(map[string]any)
		if !ok {
			return []string{fmt.Sprintf("%s: object, got %T", path, got)}
		}
		return objectDiff(path, w, g)
	case []any:
		g, ok := got.([]any)
		if !ok {
			return []string{fmt.Sprintf("%s: array, got %T", path, got)}
		}
		var out []string
		for i := range g {
			if len(w) > 0 {
				out = append(out, shapeDiff(fmt.Sprintf("%s[%d]", path, i), w[min(i, len(w)-1)], g[i])...)
			}
		}
		return out
	}
	if fmt.Sprintf("%T", want) != fmt.Sprintf("%T", got) {
		return []string{fmt.Sprintf("%s: %T, got %T", path, want, got)}
	}
	return nil
}

func objectDiff(path string, want, got map[string]any) []string {
	keys := map[string]bool{}
	for k := range want {
		keys[k] = true
	}
	for k := range got {
		keys[k] = true
	}
	sorted := make([]string, 0, len(keys))
	for k := range keys {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	var out []string
	for _, k := range sorted {
		w, inWant := want[k]
		g, inGot := got[k]
		switch {
		case !inGot && optionalKeys[k]:
		case !inGot:
			out = append(out, fmt.Sprintf("%s.%s: missing", path, k))
		case !inWant:
			out = append(out, fmt.Sprintf("%s.%s: not in the fixture", path, k))
		default:
			out = append(out, shapeDiff(path+"."+k, w, g)...)
		}
	}
	return out
}
