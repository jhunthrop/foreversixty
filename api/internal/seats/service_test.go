package seats

import (
	"context"
	"errors"
	"strings"
	"sync"
	"testing"
	"time"
)

var bg = context.Background()

func TestGive_NonPayerHeldWithLock(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	bob := h.member(t, gid, "Bob", "member")
	seat, err := h.svc.Give(bg, gid, 1, leader, bob)
	if err != nil {
		t.Fatal(err)
	}
	if seat.State != StateHeld || seat.Holder == nil || seat.Holder.Name != "Bob" {
		t.Fatalf("seat = %+v", seat)
	}
	if seat.MovableAt == nil || !seat.MovableAt.Equal(testNow.Add(30*24*time.Hour)) {
		t.Fatalf("movable_at = %v", seat.MovableAt)
	}
	if len(h.mailer.Sent) != 1 || !strings.Contains(h.mailer.Sent[0].Subject, "You have Premium from Olympus") {
		t.Fatalf("mail = %+v", h.mailer.Sent)
	}
}

func TestGive_PayerOffered(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	amy := h.member(t, gid, "Amy", "member")
	h.payingPremium(t, amy, testNow.Add(5*24*time.Hour))
	seat, err := h.svc.Give(bg, gid, 1, leader, amy)
	if err != nil {
		t.Fatal(err)
	}
	if seat.State != StateOffered || seat.MovableAt != nil {
		t.Fatalf("seat = %+v", seat)
	}
	if len(h.mailer.Sent) != 1 || !strings.Contains(h.mailer.Sent[0].Subject, "offers you a Premium seat") {
		t.Fatalf("mail = %+v", h.mailer.Sent)
	}
}

func TestGive_ConcurrentSameSlot(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	a, b := h.member(t, gid, "Aa", "member"), h.member(t, gid, "Bb", "member")
	var wg sync.WaitGroup
	errs := make([]error, 2)
	for i, to := range []int64{a, b} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, errs[i] = h.svc.Give(bg, gid, 1, leader, to)
		}()
	}
	wg.Wait()
	ok, taken := 0, 0
	for _, e := range errs {
		switch {
		case e == nil:
			ok++
		case errors.Is(e, ErrTaken):
			taken++
		default:
			t.Fatal(e)
		}
	}
	if ok != 1 || taken != 1 {
		t.Fatalf("ok=%d taken=%d", ok, taken)
	}
}

func TestGive_RefusesNonPayerActor(t *testing.T) {
	h := newHarness(t)
	gid, _ := h.guild(t, "guild")
	off := h.member(t, gid, "Off", "officer")
	bob := h.member(t, gid, "Bob", "member")
	if _, err := h.svc.Give(bg, gid, 1, off, bob); !errors.Is(err, ErrNotPayer) {
		t.Fatalf("err = %v", err)
	}
}

func TestGive_RefusesNonMember(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	outsider := h.user(t, "out")
	if _, err := h.svc.Give(bg, gid, 1, leader, outsider); !errors.Is(err, ErrNotMember) {
		t.Fatalf("err = %v", err)
	}
	if _, err := h.svc.Give(bg, gid, 2, leader, outsider); !errors.Is(err, ErrNoSuchSlot) {
		t.Fatalf("slot err = %v", err)
	}
}

func TestMove_LockedReturnsMovableAt(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	bob, cy := h.member(t, gid, "Bob", "member"), h.member(t, gid, "Cy", "member")
	if _, err := h.svc.Give(bg, gid, 1, leader, bob); err != nil {
		t.Fatal(err)
	}
	_, err := h.svc.Move(bg, gid, 1, leader, cy)
	var locked ErrLocked
	if !errors.As(err, &locked) || !locked.MovableAt.Equal(testNow.Add(30*24*time.Hour)) {
		t.Fatalf("err = %v", err)
	}
}

func TestMove_OldHolderKeepsUntilMove(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	bob, cy := h.member(t, gid, "Bob", "member"), h.member(t, gid, "Cy", "member")
	if _, err := h.svc.Give(bg, gid, 1, leader, bob); err != nil {
		t.Fatal(err)
	}
	_, _ = h.svc.Move(bg, gid, 1, leader, cy) // locked
	if seat, _ := h.svc.Store.List(bg, gid, testNow); seat[0].Holder.UserID != bob {
		t.Fatalf("holder = %+v", seat[0].Holder)
	}
	later := testNow.Add(31 * 24 * time.Hour)
	h.svc.Now = func() time.Time { return later }
	seat, err := h.svc.Move(bg, gid, 1, leader, cy)
	if err != nil || seat.Holder.UserID != cy || seat.State != StateHeld {
		t.Fatalf("%+v %v", seat, err)
	}
}

func TestWithdraw_StartsNoLock(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	amy := h.member(t, gid, "Amy", "member")
	bob := h.member(t, gid, "Bob", "member")
	h.payingPremium(t, amy, testNow.Add(5*24*time.Hour))
	if _, err := h.svc.Give(bg, gid, 1, leader, amy); err != nil {
		t.Fatal(err)
	}
	seat, err := h.svc.Withdraw(bg, gid, 1, leader)
	if err != nil || seat.State != StateEmpty || seat.MovableAt != nil {
		t.Fatalf("%+v %v", seat, err)
	}
	if _, err := h.svc.Give(bg, gid, 1, leader, bob); err != nil {
		t.Fatalf("give right after withdraw: %v", err)
	}
	if _, err := h.svc.Withdraw(bg, gid, 1, leader); !errors.Is(err, ErrNoOffer) {
		t.Fatalf("withdraw held: %v", err)
	}
}

func TestAccept_CancelsOwnAtPeriodEnd(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	amy := h.member(t, gid, "Amy", "member")
	end := testNow.Add(5 * 24 * time.Hour)
	h.payingPremium(t, amy, end)
	if _, err := h.svc.Give(bg, gid, 1, leader, amy); err != nil {
		t.Fatal(err)
	}
	if err := h.svc.Accept(bg, gid, amy); err != nil {
		t.Fatal(err)
	}
	if len(h.cancel.ids) != 1 || h.cancel.ids[0] == "" {
		t.Fatalf("cancels = %v", h.cancel.ids)
	}
	seat, _ := h.svc.Store.List(bg, gid, testNow)
	if seat[0].State != StateAccepted || !seat[0].StartsAt.Equal(end) {
		t.Fatalf("seat = %+v", seat[0])
	}
	if !seat[0].MovableAt.Equal(testNow.Add(30 * 24 * time.Hour)) {
		t.Fatalf("lock = %v", seat[0].MovableAt)
	}
	var key string
	if err := h.pool.QueryRow(bg, `select holder_gave_up_lookup_key from premium_seats where guild_id=$1`, gid).Scan(&key); err != nil || key != "premium_founder_monthly" {
		t.Fatalf("key = %q %v", key, err)
	}
}

func TestAccept_NoStripeUnavailable(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	amy := h.member(t, gid, "Amy", "member")
	h.payingPremium(t, amy, testNow.Add(24*time.Hour))
	_, _ = h.svc.Give(bg, gid, 1, leader, amy)
	h.svc.Canceller = nil
	if err := h.svc.Accept(bg, gid, amy); !errors.Is(err, ErrBillingUnavailable) {
		t.Fatalf("err = %v", err)
	}
}

func TestDecline_EmptiesSlot(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	amy := h.member(t, gid, "Amy", "member")
	h.payingPremium(t, amy, testNow.Add(24*time.Hour))
	_, _ = h.svc.Give(bg, gid, 1, leader, amy)
	if err := h.svc.Decline(bg, gid, amy); err != nil {
		t.Fatal(err)
	}
	seat, _ := h.svc.Store.List(bg, gid, testNow)
	if seat[0].State != StateEmpty || seat[0].Holder != nil {
		t.Fatalf("seat = %+v", seat[0])
	}
	if err := h.svc.Decline(bg, gid, amy); !errors.Is(err, ErrNoOffer) {
		t.Fatalf("second decline: %v", err)
	}
}

func TestSweep_DepartedHolderReleasedOnce(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	amy := h.member(t, gid, "Amy", "member")
	h.payingPremium(t, amy, testNow.Add(5*24*time.Hour))
	_, _ = h.svc.Give(bg, gid, 1, leader, amy)
	if err := h.svc.Accept(bg, gid, amy); err != nil {
		t.Fatal(err)
	}
	h.mailer.Sent = nil
	h.exec(t, `delete from guild_members where user_id = $1`, amy)
	h.exec(t, `delete from guild_characters where user_id = $1`, amy)
	for i := 0; i < 2; i++ {
		if _, err := h.svc.Sweep(bg, testNow); err != nil {
			t.Fatal(err)
		}
	}
	seat, _ := h.svc.Store.List(bg, gid, testNow)
	if seat[0].State != StateEmpty || seat[0].Returned == nil || seat[0].Returned.Name != "Amy" || !seat[0].Returned.At.Equal(testNow) {
		t.Fatalf("seat = %+v", seat[0])
	}
	var n int
	_ = h.pool.QueryRow(bg, `select count(*) from seat_comebacks where user_id = $1`, amy).Scan(&n)
	if n != 1 || len(h.mailer.Sent) != 1 {
		t.Fatalf("comebacks=%d mails=%d", n, len(h.mailer.Sent))
	}
	text := h.mailer.Sent[0].Text
	if !strings.Contains(text, "/premium/checkout?") || !strings.Contains(text, "comeback=") ||
		!strings.Contains(text, "interval=monthly") {
		t.Fatalf("mail = %s", text)
	}
}

func TestSweep_LapsedPlanReleasesSeats(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "raid_guild")
	bob := h.member(t, gid, "Bob", "member")
	_, _ = h.svc.Give(bg, gid, 2, leader, bob)
	h.exec(t, `update entitlements set status = 'canceled' where guild_id = $1`, gid)
	n, err := h.svc.Sweep(bg, testNow)
	if err != nil || n != 1 {
		t.Fatalf("n=%d err=%v", n, err)
	}
	var state string
	_ = h.pool.QueryRow(bg, `select state from premium_seats where guild_id=$1 and slot=2`, gid).Scan(&state)
	if state != "empty" {
		t.Fatalf("state = %s", state)
	}
}

func TestFrozenClaimantRefused(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	bob := h.member(t, gid, "Bob", "member")
	h.svc.Guilds = fakeFrozen{frozen: true}
	if _, err := h.svc.Give(bg, gid, 1, leader, bob); !errors.Is(err, ErrClaimContested) {
		t.Fatalf("err = %v", err)
	}
}

func TestSeatFactsForUser(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	bob := h.member(t, gid, "Bob", "member")
	_, _ = h.svc.Give(bg, gid, 1, leader, bob)
	f, err := h.svc.Store.SeatFactsForUser(bg, bob, testNow)
	if err != nil || f.Seat == nil || f.Seat.GuildName != "Olympus" || f.Seat.GivenBy != "Leader" || f.Offer != nil {
		t.Fatalf("%+v %v", f, err)
	}
}
