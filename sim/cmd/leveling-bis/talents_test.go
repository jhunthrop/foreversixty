package main

import (
	"reflect"
	"testing"
)

func TestTruncateTalentsDropsFromLastTreeLastSlotFirst(t *testing.T) {
	// The full-build fixture from marksmanship.md's build: frontmatter,
	// which sums to 51 (60-9).
	full := []string{"5522", "35305500115003", "51"}

	t.Run("budget above the total is a no-op", func(t *testing.T) {
		got := truncateTalents(full, 100)
		if !reflect.DeepEqual(got, full) {
			t.Fatalf("truncateTalents(full, 100) = %v, want unchanged %v", got, full)
		}
	})

	t.Run("budget equal to the total is a no-op", func(t *testing.T) {
		got := truncateTalents(full, 51)
		if !reflect.DeepEqual(got, full) {
			t.Fatalf("truncateTalents(full, 51) = %v, want unchanged %v", got, full)
		}
	})

	t.Run("dropping to zero clears every tree", func(t *testing.T) {
		got := truncateTalents(full, 0)
		want := []string{"", "", ""}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("truncateTalents(full, 0) = %v, want %v", got, want)
		}
		if pointTotal(got) != 0 {
			t.Fatalf("point total = %d, want 0", pointTotal(got))
		}
	})

	t.Run("a budget under the total drops from the last tree's last slot first", func(t *testing.T) {
		// The third tree ("51") has 6 points, digits [5,1]. Dropping 3
		// takes the last slot's 1 point, then 2 of the first slot's 5,
		// leaving "3" ("30" with the trailing zero trimmed) - the
		// second tree untouched because the third still had points to
		// give.
		got := truncateTalents(full, 48)
		if pointTotal(got) != 48 {
			t.Fatalf("point total = %d, want 48", pointTotal(got))
		}
		if got[2] != "3" {
			t.Fatalf("third tree = %q, want %q", got[2], "3")
		}
		if got[1] != full[1] {
			t.Fatalf("second tree = %q, want untouched %q while the third still had points to give", got[1], full[1])
		}
		if got[0] != full[0] {
			t.Fatalf("first tree = %q, want untouched %q while later trees still had points to give", got[0], full[0])
		}
	})

	t.Run("a budget that spans two trees empties the last before touching the middle", func(t *testing.T) {
		// Third tree has 6 points; asking for 51-10=41 must remove 10:
		// all 6 from tree 3, then 4 from the end of tree 2.
		got := truncateTalents(full, 41)
		if pointTotal(got) != 41 {
			t.Fatalf("point total = %d, want 41", pointTotal(got))
		}
		if got[2] != "" {
			t.Fatalf("third tree = %q, want empty", got[2])
		}
		if got[0] != full[0] {
			t.Fatalf("first tree = %q, want untouched %q", got[0], full[0])
		}
		if pointTotal([]string{got[1]}) != pointTotal([]string{full[1]})-4 {
			t.Fatalf("second tree lost %d points, want 4", pointTotal([]string{full[1]})-pointTotal([]string{got[1]}))
		}
	})

	t.Run("does not mutate its input", func(t *testing.T) {
		before := append([]string(nil), full...)
		truncateTalents(full, 10)
		if !reflect.DeepEqual(full, before) {
			t.Fatalf("truncateTalents mutated its input: %v -> %v", before, full)
		}
	})
}

func pointTotal(trees []string) int {
	total := 0
	for _, s := range trees {
		for _, c := range s {
			total += int(c - '0')
		}
	}
	return total
}

func TestTalentString(t *testing.T) {
	got := talentString([]string{"5522", "35305500115003", "51"})
	want := "5522-35305500115003-51"
	if got != want {
		t.Fatalf("talentString = %q, want %q", got, want)
	}
}

func TestLevelBudget(t *testing.T) {
	cases := map[int]int{20: 11, 30: 21, 40: 31, 60: 51}
	for level, want := range cases {
		if got := levelBudget(level); got != want {
			t.Errorf("levelBudget(%d) = %d, want %d", level, got, want)
		}
	}
}
