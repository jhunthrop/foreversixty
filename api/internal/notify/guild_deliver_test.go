package notify

import (
	"context"
	"errors"
	"strings"
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

func TestGuildStoreSetDefaults(t *testing.T) {
	ctx := context.Background()
	st, gid, _ := savedGuild(t)
	if err := st.SetDefaults(ctx, gid, true, false); err != nil {
		t.Fatal(err)
	}
	if w, _, _ := st.Get(ctx, gid); !w.PostDefault || w.RemindDefault {
		t.Fatalf("defaults not stored: post %v remind %v", w.PostDefault, w.RemindDefault)
	}
	if err := st.SetDefaults(ctx, gid+999, true, true); err != nil {
		t.Fatalf("no channel = %v", err)
	}
}

// The raid-posts job's sends return a refusal so the schedule retries it; the save-time
// paths (RaidPosted, ReportReady, Deliver) keep reporting it as not posted.
func TestJobSendsReturnARefusal(t *testing.T) {
	ctx := context.Background()
	r := newGuildRig(t)
	refused := &DiscordError{Status: 500, Body: "boom"}
	r.poster.PostErr, r.poster.ContentErr = refused, refused

	sent, err := r.d.SendRaidPost(ctx, r.gid, onyxiaThursday())
	var notSent *NotSentError
	if sent || !errors.As(err, &notSent) || !errors.Is(err, refused) {
		t.Fatalf("SendRaidPost = %v, %v", sent, err)
	}
	_, sent, err = r.d.RaidReminder(ctx, r.gid, onyxiaThursday())
	if sent || !errors.As(err, &notSent) {
		t.Fatalf("RaidReminder = %v, %v", sent, err)
	}
	if strings.Contains(err.Error(), "guild-token") {
		t.Fatal("the refusal holds the webhook token")
	}
	if w := r.hook(t); w.Failures != 2 {
		t.Fatalf("failures = %d, want 2", w.Failures)
	}
	if posted, err := r.d.ReportReady(ctx, r.gid, onyxiaReport()); posted || err != nil {
		t.Fatalf("ReportReady = %v, %v", posted, err)
	}
	if err := r.d.Deliver(ctx, r.gid, eventRaidPosted, Embed{}); err != nil {
		t.Fatalf("Deliver = %v", err)
	}

	r.poster.PostErr = nil
	if sent, err := r.d.SendRaidPost(ctx, r.gid, onyxiaThursday()); !sent || err != nil {
		t.Fatalf("landed: SendRaidPost = %v, %v", sent, err)
	}
	var none *GuildDeliverer
	if sent, err := none.SendRaidPost(ctx, r.gid, onyxiaThursday()); sent || err != nil {
		t.Fatalf("nil deliverer: %v, %v", sent, err)
	}
}
