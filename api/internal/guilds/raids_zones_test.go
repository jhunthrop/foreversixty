package guilds

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/phase"
)

const (
	firstTierPath = "../../../data/curated/raids/first-tier.json"
	eventsPath    = "../../../data/curated/raids/events.json"
)

func readJSON(t *testing.T, path string, into any) {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(raw, into); err != nil {
		t.Fatal(err)
	}
}

func TestRaidZonesAreFirstTierJSON(t *testing.T) {
	var file struct {
		Zones []struct {
			Zone
			Opens string `json:"opens"`
		} `json:"zones"`
	}
	readJSON(t, firstTierPath, &file)
	if len(file.Zones) != len(raidZones) {
		t.Fatalf("first-tier.json has %d zones, Go %d", len(file.Zones), len(raidZones))
	}
	for i, f := range file.Zones {
		want := f.Zone
		if !phase.Valid(f.Opens) {
			t.Fatalf("%s opens %q, not a phase", f.Key, f.Opens)
		}
		want.Opens = phaseStart(f.Opens)
		if !reflect.DeepEqual(raidZones[i], want) {
			t.Errorf("raid zone %d drifted from first-tier.json:\n got %+v\nwant %+v", i, raidZones[i], want)
		}
	}
}

func TestEventZonesAreEventsJSON(t *testing.T) {
	var file struct {
		Zones []Zone `json:"zones"`
	}
	readJSON(t, eventsPath, &file)
	if !reflect.DeepEqual(eventZones, file.Zones) {
		t.Fatalf("eventZones drifted from events.json; run `python -m pipeline guild-events`")
	}
}

// Spec §3.1: dungeons in level order, then the six guild events, then the raids; every
// dungeon is a party or raid instance with a sourced size and level range.
func TestCuratedZoneOrderAndFacts(t *testing.T) {
	kinds := []string{}
	for _, z := range Zones() {
		if len(kinds) == 0 || kinds[len(kinds)-1] != z.Kind {
			kinds = append(kinds, z.Kind)
		}
		if z.Kind == ZoneDungeon && (z.Size == nil || z.MinLevel == nil || z.MaxLevel == nil) {
			t.Errorf("%s lacks a sourced size or level range", z.Key)
		}
	}
	if !reflect.DeepEqual(kinds, []string{ZoneDungeon, ZoneEvent, ZoneRaid}) {
		t.Fatalf("kind order = %v", kinds)
	}
	sizes := map[string]*int{}
	for _, z := range Zones() {
		if z.Kind == ZoneEvent {
			if z.Name != "Guild event" {
				t.Errorf("%s is named %q", z.Key, z.Name)
			}
			sizes[z.Key] = z.Size
		}
	}
	want := map[string]*int{
		"event:guild": nil, "event:guild-5": intPtr(5), "event:guild-10": intPtr(10),
		"event:guild-15": intPtr(15), "event:guild-20": intPtr(20), "event:guild-40": intPtr(40),
	}
	if !reflect.DeepEqual(sizes, want) {
		t.Fatalf("event sizes = %v", sizes)
	}
}

func TestZoneByKey(t *testing.T) {
	cases := []struct {
		key  string
		name string
		kind string
		ok   bool
	}{
		{"raid:onyxias-lair", "Onyxia's Lair", ZoneRaid, true},
		{"raid:barrow-deeps", "Barrow Deeps", ZoneRaid, true},
		{"dungeon:the-deadmines", "The Deadmines", ZoneDungeon, true},
		{"dungeon:scarlet-monastery-library", "Scarlet Monastery: Library", ZoneDungeon, true},
		{"event:guild-10", "Guild event", ZoneEvent, true},
		{"raid:nope", "", "", false},
	}
	for _, c := range cases {
		z, ok := ZoneByKey(c.key)
		if ok != c.ok || z.Name != c.name || z.Kind != c.kind {
			t.Errorf("ZoneByKey(%q) = %+v, %v", c.key, z, ok)
		}
	}
}

func TestZoneOpenedByReadsTheGuildCalendar(t *testing.T) {
	ony, _ := ZoneByKey("raid:onyxias-lair")
	chicago, err := time.LoadLocation("America/Chicago")
	if err != nil {
		t.Fatal(err)
	}
	// 8 Dec 8 pm Chicago is 9 Dec 02:00 UTC: still the 8th on the guild's calendar.
	if ony.OpenedBy(time.Date(2026, 12, 9, 2, 0, 0, 0, time.UTC), chicago) {
		t.Error("8 Dec guild time opened Onyxia")
	}
	if !ony.OpenedBy(time.Date(2026, 12, 10, 2, 0, 0, 0, time.UTC), chicago) {
		t.Error("9 Dec guild time did not open Onyxia")
	}
	dm, _ := ZoneByKey("dungeon:the-deadmines")
	if !dm.OpenedBy(time.Date(2026, 11, 5, 0, 0, 0, 0, time.UTC), time.UTC) {
		t.Error("a dungeon with no Opens was closed")
	}
}
