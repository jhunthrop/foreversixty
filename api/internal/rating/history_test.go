package rating

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

const (
	histKey       = "us/normal/histmain"
	histPrefix    = "hist-"
	histEncounter = 990701 // 990701..990702: this file's own encounter ids
)

var histBase = time.Date(2026, 3, 2, 20, 0, 0, 0, time.UTC)

func histPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	pool := testPool(t)
	clean := func() {
		bg := context.Background()
		pool.Exec(bg, `delete from rating_scores where report_id like 'hist-%'`)
		pool.Exec(bg, `delete from fight_metrics where report_id like 'hist-%'`)
		pool.Exec(bg, `delete from fights where report_id like 'hist-%'`)
		pool.Exec(bg, `delete from reports where id like 'hist-%'`)
	}
	clean()
	t.Cleanup(clean)
	for _, at := range []time.Time{histBase, histBase.AddDate(0, 1, 0), histBase.AddDate(0, 2, 0)} {
		if err := db.EnsureRatingsPartition(context.Background(), pool, at); err != nil {
			t.Fatal(err)
		}
		if err := db.EnsureMetricsPartition(context.Background(), pool, at); err != nil {
			t.Fatal(err)
		}
	}
	return pool
}

func histReport(t *testing.T, pool *pgxpool.Pool, id, visibility string, flagged *string) string {
	t.Helper()
	id = histPrefix + id
	mustExec(t, pool, `insert into reports (id, title, visibility, zone, status, flagged)
		values ($1, 'History', $2, 'Molten Core', 'complete', $3)`, id, visibility, flagged)
	return id
}

// histRate stores one rated fight of histKey; a nil score is a row that kept no snapshot.
func histRate(t *testing.T, pool *pgxpool.Pool, report string, index int, at time.Time, kill bool, score *int, outputPct string) {
	t.Helper()
	histRateOn(t, pool, report, index, at, histEncounter, 3, kill, score, outputPct)
}

// histRateOn is histRate on a chosen boss and difficulty id.
func histRateOn(t *testing.T, pool *pgxpool.Pool, report string, index int, at time.Time, encounter, difficulty int64, kill bool, score *int, outputPct string) {
	t.Helper()
	tier := (*string)(nil)
	if score != nil {
		tier = new(string)
		*tier = "epic"
	}
	components := fmt.Sprintf(`[{"name":"output","percentile":%s,"excluded":false},{"name":"survival","percentile":12,"excluded":false}]`, outputPct)
	mustExec(t, pool, `insert into rating_scores (report_id, fight_index, player_key, kill, components, model_version,
		fought_at, encounter_id, difficulty, score, tier) values ($1, $2, $3, $4, $5::jsonb, 'test', $6, $7, $8, $9, $10)`,
		report, index, histKey, kill, components, at, encounter, difficulty, score, tier)
}

func histMetric(t *testing.T, pool *pgxpool.Pool, report string, index int, at time.Time) {
	t.Helper()
	mustExec(t, pool, `insert into fight_metrics (report_id, fight_index, player_key, player_name, class, spec, role,
		metric_dps, encounter_id, difficulty, kill, phase, fought_at) values ($1, $2, $3, 'Histmain', 'Warrior', 'Arms',
		'dps', 100, $4, 3, true, '', $5)`, report, index, histKey, histEncounter, at)
}

func intp(n int) *int { return &n }

func TestReadOwnerHistoryIncludesGuildAndPrivateButNeverFlagged(t *testing.T) {
	pool := histPool(t)
	flag := "tampered"
	for i, c := range []struct {
		id, visibility string
		flagged        *string
	}{
		{"pub", "public", nil}, {"unl", "unlisted", nil}, {"gld", "guild", nil},
		{"prv", "private", nil}, {"flg", "public", &flag},
	} {
		id := histReport(t, pool, c.id, c.visibility, c.flagged)
		histRate(t, pool, id, 1, histBase.Add(time.Duration(i)*time.Hour), true, intp(1500+i), "50")
	}
	h, err := (&Store{Pool: pool}).ReadOwnerHistory(context.Background(), histKey)
	if err != nil {
		t.Fatal(err)
	}
	var got []string
	for _, p := range h.Score {
		got = append(got, p.Visibility)
	}
	want := []string{"public", "public", "guild", "private"}
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("visibilities = %v, want %v (oldest first, unlisted as public, flagged never)", got, want)
	}
	if h.Score[0].Score != 1500 || h.Score[0].Tier != "epic" || h.Score[0].Difficulty != 3 {
		t.Errorf("first point = %+v", h.Score[0])
	}
}

func TestReadOwnerHistoryParsePointsComeFromTheOutputComponentOnKillsOnly(t *testing.T) {
	pool := histPool(t)
	id := histReport(t, pool, "parse", "public", nil)
	histRate(t, pool, id, 1, histBase, true, intp(1000), "40")
	histRate(t, pool, id, 2, histBase.Add(time.Hour), false, intp(1010), "90")
	histRate(t, pool, id, 3, histBase.Add(2*time.Hour), true, intp(1020), "null")
	histRate(t, pool, id, 4, histBase.Add(3*time.Hour), true, intp(1030), "88")
	h, err := (&Store{Pool: pool}).ReadOwnerHistory(context.Background(), histKey)
	if err != nil {
		t.Fatal(err)
	}
	if p := h.Score[1].ParsePercentile; p == nil || *p != 90 {
		t.Errorf("a wipe still carries its output percentile on the score line, got %v", p)
	}
	if h.Score[2].ParsePercentile != nil {
		t.Errorf("a kill with no output percentile must be nil, got %v", *h.Score[2].ParsePercentile)
	}
	if len(h.Encounters) != 1 || len(h.Encounters[0].Points) != 2 {
		t.Fatalf("encounters = %+v, want one boss with the two kills that have a parse", h.Encounters)
	}
	pts := h.Encounters[0].Points
	if pts[0].ParsePercentile != 40 || pts[1].ParsePercentile != 88 || pts[1].FightIndex != 4 {
		t.Errorf("points = %+v, want 40 then 88 (oldest first)", pts)
	}
}

func TestReadOwnerHistoryGapsUseTheClosedReasonSetInPrecedence(t *testing.T) {
	pool := histPool(t)
	flag := "tampered"
	open := histReport(t, pool, "open", "public", nil)
	flagged := histReport(t, pool, "flagged", "public", &flag)
	day := func(month time.Month, d int) time.Time { return time.Date(2026, month, d, 20, 0, 0, 0, time.UTC) }
	histRate(t, pool, open, 1, day(3, 1), true, intp(1000), "50")
	histRate(t, pool, open, 2, day(3, 20), true, intp(1000), "50") // 19 days: no rated fights
	histRate(t, pool, open, 3, day(3, 30), false, nil, "50")       // a row that kept no score
	histRate(t, pool, open, 4, day(4, 10), true, intp(1000), "50") // 21 days: score not kept
	histRate(t, pool, flagged, 1, day(4, 20), true, nil, "50")     // flagged, and also unkept
	histRate(t, pool, open, 5, day(5, 3), true, intp(1000), "50")  // 23 days: flagged wins
	histRate(t, pool, open, 6, day(5, 16), true, intp(1000), "50") // 13 days: no gap
	h, err := (&Store{Pool: pool}).ReadOwnerHistory(context.Background(), histKey)
	if err != nil {
		t.Fatal(err)
	}
	want := []HistoryGap{
		{From: "2026-03-01", To: "2026-03-20", Reason: GapNoRatedFights},
		{From: "2026-03-20", To: "2026-04-10", Reason: GapScoreNotKept},
		{From: "2026-04-10", To: "2026-05-03", Reason: GapReportFlagged},
	}
	if fmt.Sprint(h.Gaps) != fmt.Sprint(want) {
		t.Fatalf("gaps = %v, want %v", h.Gaps, want)
	}
}

func TestReadOwnerHistoryFightsAreDedupedNewestFirstWithVisibility(t *testing.T) {
	pool := histPool(t)
	pub := histReport(t, pool, "fpub", "public", nil)
	prv := histReport(t, pool, "fprv", "private", nil)
	histRate(t, pool, pub, 1, histBase, true, intp(1000), "50")
	histRate(t, pool, prv, 1, histBase.Add(time.Hour), true, intp(1100), "60")
	unl := histReport(t, pool, "funl", "unlisted", nil)
	histRate(t, pool, unl, 1, histBase.Add(2*time.Hour), true, intp(1200), "70")
	histMetric(t, pool, unl, 1, histBase.Add(2*time.Hour))
	histMetric(t, pool, pub, 1, histBase)
	histMetric(t, pool, prv, 1, histBase.Add(time.Hour))
	h, err := (&Store{Pool: pool}).ReadOwnerHistory(context.Background(), histKey)
	if err != nil {
		t.Fatal(err)
	}
	if len(h.Fights) != 3 {
		t.Fatalf("fights = %d, want 3 (one per fight)", len(h.Fights))
	}
	if h.Fights[0].ReportID != unl || h.Fights[0].Visibility != "public" ||
		h.Fights[1].ReportID != prv || h.Fights[1].Visibility != "private" ||
		h.Fights[2].ReportID != pub || h.Fights[2].Visibility != "public" {
		t.Errorf("fights = %+v, want unlisted (as public), private, public, newest first", h.Fights)
	}
	if h.Fights[1].Metric != "dps" || h.Fights[1].Value != 100 {
		t.Errorf("fight row shape = %+v", h.Fights[1].CharacterFight)
	}
}

func TestReadOwnerHistoryCapsAtTheNewestThousandAndReturnsNonNilSlices(t *testing.T) {
	pool := histPool(t)
	id := histReport(t, pool, "cap", "public", nil)
	mustExec(t, pool, `insert into rating_scores (report_id, fight_index, player_key, kill, components, model_version,
		fought_at, encounter_id, difficulty, score, tier)
		select $1, i, $2, true, '[]'::jsonb, 'test', $3::timestamptz + i * interval '1 minute', $4, 3, 1000, 'epic'
		from generate_series(1, 1001) i`, id, histKey, histBase, histEncounter)
	store := &Store{Pool: pool}
	h, err := store.ReadOwnerHistory(context.Background(), histKey)
	if err != nil {
		t.Fatal(err)
	}
	if len(h.Score) != HistoryMaxFights || h.Score[0].FightIndex != 2 || h.Score[len(h.Score)-1].FightIndex != 1001 {
		t.Fatalf("score = %d points from fight %d, want the newest %d (oldest dropped)",
			len(h.Score), h.Score[0].FightIndex, HistoryMaxFights)
	}
	empty, err := store.ReadOwnerHistory(context.Background(), "us/normal/histnobody")
	if err != nil {
		t.Fatal(err)
	}
	if empty.Score == nil || empty.Encounters == nil || empty.Gaps == nil || empty.Fights == nil {
		t.Errorf("an empty history must carry empty, not nil, slices: %+v", empty)
	}
}

func TestReadOwnerHistoryEncounterGroupsSplitByDifficultyIdAndOrderByNewestKill(t *testing.T) {
	pool := histPool(t)
	id := histReport(t, pool, "groups", "public", nil)
	other := int64(histEncounter + 1)
	at := func(h int) time.Time { return histBase.Add(time.Duration(h) * time.Hour) }
	histRateOn(t, pool, id, 1, at(0), histEncounter, 3, true, intp(1000), "10")
	histRateOn(t, pool, id, 2, at(1), histEncounter, 4, true, intp(1000), "20")
	histRateOn(t, pool, id, 3, at(2), other, 3, true, intp(1000), "30")
	histRateOn(t, pool, id, 4, at(3), histEncounter, 3, true, intp(1000), "40")
	histRateOn(t, pool, id, 5, at(4), other, 3, true, intp(1000), "50")
	histRateOn(t, pool, id, 6, at(5), histEncounter, 4, true, intp(1000), "60")
	h, err := (&Store{Pool: pool}).ReadOwnerHistory(context.Background(), histKey)
	if err != nil {
		t.Fatal(err)
	}
	type group struct {
		enc, diff int64
		pcts      string
	}
	var got []group
	for _, g := range h.Encounters {
		var pcts []float64
		for _, p := range g.Points {
			pcts = append(pcts, p.ParsePercentile)
		}
		got = append(got, group{g.EncounterID, g.Difficulty, fmt.Sprint(pcts)})
	}
	want := []group{
		{histEncounter, 4, "[20 60]"}, {other, 3, "[30 50]"}, {histEncounter, 3, "[10 40]"},
	}
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("groups = %v, want %v (newest kill first; points oldest first; ids 3 and 4 apart)", got, want)
	}
}

func TestGapsForThresholdIsMoreThanFourteenDays(t *testing.T) {
	d := func(day int) time.Time { return time.Date(2026, 3, day, 0, 0, 0, 0, time.UTC) }
	if gaps := gapsFor([]time.Time{d(1), d(15)}, nil); len(gaps) != 0 {
		t.Errorf("exactly 14 days apart is no gap, got %v", gaps)
	}
	gaps := gapsFor([]time.Time{d(1), d(16)}, nil)
	if len(gaps) != 1 || gaps[0].From != "2026-03-01" || gaps[0].To != "2026-03-16" {
		t.Errorf("15 days apart is a gap, got %v", gaps)
	}
}
