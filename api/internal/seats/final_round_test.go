package seats

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"
)

// A client that goes away while Stripe is cancelling must not leave the
// subscription cancelling with no seat: the step and the commit both finish.
func TestAccept_ClientGoneDuringCancelStillCommits(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	amy := h.member(t, gid, "Amy", "member")
	h.payingPremium(t, amy, testNow.Add(5*24*time.Hour))
	_, _ = h.svc.Give(bg, gid, 1, leader, amy)
	ctx, cancel := context.WithCancel(bg)
	defer cancel()
	var stepDeadline bool
	var stepErr error
	h.cancel.during = func(stepCtx context.Context) {
		cancel()
		_, stepDeadline = stepCtx.Deadline()
		stepErr = stepCtx.Err()
	}
	if err := h.svc.Accept(ctx, gid, amy); err != nil {
		t.Fatalf("accept: %v", err)
	}
	if !stepDeadline || stepErr != nil {
		t.Fatalf("step ctx: deadline=%v err=%v", stepDeadline, stepErr)
	}
	if seat := mustList(t, h, gid)[0]; seat.State != StateAccepted {
		t.Fatalf("seat = %+v", seat)
	}
	if len(h.cancel.resumed) != 0 {
		t.Fatalf("resumed %v", h.cancel.resumed)
	}
}

// A cancel that errors may still have landed at Stripe, so it is undone.
func TestAccept_CancelErrorResumes(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	amy := h.member(t, gid, "Amy", "member")
	h.payingPremium(t, amy, testNow.Add(5*24*time.Hour))
	_, _ = h.svc.Give(bg, gid, 1, leader, amy)
	h.cancel.err = errors.New("timeout")
	if err := h.svc.Accept(bg, gid, amy); err == nil {
		t.Fatal("accept should fail")
	}
	if sub := fmt.Sprintf("sub_user_%d", amy); len(h.cancel.resumed) != 1 || h.cancel.resumed[0] != sub {
		t.Fatalf("resumed %v", h.cancel.resumed)
	}
}

// acceptedUnstarted is Amy holding slot 1, accepted while her own Premium
// (ending in 5 days) runs out, set to end as the webhook would record it.
func acceptedUnstarted(t *testing.T, h *harness) (gid, amy int64) {
	t.Helper()
	gid, leader := h.guild(t, "guild")
	amy = h.member(t, gid, "Amy", "member")
	h.payingPremium(t, amy, testNow.Add(5*24*time.Hour))
	if _, err := h.svc.Give(bg, gid, 1, leader, amy); err != nil {
		t.Fatal(err)
	}
	if err := h.svc.Accept(bg, gid, amy); err != nil {
		t.Fatal(err)
	}
	h.exec(t, `update entitlements set cancel_at_period_end = true where user_id = $1`, amy)
	h.mailer.Sent = nil
	return gid, amy
}

// A seat released before it starts resumes the holder's own Premium instead
// of opening a comeback their Checkout would refuse.
func TestSweep_UnstartedSeatResumesOwnPremium(t *testing.T) {
	h := newHarness(t)
	gid, amy := acceptedUnstarted(t, h)
	h.exec(t, `update entitlements set status = 'canceled' where guild_id = $1`, gid)
	if res, err := h.svc.Sweep(bg, testNow); err != nil || res.Released != 1 {
		t.Fatalf("sweep = %+v, %v", res, err)
	}
	sub := fmt.Sprintf("sub_user_%d", amy)
	if len(h.cancel.resumed) != 1 || h.cancel.resumed[0] != sub {
		t.Fatalf("resumed %v", h.cancel.resumed)
	}
	var n int
	if err := h.pool.QueryRow(bg, `select count(*) from seat_comebacks where user_id = $1`, amy).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 0 || len(h.mailer.Sent) != 0 {
		t.Fatalf("comebacks=%d mails=%d", n, len(h.mailer.Sent))
	}
	if state := slotState(t, h, gid, 1); state != string(StateEmpty) {
		t.Fatalf("state = %s", state)
	}
}

func slotState(t *testing.T, h *harness, gid int64, slot int) string {
	t.Helper()
	var state string
	if err := h.pool.QueryRow(bg, `select state from premium_seats where guild_id = $1 and slot = $2`, gid, slot).
		Scan(&state); err != nil {
		t.Fatal(err)
	}
	return state
}

// A resume Stripe refuses falls back to the comeback: the sweep succeeds,
// the seat empties and the seat-removed email goes out.
func TestSweep_FailedResumeFallsBackToTheComeback(t *testing.T) {
	h := newHarness(t)
	gid, amy := acceptedUnstarted(t, h)
	h.exec(t, `update entitlements set status = 'canceled' where guild_id = $1`, gid)
	h.svc.Canceller = failingResume{h.cancel}
	if res, err := h.svc.Sweep(bg, testNow); err != nil || res.Released != 1 || res.Failed != 0 {
		t.Fatalf("sweep = %+v, %v", res, err)
	}
	if state := slotState(t, h, gid, 1); state != string(StateEmpty) {
		t.Fatalf("state = %s", state)
	}
	var n int
	if err := h.pool.QueryRow(bg, `select count(*) from seat_comebacks where user_id = $1`, amy).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 1 || len(h.mailer.Sent) != 1 || !strings.Contains(h.mailer.Sent[0].Text, "comeback=") {
		t.Fatalf("comebacks=%d mails=%+v", n, h.mailer.Sent)
	}
}

// A resumed subscription is set to end again when the sweep's commit fails.
func TestSweep_CommitFailureRecancelsTheResume(t *testing.T) {
	h := newHarness(t)
	gid, amy := acceptedUnstarted(t, h)
	h.exec(t, `update entitlements set status = 'canceled' where guild_id = $1`, gid)
	h.exec(t, `create function seats_t_fail2() returns trigger as $$ begin raise exception 'boom'; end $$ language plpgsql`)
	h.exec(t, `create constraint trigger seats_t_fail2 after update on premium_seats deferrable initially deferred
		for each row execute function seats_t_fail2()`)
	t.Cleanup(func() {
		h.exec(t, `drop trigger seats_t_fail2 on premium_seats`)
		h.exec(t, `drop function seats_t_fail2()`)
	})
	h.cancel.ids = nil
	if _, err := h.svc.Sweep(bg, testNow); err == nil {
		t.Fatal("sweep should report the failed guild")
	}
	sub := fmt.Sprintf("sub_user_%d", amy)
	if len(h.cancel.resumed) != 1 || len(h.cancel.ids) != 1 || h.cancel.ids[0] != sub {
		t.Fatalf("resumed %v re-cancelled %v", h.cancel.resumed, h.cancel.ids)
	}
	if state := slotState(t, h, gid, 1); state != string(StateAccepted) {
		t.Fatalf("state = %s", state)
	}
}

type failingResume struct{ *fakeCanceller }

func (failingResume) ResumeSubscription(context.Context, string) error {
	return errors.New("stripe down")
}

func TestSweepResult_TotalFailure(t *testing.T) {
	cases := []struct {
		name string
		r    SweepResult
		want bool
	}{
		{"nothing to sweep", SweepResult{}, false},
		{"list failed", SweepResult{ListFailed: true}, true},
		{"every guild failed", SweepResult{Attempted: 2, Failed: 2}, true},
		{"some guilds failed", SweepResult{Attempted: 3, Failed: 2, Released: 1}, false},
		{"all succeeded", SweepResult{Attempted: 2, Released: 4}, false},
	}
	for _, c := range cases {
		if got := c.r.TotalFailure(); got != c.want {
			t.Errorf("%s: TotalFailure() = %v, want %v", c.name, got, c.want)
		}
	}
}

func TestGive_RefusesMemberPendingElsewhere(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "raid_guild")
	bob, amy := h.member(t, gid, "Bob", "member"), h.member(t, gid, "Amy", "member")
	h.payingPremium(t, amy, testNow.Add(5*24*time.Hour))
	if _, err := h.svc.Give(bg, gid, 1, leader, bob); err != nil {
		t.Fatal(err)
	}
	h.svc.Now = func() time.Time { return testNow.Add(31 * 24 * time.Hour) }
	if _, err := h.svc.Move(bg, gid, 1, leader, amy); err != nil {
		t.Fatal(err)
	}
	if _, err := h.svc.Give(bg, gid, 2, leader, amy); !errors.Is(err, ErrTaken) {
		t.Fatalf("give to a pending offeree: %v", err)
	}
}

func TestMove_RefusesPayerHoldingElsewhere(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "raid_guild")
	bob, amy := h.member(t, gid, "Bob", "member"), h.member(t, gid, "Amy", "member")
	h.payingPremium(t, amy, testNow.Add(5*24*time.Hour))
	if _, err := h.svc.Give(bg, gid, 1, leader, bob); err != nil {
		t.Fatal(err)
	}
	if _, err := h.svc.Give(bg, gid, 2, leader, amy); err != nil {
		t.Fatal(err)
	}
	h.svc.Now = func() time.Time { return testNow.Add(31 * 24 * time.Hour) }
	if _, err := h.svc.Move(bg, gid, 1, leader, amy); !errors.Is(err, ErrTaken) {
		t.Fatalf("move to a payer offered slot 2: %v", err)
	}
}
