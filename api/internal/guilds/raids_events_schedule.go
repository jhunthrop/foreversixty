// api/internal/guilds/raids_events_schedule.go
//
// The time-based raid events, run by the raid-posts job every ten minutes: raid.posted as a
// raid enters the posting horizon, raid.signup_nudge and raid.short 24 hours out, and
// raid.starting at the reminder lead. Each fires while its due time is inside the catch-up
// window, so a missed run is made up; the events' dedupe keys hold each to once per event,
// raid and player.
package guilds

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/events"
)

const (
	// eventCatchUp is how long after its due time a scheduled event still fires.
	eventCatchUp = time.Hour
	// startingLead is when raid.starting goes out. Guilds have no lead setting yet, so every
	// guild has the catalogue's 60-minute default.
	startingLead = 60 * time.Minute
	// scheduleRaidLimit bounds the raids one run reads.
	scheduleRaidLimit = 1000
)

// EventScheduleRun counts the events one run published.
type EventScheduleRun struct{ Posted, Nudged, Short, Starting, Pending int }

// scheduledRaid is what the run reads to decide which events are due.
type scheduledRaid struct {
	ID       int64
	StartsAt time.Time
	Status   string
}

// dueWithin reports whether a due time at falls in (now - catch-up, now].
func dueWithin(at, now time.Time) bool {
	return !at.After(now) && at.After(now.Add(-eventCatchUp))
}

// dueKinds is the scheduled events due for r at now.
func dueKinds(r scheduledRaid, now time.Time) []events.Kind {
	var out []events.Kind
	if r.Status == RaidPlanned && dueWithin(r.StartsAt.Add(-postHorizon), now) {
		out = append(out, events.RaidPosted)
	}
	if r.Status == RaidPlanned && dueWithin(r.StartsAt.Add(-reminderLead), now) {
		out = append(out, events.RaidSignupNudge, events.RaidShort)
	}
	if r.StartsAt.After(now) && dueWithin(r.StartsAt.Add(-startingLead), now) {
		out = append(out, events.RaidStarting)
	}
	return out
}

// RunEventSchedule publishes every scheduled raid event due at now. Without s.Events it does
// nothing.
func (s *Service) RunEventSchedule(ctx context.Context, now time.Time) (EventScheduleRun, error) {
	var run EventScheduleRun
	if s.Events == nil {
		return run, nil
	}
	raids, err := s.Store.scheduledRaids(ctx, now)
	if err != nil {
		return run, err
	}
	for _, r := range raids {
		if ctx.Err() != nil {
			break
		}
		kinds := dueKinds(r, now)
		if len(kinds) == 0 {
			continue
		}
		run = run.add(kinds, s.publishDue(ctx, r.ID, kinds))
	}
	pending, err := s.publishPending(ctx, now)
	run.Pending = pending
	return run, err
}

// pendingCopy is the words of roster.pending.
var pendingCopy = struct{ Title, One, Many string }{
	Title: "%s: roster approvals",
	One:   "1 character is waiting for approval.",
	Many:  "%d characters are waiting for approval.",
}

// publishPending tells each guild's officers, once a day (UTC), how many characters wait for
// approval. A run inside an officer's quiet hours leaves the push for a later run that day.
func (s *Service) publishPending(ctx context.Context, now time.Time) (int, error) {
	pending, err := s.Store.pendingApprovals(ctx)
	if err != nil {
		return 0, err
	}
	day := now.UTC().Format(time.DateOnly)
	for _, p := range pending {
		e := s.guildBase(p.Guild)
		e.Kind, e.Audience = events.RosterPending, events.Audience{Kind: events.ToGuildOfficers, GuildID: p.Guild.ID}
		e.DedupeKey = fmt.Sprintf("roster.pending:%d:%s", p.Guild.ID, day)
		e.Title = fmt.Sprintf(pendingCopy.Title, p.Guild.Name)
		e.Body = pendingCopy.One
		if p.Count > 1 {
			e.Body = fmt.Sprintf(pendingCopy.Many, p.Count)
		}
		s.publishAll(ctx, []events.Event{e})
	}
	return len(pending), nil
}

// pendingGuild is a guild with characters waiting for approval.
type pendingGuild struct {
	Guild Guild
	Count int
}

// pendingApprovals is every guild with unverified roster characters and how many.
func (s *Store) pendingApprovals(ctx context.Context) ([]pendingGuild, error) {
	rows, err := s.Pool.Query(ctx, `
		select guild_id, count(*) from guild_characters where verified_at is null group by guild_id order by guild_id`)
	if err != nil {
		return nil, fmt.Errorf("guilds: pending approvals: %w", err)
	}
	type count struct {
		GuildID int64
		Count   int
	}
	counts, err := pgx.CollectRows(rows, pgx.RowToStructByPos[count])
	if err != nil {
		return nil, fmt.Errorf("guilds: pending approvals: %w", err)
	}
	out := make([]pendingGuild, 0, len(counts))
	for _, c := range counts {
		g, err := s.getGuild(ctx, c.GuildID)
		if err != nil {
			return nil, err
		}
		out = append(out, pendingGuild{Guild: g, Count: c.Count})
	}
	return out, nil
}

// publishDue publishes the raid's due kinds from one read of the raid; false when it could
// not be read.
func (s *Service) publishDue(ctx context.Context, raidID int64, kinds []events.Kind) bool {
	facts, ok := s.raidFacts(ctx, raidID, "event_schedule")
	if !ok {
		return false
	}
	p := facts.Post.RaidPosted
	var out []events.Event
	for _, k := range kinds {
		switch k {
		case events.RaidPosted:
			out = append(out, raidPostedEvent(facts.GuildID, p))
		case events.RaidSignupNudge:
			out = append(out, raidNudgeEvent(facts.GuildID, p))
		case events.RaidShort:
			if p.SizeKnown && p.Counts.In+p.Counts.Late < p.Size {
				out = append(out, raidShortEvent(facts.GuildID, p))
			}
		case events.RaidStarting:
			out = append(out, raidStartingEvent(facts.GuildID, p, startingLead))
		}
	}
	s.publishAll(ctx, out)
	return true
}

func (r EventScheduleRun) add(kinds []events.Kind, ok bool) EventScheduleRun {
	if !ok {
		return r
	}
	for _, k := range kinds {
		switch k {
		case events.RaidPosted:
			r.Posted++
		case events.RaidSignupNudge:
			r.Nudged++
		case events.RaidShort:
			r.Short++
		case events.RaidStarting:
			r.Starting++
		}
	}
	return r
}

// scheduledRaids is every planned or locked raid whose events can still come due: starting
// after now and within the posting horizon.
func (s *Store) scheduledRaids(ctx context.Context, now time.Time) ([]scheduledRaid, error) {
	rows, err := s.Pool.Query(ctx, `
		select id, starts_at, status from guild_raids
		where status in ('planned', 'locked') and starts_at > $1 and starts_at <= $1 + make_interval(secs => $2)
		order by starts_at limit $3`,
		now, postHorizon.Seconds(), scheduleRaidLimit)
	if err != nil {
		return nil, fmt.Errorf("guilds: scheduled raids: %w", err)
	}
	out, err := pgx.CollectRows(rows, pgx.RowToStructByPos[scheduledRaid])
	if err != nil {
		return nil, fmt.Errorf("guilds: scheduled raids: %w", err)
	}
	return out, nil
}
