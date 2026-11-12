package activity

import "testing"

func TestBandsAndWords(t *testing.T) {
	cases := map[int]int{1: 20, 19: 20, 20: 20, 29: 20, 34: 30, 59: 50, 60: 60}
	for level, want := range cases {
		if got := BandOf(level); got != want {
			t.Errorf("BandOf(%d) = %d, want %d", level, got, want)
		}
	}
	if BandWords(20) != "20 to 29" || BandWords(60) != "60" {
		t.Fatalf("band words: %q %q", BandWords(20), BandWords(60))
	}
	if LevelWords(13, 20) != "level 13" || LevelWords(20, 20) != "20 to 29" || LevelWords(0, 60) != "60" {
		t.Fatalf("level words: %q %q %q", LevelWords(13, 20), LevelWords(20, 20), LevelWords(0, 60))
	}
	if ValidBand(25) || !ValidBand(40) {
		t.Fatal("ValidBand")
	}
}

func TestSpecKeysAndNames(t *testing.T) {
	for _, c := range [][3]string{{"Hunter", "BeastMastery", "hunter-beast-mastery"}, {"Hunter", "Beast Mastery", "hunter-beast-mastery"}, {"Warrior", "Fury", "warrior-fury"}} {
		got, ok := SpecKey(c[0], c[1])
		if !ok || got != c[2] {
			t.Errorf("SpecKey(%q, %q) = %q %v", c[0], c[1], got, ok)
		}
	}
	if _, ok := SpecKey("Warrior", "Nope"); ok {
		t.Fatal("unknown spec matched")
	}
	if SpecName("warrior-fury") != "Fury Warrior" || SpecRole("priest-holy") != "healer" {
		t.Fatalf("names: %q %q", SpecName("warrior-fury"), SpecRole("priest-holy"))
	}
}
