// api/internal/guilds/home_addon_test.go
package guilds

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// seedRosterSnapshot writes the guild's roster snapshot the way POST /v1/addon/exports
// does (addon.Store.PutRosterSnapshot: one row per guild, a newer one replaces it), with an
// explicit received_at. The sender's characters row is the table's foreign key.
func seedRosterSnapshot(t *testing.T, pool *pgxpool.Pool, guildID int64, senderKey, members string, receivedAt time.Time) {
	t.Helper()
	ctx := context.Background()
	if _, err := pool.Exec(ctx,
		`insert into characters (key, region, ruleset, name) values ($1, 'us', 'hardcore', $1)
		 on conflict (key) do nothing`, senderKey); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into guild_roster_snapshots (guild_id, character_key, members, received_at)
		 values ($1, $2, $3::jsonb, $4)
		 on conflict (guild_id) do update set
		   character_key = excluded.character_key, members = excluded.members,
		   received_at = excluded.received_at`,
		guildID, senderKey, members, receivedAt); err != nil {
		t.Fatal(err)
	}
}

// addonGuild seeds a guild with three verified characters (Thorgrim, Healbot, Loner) whose
// exports put them on the home's roster, and returns the guild and the account reading it.
func addonGuild(t *testing.T, h *httpHarness) (int64, int64) {
	t.Helper()
	gid := seedGuild(t, h.pool, "Forever")
	var reader int64
	for i, name := range []string{"Thorgrim", "Healbot", "Loner"} {
		user := seedUser(t, h.pool, fmt.Sprintf("addon-%d@example.com", i))
		key := "us/hardcore/" + name
		seedCharacter(t, h.pool, gid, user, key, "member", true)
		seedExport(t, h.pool, user, key, "us", "hardcore", name)
		recomputeMembership(t, h.pool, gid, user)
		if i == 0 {
			reader = user
		}
	}
	return gid, reader
}

// readHomeRaw reads the home as actor and returns both the typed view and the raw data
// object, so a test can tell a null key from an absent one.
func readHomeRaw(t *testing.T, h *httpHarness, gid, actor int64) (HomeView, map[string]json.RawMessage) {
	t.Helper()
	h.actor = auth.Actor{UserID: actor, Role: "user", Method: "session"}
	var raw json.RawMessage
	h.data(h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), ""), &raw)
	var view HomeView
	if err := json.Unmarshal(raw, &view); err != nil {
		t.Fatal(err)
	}
	var top map[string]json.RawMessage
	if err := json.Unmarshal(raw, &top); err != nil {
		t.Fatal(err)
	}
	return view, top
}

// rosterAddonRaw is each roster row's raw "addon" value by name; a row without the key at
// all reads as "absent".
func rosterAddonRaw(t *testing.T, top map[string]json.RawMessage) map[string]string {
	t.Helper()
	var rows []map[string]json.RawMessage
	if err := json.Unmarshal(top["roster"], &rows); err != nil {
		t.Fatal(err)
	}
	out := map[string]string{}
	for _, row := range rows {
		var name string
		if err := json.Unmarshal(row["name"], &name); err != nil {
			t.Fatal(err)
		}
		v, ok := row["addon"]
		if !ok {
			out[name] = "absent"
			continue
		}
		out[name] = string(v)
	}
	return out
}

func TestHomeRosterCarriesTheNewestAddonSnapshot(t *testing.T) {
	h := newHTTPHarness(t)
	gid, reader := addonGuild(t, h)

	older := time.Now().Add(-2 * time.Hour).UTC().Truncate(time.Second)
	seedRosterSnapshot(t, h.pool, gid, "us/hardcore/Healbot", `[
	  {"name":"Thorgrim-Yell","level":48,"spec":"warrior-fury","gear_gap":20,"build":"Fury leveling","seen_at":"2026-11-04 12:00"},
	  {"name":"Loner-Yell","level":30,"spec":"rogue-combat","gear_gap":9,"build":"Combat","seen_at":"2026-11-04 12:00"}]`, older)
	newer := time.Now().Add(-3 * time.Minute).UTC().Truncate(time.Second)
	seedRosterSnapshot(t, h.pool, gid, "us/hardcore/Thorgrim", `[
	  {"name":"Thorgrim-Yell","level":52,"spec":"warrior-arms","gear_gap":3,"build":"Arms raid","seen_at":"2026-11-04 14:01"},
	  {"name":"healbot","level":40,"spec":"priest-holy","gear_gap":-2,"build":"","seen_at":"2026-11-04 13:55"},
	  {"name":"Stranger-Yell","level":60,"spec":"mage-frost","gear_gap":0,"build":"Frost AoE","seen_at":"2026-11-04 14:00"}]`, newer)

	view, top := readHomeRaw(t, h, gid, reader)

	if view.AddonSnapshot == nil {
		t.Fatalf("addon_snapshot = null, want the newer snapshot")
	}
	if snap := *view.AddonSnapshot; !snap.ReceivedAt.Equal(newer) || snap.ByCharacter != "us/hardcore/Thorgrim" || snap.Members != 3 {
		t.Fatalf("addon_snapshot = %+v, want received %s by us/hardcore/Thorgrim with 3 members", snap, newer)
	}
	var snapRaw map[string]json.RawMessage
	if err := json.Unmarshal(top["addon_snapshot"], &snapRaw); err != nil {
		t.Fatal(err)
	}
	if got, want := string(snapRaw["received_at"]), `"`+newer.Format(time.RFC3339)+`"`; got != want {
		t.Fatalf("addon_snapshot.received_at = %s, want %s (RFC 3339 UTC)", got, want)
	}

	byName := map[string]RosterRow{}
	for _, row := range view.Roster {
		byName[row.Name] = row
	}
	if len(byName) != 3 {
		t.Fatalf("roster = %d rows, want 3 (an unmatched snapshot member is never a row)", len(byName))
	}
	thorgrim := byName["Thorgrim"].Addon
	if thorgrim == nil || thorgrim.Level != 52 || thorgrim.Spec != "warrior-arms" || thorgrim.GearGap != 3 ||
		thorgrim.Build == nil || *thorgrim.Build != "Arms raid" || thorgrim.SeenAt != "2026-11-04 14:01" ||
		!thorgrim.SnapshotAt.Equal(newer) {
		t.Fatalf("Thorgrim's addon = %+v, want the newer snapshot's line", thorgrim)
	}
	healbot := byName["Healbot"].Addon
	if healbot == nil || healbot.Level != 40 || healbot.GearGap != -2 || healbot.Build != nil || healbot.SeenAt != "2026-11-04 13:55" {
		t.Fatalf("Healbot's addon = %+v, want level 40, gear gap -2 and a null build (matched by name, any case)", healbot)
	}
	if loner := byName["Loner"].Addon; loner != nil {
		t.Fatalf("Loner's addon = %+v, want null: only the older snapshot carried him", loner)
	}

	raws := rosterAddonRaw(t, top)
	if raws["Loner"] != "null" {
		t.Fatalf("Loner's raw addon = %s, want an explicit null", raws["Loner"])
	}
	var thorgrimRaw map[string]json.RawMessage
	if err := json.Unmarshal([]byte(raws["Thorgrim"]), &thorgrimRaw); err != nil {
		t.Fatal(err)
	}
	if got, want := string(thorgrimRaw["snapshot_at"]), `"`+newer.Format(time.RFC3339)+`"`; got != want {
		t.Fatalf("Thorgrim's addon.snapshot_at = %s, want %s", got, want)
	}
	for _, key := range []string{"level", "spec", "gear_gap", "build", "seen_at", "snapshot_at"} {
		if _, ok := thorgrimRaw[key]; !ok {
			t.Errorf("Thorgrim's addon has no %q key: %s", key, raws["Thorgrim"])
		}
	}
	if string(byNameRaw(t, top, "Healbot")["build"]) != "null" {
		t.Fatalf("Healbot's addon.build = %s, want null", byNameRaw(t, top, "Healbot")["build"])
	}
}

// byNameRaw is one roster row's raw addon object.
func byNameRaw(t *testing.T, top map[string]json.RawMessage, name string) map[string]json.RawMessage {
	t.Helper()
	var obj map[string]json.RawMessage
	if err := json.Unmarshal([]byte(rosterAddonRaw(t, top)[name]), &obj); err != nil {
		t.Fatal(err)
	}
	return obj
}

func TestHomeWithNoAddonSnapshotIsNullEverywhere(t *testing.T) {
	h := newHTTPHarness(t)
	gid, reader := addonGuild(t, h)

	view, top := readHomeRaw(t, h, gid, reader)

	if raw, ok := top["addon_snapshot"]; !ok || string(raw) != "null" {
		t.Fatalf("addon_snapshot = %s (present %v), want an explicit null", raw, ok)
	}
	if view.AddonSnapshot != nil {
		t.Fatalf("addon_snapshot = %+v, want nil", view.AddonSnapshot)
	}
	raws := rosterAddonRaw(t, top)
	if len(raws) != 3 {
		t.Fatalf("roster = %d rows, want 3", len(raws))
	}
	for name, raw := range raws {
		if raw != "null" {
			t.Errorf("%s's addon = %s, want null", name, raw)
		}
	}
}

func TestHomeWithAnEmptyAddonSnapshotCountsZeroMembers(t *testing.T) {
	h := newHTTPHarness(t)
	gid, reader := addonGuild(t, h)

	seedRosterSnapshot(t, h.pool, gid, "us/hardcore/Healbot", `[
	  {"name":"Thorgrim-Yell","level":48,"spec":"warrior-fury","gear_gap":20,"build":"Fury leveling","seen_at":"2026-11-04 12:00"}]`,
		time.Now().Add(-time.Hour))
	newest := time.Now().Add(-time.Minute).UTC().Truncate(time.Second)
	seedRosterSnapshot(t, h.pool, gid, "us/hardcore/Thorgrim", `[]`, newest)

	view, top := readHomeRaw(t, h, gid, reader)

	if view.AddonSnapshot == nil || view.AddonSnapshot.Members != 0 || !view.AddonSnapshot.ReceivedAt.Equal(newest) ||
		view.AddonSnapshot.ByCharacter != "us/hardcore/Thorgrim" {
		t.Fatalf("addon_snapshot = %+v, want the empty newest snapshot (members 0)", view.AddonSnapshot)
	}
	for name, raw := range rosterAddonRaw(t, top) {
		if raw != "null" {
			t.Errorf("%s's addon = %s, want null: the newest snapshot carries nobody", name, raw)
		}
	}
}
