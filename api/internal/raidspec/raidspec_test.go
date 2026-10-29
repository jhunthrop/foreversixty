package raidspec

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"
	"unicode/utf8"
)

// activeIconBuild is the build folder the emoji are drawn from today (§P.9.1).
const activeIconBuild = "1.60.1.70658"

var emojiNameRE = regexp.MustCompile(`^fs_[a-z]+(_[a-z]+)+$`)

func TestTheEmojiAreTheTwentyEightSpecsJSONKeys(t *testing.T) {
	all := All()
	if len(all) != 28 {
		t.Fatalf("%d specs", len(all))
	}
	seen := map[string]bool{}
	for _, s := range all {
		name := s.Emoji()
		if !emojiNameRE.MatchString(name) || utf8.RuneCountInString(name) > 32 || seen[name] {
			t.Errorf("%s: emoji name %q", s.Key, name)
		}
		seen[name] = true
	}
	if !seen["fs_hunter_beast_mastery"] || !seen["fs_druid_feral_bear"] {
		t.Errorf("names %v", seen)
	}
}

func TestFeralAndFeralBearHaveTheirOwnIcons(t *testing.T) {
	cat, _ := ByKey("druid-feral")
	bear, _ := ByKey("druid-feral-bear")
	if cat.Icon == bear.Icon || cat.Icon != "ability_druid_catformattack" {
		t.Fatalf("cat %q bear %q", cat.Icon, bear.Icon)
	}
}

// Every icon exists in the active build's folder. The folder is generated data that a fixture
// sync (FOREVER_DATA=fixture) replaces with a 200-file subset, so the check runs only when
// RAIDSPEC_ICONS=check names a checkout with the real build synced; CI's API job has no
// web data at all and skips.
func TestEveryIconExistsInTheActiveBuild(t *testing.T) {
	if os.Getenv("RAIDSPEC_ICONS") != "check" {
		t.Skip("set RAIDSPEC_ICONS=check on a checkout with the real build's web data synced")
	}
	dir := filepath.Join("..", "..", "..", "web", "public", "data", activeIconBuild, "icons")
	if _, err := os.Stat(dir); err != nil {
		t.Skipf("no icon folder at %s", dir)
	}
	for _, s := range All() {
		if _, err := os.Stat(filepath.Join(dir, s.Icon+".webp")); err != nil {
			t.Errorf("%s: %v", s.Key, err)
		}
	}
}

func TestResolveReadsLoggedNamesWithTheRole(t *testing.T) {
	cases := []struct{ class, name, role, want string }{
		{"paladin", "Holy", "healer", "paladin-holy"},
		{"priest", "holy", "healer", "priest-holy"},
		{"druid", "Feral", "tank", "druid-feral-bear"},
		{"druid", "Feral Combat", "dps", "druid-feral"},
		{"druid", "Feral Bear", "tank", "druid-feral-bear"},
		{"hunter", "Beast Mastery", "dps", "hunter-beast-mastery"},
		{"mage", "Holy", "healer", ""},
		{"mage", "", "dps", ""},
	}
	for _, c := range cases {
		got, ok := Resolve(c.class, c.name, c.role)
		if got.Key != c.want || ok != (c.want != "") {
			t.Errorf("Resolve(%q,%q,%q) = %q %v", c.class, c.name, c.role, got.Key, ok)
		}
	}
}

func TestMeleeFollowsTheSpecThenTheClass(t *testing.T) {
	for key, melee := range map[string]bool{"hunter-survival": true, "hunter-marksmanship": false, "druid-feral": true, "druid-balance": false, "warrior-fury": true} {
		s, _ := ByKey(key)
		if s.Melee() != melee {
			t.Errorf("%s melee %v", key, s.Melee())
		}
	}
	if !MeleeClass("rogue") || MeleeClass("mage") {
		t.Error("class fallback")
	}
	if ClassRank("druid") != 0 || ClassRank("") != len(ClassOrder) {
		t.Error("class rank")
	}
}
