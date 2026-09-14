// companion/internal/fixture/fixture.go
// Package fixture writes the combat log the companion's tests parse.
// Every line is hand-written in the v16 shapes the engine verified, and
// uses invented characters — the engine's fixture policy applies here
// too: nothing is copied out of a downloaded log.
package fixture

import (
	"fmt"
	"strings"
	"time"
)

// Header is the log's first line: v16, advanced logging on.
const Header = "9/26 20:10:00.000  COMBAT_LOG_VERSION,16,ADVANCED_LOG_ENABLED,1," +
	"BUILD_VERSION,9.0.2,PROJECT_ID,1\n"

// Zone is a zone change into the fixture's instance.
const Zone = "9/26 20:10:00.500  ZONE_CHANGE,2284,\"Sanguine Depths\",8\n"

// player and boss are invented; see logs/README.md for the cast.
const (
	player   = `Player-4184-000000A3,"Morrowlyn-Nightslayer",0x512,0x0`
	boss     = `Creature-0-2085-2284-7855-169753-0000AA0001,"Hollow Sentinel",0xa48,0x0`
	advanced = `Creature-0-2085-2284-7855-169753-0000AA0001,0000000000000000,41320,44000,` +
		`612,0,1955,0,0,0,0,0,-1487.02,6409.71,1675,1.2044,45`
	damage = `1484,1390,-1,16,0,0,0,1,nil,nil`
)

// Encounter is one boss fight: a start, a few damage events, a kill and
// an end. Index n moves it along the clock so the fights do not
// overlap, and the encounter id changes with it.
func Encounter(n int) string {
	start := time.Date(2026, 9, 26, 20, 12, 0, 0, time.UTC).Add(time.Duration(n) * 5 * time.Minute)
	at := func(d time.Duration) string {
		t := start.Add(d)
		return fmt.Sprintf("9/26 %02d:%02d:%02d.000  ", t.Hour(), t.Minute(), t.Second())
	}
	id := 9001 + n
	name := fmt.Sprintf("Warden Kelthas %d", n)
	var b strings.Builder
	fmt.Fprintf(&b, "%sENCOUNTER_START,%d,%q,8,5,2284\n", at(0), id, name)
	for i := range 3 {
		fmt.Fprintf(&b, "%sSPELL_DAMAGE,%s,%s,116,\"Frostbolt\",0x10,%s,%s\n",
			at(time.Duration(i+1)*time.Second), player, boss, advanced, damage)
	}
	fmt.Fprintf(&b, "%sUNIT_DIED,0000000000000000,nil,0x80000000,0x80000000,%s,0\n", at(30*time.Second), boss)
	fmt.Fprintf(&b, "%sENCOUNTER_END,%d,%q,8,5,1\n", at(31*time.Second), id, name)
	return b.String()
}

// Heartbeat is one line after a fight. The engine emits a closed
// fight when the line after its last one arrives, so a test that
// writes an encounter and stops would leave that fight open; a real
// raid's log never stops for long enough to notice.
func Heartbeat(n int) string {
	t := time.Date(2026, 9, 26, 20, 12, 45, 0, time.UTC).Add(time.Duration(n) * 5 * time.Minute)
	return fmt.Sprintf("9/26 %02d:%02d:%02d.000  ZONE_CHANGE,2284,\"Sanguine Depths\",8\n",
		t.Hour(), t.Minute(), t.Second())
}

// Log is a whole night: the header, a zone change and n encounters.
func Log(n int) string {
	var b strings.Builder
	b.WriteString(Header)
	b.WriteString(Zone)
	for i := range n {
		b.WriteString(Encounter(i))
		b.WriteString(Heartbeat(i))
	}
	return b.String()
}
