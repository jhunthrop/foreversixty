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

	"github.com/jhunthrop/foreversixty/api/internal/fs1"
	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

// movesCap is how many changes the report lists.
const movesCap = 5

// consumableOrder ranks the Before pills worst first.
var consumableOrder = map[string]int{"short": 0, "unknown": 1, "stocked": 2}

func (s *Store) reportConsumables(ctx context.Context, vc raidViewContext, reportID string) (ReportConsumablesView, error) {
	out := ReportConsumablesView{Rows: []ConsumableRow{}}
	inLog, err := s.reportPlayerKeys(ctx, reportID)
	if err != nil || len(inLog) == 0 {
		return out, err
	}
	keys := make([]string, 0, len(inLog))
	for k := range inLog {
		keys = append(keys, k)
	}
	if out.Rows, err = s.consumableRows(ctx, vc, keys); err != nil {
		return out, err
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

// consumableRows reads the bags state of the guild's verified roster characters among keys in
// one query, with the roster's consent rule: the class and the export are read only for a
// character whose account shares its gear, the class from its latest evidence fight, else from
// the export.
func (s *Store) consumableRows(ctx context.Context, vc raidViewContext, keys []string) ([]ConsumableRow, error) {
	rows, err := s.Pool.Query(ctx, `
		select gc.character_key, ae.name, gm.consent,
		       case when gm.consent in ('gear', 'gear_bags') then fm.class end,
		       case when gm.consent in ('gear', 'gear_bags') then ae.export end
		from guild_characters gc
		join addon_exports ae on ae.character_key = gc.character_key
		join guild_members gm on gm.guild_id = gc.guild_id and gm.user_id = gc.user_id
		left join lateral (
		  select m.class from fight_metrics m join reports r on r.id = m.report_id
		  where m.player_key = gc.character_key and `+presence.EvidenceReport("r")+`
		  order by m.fought_at desc limit 1
		) fm on true
		where gc.guild_id = $1 and gc.verified_at is not null and gc.character_key = any($2)`, vc.guildID, keys)
	if err != nil {
		return nil, fmt.Errorf("guilds: report consumables: %w", err)
	}
	defer rows.Close()
	out := []ConsumableRow{}
	for rows.Next() {
		var c ConsumableRow
		var consent string
		var class, export *string
		if err := rows.Scan(&c.CharacterKey, &c.Name, &consent, &class, &export); err != nil {
			return nil, fmt.Errorf("guilds: report consumables: %w", err)
		}
		var bags []int
		hasBags := false
		if class != nil {
			c.Class = *class
		}
		if export != nil {
			if decoded, ok := fs1.Decode(*export); ok {
				bags, hasBags = decoded.Bags, hasBagsSection(*export)
				if c.Class == "" {
					c.Class = decoded.ClassSlug
				}
			}
		}
		c.Before = consumablesState(consent, bags, hasBags)
		c.own = vc.ownsKey(c.CharacterKey)
		out = append(out, c)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: report consumables: %w", err)
	}
	return out, nil
}
func (s *Store) reportPlayerKeys(ctx context.Context, reportID string) (map[string]bool, error) {
	rows, err := s.Pool.Query(ctx,
		`select distinct m.player_key from fight_metrics m join reports r on r.id = m.report_id
		  where m.report_id = $1 and m.state = 'ok' and `+presence.EvidenceReport("r"), reportID)
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
	type award struct {
		encounterID int64
		a           ReportLootAward
		by          *int64
	}
	var read []award
	var awarders []int64
	for rows.Next() {
		var w award
		var at time.Time
		if err := rows.Scan(&w.encounterID, &w.a.ItemID, &w.a.CharacterKey, &w.a.Name, &w.a.Class, &at, &w.by); err != nil {
			return nil, fmt.Errorf("guilds: report loot: %w", err)
		}
		w.a.At = at.UTC().Format(time.RFC3339)
		w.a.ItemName, w.a.Quality, w.a.Icon = lootItemFacts(w.a.ItemID)
		if w.by != nil {
			awarders = append(awarders, *w.by)
		}
		read = append(read, w)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: report loot: %w", err)
	}
	names, err := s.awarderNames(ctx, vc.guildID, awarders)
	if err != nil {
		return nil, err
	}
	bosses := []ReportLootBoss{}
	for _, w := range read {
		if w.by != nil {
			w.a.ByName = names.of(*w.by)
		}
		bosses = addLootAward(bosses, w.encounterID, w.a)
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

// lootItemFacts is an award's item name, quality and icon from the loot table (the Loot tab's).
func lootItemFacts(itemID int) (string, int, string) {
	for _, it := range onyxiaLoot {
		if it.ItemID == itemID {
			return it.Name, it.Quality, it.Icon
		}
	}
	return fmt.Sprintf("Item %d", itemID), 0, ""
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
	dps, hps         *float64 // null when the log carried no such metric for the kill
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
	out.Rows = topMoves(rows)
	return out, nil
}

// topMoves is the movesCap biggest changes, largest first, ties by name then boss.
func topMoves(rows []MoveRow) []MoveRow {
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
	return rows
}

// moveOf is the change of cur against old in the raider's metric (hps for a healer, else
// dps); moved is false for no change.
// A switch of role between the raids, or a metric the log lacks on either side, is no change.
func moveOf(cur, old killStat, encounter string) (MoveRow, bool) {
	if cur.healer != old.healer {
		return MoveRow{}, false
	}
	metric, after, before := "dps", cur.dps, old.dps
	if cur.healer {
		metric, after, before = "hps", cur.hps, old.hps
	}
	if after == nil || before == nil {
		return MoveRow{}, false
	}
	shown, was := int(math.Round(*after)), int(math.Round(*before))
	if shown == was {
		return MoveRow{}, false
	}
	return MoveRow{CharacterKey: cur.key, Name: cur.name, Class: cur.class, Metric: metric, Encounter: encounter,
		Before: was, After: shown, delta: *after - *before}, true
}

func (s *Store) previousDoneRaid(ctx context.Context, guildID int64, v RaidView) (*RaidRefView, string, error) {
	var ref RaidRefView
	var report string
	err := s.Pool.QueryRow(ctx, `
		select g.id, g.title, g.starts_at, g.report_id from guild_raids g
		join reports r on r.id = g.report_id and `+presence.EvidenceReport("r")+`
		where g.guild_id = $1 and g.zone = $2 and g.status = 'done' and g.starts_at < $3
		order by g.starts_at desc, g.id desc limit 1`, guildID, v.Zone, v.StartsAt).
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
		select m.player_key, coalesce(max(m.player_name), ''), coalesce(lower(max(m.class)), ''), m.encounter_id,
		       max(m.metric_dps)::float8, max(m.metric_hps)::float8, coalesce(bool_or(m.role = 'healer'), false)
		from fight_metrics m join reports r on r.id = m.report_id
		where m.report_id = $1 and m.kill and m.state = 'ok' and m.encounter_id is not null
		  and `+presence.EvidenceReport("r")+`
		group by m.player_key, m.encounter_id`, reportID)
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
