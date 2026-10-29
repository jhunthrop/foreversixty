// api/internal/guilds/raids_schedule_test.go
package guilds

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// startIn is a start d from now, on the minute.
func startIn(d time.Duration) time.Time { return time.Now().Add(d).UTC().Truncate(time.Minute) }

// subMicro is the sub-microsecond part nanoClock adds.
const subMicro = 789 * time.Nanosecond

// nanoClock is the wall clock with a sub-microsecond part on every platform. Linux's clock
// carries nanoseconds and darwin's stops at microseconds, so a test that reads back a time
// it wrote fails in CI only, unless its clock carries nanoseconds on darwin too.
func nanoClock() time.Time { return time.Now().UTC().Truncate(time.Microsecond).Add(subMicro) }

// stored is t as a timestamptz column reads it back: microseconds, the rest truncated.
func stored(t time.Time) time.Time { return t.Truncate(time.Microsecond) }

const day = 24 * time.Hour

// scheduleBody is a create body at start with both Discord boxes; an empty value leaves the
// field out.
func scheduleBody(start time.Time, repeat int, post, remind string) string {
	body := fmt.Sprintf(`{"zone":%q,"title":"","starts_at":%q,"duration_minutes":180,"notes":"",`+
		`"timezone":"America/Chicago","repeat_weeks":%d`, raidZoneOny, start.Format(time.RFC3339), repeat)
	return body + optionalBoxes(post, remind) + "}"
}

func optionalBoxes(post, remind string) string {
	out := ""
	if post != "" {
		out += `,"post_to_discord":` + post
	}
	if remind != "" {
		out += `,"remind_day_before":` + remind
	}
	return out
}

// patchTo moves raid id to start as the officer, carrying the boxes given.
func (w raidWorld) patchTo(t *testing.T, id int64, start time.Time, post, remind string) RaidView {
	t.Helper()
	w.as(w.officer)
	body := fmt.Sprintf(`{"zone":%q,"title":"Ony","starts_at":%q,"duration_minutes":180,"notes":"",`+
		`"timezone":"America/Chicago"`, raidZoneOny, start.Format(time.RFC3339)) + optionalBoxes(post, remind) + "}"
	res := w.h.do(http.MethodPatch, w.path("/raids/%d", id), body)
	if res.StatusCode != http.StatusOK {
		t.Fatalf("patch = %d: %s", res.StatusCode, readBody(res))
	}
	var v RaidView
	w.h.data(res, &v)
	return v
}

func (w raidWorld) statusAction(t *testing.T, id int64, action string) {
	t.Helper()
	w.as(w.officer)
	if res := w.h.do(http.MethodPost, w.path("/raids/%d/%s", id, action), ""); res.StatusCode != http.StatusOK {
		t.Fatalf("%s = %d: %s", action, res.StatusCode, readBody(res))
	}
}

// scheduleOf reads raid id's schedule row; ok is false without one.
func (w raidWorld) scheduleOf(t *testing.T, id int64) (RaidSchedule, bool) {
	t.Helper()
	rows, err := w.h.svc.Store.SchedulesOf(context.Background(), []int64{id})
	if err != nil {
		t.Fatal(err)
	}
	rs, ok := rows[id]
	return rs, ok
}

func (w raidWorld) run(t *testing.T, now time.Time) ScheduleRun {
	t.Helper()
	got, err := w.h.svc.RunRaidSchedule(context.Background(), now)
	if err != nil {
		t.Fatal(err)
	}
	return got
}

// sentFor counts the events about raid id.
func sentFor(events []notify.RaidPosted, id int64) int {
	n := 0
	for _, e := range events {
		if e.RaidID == id {
			n++
		}
	}
	return n
}

func (f *recordingRaidEvents) snapshot() (posted, reminders []notify.RaidPosted, changed []notify.RaidChanged, withdrawn []notify.ReminderRef) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]notify.RaidPosted(nil), f.posted...), append([]notify.RaidPosted(nil), f.reminders...),
		append([]notify.RaidChanged(nil), f.changed...), append([]notify.ReminderRef(nil), f.withdrawn...)
}

func TestSaveWithinAWeekPostsAtOnce(t *testing.T) {
	w, events, _ := postingWorld(t)
	created := w.createRaids(t, scheduleBody(startIn(3*day), 0, "true", "true"))
	if len(events.posted) != 1 || events.posted[0].RaidID != created[0].ID {
		t.Fatalf("posted = %+v", events.posted)
	}
	for _, v := range []RaidView{created[0], w.getRaid(t, created[0].ID)} {
		p := v.DiscordPost
		if p == nil || p.PostedAt == nil || p.PostAt == nil || !p.PostAt.Equal(*p.PostedAt) || !p.Remind || p.RemindedAt != nil {
			t.Fatalf("discord_post = %+v", p)
		}
	}
}

func TestSaveFurtherOutPostsAWeekBefore(t *testing.T) {
	w, events, _ := postingWorld(t)
	start := startIn(10 * day)
	created := w.createRaids(t, scheduleBody(start, 0, "true", "false"))
	if len(events.posted) != 0 {
		t.Fatalf("posted = %+v", events.posted)
	}
	p := w.getRaid(t, created[0].ID).DiscordPost
	if p == nil || p.PostAt == nil || !p.PostAt.Equal(start.Add(-postHorizon)) || p.PostedAt != nil || p.Remind {
		t.Fatalf("discord_post = %+v", p)
	}
	// Without the box there is no schedule at all.
	quiet := w.createRaids(t, scheduleBody(startIn(2*day), 0, "false", "true"))
	if got := w.getRaid(t, quiet[0].ID).DiscordPost; got != nil || len(events.posted) != 0 {
		t.Fatalf("unticked: discord_post = %+v, posted %d", got, len(events.posted))
	}
}

func TestRepeatPostsOneWeekAtATime(t *testing.T) {
	w, events, _ := postingWorld(t)
	now := time.Now()
	views := w.createRaids(t, scheduleBody(startIn(2*day), raidRepeatWeeks, "true", "true"))
	if len(views) != 4 || len(events.posted) != 1 || events.posted[0].RaidID != views[0].ID {
		t.Fatalf("posted %d for %d raids", len(events.posted), len(views))
	}
	if got := w.run(t, now.Add(6*day)); got.Posted != 1 || len(events.posted) != 2 || events.posted[1].RaidID != views[1].ID {
		t.Fatalf("day 6: run %+v, posted %+v", got, events.posted)
	}
	if got := w.run(t, now.Add(13*day)); got.Posted != 1 || len(events.posted) != 3 || events.posted[2].RaidID != views[2].ID {
		t.Fatalf("day 13: run %+v, posted %+v", got, events.posted)
	}
}

func TestReminderRunsTheDayBefore(t *testing.T) {
	w, events, _ := postingWorld(t)
	on := w.createRaids(t, scheduleBody(startIn(5*day), 0, "true", "true"))[0]
	off := w.createRaids(t, scheduleBody(startIn(5*day+time.Hour), 0, "true", "false"))[0]
	late := w.createRaids(t, scheduleBody(startIn(30*time.Hour), 0, "true", "true"))[0]

	at := on.StartsAt.Add(-23 * time.Hour)
	if got := w.run(t, at); got.Reminded != 1 || sentFor(events.reminders, on.ID) != 1 {
		t.Fatalf("first run %+v, reminders %+v", got, events.reminders)
	}
	if rs, _ := w.scheduleOf(t, on.ID); rs.RemindedAt == nil || rs.Reminder == nil || rs.Reminder.MessageID != "1" {
		t.Fatalf("reminded row = %+v", rs)
	}
	if got := w.run(t, at.Add(time.Minute)); got.Reminded != 0 || sentFor(events.reminders, on.ID) != 1 {
		t.Fatalf("second run %+v", got)
	}
	w.run(t, off.StartsAt.Add(-time.Hour))
	if sentFor(events.reminders, off.ID) != 0 {
		t.Fatal("remind off still reminded")
	}
	w.run(t, late.StartsAt.Add(-23*time.Hour))
	if sentFor(events.reminders, late.ID) != 0 {
		t.Fatal("a post 30 h ahead still reminded (12 h rule)")
	}
}

func TestScheduleFollowsMovesCancelsAndLocks(t *testing.T) {
	w, events, _ := postingWorld(t)
	ctx := context.Background()
	now := time.Now()

	// Unposted: a move follows post_at; a move inside the week posts at once, no RaidChanged.
	far := w.createRaids(t, scheduleBody(startIn(30*day), 0, "true", "true"))[0]
	moved := w.patchTo(t, far.ID, startIn(20*day), "true", "true")
	if p := moved.DiscordPost; p == nil || p.PostedAt != nil || !p.PostAt.Equal(moved.StartsAt.Add(-postHorizon)) {
		t.Fatalf("moved unposted: discord_post = %+v", p)
	}
	w.patchTo(t, far.ID, startIn(3*day), "true", "true")
	posted, _, changed, _ := events.snapshot()
	if sentFor(posted, far.ID) != 1 || len(changed) != 0 {
		t.Fatalf("posted %+v, changed %+v", posted, changed)
	}

	// Reminded through the webhook, then moved: withdrawn, re-armed, re-sent at the new time.
	soon := w.createRaids(t, scheduleBody(startIn(5*day), 0, "true", "true"))[0]
	m1 := notify.ReminderRef{Via: notify.ReminderViaWebhook, MessageID: "m1"}
	if err := w.h.svc.Store.MarkScheduleReminded(ctx, soon.ID, m1, now); err != nil {
		t.Fatal(err)
	}
	newStart := soon.StartsAt.Add(time.Hour)
	w.patchTo(t, soon.ID, newStart, "true", "true")
	if _, _, _, withdrawn := events.snapshot(); len(withdrawn) != 1 || withdrawn[0] != m1 {
		t.Fatalf("withdrawn = %+v", withdrawn)
	}
	if rs, _ := w.scheduleOf(t, soon.ID); rs.RemindedAt != nil || rs.Reminder != nil || !rs.Remind {
		t.Fatalf("moved row = %+v", rs)
	}
	w.run(t, newStart.Add(-23*time.Hour))
	if _, reminders, _, _ := events.snapshot(); sentFor(reminders, soon.ID) != 1 {
		t.Fatalf("re-armed reminder not sent: %+v", reminders)
	}

	// Cancel withdraws the sent reminder; a cancelled or locked raid never posts or reminds.
	m2 := notify.ReminderRef{Via: notify.ReminderViaWebhook, MessageID: "m2"}
	if err := w.h.svc.Store.MarkScheduleReminded(ctx, soon.ID, m2, now); err != nil {
		t.Fatal(err)
	}
	w.statusAction(t, soon.ID, "cancel")
	if _, _, _, withdrawn := events.snapshot(); len(withdrawn) != 2 || withdrawn[1] != m2 {
		t.Fatalf("cancel withdrew %+v", withdrawn)
	}
	cancelled := w.createRaids(t, scheduleBody(startIn(10*day), 0, "true", "true"))[0]
	locked := w.createRaids(t, scheduleBody(startIn(10*day+time.Hour), 0, "true", "true"))[0]
	w.statusAction(t, cancelled.ID, "cancel")
	w.statusAction(t, locked.ID, "lock")
	w.run(t, now.Add(4*day))
	w.run(t, newStart.Add(-23*time.Hour))
	w.run(t, locked.StartsAt.Add(-23*time.Hour))
	posted, reminders, _, _ := events.snapshot()
	for _, id := range []int64{cancelled.ID, locked.ID} {
		if sentFor(posted, id) != 0 || sentFor(reminders, id) != 0 {
			t.Fatalf("raid %d posted or reminded while cancelled or locked", id)
		}
	}
	if sentFor(reminders, soon.ID) != 1 {
		t.Fatalf("cancelled raid reminded again: %+v", reminders)
	}

	// Switching the post off on an unposted raid drops its row.
	off := w.createRaids(t, scheduleBody(startIn(12*day), 0, "true", "true"))[0]
	w.patchTo(t, off.ID, off.StartsAt, "false", "true")
	if _, ok := w.scheduleOf(t, off.ID); ok {
		t.Fatal("unposted row survived post_to_discord false")
	}
}

func TestTwoScheduleRunsSendOnce(t *testing.T) {
	w, events, _ := postingWorld(t)
	raid := w.createRaids(t, scheduleBody(startIn(10*day), 0, "true", "true"))[0]
	race := func(now time.Time) int {
		var wg sync.WaitGroup
		runs := make([]ScheduleRun, 2)
		for i := range runs {
			wg.Add(1)
			go func() {
				defer wg.Done()
				got, err := w.h.svc.RunRaidSchedule(context.Background(), now)
				if err != nil {
					t.Error(err)
				}
				runs[i] = got
			}()
		}
		wg.Wait()
		return runs[0].Posted + runs[1].Posted + runs[0].Reminded + runs[1].Reminded
	}
	if n := race(time.Now().Add(4 * day)); n != 1 {
		t.Fatalf("post runs sent %d", n)
	}
	if n := race(raid.StartsAt.Add(-23 * time.Hour)); n != 1 {
		t.Fatalf("reminder runs sent %d", n)
	}
	posted, reminders, _, _ := events.snapshot()
	if sentFor(posted, raid.ID) != 1 || sentFor(reminders, raid.ID) != 1 {
		t.Fatalf("posted %d, reminded %d", sentFor(posted, raid.ID), sentFor(reminders, raid.ID))
	}
}

// discordOf reads the raw discord block of raid id's view.
func (w raidWorld) discordOf(t *testing.T, id int64) map[string]json.RawMessage {
	t.Helper()
	data := rawData(t, w.h.do(http.MethodGet, w.path("/raids/%d", id), ""))
	var out map[string]json.RawMessage
	if err := json.Unmarshal(data["discord"], &out); err != nil {
		t.Fatal(err)
	}
	return out
}

func TestDefaultsAreRememberedAndRead(t *testing.T) {
	w, _, writes := postingWorld(t)
	hook := &notify.GuildWebhook{ChannelID: "1", PostDefault: true, RemindDefault: true}
	w.h.svc.Discord = fakeDiscord{enabled: true, hook: hook, defaults: writes}
	fresh := w.createRaids(t, scheduleBody(startIn(3*day), 0, "", ""))[0]
	if d := w.getRaid(t, fresh.ID).Discord; !d.PostDefault || !d.RemindDefault || d.ChannelName != nil {
		t.Fatalf("fresh guild: discord = %+v", d)
	}
	w.createRaids(t, scheduleBody(startIn(3*day), 0, "false", "false"))
	if d := w.getRaid(t, fresh.ID).Discord; d.PostDefault || d.RemindDefault {
		t.Fatalf("after a save off/off: discord = %+v", d)
	}
	if got := string(w.discordOf(t, fresh.ID)["channel_name"]); got != "null" {
		t.Fatalf("webhook channel_name = %s", got)
	}
	// Only the remind box carried: the post default is kept.
	w.createRaids(t, scheduleBody(startIn(3*day), 0, "", "true"))
	if d := w.getRaid(t, fresh.ID).Discord; d.PostDefault || !d.RemindDefault {
		t.Fatalf("after a remind-only save: discord = %+v", d)
	}

	botWrites := &[][2]bool{}
	w.h.svc.Bot = fakeBot{on: true, active: true, postDefault: true, remindDefault: true,
		channelName: "raid-signups", defaults: botWrites}
	d := w.getRaid(t, fresh.ID).Discord
	if !d.Connected || d.ChannelName == nil || *d.ChannelName != "raid-signups" || !d.PostDefault || !d.RemindDefault {
		t.Fatalf("bot active: discord = %+v", d)
	}
	w.createRaids(t, scheduleBody(startIn(3*day), 0, "true", "false"))
	if len(*botWrites) != 1 || (*botWrites)[0] != [2]bool{true, false} {
		t.Fatalf("bot default writes = %v", *botWrites)
	}
	// An absent remind box reads as the guild's remind default.
	created := w.createRaids(t, scheduleBody(startIn(4*day), 0, "true", ""))
	if p := created[0].DiscordPost; p == nil || !p.Remind {
		t.Fatalf("absent remind box: discord_post = %+v", p)
	}
}

// An edit never rewrites the remembered defaults: editing a switch-off raid keeps the guild's.
func TestAnEditKeepsTheRememberedDefaults(t *testing.T) {
	w, _, writes := postingWorld(t)
	hook := &notify.GuildWebhook{ChannelID: "1", PostDefault: true, RemindDefault: true}
	w.h.svc.Discord = fakeDiscord{enabled: true, hook: hook, defaults: writes}
	botWrites := &[][2]bool{}
	w.h.svc.Bot = fakeBot{on: true, active: true, postDefault: true, remindDefault: true, defaults: botWrites}
	raid := w.createRaids(t, scheduleBody(startIn(3*day), 0, "", ""))[0]
	hookBefore, botBefore := len(*writes), len(*botWrites)
	w.patchTo(t, raid.ID, startIn(3*day), "false", "false")
	if len(*writes) != hookBefore || len(*botWrites) != botBefore {
		t.Fatalf("an edit wrote defaults: webhook %v, bot %v", *writes, *botWrites)
	}
	if d := w.getRaid(t, raid.ID).Discord; !d.PostDefault || !d.RemindDefault {
		t.Fatalf("after an off/off edit: discord = %+v", d)
	}
}
