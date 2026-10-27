package entitlements

import (
	"context"
	"testing"
	"time"
)

func TestUpsertStripe_WritesPriceFacts(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool)
	officer := seedUser(t, pool)
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	trialEnd := time.Now().Add(14 * 24 * time.Hour).UTC().Truncate(time.Second)
	if _, err := s.UpsertStripe(context.Background(), StripeUpsert{
		GuildID: i64(gid), Plan: PlanGuild, Status: "trialing", CurrentPeriodEnd: &end,
		StripeSubscriptionID: "sub_pf", BillingUserID: officer, Actor: "t",
		Tier: TierGuild, PriceLookupKey: "guild_founder_yearly", UnitAmountCents: 12000, TrialEnd: &trialEnd,
	}); err != nil {
		t.Fatal(err)
	}
	gb, err := s.GuildBilling(context.Background(), gid)
	if err != nil || gb == nil {
		t.Fatalf("guild billing = %+v, %v", gb, err)
	}
	if gb.Tier != TierGuild || gb.PriceLookupKey != "guild_founder_yearly" || gb.UnitAmountCents != 12000 ||
		gb.TrialEnd == nil || !gb.TrialEnd.Equal(trialEnd) || gb.StripeSubscriptionID != "sub_pf" || gb.ID == 0 {
		t.Fatalf("price facts not read back: %+v", gb)
	}
}

func TestUpsertStripe_GuildWithoutTierFails(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool)
	officer := seedUser(t, pool)
	_, err := s.UpsertStripe(context.Background(), StripeUpsert{
		GuildID: i64(gid), Plan: PlanGuild, Status: "active",
		StripeSubscriptionID: "sub_nt", BillingUserID: officer, Actor: "t",
	})
	if err == nil {
		t.Fatal("a guild upsert without a tier must fail fast")
	}
	if gb, _ := s.GuildBilling(context.Background(), gid); gb != nil {
		t.Fatalf("nothing may be written, got %+v", gb)
	}
}

func TestUpsertStripe_PremiumIgnoresTier(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	uid := seedUser(t, pool)
	if _, err := s.UpsertStripe(context.Background(), StripeUpsert{
		UserID: i64(uid), Plan: PlanPremium, Status: "active",
		StripeSubscriptionID: "sub_pt", BillingUserID: uid, Actor: "t", Tier: TierGuild,
	}); err != nil {
		t.Fatal(err)
	}
	var tier *string
	if err := pool.QueryRow(context.Background(),
		`select tier from entitlements where user_id = $1`, uid).Scan(&tier); err != nil {
		t.Fatal(err)
	}
	if tier != nil {
		t.Fatalf("premium tier = %q, want null", *tier)
	}
}

func TestPersonalBilling_ReturnsSubscriptionAndKey(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	uid := seedUser(t, pool)
	if _, err := s.UpsertStripe(context.Background(), StripeUpsert{
		UserID: i64(uid), Plan: PlanPremium, Status: "active",
		StripeSubscriptionID: "sub_pb", BillingUserID: uid, Actor: "t",
		PriceLookupKey: "premium_yearly", UnitAmountCents: 4000,
	}); err != nil {
		t.Fatal(err)
	}
	b, err := s.PersonalBilling(context.Background(), uid)
	if err != nil || b == nil {
		t.Fatalf("billing = %+v, %v", b, err)
	}
	if b.Source != "stripe" || b.StripeSubscriptionID != "sub_pb" ||
		b.PriceLookupKey != "premium_yearly" || b.UnitAmountCents != 4000 {
		t.Fatalf("billing = %+v", b)
	}
}

func TestGrant_GuildRowIsRaidGuild(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	whole, scoped := seedGuild(t, pool), seedGuild(t, pool)
	ctx := context.Background()
	if err := s.Grant(ctx, Subject{GuildID: i64(whole)}, PlanGuild, nil, nil, "n", "a", nil); err != nil {
		t.Fatal(err)
	}
	if err := s.Grant(ctx, Subject{GuildID: i64(scoped)}, PlanGuild, nil, nil, "n", "a",
		[]string{string(FeatureOfficerViews)}); err != nil {
		t.Fatal(err)
	}
	for _, gid := range []int64{whole, scoped} {
		gb, err := s.GuildBilling(ctx, gid)
		if err != nil || gb == nil || gb.Tier != TierRaidGuild {
			t.Fatalf("guild %d tier = %+v, %v", gid, gb, err)
		}
	}
}
