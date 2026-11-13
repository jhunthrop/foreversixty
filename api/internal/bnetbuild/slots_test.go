// api/internal/bnetbuild/slots_test.go
package bnetbuild

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

// The planner's SLOTS (web/src/lib/planner/types.ts) is derived since the one slot vocabulary
// (1ac516c28): web/src/lib/gear/slots.ts's FRAME_LEFT, FRAME_RIGHT and FRAME_BOTTOM in that
// order, less COSMETIC. This test derives the same list from those literal arrays, so it fails
// the moment that file's slot list changes without this package's copy changing with it.
var (
	frameArrayPattern = regexp.MustCompile(`export const (FRAME_LEFT|FRAME_RIGHT|FRAME_BOTTOM|COSMETIC)\b[^=]*= \[([\s\S]*?)\]`)
	quotedSlotPattern = regexp.MustCompile(`'([a-z0-9_]+)'`)
)

func plannerSlots(t *testing.T) []string {
	t.Helper()
	path := filepath.Join("..", "..", "..", "web", "src", "lib", "gear", "slots.ts")
	src, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	arrays := map[string][]string{}
	for _, m := range frameArrayPattern.FindAllSubmatch(src, -1) {
		for _, q := range quotedSlotPattern.FindAllSubmatch(m[2], -1) {
			arrays[string(m[1])] = append(arrays[string(m[1])], string(q[1]))
		}
	}
	for _, name := range []string{"FRAME_LEFT", "FRAME_RIGHT", "FRAME_BOTTOM", "COSMETIC"} {
		if len(arrays[name]) == 0 {
			t.Fatalf("parsed no %s out of %s; the regexp or the file format changed", name, path)
		}
	}
	cosmetic := map[string]bool{}
	for _, slot := range arrays["COSMETIC"] {
		cosmetic[slot] = true
	}
	var want []string
	for _, name := range []string{"FRAME_LEFT", "FRAME_RIGHT", "FRAME_BOTTOM"} {
		for _, slot := range arrays[name] {
			if !cosmetic[slot] {
				want = append(want, slot)
			}
		}
	}
	return want
}

func TestSlotsMatchesThePlannerTypesFile(t *testing.T) {
	want := plannerSlots(t)
	if len(SLOTS) != len(want) {
		t.Fatalf("bnetbuild.SLOTS has %d entries, the planner has %d: %v vs %v", len(SLOTS), len(want), SLOTS, want)
	}
	for i, slot := range SLOTS {
		if slot != want[i] {
			t.Fatalf("SLOTS[%d] = %q, the planner has %q at the same position", i, slot, want[i])
		}
	}
}
