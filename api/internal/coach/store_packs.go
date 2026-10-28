package coach

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

const savePackSQL = `
insert into coach_packs (report_id, fight_index, player_key, pack_version, pack)
values ($1, $2, $3, $4, $5)
on conflict (report_id, fight_index, player_key, pack_version) do update
  set pack = excluded.pack, built_at = now()`

const packSQL = `
select pack from coach_packs
where report_id = $1 and fight_index = $2 and player_key = $3 and pack_version = $4`

// latestPackSQL is the newest pack of any version: its context names the spec whose rule
// file decides the current version, so a hit needs no bucket read.
const latestPackSQL = `
select pack from coach_packs
where report_id = $1 and fight_index = $2 and player_key = $3
order by built_at desc, pack_version
limit 1`

// SavePack caches p under its version. A pack without a version is refused: the version
// is the caller's (PackVersion) and an empty one would never be found again.
func (s *Store) SavePack(ctx context.Context, reportID string, fightIndex int, playerKey string, p Pack) error {
	if p.Version == "" {
		return errors.New("coach: a pack needs its version before it is cached")
	}
	raw, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("coach: encode pack: %w", err)
	}
	if _, err := s.DB.Exec(ctx, savePackSQL, reportID, fightIndex, playerKey, p.Version, raw); err != nil {
		return fmt.Errorf("coach: save pack %s/%d: %w", reportID, fightIndex, err)
	}
	return nil
}

// Pack reads the cached pack of one version; false when there is none.
func (s *Store) Pack(ctx context.Context, reportID string, fightIndex int, playerKey, version string) (Pack, bool, error) {
	return s.readPack(ctx, packSQL, reportID, fightIndex, playerKey, version)
}

// latestPack reads the newest cached pack of any version; false when there is none.
func (s *Store) latestPack(ctx context.Context, reportID string, fightIndex int, playerKey string) (Pack, bool, error) {
	return s.readPack(ctx, latestPackSQL, reportID, fightIndex, playerKey)
}

func (s *Store) readPack(ctx context.Context, sql, reportID string, fightIndex int, args ...any) (Pack, bool, error) {
	var raw []byte
	err := s.DB.QueryRow(ctx, sql, append([]any{reportID, fightIndex}, args...)...).Scan(&raw)
	if errors.Is(err, pgx.ErrNoRows) {
		return Pack{}, false, nil
	}
	if err != nil {
		return Pack{}, false, fmt.Errorf("coach: read pack %s/%d: %w", reportID, fightIndex, err)
	}
	var p Pack
	if err := json.Unmarshal(raw, &p); err != nil {
		return Pack{}, false, fmt.Errorf("coach: decode pack %s/%d: %w", reportID, fightIndex, err)
	}
	return p, true, nil
}
