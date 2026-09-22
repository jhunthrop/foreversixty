// api/internal/bnetimport/refresh_test.go
package bnetimport

import (
	"context"
	"net/http"
	"testing"
)

func TestStaleBnetCharactersFindsOnlyOldBnetRows(t *testing.T) {
	pool := testPool(t)
	uid := seedUser(t, pool)
	ctx := context.Background()
	if _, err := pool.Exec(ctx,
		`insert into characters (key, region, ruleset, name, user_id, realm_slug, bnet_character_id, source, refreshed_at)
		 values ('us/pvp/stale', 'us', 'pvp', 'Stale', $1, 'whitemane', 101, 'bnet', now() - interval '25 hours')`, uid); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into characters (key, region, ruleset, name, user_id, realm_slug, bnet_character_id, source, refreshed_at)
		 values ('us/pvp/fresh', 'us', 'pvp', 'Fresh', $1, 'whitemane', 102, 'bnet', now())`, uid); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into characters (key, region, ruleset, name, user_id, source, refreshed_at)
		 values ('us/pvp/exported', 'us', 'pvp', 'Exported', $1, 'export', now() - interval '25 hours')`, uid); err != nil {
		t.Fatal(err)
	}

	svc := &Service{Pool: pool}
	stale, err := svc.staleBnetCharacters(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(stale) != 1 || stale[0].BnetCharacterID != 101 || stale[0].RealmSlug != "whitemane" {
		t.Fatalf("stale = %+v, want just bnet_character_id=101/whitemane", stale)
	}
}

func TestRunRefreshStopsOnRateLimit(t *testing.T) {
	pool := testPool(t)
	uid := seedUser(t, pool)
	ctx := context.Background()
	// "one" is staler than "two", so staleBnetCharacters (ordered oldest
	// first) processes it first — its 429 must stop the run before "two"
	// (which would otherwise succeed) is ever reached.
	ages := map[string]string{"us/pvp/one": "26 hours", "us/pvp/two": "25 hours"}
	names := map[string]string{"us/pvp/one": "One", "us/pvp/two": "Two"}
	bnetIDs := map[string]int{"us/pvp/one": 201, "us/pvp/two": 202}
	for _, key := range []string{"us/pvp/one", "us/pvp/two"} {
		if _, err := pool.Exec(ctx,
			`insert into characters (key, region, ruleset, name, user_id, realm_slug, bnet_character_id, source, refreshed_at)
			 values ($1, 'us', 'pvp', $2, $3, 'whitemane', $4, 'bnet', now() - $5::interval)`,
			key, names[key], uid, bnetIDs[key], ages[key]); err != nil {
			t.Fatal(err)
		}
	}

	f := newBlizzardFixture(t)
	f.json(http.MethodGet, "/profile/wow/character/whitemane/one?namespace=profile-classic1x-us",
		http.StatusTooManyRequests, nil)
	f.json(http.MethodGet, "/profile/wow/character/whitemane/two?namespace=profile-classic1x-us",
		http.StatusOK, map[string]any{"name": "Two", "level": 60, "faction": map[string]string{"type": "HORDE"},
			"character_class": map[string]string{"name": "Rogue"}, "realm": map[string]string{"slug": "whitemane"}})
	f.json(http.MethodGet, "/profile/wow/character/whitemane/two/equipment?namespace=profile-classic1x-us",
		http.StatusOK, map[string]any{"equipped_items": []map[string]any{}})

	svc := newTestService(t, pool, f)
	result, err := svc.RunRefresh(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !result.RateLimited {
		t.Fatalf("result = %+v, want RateLimited", result)
	}
	if result.Refreshed != 0 {
		t.Fatalf("result.Refreshed = %d, want 0 (the run stops on the first rate limit)", result.Refreshed)
	}
}

func TestRunRefreshClearsAWithdrawnBnetMembershipButLeavesExportSourcedRowsAlone(t *testing.T) {
	pool := testPool(t)
	uid := seedUser(t, pool)
	ctx := context.Background()
	if _, err := pool.Exec(ctx,
		`insert into characters (key, region, ruleset, name, user_id, realm_slug, bnet_character_id, source, refreshed_at)
		 values ('us/pvp/left', 'us', 'pvp', 'Left', $1, 'whitemane', 301, 'bnet', now() - interval '25 hours')`, uid); err != nil {
		t.Fatal(err)
	}
	var gid int64
	if err := pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'pvp', 'Old Guild') returning id`).Scan(&gid); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, rank, source, verified_by, verified_at)
		 values ($1, 'us/pvp/left', $2, 'member', 'bnet', 'bnet', now())`, gid, uid); err != nil {
		t.Fatal(err)
	}

	f := newBlizzardFixture(t)
	f.json(http.MethodGet, "/profile/wow/character/whitemane/left?namespace=profile-classic1x-us", http.StatusOK,
		map[string]any{"name": "Left", "level": 60, "faction": map[string]string{"type": "HORDE"},
			"character_class": map[string]string{"name": "Rogue"}, "realm": map[string]string{"slug": "whitemane"}})
	f.json(http.MethodGet, "/profile/wow/character/whitemane/left/equipment?namespace=profile-classic1x-us",
		http.StatusOK, map[string]any{"equipped_items": []map[string]any{}})

	svc := newTestService(t, pool, f)
	if _, err := svc.RunRefresh(ctx, nil); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := pool.QueryRow(ctx, `select count(*) from guild_characters where character_key = 'us/pvp/left'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("guild_characters rows = %d, want 0 (the character left the guild)", n)
	}
}
