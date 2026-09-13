package subscribe

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"net/mail"
	"strings"

	mailer "github.com/PLACEHOLDER/forever/api/internal/mail"
)

var ErrInvalidEmail = errors.New("invalid email")

type Storer interface {
	Upsert(ctx context.Context, email, normalized, token string) (existing bool, err error)
	Confirm(ctx context.Context, token string) (found bool, err error)
	Unsubscribe(ctx context.Context, token string) (found bool, err error)
}

type Service struct {
	Store         Storer
	Mailer        mailer.Mailer
	PublicBaseURL string
	APIBaseURL    string
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

func (s *Service) Subscribe(ctx context.Context, email string) error {
	display, normalized, err := normalize(email)
	if err != nil {
		return err
	}
	token := newToken()
	existing, err := s.Store.Upsert(ctx, display, normalized, token)
	if err != nil {
		return fmt.Errorf("subscribe: store: %w", err)
	}
	if existing {
		return nil
	}
	msg := mailer.Message{
		To:      display,
		Subject: "Confirm your Forever Sixty subscription",
		Text: "Confirm to get Forever Sixty updates (a few emails before launch, then only when something changes):\n\n" +
			s.APIBaseURL + "/v1/subscribe/confirm?token=" + token + "\n\n" +
			"If you did not ask for this, ignore it and nothing happens.\n" +
			"Unsubscribe any time: " + s.APIBaseURL + "/v1/subscribe/unsubscribe?token=" + token + "\n",
	}
	if err := s.Mailer.Send(ctx, msg); err != nil {
		return fmt.Errorf("subscribe: mail: %w", err)
	}
	return nil
}

func (s *Service) Confirm(ctx context.Context, token string) (bool, error) {
	return s.Store.Confirm(ctx, token)
}

func (s *Service) Unsubscribe(ctx context.Context, token string) (bool, error) {
	return s.Store.Unsubscribe(ctx, token)
}
