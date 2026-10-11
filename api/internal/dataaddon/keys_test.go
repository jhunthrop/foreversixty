// api/internal/dataaddon/keys_test.go
package dataaddon

import "testing"

func TestCharacterKeyIsNameRealm(t *testing.T) {
	for _, c := range []struct{ name, ruleset, want string }{
		{"Obnoxious Yell", "normal", "Obnoxious Yell-Normal"},
		{"Mörk", "pvp", "Mörk-Pvp"},
		{"Thoradin", "hardcore", "Thoradin-Hardcore"},
		{"Thoradin", "nightslayer", ""}, // not a ruleset: no realm, no key
		{"", "normal", ""},
	} {
		if got := characterKey(c.name, c.ruleset); got != c.want {
			t.Errorf("characterKey(%q, %q) = %q, want %q", c.name, c.ruleset, got, c.want)
		}
	}
}

func TestDisplayNameCutsTheLoggedRealmAndFallsBackToTheSlug(t *testing.T) {
	for _, c := range []struct{ stored, slug, want string }{
		{"Obnoxious Yell", "obnoxious-yell", "Obnoxious Yell"},
		{"Mörk-Pvp", "mörk", "Mörk"},
		{"", "thoradin", "thoradin"},
	} {
		if got := displayName(c.stored, c.slug); got != c.want {
			t.Errorf("displayName(%q, %q) = %q, want %q", c.stored, c.slug, got, c.want)
		}
	}
}

func TestProfilePathIsTheCharacterPage(t *testing.T) {
	if got := profilePath("us", "normal", "obnoxious-yell"); got != "/character/us/normal/obnoxious-yell" {
		t.Errorf("profilePath = %q", got)
	}
}

func TestSplitCharacterKey(t *testing.T) {
	region, ruleset, slug, ok := splitCharacterKey("eu/pvp/mörk")
	if !ok || region != "eu" || ruleset != "pvp" || slug != "mörk" {
		t.Errorf("splitCharacterKey = %q %q %q %v", region, ruleset, slug, ok)
	}
	if _, _, _, ok := splitCharacterKey("not-a-key"); ok {
		t.Error("a key without three parts should not split")
	}
}
