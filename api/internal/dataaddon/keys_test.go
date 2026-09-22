// api/internal/dataaddon/keys_test.go
package dataaddon

import "testing"

func TestSlugLowercasesAndCollapsesWhitespaceToOneHyphen(t *testing.T) {
	cases := map[string]string{
		"Thoradin":       "thoradin",
		"Iron  Vanguard": "iron-vanguard", // a doubled space still collapses to one hyphen
		"O'Malley":       "o'malley",      // an apostrophe is not whitespace: untouched
		"Mörk":           "mörk",          // non-ASCII passes through unescaped
	}
	for in, want := range cases {
		if got := slug(in); got != want {
			t.Errorf("slug(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestCharacterKeyMatchesTheAddonReadersOwnFormat(t *testing.T) {
	// addon/ForeverSixty/Ratings.lua: characterKey(region, realm, name) =
	// slug(region) .. ":" .. slug(realm) .. ":" .. slug(name); this job's
	// realm position is the site's ruleset (see the plan's "Key-format
	// reconciliation" section).
	got := characterKey("US", "normal", "Thoradin")
	if want := "us:normal:thoradin"; got != want {
		t.Errorf("characterKey = %q, want %q", got, want)
	}
	got = characterKey("EU", "pvp", "Mörk")
	if want := "eu:pvp:mörk"; got != want {
		t.Errorf("characterKey = %q, want %q", got, want)
	}
}

func TestGuildKeyCollapsesTheGuildNamesSpaces(t *testing.T) {
	got := guildKey("US", "normal", "Iron Vanguard")
	if want := "us:normal:iron-vanguard"; got != want {
		t.Errorf("guildKey = %q, want %q", got, want)
	}
}
