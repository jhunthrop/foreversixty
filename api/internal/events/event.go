package events

import (
	"context"
	"errors"
	"fmt"
)

// AudienceKind is how an event's recipients are found.
type AudienceKind string

const (
	// ToUsers is the listed accounts.
	ToUsers AudienceKind = "users"
	// ToGuildMembers is every verified member account of the guild.
	ToGuildMembers AudienceKind = "guild_members"
	// ToUnanswered is the guild's verified member accounts with no sign-up on the raid.
	ToUnanswered AudienceKind = "unanswered"
	// ToRaidSigned is every account with a sign-up row on the raid.
	ToRaidSigned AudienceKind = "raid_signed"
	// ToRaidInLate is every account signed In or Late on the raid.
	ToRaidInLate AudienceKind = "raid_in_late"
	// ToGuildOfficers is the guild's verified officer and leader accounts.
	ToGuildOfficers AudienceKind = "guild_officers"
)

// Audience is who an event goes to. GuildID and RaidID name the guild and raid the guild and
// raid audiences read.
type Audience struct {
	Kind    AudienceKind
	UserIDs []int64
	GuildID int64
	RaidID  int64
}

// Users is the audience of the listed accounts.
func Users(ids ...int64) Audience { return Audience{Kind: ToUsers, UserIDs: ids} }

// Action is a button on the notification: In, Late or Out on a raid.
type Action struct {
	// ID is the action's identifier the apps register ("signup_in").
	ID string `json:"id"`
	// Title is the button's label.
	Title string `json:"title"`
	// Method and Path are the API request the app makes with its device token.
	Method string `json:"method"`
	Path   string `json:"path"`
	// Body is the request's JSON body.
	Body map[string]string `json:"body,omitempty"`
}

// Event is one thing that happened, typed by Kind. DedupeKey makes Publish idempotent per
// recipient and channel: a key a recipient was already sent on a channel is not sent there
// again. Thread groups notifications (one per raid); Collapse lets a later notification
// replace an earlier one of the same kind for the same raid.
type Event struct {
	Kind      Kind
	Audience  Audience
	DedupeKey string
	Title     string
	Body      string
	// Link is the deep link, an https://foreversixty.gg/... URL the apps open.
	Link     string
	Thread   string
	Collapse string
	Actions  []Action
	// Data is extra key-value data the app reads (raid id, guild id).
	Data map[string]string
	// Source is the producer's typed fact, for a channel that renders it in full (the
	// Discord embeds read notify.SimFinished and notify.UploadParsed). nil otherwise.
	Source any
}

// Publisher is what a producer calls; Router implements it.
type Publisher interface {
	Publish(ctx context.Context, e Event) error
}

var errNoDedupe = errors.New("events: an event needs a dedupe key")

// Validate fails fast on an event the router cannot route.
func (e Event) Validate() error {
	if _, ok := Lookup(e.Kind); !ok {
		return fmt.Errorf("events: unknown kind %q", e.Kind)
	}
	if e.DedupeKey == "" {
		return errNoDedupe
	}
	switch e.Audience.Kind {
	case ToUsers:
		return nil
	case ToGuildMembers, ToGuildOfficers:
		if e.Audience.GuildID == 0 {
			return fmt.Errorf("events: %s needs a guild", e.Audience.Kind)
		}
		return nil
	case ToUnanswered, ToRaidSigned, ToRaidInLate:
		if e.Audience.RaidID == 0 || e.Audience.GuildID == 0 {
			return fmt.Errorf("events: %s needs a guild and a raid", e.Audience.Kind)
		}
		return nil
	}
	return fmt.Errorf("events: unknown audience %q", e.Audience.Kind)
}

// RaidThread is the thread and collapse id family of one raid.
func RaidThread(raidID int64) string { return fmt.Sprintf("raid-%d", raidID) }

// RaidCollapse is the collapse id of kind on one raid: a later one replaces it.
func RaidCollapse(kind Kind, raidID int64) string { return fmt.Sprintf("%s-%d", kind, raidID) }
