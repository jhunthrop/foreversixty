// api/internal/directory/claims_test.go
package directory

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// claimWorld is a guild with its Blizzard id, three directory members (one the guild master)
// and one Battle.net account already holding the third.
type claimWorld struct {
	pool    *pgxpool.Pool
	c       *Claims
	guildID int64
	bnetOwn int64
}

func newClaimWorld(t *testing.T) claimWorld {
	t.Helper()
	pool := testPool(t)
	ctx := context.Background()
	w := claimWorld{pool: pool, c: &Claims{Pool: pool}}
	if err := pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name, bnet_guild_id) values ('us', 'pvp', 'Ember Wardens', 77) returning id`).
		Scan(&w.guildID); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `
		insert into character_directory (region, bnet_character_id, name, slug, class, level, guild_bnet_id, guild_name, guild_rank, ruleset, ruleset_source)
		values ('us', 101, 'Vela Thornbrook', 'vela-thornbrook', 'mage', 41, 77, 'Ember Wardens', 0, 'pvp', 'guild'),
		       ('us', 102, 'Orrin Ashmantle', 'orrin-ashmantle', 'warrior', 12, 77, 'Ember Wardens', 3, null, null),
		       ('us', 103, 'Sel Duskwater', 'sel-duskwater', 'priest', 30, 77, 'Ember Wardens', 4, 'pvp', 'guild'),
		       ('us', 104, 'Ada Quill', 'ada-quill', 'rogue', 20, 88, 'Other Guild', 4, 'pvp', 'guild')`); err != nil {
		t.Fatal(err)
	}
	if err := pool.QueryRow(ctx, `insert into users (email) values ('owner@example.com') returning id`).Scan(&w.bnetOwn); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `insert into characters (key, region, ruleset, name, user_id, bnet_character_id, source)
		values ('us/pvp/sel-duskwater', 'us', 'pvp', 'Sel Duskwater', $1, 103, 'bnet')`, w.bnetOwn); err != nil {
		t.Fatal(err)
	}
	return w
}

func (w claimWorld) discordUser(t *testing.T, id string) int64 {
	t.Helper()
	uid, created, err := w.c.EnsureDiscordAccount(context.Background(), id, "a raider")
	if err != nil || !created {
		t.Fatalf("discord account = %d %v %v", uid, created, err)
	}
	return uid
}

func (w claimWorld) one(t *testing.T, q string, args ...any) string {
	t.Helper()
	var s string
	if err := w.pool.QueryRow(context.Background(), q, args...).Scan(&s); err != nil {
		t.Fatalf("%s: %v", q, err)
	}
	return s
}

// The picks are the guild's directory roster no account holds, highest level first.
func TestGuildPicksLeaveOutHeldCharacters(t *testing.T) {
	w := newClaimWorld(t)
	picks, err := w.c.GuildPicks(context.Background(), w.guildID, "", 25)
	if err != nil {
		t.Fatal(err)
	}
	if len(picks) != 2 || picks[0].BnetID != 101 || picks[1].BnetID != 102 {
		t.Fatalf("picks = %+v", picks)
	}
	if typed, _ := w.c.GuildPicks(context.Background(), w.guildID, "ORR", 25); len(typed) != 1 || typed[0].Name != "Orrin Ashmantle" {
		t.Fatalf("typed picks = %+v", typed)
	}
}

// A Discord-only account claims a roster character at member rank, even the guild master's;
// a second claimer is refused; a Battle.net account cannot pick; another guild's member is
// not on the roster.
func TestClaimForDiscord(t *testing.T) {
	w := newClaimWorld(t)
	ctx := context.Background()
	uid := w.discordUser(t, "d-100")
	again, created, err := w.c.EnsureDiscordAccount(ctx, "d-100", "a raider")
	if err != nil || created || again != uid {
		t.Fatalf("second ensure = %d %v %v", again, created, err)
	}
	claim, err := w.c.ClaimForDiscord(ctx, w.guildID, uid, 101)
	if err != nil || claim.Key != "us/pvp/vela-thornbrook" || !claim.Created || claim.Name != "Vela Thornbrook" {
		t.Fatalf("claim = %+v, %v", claim, err)
	}
	if rank := w.one(t, `select rank || '/' || source || '/' || verified_by from guild_characters where character_key = $1`, claim.Key); rank != "member/discord/discord" {
		t.Fatalf("guild row = %q, want a verified member row from Discord", rank)
	}
	if m := w.one(t, `select rank from guild_members where guild_id = $1 and user_id = $2 and verified_at is not null`, w.guildID, uid); m != "member" {
		t.Fatalf("membership = %q", m)
	}
	if again, err := w.c.ClaimForDiscord(ctx, w.guildID, uid, 101); err != nil || again.Created {
		t.Fatalf("re-pick = %+v, %v; want held, not created", again, err)
	}
	other := w.discordUser(t, "d-200")
	if _, err := w.c.ClaimForDiscord(ctx, w.guildID, other, 101); !errors.Is(err, ErrTaken) {
		t.Fatalf("second claimer: %v, want ErrTaken", err)
	}
	if _, err := w.c.ClaimForDiscord(ctx, w.guildID, other, 103); !errors.Is(err, ErrTaken) {
		t.Fatalf("a Battle.net character: %v, want ErrTaken", err)
	}
	if _, err := w.c.ClaimForDiscord(ctx, w.guildID, other, 104); !errors.Is(err, ErrNotInGuild) {
		t.Fatalf("another guild's member: %v, want ErrNotInGuild", err)
	}
	if _, err := w.c.ClaimForDiscord(ctx, w.guildID, w.bnetOwn, 102); !errors.Is(err, ErrNotDiscordOnly) {
		t.Fatalf("a Battle.net account: %v, want ErrNotDiscordOnly", err)
	}
	// The guild's ruleset keys a character the directory has none for.
	if c, err := w.c.ClaimForDiscord(ctx, w.guildID, other, 102); err != nil || c.Key != "us/pvp/orrin-ashmantle" {
		t.Fatalf("claim without its own ruleset = %+v, %v", c, err)
	}
}

// A Discord-only account holds at most MaxDiscordPicks characters, and an officer's remove
// frees a pick for its real owner.
func TestPickCapAndOfficerRelease(t *testing.T) {
	w := newClaimWorld(t)
	ctx := context.Background()
	if _, err := w.pool.Exec(ctx, `insert into character_directory (region, bnet_character_id, name, slug, guild_bnet_id, guild_name, ruleset)
		select 'us', 200 + n, 'Alt Number' || chr(64 + n), 'alt-number' || chr(96 + n), 77, 'Ember Wardens', 'pvp'
		  from generate_series(1, 3) n`); err != nil {
		t.Fatal(err)
	}
	uid := w.discordUser(t, "d-100")
	for _, id := range []int64{201, 202, 203} {
		if _, err := w.c.ClaimForDiscord(ctx, w.guildID, uid, id); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := w.c.ClaimForDiscord(ctx, w.guildID, uid, 101); !errors.Is(err, ErrTooManyPicks) {
		t.Fatalf("fourth pick: %v, want ErrTooManyPicks", err)
	}
	store := &guilds.Store{Pool: w.pool}
	if err := store.RemoveCharacter(ctx, w.guildID, "us/pvp/alt-numbera"); err != nil {
		t.Fatal(err)
	}
	if n := w.one(t, `select count(*)::text from characters where key = 'us/pvp/alt-numbera'`); n != "0" {
		t.Fatal("an officer's remove left the pick held")
	}
	other := w.discordUser(t, "d-200")
	if _, err := w.c.ClaimForDiscord(ctx, w.guildID, other, 201); err != nil {
		t.Fatalf("the freed character: %v", err)
	}
}
