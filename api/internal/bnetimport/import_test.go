// api/internal/bnetimport/import_test.go
package bnetimport

import (
	"context"
	"net/http"
	"testing"
)

const accountCharactersPath = "/profile/user/wow?namespace=profile-classic1x-us"

func TestImportAccountWritesAGuildedCharacterAndItsMembership(t *testing.T) {
	pool := testPool(t)
	uid := seedUser(t, pool)
	f := newBlizzardFixture(t)
	f.handlers[http.MethodGet+" "+accountCharactersPath] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"wow_accounts":[{"id":1,"characters":[
			{"name":"Thoradin","id":501,"realm":{"slug":"whitemane","name":"Whitemane"},
			 "playable_class":{"name":"Warrior"},"faction":{"type":"ALLIANCE"},"level":60}
		]}]}`))
	}
	f.json(http.MethodGet, "/data/wow/realm/index?namespace=dynamic-classic1x-us", http.StatusOK,
		map[string]any{"realms": []map[string]any{{"id": 1, "name": "Whitemane", "slug": "whitemane"}}})
	f.json(http.MethodGet, "/data/wow/realm/whitemane?namespace=dynamic-classic1x-us", http.StatusOK,
		map[string]any{"type": map[string]string{"type": "PVP"}, "category": "PvP"})
	f.handlers[http.MethodGet+" /profile/wow/character/whitemane/thoradin?namespace=profile-classic1x-us"] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"name":"Thoradin","level":60,"faction":{"type":"ALLIANCE"},
			"character_class":{"name":"Warrior"},"realm":{"slug":"whitemane","name":"Whitemane"},
			"guild":{"name":"Iron Vanguard","id":12}}`))
	}
	f.handlers[http.MethodGet+" /data/wow/guild/whitemane/iron-vanguard/roster?namespace=profile-classic1x-us"] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"guild":{"name":"Iron Vanguard","id":12,"realm":{"slug":"whitemane"}},
			"members":[{"character":{"name":"Thoradin","realm":{"slug":"whitemane"},"level":60,
			"playable_class":{"id":1}},"rank":1}]}`))
	}

	svc := newTestService(t, pool, f)
	summary, err := svc.ImportAccount(context.Background(), uid, "user-oauth-token")
	if err != nil {
		t.Fatal(err)
	}
	if summary.Characters != 1 || summary.Guilds != 1 || summary.Skipped != 0 {
		t.Fatalf("summary = %+v", summary)
	}

	var region, ruleset, source string
	var level int
	if err := pool.QueryRow(context.Background(),
		`select region, ruleset, level, source from characters where key = 'us/pvp/thoradin'`).
		Scan(&region, &ruleset, &level, &source); err != nil {
		t.Fatal(err)
	}
	if region != "us" || ruleset != "pvp" || level != 60 || source != "bnet" {
		t.Fatalf("character row = %s/%s level=%d source=%s", region, ruleset, level, source)
	}

	var rank, verifiedBy string
	var verified bool
	if err := pool.QueryRow(context.Background(),
		`select rank, verified_by, verified_at is not null from guild_characters where character_key = 'us/pvp/thoradin'`).
		Scan(&rank, &verifiedBy, &verified); err != nil {
		t.Fatal(err)
	}
	if rank != "officer" || verifiedBy != "bnet" || !verified {
		t.Fatalf("membership = rank=%s verified_by=%s verified=%v", rank, verifiedBy, verified)
	}

	var importedAt *string
	if err := pool.QueryRow(context.Background(),
		`select bnet_imported_at::text from users where id = $1`, uid).Scan(&importedAt); err != nil {
		t.Fatal(err)
	}
	if importedAt == nil {
		t.Fatal("users.bnet_imported_at was not stamped")
	}
}

func TestImportAccountSkipsCharactersBelowMinLevel(t *testing.T) {
	pool := testPool(t)
	uid := seedUser(t, pool)
	f := newBlizzardFixture(t)
	f.handlers[http.MethodGet+" "+accountCharactersPath] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"wow_accounts":[{"id":1,"characters":[
			{"name":"Bankalt","id":9,"realm":{"slug":"whitemane","name":"Whitemane"},
			 "playable_class":{"name":"Mage"},"faction":{"type":"HORDE"},"level":5}
		]}]}`))
	}
	f.json(http.MethodGet, "/data/wow/realm/index?namespace=dynamic-classic1x-us", http.StatusOK,
		map[string]any{"realms": []any{}})

	svc := newTestService(t, pool, f)
	summary, err := svc.ImportAccount(context.Background(), uid, "user-oauth-token")
	if err != nil {
		t.Fatal(err)
	}
	if summary.Characters != 0 || summary.Skipped != 1 {
		t.Fatalf("summary = %+v, want 0 characters, 1 skipped", summary)
	}
	var n int
	if err := pool.QueryRow(context.Background(), `select count(*) from characters`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("characters rows = %d, want 0", n)
	}
}

func TestImportAccountNeverStealsACharacterKeyFromAnotherAccount(t *testing.T) {
	pool := testPool(t)
	owner := seedUser(t, pool)
	importer := seedUser(t, pool)
	if _, err := pool.Exec(context.Background(),
		`insert into characters (key, region, ruleset, name, user_id) values ('us/pvp/thoradin', 'us', 'pvp', 'Thoradin', $1)`,
		owner); err != nil {
		t.Fatal(err)
	}

	f := newBlizzardFixture(t)
	f.handlers[http.MethodGet+" "+accountCharactersPath] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"wow_accounts":[{"id":1,"characters":[
			{"name":"Thoradin","id":501,"realm":{"slug":"whitemane","name":"Whitemane"},
			 "playable_class":{"name":"Warrior"},"faction":{"type":"ALLIANCE"},"level":60}
		]}]}`))
	}
	f.json(http.MethodGet, "/data/wow/realm/index?namespace=dynamic-classic1x-us", http.StatusOK,
		map[string]any{"realms": []map[string]any{{"id": 1, "name": "Whitemane", "slug": "whitemane"}}})
	f.json(http.MethodGet, "/data/wow/realm/whitemane?namespace=dynamic-classic1x-us", http.StatusOK,
		map[string]any{"type": map[string]string{"type": "PVP"}})

	svc := newTestService(t, pool, f)
	summary, err := svc.ImportAccount(context.Background(), importer, "user-oauth-token")
	if err != nil {
		t.Fatal(err)
	}
	if summary.Characters != 0 || summary.Skipped != 1 {
		t.Fatalf("summary = %+v, want the collision skipped, nothing written", summary)
	}
	var stillOwner int64
	if err := pool.QueryRow(context.Background(),
		`select user_id from characters where key = 'us/pvp/thoradin'`).Scan(&stillOwner); err != nil {
		t.Fatal(err)
	}
	if stillOwner != owner {
		t.Fatalf("owner = %d, want the original owner %d (never stolen)", stillOwner, owner)
	}
}

func TestImportAccountStampsBnetImportedAtEvenWithNoAccountInAnyRegion(t *testing.T) {
	pool := testPool(t)
	uid := seedUser(t, pool)
	f := newBlizzardFixture(t)
	f.json(http.MethodGet, accountCharactersPath, http.StatusForbidden, nil)

	svc := newTestService(t, pool, f)
	summary, err := svc.ImportAccount(context.Background(), uid, "user-oauth-token")
	if err != nil {
		t.Fatal(err)
	}
	if len(summary.Regions) != 0 {
		t.Fatalf("regions = %v, want none (403 everywhere)", summary.Regions)
	}
	var importedAt *string
	if err := pool.QueryRow(context.Background(),
		`select bnet_imported_at::text from users where id = $1`, uid).Scan(&importedAt); err != nil {
		t.Fatal(err)
	}
	if importedAt == nil {
		t.Fatal("bnet_imported_at should still be stamped — the login tried")
	}
}
