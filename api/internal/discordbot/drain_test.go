package discordbot

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

const (
	postChannel = "chan1"
	webhookURL  = "https://discord.com/api/webhooks/9/hook-secret"
)

// RaidPostFacts lets the install harness's fake stand in for the guilds service.
func (f *fakeGuilds) RaidPostFacts(context.Context, int64) (guilds.RaidPostFacts, bool, error) {
	return guilds.RaidPostFacts{}, false, nil
}

// fakeFacts answers RaidPostFacts from a map the test edits.
type fakeFacts struct {
	mu sync.Mutex
	by map[int64]guilds.RaidPostFacts
}

func (f *fakeFacts) RaidPostFacts(_ context.Context, raidID int64) (guilds.RaidPostFacts, bool, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	facts, ok := f.by[raidID]
	return facts, ok, nil
}

func (f *fakeFacts) edit(raidID int64, change func(*guilds.RaidPostFacts)) {
	f.mu.Lock()
	defer f.mu.Unlock()
	facts := f.by[raidID]
	change(&facts)
	f.by[raidID] = facts
}

// testClock is a settable clock; it starts just after the database's now().
type testClock struct {
	mu sync.Mutex
	t  time.Time
}

func (c *testClock) now() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.t
}

func (c *testClock) advance(d time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.t = c.t.Add(d)
}

// postWorld is a guild whose bot posts in postChannel, one raid, a webhook beside it, and
// a drainer on a fake clock.
type postWorld struct {
	st     *Store
	svc    *Service
	rest   *FakeREST
	facts  *fakeFacts
	clock  *testClock
	poster *notify.FakePoster
	router *RaidRouter
	log    *bytes.Buffer
	gid    int64
	uid    int64
	raidID int64
}

func newPostWorld(t *testing.T) *postWorld {
	t.Helper()
	st := testStore(t)
	w := &postWorld{
		st: st, rest: &FakeREST{}, facts: &fakeFacts{by: map[int64]guilds.RaidPostFacts{}},
		clock: &testClock{t: time.Now().Add(time.Second)}, poster: &notify.FakePoster{}, log: &bytes.Buffer{},
	}
	w.uid = seedUser(t, st.Pool)
	w.gid = seedGuild(t, st.Pool, "Olympus XXVII")
	ctx := context.Background()
	if _, err := st.BindServer(ctx, w.gid, "srv1", "Olympus", w.uid); err != nil {
		t.Fatal(err)
	}
	if err := st.SetChannel(ctx, w.gid, postChannel, "raid-signups"); err != nil {
		t.Fatal(err)
	}
	w.rest.Layout = postableLayout()
	w.svc = New(Deps{Store: st, REST: w.rest, Log: slog_(w.log)})
	w.svc.now = w.clock.now
	w.router = &RaidRouter{Bot: w.svc, Webhook: webhookOf(t, st.Pool, w.gid, w.uid, w.poster)}
	w.raidID = w.addRaid(t)
	return w
}

// postableLayout is srv1 with postChannel and chan2, where the bot holds every permission.
func postableLayout() ServerLayout {
	return ServerLayout{
		BotUserID: "bot", Roles: []Role{{ID: "srv1", Permissions: InstallPermissions}},
		Channels: []RawChannel{{ID: postChannel, Name: "raid-signups"}, {ID: "chan2", Name: "general"}},
	}
}

// webhookOf saves a sealed webhook channel for the guild and delivers through poster.
func webhookOf(t *testing.T, pool *pgxpool.Pool, gid, uid int64, poster *notify.FakePoster) *notify.GuildDeliverer {
	t.Helper()
	sealer, err := notify.NewSealer(make([]byte, 32))
	if err != nil {
		t.Fatal(err)
	}
	box, err := sealer.Seal(webhookURL)
	if err != nil {
		t.Fatal(err)
	}
	store := &notify.GuildStore{Pool: pool}
	if err := store.Save(context.Background(), notify.GuildSaveInput{GuildID: gid, ChannelID: "555", TokenBox: box, ConnectedBy: uid}); err != nil {
		t.Fatal(err)
	}
	return &notify.GuildDeliverer{Store: store, Sealer: sealer, Poster: poster}
}

// addRaid inserts a raid starting in two days, three hours long, with open facts.
func (w *postWorld) addRaid(t *testing.T) int64 {
	t.Helper()
	starts := w.clock.now().Add(48 * time.Hour).Truncate(time.Minute)
	var id int64
	err := w.st.Pool.QueryRow(context.Background(), `insert into guild_raids
		(guild_id, zone, title, starts_at, duration_minutes, created_by)
		values ($1, 'raid:onyxias-lair', 'Onyxia''s Lair', $2, 180, $3) returning id`, w.gid, starts, w.uid).Scan(&id)
	if err != nil {
		t.Fatal(err)
	}
	raidURL := "https://site.test/guild/us/normal/olympus-xxvii/raids/1"
	w.facts.edit(id, func(f *guilds.RaidPostFacts) {
		*f = guilds.RaidPostFacts{
			GuildID: w.gid, Status: guilds.RaidPlanned, EndsAt: starts.Add(3 * time.Hour), RaidURL: raidURL,
			Post: notify.RaidPost{RaidPosted: notify.RaidPosted{
				RaidID: id, GuildName: "Olympus XXVII", Title: "Onyxia's Lair", StartsAt: starts, RaidURL: raidURL,
				Size: 40, SizeKnown: true, Counts: notify.RaidCounts{In: 1, DPS: 1},
			}},
		}
	})
	return id
}

func (w *postWorld) drainer() *Drainer {
	d := w.svc.NewDrainer(w.facts)
	d.Now = w.clock.now
	return d
}

func (w *postWorld) drain(t *testing.T) {
	t.Helper()
	if err := w.drainer().DrainOnce(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func (w *postWorld) enqueue(t *testing.T, raidID int64) {
	t.Helper()
	if queued, err := w.st.EnqueuePost(context.Background(), raidID, w.gid); err != nil || !queued {
		t.Fatalf("enqueue = %v, %v", queued, err)
	}
}

// enqueueIn writes a post row in channel directly (the store always uses the install's).
func (w *postWorld) enqueueIn(t *testing.T, raidID int64, channel string) {
	t.Helper()
	if _, err := w.st.Pool.Exec(context.Background(), `insert into guild_raid_discord_posts (raid_id, guild_id, channel_id)
		values ($1, $2, $3)`, raidID, w.gid, channel); err != nil {
		t.Fatal(err)
	}
}

func (w *postWorld) callsOf(method string) []RESTCall {
	var out []RESTCall
	for _, c := range w.rest.Calls() {
		if c.Method == method {
			out = append(out, c)
		}
	}
	return out
}

// postRow is a post row as the tests read it.
type postRow struct {
	messageID   *string
	movedFrom   *time.Time
	dirtySeq    int64
	dirty, gone bool
	attempts    int
	nextAttempt time.Time
	wake        *time.Time
}

func (w *postWorld) row(t *testing.T, raidID int64) (postRow, bool) {
	t.Helper()
	var r postRow
	err := w.st.Pool.QueryRow(context.Background(), `select message_id, moved_from, dirty_seq,
		dirty_at is not null, gone, attempts, next_attempt_at, wake_at
		from guild_raid_discord_posts where raid_id = $1`, raidID).
		Scan(&r.messageID, &r.movedFrom, &r.dirtySeq, &r.dirty, &r.gone, &r.attempts, &r.nextAttempt, &r.wake)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return postRow{}, false
		}
		t.Fatal(err)
	}
	return r, true
}

func signedUp(t *testing.T, m Message) string {
	t.Helper()
	for _, f := range m.Embeds[0].Fields {
		if f.Name == "Signed up" {
			return f.Value
		}
	}
	t.Fatalf("no Signed up field in %+v", m.Embeds[0].Fields)
	return ""
}

func TestDrainCreatesThenCoalesces(t *testing.T) {
	w := newPostWorld(t)
	ctx := context.Background()
	w.enqueue(t, w.raidID)
	w.drain(t)
	creates := w.callsOf("CreateMessage")
	if len(creates) != 1 || creates[0].ChannelID != postChannel || len(creates[0].Message.Components) != 1 {
		t.Fatalf("creates = %+v", creates)
	}
	if r, _ := w.row(t, w.raidID); r.messageID == nil || *r.messageID != "m1" || r.dirty || r.wake == nil {
		t.Fatalf("row after create = %+v", r)
	}
	for i := 2; i <= 6; i++ {
		w.facts.edit(w.raidID, func(f *guilds.RaidPostFacts) { f.Post.Counts.In, f.Post.Counts.DPS = i, i })
		w.st.MarkRaidDirty(ctx, w.raidID)
		w.clock.advance(200 * time.Millisecond)
		w.drain(t)
	}
	if n := len(w.rest.Calls()); n != 1 {
		t.Fatalf("calls inside the debounce = %d, want 1", n)
	}
	w.clock.advance(2 * time.Second)
	w.drain(t)
	w.drain(t)
	edits := w.callsOf("EditMessage")
	if len(edits) != 1 || edits[0].MessageID != "m1" || edits[0].ChannelID != postChannel {
		t.Fatalf("edits = %+v", edits)
	}
	if got := signedUp(t, edits[0].Message); !strings.HasPrefix(got, "**6 of 40**") {
		t.Fatalf("edit counts = %q, want the final 6", got)
	}
	if r, _ := w.row(t, w.raidID); r.dirty || r.attempts != 0 {
		t.Fatalf("row after edit = %+v", r)
	}
}

func TestTwoDrainersPostOnce(t *testing.T) {
	w := newPostWorld(t)
	w.enqueue(t, w.raidID)
	// Discord is slow, so both drainers are inside DrainOnce at once.
	w.rest.OnCall = func(RESTCall) { time.Sleep(100 * time.Millisecond) }
	start := make(chan struct{})
	var wg sync.WaitGroup
	errs := make(chan error, 2)
	for range 2 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			errs <- w.drainer().DrainOnce(context.Background())
		}()
	}
	close(start)
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			t.Fatal(err)
		}
	}
	if n := len(w.callsOf("CreateMessage")); n != 1 {
		t.Fatalf("creates = %d, want exactly 1", n)
	}
}

func TestClaimDueHoldsTheLease(t *testing.T) {
	w := newPostWorld(t)
	ctx := context.Background()
	w.enqueue(t, w.raidID)
	now := w.clock.now()
	first, err := w.st.ClaimDue(ctx, now, 2*time.Second, 30*time.Second, 10)
	if err != nil || len(first) != 1 || first[0].MessageID != nil || first[0].DirtySeq != 1 {
		t.Fatalf("first claim = %+v, %v", first, err)
	}
	if again, err := w.st.ClaimDue(ctx, now.Add(29*time.Second), 2*time.Second, 30*time.Second, 10); err != nil || len(again) != 0 {
		t.Fatalf("claim inside the lease = %+v, %v", again, err)
	}
	if lapsed, err := w.st.ClaimDue(ctx, now.Add(31*time.Second), 2*time.Second, 30*time.Second, 10); err != nil || len(lapsed) != 1 {
		t.Fatalf("claim after the lease = %+v, %v", lapsed, err)
	}
}

func TestMarkDuringEditIsKept(t *testing.T) {
	w := newPostWorld(t)
	ctx := context.Background()
	w.enqueue(t, w.raidID)
	w.drain(t)
	w.st.MarkRaidDirty(ctx, w.raidID)
	w.clock.advance(3 * time.Second)
	w.rest.OnCall = func(c RESTCall) {
		if c.Method == "EditMessage" {
			w.st.MarkRaidDirty(ctx, w.raidID)
		}
	}
	w.drain(t)
	w.rest.OnCall = nil
	if r, _ := w.row(t, w.raidID); !r.dirty {
		t.Fatalf("the mark made during the edit was lost: %+v", r)
	}
	w.clock.advance(3 * time.Second)
	w.drain(t)
	if n := len(w.callsOf("EditMessage")); n != 2 {
		t.Fatalf("edits = %d, want 2", n)
	}
	if r, _ := w.row(t, w.raidID); r.dirty {
		t.Fatalf("row still dirty after the second edit: %+v", r)
	}
}

func TestDrainHonoursDiscordRefusals(t *testing.T) {
	t.Run("429 waits retry_after", func(t *testing.T) {
		w := newPostWorld(t)
		w.enqueue(t, w.raidID)
		w.rest.ErrFor = map[string]error{"CreateMessage": &APIError{Status: http.StatusTooManyRequests, RetryAfter: 7 * time.Second}}
		w.drain(t)
		r, _ := w.row(t, w.raidID)
		// A 429 is a wait, not a failed attempt.
		if r.nextAttempt.Before(w.clock.now().Add(7*time.Second)) || r.attempts != 0 || r.messageID != nil || !r.dirty {
			t.Fatalf("row = %+v, now %v", r, w.clock.now())
		}
		w.clock.advance(5 * time.Second)
		w.drain(t)
		if n := len(w.rest.Calls()); n != 1 {
			t.Fatalf("retried before retry_after: %d calls", n)
		}
	})
	t.Run("403 50001 removes the install and the router falls back", func(t *testing.T) {
		w := newPostWorld(t)
		w.enqueue(t, w.raidID)
		missing := &APIError{Status: http.StatusForbidden, Code: codeMissingAccess}
		w.rest.ErrFor = map[string]error{"CreateMessage": missing, "ServerLayout": missing}
		w.drain(t)
		in, _, err := w.st.InstallFor(context.Background(), w.gid)
		if err != nil || !in.RemovedInDiscord {
			t.Fatalf("install = %+v, %v", in, err)
		}
		w.clock.advance(time.Hour)
		w.drain(t)
		if n := len(w.callsOf("CreateMessage")); n != 1 {
			t.Fatalf("the refused row was retried: %d creates", n)
		}
		other := w.addRaid(t)
		w.router.RaidPosted(context.Background(), w.gid, notify.RaidPosted{RaidID: other, GuildName: "Olympus XXVII", Title: "Next"})
		if _, found := w.row(t, other); found || len(w.poster.Posts()) != 1 {
			t.Fatalf("router after removal: row %v, webhook posts %d", found, len(w.poster.Posts()))
		}
	})
	t.Run("404 10008 is gone and never re-posted", func(t *testing.T) {
		w := newPostWorld(t)
		ctx := context.Background()
		w.enqueue(t, w.raidID)
		w.drain(t)
		w.rest.ErrFor = map[string]error{"EditMessage": &APIError{Status: http.StatusNotFound, Code: codeUnknownMessage}}
		w.st.MarkRaidDirty(ctx, w.raidID)
		w.clock.advance(3 * time.Second)
		w.drain(t)
		if r, _ := w.row(t, w.raidID); !r.gone {
			t.Fatalf("row = %+v, want gone", r)
		}
		w.st.MarkRaidDirty(ctx, w.raidID)
		was := w.clock.now()
		w.router.RaidChanged(ctx, w.gid, notify.RaidChanged{RaidPosted: notify.RaidPosted{RaidID: w.raidID}, Was: &was})
		w.clock.advance(time.Hour)
		w.drain(t)
		if len(w.callsOf("CreateMessage")) != 1 || len(w.callsOf("EditMessage")) != 1 || len(w.poster.Posts()) != 0 {
			t.Fatalf("a gone post was touched again: %+v, webhook %d", w.rest.Calls(), len(w.poster.Posts()))
		}
	})
	t.Run("an expired-token code is not a deleted post", func(t *testing.T) {
		w := newPostWorld(t)
		w.enqueue(t, w.raidID)
		w.rest.ErrFor = map[string]error{"CreateMessage": &APIError{Status: http.StatusNotFound, Code: codeUnknownWebhook}}
		w.drain(t)
		if r, _ := w.row(t, w.raidID); r.gone || r.attempts != 1 {
			t.Fatalf("row = %+v, want a retry", r)
		}
	})
	t.Run("a global 429 holds the rest of the batch", func(t *testing.T) {
		w := newPostWorld(t)
		second := w.addRaid(t)
		w.enqueue(t, w.raidID)
		w.enqueueIn(t, second, "chan2")
		w.rest.ErrFor = map[string]error{"CreateMessage": &APIError{Status: http.StatusTooManyRequests, Global: true, RetryAfter: time.Second}}
		w.drain(t)
		if n := len(w.rest.Calls()); n != 1 {
			t.Fatalf("calls after a global 429 = %d, want 1", n)
		}
		for _, id := range []int64{w.raidID, second} {
			// Neither counts an attempt: one was a 429, the other held without a request.
			if r, _ := w.row(t, id); r.attempts != 0 || !r.nextAttempt.After(w.clock.now()) {
				t.Fatalf("raid %d row = %+v", id, r)
			}
		}
	})
}

func TestThePostEndsOnItsOwn(t *testing.T) {
	w := newPostWorld(t)
	w.enqueue(t, w.raidID)
	w.drain(t)
	ends := w.facts.by[w.raidID].EndsAt
	w.clock.advance(ends.Sub(w.clock.now()) - time.Minute)
	w.drain(t)
	if n := len(w.rest.Calls()); n != 1 {
		t.Fatalf("calls before the end = %d", n)
	}
	w.clock.advance(2 * time.Minute)
	w.drain(t)
	edits := w.callsOf("EditMessage")
	if len(edits) != 1 {
		t.Fatalf("edits = %d, want the end-time edit", len(edits))
	}
	m := edits[0].Message
	if m.Embeds[0].Description != "This raid has ended." {
		t.Fatalf("description = %q", m.Embeds[0].Description)
	}
	for _, c := range m.Components[0] {
		if c.Style != StyleLink && !c.Disabled {
			t.Fatalf("button %q is live on an ended raid", c.Label)
		}
	}
	if r, _ := w.row(t, w.raidID); r.wake != nil {
		t.Fatalf("an ended post still wakes: %+v", r)
	}
	w.clock.advance(24 * time.Hour)
	w.drain(t)
	if n := len(w.rest.Calls()); n != 2 {
		t.Fatalf("calls after the end = %d, want 2", n)
	}
}

func TestAMovedPostWakesAtTheStart(t *testing.T) {
	w := newPostWorld(t)
	ctx := context.Background()
	was := w.clock.now().Add(24 * time.Hour)
	w.router.RaidChanged(ctx, w.gid, notify.RaidChanged{RaidPosted: notify.RaidPosted{RaidID: w.raidID}, Was: &was})
	w.drain(t)
	creates := w.callsOf("CreateMessage")
	if len(creates) != 1 || !strings.HasPrefix(creates[0].Message.Embeds[0].Title, "Moved: ") {
		t.Fatalf("creates = %+v", creates)
	}
	r, _ := w.row(t, w.raidID)
	if starts := w.facts.by[w.raidID].Post.StartsAt; r.wake == nil || !r.wake.Equal(starts) {
		t.Fatalf("wake = %v, want the start %v", r.wake, starts)
	}
}

func TestADeletedRaidDropsItsRow(t *testing.T) {
	w := newPostWorld(t)
	w.enqueue(t, w.raidID)
	w.facts.mu.Lock()
	delete(w.facts.by, w.raidID)
	w.facts.mu.Unlock()
	w.drain(t)
	if _, found := w.row(t, w.raidID); found || len(w.rest.Calls()) != 0 {
		t.Fatalf("row kept (%v) or Discord called: %+v", found, w.rest.Calls())
	}
}

func TestAnotherFailureBacksOffAndCounts(t *testing.T) {
	w := newPostWorld(t)
	w.enqueue(t, w.raidID)
	w.rest.ErrFor = map[string]error{"CreateMessage": ErrTimeout}
	w.drain(t)
	r, _ := w.row(t, w.raidID)
	if r.attempts != 1 || r.nextAttempt.Before(w.clock.now().Add(2*time.Second)) {
		t.Fatalf("row = %+v", r)
	}
	in, _, _ := w.st.InstallFor(context.Background(), w.gid)
	if in.Failures != 1 || in.RemovedInDiscord {
		t.Fatalf("install = %+v", in)
	}
}

func TestBackoff(t *testing.T) {
	cases := []struct {
		attempts   int
		retryAfter time.Duration
		want       time.Duration
	}{
		{0, 0, 2 * time.Second},
		{1, 0, 4 * time.Second},
		{3, 0, 16 * time.Second},
		{8, 0, 5 * time.Minute},
		{60, 0, 5 * time.Minute},
		{0, 7 * time.Second, 7 * time.Second},
		{2, 3 * time.Second, 8 * time.Second},
		{20, 10 * time.Minute, 10 * time.Minute},
	}
	for _, c := range cases {
		if got := backoff(c.attempts, c.retryAfter); got != c.want {
			t.Errorf("backoff(%d, %v) = %v, want %v", c.attempts, c.retryAfter, got, c.want)
		}
	}
}

func TestAPIErrorReadsTheGlobalLimit(t *testing.T) {
	rest, _ := serve(t, http.StatusTooManyRequests, `{"message":"You are being rate limited.","retry_after":0.5,"global":true}`)
	_, err := rest.CreateMessage(context.Background(), "c1", Message{})
	var api *APIError
	if !errors.As(err, &api) || !api.Global || !api.RateLimited() || api.RetryAfter != 500*time.Millisecond {
		t.Fatalf("err = %#v", err)
	}
}

func TestDrainStopsTheDrainer(t *testing.T) {
	w := newPostWorld(t)
	w.enqueue(t, w.raidID)
	d := w.drainer()
	d.Tick = 10 * time.Millisecond
	w.svc.StartDrainer(context.Background(), d)
	deadline := time.Now().Add(5 * time.Second)
	for len(w.callsOf("CreateMessage")) == 0 && time.Now().Before(deadline) {
		time.Sleep(10 * time.Millisecond)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := w.svc.Drain(ctx); err != nil {
		t.Fatal(err)
	}
	if n := len(w.callsOf("CreateMessage")); n != 1 {
		t.Fatalf("creates = %d", n)
	}
}
