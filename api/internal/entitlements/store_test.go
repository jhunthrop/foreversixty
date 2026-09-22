// api/internal/entitlements/store_test.go
package entitlements

import (
	"context"
	"os"
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
		`insert into entitlements (user_id, guild_id, plan, source, status)
		 values ($1, $2, $3, 'grant', $4)`, subj.UserID, subj.GuildID, plan, status); err != nil {
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
		GuildID: i64(gid), Plan: PlanGuild, Status: "active", CurrentPeriodEnd: &end,
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
		GuildID: i64(gid), Plan: PlanGuild, Status: "active", CurrentPeriodEnd: &end,
		StripeSubscriptionID: "sub_a", BillingUserID: officerA, Actor: "stripe_webhook:evt_a",
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.Duplicate {
		t.Fatalf("first write must not be flagged a duplicate: %+v", res)
	}

	res, err = s.UpsertStripe(context.Background(), StripeUpsert{
		GuildID: i64(gid), Plan: PlanGuild, Status: "active", CurrentPeriodEnd: &end,
		StripeSubscriptionID: "sub_b", BillingUserID: officerB, Actor: "stripe_webhook:evt_b",
	})
	if err != nil {
		t.Fatal(err)
	}
	if !res.Duplicate || res.ExistingSubscriptionID != "sub_a" {
		t.Fatalf("second write = %+v, want Duplicate=true ExistingSubscriptionID=sub_a", res)
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
	if res.Duplicate {
		t.Fatalf("a renewal on the same subscription id must not be flagged a duplicate: %+v", res)
	}
	b, err := s.PersonalBilling(context.Background(), uid)
	if err != nil || b == nil || !b.CurrentPeriodEnd.Equal(end2) {
		t.Fatalf("after renewal: %+v, %v", b, err)
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
		&grantedBy, "beta tester", "cli_grant:tester beta tester"); err != nil {
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
	if err := s.Grant(context.Background(), Subject{UserID: i64(uid3)}, PlanPremium, nil, nil, "n", "a"); err != nil {
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
