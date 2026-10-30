package notify

import (
	"context"
	"errors"
	"log/slog"
	"time"

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
	// postTimeout bounds one post; zero means posterTimeout.
	postTimeout time.Duration
}

var errNoEvents = errors.New("notify: Deliver needs at least one event")

// What a member reads in last_error when the post never got an answer.
const (
	timeoutText     = "Discord did not answer in time."
	unreachableText = "Discord could not be reached."
)

// Deliverer returns a Deliverer sharing the service's wiring.
func (s *Service) Deliverer() *Deliverer {
	return &Deliverer{
		Store: s.Store, Sealer: s.Sealer, Entitlements: s.Entitlements,
		Poster: s.Poster, PublicBaseURL: s.PublicBaseURL, Log: s.Log,
	}
}

// effectiveTimeout is the injected post timeout, or the default.
func (d *Deliverer) effectiveTimeout() time.Duration {
	if d.postTimeout == 0 {
		return posterTimeout
	}
	return d.postTimeout
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
	d.deliverQuietly(ctx, e.UserID, SimFinishedEmbed(d.PublicBaseURL, e), EventSimFinished)
}

// UploadParsed posts the log message; see SimFinished. A private report
// also needs the upload_private opt-in.
func (d *Deliverer) UploadParsed(ctx context.Context, e UploadParsed) {
	d.deliverQuietly(ctx, e.UserID, UploadParsedEmbed(d.PublicBaseURL, e), uploadNeeds(e)...)
}

// uploadNeeds is the events an upload's post needs on.
func uploadNeeds(e UploadParsed) []EventName {
	if e.Private {
		return []EventName{EventUploadParsed, EventUploadPrivate}
	}
	return []EventName{EventUploadParsed}
}

func (d *Deliverer) deliverQuietly(ctx context.Context, userID int64, embed Embed, needs ...EventName) {
	if err := d.Deliver(ctx, userID, embed, needs...); err != nil {
		d.logger().Error("notify", "user", userID, "events", needs, "outcome", "error", "err", err)
	}
}

// Deliver posts embed when the member has a live webhook with every needed
// event on (the first names the event in the log) and still holds the
// entitlement. Skips are silent successes. needs must name at least one
// event. The error it returns never holds the URL, token or message text.
func (d *Deliverer) Deliver(ctx context.Context, userID int64, embed Embed, needs ...EventName) error {
	if len(needs) == 0 {
		return errNoEvents
	}
	hook, found, err := d.Store.Get(ctx, userID, KindDiscord)
	if err != nil {
		return err
	}
	if !found || hook.DisabledAt != nil || !allOn(hook.Events, needs) {
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
	postCtx, cancel := context.WithTimeout(ctx, d.effectiveTimeout())
	defer cancel()
	return d.record(ctx, userID, needs[0], d.Poster.Post(postCtx, url, embed))
}

func allOn(events map[EventName]bool, needs []EventName) bool {
	for _, need := range needs {
		if !events[need] {
			return false
		}
	}
	return true
}

// record writes the outcome on a context that outlives a cancelled job, so
// a post that happened is never left unrecorded.
func (d *Deliverer) record(ctx context.Context, userID int64, event EventName, postErr error) error {
	ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), posterTimeout)
	defer cancel()
	if postErr == nil {
		d.logger().Info("notify", "user", userID, "event", string(event), "outcome", "posted")
		return d.Store.RecordSuccess(ctx, userID, KindDiscord)
	}
	f := failureOf(postErr)
	disabled, err := d.Store.RecordFailure(ctx, userID, KindDiscord, f.text, f.counts)
	d.logger().Warn("notify", "user", userID, "event", string(event), "outcome", "failed",
		"status", f.status, "disabled", disabled)
	return err
}

// postFailure is a failed post as the five-failure rule reads it.
type postFailure struct {
	// counts is false for a 429: a rate limit never counts toward the five.
	counts bool
	status int
	// text is what last_error stores; never the URL, token or message.
	text string
}

// failureOf classifies a failed post for the account and guild rules alike.
func failureOf(postErr error) postFailure {
	var discordErr *DiscordError
	f := postFailure{counts: true, text: sanitize(postErr.Error())}
	switch {
	case errors.As(postErr, &discordErr):
		f.counts, f.status, f.text = !discordErr.RateLimited(), discordErr.Status, discordErr.Body
	case errors.Is(postErr, ErrTimeout), errors.Is(postErr, context.DeadlineExceeded):
		f.text = timeoutText
	case errors.Is(postErr, ErrUnreachable):
		f.text = unreachableText
	}
	return f
}
