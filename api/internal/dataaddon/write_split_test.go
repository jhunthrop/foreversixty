// api/internal/dataaddon/write_split_test.go
package dataaddon

import (
	"strings"
	"testing"
	"time"
)

func TestRenderWithLimitSplitsByRegionWhenOverTheLimit(t *testing.T) {
	d := Data{
		Generated: time.Date(2026, 9, 22, 4, 0, 0, 0, time.UTC),
		Build:     "1.60.1.69893",
		Characters: map[string]characterRow{
			"us:normal:thoradin": {Rating: 81, Mean90: 81, Components: map[string]int{"output": 91}, Fights: 2},
			"eu:pvp:mörk":        {Rating: 88, Mean90: 88, Components: map[string]int{"output": 70}, Fights: 1},
		},
		Guilds: map[string]guildRow{
			"us:normal:iron-vanguard": {Name: "Iron Vanguard", Nights: 1, Roster: 1, Members: []string{"a"}},
		},
	}
	// A limit of 1 byte guarantees the single-file render (never under 1
	// byte) always exceeds it, forcing the split path deterministically.
	files, err := renderWithLimit(d, 1)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := files["Data.lua"]; !ok {
		t.Fatal("Data.lua (the header/init file) is missing from a split render")
	}
	if strings.Contains(string(files["Data.lua"]), "thoradin") {
		t.Error("the header file should carry no rows, only format/generated/build and empty tables")
	}
	us, ok := files["Data-us.lua"]
	if !ok {
		t.Fatal("Data-us.lua is missing")
	}
	if !strings.Contains(string(us), "thoradin") || strings.Contains(string(us), "mörk") {
		t.Errorf("Data-us.lua should carry only us: rows, got:\n%s", us)
	}
	if !strings.Contains(string(us), "ForeverSixtyData.characters[key] = row") {
		t.Errorf("Data-us.lua should merge into the shared global, got:\n%s", us)
	}
	eu, ok := files["Data-eu.lua"]
	if !ok {
		t.Fatal("Data-eu.lua is missing")
	}
	if !strings.Contains(string(eu), "mörk") || strings.Contains(string(eu), "thoradin") {
		t.Errorf("Data-eu.lua should carry only eu: rows, got:\n%s", eu)
	}
}

func TestRegionsOfListsEveryRegionAKeyMentions(t *testing.T) {
	d := Data{
		Characters: map[string]characterRow{"us:normal:a": {}, "eu:pvp:b": {}},
		Guilds:     map[string]guildRow{"kr:hardcore:c": {}},
	}
	got := regionsOf(d)
	want := []string{"eu", "kr", "us"}
	if len(got) != len(want) {
		t.Fatalf("regions = %v, want %v", got, want)
	}
	for i, r := range want {
		if got[i] != r {
			t.Errorf("regions[%d] = %q, want %q", i, got[i], r)
		}
	}
}
