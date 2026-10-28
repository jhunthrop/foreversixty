// api/internal/guilds/standing_queries_test.go
package guilds

import (
	"context"
	"fmt"
	"os"
	"sync/atomic"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/billing"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

// statementCounter counts every statement a pool runs.
type statementCounter struct{ n atomic.Int64 }

func (c *statementCounter) TraceQueryStart(ctx context.Context, _ *pgx.Conn, _ pgx.TraceQueryStartData) context.Context {
	c.n.Add(1)
	return ctx
}

func (c *statementCounter) TraceQueryEnd(context.Context, *pgx.Conn, pgx.TraceQueryEndData) {}

// countedStore is a guild store on its own pool whose statements counter counts, with the
// subscribe and covered reads wired as production wires them.
func countedStore(t *testing.T) (*Store, *statementCounter) {
	t.Helper()
	cfg, err := pgxpool.ParseConfig(os.Getenv("TEST_DATABASE_URL"))
	if err != nil {
		t.Fatal(err)
	}
	counter := &statementCounter{}
	cfg.ConnConfig.Tracer = counter
	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	return &Store{Pool: pool, Billing: &billing.Store{Pool: pool}, Plans: &entitlements.Store{Pool: pool}}, counter
}

// standingStatements is how many statements GuildStanding runs for a verified officer of n
// claimed guilds, each on a Stripe plan paid by another account, so both name reads run.
func standingStatements(t *testing.T, n int) int64 {
	t.Helper()
	h := newHTTPHarness(t)
	viewer := seedUser(t, h.pool, "counted-viewer@example.com")
	ids := make([]int64, n)
	for i := range ids {
		gid := seedGuild(t, h.pool, fmt.Sprintf("Counted %d", i))
		payer := seedRanked(t, h, gid, fmt.Sprintf("Payer%d", i), "leader", true)
		seedCharacter(t, h.pool, gid, viewer, fmt.Sprintf("us/hardcore/Viewer%d", i), "officer", true)
		syncMembership(t, h.pool, gid, viewer)
		setGuild(t, h, `update guilds set claimed_by = $2, claimed_at = now() where id = $1`, gid, payer)
		seedStripePlan(t, h, gid, "guild", "active", daysFromNow(20), false, &payer)
		ids[i] = gid
	}
	store, counter := countedStore(t)
	got, err := store.GuildStanding(context.Background(), viewer, ids)
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range ids {
		if got[id].Covered == nil || got[id].Covered.BilledBy == "" || got[id].Claim.HolderName == nil {
			t.Fatalf("guild %d: %+v, want covered with billed_by and a holder name (both name reads run)", id, got[id])
		}
	}
	return counter.n.Load()
}

// TestGuildStanding_QueryCountFixed: GuildStanding's reads do not grow with the guild count;
// twelve fixed plus the two name reads (me_standing.go's GuildStanding comment).
func TestGuildStanding_QueryCountFixed(t *testing.T) {
	const want = 14
	one, three := standingStatements(t, 1), standingStatements(t, 3)
	if one != want || three != want {
		t.Fatalf("GuildStanding ran %d statements for 1 guild and %d for 3, want %d for both", one, three, want)
	}
}
