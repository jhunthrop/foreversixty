// api/internal/guilds/raids_schedule_retry_test.go
package guilds

import (
	"context"
	"errors"
	"testing"
	"time"
)

// attemptsOf reads raid id's schedule attempts and whether it holds a lease.
func (w raidWorld) attemptsOf(t *testing.T, id int64) (attempts int, leased bool) {
	t.Helper()
	if err := w.h.pool.QueryRow(context.Background(),
		`select attempts, lease_until is not null from guild_raid_discord_schedule where raid_id = $1`, id).
		Scan(&attempts, &leased); err != nil {
		t.Fatal(err)
	}
	return attempts, leased
}

func TestARowLeasedByAnotherClaimIsNotSent(t *testing.T) {
	w, events, _ := postingWorld(t)
	ctx := context.Background()
	raid := w.createRaids(t, scheduleBody(startIn(10*day), 0, "true", "true"))[0]
	at := time.Now().Add(4 * day)
	mine := claimed(t, w.h.svc.Store, at)
	// This claim's lease runs out mid-batch and another runner claims the row.
	later := at.Add(scheduleLease + time.Minute)
	theirs := claimed(t, w.h.svc.Store, later)
	if len(mine) != 1 || len(theirs) != 1 || mine[0].RaidID != raid.ID {
		t.Fatalf("claims %+v and %+v", mine, theirs)
	}
	if got := w.h.svc.runDue(ctx, mine[0], at); got != outcomeSkipped || len(events.posted) != 0 {
		t.Fatalf("the stale claim sent: outcome %v, posted %d", got, len(events.posted))
	}
	if got := w.h.svc.runDue(ctx, theirs[0], later); got != outcomePosted || len(events.posted) != 1 {
		t.Fatalf("the live claim: outcome %v, posted %d", got, len(events.posted))
	}
}

func TestARaidLockedAfterTheClaimIsNotSent(t *testing.T) {
	w, events, _ := postingWorld(t)
	ctx := context.Background()
	raid := w.createRaids(t, scheduleBody(startIn(10*day), 0, "true", "true"))[0]
	at := time.Now().Add(4 * day)
	due := claimed(t, w.h.svc.Store, at)
	if len(due) != 1 {
		t.Fatalf("claimed %+v", due)
	}
	if _, err := w.h.pool.Exec(ctx, `update guild_raids set status = 'locked' where id = $1`, raid.ID); err != nil {
		t.Fatal(err)
	}
	if got := w.h.svc.runDue(ctx, due[0], at); got != outcomeSkipped || len(events.posted) != 0 {
		t.Fatalf("a locked raid: outcome %v, posted %d", got, len(events.posted))
	}
	if _, leased := w.attemptsOf(t, raid.ID); leased {
		t.Fatal("the skipped row kept its lease")
	}
}

func TestARefusedPostIsRetriedNotMarked(t *testing.T) {
	w, events, _ := postingWorld(t)
	raid := w.createRaids(t, scheduleBody(startIn(10*day), 0, "true", "true"))[0]
	at := time.Now().Add(4 * day)
	events.refuse(errors.New("refused"), nil)
	got, err := w.h.svc.RunRaidSchedule(context.Background(), at)
	if err == nil || got.Failed != 1 || got.Posted != 0 {
		t.Fatalf("refused run = %+v, %v", got, err)
	}
	if rs, _ := w.scheduleOf(t, raid.ID); rs.PostedAt != nil {
		t.Fatal("a refused post was marked posted")
	}
	if attempts, leased := w.attemptsOf(t, raid.ID); attempts != 1 || leased {
		t.Fatalf("attempts %d, leased %v", attempts, leased)
	}
	if got := w.run(t, at.Add(time.Minute)); got.Failed+got.Posted != 0 {
		t.Fatalf("retried inside the backoff: %+v", got)
	}
	events.refuse(nil, nil)
	if got := w.run(t, at.Add(scheduleRetryBase+time.Minute)); got.Posted != 1 || len(events.posted) != 1 {
		t.Fatalf("retry = %+v", got)
	}
}

func TestARefusedReminderIsRetriedNotMarked(t *testing.T) {
	w, events, _ := postingWorld(t)
	raid := w.createRaids(t, scheduleBody(startIn(5*day), 0, "true", "true"))[0]
	at := raid.StartsAt.Add(-23 * time.Hour)
	events.refuse(nil, errors.New("refused"))
	if got, err := w.h.svc.RunRaidSchedule(context.Background(), at); err == nil || got.Failed != 1 {
		t.Fatalf("refused run = %+v, %v", got, err)
	}
	if rs, _ := w.scheduleOf(t, raid.ID); rs.RemindedAt != nil {
		t.Fatal("a refused reminder was marked reminded")
	}
	events.refuse(nil, nil)
	if got := w.run(t, at.Add(scheduleRetryBase+time.Minute)); got.Reminded != 1 {
		t.Fatalf("retry = %+v", got)
	}
}

// leaseOf reads raid id's lease and last error.
func (w raidWorld) leaseOf(t *testing.T, id int64) (lease *time.Time, lastErr *string) {
	t.Helper()
	if err := w.h.pool.QueryRow(context.Background(),
		`select lease_until, last_error from guild_raid_discord_schedule where raid_id = $1`, id).
		Scan(&lease, &lastErr); err != nil {
		t.Fatal(err)
	}
	return lease, lastErr
}

func TestAReleaseAfterAnotherClaimLeavesItsLease(t *testing.T) {
	w, _, _ := postingWorld(t)
	ctx := context.Background()
	raid := w.createRaids(t, scheduleBody(startIn(10*day), 0, "true", "true"))[0]
	at := time.Now().Add(4 * day)
	mine := claimed(t, w.h.svc.Store, at)
	theirs := claimed(t, w.h.svc.Store, at.Add(scheduleLease+time.Minute))
	if len(mine) != 1 || len(theirs) != 1 {
		t.Fatalf("claims %+v and %+v", mine, theirs)
	}
	if ok, err := w.h.svc.Store.ReleaseSchedule(ctx, raid.ID, mine[0].LeaseUntil); ok || err != nil {
		t.Fatalf("stale release = %v, %v", ok, err)
	}
	w.h.svc.releaseDue(ctx, raid.ID, mine[0].LeaseUntil)
	w.h.svc.failDue(ctx, raid.ID, 0, mine[0].LeaseUntil, at, errors.New("boom"))
	lease, _ := w.leaseOf(t, raid.ID)
	if lease == nil || !lease.Equal(theirs[0].LeaseUntil) {
		t.Fatalf("lease = %v, want the other runner's %v", lease, theirs[0].LeaseUntil)
	}
	if attempts, _ := w.attemptsOf(t, raid.ID); attempts != 0 {
		t.Fatalf("a stale fail counted: attempts %d", attempts)
	}
}

// refuseMarks makes every write that marks a schedule row posted fail, until cleared.
func refuseMarks(t *testing.T, w raidWorld) (clear func()) {
	t.Helper()
	ctx := context.Background()
	for _, stmt := range []string{
		`create or replace function t12_refuse_mark() returns trigger language plpgsql as $$
		 begin
		   if new.posted_at is not null and old.posted_at is null then raise exception 'mark refused'; end if;
		   return new;
		 end $$`,
		`create trigger t12_refuse_mark before update on guild_raid_discord_schedule
		 for each row execute function t12_refuse_mark()`,
	} {
		if _, err := w.h.pool.Exec(ctx, stmt); err != nil {
			t.Fatal(err)
		}
	}
	done := false
	clear = func() {
		if done {
			return
		}
		done = true
		if _, err := w.h.pool.Exec(ctx, `drop trigger if exists t12_refuse_mark on guild_raid_discord_schedule;
			drop function if exists t12_refuse_mark()`); err != nil {
			t.Error(err)
		}
	}
	t.Cleanup(clear)
	return clear
}

func TestAFailedMarkHoldsTheSentPost(t *testing.T) {
	w, events, _ := postingWorld(t)
	raid := w.createRaids(t, scheduleBody(startIn(10*day), 0, "true", "true"))[0]
	at := nanoClock().Add(4 * day)
	clear := refuseMarks(t, w)
	if got := w.run(t, at); got.Posted != 1 || len(events.posted) != 1 {
		t.Fatalf("run = %+v, posted %d", got, len(events.posted))
	}
	rs, _ := w.scheduleOf(t, raid.ID)
	lease, lastErr := w.leaseOf(t, raid.ID)
	if rs.PostedAt != nil || lease == nil || !lease.Equal(raid.StartsAt) || lastErr == nil || *lastErr != scheduleSentUnmarked {
		t.Fatalf("held row: posted %v, lease %v, last error %v", rs.PostedAt, lease, lastErr)
	}
	// The lease has long expired by now, but the held row is never claimed or posted again.
	w.run(t, at.Add(2*day))
	if len(events.posted) != 1 {
		t.Fatalf("re-posted: %d", len(events.posted))
	}
	clear()
	recoverAt := at.Add(2*day + time.Hour)
	w.run(t, recoverAt)
	rs, _ = w.scheduleOf(t, raid.ID)
	if rs.PostedAt == nil || !rs.PostedAt.Equal(stored(recoverAt)) || len(events.posted) != 1 {
		t.Fatalf("recovered row: posted %v, posts %d", rs.PostedAt, len(events.posted))
	}
}

// A save carrying one box keeps the install's other default, even with the bot inactive.
func TestAOneBoxSaveKeepsTheOtherDefault(t *testing.T) {
	w, _, _ := postingWorld(t)
	w.h.svc.Discord = nil
	writes := &[][2]bool{}
	w.h.svc.Bot = fakeBot{on: true, postDefault: true, remindDefault: true, defaults: writes}
	w.createRaids(t, scheduleBody(startIn(3*day), 0, "", "false"))
	if len(*writes) != 1 || (*writes)[0] != [2]bool{true, false} {
		t.Fatalf("install default writes = %v", *writes)
	}
}
