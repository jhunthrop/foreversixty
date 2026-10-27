package seats

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

func mustList(t *testing.T, h *harness, gid int64) []Seat {
	t.Helper()
	seats, err := h.svc.Store.List(bg, gid, h.svc.now())
	if err != nil {
		t.Fatal(err)
	}
	return seats
}

func TestAccept_LockStartsWhenSeatStarts(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	amy := h.member(t, gid, "Amy", "member")
	end := testNow.Add(200 * 24 * time.Hour)
	h.payingPremium(t, amy, end)
	_, _ = h.svc.Give(bg, gid, 1, leader, amy)
	if err := h.svc.Accept(bg, gid, amy); err != nil {
		t.Fatal(err)
	}
	seat := mustList(t, h, gid)[0]
	if want := end.Add(30 * 24 * time.Hour); !seat.MovableAt.Equal(want) {
		t.Fatalf("movable_at = %v, want %v", seat.MovableAt, want)
	}
}

func TestAccept_AlreadyCancelingSkipsCancelAndKey(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	amy := h.member(t, gid, "Amy", "member")
	h.payingPremiumCanceling(t, amy, testNow.Add(5*24*time.Hour))
	_, _ = h.svc.Give(bg, gid, 1, leader, amy)
	if err := h.svc.Accept(bg, gid, amy); err != nil {
		t.Fatal(err)
	}
	var key *string
	_ = h.pool.QueryRow(bg, `select holder_gave_up_lookup_key from premium_seats where guild_id=$1`, gid).Scan(&key)
	if len(h.cancel.ids) != 0 || key != nil {
		t.Fatalf("cancels=%v key=%v", h.cancel.ids, key)
	}
}

func TestAccept_StalePlanTouchesNothing(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	amy := h.member(t, gid, "Amy", "member")
	h.payingPremium(t, amy, testNow.Add(5*24*time.Hour))
	_, _ = h.svc.Give(bg, gid, 1, leader, amy)
	h.exec(t, `update entitlements set status = 'canceled' where guild_id = $1`, gid)
	if err := h.svc.Accept(bg, gid, amy); !errors.Is(err, ErrNoPlan) {
		t.Fatalf("lapsed: %v", err)
	}
	h.exec(t, `update entitlements set status = 'active' where guild_id = $1`, gid)
	h.exec(t, `update guild_members set verified_at = null where user_id = $1`, amy)
	if err := h.svc.Accept(bg, gid, amy); !errors.Is(err, ErrNotMember) {
		t.Fatalf("left: %v", err)
	}
	if len(h.cancel.ids) != 0 {
		t.Fatalf("cancelled %v", h.cancel.ids)
	}
}

func TestAccept_NoOfferBeforeStripeCheck(t *testing.T) {
	h := newHarness(t)
	gid, _ := h.guild(t, "guild")
	amy := h.member(t, gid, "Amy", "member")
	h.payingPremium(t, amy, testNow.Add(24*time.Hour))
	h.svc.Canceller = nil
	if err := h.svc.Accept(bg, gid, amy); !errors.Is(err, ErrNoOffer) {
		t.Fatalf("err = %v", err)
	}
}

func TestAccept_CommitFailureResumesSubscription(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	amy := h.member(t, gid, "Amy", "member")
	h.payingPremium(t, amy, testNow.Add(5*24*time.Hour))
	_, _ = h.svc.Give(bg, gid, 1, leader, amy)
	h.exec(t, `create function seats_t_fail() returns trigger as $$ begin raise exception 'boom'; end $$ language plpgsql`)
	h.exec(t, `create constraint trigger seats_t_fail after update on premium_seats deferrable initially deferred
		for each row execute function seats_t_fail()`)
	t.Cleanup(func() {
		h.exec(t, `drop trigger seats_t_fail on premium_seats`)
		h.exec(t, `drop function seats_t_fail()`)
	})
	if err := h.svc.Accept(bg, gid, amy); err == nil {
		t.Fatal("accept should fail")
	}
	sub := fmt.Sprintf("sub_user_%d", amy)
	if len(h.cancel.ids) != 1 || len(h.cancel.resumed) != 1 || h.cancel.resumed[0] != sub {
		t.Fatalf("cancelled %v resumed %v", h.cancel.ids, h.cancel.resumed)
	}
}

func TestAccept_CancelFailureWritesNothing(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	amy := h.member(t, gid, "Amy", "member")
	h.payingPremium(t, amy, testNow.Add(5*24*time.Hour))
	_, _ = h.svc.Give(bg, gid, 1, leader, amy)
	h.cancel.err = errors.New("stripe down")
	if err := h.svc.Accept(bg, gid, amy); err == nil {
		t.Fatal("accept should fail")
	}
	if seat := mustList(t, h, gid)[0]; seat.State != StateOffered {
		t.Fatalf("seat = %+v", seat)
	}
}

func TestMove_ToPayerKeepsHolderUntilAccept(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	bob, amy := h.member(t, gid, "Bob", "member"), h.member(t, gid, "Amy", "member")
	h.payingPremium(t, amy, testNow.Add(5*24*time.Hour))
	_, _ = h.svc.Give(bg, gid, 1, leader, bob)
	h.svc.Now = func() time.Time { return testNow.Add(31 * 24 * time.Hour) }
	seat, err := h.svc.Move(bg, gid, 1, leader, amy)
	if err != nil || seat.Holder.UserID != bob || seat.State != StateHeld || seat.Offeree == nil || seat.Offeree.Name != "Amy" {
		t.Fatalf("%+v %v", seat, err)
	}
	if _, err := h.svc.Move(bg, gid, 1, leader, amy); !errors.Is(err, ErrOfferPending) {
		t.Fatalf("second move: %v", err)
	}
	if _, err := h.svc.Give(bg, gid, 1, leader, amy); !errors.Is(err, ErrOfferPending) {
		t.Fatalf("give over offer: %v", err)
	}
	if f, _ := h.svc.Store.SeatFactsForUser(bg, amy, h.svc.now()); f.Offer == nil {
		t.Fatal("offeree sees no offer")
	}
	if err := h.svc.Accept(bg, gid, amy); err != nil {
		t.Fatal(err)
	}
	seat = mustList(t, h, gid)[0]
	if seat.Holder.UserID != amy || seat.State != StateAccepted || seat.Offeree != nil {
		t.Fatalf("after accept %+v", seat)
	}
}

func TestMove_ToPayerDeclineAndWithdrawKeepHolder(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	bob, amy := h.member(t, gid, "Bob", "member"), h.member(t, gid, "Amy", "member")
	h.payingPremium(t, amy, testNow.Add(5*24*time.Hour))
	_, _ = h.svc.Give(bg, gid, 1, leader, bob)
	h.svc.Now = func() time.Time { return testNow.Add(31 * 24 * time.Hour) }
	_, _ = h.svc.Move(bg, gid, 1, leader, amy)
	if err := h.svc.Decline(bg, gid, amy); err != nil {
		t.Fatal(err)
	}
	if s := mustList(t, h, gid)[0]; s.Holder.UserID != bob || s.Offeree != nil {
		t.Fatalf("after decline %+v", s)
	}
	_, _ = h.svc.Move(bg, gid, 1, leader, amy)
	seat, err := h.svc.Withdraw(bg, gid, 1, leader)
	if err != nil || seat.Holder.UserID != bob || seat.Offeree != nil {
		t.Fatalf("withdraw %+v %v", seat, err)
	}
}

func TestMove_Concurrent(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	bob := h.member(t, gid, "Bob", "member")
	a, b := h.member(t, gid, "Aa", "member"), h.member(t, gid, "Bb", "member")
	_, _ = h.svc.Give(bg, gid, 1, leader, bob)
	h.svc.Now = func() time.Time { return testNow.Add(31 * 24 * time.Hour) }
	var wg sync.WaitGroup
	errs := make([]error, 2)
	for i, to := range []int64{a, b} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, errs[i] = h.svc.Move(bg, gid, 1, leader, to)
		}()
	}
	wg.Wait()
	ok, locked := 0, 0
	for _, e := range errs {
		switch {
		case e == nil:
			ok++
		case errors.As(e, new(ErrLocked)):
			locked++
		default:
			t.Fatal(e)
		}
	}
	if ok != 1 || locked != 1 {
		t.Fatalf("ok=%d locked=%d", ok, locked)
	}
}

func TestSweep_DowngradeReleasesUpperSlots(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "raid_guild")
	a, b := h.member(t, gid, "Aa", "member"), h.member(t, gid, "Bb", "member")
	_, _ = h.svc.Give(bg, gid, 1, leader, a)
	_, _ = h.svc.Give(bg, gid, 2, leader, b)
	h.exec(t, `update entitlements set tier = 'guild' where guild_id = $1`, gid)
	n, err := h.svc.Sweep(bg, testNow)
	if err != nil || n.Released != 1 {
		t.Fatalf("n=%+v err=%v", n, err)
	}
	var holder *int64
	_ = h.pool.QueryRow(bg, `select holder_user_id from premium_seats where guild_id=$1 and slot=2`, gid).Scan(&holder)
	if holder != nil {
		t.Fatal("slot 2 still held")
	}
}

func TestIsPayer_DemotedSubscriberStillPays(t *testing.T) {
	h := newHarness(t)
	gid, _ := h.guild(t, "guild")
	sub := h.member(t, gid, "Sub", "member")
	h.exec(t, `update entitlements set billing_user_id = $2 where guild_id = $1`, gid, sub)
	if ok, err := h.svc.Store.IsPayer(bg, gid, sub); err != nil || !ok {
		t.Fatalf("%v %v", ok, err)
	}
}
