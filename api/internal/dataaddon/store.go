// api/internal/dataaddon/store.go
package dataaddon

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

// Store is every database read this job makes itself; the scores come from rating.Store.
type Store struct {
	Pool *pgxpool.Pool
}

// windowParses reads, for every character with a rated fight in the window, the best
// overall percentile it has on each boss there: player_key -> boss name -> best overall.
// A rated fight counts when its report is public and not flagged (presence.EvidenceReport)
// and the engine did not mark the card insufficient (migration 0023: an insufficient card's
// overall is a 0 that is not a score; a NULL reads as sufficient, as every other read of the
// column does). The boss's name is the fight's own name, else the newest name any fight
// recorded for the encounter (the rule the inbox's last_raid names bosses by); a fight with
// no name at all still puts the character in the window but adds no boss. Anonymised
// characters are left out by the caller, with rating.Store.Anonymized.
func (s *Store) windowParses(ctx context.Context, since time.Time) (map[string]map[string]float64, error) {
	rows, err := s.Pool.Query(ctx,
		`with rated as (
		   select rs.player_key, rs.overall,
		          coalesce(nullif(f.name, ''),
		            (select f2.name from fights f2 where f2.encounter_id = rs.encounter_id and f2.name <> ''
		              order by f2.start_ms desc limit 1), '') as boss
		     from rating_scores rs
		     join reports r on r.id = rs.report_id
		     left join fights f on f.report_id = rs.report_id and f.fight_index = rs.fight_index
		    where r.visibility = 'public' and `+presence.EvidenceReport("r")+` and rs.fought_at >= $1
		      and not coalesce(rs.insufficient, false)
		 )
		 select player_key, boss, max(coalesce(overall, 0))::float8
		   from rated group by player_key, boss
		  order by player_key, boss`, since)
	if err != nil {
		return nil, fmt.Errorf("dataaddon: read rated fights: %w", err)
	}
	defer rows.Close()
	out := map[string]map[string]float64{}
	for rows.Next() {
		var key, boss string
		var best float64
		if err := rows.Scan(&key, &boss, &best); err != nil {
			return nil, fmt.Errorf("dataaddon: scan rated fight: %w", err)
		}
		if out[key] == nil {
			out[key] = map[string]float64{}
		}
		if boss != "" {
			out[key][boss] = best
		}
	}
	return out, rows.Err()
}

// identities reads who each of keys is in game (identity): the characters row's display
// name and Battle.net realm, and the newest rated row's player_name as the log printed it.
// A key with none of them reads as an empty identity.
func (s *Store) identities(ctx context.Context, keys []string) (map[string]identity, error) {
	rows, err := s.Pool.Query(ctx,
		`select k.key, coalesce(c.name, ''), coalesce(c.realm_name, ''),
		   coalesce((select rs.player_name from rating_scores rs where rs.player_key = k.key and rs.player_name <> ''
		     order by rs.fought_at desc limit 1), '')
		 from unnest($1::text[]) as k(key)
		 left join characters c on c.key = k.key`, keys)
	if err != nil {
		return nil, fmt.Errorf("dataaddon: read identities: %w", err)
	}
	defer rows.Close()
	out := map[string]identity{}
	for rows.Next() {
		var key string
		var id identity
		if err := rows.Scan(&key, &id.Name, &id.RealmName, &id.PlayerName); err != nil {
			return nil, fmt.Errorf("dataaddon: scan identity: %w", err)
		}
		out[key] = id
	}
	return out, rows.Err()
}

// exports reads the stored addon export of each of keys that has one.
func (s *Store) exports(ctx context.Context, keys []string) (map[string]string, error) {
	rows, err := s.Pool.Query(ctx,
		`select character_key, export from addon_exports where character_key = any($1)`, keys)
	if err != nil {
		return nil, fmt.Errorf("dataaddon: read exports: %w", err)
	}
	defer rows.Close()
	out := map[string]string{}
	for rows.Next() {
		var key, export string
		if err := rows.Scan(&key, &export); err != nil {
			return nil, fmt.Errorf("dataaddon: scan export: %w", err)
		}
		out[key] = export
	}
	return out, rows.Err()
}

// guildsWithVerifiedMembers reads every guild that has at least one verified
// guild_characters row -- a guild with none has nothing corroborated to publish, so it
// is left out of this list entirely and never gets a record.
func (s *Store) guildsWithVerifiedMembers(ctx context.Context) ([]guildIdentity, error) {
	rows, err := s.Pool.Query(ctx,
		`select g.id, g.region, g.ruleset, g.name
		 from guilds g
		 where exists (
		   select 1 from guild_characters gc
		   where gc.guild_id = g.id and gc.verified_at is not null
		 )
		 order by g.id`)
	if err != nil {
		return nil, fmt.Errorf("dataaddon: read guilds: %w", err)
	}
	defer rows.Close()
	var out []guildIdentity
	for rows.Next() {
		var g guildIdentity
		if err := rows.Scan(&g.ID, &g.Region, &g.Ruleset, &g.Name); err != nil {
			return nil, fmt.Errorf("dataaddon: scan guild: %w", err)
		}
		out = append(out, g)
	}
	return out, rows.Err()
}

// verifiedMembers reads every verified guild_characters row for the given guild ids,
// keyed by guild_id.
func (s *Store) verifiedMembers(ctx context.Context, guildIDs []int64) (map[int64][]string, error) {
	rows, err := s.Pool.Query(ctx,
		`select guild_id, character_key from guild_characters
		 where verified_at is not null and guild_id = any($1)
		 order by guild_id, character_key`, guildIDs)
	if err != nil {
		return nil, fmt.Errorf("dataaddon: read verified members: %w", err)
	}
	defer rows.Close()
	out := map[int64][]string{}
	for rows.Next() {
		var id int64
		var key string
		if err := rows.Scan(&id, &key); err != nil {
			return nil, fmt.Errorf("dataaddon: scan verified member: %w", err)
		}
		out[id] = append(out[id], key)
	}
	return out, rows.Err()
}

// nightsByGuild counts, per guild id, the distinct UTC calendar dates of that guild's
// public reports created at or after since -- this job's own definition of "a raid
// night". A flagged report (presence.EvidenceReport) is no night.
func (s *Store) nightsByGuild(ctx context.Context, guildIDs []int64, since time.Time) (map[int64]int, error) {
	rows, err := s.Pool.Query(ctx,
		`select r.guild_id, count(distinct (r.created_at at time zone 'UTC')::date)
		 from reports r
		 where r.guild_id = any($1) and r.visibility = 'public' and `+presence.EvidenceReport("r")+`
		   and r.created_at >= $2
		 group by r.guild_id`, guildIDs, since)
	if err != nil {
		return nil, fmt.Errorf("dataaddon: read guild nights: %w", err)
	}
	defer rows.Close()
	out := map[int64]int{}
	for rows.Next() {
		var id int64
		var n int
		if err := rows.Scan(&id, &n); err != nil {
			return nil, fmt.Errorf("dataaddon: scan guild nights: %w", err)
		}
		out[id] = n
	}
	return out, rows.Err()
}

// bossesKilledByGuild counts, per guild id, the distinct encounters killed at least once
// in the window, from fights attached to a report that is evidence (presence.EvidenceReport:
// not private, not flagged), limited to reports created at or after since, so a forged report
// cannot lift a guild's standing. The guild page's own progression (api/internal/rankings/
// guilds.go's Guild()) reads the same rule.
func (s *Store) bossesKilledByGuild(ctx context.Context, guildIDs []int64, since time.Time) (map[int64]int, error) {
	rows, err := s.Pool.Query(ctx,
		`select r.guild_id, count(distinct f.encounter_id)
		 from fights f join reports r on r.id = f.report_id
		 where r.guild_id = any($1) and f.encounter_id is not null and f.kill
		   and `+presence.EvidenceReport("r")+` and r.created_at >= $2
		 group by r.guild_id`, guildIDs, since)
	if err != nil {
		return nil, fmt.Errorf("dataaddon: read guild kills: %w", err)
	}
	defer rows.Close()
	out := map[int64]int{}
	for rows.Next() {
		var id int64
		var n int
		if err := rows.Scan(&id, &n); err != nil {
			return nil, fmt.Errorf("dataaddon: scan guild kills: %w", err)
		}
		out[id] = n
	}
	return out, rows.Err()
}
