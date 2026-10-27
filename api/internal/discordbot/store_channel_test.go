package discordbot

import (
	"context"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// TestRepointFollowsRaidSignable pins repointSQL to guilds.Raid.Signable: a channel save
// moves exactly the posts of the raids Signable holds open, and drops the rest.
func TestRepointFollowsRaidSignable(t *testing.T) {
	now := time.Now().Truncate(time.Minute)
	cases := []struct {
		name     string
		status   string
		startsIn time.Duration
		minutes  int
	}{
		{"planned, ahead", guilds.RaidPlanned, 48 * time.Hour, 180},
		{"locked, ahead", guilds.RaidLocked, 2 * time.Hour, 180},
		{"planned, under way", guilds.RaidPlanned, -time.Hour, 180},
		{"locked, ended", guilds.RaidLocked, -5 * time.Hour, 180},
		{"planned, ended", guilds.RaidPlanned, -4 * time.Hour, 180},
		{"cancelled, ahead", guilds.RaidCancelled, 48 * time.Hour, 180},
		{"done, ended", guilds.RaidDone, -5 * time.Hour, 180},
	}
	w := newPostWorld(t)
	ctx := context.Background()
	ids := make([]int64, len(cases))
	for i, c := range cases {
		if err := w.st.Pool.QueryRow(ctx, `insert into guild_raids
			(guild_id, zone, title, starts_at, duration_minutes, status, created_by)
			values ($1, 'raid:onyxias-lair', $2, $3, $4, $5, $6) returning id`,
			w.gid, c.name, now.Add(c.startsIn), c.minutes, c.status, w.uid).Scan(&ids[i]); err != nil {
			t.Fatal(err)
		}
		w.enqueueIn(t, ids[i], postChannel)
	}
	if err := w.st.SetChannel(ctx, w.gid, "chan2", "general"); err != nil {
		t.Fatal(err)
	}
	for i, c := range cases {
		raid := guilds.Raid{Status: c.status, StartsAt: now.Add(c.startsIn), DurationMinutes: c.minutes}
		var channel string
		err := w.st.Pool.QueryRow(ctx, `select channel_id from guild_raid_discord_posts where raid_id = $1`, ids[i]).Scan(&channel)
		repointed := err == nil && channel == "chan2"
		if want := raid.Signable(time.Now()); repointed != want {
			t.Errorf("%s: re-pointed %v (channel %q, err %v), Signable %v", c.name, repointed, channel, err, want)
		}
	}
}
