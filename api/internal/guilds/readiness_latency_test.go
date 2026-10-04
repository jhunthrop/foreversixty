// api/internal/guilds/readiness_latency_test.go
//
// Measures GET /v1/guilds/{id}/readiness against 25 raiders on the local harness - the
// task brief's own "keep it under 300ms for 25 raiders on the harness (measure, print in
// the report)". Not a pass/fail assertion against a hard budget (a loaded CI runner's own
// Postgres container is not a production latency SLA), but it fails loudly if the endpoint
// regresses to a clearly non-constant-per-row cost (N+1 query explosion) by comparing
// against a generous ceiling.
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

const readinessLatencyCeiling = 2 * time.Second

func TestReadinessLatencyWith25Raiders(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")

	classes := []string{"warrior", "paladin", "hunter", "rogue", "priest", "shaman", "mage", "warlock", "druid"}
	var viewerID int64
	for i := 0; i < 25; i++ {
		uid := seedUser(t, h.pool, fmt.Sprintf("readiness-latency-%d@example.com", i))
		key := fmt.Sprintf("us/hardcore/readinesslatency%d", i)
		seedCharacter(t, h.pool, gid, uid, key, "member", true)
		export := fmt.Sprintf(
			`FS1:1.60.1.70009:%s:human:555555555/33/0:head=12640,chest=11726:41,wrist=1234:724|professions=skinning,leatherworking|bags=13510,13444|level=60`,
			classes[i%len(classes)])
		if _, err := h.pool.Exec(ctx, `
			insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at, updated_at)
			values ($1, $2, 'us', 'hardcore', $3, $4, now(), now())`,
			key, uid, fmt.Sprintf("Latency%d", i), export); err != nil {
			t.Fatal(err)
		}
		recomputeMembership(t, h.pool, gid, uid)
		if _, err := h.pool.Exec(ctx, `
			update guild_members set consent = 'gear_bags' where guild_id = $1 and user_id = $2`, gid, uid); err != nil {
			t.Fatal(err)
		}
		if _, err := h.pool.Exec(ctx, `
			insert into fight_metrics (report_id, fight_index, player_key, class, spec, role, ilvl, faction, fought_at)
			values ($1, 0, $2, $3, 'Fury', 'dps', 60, 'alliance', now())`,
			fmt.Sprintf("latency-report-%d", i), key, classes[i%len(classes)]); err != nil {
			t.Fatal(err)
		}
		if i == 0 {
			viewerID = uid
		}
	}

	h.actor = auth.Actor{UserID: viewerID, Role: "user", Method: "session"}
	start := time.Now()
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/readiness", gid), "")
	elapsed := time.Since(start)
	var view ReadinessView
	h.data(res, &view)

	t.Logf("GET /v1/guilds/%d/readiness with 25 raiders took %s (%d rows)", gid, elapsed, len(view.Rows))
	if len(view.Rows) != 25 {
		t.Fatalf("rows = %d, want 25", len(view.Rows))
	}
	if elapsed > readinessLatencyCeiling {
		t.Errorf("readiness latency = %s, want under %s", elapsed, readinessLatencyCeiling)
	}
}
