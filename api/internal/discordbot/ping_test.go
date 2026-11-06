package discordbot

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

func TestPingRaidPostsInTheBotChannelPingingOnlyTheNamed(t *testing.T) {
	w := newPostWorld(t)
	ping := guilds.RaidPing{RaidID: w.raidID, Content: "Still need your answer: <@111111111111111111>", MentionIDs: []string{"111111111111111111"}}
	posted, err := w.router.PingRaid(context.Background(), w.gid, ping)
	if err != nil || !posted {
		t.Fatalf("posted %v, err %v", posted, err)
	}
	calls := w.callsOf("CreateMessage")
	if len(calls) != 1 || calls[0].ChannelID != postChannel || calls[0].Message.ReplyTo != "" {
		t.Fatalf("calls = %+v", calls)
	}
	raw, err := json.Marshal(calls[0].Message)
	if err != nil {
		t.Fatal(err)
	}
	var body struct {
		AllowedMentions struct {
			Parse []string `json:"parse"`
			Users []string `json:"users"`
		} `json:"allowed_mentions"`
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		t.Fatal(err)
	}
	if len(body.AllowedMentions.Parse) != 0 || len(body.AllowedMentions.Users) != 1 || body.AllowedMentions.Users[0] != "111111111111111111" {
		t.Fatalf("allowed_mentions = %+v (%s)", body.AllowedMentions, raw)
	}
}

func TestPingRaidRepliesUnderTheLivePost(t *testing.T) {
	w := newPostWorld(t)
	w.enqueueIn(t, w.raidID, "chan2")
	if _, err := w.st.Pool.Exec(context.Background(),
		`update guild_raid_discord_posts set message_id = 'post1' where raid_id = $1`, w.raidID); err != nil {
		t.Fatal(err)
	}
	if posted, err := w.router.PingRaid(context.Background(), w.gid, guilds.RaidPing{RaidID: w.raidID, Content: "x"}); err != nil || !posted {
		t.Fatalf("posted %v, err %v", posted, err)
	}
	calls := w.callsOf("CreateMessage")
	if len(calls) != 1 || calls[0].ChannelID != "chan2" || calls[0].Message.ReplyTo != "post1" {
		t.Fatalf("calls = %+v", calls)
	}
}

func TestPingRaidWithoutAnActiveBotPostsNothing(t *testing.T) {
	w := newPostWorld(t)
	if err := w.st.RemoveBot(context.Background(), w.gid); err != nil {
		t.Fatal(err)
	}
	posted, err := w.router.PingRaid(context.Background(), w.gid, guilds.RaidPing{RaidID: w.raidID, Content: "x"})
	if err != nil || posted || len(w.rest.Calls()) != 0 {
		t.Fatalf("posted %v, err %v, calls %+v", posted, err, w.rest.Calls())
	}
}

func TestMessagesWithoutMentionUsersPingNoOne(t *testing.T) {
	raw, err := json.Marshal(Message{Content: "hi <@1>"})
	if err != nil {
		t.Fatal(err)
	}
	if want := `"allowed_mentions":{"parse":[]}`; !strings.Contains(string(raw), want) {
		t.Fatalf("body = %s", raw)
	}
}
