package seats

import (
	"context"
	"testing"
	"time"
)

func TestList_SlotsByTier(t *testing.T) {
	for tier, want := range map[string]int{"guild": 1, "raid_guild": 3} {
		h := newHarness(t)
		gid, _ := h.guild(t, tier)
		got, err := h.svc.Store.List(context.Background(), gid, testNow)
		if err != nil || len(got) != want {
			t.Fatalf("%s: %d slots, %v; want %d", tier, len(got), err, want)
		}
		if got[0].State != StateEmpty {
			t.Fatalf("slot 1 = %s", got[0].State)
		}
	}
}

func TestList_ScopedGrantIsEmpty(t *testing.T) {
	h := newHarness(t)
	gid, _ := h.guild(t, "guild")
	h.exec(t, `update entitlements set source = 'grant', features = array['officer_views'] where guild_id = $1`, gid)
	got, err := h.svc.Store.List(context.Background(), gid, testNow)
	if err != nil || len(got) != 0 {
		t.Fatalf("%v, %v", got, err)
	}
}

func TestList_ReturnedShownFor14Days(t *testing.T) {
	h := newHarness(t)
	gid, _ := h.guild(t, "guild")
	h.exec(t, `insert into premium_seats (guild_id, slot, state, returned_name, returned_at) values ($1,1,'empty','Gone',$2)`,
		gid, testNow.Add(-13*24*time.Hour))
	got, _ := h.svc.Store.List(context.Background(), gid, testNow)
	if got[0].Returned == nil || got[0].Returned.Name != "Gone" {
		t.Fatalf("day 13: %+v", got[0].Returned)
	}
	got, _ = h.svc.Store.List(context.Background(), gid, testNow.Add(2*24*time.Hour))
	if got[0].Returned != nil {
		t.Fatalf("day 15 still shown: %+v", got[0].Returned)
	}
}

func TestCandidates_PaysPremiumFlag(t *testing.T) {
	h := newHarness(t)
	gid, _ := h.guild(t, "raid_guild")
	zed := h.member(t, gid, "Zed", "member")
	amy := h.member(t, gid, "Amy", "member")
	holder := h.member(t, gid, "Holder", "member")
	h.payingPremium(t, amy, testNow.Add(10*24*time.Hour))
	h.exec(t, `insert into premium_seats (guild_id, slot, state, holder_user_id, given_at, movable_at) values ($1,1,'held',$2,now(),now())`, gid, holder)
	got, err := h.svc.Store.Candidates(context.Background(), gid)
	if err != nil {
		t.Fatal(err)
	}
	names := []string{}
	for _, c := range got {
		names = append(names, c.Name)
		if c.UserID == amy && !c.PaysPremium || c.UserID == zed && c.PaysPremium {
			t.Fatalf("flags wrong: %+v", c)
		}
	}
	if len(got) != 3 || names[0] != "Amy" || names[1] != "Leader" || names[2] != "Zed" {
		t.Fatalf("candidates = %v", names)
	}
}

func TestIsPayer(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.guild(t, "guild")
	officer := h.member(t, gid, "Off", "officer")
	member := h.member(t, gid, "Mem", "member")
	for uid, want := range map[int64]bool{leader: true, officer: false, member: false} {
		if got, err := h.svc.Store.IsPayer(context.Background(), gid, uid); err != nil || got != want {
			t.Fatalf("user %d: %v %v", uid, got, err)
		}
	}
	if ok, _ := h.svc.Store.IsSubscriber(context.Background(), gid, leader); !ok {
		t.Fatal("leader should be the subscriber")
	}
}
