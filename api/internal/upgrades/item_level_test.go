package upgrades

import "testing"

func TestEquippedItemLevelIsTheRoundedMean(t *testing.T) {
	cases := []struct {
		name   string
		levels []*int
		want   int
		ok     bool
	}{
		{"rounds down under a half", []*int{ptr(60), ptr(59), ptr(59), ptr(60), ptr(59)}, 59, true}, // 59.4
		{"rounds a half up", []*int{ptr(60), ptr(61)}, 61, true},                                    // 60.5
		{"nothing equipped", nil, 0, false},
		{"an unknown item", []*int{ptr(60), nil}, 0, false},
	}
	for _, c := range cases {
		got, ok := EquippedItemLevel(c.levels)
		if got != c.want || ok != c.ok {
			t.Errorf("%s: = %d %v, want %d %v", c.name, got, ok, c.want, c.ok)
		}
	}
}
