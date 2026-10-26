// api/internal/guilds/raids_store.go
//
// The guild raid calendar's data layer (migration 0050): scheduled raids and the guild's
// raid timezone. Sign-ups live in raids_signups_store.go.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// Raid statuses.
const (
	RaidPlanned   = "planned"
	RaidLocked    = "locked"
	RaidDone      = "done"
	RaidCancelled = "cancelled"
)

// RaidCap is the most planned or locked raids one guild may hold at once.
const RaidCap = 12

// raidLingerHours is how long a planned or locked raid stays listed after it ends.
const raidLingerHours = 6

var (
	ErrRaidCap       = errors.New("guilds: raid cap reached")
	ErrRaidLocked    = errors.New("guilds: raid is locked")
	ErrRaidCancelled = errors.New("guilds: raid is cancelled")
)

// Raid is one scheduled raid night.
type Raid struct {
	ID              int64
	GuildID         int64
	Zone            string
	Title           string
	StartsAt        time.Time
	DurationMinutes int
	Notes           string
	Status          string
	ReportID        *string
	PostDefault     bool
	CreatedBy       int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// RaidInput carries the editable fields of a raid. GuildID and CreatedBy matter on create only.
type RaidInput struct {
	GuildID         int64
	Zone            string
	Title           string
	StartsAt        time.Time
	DurationMinutes int
	Notes           string
	PostDefault     bool
	CreatedBy       int64
}

const raidColumns = `id, guild_id, zone, title, starts_at, duration_minutes, notes, status,
	report_id, post_default, created_by, created_at, updated_at`

func scanRaid(row pgx.Row) (Raid, error) {
	var r Raid
	err := row.Scan(&r.ID, &r.GuildID, &r.Zone, &r.Title, &r.StartsAt, &r.DurationMinutes, &r.Notes,
		&r.Status, &r.ReportID, &r.PostDefault, &r.CreatedBy, &r.CreatedAt, &r.UpdatedAt)
	return r, err
}

func scanRaids(rows pgx.Rows) ([]Raid, error) {
	defer rows.Close()
	out := []Raid{}
	for rows.Next() {
		r, err := scanRaid(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	return out, rows.Err()
}

// CreateRaid inserts a planned raid unless the guild already holds RaidCap planned or locked ones.
func (s *Store) CreateRaid(ctx context.Context, in RaidInput) (Raid, error) {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return Raid{}, fmt.Errorf("guilds: create raid: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if _, err := tx.Exec(ctx, `select 1 from guilds where id = $1 for update`, in.GuildID); err != nil {
		return Raid{}, fmt.Errorf("guilds: create raid: lock guild: %w", err)
	}
	var open int
	if err := tx.QueryRow(ctx,
		`select count(*) from guild_raids where guild_id = $1 and status in ('planned','locked')`,
		in.GuildID).Scan(&open); err != nil {
		return Raid{}, fmt.Errorf("guilds: create raid: count: %w", err)
	}
	if open >= RaidCap {
		return Raid{}, ErrRaidCap
	}
	raid, err := scanRaid(tx.QueryRow(ctx, `
		insert into guild_raids (guild_id, zone, title, starts_at, duration_minutes, notes, post_default, created_by)
		values ($1, $2, $3, $4, $5, $6, $7, $8) returning `+raidColumns,
		in.GuildID, in.Zone, in.Title, in.StartsAt, in.DurationMinutes, in.Notes, in.PostDefault, in.CreatedBy))
	if err != nil {
		return Raid{}, fmt.Errorf("guilds: create raid: insert: %w", err)
	}
	if err := tx.Commit(ctx); err != nil {
		return Raid{}, fmt.Errorf("guilds: create raid: commit: %w", err)
	}
	return raid, nil
}

// UpdateRaid rewrites a raid's editable fields.
func (s *Store) UpdateRaid(ctx context.Context, id int64, in RaidInput) (Raid, error) {
	raid, err := scanRaid(s.Pool.QueryRow(ctx, `
		update guild_raids set zone = $2, title = $3, starts_at = $4, duration_minutes = $5,
			notes = $6, post_default = $7, updated_at = now()
		where id = $1 returning `+raidColumns,
		id, in.Zone, in.Title, in.StartsAt, in.DurationMinutes, in.Notes, in.PostDefault))
	return raid, raidWriteErr("update raid", err)
}

// SetRaidStatus moves a raid to a new status.
func (s *Store) SetRaidStatus(ctx context.Context, id int64, status string) (Raid, error) {
	raid, err := scanRaid(s.Pool.QueryRow(ctx,
		`update guild_raids set status = $2, updated_at = now() where id = $1 returning `+raidColumns,
		id, status))
	return raid, raidWriteErr("set raid status", err)
}

func raidWriteErr(op string, err error) error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, pgx.ErrNoRows):
		return ErrNotFound
	default:
		return fmt.Errorf("guilds: %s: %w", op, err)
	}
}

// GetRaid reads one raid of a guild.
func (s *Store) GetRaid(ctx context.Context, guildID, id int64) (Raid, bool, error) {
	raid, err := scanRaid(s.Pool.QueryRow(ctx,
		`select `+raidColumns+` from guild_raids where guild_id = $1 and id = $2`, guildID, id))
	if errors.Is(err, pgx.ErrNoRows) {
		return Raid{}, false, nil
	}
	if err != nil {
		return Raid{}, false, fmt.Errorf("guilds: get raid: %w", err)
	}
	return raid, true, nil
}

// PlannedRaids lists the raids a guild's calendar shows at now: planned or locked ones until
// raidLingerHours after they end, cancelled ones until they would have started.
func (s *Store) PlannedRaids(ctx context.Context, guildID int64, now time.Time) ([]Raid, error) {
	rows, err := s.Pool.Query(ctx, `
		select `+raidColumns+` from guild_raids
		where guild_id = $1 and (
			(status in ('planned','locked')
				and starts_at + make_interval(mins => duration_minutes) + make_interval(hours => $3) > $2)
			or (status = 'cancelled' and starts_at > $2))
		order by starts_at, id`, guildID, now, raidLingerHours)
	if err != nil {
		return nil, fmt.Errorf("guilds: planned raids: %w", err)
	}
	raids, err := scanRaids(rows)
	if err != nil {
		return nil, fmt.Errorf("guilds: planned raids: %w", err)
	}
	return raids, nil
}

// NextRaid is the first planned or locked raid that has not yet ended.
func (s *Store) NextRaid(ctx context.Context, guildID int64, now time.Time) (Raid, bool, error) {
	raid, err := scanRaid(s.Pool.QueryRow(ctx, `
		select `+raidColumns+` from guild_raids
		where guild_id = $1 and status in ('planned','locked')
			and starts_at + make_interval(mins => duration_minutes) > $2
		order by starts_at, id limit 1`, guildID, now))
	if errors.Is(err, pgx.ErrNoRows) {
		return Raid{}, false, nil
	}
	if err != nil {
		return Raid{}, false, fmt.Errorf("guilds: next raid: %w", err)
	}
	return raid, true, nil
}

// SetGuildRaidTimezone stores the guild's IANA raid timezone.
func (s *Store) SetGuildRaidTimezone(ctx context.Context, guildID int64, tz string) error {
	tag, err := s.Pool.Exec(ctx, `update guilds set raid_timezone = $2 where id = $1`, guildID, tz)
	if err != nil {
		return fmt.Errorf("guilds: set raid timezone: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

// GuildRaidTimezone reads the guild's raid timezone; empty when never set.
func (s *Store) GuildRaidTimezone(ctx context.Context, guildID int64) (string, error) {
	var tz *string
	err := s.Pool.QueryRow(ctx, `select raid_timezone from guilds where id = $1`, guildID).Scan(&tz)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", ErrNotFound
	}
	if err != nil {
		return "", fmt.Errorf("guilds: raid timezone: %w", err)
	}
	if tz == nil {
		return "", nil
	}
	return *tz, nil
}
