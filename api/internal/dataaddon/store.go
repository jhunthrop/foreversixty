// api/internal/dataaddon/store.go
package dataaddon

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Store is every database read this job needs.
type Store struct {
	Pool *pgxpool.Pool
}

// characterFightRow is one public, in-window, non-anonymized, sufficient
// rated fight -- "sufficient" meaning the engine did not mark it
// insufficient (migration 0023: a card built from under half a role's
// weight table gets Overall = 0, and that 0 is not a score).
type characterFightRow struct {
	PlayerKey  string
	Overall    float64
	Components json.RawMessage
	FoughtAt   time.Time
}

// characterFights reads every rating_scores row fought_at or after since,
// whose report is public, for a player_key whose owning account has not
// asked to be anonymized -- the same rule api/internal/rating.Store's
// anonymized and ReadCharacterRating apply per character, applied here
// across every character at once. A player_key with no characters row at
// all (never linked to an account) is not anonymized, matching
// rating.Store.anonymized's own "nothing to hide" rule exactly. It also
// excludes insufficient cards at this same SQL layer, with the identical
// predicate rating.Store.ReadCharacterRating uses: `not coalesce(rs.
// insufficient, false)`. A NULL insufficient (a row written before
// migration 0023 added the column) reads as sufficient, matching every
// other stale-row gap in this codebase.
func (s *Store) characterFights(ctx context.Context, since time.Time) ([]characterFightRow, error) {
	rows, err := s.Pool.Query(ctx,
		`select rs.player_key, rs.overall, rs.components, rs.fought_at
		 from rating_scores rs
		 join reports r on r.id = rs.report_id
		 where r.visibility = 'public' and rs.fought_at >= $1
		   and not coalesce(rs.insufficient, false)
		   and not exists (
		     select 1 from characters c join users u on u.id = c.user_id
		     where c.key = rs.player_key and u.anonymize
		   )
		 order by rs.player_key`, since)
	if err != nil {
		return nil, fmt.Errorf("dataaddon: read character fights: %w", err)
	}
	defer rows.Close()
	var out []characterFightRow
	for rows.Next() {
		var r characterFightRow
		if err := rows.Scan(&r.PlayerKey, &r.Overall, &r.Components, &r.FoughtAt); err != nil {
			return nil, fmt.Errorf("dataaddon: scan character fight: %w", err)
		}
		out = append(out, r)
	}
	return out, rows.Err()
}

// guildsWithVerifiedMembers reads every guild that has at least one
// verified guild_characters row -- a guild with none has nothing
// corroborated to publish (dispatch: "only VERIFIED members count"), so it
// is left out of this list entirely and never gets a Data.lua row.
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

// verifiedMembers reads every verified guild_characters row for the given
// guild ids, keyed by guild_id.
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

// nightsByGuild counts, per guild id, the distinct UTC calendar dates of
// that guild's public reports created at or after since -- this job's own
// definition of "a raid night" (see the plan's aggregation-rule section).
func (s *Store) nightsByGuild(ctx context.Context, guildIDs []int64, since time.Time) (map[int64]int, error) {
	rows, err := s.Pool.Query(ctx,
		`select guild_id, count(distinct date(created_at))
		 from reports
		 where guild_id = any($1) and visibility = 'public' and created_at >= $2
		 group by guild_id`, guildIDs, since)
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

// progressionByGuild counts, per guild id, encounters killed at least once
// against every encounter attempted, from fights attached to a report the
// guild's own public page would show
// (api/internal/rankings/guilds.go's own Guild() progression query:
// r.visibility <> 'private'). Duplicated here rather than imported:
// api/internal/rankings is outside this lane's file ownership, and
// api/internal/rating/backfill.go's own readSummary already sets the
// precedent in this codebase for duplicating one small, owned query across
// a lane boundary rather than reaching across it.
func (s *Store) progressionByGuild(ctx context.Context, guildIDs []int64) (killed, total map[int64]int, err error) {
	rows, err := s.Pool.Query(ctx,
		`select r.guild_id, f.encounter_id, count(*) filter (where f.kill)
		 from fights f join reports r on r.id = f.report_id
		 where r.guild_id = any($1) and f.encounter_id is not null and r.visibility <> 'private'
		 group by r.guild_id, f.encounter_id`, guildIDs)
	if err != nil {
		return nil, nil, fmt.Errorf("dataaddon: read guild progression: %w", err)
	}
	defer rows.Close()
	killed, total = map[int64]int{}, map[int64]int{}
	for rows.Next() {
		var id, encounterID int64
		var kills int
		if err := rows.Scan(&id, &encounterID, &kills); err != nil {
			return nil, nil, fmt.Errorf("dataaddon: scan guild progression: %w", err)
		}
		total[id]++
		if kills > 0 {
			killed[id]++
		}
	}
	return killed, total, rows.Err()
}
