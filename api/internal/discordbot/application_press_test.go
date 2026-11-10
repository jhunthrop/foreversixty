package discordbot

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

func applicationPressOf(customID, discordUser string) Interaction {
	data, _ := json.Marshal(map[string]string{"custom_id": customID})
	return Interaction{Type: typeComponent, GuildID: "srv1", DiscordUserID: discordUser, Data: data}
}

func TestApplicationButtonsCarryTheID(t *testing.T) {
	row := applicationButtons(42, "https://site.test/apps")
	if len(row) != 3 || row[0].CustomID != "gapp:accept:42" || row[1].CustomID != "gapp:decline:42" ||
		row[2].URL != "https://site.test/apps" {
		t.Fatalf("row = %+v", row)
	}
	for _, c := range []struct {
		id     string
		accept bool
		ok     bool
	}{{"gapp:accept:42", true, true}, {"gapp:decline:7", false, true}, {"gapp:maybe:7", false, false}, {"gapp:accept:x", false, false}} {
		if _, accept, ok := parseApplicationID(c.id); ok != c.ok || accept != c.accept {
			t.Errorf("%s: accept %v ok %v", c.id, accept, ok)
		}
	}
}

// A linked officer accepts from Discord through the site's Decide; anyone else is refused.
func TestApplicationPressDecidesForALinkedOfficer(t *testing.T) {
	h := newInstallHarness(t)
	h.callback(h.gid, "code=c", h.officerActor())
	ctx := context.Background()

	rep, err := h.svc.applicationPress(ctx, applicationPressOf("gapp:accept:9", "d-unlinked"))
	if err != nil || rep.Message.Content != applicationCopy.Unlinked {
		t.Fatalf("unlinked: %q %v", rep.Message.Content, err)
	}
	if err := h.svc.d.Store.SaveLink(ctx, h.uid, "d-officer", "Thornhide"); err != nil {
		t.Fatal(err)
	}
	rep, err = h.svc.applicationPress(ctx, applicationPressOf("gapp:accept:9", "d-officer"))
	if err != nil || rep.Message.Content != "Accepted Mirelle." || !rep.Message.Ephemeral {
		t.Fatalf("accept: %+v %v", rep.Message, err)
	}
	if len(h.guilds.decided) != 1 || h.guilds.decided[0] != (fakeDecision{h.gid, h.uid, 9, true}) {
		t.Fatalf("decided = %+v", h.guilds.decided)
	}
	h.guilds.decideErr = guilds.ErrApplicationClosed
	if rep, _ := h.svc.applicationPress(ctx, applicationPressOf("gapp:decline:9", "d-officer")); rep.Message.Content != applicationCopy.Closed {
		t.Fatalf("closed: %q", rep.Message.Content)
	}
	h.setStanding(func(s *notify.GuildStanding) { s.Officer = false })
	if rep, _ := h.svc.applicationPress(ctx, applicationPressOf("gapp:decline:9", "d-officer")); rep.Message.Content != applicationCopy.OfficersOnly {
		t.Fatalf("member: %q", rep.Message.Content)
	}
	if len(h.guilds.decided) != 2 {
		t.Fatalf("a member decided: %+v", h.guilds.decided)
	}
}
