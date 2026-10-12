// api/internal/dataaddon/keys_test.go
package dataaddon

import "testing"

// The key is Name-<the realm with its spaces and hyphens removed>, the second form the
// addon's Score.forCharacter looks a character up by.
func TestCharacterKeyIsNameAndTheNormalisedRealm(t *testing.T) {
	for _, c := range []struct{ name, realm, want string }{
		{"Obnoxious Yell", "Classic Beta PvE 2", "Obnoxious Yell-ClassicBetaPvE2"},
		{"Mörk", "ClassicBetaPvP", "Mörk-ClassicBetaPvP"},
		{"Thoradin", "Classic-Beta PvP", "Thoradin-ClassicBetaPvP"},
		{"Thoradin", "", ""},
		{"", "Classic Beta PvP", ""},
	} {
		if got := characterKey(c.name, c.realm); got != c.want {
			t.Errorf("characterKey(%q, %q) = %q, want %q", c.name, c.realm, got, c.want)
		}
	}
}

// The realm comes from the Battle.net import, else the export's who=, else the newest rated
// row's unit; the name from the characters row, else that unit. Nothing is guessed.
func TestNameAndRealmTakeTheFirstSourceThatHasOne(t *testing.T) {
	const who = "FS1:1.60.1.70291:warrior:human:0/05320/0:|who=Quiet%20Fern:Classic%20Beta%20PvE%202"
	for _, c := range []struct {
		id                  identity
		export, name, realm string
	}{
		{identity{Name: "Obnoxious Yell", RealmName: "Classic Beta PvE 2", PlayerName: "Obnoxious Yell-ClassicBetaPvP"}, who,
			"Obnoxious Yell", "Classic Beta PvE 2"},
		{identity{Name: "Quiet Fern", PlayerName: "Quiet Fern-ClassicBetaPvP"}, who, "Quiet Fern", "Classic Beta PvE 2"},
		{identity{PlayerName: "Mörk-ClassicBetaPvP"}, "", "Mörk", "ClassicBetaPvP"},
		{identity{Name: "Thoradin"}, "FS1:1.60.1.70291:warrior:human:0/05320/0:", "Thoradin", ""},
		{identity{}, "", "", ""},
	} {
		name, realm := nameAndRealm(c.id, c.export)
		if name != c.name || realm != c.realm {
			t.Errorf("nameAndRealm(%+v) = %q, %q; want %q, %q", c.id, name, realm, c.name, c.realm)
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
