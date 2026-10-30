package events

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"
)

// Sink delivers one event to one account on one channel. A sink that has nothing to deliver
// for the event (no push token, no Discord webhook, a kind it does not render) returns nil.
type Sink interface {
	Deliver(ctx context.Context, userID int64, e Event) error
}

// KindFilter is a Sink that renders only some kinds; the router skips it for the rest
// before claiming anything.
type KindFilter interface {
	Renders(kind Kind) bool
}

// RouterStore is what the router reads and writes (Store).
type RouterStore interface {
	// Recipients resolves an audience to account ids, each once.
	Recipients(ctx context.Context, a Audience) ([]int64, error)
	// PrefsOf reads each account's preferences; an account with no row has the defaults.
	PrefsOf(ctx context.Context, userIDs []int64) (map[int64]Prefs, error)
	// Claim records that key goes to the account on the channel; false when it already went.
	Claim(ctx context.Context, userID int64, key string, channel Channel) (bool, error)
	// Unclaim drops a claim whose delivery failed, so a later publish may send it.
	Unclaim(ctx context.Context, userID int64, key string, channel Channel) error
}

// Router sends a published event to each recipient's channels, per preference. A channel
// with no sink is off for everyone.
type Router struct {
	Store RouterStore
	Sinks map[Channel]Sink
	// Now is the clock quiet hours are read against; nil is time.Now.
	Now func() time.Time
	Log *slog.Logger
}

var _ Publisher = (*Router)(nil)

var errOfficerAudience = errors.New("events: an officer event goes only to a guild's officers")

func (r *Router) now() time.Time {
	if r.Now != nil {
		return r.Now()
	}
	return time.Now()
}

func (r *Router) logger() *slog.Logger {
	if r.Log != nil {
		return r.Log
	}
	return slog.Default()
}

// Publish routes e. It fails on an event it cannot route or an audience it cannot read; a
// failed delivery is logged, its claim dropped, and joined into the error returned.
func (r *Router) Publish(ctx context.Context, e Event) error {
	if err := e.Validate(); err != nil {
		return err
	}
	spec, _ := Lookup(e.Kind)
	if spec.Officer && e.Audience.Kind != ToGuildOfficers {
		return errOfficerAudience
	}
	if !r.anySinkRenders(e.Kind) {
		return nil
	}
	users, err := r.Store.Recipients(ctx, e.Audience)
	if err != nil {
		return fmt.Errorf("events: recipients of %s: %w", e.Kind, err)
	}
	if len(users) == 0 {
		return nil
	}
	prefs, err := r.Store.PrefsOf(ctx, users)
	if err != nil {
		// The defaults still deliver: an unreadable preference never drops a notification.
		r.logger().Error("events", "op", "prefs", "kind", string(e.Kind), "err", err)
		prefs = map[int64]Prefs{}
	}
	now := r.now()
	var errs []error
	for _, uid := range users {
		for _, ch := range Channels {
			if err := r.deliverOne(ctx, uid, ch, spec, prefs[uid], e, now); err != nil {
				errs = append(errs, err)
			}
		}
	}
	return errors.Join(errs...)
}

// deliverOne sends e to one account on one channel when the preference and the quiet hours
// allow it and the dedupe key is not spent. Quiet hours hold back push only.
func (r *Router) deliverOne(ctx context.Context, uid int64, ch Channel, spec Spec, p Prefs, e Event, now time.Time) error {
	sink, ok := r.Sinks[ch]
	if !ok || !p.Allows(e.Kind, ch) {
		return nil
	}
	if f, ok := sink.(KindFilter); ok && !f.Renders(e.Kind) {
		return nil
	}
	if ch == ChannelPush && !spec.QuietExempt && p.InQuiet(now) {
		return nil
	}
	claimed, err := r.Store.Claim(ctx, uid, e.DedupeKey, ch)
	if err != nil {
		// Fail open: a notification sent twice beats one never sent.
		r.logger().Error("events", "op", "claim", "kind", string(e.Kind), "user", uid, "err", err)
		claimed = true
	}
	if !claimed {
		return nil
	}
	if err := sink.Deliver(ctx, uid, e); err != nil {
		r.logger().Warn("events", "op", "deliver", "kind", string(e.Kind), "channel", string(ch), "user", uid, "err", err)
		if uerr := r.Store.Unclaim(context.WithoutCancel(ctx), uid, e.DedupeKey, ch); uerr != nil {
			r.logger().Error("events", "op", "unclaim", "kind", string(e.Kind), "user", uid, "err", uerr)
		}
		return fmt.Errorf("events: %s to %d on %s: %w", e.Kind, uid, ch, err)
	}
	return nil
}

// Quietly publishes e through p and logs a failure; producers whose own work already
// succeeded use it so a notification never fails them. A nil p publishes nothing.
func Quietly(ctx context.Context, p Publisher, log *slog.Logger, e Event) {
	if p == nil {
		return
	}
	if err := p.Publish(ctx, e); err != nil {
		if log == nil {
			log = slog.Default()
		}
		log.Error("events", "op", "publish", "kind", string(e.Kind), "err", err)
	}
}

// anySinkRenders reports whether some sink can deliver kind, so an event no channel renders
// costs no read.
func (r *Router) anySinkRenders(kind Kind) bool {
	for _, sink := range r.Sinks {
		if f, ok := sink.(KindFilter); !ok || f.Renders(kind) {
			return true
		}
	}
	return false
}
