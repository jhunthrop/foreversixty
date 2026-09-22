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
			 "playable_class":{"name":"Warrior"},"playable_race":{"name":"Human"},
			 "gender":{"type":"MALE"},"faction":{"type":"ALLIANCE"},"level":60}
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
			"guild":{"name":"Iron Vanguard","id":12},"average_item_level":40,"equipped_item_level":38,
			"last_login_timestamp":1700000000000}`))
	}
	f.handlers[http.MethodGet+" /data/wow/guild/whitemane/iron-vanguard/roster?namespace=profile-classic1x-us"] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"guild":{"name":"Iron Vanguard","id":12,"realm":{"slug":"whitemane"}},
			"members":[{"character":{"name":"Thoradin","realm":{"slug":"whitemane"},"level":60,
			"playable_class":{"id":1}},"rank":1}]}`))
	}
	f.json(http.MethodGet, "/profile/wow/character/whitemane/thoradin/equipment?namespace=profile-classic1x-us",
		http.StatusOK, map[string]any{"equipped_items": []map[string]any{}})
	f.json(http.MethodGet, "/profile/wow/character/whitemane/thoradin/character-media?namespace=profile-classic1x-us",
		http.StatusOK, map[string]any{"assets": []map[string]string{
			{"key": "avatar", "value": "https://render.worldofwarcraft.com/us/character/whitemane/1-thoradin-avatar.jpg"},
			{"key": "main-raw", "value": "https://render.worldofwarcraft.com/character/whitemane/1-thoradin-main-raw.png"},
		}})

	svc := newTestService(t, pool, f)
	summary, err := svc.ImportAccount(context.Background(), uid, "user-oauth-token")
	if err != nil {
		t.Fatal(err)
	}
	if summary.Characters != 1 || summary.Guilds != 1 || summary.Skipped != 0 || summary.Unavailable != 0 {
		t.Fatalf("summary = %+v", summary)
	}

	var region, ruleset, source, race, gender string
	var level, avgIL, equipIL int
	var bnetAccount, bnetProfile, bnetEquipment, bnetMedia []byte
	var avatarURL, renderURL *string
	if err := pool.QueryRow(context.Background(),
		`select region, ruleset, level, source, race, gender, average_item_level, equipped_item_level,
		        bnet_account, bnet_profile, bnet_equipment, avatar_url, render_url, bnet_media
		 from characters where key = 'us/pvp/thoradin'`).
		Scan(&region, &ruleset, &level, &source, &race, &gender, &avgIL, &equipIL,
			&bnetAccount, &bnetProfile, &bnetEquipment, &avatarURL, &renderURL, &bnetMedia); err != nil {
		t.Fatal(err)
	}
	if avatarURL == nil || *avatarURL != "https://render.worldofwarcraft.com/us/character/whitemane/1-thoradin-avatar.jpg" {
		t.Fatalf("avatar_url = %v", avatarURL)
	}
	if renderURL == nil || *renderURL != "https://render.worldofwarcraft.com/character/whitemane/1-thoradin-main-raw.png" {
		t.Fatalf("render_url = %v", renderURL)
	}
	if len(bnetMedia) == 0 {
		t.Fatal("bnet_media was not captured")
	}
	if region != "us" || ruleset != "pvp" || level != 60 || source != "bnet" {
		t.Fatalf("character row = %s/%s level=%d source=%s", region, ruleset, level, source)
	}
	if race != "Human" || gender != "male" {
		t.Fatalf("race=%q gender=%q, want Human/male", race, gender)
	}
	if avgIL != 40 || equipIL != 38 {
		t.Fatalf("average_item_level=%d equipped_item_level=%d, want 40/38", avgIL, equipIL)
	}
	if len(bnetAccount) == 0 || len(bnetProfile) == 0 || len(bnetEquipment) == 0 {
		t.Fatalf("capture columns = account=%d profile=%d equipment=%d bytes, want all non-empty",
			len(bnetAccount), len(bnetProfile), len(bnetEquipment))
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

func TestImportAccountRekeysWhenARealmsRulesetResolvesDifferently(t *testing.T) {
	pool := testPool(t)
	uid := seedUser(t, pool)
	ctx := context.Background()

	var oldGuildID int64
	if err := pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'normal', 'Old Flame') returning id`).
		Scan(&oldGuildID); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into characters (key, region, ruleset, name, user_id, realm_slug, bnet_character_id, source, refreshed_at)
		 values ('us/normal/dottzz', 'us', 'normal', 'Dottzz', $1, 'whitemane', 777, 'bnet', now())`, uid); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, rank, source, verified_by, verified_at)
		 values ($1, 'us/normal/dottzz', $2, 'member', 'bnet', 'bnet', now())`, oldGuildID, uid); err != nil {
		t.Fatal(err)
	}

	f := newBlizzardFixture(t)
	f.handlers[http.MethodGet+" "+accountCharactersPath] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"wow_accounts":[{"id":1,"characters":[
			{"name":"Dottzz","id":777,"realm":{"slug":"whitemane","name":"Whitemane"},
			 "playable_class":{"name":"Rogue"},"faction":{"type":"HORDE"},"level":60}
		]}]}`))
	}
	f.json(http.MethodGet, "/data/wow/realm/index?namespace=dynamic-classic1x-us", http.StatusOK,
		map[string]any{"realms": []map[string]any{{"id": 1, "name": "Whitemane", "slug": "whitemane"}}})
	// The realm detail this time resolves correctly to PVP (spec A1's own
	// fix): the character's key must move from us/normal/dottzz to
	// us/pvp/dottzz rather than duplicate (spec A3).
	f.json(http.MethodGet, "/data/wow/realm/whitemane?namespace=dynamic-classic1x-us", http.StatusOK,
		map[string]any{"type": map[string]string{"type": "PVP"}, "category": "PvP"})
	f.handlers[http.MethodGet+" /profile/wow/character/whitemane/dottzz?namespace=profile-classic1x-us"] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"name":"Dottzz","level":60,"faction":{"type":"HORDE"},
			"character_class":{"name":"Rogue"},"realm":{"slug":"whitemane","name":"Whitemane"}}`))
	}
	f.json(http.MethodGet, "/profile/wow/character/whitemane/dottzz/equipment?namespace=profile-classic1x-us",
		http.StatusOK, map[string]any{"equipped_items": []map[string]any{}})
	f.json(http.MethodGet, "/profile/wow/character/whitemane/dottzz/character-media?namespace=profile-classic1x-us",
		http.StatusOK, map[string]any{"assets": []map[string]string{}})

	svc := newTestService(t, pool, f)
	if _, err := svc.ImportAccount(ctx, uid, "user-oauth-token"); err != nil {
		t.Fatal(err)
	}

	var n int
	if err := pool.QueryRow(ctx, `select count(*) from characters where bnet_character_id = 777`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("characters rows for bnet_character_id=777 = %d, want 1 (moved, not duplicated)", n)
	}
	var key string
	if err := pool.QueryRow(ctx, `select key from characters where bnet_character_id = 777`).Scan(&key); err != nil {
		t.Fatal(err)
	}
	if key != "us/pvp/dottzz" {
		t.Fatalf("key = %q, want us/pvp/dottzz", key)
	}
	if err := pool.QueryRow(ctx, `select count(*) from characters where key = 'us/normal/dottzz'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("old key rows = %d, want 0", n)
	}
	if err := pool.QueryRow(ctx, `select count(*) from guild_characters where character_key = 'us/normal/dottzz'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("old guild_characters rows = %d, want 0 (deleted by the rekey)", n)
	}
}

func TestImportAccountLogsAndCountsA404CharacterProfileAsUnavailable(t *testing.T) {
	pool := testPool(t)
	uid := seedUser(t, pool)
	f := newBlizzardFixture(t)
	f.handlers[http.MethodGet+" "+accountCharactersPath] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"wow_accounts":[{"id":1,"characters":[
			{"name":"Sodpop","id":901,"realm":{"slug":"whitemane","name":"Whitemane"},
			 "playable_class":{"name":"Priest"},"faction":{"type":"ALLIANCE"},"level":25}
		]}]}`))
	}
	f.json(http.MethodGet, "/data/wow/realm/index?namespace=dynamic-classic1x-us", http.StatusOK,
		map[string]any{"realms": []map[string]any{{"id": 1, "name": "Whitemane", "slug": "whitemane"}}})
	f.json(http.MethodGet, "/data/wow/realm/whitemane?namespace=dynamic-classic1x-us", http.StatusOK,
		map[string]any{"type": map[string]string{"type": "NORMAL"}, "category": "PvE"})
	// Season of Discovery characters live in the classic1x account
	// profile but every character sub-endpoint answers 404 under that
	// namespace (spec A2).
	f.json(http.MethodGet, "/profile/wow/character/whitemane/sodpop?namespace=profile-classic1x-us",
		http.StatusNotFound, nil)
	f.json(http.MethodGet, "/profile/wow/character/whitemane/sodpop/equipment?namespace=profile-classic1x-us",
		http.StatusNotFound, nil)
	f.json(http.MethodGet, "/profile/wow/character/whitemane/sodpop/character-media?namespace=profile-classic1x-us",
		http.StatusNotFound, nil)

	svc := newTestService(t, pool, f)
	summary, err := svc.ImportAccount(context.Background(), uid, "user-oauth-token")
	if err != nil {
		t.Fatal(err)
	}
	if summary.Characters != 1 || summary.Unavailable != 1 || summary.Guilds != 0 {
		t.Fatalf("summary = %+v, want 1 character written, 1 unavailable, 0 guilds", summary)
	}
	var n int
	if err := pool.QueryRow(context.Background(), `select count(*) from guild_characters`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("guild_characters rows = %d, want 0", n)
	}
}
