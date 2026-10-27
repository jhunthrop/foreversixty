package discordbot

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"slices"
	"strings"
	"testing"
	"time"
)

func (w *postWorld) install(t *testing.T) Install {
	t.Helper()
	in, _, err := w.st.InstallFor(context.Background(), w.gid)
	if err != nil {
		t.Fatal(err)
	}
	return in
}

func (w *postWorld) parked(t *testing.T, raidID int64) bool {
	t.Helper()
	var parked bool
	if err := w.st.Pool.QueryRow(context.Background(), `select next_attempt_at = 'infinity'
		from guild_raid_discord_posts where raid_id = $1`, raidID).Scan(&parked); err != nil {
		t.Fatal(err)
	}
	return parked
}

func TestAMissingPermissionMarksTheChannelUnusable(t *testing.T) {
	w := newPostWorld(t)
	ctx := context.Background()
	w.enqueue(t, w.raidID)
	w.rest.Layout.Channels[0].Overwrites = []Overwrite{{ID: "srv1", Type: OverwriteRole, Deny: PermSendMessages | PermEmbedLinks}}
	w.rest.ErrFor = map[string]error{"CreateMessage": &APIError{Status: http.StatusForbidden, Code: codeMissingPerms}}
	w.drain(t)
	in := w.install(t)
	if in.RemovedInDiscord || in.ChannelUnusable == nil || in.ChannelUnusable.Reason != UnusablePermissions ||
		!slices.Equal(in.ChannelUnusable.Missing, []string{"Send Messages", "Embed Links"}) {
		t.Fatalf("install = %+v unusable %+v", in, in.ChannelUnusable)
	}
	if _, active, _ := w.svc.Active(ctx, w.gid); active {
		t.Fatal("an unusable channel is still active")
	}
	other := w.addRaid(t)
	w.router.RaidPosted(ctx, w.gid, posted(w, other))
	if len(w.poster.Posts()) != 1 {
		t.Fatalf("the webhook posted %d, want 1", len(w.poster.Posts()))
	}
	w.clock.advance(time.Hour)
	w.drain(t)
	if n := len(w.callsOf("CreateMessage")); n != 1 {
		t.Fatalf("the paused row was tried again: %d creates", n)
	}
	// The officer fixes the permission and saves the same channel: the row drains.
	w.rest.ErrFor = nil
	if err := w.st.SetChannel(ctx, w.gid, postChannel, "raid-signups"); err != nil {
		t.Fatal(err)
	}
	if in := w.install(t); in.ChannelUnusable != nil {
		t.Fatalf("save kept the unusable state: %+v", in.ChannelUnusable)
	}
	w.clock.advance(time.Second)
	w.drain(t)
	if r, _ := w.row(t, w.raidID); r.messageID == nil {
		t.Fatalf("row after the save = %+v", r)
	}
}

func TestADeletedChannelIsUnusableNotGone(t *testing.T) {
	w := newPostWorld(t)
	w.enqueue(t, w.raidID)
	w.rest.Layout.Channels = w.rest.Layout.Channels[1:]
	w.rest.ErrFor = map[string]error{"CreateMessage": &APIError{Status: http.StatusNotFound, Code: codeUnknownChannel}}
	w.drain(t)
	in := w.install(t)
	if in.ChannelUnusable == nil || in.ChannelUnusable.Reason != UnusableDeleted || len(in.ChannelUnusable.Missing) != 0 {
		t.Fatalf("unusable = %+v", in.ChannelUnusable)
	}
	if r, _ := w.row(t, w.raidID); r.gone {
		t.Fatal("a deleted channel marked the post gone")
	}
}

func TestAChannelRefusalTheServerReadCannotExplainRetries(t *testing.T) {
	w := newPostWorld(t)
	w.enqueue(t, w.raidID)
	w.rest.ErrFor = map[string]error{"CreateMessage": &APIError{Status: http.StatusForbidden, Code: codeMissingPerms}}
	w.drain(t)
	in := w.install(t)
	if in.ChannelUnusable != nil || in.RemovedInDiscord {
		t.Fatalf("install = %+v", in)
	}
	if r, _ := w.row(t, w.raidID); r.attempts != 1 {
		t.Fatalf("row = %+v, want a retry", r)
	}
}

func TestTheSettingsReadCarriesChannelUnusable(t *testing.T) {
	w := newPostWorld(t)
	ctx := context.Background()
	if err := w.st.MarkChannelUnusable(ctx, w.gid, ChannelUnusable{Reason: UnusablePermissions, Missing: []string{"View Channel"}}, "x"); err != nil {
		t.Fatal(err)
	}
	view, err := w.svc.installView(ctx, w.gid)
	if err != nil {
		t.Fatal(err)
	}
	raw, err := json.Marshal(view)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), `"channel_unusable":{"reason":"permissions","missing":["View Channel"]}`) {
		t.Fatalf("settings bot = %s", raw)
	}
	if err := w.st.SetChannel(ctx, w.gid, postChannel, "raid-signups"); err != nil {
		t.Fatal(err)
	}
	view, _ = w.svc.installView(ctx, w.gid)
	raw, _ = json.Marshal(view)
	if !strings.Contains(string(raw), `"channel_unusable":null`) {
		t.Fatalf("after save = %s", raw)
	}
}

func TestTheTestMessageUsesTheDrainMapping(t *testing.T) {
	h := newInstallHarness(t)
	h.callback(h.gid, "code=c", h.officerActor())
	if rec := h.do(http.MethodPut, h.path(h.gid, "/channel"), `{"channel_id":"c1"}`); rec.Code != http.StatusOK {
		t.Fatalf("save = %d", rec.Code)
	}
	h.rest.Layout.Channels = h.rest.Layout.Channels[:1]
	h.rest.ErrFor = map[string]error{"CreateMessage": &APIError{Status: http.StatusNotFound, Code: codeUnknownChannel}}
	h.do(http.MethodPost, h.path(h.gid, "/test"), "")
	in, _ := h.installFor(h.gid)
	if in.ChannelUnusable == nil || in.ChannelUnusable.Reason != UnusableDeleted || in.RemovedInDiscord {
		t.Fatalf("install = %+v", in)
	}
}

func TestTheCreateCarriesTheRowNonce(t *testing.T) {
	w := newPostWorld(t)
	w.enqueue(t, w.raidID)
	w.drain(t)
	creates := w.callsOf("CreateMessage")
	if len(creates) != 1 {
		t.Fatalf("creates = %d", len(creates))
	}
	nonce := creates[0].Message.Nonce
	if !strings.HasPrefix(nonce, "r") || len([]rune(nonce)) > maxNonceRunes {
		t.Fatalf("nonce = %q", nonce)
	}
	raw, err := json.Marshal(creates[0].Message)
	if err != nil || !strings.Contains(string(raw), `"nonce":"`+nonce+`","enforce_nonce":true`) {
		t.Fatalf("create body = %s, %v", raw, err)
	}
	if long := postNonce(Post{RaidID: 1<<63 - 1, CreatedAt: time.UnixMilli(1 << 46)}); len(long) > maxNonceRunes {
		t.Fatalf("the largest nonce %q is over the limit", long)
	}
	if _, err := json.Marshal(Message{Nonce: strings.Repeat("n", maxNonceRunes+1)}); err == nil {
		t.Fatal("an over-long nonce marshalled")
	}
	w.st.MarkRaidDirty(context.Background(), w.raidID)
	w.clock.advance(3 * time.Second)
	w.drain(t)
	if edits := w.callsOf("EditMessage"); len(edits) != 1 || edits[0].Message.Nonce != "" {
		t.Fatalf("edits = %+v", edits)
	}
}

func TestClaimWritesNeedTheLease(t *testing.T) {
	w := newPostWorld(t)
	ctx := context.Background()
	w.enqueue(t, w.raidID)
	now := w.clock.now()
	claimed, err := w.st.ClaimDue(ctx, now, 2*time.Second, 30*time.Second, 10)
	if err != nil || len(claimed) != 1 {
		t.Fatalf("claim = %+v, %v", claimed, err)
	}
	stale := claimed[0]
	// The lease lapses and a second drainer takes the row.
	if again, err := w.st.ClaimDue(ctx, now.Add(31*time.Second), 2*time.Second, 30*time.Second, 10); err != nil || len(again) != 1 {
		t.Fatalf("second claim = %+v, %v", again, err)
	}
	for name, write := range map[string]func() error{
		"finish": func() error { return w.st.FinishEdit(ctx, stale, "m9", nil, now) },
		"fail":   func() error { return w.st.FailEdit(ctx, stale, now, "x") },
		"gone":   func() error { return w.st.MarkGone(ctx, stale) },
		"defer":  func() error { return w.st.DeferPost(ctx, stale, now) },
		"park":   func() error { return w.st.ParkPost(ctx, stale, "x") },
	} {
		if err := write(); !errors.Is(err, ErrLeaseLost) {
			t.Fatalf("%s with a lost lease = %v", name, err)
		}
	}
	if r, _ := w.row(t, w.raidID); r.messageID != nil || r.gone {
		t.Fatalf("a stale write landed: %+v", r)
	}
}

func TestRowsReachedNearTheLeaseEndAreReleased(t *testing.T) {
	w := newPostWorld(t)
	second := w.addRaid(t)
	w.enqueue(t, w.raidID)
	w.enqueue(t, second)
	// The first create takes so long that the second row would start past its lease.
	w.rest.OnCall = func(RESTCall) { w.clock.advance(25 * time.Second) }
	w.drain(t)
	if n := len(w.callsOf("CreateMessage")); n != 1 {
		t.Fatalf("creates = %d, want 1", n)
	}
	var leased int
	if err := w.st.Pool.QueryRow(context.Background(),
		`select count(*) from guild_raid_discord_posts where lease_until is not null`).Scan(&leased); err != nil || leased != 0 {
		t.Fatalf("leased rows = %d, %v", leased, err)
	}
}

func TestAStopReleasesTheUnreachedRows(t *testing.T) {
	w := newPostWorld(t)
	second := w.addRaid(t)
	w.enqueue(t, w.raidID)
	w.enqueue(t, second)
	ctx, cancel := context.WithCancel(context.Background())
	w.rest.OnCall = func(RESTCall) { cancel() }
	if err := w.drainer().DrainOnce(ctx); err != nil {
		t.Fatal(err)
	}
	if n := len(w.callsOf("CreateMessage")); n != 1 {
		t.Fatalf("creates = %d, want the row in flight only", n)
	}
	r1, _ := w.row(t, w.raidID)
	r2, _ := w.row(t, second)
	if (r1.messageID == nil) == (r2.messageID == nil) {
		t.Fatalf("rows = %+v, %+v, want exactly one sent", r1, r2)
	}
	w.rest.OnCall = nil
	w.drain(t)
	if n := len(w.callsOf("CreateMessage")); n != 2 {
		t.Fatalf("the released row was not taken at once: %d creates", n)
	}
}

func TestA429HoldsTheChannelAcrossTicks(t *testing.T) {
	w := newPostWorld(t)
	d := w.drainer()
	w.enqueue(t, w.raidID)
	w.rest.ErrFor = map[string]error{"CreateMessage": &APIError{Status: http.StatusTooManyRequests, RetryAfter: 10 * time.Second}}
	if err := d.DrainOnce(context.Background()); err != nil {
		t.Fatal(err)
	}
	w.rest.ErrFor = nil
	second := w.addRaid(t)
	w.enqueue(t, second)
	w.clock.advance(time.Second)
	if err := d.DrainOnce(context.Background()); err != nil {
		t.Fatal(err)
	}
	if n := len(w.rest.Calls()); n != 1 {
		t.Fatalf("calls in the held channel = %d, want 1", n)
	}
	if r, _ := w.row(t, second); r.attempts != 0 || !r.nextAttempt.After(w.clock.now()) {
		t.Fatalf("held row = %+v", r)
	}
	w.clock.advance(15 * time.Second)
	if err := d.DrainOnce(context.Background()); err != nil {
		t.Fatal(err)
	}
	if n := len(w.callsOf("CreateMessage")); n != 3 {
		t.Fatalf("creates after the hold = %d, want 3", n)
	}
}

func TestANonRetryable4xxParksUntilTheNextMark(t *testing.T) {
	w := newPostWorld(t)
	w.enqueue(t, w.raidID)
	w.rest.ErrFor = map[string]error{"CreateMessage": &APIError{Status: http.StatusBadRequest, Code: 50035}}
	w.drain(t)
	w.clock.advance(24 * time.Hour)
	w.drain(t)
	if n := len(w.rest.Calls()); n != 1 || !strings.Contains(w.log.String(), "post_parked") {
		t.Fatalf("calls = %d, log %s", n, w.log.String())
	}
	if !w.parked(t, w.raidID) {
		t.Fatal("row not parked")
	}
	// The raid changes (a sign-up): the mark unparks the row and the post goes out.
	w.rest.ErrFor = nil
	w.st.MarkRaidDirty(context.Background(), w.raidID)
	w.clock.advance(time.Second)
	w.drain(t)
	if r, _ := w.row(t, w.raidID); r.messageID == nil || r.attempts != 0 {
		t.Fatalf("the mark did not unpark the row: %+v", r)
	}
}

func TestAChannelSaveUnparksTheGuildsRows(t *testing.T) {
	w := newPostWorld(t)
	w.enqueue(t, w.raidID)
	w.rest.ErrFor = map[string]error{"CreateMessage": &APIError{Status: http.StatusRequestEntityTooLarge}}
	w.drain(t)
	w.rest.ErrFor = nil
	if err := w.st.SetChannel(context.Background(), w.gid, postChannel, "raid-signups"); err != nil {
		t.Fatal(err)
	}
	w.clock.advance(time.Second)
	w.drain(t)
	if r, _ := w.row(t, w.raidID); r.messageID == nil {
		t.Fatalf("a save did not unpark the row: %+v", r)
	}
}

func TestTransientFailuresNeverPark(t *testing.T) {
	for name, err := range map[string]error{
		"timeout":         ErrTimeout,
		"5xx":             &APIError{Status: http.StatusBadGateway},
		"unexplained 403": &APIError{Status: http.StatusForbidden, Code: codeMissingPerms},
		"other 404":       &APIError{Status: http.StatusNotFound, Code: codeUnknownWebhook},
	} {
		t.Run(name, func(t *testing.T) {
			w := newPostWorld(t)
			w.enqueue(t, w.raidID)
			w.rest.ErrFor = map[string]error{"CreateMessage": err}
			for range 10 {
				w.drain(t)
				w.clock.advance(backoffCap + time.Second)
			}
			if n := len(w.callsOf("CreateMessage")); n != 10 {
				t.Fatalf("creates = %d, want a try every round", n)
			}
			r, _ := w.row(t, w.raidID)
			if r.attempts != 10 || r.nextAttempt.After(w.clock.now()) || strings.Contains(w.log.String(), "post_parked") {
				t.Fatalf("row = %+v, now %v", r, w.clock.now())
			}
		})
	}
}

func TestA401HoldsEveryPostWithoutParking(t *testing.T) {
	w := newPostWorld(t)
	second := w.addRaid(t)
	w.enqueue(t, w.raidID)
	w.enqueueIn(t, second, "chan2")
	d := w.drainer()
	w.rest.ErrFor = map[string]error{"CreateMessage": &APIError{Status: http.StatusUnauthorized, Code: 0}}
	for range 3 {
		if err := d.DrainOnce(context.Background()); err != nil {
			t.Fatal(err)
		}
		w.clock.advance(time.Minute)
	}
	if n := len(w.rest.Calls()); n != 1 {
		t.Fatalf("calls under a 401 hold = %d, want 1", n)
	}
	if got := strings.Count(w.log.String(), "post_unauthorized"); got != 1 {
		t.Fatalf("401 log lines = %d, want 1", got)
	}
	for _, id := range []int64{w.raidID, second} {
		if r, _ := w.row(t, id); w.parked(t, id) || r.attempts != 0 {
			t.Fatalf("raid %d row = %+v", id, r)
		}
	}
	w.rest.ErrFor = nil
	w.clock.advance(authHold)
	if err := d.DrainOnce(context.Background()); err != nil {
		t.Fatal(err)
	}
	if n := len(w.callsOf("CreateMessage")); n != 3 {
		t.Fatalf("creates after the hold = %d, want 3", n)
	}
}

func TestEnqueueReadsTheInstallsChannel(t *testing.T) {
	w := newPostWorld(t)
	ctx := context.Background()
	if err := w.st.SetChannel(ctx, w.gid, "chan2", "general"); err != nil {
		t.Fatal(err)
	}
	w.enqueue(t, w.raidID)
	var channel string
	if err := w.st.Pool.QueryRow(ctx, `select channel_id from guild_raid_discord_posts where raid_id = $1`, w.raidID).Scan(&channel); err != nil || channel != "chan2" {
		t.Fatalf("channel = %q, %v", channel, err)
	}
	if err := w.st.MarkRemovedInDiscord(ctx, w.gid, "gone"); err != nil {
		t.Fatal(err)
	}
	other := w.addRaid(t)
	if queued, err := w.st.EnqueuePost(ctx, other, w.gid); err != nil || queued {
		t.Fatalf("enqueue without a usable channel = %v, %v", queued, err)
	}
	if _, found := w.row(t, other); found {
		t.Fatal("a row was queued without a usable channel")
	}
}

func TestChangingTheChannelRepointsOpenRaids(t *testing.T) {
	w := newPostWorld(t)
	ctx := context.Background()
	w.enqueue(t, w.raidID)
	w.drain(t)
	ended := w.addRaid(t)
	if _, err := w.st.Pool.Exec(ctx, `update guild_raids set starts_at = now() - interval '1 day' where id = $1`, ended); err != nil {
		t.Fatal(err)
	}
	w.enqueue(t, ended)
	was := w.clock.now().Truncate(time.Second)
	w.st.MarkRaidMoved(ctx, w.raidID, was)
	if err := w.st.SetChannel(ctx, w.gid, "chan2", "general"); err != nil {
		t.Fatal(err)
	}
	r, found := w.row(t, w.raidID)
	if !found || r.messageID != nil || r.movedFrom == nil || !r.movedFrom.Equal(was) || !r.dirty {
		t.Fatalf("open raid row = %+v, found %v (moved_from belongs to the raid)", r, found)
	}
	var channel string
	if err := w.st.Pool.QueryRow(ctx, `select channel_id from guild_raid_discord_posts where raid_id = $1`, w.raidID).Scan(&channel); err != nil || channel != "chan2" {
		t.Fatalf("channel = %q, %v", channel, err)
	}
	if _, found := w.row(t, ended); found {
		t.Fatal("the ended raid's old post was kept")
	}
	w.clock.advance(time.Second)
	w.drain(t)
	creates := w.callsOf("CreateMessage")
	if len(creates) != 2 || creates[1].ChannelID != "chan2" || creates[1].Message.Nonce == creates[0].Message.Nonce {
		t.Fatalf("creates = %+v", creates)
	}
}

func TestAMoveMarksMovedFrom(t *testing.T) {
	w := newPostWorld(t)
	ctx := context.Background()
	w.enqueue(t, w.raidID)
	was := w.clock.now().Add(time.Hour).Truncate(time.Second)
	before, _ := w.row(t, w.raidID)
	w.st.MarkRaidMoved(ctx, w.raidID, was)
	r, _ := w.row(t, w.raidID)
	if r.movedFrom == nil || !r.movedFrom.Equal(was) || r.dirtySeq != before.dirtySeq+1 {
		t.Fatalf("row = %+v", r)
	}
}
