// api/internal/entitlements/store_test.go
package entitlements

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

func testPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set; start docker-compose.test.yml")
	}
	if err := db.Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := db.Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	if _, err := pool.Exec(context.Background(),
		`truncate users, guilds, entitlements, entitlement_audit, entitlement_anomalies cascade`); err != nil {
		t.Fatal(err)
	}
	return pool
}

func seedUser(t *testing.T, pool *pgxpool.Pool) int64 {
	t.Helper()
	var id int64
	if err := pool.QueryRow(context.Background(),
		`insert into users (email) values ($1) returning id`,
		"e-"+time.Now().Format("150405.000000000")+"@example.com").Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

func seedGuild(t *testing.T, pool *pgxpool.Pool) int64 {
	t.Helper()
	var id int64
	if err := pool.QueryRow(context.Background(),
		`insert into guilds (region, ruleset, name) values ('us', 'hardcore', $1) returning id`,
		"guild-"+time.Now().Format("150405.000000000")).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

func seedGuildMember(t *testing.T, pool *pgxpool.Pool, guildID, userID int64, rank string, verified bool) {
	t.Helper()
	var verifiedAt any
	if verified {
		verifiedAt = time.Now()
	}
	if _, err := pool.Exec(context.Background(),
		`insert into guild_members (guild_id, user_id, rank, verified_at, refreshed_at)
		 values ($1, $2, $3, $4, now())`, guildID, userID, rank, verifiedAt); err != nil {
		t.Fatal(err)
	}
}

func seedEntitlement(t *testing.T, pool *pgxpool.Pool, subj Subject, plan, status string) {
	t.Helper()
	if _, err := pool.Exec(context.Background(),
		`insert into entitlements (user_id, guild_id, plan, source, status, tier)
		 values ($1, $2, $3, 'grant', $4, case when $3 = 'guild' then 'raid_guild' end)`, subj.UserID, subj.GuildID, plan, status); err != nil {
		t.Fatal(err)
	}
}

func i64(v int64) *int64 { return &v }

func TestCanSignedOut(t *testing.T) {
	s := &Store{Pool: testPool(t)}
	ok, reason, err := s.Can(context.Background(), 0, FeatureServerSims)
	if err != nil {
		t.Fatal(err)
	}
	if ok || reason != ReasonSignInRequired {
		t.Fatalf("Can(0) = %v, %v, want false, %v", ok, reason, ReasonSignInRequired)
	}
}

func TestCanPersonalPremiumStatuses(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	for _, tc := range []struct {
		status string
		want   bool
	}{
		{"active", true}, {"trialing", true}, {"past_due", true},
		{"canceled", false}, {"unpaid", false}, {"incomplete", false}, {"incomplete_expired", false},
	} {
		t.Run(tc.status, func(t *testing.T) {
			uid := seedUser(t, pool)
			seedEntitlement(t, pool, Subject{UserID: i64(uid)}, PlanPremium, tc.status)
			ok, reason, err := s.Can(context.Background(), uid, FeatureServerSims)
			if err != nil {
				t.Fatal(err)
			}
			if ok != tc.want {
				t.Fatalf("Can(%s) = %v, %v, want ok=%v", tc.status, ok, reason, tc.want)
			}
			if !ok && reason != ReasonNoPlan {
				t.Fatalf("reason = %v, want %v", reason, ReasonNoPlan)
			}
		})
	}
}

func TestCanGuildPlanRequiresVerifiedMembership(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool)
	seedEntitlement(t, pool, Subject{GuildID: i64(gid)}, PlanGuild, "active")

	verified := seedUser(t, pool)
	seedGuildMember(t, pool, gid, verified, "member", true)
	if ok, _, err := s.Can(context.Background(), verified, FeatureHistory); err != nil || !ok {
		t.Fatalf("verified member: Can = %v, %v", ok, err)
	}

	unverified := seedUser(t, pool)
	seedGuildMember(t, pool, gid, unverified, "member", false)
	ok, reason, err := s.Can(context.Background(), unverified, FeatureHistory)
	if err != nil {
		t.Fatal(err)
	}
	if ok || reason != ReasonNoPlan {
		// RULING 2's regression case: an unverified guild_members row must
		// answer exactly like no row at all.
		t.Fatalf("unverified member: Can = %v, %v, want false, %v", ok, reason, ReasonNoPlan)
	}
}

func TestCanOfficerViewsRequiresGuildPlanAndVerifiedOfficerRank(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool)

	// No guild plan at all: a verified officer still gets ReasonNoPlan.
	officerNoPlan := seedUser(t, pool)
	seedGuildMember(t, pool, gid, officerNoPlan, "officer", true)
	if ok, reason, err := s.Can(context.Background(), officerNoPlan, FeatureOfficerViews); err != nil || ok || reason != ReasonNoPlan {
		t.Fatalf("officer, no plan: %v, %v, %v", ok, reason, err)
	}

	seedEntitlement(t, pool, Subject{GuildID: i64(gid)}, PlanGuild, "active")

	member := seedUser(t, pool)
	seedGuildMember(t, pool, gid, member, "member", true)
	if ok, reason, err := s.Can(context.Background(), member, FeatureOfficerViews); err != nil || ok || reason != ReasonNotOfficer {
		t.Fatalf("verified member, guild has plan: %v, %v, %v", ok, reason, err)
	}

	officer := seedUser(t, pool)
	seedGuildMember(t, pool, gid, officer, "officer", true)
	if ok, _, err := s.Can(context.Background(), officer, FeatureOfficerViews); err != nil || !ok {
		t.Fatalf("verified officer, guild has plan: %v, %v", ok, err)
	}

	// A premium (not guild) entitlement never unlocks officer_views (RULING 3).
	premiumOnly := seedUser(t, pool)
	seedEntitlement(t, pool, Subject{UserID: i64(premiumOnly)}, PlanPremium, "active")
	if ok, reason, err := s.Can(context.Background(), premiumOnly, FeatureOfficerViews); err != nil || ok || reason != ReasonNoPlan {
		t.Fatalf("premium-only: %v, %v, %v", ok, reason, err)
	}
}

func TestIsSupporterMatchesTheStandardFeatureRule(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	uid := seedUser(t, pool)
	if ok, err := s.IsSupporter(context.Background(), uid); err != nil || ok {
		t.Fatalf("no entitlement: %v, %v", ok, err)
	}
	seedEntitlement(t, pool, Subject{UserID: i64(uid)}, PlanPremium, "active")
	if ok, err := s.IsSupporter(context.Background(), uid); err != nil || !ok {
		t.Fatalf("active premium: %v, %v", ok, err)
	}
	if ok, err := s.IsSupporter(context.Background(), 0); err != nil || ok {
		t.Fatalf("signed out: %v, %v", ok, err)
	}
}

func TestUpsertStripeInsertsThenUpdatesTheSameRowAndAudits(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	uid := seedUser(t, pool)
	end := time.Now().Add(30 * 24 * time.Hour).UTC().Truncate(time.Second)
	if _, err := s.UpsertStripe(context.Background(), StripeUpsert{
		UserID: i64(uid), Plan: PlanPremium, Status: "active", CurrentPeriodEnd: &end,
		StripeSubscriptionID: "sub_1", BillingUserID: uid, Actor: "stripe_webhook:evt_1",
	}); err != nil {
		t.Fatal(err)
	}
	b, err := s.PersonalBilling(context.Background(), uid)
	if err != nil || b == nil || b.Status != "active" {
		t.Fatalf("after insert: %+v, %v", b, err)
	}

	// A later event (renewal) updates the same row, not a second one.
	end2 := end.Add(30 * 24 * time.Hour)
	if _, err := s.UpsertStripe(context.Background(), StripeUpsert{
		UserID: i64(uid), Plan: PlanPremium, Status: "active", CurrentPeriodEnd: &end2,
		StripeSubscriptionID: "sub_1", BillingUserID: uid, Actor: "stripe_webhook:evt_2",
	}); err != nil {
		t.Fatal(err)
	}
	var rows int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from entitlements where user_id = $1 and plan = 'premium'`, uid).Scan(&rows); err != nil {
		t.Fatal(err)
	}
	if rows != 1 {
		t.Fatalf("row count = %d, want 1 (upsert, not a new row)", rows)
	}
	var audits int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from entitlement_audit where user_id = $1`, uid).Scan(&audits); err != nil {
		t.Fatal(err)
	}
	if audits != 2 {
		t.Fatalf("audit rows = %d, want 2 (one per UpsertStripe call)", audits)
	}
}

func TestUpsertStripeGuildPlanUpsertsByGuildID(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool)
	officer := seedUser(t, pool)
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	up := StripeUpsert{
		GuildID: i64(gid), Plan: PlanGuild, Tier: TierGuild, Status: "active", CurrentPeriodEnd: &end,
		StripeSubscriptionID: "sub_guild_1", BillingUserID: officer, Actor: "stripe_webhook:evt_1",
	}
	if _, err := s.UpsertStripe(context.Background(), up); err != nil {
		t.Fatal(err)
	}
	up.Status, up.Actor = "past_due", "stripe_webhook:evt_2"
	if _, err := s.UpsertStripe(context.Background(), up); err != nil {
		t.Fatal(err)
	}
	gb, err := s.GuildBilling(context.Background(), gid)
	if err != nil || gb == nil || gb.Status != "past_due" || gb.BillingUserID == nil || *gb.BillingUserID != officer {
		t.Fatalf("guild billing = %+v, %v", gb, err)
	}
}

// TestUpsertStripeKeepsAnExistingActiveRowOnADifferentSubscriptionID is the
// entitlements-layer half of the security review's double-billing fix
// (2026-09-21): two Checkout Sessions for the same guild, each completing
// its own subscription, must never let the second silently overwrite the
// first — the first officer's subscription would otherwise keep billing
// with nothing in our database pointing at it.
func TestUpsertStripeKeepsAnExistingActiveRowOnADifferentSubscriptionID(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool)
	officerA := seedUser(t, pool)
	officerB := seedUser(t, pool)
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)

	res, err := s.UpsertStripe(context.Background(), StripeUpsert{
		GuildID: i64(gid), Plan: PlanGuild, Tier: TierGuild, Status: "active", CurrentPeriodEnd: &end,
		StripeSubscriptionID: "sub_a", BillingUserID: officerA, Actor: "stripe_webhook:evt_a",
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.Outcome != OutcomeWritten {
		t.Fatalf("first write must not be flagged a duplicate: %+v", res)
	}

	res, err = s.UpsertStripe(context.Background(), StripeUpsert{
		GuildID: i64(gid), Plan: PlanGuild, Tier: TierGuild, Status: "active", CurrentPeriodEnd: &end,
		StripeSubscriptionID: "sub_b", BillingUserID: officerB, Actor: "stripe_webhook:evt_b",
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.Outcome != OutcomeDuplicate || res.ExistingSubscriptionID != "sub_a" {
		t.Fatalf("second write = %+v, want OutcomeDuplicate ExistingSubscriptionID=sub_a", res)
	}

	gb, err := s.GuildBilling(context.Background(), gid)
	if err != nil || gb == nil || gb.BillingUserID == nil || *gb.BillingUserID != officerA {
		t.Fatalf("guild billing after duplicate write = %+v, %v — sub_a's row must be untouched", gb, err)
	}
	var subID string
	if err := pool.QueryRow(context.Background(),
		`select stripe_subscription_id from entitlements where guild_id = $1 and plan = 'guild'`, gid).Scan(&subID); err != nil {
		t.Fatal(err)
	}
	if subID != "sub_a" {
		t.Fatalf("stripe_subscription_id = %q, want sub_a (untouched)", subID)
	}

	var anomalies int
	var kind, gotSubID, actor string
	if err := pool.QueryRow(context.Background(),
		`select count(*) from entitlement_anomalies where guild_id = $1`, gid).Scan(&anomalies); err != nil {
		t.Fatal(err)
	}
	if anomalies != 1 {
		t.Fatalf("anomaly rows = %d, want 1", anomalies)
	}
	if err := pool.QueryRow(context.Background(),
		`select kind, stripe_subscription_id, actor from entitlement_anomalies where guild_id = $1`, gid).
		Scan(&kind, &gotSubID, &actor); err != nil {
		t.Fatal(err)
	}
	if kind != "duplicate_subscription" || gotSubID != "sub_b" || actor != "stripe_webhook:evt_b" {
		t.Fatalf("anomaly = kind=%q sub=%q actor=%q, want duplicate_subscription/sub_b/stripe_webhook:evt_b", kind, gotSubID, actor)
	}

	// The audit trail records only the write that actually happened.
	var audits int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from entitlement_audit where guild_id = $1`, gid).Scan(&audits); err != nil {
		t.Fatal(err)
	}
	if audits != 1 {
		t.Fatalf("audit rows = %d, want 1 (the skipped duplicate write leaves no audit row)", audits)
	}
}

// TestUpsertStripeAllowsARenewalOnTheSameSubscriptionID confirms the
// duplicate guard keys off a *different* subscription id, not merely an
// already-active row — a renewal event for the same subscription (the
// ordinary, overwhelmingly common case) must keep upserting in place.
func TestUpsertStripeAllowsARenewalOnTheSameSubscriptionID(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	uid := seedUser(t, pool)
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	if _, err := s.UpsertStripe(context.Background(), StripeUpsert{
		UserID: i64(uid), Plan: PlanPremium, Status: "active", CurrentPeriodEnd: &end,
		StripeSubscriptionID: "sub_renew", BillingUserID: uid, Actor: "stripe_webhook:evt_1",
	}); err != nil {
		t.Fatal(err)
	}
	end2 := end.Add(24 * time.Hour)
	res, err := s.UpsertStripe(context.Background(), StripeUpsert{
		UserID: i64(uid), Plan: PlanPremium, Status: "active", CurrentPeriodEnd: &end2,
		StripeSubscriptionID: "sub_renew", BillingUserID: uid, Actor: "stripe_webhook:evt_2",
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.Outcome != OutcomeWritten {
		t.Fatalf("a renewal on the same subscription id must not be flagged a duplicate: %+v", res)
	}
	b, err := s.PersonalBilling(context.Background(), uid)
	if err != nil || b == nil || !b.CurrentPeriodEnd.Equal(end2) {
		t.Fatalf("after renewal: %+v, %v", b, err)
	}
}

// TestUpsertStripeTransferOverwritesInsteadOfFlaggingADuplicate is a
// regression test for a gap an independent review found in an earlier
// version of the double-billing fix (2026-09-21): the guild-billing
// "Take over billing" handoff (spec §2.9 RULING 10) deliberately creates
// a second, different subscription for a guild that already has an
// active one, and *intends* the new one to become the entitled
// subscription — the opposite of the duplicate guard's default behavior.
// A StripeUpsert carrying Transfer: true must overwrite, not flag.
func TestUpsertStripeTransferOverwritesInsteadOfFlaggingADuplicate(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool)
	oldOfficer := seedUser(t, pool)
	newOfficer := seedUser(t, pool)
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)

	if _, err := s.UpsertStripe(context.Background(), StripeUpsert{
		GuildID: i64(gid), Plan: PlanGuild, Tier: TierGuild, Status: "active", CurrentPeriodEnd: &end,
		StripeSubscriptionID: "sub_old_officer", BillingUserID: oldOfficer, Actor: "stripe_webhook:evt_old",
		SubscriptionCreated: time.Now().Add(-2 * time.Hour),
	}); err != nil {
		t.Fatal(err)
	}

	res, err := s.UpsertStripe(context.Background(), StripeUpsert{
		GuildID: i64(gid), Plan: PlanGuild, Tier: TierGuild, Status: "active", CurrentPeriodEnd: &end,
		StripeSubscriptionID: "sub_new_officer", BillingUserID: newOfficer, Actor: "stripe_webhook:evt_transfer",
		Transfer: true, SubscriptionCreated: time.Now().Add(-time.Hour),
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.Outcome != OutcomeWritten {
		t.Fatalf("a transfer upsert must not be flagged a duplicate: %+v", res)
	}

	gb, err := s.GuildBilling(context.Background(), gid)
	if err != nil || gb == nil || gb.BillingUserID == nil || *gb.BillingUserID != newOfficer {
		t.Fatalf("guild billing after transfer = %+v, %v, want billing_user_id = newOfficer", gb, err)
	}
	var subID string
	if err := pool.QueryRow(context.Background(),
		`select stripe_subscription_id from entitlements where guild_id = $1 and plan = 'guild'`, gid).Scan(&subID); err != nil {
		t.Fatal(err)
	}
	if subID != "sub_new_officer" {
		t.Fatalf("stripe_subscription_id = %q, want sub_new_officer (the transfer must overwrite)", subID)
	}
	var anomalies int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from entitlement_anomalies where guild_id = $1`, gid).Scan(&anomalies); err != nil {
		t.Fatal(err)
	}
	if anomalies != 0 {
		t.Fatalf("anomaly rows = %d, want 0 for a deliberate transfer", anomalies)
	}
}

func TestSetCanceledSetsGraceUntilFromTheLastKnownPeriodEnd(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	uid := seedUser(t, pool)
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	if _, err := s.UpsertStripe(context.Background(), StripeUpsert{
		UserID: i64(uid), Plan: PlanPremium, Status: "active", CurrentPeriodEnd: &end,
		StripeSubscriptionID: "sub_cancel_1", BillingUserID: uid, Actor: "stripe_webhook:evt_1",
	}); err != nil {
		t.Fatal(err)
	}
	grace := end.Add(30 * 24 * time.Hour)
	if err := s.SetCanceled(context.Background(), "sub_cancel_1", &grace, "stripe_webhook:evt_2"); err != nil {
		t.Fatal(err)
	}
	var status string
	var got time.Time
	if err := pool.QueryRow(context.Background(),
		`select status, grace_until from entitlements where user_id = $1`, uid).Scan(&status, &got); err != nil {
		t.Fatal(err)
	}
	if status != "canceled" || !got.Equal(grace) {
		t.Fatalf("status=%q grace=%v, want canceled, %v", status, got, grace)
	}
	if ok, _, _ := s.Can(context.Background(), uid, FeatureServerSims); ok {
		t.Fatal("a canceled entitlement must not grant a feature")
	}
}

func TestGrantAndRevokeRoundTrip(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	uid := seedUser(t, pool)
	grantedBy := seedUser(t, pool)
	if err := s.Grant(context.Background(), Subject{UserID: i64(uid)}, PlanPremium, nil,
		&grantedBy, "beta tester", "cli_grant:tester beta tester", nil); err != nil {
		t.Fatal(err)
	}
	if ok, _, err := s.Can(context.Background(), uid, FeatureServerSims); err != nil || !ok {
		t.Fatalf("after grant: %v, %v", ok, err)
	}
	if err := s.Revoke(context.Background(), Subject{UserID: i64(uid)}, PlanPremium,
		"cli_revoke:tester requested by support"); err != nil {
		t.Fatal(err)
	}
	if ok, _, err := s.Can(context.Background(), uid, FeatureServerSims); err != nil || ok {
		t.Fatalf("after revoke: %v, %v", ok, err)
	}
	if err := s.Revoke(context.Background(), Subject{UserID: i64(999999)}, PlanPremium, "cli_revoke:x"); err != ErrNotFound {
		t.Fatalf("revoke unknown subject: %v, want ErrNotFound", err)
	}
}

func TestGrantFeaturesScopeAndClear(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool)
	until := tomorrow().UTC().Truncate(time.Second)
	readFeatures := func() []string {
		var f []string
		if err := pool.QueryRow(ctx, `select features from entitlements where guild_id = $1 and plan = 'guild'`, gid).Scan(&f); err != nil {
			t.Fatal(err)
		}
		return f
	}
	if err := s.Grant(ctx, Subject{GuildID: i64(gid)}, PlanGuild, &until, nil, "n", "a",
		[]string{FeatureScopeOfficerViews}); err != nil {
		t.Fatal(err)
	}
	if f := readFeatures(); len(f) != 1 || f[0] != FeatureScopeOfficerViews {
		t.Fatalf("features = %v, want {officer_views}", f)
	}
	if err := s.Grant(ctx, Subject{GuildID: i64(gid)}, PlanGuild, &until, nil, "n", "a", nil); err != nil {
		t.Fatal(err)
	}
	if f := readFeatures(); f != nil {
		t.Fatalf("features = %v, want null after a whole-plan re-grant", f)
	}
}

func TestClaimedGuildsWithoutActivePlan(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	bare, planned, lapsed := seedGuild(t, pool), seedGuild(t, pool), seedGuild(t, pool)
	seedGuild(t, pool) // unclaimed
	for _, gid := range []int64{bare, planned, lapsed} {
		if _, err := pool.Exec(ctx, `update guilds set claimed_by = $2 where id = $1`, gid, seedUser(t, pool)); err != nil {
			t.Fatal(err)
		}
	}
	seedEntitlement(t, pool, Subject{GuildID: i64(planned)}, PlanGuild, "active")
	seedEntitlement(t, pool, Subject{GuildID: i64(lapsed)}, PlanGuild, "canceled")
	got, err := s.ClaimedGuildsWithoutActivePlan(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0].ID != bare || got[1].ID != lapsed || got[0].Name == "" {
		t.Fatalf("got %+v, want [%d %d]", got, bare, lapsed)
	}
}

func TestStripeSubscriptionIDsListsOnlyNonCanceledStripeSourcedRows(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	uid := seedUser(t, pool)
	end := time.Now().Add(time.Hour)
	if _, err := s.UpsertStripe(context.Background(), StripeUpsert{
		UserID: i64(uid), Plan: PlanPremium, Status: "active", CurrentPeriodEnd: &end,
		StripeSubscriptionID: "sub_reconcile_active", BillingUserID: uid, Actor: "test",
	}); err != nil {
		t.Fatal(err)
	}
	uid2 := seedUser(t, pool)
	if _, err := s.UpsertStripe(context.Background(), StripeUpsert{
		UserID: i64(uid2), Plan: PlanPremium, Status: "active", CurrentPeriodEnd: &end,
		StripeSubscriptionID: "sub_reconcile_canceled", BillingUserID: uid2, Actor: "test",
	}); err != nil {
		t.Fatal(err)
	}
	if err := s.SetCanceled(context.Background(), "sub_reconcile_canceled", nil, "test"); err != nil {
		t.Fatal(err)
	}
	// A grant (source = 'grant', no stripe_subscription_id) must never appear.
	uid3 := seedUser(t, pool)
	if err := s.Grant(context.Background(), Subject{UserID: i64(uid3)}, PlanPremium, nil, nil, "n", "a", nil); err != nil {
		t.Fatal(err)
	}

	ids, err := s.StripeSubscriptionIDs(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(ids) != 1 || ids[0] != "sub_reconcile_active" {
		t.Fatalf("ids = %v, want [sub_reconcile_active]", ids)
	}
}

func guildRowSubscription(t *testing.T, pool *pgxpool.Pool, gid int64) (subID string, cancelAtEnd bool) {
	t.Helper()
	if err := pool.QueryRow(context.Background(),
		`select stripe_subscription_id, cancel_at_period_end from entitlements where guild_id = $1 and plan = 'guild'`,
		gid).Scan(&subID, &cancelAtEnd); err != nil {
		t.Fatal(err)
	}
	return subID, cancelAtEnd
}

func countRows(t *testing.T, pool *pgxpool.Pool, table string, gid int64) int {
	t.Helper()
	var n int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from `+table+` where guild_id = $1`, gid).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

func upsertGuildSub(t *testing.T, s *Store, gid, billing int64, subID, status string, created time.Time, transfer bool) UpsertResult {
	t.Helper()
	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	res, err := s.UpsertStripe(context.Background(), StripeUpsert{
		GuildID: i64(gid), Plan: PlanGuild, Tier: TierGuild, Status: status, CurrentPeriodEnd: &end,
		StripeSubscriptionID: subID, BillingUserID: billing, Actor: "stripe_webhook:evt_" + subID,
		Transfer: transfer, SubscriptionCreated: created,
	})
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func TestUpsertStripeStaleTransferLeavesTheNewerRowUntouched(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool)
	officer := seedUser(t, pool)
	upsertGuildSub(t, s, gid, officer, "sub_c", "active", time.Now().Add(-time.Hour), true)

	end := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	res, err := s.UpsertStripe(context.Background(), StripeUpsert{
		GuildID: i64(gid), Plan: PlanGuild, Tier: TierGuild, Status: "active", CurrentPeriodEnd: &end, CancelAtPeriodEnd: true,
		StripeSubscriptionID: "sub_b", BillingUserID: officer, Actor: "stripe_webhook:evt_b",
		Transfer: true, SubscriptionCreated: time.Now().Add(-2 * time.Hour),
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.Outcome != OutcomeStaleTransfer || res.ExistingSubscriptionID != "sub_c" {
		t.Fatalf("result = %+v, want OutcomeStaleTransfer / sub_c", res)
	}
	if subID, cancel := guildRowSubscription(t, pool, gid); subID != "sub_c" || cancel {
		t.Fatalf("row = %q cancel=%v, want sub_c untouched", subID, cancel)
	}
	if n := countRows(t, pool, "entitlement_anomalies", gid); n != 0 {
		t.Fatalf("anomaly rows = %d, want 0", n)
	}
	if n := countRows(t, pool, "entitlement_audit", gid); n != 1 {
		t.Fatalf("audit rows = %d, want 1", n)
	}
}

func TestUpsertStripeTransferWinsOverARowWithUnknownCreated(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool)
	officer := seedUser(t, pool)
	upsertGuildSub(t, s, gid, officer, "sub_old", "active", time.Time{}, false)

	res := upsertGuildSub(t, s, gid, officer, "sub_new", "active", time.Now(), true)
	if res.Outcome != OutcomeWritten {
		t.Fatalf("result = %+v, want OutcomeWritten", res)
	}
	if subID, _ := guildRowSubscription(t, pool, gid); subID != "sub_new" {
		t.Fatalf("row points at %q, want sub_new", subID)
	}
	var notNull bool
	if err := pool.QueryRow(context.Background(),
		`select stripe_subscription_created is not null from entitlements where guild_id = $1`, gid).Scan(&notNull); err != nil || !notNull {
		t.Fatalf("stripe_subscription_created not null = %v, %v", notNull, err)
	}
}

func TestUpsertStripeEqualCreatedIsNotNewer(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool)
	officer := seedUser(t, pool)
	created := time.Now().Add(-time.Hour).UTC().Truncate(time.Second)
	upsertGuildSub(t, s, gid, officer, "sub_a", "active", created, true)

	res := upsertGuildSub(t, s, gid, officer, "sub_b", "active", created, true)
	if res.Outcome != OutcomeStaleTransfer {
		t.Fatalf("result = %+v, want OutcomeStaleTransfer", res)
	}
}

func TestUpsertStripeDeadNewcomerIsIgnoredWithoutAnAnomaly(t *testing.T) {
	for _, status := range []string{"canceled", "incomplete", "incomplete_expired", "unpaid"} {
		t.Run(status, func(t *testing.T) {
			pool := testPool(t)
			s := &Store{Pool: pool}
			gid := seedGuild(t, pool)
			officer := seedUser(t, pool)
			upsertGuildSub(t, s, gid, officer, "sub_a", "active", time.Now().Add(-time.Hour), false)

			res := upsertGuildSub(t, s, gid, officer, "sub_x", status, time.Now(), false)
			if res.Outcome != OutcomeDeadNewcomer || res.ExistingSubscriptionID != "sub_a" {
				t.Fatalf("result = %+v, want OutcomeDeadNewcomer / sub_a", res)
			}
			if subID, _ := guildRowSubscription(t, pool, gid); subID != "sub_a" {
				t.Fatalf("row points at %q, want sub_a unchanged", subID)
			}
			if n := countRows(t, pool, "entitlement_anomalies", gid); n != 0 {
				t.Fatalf("anomaly rows = %d, want 0", n)
			}
		})
	}
}

func TestUpsertStripeStoresTheSubscriptionCreated(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool)
	officer := seedUser(t, pool)
	created := time.Now().Add(-time.Hour).UTC().Truncate(time.Second)
	upsertGuildSub(t, s, gid, officer, "sub_a", "active", created, false)

	var got time.Time
	if err := pool.QueryRow(context.Background(),
		`select stripe_subscription_created from entitlements where guild_id = $1`, gid).Scan(&got); err != nil {
		t.Fatal(err)
	}
	if !got.Equal(created) {
		t.Fatalf("stripe_subscription_created = %v, want %v", got, created)
	}

	upsertGuildSub(t, s, gid, officer, "sub_a", "active", time.Time{}, false)
	if err := pool.QueryRow(context.Background(),
		`select stripe_subscription_created from entitlements where guild_id = $1`, gid).Scan(&got); err != nil {
		t.Fatal(err)
	}
	if !got.Equal(created) {
		t.Fatalf("after a zero-created renewal stripe_subscription_created = %v, want %v kept", got, created)
	}
}

func TestUpsertStripeNewSubscriptionWithUnknownCreatedClearsAStoredCreated(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool)
	officer := seedUser(t, pool)
	upsertGuildSub(t, s, gid, officer, "sub_a", "canceled", time.Now().Add(-time.Hour), false)

	upsertGuildSub(t, s, gid, officer, "sub_b", "active", time.Time{}, false)
	var notNull bool
	if err := pool.QueryRow(context.Background(),
		`select stripe_subscription_created is not null from entitlements where guild_id = $1`, gid).Scan(&notNull); err != nil {
		t.Fatal(err)
	}
	if notNull {
		t.Fatal("stripe_subscription_created kept sub_a's value for sub_b, want NULL")
	}
}

func TestRecordAnomalyIsIdempotentPerKindAndSubscription(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool)
	record := func(kind AnomalyKind) {
		t.Helper()
		if err := s.RecordAnomaly(context.Background(), nil, i64(gid), PlanGuild, kind, "sub_z", "test"); err != nil {
			t.Fatal(err)
		}
	}
	record(AnomalyDuplicateSubscription)
	record(AnomalyDuplicateSubscription)
	if n := countRows(t, pool, "entitlement_anomalies", gid); n != 1 {
		t.Fatalf("anomaly rows after repeat = %d, want 1", n)
	}
	record(AnomalyOrphanSubscription)
	if n := countRows(t, pool, "entitlement_anomalies", gid); n != 2 {
		t.Fatalf("anomaly rows after a different kind = %d, want 2", n)
	}
}

func seedEntitlementUntil(t *testing.T, pool *pgxpool.Pool, subj Subject, plan, status string, features []string, until *time.Time) {
	t.Helper()
	if _, err := pool.Exec(context.Background(),
		`insert into entitlements (user_id, guild_id, plan, source, status, features, current_period_end)
		 values ($1, $2, $3, 'grant', $4, $5, $6)`,
		subj.UserID, subj.GuildID, plan, status, features, until); err != nil {
		t.Fatal(err)
	}
}

func seedStripeEntitlement(t *testing.T, pool *pgxpool.Pool, subj Subject, plan, status string, periodEnd time.Time) {
	t.Helper()
	if _, err := pool.Exec(context.Background(),
		`insert into entitlements (user_id, guild_id, plan, source, status, stripe_subscription_id, current_period_end)
		 values ($1, $2, $3, 'stripe', $4, $5, $6)`,
		subj.UserID, subj.GuildID, plan, status, "sub_"+time.Now().Format("150405.000000000"), periodEnd); err != nil {
		t.Fatal(err)
	}
}

func timePtr(v time.Time) *time.Time { return &v }

func yesterday() time.Time { return time.Now().Add(-24 * time.Hour) }
func tomorrow() time.Time  { return time.Now().Add(24 * time.Hour) }

func TestExpiredGrantIsOff(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	uid := seedUser(t, pool)
	seedEntitlementUntil(t, pool, Subject{UserID: i64(uid)}, PlanPremium, "active", nil, timePtr(yesterday()))
	ok, reason, err := s.Can(context.Background(), uid, FeatureServerSims)
	if err != nil {
		t.Fatal(err)
	}
	if ok || reason != ReasonNoPlan {
		t.Fatalf("Can = %v, %v, want false, %v", ok, reason, ReasonNoPlan)
	}
	if sup, err := s.IsSupporter(context.Background(), uid); err != nil || sup {
		t.Fatalf("IsSupporter = %v, %v, want false", sup, err)
	}
}

func TestPastDueStripeRowStaysOn(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool)
	seedStripeEntitlement(t, pool, Subject{GuildID: i64(gid)}, PlanGuild, "past_due", yesterday())
	officer := seedUser(t, pool)
	seedGuildMember(t, pool, gid, officer, "officer", true)
	ok, reason, err := s.Can(context.Background(), officer, FeatureOfficerViews)
	if err != nil || !ok {
		t.Fatalf("Can = %v, %v, %v, want true", ok, reason, err)
	}
}

func TestScopedGrantGivesMembersNothing(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool)
	seedEntitlementUntil(t, pool, Subject{GuildID: i64(gid)}, PlanGuild, "active", []string{FeatureScopeOfficerViews}, timePtr(tomorrow()))
	member := seedUser(t, pool)
	seedGuildMember(t, pool, gid, member, "member", true)
	officer := seedUser(t, pool)
	seedGuildMember(t, pool, gid, officer, "officer", true)

	if ok, _, err := s.Can(context.Background(), member, FeatureServerSims); err != nil || ok {
		t.Fatalf("member Can(server_sims) = %v, %v, want false", ok, err)
	}
	if sup, err := s.IsSupporter(context.Background(), member); err != nil || sup {
		t.Fatalf("member IsSupporter = %v, %v, want false", sup, err)
	}
	if ok, _, err := s.Can(context.Background(), officer, FeatureOfficerViews); err != nil || !ok {
		t.Fatalf("officer Can(officer_views) = %v, %v, want true", ok, err)
	}
}

func TestCanForGuild(t *testing.T) {
	scoped := []string{FeatureScopeOfficerViews}
	for _, tc := range []struct {
		name       string
		until      time.Time
		rank       string
		verified   bool
		otherGuild bool
		wholePlan  bool
		feature    Feature
		wantOK     bool
		wantReason Reason
		wantErr    bool
	}{
		{name: "officer scoped grant", until: tomorrow(), rank: "officer", verified: true, feature: FeatureOfficerViews, wantOK: true, wantReason: ReasonEntitled},
		{name: "whole plan active row", until: tomorrow(), rank: "officer", verified: true, feature: FeatureOfficerViews, wantOK: true, wantReason: ReasonEntitled, wholePlan: true},
		{name: "officer expired grant", until: yesterday(), rank: "officer", verified: true, feature: FeatureOfficerViews, wantReason: ReasonNoPlan},
		{name: "member with grant", until: tomorrow(), rank: "member", verified: true, feature: FeatureOfficerViews, wantReason: ReasonNotOfficer},
		{name: "officer of another guild", until: tomorrow(), rank: "officer", verified: true, otherGuild: true, feature: FeatureOfficerViews, wantReason: ReasonNoPlan},
		{name: "unverified officer", until: tomorrow(), rank: "officer", verified: false, feature: FeatureOfficerViews, wantReason: ReasonNoPlan},
		{name: "other feature", until: tomorrow(), rank: "officer", verified: true, feature: FeatureHistory, wantErr: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pool := testPool(t)
			s := &Store{Pool: pool}
			gid := seedGuild(t, pool)
			features := scoped
			if tc.wholePlan {
				features = nil
			}
			seedEntitlementUntil(t, pool, Subject{GuildID: i64(gid)}, PlanGuild, "active", features, timePtr(tc.until))
			uid := seedUser(t, pool)
			memberOf := gid
			if tc.otherGuild {
				memberOf = seedGuild(t, pool)
			}
			seedGuildMember(t, pool, memberOf, uid, tc.rank, tc.verified)
			ok, reason, err := s.CanForGuild(context.Background(), gid, uid, tc.feature)
			if (err != nil) != tc.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tc.wantErr)
			}
			if ok != tc.wantOK || reason != tc.wantReason {
				t.Fatalf("CanForGuild = %v, %q, want %v, %q", ok, reason, tc.wantOK, tc.wantReason)
			}
		})
	}
}

func TestGuildBillingCarriesFeaturesAndActive(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	live, ended := seedGuild(t, pool), seedGuild(t, pool)
	scoped := []string{FeatureScopeOfficerViews}
	seedEntitlementUntil(t, pool, Subject{GuildID: i64(live)}, PlanGuild, "active", scoped, timePtr(tomorrow()))
	seedEntitlementUntil(t, pool, Subject{GuildID: i64(ended)}, PlanGuild, "active", scoped, timePtr(yesterday()))

	g, err := s.GuildBilling(context.Background(), live)
	if err != nil || g == nil {
		t.Fatalf("GuildBilling live = %v, %v", g, err)
	}
	if len(g.Features) != 1 || g.Features[0] != FeatureScopeOfficerViews || !g.Scoped() || !g.Active {
		t.Fatalf("live = %+v, want scoped and active", g)
	}
	g, err = s.GuildBilling(context.Background(), ended)
	if err != nil || g == nil {
		t.Fatalf("GuildBilling ended = %v, %v", g, err)
	}
	if g.Active {
		t.Fatalf("ended = %+v, want inactive", g)
	}
}

func TestUpsertStripeReplacesScopedGrant(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool)
	uid := seedUser(t, pool)
	seedEntitlementUntil(t, pool, Subject{GuildID: i64(gid)}, PlanGuild, "active",
		[]string{FeatureScopeOfficerViews}, timePtr(tomorrow()))
	end := tomorrow().UTC().Truncate(time.Second)
	if _, err := s.UpsertStripe(context.Background(), StripeUpsert{
		GuildID: i64(gid), Plan: PlanGuild, Tier: TierGuild, Status: "active", CurrentPeriodEnd: &end,
		StripeSubscriptionID: "sub_scoped", BillingUserID: uid, Actor: "stripe_webhook:evt_s",
	}); err != nil {
		t.Fatal(err)
	}
	var source string
	var features []string
	if err := pool.QueryRow(context.Background(),
		`select source, features from entitlements where guild_id = $1 and plan = 'guild'`, gid).Scan(&source, &features); err != nil {
		t.Fatal(err)
	}
	if source != "stripe" || features != nil {
		t.Fatalf("source = %q, features = %v, want stripe and null", source, features)
	}
	var oldStatus, newStatus string
	if err := pool.QueryRow(context.Background(),
		`select old_status, new_status from entitlement_audit where guild_id = $1 order by id desc limit 1`, gid).Scan(&oldStatus, &newStatus); err != nil {
		t.Fatal(err)
	}
	if oldStatus != "active" || newStatus != "active" {
		t.Fatalf("audit = %q -> %q, want active -> active", oldStatus, newStatus)
	}
	member := seedUser(t, pool)
	seedGuildMember(t, pool, gid, member, "member", true)
	if ok, _, err := s.Can(context.Background(), member, FeatureServerSims); err != nil || !ok {
		t.Fatalf("member Can(server_sims) = %v, %v, want true on the whole plan", ok, err)
	}
	if sup, err := s.IsSupporter(context.Background(), member); err != nil || !sup {
		t.Fatalf("member IsSupporter = %v, %v, want true on the whole plan", sup, err)
	}
}

// guildRowState is the part of a guild row a scoped Grant may or may not touch.
type guildRowState struct {
	Source            string
	Features          []string
	BillingUserID     *int64
	CancelAtPeriodEnd bool
	SubscriptionID    *string
	GraceUntil        *time.Time
}

func readGuildRowState(t *testing.T, pool *pgxpool.Pool, gid int64) guildRowState {
	t.Helper()
	var r guildRowState
	if err := pool.QueryRow(context.Background(), `
		select source, features, billing_user_id, cancel_at_period_end, stripe_subscription_id, grace_until
		from entitlements where guild_id = $1 and plan = 'guild'`, gid).
		Scan(&r.Source, &r.Features, &r.BillingUserID, &r.CancelAtPeriodEnd, &r.SubscriptionID, &r.GraceUntil); err != nil {
		t.Fatal(err)
	}
	return r
}

// seedFullStripeRow writes a Stripe guild row with every Stripe-only column set.
func seedFullStripeRow(t *testing.T, pool *pgxpool.Pool, gid int64, status string) {
	t.Helper()
	if _, err := pool.Exec(context.Background(), `
		insert into entitlements (guild_id, tier, plan, source, status, current_period_end, billing_user_id,
		                          cancel_at_period_end, stripe_subscription_id, grace_until)
		values ($1, 'raid_guild', 'guild', 'stripe', $2, now() + interval '10 days', $3, true, $4, now() + interval '30 days')`,
		gid, status, seedUser(t, pool), fmt.Sprintf("sub_full_%d", gid)); err != nil {
		t.Fatal(err)
	}
}

func TestScopedGrantNeverDowngradesALiveRow(t *testing.T) {
	scoped := []string{FeatureScopeOfficerViews}
	for _, tc := range []struct {
		name       string
		seed       func(t *testing.T, pool *pgxpool.Pool, gid int64)
		wantRefuse bool
	}{
		{"active stripe row", func(t *testing.T, pool *pgxpool.Pool, gid int64) { seedFullStripeRow(t, pool, gid, "active") }, true},
		{"active whole-plan grant", func(t *testing.T, pool *pgxpool.Pool, gid int64) {
			seedEntitlementUntil(t, pool, Subject{GuildID: i64(gid)}, PlanGuild, "active", nil, timePtr(tomorrow()))
		}, true},
		{"expired whole-plan grant", func(t *testing.T, pool *pgxpool.Pool, gid int64) {
			seedEntitlementUntil(t, pool, Subject{GuildID: i64(gid)}, PlanGuild, "active", nil, timePtr(yesterday()))
		}, false},
		{"canceled stripe row", func(t *testing.T, pool *pgxpool.Pool, gid int64) { seedFullStripeRow(t, pool, gid, "canceled") }, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pool := testPool(t)
			s := &Store{Pool: pool}
			gid := seedGuild(t, pool)
			tc.seed(t, pool, gid)
			before := readGuildRowState(t, pool, gid)
			until := tomorrow().UTC().Truncate(time.Second)
			err := s.Grant(context.Background(), Subject{GuildID: i64(gid)}, PlanGuild, &until, nil, "n", "a", scoped)
			after := readGuildRowState(t, pool, gid)
			if tc.wantRefuse {
				if !errors.Is(err, ErrWouldDowngrade) {
					t.Fatalf("err = %v, want ErrWouldDowngrade", err)
				}
				if !reflect.DeepEqual(before, after) {
					t.Fatalf("row changed: %+v -> %+v", before, after)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			want := guildRowState{Source: "grant", Features: scoped}
			if !reflect.DeepEqual(after, want) {
				t.Fatalf("row = %+v, want %+v (Stripe columns cleared)", after, want)
			}
		})
	}
}
