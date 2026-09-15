// logs/engine/fight/fight_test.go
package fight

import (
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
)

var t0 = time.Date(2026, 9, 26, 20, 10, 0, 0, time.UTC)

const (
	player = "Player-4184-000000A1"
	mage   = "Player-4184-000000A3"
	mob    = "Creature-0-2085-2284-7855-169753-0000AA0001"
)

func at(sec float64) time.Time {
	return t0.Add(time.Duration(sec * float64(time.Second)))
}

func hit(sec float64, src, dst string, srcFlags, dstFlags uint32) event.Event {
	return event.Event{
		Time: at(sec), Kind: event.Damage, Name: "SPELL_DAMAGE",
		Line: int64(sec * 10), Offset: int64(sec * 100),
		Source: event.Unit{GUID: src, Flags: srcFlags},
		Dest:   event.Unit{GUID: dst, Flags: dstFlags},
		Amount: event.OptInt{V: 100, OK: true},
	}
}

// playerHit is a player hitting a hostile NPC.
func playerHit(sec float64) event.Event { return hit(sec, mage, mob, 0x512, 0xa48) }

func run(s *Segmenter, evs []event.Event) []*Fight {
	var closed []*Fight
	for _, e := range evs {
		if st := s.Feed(e); st.Closed != nil {
			closed = append(closed, st.Closed)
		}
	}
	return closed
}

func TestEncounterStartAndEndNameTheFight(t *testing.T) {
	s := NewSegmenter(Options{Trailing: 0})
	evs := []event.Event{
		{Time: at(0), Kind: event.ZoneChange, Zone: &event.Zone{ID: 2284, Name: "Sanguine Depths"}},
		{Time: at(1), Kind: event.EncounterStart, Line: 10,
			Encounter: &event.Encounter{ID: 9001, Name: "Warden Kelthas", Difficulty: 8, Size: 5}},
		playerHit(2),
		{Time: at(40), Kind: event.EncounterEnd, Line: 400,
			Encounter: &event.Encounter{ID: 9001, Name: "Warden Kelthas", Kill: true}},
	}
	closed := run(s, evs)
	if len(closed) != 1 {
		t.Fatalf("closed %d fights, want 1", len(closed))
	}
	f := closed[0]
	if f.Kind != Encounter || f.Name != "Warden Kelthas" || f.EncounterID != 9001 {
		t.Errorf("fight = %+v", *f)
	}
	if !f.Kill || f.Difficulty != 8 || f.Size != 5 {
		t.Errorf("kill=%v difficulty=%d size=%d", f.Kill, f.Difficulty, f.Size)
	}
	if f.Zone != "Sanguine Depths" || f.ZoneID != 2284 {
		t.Errorf("zone = %q %d", f.Zone, f.ZoneID)
	}
	if f.Duration() != 39*time.Second {
		t.Errorf("duration = %s, want 39s", f.Duration())
	}
	if len(f.Players) != 1 || f.Players[0] != mage {
		t.Errorf("players = %v", f.Players)
	}
	if f.Index != 1 || f.InProgress {
		t.Errorf("index = %d in progress = %v", f.Index, f.InProgress)
	}
}

func TestTrailingWindowKeepsDotTicksInsideTheEncounter(t *testing.T) {
	s := NewSegmenter(Options{Trailing: 2 * time.Second})
	evs := []event.Event{
		{Time: at(1), Kind: event.EncounterStart,
			Encounter: &event.Encounter{ID: 9001, Name: "Warden Kelthas"}},
		playerHit(2),
		{Time: at(40), Kind: event.EncounterEnd,
			Encounter: &event.Encounter{ID: 9001, Name: "Warden Kelthas", Kill: true}},
		playerHit(41), // a tick inside the window
		playerHit(50), // past the window: this one closes the fight
	}
	var closed []*Fight
	var inside, outside bool
	for i, e := range evs {
		st := s.Feed(e)
		if st.Closed != nil {
			closed = append(closed, st.Closed)
		}
		if i == 3 && st.Fight != nil {
			inside = true
		}
		if i == 4 && st.Fight != nil && st.Fight.Kind == Trash {
			outside = true
		}
	}
	if !inside {
		t.Error("a tick inside the trailing window must still belong to the encounter")
	}
	if len(closed) != 1 || closed[0].Kind != Encounter {
		t.Fatalf("closed = %d fights", len(closed))
	}
	if closed[0].End != at(42) {
		t.Errorf("encounter ended at %s, want the end of the trailing window", closed[0].End)
	}
	if !outside {
		t.Error("the event past the window opens a new trash fight")
	}
}

func TestTrashIsSegmentedByGapsWithNoEncounterEvents(t *testing.T) {
	s := NewSegmenter(Options{Gap: 5 * time.Second, MinTrash: time.Second})
	var evs []event.Event
	for _, sec := range []float64{0, 1, 2, 3} { // pull one
		evs = append(evs, playerHit(sec))
	}
	for _, sec := range []float64{20, 21, 22} { // pull two, after a long gap
		evs = append(evs, playerHit(sec))
	}
	closed := run(s, evs)
	if len(closed) != 1 {
		t.Fatalf("closed %d fights mid-stream, want 1", len(closed))
	}
	if closed[0].Kind != Trash || closed[0].Name != "Trash" {
		t.Errorf("first pull = %+v", *closed[0])
	}
	if closed[0].Duration() != 3*time.Second {
		t.Errorf("first pull duration = %s, want 3s", closed[0].Duration())
	}
	last := s.Flush(at(25))
	if last == nil || last.Index != 2 {
		t.Fatalf("flush = %+v", last)
	}
	if last.Duration() != 2*time.Second {
		t.Errorf("second pull duration = %s, want 2s", last.Duration())
	}
}

func TestAStraySwingIsNotAFight(t *testing.T) {
	s := NewSegmenter(Options{Gap: 5 * time.Second, MinTrash: 3 * time.Second})
	run(s, []event.Event{playerHit(0)})
	if f := s.Flush(at(1)); f != nil {
		t.Fatalf("a one-second segment became fight %+v", *f)
	}
	// The index was not consumed, so the next real fight is still 1.
	run(s, []event.Event{playerHit(100), playerHit(104)})
	f := s.Flush(at(105))
	if f == nil || f.Index != 1 {
		t.Fatalf("next fight = %+v, want index 1", f)
	}
}

func TestFriendlyOnlyEventsDoNotOpenAFight(t *testing.T) {
	s := NewSegmenter(Options{})
	evs := []event.Event{
		{Time: at(0), Kind: event.Heal,
			Source: event.Unit{GUID: player, Flags: 0x512},
			Dest:   event.Unit{GUID: mage, Flags: 0x512}},
		hit(1, player, mage, 0x512, 0x512), // a duel-free friendly fire line
		{Time: at(2), Kind: event.AuraApplied,
			Source: event.Unit{GUID: player, Flags: 0x512},
			Dest:   event.Unit{GUID: mage, Flags: 0x512}},
	}
	for _, e := range evs {
		if st := s.Feed(e); st.Fight != nil {
			t.Fatalf("event %s opened a fight", e.Kind)
		}
	}
	if s.Open() != nil {
		t.Error("no fight should be open")
	}
}

func TestKillsAndDeathsAreCounted(t *testing.T) {
	s := NewSegmenter(Options{MinTrash: 0})
	evs := []event.Event{
		playerHit(0),
		{Time: at(1), Kind: event.Death, Dest: event.Unit{GUID: mob, Flags: 0xa48}},
		{Time: at(2), Kind: event.Death, Dest: event.Unit{GUID: player, Flags: 0x512}},
		playerHit(3),
	}
	run(s, evs)
	f := s.Flush(at(4))
	if f == nil {
		t.Fatal("no fight")
	}
	if f.NPCKills != 1 {
		t.Errorf("npc kills = %d, want 1", f.NPCKills)
	}
	if f.Deaths != 1 {
		t.Errorf("player deaths = %d, want 1", f.Deaths)
	}
}

func TestRaidMarkersAreRecordedOnce(t *testing.T) {
	s := NewSegmenter(Options{MinTrash: 0})
	marked := playerHit(0)
	marked.Dest.Raid = 0x08 // triangle
	marked.Dest.Name = "Hollow Sentinel"
	again := playerHit(1)
	again.Dest.Raid = 0x08
	run(s, []event.Event{marked, again})
	f := s.Flush(at(2))
	if f == nil || len(f.Markers) != 1 {
		t.Fatalf("markers = %+v", f)
	}
	if f.Markers[0].Flag != 0x08 || f.Markers[0].Name != "Hollow Sentinel" {
		t.Errorf("marker = %+v", f.Markers[0])
	}
}

func TestOpenIsACopySoCallersCannotMutateTheSegmenter(t *testing.T) {
	s := NewSegmenter(Options{})
	s.Feed(playerHit(0))
	open := s.Open()
	if open == nil || !open.InProgress {
		t.Fatal("expected an open fight")
	}
	open.Name = "tampered"
	if again := s.Open(); again.Name == "tampered" {
		t.Error("Open must return a copy")
	}
}

func TestStateRoundTripKeepsTheOpenFight(t *testing.T) {
	s := NewSegmenter(Options{MinTrash: 0})
	s.Feed(event.Event{Time: at(0), Kind: event.ZoneChange, Zone: &event.Zone{ID: 2284, Name: "Sanguine Depths"}})
	s.Feed(event.Event{Time: at(1), Kind: event.EncounterStart,
		Encounter: &event.Encounter{ID: 9001, Name: "Warden Kelthas"}})
	s.Feed(playerHit(2))

	revived := RestoreSegmenter(Options{MinTrash: 0}, s.State())
	revived.Feed(playerHit(3))
	end := revived.Feed(event.Event{Time: at(10), Kind: event.EncounterEnd,
		Encounter: &event.Encounter{ID: 9001, Name: "Warden Kelthas", Kill: true}})
	f := end.Closed
	if f == nil {
		f = revived.Flush(at(12))
	}
	if f == nil {
		t.Fatal("the open fight was lost across restore")
	}
	if f.Name != "Warden Kelthas" || !f.Kill || f.Zone != "Sanguine Depths" {
		t.Errorf("fight = %+v", *f)
	}
	if len(f.Players) != 1 {
		t.Errorf("players lost across restore: %v", f.Players)
	}
}

func TestFlushOnAnEmptySegmenterIsNil(t *testing.T) {
	if f := NewSegmenter(Options{}).Flush(at(0)); f != nil {
		t.Fatalf("flush = %+v, want nil", f)
	}
}

// TestADiscardedTrashSegmentIsReportedAsDiscarded covers the signal a
// caller needs to drop the accumulator it opened for a segment that
// closes without being reported.
func TestADiscardedTrashSegmentIsReportedAsDiscarded(t *testing.T) {
	s := NewSegmenter(Options{Gap: 5 * time.Second, MinTrash: 3 * time.Second})
	// One stray swing, then a long enough quiet period to close it.
	if st := s.Feed(playerHit(0)); !st.Opened {
		t.Fatal("the first hit did not open a fight")
	}
	st := s.Feed(playerHit(100))
	if st.Closed != nil {
		t.Fatalf("a one-second trash segment was reported as fight %+v", *st.Closed)
	}
	if !st.Discarded {
		t.Fatal("the discarded segment was not reported as discarded")
	}
	// A fight that closes normally is reported, not discarded.
	next := s.Feed(playerHit(104))
	if next.Discarded {
		t.Error("nothing was discarded on this event")
	}
	closed := s.Feed(playerHit(200))
	if closed.Closed == nil {
		t.Fatal("the four-second segment should have been reported")
	}
	if closed.Discarded {
		t.Error("a reported fight must not also be marked discarded")
	}
}

func TestAWipeRemembersTheBossHealth(t *testing.T) {
	s := NewSegmenter(Options{Trailing: 0})
	boss := "Creature-0-1-2-3-9001-0000000001"
	evs := []event.Event{
		{Time: at(1), Kind: event.EncounterStart, Line: 10,
			Encounter: &event.Encounter{ID: 9001, Name: "Warden Kelthas", Difficulty: 8, Size: 5}},
		playerHit(2),
		// A player's spell on the boss: the advanced block describes the caster, so it
		// says nothing about the boss.
		{Time: at(3), Kind: event.Damage, Name: "SPELL_DAMAGE",
			Source: event.Unit{GUID: "Player-1-A", Name: "Mage", Flags: 0x511},
			Dest:   event.Unit{GUID: boss, Name: "Warden Kelthas", Flags: 0xa48},
			Adv:    event.Advanced{OK: true, InfoGUID: "Player-1-A", CurrentHP: 100, MaxHP: 100}},
		// The boss's own swing carries the boss's health.
		{Time: at(4), Kind: event.Damage, Name: "SWING_DAMAGE",
			Source: event.Unit{GUID: boss, Name: "Warden Kelthas", Flags: 0xa48},
			Dest:   event.Unit{GUID: "Player-1-A", Name: "Mage", Flags: 0x511},
			Adv:    event.Advanced{OK: true, InfoGUID: boss, CurrentHP: 2300, MaxHP: 10000}},
		{Time: at(40), Kind: event.EncounterEnd, Line: 400,
			Encounter: &event.Encounter{ID: 9001, Name: "Warden Kelthas", Kill: false}},
	}
	closed := run(s, evs)
	if len(closed) != 1 {
		t.Fatalf("closed %d fights, want 1", len(closed))
	}
	if got := closed[0].BossHealthPct; got != 23 {
		t.Errorf("boss health = %v, want 23 (the last block that named the boss)", got)
	}
}
