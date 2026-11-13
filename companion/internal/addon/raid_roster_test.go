package addon

import (
	"encoding/json"
	"testing"
)

// The API's raid_roster message (docs/contracts/fixtures/inbox-raid-roster.json) is nested, so
// the companion knows the type: it decodes field by field and renders every table, the officer's
// members and assignments included.
func TestTheRaidRosterFixtureIsRendered(t *testing.T) {
	raw := contractFixture(t, "inbox-raid-roster.json")
	var strict Message
	contractStrict(t, raw, &strict) // every field of the API's message has a home in Message
	var in Inbox
	if err := json.Unmarshal([]byte(`{"format":2,"generated_at":"2026-11-04T21:03:00Z","builds":[],"messages":[`+string(raw)+`]}`), &in); err != nil {
		t.Fatal(err)
	}
	if len(in.Messages) != 1 || len(in.Skipped) != 0 {
		t.Fatalf("messages = %+v, skipped = %+v", in.Messages, in.Skipped)
	}
	row := parseInbox(t, string(RenderInbox(in, t0))).Fields["messages"].(*Table).Items[0].(*Table)
	if row.Get("type") != "raid_roster" || row.Get("character") != "us/normal/obnoxious" || keyOf(row.Fields["raid_id"]) != "42" ||
		row.Get("raid") != "Onyxia's Lair" || row.Get("starts_at") != "2026-12-10T02:00:00Z" || keyOf(row.Fields["size"]) != "40" {
		t.Errorf("rendered = %+v", row.Fields)
	}
	signups := row.Fields["signups"].(*Table)
	if keyOf(signups.Fields["in"]) != "10" || keyOf(signups.Fields["none"]) != "16" {
		t.Errorf("signups = %+v", signups.Fields)
	}
	roles := row.Fields["roles"].(*Table).Items
	if len(roles) != 4 || roles[1].(*Table).Get("role") != "healer" || keyOf(roles[1].(*Table).Fields["wanted"]) != "8" {
		t.Errorf("roles = %+v", roles)
	}
	members := row.Fields["members"].(*Table).Items
	if len(members) != 3 || members[2].(*Table).Get("answer") != "none" || members[1].(*Table).Get("class") != "hunter" {
		t.Errorf("members = %+v", members)
	}
	jobs := row.Fields["assignments"].(*Table).Items
	open := jobs[1].(*Table).Fields["holders"].(*Table)
	if len(jobs) != 2 || jobs[0].(*Table).Fields["holders"].(*Table).Items[0] != "Thornhide" || len(open.Items) != 0 {
		t.Errorf("assignments = %+v", jobs)
	}
}

// A member's roster has no names: the companion writes neither table.
func TestAMembersRaidRosterHasNoNames(t *testing.T) {
	in := Inbox{Messages: []Message{{Type: MessageRaidRoster, Character: "us/normal/obnoxious", RaidID: 7,
		Raid: "Onyxia's Lair", Signups: &RosterSignups{In: 1}, Roles: []RosterRole{{Role: "tank", Coming: 1}},
		Jobs: &RosterJobs{Held: 0, Total: 2}}}}
	row := parseInbox(t, string(RenderInbox(in, t0))).Fields["messages"].(*Table).Items[0].(*Table)
	for _, absent := range []string{"members", "assignments", "size"} {
		if _, ok := row.Fields[absent]; ok {
			t.Errorf("a member's roster carries %s: %+v", absent, row.Fields)
		}
	}
	if _, ok := row.Fields["roles"].(*Table).Items[0].(*Table).Fields["wanted"]; ok {
		t.Error("a role without a target carries wanted")
	}
}
