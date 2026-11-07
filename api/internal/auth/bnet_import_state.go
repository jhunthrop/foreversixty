package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// The Battle.net import states GET /v1/me reports as bnet_import.state
// (design/specs/2026-11-07-setup.md section 9, G1). They live on the users row
// (migration 0078), so a poll that lands on another Cloud Run instance reads the same state.
const (
	BnetImportRunning = "running"
	BnetImportDone    = "done"
	BnetImportFailed  = "failed"
)

// bnetImportStaleMargin is the slack past the longest an import can run (the sign-in
// budget plus the background continuation) before a running state is presumed lost.
const bnetImportStaleMargin = time.Minute

// bnetImportStaleAfter is how old a running state may be and still read as running. Past
// it, the instance that ran the import was stopped before it could record the end (a
// deploy, a scale-in), so the read reports failed: the import did not finish, and the
// player is offered Refresh rather than a spinner that never ends. The row is not rewritten.
const bnetImportStaleAfter = importBudget + importContinueBudget + bnetImportStaleMargin

// bnetImportStateWriteBudget bounds one state write. It runs on a fresh context, since the
// request's (or the continuation's) may already be spent by the time the import ends.
const bnetImportStateWriteBudget = 5 * time.Second

// BnetImportView is /v1/me's bnet_import object: the last tracked import's state and when
// that state was recorded (RFC 3339).
type BnetImportView struct {
	State string `json:"state"`
	At    string `json:"at"`
}

// bnetImportViewOf is the bnet_import /v1/me answers for an account, nil when no import
// has ever been tracked. A running state older than bnetImportStaleAfter reads as failed,
// at its recorded time.
func bnetImportViewOf(u User, now time.Time) *BnetImportView {
	if u.BnetImportState == nil || u.BnetImportStateAt == nil {
		return nil
	}
	state := *u.BnetImportState
	if state == BnetImportRunning && now.Sub(*u.BnetImportStateAt) > bnetImportStaleAfter {
		state = BnetImportFailed
	}
	return &BnetImportView{State: state, At: u.BnetImportStateAt.UTC().Format(time.RFC3339)}
}

// SetBnetImportState records the account's Battle.net import state, stamped now.
func (s *Store) SetBnetImportState(ctx context.Context, userID int64, state string) error {
	if _, err := s.Pool.Exec(ctx,
		`update users set bnet_import_state = $2, bnet_import_state_at = now() where id = $1`,
		userID, state); err != nil {
		return fmt.Errorf("auth: set bnet import state: %w", err)
	}
	return nil
}

// recordImportState writes the state and logs a failure to: it never fails the sign-in.
func (s *Service) recordImportState(id string, userID int64, state string) {
	ctx, cancel := context.WithTimeout(context.Background(), bnetImportStateWriteBudget)
	defer cancel()
	if err := s.Store.SetBnetImportState(ctx, userID, state); err != nil {
		s.logger().Warn("auth", "id", id, "op", "battlenet_import_state", "user_id", userID,
			"state", state, "err", err)
	}
}

// importEndState is the state an import run that ended this way records.
func importEndState(err error) string {
	if err != nil {
		return BnetImportFailed
	}
	return BnetImportDone
}

// runImport calls the Battle.net importer, bounded by importBudget, and
// logs the outcome. It never fails the login (spec §4.1): an error is
// logged at WARN and the caller redirects exactly as it would have with
// no importer at all. The account reads running from the start until the
// run ends, or until the background continuation it hands off to ends.
func (s *Service) runImport(r *http.Request, userID int64, accessToken string, expiresAt time.Time) {
	if s.Importer == nil || accessToken == "" {
		return
	}
	id := httpx.RequestIDFrom(r.Context())
	if keeper, ok := s.Importer.(TokenKeeper); ok {
		keeper.KeepToken(r.Context(), userID, accessToken, expiresAt)
	}
	s.recordImportState(id, userID, BnetImportRunning)
	ctx, cancel := context.WithTimeout(r.Context(), importBudget)
	defer cancel()
	summary, err := s.Importer.ImportAccount(ctx, userID, accessToken)
	s.logImport(id, "battlenet_import", userID, summary, err)
	if summary.Pending > 0 && !summary.RateLimited {
		s.continueImport(id, userID, accessToken)
		return
	}
	s.recordImportState(id, userID, importEndState(err))
}

// continueImport finishes, in the background, an import the sign-in budget
// cut short (importContinueBudget). The access token is never logged; the only copy
// persisted is the TokenKeeper's sealed one.
func (s *Service) continueImport(id string, userID int64, accessToken string) {
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), importContinueBudget)
		defer cancel()
		summary, err := s.Importer.ImportAccount(ctx, userID, accessToken)
		s.logImport(id, "battlenet_import_continued", userID, summary, err)
		s.recordImportState(id, userID, importEndState(err))
	}()
}

func (s *Service) logImport(id, op string, userID int64, summary ImportSummary, err error) {
	if err != nil {
		s.logger().Warn("auth", "id", id, "op", op, "user_id", userID, "pending", summary.Pending,
			"rate_limited", summary.RateLimited, "err", err)
		return
	}
	s.logger().Info("auth", "id", id, "op", op, "user_id", userID,
		"regions", summary.Regions, "characters", summary.Characters,
		"guilds", summary.Guilds, "skipped", summary.Skipped, "unavailable", summary.Unavailable)
}
