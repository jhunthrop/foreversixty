package addon

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// fakeRaidRosters answers AddonRaidRosters from a fixed list, or fails.
type fakeRaidRosters struct {
	rows   []guilds.AddonRaidRoster
	err    error
	userID int64
}

func (f *fakeRaidRosters) AddonRaidRosters(_ context.Context, userID int64, _ []string, _ time.Time) ([]guilds.AddonRaidRoster, error) {
	f.userID = userID
	return f.rows, f.err
}

// fixtureRoster is the row the inbox-raid-roster.json fixture is the officer's message of.
func fixtureRoster(officer bool) guilds.AddonRaidRoster {
	size := 40
	return guilds.AddonRaidRoster{
		CharacterKey: contractCharacter, RaidID: 42, ZoneName: "Onyxia's Lair", Title: "Onyxia",
		StartsAt: time.Date(2026, 12, 10, 2, 0, 0, 0, time.UTC), Size: &size, Officer: officer,
		Counts: guilds.AddonRosterCounts{In: 10, Late: 2, Bench: 1, Out: 3, NotSigned: 16},
		Coming: guilds.RoleCounts{Tanks: 2, Healers: 4, Melee: 2, Ranged: 4},
		Wanted: &guilds.RoleCounts{Tanks: 2, Healers: 8, Melee: 10, Ranged: 20},
		JobsFilled: 1, Jobs: 2,
		Members: []guilds.AddonRosterMember{
			{Name: "Thornhide", Class: "warrior", Answer: "in", Group: "tank"},
			{Name: "Velisandre", Class: "hunter", Answer: "late", Group: "ranged"},
			{Name: "Quenna", Class: "priest", Answer: "none", Group: "healer"},
		},
		JobList: []guilds.AddonRosterJob{
			{Encounter: "Onyxia", Label: "Main tank", Count: 1, Holders: []string{"Thornhide"}},
			{Encounter: "Onyxia", Label: "Off tank (phase 3, fear)", Count: 1},
		},
	}
}

// The raid_roster fixture is an officer's message that MessageEntry decodes strictly and writes
// back unchanged.
func TestRaidRosterFixtureGolden(t *testing.T) {
	raw := contractFixture(t, "inbox-raid-roster.json")
	var m MessageEntry
	contractStrict(t, raw, &m)
	contractRoundTrip(t, raw, m)
	if m.Type != messageRaidRoster || m.Signups == nil || m.Signups.NotSigned != 16 || m.Members == nil || len(*m.Members) != 3 {
		t.Fatalf("message = %+v", m)
	}
}

func TestRaidRosterMessagesNamesForOfficersCountsForMembers(t *testing.T) {
	at := time.Date(2026, 12, 9, 2, 0, 0, 0, time.UTC)
	src := &fakeRaidRosters{rows: []guilds.AddonRaidRoster{fixtureRoster(true)}}
	s := &Service{Log: quietLog(), RaidRosters: src}
	got := s.raidRosterMessages(context.Background(), 7, []string{contractCharacter}, at)
	if len(got) != 1 || src.userID != 7 {
		t.Fatalf("messages = %+v, asked as %d", got, src.userID)
	}
	contractRoundTrip(t, contractFixture(t, "inbox-raid-roster.json"), got[0])

	member := fixtureRoster(false)
	member.Wanted = nil
	src.rows = []guilds.AddonRaidRoster{member}
	raw, _ := json.Marshal(s.raidRosterMessages(context.Background(), 7, []string{contractCharacter}, at)[0])
	var fields map[string]any
	_ = json.Unmarshal(raw, &fields)
	for _, absent := range []string{"members", "assignments"} {
		if _, ok := fields[absent]; ok {
			t.Errorf("a member's roster carries %s: %s", absent, raw)
		}
	}
	roles := fields["roles"].([]any)
	if _, ok := roles[0].(map[string]any)["wanted"]; ok || len(roles) != 4 {
		t.Errorf("roles without a target = %v", roles)
	}
	if fields["jobs"].(map[string]any)["total"] != float64(2) {
		t.Errorf("jobs = %v", fields["jobs"])
	}
}

func TestRaidRosterMessagesDegrade(t *testing.T) {
	at := time.Date(2026, 12, 9, 2, 0, 0, 0, time.UTC)
	if none := (&Service{Log: quietLog()}).raidRosterMessages(context.Background(), 7, []string{contractCharacter}, at); none != nil {
		t.Errorf("no source sent %+v", none)
	}
	s := &Service{Log: quietLog(), RaidRosters: &fakeRaidRosters{err: errors.New("down")}}
	if failed := s.raidRosterMessages(context.Background(), 7, []string{contractCharacter}, at); failed != nil {
		t.Errorf("a failed read sent %+v", failed)
	}
}

// Every other message type stays free of the roster's keys.
func TestOtherMessagesCarryNoRosterKeys(t *testing.T) {
	raw, _ := json.Marshal(MessageEntry{Type: messageRaidSignup, Character: contractCharacter, RaidID: 42})
	var fields map[string]any
	_ = json.Unmarshal(raw, &fields)
	for _, key := range []string{"size", "signups", "roles", "jobs", "members", "assignments"} {
		if _, ok := fields[key]; ok {
			t.Errorf("raid_signup carries %s: %s", key, raw)
		}
	}
}
