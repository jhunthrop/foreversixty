package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// ConnectTTL is how long a companion connect can be approved and
// collected, from its creation. A decided row expires on the same clock.
const ConnectTTL = 10 * time.Minute

// ConnectPollInterval is the seconds the companion is told to wait
// between polls.
const ConnectPollInterval = 2

// connectKeep is how long a finished row stays after it expires before the
// sweep takes it, so a late poll still reads "used" or "expired".
const connectKeep = time.Hour

// maxCodeDraws bounds the retries when a drawn user code is already live.
const maxCodeDraws = 5

// ErrConnectDecided is returned when a connect is approved or denied after
// it was already decided.
var ErrConnectDecided = errors.New("auth: connect already decided")

// newUserCode is a variable so a test can force a collision.
var newUserCode = NewUserCode

// Connect is a freshly minted connect: the id the companion holds and the
// code the person sees. Neither is stored in the clear (the id as a hash,
// the code in its hyphen-less form).
type Connect struct {
	ConnectID string
	UserCode  string // formatted, KQ4-7TM
}

// ConnectInfo is what the site shows before the person decides.
type ConnectInfo struct {
	Name, Platform string
	ExpiresIn      int
}

// ConnectState is what a poll reads.
type ConnectState struct {
	Status string // pending, approved, denied, expired, used
	// Approved only: the device made by this poll, its owner's label.
	Device       Device
	AccountLabel string
}

// CreateConnect mints a connect row. A drawn code that collides with a live
// one is refused by the partial unique index and drawn again.
func (s *Store) CreateConnect(ctx context.Context, name, platform string) (Connect, error) {
	// Lazy sweep: rows long past their expiry are of no use to anyone.
	if _, err := s.Pool.Exec(ctx,
		`delete from device_connects where expires_at < now() - $1::interval`,
		fmt.Sprintf("%d seconds", int(connectKeep.Seconds()))); err != nil {
		return Connect{}, fmt.Errorf("auth: sweep connects: %w", err)
	}
	id := NewConnectID()
	for range maxCodeDraws {
		code := newUserCode()
		_, err := s.Pool.Exec(ctx,
			`insert into device_connects (id_hash, user_code, name, platform, expires_at)
			 values ($1, $2, $3, $4, now() + $5::interval)`,
			TokenHash(id), code, name, platform, fmt.Sprintf("%d seconds", int(ConnectTTL.Seconds())))
		var pg *pgconn.PgError
		if errors.As(err, &pg) && pg.Code == "23505" {
			continue
		}
		if err != nil {
			return Connect{}, fmt.Errorf("auth: create connect: %w", err)
		}
		return Connect{ConnectID: id, UserCode: FormatUserCode(code)}, nil
	}
	return Connect{}, errors.New("auth: create connect: no free user code")
}

// LookupConnect reads a live, undecided connect by its code.
func (s *Store) LookupConnect(ctx context.Context, code string) (ConnectInfo, error) {
	var in ConnectInfo
	err := s.Pool.QueryRow(ctx,
		`select name, platform, greatest(0, ceil(extract(epoch from expires_at - now())))::int
		 from device_connects
		 where user_code = $1 and collected_at is null and status = 'pending' and expires_at > now()`,
		code).Scan(&in.Name, &in.Platform, &in.ExpiresIn)
	if errors.Is(err, pgx.ErrNoRows) {
		return ConnectInfo{}, ErrNotFound
	}
	if err != nil {
		return ConnectInfo{}, fmt.Errorf("auth: look up connect: %w", err)
	}
	return in, nil
}

// DecideConnect approves (binding the row to userID) or denies a live
// connect. ErrNotFound is an unknown or expired code; ErrConnectDecided is a
// code that was already approved or denied.
func (s *Store) DecideConnect(ctx context.Context, code string, userID int64, approve bool) (ConnectInfo, error) {
	status := "denied"
	if approve {
		status = "approved"
	}
	var in ConnectInfo
	err := s.Pool.QueryRow(ctx,
		`update device_connects set status = $2, user_id = $3
		 where user_code = $1 and collected_at is null and status = 'pending' and expires_at > now()
		 returning name, platform`, code, status, userID).Scan(&in.Name, &in.Platform)
	if err == nil {
		return in, nil
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return ConnectInfo{}, fmt.Errorf("auth: decide connect: %w", err)
	}
	var decided bool
	err = s.Pool.QueryRow(ctx,
		`select true from device_connects where user_code = $1 and collected_at is null and expires_at > now()`,
		code).Scan(&decided)
	if errors.Is(err, pgx.ErrNoRows) {
		return ConnectInfo{}, ErrNotFound
	}
	if err != nil {
		return ConnectInfo{}, fmt.Errorf("auth: decide connect: %w", err)
	}
	return ConnectInfo{}, ErrConnectDecided
}

// PollConnect answers a poll. An approved row is collected here: the device
// is created, with the caller's token hash, in the same transaction that
// stamps the row collected, so a token is handed out at most once and never
// without a device behind it. ErrNotFound is an unknown id.
func (s *Store) PollConnect(ctx context.Context, id string, tokenHash []byte) (ConnectState, error) {
	idHash := TokenHash(id)
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return ConnectState{}, fmt.Errorf("auth: poll connect: %w", err)
	}
	defer tx.Rollback(ctx) //nolint:errcheck // a no-op after commit

	var (
		status, name, platform string
		userID                 *int64
		collected, expired     bool
	)
	err = tx.QueryRow(ctx,
		`select status, name, platform, user_id, collected_at is not null, expires_at <= now()
		 from device_connects where id_hash = $1 for update`, idHash).
		Scan(&status, &name, &platform, &userID, &collected, &expired)
	if errors.Is(err, pgx.ErrNoRows) {
		return ConnectState{}, ErrNotFound
	}
	if err != nil {
		return ConnectState{}, fmt.Errorf("auth: poll connect: %w", err)
	}
	switch {
	case collected:
		return ConnectState{Status: "used"}, nil
	case expired:
		return ConnectState{Status: "expired"}, nil
	case status != "approved" || userID == nil:
		return ConnectState{Status: status}, nil
	}

	d := Device{ID: NewDeviceID(), UserID: *userID, Name: name, Platform: platform, Kind: DeviceKindCompanion}
	if err := tx.QueryRow(ctx,
		`insert into devices (id, user_id, name, platform, token_hash) values ($1, $2, $3, $4, $5)
		 returning created_at`, d.ID, d.UserID, name, platform, tokenHash).Scan(&d.CreatedAt); err != nil {
		return ConnectState{}, fmt.Errorf("auth: create device: %w", err)
	}
	if _, err := tx.Exec(ctx, `update device_connects set collected_at = now() where id_hash = $1`, idHash); err != nil {
		return ConnectState{}, fmt.Errorf("auth: collect connect: %w", err)
	}
	var label string
	if err := tx.QueryRow(ctx,
		`select coalesce(nullif(battletag, ''), email, '') from users where id = $1`, d.UserID).Scan(&label); err != nil {
		return ConnectState{}, fmt.Errorf("auth: connect account label: %w", err)
	}
	if err := tx.Commit(ctx); err != nil {
		return ConnectState{}, fmt.Errorf("auth: poll connect: %w", err)
	}
	return ConnectState{Status: "approved", Device: d, AccountLabel: label}, nil
}

// AccountLabel is who a device connected as: the battletag, else the email.
func (s *Store) AccountLabel(ctx context.Context, userID int64) (string, error) {
	var label string
	err := s.Pool.QueryRow(ctx,
		`select coalesce(nullif(battletag, ''), email, '') from users where id = $1`, userID).Scan(&label)
	if err != nil {
		return "", fmt.Errorf("auth: account label: %w", err)
	}
	return label, nil
}
