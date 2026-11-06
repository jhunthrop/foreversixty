// api/internal/guilds/raids_attendance_override.go
//
// An officer's manual attendance fix (spec 2026-11-05-web-layout §7.4, migration 0074). The fix
// lives in guild_raid_attendance_overrides; guild_raid_attendance keeps the effective outcome and
// excused flag, so every reader (raid view, roster nights, CSV export, Discord reads, the recent
// grid) reads the fix with no change of its own. writeAttendanceRows, the one writer of computed
// rows, calls applyOverrides after every recompute, so a fix survives a re-read of the log.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// The statuses an officer can set (guild_raid_attendance_overrides.status).
const (
	OverridePresent = "present"
	OverrideAbsent  = "absent"
	OverrideBench   = "bench"
	OverrideExcused = "excused"
)

// ErrNoLog is a write that needs the raid's matched log, on a raid that has none.
var ErrNoLog = errors.New("guilds: the raid has no matched log")

// validOverride reports whether status is one an officer can set.
func validOverride(status string) bool {
	switch status {
	case OverridePresent, OverrideAbsent, OverrideBench, OverrideExcused:
		return true
	}
	return false
}

// AttendanceOverride is the officer's fix on one row, as stored.
type AttendanceOverride struct {
	Status      string
	AutoOutcome string
	AutoExcused bool
	SetBy       *int64
	SetAt       time.Time
}

// effectiveOverrideSQL writes the effective outcome and excused flag of every fixed row of raid
// $1. Present and bench earn the night (a benched raider gets the credit, as a raid leader
// gives it); absent is a no-show for a raider who signed In or Late, else plain absent; excused
// keeps the log's outcome and sets the flag, so the night counts for nobody.
const effectiveOverrideSQL = `
	update guild_raid_attendance a set
	  outcome = case o.status
	    when 'present' then 'present'
	    when 'bench' then 'present'
	    when 'absent' then case when a.signed in ('in', 'late') then 'no_show' else 'absent' end
	    else o.auto_outcome end,
	  excused = (o.status = 'excused')
	from guild_raid_attendance_overrides o
	where o.raid_id = a.raid_id and o.character_key = a.character_key and a.raid_id = $1`

// applyOverrides runs inside writeAttendanceRows once the computed rows are written: a fixed
// character the log and the sign-ups no longer name keeps a bare row, every fix's auto value is
// refreshed from the rows just computed, then the effective values are written.
func applyOverrides(ctx context.Context, tx pgx.Tx, raidID int64, computedKeys []string, pullsTotal int) error {
	if _, err := tx.Exec(ctx, `
		update guild_raid_attendance a set outcome = 'absent', signed = null, role = null, came_as = null,
		       came_as_key = null, first_seen_offset_s = null, pulls_present = 0, pulls_total = $3, computed_at = now()
		where a.raid_id = $1 and not (a.character_key = any($2))
		  and exists (select 1 from guild_raid_attendance_overrides o
		              where o.raid_id = a.raid_id and o.character_key = a.character_key)`,
		raidID, computedKeys, pullsTotal); err != nil {
		return fmt.Errorf("guilds: attendance overrides: bare rows: %w", err)
	}
	if _, err := tx.Exec(ctx, `
		update guild_raid_attendance_overrides o set auto_outcome = a.outcome
		from guild_raid_attendance a
		where a.raid_id = o.raid_id and a.character_key = o.character_key and o.raid_id = $1`, raidID); err != nil {
		return fmt.Errorf("guilds: attendance overrides: refresh auto: %w", err)
	}
	if _, err := tx.Exec(ctx, effectiveOverrideSQL, raidID); err != nil {
		return fmt.Errorf("guilds: attendance overrides: apply: %w", err)
	}
	return nil
}

// dropOverrides removes every fix of raidID: the night they fixed is gone (the match was
// cleared).
func dropOverrides(ctx context.Context, tx pgx.Tx, raidID int64) error {
	if _, err := tx.Exec(ctx, `delete from guild_raid_attendance_overrides where raid_id = $1`, raidID); err != nil {
		return fmt.Errorf("guilds: drop attendance overrides: %w", err)
	}
	return nil
}

// SetAttendanceOverride fixes one character's night on raidID of guildID. A verified roster
// character with no row (not signed up, not in the log) gets a bare one first. ErrNoLog when
// the raid holds no log; ErrNotFound when the character is not a verified roster character.
func (s *Store) SetAttendanceOverride(ctx context.Context, guildID, raidID int64, key, status string, by int64) error {
	if !validOverride(status) {
		return fmt.Errorf("guilds: attendance override: invalid status %q", status)
	}
	return s.inRaidTx(ctx, guildID, raidID, func(tx pgx.Tx) error {
		if err := ensureAttendanceRow(ctx, tx, guildID, raidID, key); err != nil {
			return err
		}
		if _, err := tx.Exec(ctx, `
			insert into guild_raid_attendance_overrides (raid_id, character_key, status, auto_outcome, auto_excused, set_by)
			select a.raid_id, a.character_key, $3, a.outcome, a.excused, $4
			from guild_raid_attendance a where a.raid_id = $1 and a.character_key = $2
			on conflict (raid_id, character_key) do update
			  set status = excluded.status, set_by = excluded.set_by, set_at = now()`,
			raidID, key, status, by); err != nil {
			return fmt.Errorf("guilds: attendance override: write: %w", err)
		}
		if _, err := tx.Exec(ctx, effectiveOverrideSQL, raidID); err != nil {
			return fmt.Errorf("guilds: attendance override: apply: %w", err)
		}
		return nil
	})
}

// ClearAttendanceOverride drops one character's fix and puts back the auto value. Clearing a
// row with no fix is a no-op. ErrNoLog when the raid holds no log.
func (s *Store) ClearAttendanceOverride(ctx context.Context, guildID, raidID int64, key string) error {
	return s.inRaidTx(ctx, guildID, raidID, func(tx pgx.Tx) error { return clearOverride(ctx, tx, raidID, key) })
}

// clearOverride restores the row under a fix. A bare row (no sign-up, not in the log: the
// recompute never writes one) goes with its fix.
func clearOverride(ctx context.Context, tx pgx.Tx, raidID int64, key string) error {
	var auto string
	var autoExcused bool
	err := tx.QueryRow(ctx, `
		delete from guild_raid_attendance_overrides where raid_id = $1 and character_key = $2
		returning auto_outcome, auto_excused`, raidID, key).Scan(&auto, &autoExcused)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("guilds: clear attendance override: %w", err)
	}
	if _, err := tx.Exec(ctx, `
		update guild_raid_attendance set outcome = $3, excused = $4 where raid_id = $1 and character_key = $2`,
		raidID, key, auto, autoExcused); err != nil {
		return fmt.Errorf("guilds: clear attendance override: restore: %w", err)
	}
	if _, err := tx.Exec(ctx, `
		delete from guild_raid_attendance where raid_id = $1 and character_key = $2
		  and signed is null and outcome = 'absent' and pulls_present = 0 and not excused`, raidID, key); err != nil {
		return fmt.Errorf("guilds: clear attendance override: bare row: %w", err)
	}
	return nil
}

// inRaidTx runs fn in a transaction holding raidID's lock, the lock a recompute takes, so a fix
// and a recompute never interleave. ErrNotFound for a raid not of guildID; ErrNoLog for a raid
// with no log.
func (s *Store) inRaidTx(ctx context.Context, guildID, raidID int64, fn func(pgx.Tx) error) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("guilds: attendance override: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	var report *string
	err = tx.QueryRow(ctx, `select report_id from guild_raids where id = $1 and guild_id = $2 for update`,
		raidID, guildID).Scan(&report)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotFound
	}
	if err != nil {
		return fmt.Errorf("guilds: attendance override: lock raid: %w", err)
	}
	if report == nil {
		return ErrNoLog
	}
	if err := fn(tx); err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("guilds: attendance override: commit: %w", err)
	}
	return nil
}

// ensureAttendanceRow gives a verified roster character with no row on raidID a bare one.
func ensureAttendanceRow(ctx context.Context, tx pgx.Tx, guildID, raidID int64, key string) error {
	var exists bool
	if err := tx.QueryRow(ctx, `select exists (select 1 from guild_raid_attendance where raid_id = $1 and character_key = $2)`,
		raidID, key).Scan(&exists); err != nil {
		return fmt.Errorf("guilds: attendance override: row: %w", err)
	}
	if exists {
		return nil
	}
	chars, err := raidRosterOn(ctx, tx, guildID, key)
	if err != nil {
		return err
	}
	if len(chars) == 0 || !chars[0].verified {
		return ErrNotFound
	}
	c := chars[0]
	if _, err := tx.Exec(ctx, `
		insert into guild_raid_attendance (raid_id, character_key, name, class, spec, outcome, pulls_total)
		values ($1, $2, $3, $4, $5, 'absent',
		        coalesce((select max(pulls_total) from guild_raid_attendance where raid_id = $1), 0))`,
		raidID, key, c.Name, c.Class, c.Spec); err != nil {
		return fmt.Errorf("guilds: attendance override: bare row: %w", err)
	}
	return nil
}

// attendanceOverridesFor reads every fix of the raids in raidIDs, by raid then character key.
func (s *Store) attendanceOverridesFor(ctx context.Context, raidIDs []int64) (map[int64]map[string]AttendanceOverride, error) {
	rows, err := s.Pool.Query(ctx, `
		select raid_id, character_key, status, auto_outcome, auto_excused, set_by, set_at
		from guild_raid_attendance_overrides where raid_id = any($1)`, raidIDs)
	if err != nil {
		return nil, fmt.Errorf("guilds: attendance overrides: %w", err)
	}
	defer rows.Close()
	out := map[int64]map[string]AttendanceOverride{}
	for rows.Next() {
		var raidID int64
		var key string
		var o AttendanceOverride
		if err := rows.Scan(&raidID, &key, &o.Status, &o.AutoOutcome, &o.AutoExcused, &o.SetBy, &o.SetAt); err != nil {
			return nil, fmt.Errorf("guilds: attendance overrides: %w", err)
		}
		if out[raidID] == nil {
			out[raidID] = map[string]AttendanceOverride{}
		}
		o.SetAt = o.SetAt.UTC()
		out[raidID][key] = o
	}
	return out, rows.Err()
}

// overrideView is one fix as the raid view shows it, with the auto value under it.
type overrideView struct {
	view        AttendanceOverrideView
	autoOutcome string
	autoExcused bool
}

// overrideViewsFor is attendanceOverridesFor with each fixer's name (the loot record's rule).
func (s *Store) overrideViewsFor(ctx context.Context, guildID int64, raidIDs []int64) (map[int64]map[string]overrideView, error) {
	out := map[int64]map[string]overrideView{}
	if len(raidIDs) == 0 {
		return out, nil
	}
	fixes, err := s.attendanceOverridesFor(ctx, raidIDs)
	if err != nil {
		return nil, err
	}
	var by []int64
	for _, raid := range fixes {
		for _, o := range raid {
			if o.SetBy != nil {
				by = append(by, *o.SetBy)
			}
		}
	}
	names, err := s.awarderNames(ctx, guildID, by)
	if err != nil {
		return nil, err
	}
	for raidID, raid := range fixes {
		out[raidID] = map[string]overrideView{}
		for key, o := range raid {
			name := anOfficer
			if o.SetBy != nil {
				name = names.of(*o.SetBy)
			}
			out[raidID][key] = overrideView{view: AttendanceOverrideView{Status: o.Status, ByName: name, At: o.SetAt},
				autoOutcome: o.AutoOutcome, autoExcused: o.AutoExcused}
		}
	}
	return out, nil
}
