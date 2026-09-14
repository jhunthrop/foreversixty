package character

import (
	"strings"
	"testing"
)

func TestKeyIsRegionRulesetAndNameSlug(t *testing.T) {
	for _, tc := range []struct{ region, ruleset, name, want string }{
		{"us", "normal", "Baelgrim", "us/normal/baelgrim"},
		{"US", "Hardcore", "Baelgrim", "us/hardcore/baelgrim"},
		{"eu", "rp", "Lady Sunwick", "eu/rp/lady-sunwick"},
	} {
		if got := Key(tc.region, tc.ruleset, tc.name); got != tc.want {
			t.Errorf("Key(%q, %q, %q) = %q, want %q", tc.region, tc.ruleset, tc.name, got, tc.want)
		}
	}
}

func TestValidRulesetAndRegion(t *testing.T) {
	for _, r := range []string{"normal", "pvp", "rp", "hardcore"} {
		if !ValidRuleset(r) {
			t.Errorf("%q should be a ruleset", r)
		}
	}
	if ValidRuleset("nightslayer") {
		t.Error("a realm name is not a ruleset")
	}
	if !ValidRegion("eu") || ValidRegion("mars") {
		t.Error("regions are the five the contract names")
	}
}

func TestRulesetFromRealmTakesTheSegmentThenTheFallback(t *testing.T) {
	if got := RulesetFromRealm("Hardcore", "normal"); got != "hardcore" {
		t.Errorf("a segment that names a ruleset wins: %q", got)
	}
	if got := RulesetFromRealm("Nightslayer", "pvp"); got != "pvp" {
		t.Errorf("an unknown segment takes the report's ruleset: %q", got)
	}
	if got := RulesetFromRealm("Nightslayer", ""); got != RulesetNormal {
		t.Errorf("with no fallback the default is normal: %q", got)
	}
}

func TestKeyFromUnitSplitsTheLoggedName(t *testing.T) {
	if got := KeyFromUnit("us", "hardcore", "Baelgrim-Nightslayer"); got != "us/hardcore/baelgrim" {
		t.Errorf("KeyFromUnit = %q", got)
	}
	if got := KeyFromUnit("us", "pvp", "Baelgrim"); got != "us/pvp/baelgrim" {
		t.Errorf("a unit with no realm segment is all name: %q", got)
	}
}

func TestValidKeyAcceptsTheContractsShapeAndNothingElse(t *testing.T) {
	for _, c := range []struct {
		key  string
		want bool
	}{
		{"us/normal/baelgrim", true},
		{"eu/rp/lady-sunwick", true},
		{"kr/hardcore/" + "사실", true},
		{"us/normal/bael9", true},
		{"", false},
		{"us/normal", false},
		{"us/normal/", false},
		{"mars/normal/baelgrim", false},
		{"us/nightslayer/baelgrim", false},
		{"us/normal/Baelgrim", false},
		{"us/normal/bael grim", false}, // Slug would have hyphenated it
		{"us/normal/bael/grim", false},
		// The slug crosses into a Lua file the addon loads, so the
		// characters that would close a string or a long bracket there
		// are refused at the door.
		{`us/normal/bael"grim`, false},
		{"us/normal/bael'grim", false},
		{`us/normal/bael\grim`, false},
		{"us/normal/bael]]grim", false},
		{"us/normal/bael\ngrim", false},
		{"us/normal/bael\x00grim", false},
		{"us/normal/bael\u00a0grim", false}, // a non-breaking space is still a space
		{"us/normal/bael\xffgrim", false},
		{"us/normal/" + strings.Repeat("a", MaxSlugBytes), true},
		{"us/normal/" + strings.Repeat("a", MaxSlugBytes+1), false},
	} {
		if got := ValidKey(c.key); got != c.want {
			t.Errorf("ValidKey(%q) = %v, want %v", c.key, got, c.want)
		}
	}
}
