// Package notify owns the outbound notification channels a member can switch
// on (Discord webhooks now): the sealed-credential store, the URL rule, the
// poster and the failure counter.
package notify

import (
	"errors"
	"time"
)

// KindDiscord is the channel kind of a Discord incoming webhook.
const KindDiscord = "discord"

// EventName is one thing a channel can be told about.
type EventName string

const (
	EventSimFinished  EventName = "sim_finished"
	EventUploadParsed EventName = "upload_parsed"
	// EventUploadPrivate is a modifier of EventUploadParsed, not an event: a
	// private report posts only when it is on.
	EventUploadPrivate EventName = "upload_private"
)

// MaxFailures is how many counted failures in a row switch a webhook off.
const MaxFailures = 5

// ErrNotFound is returned when a member has no webhook of the asked kind.
var ErrNotFound = errors.New("notify: no webhook saved")

// Webhook is a saved channel as the page and the jobs see it. The token is
// deliberately absent; Store.TokenBox is the one read that returns it.
type Webhook struct {
	UserID    int64
	Kind      string
	ChannelID string
	// Name is the webhook's own name at Discord; nil when Discord gave none.
	Name       *string
	Events     map[EventName]bool
	Failures   int
	LastError  string
	DisabledAt *time.Time
	UpdatedAt  time.Time
}
