package notify

import (
	"context"
	"testing"
	"time"
)

const guildHookURL = "https://discord.com/api/webhooks/1/guild-token"

// guildRig is a guild with a saved, sealed channel and a deliverer over a FakePoster.
type guildRig struct {
	d      *GuildDeliverer
	poster *FakePoster
	store  *GuildStore
	gid    int64
}

func newGuildRig(t *testing.T) guildRig {
	t.Helper()
	st, gid, uid := savedGuild(t)
	sealer, err := NewSealer(make([]byte, 32))
	if err != nil {
		t.Fatal(err)
	}
	box, err := sealer.Seal(guildHookURL)
	if err != nil {
		t.Fatal(err)
	}
	if err := st.Save(context.Background(), GuildSaveInput{GuildID: gid, ChannelID: "555", TokenBox: box, ConnectedBy: uid}); err != nil {
		t.Fatal(err)
	}
	poster := &FakePoster{}
	d := &GuildDeliverer{Store: st, Sealer: sealer, Poster: poster, now: func() time.Time { return fixedFinish }}
	return guildRig{d: d, poster: poster, store: st, gid: gid}
}

func (r guildRig) hook(t *testing.T) GuildWebhook {
	t.Helper()
	w, ok, err := r.store.Get(context.Background(), r.gid)
	if err != nil || !ok {
		t.Fatalf("Get = %v, %v", ok, err)
	}
	return w
}

func TestGuildDelivererPostsTheRaidEmbeds(t *testing.T) {
	r := newGuildRig(t)
	ctx := context.Background()
	r.d.RaidPosted(ctx, r.gid, onyxiaThursday())
	r.d.RaidChanged(ctx, r.gid, RaidChanged{RaidPosted: onyxiaThursday(), Cancelled: true})
	posts := r.poster.Posts()
	if len(posts) != 2 || posts[0].URL != guildHookURL {
		t.Fatalf("posts = %+v", posts)
	}
	if posts[0].Embed.Title != "Onyxia's Lair · Thursday 10 December" || posts[1].Embed.Color != cancelColor ||
		!posts[0].Embed.Timestamp.Equal(fixedFinish) {
		t.Fatalf("embeds = %+v", posts)
	}
}

func TestGuildDelivererFiveFailuresSwitchTheChannelOff(t *testing.T) {
	r := newGuildRig(t)
	ctx := context.Background()
	r.poster.PostErr = &DiscordError{Status: 429, Body: "slow down"}
	r.d.RaidPosted(ctx, r.gid, onyxiaThursday())
	if w := r.hook(t); w.Failures != 0 || w.LastError == nil || *w.LastError != "slow down" {
		t.Fatalf("a 429 counted: %+v", w)
	}
	r.poster.PostErr = &DiscordError{Status: 404, Body: "Unknown Webhook"}
	for i := 0; i < MaxFailures; i++ {
		r.d.RaidPosted(ctx, r.gid, onyxiaThursday())
	}
	w := r.hook(t)
	if w.Failures != MaxFailures || w.DisabledReason == nil || *w.DisabledReason != reasonFailedFiveTimes {
		t.Fatalf("after five = %+v", w)
	}
	before := len(r.poster.Posts())
	r.d.RaidPosted(ctx, r.gid, onyxiaThursday())
	if len(r.poster.Posts()) != before {
		t.Fatal("a switched-off channel was posted to")
	}
}

func TestGuildDelivererSuccessResetsTheCounter(t *testing.T) {
	r := newGuildRig(t)
	ctx := context.Background()
	r.poster.PostErr = &DiscordError{Status: 500, Body: "boom"}
	r.d.RaidPosted(ctx, r.gid, onyxiaThursday())
	r.d.RaidPosted(ctx, r.gid, onyxiaThursday())
	if w := r.hook(t); w.Failures != 2 {
		t.Fatalf("failures = %d", w.Failures)
	}
	r.poster.PostErr = nil
	r.d.RaidPosted(ctx, r.gid, onyxiaThursday())
	if w := r.hook(t); w.Failures != 0 || w.LastError != nil {
		t.Fatalf("after success = %+v", w)
	}
}

func TestGuildDelivererSkipsAGuildWithNoChannelAndANilDeliverer(t *testing.T) {
	r := newGuildRig(t)
	ctx := context.Background()
	if err := r.store.Delete(ctx, r.gid); err != nil {
		t.Fatal(err)
	}
	r.d.RaidPosted(ctx, r.gid, onyxiaThursday())
	if len(r.poster.Posts()) != 0 {
		t.Fatal("posted with no channel")
	}
	var none *GuildDeliverer
	none.RaidPosted(ctx, r.gid, onyxiaThursday())
	none.RaidChanged(ctx, r.gid, RaidChanged{})
	if (&Service{}).GuildDeliverer() != nil {
		t.Fatal("a service with no guild store built a deliverer")
	}
}

func TestGuildStoreSetPostDefault(t *testing.T) {
	ctx := context.Background()
	st, gid, _ := savedGuild(t)
	if err := st.SetPostDefault(ctx, gid, true); err != nil {
		t.Fatal(err)
	}
	if w, _, _ := st.Get(ctx, gid); !w.PostDefault {
		t.Fatal("post_default not stored")
	}
	if err := st.SetPostDefault(ctx, gid+999, true); err != nil {
		t.Fatalf("no channel = %v", err)
	}
}
