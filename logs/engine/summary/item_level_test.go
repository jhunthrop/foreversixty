package summary

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/layout"
)

func itemLevelByGUID(s Summary) map[string]int64 {
	out := map[string]int64{}
	for _, r := range s.Roster {
		out[r.GUID] = r.ItemLevel
	}
	return out
}

func TestRosterRowItemLevelFromAdvancedBlock(t *testing.T) {
	_, s := excerptSummary(t, filepath.Join("..", "event", "testdata", "v22.log"),
		layout.RetailV22(), time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC))
	if len(s.Roster) == 0 {
		t.Fatal("empty roster")
	}
	for _, r := range s.Roster {
		if r.ItemLevel <= 0 {
			t.Errorf("%s item level = %d, want the advanced block's item level", r.Name, r.ItemLevel)
		}
	}
}

func TestRosterRowItemLevelKeepsTheHighestOwnLine(t *testing.T) {
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
	if got := itemLevelByGUID(s)[mage]; got != 60 {
		t.Fatalf("mage item level = %d, want 60", got)
	}
}

func TestRosterRowItemLevelZeroWithoutAdvancedLogging(t *testing.T) {
	o, _ := opts(t)
	a := New(o)
	a.Start(at(0))
	a.Add(dmg(1, mage, "Creature-0-1-2-3-4-5", 133, "Fireball", 100, 0))
	s := a.Snapshot(fixtureFight(), "test")
	if got := itemLevelByGUID(s)[mage]; got != 0 {
		t.Fatalf("mage item level = %d, want 0 with no advanced block", got)
	}
}

func TestRosterRowLevelIsNeverSetByTheEngine(t *testing.T) {
	_, _, s := build(t)
	for _, r := range s.Roster {
		if r.Level != 0 {
			t.Errorf("%s level = %d, want 0: only the API overlays the character level", r.Name, r.Level)
		}
	}
}
