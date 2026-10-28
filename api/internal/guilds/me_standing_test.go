// api/internal/guilds/me_standing_test.go
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

func TestGuildStanding_IdentityFacts(t *testing.T) {
	h := newHTTPHarness(t)
	h.store.APIBaseURL = "https://api.example.test"
	gid, reader := addonGuild(t, h)
	if _, err := h.pool.Exec(context.Background(),
		`update guilds set faction = 'horde', crest_key = $2 where id = $1`,
		gid, "guilds/1/crest/0123456789abcdef.webp"); err != nil {
		t.Fatal(err)
	}
	// Three accounts on the roster plus two names no account has claimed: on_roster 5,
	// in_game_no_account 2, so the card reads 7.
	seedRosterSnapshot(t, h.pool, gid, "us/hardcore/Thorgrim", `[
	  {"name":"Thorgrim-Yell","level":52,"spec":"warrior-arms","gear_gap":3,"build":"","seen_at":"2026-11-04 14:01"},
	  {"name":"Stranger-Yell","level":60,"spec":"mage-frost","gear_gap":0,"build":"","seen_at":"2026-11-04 14:00"},
	  {"name":"Drifter-Yell","level":55,"spec":"rogue-combat","gear_gap":0,"build":"","seen_at":"2026-11-04 14:00"}]`,
		time.Now().UTC())

	got, err := h.store.GuildStanding(context.Background(), reader, []int64{gid})
	if err != nil {
		t.Fatal(err)
	}
	st, ok := got[gid]
	if !ok {
		t.Fatalf("no standing for guild %d: %+v", gid, got)
	}
	if st.RulesetLabel != "Hardcore" {
		t.Errorf("RulesetLabel = %q, want Hardcore", st.RulesetLabel)
	}
	if st.Slug != "forever" {
		t.Errorf("Slug = %q, want forever", st.Slug)
	}
	if st.CrestURL == nil || !strings.HasPrefix(*st.CrestURL, "https://api.example.test/v1/guilds/") ||
		!strings.Contains(*st.CrestURL, "?v=0123456789abcdef") {
		t.Errorf("CrestURL = %v, want a versioned crest url", st.CrestURL)
	}
	if st.Faction == nil || *st.Faction != "horde" {
		t.Errorf("Faction = %v, want horde", st.Faction)
	}
	if st.RosterCount == nil || *st.RosterCount != 7 {
		t.Errorf("RosterCount = %v, want 7", st.RosterCount)
	}
}

func TestGuildStanding_NoSnapshotNoCrest(t *testing.T) {
	h := newHTTPHarness(t)
	gid, reader := addonGuild(t, h)

	got, err := h.store.GuildStanding(context.Background(), reader, []int64{gid})
	if err != nil {
		t.Fatal(err)
	}
	st := got[gid]
	if st.RosterCount != nil || st.CrestURL != nil || st.Faction != nil {
		t.Errorf("standing = %+v, want nil roster count, crest and faction", st)
	}
	if st.RulesetLabel != "Hardcore" {
		t.Errorf("RulesetLabel = %q, want Hardcore", st.RulesetLabel)
	}
}

// TestGuildStanding_ManyGuildsEachItsOwnFacts: three guilds on three rulesets answer their
// own label and slug, and a guild id that does not exist is simply absent.
func TestGuildStanding_ManyGuildsEachItsOwnFacts(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	user := seedUser(t, h.pool, "many@example.com")
	want := map[int64][2]string{}
	for _, g := range []struct{ name, ruleset, label, slug string }{
		{"Alpha Wolves", "hardcore", "Hardcore", "alpha-wolves"},
		{"Beta Team", "pvp", "PvP", "beta-team"},
		{"Gamma", "rp", "Roleplay", "gamma"},
	} {
		var id int64
		if err := h.pool.QueryRow(ctx,
			`insert into guilds (region, ruleset, name) values ('us', $1, $2) returning id`, g.ruleset, g.name).Scan(&id); err != nil {
			t.Fatal(err)
		}
		want[id] = [2]string{g.label, g.slug}
	}
	ids := []int64{}
	for id := range want {
		ids = append(ids, id)
	}
	const missingID = int64(999999999)

	got, err := h.store.GuildStanding(ctx, user, append(ids, missingID))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 3 {
		t.Fatalf("got %d entries, want 3: %+v", len(got), got)
	}
	if _, ok := got[missingID]; ok {
		t.Errorf("a nonexistent guild must be absent: %+v", got[missingID])
	}
	for id, w := range want {
		if got[id].RulesetLabel != w[0] || got[id].Slug != w[1] {
			t.Errorf("guild %d = %q/%q, want %q/%q", id, got[id].RulesetLabel, got[id].Slug, w[0], w[1])
		}
	}
}

func TestGuildStanding_NoGuildsAnswersEmpty(t *testing.T) {
	h := newHTTPHarness(t)
	got, err := h.store.GuildStanding(context.Background(), 1, nil)
	if err != nil || len(got) != 0 {
		t.Fatalf("got %+v, %v; want an empty map", got, err)
	}
}

// TestGuildStanding_RosterCountMatchesHome: the count equals the home's on_roster plus
// in_game_no_account for a guild with a waiting account, an invitee with a character, an
// invitee with none (named by its public name, which the snapshot also names) and two names
// no account has claimed.
func TestGuildStanding_RosterCountMatchesHome(t *testing.T) {
	h := newHTTPHarness(t)
	h.store.Accounts = h.svc.Accounts
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")

	master := seedUser(t, h.pool, "std-master@example.com")
	seedCharacter(t, h.pool, gid, master, "us/hardcore/stdmaster", "leader", true)
	seedExport(t, h.pool, master, "us/hardcore/stdmaster", "us", "hardcore", "Stdmaster")
	recomputeMembership(t, h.pool, gid, master)
	pending := seedUser(t, h.pool, "std-pending@example.com")
	seedCharacter(t, h.pool, gid, pending, "us/hardcore/stdpending", "member", false)
	seedExport(t, h.pool, pending, "us/hardcore/stdpending", "us", "hardcore", "Stdpending")
	recomputeMembership(t, h.pool, gid, pending)

	withChar := seedUser(t, h.pool, "std-invitee@example.com")
	seedCharacter(t, h.pool, gid, withChar, syntheticKey(withChar), "member", true)
	recomputeMembership(t, h.pool, gid, withChar)
	if _, err := h.pool.Exec(ctx, `insert into characters (key, user_id, region, ruleset, name, class)
		values ('us/hardcore/driftwood', $1, 'us', 'hardcore', 'Driftwood', 'rogue')`, withChar); err != nil {
		t.Fatal(err)
	}
	bare := seedUser(t, h.pool, "std-bare@example.com")
	seedCharacter(t, h.pool, gid, bare, syntheticKey(bare), "member", true)
	recomputeMembership(t, h.pool, gid, bare)
	u, err := h.svc.Accounts.User(ctx, bare)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := h.pool.Exec(ctx, `insert into characters (key, user_id, region, ruleset, name)
		values ('us/hardcore/stdmaster', $1, 'us', 'hardcore', 'Stdmaster') on conflict do nothing`, master); err != nil {
		t.Fatal(err)
	}
	putSnapshot(t, h, gid, "us/hardcore/stdmaster", fmt.Sprintf(`[
		{"name":"Stdmaster","level":60,"spec":"hunter-beast-mastery","gear_gap":0,"seen_at":"2026-10-20 18:00"},
		{"name":%q,"level":60,"spec":"mage-frost","gear_gap":0,"seen_at":"2026-10-20 18:01"},
		{"name":"Strayhound","level":60,"spec":"warrior-protection","gear_gap":0,"seen_at":"2026-10-20 17:00"},
		{"name":"Oldmarrow","level":58,"spec":"warrior-arms","gear_gap":0,"seen_at":"2026-10-19 21:00"}]`, u.PublicName()))

	h.actor = auth.Actor{UserID: master, Role: "user", Method: "session"}
	var view HomeView
	h.data(h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), ""), &view)
	want := view.Counts.OnRoster + view.Counts.InGameNoAccount

	got, err := h.store.GuildStanding(ctx, master, []int64{gid})
	if err != nil {
		t.Fatal(err)
	}
	if got[gid].RosterCount == nil || *got[gid].RosterCount != want {
		t.Fatalf("RosterCount = %v, want the home's %d (counts %+v)", got[gid].RosterCount, want, view.Counts)
	}
}
