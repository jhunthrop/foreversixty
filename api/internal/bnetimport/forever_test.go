// api/internal/bnetimport/forever_test.go
package bnetimport

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

const foreverAccountPath = "/profile/user/wow?namespace=profile-forever-us"

// foreverFixture stands up the Forever account the shared modeled fixture
// describes (bnetapi/testdata/forever): Elyra Duskvale (60, realm 6798 PvP,
// guild leader of Quiet Hours), Tor Ashbrand (42, realm 6928 RP, profile not
// public yet) and Bank Holder (1, realm 6798, profile not public yet).
func foreverFixture(t *testing.T) *blizzardFixture {
	t.Helper()
	f := newBlizzardFixture(t)
	account, err := os.ReadFile("../bnetapi/testdata/forever/account-profile-modeled.json")
	if err != nil {
		t.Fatal(err)
	}
	f.rawJSON(foreverAccountPath, account)
	f.json(http.MethodGet, "/data/wow/realm/index?namespace=dynamic-forever-us", http.StatusOK, map[string]any{"realms": []map[string]any{
		{"id": 6798, "name": "CAMWOW US1 Realm 2 PvP", "slug": "camwow-us1-realm-2-pvp"},
		{"id": 6928, "name": "CAMWOW US1 Realm 10 RP", "slug": "camwow-us1-realm-10-rp"},
	}})
	f.json(http.MethodGet, "/data/wow/realm/camwow-us1-realm-2-pvp?namespace=dynamic-forever-us", http.StatusOK,
		map[string]any{"type": map[string]string{"type": "PVP"}, "category": "United States"})
	f.json(http.MethodGet, "/data/wow/realm/camwow-us1-realm-10-rp?namespace=dynamic-forever-us", http.StatusOK,
		map[string]any{"type": map[string]string{"type": "RP"}, "category": "United States"})

	f.rawJSON("/profile/wow/character/elyra-duskvale?namespace=profile-forever-us",
		[]byte(`{"name":"Elyra","surname":"Duskvale","level":60,"faction":{"type":"ALLIANCE"},
			"character_class":{"name":"Mage","key":{"href":"https://us.api.blizzard.com/data/wow/playable-class/8?namespace=static-1.60.1_70399-forever-us"}},
			"race":{"name":"Human"},"guild":{"name":"Quiet Hours","id":3},"last_login_timestamp":1762200000000}`))
	f.rawJSON("/data/wow/guild/quiet-hours/roster?namespace=profile-forever-us",
		[]byte(`{"guild":{"name":"Quiet Hours","id":3},"members":[
			{"character":{"id":7001,"name":"Elyra","surname":"Duskvale","level":60,"playable_class":{"id":8}},"rank":0}]}`))
	f.json(http.MethodGet, "/profile/wow/character/elyra-duskvale/equipment?namespace=profile-forever-us",
		http.StatusOK, map[string]any{"equipped_items": []map[string]any{}})
	f.json(http.MethodGet, "/profile/wow/character/elyra-duskvale/character-media?namespace=profile-forever-us",
		http.StatusOK, map[string]any{"assets": []map[string]string{}})

	for _, name := range []string{"tor-ashbrand", "bank-holder"} {
		for _, sub := range []string{"", "/equipment", "/character-media"} {
			f.json(http.MethodGet, "/profile/wow/character/"+name+sub+"?namespace=profile-forever-us", http.StatusNotFound, nil)
		}
	}
	return f
}

func (f *blizzardFixture) rawJSON(path string, body []byte) {
	f.handlers[http.MethodGet+" "+path] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(body)
	}
}

func foreverService(t *testing.T, f *blizzardFixture) *Service {
	svc := newTestService(t, testPool(t), f)
	svc.Client.Game = "forever"
	return svc
}

func TestImportAccountImportsForeverCharactersByNameSurname(t *testing.T) {
	f := foreverFixture(t)
	svc := foreverService(t, f)
	uid := seedUser(t, svc.Pool)

	summary, err := svc.ImportAccount(context.Background(), uid, "user-oauth-token")
	if err != nil {
		t.Fatal(err)
	}
	if summary.Characters != 3 || summary.Skipped != 0 || summary.Guilds != 1 || summary.Unavailable != 2 || summary.Pending != 0 {
		t.Fatalf("summary = %+v", summary)
	}

	rows := map[string]string{}
	r, err := svc.Pool.Query(context.Background(),
		`select key, name from characters where user_id = $1 and realm_slug is null and source = 'bnet'`, uid)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	for r.Next() {
		var key, name string
		if err := r.Scan(&key, &name); err != nil {
			t.Fatal(err)
		}
		rows[key] = name
	}
	if rows["us/pvp/elyra-duskvale"] != "Elyra Duskvale" || rows["us/rp/tor-ashbrand"] != "Tor Ashbrand" ||
		rows["us/pvp/bank-holder"] != "Bank Holder" || len(rows) != 3 {
		t.Fatalf("rows = %v, want all three, the level 1 included, keyed by their realm's ruleset", rows)
	}

	var rank string
	if err := svc.Pool.QueryRow(context.Background(),
		`select rank from guild_characters where character_key = 'us/pvp/elyra-duskvale'`).Scan(&rank); err != nil {
		t.Fatal(err)
	}
	if rank != "leader" {
		t.Fatalf("rank = %q, want leader (matched on the roster by character id)", rank)
	}
}

// Since 2026-11-05 Blizzard answers 404 for Forever's realm index, so the
// realm id resolves nothing; the character profile's own ruleset is the
// truth. Production keyed such a PvP character us/normal and made its guild
// a us/normal twin of the real us/pvp one.
func TestImportAccountKeysByTheProfilesRulesetWhenRealmsAreGone(t *testing.T) {
	f := foreverFixture(t)
	f.json(http.MethodGet, "/data/wow/realm/index?namespace=dynamic-forever-us", http.StatusNotFound, nil)
	f.rawJSON("/profile/wow/character/elyra-duskvale?namespace=profile-forever-us",
		[]byte(`{"name":"Elyra","surname":"Duskvale","level":60,"faction":{"type":"ALLIANCE"},
			"character_class":{"name":"Mage"},"race":{"name":"Human"},
			"ruleset":{"id":1,"name":"PvP"},
			"guild":{"name":"Quiet Hours","id":3},"last_login_timestamp":1762200000000}`))
	svc := foreverService(t, f)
	uid := seedUser(t, svc.Pool)

	if _, err := svc.ImportAccount(context.Background(), uid, "user-oauth-token"); err != nil {
		t.Fatal(err)
	}
	var ruleset, guildRuleset string
	if err := svc.Pool.QueryRow(context.Background(),
		`select c.ruleset, g.ruleset from characters c
		   join guild_characters gc on gc.character_key = c.key join guilds g on g.id = gc.guild_id
		  where c.user_id = $1 and c.name = 'Elyra Duskvale'`, uid).Scan(&ruleset, &guildRuleset); err != nil {
		t.Fatal(err)
	}
	if ruleset != "pvp" || guildRuleset != "pvp" {
		t.Fatalf("character ruleset %q, guild ruleset %q; want both pvp, from the profile", ruleset, guildRuleset)
	}
	var stale int
	if err := svc.Pool.QueryRow(context.Background(),
		`select count(*) from characters where user_id = $1 and key = 'us/normal/elyra-duskvale'`, uid).Scan(&stale); err != nil {
		t.Fatal(err)
	}
	if stale != 0 {
		t.Fatal("the normal-keyed row was left behind")
	}
}

// Blizzard lists no Forever realm at all (the realm index, the connected-realm
// index and every realm by id or slug answered 404 on 2026-11-12; the realm
// search answers an empty page). The realm id the protected-character link
// carries still names the realm, and the committed Forever realm table
// (bnetapi.KnownRealm) names its type, so every character is keyed without a
// profile call: Elyra and Bank on 6798 (PvP), Tor on 6928 (RP).
func TestImportAccountKeysByTheRealmTableWhenBlizzardListsNoRealm(t *testing.T) {
	f := foreverFixture(t)
	f.json(http.MethodGet, "/data/wow/realm/index?namespace=dynamic-forever-us", http.StatusNotFound, nil)
	svc := foreverService(t, f)
	uid := seedUser(t, svc.Pool)

	summary, err := svc.ImportAccount(context.Background(), uid, "user-oauth-token")
	if err != nil {
		t.Fatal(err)
	}
	if summary.Characters != 3 || summary.Skipped != 0 {
		t.Fatalf("summary = %+v, want all three keyed by the realm table", summary)
	}
	rows, err := svc.Pool.Query(context.Background(), `select key from characters where user_id = $1 order by key`, uid)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	var keys []string
	for rows.Next() {
		var k string
		if err := rows.Scan(&k); err != nil {
			t.Fatal(err)
		}
		keys = append(keys, k)
	}
	want := []string{"us/pvp/bank-holder", "us/pvp/elyra-duskvale", "us/rp/tor-ashbrand"}
	if strings.Join(keys, " ") != strings.Join(want, " ") {
		t.Fatalf("keys = %v, want %v", keys, want)
	}
}


func TestImportAccountListsEveryCharacterBeforeTheBudgetRunsOut(t *testing.T) {
	f := foreverFixture(t)
	slow := f.handlers[http.MethodGet+" /profile/wow/character/elyra-duskvale?namespace=profile-forever-us"]
	f.handlers[http.MethodGet+" /profile/wow/character/elyra-duskvale?namespace=profile-forever-us"] = func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
		case <-time.After(2 * time.Second):
		}
		slow(w, r)
	}
	svc := foreverService(t, f)
	uid := seedUser(t, svc.Pool)

	ctx, cancel := context.WithTimeout(context.Background(), 700*time.Millisecond)
	defer cancel()
	summary, err := svc.ImportAccount(ctx, uid, "user-oauth-token")
	if !errors.Is(err, ErrBudget) {
		t.Fatalf("err = %v, want ErrBudget", err)
	}
	if summary.Characters != 3 || summary.Pending != 3 {
		t.Fatalf("summary = %+v, want every row written and every one pending enrichment", summary)
	}
	var n int
	if err := svc.Pool.QueryRow(context.Background(),
		`select count(*) from characters where user_id = $1`, uid).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 3 {
		t.Fatalf("characters = %d, want 3 listed even though enrichment ran out of budget", n)
	}
}

func TestImportAccountKeepsAnAddonRowsRulesetWhenNoRealmIsKnown(t *testing.T) {
	f := newBlizzardFixture(t)
	f.rawJSON(foreverAccountPath, []byte(`{"wow_accounts":[{"id":1,"characters":[
		{"name":"Bow","surname":"Jackzon","id":9001,"playable_class":{"name":"Hunter"},
		 "playable_race":{"name":"Dwarf"},"gender":{"type":"MALE"},"faction":{"type":"ALLIANCE"},"level":60}]}]}`))
	f.json(http.MethodGet, "/data/wow/realm/index?namespace=dynamic-forever-us", http.StatusOK, map[string]any{"realms": []any{}})
	for _, sub := range []string{"", "/equipment", "/character-media"} {
		f.json(http.MethodGet, "/profile/wow/character/bow-jackzon"+sub+"?namespace=profile-forever-us", http.StatusNotFound, nil)
	}
	svc := foreverService(t, f)
	uid := seedUser(t, svc.Pool)
	if _, err := svc.Pool.Exec(context.Background(),
		`insert into characters (key, region, ruleset, name, class, user_id) values ('us/pvp/bow-jackzon', 'us', 'pvp', 'Bow Jackzon', 'hunter', $1)`,
		uid); err != nil {
		t.Fatal(err)
	}

	if _, err := svc.ImportAccount(context.Background(), uid, "user-oauth-token"); err != nil {
		t.Fatal(err)
	}
	var key string
	var bnetID int64
	if err := svc.Pool.QueryRow(context.Background(),
		`select key, bnet_character_id from characters where user_id = $1`, uid).Scan(&key, &bnetID); err != nil {
		t.Fatal(err)
	}
	if key != "us/pvp/bow-jackzon" || bnetID != 9001 {
		t.Fatalf("key = %q, bnet id = %d: want the addon row adopted, not a second row under normal", key, bnetID)
	}
}

// Sim audit round 17: a Forever race's base stats (Skyborne) come only from the profile
// API's /statistics resource, which the import never read. It is read when it answers and
// stored verbatim in bnet_statistics.
func TestImportAccountRecordsTheStatisticsWhenTheyAnswer(t *testing.T) {
	f := foreverFixture(t)
	f.rawJSON("/profile/wow/character/elyra-duskvale/statistics?namespace=profile-forever-us",
		[]byte(`{"strength":{"base":20,"effective":31},"intellect":{"base":125,"effective":290}}`))
	svc := foreverService(t, f)
	uid := seedUser(t, svc.Pool)
	if _, err := svc.ImportAccount(context.Background(), uid, "user-oauth-token"); err != nil {
		t.Fatal(err)
	}
	var base int
	if err := svc.Pool.QueryRow(context.Background(),
		`select (bnet_statistics->'intellect'->>'base')::int from characters where key = 'us/pvp/elyra-duskvale'`).
		Scan(&base); err != nil {
		t.Fatal(err)
	}
	if base != 125 {
		t.Fatalf("intellect base = %d, want 125 from /statistics", base)
	}
}

// The read is guarded: a /statistics that fails outright never costs the character its
// import, its equipment or its export.
func TestImportAccountKeepsGoingWhenTheStatisticsFail(t *testing.T) {
	f := foreverFixture(t)
	f.json(http.MethodGet, "/profile/wow/character/elyra-duskvale/statistics?namespace=profile-forever-us",
		http.StatusInternalServerError, nil)
	svc := foreverService(t, f)
	uid := seedUser(t, svc.Pool)
	summary, err := svc.ImportAccount(context.Background(), uid, "user-oauth-token")
	if err != nil {
		t.Fatal(err)
	}
	if summary.Characters != 3 || summary.Pending != 0 {
		t.Fatalf("summary = %+v, want every character imported and enriched", summary)
	}
	var stats *string
	var captured bool
	if err := svc.Pool.QueryRow(context.Background(),
		`select bnet_statistics::text, bnet_equipment is not null from characters where key = 'us/pvp/elyra-duskvale'`).
		Scan(&stats, &captured); err != nil {
		t.Fatal(err)
	}
	if stats != nil || !captured {
		t.Fatalf("bnet_statistics = %v, equipment captured = %v: want no statistics and the rest intact", stats, captured)
	}
}
