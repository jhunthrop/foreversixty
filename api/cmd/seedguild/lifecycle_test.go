// api/cmd/seedguild/lifecycle_test.go
//
// The local walkthrough world (lifecycle.go): its cast is pure data and checked without a
// database; applying it twice and removing it is checked against the test database, which
// needs TEST_DATABASE_URL like the other integration tests here.
package main

import (
	"context"
	"strings"
	"testing"
)

func TestLifecycleCastHasOneMasterOneOfficerAndThreeMembers(t *testing.T) {
	var master, officer, members, outside int
	accounts := map[string]bool{}
	for _, acct := range lifecycleAccounts() {
		if accounts[acct.Battletag] || acct.Sub == "" {
			t.Fatalf("account %q is duplicated or has no bnet sub", acct.Battletag)
		}
		accounts[acct.Battletag] = true
		for _, c := range acct.Characters {
			switch {
			case !c.InGuild:
				outside++
			case c.RankIndex == 0:
				master++
			case c.RankIndex == 1:
				officer++
			default:
				members++
			}
		}
	}
	if master != 1 || officer != 1 || members != 3 || outside != 2 {
		t.Fatalf("master=%d officer=%d members=%d outside=%d, want 1/1/3/2", master, officer, members, outside)
	}
	if len(accounts) != 4 {
		t.Fatalf("%d accounts, want 4", len(accounts))
	}
}

func TestLifecycleExportCarriesTheGuildSectionOnlyForMembers(t *testing.T) {
	for _, acct := range lifecycleAccounts() {
		for _, c := range acct.Characters {
			export, err := lifecycleExport(c)
			if err != nil {
				t.Fatalf("%s: %v", c.Mock.Name, err)
			}
			has := strings.Contains(export, "|guild=Lifecycle%20Test:")
			if has != c.InGuild {
				t.Errorf("%s: guild section present = %v, want %v (%s)", c.Mock.Name, has, c.InGuild, export)
			}
			if !strings.HasPrefix(export, "FS1:") || !strings.Contains(export, "|level=60") {
				t.Errorf("%s: not a level-60 FS1 export: %s", c.Mock.Name, export)
			}
		}
	}
}

func TestLifecycleReportHasTwoBossKills(t *testing.T) {
	kills := lifecycleReport().kills()
	if kills != 2 {
		t.Fatalf("report has %d boss kills, want 2", kills)
	}
}

func TestApplyLifecycleThenRemoveRoundTrips(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	count := func(query string, args ...any) int {
		t.Helper()
		var n int
		if err := pool.QueryRow(ctx, query, args...).Scan(&n); err != nil {
			t.Fatal(err)
		}
		return n
	}

	if err := ApplyLifecycle(ctx, pool); err != nil {
		t.Fatal(err)
	}
	// A second run must not duplicate anything.
	if err := ApplyLifecycle(ctx, pool); err != nil {
		t.Fatal(err)
	}

	if n := count(`select count(*) from users`); n != 4 {
		t.Errorf("%d users, want 4", n)
	}
	if n := count(`select count(*) from users where bnet_sub is not null`); n != 4 {
		t.Errorf("%d users with a bnet sub, want 4 (the guild master claim needs one)", n)
	}
	if n := count(`select count(*) from characters where level = 60`); n != 7 {
		t.Errorf("%d level-60 characters, want 7", n)
	}
	if n := count(`select count(*) from addon_exports`); n != 7 {
		t.Errorf("%d exports, want one per character (7)", n)
	}
	if n := count(`select count(*) from guilds where lower(name) = 'lifecycle test' and claimed_by is null and claim_pending_by is null`); n != 1 {
		t.Errorf("%d unclaimed Lifecycle Test guilds, want 1", n)
	}
	if n := count(`select count(*) from guild_characters where rank = 'leader' and character_key = 'us/pvp/bow-jackzon'`); n != 1 {
		t.Errorf("Bow Jackzon is not the guild's leader")
	}
	if n := count(`select count(*) from guild_characters where verified_at is not null`); n != 1 {
		t.Errorf("%d rows verified, want only the guild master's (as a Battle.net import leaves it)", n)
	}
	if n := count(`select count(*) from guild_members where rank = 'leader' and verified_at is not null`); n != 1 {
		t.Errorf("the guild master is not a verified leader in guild_members, so the claim page would refuse them")
	}
	if n := count(`select count(*) from guild_characters where rank = 'officer'`); n != 1 {
		t.Errorf("%d officers, want 1", n)
	}
	if n := count(`select jsonb_array_length(members) from guild_roster_snapshots`); n != 7 {
		t.Errorf("roster snapshot has %d lines, want 7", n)
	}
	if n := count(`select count(*) from reports where status = 'complete' and owner_id = (select id from users where battletag = 'hunthrop#1894')`); n != 1 {
		t.Errorf("%d complete reports owned by the guild master, want 1", n)
	}
	if n := count(`select count(*) from fights where kill and encounter_id is not null`); n != 2 {
		t.Errorf("%d boss kills, want 2", n)
	}
	if n := count(`select count(*) from fight_metrics where player_key = 'us/pvp/bow-jackzon' and kill and encounter_id is not null`); n != 2 {
		t.Errorf("%d boss kills credited to the guild master, want 2", n)
	}

	if err := RemoveLifecycle(ctx, pool); err != nil {
		t.Fatal(err)
	}
	for _, table := range []string{"users", "characters", "addon_exports", "guild_characters", "guild_members",
		"guild_roster_snapshots", "reports", "fights", "fight_metrics", "seed_rows"} {
		if n := count(`select count(*) from ` + table); n != 0 {
			t.Errorf("%s still has %d rows after --remove", table, n)
		}
	}
	if n := count(`select count(*) from guilds`); n != 0 {
		t.Errorf("%d guilds left after --remove", n)
	}
}
