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

func TestAnOfficerCannotRemoveTheGuildMastersCharacter(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	gm := seedUser(t, h.pool, "protected-gm@example.com")
	seedCharacter(t, h.pool, gid, gm, "us/hardcore/protectedgm", "leader", true)
	if _, err := h.pool.Exec(ctx, `update guilds set claimed_by = $1 where id = $2`, gm, gid); err != nil {
		t.Fatal(err)
	}
	officer := seedUser(t, h.pool, "unprivileged-officer@example.com")
	seedCharacter(t, h.pool, gid, officer, "us/hardcore/unprivilegedofficer", "officer", true)

	h.actor = auth.Actor{UserID: officer, Role: "user", Method: "session"}
	res := h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/characters/us/hardcore/protectedgm", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("an officer removing the guild master = %d, want 403", res.StatusCode)
	}
	var claimedBy *int64
	h.pool.QueryRow(ctx, `select claimed_by from guilds where id = $1`, gid).Scan(&claimedBy)
	if claimedBy == nil || *claimedBy != gm {
		t.Fatalf("claimed_by = %v, want unchanged (%d) - the claim must stay intact", claimedBy, gm)
	}
	var n int
	h.pool.QueryRow(ctx, `select count(*) from guild_characters where character_key = 'us/hardcore/protectedgm'`).Scan(&n)
	if n != 1 {
		t.Fatal("the guild master's row must still exist")
	}
}

func TestTheAccountHoldingTheClaimCanRemoveAnOfficerButNotAnotherLeaderRow(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	gm := seedUser(t, h.pool, "gm-removes@example.com")
	seedCharacter(t, h.pool, gid, gm, "us/hardcore/gmremoves", "leader", true)
	if _, err := h.pool.Exec(ctx, `update guilds set claimed_by = $1 where id = $2`, gm, gid); err != nil {
		t.Fatal(err)
	}
	officer := seedUser(t, h.pool, "removable-officer@example.com")
	seedCharacter(t, h.pool, gid, officer, "us/hardcore/removableofficer", "officer", true)
	otherLeader := seedUser(t, h.pool, "other-leader@example.com")
	seedCharacter(t, h.pool, gid, otherLeader, "us/hardcore/otherleader", "leader", true)

	h.actor = auth.Actor{UserID: gm, Role: "user", Method: "session"}
	res := h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/characters/us/hardcore/removableofficer", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("the claim-holding account removing an officer = %d, want 200", res.StatusCode)
	}

	res = h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/characters/us/hardcore/otherleader", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("the claim-holding account removing a DIFFERENT leader-rank row = %d, want 403", res.StatusCode)
	}
}

func TestAContestedClaimFreezesApproveAndRemove(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	claimant := seedUser(t, h.pool, "frozen-claimant@example.com")
	seedCharacter(t, h.pool, gid, claimant, "us/hardcore/frozenclaimant", "leader", true)
	// verifiedOfficerOrLeader reads the derived guild_members table, not
	// guild_characters directly (see syncMembership's doc comment above);
	// the claimant needs its leader standing synced there before an
	// approve call can reach the contested check this test is after.
	syncMembership(t, h.pool, gid, claimant)
	if _, err := h.pool.Exec(ctx, `update guilds set claimed_by = $1, claim_contested_at = now() where id = $2`, claimant, gid); err != nil {
		t.Fatal(err)
	}
	unverified := seedUser(t, h.pool, "frozen-unverified@example.com")
	seedCharacter(t, h.pool, gid, unverified, "us/hardcore/frozenunverified", "member", false)

	h.actor = auth.Actor{UserID: claimant, Role: "user", Method: "session"}
	res := h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/characters/us/hardcore/frozenunverified/approve", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusConflict {
		t.Fatalf("approve while contested = %d, want 409", res.StatusCode)
	}

	otherOfficer := seedUser(t, h.pool, "frozen-other-officer@example.com")
	seedCharacter(t, h.pool, gid, otherOfficer, "us/hardcore/frozenotherofficer", "officer", true)
	// Same reason as claimant above: mayRemoveCharacter's member-rank leg
	// needs verifiedOfficerOrLeader to see otherOfficer as a real officer
	// before the handler ever reaches the contested-freeze check, or this
	// assertion would see 403 (unauthorized) instead of 409 (frozen).
	syncMembership(t, h.pool, gid, otherOfficer)
	h.actor = auth.Actor{UserID: otherOfficer, Role: "user", Method: "session"}
	res = h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/characters/us/hardcore/frozenunverified", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusConflict {
		t.Fatalf("remove by an officer while contested = %d, want 409", res.StatusCode)
	}

	// Self-removal still works during a contest - leaving is never frozen.
	h.actor = auth.Actor{UserID: unverified, Role: "user", Method: "session"}
	res = h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/characters/us/hardcore/frozenunverified", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("self-removal while contested = %d, want 200 (never frozen)", res.StatusCode)
	}
}

func TestApproveCharacterRecordsVerifiedByOfficer(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	uid := seedUser(t, h.pool, "verify-source@example.com")
	seedCharacter(t, h.pool, gid, uid, "us/hardcore/verifysource", "member", false)
	officer := seedUser(t, h.pool, "verify-source-officer@example.com")
	seedCharacter(t, h.pool, gid, officer, "us/hardcore/verifysourceofficer", "officer", true)
	// verifiedOfficerOrLeader reads guild_members, not guild_characters;
	// sync it so this officer's approve call actually clears that gate
	// (see TestApproveCharacterNeedsAVerifiedOfficer above for the same
	// established pattern).
	syncMembership(t, h.pool, gid, officer)

	h.actor = auth.Actor{UserID: officer, Role: "user", Method: "session"}
	res := h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/characters/us/hardcore/verifysource/approve", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("approve = %d, want 200", res.StatusCode)
	}
	var by string
	h.pool.QueryRow(ctx, `select verified_by from guild_characters where character_key = 'us/hardcore/verifysource'`).Scan(&by)
	if by != "officer" {
		t.Fatalf("verified_by = %q, want officer", by)
	}
}
