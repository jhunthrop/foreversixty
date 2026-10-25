// api/cmd/api/grandfather_test.go
package main

import (
	"bytes"
	"context"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

func grandfatherPool(t *testing.T) *pgxpool.Pool {
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

func seedClaimedGuild(t *testing.T, pool *pgxpool.Pool, name string, claimed bool) int64 {
	t.Helper()
	ctx := context.Background()
	var id int64
	if err := pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'hardcore', $1) returning id`, name).Scan(&id); err != nil {
		t.Fatal(err)
	}
	if claimed {
		var uid int64
		if err := pool.QueryRow(ctx, `insert into users (email) values ($1) returning id`, name+"@example.com").Scan(&uid); err != nil {
			t.Fatal(err)
		}
		if _, err := pool.Exec(ctx, `update guilds set claimed_by = $2 where id = $1`, id, uid); err != nil {
			t.Fatal(err)
		}
	}
	return id
}

func itoa(n int64) string { return strconv.FormatInt(n, 10) }

func seedPlan(t *testing.T, pool *pgxpool.Pool, guildID int64, source, features string) {
	t.Helper()
	if _, err := pool.Exec(context.Background(), `
		insert into entitlements (guild_id, plan, source, status, current_period_end, features)
		values ($1, 'guild', $2, 'active', now() + interval '30 days', nullif($3, '')::text[])`,
		guildID, source, features); err != nil {
		t.Fatal(err)
	}
}

func runGrandfather(t *testing.T, pool *pgxpool.Pool, apply bool) string {
	t.Helper()
	var out bytes.Buffer
	opts := grandfatherOptions{
		Until: time.Date(2026, 12, 9, 0, 0, 0, 0, time.UTC), Note: "launch", Operator: "tester", Apply: apply,
	}
	if err := grandfatherOfficerTools(context.Background(), pool, opts, &out); err != nil {
		t.Fatal(err)
	}
	return out.String()
}

func featuresOf(t *testing.T, pool *pgxpool.Pool, guildID int64) (features []string, rows int) {
	t.Helper()
	if err := pool.QueryRow(context.Background(),
		`select count(*) from entitlements where guild_id = $1`, guildID).Scan(&rows); err != nil {
		t.Fatal(err)
	}
	if rows > 0 {
		if err := pool.QueryRow(context.Background(),
			`select features from entitlements where guild_id = $1`, guildID).Scan(&features); err != nil {
			t.Fatal(err)
		}
	}
	return features, rows
}

func TestGrandfatherOfficerTools(t *testing.T) {
	pool := grandfatherPool(t)
	bare := seedClaimedGuild(t, pool, "Bare", true)
	scoped := seedClaimedGuild(t, pool, "Scoped", true)
	seedPlan(t, pool, scoped, "grant", "{officer_views}")
	paid := seedClaimedGuild(t, pool, "Paid", true)
	seedPlan(t, pool, paid, "stripe", "")
	loose := seedClaimedGuild(t, pool, "Loose", false)

	dry := runGrandfather(t, pool, false)
	for _, want := range []string{
		"would grant " + itoa(bare) + " Bare until 2026-12-09",
		"skipped " + itoa(scoped) + " Scoped: already granted",
		"skipped " + itoa(paid) + " Paid: on the plan",
		"1 would be granted, 2 skipped",
	} {
		if !strings.Contains(dry, want) {
			t.Fatalf("dry run output missing %q:\n%s", want, dry)
		}
	}
	if _, rows := featuresOf(t, pool, bare); rows != 0 {
		t.Fatal("a dry run must write nothing")
	}

	applied := runGrandfather(t, pool, true)
	if !strings.Contains(applied, "granted "+itoa(bare)+" Bare until 2026-12-09") || !strings.Contains(applied, "1 granted, 2 skipped") {
		t.Fatalf("apply output:\n%s", applied)
	}
	if f, rows := featuresOf(t, pool, bare); rows != 1 || len(f) != 1 || f[0] != "officer_views" {
		t.Fatalf("bare guild: features %v, rows %d", f, rows)
	}
	if _, rows := featuresOf(t, pool, loose); rows != 0 {
		t.Fatal("an unclaimed guild must not be granted")
	}

	again := runGrandfather(t, pool, true)
	if !strings.Contains(again, "0 granted, 3 skipped") {
		t.Fatalf("second apply output:\n%s", again)
	}
}
