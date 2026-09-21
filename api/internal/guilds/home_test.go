// api/internal/guilds/home_test.go
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/db"
)

// recomputeMembership seeds the derived guild_members row (consent, rank,
// verified_at) for userID from its guild_characters rows. seedCharacter
// only writes guild_characters; a test that needs guild_members to exist
// first — e.g. to set a non-default consent, or because HomeRoster inner
// joins on it — calls this after seeding the character.
func recomputeMembership(t *testing.T, pool *pgxpool.Pool, guildID, userID int64) {
	t.Helper()
	ctx := context.Background()
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback(ctx)
	if err := RecomputeMembership(ctx, tx, guildID, &userID); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatal(err)
	}
}

// seedExport inserts the addon_exports row HomeRoster's query inner joins
// guild_characters against — seedCharacter never creates one, and no
// existing helper in this package does either.
func seedExport(t *testing.T, pool *pgxpool.Pool, userID int64, key, region, ruleset, name string) {
	t.Helper()
	if _, err := pool.Exec(context.Background(), `
		insert into addon_exports (character_key, user_id, region, ruleset, name, export, updated_at)
		values ($1, $2, $3, $4, $5, '', now())`,
		key, userID, region, ruleset, name); err != nil {
		t.Fatal(err)
	}
}

func TestHomeRequiresMembership(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	stranger := seedUser(t, h.pool, "home-stranger@example.com")
	h.actor = auth.Actor{UserID: stranger, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("a non-member reading home = %d, want 403", res.StatusCode)
	}

	member := seedUser(t, h.pool, "home-member@example.com")
	seedCharacter(t, h.pool, gid, member, "us/hardcore/homemember", "member", false)
	h.actor = auth.Actor{UserID: member, Role: "user", Method: "session"}
	res = h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("an unverified member reading home = %d, want 200 (the home shell is open to any member)", res.StatusCode)
	}
}

func TestHomeRosterHonoursConsent(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")

	// fight_metrics is partitioned by month; unlike reports/sims/rankings'
	// harnesses, this package's harness_test.go never ensures a partition
	// exists, so a bare insert with fought_at = now() has nowhere to land.
	if err := db.EnsureMetricsPartitions(ctx, h.pool, time.Now(), 0); err != nil {
		t.Fatal(err)
	}

	rosterMember := seedUser(t, h.pool, "roster-consent@example.com")
	seedCharacter(t, h.pool, gid, rosterMember, "us/hardcore/rosterconsent", "member", true)
	seedExport(t, h.pool, rosterMember, "us/hardcore/rosterconsent", "us", "hardcore", "rosterconsent")
	recomputeMembership(t, h.pool, gid, rosterMember)
	if _, err := h.pool.Exec(ctx,
		`update guild_members set consent = 'roster' where guild_id = $1 and user_id = $2`, gid, rosterMember); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx,
		`insert into fight_metrics (report_id, fight_index, player_key, class, spec, ilvl, fought_at)
		 values ('homefight1', 0, 'us/hardcore/rosterconsent', 'Warrior', 'Fury', 60, now())`); err != nil {
		t.Fatal(err)
	}

	gearMember := seedUser(t, h.pool, "gear-consent@example.com")
	seedCharacter(t, h.pool, gid, gearMember, "us/hardcore/gearconsent", "member", true)
	seedExport(t, h.pool, gearMember, "us/hardcore/gearconsent", "us", "hardcore", "gearconsent")
	recomputeMembership(t, h.pool, gid, gearMember)
	if _, err := h.pool.Exec(ctx,
		`insert into fight_metrics (report_id, fight_index, player_key, class, spec, ilvl, fought_at)
		 values ('homefight2', 0, 'us/hardcore/gearconsent', 'Mage', 'Frost', 55, now())`); err != nil {
		t.Fatal(err)
	}

	h.actor = auth.Actor{UserID: rosterMember, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	var view HomeView
	h.data(res, &view)

	byKey := map[string]RosterRow{}
	for _, row := range view.Roster {
		byKey[row.CharacterKey] = row
	}
	if row := byKey["us/hardcore/rosterconsent"]; row.ItemLevel != nil {
		t.Fatalf("roster-consent row leaked item level: %+v", row)
	}
	if row := byKey["us/hardcore/gearconsent"]; row.ItemLevel == nil || *row.ItemLevel != 55 {
		t.Fatalf("gear-consent row = %+v, want item_level 55", row)
	}
}

func TestHomeReportsListsOnlyTheTrailingWeek(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	uid := seedUser(t, h.pool, "home-reports@example.com")
	seedCharacter(t, h.pool, gid, uid, "us/hardcore/homereports", "member", true)

	if _, err := h.pool.Exec(ctx,
		`insert into reports (id, owner_id, guild_id, visibility, status, created_at)
		 values ('recenthome', $1, $2, 'guild', 'complete', now() - interval '2 days')`, uid, gid); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx,
		`insert into reports (id, owner_id, guild_id, visibility, status, created_at)
		 values ('oldhome', $1, $2, 'guild', 'complete', now() - interval '9 days')`, uid, gid); err != nil {
		t.Fatal(err)
	}

	h.actor = auth.Actor{UserID: uid, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	var view HomeView
	h.data(res, &view)
	if len(view.Reports) != 1 || view.Reports[0].ID != "recenthome" {
		t.Fatalf("reports = %+v, want only the 2-day-old one", view.Reports)
	}
}

func TestHomeReportsHidesPrivateAndUnlistedFromOthers(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	owner := seedUser(t, h.pool, "home-visibility-owner@example.com")
	seedCharacter(t, h.pool, gid, owner, "us/hardcore/homevisibilityowner", "member", true)
	viewer := seedUser(t, h.pool, "home-visibility-viewer@example.com")
	seedCharacter(t, h.pool, gid, viewer, "us/hardcore/homevisibilityviewer", "member", true)
	// GuildRank (which the home handler calls to determine `verified`)
	// reads guild_members, not guild_characters; seedCharacter alone
	// leaves guild_members empty, so the viewer would read as
	// unverified without this.
	recomputeMembership(t, h.pool, gid, viewer)

	for _, r := range []struct{ id, visibility string }{
		{"homevispublic", "public"}, {"homevisguild", "guild"},
		{"homevisprivate", "private"}, {"homevisunlisted", "unlisted"},
	} {
		if _, err := h.pool.Exec(ctx,
			`insert into reports (id, owner_id, guild_id, visibility, status, created_at)
			 values ($1, $2, $3, $4, 'complete', now())`, r.id, owner, gid, r.visibility); err != nil {
			t.Fatal(err)
		}
	}

	h.actor = auth.Actor{UserID: viewer, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	var view HomeView
	h.data(res, &view)
	seen := map[string]bool{}
	for _, r := range view.Reports {
		seen[r.ID] = true
	}
	if !seen["homevispublic"] || !seen["homevisguild"] {
		t.Fatalf("a verified viewer should see public and guild reports: %+v", view.Reports)
	}
	if seen["homevisprivate"] || seen["homevisunlisted"] {
		t.Fatalf("a verified viewer who is not the owner must never see private or unlisted reports: %+v", view.Reports)
	}

	h.actor = auth.Actor{UserID: owner, Role: "user", Method: "session"}
	res = h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	h.data(res, &view)
	seen = map[string]bool{}
	for _, r := range view.Reports {
		seen[r.ID] = true
	}
	for _, id := range []string{"homevispublic", "homevisguild", "homevisprivate", "homevisunlisted"} {
		if !seen[id] {
			t.Fatalf("the owner should see every one of their own reports regardless of visibility: %+v", view.Reports)
		}
	}
}

func TestHomeReportsHidesGuildVisibilityFromAnUnverifiedMember(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	owner := seedUser(t, h.pool, "home-unverified-owner@example.com")
	seedCharacter(t, h.pool, gid, owner, "us/hardcore/homeunverifiedowner", "member", true)
	unverified := seedUser(t, h.pool, "home-unverified-viewer@example.com")
	seedCharacter(t, h.pool, gid, unverified, "us/hardcore/homeunverifiedviewer", "member", false)

	for _, r := range []struct{ id, visibility string }{
		{"unverifiedpublic", "public"}, {"unverifiedguild", "guild"},
	} {
		if _, err := h.pool.Exec(ctx,
			`insert into reports (id, owner_id, guild_id, visibility, status, created_at)
			 values ($1, $2, $3, $4, 'complete', now())`, r.id, owner, gid, r.visibility); err != nil {
			t.Fatal(err)
		}
	}

	h.actor = auth.Actor{UserID: unverified, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	var view HomeView
	h.data(res, &view)
	seen := map[string]bool{}
	for _, r := range view.Reports {
		seen[r.ID] = true
	}
	if !seen["unverifiedpublic"] {
		t.Fatal("an unverified member should still see public reports")
	}
	if seen["unverifiedguild"] {
		t.Fatal("an unverified member must not see guild-visibility reports")
	}
}

func TestHomeExposesClaimState(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	member := seedUser(t, h.pool, "home-claim-state@example.com")
	seedCharacter(t, h.pool, gid, member, "us/hardcore/homeclaimstate", "member", true)
	claimant := seedUser(t, h.pool, "home-claim-state-claimant@example.com")
	seedCharacter(t, h.pool, gid, claimant, "us/hardcore/homeclaimstateclaimant", "leader", true)
	if _, err := h.pool.Exec(ctx, `update guilds set claimed_by = $1 where id = $2`, claimant, gid); err != nil {
		t.Fatal(err)
	}
	h.actor = auth.Actor{UserID: member, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	var view HomeView
	h.data(res, &view)
	if view.Claim.State != "claimed" {
		t.Fatalf("claim.state = %q, want claimed", view.Claim.State)
	}
}

// TestHomeReportsExposeZone is item 7 (fourth security review
// response): each report in the guild home list carries its zone, the
// same way every other report list in this codebase already does.
func TestHomeReportsExposeZone(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	uid := seedUser(t, h.pool, "home-zone@example.com")
	seedCharacter(t, h.pool, gid, uid, "us/hardcore/homezone", "member", true)

	if _, err := h.pool.Exec(ctx,
		`insert into reports (id, owner_id, guild_id, visibility, status, zone, created_at)
		 values ('zonedhome', $1, $2, 'guild', 'complete', 'Blackrock Spire', now() - interval '1 day')`,
		uid, gid); err != nil {
		t.Fatal(err)
	}

	h.actor = auth.Actor{UserID: uid, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	var view HomeView
	h.data(res, &view)
	if len(view.Reports) != 1 || view.Reports[0].Zone != "Blackrock Spire" {
		t.Fatalf("reports = %+v, want one report with zone Blackrock Spire", view.Reports)
	}
}

// TestHomeRosterMayRemoveMatchesTheRemoveRule is item 7 (fourth
// security review response): may_remove on each roster row is computed
// from the exact rank-protects-rank rule the DELETE route enforces, from
// the viewpoint of whoever is asking - a verified officer viewer may
// remove a plain member's row (verifiedOfficer), never the guild
// master's leader-rank row (rank protects rank), and always their own
// row regardless of rank (self).
func TestHomeRosterMayRemoveMatchesTheRemoveRule(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")

	viewer := seedUser(t, h.pool, "may-remove-viewer@example.com")
	seedCharacter(t, h.pool, gid, viewer, "us/hardcore/mayremoveviewer", "officer", true)
	seedExport(t, h.pool, viewer, "us/hardcore/mayremoveviewer", "us", "hardcore", "mayremoveviewer")
	recomputeMembership(t, h.pool, gid, viewer)

	member := seedUser(t, h.pool, "may-remove-member@example.com")
	seedCharacter(t, h.pool, gid, member, "us/hardcore/mayremovemember", "member", false)
	seedExport(t, h.pool, member, "us/hardcore/mayremovemember", "us", "hardcore", "mayremovemember")
	recomputeMembership(t, h.pool, gid, member)

	leader := seedUser(t, h.pool, "may-remove-leader@example.com")
	seedCharacter(t, h.pool, gid, leader, "us/hardcore/mayremoveleader", "leader", true)
	seedExport(t, h.pool, leader, "us/hardcore/mayremoveleader", "us", "hardcore", "mayremoveleader")
	recomputeMembership(t, h.pool, gid, leader)
	if _, err := h.pool.Exec(ctx, `update guilds set claimed_by = $1 where id = $2`, leader, gid); err != nil {
		t.Fatal(err)
	}

	h.actor = auth.Actor{UserID: viewer, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	var view HomeView
	h.data(res, &view)
	byKey := map[string]RosterRow{}
	for _, row := range view.Roster {
		byKey[row.CharacterKey] = row
	}
	if row := byKey["us/hardcore/mayremoveviewer"]; !row.MayRemove {
		t.Fatalf("the viewer's own row: may_remove = %v, want true (self)", row.MayRemove)
	}
	if row := byKey["us/hardcore/mayremovemember"]; !row.MayRemove {
		t.Fatalf("a plain member row seen by a verified officer: may_remove = %v, want true", row.MayRemove)
	}
	if row := byKey["us/hardcore/mayremoveleader"]; row.MayRemove {
		t.Fatalf("the guild master's leader-rank row: may_remove = %v, want false (rank protects rank)", row.MayRemove)
	}
}
