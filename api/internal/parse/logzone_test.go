package parse

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/fight"
)

func TestLogZoneNameFallsBackToTheHeaderZone(t *testing.T) {
	cases := []struct {
		name   string
		fights []fight.Fight
		header string
		want   string
	}{
		{"the first encounter fight wins", []fight.Fight{{Zone: "Trash Zone"}, {Zone: "Boss Zone", EncounterID: 7}}, "Header", "Boss Zone"},
		{"else the first fight that named one", []fight.Fight{{}, {Zone: "Trash Zone"}}, "Header", "Trash Zone"},
		{"no fight named a zone: the header's", []fight.Fight{{}, {EncounterID: 7}}, "Header", "Header"},
		{"nothing at all", nil, "", ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var z logZone
			for _, f := range c.fights {
				z.see(f)
			}
			if got := z.name(c.header); got != c.want {
				t.Fatalf("name(%q) = %q, want %q", c.header, got, c.want)
			}
		})
	}
}
