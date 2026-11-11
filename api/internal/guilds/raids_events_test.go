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
	return notify.RaidPosted{RaidID: 9, Zone: "raid:onyxias-lair", Title: "Onyxia", RaidURL: "https://foreversixty.gg/guild/us/hardcore/olympus/raids/9",
		StartsAt: time.Date(2026, 11, 5, 2, 0, 0, 0, time.UTC), Timezone: "America/Chicago",
		Counts: notify.RaidCounts{In: 10, Late: 2}, Size: 40, SizeKnown: true}
}

func TestRaidEventsReadPlainlyAndThreadPerRaid(t *testing.T) {
	p := samplePosted()
	posted := raidPostedEvent(3, p)
	if posted.Body != "12 of 40 signed. Wed 4 Nov, 8 pm CST." || posted.Thread != "raid-9" || posted.Collapse != "raid.posted-9" {
		t.Fatalf("posted %+v", posted)
	}
	if len(posted.Actions) != 4 || posted.Actions[0].Path != "/v1/guilds/3/raids/9/signups/{character}" ||
		posted.Actions[2].Body["state"] != SignupBench || posted.Actions[3].Body["state"] != SignupOut ||
		posted.Actions[1].Method != http.MethodPut {
		t.Fatalf("actions %+v", posted.Actions)
	}
	if posted.Icon != "inv_misc_head_dragon_black" {
		t.Fatalf("posted icon %q", posted.Icon)
	}
	if posted.Audience.Kind != events.ToGuildMembers || posted.Data["raid_id"] != "9" {
		t.Fatalf("posted %+v", posted)
	}
	if e := raidShortEvent(3, p, reminderLead, unansweredUnknown, false); e.Body != "12 of 40 signed, 24 hours out." || e.Audience.Kind != events.ToGuildOfficers {
		t.Fatalf("short %+v", e)
	}
	moved := raidChangedEvent(3, p, false)
	cancelled := raidChangedEvent(3, p, true)
	if moved.DedupeKey == cancelled.DedupeKey || cancelled.Body != "This raid was cancelled." || moved.Collapse != cancelled.Collapse {
		t.Fatalf("moved %+v cancelled %+v", moved, cancelled)
	}
	later := p
	later.StartsAt = later.StartsAt.Add(time.Hour)
	if raidNudgeEvent(3, p, []int64{1}).DedupeKey == raidNudgeEvent(3, later, []int64{1}).DedupeKey {
		t.Fatal("a moved raid is not nudged again")
	}
	if e := lateChangeEvent(3, p, "k", "Kraggor", 5*time.Hour); e.Body != "Kraggor changed to Out, 5 hours before the start." {
		t.Fatalf("late %q", e.Body)
	}
	if e := rosterSetEvent(3, p, 7, SignupBench, time.Unix(1, 0)); e.Body != raidEventCopy.RosterBench || e.Audience.UserIDs[0] != 7 {
		t.Fatalf("roster %+v", e)
	}
	for _, e := range []events.Event{posted, moved, raidNudgeEvent(3, p, []int64{1}), raidStartingEvent(3, p, time.Hour, events.Users(1), ""), raidShortEvent(3, p, reminderLead, 0, true)} {
		if err := e.Validate(); err != nil {
			t.Fatalf("%s: %v", e.Kind, err)
		}
	}
}

func TestRaidPostedBodyCountsTheSigned(t *testing.T) {
	open := samplePosted()
	open.SizeKnown, open.Size = false, 0
	afternoon := samplePosted()
	afternoon.StartsAt = time.Date(2026, 11, 5, 15, 30, 0, 0, time.UTC)
	cases := []struct {
		name string
		p    notify.RaidPosted
		want string
	}{
		{"size known", samplePosted(), "12 of 40 signed. Wed 4 Nov, 8 pm CST."},
		{"open size", open, "12 signed. Wed 4 Nov, 8 pm CST."},
		{"morning reads am", afternoon, "12 of 40 signed. Thu 5 Nov, 9:30 am CST."},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := raidPostedEvent(3, c.p).Body; got != c.want {
				t.Fatalf("body %q, want %q", got, c.want)
			}
		})
	}
}

func TestRaidShortBodyNamesShortRolesAndUnanswered(t *testing.T) {
	withTarget := samplePosted()
	withTarget.Target = &notify.RoleTarget{Tanks: 1, Healers: 8, Melee: 0, Ranged: 0}
	withTarget.Roster = []notify.RosterLine{
		{Name: "Brakka", Class: "warrior", Spec: "Protection", Role: "tank", State: "in"},
		{Name: "Mirelle", Class: "priest", Spec: "Holy", Role: "healer", State: "in"},
	}
	cases := []struct {
		name       string
		p          notify.RaidPosted
		unanswered int
		want       string
	}{
		{"roles and unanswered", withTarget, 5, "Healers 1 of 8, 5 not answered. 12 of 40 signed, 24 hours out."},
		{"roles only, count unknown", withTarget, unansweredUnknown, "Healers 1 of 8. 12 of 40 signed, 24 hours out."},
		{"no target, unanswered", samplePosted(), 3, "3 not answered. 12 of 40 signed, 24 hours out."},
		{"nothing more known", samplePosted(), 0, "12 of 40 signed, 24 hours out."},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := raidShortEvent(3, c.p, reminderLead, c.unanswered, false).Body; got != c.want {
				t.Fatalf("body %q, want %q", got, c.want)
			}
		})
	}
}

func TestRaidShortMentionsTheShortestRole(t *testing.T) {
	short := samplePosted()
	short.Target = &notify.RoleTarget{Tanks: 2, Healers: 8, Ranged: 4}
	short.Roster = []notify.RosterLine{{Name: "Mirelle", Class: "priest", Spec: "Holy", Role: "healer", State: "in"}}
	healerShort := short
	healerShort.Target = &notify.RoleTarget{Healers: 8, Ranged: 4}
	dpsShort := samplePosted()
	dpsShort.Target = &notify.RoleTarget{Ranged: 3}
	cases := []struct {
		name    string
		p       notify.RaidPosted
		canPing bool
		want    string // the ping body's role; "" for no action
		title   string
	}{
		{"tanks first though healers are further below", short, true, RoleTank, "Mention tanks"},
		{"healers before dps", healerShort, true, RoleHealer, "Mention healers"},
		{"ranged pings dps", dpsShort, true, RoleDPS, "Mention DPS"},
		{"inside the cooldown", short, false, "", ""},
		{"no target", samplePosted(), true, "", ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			e := raidShortEvent(3, c.p, reminderLead, 0, c.canPing)
			if c.want == "" {
				if len(e.Actions) != 0 {
					t.Fatalf("actions %+v, want none", e.Actions)
				}
				return
			}
			if len(e.Actions) != 1 {
				t.Fatalf("actions %+v", e.Actions)
			}
			a := e.Actions[0]
			if a.ID != MentionShortActionID || a.Method != http.MethodPost || a.Path != "/v1/guilds/3/raids/9/ping" ||
				a.Body["role"] != c.want || a.Title != c.title || e.Data[MentionWhoDataKey] != strings.TrimPrefix(c.title, mentionVerb) {
				t.Fatalf("action %+v", a)
			}
		})
	}
}

func TestScheduledEventsSayTheActualLead(t *testing.T) {
	p := samplePosted()
	cases := []struct {
		name string
		got  string
		want string
	}{
		{"short on time", raidShortEvent(3, p, reminderLead, 0, false).Body, "12 of 40 signed, 24 hours out."},
		{"short a run late", raidShortEvent(3, p, 23*time.Hour+10*time.Minute, 0, false).Body, "12 of 40 signed, 23 hours out."},
		{"starting on time", raidStartingEvent(3, p, startingLead, events.Users(1), "").Body, "Starts in 60 minutes."},
		{"starting late", raidStartingEvent(3, p, 42*time.Minute+20*time.Second, events.Users(1), "").Body, "Starts in 42 minutes."},
		{"one minute", raidStartingEvent(3, p, time.Minute, events.Users(1), "").Body, "Starts in 1 minute."},
	}
	for _, c := range cases {
		if c.got != c.want {
			t.Errorf("%s: %q, want %q", c.name, c.got, c.want)
		}
	}
}

func TestPingOpenHonoursNextAllowedAt(t *testing.T) {
	now := time.Date(2026, 11, 4, 12, 0, 0, 0, time.UTC)
	cases := []struct {
		name string
		last *time.Time
		want bool
	}{
		{"never pinged", nil, true},
		{"inside the cooldown", ptrTime(now.Add(-time.Hour)), false},
		{"cooldown over", ptrTime(now.Add(-pingEvery)), true},
	}
	for _, c := range cases {
		if got := pingOpen(pingStateOf(c.last), now); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}

func ptrTime(t time.Time) *time.Time { return &t }

func TestRaidStartingBodyNamesTheReadyMiss(t *testing.T) {
	ok := ReadyItem{Key: "flask", State: ReadyOK, Title: "Flask of the Titans"}
	cases := []struct {
		name  string
		items []ReadyItem
		want  string
	}{
		{"no check", nil, "Starts in 60 minutes."},
		{"all ok", []ReadyItem{ok, {Key: "enchant.feet", State: ReadyWarn}}, "Starts in 60 minutes."},
		{"flask", []ReadyItem{{Key: "flask", State: ReadyMiss}, {Key: "food", State: ReadyMiss}}, "Starts in 60 minutes. No flask in your bags."},
		{"food", []ReadyItem{ok, {Key: "food", State: ReadyMiss}}, "Starts in 60 minutes. No buff food in your bags."},
		{"attunement first", []ReadyItem{{Key: "attune.onyxias-lair", State: ReadyMiss}, {Key: "food", State: ReadyMiss}}, "Starts in 60 minutes. Not attuned."},
		{"level", []ReadyItem{{Key: "level", State: ReadyMiss, Line: "You are level 58"}}, "Starts in 60 minutes. You are level 58."},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			e := raidStartingEvent(3, samplePosted(), startingLead, events.Users(7), readyMissLine(c.items))
			if e.Body != c.want || e.Audience.UserIDs[0] != 7 {
				t.Fatalf("starting %q to %+v, want %q", e.Body, e.Audience, c.want)
			}
		})
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
		// raid.posted goes out on save, never from the schedule.
		{at(7*24*time.Hour, RaidPlanned), nil},
		{at(30*24*time.Hour, RaidPlanned), nil},
		{at(24*time.Hour, RaidPlanned), []events.Kind{events.RaidShort}},
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
	// The nudge goes at each player's reminder (raids_nudge_test.go), never 24 hours out.
	if run.Nudged != 0 || run.Short != 1 || len(pub.ofKind(events.RaidSignupNudge)) != 0 {
		t.Fatalf("run %+v", run)
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
