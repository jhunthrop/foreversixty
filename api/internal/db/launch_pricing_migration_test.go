package db

import (
	"context"
	"testing"
)

const launchPricingVersion = 54

// launchPricingTables counts the five tables 0054 adds.
func launchPricingTables(t *testing.T, url string) int {
	t.Helper()
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	var n int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from information_schema.tables
		  where table_name in ('premium_seats','guild_trials','plan_transfers','seat_comebacks','billing_notices')`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

// 0054 puts every existing guild row on the raid tier (grants and scoped grants keep the
// loot council), leaves premium rows tierless and keeps the owner's open-ended grant.
func TestLaunchPricingMigration_TierBackfill(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	migrateTo(t, url, launchPricingVersion-1)

	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	ctx := context.Background()
	if _, err := pool.Exec(ctx, `truncate users, guilds, entitlements cascade`); err != nil {
		t.Fatal(err)
	}
	var uid, gid int64
	if err := pool.QueryRow(ctx, `insert into users (email) values ('owner@example.com') returning id`).Scan(&uid); err != nil {
		t.Fatal(err)
	}
	if err := pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us','normal','G') returning id`).Scan(&gid); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into entitlements (guild_id, plan, source, status) values ($1,'guild','grant','active')`, gid); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into entitlements (user_id, plan, source, status) values ($1,'premium','grant','active')`, uid); err != nil {
		t.Fatal(err)
	}

	if err := Migrate(url); err != nil {
		t.Fatalf("migrating up: %v", err)
	}
	assertTiers := func(label string) {
		t.Helper()
		var guildTier, premiumTier *string
		var premiumEnd any
		var premiumStatus string
		if err := pool.QueryRow(ctx, `select tier from entitlements where plan='guild'`).Scan(&guildTier); err != nil {
			t.Fatal(err)
		}
		if err := pool.QueryRow(ctx,
			`select tier, status, current_period_end from entitlements where plan='premium'`).
			Scan(&premiumTier, &premiumStatus, &premiumEnd); err != nil {
			t.Fatal(err)
		}
		if guildTier == nil || *guildTier != "raid_guild" {
			t.Fatalf("%s: guild tier = %v, want raid_guild", label, guildTier)
		}
		if premiumTier != nil || premiumStatus != "active" || premiumEnd != nil {
			t.Fatalf("%s: premium row changed: tier %v status %s end %v", label, premiumTier, premiumStatus, premiumEnd)
		}
	}
	assertTiers("after up")
	if n := launchPricingTables(t, url); n != 5 {
		t.Fatalf("after up: %d tables, want 5", n)
	}

	migrateTo(t, url, launchPricingVersion-1)
	if n := launchPricingTables(t, url); n != 0 {
		t.Fatalf("after down: %d tables, want 0", n)
	}
	if err := Migrate(url); err != nil {
		t.Fatalf("migrating up again: %v", err)
	}
	assertTiers("after up again")
}
