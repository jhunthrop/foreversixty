// api/internal/guilds/raids_discord_test.go
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// recordingRaidEvents is the guild channel's poster, recording every post in order.
type recordingRaidEvents struct {
	mu      sync.Mutex
	posted  []notify.RaidPosted
	changed []notify.RaidChanged
	reports []notify.ReportReady
	// notPosted makes ReportReady answer that the post reached no channel.
	notPosted bool
	// reportErr makes ReportReady fail outright.
	reportErr error
}

func (f *recordingRaidEvents) ReportReady(_ context.Context, _ int64, e notify.ReportReady) (bool, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.reports = append(f.reports, e)
	if f.reportErr != nil {
		return false, f.reportErr
	}
	return !f.notPosted, nil
}

func (f *recordingRaidEvents) RaidPosted(_ context.Context, _ int64, e notify.RaidPosted) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.posted = append(f.posted, e)
}

func (f *recordingRaidEvents) RaidChanged(_ context.Context, _ int64, e notify.RaidChanged) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.changed = append(f.changed, e)
}

const testSiteBase = "https://foreversixty.gg"

// postingWorld is a raid world whose guild has a connected channel and a recording poster.
func postingWorld(t *testing.T) (raidWorld, *recordingRaidEvents, *[]bool) {
	t.Helper()
	w := newRaidWorld(t, true)
	events, defaults := &recordingRaidEvents{}, &[]bool{}
	w.h.svc.Discord = fakeDiscord{enabled: true, hook: &notify.GuildWebhook{ChannelID: "1"}, defaults: defaults}
	w.h.svc.RaidEvents = events
	w.h.svc.PublicBaseURL = testSiteBase + "/"
	return w, events, defaults
}

func postingBody(startsAt string, repeat int, post string) string {
	body := raidBody(startsAt, repeat)
	if post == "" {
		return strings.Replace(body, `,"post_to_discord":false`, "", 1)
	}
	return strings.Replace(body, `"post_to_discord":false`, `"post_to_discord":`+post, 1)
}

func patchBody(title, startsAt, post string) string {
	return fmt.Sprintf(`{"zone":%q,"title":%q,"starts_at":%q,"duration_minutes":180,"notes":"",`+
		`"timezone":"America/Chicago","post_to_discord":%s}`, raidZoneOny, title, startsAt, post)
}

func TestCreateRaidPostsOneEmbedPerRaidSoonestFirst(t *testing.T) {
	w, events, defaults := postingWorld(t)
	views := w.createRaids(t, postingBody(futureStart(), 3, "true"))
	if len(events.posted) != 4 || len(views) != 4 {
		t.Fatalf("posted %d for %d raids", len(events.posted), len(views))
	}
	for i, e := range events.posted {
		if i > 0 && !e.StartsAt.After(events.posted[i-1].StartsAt) {
			t.Fatalf("post %d not after post %d", i, i-1)
		}
	}
	first := events.posted[0]
	wantURL := fmt.Sprintf("%s/guild/us/hardcore/olympus/raids/%d", testSiteBase, views[0].ID)
	if first.RaidURL != wantURL || first.GuildName != "Olympus" || first.Title != "Onyxia's Lair" ||
		first.Timezone != "America/Chicago" || first.SignedUp != 0 || !first.SizeKnown || first.Size != 40 ||
		first.CrestURL != "" {
		t.Fatalf("first post = %+v, want url %s", first, wantURL)
	}
	if len(*defaults) != 1 || !(*defaults)[0] {
		t.Fatalf("post_default writes = %v", *defaults)
	}
}

func TestCreateRaidWithoutTheBoxPostsNothing(t *testing.T) {
	w, events, defaults := postingWorld(t)
	w.createRaids(t, postingBody(futureStart(), 0, "false"))
	w.createRaids(t, postingBody(time.Now().Add(9*24*time.Hour).UTC().Truncate(time.Minute).Format(time.RFC3339), 0, ""))
	if len(events.posted) != 0 {
		t.Fatalf("posted %+v", events.posted)
	}
	// Only the save that carried the box remembers it.
	if len(*defaults) != 1 || (*defaults)[0] {
		t.Fatalf("post_default writes = %v", *defaults)
	}
}

func TestPatchRaidPostsOnlyATimeChange(t *testing.T) {
	w, events, defaults := postingWorld(t)
	raid := w.createRaid(t)
	old := raid.StartsAt.UTC()
	w.as(w.officer)
	patch := func(body string) {
		t.Helper()
		if res := w.h.do(http.MethodPatch, w.path("/raids/%d", raid.ID), body); res.StatusCode != http.StatusOK {
			t.Fatalf("patch = %d: %s", res.StatusCode, readBody(res))
		}
	}
	patch(patchBody("Ony night", old.Format(time.RFC3339), "true"))
	if len(events.changed) != 0 {
		t.Fatalf("a title edit posted %+v", events.changed)
	}
	moved := old.Add(24 * time.Hour)
	patch(patchBody("Ony night", moved.Format(time.RFC3339), "false"))
	if len(events.changed) != 0 {
		t.Fatalf("a move without the box posted %+v", events.changed)
	}
	patch(patchBody("Ony night", moved.Add(time.Hour).Format(time.RFC3339), "true"))
	if len(events.changed) != 1 {
		t.Fatalf("changed = %+v", events.changed)
	}
	got := events.changed[0]
	if got.Cancelled || got.Was == nil || !got.Was.Equal(moved) || !got.StartsAt.Equal(moved.Add(time.Hour)) ||
		got.Title != "Ony night" {
		t.Fatalf("moved post = %+v", got)
	}
	if want := []bool{false, true, false, true}; fmt.Sprint(*defaults) != fmt.Sprint(want) {
		t.Fatalf("post_default writes = %v, want %v", *defaults, want)
	}
}

func TestCancelRaidPostsTheCancellationWhenAsked(t *testing.T) {
	w, events, defaults := postingWorld(t)
	quiet, loud := w.createRaid(t), w.createRaid(t)
	writesBefore := len(*defaults)
	w.as(w.officer)
	if res := w.h.do(http.MethodPost, w.path("/raids/%d/cancel", quiet.ID), ""); res.StatusCode != http.StatusOK {
		t.Fatalf("cancel without body = %d: %s", res.StatusCode, readBody(res))
	}
	if len(events.changed) != 0 {
		t.Fatalf("a cancel without the box posted %+v", events.changed)
	}
	w.as(w.member)
	if res := w.putSignup(loud.ID, memberKeyA, `{"state":"in","role":"dps","note":""}`); res.StatusCode != http.StatusOK {
		t.Fatalf("sign-up = %d: %s", res.StatusCode, readBody(res))
	}
	w.as(w.officer)
	res := w.h.do(http.MethodPost, w.path("/raids/%d/cancel", loud.ID), `{"post_to_discord":true}`)
	if res.StatusCode != http.StatusOK {
		t.Fatalf("cancel = %d: %s", res.StatusCode, readBody(res))
	}
	if len(events.changed) != 1 || !events.changed[0].Cancelled || events.changed[0].Was != nil ||
		events.changed[0].SignedUp != 1 {
		t.Fatalf("changed = %+v", events.changed)
	}
	// The cancel box pre-ticks from the form's default and never writes it.
	if len(*defaults) != writesBefore {
		t.Fatalf("cancel wrote post_default: %v", *defaults)
	}
}
