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
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
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
	seedPlanRow(t, pool, guildID, source, "active", "30 days", features)
}

// seedPlanRow writes a guild row ending at now() + offset (a Postgres interval).
func seedPlanRow(t *testing.T, pool *pgxpool.Pool, guildID int64, source, status, offset, features string) {
	t.Helper()
	if _, err := pool.Exec(context.Background(), `
		insert into entitlements (guild_id, plan, source, status, current_period_end, features)
		values ($1, 'guild', $2, $3, now() + $4::interval, nullif($5, '')::text[])`,
		guildID, source, status, offset, features); err != nil {
		t.Fatal(err)
	}
}

// testGrandfatherUntil stays in the future for the life of these tests, so
// the grants the command writes stay live between runs.
var testGrandfatherUntil = time.Date(2099, 12, 11, 0, 0, 0, 0, time.UTC)

func testGrandfatherOptions(apply bool) grandfatherOptions {
	return grandfatherOptions{Until: testGrandfatherUntil, Note: "launch", Operator: "tester", Apply: apply}
}

func runGrandfather(t *testing.T, pool *pgxpool.Pool, apply bool) string {
	t.Helper()
	var out bytes.Buffer
	opts := testGrandfatherOptions(apply)
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
	lapsed := seedClaimedGuild(t, pool, "Lapsed", true)
	seedPlanRow(t, pool, lapsed, "grant", "active", "-1 day", "")
	churned := seedClaimedGuild(t, pool, "Churned", true)
	seedPlanRow(t, pool, churned, "stripe", "canceled", "-1 day", "")
	todo := map[int64]string{bare: "Bare", lapsed: "Lapsed", churned: "Churned"}

	dry := runGrandfather(t, pool, false)
	wantDry := []string{
		"skipped " + itoa(scoped) + " Scoped: already granted",
		"skipped " + itoa(paid) + " Paid: on the plan",
		"3 would be granted, 2 skipped",
	}
	for id, name := range todo {
		wantDry = append(wantDry, "would grant "+itoa(id)+" "+name+" until 2099-12-11")
	}
	for _, want := range wantDry {
		if !strings.Contains(dry, want) {
			t.Fatalf("dry run output missing %q:\n%s", want, dry)
		}
	}
	if _, rows := featuresOf(t, pool, bare); rows != 0 {
		t.Fatal("a dry run must write nothing")
	}

	applied := runGrandfather(t, pool, true)
	if !strings.Contains(applied, "3 granted, 2 skipped") {
		t.Fatalf("apply output:\n%s", applied)
	}
	for id, name := range todo {
		if !strings.Contains(applied, "granted "+itoa(id)+" "+name+" until 2099-12-11") {
			t.Fatalf("apply output missing %s:\n%s", name, applied)
		}
		if f, rows := featuresOf(t, pool, id); rows != 1 || len(f) != 1 || f[0] != "officer_views" {
			t.Fatalf("%s: features %v, rows %d", name, f, rows)
		}
	}
	if _, rows := featuresOf(t, pool, loose); rows != 0 {
		t.Fatal("an unclaimed guild must not be granted")
	}

	again := runGrandfather(t, pool, true)
	if !strings.Contains(again, "0 granted, 5 skipped") {
		t.Fatalf("second apply output:\n%s", again)
	}
}

// TestGrandfatherSkipsAGuildThatGainedAPlan covers the window between the
// list read and the write: a guild listed as having no plan whose Stripe row
// is live by the time Grant runs is reported as skipped, and left alone.
func TestGrandfatherSkipsAGuildThatGainedAPlan(t *testing.T) {
	pool := grandfatherPool(t)
	paid := seedClaimedGuild(t, pool, "Paid", true)
	seedPlan(t, pool, paid, "stripe", "")
	store := &entitlements.Store{Pool: pool}
	var out bytes.Buffer
	granted, skipped, err := grantAll(context.Background(), store,
		[]entitlements.GuildRef{{ID: paid, Name: "Paid"}}, testGrandfatherOptions(true), &out)
	if err != nil {
		t.Fatal(err)
	}
	if granted != 0 || skipped != 1 || !strings.Contains(out.String(), "skipped "+itoa(paid)+" Paid: on the plan") {
		t.Fatalf("granted %d, skipped %d, output:\n%s", granted, skipped, out.String())
	}
	if f, _ := featuresOf(t, pool, paid); f != nil {
		t.Fatalf("the paid row was scoped: %v", f)
	}
}

func TestParseGrandfatherFlagsRefusesAPastUntil(t *testing.T) {
	_, err := parseGrandfatherFlags([]string{"--until", "2020-01-01", "--note", "x"})
	if err == nil || err.Error() != "--until must be a future date" {
		t.Fatalf("err = %v, want --until must be a future date", err)
	}
	if _, err := parseGrandfatherFlags([]string{"--until", "2099-12-11", "--note", "x"}); err != nil {
		t.Fatal(err)
	}
}
