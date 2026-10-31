// api/internal/guilds/widget_raid.go
//
// The raid half of GET /v1/me/widget (spec 2026-10-31-mobile-wave3 §5.6, W3-C12): one
// character's guild, the next raid on the guild's calendar as the raids tab lists it, and the
// character's own answer. It reads the calendar (PlannedRaids), the sign-ups (RaidSignups) and
// the roster (raidRoster) the raid page reads, so the widget and the raid page never disagree.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
)

// Widget raid statuses: the raid page's planned reads "open" on a widget.
const (
	WidgetRaidOpen      = "open"
	WidgetRaidLocked    = "locked"
	WidgetRaidCancelled = "cancelled"
)

// WidgetGuild is the character's guild.
type WidgetGuild struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// WidgetMine is the character's own answer. State is null until it signs up; Spec is the
// picked raid spec's name, null when the class has specs and none is picked (the widget then
// says `Choose a spec`); Role is the sign-up's role, else the character's default.
type WidgetMine struct {
	State *string `json:"state"`
	Spec  *string `json:"spec"`
	Role  string  `json:"role"`
}

// WidgetNextRaid is the next raid on the guild's calendar for the widget. InviteAt is null
// until raids carry an invite time (app-v1 `invite_at`, handler pending); PullAt is the raid's
// start, which is the pull. SignedUp is In plus Late, the band line's count.
type WidgetNextRaid struct {
	RaidID    int64      `json:"raid_id"`
	GuildID   int64      `json:"guild_id"`
	GuildName string     `json:"guild_name"`
	Zone      string     `json:"zone"`
	ZoneName  string     `json:"zone_name"`
	StartsAt  time.Time  `json:"starts_at"`
	InviteAt  *time.Time `json:"invite_at"`
	PullAt    time.Time  `json:"pull_at"`
	EndsAt    time.Time  `json:"ends_at"`
	Size      *int       `json:"size"`
	SignedUp  int        `json:"signed_up"`
	Status    string     `json:"status"`
	My        WidgetMine `json:"my"`
	CanSignUp bool       `json:"can_sign_up"`
}

// WidgetLastRaid is the guild's most recent raid that was not cancelled, for `Last: {zone}`.
type WidgetLastRaid struct {
	ZoneName string    `json:"zone_name"`
	StartsAt time.Time `json:"starts_at"`
}

// WidgetRaid is everything the raid widget shows. Guild is nil when the character is on no
// guild's roster; NextRaid nil when the guild has nothing planned.
type WidgetRaid struct {
	Guild    *WidgetGuild
	NextRaid *WidgetNextRaid
	LastRaid *WidgetLastRaid
}

// WidgetRaidFor reads the raid widget for one of userID's characters at now.
func (s *Service) WidgetRaidFor(ctx context.Context, userID int64, characterKey string, now time.Time) (WidgetRaid, error) {
	guildID, found, err := s.Store.characterGuild(ctx, userID, characterKey)
	if err != nil || !found {
		return WidgetRaid{}, err
	}
	name, err := s.GuildName(ctx, guildID)
	if err != nil {
		return WidgetRaid{}, err
	}
	out := WidgetRaid{Guild: &WidgetGuild{ID: guildID, Name: name}}
	if out.LastRaid, err = s.Store.lastRaid(ctx, guildID, now); err != nil {
		return WidgetRaid{}, err
	}
	raid, found, err := s.nextCalendarRaid(ctx, guildID, now)
	if err != nil || !found {
		return out, err
	}
	out.NextRaid, err = s.widgetNextRaid(ctx, raid, name, userID, characterKey, now)
	return out, err
}

// nextCalendarRaid is the first raid of the guild's calendar (PlannedRaids: planned, locked and
// cancelled ones still ahead) that has not ended.
func (s *Service) nextCalendarRaid(ctx context.Context, guildID int64, now time.Time) (Raid, bool, error) {
	raids, err := s.Store.PlannedRaids(ctx, guildID, now)
	if err != nil {
		return Raid{}, false, err
	}
	for _, r := range raids {
		if r.EndsAt().After(now) {
			return r, true, nil
		}
	}
	return Raid{}, false, nil
}

func (s *Service) widgetNextRaid(ctx context.Context, raid Raid, guildName string, userID int64, key string, now time.Time) (*WidgetNextRaid, error) {
	signups, err := s.Store.RaidSignups(ctx, raid.ID)
	if err != nil {
		return nil, err
	}
	chars, err := s.Store.raidRoster(ctx, raid.GuildID, key)
	if err != nil {
		return nil, err
	}
	_, memberVerified, err := s.Accounts.GuildRank(ctx, raid.GuildID, userID)
	if err != nil {
		return nil, err
	}
	zone := zoneOf(raid.Zone)
	c := countSignups(signups)
	view := &WidgetNextRaid{
		RaidID: raid.ID, GuildID: raid.GuildID, GuildName: guildName, Zone: raid.Zone, ZoneName: zone.Name,
		StartsAt: raid.StartsAt.UTC(), PullAt: raid.StartsAt.UTC(), EndsAt: raid.EndsAt().UTC(),
		Size: zone.Size, SignedUp: c.In + c.Late, Status: widgetStatus(raid.Status),
	}
	if len(chars) == 0 {
		return view, nil
	}
	char := chars[0]
	view.My = widgetMine(char, signups)
	view.CanSignUp = memberVerified && char.verified && char.userID == userID &&
		raid.Status == RaidPlanned && raid.EndsAt().After(now)
	return view, nil
}

// widgetMine is the character's answer and spec as the raid page shows them.
func widgetMine(char raidRosterChar, signups []Signup) WidgetMine {
	mine := WidgetMine{Role: char.defaultRole()}
	for _, sg := range signups {
		if sg.CharacterKey == char.CharacterKey {
			state := sg.State
			mine.State, mine.Role = &state, sg.Role
			break
		}
	}
	if sp, ok := raidspec.ByKey(char.raidSpec); ok {
		name := sp.Name
		mine.Spec = &name
	} else if len(raidspec.ForClass(char.Class)) == 0 && char.Spec != "" {
		// A class with no spec to pick answers with the spec its logs name.
		spec := char.Spec
		mine.Spec = &spec
	}
	return mine
}

func widgetStatus(status string) string {
	switch status {
	case RaidLocked:
		return WidgetRaidLocked
	case RaidCancelled:
		return WidgetRaidCancelled
	default:
		return WidgetRaidOpen
	}
}

// characterGuild is the guild whose roster holds userID's character, a verified row first.
func (s *Store) characterGuild(ctx context.Context, userID int64, key string) (int64, bool, error) {
	var id int64
	err := s.Pool.QueryRow(ctx, `
		select guild_id from guild_characters
		where character_key = $1 and user_id = $2
		order by verified_at is null, verified_at limit 1`, key, userID).Scan(&id)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, false, nil
	}
	if err != nil {
		return 0, false, fmt.Errorf("guilds: character guild: %w", err)
	}
	return id, true, nil
}

// lastRaid is the guild's most recent raid that started before now and was not cancelled.
func (s *Store) lastRaid(ctx context.Context, guildID int64, now time.Time) (*WidgetLastRaid, error) {
	var zone string
	var starts time.Time
	err := s.Pool.QueryRow(ctx, `
		select zone, starts_at from guild_raids
		where guild_id = $1 and status <> 'cancelled' and starts_at < $2
		order by starts_at desc, id desc limit 1`, guildID, now).Scan(&zone, &starts)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("guilds: last raid: %w", err)
	}
	return &WidgetLastRaid{ZoneName: zoneOf(zone).Name, StartsAt: starts.UTC()}, nil
}

// LiveRaidEnd is when a raid's Live Activity ends, for a verified member of the raid's guild:
// the raid's end. found is false for an unknown raid or one the caller may not see, so a
// registration can never subscribe a stranger to a guild's raid.
func (s *Service) LiveRaidEnd(ctx context.Context, userID, raidID int64) (time.Time, bool, error) {
	var guildID int64
	err := s.Store.Pool.QueryRow(ctx, `select guild_id from guild_raids where id = $1`, raidID).Scan(&guildID)
	if errors.Is(err, pgx.ErrNoRows) {
		return time.Time{}, false, nil
	}
	if err != nil {
		return time.Time{}, false, fmt.Errorf("guilds: live raid: %w", err)
	}
	if _, verified, err := s.Accounts.GuildRank(ctx, guildID, userID); err != nil || !verified {
		return time.Time{}, false, err
	}
	raid, found, err := s.Store.GetRaid(ctx, guildID, raidID)
	if err != nil || !found {
		return time.Time{}, false, err
	}
	return raid.EndsAt(), true, nil
}
