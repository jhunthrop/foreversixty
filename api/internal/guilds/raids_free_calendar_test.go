// api/internal/guilds/raids_free_calendar_test.go
//
// Owner ruling 2026-11-09: the Discord bot is free for any claimed guild's officers, and its
// raid posts need raids, so the raid calendar (create, edit, lock, unlock, cancel, Post now,
// answering for a member) runs without a plan. The raid tools beyond it keep the guild plan.
package guilds

import (
	"net/http"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

func TestTheRaidCalendarIsFreeForAClaimedGuildsOfficers(t *testing.T) {
	w := newRaidWorld(t, false)
	events := &recordingRaidEvents{}
	w.h.svc.Discord = fakeDiscord{enabled: true, hook: &notify.GuildWebhook{ChannelID: "1"}}
	w.h.svc.RaidEvents = events
	w.h.svc.PublicBaseURL = testSiteBase + "/"

	raid := w.createRaids(t, scheduleBody(startIn(20*day), 0, "true", "true"))[0]
	if sentFor(events.posted, raid.ID) != 1 {
		t.Fatalf("a free guild's raid did not post: %+v", events.posted)
	}
	w.patchTo(t, raid.ID, startIn(21*day), "true", "true")
	for _, action := range []string{"lock", "unlock"} {
		w.statusAction(t, raid.ID, action)
	}
	w.as(w.officer)
	if res := w.postNow(t, raid.ID); res.StatusCode != http.StatusOK {
		t.Fatalf("post now = %d: %s", res.StatusCode, readBody(res))
	}
	w.statusAction(t, raid.ID, "cancel")
}

// The raid tools beyond the calendar still answer guild_plan_required on a free guild.
func TestTheRaidToolsKeepTheGuildPlan(t *testing.T) {
	w := newRaidWorld(t, false)
	raid := w.createRaids(t, raidBody(futureStart(), 0))[0]
	w.as(w.officer)
	for _, c := range []struct{ method, path, body string }{
		{http.MethodPut, "/raids/%d/comp-target", `{"tanks":2,"healers":8,"melee":10,"ranged":20}`},
		{http.MethodPost, "/raids/%d/ping", ""},
	} {
		res := w.h.do(c.method, w.path(c.path, raid.ID), c.body)
		if body := readBody(res); res.StatusCode != http.StatusForbidden || !strings.Contains(body, "guild_plan_required") {
			t.Errorf("%s %s = %d: %s", c.method, c.path, res.StatusCode, body)
		}
	}
	// A member still cannot write the calendar.
	w.as(w.member)
	if res := w.h.do(http.MethodPost, w.path("/raids/%d/lock", raid.ID), ""); res.StatusCode != http.StatusForbidden {
		t.Fatalf("member lock = %d", res.StatusCode)
	}
}
