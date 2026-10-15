// api/internal/dataaddon/write_test.go
package dataaddon

import (
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/rating"
)

// sampleData is a hand-built night: one scored character, one no-logs guild member, and
// their guild.
func sampleData() Data {
	asOf := time.Date(2026, 11, 2, 21, 30, 0, 0, time.UTC)
	return Data{
		Generated: time.Date(2026, 11, 4, 4, 0, 0, 0, time.UTC),
		Characters: map[string]characterRecord{
			"Obnoxious Yell-ClassicBetaPvE2": {
				Region: "us",
				Score: &rating.ScoreV2{
					Score: 1230, Tier: "rare", AsOf: asOf,
					Components: rating.ScoreComponents{Gear: 400, Experience: 380, Parses: 250, Attendance: 150, Discipline: 50},
					GearScore:  400, ItemLevel: 31, Attendance: 0.75, Spec: "warrior-arms", Fights: 4,
					Experience: rating.ScoreExperience{RaidKills: 3, Alts: 1},
				},
				BestParses: map[string]int{"Magmadar": 58, "Lucifron": 71},
				ItemLevel:  31, Spec: "warrior-arms",
				Profile: "/character/us/normal/obnoxious-yell",
			},
			"Quiet Fern-ClassicBetaPvE2": {Region: "us", Spec: "warrior-fury", Profile: "/character/us/normal/quiet-fern"},
		},
		Guilds: map[string]guildRecord{
			"OLYMPUS XXVII": {
				Region: "us", Standing: 1, Nights: 4, Roster: 3, Attendance: fraction(2.0 / 3),
				Members: []string{"Obnoxious Yell-ClassicBetaPvE2", "Quiet Fern-ClassicBetaPvE2"},
			},
		},
	}
}

func TestRenderWritesContractFivesLayout(t *testing.T) {
	files, err := Render(sampleData())
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 {
		t.Fatalf("files = %v, want exactly Data.lua", mapKeysOf(files))
	}
	want := `ForeverSixtyData = {
	format = 2,
	generated = "2026-11-04",
	characters = {
		["Obnoxious Yell-ClassicBetaPvE2"] = {
			score = 1230,
			tier = "rare",
			as_of = "2026-11-02",
			components = { gear = 400, experience = 380, parses = 250, attendance = 150, discipline = 50 },
			gear_score = 400,
			item_level = 31,
			spec = "warrior-arms",
			experience = { dungeons = 0, raid_kills = 3, alts = 1 },
			best_parses = { ["Lucifron"] = 71, ["Magmadar"] = 58 },
			attendance = 0.75,
			fights = 4,
			profile = "/character/us/normal/obnoxious-yell",
		},
		["Quiet Fern-ClassicBetaPvE2"] = {
			score = nil,
			spec = "warrior-fury",
			profile = "/character/us/normal/quiet-fern",
		},
	},
	guilds = {
		["OLYMPUS XXVII"] = {
			standing = 1,
			nights = 4,
			roster = 3,
			attendance = 0.67,
			members = { "Obnoxious Yell-ClassicBetaPvE2", "Quiet Fern-ClassicBetaPvE2" },
		},
	},
}
`
	if got := string(files["Data.lua"]); got != want {
		t.Errorf("Data.lua:\n%s\nwant:\n%s", got, want)
	}
}

func TestRenderOmitsWhatIsNotKnown(t *testing.T) {
	d := Data{
		Generated:  time.Date(2026, 11, 4, 0, 0, 0, 0, time.UTC),
		Characters: map[string]characterRecord{"Thoradin-ClassicBetaPvE2": {Region: "us", Profile: "/character/us/normal/thoradin"}},
		Guilds:     map[string]guildRecord{"Iron Vanguard": {Region: "us", Standing: 2, Roster: 1, Members: []string{}}},
	}
	files, err := Render(d)
	if err != nil {
		t.Fatal(err)
	}
	got := string(files["Data.lua"])
	if !strings.Contains(got, "\t\t[\"Thoradin-ClassicBetaPvE2\"] = {\n\t\t\tscore = nil,\n\t\t\tprofile = \"/character/us/normal/thoradin\",\n\t\t},\n") {
		t.Errorf("a no-logs record with nothing else known should carry score = nil and its profile only:\n%s", got)
	}
	if strings.Contains(got, "attendance") {
		t.Errorf("a guild with no scored member has no attendance to write:\n%s", got)
	}
	if !strings.Contains(got, "members = {},") {
		t.Errorf("a guild with no listed member writes an empty list:\n%s", got)
	}
}

func TestLuaStringEscapesOnlyWhatLuaNeeds(t *testing.T) {
	for in, want := range map[string]string{
		"Mörk":         `"Mörk"`,
		`O"Malley\`:    `"O\"Malley\\"`,
		"two\nlines":   `"two\nlines"`,
		"tab\t1":       `"tab\0091"`,
		"Lady Sunwick": `"Lady Sunwick"`,
	} {
		if got := luaString(in); got != want {
			t.Errorf("luaString(%q) = %s, want %s", in, got, want)
		}
	}
}

func TestLuaFractionIsAWholePercent(t *testing.T) {
	for v, want := range map[float64]string{0.75: "0.75", 2.0 / 3: "0.67", 1: "1", 0: "0", 0.9: "0.9"} {
		if got := luaFraction(v); got != want {
			t.Errorf("luaFraction(%v) = %s, want %s", v, got, want)
		}
	}
}

func mapKeysOf[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}
