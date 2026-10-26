// api/internal/guilds/raids_report_view.go
//
// The report's consumables, loot and moves blocks (design/specs/2026-10-27-guild-attendance.md
// §3.4): the readiness board's bags rule, the awards made for the raid, and the biggest
// changes against the previous done raid of the zone.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/jackc/pgx/v5"
)

// movesCap is how many changes the report lists.
const movesCap = 5

// consumableOrder ranks the Before pills worst first.
var consumableOrder = map[string]int{"short": 0, "unknown": 1, "stocked": 2}

func (s *Store) reportConsumables(ctx context.Context, vc raidViewContext, reportID string) (ReportConsumablesView, error) {
	out := ReportConsumablesView{Rows: []ConsumableRow{}}
	inLog, err := s.reportPlayerKeys(ctx, reportID)
	if err != nil {
		return out, err
	}
	g, err := s.getGuild(ctx, vc.guildID)
	if err != nil {
		return out, err
	}
	roster, err := s.HomeRoster(ctx, vc.guildID, g, 0, false, false)
	if err != nil {
		return out, err
	}
	for _, row := range roster {
		if !row.Verified || !inLog[row.CharacterKey] {
			continue
		}
		out.Rows = append(out.Rows, ConsumableRow{CharacterKey: row.CharacterKey, Name: row.Name, Class: row.className,
			Before: consumablesState(row.Consent, row.bags, row.hasBagsSection), own: vc.ownsKey(row.CharacterKey)})
	}
	sort.SliceStable(out.Rows, func(i, j int) bool {
		a, b := out.Rows[i], out.Rows[j]
		if consumableOrder[a.Before] != consumableOrder[b.Before] {
			return consumableOrder[a.Before] < consumableOrder[b.Before]
		}
		return a.Name < b.Name
	})
	return out, nil
}

func (s *Store) reportPlayerKeys(ctx context.Context, reportID string) (map[string]bool, error) {
	rows, err := s.Pool.Query(ctx,
		`select distinct player_key from fight_metrics where report_id = $1 and state = 'ok'`, reportID)
	if err != nil {
		return nil, fmt.Errorf("guilds: report players: %w", err)
	}
	defer rows.Close()
	out := map[string]bool{}
	for rows.Next() {
		var key string
		if err := rows.Scan(&key); err != nil {
			return nil, fmt.Errorf("guilds: report players: %w", err)
		}
		out[key] = true
	}
	return out, rows.Err()
}

func (s *Store) reportLoot(ctx context.Context, vc raidViewContext, raidID int64) ([]ReportLootBoss, error) {
	rows, err := s.Pool.Query(ctx, `
		select la.encounter_id, la.item_id, la.character_key,
		       coalesce(nullif(c.name, ''), nullif(ae.name, ''), split_part(la.character_key, '/', 3)),
		       lower(coalesce(c.class, '')), la.awarded_at, la.awarded_by
		from loot_awards la
		left join characters c on c.key = la.character_key
		left join addon_exports ae on ae.character_key = la.character_key
		where la.raid_id = $1 order by la.encounter_id, la.awarded_at, la.id`, raidID)
	if err != nil {
		return nil, fmt.Errorf("guilds: report loot: %w", err)
	}
	defer rows.Close()
	bosses := []ReportLootBoss{}
	for rows.Next() {
		var encounterID int64
		var a ReportLootAward
		var at time.Time
		var by *int64
		if err := rows.Scan(&encounterID, &a.ItemID, &a.CharacterKey, &a.Name, &a.Class, &at, &by); err != nil {
			return nil, fmt.Errorf("guilds: report loot: %w", err)
		}
		a.At = at.UTC().Format(time.RFC3339)
		a.ItemName, a.Quality = lootItemFacts(a.ItemID)
		if by != nil && s.Accounts != nil {
			a.ByName = s.awarderName(ctx, vc.guildID, *by)
		}
		bosses = addLootAward(bosses, encounterID, a)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: report loot: %w", err)
	}
	for _, b := range bosses {
		sort.SliceStable(b.Awards, func(i, j int) bool {
			if b.Awards[i].Quality != b.Awards[j].Quality {
				return b.Awards[i].Quality > b.Awards[j].Quality
			}
			return b.Awards[i].ItemName < b.Awards[j].ItemName
		})
	}
	return bosses, nil
}

// lootItemFacts is an award's item name and quality from the loot table.
func lootItemFacts(itemID int) (string, int) {
	for _, it := range onyxiaLoot {
		if it.ItemID == itemID {
			return it.Name, it.Quality
		}
	}
	return fmt.Sprintf("Item %d", itemID), 0
}

func addLootAward(bosses []ReportLootBoss, encounterID int64, a ReportLootAward) []ReportLootBoss {
	for i := range bosses {
		if bosses[i].EncounterID == encounterID {
			bosses[i].Awards = append(bosses[i].Awards, a)
			return bosses
		}
	}
	boss := fmt.Sprintf("Encounter %d", encounterID)
	if encounterID == onyxiaEncounterID {
		boss = encounterName
	}
	return append(bosses, ReportLootBoss{EncounterID: encounterID, Boss: boss, Awards: []ReportLootAward{a}})
}

// killStat is one raider's best kill on one boss.
type killStat struct {
	key, name, class string
	encounterID      int
	dps, hps         float64
	healer           bool
}

type statKey struct {
	key         string
	encounterID int
}

func (s *Store) reportMoves(ctx context.Context, vc raidViewContext, v RaidView, reportID string, kills []KillOrderView) (MovesView, error) {
	out := MovesView{Rows: []MoveRow{}}
	prev, prevReport, err := s.previousDoneRaid(ctx, vc.guildID, v)
	if err != nil || prev == nil {
		return out, err
	}
	out.Previous = prev
	now, err := s.killStats(ctx, reportID)
	if err != nil {
		return out, err
	}
	before, err := s.killStats(ctx, prevReport)
	if err != nil {
		return out, err
	}
	names := map[int]string{}
	for _, k := range kills {
		names[k.EncounterID] = k.Name
	}
	rows := []MoveRow{}
	for k, cur := range now {
		old, ok := before[k]
		if !ok {
			continue
		}
		if row, moved := moveOf(cur, old, names[cur.encounterID]); moved {
			row.own = vc.ownsKey(cur.key)
			rows = append(rows, row)
		}
	}
	sort.SliceStable(rows, func(i, j int) bool {
		if a, b := math.Abs(rows[i].delta), math.Abs(rows[j].delta); a != b {
			return a > b
		}
		if rows[i].Name != rows[j].Name {
			return rows[i].Name < rows[j].Name
		}
		return rows[i].Encounter < rows[j].Encounter
	})
	if len(rows) > movesCap {
		rows = rows[:movesCap]
	}
	out.Rows = rows
	return out, nil
}

// moveOf is the change of cur against old in the raider's metric (hps for a healer, else
// dps); moved is false for no change.
func moveOf(cur, old killStat, encounter string) (MoveRow, bool) {
	metric, after, before := "dps", cur.dps, old.dps
	if cur.healer {
		metric, after, before = "hps", cur.hps, old.hps
	}
	if after == before {
		return MoveRow{}, false
	}
	return MoveRow{CharacterKey: cur.key, Name: cur.name, Class: cur.class, Metric: metric, Encounter: encounter,
		Before: int(math.Round(before)), After: int(math.Round(after)), delta: after - before}, true
}

func (s *Store) previousDoneRaid(ctx context.Context, guildID int64, v RaidView) (*RaidRefView, string, error) {
	var ref RaidRefView
	var report string
	err := s.Pool.QueryRow(ctx, `
		select id, title, starts_at, report_id from guild_raids
		where guild_id = $1 and zone = $2 and status = 'done' and report_id is not null and starts_at < $3
		order by starts_at desc, id desc limit 1`, guildID, v.Zone, v.StartsAt).
		Scan(&ref.ID, &ref.Title, &ref.StartsAt, &report)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, "", nil
	}
	if err != nil {
		return nil, "", fmt.Errorf("guilds: previous raid: %w", err)
	}
	ref.StartsAt = ref.StartsAt.UTC()
	return &ref, report, nil
}

func (s *Store) killStats(ctx context.Context, reportID string) (map[statKey]killStat, error) {
	rows, err := s.Pool.Query(ctx, `
		select player_key, max(player_name), lower(max(class)), encounter_id,
		       coalesce(max(metric_dps), 0)::float8, coalesce(max(metric_hps), 0)::float8, bool_or(role = 'healer')
		from fight_metrics
		where report_id = $1 and kill and state = 'ok' and encounter_id is not null
		group by player_key, encounter_id`, reportID)
	if err != nil {
		return nil, fmt.Errorf("guilds: report kill stats: %w", err)
	}
	defer rows.Close()
	out := map[statKey]killStat{}
	for rows.Next() {
		var k killStat
		if err := rows.Scan(&k.key, &k.name, &k.class, &k.encounterID, &k.dps, &k.hps, &k.healer); err != nil {
			return nil, fmt.Errorf("guilds: report kill stats: %w", err)
		}
		out[statKey{k.key, k.encounterID}] = k
	}
	return out, rows.Err()
}
