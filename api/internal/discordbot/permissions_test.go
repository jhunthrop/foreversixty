package discordbot

import (
	"reflect"
	"testing"
)

func ptr(s string) *string { return &s }

func TestPostableChannels(t *testing.T) {
	const srv, bot, botRole = "srv", "bot", "r-bot"
	base := PermViewChannel | PermSendMessages | PermEmbedLinks | PermReadMessageHistory
	layout := ServerLayout{
		BotUserID: bot, BotRoleIDs: []string{botRole},
		Roles: []Role{{ID: srv, Permissions: base}, {ID: botRole, Permissions: 0}},
		Channels: []RawChannel{
			{ID: "9", Name: "voice", Type: 2},
			{ID: "cat2", Name: "Raids", Type: 4, Position: 1},
			{ID: "cat1", Name: "Info", Type: 4, Position: 0},
			{ID: "5", Name: "raids", Type: 0, Position: 0, ParentID: ptr("cat2"),
				Overwrites: []Overwrite{{ID: srv, Type: OverwriteRole, Deny: PermSendMessages}, {ID: botRole, Type: OverwriteRole, Allow: PermSendMessages}}},
			{ID: "4", Name: "rules", Type: 0, Position: 3, ParentID: ptr("cat1")},
			{ID: "3", Name: "general", Type: 0, Position: 1,
				Overwrites: []Overwrite{{ID: srv, Type: OverwriteRole, Deny: PermSendMessages}}},
			{ID: "2", Name: "welcome", Type: 0, Position: 5},
			{ID: "6", Name: "tied-b", Type: 0, Position: 5},
			{ID: "7", Name: "hidden", Type: 0, Position: 6,
				Overwrites: []Overwrite{{ID: bot, Type: OverwriteMember, Deny: PermViewChannel}}},
			{ID: "8", Name: "granted", Type: 0, Position: 7,
				Overwrites: []Overwrite{{ID: srv, Type: OverwriteRole, Deny: PermViewChannel}, {ID: bot, Type: OverwriteMember, Allow: PermViewChannel}}},
		},
	}
	got := PostableChannels(srv, layout)
	want := []Channel{
		{ID: "2", Name: "welcome"}, {ID: "6", Name: "tied-b"}, {ID: "8", Name: "granted"},
		{ID: "4", Name: "rules", Category: ptr("Info")}, {ID: "5", Name: "raids", Category: ptr("Raids")},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v\nwant %+v", got, want)
	}
}

func TestPostableChannelsAdministratorGrantsEverything(t *testing.T) {
	layout := ServerLayout{
		BotUserID: "bot", BotRoleIDs: []string{"adm"},
		Roles: []Role{{ID: "srv", Permissions: 0}, {ID: "adm", Permissions: PermAdministrator}},
		Channels: []RawChannel{
			{ID: "1", Name: "locked", Type: 0, Overwrites: []Overwrite{{ID: "srv", Type: OverwriteRole, Deny: PermViewChannel | PermSendMessages}}},
			{ID: "2", Name: "voice", Type: 2},
		},
	}
	got := PostableChannels("srv", layout)
	if len(got) != 1 || got[0].ID != "1" {
		t.Fatalf("got %+v", got)
	}
}

func TestPostableChannelsMissingPermissionsYieldsNothing(t *testing.T) {
	layout := ServerLayout{BotUserID: "bot", Channels: []RawChannel{{ID: "1", Name: "a", Type: 0}}}
	if got := PostableChannels("srv", layout); len(got) != 0 {
		t.Fatalf("got %+v", got)
	}
}
