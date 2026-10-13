// api/internal/guilds/flagged_readers_test.go
package guilds

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Task L6.12: a report flagged tampered or unverified counts nowhere a player or guild looks.
// Every guild reader of a night's rated or ranked rows picks its reports through
// presence.EvidenceReport. Each test seeds one member at two guild nights: a clean one two days
// ago and a flagged one yesterday, newer and better on every figure, and asks one reader for
// what it shows. The flagged night's row must not appear.

const l612Key = "us/hardcore/l612member"

type l612Guild struct {
	pool                   *pgxpool.Pool
	store                  *Store
	guildID, userID        int64
	cleanReport, flaggedID string
}

func seedL612Guild(t *testing.T) l612Guild {
	t.Helper()
	pool := testPool(t)
	g := l612Guild{pool: pool, store: &Store{Pool: pool}, cleanReport: "l612-clean", flaggedID: "l612-flagged"}
	g.guildID = seedGuild(t, pool, "L612 Guild")
	g.userID = seedUser(t, pool, "l612-member@example.com")
	seedCharacter(t, pool, g.guildID, g.userID, l612Key, "member", true)
	seedExport(t, pool, g.userID, l612Key, "us", "hardcore", "L612member")
	recomputeMembership(t, pool, g.guildID, g.userID)
	now := time.Now()
	seedL68Night(t, pool, g.guildID, l68Night{id: g.cleanReport, visibility: "public", at: now.Add(-48 * time.Hour),
		keys: []string{l612Key}, rated: []string{l612Key}})
	seedL68Night(t, pool, g.guildID, l68Night{id: g.flaggedID, visibility: "public", at: now.Add(-24 * time.Hour),
		keys: []string{l612Key}, rated: []string{l612Key}, flagged: "tampered"})
	ctx := context.Background()
	for _, q := range []struct {
		sql string
		id  string
	}{
		{`update fight_metrics set encounter_id = 1084, spec = 'Arms', ilvl = 50 where report_id = $1`, g.cleanReport},
		{`update fight_metrics set encounter_id = 1084, class = 'mage', spec = 'Frost', ilvl = 99, metric_dps = 999
		  where report_id = $1`, g.flaggedID},
		{`update rating_scores set overall = 99, components = '[{"name":"output","score":99}]' where report_id = $1`,
			g.flaggedID},
	} {
		if _, err := pool.Exec(ctx, q.sql, q.id); err != nil {
			t.Fatal(err)
		}
	}
	return g
}

func (g l612Guild) rosterRow(t *testing.T) RosterRow {
	t.Helper()
	roster, err := g.store.HomeRoster(context.Background(), g.guildID, Guild{}, g.userID, false, false)
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range roster {
		if row.CharacterKey == l612Key {
			return row
		}
	}
	t.Fatalf("%s is not on the roster: %+v", l612Key, roster)
	return RosterRow{}
}

// The roster's best parse (bestParses) is never a flagged report's.
func TestRosterBestParseIgnoresAFlaggedReport(t *testing.T) {
	g := seedL612Guild(t)
	row := g.rosterRow(t)
	if row.BestParse == nil || row.BestParse.ReportID != g.cleanReport || row.BestParse.Value != 100 {
		t.Fatalf("best_parse = %+v, want the clean report's 100 dps (the flagged 999 is no parse)", row.BestParse)
	}
}

// The roster's rating (latestRatings) is never a flagged report's card, even the newest.
func TestRosterRatingIgnoresAFlaggedReport(t *testing.T) {
	g := seedL612Guild(t)
	row := g.rosterRow(t)
	if row.Rating == nil || row.Rating.Overall != 50 || row.Rating.Output != 0 {
		t.Fatalf("rating = %+v, want the clean report's overall 50 (the flagged 99 is no card)", row.Rating)
	}
}

// The roster's class, spec and item level (HomeRoster's newest ranked row) are never read
// off a flagged report, even the newest.
func TestRosterClassSpecAndItemLevelIgnoreAFlaggedReport(t *testing.T) {
	g := seedL612Guild(t)
	row := g.rosterRow(t)
	if row.Spec == nil || *row.Spec != "Arms" || row.Class == nil || *row.Class != "warrior" ||
		row.ItemLevel == nil || *row.ItemLevel != 50 {
		t.Fatalf("class %v spec %v item level %v, want the clean report's warrior, Arms, 50",
			deref(row.Class), deref(row.Spec), deref(row.ItemLevel))
	}
}

// The Progression tab's best parse per role (bestByRole) is never a flagged report's.
func TestProgressionBestByRoleIgnoresAFlaggedReport(t *testing.T) {
	g := seedL612Guild(t)
	view, err := g.store.Progression(context.Background(), g.guildID)
	if err != nil {
		t.Fatal(err)
	}
	if len(view.Encounters) != 1 {
		t.Fatalf("encounters = %+v", view.Encounters)
	}
	dps := view.Encounters[0].BestByRole["dps"]
	if dps == nil || dps.ReportID != g.cleanReport || dps.Value != 100 {
		t.Fatalf("best dps = %+v, want the clean report's 100 (the flagged 999 is no parse)", dps)
	}
}

// The Raids tab's top parse (fillTopParse) is never a flagged report's: the flagged night is
// still listed (visibility decides that), with no top parse, as it lists nobody present.
func TestRaidsTopParseIgnoresAFlaggedReport(t *testing.T) {
	g := seedL612Guild(t)
	page, err := g.store.Raids(context.Background(), g.guildID, g.userID, true, RaidsDefaultLimit, nil)
	if err != nil {
		t.Fatal(err)
	}
	seen := map[string]bool{}
	for _, row := range page.Rows {
		seen[row.ID] = true
		switch row.ID {
		case g.flaggedID:
			if row.TopParse != nil {
				t.Errorf("the flagged night's top_parse = %+v, want none", *row.TopParse)
			}
		case g.cleanReport:
			if row.TopParse == nil || row.TopParse.Value != 100 {
				t.Errorf("the clean night's top_parse = %+v, want its 100 dps", row.TopParse)
			}
		}
	}
	if !seen[g.flaggedID] || !seen[g.cleanReport] {
		t.Fatalf("raids rows = %+v, want both nights listed", page.Rows)
	}
}

func deref[T any](p *T) any {
	if p == nil {
		return nil
	}
	return *p
}
