// api/internal/guilds/raids_schedule_giveup_test.go
package guilds

import (
	"context"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// nextAttemptInfinite reads whether raid id's row was given up (never due again).
func (w raidWorld) nextAttemptInfinite(t *testing.T, id int64) bool {
	t.Helper()
	var infinite bool
	if err := w.h.pool.QueryRow(context.Background(),
		`select next_attempt_at = 'infinity' from guild_raid_discord_schedule where raid_id = $1`, id).
		Scan(&infinite); err != nil {
		t.Fatal(err)
	}
	return infinite
}

// A due post that reached no channel is held back and retried, never marked or counted posted.
func TestAPostThatReachedNoChannelIsRetriedNotMarked(t *testing.T) {
	w, events, _ := postingWorld(t)
	raid := w.createRaids(t, scheduleBody(startIn(10*day), 0, "true", "true"))[0]
	at := time.Now().Add(4 * day)
	events.mu.Lock()
	events.noChannel = true
	events.mu.Unlock()
	got, err := w.h.svc.RunRaidSchedule(context.Background(), at)
	if err == nil || got.Posted != 0 || got.Failed != 1 {
		t.Fatalf("no-channel run = %+v, %v", got, err)
	}
	if rs, _ := w.scheduleOf(t, raid.ID); rs.PostedAt != nil {
		t.Fatal("a post that reached no channel was marked posted")
	}
	if attempts, leased := w.attemptsOf(t, raid.ID); attempts != 1 || leased {
		t.Fatalf("attempts %d, leased %v", attempts, leased)
	}
	if _, lastErr := w.leaseOf(t, raid.ID); lastErr == nil || *lastErr != errPostReachedNoChannel.Error() {
		t.Fatalf("last error = %v", lastErr)
	}
}

// A row that keeps failing is given up after scheduleMaxAttempts: failed, kept, never retried.
func TestARowIsGivenUpAfterTheAttemptCap(t *testing.T) {
	w, events, _ := postingWorld(t)
	raid := w.createRaids(t, scheduleBody(startIn(10*day), 0, "true", "true"))[0]
	at := time.Now().Add(4 * day)
	events.refuse(&notify.NotSentError{Err: &notify.DiscordError{Status: 500, Body: "boom"}}, nil)
	for i := 0; i < scheduleMaxAttempts; i++ {
		at = at.Add(scheduleRetryCap + time.Minute)
		if got, _ := w.h.svc.RunRaidSchedule(context.Background(), at); got.Failed != 1 {
			t.Fatalf("attempt %d: run = %+v", i+1, got)
		}
	}
	if attempts, _ := w.attemptsOf(t, raid.ID); attempts != scheduleMaxAttempts {
		t.Fatalf("attempts = %d", attempts)
	}
	if !w.nextAttemptInfinite(t, raid.ID) {
		t.Fatal("the capped row is still due again")
	}
	if _, lastErr := w.leaseOf(t, raid.ID); lastErr == nil || *lastErr == "" {
		t.Fatal("the capped row lost its error")
	}
	events.refuse(nil, nil)
	if got := w.run(t, at.Add(2*scheduleRetryCap)); got.Posted+got.Failed != 0 {
		t.Fatalf("a given-up row was retried: %+v", got)
	}
}

// A refusal Discord will repeat (a non-429 4xx) gives the row up at once; a 429 is retried.
func TestARefusalGivesTheRowUpAtOnce(t *testing.T) {
	w, events, _ := postingWorld(t)
	raid := w.createRaids(t, scheduleBody(startIn(10*day), 0, "true", "true"))[0]
	at := time.Now().Add(4 * day)
	events.refuse(&notify.NotSentError{Err: &notify.DiscordError{Status: 429, Body: "slow down"}}, nil)
	if got, _ := w.h.svc.RunRaidSchedule(context.Background(), at); got.Failed != 1 {
		t.Fatalf("429 run = %+v", got)
	}
	if w.nextAttemptInfinite(t, raid.ID) {
		t.Fatal("a 429 gave the row up")
	}
	events.refuse(&notify.NotSentError{Err: &notify.DiscordError{Status: 403, Body: "Missing Access"}}, nil)
	at = at.Add(scheduleRetryCap + time.Minute)
	if got, _ := w.h.svc.RunRaidSchedule(context.Background(), at); got.Failed != 1 {
		t.Fatalf("403 run = %+v", got)
	}
	if attempts, _ := w.attemptsOf(t, raid.ID); attempts != 2 || !w.nextAttemptInfinite(t, raid.ID) {
		t.Fatalf("after a 403: attempts %d, given up %v", attempts, w.nextAttemptInfinite(t, raid.ID))
	}
	events.refuse(nil, nil)
	if got := w.run(t, at.Add(2*scheduleRetryCap)); got.Posted != 0 {
		t.Fatalf("a refused row was retried: %+v", got)
	}
	// An edit plans the unposted raid afresh, so it is due again.
	w.patchTo(t, raid.ID, raid.StartsAt, "true", "true")
	if w.nextAttemptInfinite(t, raid.ID) {
		t.Fatal("an edit left the row given up")
	}
}
