package push

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jhunthrop/foreversixty/api/internal/events"
)

// Sink is the push channel: one message per registered token of the account. A token the
// service calls invalid is deleted.
type Sink struct {
	Store   TokenStore
	Senders map[Platform]Sender
	Log     *slog.Logger
}

var _ events.Sink = (*Sink)(nil)

func (s *Sink) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

// MessageOf is e as a message to t.
func MessageOf(e events.Event, t Token) Message {
	m := Message{
		Token: t.Token, Platform: t.Platform, Kind: e.Kind, Title: e.Title, Body: e.Body, Link: e.Link,
		Thread: e.Thread, Collapse: e.Collapse, Actions: e.Actions, Data: e.Data,
	}
	if len(e.Actions) > 0 {
		m.Category = RaidSignupCategory
	}
	return m
}

// Deliver sends e to every token of userID. It fails only when no token took it and at
// least one send failed, so the router retries; a token that took it is never sent twice by
// a partial failure.
func (s *Sink) Deliver(ctx context.Context, userID int64, e events.Event) error {
	tokens, err := s.Store.TokensOf(ctx, userID)
	if err != nil {
		return err
	}
	var errs []error
	delivered := false
	for _, t := range tokens {
		sender, ok := s.Senders[t.Platform]
		if !ok {
			continue
		}
		err := sender.Send(ctx, MessageOf(e, t))
		switch {
		case err == nil:
			delivered = true
		case errors.Is(err, ErrInvalidToken):
			s.dropInvalid(ctx, t)
		default:
			errs = append(errs, fmt.Errorf("push: %s: %w", t.Platform, err))
		}
	}
	if delivered {
		return nil
	}
	return errors.Join(errs...)
}

func (s *Sink) dropInvalid(ctx context.Context, t Token) {
	s.logger().Info("push", "op", "drop_invalid", "user", t.UserID, "device", t.DeviceID, "platform", string(t.Platform))
	if err := s.Store.Delete(context.WithoutCancel(ctx), t.Token); err != nil {
		s.logger().Error("push", "op", "drop_invalid", "device", t.DeviceID, "err", err)
	}
}
