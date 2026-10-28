package auth

import (
	"context"
	"errors"
	"net/http"
	"testing"
)

// fakeGuildStanding answers GuildStandingSource from a table keyed by guild, or fails.
type fakeGuildStanding struct {
	byGuild map[int64]GuildStanding
	err     error
}

func (f fakeGuildStanding) GuildStanding(_ context.Context, _ int64, _ []int64) (map[int64]GuildStanding, error) {
	return f.byGuild, f.err
}

type standingMeBody struct {
	Guilds []struct {
		RulesetLabel *string `json:"ruleset_label"`
		Slug         *string `json:"slug"`
		CrestURL     *string `json:"crest_url"`
		Faction      *string `json:"faction"`
		RosterCount  *int    `json:"roster_count"`
		Claim        *struct {
			State        string  `json:"state"`
			Yours        bool    `json:"yours"`
			PendingYours bool    `json:"pending_yours"`
			HolderName   *string `json:"holder_name"`
		} `json:"claim"`
		Verify *struct {
			HasApprover bool `json:"has_approver"`
		} `json:"verify"`
	} `json:"guilds"`
}

func readStandingMe(t *testing.T, h *harness, user int64) standingMeBody {
	t.Helper()
	res := h.sessionRequest(t, http.MethodGet, "/v1/me", user)
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", res.StatusCode)
	}
	var b standingMeBody
	h.decode(t, res, &b)
	return b
}

// TestMe_GuildsCarryIdentity: with a Standing source every guild entry carries ruleset_label,
// slug, crest_url, faction and roster_count; with none, or a source that fails, /v1/me still
// answers 200 and the string facts are absent.
func TestMe_GuildsCarryIdentity(t *testing.T) {
	h := newHarness(t)
	gid := h.seedGuild(t, "me-guild-identity")
	user := h.seedUser(t, "me-identity@example.com")
	h.seedGuildMember(t, gid, user, "member", true)

	assertOlder := func(label string) {
		t.Helper()
		b := readStandingMe(t, h, user)
		if len(b.Guilds) != 1 || b.Guilds[0].RulesetLabel != nil || b.Guilds[0].Slug != nil || b.Guilds[0].CrestURL != nil ||
			b.Guilds[0].Faction != nil || b.Guilds[0].RosterCount != nil {
			t.Fatalf("%s: guilds = %+v, want no identity facts", label, b.Guilds)
		}
	}
	assertOlder("no Standing wired")

	h.svc.Standing = fakeGuildStanding{err: errors.New("standing read failed")}
	assertOlder("failing Standing")

	crest, faction, roster, holder := "https://api.example.test/crest.webp?v=abc", "horde", 38, "Holder"
	h.svc.Standing = fakeGuildStanding{byGuild: map[int64]GuildStanding{gid: {
		RulesetLabel: "Hardcore", Slug: "me-guild-identity", CrestURL: &crest, Faction: &faction, RosterCount: &roster,
		Claim:  &GuildClaimView{State: "claimed", Yours: true, PendingYours: true, HolderName: &holder},
		Verify: &GuildVerifyView{HasApprover: true},
	}}}
	b := readStandingMe(t, h, user)
	if len(b.Guilds) != 1 {
		t.Fatalf("guilds = %+v", b.Guilds)
	}
	g := b.Guilds[0]
	if g.RulesetLabel == nil || *g.RulesetLabel != "Hardcore" || g.Slug == nil || *g.Slug != "me-guild-identity" ||
		g.CrestURL == nil || *g.CrestURL != crest || g.Faction == nil || *g.Faction != "horde" ||
		g.RosterCount == nil || *g.RosterCount != 38 {
		t.Fatalf("guild = %+v, want the standing's identity facts", g)
	}
	if g.Claim == nil || g.Claim.State != "claimed" || !g.Claim.Yours || !g.Claim.PendingYours ||
		g.Claim.HolderName == nil || *g.Claim.HolderName != "Holder" {
		t.Errorf("claim = %+v, want state claimed, yours, pending_yours, holder_name Holder", g.Claim)
	}
	if g.Verify == nil || !g.Verify.HasApprover {
		t.Errorf("verify = %+v, want has_approver true", g.Verify)
	}
}
