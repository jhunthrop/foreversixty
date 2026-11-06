// api/internal/guilds/raid_assign_store.go
//
// The raid_assignments rows (migration 0077): read by raid for the raid view, a slot written
// whole (its names in order) by an officer's PUT, cleared by DELETE, and the re-post throttle
// of "Post to Discord".
package guilds

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// assignmentRow is one stored name of one slot.
type assignmentRow struct {
	RaidID       int64
	EncounterID  int64
	SlotKey      string
	CharacterKey string
}

// slotRef names one slot of one raid's encounter.
type slotRef struct {
	EncounterID int64
	SlotKey     string
}

// assignmentsByRaid reads every stored name of raidIDs by raid, then slot, in position order.
func (s *Store) assignmentsByRaid(ctx context.Context, raidIDs []int64) (map[int64]map[slotRef][]string, error) {
	out := map[int64]map[slotRef][]string{}
	if len(raidIDs) == 0 {
		return out, nil
	}
	rows, err := s.Pool.Query(ctx, `select raid_id, encounter_id, slot_key, character_key from raid_assignments
		where raid_id = any($1) order by raid_id, encounter_id, slot_key, position`, raidIDs)
	if err != nil {
		return nil, fmt.Errorf("guilds: raid assignments: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var a assignmentRow
		if err := rows.Scan(&a.RaidID, &a.EncounterID, &a.SlotKey, &a.CharacterKey); err != nil {
			return nil, fmt.Errorf("guilds: raid assignments: %w", err)
		}
		if out[a.RaidID] == nil {
			out[a.RaidID] = map[slotRef][]string{}
		}
		ref := slotRef{EncounterID: a.EncounterID, SlotKey: a.SlotKey}
		out[a.RaidID][ref] = append(out[a.RaidID][ref], a.CharacterKey)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: raid assignments: %w", err)
	}
	return out, nil
}

// SetSlot replaces the slot's names with keys, in order, in one transaction.
func (s *Store) SetSlot(ctx context.Context, raidID int64, ref slotRef, keys []string, by int64) error {
	return pgx.BeginFunc(ctx, s.Pool, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, `delete from raid_assignments where raid_id = $1 and encounter_id = $2 and slot_key = $3`,
			raidID, ref.EncounterID, ref.SlotKey); err != nil {
			return fmt.Errorf("guilds: set slot: %w", err)
		}
		for i, key := range keys {
			if _, err := tx.Exec(ctx, `insert into raid_assignments
				(raid_id, encounter_id, slot_key, position, character_key, set_by) values ($1, $2, $3, $4, $5, $6)`,
				raidID, ref.EncounterID, ref.SlotKey, i, key, by); err != nil {
				return fmt.Errorf("guilds: set slot: %w", err)
			}
		}
		return nil
	})
}

// ClearSlot removes every name of the slot.
func (s *Store) ClearSlot(ctx context.Context, raidID int64, ref slotRef) error {
	if _, err := s.Pool.Exec(ctx, `delete from raid_assignments where raid_id = $1 and encounter_id = $2 and slot_key = $3`,
		raidID, ref.EncounterID, ref.SlotKey); err != nil {
		return fmt.Errorf("guilds: clear slot: %w", err)
	}
	return nil
}

// lastAssignmentPosts reads when each of raidIDs last posted its assignments; never posted is absent.
func (s *Store) lastAssignmentPosts(ctx context.Context, raidIDs []int64) (map[int64]time.Time, error) {
	out := map[int64]time.Time{}
	if len(raidIDs) == 0 {
		return out, nil
	}
	rows, err := s.Pool.Query(ctx, `select raid_id, last_posted_at from raid_assignment_posts where raid_id = any($1)`, raidIDs)
	if err != nil {
		return nil, fmt.Errorf("guilds: assignment posts: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		var at time.Time
		if err := rows.Scan(&id, &at); err != nil {
			return nil, fmt.Errorf("guilds: assignment posts: %w", err)
		}
		out[id] = at
	}
	return out, rows.Err()
}

// claimAssignmentPost stamps the raid's post at at when its last one is every old or there is
// none; claimed is false otherwise. prev is the last post before this call (nil: never).
func (s *Store) claimAssignmentPost(ctx context.Context, raidID, by int64, at time.Time, every time.Duration) (bool, *time.Time, error) {
	var claimed bool
	var prev *time.Time
	err := s.Pool.QueryRow(ctx, `
		with prev as (select last_posted_at from raid_assignment_posts where raid_id = $1),
		up as (
		  insert into raid_assignment_posts (raid_id, last_posted_at, posted_by) values ($1, $2, $3)
		  on conflict (raid_id) do update set last_posted_at = excluded.last_posted_at, posted_by = excluded.posted_by
		  where raid_assignment_posts.last_posted_at <= $2::timestamptz - make_interval(secs => $4)
		  returning 1)
		select exists (select 1 from up), (select last_posted_at from prev)`,
		raidID, at, by, every.Seconds()).Scan(&claimed, &prev)
	if err != nil {
		return false, nil, fmt.Errorf("guilds: claim assignment post: %w", err)
	}
	return claimed, prev, nil
}

// releaseAssignmentPost gives back a claim made at at: the row returns to prev, or goes.
func (s *Store) releaseAssignmentPost(ctx context.Context, raidID int64, at time.Time, prev *time.Time) error {
	var err error
	if prev == nil {
		_, err = s.Pool.Exec(ctx, `delete from raid_assignment_posts where raid_id = $1 and last_posted_at = $2`, raidID, at)
	} else {
		_, err = s.Pool.Exec(ctx, `update raid_assignment_posts set last_posted_at = $3
			where raid_id = $1 and last_posted_at = $2`, raidID, at, *prev)
	}
	if err != nil {
		return fmt.Errorf("guilds: release assignment post: %w", err)
	}
	return nil
}
