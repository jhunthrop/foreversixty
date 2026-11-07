package account

import (
	"context"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	stripe "github.com/stripe/stripe-go/v82"

	"github.com/jhunthrop/foreversixty/api/internal/billing"
	"github.com/jhunthrop/foreversixty/api/internal/db"
)

// testPool is a migrated, emptied database (TEST_DATABASE_URL: a throwaway local database,
// never production).
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
		`truncate users, guilds, reports, fight_metrics, rating_scores, characters, sims, activity_events,
		 builds, login_tokens, uploads cascade`); err != nil {
		t.Fatal(err)
	}
	if err := db.EnsureMetricsPartition(context.Background(), pool, foughtAt); err != nil {
		t.Fatal(err)
	}
	if err := db.EnsureRatingsPartition(context.Background(), pool, foughtAt); err != nil {
		t.Fatal(err)
	}
	return pool
}

// foughtAt is every seeded fight's time.
var foughtAt = time.Date(2026, 11, 5, 20, 0, 0, 0, time.UTC)

// exec runs one seeding statement.
func exec(t *testing.T, pool *pgxpool.Pool, sql string, args ...any) {
	t.Helper()
	if _, err := pool.Exec(context.Background(), sql, args...); err != nil {
		t.Fatalf("%s: %v", sql, err)
	}
}

// count answers a count(*) query.
func count(t *testing.T, pool *pgxpool.Pool, sql string, args ...any) int {
	t.Helper()
	var n int
	if err := pool.QueryRow(context.Background(), sql, args...).Scan(&n); err != nil {
		t.Fatalf("%s: %v", sql, err)
	}
	return n
}

// seedUser makes an account with a BattleTag and an email.
func seedUser(t *testing.T, pool *pgxpool.Pool, battletag, role string) int64 {
	t.Helper()
	var id int64
	if err := pool.QueryRow(context.Background(),
		`insert into users (bnet_sub, battletag, email, role) values ($1, $2, $3, $4) returning id`,
		"sub-"+battletag, battletag, battletag+"@example.com", role).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

// seedGuild makes a guild, claimed by claimant when it is not zero.
func seedGuild(t *testing.T, pool *pgxpool.Pool, name string, claimant int64) int64 {
	t.Helper()
	var claimedBy *int64
	if claimant != 0 {
		claimedBy = &claimant
	}
	var id int64
	if err := pool.QueryRow(context.Background(),
		`insert into guilds (region, ruleset, name, claimed_by) values ('us', 'normal', $1, $2) returning id`,
		name, claimedBy).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

// seedGuildPlan gives the guild a Stripe Guild plan billed to payer.
func seedGuildPlan(t *testing.T, pool *pgxpool.Pool, guildID, payer int64, subID string, endingAtPeriodEnd bool) {
	t.Helper()
	exec(t, pool, `insert into entitlements (guild_id, plan, tier, source, status, stripe_subscription_id, billing_user_id, cancel_at_period_end)
		values ($1, 'guild', 'guild', 'stripe', 'active', $2, $3, $4)`, guildID, subID, payer, endingAtPeriodEnd)
}

// seedPremium gives the account its own Stripe Premium.
func seedPremium(t *testing.T, pool *pgxpool.Pool, userID int64, subID string) {
	t.Helper()
	exec(t, pool, `insert into entitlements (user_id, plan, source, status, stripe_subscription_id, billing_user_id)
		values ($1, 'premium', 'stripe', 'active', $2, $1)`, userID, subID)
	exec(t, pool, `insert into stripe_customers (user_id, stripe_customer_id) values ($1, $2)`, userID, "cus_"+subID)
}

// seedReport makes a report owned by owner with one fight, one raw chunk and one metric row
// for playerKey.
func seedReport(t *testing.T, pool *pgxpool.Pool, id string, owner int64, visibility string, guildID *int64, playerKey string) {
	t.Helper()
	exec(t, pool, `insert into reports (id, owner_id, guild_id, visibility, status) values ($1, $2, $3, $4, 'complete')`,
		id, owner, guildID, visibility)
	exec(t, pool, `insert into fights (report_id, fight_index) values ($1, 0)`, id)
	exec(t, pool, `insert into raw_chunks (report_id, start_offset, end_offset, sha256) values ($1, 4096, 8192, '\x00')`, id)
	exec(t, pool, `insert into fight_metrics (report_id, fight_index, player_key, fought_at) values ($1, 0, $2, $3)`,
		id, playerKey, foughtAt)
	exec(t, pool, `insert into rating_scores (report_id, fight_index, player_key, kill, components, model_version, fought_at)
		values ($1, 0, $2, true, '{}', 'v1', $3)`, id, playerKey, foughtAt)
	exec(t, pool, `insert into uploads (id, user_id, object_key, r2_upload_id, size_bytes, report_id)
		values ($1, $2, $3, 'mp', 1, $4)`, "up-"+id, owner, "uploads/up-"+id+"/raw.txt.zst", id)
}

// seedCharacter links a character to owner; source is 'export' or 'bnet'.
func seedCharacter(t *testing.T, pool *pgxpool.Pool, key string, owner int64, source string, bnetID *int64) {
	t.Helper()
	exec(t, pool, `insert into characters (key, region, ruleset, name, user_id, source, bnet_character_id, bnet_profile)
		values ($1, 'us', 'normal', $1, $2, $3, $4, case when $4::bigint is null then null else '{"level":60}'::jsonb end)`,
		key, owner, source, bnetID)
	exec(t, pool, `insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at)
		values ($1, $2, 'us', 'normal', $1, 'FS1:x', now())`, key, owner)
	exec(t, pool, `insert into character_syncs (character_key, source, outcome) values ($1, 'addon', 'ok')`, key)
	exec(t, pool, `insert into character_syncs (character_key, source, outcome) values ($1, 'blizzard', 'ok')`, key)
}

// fakeRevoker records the accounts whose Discord role it was asked to revoke.
type fakeRevoker struct {
	mu      sync.Mutex
	revoked []int64
	err     error
}

func (f *fakeRevoker) Revoke(_ context.Context, userID int64) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.revoked = append(f.revoked, userID)
	return f.err
}

// fakeObjects records deleted object keys.
type fakeObjects struct {
	mu      sync.Mutex
	deleted []string
}

func (f *fakeObjects) Delete(_ context.Context, key string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.deleted = append(f.deleted, key)
	return nil
}

// rig is a Service over the test database with Stripe's fake gateway (never live keys).
type rig struct {
	pool    *pgxpool.Pool
	svc     *Service
	gateway *billing.FakeGateway
	discord *fakeRevoker
	objects *fakeObjects
}

func newRig(t *testing.T) rig {
	t.Helper()
	pool := testPool(t)
	gateway := &billing.FakeGateway{Subscriptions: map[string]*stripe.Subscription{}}
	r := rig{pool: pool, gateway: gateway, discord: &fakeRevoker{}, objects: &fakeObjects{}}
	r.svc = &Service{
		Store:   &Store{Pool: pool},
		Billing: &billing.Service{Store: &billing.Store{Pool: pool}, Gateway: gateway},
		Discord: r.discord, Objects: r.objects,
		Now: func() time.Time { return time.Date(2026, 11, 7, 21, 41, 0, 0, time.UTC) },
	}
	return r
}

// liveSubscription registers an active subscription at the fake Stripe.
func (r rig) liveSubscription(id string) {
	r.gateway.Subscriptions[id] = &stripe.Subscription{ID: id, Status: stripe.SubscriptionStatusActive}
}

func userExists(t *testing.T, pool *pgxpool.Pool, id int64) bool {
	t.Helper()
	return count(t, pool, `select count(*) from users where id = $1`, id) == 1
}

func idText(id int64) string { return fmt.Sprint(id) }
