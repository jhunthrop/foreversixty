package discordbot

import (
	"context"
	"encoding/json"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// reminderRow is a post row's reminder columns as the tests read them.
type reminderRow struct {
	due        bool
	messageID  *string
	remindedAt *time.Time
	wake       *time.Time
}

func (w *postWorld) reminderRow(t *testing.T) reminderRow {
	t.Helper()
	var r reminderRow
	if err := w.st.Pool.QueryRow(context.Background(), `select remind_due, reminder_message_id, reminded_at, wake_at
		from guild_raid_discord_posts where raid_id = $1`, w.raidID).Scan(&r.due, &r.messageID, &r.remindedAt, &r.wake); err != nil {
		t.Fatal(err)
	}
	return r
}

// posted creates the raid's post (message m1) and lets the debounce pass.
func (w *postWorld) postedOnce(t *testing.T) {
	t.Helper()
	w.enqueue(t, w.raidID)
	w.drain(t)
	w.clock.advance(3 * time.Second)
}

// markAndDrain runs mark, lets the debounce pass and drains.
func (w *postWorld) markAndDrain(t *testing.T, mark func(ctx context.Context)) {
	t.Helper()
	mark(context.Background())
	w.clock.advance(3 * time.Second)
	w.drain(t)
}

func (w *postWorld) remindDue(t *testing.T) {
	t.Helper()
	w.markAndDrain(t, func(ctx context.Context) {
		if err := w.st.MarkReminderDue(ctx, w.raidID); err != nil {
			t.Fatal(err)
		}
	})
}

func (w *postWorld) lastEditTitle(t *testing.T) string {
	t.Helper()
	edits := w.callsOf("EditMessage")
	if len(edits) == 0 {
		t.Fatal("no edit")
	}
	return edits[len(edits)-1].Message.Embeds[0].Title
}

func (w *postWorld) replies() []RESTCall {
	var out []RESTCall
	for _, c := range w.callsOf("CreateMessage") {
		if c.Message.ReplyTo != "" {
			out = append(out, c)
		}
	}
	return out
}

func TestReminderEditsTheNewestPost(t *testing.T) {
	w := newPostWorld(t)
	w.postedOnce(t)
	w.rest.Latest = "m1"
	w.remindDue(t)
	edits := w.callsOf("EditMessage")
	if len(edits) != 1 || edits[0].MessageID != "m1" || !strings.HasPrefix(edits[0].Message.Embeds[0].Title, "Tomorrow: ") {
		t.Fatalf("edits = %+v", edits)
	}
	if n := len(w.callsOf("CreateMessage")); n != 1 {
		t.Fatalf("creates = %d, want only the post", n)
	}
	if latest := w.callsOf("LatestMessageID"); len(latest) != 1 || latest[0].ChannelID != postChannel {
		t.Fatalf("latest reads = %+v", latest)
	}
	r := w.reminderRow(t)
	if r.due || r.remindedAt == nil || r.messageID != nil {
		t.Fatalf("row = %+v", r)
	}
	if starts := w.facts.by[w.raidID].Post.StartsAt; r.wake == nil || !r.wake.Equal(starts) {
		t.Fatalf("wake = %v, want the start %v", r.wake, starts)
	}
}

func TestReminderRepliesWhenTheChannelMovedOn(t *testing.T) {
	w := newPostWorld(t)
	w.postedOnce(t)
	w.rest.Latest, w.rest.ReplyID = "other", "r1"
	w.remindDue(t)
	if !strings.HasPrefix(w.lastEditTitle(t), "Tomorrow: ") {
		t.Fatalf("post title = %q", w.lastEditTitle(t))
	}
	replies := w.replies()
	if len(replies) != 1 {
		t.Fatalf("replies = %+v", replies)
	}
	m := replies[0].Message
	if want := notify.ReminderLine(w.facts.by[w.raidID].Post.RaidPosted); m.Content != want || m.ReplyTo != "m1" || len(m.Embeds) != 0 {
		t.Fatalf("reply = %+v, want content %q", m, want)
	}
	if len(m.Components) != 1 || len(m.Components[0]) != 5 {
		t.Fatalf("reply components = %+v", m.Components)
	}
	for _, c := range m.Components[0][:4] {
		if !strings.HasPrefix(c.CustomID, "raid:") || c.Disabled {
			t.Fatalf("reply button = %+v", c)
		}
	}
	raw, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`"allowed_mentions":{"parse":[]}`, `"message_reference":{"message_id":"m1","fail_if_not_exists":false}`} {
		if !strings.Contains(string(raw), want) {
			t.Fatalf("reply json %s lacks %s", raw, want)
		}
	}
	if r := w.reminderRow(t); r.due || r.remindedAt == nil || r.messageID == nil || *r.messageID != "r1" {
		t.Fatalf("row = %+v", r)
	}
}

func TestTomorrowStaysOnLaterEdits(t *testing.T) {
	w := newPostWorld(t)
	w.postedOnce(t)
	w.rest.Latest = "m1"
	w.remindDue(t)
	w.facts.edit(w.raidID, func(f *guilds.RaidPostFacts) { f.Post.Counts.In = 2 })
	w.markAndDrain(t, func(ctx context.Context) { w.st.MarkRaidDirty(ctx, w.raidID) })
	if title := w.lastEditTitle(t); !strings.HasPrefix(title, "Tomorrow: ") {
		t.Fatalf("a later edit dropped Tomorrow: %q", title)
	}
	starts := w.facts.by[w.raidID].Post.StartsAt
	if r := w.reminderRow(t); r.wake == nil || !r.wake.Equal(starts) {
		t.Fatalf("wake = %v, want %v", r.wake, starts)
	}
	w.clock.advance(starts.Sub(w.clock.now()) + time.Minute)
	w.drain(t)
	if title := w.lastEditTitle(t); strings.HasPrefix(title, "Tomorrow: ") || len(w.callsOf("EditMessage")) != 3 {
		t.Fatalf("at the start the title = %q after %d edits", title, len(w.callsOf("EditMessage")))
	}
	if ends := w.facts.by[w.raidID].EndsAt; !w.reminderRow(t).wake.Equal(ends) {
		t.Fatalf("after the start the post wakes at %v, want the end", w.reminderRow(t).wake)
	}
}

// replied is a world whose reminder went out as reply r1 under post m1.
func replied(t *testing.T) *postWorld {
	t.Helper()
	w := newPostWorld(t)
	w.postedOnce(t)
	w.rest.Latest, w.rest.ReplyID = "other", "r1"
	w.remindDue(t)
	return w
}

func (w *postWorld) moveBy(t *testing.T, d time.Duration) {
	t.Helper()
	was := w.facts.by[w.raidID].Post.StartsAt
	w.facts.edit(w.raidID, func(f *guilds.RaidPostFacts) {
		f.Post.StartsAt = f.Post.StartsAt.Add(d)
		f.EndsAt = f.EndsAt.Add(d)
	})
	w.markAndDrain(t, func(ctx context.Context) { w.st.MarkRaidMoved(ctx, w.raidID, was) })
}

func (w *postWorld) setStatus(t *testing.T, status string) {
	t.Helper()
	w.facts.edit(w.raidID, func(f *guilds.RaidPostFacts) { f.Status = status })
	w.markAndDrain(t, func(ctx context.Context) { w.st.MarkRaidDirty(ctx, w.raidID) })
}

func TestDrainerWithdrawsTheReminder(t *testing.T) {
	cases := []struct {
		name   string
		change func(t *testing.T, w *postWorld)
		title  string
		rearm  bool
	}{
		{"cancel", func(t *testing.T, w *postWorld) { w.setStatus(t, guilds.RaidCancelled) }, "Cancelled: ", false},
		{"lock", func(t *testing.T, w *postWorld) { w.setStatus(t, guilds.RaidLocked) }, "", false},
		{"move", func(t *testing.T, w *postWorld) { w.moveBy(t, 48*time.Hour) }, "Moved: ", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := replied(t)
			tc.change(t, w)
			deletes := w.callsOf("DeleteMessage")
			if len(deletes) != 1 || deletes[0].ChannelID != postChannel || deletes[0].MessageID != "r1" {
				t.Fatalf("deletes = %+v", deletes)
			}
			if title := w.lastEditTitle(t); !strings.HasPrefix(title, tc.title) || strings.HasPrefix(title, "Tomorrow: ") {
				t.Fatalf("post title = %q, want prefix %q", title, tc.title)
			}
			r := w.reminderRow(t)
			if r.messageID != nil || (r.remindedAt == nil) != tc.rearm {
				t.Fatalf("row = %+v, rearm %v", r, tc.rearm)
			}
			w.markAndDrain(t, func(ctx context.Context) { w.st.MarkRaidDirty(ctx, w.raidID) })
			if n := len(w.callsOf("DeleteMessage")); n != 1 {
				t.Fatalf("deletes after a later edit = %d", n)
			}
		})
	}
	t.Run("already deleted in Discord", func(t *testing.T) {
		w := replied(t)
		w.rest.ErrFor = map[string]error{"DeleteMessage": &APIError{Status: 404, Code: codeUnknownMessage}}
		w.moveBy(t, 48*time.Hour)
		if r := w.reminderRow(t); r.messageID != nil || r.remindedAt != nil {
			t.Fatalf("row = %+v", r)
		}
		if r, _ := w.row(t, w.raidID); r.gone {
			t.Fatal("a reminder already deleted marked the post gone")
		}
	})
	t.Run("an edited post loses Tomorrow on a move", func(t *testing.T) {
		w := newPostWorld(t)
		w.postedOnce(t)
		w.rest.Latest = "m1"
		w.remindDue(t)
		w.moveBy(t, 48*time.Hour)
		if n := len(w.callsOf("DeleteMessage")); n != 0 {
			t.Fatalf("deletes = %d with no reply", n)
		}
		if r := w.reminderRow(t); r.remindedAt != nil {
			t.Fatalf("reminded_at kept after a move: %+v", r)
		}
		if title := w.lastEditTitle(t); !strings.HasPrefix(title, "Moved: ") {
			t.Fatalf("title = %q", title)
		}
	})
}

func TestAMovedRaidStillGetsItsReminder(t *testing.T) {
	w := replied(t)
	w.moveBy(t, 48*time.Hour)
	w.rest.Latest, w.rest.ReplyID = "m1", "r2"
	w.remindDue(t)
	if title := w.lastEditTitle(t); !strings.HasPrefix(title, "Moved: ") {
		t.Fatalf("moved beats Tomorrow, title = %q", title)
	}
	replies := w.replies()
	if len(replies) != 2 || replies[1].Message.ReplyTo != "m1" {
		t.Fatalf("replies = %+v", replies)
	}
	if r := w.reminderRow(t); r.messageID == nil || *r.messageID != "r2" || r.remindedAt == nil {
		t.Fatalf("row = %+v", r)
	}
	w.markAndDrain(t, func(ctx context.Context) { w.st.MarkRaidDirty(ctx, w.raidID) })
	if n := len(w.callsOf("DeleteMessage")); n != 1 {
		t.Fatalf("the re-armed reminder was withdrawn: deletes = %d", n)
	}
	w.moveBy(t, time.Hour)
	if deletes := w.callsOf("DeleteMessage"); len(deletes) != 2 || deletes[1].MessageID != "r2" {
		t.Fatalf("a second move kept the reminder: %+v", deletes)
	}
}

func TestAReminderOnAClosedRaidIsDroppedUnsent(t *testing.T) {
	w := newPostWorld(t)
	w.postedOnce(t)
	w.facts.edit(w.raidID, func(f *guilds.RaidPostFacts) { f.Status = guilds.RaidLocked })
	w.rest.Latest = "other"
	w.remindDue(t)
	if len(w.replies()) != 0 || len(w.callsOf("LatestMessageID")) != 0 {
		t.Fatalf("a locked raid was reminded: %+v", w.rest.Calls())
	}
	if r := w.reminderRow(t); r.due || r.messageID != nil {
		t.Fatalf("row = %+v", r)
	}
}

func TestReminderDueIsDrainedOnce(t *testing.T) {
	w := newPostWorld(t)
	w.postedOnce(t)
	w.rest.Latest = "other"
	if err := w.st.MarkReminderDue(context.Background(), w.raidID); err != nil {
		t.Fatal(err)
	}
	w.clock.advance(3 * time.Second)
	w.rest.OnCall = func(RESTCall) { time.Sleep(50 * time.Millisecond) }
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
	w.clock.advance(time.Minute)
	w.drain(t)
	if n := len(w.replies()); n != 1 {
		t.Fatalf("replies = %d, want exactly 1", n)
	}
}

func TestAMoveWhileTheReminderIsInFlightWithdrawsIt(t *testing.T) {
	w := newPostWorld(t)
	w.postedOnce(t)
	w.rest.Latest, w.rest.ReplyID = "other", "r1"
	var once sync.Once
	w.rest.OnCall = func(c RESTCall) {
		if c.Method != "CreateMessage" || c.Message.ReplyTo == "" {
			return
		}
		once.Do(func() {
			was := w.facts.by[w.raidID].Post.StartsAt
			w.facts.edit(w.raidID, func(f *guilds.RaidPostFacts) {
				f.Post.StartsAt = f.Post.StartsAt.Add(48 * time.Hour)
				f.EndsAt = f.EndsAt.Add(48 * time.Hour)
			})
			w.st.MarkRaidMoved(context.Background(), w.raidID, was)
		})
	}
	w.remindDue(t)
	if r := w.reminderRow(t); r.remindedAt != nil || r.messageID == nil {
		t.Fatalf("row after an in-flight move = %+v", r)
	}
	w.clock.advance(3 * time.Second)
	w.drain(t)
	if d := w.callsOf("DeleteMessage"); len(d) != 1 || d[0].MessageID != "r1" {
		t.Fatalf("deletes = %+v", d)
	}
}

func TestTheReminderReplyCarriesANonce(t *testing.T) {
	w := replied(t)
	r := w.replies()[0].Message
	if r.Nonce == "" || r.Nonce == w.callsOf("CreateMessage")[0].Message.Nonce || len(r.Nonce) > maxNonceRunes {
		t.Fatalf("reply nonce = %q", r.Nonce)
	}
}

func TestRouterSendsTheReminderToTheBotOrTheWebhook(t *testing.T) {
	ctx := context.Background()
	t.Run("bot", func(t *testing.T) {
		w := newPostWorld(t)
		w.enqueue(t, w.raidID)
		ref, sent, err := w.router.RaidReminder(ctx, w.gid, posted(w, w.raidID))
		if err != nil || !sent || ref.Via != notify.ReminderViaBot {
			t.Fatalf("RaidReminder = %+v, %v, %v", ref, sent, err)
		}
		if r := w.reminderRow(t); !r.due {
			t.Fatalf("row = %+v", r)
		}
		if len(w.poster.Contents()) != 0 {
			t.Fatalf("webhook reminded: %+v", w.poster.Contents())
		}
		w.router.RaidReminderWithdrawn(ctx, w.gid, ref)
		if len(w.poster.MessageDeletes()) != 0 || len(w.rest.Calls()) != 0 {
			t.Fatal("a bot reference was withdrawn by the router")
		}
	})
	webhook := map[string]func(t *testing.T, w *postWorld){
		"no row": func(*testing.T, *postWorld) {},
		"gone row": func(t *testing.T, w *postWorld) {
			w.enqueue(t, w.raidID)
			if _, err := w.st.Pool.Exec(ctx, `update guild_raid_discord_posts set gone = true where raid_id = $1`, w.raidID); err != nil {
				t.Fatal(err)
			}
		},
		"no bot": func(t *testing.T, w *postWorld) {
			w.enqueue(t, w.raidID)
			if err := w.st.RemoveBot(ctx, w.gid); err != nil {
				t.Fatal(err)
			}
		},
	}
	for name, setup := range webhook {
		t.Run(name, func(t *testing.T) {
			w := newPostWorld(t)
			setup(t, w)
			w.poster.ContentID = "555"
			ref, sent, err := w.router.RaidReminder(ctx, w.gid, posted(w, w.raidID))
			if err != nil || !sent || ref.Via != notify.ReminderViaWebhook || len(w.poster.Contents()) != 1 {
				t.Fatalf("RaidReminder = %+v, %v, %v; contents %+v", ref, sent, err, w.poster.Contents())
			}
			w.router.RaidReminderWithdrawn(ctx, w.gid, ref)
			if d := w.poster.MessageDeletes(); len(d) != 1 || d[0].MessageID != "555" {
				t.Fatalf("webhook deletes = %+v", d)
			}
		})
	}
}

func TestTomorrowKeepsTheSignUpButtonsAndMovedOutranksIt(t *testing.T) {
	post := notify.RaidPost{RaidPosted: notify.RaidPosted{RaidID: 7}, State: notify.PostTomorrow}
	for _, c := range postRows(post, "https://site.test/r", "")[0] {
		if c.Disabled {
			t.Fatalf("Tomorrow button %q is disabled", c.Label)
		}
	}
	now := time.Now()
	start := now.Add(20 * time.Hour)
	reminded := Post{RemindedAt: &now}
	if got := lookOf(notify.PostMoved, reminded, start, now); got != notify.PostMoved {
		t.Fatalf("moved reminded look = %q", got)
	}
	if got := lookOf(notify.PostOpen, reminded, start, now); got != notify.PostTomorrow {
		t.Fatalf("open reminded look = %q", got)
	}
	if got := lookOf(notify.PostOpen, reminded, start, start); got != notify.PostOpen {
		t.Fatalf("at the start look = %q", got)
	}
	if got := lookOf(notify.PostOpen, Post{}, start, now); got != notify.PostOpen {
		t.Fatalf("unreminded look = %q", got)
	}
}

func TestRESTLatestMessageAndDelete(t *testing.T) {
	ctx := context.Background()
	rest, got := serve(t, 200, `[{"id":"m9"}]`)
	if id, err := rest.LatestMessageID(ctx, "c1"); err != nil || id != "m9" || got.method != "GET" ||
		got.path != "/channels/c1/messages" || got.auth != "Bot t0k" {
		t.Fatalf("latest = %q, %v, seen %+v", id, err, got)
	}
	empty, _ := serve(t, 200, `[]`)
	if id, err := empty.LatestMessageID(ctx, "c1"); err != nil || id != "" {
		t.Fatalf("empty latest = %q, %v", id, err)
	}
	del, seenDel := serve(t, 204, ``)
	if err := del.DeleteMessage(ctx, "c1", "r1"); err != nil || seenDel.method != "DELETE" || seenDel.path != "/channels/c1/messages/r1" {
		t.Fatalf("delete = %v, seen %+v", err, seenDel)
	}
	gone, _ := serve(t, 404, `{"code":10008,"message":"Unknown Message"}`)
	if err := gone.DeleteMessage(ctx, "c1", "r1"); err != nil {
		t.Fatalf("10008 on delete = %v, want success", err)
	}
	refused, _ := serve(t, 403, `{"code":50013,"message":"Missing Permissions"}`)
	if err := refused.DeleteMessage(ctx, "c1", "r1"); err == nil {
		t.Fatal("a refused delete answered success")
	}
}
