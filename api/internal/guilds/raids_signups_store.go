// api/internal/guilds/raids_signups_store.go
//
// Raid sign-ups and rule R3: one character per account is `in` or `late` for a raid.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// Sign-up roles and states.
const (
	RoleTank   = "tank"
	RoleHealer = "healer"
	RoleDPS    = "dps"

	SignupIn    = "in"
	SignupLate  = "late"
	SignupBench = "bench"
	SignupOut   = "out"
)

// Signup is one character's answer for a raid. Name, Class and Spec are a snapshot the
// caller supplies. AsOfficer is a write-time flag, never stored: it lets an officer write to
// a locked raid.
type Signup struct {
	RaidID       int64
	CharacterKey string
	Name         string
	Class        string
	Spec         string
	Role         string
	State        string
	Note         string
	SignedBy     int64
	UpdatedAt    time.Time
	AsOfficer    bool
}

func isAttending(state string) bool { return state == SignupIn || state == SignupLate }

// UpsertSignup writes a sign-up. When the new state is in or late, every other character the
// same account has in or late for this raid moves to bench; the first moved name is returned.
func (s *Store) UpsertSignup(ctx context.Context, sg Signup) (*string, error) {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("guilds: upsert signup: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if err := lockOwnSignups(ctx, tx, sg.RaidID, sg.SignedBy); err != nil {
		return nil, err
	}
	if err := checkRaidOpen(ctx, tx, sg); err != nil {
		return nil, err
	}
	var moved *string
	if isAttending(sg.State) {
		if moved, err = benchOtherOwn(ctx, tx, sg); err != nil {
			return nil, err
		}
	}
	if err := writeSignup(ctx, tx, sg); err != nil {
		return nil, err
	}
	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("guilds: upsert signup: commit: %w", err)
	}
	return moved, nil
}

// lockOwnSignups serialises one account's writes to one raid: the advisory lock covers the
// case where the account has no row yet, the row lock covers existing ones.
func lockOwnSignups(ctx context.Context, tx pgx.Tx, raidID, userID int64) error {
	if _, err := tx.Exec(ctx,
		`select pg_advisory_xact_lock(hashtextextended($1::bigint::text || ':' || $2::bigint::text, 0))`, raidID, userID); err != nil {
		return fmt.Errorf("guilds: upsert signup: advisory lock: %w", err)
	}
	if _, err := tx.Exec(ctx,
		`select 1 from guild_raid_signups where raid_id = $1 and signed_by = $2 for update`, raidID, userID); err != nil {
		return fmt.Errorf("guilds: upsert signup: lock own signups: %w", err)
	}
	return nil
}

func checkRaidOpen(ctx context.Context, tx pgx.Tx, sg Signup) error {
	var status string
	err := tx.QueryRow(ctx, `select status from guild_raids where id = $1 for share`, sg.RaidID).Scan(&status)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotFound
	}
	if err != nil {
		return fmt.Errorf("guilds: upsert signup: read raid: %w", err)
	}
	switch {
	case status == RaidCancelled:
		return ErrRaidCancelled
	case status == RaidLocked && !sg.AsOfficer:
		return ErrRaidLocked
	}
	return nil
}

func benchOtherOwn(ctx context.Context, tx pgx.Tx, sg Signup) (*string, error) {
	rows, err := tx.Query(ctx, `
		update guild_raid_signups set state = 'bench', updated_at = now()
		where raid_id = $1 and signed_by = $2 and character_key <> $3 and state in ('in','late')
		returning name`, sg.RaidID, sg.SignedBy, sg.CharacterKey)
	if err != nil {
		return nil, fmt.Errorf("guilds: upsert signup: bench others: %w", err)
	}
	defer rows.Close()
	var moved *string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("guilds: upsert signup: bench others: %w", err)
		}
		if moved == nil {
			moved = &name
		}
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: upsert signup: bench others: %w", err)
	}
	return moved, nil
}

func writeSignup(ctx context.Context, tx pgx.Tx, sg Signup) error {
	if _, err := tx.Exec(ctx, `
		insert into guild_raid_signups (raid_id, character_key, name, class, spec, role, state, note, signed_by)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		on conflict (raid_id, character_key) do update set
			name = excluded.name, class = excluded.class, spec = excluded.spec, role = excluded.role,
			state = excluded.state, note = excluded.note, signed_by = excluded.signed_by, updated_at = now()`,
		sg.RaidID, sg.CharacterKey, sg.Name, sg.Class, sg.Spec, sg.Role, sg.State, sg.Note, sg.SignedBy); err != nil {
		return fmt.Errorf("guilds: upsert signup: write: %w", err)
	}
	return nil
}

// RaidSignups lists a raid's sign-ups in a stable order.
func (s *Store) RaidSignups(ctx context.Context, raidID int64) ([]Signup, error) {
	rows, err := s.Pool.Query(ctx, `
		select raid_id, character_key, name, class, spec, role, state, note, signed_by, updated_at
		from guild_raid_signups where raid_id = $1 order by name, character_key`, raidID)
	if err != nil {
		return nil, fmt.Errorf("guilds: raid signups: %w", err)
	}
	defer rows.Close()
	out := []Signup{}
	for rows.Next() {
		var sg Signup
		if err := rows.Scan(&sg.RaidID, &sg.CharacterKey, &sg.Name, &sg.Class, &sg.Spec, &sg.Role,
			&sg.State, &sg.Note, &sg.SignedBy, &sg.UpdatedAt); err != nil {
			return nil, fmt.Errorf("guilds: raid signups: %w", err)
		}
		out = append(out, sg)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: raid signups: %w", err)
	}
	return out, nil
}
