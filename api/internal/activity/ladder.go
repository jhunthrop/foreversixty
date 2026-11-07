package activity

import (
	"context"
	"fmt"
	"time"
)

// LadderMaxAge drops a sim older than this from the guild ladder (ruling R2).
const LadderMaxAge = 30 * 24 * time.Hour

// Ladder sizes (rulings P7, R2): no ladder below ladderMin members in the group, and no
// "#n of N" in a sentence below ladderRankWords.
const (
	ladderMin       = 3
	ladderRankWords = 5
)

// LadderRow is one member's latest sim on a guild ladder.
type LadderRow struct {
	Rank     int       `json:"rank"`
	Name     string    `json:"name"`
	Class    string    `json:"class"`
	Key      string    `json:"key"`
	Value    float64   `json:"value"`
	Delta    *int      `json:"delta"`
	SimmedAt time.Time `json:"simmed_at"`
	SimID    string    `json:"sim_id"`
	You      bool      `json:"you"`
	userID   int64
}

// LadderGroup is what a ladder compares like with like on.
type LadderGroup struct {
	GuildID int64
	Role    string
	Band    int
	Preset  string
}

// ladderSQL is each verified, non-anonymised member's latest finished sim in the group,
// within LadderMaxAge, best first.
const ladderSQL = `with members as (
  select gc.character_key from guild_characters gc
  join characters c on c.key = gc.character_key
  left join users u on u.id = c.user_id
  where gc.guild_id = $1 and gc.verified_at is not null and not coalesce(u.anonymize, false)
), latest as (
  select distinct on (cs.character_key) cs.character_key, cs.sim_id, cs.user_id, sm.dps_mean,
         coalesce(sm.finished_at, sm.created_at) as at
  from character_sims cs
  join members m on m.character_key = cs.character_key
  join sims sm on sm.id = cs.sim_id
  where cs.role = $2 and cs.band = $3 and cs.preset = $4 and sm.state = 'done'
    and coalesce(sm.finished_at, sm.created_at) > $5
  order by cs.character_key, coalesce(sm.finished_at, sm.created_at) desc
)
select l.character_key, c.name, lower(coalesce(c.class, '')), l.dps_mean, l.at, l.sim_id, l.user_id
from latest l join characters c on c.key = l.character_key
order by l.dps_mean desc, l.character_key`

// GuildLadder reads a group's ladder as it stands at now.
func (s *Store) GuildLadder(ctx context.Context, g LadderGroup, now time.Time) ([]LadderRow, error) {
	rows, err := s.Pool.Query(ctx, ladderSQL, g.GuildID, g.Role, g.Band, g.Preset, now.Add(-LadderMaxAge))
	if err != nil {
		return nil, fmt.Errorf("activity: guild ladder %d: %w", g.GuildID, err)
	}
	defer rows.Close()
	out := []LadderRow{}
	for rows.Next() {
		var r LadderRow
		if err := rows.Scan(&r.Key, &r.Name, &r.Class, &r.Value, &r.SimmedAt, &r.SimID, &r.userID); err != nil {
			return nil, fmt.Errorf("activity: scan ladder: %w", err)
		}
		r.Rank = len(out) + 1
		out = append(out, r)
	}
	return out, rows.Err()
}

// RoleWords names a role's members in a sentence ("DPS", "healers", "tanks") and its unit.
func RoleWords(role string) (members, unit string) {
	switch role {
	case "healer":
		return "healers", "HPS"
	case "tank":
		return "tanks", "DPS"
	default:
		return "DPS", "DPS"
	}
}
