package subscribe

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"
	"net/mail"
	"strings"
	"sync"
	"time"

	"github.com/PLACEHOLDER/forever/api/internal/httpx"
	mailer "github.com/PLACEHOLDER/forever/api/internal/mail"
)

var ErrInvalidEmail = errors.New("invalid email")

// mailSendTimeout bounds how long a single confirmation send (which now
// happens off the request path, see Service.Subscribe) may run.
const mailSendTimeout = 15 * time.Second

type Storer interface {
	// Upsert inserts a new subscriber row when normalized is not already
	// present. When a row already exists, no row is inserted and the
	// returned UpsertResult describes that row so the caller can decide
	// whether, and how, to resend the confirmation instead of leaving the
	// address stuck unconfirmed after a prior mail-send failure, or
	// re-mailing an address that is already confirmed and still subscribed.
	Upsert(ctx context.Context, email, normalized, token, unsubscribeToken string) (UpsertResult, error)
	// MarkConfirmationSent records that a confirmation email was just sent
	// for normalized, so future sends can be rate-limited against it.
	MarkConfirmationSent(ctx context.Context, normalized string) error
	Confirm(ctx context.Context, token string) (found bool, err error)
	Unsubscribe(ctx context.Context, unsubscribeToken string) (found bool, err error)
}

// Service implements the double opt-in subscribe flow. Mail is sent off the
// request path (see Subscribe): callers that need to observe a send (tests,
// and graceful shutdown) must call Wait first.
type Service struct {
	Store         Storer
	Mailer        mailer.Mailer
	PublicBaseURL string
	APIBaseURL    string
	Logger        *slog.Logger

	// ResendCooldown, when positive, blocks resending a confirmation email
	// to the same address more often than this. Zero disables the cooldown.
	ResendCooldown time.Duration
	// MaxSendsPerHour, when positive, caps how many confirmation emails this
	// Service instance will send in any trailing 60-minute window,
	// regardless of address. Zero disables the cap.
	MaxSendsPerHour int

	wg sync.WaitGroup

	sendMu    sync.Mutex
	sendTimes []time.Time // sliding window of send timestamps, for MaxSendsPerHour
}

func normalize(email string) (display, normalized string, err error) {
	display = strings.TrimSpace(email)
	addr, err := mail.ParseAddress(display)
	if err != nil || addr.Name != "" || !strings.Contains(addr.Address, "@") {
		return "", "", ErrInvalidEmail
	}
	return addr.Address, strings.ToLower(addr.Address), nil
}

func newToken() string {
	b := make([]byte, 24)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func (s *Service) log() *slog.Logger {
	if s.Logger != nil {
		return s.Logger
	}
	return slog.Default()
}

// withinCooldown reports whether a resend must be blocked because the last
// confirmation for this address went out too recently. sentAt is nil for an
// address that has never had a confirmation sent, which is never blocked.
func (s *Service) withinCooldown(sentAt *time.Time) bool {
	if s.ResendCooldown <= 0 || sentAt == nil {
		return false
	}
	return time.Since(*sentAt) < s.ResendCooldown
}

// reserveSendSlot enforces MaxSendsPerHour as a sliding window over the
// trailing hour, shared across all addresses on this Service instance. It
// returns false (reserving nothing) when the cap has already been reached.
func (s *Service) reserveSendSlot(now time.Time) bool {
	if s.MaxSendsPerHour <= 0 {
		return true
	}
	s.sendMu.Lock()
	defer s.sendMu.Unlock()
	cutoff := now.Add(-time.Hour)
	kept := s.sendTimes[:0]
	for _, t := range s.sendTimes {
		if t.After(cutoff) {
			kept = append(kept, t)
		}
	}
	s.sendTimes = kept
	if len(s.sendTimes) >= s.MaxSendsPerHour {
		return false
	}
	s.sendTimes = append(s.sendTimes, now)
	return true
}

func (s *Service) logSkippedSend(ctx context.Context, reason string) {
	s.log().Info("subscribe", "id", httpx.RequestIDFrom(ctx), "op", "subscribe", "reason", reason)
}

// Subscribe records email and, depending on the existing row (if any),
// sends a confirmation email:
//
//   - new address: send a confirmation for the freshly issued token.
//   - existing, not yet confirmed: resend the original confirmation token.
//   - existing, confirmed, and unsubscribed: resend the original
//     confirmation token; visiting its link (Confirm) clears unsubscribed_at
//     and resubscribes the address.
//   - existing, confirmed, and not unsubscribed: send nothing.
//
// Any send is subject to ResendCooldown and MaxSendsPerHour: when either
// blocks it, Subscribe sends nothing, logs why, and still returns nil so the
// handler returns 202 regardless.
//
// The send itself happens after Subscribe returns (see Wait) so a slow or
// failing mail provider cannot turn a subscribe request into a slow or
// failing response.
func (s *Service) Subscribe(ctx context.Context, email string) error {
	display, normalized, err := normalize(email)
	if err != nil {
		return err
	}
	token := newToken()
	unsubscribeToken := newToken()
	result, err := s.Store.Upsert(ctx, display, normalized, token, unsubscribeToken)
	if err != nil {
		return fmt.Errorf("subscribe: store: %w", err)
	}

	if result.Existing {
		if result.Confirmed && !result.Unsubscribed {
			return nil // already confirmed and subscribed: nothing to do
		}
		token = result.Token // never invent a new token for an existing row
	}

	if s.withinCooldown(result.ConfirmationSentAt) {
		s.logSkippedSend(ctx, "cooldown")
		return nil
	}
	if !s.reserveSendSlot(time.Now()) {
		s.logSkippedSend(ctx, "cap")
		return nil
	}

	msg := mailer.Message{
		To:      display,
		Subject: "Confirm your Forever Sixty subscription",
		Text: "Confirm to get Forever Sixty updates (a few emails before launch, then only when something changes):\n\n" +
			s.APIBaseURL + "/v1/subscribe/confirm?token=" + token + "\n\n" +
			"If you did not ask for this, ignore it and nothing happens.\n",
	}
	requestID := httpx.RequestIDFrom(ctx)

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		sendCtx, cancel := context.WithTimeout(context.Background(), mailSendTimeout)
		defer cancel()
		if err := s.Mailer.Send(sendCtx, msg); err != nil {
			s.log().Error("subscribe", "id", requestID, "op", "subscribe", "stage", "mail", "err", err)
			return
		}
		if err := s.Store.MarkConfirmationSent(sendCtx, normalized); err != nil {
			s.log().Error("subscribe", "id", requestID, "op", "subscribe", "stage", "mail", "err", err)
		}
	}()
	return nil
}

// Wait blocks until every confirmation send started by Subscribe has
// finished. Tests must call it before asserting on a fake mailer's sent
// messages; shutdown calls it before closing the database pool so an
// in-flight send's MarkConfirmationSent has somewhere to write.
func (s *Service) Wait() {
	s.wg.Wait()
}

func (s *Service) Confirm(ctx context.Context, token string) (bool, error) {
	return s.Store.Confirm(ctx, token)
}

func (s *Service) Unsubscribe(ctx context.Context, unsubscribeToken string) (bool, error) {
	return s.Store.Unsubscribe(ctx, unsubscribeToken)
}
