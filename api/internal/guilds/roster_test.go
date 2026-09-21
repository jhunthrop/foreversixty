// api/internal/guilds/roster_test.go
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// syncMembership runs RecomputeMembership for (guildID, userID) outside
// any other transaction, the same way store_test.go does. seedCharacter
// only inserts the raw guild_characters row it's given; guild_members —
// which auth.Store.GuildRank (and so verifiedOfficerOrLeader) reads — is
// derived, never written directly, so an HTTP-level test that depends on
// an account's officer/leader standing must call this after seeding a
// verified character for that account.
func syncMembership(t *testing.T, pool *pgxpool.Pool, guildID, userID int64) {
	t.Helper()
	ctx := context.Background()
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if err := RecomputeMembership(ctx, tx, guildID, &userID); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestApproveCharacterNeedsAVerifiedOfficer(t *testing.T) {
	h := newHTTPHarness(t)
	uid := seedUser(t, h.pool, "roster-owner@example.com")
	gid := seedGuild(t, h.pool, "Forever")
	seedCharacter(t, h.pool, gid, uid, "us/hardcore/rosterowner", "member", false)

	h.actor = auth.Actor{UserID: uid, Role: "user", Method: "session"}
	res := h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/characters/us/hardcore/rosterowner/approve", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("a plain member approving = %d, want 403", res.StatusCode)
	}

	officer := seedUser(t, h.pool, "roster-officer@example.com")
	seedCharacter(t, h.pool, gid, officer, "us/hardcore/rosterofficer", "officer", true)
	syncMembership(t, h.pool, gid, officer)
	h.actor = auth.Actor{UserID: officer, Role: "user", Method: "session"}
	res = h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/characters/us/hardcore/rosterowner/approve", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("a verified officer approving = %d, want 200", res.StatusCode)
	}
	var verified bool
	h.pool.QueryRow(context.Background(),
		`select verified_at is not null from guild_characters where character_key = 'us/hardcore/rosterowner'`).Scan(&verified)
	if !verified {
		t.Fatal("the character should be verified after approve")
	}
}

func TestRemoveCharacterIsTheOwnerOrAVerifiedOfficer(t *testing.T) {
	h := newHTTPHarness(t)
	uid := seedUser(t, h.pool, "self-remove@example.com")
	gid := seedGuild(t, h.pool, "Forever")
	seedCharacter(t, h.pool, gid, uid, "us/hardcore/selfremove", "member", true)

	stranger := seedUser(t, h.pool, "remove-stranger@example.com")
	h.actor = auth.Actor{UserID: stranger, Role: "user", Method: "session"}
	res := h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/characters/us/hardcore/selfremove", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("a stranger removing = %d, want 403", res.StatusCode)
	}

	h.actor = auth.Actor{UserID: uid, Role: "user", Method: "session"}
	res = h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/characters/us/hardcore/selfremove", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("the character's own account removing itself = %d, want 200", res.StatusCode)
	}
	var n int
	h.pool.QueryRow(context.Background(),
		`select count(*) from guild_characters where character_key = 'us/hardcore/selfremove'`).Scan(&n)
	if n != 0 {
		t.Fatal("the row should be gone")
	}
}

func TestConsentPatchAndLeave(t *testing.T) {
	h := newHTTPHarness(t)
	uid := seedUser(t, h.pool, "consent@example.com")
	gid := seedGuild(t, h.pool, "Forever")
	seedCharacter(t, h.pool, gid, uid, "us/hardcore/consenter", "member", true)
	syncMembership(t, h.pool, gid, uid)
	h.actor = auth.Actor{UserID: uid, Role: "user", Method: "session"}

	res := h.do(http.MethodPatch, fmt.Sprintf("/v1/guilds/%d/members/me", gid), `{"consent":"gear_bags"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("consent patch = %d, want 200", res.StatusCode)
	}
	var consent string
	h.pool.QueryRow(context.Background(),
		`select consent from guild_members where guild_id = $1 and user_id = $2`, gid, uid).Scan(&consent)
	if consent != "gear_bags" {
		t.Fatalf("consent = %q, want gear_bags", consent)
	}

	res = h.do(http.MethodPatch, fmt.Sprintf("/v1/guilds/%d/members/me", gid), `{"consent":"nonsense"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("bad consent = %d, want 400", res.StatusCode)
	}

	res = h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/members/me", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("leave = %d, want 200", res.StatusCode)
	}
	var n int
	h.pool.QueryRow(context.Background(),
		`select count(*) from guild_members where guild_id = $1 and user_id = $2`, gid, uid).Scan(&n)
	if n != 0 {
		t.Fatal("guild_members row should be gone after leaving")
	}
	res = h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/members/me", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("leaving twice = %d, want 404", res.StatusCode)
	}
}
