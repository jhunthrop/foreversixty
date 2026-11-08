package guilds

import (
	"net/http"
	"slices"
	"testing"
)

// A ping with a role mentions only the not-signed raiders of that role; an unknown role is 422.
func TestPingMentionsOnlyTheRoleAsked(t *testing.T) {
	w, raid, pinger := pingWorld(t)
	if _, err := w.h.pool.Exec(t.Context(), `update guild_characters set raid_spec = 'priest-holy' where character_key = $1`,
		member2Key); err != nil {
		t.Fatal(err)
	}
	w.as(w.officer)
	if res := w.h.do(http.MethodPost, w.path("/raids/%d/ping", raid.ID), `{"role":"mage"}`); res.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("unknown role = %d", res.StatusCode)
	}
	res := w.h.do(http.MethodPost, w.path("/raids/%d/ping", raid.ID), `{"role":"healer"}`)
	if res.StatusCode != http.StatusOK {
		t.Fatalf("healer ping = %d: %s", res.StatusCode, readBody(res))
	}
	var out PingResult
	w.h.data(res, &out)
	if !out.Posted || out.Mentioned != 1 || len(pinger.pings) != 1 || !slices.Equal(pinger.pings[0].MentionIDs, []string{member2Discord}) {
		t.Fatalf("result %+v pings %+v", out, pinger.pings)
	}
}

// A ping with no body mentions every linked not-signed raider, as before.
func TestPingWithNoRoleMentionsEveryone(t *testing.T) {
	w, raid, pinger := pingWorld(t)
	w.as(w.officer)
	if res := w.h.do(http.MethodPost, w.path("/raids/%d/ping", raid.ID), ""); res.StatusCode != http.StatusOK {
		t.Fatalf("ping = %d: %s", res.StatusCode, readBody(res))
	}
	if len(pinger.pings) != 1 || len(pinger.pings[0].MentionIDs) != 2 {
		t.Fatalf("pings %+v", pinger.pings)
	}
}

// The DM link shows only for a member who opted into officer DMs; never a raw id otherwise.
func TestNotSignedDMLinkNeedsTheOptIn(t *testing.T) {
	w, raid, _ := pingWorld(t)
	if _, err := w.h.pool.Exec(t.Context(), `insert into notification_settings (user_id, time_zone, officer_dms)
		values ($1, 'UTC', true)`, w.member2); err != nil {
		t.Fatal(err)
	}
	w.as(w.officer)
	got := map[string]*string{}
	for _, r := range *w.getRaid(t, raid.ID).NotSignedRaiders {
		got[r.Name] = r.DiscordDMURL
	}
	if u := got["Mirelle"]; u == nil || *u != "https://discord.com/users/"+member2Discord {
		t.Fatalf("Mirelle's link %v", u)
	}
	if got["Thornhide"] != nil || got["Kraggor"] != nil {
		t.Fatalf("a link without the opt-in: %+v", got)
	}
}
