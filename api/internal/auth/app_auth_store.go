package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// The device kinds the devices table holds (migration 0059).
const (
	DeviceKindCompanion = "companion"
	DeviceKindApp       = "app"
)

const (
	// AppFlowTTL is how long an app sign-in may sit at Battle.net: the web flow's own limit.
	AppFlowTTL = bnetStateTTL
	// AppCodeTTL is how long the one-time code on the app link can be exchanged.
	AppCodeTTL = 2 * time.Minute
	// appFlowKeep is how long a finished or abandoned flow stays before the sweep takes it.
	appFlowKeep = time.Hour
	// DefaultAppTokenIdle and DefaultAppTokenMaxAge are config's defaults, used when a Store
	// is built without them (tests).
	DefaultAppTokenIdle   = 30 * 24 * time.Hour
	DefaultAppTokenMaxAge = 180 * 24 * time.Hour
)

// ErrDeviceExpired is an app device token past its idle limit or maximum age.
var ErrDeviceExpired = errors.New("auth: app sign-in expired")

// liveAppClause keeps a device row only while it is a companion or an app inside both of
// its limits; $2 is the idle interval and $3 the maximum age.
const liveAppClause = `(kind <> 'app' or (coalesce(last_seen_at, created_at) > now() - $2::interval
	   and created_at > now() - $3::interval))`

func (s *Store) appTokenLimits() (idle, maxAge time.Duration) {
	idle, maxAge = s.AppTokenIdle, s.AppTokenMaxAge
	if idle <= 0 {
		idle = DefaultAppTokenIdle
	}
	if maxAge <= 0 {
		maxAge = DefaultAppTokenMaxAge
	}
	return idle, maxAge
}

// lapseAppDevice runs when a token matched no live device. A live row that is left can only
// be an app past its limits: it is revoked, and the caller hears ErrDeviceExpired. Anything
// else is ErrNotFound.
func (s *Store) lapseAppDevice(ctx context.Context, hash []byte) error {
	tag, err := s.Pool.Exec(ctx,
		`update devices set revoked_at = now() where token_hash = $1 and revoked_at is null and kind = 'app'`, hash)
	if err != nil {
		return fmt.Errorf("auth: lapse app device: %w", err)
	}
	if tag.RowsAffected() == 1 {
		return ErrDeviceExpired
	}
	return ErrNotFound
}

// AppFlow is what an app sign-in carries from start to exchange.
type AppFlow struct {
	Platform      string
	CodeChallenge string
	AppState      string
}

// interval renders a duration as a Postgres interval literal, the shape every TTL here uses.
func interval(d time.Duration) string { return fmt.Sprintf("%d seconds", int(d.Seconds())) }

// CreateAppFlow records an app sign-in under the hash of the state sent to Battle.net, and
// sweeps flows long past their expiry.
func (s *Store) CreateAppFlow(ctx context.Context, stateHash []byte, f AppFlow) error {
	if _, err := s.Pool.Exec(ctx,
		`delete from app_auth_flows where expires_at < now() - $1::interval`, interval(appFlowKeep)); err != nil {
		return fmt.Errorf("auth: sweep app flows: %w", err)
	}
	if _, err := s.Pool.Exec(ctx,
		`insert into app_auth_flows (state_hash, platform, code_challenge, app_state, expires_at)
		 values ($1, $2, $3, $4, now() + $5::interval)`,
		stateHash, f.Platform, f.CodeChallenge, f.AppState, interval(AppFlowTTL)); err != nil {
		return fmt.Errorf("auth: create app flow: %w", err)
	}
	return nil
}

// IssueAppCode binds a live flow to the account Battle.net confirmed and stores the hash of
// its one-time code. A flow that is unknown, expired or already issued is ErrNotFound, so a
// Battle.net callback replayed against the same state mints nothing.
func (s *Store) IssueAppCode(ctx context.Context, stateHash []byte, userID int64, codeHash []byte) (AppFlow, error) {
	var f AppFlow
	err := s.Pool.QueryRow(ctx,
		`update app_auth_flows set code_hash = $3, user_id = $2, code_expires_at = now() + $4::interval
		 where state_hash = $1 and code_hash is null and expires_at > now()
		 returning platform, code_challenge, app_state`,
		stateHash, userID, codeHash, interval(AppCodeTTL)).Scan(&f.Platform, &f.CodeChallenge, &f.AppState)
	if errors.Is(err, pgx.ErrNoRows) {
		return AppFlow{}, ErrNotFound
	}
	if err != nil {
		return AppFlow{}, fmt.Errorf("auth: issue app code: %w", err)
	}
	return f, nil
}

// SpendAppCode spends a one-time code and returns whose it was. The update is the check: a
// code already spent, or expired, matches no row and is ErrNotFound. It is spent before the
// caller checks the PKCE verifier, so a stolen code tried without the verifier is gone.
func (s *Store) SpendAppCode(ctx context.Context, codeHash []byte) (int64, AppFlow, error) {
	var userID int64
	var f AppFlow
	err := s.Pool.QueryRow(ctx,
		`update app_auth_flows set used_at = now()
		 where code_hash = $1 and used_at is null and code_expires_at > now()
		 returning user_id, platform, code_challenge, app_state`, codeHash).
		Scan(&userID, &f.Platform, &f.CodeChallenge, &f.AppState)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, AppFlow{}, ErrNotFound
	}
	if err != nil {
		return 0, AppFlow{}, fmt.Errorf("auth: spend app code: %w", err)
	}
	return userID, f, nil
}

// CreateAppDevice stores a signed-in native app as a device of kind app.
func (s *Store) CreateAppDevice(ctx context.Context, userID int64, name, platform string, hash []byte) (Device, error) {
	return s.createDevice(ctx, DeviceKindApp, userID, name, platform, hash)
}
