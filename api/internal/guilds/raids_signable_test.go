package guilds

import (
	"fmt"
	"testing"
	"time"
)

// TestNextRaidAndSignableAgree pins Raid.Signable (the bot's /signup list) to NextRaid's SQL:
// for one raid per guild, NextRaid finds it exactly when Signable says so.
func TestNextRaidAndSignableAgree(t *testing.T) {
	pool := testPool(t)
	store := &Store{Pool: pool}
	uid := seedUser(t, pool, "signable@example.com")
	now := time.Now().Truncate(time.Second)
	cases := []struct {
		name   string
		status string
		starts time.Time
	}{
		{"planned ahead", RaidPlanned, now.Add(24 * time.Hour)},
		{"planned, running", RaidPlanned, now.Add(-time.Hour)},
		{"locked ahead", RaidLocked, now.Add(24 * time.Hour)},
		{"planned, ended", RaidPlanned, now.Add(-4 * time.Hour)},
		{"locked, ended", RaidLocked, now.Add(-4 * time.Hour)},
		{"planned, ends exactly now", RaidPlanned, now.Add(-3 * time.Hour)},
		{"cancelled ahead", RaidCancelled, now.Add(24 * time.Hour)},
		{"done", RaidDone, now.Add(-4 * time.Hour)},
	}
	for i, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gid := seedGuild(t, pool, fmt.Sprintf("Signable %d", i))
			var raid Raid
			err := pool.QueryRow(t.Context(), `insert into guild_raids
				(guild_id, zone, title, starts_at, duration_minutes, status, created_by)
				values ($1, 'raid:onyxias-lair', 'Onyxia''s Lair', $2, 180, $3, $4) returning id`,
				gid, c.starts, c.status, uid).Scan(&raid.ID)
			if err != nil {
				t.Fatal(err)
			}
			raid.Status, raid.StartsAt, raid.DurationMinutes = c.status, c.starts, 180
			_, found, err := store.NextRaid(t.Context(), gid, now)
			if err != nil {
				t.Fatal(err)
			}
			if found != raid.Signable(now) {
				t.Fatalf("NextRaid found=%v, Signable=%v", found, raid.Signable(now))
			}
		})
	}
}
