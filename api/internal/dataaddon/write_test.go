// api/internal/dataaddon/write_test.go
package dataaddon

import (
	"strings"
	"testing"
	"time"
)

func TestRenderProducesOneSortedDeterministicFile(t *testing.T) {
	d := Data{
		Generated: time.Date(2026, 9, 22, 4, 0, 0, 0, time.UTC),
		Build:     "1.60.1.69893",
		Characters: map[string]characterRow{
			"us:normal:thoradin": {Rating: 81, Mean90: 81, Components: map[string]int{
				"output": 91, "survival": 72, "mechanics": 84, "utility": 62, "preparation": 94, "activity": 91,
			}, Fights: 2},
			"eu:pvp:mörk": {Rating: 88, Mean90: 88, Components: map[string]int{"output": 70, "mechanics": 70}, Fights: 1},
		},
		Guilds: map[string]guildRow{
			"us:normal:iron-vanguard": {
				Name: "Iron Vanguard", Progress: "2/3", Nights: 2, Roster: 2,
				Members: []string{"o'malley", "thoradin"},
			},
		},
	}
	files, err := Render(d)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 {
		t.Fatalf("files = %v, want exactly Data.lua", mapKeysOf(files))
	}
	got := string(files["Data.lua"])
	if !strings.Contains(got, `generated = "2026-09-22T04:00:00Z"`) {
		t.Error("missing generated field")
	}
	if !strings.Contains(got, `build = "1.60.1.69893"`) {
		t.Error("missing build field")
	}
	// eu:pvp:mörk sorts before us:normal:thoradin (byte order: 'e' < 'u').
	if strings.Index(got, "eu:pvp:mörk") > strings.Index(got, "us:normal:thoradin") {
		t.Error("characters are not sorted by key")
	}
	if !strings.Contains(got, `["us:normal:thoradin"] = { rating = 81, mean90 = 81, output = 91, survival = 72, mechanics = 84, utility = 62, preparation = 94, activity = 91, fights = 2 },`) {
		t.Errorf("thoradin row wrong, got:\n%s", got)
	}
	if !strings.Contains(got, `["eu:pvp:mörk"] = { rating = 88, mean90 = 88, output = 70, mechanics = 70, fights = 1 },`) {
		t.Errorf("mörk row wrong (should omit the four unscored components), got:\n%s", got)
	}
	if !strings.Contains(got, `["us:normal:iron-vanguard"] = { name = "Iron Vanguard", progress = "2/3", nights = 2, roster = 2, members = { "o'malley", "thoradin" } },`) {
		t.Errorf("guild row wrong, got:\n%s", got)
	}
}

func TestRenderOmitsProgressWhenEmpty(t *testing.T) {
	d := Data{
		Generated: time.Now(), Build: "x",
		Guilds: map[string]guildRow{"us:normal:new-guild": {Name: "New Guild", Nights: 0, Roster: 1, Members: []string{"a"}}},
	}
	files, err := Render(d)
	if err != nil {
		t.Fatal(err)
	}
	got := string(files["Data.lua"])
	if strings.Contains(got, "progress") {
		t.Errorf("progress should be omitted entirely, got:\n%s", got)
	}
}

func mapKeysOf[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}
