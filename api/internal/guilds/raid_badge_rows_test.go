package guilds

import (
	"context"
	"testing"
	"time"
)

// fakeSimsRanks answers one rank for one character and records what it was asked.
type fakeSimsRanks struct {
	key      string
	pos, of  int
	askedKey string
	askedFor int64
}

func (f *fakeSimsRanks) GuildSimsRank(_ context.Context, guildID int64, key string) (int, int, bool, error) {
	f.askedKey, f.askedFor = key, guildID
	if key != f.key {
		return 0, 0, false, nil
	}
	return f.pos, f.of, true, nil
}

// The Raid tab's rows: the viewer's main on the guild's sims ladder, and how much loot the
// guild's last done raid gave out (spec 2026-11-08-app-flows, the Raid tab's guild pages).
func TestRaidBadgeCarriesSimsRankAndLastRaidLoot(t *testing.T) {
	w := newRaidWorld(t, true)
	if b := w.badgeOf(t, w.member); b.SimsRank != nil || b.LastRaidLoot != nil {
		t.Fatalf("nothing yet: %+v", b)
	}
	ranks := &fakeSimsRanks{key: memberKeyA, pos: 2, of: 7}
	w.h.svc.SimsRanks = ranks
	b := w.badgeOf(t, w.member)
	if ranks.askedKey != memberKeyA || ranks.askedFor != w.gid {
		t.Fatalf("asked %q in %d", ranks.askedKey, ranks.askedFor)
	}
	if b.SimsRank == nil || b.SimsRank.Position != 2 || b.SimsRank.Of != 7 {
		t.Fatalf("sims rank %+v", b.SimsRank)
	}

	ctx := context.Background()
	exec := func(sql string, args ...any) {
		t.Helper()
		if _, err := w.h.pool.Exec(ctx, sql, args...); err != nil {
			t.Fatal(err)
		}
	}
	doneRaid := func(daysAgo int) int64 {
		var id int64
		at := time.Now().Add(-time.Duration(daysAgo) * 24 * time.Hour)
		if err := w.h.pool.QueryRow(ctx, `insert into guild_raids (guild_id, zone, title, starts_at, duration_minutes, status, created_by)
			values ($1, 'raid:onyxias-lair', 'Onyxia', $2, 180, 'done', $3) returning id`, w.gid, at, w.officer).Scan(&id); err != nil {
			t.Fatal(err)
		}
		return id
	}
	older, last := doneRaid(14), doneRaid(7)
	exec(`insert into loot_awards (guild_id, encounter_id, item_id, character_key, raid_id) values
		($1, 1084, 16900, $2, $3), ($1, 1084, 16901, $2, $3), ($1, 1084, 16902, $2, $4)`, w.gid, memberKeyA, last, older)
	// A recorded drop nobody was awarded counts its copies; one that was awarded counts once.
	exec(`insert into raid_drops (guild_id, raid_id, encounter_id, item_id, count, corpse_key, dropped_at, reported_character)
		values ($1, $2, 1084, 18423, 2, 'c1', now(), 'kraggor'), ($1, $2, 1084, 16900, 1, 'c1', now(), 'kraggor')`, w.gid, last)
	// A cancelled raid after it is not the last done raid.
	exec(`insert into guild_raids (guild_id, zone, title, starts_at, duration_minutes, status, created_by)
		values ($1, 'raid:onyxias-lair', 'Onyxia', now() - interval '1 day', 180, 'cancelled', $2)`, w.gid, w.officer)

	b = w.badgeOf(t, w.member)
	if b.LastRaidLoot == nil || b.LastRaidLoot.RaidID != last || b.LastRaidLoot.Count != 4 {
		t.Fatalf("last raid loot %+v", b.LastRaidLoot)
	}
}

// The main is the account's marked main when it is verified in the guild, else its highest
// level verified character there.
func TestGuildMainOf(t *testing.T) {
	w := newRaidWorld(t, true)
	ctx := context.Background()
	key, ok, err := w.h.svc.Store.guildMainOf(ctx, w.gid, w.member)
	if err != nil || !ok || key == "" {
		t.Fatalf("main %q %v %v", key, ok, err)
	}
	if _, err := w.h.pool.Exec(ctx, `update users set main_character_key = $1 where id = $2`, memberKeyB, w.member); err != nil {
		t.Fatal(err)
	}
	if key, _, _ := w.h.svc.Store.guildMainOf(ctx, w.gid, w.member); key != memberKeyB {
		t.Fatalf("marked main %q", key)
	}
	stranger := seedUser(t, w.h.pool, "nomain@example.com")
	if _, ok, _ := w.h.svc.Store.guildMainOf(ctx, w.gid, stranger); ok {
		t.Fatal("a stranger has a main in the guild")
	}
}
