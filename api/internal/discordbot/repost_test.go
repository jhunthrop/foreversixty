package discordbot

import (
	"context"
	"testing"
)

// A repost turns the sent post into a fresh one in the install's channel and deletes the old
// message and its reminder reply.
func TestRepostQueuesAFreshPostAndDeletesTheOld(t *testing.T) {
	ctx := context.Background()
	w := newPostWorld(t)
	w.router.RaidPosted(ctx, w.gid, posted(w, w.raidID))
	if _, err := w.st.Pool.Exec(ctx, `update guild_raid_discord_posts set message_id = 'm1',
		reminder_message_id = 'r1', dirty_at = null, created_at = now() - interval '1 day'
		where raid_id = $1`, w.raidID); err != nil {
		t.Fatal(err)
	}
	var channel string
	if err := w.st.Pool.QueryRow(ctx, `select channel_id from guild_raid_discord_posts where raid_id = $1`,
		w.raidID).Scan(&channel); err != nil {
		t.Fatal(err)
	}
	sent, err := w.router.RepostRaid(ctx, w.gid, posted(w, w.raidID))
	if !sent || err != nil {
		t.Fatalf("repost: sent %v, err %v", sent, err)
	}
	r, found := w.row(t, w.raidID)
	if !found || r.messageID != nil || !r.dirty || r.gone {
		t.Fatalf("row after repost = %+v, found %v", r, found)
	}
	deleted := map[string]bool{}
	for _, c := range w.rest.Calls() {
		if c.Method == "DeleteMessage" && c.ChannelID == channel {
			deleted[c.MessageID] = true
		}
	}
	if !deleted["m1"] || !deleted["r1"] || len(w.poster.Posts()) != 0 {
		t.Fatalf("deleted %v, webhook posts %d", deleted, len(w.poster.Posts()))
	}
}

// A raid with no bot row yet is queued; without the bot the webhook posts it again.
func TestRepostQueuesOrFallsBackToTheWebhook(t *testing.T) {
	ctx := context.Background()
	w := newPostWorld(t)
	if sent, err := w.router.RepostRaid(ctx, w.gid, posted(w, w.raidID)); !sent || err != nil {
		t.Fatalf("no row: sent %v, err %v", sent, err)
	}
	if r, found := w.row(t, w.raidID); !found || !r.dirty {
		t.Fatalf("no row: queued %+v, found %v", r, found)
	}

	w = newPostWorld(t)
	if err := w.st.RemoveBot(ctx, w.gid); err != nil {
		t.Fatal(err)
	}
	if sent, err := w.router.RepostRaid(ctx, w.gid, posted(w, w.raidID)); !sent || err != nil {
		t.Fatalf("webhook: sent %v, err %v", sent, err)
	}
	if len(w.poster.Posts()) != 1 {
		t.Fatalf("webhook posts %d", len(w.poster.Posts()))
	}
}
