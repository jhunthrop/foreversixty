// api/internal/guilds/progression_test.go
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/db"
)

// TestProgressionCountsOnyxiaAndUnnamedZones seeds two Onyxia pulls (one wipe, one kill)
// and a Barrow Deeps night with no named encounter, and checks the tier bar, Onyxia's own
// depth and the unnamed zone's pull count.
func TestProgressionCountsOnyxiaAndUnnamedZones(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")

	if err := db.EnsureMetricsPartitions(ctx, h.pool, time.Now(), 0); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx,
		`insert into reports (id, guild_id, visibility, status, zone, created_at)
		 values ('prog-onyxia', $1, 'public', 'complete', 'Onyxia''s Lair', now())`, gid); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx,
		`insert into fights (report_id, fight_index, name, encounter_id, kill, duration_ms, start_ms, deaths)
		 values ('prog-onyxia', 0, 'Onyxia', 1084, false, 200000, 1000, 2)`); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx,
		`insert into fights (report_id, fight_index, name, encounter_id, kill, duration_ms, start_ms, deaths, players)
		 values ('prog-onyxia', 1, 'Onyxia', 1084, true, 260000, 201000, 1, $1)`, []string{"us/hardcore/progtank"}); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx,
		`insert into fight_metrics (report_id, fight_index, player_key, player_name, class, spec, role, metric_dps, kill, encounter_id, fought_at)
		 values ('prog-onyxia', 1, 'us/hardcore/progtank', 'Progtank', 'warrior', 'Protection', 'tank', 42.0, true, 1084, now())`); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx,
		`insert into reports (id, guild_id, visibility, status, zone, created_at)
		 values ('prog-barrow', $1, 'public', 'complete', 'Barrow Deeps', now())`, gid); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx,
		`insert into fights (report_id, fight_index, name, kill, duration_ms, start_ms)
		 values ('prog-barrow', 0, 'Pull 1', true, 180000, 1000)`); err != nil {
		t.Fatal(err)
	}

	h.actor = auth.Actor{}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/progression", gid), "")
	var view ProgressionView
	h.data(res, &view)

	if view.Tier.Name != "First tier" || view.Tier.Down != 1 || view.Tier.NamedEncounters != 1 {
		t.Fatalf("tier = %+v, want First tier/down=1/named=1", view.Tier)
	}
	if len(view.Tier.Raids) != 3 || view.Tier.Raids[2] != "Onyxia's Lair" {
		t.Fatalf("tier.raids = %v, want the 3 real raids ending in Onyxia's Lair", view.Tier.Raids)
	}
	if len(view.Encounters) != 1 {
		t.Fatalf("encounters = %+v, want exactly Onyxia", view.Encounters)
	}
	onyxia := view.Encounters[0]
	if onyxia.Pulls != 2 || onyxia.Kills != 1 {
		t.Errorf("onyxia pulls/kills = %d/%d, want 2/1", onyxia.Pulls, onyxia.Kills)
	}
	if onyxia.BestKillMS == nil || *onyxia.BestKillMS != 260000 {
		t.Errorf("bestKillMS = %v, want 260000", onyxia.BestKillMS)
	}
	if onyxia.FirstKillAt == nil {
		t.Error("firstKillAt = nil, want set")
	}
	if onyxia.DeathsPerPull != 1.5 {
		t.Errorf("deathsPerPull = %v, want 1.5 ((2+1)/2)", onyxia.DeathsPerPull)
	}
	if len(onyxia.PullsByNight) != 1 || onyxia.PullsByNight[0].Pulls != 2 || !onyxia.PullsByNight[0].Killed {
		t.Fatalf("pullsByNight = %+v, want one night, 2 pulls, killed", onyxia.PullsByNight)
	}
	tank := onyxia.BestByRole["tank"]
	if tank == nil || tank.Name != "Progtank" || tank.Metric != "dps" || tank.Value != 42.0 {
		t.Errorf("best tank = %+v, want Progtank dps 42.0", tank)
	}

	if len(view.Unnamed) != 2 {
		t.Fatalf("unnamed = %+v, want 2 zones", view.Unnamed)
	}
	var barrow *UnnamedZoneProgress
	for i := range view.Unnamed {
		if view.Unnamed[i].Zone == "Barrow Deeps" {
			barrow = &view.Unnamed[i]
		}
	}
	if barrow == nil || barrow.Pulls != 1 || barrow.Nights != 1 {
		t.Fatalf("barrow deeps = %+v, want 1 pull, 1 night", barrow)
	}
}

// TestProgressionWithNoDataReadsAllZero checks a guild with no reports at all still
// answers a complete, zero-valued shape rather than an error.
func TestProgressionWithNoDataReadsAllZero(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	h.actor = auth.Actor{}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/progression", gid), "")
	var view ProgressionView
	h.data(res, &view)
	if len(view.Encounters) != 1 || view.Encounters[0].Pulls != 0 || view.Tier.Down != 0 {
		t.Fatalf("encounters = %+v, want one zero-pull Onyxia row and tier.down 0", view.Encounters)
	}
	if len(view.Unnamed) != 2 || view.Unnamed[0].Pulls != 0 {
		t.Fatalf("unnamed = %+v, want 2 zero-pull zones", view.Unnamed)
	}
}
