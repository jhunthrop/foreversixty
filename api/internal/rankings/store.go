// Package rankings is the ranked side of the logs product: the rows
// written at fight close, the percentile digests kept beside them, and
// the character, guild, and leaderboard reads the site shows.
package rankings

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"slices"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/digest"
	"github.com/jhunthrop/foreversixty/api/internal/phase"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/api/internal/spec"
	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// The metrics a row can be ranked on.
const (
	MetricDPS         = "dps"
	MetricHPS         = "hps"
	MetricDamageTaken = "damage_taken"
	// MetricExecution ranks on how much of what a player's gear can
	// do they actually did, rather than on the raw number. It is the
	// one metric that is not a column the ingest writes: the
	// simulator fills it in per fight.
	MetricExecution = "execution"
)

// Metrics is every metric a leaderboard can be sorted by.
var Metrics = []string{MetricDPS, MetricHPS, MetricDamageTaken, MetricExecution}

// DigestedMetrics is every metric percentile_digests holds a curve
// for, which is the smaller list: a percentile is a place on a
// distribution the digest job folded, and nothing folds execution
// scores. GET /v1/rankings/percentile reads this one, so asking it to
// place an execution score is the 400 it was before this task rather
// than a silent "no percentile for that".
var DigestedMetrics = []string{MetricDPS, MetricHPS, MetricDamageTaken}

// ValidMetric reports whether m is one of them.
func ValidMetric(m string) bool {
	for _, v := range Metrics {
		if v == m {
			return true
		}
	}
	return false
}

// ValidDigestMetric reports whether a metric has a percentile curve.
// Written as a loop, the way ValidMetric beside it already is.
func ValidDigestMetric(m string) bool {
	for _, v := range DigestedMetrics {
		if v == m {
			return true
		}
	}
	return false
}

// The moderation states a row can be in.
const (
	StateOK      = "ok"
	StateAtRisk  = "at_risk"
	StateRemoved = "removed"
)

// maxSinceDays is the longest "last n days" window a read may ask for,
// which is ten years: past that the parameter is a typo, not a window.
const maxSinceDays = 3650

// trinketSlots are the indexes of the two trinkets in a COMBATANT_INFO
// gear list. This is the retail layout's order; the first beta log
// settles Forever's, and this is the only place the API reads a slot
// position out of a gear list.
var trinketSlots = [2]int{12, 13}

// Store is every ranking read and write.
type Store struct {
	Pool *pgxpool.Pool
	// Specs names a build's spec from its talents. It may be nil, in
	// which case rows are stored with the engine's own spec string.
	Specs *spec.Inferrer
}

// Store is the ingest's ranker.
var _ reports.Ranker = (*Store)(nil)

// metricOf is the metric a role's headline figure is read on: healing
// for a healer, damage for everyone else, tanks included, the way the
// parse culture around Warcraft Logs reads them. The digests hold every
// metric for every spec, so any other reading is a query away.
func metricOf(role string) string {
	if role == "healer" {
		return MetricHPS
	}
	return MetricDPS
}

// digestKey is one percentile bracket: a metric and the spec it was
// posted by, within the fight's own encounter, difficulty and phase.
type digestKey struct {
	metric string
	spec   string
}

// WriteFight stores one verified fight's rows and folds their values
// into the percentile digests, in one transaction: a row is never
// visible without the digest that ranks it.
//
// It is idempotent per (report_id, fight_index). A fight can arrive
// twice - a second verifying bundle for the same index with a different
// raw range does not hit the ingest's hash short-circuit - so the
// fight's rows are withdrawn and rewritten rather than added to, and a
// rewrite leaves the digests alone: its values are already in them, a
// t-digest cannot have a value taken back out, and folding them again
// would count the fight twice.
func (s *Store) WriteFight(ctx context.Context, f reports.RankedFight) error {
	if len(f.Rows) == 0 {
		return nil
	}
	if err := db.EnsureMetricsPartition(ctx, s.Pool, f.FoughtAt); err != nil {
		return err
	}
	at := phase.At(f.FoughtAt)
	combatants := byGUID(f.Combatants)

	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("rankings: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	// Serialise writers of the same fight. Without this, two of them
	// race: both run the delete below before either has committed an
	// insert, both count zero rows withdrawn, both call themselves a
	// first write, and the fight lands in the digests twice. The lock
	// is transaction-scoped, so it is released by the commit or the
	// rollback below. A hash collision between two different fights
	// only makes them wait for each other.
	if _, err := tx.Exec(ctx, `select pg_advisory_xact_lock(hashtext($1), $2)`,
		f.ReportID, f.FightIndex); err != nil {
		return fmt.Errorf("rankings: lock %s/%d: %w", f.ReportID, f.FightIndex, err)
	}

	// Withdraw whatever this fight wrote before, so a rewrite replaces
	// its rows instead of leaving a stale player behind, and so the
	// digests can tell a first write from a second. Under the lock
	// above, this count is authoritative.
	withdrawn, err := tx.Exec(ctx,
		`delete from fight_metrics where report_id = $1 and fight_index = $2`, f.ReportID, f.FightIndex)
	if err != nil {
		return fmt.Errorf("rankings: withdraw %s/%d: %w", f.ReportID, f.FightIndex, err)
	}
	rewrite := withdrawn.RowsAffected() > 0

	pending := map[digestKey][]float64{}

	for _, row := range f.Rows {
		name, _ := character.SplitUnit(row.Name)
		key := character.KeyFromUnit(f.Region, f.Ruleset, row.Name)
		c := combatants[row.PlayerGUID]
		specName, split := s.Specs.Of(row.Class, talentsOf(c))
		if specName == "" {
			specName = row.Spec
		}
		if _, err := tx.Exec(ctx,
			`insert into fight_metrics (report_id, fight_index, player_key, player_name, class, spec,
			   role, ilvl, metric_dps, metric_hps, damage_taken, active_ms, deaths, encounter_id,
			   difficulty, size, duration_ms, kill, phase, fought_at, talent_split, trinkets,
			   buff_count, faction, state)
			 values ($1, $2, $3, $4, nullif($5, ''), nullif($6, ''), $7, $8, $9, $10, $11, $12, $13,
			   $14, $15, $16, $17, $18, $19, $20, nullif($21, ''), $22, $23, nullif($24, ''), $25)
			 on conflict (report_id, fight_index, player_key, fought_at) do update set
			   player_name = excluded.player_name, class = excluded.class, spec = excluded.spec,
			   role = excluded.role, ilvl = excluded.ilvl, metric_dps = excluded.metric_dps,
			   metric_hps = excluded.metric_hps, damage_taken = excluded.damage_taken,
			   active_ms = excluded.active_ms, deaths = excluded.deaths,
			   duration_ms = excluded.duration_ms, kill = excluded.kill,
			   talent_split = excluded.talent_split, trinkets = excluded.trinkets,
			   buff_count = excluded.buff_count, faction = excluded.faction`,
			f.ReportID, f.FightIndex, key, name, row.Class, specName, row.Role, row.Ilvl,
			row.MetricDPS, row.MetricHPS, row.DamageTaken, row.ActiveMS, row.Deaths,
			row.EncounterID, row.Difficulty, row.Size, row.DurationMS, row.Kill, at, f.FoughtAt,
			split, trinketsOf(c), buffCountOf(c), f.Factions[row.PlayerGUID], StateOK); err != nil {
			return fmt.Errorf("rankings: write row %s/%d/%s: %w", f.ReportID, f.FightIndex, key, err)
		}
		if row.Kill && !rewrite {
			// Only kills feed the digests: a wipe's numbers are not
			// comparable with a kill's, and the percentile a player is
			// shown is against kills. Every row feeds every metric's
			// bracket for its spec, the way Warcraft Logs ranks: a
			// healer has a damage parse among healers of that spec and
			// a tank a damage parse among tanks, not only the metric
			// the leaderboard sorts their role by.
			for _, metric := range Metrics {
				value := valueOf(row.MetricDPS, row.MetricHPS, float64(row.DamageTaken), metric)
				pending[digestKey{metric, specName}] = append(pending[digestKey{metric, specName}], value)
			}
		}
	}

	// In a fixed order, so two fights that touch an overlapping set of
	// brackets take the digest row locks in the same order and cannot
	// deadlock against each other.
	for _, k := range sortedKeys(pending) {
		if err := updateDigest(ctx, tx, f.Rows[0].EncounterID, f.Rows[0].Difficulty,
			k.spec, at, k.metric, pending[k]); err != nil {
			return err
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("rankings: commit: %w", err)
	}
	return nil
}

// sortedKeys is the brackets of one fight in a stable order: by metric,
// then by spec. Go randomises map iteration, and these keys name rows
// that are about to be locked.
func sortedKeys(pending map[digestKey][]float64) []digestKey {
	keys := slices.Collect(maps.Keys(pending))
	slices.SortFunc(keys, func(a, b digestKey) int {
		if c := strings.Compare(a.metric, b.metric); c != 0 {
			return c
		}
		return strings.Compare(a.spec, b.spec)
	})
	return keys
}

// updateDigest folds values into one digest row, holding that row's
// lock across the read and the write so two fights closing at once
// cannot lose each other's values.
//
// The lock is taken by creating the row rather than by selecting it:
// `select ... for update` locks nothing when nothing matches, so two
// transactions first-populating the same brand-new bracket would both
// read an empty digest, both fold only their own values, and the
// second would overwrite the first's committed values. The insert
// below always leaves a locked row behind - a new empty one, or the
// existing one, whose digest it returns unchanged - and the update
// that follows is protected by that lock.
func updateDigest(ctx context.Context, tx pgx.Tx, encounterID, difficulty int64,
	specName, at, metric string, values []float64) error {
	var raw []byte
	if err := tx.QueryRow(ctx,
		`insert into percentile_digests (encounter_id, difficulty, spec, phase, metric, digest, n)
		 values ($1, $2, $3, $4, $5, $6, 0)
		 on conflict (encounter_id, difficulty, spec, phase, metric) do update set
		   digest = percentile_digests.digest
		 returning digest`,
		encounterID, difficulty, specName, at, metric, []byte{}).Scan(&raw); err != nil {
		return fmt.Errorf("rankings: lock digest: %w", err)
	}
	d, err := digest.Unmarshal(raw)
	if err != nil {
		return err
	}
	for _, v := range values {
		d.Add(v)
	}
	encoded, err := d.MarshalBinary()
	if err != nil {
		return err
	}
	if _, err := tx.Exec(ctx,
		`update percentile_digests set digest = $6, n = $7, updated_at = now()
		 where encounter_id = $1 and difficulty = $2 and spec = $3 and phase = $4 and metric = $5`,
		encounterID, difficulty, specName, at, metric, encoded, d.Count()); err != nil {
		return fmt.Errorf("rankings: write digest: %w", err)
	}
	return nil
}

// RemoveReport withdraws a report's ranking rows and records why. A
// report with no rows - one whose fights never passed verification - is
// not an error: there is nothing to take back, and the withdrawal is
// still recorded, because moderation is an append-only log of what was
// decided about a report.
//
// reason is the caller's own words for the withdrawal, because the
// three callers withdraw for three different reasons: a tampered
// report, a re-sent bundle that stopped verifying, and an owner making
// their report private. The log would be a lie if it named only the
// first.
//
// The digests are left as they are: a t-digest cannot have a value
// taken back out, and the alternative - rebuilding every affected
// digest from the rows - would cost far more than one tampered report
// distorts.
func (s *Store) RemoveReport(ctx context.Context, reportID, reason string) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("rankings: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if _, err := tx.Exec(ctx,
		`insert into moderation (target_kind, target_id, state, reason)
		 values ('report', $1, 'removed', $2)`, reportID, reason); err != nil {
		return fmt.Errorf("rankings: record moderation: %w", err)
	}
	if _, err := tx.Exec(ctx, `delete from fight_metrics where report_id = $1`, reportID); err != nil {
		return fmt.Errorf("rankings: remove rows of %s: %w", reportID, err)
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("rankings: commit: %w", err)
	}
	return nil
}

// Percentile is where value stands among the kills in its bracket, as a
// number from 0 to 100. The second return is false when the bracket has
// no digest yet.
// Percentile places value among the bracket's ranked kills and says how many
// there are: a 0 among one kill and a 0 among a thousand are different
// answers, and the page has to be able to tell them apart.
func (s *Store) Percentile(ctx context.Context, encounterID, difficulty int64,
	specName, at, metric string, value float64) (float64, int64, bool, error) {
	var raw []byte
	err := s.Pool.QueryRow(ctx,
		`select digest from percentile_digests
		 where encounter_id = $1 and difficulty = $2 and spec = $3 and phase = $4 and metric = $5`,
		encounterID, difficulty, specName, at, metric).Scan(&raw)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, 0, false, nil
	}
	if err != nil {
		return 0, 0, false, fmt.Errorf("rankings: read digest: %w", err)
	}
	d, err := digest.Unmarshal(raw)
	if err != nil {
		return 0, 0, false, err
	}
	if d.Count() == 0 {
		return 0, 0, false, nil
	}
	return d.Placement(value) * 100, d.Count(), true, nil
}

// byGUID indexes the combatant rows a fight carried.
func byGUID(rows []summary.CombatantRow) map[string]summary.CombatantRow {
	out := make(map[string]summary.CombatantRow, len(rows))
	for _, r := range rows {
		out[r.GUID] = r
	}
	return out
}

func talentsOf(c summary.CombatantRow) []int64 { return c.Talents }

// trinketsOf reads the two trinket item ids out of a gear list.
func trinketsOf(c summary.CombatantRow) []int64 {
	out := []int64{}
	for _, slot := range trinketSlots {
		if slot < len(c.Gear) && c.Gear[slot].ID != 0 {
			out = append(out, c.Gear[slot].ID)
		}
	}
	return out
}

// buffCountOf is how many raid buffs and consumables a player pulled
// with.
//
// FOLLOW-UP (data): the engine counts these against the spell tables it
// is given, and the API gives it none yet, so this is zero until
// data/curated names the raid buffs and consumables. The column is
// written from the start so the rankings row shape never changes.
func buffCountOf(c summary.CombatantRow) int {
	return len(c.RaidBuffs) + len(c.Consumables)
}

// valueOf picks the number a row is ranked on.
func valueOf(dps, hps, taken float64, metric string) float64 {
	switch metric {
	case MetricHPS:
		return hps
	case MetricDamageTaken:
		return taken
	default:
		return dps
	}
}

// FactionsFromEvents reads each player's faction out of the
// COMBATANT_INFO lines of a fight. The Parquet schema does not carry
// the payload, so only a parse that sees the original log text - the
// whole-file job - can fill this in; live bundles leave it empty until
// the engine's Parquet rows carry the combatant block.
func FactionsFromEvents(events []event.Event) map[string]string {
	out := map[string]string{}
	for _, e := range events {
		if e.Kind != event.CombatantInfo || e.Combatant == nil {
			continue
		}
		if name := factionName(e.Combatant.Faction); name != "" {
			out[e.Combatant.GUID] = name
		}
	}
	return out
}

// factionName maps COMBATANT_INFO's faction number to a name. Zero is
// Horde and one is Alliance, as the game writes them.
func factionName(faction int64) string {
	switch faction {
	case 0:
		return "horde"
	case 1:
		return "alliance"
	default:
		return ""
	}
}

// since parses the all-time versus today parameter: "today" or a
// duration like "7d". An empty value means all time.
func since(v string, now time.Time) (time.Time, error) {
	switch v = strings.TrimSpace(strings.ToLower(v)); v {
	case "":
		return time.Time{}, nil
	case "today":
		return now.UTC().Truncate(24 * time.Hour), nil
	}
	if days, ok := strings.CutSuffix(v, "d"); ok {
		var n int
		if _, err := fmt.Sscanf(days, "%d", &n); err == nil && n > 0 && n <= maxSinceDays {
			return now.UTC().AddDate(0, 0, -n), nil
		}
	}
	return time.Time{}, fmt.Errorf("since must be empty, today, or a number of days like 7d")
}

// The bounds a stored execution score is clamped to. A ratio outside
// them is a sim that went wrong rather than a player who played twice
// as well as their gear allows, and the contract pins the range.
const (
	ExecutionMin = 0.0
	ExecutionMax = 2.0
)

// SetExecutionScore writes one player's execution score on one fight,
// clamped to [ExecutionMin, ExecutionMax]. The fight-close scorer and
// the nightly job are the callers; both are idempotent, so this
// simply overwrites whatever was there.
func (s *Store) SetExecutionScore(ctx context.Context, reportID string, fightIndex int,
	playerKey string, score float64) error {
	score = min(max(score, ExecutionMin), ExecutionMax)
	if _, err := s.Pool.Exec(ctx,
		`update fight_metrics set execution_score = $4
		 where report_id = $1 and fight_index = $2 and player_key = $3`,
		reportID, fightIndex, playerKey, score); err != nil {
		return fmt.Errorf("rankings: execution score %s/%d/%s: %w", reportID, fightIndex, playerKey, err)
	}
	return nil
}
