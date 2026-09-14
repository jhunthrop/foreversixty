package character

import "testing"

func TestTheKeyIsRegionRulesetAndTheNameSlug(t *testing.T) {
	for _, tc := range []struct {
		in   Character
		want string
	}{
		{Character{"US", "Hardcore", "Morrowlyn"}, "us/hardcore/morrowlyn"},
		{Character{"eu", "pvp", "Grim Batol"}, "eu/pvp/grim-batol"},
		{Character{"us", "normal", " Thalgrit "}, "us/normal/thalgrit"},
	} {
		if got := tc.in.Key(); got != tc.want {
			t.Errorf("%+v.Key() = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestValidNeedsAllThreeSegments(t *testing.T) {
	if !(Character{"us", "normal", "Morrowlyn"}).Valid() {
		t.Error("a filled-in character is not valid")
	}
	for _, c := range []Character{
		{"", "normal", "Morrowlyn"},
		{"us", "", "Morrowlyn"},
		{"us", "normal", "  "},
	} {
		if c.Valid() {
			t.Errorf("%+v is valid", c)
		}
	}
}
