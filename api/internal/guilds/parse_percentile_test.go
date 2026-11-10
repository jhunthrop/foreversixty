package guilds

import (
	"math"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/digest"
)

// seedKillDigest writes a dps digest of values under the bracket of reportID's first boss-kill
// row (read back through bracketColumns, the key every reader places by) and returns it.
func seedKillDigest(t *testing.T, pool *pgxpool.Pool, reportID string, values []float64) *digest.Digest {
	t.Helper()
	var b parseBracket
	if err := pool.QueryRow(t.Context(), `select m.encounter_id, `+bracketColumns+` from fight_metrics m
		where m.report_id = $1 and `+bossKillMetrics+` limit 1`, reportID).
		Scan(&b.EncounterID, &b.Difficulty, &b.Spec, &b.Phase); err != nil {
		t.Fatal(err)
	}
	d := digest.New()
	for _, v := range values {
		d.Add(v)
	}
	raw, err := d.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(t.Context(), `insert into percentile_digests (encounter_id, difficulty, spec, phase, metric, digest, n)
		values ($1, $2, $3, $4, 'dps', $5, $6)`, b.EncounterID, b.Difficulty, b.Spec, b.Phase, raw, d.Count()); err != nil {
		t.Fatal(err)
	}
	return d
}

// percentileMatches reports whether got is digest d's placement of value, or null when d is nil.
func percentileMatches(d *digest.Digest, value float64, got *float64) bool {
	if d == nil {
		return got == nil
	}
	return got != nil && math.Abs(*got-d.Placement(value)*100) <= 1e-9
}

// placeParse is rankings' percentileOf rule: a bracket with no digest, or an empty one, is
// null; otherwise the digest's placement scaled to 0-100.
func TestPlaceParseReadsItsOwnBracket(t *testing.T) {
	full := digest.New()
	for _, v := range []float64{100, 200, 300, 400} {
		full.Add(v)
	}
	k := parseBracket{EncounterID: 1084, Spec: "Fire", Metric: "dps"}
	other := parseBracket{EncounterID: 1084, Spec: "Frost", Metric: "dps"}
	cases := []struct {
		name    string
		digests map[parseBracket]*digest.Digest
		key     parseBracket
		value   float64
		want    *float64
	}{
		{name: "no digest", digests: map[parseBracket]*digest.Digest{}, key: k, value: 300},
		{name: "empty digest", digests: map[parseBracket]*digest.Digest{k: digest.New()}, key: k, value: 300},
		{name: "another spec's digest", digests: map[parseBracket]*digest.Digest{other: full}, key: k, value: 300},
		{name: "own digest", digests: map[parseBracket]*digest.Digest{k: full}, key: k, value: 300,
			want: ptrFloat(full.Placement(300) * 100)},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := placeParse(tc.digests, tc.key, tc.value)
			switch {
			case tc.want == nil && got != nil:
				t.Fatalf("got %v, want null", *got)
			case tc.want != nil && (got == nil || math.Abs(*got-*tc.want) > 1e-9):
				t.Fatalf("got %v, want %v", got, *tc.want)
			}
		})
	}
}

func ptrFloat(v float64) *float64 { return &v }

func TestTheReportsParsesCarryTheirRankingsPercentile(t *testing.T) {
	cases := []struct {
		name   string
		digest []float64 // nil: no digest row for the bracket
	}{
		{name: "digest present", digest: []float64{50, 100, 150, 200, 250, 300, 350, 400}},
		{name: "digest absent", digest: nil},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w, raid, _ := digestedNight(t)
			if _, err := w.h.pool.Exec(t.Context(), `truncate percentile_digests`); err != nil {
				t.Fatal(err)
			}
			if _, err := w.h.pool.Exec(t.Context(), `update fight_metrics set metric_dps = 300 where player_key = $1`, memberKeyB); err != nil {
				t.Fatal(err)
			}
			var d *digest.Digest
			if tc.digest != nil {
				d = seedKillDigest(t, w.h.pool, "night", tc.digest)
			}
			w.as(w.officer)
			rep := w.getRaid(t, raid.ID).Report
			if len(rep.Parses) == 0 {
				t.Fatal("no parses")
			}
			for _, p := range rep.Parses {
				if !percentileMatches(d, p.Value, p.Percentile) {
					t.Fatalf("%s percentile = %v, want the digest's placement of %v (null with none)",
						p.CharacterKey, p.Percentile, p.Value)
				}
			}
			if d != nil && *rep.Parses[0].Percentile <= *rep.Parses[len(rep.Parses)-1].Percentile {
				t.Fatalf("the best parse does not place above the worst: %+v", rep.Parses)
			}
		})
	}
}

// The roster's best_parse carries the same placement as the report, null with no digest.
func TestRosterBestParseCarriesItsRankingsPercentile(t *testing.T) {
	cases := []struct {
		name   string
		digest []float64 // nil: no digest row for the bracket
	}{
		{name: "digest present", digest: []float64{25, 50, 75, 100, 125, 150}},
		{name: "digest absent", digest: nil},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g := seedL612Guild(t)
			if _, err := g.pool.Exec(t.Context(), `truncate percentile_digests`); err != nil {
				t.Fatal(err)
			}
			var d *digest.Digest
			if tc.digest != nil {
				d = seedKillDigest(t, g.pool, g.cleanReport, tc.digest)
			}
			bp := g.rosterRow(t).BestParse
			if bp == nil || bp.ReportID != g.cleanReport {
				t.Fatalf("best_parse = %+v, want the clean report's", bp)
			}
			if !percentileMatches(d, bp.Value, bp.Percentile) {
				t.Fatalf("best_parse percentile = %v, want the digest's placement of %v (null with none)", bp.Percentile, bp.Value)
			}
		})
	}
}

// bracketOf keeps the row's difficulty, spec and phase and sets the encounter and metric.
func TestBracketOfKeepsTheRowsAxes(t *testing.T) {
	cases := []struct {
		name      string
		row       parseBracket
		encounter int64
		metric    string
		want      parseBracket
	}{
		{name: "dps row", row: parseBracket{Difficulty: 3, Spec: "Arms", Phase: "p1"}, encounter: 1084, metric: "dps",
			want: parseBracket{EncounterID: 1084, Difficulty: 3, Spec: "Arms", Phase: "p1", Metric: "dps"}},
		{name: "healer row with no axes", row: parseBracket{}, encounter: 610, metric: "hps",
			want: parseBracket{EncounterID: 610, Metric: "hps"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := bracketOf(tc.row, tc.encounter, tc.metric); got != tc.want {
				t.Fatalf("bracketOf = %+v, want %+v", got, tc.want)
			}
		})
	}
}
