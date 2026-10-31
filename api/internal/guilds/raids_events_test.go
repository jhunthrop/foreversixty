package guilds

import (
	"context"
	"net/http"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/events"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// recordingPublisher records every published event; the request producers publish off the
// request, so a test waits for them.
type recordingPublisher struct {
	mu  sync.Mutex
	got []events.Event
}

func (p *recordingPublisher) Publish(_ context.Context, e events.Event) error {
	if err := e.Validate(); err != nil {
		return err
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.got = append(p.got, e)
	return nil
}

func (p *recordingPublisher) ofKind(kind events.Kind) []events.Event {
	p.mu.Lock()
	defer p.mu.Unlock()
	var out []events.Event
	for _, e := range p.got {
		if e.Kind == kind {
			out = append(out, e)
		}
	}
	return out
}

// waitFor waits until n events of kind arrived and returns them.
func (p *recordingPublisher) waitFor(t *testing.T, kind events.Kind, n int) []events.Event {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		if got := p.ofKind(kind); len(got) >= n {
			return got
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatalf("%s: got %d, want %d", kind, len(p.ofKind(kind)), n)
	return nil
}

func samplePosted() notify.RaidPosted {
	return notify.RaidPosted{RaidID: 9, Title: "Onyxia", RaidURL: "https://foreversixty.gg/guild/us/hardcore/olympus/raids/9",
		StartsAt: time.Date(2026, 11, 5, 2, 0, 0, 0, time.UTC), Timezone: "America/Chicago",
		Counts: notify.RaidCounts{In: 10, Late: 2}, Size: 40, SizeKnown: true}
}

func TestRaidEventsReadPlainlyAndThreadPerRaid(t *testing.T) {
	p := samplePosted()
	posted := raidPostedEvent(3, p)
	if posted.Body != "Sign-ups are open: Wed 4 Nov, 20:00 CST." || posted.Thread != "raid-9" || posted.Collapse != "raid.posted-9" {
		t.Fatalf("posted %+v", posted)
	}
	if len(posted.Actions) != 3 || posted.Actions[0].Path != "/v1/guilds/3/raids/9/signups/{character}" ||
		posted.Actions[2].Body["state"] != SignupOut || posted.Actions[1].Method != http.MethodPut {
		t.Fatalf("actions %+v", posted.Actions)
	}
	if posted.Audience.Kind != events.ToGuildMembers || posted.Data["raid_id"] != "9" {
		t.Fatalf("posted %+v", posted)
	}
	if e := raidShortEvent(3, p); e.Body != "12 of 40 signed, 24 hours out." || e.Audience.Kind != events.ToGuildOfficers {
		t.Fatalf("short %+v", e)
	}
	moved := raidChangedEvent(3, p, false)
	cancelled := raidChangedEvent(3, p, true)
	if moved.DedupeKey == cancelled.DedupeKey || cancelled.Body != "This raid was cancelled." || moved.Collapse != cancelled.Collapse {
		t.Fatalf("moved %+v cancelled %+v", moved, cancelled)
	}
	later := p
	later.StartsAt = later.StartsAt.Add(time.Hour)
	if raidNudgeEvent(3, p).DedupeKey == raidNudgeEvent(3, later).DedupeKey {
		t.Fatal("a moved raid is not nudged again")
	}
	if e := lateChangeEvent(3, p, "k", "Kraggor", 5*time.Hour); e.Body != "Kraggor changed to Out, 5 hours before the start." {
		t.Fatalf("late %q", e.Body)
	}
	if e := rosterSetEvent(3, p, 7, SignupBench, time.Unix(1, 0)); e.Body != raidEventCopy.RosterBench || e.Audience.UserIDs[0] != 7 {
		t.Fatalf("roster %+v", e)
	}
	for _, e := range []events.Event{posted, moved, raidNudgeEvent(3, p), raidStartingEvent(3, p, time.Hour), raidShortEvent(3, p)} {
		if err := e.Validate(); err != nil {
			t.Fatalf("%s: %v", e.Kind, err)
		}
	}
}

func TestDueKinds(t *testing.T) {
	now := time.Date(2026, 11, 4, 12, 0, 0, 0, time.UTC)
	at := func(d time.Duration, status string) scheduledRaid {
		return scheduledRaid{ID: 1, StartsAt: now.Add(d), Status: status}
	}
	cases := []struct {
		r    scheduledRaid
		want []events.Kind
	}{
		{at(postHorizon, RaidPlanned), []events.Kind{events.RaidPosted}},
		{at(postHorizon-30*time.Minute, RaidPlanned), []events.Kind{events.RaidPosted}},
		{at(postHorizon-2*time.Hour, RaidPlanned), nil},
		{at(24*time.Hour, RaidPlanned), []events.Kind{events.RaidSignupNudge, events.RaidShort}},
		{at(24*time.Hour, RaidLocked), nil},
		{at(30*time.Minute, RaidLocked), []events.Kind{events.RaidStarting}},
		{at(60*time.Minute, RaidPlanned), []events.Kind{events.RaidStarting}},
		{at(61*time.Minute, RaidPlanned), nil},
	}
	for _, c := range cases {
		if got := dueKinds(c.r, now); !slices.Equal(got, c.want) {
			t.Errorf("%v %s: %v, want %v", c.r.StartsAt.Sub(now), c.r.Status, got, c.want)
		}
	}
}

func eventsWorld(t *testing.T) (raidWorld, *recordingPublisher) {
	t.Helper()
	w := newRaidWorld(t, true)
	pub := &recordingPublisher{}
	w.h.svc.Events = pub
	return w, pub
}

func TestRaidWritesPublishTheirEvents(t *testing.T) {
	w, pub := eventsWorld(t)
	raid := w.createRaid(t)
	posted := pub.waitFor(t, events.RaidPosted, 1)[0]
	if posted.Audience.GuildID != w.gid || posted.Data["raid_id"] == "" || !strings.Contains(posted.Link, "/raids/") {
		t.Fatalf("posted %+v", posted)
	}
	w.as(w.member)
	w.signup(t, raid.ID, memberKeyA, `{"state":"bench"}`)

	w.as(w.officer)
	moved := time.Now().Add(8 * 24 * time.Hour).UTC().Truncate(time.Minute).Format(time.RFC3339)
	res := w.h.do(http.MethodPatch, w.path("/raids/%d", raid.ID), raidBody(moved, 0))
	if res.StatusCode != http.StatusOK {
		t.Fatalf("patch %d %s", res.StatusCode, readBody(res))
	}
	if e := pub.waitFor(t, events.RaidChanged, 1)[0]; e.Audience.Kind != events.ToRaidSigned {
		t.Fatalf("changed %+v", e)
	}
	if res := w.h.do(http.MethodPost, w.path("/raids/%d/lock", raid.ID), ""); res.StatusCode != http.StatusOK {
		t.Fatalf("lock %d", res.StatusCode)
	}
	roster := pub.waitFor(t, events.RaidRosterSet, 1)
	if len(roster) != 1 || roster[0].Audience.UserIDs[0] != w.member || roster[0].Body != raidEventCopy.RosterBench {
		t.Fatalf("roster %+v", roster)
	}
	if res := w.h.do(http.MethodPost, w.path("/raids/%d/cancel", raid.ID), ""); res.StatusCode != http.StatusOK {
		t.Fatalf("cancel %d", res.StatusCode)
	}
	if got := pub.waitFor(t, events.RaidChanged, 2); got[1].Body != raidEventCopy.Cancelled {
		t.Fatalf("cancelled %+v", got[1])
	}
}

func TestAnOutInsideADayTellsTheOfficers(t *testing.T) {
	w, pub := eventsWorld(t)
	soon := time.Now().Add(5 * time.Hour).UTC().Truncate(time.Minute).Format(time.RFC3339)
	raid := w.createRaids(t, raidBody(soon, 0))[0]
	w.as(w.member)
	w.signup(t, raid.ID, memberKeyA, `{"state":"in"}`)
	w.signup(t, raid.ID, memberKeyA, `{"state":"out"}`)
	late := pub.waitFor(t, events.RaidLateChange, 1)[0]
	if late.Audience.Kind != events.ToGuildOfficers || !strings.HasPrefix(late.Body, "Kraggor changed to Out") {
		t.Fatalf("late %+v", late)
	}
	// A week out, an Out tells nobody.
	far := w.createRaid(t)
	w.as(w.member)
	w.signup(t, far.ID, memberKeyA, `{"state":"out"}`)
	time.Sleep(100 * time.Millisecond)
	if n := len(pub.ofKind(events.RaidLateChange)); n != 1 {
		t.Fatalf("late changes %d", n)
	}
}

func TestRunEventSchedulePublishesWhatIsDue(t *testing.T) {
	w, pub := eventsWorld(t)
	soon := time.Now().Add(24 * time.Hour).UTC().Truncate(time.Minute).Format(time.RFC3339)
	raid := w.createRaids(t, raidBody(soon, 0))[0]
	pub.waitFor(t, events.RaidPosted, 1)
	run, err := w.h.svc.RunEventSchedule(context.Background(), raid.StartsAt.Add(-24*time.Hour))
	if err != nil {
		t.Fatal(err)
	}
	if run.Nudged != 1 || run.Short != 1 {
		t.Fatalf("run %+v", run)
	}
	if e := pub.ofKind(events.RaidSignupNudge); len(e) != 1 || e[0].Audience.Kind != events.ToUnanswered || len(e[0].Actions) != 3 {
		t.Fatalf("nudge %+v", e)
	}
	if len(pub.ofKind(events.RaidShort)) != 1 {
		t.Fatal("an empty Onyxia is short")
	}
	run, err = w.h.svc.RunEventSchedule(context.Background(), raid.StartsAt.Add(-45*time.Minute))
	if err != nil || run.Starting != 1 || pub.ofKind(events.RaidStarting)[0].Audience.Kind != events.ToRaidInLate {
		t.Fatalf("starting %+v %v", run, err)
	}
	if run.Pending != 1 || len(pub.ofKind(events.RosterPending)) == 0 {
		t.Fatalf("pending %+v", run)
	}
	if e := pub.ofKind(events.RosterPending)[0]; e.Body != "1 character is waiting for approval." || e.Audience.Kind != events.ToGuildOfficers {
		t.Fatalf("pending %+v", e)
	}
	w.h.svc.Events = nil
	if run, err := w.h.svc.RunEventSchedule(context.Background(), time.Now()); err != nil || run != (EventScheduleRun{}) {
		t.Fatalf("without events %+v %v", run, err)
	}
}

func TestAnOfficerMovingAPlayerOnASetRosterTellsThem(t *testing.T) {
	w, pub := eventsWorld(t)
	raid := w.createRaid(t)
	w.as(w.member)
	w.signup(t, raid.ID, memberKeyA, `{"state":"in"}`)
	// Before the roster is set an officer's change tells nothing.
	w.as(w.officer)
	w.signup(t, raid.ID, memberKeyA, `{"state":"late"}`)
	if res := w.h.do(http.MethodPost, w.path("/raids/%d/lock", raid.ID), ""); res.StatusCode != http.StatusOK {
		t.Fatalf("lock %d", res.StatusCode)
	}
	w.signup(t, raid.ID, memberKeyA, `{"state":"bench"}`)
	w.signup(t, raid.ID, memberKeyA, `{"state":"bench"}`) // the same place again tells nothing
	w.signup(t, raid.ID, member2Key, `{"state":"in"}`)    // added by the officer
	got := pub.waitFor(t, events.RaidRosterMoved, 2)
	time.Sleep(100 * time.Millisecond)
	got = pub.ofKind(events.RaidRosterMoved)
	if len(got) != 2 {
		t.Fatalf("moves %+v", got)
	}
	bodies := map[int64]string{got[0].Audience.UserIDs[0]: got[0].Body, got[1].Audience.UserIDs[0]: got[1].Body}
	if bodies[w.member] != "An officer moved Kraggor to the bench." || bodies[w.member2] != "An officer moved Mirelle to In." {
		t.Fatalf("bodies %v", bodies)
	}
}
