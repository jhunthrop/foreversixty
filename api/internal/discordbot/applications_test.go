package discordbot

import (
	"context"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

func TestGrantGuildRoleNeedsAnInstallAndALink(t *testing.T) {
	ctx := context.Background()
	st := testStore(t)
	rest := &FakeREST{}
	svc := New(Deps{Store: st, REST: rest})
	officer, applicant := seedUser(t, st.Pool), seedUser(t, st.Pool)
	g := seedGuild(t, st.Pool, "Recruiters")

	if got, err := svc.GrantGuildRole(ctx, g, applicant, "123456789012345678"); err != nil || got != guilds.OutcomeNoBot {
		t.Fatalf("no install: %q, %v", got, err)
	}
	if _, err := st.BindServer(ctx, g, "srv1", "S", officer); err != nil {
		t.Fatal(err)
	}
	if got, err := svc.GrantGuildRole(ctx, g, applicant, "123456789012345678"); err != nil || got != guilds.OutcomeNotLinked {
		t.Fatalf("no link: %q, %v", got, err)
	}
	if err := st.SaveLink(ctx, applicant, "998877", "Ash"); err != nil {
		t.Fatal(err)
	}
	if got, err := svc.GrantGuildRole(ctx, g, applicant, "123456789012345678"); err != nil || got != guilds.OutcomeGranted {
		t.Fatalf("linked: %q, %v", got, err)
	}
	calls := rest.Calls()
	last := calls[len(calls)-1]
	if last.Method != "AddMemberRole" || last.ServerID != "srv1" || last.UserID != "998877" || last.RoleID != "123456789012345678" {
		t.Fatalf("call = %+v", last)
	}
	rest.ErrFor = map[string]error{"AddMemberRole": &APIError{Status: 403, Code: 50013, Message: "Missing Permissions"}}
	if got, err := svc.GrantGuildRole(ctx, g, applicant, "123456789012345678"); err != nil || got != guilds.OutcomeDiscordRefuse {
		t.Fatalf("refused: %q, %v", got, err)
	}
}

func TestDirectMessageOpensADMAndSendsTheText(t *testing.T) {
	ctx := context.Background()
	st := testStore(t)
	rest := &FakeREST{}
	svc := New(Deps{Store: st, REST: rest})
	applicant := seedUser(t, st.Pool)

	if got, err := svc.DirectMessage(ctx, applicant, "hello"); err != nil || got != guilds.OutcomeNotLinked {
		t.Fatalf("no link: %q, %v", got, err)
	}
	if err := st.SaveLink(ctx, applicant, "998877", "Ash"); err != nil {
		t.Fatal(err)
	}
	if got, err := svc.DirectMessage(ctx, applicant, "hello"); err != nil || got != guilds.OutcomeSent {
		t.Fatalf("linked: %q, %v", got, err)
	}
	calls := rest.Calls()
	if len(calls) != 2 || calls[0].Method != "OpenDM" || calls[1].ChannelID != "dm-998877" || calls[1].Message.Content != "hello" {
		t.Fatalf("calls = %+v", calls)
	}
}
