// api/internal/guilds/raids_events.go
//
// The raid events of the player notification catalogue (docs/superpowers/plans/
// 2026-10-30-mobile-apps-lanes.md): raid.posted, raid.changed, raid.roster_set and
// raid.late_change from the raid writes, and the time-based raid.signup_nudge,
// raid.starting and raid.short from the scheduler (raids_events_schedule.go). They go
// through s.Events beside, never instead of, the guild's Discord channel post. Each is
// published after its write commits and never fails the request.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/events"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// raidCopy is the words of the raid notifications. Every line names only what the raid page
// shows a member.
var raidEventCopy = struct {
	Posted, Nudge, Moved, Cancelled, RosterIn, RosterBench, RosterOther string
	Starting, Short, ShortOpen, LateChange                              string
	MovedIn, MovedBench, MovedLate, MovedOut                            string
	ActionIn, ActionLate, ActionOut                                     string
}{
	Posted:      "Sign-ups are open: %s.",
	Nudge:       "%s starts in 24 hours. Are you coming?",
	Moved:       "Moved to %s.",
	Cancelled:   "This raid was cancelled.",
	RosterIn:    "The roster is set. You are In.",
	RosterBench: "The roster is set. You are on the bench.",
	RosterOther: "The roster is set.",
	Starting:    "Starts in %d minutes.",
	Short:       "%d of %d signed, 24 hours out.",
	ShortOpen:   "%d signed, 24 hours out.",
	LateChange:  "%s changed to Out, %s before the start.",
	MovedIn:     "An officer moved %s to In.",
	MovedBench:  "An officer moved %s to the bench.",
	MovedLate:   "An officer marked %s Late.",
	MovedOut:    "An officer took %s off the roster.",
	ActionIn:    "In",
	ActionLate:  "Late",
	ActionOut:   "Out",
}

// raidWhenLayout is the start as a notification words it, in the guild's raid zone.
const raidWhenLayout = "Mon 2 Jan, 15:04 MST"

// SignupCharacterPlaceholder stands for the account's character in an action's path: the app
// puts the {region}/{ruleset}/{name} of the character it answers for in its place.
const SignupCharacterPlaceholder = "{character}"

// raidWhen is p's start in its zone (UTC when the guild set none).
func raidWhen(p notify.RaidPosted) string {
	loc := time.UTC
	if p.Timezone != "" {
		if l, err := time.LoadLocation(p.Timezone); err == nil {
			loc = l
		}
	}
	return p.StartsAt.In(loc).Format(raidWhenLayout)
}

// signupActions are the In, Late and Out buttons: the sign-up PUT the app calls with its
// device token.
func signupActions(guildID, raidID int64) []events.Action {
	path := fmt.Sprintf("/v1/guilds/%d/raids/%d/signups/%s", guildID, raidID, SignupCharacterPlaceholder)
	action := func(state, title string) events.Action {
		return events.Action{ID: "signup_" + state, Title: title, Method: "PUT", Path: path, Body: map[string]string{"state": state}}
	}
	return []events.Action{
		action(SignupIn, raidEventCopy.ActionIn), action(SignupLate, raidEventCopy.ActionLate), action(SignupOut, raidEventCopy.ActionOut),
	}
}

// raidEvent is the parts every raid event shares: the raid's title, its page, one thread per
// raid and a collapse id per kind and raid.
func raidEvent(kind events.Kind, guildID int64, p notify.RaidPosted, key string, aud events.Audience) events.Event {
	return events.Event{
		Kind: kind, Audience: aud, DedupeKey: key, Title: p.Title, Link: p.RaidURL,
		Thread: events.RaidThread(p.RaidID), Collapse: events.RaidCollapse(kind, p.RaidID),
		Data: map[string]string{"guild_id": strconv.FormatInt(guildID, 10), "raid_id": strconv.FormatInt(p.RaidID, 10)},
	}
}

// raidKey is a dedupe key of kind for the raid at its current start, so a moved raid is told
// again.
func raidKey(kind events.Kind, p notify.RaidPosted, extra ...string) string {
	key := fmt.Sprintf("%s:%d:%d", kind, p.RaidID, p.StartsAt.Unix())
	for _, e := range extra {
		key += ":" + e
	}
	return key
}

func raidPostedEvent(guildID int64, p notify.RaidPosted) events.Event {
	e := raidEvent(events.RaidPosted, guildID, p, fmt.Sprintf("%s:%d", events.RaidPosted, p.RaidID),
		events.Audience{Kind: events.ToGuildMembers, GuildID: guildID})
	e.Body = fmt.Sprintf(raidEventCopy.Posted, raidWhen(p))
	e.Actions = signupActions(guildID, p.RaidID)
	return e
}

func raidNudgeEvent(guildID int64, p notify.RaidPosted) events.Event {
	e := raidEvent(events.RaidSignupNudge, guildID, p, raidKey(events.RaidSignupNudge, p),
		events.Audience{Kind: events.ToUnanswered, GuildID: guildID, RaidID: p.RaidID})
	e.Body = fmt.Sprintf(raidEventCopy.Nudge, p.Title)
	e.Actions = signupActions(guildID, p.RaidID)
	return e
}

// raidChangedEvent is a move (was set) or a cancel (cancelled).
func raidChangedEvent(guildID int64, p notify.RaidPosted, cancelled bool) events.Event {
	extra, body := "moved", fmt.Sprintf(raidEventCopy.Moved, raidWhen(p))
	if cancelled {
		extra, body = "cancelled", raidEventCopy.Cancelled
	}
	e := raidEvent(events.RaidChanged, guildID, p, raidKey(events.RaidChanged, p, extra),
		events.Audience{Kind: events.ToRaidSigned, GuildID: guildID, RaidID: p.RaidID})
	e.Body = body
	return e
}

// rosterSetEvent tells one account its place; lockedAt tells one lock from the next.
func rosterSetEvent(guildID int64, p notify.RaidPosted, userID int64, state string, lockedAt time.Time) events.Event {
	e := raidEvent(events.RaidRosterSet, guildID, p, raidKey(events.RaidRosterSet, p, strconv.FormatInt(lockedAt.Unix(), 10)),
		events.Users(userID))
	switch state {
	case SignupIn, SignupLate:
		e.Body = raidEventCopy.RosterIn
	case SignupBench:
		e.Body = raidEventCopy.RosterBench
	default:
		e.Body = raidEventCopy.RosterOther
	}
	return e
}

func raidStartingEvent(guildID int64, p notify.RaidPosted, lead time.Duration) events.Event {
	e := raidEvent(events.RaidStarting, guildID, p, raidKey(events.RaidStarting, p),
		events.Audience{Kind: events.ToRaidInLate, GuildID: guildID, RaidID: p.RaidID})
	e.Body = fmt.Sprintf(raidEventCopy.Starting, int(lead.Minutes()))
	return e
}

func raidShortEvent(guildID int64, p notify.RaidPosted) events.Event {
	e := raidEvent(events.RaidShort, guildID, p, raidKey(events.RaidShort, p),
		events.Audience{Kind: events.ToGuildOfficers, GuildID: guildID})
	signed := p.Counts.In + p.Counts.Late
	e.Body = fmt.Sprintf(raidEventCopy.ShortOpen, signed)
	if p.SizeKnown {
		e.Body = fmt.Sprintf(raidEventCopy.Short, signed, p.Size)
	}
	return e
}

func lateChangeEvent(guildID int64, p notify.RaidPosted, characterKey, name string, before time.Duration) events.Event {
	e := raidEvent(events.RaidLateChange, guildID, p, raidKey(events.RaidLateChange, p, characterKey),
		events.Audience{Kind: events.ToGuildOfficers, GuildID: guildID})
	e.Body = fmt.Sprintf(raidEventCopy.LateChange, name, roundedHours(before))
	return e
}

// roundedHours is d as "5 hours" ("1 hour", "under an hour").
func roundedHours(d time.Duration) string {
	h := int(d.Hours())
	switch {
	case h < 1:
		return "under an hour"
	case h == 1:
		return "1 hour"
	}
	return fmt.Sprintf("%d hours", h)
}

// eventBudget bounds one request's publish: its reads, audience and every recipient's sends.
const eventBudget = 30 * time.Second

// background builds a request's events and publishes them off the request, on a detached
// context bounded by the budget (serve runs with always-on CPU, as the post drainer does).
// Failures are logged. Without s.Events nothing runs.
func (s *Service) background(ctx context.Context, build func(context.Context) []events.Event) {
	if s.Events == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), eventBudget)
	go func() {
		defer cancel()
		s.publishAll(ctx, build(ctx))
	}()
}

// publishAll publishes each event in turn; a failure is logged.
func (s *Service) publishAll(ctx context.Context, es []events.Event) {
	for _, e := range es {
		events.Quietly(ctx, s.Events, s.logger(), e)
	}
}

// raidFacts is the raid as its notifications word it; ok is false (logged) when it cannot be read.
func (s *Service) raidFacts(ctx context.Context, raidID int64, op string) (RaidPostFacts, bool) {
	facts, found, err := s.RaidPostFacts(ctx, raidID)
	if err != nil {
		s.logger().Error("guilds", "op", op, "raid_id", raidID, "err", err)
		return RaidPostFacts{}, false
	}
	return facts, found
}

// postedEvents is raid.posted for a planned raid: its sign-ups are open. The dedupe key holds
// it to once per raid, whether a save or the scheduler gets there first.
func (s *Service) postedEvents(ctx context.Context, raidID int64) []events.Event {
	facts, ok := s.raidFacts(ctx, raidID, "event_raid_posted")
	if !ok || facts.Status != RaidPlanned {
		return nil
	}
	return []events.Event{raidPostedEvent(facts.GuildID, facts.Post.RaidPosted)}
}

// publishCreated tells raid.posted for each created raid already inside the posting horizon;
// later raids are told by the scheduler as they enter it.
func (s *Service) publishCreated(ctx context.Context, views []RaidView, now time.Time) {
	var due []int64
	for _, v := range views {
		if !postAtFor(v.StartsAt, now).After(now) {
			due = append(due, v.ID)
		}
	}
	if len(due) == 0 {
		return
	}
	s.background(ctx, func(ctx context.Context) []events.Event {
		var out []events.Event
		for _, id := range due {
			out = append(out, s.postedEvents(ctx, id)...)
		}
		return out
	})
}

// publishChanged tells the signed-up players a raid moved or was cancelled.
func (s *Service) publishChanged(ctx context.Context, raidID int64, cancelled bool) {
	s.background(ctx, func(ctx context.Context) []events.Event {
		facts, ok := s.raidFacts(ctx, raidID, "event_raid_changed")
		if !ok {
			return nil
		}
		return []events.Event{raidChangedEvent(facts.GuildID, facts.Post.RaidPosted, cancelled)}
	})
}

// publishStatus tells a status move's event: a cancel is raid.changed, a lock sets the roster.
func (s *Service) publishStatus(ctx context.Context, raid Raid, to string) {
	switch to {
	case RaidCancelled:
		s.publishChanged(ctx, raid.ID, true)
	case RaidLocked:
		s.background(ctx, func(ctx context.Context) []events.Event { return s.rosterSetEvents(ctx, raid) })
	}
}

// rosterSetEvents tells each signed-up account its place when the officer locks the raid.
func (s *Service) rosterSetEvents(ctx context.Context, raid Raid) []events.Event {
	facts, ok := s.raidFacts(ctx, raid.ID, "event_roster_set")
	if !ok {
		return nil
	}
	places, err := s.Store.accountPlaces(ctx, raid.ID)
	if err != nil {
		s.logger().Error("guilds", "op", "event_roster_set", "raid_id", raid.ID, "err", err)
		return nil
	}
	out := make([]events.Event, 0, len(places))
	for uid, state := range places {
		out = append(out, rosterSetEvent(facts.GuildID, facts.Post.RaidPosted, uid, state, raid.UpdatedAt))
	}
	return out
}

// lateChangeWindow is how close to the start an Out tells the officers.
const lateChangeWindow = 24 * time.Hour

// publishLateChange tells the officers a raider answered Out inside the window.
func (s *Service) publishLateChange(ctx context.Context, raid Raid, characterKey, name string, now time.Time) {
	before := raid.StartsAt.Sub(now)
	if before <= 0 || before > lateChangeWindow {
		return
	}
	s.background(ctx, func(ctx context.Context) []events.Event {
		facts, ok := s.raidFacts(ctx, raid.ID, "event_late_change")
		if !ok {
			return nil
		}
		return []events.Event{lateChangeEvent(facts.GuildID, facts.Post.RaidPosted, characterKey, name, before)}
	})
}

// rosterMovedEvent tells one account an officer changed its character's place on a set
// roster; was is "" for a character the officer added.
func rosterMovedEvent(guildID int64, p notify.RaidPosted, userID int64, name, was, now string, at time.Time) events.Event {
	e := raidEvent(events.RaidRosterMoved, guildID, p,
		raidKey(events.RaidRosterMoved, p, name, was, now, strconv.FormatInt(at.UnixNano(), 10)), events.Users(userID))
	format := map[string]string{
		SignupIn: raidEventCopy.MovedIn, SignupBench: raidEventCopy.MovedBench,
		SignupLate: raidEventCopy.MovedLate, SignupOut: raidEventCopy.MovedOut,
	}[now]
	e.Body = fmt.Sprintf(format, name)
	return e
}

// rosterPlaceBefore is the character's place before an officer's write to a set (locked)
// roster for another account: "" when the officer adds it. nil for every other write, which
// raid.roster_moved does not tell.
func (s *Service) rosterPlaceBefore(ctx context.Context, raid Raid, cmd SignupCommand, char raidRosterChar, asOfficer bool) (*string, error) {
	if s.Events == nil || !asOfficer || raid.Status != RaidLocked || char.userID == cmd.ActorID {
		return nil, nil
	}
	state, err := s.Store.signupState(ctx, raid.ID, char.CharacterKey)
	if err != nil {
		return nil, err
	}
	return &state, nil
}

// publishRosterMoved tells the account its character's place on the set roster changed.
func (s *Service) publishRosterMoved(ctx context.Context, raid Raid, userID int64, name, was, now string) {
	at := time.Now()
	s.background(ctx, func(ctx context.Context) []events.Event {
		facts, ok := s.raidFacts(ctx, raid.ID, "event_roster_moved")
		if !ok {
			return nil
		}
		return []events.Event{rosterMovedEvent(facts.GuildID, facts.Post.RaidPosted, userID, name, was, now, at)}
	})
}

// signupState is a character's sign-up state on a raid; "" when it has none.
func (s *Store) signupState(ctx context.Context, raidID int64, characterKey string) (string, error) {
	var state string
	err := s.Pool.QueryRow(ctx, `select state from guild_raid_signups where raid_id = $1 and character_key = $2`,
		raidID, characterKey).Scan(&state)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("guilds: signup state: %w", err)
	}
	return state, nil
}

// accountPlaces is each signed-up account's best place on a raid: In over Late over Bench
// over Out.
func (s *Store) accountPlaces(ctx context.Context, raidID int64) (map[int64]string, error) {
	rows, err := s.Pool.Query(ctx, `
		select distinct on (signed_by) signed_by, state from guild_raid_signups where raid_id = $1
		order by signed_by, case state when 'in' then 0 when 'late' then 1 when 'bench' then 2 else 3 end`, raidID)
	if err != nil {
		return nil, fmt.Errorf("guilds: account places: %w", err)
	}
	defer rows.Close()
	out := map[int64]string{}
	for rows.Next() {
		var uid int64
		var state string
		if err := rows.Scan(&uid, &state); err != nil {
			return nil, fmt.Errorf("guilds: account places: %w", err)
		}
		out[uid] = state
	}
	return out, rows.Err()
}
