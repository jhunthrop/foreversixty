package character

import "testing"

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
