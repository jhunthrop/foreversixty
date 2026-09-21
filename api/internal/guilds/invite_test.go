// api/internal/guilds/invite_test.go
package guilds

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

func TestRotateInviteOnlyAVerifiedOfficerCanReadItAfterwards(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")

	token, rotatedAt, err := s.RotateInvite(ctx, gid)
	if err != nil || !strings.HasPrefix(token, InviteTokenPrefix) || rotatedAt.IsZero() {
		t.Fatalf("RotateInvite = %q, %v, %v", token, rotatedAt, err)
	}
	var hash []byte
	pool.QueryRow(ctx, `select invite_token_hash from guilds where id = $1`, gid).Scan(&hash)
	if len(hash) != 32 {
		t.Fatalf("stored hash length = %d, want 32 (sha256); the raw token must never be stored", len(hash))
	}
}

func TestAcceptInviteCreatesAVerifiedSyntheticRow(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	token, _, err := s.RotateInvite(ctx, gid)
	if err != nil {
		t.Fatal(err)
	}
	uid := seedUser(t, pool, "invited@example.com")

	result, err := s.AcceptInvite(ctx, token, uid)
	if err != nil || result.GuildID != gid || result.Rank != "member" {
		t.Fatalf("AcceptInvite = %+v, %v", result, err)
	}
	var source string
	var verified bool
	pool.QueryRow(ctx,
		`select source, verified_at is not null from guild_characters where character_key = $1`,
		syntheticKey(uid)).Scan(&source, &verified)
	if source != "invite" || !verified {
		t.Fatalf("source = %q verified = %v, want invite/true", source, verified)
	}
}

func TestAcceptInviteAnswersTheSame404ForUnknownAndRotatedTokens(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	oldToken, _, _ := s.RotateInvite(ctx, gid)
	if _, _, err := s.RotateInvite(ctx, gid); err != nil {
		t.Fatal(err)
	}
	uid := seedUser(t, pool, "trying-old@example.com")

	_, errRotated := s.AcceptInvite(ctx, oldToken, uid)
	_, errUnknown := s.AcceptInvite(ctx, "fsg_totallymadeup", uid)
	if !errors.Is(errRotated, ErrNotFound) || !errors.Is(errUnknown, ErrNotFound) {
		t.Fatalf("errRotated = %v, errUnknown = %v, want both ErrNotFound", errRotated, errUnknown)
	}
}

func TestAcceptInviteTwiceForDifferentGuildsTransfersTheSyntheticRow(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	g1 := seedGuild(t, pool, "First")
	g2 := seedGuild(t, pool, "Second")
	token1, _, _ := s.RotateInvite(ctx, g1)
	token2, _, _ := s.RotateInvite(ctx, g2)
	uid := seedUser(t, pool, "double-invite@example.com")

	if _, err := s.AcceptInvite(ctx, token1, uid); err != nil {
		t.Fatal(err)
	}
	if _, err := s.AcceptInvite(ctx, token2, uid); err != nil {
		t.Fatal(err)
	}
	var n int
	pool.QueryRow(ctx, `select count(*) from guild_characters where character_key = $1`, syntheticKey(uid)).Scan(&n)
	if n != 1 {
		t.Fatalf("synthetic rows for the account = %d, want 1 (the unique index allows only one guild at a time)", n)
	}
	var guildID int64
	pool.QueryRow(ctx, `select guild_id from guild_characters where character_key = $1`, syntheticKey(uid)).Scan(&guildID)
	if guildID != g2 {
		t.Fatalf("guild_id = %d, want %d (the second, most recent invite)", guildID, g2)
	}
}

func TestAcceptInviteRecordsVerifiedByInvite(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	token, _, err := s.RotateInvite(ctx, gid)
	if err != nil {
		t.Fatal(err)
	}
	uid := seedUser(t, pool, "invite-source@example.com")
	if _, err := s.AcceptInvite(ctx, token, uid); err != nil {
		t.Fatal(err)
	}
	var by string
	if err := pool.QueryRow(ctx,
		`select verified_by from guild_characters where character_key = $1`, syntheticKey(uid)).Scan(&by); err != nil {
		t.Fatal(err)
	}
	if by != "invite" {
		t.Fatalf("verified_by = %q, want invite", by)
	}
}

func TestRotateInviteIsFrozenDuringAContestedClaim(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	officer := seedUser(t, h.pool, "frozen-rotate@example.com")
	seedCharacter(t, h.pool, gid, officer, "us/hardcore/frozenrotate", "officer", true)
	syncMembership(t, h.pool, gid, officer)
	if _, err := h.pool.Exec(ctx, `update guilds set claim_contested_at = now() where id = $1`, gid); err != nil {
		t.Fatal(err)
	}
	h.actor = auth.Actor{UserID: officer, Role: "user", Method: "session"}
	res := h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/invite/rotate", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusConflict {
		t.Fatalf("invite rotate while contested = %d, want 409", res.StatusCode)
	}
}
