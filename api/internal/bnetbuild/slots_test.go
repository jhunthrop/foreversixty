// api/internal/bnetbuild/slots_test.go
package bnetbuild

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

// slotsArrayPattern extracts the quoted strings inside web/src/lib/planner/types.ts's
// `export const SLOTS = [...] as const;` block, so this test fails the moment that file's
// slot list changes without this package's copy changing with it.
var slotsArrayPattern = regexp.MustCompile(`export const SLOTS = \[([\s\S]*?)\] as const;`)

func TestSlotsMatchesThePlannerTypesFile(t *testing.T) {
	path := filepath.Join("..", "..", "..", "web", "src", "lib", "planner", "types.ts")
	src, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	m := slotsArrayPattern.FindSubmatch(src)
	if m == nil {
		t.Fatalf("could not find the SLOTS array in %s", path)
	}
	quoted := regexp.MustCompile(`'([a-z0-9_]+)'`).FindAllStringSubmatch(string(m[1]), -1)
	var want []string
	for _, q := range quoted {
		want = append(want, q[1])
	}
	if len(want) == 0 {
		t.Fatal("parsed zero slots out of types.ts; the regexp or the file format changed")
	}
	if len(SLOTS) != len(want) {
		t.Fatalf("bnetbuild.SLOTS has %d entries, types.ts has %d: %v vs %v", len(SLOTS), len(want), SLOTS, want)
	}
	for i, slot := range SLOTS {
		if slot != want[i] {
			t.Fatalf("SLOTS[%d] = %q, types.ts has %q at the same position", i, slot, want[i])
		}
	}
}
