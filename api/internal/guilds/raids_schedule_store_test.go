package guilds

import (
	"context"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

type scheduleWorld struct {
	st  *Store
	gid int64
	uid int64
}

func newScheduleWorld(t *testing.T) scheduleWorld {
	t.Helper()
	pool := testPool(t)
	return scheduleWorld{
		st:  &Store{Pool: pool},
		gid: seedGuild(t, pool, "Sched"),
		uid: seedUser(t, pool, "sched@example.com"),
	}
}

func (w scheduleWorld) raid(t *testing.T, startsAt time.Time, status string) int64 {
	t.Helper()
	var id int64
	if err := w.st.Pool.QueryRow(context.Background(), `
		insert into guild_raids (guild_id, zone, title, starts_at, duration_minutes, status, created_by)
		values ($1, 'raid:onyxias-lair', 'Ony', $2, 180, $3, $4) returning id`,
		w.gid, startsAt, status, w.uid).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

func claimed(t *testing.T, st *Store, now time.Time) []DueSchedule {
	t.Helper()
	due, err := st.ClaimDueSchedules(context.Background(), now, scheduleBatch)
	if err != nil {
		t.Fatal(err)
	}
	return due
}

func TestPostAtFor(t *testing.T) {
	now := time.Date(2026, 11, 1, 12, 0, 0, 0, time.UTC)
	if got := postAtFor(now.Add(3*24*time.Hour), now); !got.Equal(now) {
		t.Fatalf("near start posts now, got %v", got)
	}
	far := now.Add(10 * 24 * time.Hour)
	if got := postAtFor(far, now); !got.Equal(far.Add(-postHorizon)) {
		t.Fatalf("far start posts 7 days before, got %v", got)
	}
}

func TestScheduleDueWindows(t *testing.T) {
	ctx := context.Background()
	w := newScheduleWorld(t)
	now := time.Now().UTC().Add(5 * time.Second)

	unposted := w.raid(t, now.Add(48*time.Hour), "planned")
	cancelled := w.raid(t, now.Add(48*time.Hour), "cancelled")
	locked := w.raid(t, now.Add(48*time.Hour), "locked")
	for _, id := range []int64{unposted, cancelled, locked} {
		if err := w.st.UpsertSchedule(ctx, id, w.gid, now.Add(-time.Minute), true); err != nil {
			t.Fatal(err)
		}
	}
	due := claimed(t, w.st, now)
	if len(due) != 1 || due[0].RaidID != unposted || due[0].Kind != duePost {
		t.Fatalf("due = %+v, want one post for raid %d", due, unposted)
	}
	if again := claimed(t, w.st, now); len(again) != 0 {
		t.Fatalf("a leased row was claimed again: %+v", again)
	}

	// A raid posted 3 days ahead is due its reminder 23 h before the start.
	start := now.Add(23 * time.Hour)
	reminded := w.raid(t, start, "planned")
	quiet := w.raid(t, start, "planned")
	fresh := w.raid(t, start, "planned")
	for _, c := range []struct {
		id       int64
		remind   bool
		postedAt time.Time
	}{
		{reminded, true, start.Add(-72 * time.Hour)},
		{quiet, false, start.Add(-72 * time.Hour)},
		{fresh, true, start.Add(-30 * time.Hour)},
	} {
		if err := w.st.UpsertSchedule(ctx, c.id, w.gid, c.postedAt, c.remind); err != nil {
			t.Fatal(err)
		}
		if err := w.st.MarkSchedulePosted(ctx, c.id, c.postedAt); err != nil {
			t.Fatal(err)
		}
	}
	due = claimed(t, w.st, now)
	if len(due) != 1 || due[0].RaidID != reminded || due[0].Kind != dueReminder {
		t.Fatalf("due = %+v, want one reminder for raid %d", due, reminded)
	}
}

func TestTwoClaimsSplitRows(t *testing.T) {
	ctx := context.Background()
	w := newScheduleWorld(t)
	now := time.Now().UTC().Add(5 * time.Second)
	for range 4 {
		if err := w.st.UpsertSchedule(ctx, w.raid(t, now.Add(24*time.Hour), "planned"), w.gid, now.Add(-time.Minute), false); err != nil {
			t.Fatal(err)
		}
	}
	first, second := make(chan []DueSchedule, 1), make(chan []DueSchedule, 1)
	for _, out := range []chan []DueSchedule{first, second} {
		go func() {
			due, err := w.st.ClaimDueSchedules(ctx, now, scheduleBatch)
			if err != nil {
				t.Error(err)
			}
			out <- due
		}()
	}
	a, b := <-first, <-second
	if len(a)+len(b) != 4 {
		t.Fatalf("claims took %d + %d rows, want 4 in all with none twice", len(a), len(b))
	}
	seen := map[int64]bool{}
	for _, d := range append(a, b...) {
		if seen[d.RaidID] {
			t.Fatalf("raid %d claimed twice", d.RaidID)
		}
		seen[d.RaidID] = true
	}
}

// A schedule time reads back at microseconds, the rest truncated: a test comparing a
// written time with a read one compares stored(written). nanoClock carries the
// sub-microsecond part on darwin too, so such a comparison fails locally as it does in CI.
func TestScheduleTimesReadBackAtMicroseconds(t *testing.T) {
	ctx := context.Background()
	w := newScheduleWorld(t)
	now := nanoClock()
	if now.Sub(stored(now)) != subMicro {
		t.Fatalf("nanoClock %v carries no sub-microsecond part", now)
	}
	id := w.raid(t, now.Add(5*24*time.Hour), "planned")
	if err := w.st.UpsertSchedule(ctx, id, w.gid, now, false); err != nil {
		t.Fatal(err)
	}
	got, err := w.st.SchedulesOf(ctx, []int64{id})
	if err != nil {
		t.Fatal(err)
	}
	if at := got[id].PostAt; !at.Equal(stored(now)) || at.Equal(now) {
		t.Fatalf("post_at read back %v, want %v (written %v)", at, stored(now), now)
	}
}

func TestUpsertKeepsAPostedRow(t *testing.T) {
	ctx := context.Background()
	w := newScheduleWorld(t)
	now := nanoClock().Add(5 * time.Second)
	id := w.raid(t, now.Add(5*24*time.Hour), "planned")
	first := now.Add(-time.Hour)
	if err := w.st.UpsertSchedule(ctx, id, w.gid, first, false); err != nil {
		t.Fatal(err)
	}
	if err := w.st.MarkSchedulePosted(ctx, id, first); err != nil {
		t.Fatal(err)
	}
	if err := w.st.UpsertSchedule(ctx, id, w.gid, now.Add(48*time.Hour), true); err != nil {
		t.Fatal(err)
	}
	got, err := w.st.SchedulesOf(ctx, []int64{id})
	if err != nil {
		t.Fatal(err)
	}
	rs := got[id]
	if !rs.PostAt.Equal(stored(first)) || rs.PostedAt == nil || !rs.PostedAt.Equal(stored(first)) || !rs.Remind {
		t.Fatalf("schedule = %+v", rs)
	}
	if err := w.st.DeleteUnpostedSchedule(ctx, id); err != nil {
		t.Fatal(err)
	}
	if got, _ := w.st.SchedulesOf(ctx, []int64{id}); len(got) != 1 {
		t.Fatal("a posted row was deleted")
	}
}

// A row sent but not yet marked posted survives an untick: deleting it would let a
// later re-tick create a fresh row and post the raid again.
func TestDeleteKeepsASentUnmarkedRow(t *testing.T) {
	ctx := context.Background()
	w := newScheduleWorld(t)
	now := time.Now().UTC().Add(5 * time.Second)
	id := w.raid(t, now.Add(30*time.Hour), "planned")
	if err := w.st.UpsertSchedule(ctx, id, w.gid, now, true); err != nil {
		t.Fatal(err)
	}
	if _, err := w.st.Pool.Exec(ctx,
		`update guild_raid_discord_schedule set last_error = $2 where raid_id = $1`, id, scheduleSentUnmarked); err != nil {
		t.Fatal(err)
	}
	if err := w.st.DeleteUnpostedSchedule(ctx, id); err != nil {
		t.Fatal(err)
	}
	if got, _ := w.st.SchedulesOf(ctx, []int64{id}); len(got) != 1 {
		t.Fatal("a sent, unmarked row was deleted")
	}
}

func TestResetReturnsTheOldReminder(t *testing.T) {
	ctx := context.Background()
	w := newScheduleWorld(t)
	now := time.Now().UTC().Add(5 * time.Second)
	id := w.raid(t, now.Add(30*time.Hour), "planned")
	if err := w.st.UpsertSchedule(ctx, id, w.gid, now, true); err != nil {
		t.Fatal(err)
	}
	ref := notify.ReminderRef{Via: notify.ReminderViaWebhook, MessageID: "m1"}
	if err := w.st.MarkScheduleReminded(ctx, id, ref, now); err != nil {
		t.Fatal(err)
	}
	got, err := w.st.ResetScheduleReminder(ctx, id)
	if err != nil || got == nil || *got != ref {
		t.Fatalf("first reset = %v, %v", got, err)
	}
	if got, err := w.st.ResetScheduleReminder(ctx, id); err != nil || got != nil {
		t.Fatalf("second reset = %v, %v", got, err)
	}
}

func TestFailAndReleaseSchedule(t *testing.T) {
	ctx := context.Background()
	w := newScheduleWorld(t)
	now := time.Now().UTC().Add(5 * time.Second)
	id := w.raid(t, now.Add(24*time.Hour), "planned")
	if err := w.st.UpsertSchedule(ctx, id, w.gid, now.Add(-time.Minute), false); err != nil {
		t.Fatal(err)
	}
	first := claimed(t, w.st, now)
	if len(first) != 1 {
		t.Fatal("expected a claim")
	}
	if ok, err := w.st.FailSchedule(ctx, id, first[0].LeaseUntil, now.Add(time.Hour), "boom"); err != nil || !ok {
		t.Fatalf("fail = %v, %v", ok, err)
	}
	if len(claimed(t, w.st, now)) != 0 {
		t.Fatal("a failed row was due before its retry time")
	}
	due := claimed(t, w.st, now.Add(2*time.Hour))
	if len(due) != 1 || due[0].Attempts != 1 {
		t.Fatalf("after retry time: %+v", due)
	}
	if ok, err := w.st.ReleaseSchedule(ctx, id, due[0].LeaseUntil); err != nil || !ok {
		t.Fatalf("release = %v, %v", ok, err)
	}
	if len(claimed(t, w.st, now.Add(2*time.Hour))) != 1 {
		t.Fatal("a released row was not claimable")
	}
}

func TestRemindFlippedOffAndOnIsStillClaimed(t *testing.T) {
	ctx := context.Background()
	w := newScheduleWorld(t)
	now := time.Now().UTC().Add(5 * time.Second)
	start := now.Add(23 * time.Hour)
	id := w.raid(t, start, "planned")
	postedAt := start.Add(-72 * time.Hour)
	if err := w.st.UpsertSchedule(ctx, id, w.gid, postedAt, true); err != nil {
		t.Fatal(err)
	}
	if err := w.st.MarkSchedulePosted(ctx, id, postedAt); err != nil {
		t.Fatal(err)
	}
	if err := w.st.SetScheduleRemind(ctx, id, false); err != nil {
		t.Fatal(err)
	}
	if len(claimed(t, w.st, now)) != 0 {
		t.Fatal("a row with remind off was claimed")
	}
	if err := w.st.SetScheduleRemind(ctx, id, true); err != nil {
		t.Fatal(err)
	}
	due := claimed(t, w.st, now)
	if len(due) != 1 || due[0].Kind != dueReminder {
		t.Fatalf("after re-enabling: %+v", due)
	}
	if err := w.st.UpsertSchedule(ctx, id, w.gid, postedAt, false); err != nil {
		t.Fatal(err)
	}
	if err := w.st.UpsertSchedule(ctx, id, w.gid, postedAt, true); err != nil {
		t.Fatal(err)
	}
	if ok, err := w.st.ReleaseSchedule(ctx, id, due[0].LeaseUntil); err != nil || !ok {
		t.Fatalf("release = %v, %v", ok, err)
	}
	if len(claimed(t, w.st, now)) != 1 {
		t.Fatal("an upsert flip lost the reminder")
	}
}
