// api/internal/guilds/raids_post_now_test.go
package guilds

import (
	"context"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// repostingEvents is recordingRaidEvents with the bot router's repost.
type repostingEvents struct {
	*recordingRaidEvents
	reposts []notify.RaidPosted
}

func (f *repostingEvents) RepostRaid(_ context.Context, _ int64, e notify.RaidPosted) (bool, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.reposts = append(f.reposts, e)
	return true, nil
}

func (w raidWorld) postNow(t *testing.T, id int64) *http.Response {
	t.Helper()
	return w.h.do(http.MethodPost, w.path("/raids/%d/discord/post", id), "")
}

func TestPostNowPostsAnUnpostedRaidOnce(t *testing.T) {
	w, events, _ := postingWorld(t)
	quiet := w.createRaids(t, scheduleBody(startIn(20*day), 0, "false", "true"))[0]
	if len(events.posted) != 0 {
		t.Fatalf("box off posted %+v", events.posted)
	}
	w.as(w.officer)
	res := w.postNow(t, quiet.ID)
	if res.StatusCode != http.StatusOK {
		t.Fatalf("post now = %d: %s", res.StatusCode, readBody(res))
	}
	var v RaidView
	w.h.data(res, &v)
	if p := v.DiscordPost; p == nil || p.PostedAt == nil {
		t.Fatalf("discord_post = %+v", p)
	}
	if sentFor(events.posted, quiet.ID) != 1 {
		t.Fatalf("posted %d", sentFor(events.posted, quiet.ID))
	}
	// The schedule run never sends it again.
	if got := w.run(t, startIn(time.Hour)); got.Posted != 0 {
		t.Fatalf("run posted again: %+v", got)
	}
}

func TestRepostAsksTheRouterForAFreshMessage(t *testing.T) {
	w, rec, _ := postingWorld(t)
	events := &repostingEvents{recordingRaidEvents: rec}
	w.h.svc.RaidEvents = events
	raid := w.createRaids(t, scheduleBody(startIn(3*day), 0, "true", "true"))[0]
	w.as(w.officer)
	if res := w.postNow(t, raid.ID); res.StatusCode != http.StatusOK {
		t.Fatalf("repost = %d: %s", res.StatusCode, readBody(res))
	}
	if len(events.reposts) != 1 || events.reposts[0].RaidID != raid.ID || sentFor(rec.posted, raid.ID) != 1 {
		t.Fatalf("reposts %+v, posted %d", events.reposts, sentFor(rec.posted, raid.ID))
	}
}

// A poster without a repost (the webhook) posts the raid again as a new message.
func TestRepostThroughTheWebhookPostsAgain(t *testing.T) {
	w, events, _ := postingWorld(t)
	raid := w.createRaids(t, scheduleBody(startIn(3*day), 0, "true", "false"))[0]
	w.as(w.officer)
	if res := w.postNow(t, raid.ID); res.StatusCode != http.StatusOK {
		t.Fatalf("repost = %d: %s", res.StatusCode, readBody(res))
	}
	if sentFor(events.posted, raid.ID) != 2 {
		t.Fatalf("posted %d, want 2", sentFor(events.posted, raid.ID))
	}
}

func TestPostNowRefusals(t *testing.T) {
	w, events, _ := postingWorld(t)
	raid := w.createRaids(t, scheduleBody(startIn(3*day), 0, "false", "false"))[0]

	w.as(w.member)
	if res := w.postNow(t, raid.ID); res.StatusCode != http.StatusForbidden {
		t.Fatalf("member = %d", res.StatusCode)
	}

	events.noChannel = true
	w.as(w.officer)
	if res := w.postNow(t, raid.ID); res.StatusCode != http.StatusBadGateway {
		t.Fatalf("refused send = %d: %s", res.StatusCode, readBody(res))
	}
	if rs, _ := w.scheduleOf(t, raid.ID); rs.PostedAt != nil {
		t.Fatalf("a refused send marked posted: %+v", rs)
	}
	events.noChannel = false

	w.statusAction(t, raid.ID, "cancel")
	if res := w.postNow(t, raid.ID); res.StatusCode != http.StatusConflict {
		t.Fatalf("cancelled = %d", res.StatusCode)
	}

	w.h.svc.Discord = fakeDiscord{enabled: true}
	open := w.createRaids(t, scheduleBody(startIn(3*day), 0, "false", "false"))[0]
	res := w.postNow(t, open.ID)
	if res.StatusCode != http.StatusConflict || !strings.Contains(readBody(res), "no_channel") {
		t.Fatalf("no channel = %d", res.StatusCode)
	}
}
