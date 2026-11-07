package activity

import (
	"context"
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/jackc/pgx/v5"
)

// The tier list's DPS tiers, best first: the letter starts at this many percent behind the
// top spec. The site's own rules are web/src/lib/tiers/tier-rules.ts (TIER_BANDS);
// TestTierBandsMatchTheSite holds the two together. Only damage roles get a letter.
var tierBands = []struct {
	letter string
	from   float64
}{{"S", 0}, {"A", 5}, {"B", 10}, {"C", 20}, {"D", 30}}

const gapPrecision = 1e6

// tierOf is the tier a gap falls in, rounded past float noise first (tier-list.ts tierOf).
func tierOf(gapPercent float64) string {
	gap := math.Round(gapPercent*gapPrecision) / gapPrecision
	letter := tierBands[0].letter
	for _, b := range tierBands {
		if gap >= b.from {
			letter = b.letter
		}
	}
	return letter
}

// gradeRow is one spec's place on its band, faction and role.
type gradeRow struct {
	spec, faction, role string
	band, place, of     int
	grade               *string
}

type groupKey struct {
	band          int
	faction, role string
}

type scoredBand struct {
	spec   string
	metric float64
}

// tierBandFor is the entry a band's grade reads: the bare list below 60, the raid list at
// 60 (the tier list's own), else the bare one.
func tierBandFor(f bisFile, band int, faction string) (bisBand, bool) {
	var bare, raid *bisBand
	for i := range f.Bands {
		b := &f.Bands[i]
		if b.Band != band || b.Faction != faction {
			continue
		}
		if b.preset() == barePreset {
			bare = b
		} else if b.Preset == "raid" {
			raid = b
		}
	}
	if band >= topBand && raid != nil {
		return *raid, true
	}
	if bare != nil {
		return *bare, true
	}
	return bisBand{}, false
}

// metricOf is the figure a role sorts on (tier-list.ts sortMetricOf): damage taken per
// second for a tank (lower is better), set DPS otherwise.
func metricOf(b bisBand) (float64, bool) {
	if b.role() == "tank" {
		if b.Metrics == nil || b.Metrics.DTPS <= 0 {
			return 0, false
		}
		return b.Metrics.DTPS, true
	}
	return b.SetDPS, b.SetDPS > 0
}

// gradeRows ranks every spec within each band, faction and role.
func gradeRows(files []bisFile) []gradeRow {
	groups := map[groupKey][]scoredBand{}
	var keys []groupKey
	for _, f := range files {
		for band := firstBand; band <= topBand; band += bandWidth {
			for _, faction := range []string{"alliance", "horde"} {
				b, ok := tierBandFor(f, band, faction)
				m, measured := metricOf(b)
				if !ok || !measured {
					continue
				}
				k := groupKey{band, faction, b.role()}
				if _, seen := groups[k]; !seen {
					keys = append(keys, k)
				}
				groups[k] = append(groups[k], scoredBand{spec: f.Spec, metric: m})
			}
		}
	}
	var out []gradeRow
	for _, k := range keys {
		out = append(out, rankGroup(k, groups[k])...)
	}
	return out
}

func rankGroup(k groupKey, group []scoredBand) []gradeRow {
	tank := k.role == "tank"
	sort.SliceStable(group, func(i, j int) bool {
		if tank {
			return group[i].metric < group[j].metric
		}
		return group[i].metric > group[j].metric
	})
	top := group[0].metric
	out := make([]gradeRow, len(group))
	for i, g := range group {
		row := gradeRow{spec: g.spec, band: k.band, faction: k.faction, role: k.role, place: i + 1, of: len(group)}
		if k.role == "dps" {
			letter := tierOf((1 - g.metric/top) * 100)
			row.grade = &letter
		}
		out[i] = row
	}
	return out
}

// SnapshotTiers upserts this week's tier_grade_weeks rows from the deployed BiS files.
func (s *Store) SnapshotTiers(ctx context.Context, dataDir, build string, week time.Time) (int, error) {
	files, err := loadBisFiles(dataDir, build)
	if err != nil {
		return 0, err
	}
	rows := gradeRows(files)
	batch := &pgx.Batch{}
	for _, r := range rows {
		batch.Queue(`insert into tier_grade_weeks (week, spec, band, faction, role, grade, place, of_n)
			values ($1, $2, $3, $4, $5, $6, $7, $8)
			on conflict (week, spec, band, faction) do update set role = excluded.role, grade = excluded.grade,
			  place = excluded.place, of_n = excluded.of_n, updated_at = now()`,
			week, r.spec, r.band, r.faction, r.role, r.grade, r.place, r.of)
	}
	if err := s.Pool.SendBatch(ctx, batch).Close(); err != nil {
		return 0, fmt.Errorf("activity: snapshot tiers: %w", err)
	}
	return len(rows), nil
}
