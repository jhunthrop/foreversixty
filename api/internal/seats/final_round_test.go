package seats

import (
	"context"
	"errors"
	"fmt"
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
