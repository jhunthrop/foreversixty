// Package push sends the apps' push notifications: the PushSender interface with its APNs
// (HTTP/2, token-based .p8 key) and FCM HTTP v1 implementations and a fake for tests, the
// push_tokens table (migration 0062) with its register and unregister routes, and the
// events.Sink that turns an event into one message per registered token.
package push

import (
	"context"
	"errors"
	"fmt"

	"github.com/jhunthrop/foreversixty/api/internal/events"
)

// Platform is the app a token belongs to.
type Platform string

const (
	IOS     Platform = "ios"
	Android Platform = "android"
)

// Known reports whether p is a platform the senders serve.
func (p Platform) Known() bool { return p == IOS || p == Android }

// ErrInvalidToken is a token the push service says is gone or never was: the sink deletes it.
var ErrInvalidToken = errors.New("push: the token is not valid")

// Message is one notification to one token.
type Message struct {
	Token    string
	Platform Platform
	Title    string
	Body     string
	// Link is the deep link the app opens on a tap.
	Link string
	// Thread groups the notifications of one raid; Collapse replaces an earlier one with the
	// same id (APNs apns-collapse-id, FCM collapse_key and the Android notification tag).
	Thread   string
	Collapse string
	// Category names the button set the app registered (iOS category, Android action set);
	// Actions carry each button's API request.
	Category string
	Actions  []events.Action
	Data     map[string]string
	// Kind is the event's kind, sent as data so the app can route it.
	Kind events.Kind
	// TimeSensitive breaks through Focus (APNs interruption-level time-sensitive) and goes at
	// high priority on FCM; any other message goes at normal priority there.
	TimeSensitive bool
}

// Sender sends one message. It returns ErrInvalidToken (wrapped or not) when the service
// rejected the token itself. An error never carries a key or a token.
type Sender interface {
	Send(ctx context.Context, m Message) error
}

// ServiceError is a push service's refusal: the HTTP status and the service's own reason.
type ServiceError struct {
	Service string
	Status  int
	Reason  string
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("push: %s answered %d %s", e.Service, e.Status, e.Reason)
}

// RaidSignupCategory is the button set of the raid events: In, Late and Out.
const RaidSignupCategory = "RAID_SIGNUP"

// maxCollapseBytes is APNs' limit on apns-collapse-id.
const maxCollapseBytes = 64

// collapseID is c cut to the APNs limit.
func collapseID(c string) string {
	if len(c) > maxCollapseBytes {
		return c[:maxCollapseBytes]
	}
	return c
}
