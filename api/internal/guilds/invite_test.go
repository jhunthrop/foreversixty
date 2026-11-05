// api/internal/guilds/invite_test.go
package guilds

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

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

func TestPeekInviteNamesTheGuildWithoutJoiningAndRefusesBadTokens(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	token, _, err := h.store.RotateInvite(ctx, gid)
	if err != nil {
		t.Fatal(err)
	}

	// Signed out: the zero actor still gets the card.
	h.actor = auth.Actor{}
	res := h.do(http.MethodGet, "/v1/guild-invites/"+token, "")
	var got struct {
		Guild struct {
			ID      int64  `json:"id"`
			Name    string `json:"name"`
			Region  string `json:"region"`
			Ruleset string `json:"ruleset"`
		} `json:"guild"`
	}
	h.data(res, &got)
	if got.Guild.ID != gid || got.Guild.Name != "Forever" || got.Guild.Region == "" || got.Guild.Ruleset == "" {
		t.Fatalf("peek = %+v, want guild %d Forever with a realm", got.Guild, gid)
	}
	var rows int
	if err := h.pool.QueryRow(ctx, `select count(*) from guild_characters where guild_id = $1`, gid).Scan(&rows); err != nil || rows != 0 {
		t.Fatalf("peek joined someone: rows = %d, err = %v", rows, err)
	}

	// Rotated-away and invented tokens answer the same 404.
	if _, _, err := h.store.RotateInvite(ctx, gid); err != nil {
		t.Fatal(err)
	}
	for _, bad := range []string{token, "fsg_totallymadeup"} {
		res := h.do(http.MethodGet, "/v1/guild-invites/"+bad, "")
		res.Body.Close()
		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("peek %q = %d, want 404", bad, res.StatusCode)
		}
	}
}

// peekAs reads the invite preview as actor.
func peekAs(t *testing.T, h *httpHarness, token string, actor auth.Actor) (members int, joinsAs *JoinsAs) {
	t.Helper()
	h.actor = actor
	var got struct {
		Members int      `json:"members"`
		JoinsAs *JoinsAs `json:"joins_as"`
	}
	h.data(h.do(http.MethodGet, "/v1/guild-invites/"+token, ""), &got)
	return got.Members, got.JoinsAs
}

func seedOwned(t *testing.T, h *httpHarness, userID int64, key, name, class string, level int) {
	t.Helper()
	if _, err := h.pool.Exec(context.Background(),
		`insert into characters (key, region, ruleset, name, class, race, level, user_id)
		 values ($1, 'us', 'hardcore', $2, $3, 'night-elf', $4, $5)`, key, name, class, level, userID); err != nil {
		t.Fatal(err)
	}
}

func TestPeekInviteSaysWhichCharacterJoins(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	if _, err := h.pool.Exec(ctx, `update guilds set region = 'us', ruleset = 'hardcore' where id = $1`, gid); err != nil {
		t.Fatal(err)
	}
	token, _, err := h.store.RotateInvite(ctx, gid)
	if err != nil {
		t.Fatal(err)
	}
	verified := seedUser(t, h.pool, "peek-verified@example.com")
	seedCharacter(t, h.pool, gid, verified, "us/hardcore/peekverified", "member", true)
	syncMembership(t, h.pool, gid, verified)

	members, joins := peekAs(t, h, token, auth.Actor{})
	if members != 1 || joins != nil {
		t.Fatalf("signed out: members = %d, joins_as = %+v; want the count and no character", members, joins)
	}

	// No characters at all.
	bare := seedUser(t, h.pool, "peek-bare@example.com")
	session := func(id int64) auth.Actor { return auth.Actor{UserID: id, Role: "user", Method: "session"} }
	if _, joins = peekAs(t, h, token, session(bare)); joins != nil {
		t.Fatalf("no characters: joins_as = %+v, want nil", joins)
	}

	// Characters, none on the roster: the main with an export, else the newest export.
	viewer := seedUser(t, h.pool, "peek-viewer@example.com")
	seedOwned(t, h, viewer, "us/hardcore/peekaaa", "Peekaaa", "mage", 30)
	seedOwned(t, h, viewer, "us/hardcore/peekbbb", "Peekbbb", "rogue", 20)
	seedOwned(t, h, viewer, "us/hardcore/peekccc", "Peekccc", "priest", 60)
	seedExport(t, h.pool, viewer, "us/hardcore/peekbbb", "us", "hardcore", "Peekbbb")
	_, joins = peekAs(t, h, token, session(viewer))
	if joins == nil || joins.Name != "Peekbbb" || joins.OnRoster {
		t.Fatalf("none on roster: joins_as = %+v, want the newest export Peekbbb, off the roster", joins)
	}

	// The snapshot names two of them: the higher level joins, with its race and class.
	seedRosterSnapshot(t, h.pool, gid, "us/hardcore/peeksender",
		`[{"name":"Peekaaa-Realm","level":30,"spec":"mage-frost","gear_gap":0,"build":"","seen_at":"2026-11-04 14:01"},
		  {"name":"peekccc","level":60,"spec":"priest-holy","gear_gap":0,"build":"","seen_at":"2026-11-04 14:01"}]`,
		time.Now())
	_, joins = peekAs(t, h, token, session(viewer))
	if joins == nil || joins.Name != "Peekccc" || !joins.OnRoster || *joins.Level != 60 ||
		*joins.Class != "priest" || *joins.Race != "night-elf" {
		t.Fatalf("on roster: joins_as = %+v, want Peekccc the level 60 Night Elf Priest", joins)
	}
}

// An account whose member-rank character already waits on the roster and then joins by invite is
// approved at once: the character carries the invite's verification and no placeholder is left.
func TestAcceptInviteApprovesTheAccountsWaitingCharacter(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool, "Forever")
	uid := seedUser(t, pool, "waiting@example.com")
	seedCharacter(t, pool, gid, uid, "us/hardcore/waitsalot", "member", false)
	acceptAFreshInvite(t, s, gid, uid)
	if by := inviteVerifiedBy(t, pool, "us/hardcore/waitsalot"); by == nil || *by != "invite" {
		t.Fatalf("verified_by = %v, want invite", by)
	}
	if n := placeholderCount(t, pool, uid); n != 0 {
		t.Fatalf("placeholder rows = %d, want 0", n)
	}
	if rank := verifiedRank(t, pool, gid, uid); rank != "member" {
		t.Fatalf("guild_members.rank = %q, want member", rank)
	}
}

// An officer-rank character is the export's own word and an invite link can be passed on, so it
// keeps waiting for an officer; the placeholder keeps the account a verified member meanwhile.
func TestAcceptInviteLeavesAnOfficerRankWaiting(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool, "Forever")
	uid := seedUser(t, pool, "officer@example.com")
	seedCharacter(t, pool, gid, uid, "us/hardcore/claimsofficer", "officer", false)
	acceptAFreshInvite(t, s, gid, uid)
	if by := inviteVerifiedBy(t, pool, "us/hardcore/claimsofficer"); by != nil {
		t.Fatalf("verified_by = %q, want nil until an officer approves", *by)
	}
	if n := placeholderCount(t, pool, uid); n != 1 {
		t.Fatalf("placeholder rows = %d, want 1", n)
	}
	if rank := verifiedRank(t, pool, gid, uid); rank != "member" {
		t.Fatalf("guild_members.rank = %q, want member (never officer by invite)", rank)
	}
}

func acceptAFreshInvite(t *testing.T, s *Store, gid, uid int64) {
	t.Helper()
	token, _, err := s.RotateInvite(context.Background(), gid)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := s.AcceptInvite(context.Background(), token, uid); err != nil {
		t.Fatal(err)
	}
}

func inviteVerifiedBy(t *testing.T, pool *pgxpool.Pool, key string) *string {
	t.Helper()
	var by *string
	if err := pool.QueryRow(context.Background(),
		`select verified_by from guild_characters where character_key = $1`, key).Scan(&by); err != nil {
		t.Fatal(err)
	}
	return by
}

func placeholderCount(t *testing.T, pool *pgxpool.Pool, uid int64) int {
	t.Helper()
	var n int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from guild_characters where character_key = $1`, syntheticKey(uid)).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

func verifiedRank(t *testing.T, pool *pgxpool.Pool, gid, uid int64) string {
	t.Helper()
	var rank string
	_ = pool.QueryRow(context.Background(),
		`select rank from guild_members where guild_id = $1 and user_id = $2 and verified_at is not null`,
		gid, uid).Scan(&rank)
	return rank
}
