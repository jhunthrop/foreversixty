package addon

import (
	"context"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// joinByInviteAndSync makes a guild, accepts its invite as the harness owner and syncs one
// character into it at rankIndex; it answers the guild id.
func joinByInviteAndSync(t *testing.T, h *harness, rankIndex string) int64 {
	t.Helper()
	ctx := context.Background()
	var guildID int64
	if err := h.pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'hardcore', 'Forever') returning id`).Scan(&guildID); err != nil {
		t.Fatal(err)
	}
	gs := &guilds.Store{Pool: h.pool}
	token, _, err := gs.RotateInvite(ctx, guildID, nil)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := gs.AcceptInvite(ctx, token, h.owner); err != nil {
		t.Fatal(err)
	}
	if err := h.store.PutExports(ctx, h.owner, []Export{{Name: "Baelgrim", Region: "us", Ruleset: "hardcore",
		Export: "FS1:1.60.1.70790:warrior:tauren:0/0/0:|guild=Forever:" + rankIndex}}); err != nil {
		t.Fatal(err)
	}
	return guildID
}

func characterVerifiedBy(t *testing.T, h *harness, guildID int64) *string {
	t.Helper()
	var verifiedBy *string
	if err := h.pool.QueryRow(context.Background(),
		`select verified_by from guild_characters where guild_id = $1 and character_key not like 'account:%'`,
		guildID).Scan(&verifiedBy); err != nil {
		t.Fatal(err)
	}
	return verifiedBy
}

func placeholderRows(t *testing.T, h *harness, guildID int64) int {
	t.Helper()
	var n int
	if err := h.pool.QueryRow(context.Background(),
		`select count(*) from guild_characters where guild_id = $1 and character_key like 'account:%'`,
		guildID).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

func membership(t *testing.T, h *harness, guildID int64) (string, bool) {
	t.Helper()
	var rank string
	var verified bool
	if err := h.pool.QueryRow(context.Background(),
		`select rank, verified_at is not null from guild_members where guild_id = $1 and user_id = $2`,
		guildID, h.owner).Scan(&rank, &verified); err != nil {
		t.Fatal(err)
	}
	return rank, verified
}

// An account that joined by invite and then syncs a member-rank character into the guild is
// approved, not "Waiting for approval": the invite's verification moves onto the character and the
// placeholder row goes (round 3 e2e, item 2).
func TestPutExportsCarriesAnInviteVerificationOntoTheSyncedCharacter(t *testing.T) {
	h := newHarness(t)
	guildID := joinByInviteAndSync(t, h, "5")
	if by := characterVerifiedBy(t, h, guildID); by == nil || *by != "invite" {
		t.Fatalf("verified_by = %v, want invite", by)
	}
	if n := placeholderRows(t, h, guildID); n != 0 {
		t.Fatalf("placeholder rows = %d, want 0 once a character stands in", n)
	}
	if rank, verified := membership(t, h, guildID); rank != "member" || !verified {
		t.Fatalf("guild_members = %s/%v, want member/verified", rank, verified)
	}
}

// An officer rank is the export's own word and an invite link can be passed on, so an
// officer-rank character still waits for an officer; the placeholder keeps the account a verified
// member meanwhile (round 3: an invite never hands out officer tools).
func TestPutExportsKeepsAnInvitedOfficerRankWaiting(t *testing.T) {
	h := newHarness(t)
	guildID := joinByInviteAndSync(t, h, "1")
	if by := characterVerifiedBy(t, h, guildID); by != nil {
		t.Fatalf("verified_by = %v, want nil until an officer approves the officer rank", *by)
	}
	if n := placeholderRows(t, h, guildID); n != 1 {
		t.Fatalf("placeholder rows = %d, want 1 (the account stays a verified member)", n)
	}
	if rank, verified := membership(t, h, guildID); rank != "member" || !verified {
		t.Fatalf("guild_members = %s/%v, want member/verified", rank, verified)
	}
}

// removeAndRejoin has an officer remove the harness owner's account (guilds.Store.Leave, the
// remove route's write), then accepts a fresh invite with the export resync wired.
func removeAndRejoin(t *testing.T, h *harness, guildID int64) {
	t.Helper()
	ctx := context.Background()
	gs := &guilds.Store{Pool: h.pool, Exports: h.store}
	if err := gs.Leave(ctx, guildID, h.owner); err != nil {
		t.Fatal(err)
	}
	if n := placeholderRows(t, h, guildID); n != 0 {
		t.Fatalf("placeholder rows after removal = %d", n)
	}
	token, _, err := gs.RotateInvite(ctx, guildID, nil)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := gs.AcceptInvite(ctx, token, h.owner); err != nil {
		t.Fatal(err)
	}
}

// A removed member who rejoins by invite gets his synced character back at once, verified by
// the invite, with no placeholder left (guild round 3, item 3).
func TestAcceptInviteReattachesARemovedMembersSyncedCharacter(t *testing.T) {
	h := newHarness(t)
	guildID := joinByInviteAndSync(t, h, "5")
	removeAndRejoin(t, h, guildID)
	if by := characterVerifiedBy(t, h, guildID); by == nil || *by != "invite" {
		t.Fatalf("verified_by = %v, want the rejoined character verified by invite", by)
	}
	if n := placeholderRows(t, h, guildID); n != 0 {
		t.Fatalf("placeholder rows = %d, want 0 once the character is back", n)
	}
	if rank, verified := membership(t, h, guildID); rank != "member" || !verified {
		t.Fatalf("guild_members = %s/%v, want member/verified", rank, verified)
	}
}

// An officer-rank character comes back but still waits for an officer; the placeholder keeps
// the account a verified member (the invite rule).
func TestAcceptInviteReattachesAnOfficerRankStillWaiting(t *testing.T) {
	h := newHarness(t)
	guildID := joinByInviteAndSync(t, h, "1")
	removeAndRejoin(t, h, guildID)
	if by := characterVerifiedBy(t, h, guildID); by != nil {
		t.Fatalf("verified_by = %v, want nil until an officer approves", *by)
	}
	if n := placeholderRows(t, h, guildID); n != 1 {
		t.Fatalf("placeholder rows = %d, want 1", n)
	}
}

// An export that names another guild, or another realm's guild of the same name, is left alone.
func TestAcceptInviteLeavesExportsNamingAnotherGuild(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	var guildID int64
	if err := h.pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'hardcore', 'Forever') returning id`).Scan(&guildID); err != nil {
		t.Fatal(err)
	}
	if err := h.store.PutExports(ctx, h.owner, []Export{
		{Name: "Baelgrim", Region: "us", Ruleset: "hardcore", Export: "FS1:1.60.1.70790:warrior:tauren:0/0/0:|guild=Elsewhere:5"},
		{Name: "Brightwing", Region: "eu", Ruleset: "hardcore", Export: "FS1:1.60.1.70790:priest:human:0/0/0:|guild=Forever:5"},
	}); err != nil {
		t.Fatal(err)
	}
	gs := &guilds.Store{Pool: h.pool, Exports: h.store}
	token, _, err := gs.RotateInvite(ctx, guildID, nil)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := gs.AcceptInvite(ctx, token, h.owner); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := h.pool.QueryRow(ctx, `select count(*) from guild_characters where guild_id = $1 and character_key not like 'account:%'`,
		guildID).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("characters attached = %d, want 0", n)
	}
}

// Without an invite nothing changes: a synced character still waits for an officer.
func TestPutExportsWithoutAnInviteStillWaits(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	if err := h.store.PutExports(ctx, h.owner, []Export{{Name: "Baelgrim", Region: "us", Ruleset: "hardcore",
		Export: "FS1:1.60.1.70790:warrior:tauren:0/0/0:|guild=Forever:5"}}); err != nil {
		t.Fatal(err)
	}
	var verified bool
	if err := h.pool.QueryRow(ctx, `select verified_at is not null from guild_characters`).Scan(&verified); err != nil {
		t.Fatal(err)
	}
	if verified {
		t.Fatal("a character with no invite behind it was verified")
	}
}
