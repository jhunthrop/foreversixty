package discordbot

import (
	"context"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

func posted(w *postWorld, raidID int64) notify.RaidPosted {
	return notify.RaidPosted{RaidID: raidID, GuildName: "Olympus XXVII", Title: "Onyxia's Lair", StartsAt: w.clock.now()}
}

func TestRouterPostsThroughTheBotWhenActive(t *testing.T) {
	w := newPostWorld(t)
	w.router.RaidPosted(context.Background(), w.gid, posted(w, w.raidID))
	r, found := w.row(t, w.raidID)
	if !found || r.messageID != nil || !r.dirty {
		t.Fatalf("row = %+v, found %v", r, found)
	}
	if len(w.poster.Posts()) != 0 || len(w.rest.Calls()) != 0 {
		t.Fatalf("posted at once: webhook %d, rest %+v", len(w.poster.Posts()), w.rest.Calls())
	}
}

func TestRouterFallsBackToTheWebhook(t *testing.T) {
	cases := map[string]func(t *testing.T, w *postWorld){
		"no install": func(t *testing.T, w *postWorld) {
			if err := w.st.DeleteInstall(context.Background(), w.gid); err != nil {
				t.Fatal(err)
			}
		},
		"no channel": func(t *testing.T, w *postWorld) {
			ctx := context.Background()
			if err := w.st.DeleteInstall(ctx, w.gid); err != nil {
				t.Fatal(err)
			}
			if _, err := w.st.BindServer(ctx, w.gid, "srv1", "Olympus", w.uid); err != nil {
				t.Fatal(err)
			}
		},
		"removed in Discord": func(t *testing.T, w *postWorld) {
			if err := w.st.MarkRemovedInDiscord(context.Background(), w.gid, "gone"); err != nil {
				t.Fatal(err)
			}
		},
	}
	for name, setup := range cases {
		t.Run(name, func(t *testing.T) {
			w := newPostWorld(t)
			setup(t, w)
			w.router.RaidPosted(context.Background(), w.gid, posted(w, w.raidID))
			if _, found := w.row(t, w.raidID); found || len(w.poster.Posts()) != 1 || w.poster.Posts()[0].URL != webhookURL {
				t.Fatalf("row %v, webhook posts %+v", found, w.poster.Posts())
			}
		})
	}
}

func TestRaidChangedEditsTheOnePost(t *testing.T) {
	w := newPostWorld(t)
	ctx := context.Background()
	w.router.RaidPosted(ctx, w.gid, posted(w, w.raidID))
	before, _ := w.row(t, w.raidID)
	was := w.clock.now().Add(24 * time.Hour).Truncate(time.Second)
	w.router.RaidChanged(ctx, w.gid, notify.RaidChanged{RaidPosted: posted(w, w.raidID), Was: &was})
	after, _ := w.row(t, w.raidID)
	if after.dirtySeq != before.dirtySeq+1 || after.movedFrom == nil || !after.movedFrom.Equal(was) || after.messageID != nil {
		t.Fatalf("row before %+v, after %+v", before, after)
	}
	var rows int
	if err := w.st.Pool.QueryRow(ctx, `select count(*) from guild_raid_discord_posts`).Scan(&rows); err != nil || rows != 1 {
		t.Fatalf("rows = %d, %v", rows, err)
	}

	webhookEra := w.addRaid(t)
	w.router.RaidChanged(ctx, w.gid, notify.RaidChanged{RaidPosted: posted(w, webhookEra), Was: &was})
	fresh, found := w.row(t, webhookEra)
	if !found || fresh.movedFrom == nil || !fresh.movedFrom.Equal(was) || !fresh.dirty {
		t.Fatalf("webhook-era raid row = %+v, found %v", fresh, found)
	}
	cancelled := w.addRaid(t)
	w.router.RaidChanged(ctx, w.gid, notify.RaidChanged{RaidPosted: posted(w, cancelled), Cancelled: true})
	if r, found := w.row(t, cancelled); !found || r.movedFrom != nil {
		t.Fatalf("cancelled raid row = %+v, found %v", r, found)
	}
	if len(w.poster.Posts()) != 0 {
		t.Fatalf("the webhook posted %d", len(w.poster.Posts()))
	}
}

func TestRaidChangedWithoutTheBotKeepsTheRowAndPostsTheWebhook(t *testing.T) {
	w := newPostWorld(t)
	ctx := context.Background()
	w.router.RaidPosted(ctx, w.gid, posted(w, w.raidID))
	before, _ := w.row(t, w.raidID)
	if err := w.st.MarkRemovedInDiscord(ctx, w.gid, "gone"); err != nil {
		t.Fatal(err)
	}
	was := w.clock.now()
	w.router.RaidChanged(ctx, w.gid, notify.RaidChanged{RaidPosted: posted(w, w.raidID), Was: &was})
	after, _ := w.row(t, w.raidID)
	if after.dirtySeq != before.dirtySeq+1 || after.movedFrom == nil || len(w.poster.Posts()) != 1 {
		t.Fatalf("row %+v, webhook posts %d", after, len(w.poster.Posts()))
	}
}

func TestReportReadyGoesThroughTheBot(t *testing.T) {
	w := newPostWorld(t)
	ctx := context.Background()
	report := notify.ReportReady{RaidPosted: posted(w, w.raidID), Came: 22, Signed: 23}
	ok, err := w.router.ReportReady(ctx, w.gid, report)
	creates := w.callsOf("CreateMessage")
	if err != nil || !ok || len(creates) != 1 || creates[0].ChannelID != postChannel || len(w.poster.Posts()) != 0 {
		t.Fatalf("ok %v err %v creates %+v webhook %d", ok, err, creates, len(w.poster.Posts()))
	}
	if e := creates[0].Message.Embeds; len(e) != 1 || !strings.Contains(e[0].Title, "Onyxia's Lair") || len(creates[0].Message.Components) != 0 {
		t.Fatalf("report message = %+v", creates[0].Message)
	}
	if _, found := w.row(t, w.raidID); found {
		t.Fatal("the report made a post row")
	}

	gone := &APIError{Status: http.StatusForbidden, Code: codeMissingAccess}
	w.rest.ErrFor = map[string]error{"CreateMessage": gone, "ServerLayout": gone}
	if ok, err := w.router.ReportReady(ctx, w.gid, report); ok || err != nil {
		t.Fatalf("refused report = %v, %v", ok, err)
	}
	if in, _, _ := w.st.InstallFor(ctx, w.gid); !in.RemovedInDiscord {
		t.Fatalf("install = %+v, want removed", in)
	}
	if ok, err := w.router.ReportReady(ctx, w.gid, report); !ok || err != nil || len(w.poster.Posts()) != 1 {
		t.Fatalf("webhook report = %v, %v, posts %d", ok, err, len(w.poster.Posts()))
	}
}

func TestRouterWithoutABotUsesTheWebhook(t *testing.T) {
	w := newPostWorld(t)
	router := &RaidRouter{Webhook: w.router.Webhook}
	ctx := context.Background()
	router.RaidPosted(ctx, w.gid, posted(w, w.raidID))
	router.RaidChanged(ctx, w.gid, notify.RaidChanged{RaidPosted: posted(w, w.raidID), Cancelled: true})
	if _, found := w.row(t, w.raidID); found || len(w.poster.Posts()) != 2 {
		t.Fatalf("row %v, webhook posts %d", found, len(w.poster.Posts()))
	}
}

func TestSavingAnotherChannelRepointsThePosts(t *testing.T) {
	h := newInstallHarness(t)
	h.callback(h.gid, "code=c", h.officerActor())
	ctx := context.Background()
	if rec := h.do(http.MethodPut, h.path(h.gid, "/channel"), `{"channel_id":"c1"}`); rec.Code != http.StatusOK {
		t.Fatalf("first save = %d: %s", rec.Code, rec.Body.String())
	}
	var raidID int64
	if err := h.svc.d.Store.Pool.QueryRow(ctx, `insert into guild_raids (guild_id, zone, title, starts_at, duration_minutes, created_by)
		values ($1, 'raid:onyxias-lair', 'Ony', now() + interval '1 day', 180, $2) returning id`, h.gid, h.uid).Scan(&raidID); err != nil {
		t.Fatal(err)
	}
	if err := h.svc.d.Store.EnqueuePost(ctx, raidID, h.gid, "c1"); err != nil {
		t.Fatal(err)
	}
	channelOf := func() string {
		var channel string
		if err := h.svc.d.Store.Pool.QueryRow(ctx, `select channel_id from guild_raid_discord_posts where raid_id = $1`, raidID).Scan(&channel); err != nil {
			t.Fatal(err)
		}
		return channel
	}
	if rec := h.do(http.MethodPut, h.path(h.gid, "/channel"), `{"channel_id":"c1"}`); rec.Code != http.StatusOK || channelOf() != "c1" {
		t.Fatalf("same save = %d", rec.Code)
	}
	if rec := h.do(http.MethodPut, h.path(h.gid, "/channel"), `{"channel_id":"c2"}`); rec.Code != http.StatusOK || channelOf() != "c2" {
		t.Fatalf("move = %d, channel %s", rec.Code, channelOf())
	}
}
