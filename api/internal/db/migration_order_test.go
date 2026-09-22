// api/internal/db/migration_order_test.go
package db

import (
	"strconv"
	"strings"
	"testing"
)

// mergeBaseMigrations is the exact set of migration files present on main at this lane's
// merge base (confirmed via `ls api/internal/db/migrations/*.up.sql | sort` in this
// checkout before this lane wrote anything). golang-migrate only ever moves a database
// forward from its currently-applied version, so a migration numbered at or below the
// highest version already live is silently skipped the moment it deploys after a higher
// one - the exact mistake the coordinator's numbering ruling calls out by name. This test
// cannot see the future (it cannot know what version another in-flight lane will land at),
// so it enforces the one thing it can check mechanically and permanently: nothing this
// lane adds may reuse, or fall at or below, a number that was already on main before this
// lane started.
var mergeBaseMigrations = map[string]bool{
	"0001_subscribers.up.sql":                true,
	"0002_unsubscribe_token.up.sql":          true,
	"0003_confirmation_sent_at.up.sql":       true,
	"0004_builds.up.sql":                     true,
	"0005_logs.up.sql":                       true,
	"0006_boss_health.up.sql":                true,
	"0007_refold_digests.up.sql":             true,
	"0008_refold_after_engine.up.sql":        true,
	"0009_refold_after_engine_0_2_4.up.sql":  true,
	"0010_refold_after_engine_0_2_5.up.sql":  true,
	"0011_refold_after_engine_0_2_6.up.sql":  true,
	"0012_builds_point_order_integer.up.sql": true,
	"0013_sims.up.sql":                       true,
	"0014_sim_kinds.up.sql":                  true,
	"0015_builds_owner.up.sql":               true,
	"0016_sim_headline_backfill.up.sql":      true,
	"0017_reports_recent_idx.up.sql":         true,
	"0018_guild_membership.up.sql":           true,
}

// mergeBaseMax is the highest number in mergeBaseMigrations.
const mergeBaseMax = 18

func TestMigrationNumbersDoNotRegressBelowMergeBase(t *testing.T) {
	entries, err := migrations.ReadDir("migrations")
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		name := e.Name()
		if !strings.HasSuffix(name, ".up.sql") {
			continue
		}
		numStr, _, ok := strings.Cut(name, "_")
		if !ok {
			t.Fatalf("migration file %q has no <number>_ prefix", name)
		}
		n, err := strconv.Atoi(numStr)
		if err != nil {
			t.Fatalf("migration file %q has a non-numeric prefix: %v", name, err)
		}
		if n > mergeBaseMax {
			continue // this lane's own migration(s), or any later lane's: fine by construction
		}
		if !mergeBaseMigrations[name] {
			t.Errorf("migration %q is numbered %d, at or below the merge-base max of %d, "+
				"but is not one of the migrations already on main - renumber it above %d",
				name, n, mergeBaseMax, mergeBaseMax)
		}
	}
}

func TestRatingsMigrationIsNumberedAfterMergeBase(t *testing.T) {
	entries, err := migrations.ReadDir("migrations")
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		if e.Name() == "0022_ratings.up.sql" {
			return
		}
	}
	t.Fatal("expected api/internal/db/migrations/0022_ratings.up.sql per the coordinator's " +
		"numbering ruling (0019 is reserved by the guild-membership window, 0020/0021 by " +
		"the entitlements lane)")
}
