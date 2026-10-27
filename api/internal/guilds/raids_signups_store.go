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
// a locked raid. SignedBy is the character's account (R3 groups by it); ActedBy is who made
// the write, the account itself or an officer answering for it.
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
	ActedBy      int64
	UpdatedAt    time.Time
	AsOfficer    bool
	// KeepNote is a write-time flag, never stored: an existing row keeps its note rather
	// than taking Note (a first answer still starts from Note).
	KeepNote bool
}

func isAttending(state string) bool { return state == SignupIn || state == SignupLate }

// UpsertSignup writes a sign-up. When the new state is in or late, every other character the
// same account has in or late for this raid moves to bench; the first moved name is returned.
//
// Ownership is the caller's contract: this trusts sg.SignedBy as the account the character
// belongs to (rule R3 groups by it) and checks neither that the account owns CharacterKey
// nor that an AsOfficer caller is an officer. The HTTP layer (raids_signup_handler.go) does
// both.
func (s *Store) UpsertSignup(ctx context.Context, sg Signup) (*string, error) {
	moved, _, err := s.upsertSignup(ctx, sg, false)
	return moved, err
}

// UpsertSignupIfChanged is UpsertSignup that writes nothing when the character's stored row
// already holds the asked state, role and (when a note is asked) note, judged under the same
// lock and after the same raid-open check as a write; wrote is false then.
func (s *Store) UpsertSignupIfChanged(ctx context.Context, sg Signup) (moved *string, wrote bool, err error) {
	return s.upsertSignup(ctx, sg, true)
}

func (s *Store) upsertSignup(ctx context.Context, sg Signup, skipSame bool) (*string, bool, error) {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return nil, false, fmt.Errorf("guilds: upsert signup: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if err := lockOwnSignups(ctx, tx, sg.RaidID, sg.SignedBy); err != nil {
		return nil, false, err
	}
	if err := checkRaidOpen(ctx, tx, sg); err != nil {
		return nil, false, err
	}
	if skipSame {
		same, err := signupHolds(ctx, tx, sg)
		if err != nil || same {
			return nil, false, err
		}
	}
	var moved *string
	if isAttending(sg.State) {
		if moved, err = benchOtherOwn(ctx, tx, sg); err != nil {
			return nil, false, err
		}
	}
	if err := writeSignup(ctx, tx, sg); err != nil {
		return nil, false, err
	}
	if err := tx.Commit(ctx); err != nil {
		return nil, false, fmt.Errorf("guilds: upsert signup: commit: %w", err)
	}
	return moved, true, nil
}

// signupHolds reports whether the character's stored row already has sg's state, role and
// (unless the note is kept) note.
func signupHolds(ctx context.Context, tx pgx.Tx, sg Signup) (bool, error) {
	var state, role, note string
	err := tx.QueryRow(ctx,
		`select state, role, note from guild_raid_signups where raid_id = $1 and character_key = $2`,
		sg.RaidID, sg.CharacterKey).Scan(&state, &role, &note)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("guilds: upsert signup: read row: %w", err)
	}
	return state == sg.State && role == sg.Role && (sg.KeepNote || note == sg.Note), nil
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
	case (status == RaidLocked || status == RaidDone) && !sg.AsOfficer:
		// A done raid is logged; like a locked one, only an officer corrects it.
		return ErrRaidLocked
	}
	return nil
}

func benchOtherOwn(ctx context.Context, tx pgx.Tx, sg Signup) (*string, error) {
	rows, err := tx.Query(ctx, `
		update guild_raid_signups set state = 'bench', acted_by = $4, updated_at = now()
		where raid_id = $1 and signed_by = $2 and character_key <> $3 and state in ('in','late')
		returning name`, sg.RaidID, sg.SignedBy, sg.CharacterKey, sg.ActedBy)
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
		insert into guild_raid_signups (raid_id, character_key, name, class, spec, role, state, note, signed_by, acted_by)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		on conflict (raid_id, character_key) do update set
			name = excluded.name, class = excluded.class, spec = excluded.spec, role = excluded.role,
			state = excluded.state, signed_by = excluded.signed_by,
			note = case when $11 then guild_raid_signups.note else excluded.note end,
			acted_by = excluded.acted_by, updated_at = now()`,
		sg.RaidID, sg.CharacterKey, sg.Name, sg.Class, sg.Spec, sg.Role, sg.State, sg.Note, sg.SignedBy, sg.ActedBy,
		sg.KeepNote); err != nil {
		return fmt.Errorf("guilds: upsert signup: write: %w", err)
	}
	return nil
}

// RaidSignups lists a raid's sign-ups in a stable order.
func (s *Store) RaidSignups(ctx context.Context, raidID int64) ([]Signup, error) {
	byRaid, err := s.SignupsByRaid(ctx, []int64{raidID})
	if err != nil {
		return nil, err
	}
	if out := byRaid[raidID]; out != nil {
		return out, nil
	}
	return []Signup{}, nil
}

// SignupsByRaid reads the sign-ups of every raid in raidIDs in one query, keyed by raid id,
// each list in a stable order (name, then key).
func (s *Store) SignupsByRaid(ctx context.Context, raidIDs []int64) (map[int64][]Signup, error) {
	return signupsByRaid(ctx, s.Pool, raidIDs)
}

// signupsByRaid is SignupsByRaid on q (the pool, or a match's transaction).
func signupsByRaid(ctx context.Context, q dbQuerier, raidIDs []int64) (map[int64][]Signup, error) {
	out := make(map[int64][]Signup, len(raidIDs))
	if len(raidIDs) == 0 {
		return out, nil
	}
	rows, err := q.Query(ctx, `
		select raid_id, character_key, name, class, spec, role, state, note, signed_by, acted_by, updated_at
		from guild_raid_signups where raid_id = any($1) order by raid_id, name, character_key`, raidIDs)
	if err != nil {
		return nil, fmt.Errorf("guilds: raid signups: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var sg Signup
		if err := rows.Scan(&sg.RaidID, &sg.CharacterKey, &sg.Name, &sg.Class, &sg.Spec, &sg.Role,
			&sg.State, &sg.Note, &sg.SignedBy, &sg.ActedBy, &sg.UpdatedAt); err != nil {
			return nil, fmt.Errorf("guilds: raid signups: %w", err)
		}
		out[sg.RaidID] = append(out[sg.RaidID], sg)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: raid signups: %w", err)
	}
	return out, nil
}
