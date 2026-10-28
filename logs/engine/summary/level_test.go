package summary

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/layout"
)

func levelByGUID(s Summary) map[string]int64 {
	out := map[string]int64{}
	for _, r := range s.Roster {
		out[r.GUID] = r.Level
	}
	return out
}

func TestRosterRowLevelFromAdvancedBlock(t *testing.T) {
	_, s := excerptSummary(t, filepath.Join("..", "event", "testdata", "v22.log"),
		layout.RetailV22(), time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC))
	if len(s.Roster) == 0 {
		t.Fatal("empty roster")
	}
	for _, r := range s.Roster {
		if r.Level <= 0 {
			t.Errorf("%s level = %d, want the advanced block's level", r.Name, r.Level)
		}
	}
}

func TestRosterRowLevelKeepsTheHighestOwnLine(t *testing.T) {
	o, _ := opts(t)
	a := New(o)
	a.Start(at(0))
	for _, lvl := range []int64{58, 60, 59} {
		e := dmg(1, mage, "Creature-0-1-2-3-4-5", 133, "Fireball", 100, 0)
		e.Adv = event.Advanced{OK: true, InfoGUID: mage, Level: lvl}
		a.Add(e)
	}
	// A line whose block describes another unit never sets the player's level.
	other := dmg(2, mage, "Creature-0-1-2-3-4-5", 133, "Fireball", 100, 0)
	other.Adv = event.Advanced{OK: true, InfoGUID: "Creature-0-1-2-3-4-5", Level: 63}
	a.Add(other)
	s := a.Snapshot(fixtureFight(), "test")
	if got := levelByGUID(s)[mage]; got != 60 {
		t.Fatalf("mage level = %d, want 60", got)
	}
}

func TestRosterRowLevelZeroWithoutAdvancedLogging(t *testing.T) {
	_, _, s := build(t)
	for guid, lvl := range levelByGUID(s) {
		if lvl != 0 {
			t.Errorf("%s level = %d, want 0 with no advanced block carrying one", guid, lvl)
		}
	}
}
