package discordbot

import (
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// Guild round 3, items 9 and 10: the bot's night lines say what the site says.
func TestNightLineWithNobodySignedUp(t *testing.T) {
	cases := map[string]guilds.NightView{
		"5 came, nobody signed up beforehand · 1 kill": {Joined: 5, Kills: 1},
		"Nobody signed up beforehand":                  {},
		"4 of 5 came · 1 joined · 1 kill":              {Came: 4, Signed: 5, Joined: 1, Kills: 1},
	}
	for want, n := range cases {
		if got := nightLine(n); got != want {
			t.Errorf("nightLine(%+v) = %q, want %q", n, got, want)
		}
	}
}

func TestConsumablesLineWhenNobodyShared(t *testing.T) {
	unknown := guilds.ReportConsumablesView{Rows: []guilds.ConsumableRow{{Before: "unknown"}, {Before: "unknown"}}}
	if got := consumablesLine(unknown); got != "Nobody shared their bags before the raid" {
		t.Errorf("all not shared = %q", got)
	}
	mixed := guilds.ReportConsumablesView{Rows: []guilds.ConsumableRow{{Before: "unknown"}, {Before: "stocked"}}}
	if got := consumablesLine(mixed); got != "0 short before the raid" {
		t.Errorf("one shared = %q", got)
	}
}

func TestLoggedNightLineCountsTrashApart(t *testing.T) {
	ans := guilds.LoggedNightAnswer{Present: 5, Kills: 1, TrashPacks: 2, Deaths: 3}
	if got, want := loggedNightLine(ans), "5 present · 1 kill · 2 trash packs · 3 deaths"; got != want {
		t.Errorf("line = %q, want %q", got, want)
	}
	trashOnly := guilds.LoggedNightAnswer{Present: 5, TrashPacks: 1}
	if got, want := loggedNightLine(trashOnly), "5 present · 1 trash pack"; got != want {
		t.Errorf("trash only = %q, want %q", got, want)
	}
}
