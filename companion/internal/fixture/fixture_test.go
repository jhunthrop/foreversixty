package fixture

import (
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// TestTheFixtureParsesIntoWholeFights is the guard on every other
// test in the companion: if the engine cannot read this log, no
// pipeline test below means anything.
func TestTheFixtureParsesIntoWholeFights(t *testing.T) {
	o := session.Options{
		ReportID: "fixture", KeepEvents: true,
		Base:    time.Date(2026, 9, 26, 0, 0, 0, 0, time.UTC),
		Units:   units.Options{ClassBySpec: units.RetailSpecClass},
		Fight:   fight.DefaultOptions(),
		Summary: summary.DefaultOptions(),
	}
	s := session.New(o)
	log := Log(3)
	res, err := s.Feed([]byte(log), 0)
	if err != nil {
		t.Fatal(err)
	}
	tail, err := s.Close()
	if err != nil {
		t.Fatal(err)
	}
	closed := append(res.Closed, tail.Closed...)
	var encounters []string
	for _, c := range closed {
		if c.Fight.Kind == fight.Encounter {
			encounters = append(encounters, c.Fight.Name)
		}
	}
	if len(encounters) != 3 {
		t.Fatalf("encounters = %v, want three", encounters)
	}
	if !strings.HasPrefix(encounters[0], "Warden Kelthas") {
		t.Errorf("first encounter = %q", encounters[0])
	}
	h := s.Health()
	if h.ParseErrors != 0 {
		t.Errorf("the fixture has %d parse errors", h.ParseErrors)
	}
	if !h.AdvancedLogging {
		t.Error("the fixture must have advanced logging on")
	}
	for _, c := range closed {
		if c.Fight.Kind == fight.Encounter && len(c.Summary.Roster) == 0 {
			t.Errorf("fight %d has an empty roster", c.Fight.Index)
		}
	}
}

// The logger is the one player whose flags carry MINE: the fixture
// writes its first player that way and the rest of the group as PARTY,
// which is what the companion's logging-character detection reads.
func TestTheFixtureLogsAsOnePlayerAndThePartyAroundIt(t *testing.T) {
	s := session.New(session.Options{
		ReportID: "fixture", Base: time.Date(2026, 9, 26, 0, 0, 0, 0, time.UTC),
		Units: units.Options{ClassBySpec: units.RetailSpecClass},
		Fight: fight.DefaultOptions(), Summary: summary.DefaultOptions(),
	})
	if _, err := s.Feed([]byte(Log(1)), 0); err != nil {
		t.Fatal(err)
	}
	var mine, party []string
	for _, u := range s.Units().Players() {
		switch {
		case u.Flags&units.FlagAffiliationMine != 0:
			mine = append(mine, u.Name)
		case u.Flags&units.FlagAffiliationParty != 0:
			party = append(party, u.Name)
		}
	}
	if len(mine) != 1 || !strings.HasPrefix(mine[0], Logger+"-") {
		t.Errorf("MINE players = %v, want only %s", mine, Logger)
	}
	if len(party) != 1 || !strings.HasPrefix(party[0], "Brannic-") {
		t.Errorf("PARTY players = %v, want Brannic", party)
	}
}
