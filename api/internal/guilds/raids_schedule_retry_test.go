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
