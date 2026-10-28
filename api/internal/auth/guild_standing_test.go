package auth

import (
	"context"
	"net/http"
	"testing"
)

// fakeGuildStanding answers GuildStandingSource from a table keyed by guild.
type fakeGuildStanding map[int64]GuildStanding

func (f fakeGuildStanding) GuildStanding(_ context.Context, _ int64, _ []int64) (map[int64]GuildStanding, error) {
	return f, nil
}

// TestMe_GuildsCarryIdentity: with a Standing source every guild entry carries realm,
// crest_url, faction and roster_count; with none, /v1/me still answers 200 and the entry
// carries no identity facts.
func TestMe_GuildsCarryIdentity(t *testing.T) {
	h := newHarness(t)
	gid := h.seedGuild(t, "me-guild-identity")
	user := h.seedUser(t, "me-identity@example.com")
	h.seedGuildMember(t, gid, user, "member", true)

	type meBody struct {
		Guilds []struct {
			Realm       string  `json:"realm"`
			CrestURL    *string `json:"crest_url"`
			Faction     *string `json:"faction"`
			RosterCount *int    `json:"roster_count"`
		} `json:"guilds"`
	}
	read := func() meBody {
		res := h.sessionRequest(t, http.MethodGet, "/v1/me", user)
		if res.StatusCode != http.StatusOK {
			t.Fatalf("status = %d, want 200", res.StatusCode)
		}
		var b meBody
		h.decode(t, res, &b)
		return b
	}

	if b := read(); len(b.Guilds) != 1 || b.Guilds[0].Realm != "" || b.Guilds[0].CrestURL != nil ||
		b.Guilds[0].Faction != nil || b.Guilds[0].RosterCount != nil {
		t.Fatalf("with no Standing wired, guilds = %+v, want no identity facts", b.Guilds)
	}

	crest, faction, roster := "https://api.example.test/crest.webp?v=abc", "horde", 38
	h.svc.Standing = fakeGuildStanding{gid: {Realm: "Hardcore", CrestURL: &crest, Faction: &faction, RosterCount: &roster}}
	b := read()
	if len(b.Guilds) != 1 || b.Guilds[0].Realm != "Hardcore" || b.Guilds[0].CrestURL == nil || *b.Guilds[0].CrestURL != crest ||
		b.Guilds[0].Faction == nil || *b.Guilds[0].Faction != "horde" || b.Guilds[0].RosterCount == nil || *b.Guilds[0].RosterCount != 38 {
		t.Fatalf("guilds = %+v, want the standing's identity facts", b.Guilds)
	}
}
