package activity

import (
	"context"
	"path/filepath"
	"testing"
	"time"
)

func fightMetric(t *testing.T, s *Store, report string, idx, enc int, key, name string, dps float64) {
	t.Helper()
	exec(t, s.Pool, `insert into fight_metrics (report_id, fight_index, player_key, player_name, class, spec, role, metric_dps,
		encounter_id, kill, fought_at) values ($1, $2, $3, $4, 'Warrior', 'Fury', 'dps', $5, $6, true, $7)`,
		report, idx, key, name, dps, enc, testNow.Add(-time.Second))
}

// A board ranks each character on its average percentile over the band's instances; the
// second refresh tells the climb and the new #1, and a private report never counts.
func TestRefreshBoardsRanksAndTellsMovement(t *testing.T) {
	s := testStore(t)
	u := user(t, s, "a#1", false)
	report(t, s, "r1", u, "public", "Shadowfang Keep", "")
	report(t, s, "rp", u, "private", "Shadowfang Keep", "")
	bands := map[string]int{"Shadowfang Keep": 20}
	fightMetric(t, s, "r1", 0, 100, "k/a", "Alpha", 50)
	fightMetric(t, s, "r1", 0, 100, "k/b", "Bravo", 40)
	fightMetric(t, s, "rp", 0, 100, "k/p", "Private", 99)
	if _, err := s.RefreshBoards(context.Background(), bands, testNow); err != nil {
		t.Fatal(err)
	}
	board, err := s.boardWeek(context.Background(), "warrior-fury", 20, WeekOf(testNow))
	if err != nil || len(board) != 2 || board[0].Name != "Alpha" {
		t.Fatalf("first board: %+v %v", board, err)
	}
	fightMetric(t, s, "r1", 1, 101, "k/b", "Bravo", 80)
	fightMetric(t, s, "r1", 1, 101, "k/a", "Alpha", 10)
	fightMetric(t, s, "r1", 2, 100, "k/b", "Bravo", 60)
	n, err := s.RefreshBoards(context.Background(), bands, testNow)
	if err != nil || n != 1 {
		t.Fatalf("second refresh wrote %d rows: %v", n, err)
	}
	got := rowTexts(t, s)
	if len(got) != 1 || got[0] != "Bravo took #1 for Fury Warrior, 20 to 29" {
		t.Fatalf("rows: %q", got)
	}
}

// One pass over the repository's own data: the BiS files and the curated files load, and
// every damage spec of the active build gets a grade.
func TestLoopPassOverTheRepositoryData(t *testing.T) {
	s := testStore(t)
	root := filepath.Join("..", "..", "..")
	curated, err := LoadCurated(filepath.Join(root, "data", "curated"))
	if err != nil {
		t.Skip("no data/curated: ", err)
	}
	if curated.ZoneBands["Shadowfang Keep"] != 20 || !curated.RaidZones["Onyxia's Lair"] {
		t.Fatalf("curated: %+v", curated)
	}
	l := &Loop{Store: s, DataDir: filepath.Join(root, "data", "builds"), Build: activeBuild(t, root), Curated: curated, Now: func() time.Time { return testNow }}
	if err := l.Pass(context.Background()); err != nil {
		t.Fatal(err)
	}
	var graded int
	if err := s.Pool.QueryRow(context.Background(), `select count(*) from tier_grade_weeks where grade is not null and band = 20`).Scan(&graded); err != nil || graded < 10 {
		t.Fatalf("graded %d specs at 20: %v", graded, err)
	}
}

func activeBuild(t *testing.T, root string) string {
	t.Helper()
	var f struct {
		Build string `json:"build"`
	}
	if err := readJSON(filepath.Join(root, "web", "src", "data", "active-build.json"), &f); err != nil {
		t.Skip(err)
	}
	return f.Build
}
