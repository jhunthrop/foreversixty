// api/internal/guilds/raid_night_log.go
//
// The raid night's after half, read from the report matched to the raid: the night's boss kills
// and wipes, your figure on each kill (the parse source's DPS, HPS for a healer, with the fight's
// Forever rating) and how many beat your best on that boss before the night, the loot awarded to
// your character in the report, and the coach teaser.
package guilds

import (
	"context"
	"fmt"
	"math"
)

// fillRaidNightLog fills out's log fields from reportID for char.
func (s *Service) fillRaidNightLog(ctx context.Context, out *RaidNightSummary, viewerID, guildID int64, reportID string, char ChecklistCharacter) error {
	out.ReportID = &reportID
	kills, wipes, err := s.Store.reportBossPulls(ctx, reportID)
	if err != nil {
		return err
	}
	out.Kills, out.Wipes = &kills, &wipes
	bests, err := s.Store.nightBests(ctx, reportID, char.CharacterKey)
	if err != nil {
		return err
	}
	newBests := 0
	for _, b := range bests {
		if b.newBest {
			newBests++
		}
	}
	out.Bests, out.NewBests = bests, &newBests
	if out.Loot, err = s.nightLootOf(ctx, guildID, reportID, char); err != nil {
		return err
	}
	if s.Coach != nil {
		notes, err := s.Coach.CoachNotes(ctx, viewerID, reportID, char.CharacterKey)
		if err != nil {
			s.logger().Warn("guilds", "op", "raid_night_coach", "report", reportID, "err", err)
		} else {
			out.CoachNotes = notes
		}
	}
	return nil
}

// reportBossPulls counts the report's boss fights: kills and wipes.
func (s *Store) reportBossPulls(ctx context.Context, reportID string) (kills, wipes int, err error) {
	err = s.Pool.QueryRow(ctx, `select count(*) filter (where kill), count(*) filter (where not kill)
		from fights where report_id = $1 and coalesce(encounter_id, 0) > 0`, reportID).Scan(&kills, &wipes)
	if err != nil {
		return 0, 0, fmt.Errorf("guilds: report pulls: %w", err)
	}
	return kills, wipes, nil
}

// nightBestsSQL is the character's boss kills in the report with the fight's rating, and its best
// figure on the same boss in earlier kills.
var nightBestsSQL = `
select coalesce(f.name, ''), m.encounter_id, coalesce(m.role, ''), m.metric_dps::float8, m.metric_hps::float8,
       coalesce(m.duration_ms, f.duration_ms, 0),
       (select rs.overall::float8 from rating_scores rs
         where rs.report_id = m.report_id and rs.fight_index = m.fight_index and rs.player_key = m.player_key
           and not coalesce(rs.insufficient, false) limit 1),
       (select max(p.metric_dps)::float8 from fight_metrics p where p.player_key = m.player_key
          and p.encounter_id = m.encounter_id and p.kill and p.state <> 'removed' and p.fought_at < m.fought_at),
       (select max(p.metric_hps)::float8 from fight_metrics p where p.player_key = m.player_key
          and p.encounter_id = m.encounter_id and p.kill and p.state <> 'removed' and p.fought_at < m.fought_at)
from fight_metrics m
join fights f on f.report_id = m.report_id and f.fight_index = m.fight_index
where m.report_id = $1 and m.player_key = $2 and ` + bossKillMetrics + `
order by m.fight_index`

// nightBests is the character's figure on each boss kill of the report, and whether it beat
// every earlier kill of that boss.
func (s *Store) nightBests(ctx context.Context, reportID, key string) ([]RaidNightBest, error) {
	rows, err := s.Pool.Query(ctx, nightBestsSQL, reportID, key)
	if err != nil {
		return nil, fmt.Errorf("guilds: night bests: %w", err)
	}
	defer rows.Close()
	out := []RaidNightBest{}
	for rows.Next() {
		var b RaidNightBest
		var encounter int64
		var role string
		var dps, hps, rating, priorDPS, priorHPS *float64
		if err := rows.Scan(&b.Boss, &encounter, &role, &dps, &hps, &b.KillMs, &rating, &priorDPS, &priorHPS); err != nil {
			return nil, fmt.Errorf("guilds: night bests: %w", err)
		}
		if b.Boss == "" {
			b.Boss = bossName(encounter)
		}
		value, prior := dps, priorDPS
		if roleMetric(role) == "hps" {
			value, prior = hps, priorHPS
		}
		if value != nil {
			b.DPS = *value
		}
		b.newBest = prior != nil && b.DPS > *prior
		if rating != nil {
			p := int(math.Round(*rating))
			b.Percentile = &p
		}
		out = append(out, b)
	}
	return out, rows.Err()
}

// nightLootOf is the loot awarded to char in the report, each item read from the loot table
// (else the class's item table) with its boss and the awarding officer's character.
func (s *Service) nightLootOf(ctx context.Context, guildID int64, reportID string, char ChecklistCharacter) ([]RaidNightLoot, error) {
	rows, err := s.Store.Pool.Query(ctx, `select item_id, encounter_id, awarded_by from loot_awards
		where guild_id = $1 and report_id = $2 and character_key = $3 order by id`, guildID, reportID, char.CharacterKey)
	if err != nil {
		return nil, fmt.Errorf("guilds: night loot: %w", err)
	}
	type award struct {
		item      int
		encounter int64
		by        *int64
	}
	var awards []award
	for rows.Next() {
		var a award
		if err := rows.Scan(&a.item, &a.encounter, &a.by); err != nil {
			rows.Close()
			return nil, fmt.Errorf("guilds: night loot: %w", err)
		}
		awards = append(awards, a)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: night loot: %w", err)
	}
	out := make([]RaidNightLoot, 0, len(awards))
	for _, a := range awards {
		l := s.Store.nightItem(a.item, char.Class)
		l.Boss = bossName(a.encounter)
		if a.by != nil {
			ref, err := s.Store.claimHolderRef(ctx, guildID, *a.by)
			if err != nil {
				return nil, err
			}
			if ref != nil {
				l.AwardedBy = ref.Name
			}
		}
		out = append(out, l)
	}
	return out, nil
}

// nightItem is an item's name, quality, icon and slot words: the loot table's row, else the
// class's item table's.
func (s *Store) nightItem(itemID int, class string) RaidNightLoot {
	l := RaidNightLoot{ItemID: itemID}
	if it, ok := lootItemByID(itemID); ok {
		l.Name, l.Quality, l.Icon, l.Slot = it.Name, it.Quality, it.Icon, clientSlotWords[it.InventoryType]
		return l
	}
	if it, ok := s.ItemIn(s.BisBuild(), class, itemID); ok {
		l.Name, l.Quality, l.Icon, l.Slot = it.Name, it.Quality, it.Icon, it.Slot
	}
	return l
}
