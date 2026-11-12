// api/internal/guilds/raids_test.go
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/db"
)

// TestRaidsVisibilityMatchesReportRules checks a signed-out visitor sees only the public
// report, a verified member additionally sees the guild-visibility one, and the owner sees
// their own private report too - the exact rule HomeReports already enforces.
func TestRaidsVisibilityMatchesReportRules(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	owner := seedUser(t, h.pool, "raids-owner@example.com")
	seedCharacter(t, h.pool, gid, owner, "us/hardcore/raidsowner", "member", true)
	verifiedMember := seedUser(t, h.pool, "raids-member@example.com")
	seedCharacter(t, h.pool, gid, verifiedMember, "us/hardcore/raidsmember", "member", true)
	recomputeMembership(t, h.pool, gid, verifiedMember)

	for _, r := range []struct{ id, visibility string }{
		{"raidspublic", "public"}, {"raidsguild", "guild"}, {"raidsprivate", "private"},
	} {
		if _, err := h.pool.Exec(ctx,
			`insert into reports (id, owner_id, guild_id, visibility, status, created_at)
			 values ($1, $2, $3, $4, 'complete', now())`, r.id, owner, gid, r.visibility); err != nil {
			t.Fatal(err)
		}
	}

	ids := func(res *http.Response) map[string]bool {
		var page RaidsPage
		h.data(res, &page)
		out := map[string]bool{}
		for _, row := range page.Rows {
			out[row.ID] = true
		}
		return out
	}

	// Signed out.
	h.actor = auth.Actor{}
	seen := ids(h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/raids", gid), ""))
	if !seen["raidspublic"] || seen["raidsguild"] || seen["raidsprivate"] {
		t.Fatalf("signed-out visitor saw %v, want only the public report", seen)
	}

	// Verified member.
	h.actor = auth.Actor{UserID: verifiedMember, Role: "user", Method: "session"}
	seen = ids(h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/raids", gid), ""))
	if !seen["raidspublic"] || !seen["raidsguild"] || seen["raidsprivate"] {
		t.Fatalf("verified member saw %v, want public and guild, not private", seen)
	}

	// Owner.
	h.actor = auth.Actor{UserID: owner, Role: "user", Method: "session"}
	seen = ids(h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/raids", gid), ""))
	if !seen["raidspublic"] || !seen["raidsguild"] || !seen["raidsprivate"] {
		t.Fatalf("owner saw %v, want every one of their own reports", seen)
	}
}

// TestRaidsFightsAndPresentAndTopParse seeds one report with two fights (a kill and a
// wipe) and checks the row's own fight list, present roster, Top HPS and Top DPS.
func TestRaidsFightsAndPresentAndTopParse(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	uid := seedUser(t, h.pool, "raids-detail@example.com")
	key := "us/hardcore/raidsdetail"
	seedCharacter(t, h.pool, gid, uid, key, "member", true)
	seedExport(t, h.pool, uid, key, "us", "hardcore", "raidsdetail")

	if err := db.EnsureMetricsPartitions(ctx, h.pool, time.Now(), 0); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx,
		`insert into reports (id, guild_id, visibility, status, zone, created_at)
		 values ('raids-detail-report', $1, 'public', 'complete', 'Onyxia''s Lair', now())`, gid); err != nil {
		t.Fatal(err)
	}
	// fights.players holds the log's GUIDs, as engine.go writes it; the present list reads
	// the character by key from fight_metrics (task L6.8).
	if _, err := h.pool.Exec(ctx,
		`insert into fights (report_id, fight_index, name, encounter_id, kill, duration_ms, start_ms, deaths, players)
		 values ('raids-detail-report', 0, 'Onyxia', 1084, false, 200000, 1000, 2, $1)`, []string{l68GUID}); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx,
		`insert into fights (report_id, fight_index, name, encounter_id, kill, duration_ms, start_ms, deaths, players)
		 values ('raids-detail-report', 1, 'Onyxia', 1084, true, 260000, 201000, 1, $1)`, []string{l68GUID}); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx,
		`insert into fight_metrics (report_id, fight_index, player_key, player_name, class, role, metric_dps, kill, fought_at, encounter_id)
		 values ('raids-detail-report', 1, $1, 'Raidsdetail', 'hunter', 'dps', 300.5, true, now(), 1084)`, key); err != nil {
		t.Fatal(err)
	}
	// A healer's HPS is the night's Top HPS; their lower DPS never contends for Top DPS.
	if _, err := h.pool.Exec(ctx,
		`insert into fight_metrics (report_id, fight_index, player_key, player_name, class, role, metric_dps, metric_hps, kill, fought_at, encounter_id)
		 values ('raids-detail-report', 1, 'us/hardcore/rootbound', 'Rootbound', 'druid', 'healer', 900, 449.4, true, now(), 1084)`); err != nil {
		t.Fatal(err)
	}

	h.actor = auth.Actor{}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/raids", gid), "")
	var page RaidsPage
	h.data(res, &page)
	if len(page.Rows) != 1 {
		t.Fatalf("rows = %d, want 1", len(page.Rows))
	}
	row := page.Rows[0]
	if row.FightCount != 2 || row.KillCount != 1 || row.WipeCount != 1 {
		t.Errorf("fight/kill/wipe = %d/%d/%d, want 2/1/1", row.FightCount, row.KillCount, row.WipeCount)
	}
	if row.Deaths != 3 {
		t.Errorf("deaths = %d, want 3", row.Deaths)
	}
	if len(row.Fights) != 2 || row.Fights[0].EncounterID == nil || *row.Fights[0].EncounterID != 1084 {
		t.Fatalf("fights = %+v, want 2 rows naming encounter 1084", row.Fights)
	}
	if row.Fights[1].Players != 1 {
		t.Errorf("fights[1].players = %d, want 1", row.Fights[1].Players)
	}
	// The healer row seeded for Top HPS is present too (fight_metrics names them), after the hunter.
	if len(row.Present) != 2 || row.Present[0].CharacterKey != key || row.Present[0].Name != "raidsdetail" {
		t.Fatalf("present = %+v, want %s named raidsdetail (addon_exports.name, preferred over fight_metrics.player_name) and the healer", row.Present, key)
	}
	if row.TopDPS == nil || row.TopDPS.Metric != "dps" || row.TopDPS.Value != 300.5 || row.TopDPS.Name != "Raidsdetail" {
		t.Fatalf("top_dps = %+v, want dps 300.5 for Raidsdetail", row.TopDPS)
	}
	if row.TopHPS == nil || row.TopHPS.Metric != "hps" || row.TopHPS.Value != 449.4 || row.TopHPS.Name != "Rootbound" {
		t.Fatalf("top_hps = %+v, want hps 449.4 for Rootbound", row.TopHPS)
	}
}

// TestRaidsCursorPaginates seeds 3 reports and checks a limit=2 page plus its cursor reach
// the third.
func TestRaidsCursorPaginates(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	for i := 0; i < 3; i++ {
		if _, err := h.pool.Exec(ctx,
			`insert into reports (id, guild_id, visibility, status, created_at)
			 values ($1, $2, 'public', 'complete', now() - ($3 || ' minutes')::interval)`,
			fmt.Sprintf("cursor-report-%d", i), gid, fmt.Sprint(i)); err != nil {
			t.Fatal(err)
		}
	}
	h.actor = auth.Actor{}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/raids?limit=2", gid), "")
	var page RaidsPage
	h.data(res, &page)
	if len(page.Rows) != 2 || page.NextCursor == "" {
		t.Fatalf("first page = %d rows, cursor %q; want 2 rows and a cursor", len(page.Rows), page.NextCursor)
	}
	if page.Rows[0].ID != "cursor-report-0" || page.Rows[1].ID != "cursor-report-1" {
		t.Fatalf("first page ids = %s, %s; want cursor-report-0, cursor-report-1", page.Rows[0].ID, page.Rows[1].ID)
	}
	res = h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/raids?limit=2&cursor=%s", gid, page.NextCursor), "")
	h.data(res, &page)
	if len(page.Rows) != 1 || page.Rows[0].ID != "cursor-report-2" {
		t.Fatalf("second page = %+v, want exactly cursor-report-2", page.Rows)
	}
}

// TestRaidsListAMembersAttributedReport covers the Raids tab for a report nobody set by
// hand: a member's public upload the matcher reads as the guild's (matched_guild_id) is
// one of the guild's rows, for a verified member and a signed-out visitor alike; a report
// with neither column set is nobody's.
func TestRaidsListAMembersAttributedReport(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	owner := seedUser(t, h.pool, "raids-attributed-owner@example.com")
	seedCharacter(t, h.pool, gid, owner, "us/hardcore/raidsattributedowner", "member", true)
	member := seedUser(t, h.pool, "raids-attributed-member@example.com")
	seedCharacter(t, h.pool, gid, member, "us/hardcore/raidsattributedmember", "member", true)
	recomputeMembership(t, h.pool, gid, member)

	if _, err := h.pool.Exec(ctx,
		`insert into reports (id, owner_id, matched_guild_id, visibility, status, created_at)
		 values ('raidsattributed', $1, $2, 'public', 'complete', now())`, owner, gid); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx,
		`insert into reports (id, owner_id, visibility, status, created_at)
		 values ('raidsnobodys', $1, 'public', 'complete', now())`, owner); err != nil {
		t.Fatal(err)
	}

	for _, actor := range []auth.Actor{{}, {UserID: member, Role: "user", Method: "session"}} {
		h.actor = actor
		var page RaidsPage
		h.data(h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/raids", gid), ""), &page)
		if len(page.Rows) != 1 || page.Rows[0].ID != "raidsattributed" {
			t.Fatalf("actor %+v saw %+v, want only the attributed report", actor, page.Rows)
		}
	}
}
