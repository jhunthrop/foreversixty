package addon

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// fakeAssignments answers NextRaidAssignments from a fixed list, or fails.
type fakeAssignments struct {
	held []guilds.CharacterAssignment
	err  error
	keys []string
}

func (f *fakeAssignments) NextRaidAssignments(_ context.Context, keys []string, _ time.Time) ([]guilds.CharacterAssignment, error) {
	f.keys = keys
	return f.held, f.err
}

// The raid_assignment fixture is a flat message (strings and numbers only, so the companion
// carries it through) that MessageEntry decodes strictly and writes back unchanged.
func TestRaidAssignmentFixtureGolden(t *testing.T) {
	raw := contractFixture(t, "inbox-raid-assignment.json")
	var m MessageEntry
	contractStrict(t, raw, &m)
	contractRoundTrip(t, raw, m)
	if m.Type != messageRaidAssignment || m.Character != contractCharacter || m.RaidID != 42 || m.Warning != "OT" {
		t.Fatalf("message = %+v", m)
	}
}

func TestRaidAssignmentMessagesOnePerHeldSlot(t *testing.T) {
	at := time.Date(2026, 12, 10, 2, 0, 0, 0, time.UTC)
	src := &fakeAssignments{held: []guilds.CharacterAssignment{
		{CharacterKey: contractCharacter, RaidID: 42, RaidTitle: "Onyxia", ZoneName: "Onyxia's Lair", StartsAt: at,
			Encounter: "Onyxia", SlotKey: "off_tank", Label: "Off tank (phase 3, fear)", Warning: "OT",
			Purpose: "Builds threat from the landing and takes her when Bellowing Roar or Knock Away drops the main tank.",
			With:    []string{"Thornhide"}},
		{CharacterKey: contractCharacter, RaidID: 42, ZoneName: "Onyxia's Lair", StartsAt: at, SlotKey: "soulstones", With: []string{}},
	}}
	s := &Service{Log: quietLog(), Assignments: src}
	got := s.raidAssignmentMessages(context.Background(), []string{contractCharacter}, at)
	want := contractFixture(t, "inbox-raid-assignment.json")
	if len(got) != 2 || len(src.keys) != 1 {
		t.Fatalf("messages = %+v", got)
	}
	contractRoundTrip(t, want, got[0])
	if got[1].With != "" || got[1].Slot != "soulstones" {
		t.Errorf("second = %+v", got[1])
	}
	if none := (&Service{Log: quietLog()}).raidAssignmentMessages(context.Background(), []string{contractCharacter}, at); none != nil {
		t.Errorf("no source sent %+v", none)
	}
	src.err = errors.New("down")
	if failed := s.raidAssignmentMessages(context.Background(), []string{contractCharacter}, at); failed != nil {
		t.Errorf("a failed read sent %+v", failed)
	}
}
