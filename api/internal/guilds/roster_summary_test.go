// api/internal/guilds/roster_summary_test.go
package guilds

import (
	"context"
	"testing"
)

func rated(overalls ...float64) []RosterRow {
	rows := make([]RosterRow, 0, len(overalls))
	for _, o := range overalls {
		rows = append(rows, RosterRow{Rating: &RatingView{Overall: o}})
	}
	return rows
}

func TestBelowRatingFloor(t *testing.T) {
	cases := []struct {
		name string
		rows []RosterRow
		want int
	}{
		{"two rated raiders are too few for a floor", rated(40, 90), 0},
		{"seven rated raiders are still too few", rated(10, 50, 60, 70, 80, 90, 95), 0},
		{"eight raiders within the band flag nobody", rated(80, 81, 82, 82, 83, 84, 85, 85), 0},
		{"eight spread raiders flag the two under the 25th percentile", rated(30, 40, 60, 65, 70, 75, 80, 90), 2},
		{"a tie at the bottom is not below itself", rated(50, 50, 50, 70, 75, 80, 85, 90), 0},
		{"unrated rows never count", append(rated(30, 40, 60, 65, 70, 75, 80, 90), RosterRow{}), 2},
	}
	for _, c := range cases {
		if got := belowRatingFloor(c.rows); got != c.want {
			t.Errorf("%s: belowRatingFloor = %d, want %d", c.name, got, c.want)
		}
	}
}

func TestSummaryCountsRaidersByAccountAndWaitingByCharacter(t *testing.T) {
	gear := map[string]int{"chest": 100}
	row := func(owner int64, verified bool, consent string) RosterRow {
		return RosterRow{ownerUserID: owner, Verified: verified, Consent: consent, gear: gear}
	}
	roster := []RosterRow{
		// Shares gear, but the export carried no gear list: gear unknown, not "with gear".
		{ownerUserID: 5, Verified: true, Consent: "gear"},
		// One account, three characters, all waiting: three waiting, as the panel lists them.
		row(1, false, "gear"), row(1, false, "gear"), row(1, false, "gear"),
		// A main and an alt with gear on the site: one raider.
		row(2, true, "gear"), row(2, true, "gear"),
		// Roster-only consent: approved, but no gear on the site.
		row(3, true, "roster"),
		row(4, true, "gear_bags"),
	}
	raiders, waiting := summaryCounts(roster)
	if raiders != 2 || waiting != 3 {
		t.Fatalf("summaryCounts = %d raiders, %d waiting; want 2 and 3", raiders, waiting)
	}
}

// TestEncountersDownAgreesWithProgression: a kill of a boss outside this tier's named list never
// counts as an encounter down, and the home's figure is the Progression tab's; with no logs at
// all both read zero.
func TestEncountersDownAgreesWithProgression(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")

	agree := func(wantDown, wantPulls int) {
		t.Helper()
		sv, err := h.store.homeSummary(ctx, gid, nil)
		if err != nil {
			t.Fatal(err)
		}
		prog, err := h.store.Progression(ctx, gid, true)
		if err != nil {
			t.Fatal(err)
		}
		if sv.Tier.Down != wantDown || prog.Tier.Down != wantDown || sv.NamedEncountersDown != wantDown {
			t.Fatalf("home %d (%d), progression %d encounters down; want %d", sv.Tier.Down, sv.NamedEncountersDown, prog.Tier.Down, wantDown)
		}
		if sv.Tier.PullsTotal != wantPulls || prog.Tier.PullsTotal != wantPulls || sv.PullsThisTier != wantPulls {
			t.Fatalf("home %d (%d), progression %d pulls; want %d", sv.Tier.PullsTotal, sv.PullsThisTier, prog.Tier.PullsTotal, wantPulls)
		}
	}
	agree(0, 0)

	if _, err := h.pool.Exec(ctx,
		`insert into reports (id, guild_id, visibility, status, zone, created_at)
		 values ('down-night', $1, 'guild', 'complete', 'Molten Core', now())`, gid); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx,
		`insert into fights (report_id, fight_index, name, encounter_id, kill) values ('down-night', 0, 'Lucifron', 663, true)`); err != nil {
		t.Fatal(err)
	}
	agree(0, 0)

	if _, err := h.pool.Exec(ctx,
		`insert into fights (report_id, fight_index, name, encounter_id, kill) values ('down-night', 1, 'Onyxia', $1, true)`,
		testOnyxiaEncounter); err != nil {
		t.Fatal(err)
	}
	agree(1, 1)
}
