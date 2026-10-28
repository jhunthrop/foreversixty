package coach

import (
	"context"
	"errors"
	"math"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/phase"
	"github.com/jhunthrop/foreversixty/api/internal/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const (
	writerEncounterID = 667
	writerDPS         = 1234.5
	writerFightMS     = 180_000
)

var writerFoughtAt = time.Date(2026, 12, 9, 1, 0, 0, 0, time.UTC)

type fakeTables struct{ err error }

func (f fakeTables) TablesFor(context.Context, summary.RosterRow, int64, time.Time) (Tables, error) {
	return Tables{}, f.err
}

// failAfter runs the inner hook, then fails: the rating and the profiles must both go.
type failAfter struct{ inner rating.FightHook }

func (h failAfter) InTx(ctx context.Context, tx pgx.Tx, f rating.RatedFight, cards []rating.CardRow) error {
	if err := h.inner.InTx(ctx, tx, f, cards); err != nil {
		return err
	}
	return errors.New("hook failed after the profiles")
}

func writerPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set")
	}
	if err := db.Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := db.Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	return pool
}

// seedReport inserts a report with the visibility and flag, removed again at cleanup.
func seedReport(t *testing.T, pool *pgxpool.Pool, id, visibility string, flagged *string) {
	t.Helper()
	ctx := context.Background()
	if _, err := pool.Exec(ctx,
		`insert into reports (id, visibility, status, flagged) values ($1, $2, 'complete', $3)`,
		id, visibility, flagged); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from coach_profiles where report_id = $1`, id)
		pool.Exec(ctx, `delete from rating_scores where report_id = $1`, id)
		pool.Exec(ctx, `delete from reports where id = $1`, id)
	})
}

func writerFight(id string, kill bool) rating.RatedFight {
	return rating.RatedFight{
		ReportID: id, FightIndex: 1, Region: "us", Ruleset: "normal", FoughtAt: writerFoughtAt,
		EncounterID: writerEncounterID,
		Summary: summary.Summary{
			EncounterID: writerEncounterID, Difficulty: 3, DurationMS: writerFightMS, Kill: kill,
			Roster: []summary.RosterRow{
				{GUID: "g1", Name: "Profwar", Class: "Warrior", Spec: "Fury", Role: RoleDPS, DPS: writerDPS},
				{GUID: "g2", Name: "Profpriest", Class: "Priest", Spec: "Holy", Role: RoleHealer},
				{GUID: "g3", Name: "Profnospec", Class: "Mage", Role: RoleDPS},
			},
		},
	}
}

func countProfiles(t *testing.T, pool *pgxpool.Pool, id string) int {
	t.Helper()
	var n int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from coach_profiles where report_id = $1`, id).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

func rateWith(t *testing.T, pool *pgxpool.Pool, hook rating.FightHook, f rating.RatedFight) error {
	t.Helper()
	return (&rating.Store{Pool: pool, Hook: hook}).RateFight(context.Background(), f)
}

func TestProfileWriter_KillOnRankedEvidenceReport(t *testing.T) {
	pool := writerPool(t)
	seedReport(t, pool, "prof-kill", "public", nil)
	f := writerFight("prof-kill", true)
	if err := rateWith(t, pool, &ProfileWriter{Tables: fakeTables{}}, f); err != nil {
		t.Fatal(err)
	}
	if n := countProfiles(t, pool, f.ReportID); n != 2 {
		t.Fatalf("%d profiles, want 2 (the row with no spec writes none)", n)
	}
	var warMetric, priestMetric, overall float64
	var ph, version string
	ctx := context.Background()
	if err := pool.QueryRow(ctx, `select metric, phase, model_version from coach_profiles
	  where report_id = $1 and spec = 'Fury'`, f.ReportID).Scan(&warMetric, &ph, &version); err != nil {
		t.Fatal(err)
	}
	if err := pool.QueryRow(ctx, `select p.metric, r.overall from coach_profiles p
	  join rating_scores r on r.report_id = p.report_id and r.fight_index = p.fight_index and r.player_key = p.player_key
	  where p.report_id = $1 and p.spec = 'Holy'`, f.ReportID).Scan(&priestMetric, &overall); err != nil {
		t.Fatal(err)
	}
	if warMetric != writerDPS {
		t.Errorf("warrior metric %v, want its DPS %v", warMetric, writerDPS)
	}
	if math.Abs(priestMetric-overall) > 1e-9 {
		t.Errorf("priest metric %v, want the card's overall %v", priestMetric, overall)
	}
	if ph != phase.At(writerFoughtAt) || version != RulesVersion {
		t.Errorf("phase %q version %q, want %q %q", ph, version, phase.At(writerFoughtAt), RulesVersion)
	}
}

func TestProfileWriter_WipePrivateAndFlaggedWriteNothing(t *testing.T) {
	pool := writerPool(t)
	tampered := "tampered"
	cases := []struct {
		name, visibility string
		flagged          *string
		kill             bool
	}{
		{"wipe", "public", nil, false},
		{"private", "private", nil, true},
		{"flagged", "public", &tampered, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			id := "prof-none-" + tc.name
			seedReport(t, pool, id, tc.visibility, tc.flagged)
			if err := rateWith(t, pool, &ProfileWriter{Tables: fakeTables{}}, writerFight(id, tc.kill)); err != nil {
				t.Fatal(err)
			}
			if n := countProfiles(t, pool, id); n != 0 {
				t.Fatalf("%d profiles written, want 0", n)
			}
		})
	}
}

func TestProfileWriter_RewriteWithdrawsEarlierProfiles(t *testing.T) {
	pool := writerPool(t)
	seedReport(t, pool, "prof-rewrite", "public", nil)
	w := &ProfileWriter{Tables: fakeTables{}}
	if err := rateWith(t, pool, w, writerFight("prof-rewrite", true)); err != nil {
		t.Fatal(err)
	}
	if err := rateWith(t, pool, w, writerFight("prof-rewrite", false)); err != nil {
		t.Fatal(err)
	}
	if n := countProfiles(t, pool, "prof-rewrite"); n != 0 {
		t.Fatalf("%d profiles survive the fight becoming a wipe", n)
	}
}

func TestProfileWriter_ErrorRollsTheRatingBack(t *testing.T) {
	pool := writerPool(t)
	seedReport(t, pool, "prof-fail", "public", nil)
	hook := failAfter{inner: &ProfileWriter{Tables: fakeTables{}}}
	if err := rateWith(t, pool, hook, writerFight("prof-fail", true)); err == nil {
		t.Fatal("RateFight returned nil for a failing hook")
	}
	var scores int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from rating_scores where report_id = 'prof-fail'`).Scan(&scores); err != nil {
		t.Fatal(err)
	}
	if scores != 0 || countProfiles(t, pool, "prof-fail") != 0 {
		t.Fatalf("%d rating rows and %d profiles survived the rollback", scores, countProfiles(t, pool, "prof-fail"))
	}
}

func TestProfileWriter_MissingTablesSkipTheRowNotTheRating(t *testing.T) {
	pool := writerPool(t)
	seedReport(t, pool, "prof-tables", "public", nil)
	w := &ProfileWriter{Tables: fakeTables{err: errors.New("no tables")}}
	if err := rateWith(t, pool, w, writerFight("prof-tables", true)); err != nil {
		t.Fatal(err)
	}
	if n := countProfiles(t, pool, "prof-tables"); n != 0 {
		t.Fatalf("%d profiles written without tables", n)
	}
}
