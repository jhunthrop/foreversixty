// api/internal/directory/handover_test.go
package directory

import (
	"context"
	"testing"
)

// signUp writes a sign-up for key by userID on a raid of the world's guild.
func (w claimWorld) signUp(t *testing.T, key string, userID int64) int64 {
	t.Helper()
	ctx := context.Background()
	var raidID int64
	if err := w.pool.QueryRow(ctx, `insert into guild_raids (guild_id, zone, title, starts_at, duration_minutes, created_by)
		values ($1, 'raid:onyxias-lair', 'Test raid', now() + interval '2 days', 180, $2) returning id`, w.guildID, w.bnetOwn).
		Scan(&raidID); err != nil {
		t.Fatal(err)
	}
	if _, err := w.pool.Exec(ctx, `insert into guild_raid_signups (raid_id, character_key, name, class, spec, role, state, signed_by, acted_by)
		values ($1, $2, 'Vela Thornbrook', 'mage', '', 'dps', 'in', $3, $3)`, raidID, key, userID); err != nil {
		t.Fatal(err)
	}
	return raidID
}

// A Battle.net sign-in that owns a character picked on Discord takes it over: the character,
// its guild row and its sign-up move to the Battle.net account, with nothing duplicated.
func TestTakeOverDiscordClaims(t *testing.T) {
	w := newClaimWorld(t)
	ctx := context.Background()
	picker := w.discordUser(t, "d-100")
	claim, err := w.c.ClaimForDiscord(ctx, w.guildID, picker, 101)
	if err != nil {
		t.Fatal(err)
	}
	w.signUp(t, claim.Key, picker)
	var owner int64
	if err := w.pool.QueryRow(ctx, `insert into users (email) values ('vela@example.com') returning id`).Scan(&owner); err != nil {
		t.Fatal(err)
	}
	tx, err := w.pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	moved, err := TakeOverDiscordClaims(ctx, tx, owner, "us", 101)
	if err != nil || moved != 1 {
		t.Fatalf("moved = %d, %v", moved, err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatal(err)
	}
	for _, q := range []string{
		`select count(*)::text from characters where key = $1 and user_id = $2`,
		`select count(*)::text from guild_characters where character_key = $1 and user_id = $2`,
		`select count(*)::text from guild_raid_signups where character_key = $1 and signed_by = $2 and acted_by = $2`,
	} {
		if n := w.one(t, q, claim.Key, owner); n != "1" {
			t.Fatalf("%s = %s, want 1", q, n)
		}
	}
	if n := w.one(t, `select count(*)::text from guild_members where guild_id = $1 and user_id = $2`, w.guildID, picker); n != "0" {
		t.Fatal("the Discord-only account is still a member with nothing on the roster")
	}
}

// Linking the Discord account a Discord-only account belongs to merges it into the site
// account: its picks and sign-ups move across and the empty account is gone.
func TestMergeDiscordOnly(t *testing.T) {
	w := newClaimWorld(t)
	ctx := context.Background()
	picker := w.discordUser(t, "d-100")
	claim, err := w.c.ClaimForDiscord(ctx, w.guildID, picker, 101)
	if err != nil {
		t.Fatal(err)
	}
	w.signUp(t, claim.Key, picker)
	if err := w.c.MergeDiscordOnly(ctx, w.bnetOwn, picker); err == nil {
		t.Fatal("a Battle.net account merged into another")
	}
	if err := w.c.MergeDiscordOnly(ctx, picker, w.bnetOwn); err != nil {
		t.Fatal(err)
	}
	if n := w.one(t, `select count(*)::text from guild_raid_signups where signed_by = $1`, w.bnetOwn); n != "1" {
		t.Fatalf("sign-ups moved = %s", n)
	}
	if n := w.one(t, `select count(*)::text from users where id = $1`, picker); n != "0" {
		t.Fatal("the Discord-only account survived the merge")
	}
	if n := w.one(t, `select count(*)::text from discord_links where discord_user_id = 'd-100'`); n != "0" {
		t.Fatal("the old link survived the merge")
	}
}
