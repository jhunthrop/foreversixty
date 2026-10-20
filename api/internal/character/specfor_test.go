package character

import "testing"

func TestSpecForResolutionOrder(t *testing.T) {
	cases := []struct {
		name         string
		class, fight string
		ranks        []int
		spec, source string
	}{
		{"a fight's spec wins over the talents", "warrior", "warrior-arms", []int{0, 12, 0}, "warrior-arms", "fight"},
		{"the talents name the dominant tree", "warrior", "", []int{0, 12, 0}, "warrior-fury", "talents"},
		{"a tie goes to the first tree", "warrior", "", []int{5, 5, 0}, "warrior-arms", "talents"},
		{"the last tree", "warrior", "", []int{0, 2, 9}, "warrior-protection", "talents"},
		{"no points spent is the class's first spec", "warrior", "", []int{0, 0, 0}, "warrior-arms", "default"},
		{"no tree ranks at all is the default too", "mage", "", nil, "mage-arcane", "default"},
		{"a class the data does not know has no spec", "gnome", "", []int{1, 0, 0}, "", ""},
	}
	for _, c := range cases {
		spec, source := SpecFor(c.class, c.fight, c.ranks)
		if spec != c.spec || source != c.source {
			t.Errorf("%s: SpecFor = %q, %q; want %q, %q", c.name, spec, source, c.spec, c.source)
		}
	}
}
