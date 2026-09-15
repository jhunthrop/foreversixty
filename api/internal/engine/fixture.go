package engine

import (
	"fmt"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/parquet"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// FixtureBase is the year the fixture log's year-less timestamps resolve
// against, so every test that uses the fixture sees the same dates.
var FixtureBase = time.Date(2026, 12, 9, 22, 0, 0, 0, time.UTC)

// fixtureLog is one retail-v16 encounter: a kill of encounter 9001 by
// three players, with the advanced block on. The lines are the shapes
// logs/engine/event/testdata/v16.log pins, so the decoder reads every
// field rather than falling back to a raw event.
const fixtureLog = `12/9 22:00:00.000  COMBAT_LOG_VERSION,16,ADVANCED_LOG_ENABLED,1,BUILD_VERSION,9.0.2,PROJECT_ID,1
12/9 22:00:01.000  ZONE_CHANGE,2284,"Blackrock Spire",8
12/9 22:00:10.000  ENCOUNTER_START,9001,"Warden Kelthas",8,5,2284
12/9 22:00:11.000  SPELL_DAMAGE,Player-4184-000000A3,"Morrowlyn-Nightslayer",0x512,0x0,Creature-0-2085-2284-7855-169753-0000AA0001,"Hollow Sentinel",0xa48,0x0,116,"Frostbolt",0x10,Creature-0-2085-2284-7855-169753-0000AA0001,0000000000000000,41320,44000,612,0,1955,0,0,0,0,0,-1487.02,6409.71,1675,1.2044,45,1484,1390,-1,16,0,0,0,1,nil,nil
12/9 22:00:12.000  SWING_DAMAGE_LANDED,Creature-0-2085-2284-7855-169753-0000AA0001,"Hollow Sentinel",0xa48,0x0,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0,Player-4184-000000A1,0000000000000000,10188,11000,388,0,4120,0,1,0,1000,0,-1489.90,6410.05,1675,4.1002,183,812,1290,-1,1,0,0,0,nil,nil,nil
12/9 22:00:13.000  SPELL_HEAL,Player-4184-000000A2,"Sunwick-Nightslayer",0x512,0x0,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0,2060,"Heal",0x2,Player-4184-000000A1,0000000000000000,11000,11000,388,0,4120,0,0,9330,9330,0,-1489.90,6410.05,1675,4.1002,183,1618,1618,806,0,1
12/9 22:00:14.000  SPELL_DAMAGE,Player-4184-000000A3,"Morrowlyn-Nightslayer",0x512,0x0,Creature-0-2085-2284-7855-169753-0000AA0001,"Hollow Sentinel",0xa48,0x0,116,"Frostbolt",0x10,Creature-0-2085-2284-7855-169753-0000AA0001,0000000000000000,31320,44000,612,0,1955,0,0,0,0,0,-1487.02,6409.71,1675,1.2044,45,2484,2390,-1,16,0,0,0,1,nil,nil
12/9 22:00:40.000  ENCOUNTER_END,9001,"Warden Kelthas",8,5,1
`

// Fixture is one parsed fight: what a companion would send for it, and
// what the API must be able to rebuild from the Parquet alone.
type Fixture struct {
	Log     []byte
	Fight   fight.Fight
	Summary summary.Summary
	Events  []event.Event
	Parquet []byte
	Health  session.Health
}

// FixtureLog is the raw bytes of the fixture log, for tests that feed a
// whole file through the parse job.
func FixtureLog() []byte { return []byte(fixtureLog) }

// NewFixture parses the fixture log with the API's own options and hands
// back the single fight it contains. It fails only if the engine and this
// package disagree, which is exactly what the tests using it check.
func NewFixture(reportID string) (Fixture, error) {
	s := session.New(SessionOptions(reportID, FixtureBase, true))
	res, err := s.Feed([]byte(fixtureLog), 0)
	if err != nil {
		return Fixture{}, err
	}
	closing, err := s.Close()
	if err != nil {
		return Fixture{}, err
	}
	closed := append(res.Closed, closing.Closed...)
	if len(closed) != 1 {
		return Fixture{}, fmt.Errorf("engine: the fixture log holds %d fights, want 1", len(closed))
	}
	c := closed[0]
	pq, err := parquet.Marshal(c.Events)
	if err != nil {
		return Fixture{}, err
	}
	return Fixture{
		Log: []byte(fixtureLog), Fight: c.Fight, Summary: c.Summary,
		Events: c.Events, Parquet: pq, Health: s.Health(),
	}, nil
}

// FixtureLines is the fixture log's line count, for tests that assert on
// health counters.
func FixtureLines() int { return strings.Count(fixtureLog, "\n") }
