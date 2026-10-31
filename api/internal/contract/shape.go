// Package contract pins live responses to the JSON fixtures of docs/contracts/app-v1.md.
//
// A fixture is the contract's floor, not its ceiling: every field the fixture names must be
// in the response with the same JSON kind, and a field the response carries beyond the
// fixture is allowed (a client ignores what it does not know). A null on either side matches
// any kind, because nullability is documented per field and a fixture shows one state of it.
// Counts, ids and values are never compared. This package is test support only; nothing in
// the running service imports it.
package contract

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"sort"
	"strings"
	"testing"
)

// fixtureDir is docs/contracts/fixtures/app, found from this file's compiled-in path so a
// test in any package can use it.
func fixtureDir() string {
	_, self, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(self), "..", "..", "..", "docs", "contracts", "fixtures", "app")
}

// FixturePath is the path of the fixture called name (no extension).
func FixturePath(name string) string { return filepath.Join(fixtureDir(), name+".json") }

// Options adjust one comparison.
type Options struct {
	// Optional lists the paths of fields a response may omit (Go's omitempty). A path is dotted
	// from the root, with [] after an array: "data.characters[].level".
	Optional []string
	// Open lists the paths of objects whose keys are data, not schema (a map keyed by slot or
	// by stat): only the kind of their values is compared, against the fixture's first value.
	Open []string
}

// Load reads the fixture called name as generic JSON.
func Load(t testing.TB, name string) any {
	t.Helper()
	raw, err := os.ReadFile(FixturePath(name))
	if err != nil {
		t.Fatalf("contract: read fixture %s: %v", name, err)
	}
	v, err := decode(raw)
	if err != nil {
		t.Fatalf("contract: fixture %s is not JSON: %v", name, err)
	}
	return v
}

func decode(raw []byte) (any, error) {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// AssertFixture checks the response body got against the fixture called name.
func AssertFixture(t testing.TB, name string, got []byte, opts Options) {
	t.Helper()
	gotV, err := decode(got)
	if err != nil {
		t.Fatalf("contract: %s: response is not JSON: %v\n%s", name, err, got)
	}
	if diffs := Compare(Load(t, name), gotV, opts); len(diffs) > 0 {
		t.Errorf("contract: %s: the response breaks its fixture:\n  %s\nresponse:\n%s",
			name, strings.Join(diffs, "\n  "), got)
	}
}

// AssertValue is AssertFixture for a value already in memory.
func AssertValue(t testing.TB, name string, got any, opts Options) {
	t.Helper()
	raw, err := json.Marshal(got)
	if err != nil {
		t.Fatalf("contract: %s: encode: %v", name, err)
	}
	AssertFixture(t, name, raw, opts)
}

// Compare reports every way got breaks the fixture want; empty means it holds.
func Compare(want, got any, opts Options) []string {
	c := comparer{optional: set(opts.Optional), open: set(opts.Open)}
	c.walk("", want, got)
	return unique(c.diffs)
}

// unique sorts diffs and drops repeats: one broken field in a thousand array elements is one
// finding.
func unique(diffs []string) []string {
	sort.Strings(diffs)
	return slices.Compact(diffs)
}

func set(paths []string) map[string]bool {
	m := make(map[string]bool, len(paths))
	for _, p := range paths {
		m[p] = true
	}
	return m
}

type comparer struct {
	optional, open map[string]bool
	diffs          []string
}

func (c *comparer) add(path, format string, args ...any) {
	c.diffs = append(c.diffs, fmt.Sprintf("%s: %s", display(path), fmt.Sprintf(format, args...)))
}

func display(path string) string {
	if path == "" {
		return "(root)"
	}
	return path
}

func kind(v any) string {
	switch v.(type) {
	case nil:
		return "null"
	case bool:
		return "boolean"
	case float64:
		return "number"
	case string:
		return "string"
	case []any:
		return "array"
	default:
		return "object"
	}
}

func (c *comparer) walk(path string, want, got any) {
	wk, gk := kind(want), kind(got)
	if wk == "null" || gk == "null" {
		return
	}
	if wk != gk {
		c.add(path, "is a %s, the fixture has a %s", gk, wk)
		return
	}
	switch w := want.(type) {
	case map[string]any:
		c.walkObject(path, w, got.(map[string]any))
	case []any:
		c.walkArray(path, w, got.([]any))
	}
}

func (c *comparer) walkArray(path string, want, got []any) {
	if len(want) == 0 {
		return
	}
	for _, g := range got {
		c.walk(path+"[]", want[0], g)
	}
}

func (c *comparer) walkObject(path string, want, got map[string]any) {
	if c.open[path] {
		c.walkOpen(path, want, got)
		return
	}
	for _, key := range sortedKeys(want) {
		child := join(path, key)
		g, present := got[key]
		if !present {
			if !c.optional[child] {
				c.add(child, "is missing from the response")
			}
			continue
		}
		c.walk(child, want[key], g)
	}
}

func (c *comparer) walkOpen(path string, want, got map[string]any) {
	if len(want) == 0 {
		return
	}
	first := want[sortedKeys(want)[0]]
	for _, key := range sortedKeys(got) {
		c.walk(join(path, key), first, got[key])
	}
}

func join(path, key string) string {
	if path == "" {
		return key
	}
	return path + "." + key
}

func sortedKeys(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
