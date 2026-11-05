package addon

import (
	"context"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// An account that joined by invite and then syncs a character into the guild is approved, not
// "Waiting for approval": the invite's verification moves onto the character, the placeholder
// row goes, and an officer-rank character comes back an officer (round 3 e2e, item 2).
func TestPutExportsCarriesAnInviteVerificationOntoTheSyncedCharacter(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	var guildID int64
	if err := h.pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'hardcore', 'Forever') returning id`).Scan(&guildID); err != nil {
		t.Fatal(err)
	}
	gs := &guilds.Store{Pool: h.pool}
	token, _, err := gs.RotateInvite(ctx, guildID)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := gs.AcceptInvite(ctx, token, h.owner); err != nil {
		t.Fatal(err)
	}
	if err := h.store.PutExports(ctx, h.owner, []Export{{Name: "Baelgrim", Region: "us", Ruleset: "hardcore",
		Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Forever:1"}}); err != nil {
		t.Fatal(err)
	}
	var verifiedBy *string
	if err := h.pool.QueryRow(ctx,
		`select verified_by from guild_characters where guild_id = $1 and character_key not like 'account:%'`,
		guildID).Scan(&verifiedBy); err != nil {
		t.Fatal(err)
	}
	if verifiedBy == nil || *verifiedBy != "invite" {
		t.Fatalf("verified_by = %v, want invite", verifiedBy)
	}
	var placeholders int
	if err := h.pool.QueryRow(ctx,
		`select count(*) from guild_characters where guild_id = $1 and character_key like 'account:%'`,
		guildID).Scan(&placeholders); err != nil {
		t.Fatal(err)
	}
	if placeholders != 0 {
		t.Fatalf("placeholder rows = %d, want 0 once a character stands in", placeholders)
	}
	var rank string
	var verified bool
	if err := h.pool.QueryRow(ctx,
		`select rank, verified_at is not null from guild_members where guild_id = $1 and user_id = $2`,
		guildID, h.owner).Scan(&rank, &verified); err != nil {
		t.Fatal(err)
	}
	if rank != "officer" || !verified {
		t.Fatalf("guild_members = %s/%v, want officer/verified", rank, verified)
	}
}

// Without an invite nothing changes: a synced character still waits for an officer.
func TestPutExportsWithoutAnInviteStillWaits(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	if err := h.store.PutExports(ctx, h.owner, []Export{{Name: "Baelgrim", Region: "us", Ruleset: "hardcore",
		Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Forever:5"}}); err != nil {
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
