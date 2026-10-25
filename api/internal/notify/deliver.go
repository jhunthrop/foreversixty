package notify

import (
	"context"
	"errors"
	"log/slog"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

// SimEvents is what the sims job tells; nil-safe at the call site.
type SimEvents interface {
	SimFinished(ctx context.Context, e SimFinished)
}

// UploadEvents is what the parse job tells; nil-safe at the call site.
type UploadEvents interface {
	UploadParsed(ctx context.Context, e UploadParsed)
}

// Deliverer posts one embed per event to a member's saved webhook.
type Deliverer struct {
	Store         *Store
	Sealer        *Sealer
	Entitlements  Entitlementer
	Poster        Poster
	PublicBaseURL string
	Log           *slog.Logger
}

// Deliverer returns a Deliverer sharing the service's wiring.
func (s *Service) Deliverer() *Deliverer {
	return &Deliverer{
		Store: s.Store, Sealer: s.Sealer, Entitlements: s.Entitlements,
		Poster: s.Poster, PublicBaseURL: s.PublicBaseURL, Log: s.Log,
	}
}

func (d *Deliverer) logger() *slog.Logger {
	if d.Log != nil {
		return d.Log
	}
	return slog.Default()
}

// SimFinished posts the sim message; a failure is logged, never returned,
// because the sim itself succeeded.
func (d *Deliverer) SimFinished(ctx context.Context, e SimFinished) {
	d.deliverQuietly(ctx, e.UserID, EventSimFinished, SimFinishedEmbed(d.PublicBaseURL, e))
}

// UploadParsed posts the log message; see SimFinished.
func (d *Deliverer) UploadParsed(ctx context.Context, e UploadParsed) {
	d.deliverQuietly(ctx, e.UserID, EventUploadParsed, UploadParsedEmbed(d.PublicBaseURL, e))
}

func (d *Deliverer) deliverQuietly(ctx context.Context, userID int64, event EventName, embed Embed) {
	if err := d.Deliver(ctx, userID, event, embed); err != nil {
		d.logger().Error("notify", "user", userID, "event", string(event), "outcome", "error")
	}
}

// Deliver posts embed when the member has a live webhook with the event on
// and still holds the entitlement. Skips are silent successes. The error it
// returns never holds the URL, token or message text.
func (d *Deliverer) Deliver(ctx context.Context, userID int64, event EventName, embed Embed) error {
	hook, found, err := d.Store.Get(ctx, userID, KindDiscord)
	if err != nil {
		return err
	}
	if !found || hook.DisabledAt != nil || !hook.Events[event] {
		return nil
	}
	allowed, _, err := d.Entitlements.Can(ctx, userID, entitlements.FeatureNotifications)
	if err != nil {
		return err
	}
	if !allowed {
		return nil
	}
	box, err := d.Store.TokenBox(ctx, userID, KindDiscord)
	if err != nil {
		return err
	}
	url, err := d.Sealer.Open(box)
	if err != nil {
		return errors.New("notify: the saved webhook could not be opened")
	}
	postCtx, cancel := context.WithTimeout(ctx, posterTimeout)
	defer cancel()
	return d.record(ctx, userID, event, d.Poster.Post(postCtx, url, embed))
}

func (d *Deliverer) record(ctx context.Context, userID int64, event EventName, postErr error) error {
	if postErr == nil {
		d.logger().Info("notify", "user", userID, "event", string(event), "outcome", "posted")
		return d.Store.RecordSuccess(ctx, userID, KindDiscord)
	}
	var discordErr *DiscordError
	counts, status, text := true, 0, sanitize(postErr.Error())
	if errors.As(postErr, &discordErr) {
		counts, status, text = !discordErr.RateLimited(), discordErr.Status, discordErr.Body
	}
	disabled, err := d.Store.RecordFailure(ctx, userID, KindDiscord, text, counts)
	d.logger().Warn("notify", "user", userID, "event", string(event), "outcome", "failed",
		"status", status, "disabled", disabled)
	return err
}
