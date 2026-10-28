// api/internal/guilds/me_standing_test.go
package guilds

import (
	"context"
	"strings"
	"testing"
	"time"
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
	if st.Realm != "Hardcore" {
		t.Errorf("Realm = %q, want Hardcore", st.Realm)
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
	if st.Realm != "Hardcore" {
		t.Errorf("Realm = %q, want Hardcore", st.Realm)
	}
}

func TestGuildStanding_OneQueryForManyGuilds(t *testing.T) {
	h := newHTTPHarness(t)
	user := seedUser(t, h.pool, "many@example.com")
	ids := []int64{seedGuild(t, h.pool, "One"), seedGuild(t, h.pool, "Two"), seedGuild(t, h.pool, "Three")}

	got, err := h.store.GuildStanding(context.Background(), user, ids)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 3 {
		t.Fatalf("got %d entries, want 3: %+v", len(got), got)
	}
	for _, id := range ids {
		if _, ok := got[id]; !ok {
			t.Errorf("missing guild %d", id)
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
