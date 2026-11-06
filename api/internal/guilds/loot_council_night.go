// api/internal/guilds/loot_council_night.go
//
// The loot council's drop identity (design/specs/2026-11-05-web-layout.md §7.2): a drop is the
// guild, the night, the encounter and the item, the way loot_awards keys an award. A night is a
// planned raid, else a logged night's report, else no night yet (an award made outside every
// raid with no log, which resolveAwardNight files under no night). The API names a night with
// one opaque key: "raid:<id>", "report:<id>" or "live".
package guilds

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

// Night key prefixes and the key of no night yet.
const (
	nightKeyRaid   = "raid:"
	nightKeyReport = "report:"
	nightKeyLive   = "live"
)

// ErrInvalidNight is a night key that is malformed or names no Onyxia night of the guild.
var ErrInvalidNight = errors.New("guilds: invalid loot night")

// key is the night's opaque key.
func (n awardNight) key() string {
	switch {
	case n.raidID != nil:
		return nightKeyRaid + strconv.FormatInt(*n.raidID, 10)
	case n.reportID != nil:
		return nightKeyReport + *n.reportID
	default:
		return nightKeyLive
	}
}

// parseNightKey reads a night key; false when it is malformed.
func parseNightKey(key string) (awardNight, bool) {
	switch {
	case key == nightKeyLive:
		return awardNight{}, true
	case strings.HasPrefix(key, nightKeyRaid):
		id, err := strconv.ParseInt(strings.TrimPrefix(key, nightKeyRaid), 10, 64)
		if err != nil || id <= 0 {
			return awardNight{}, false
		}
		return awardNight{raidID: &id}, true
	case strings.HasPrefix(key, nightKeyReport):
		id := strings.TrimPrefix(key, nightKeyReport)
		if id == "" {
			return awardNight{}, false
		}
		return awardNight{reportID: &id}, true
	}
	return awardNight{}, false
}

// liveCutoff is how far back a vote, a flag or an award filed under no night still counts for
// the live night: one raid night, the rule tonightsAwards applies.
func liveCutoff(now time.Time) time.Time { return now.Add(-raidlessAwardNight) }

// councilNight is the night a council read or write names: the current night (resolveAwardNight)
// for an empty key, else the key's night once it is checked to be the guild's. A raid must be the
// guild's Onyxia raid and not cancelled; a report must be the guild's unflagged log, and a report
// a planned raid holds reads as that raid.
func (s *Store) councilNight(ctx context.Context, guildID int64, key string, now time.Time) (awardNight, error) {
	if key == "" {
		return s.resolveAwardNight(ctx, guildID, onyxiaEncounterID, nil, now)
	}
	night, ok := parseNightKey(key)
	if !ok {
		return awardNight{}, ErrInvalidNight
	}
	switch {
	case night.raidID != nil:
		return night, s.checkCouncilRaid(ctx, guildID, *night.raidID)
	case night.reportID != nil:
		return s.councilReportNight(ctx, guildID, *night.reportID)
	}
	return night, nil
}

func (s *Store) checkCouncilRaid(ctx context.Context, guildID, raidID int64) error {
	var ok bool
	if err := s.Pool.QueryRow(ctx,
		`select exists(select 1 from guild_raids where id = $1 and guild_id = $2 and zone = $3 and status <> 'cancelled')`,
		raidID, guildID, onyxiaZoneKey).Scan(&ok); err != nil {
		return fmt.Errorf("guilds: council night: %w", err)
	}
	if !ok {
		return ErrInvalidNight
	}
	return nil
}

func (s *Store) councilReportNight(ctx context.Context, guildID int64, reportID string) (awardNight, error) {
	var raidID *int64
	err := s.Pool.QueryRow(ctx, `
		select (select g.id from guild_raids g
		        where g.report_id = r.id and g.guild_id = $2 and g.zone = $3 and g.status <> 'cancelled'
		        order by g.starts_at desc limit 1)
		from reports r where r.id = $1 and r.guild_id = $2 and `+presence.NotFlagged("r"),
		reportID, guildID, onyxiaZoneKey).Scan(&raidID)
	if errors.Is(err, pgx.ErrNoRows) {
		return awardNight{}, ErrInvalidNight
	}
	if err != nil {
		return awardNight{}, fmt.Errorf("guilds: council night: %w", err)
	}
	if raidID != nil {
		return awardNight{raidID: raidID}, nil
	}
	return awardNight{reportID: &reportID}, nil
}

// nightArgs is the night as the three bind values nightMatchSQL reads: raid id, report id and
// the live cutoff.
func (n awardNight) args(now time.Time) []any {
	return []any{n.raidID, n.reportID, liveCutoff(now)}
}

// nightMatchSQL is the filter for a row of alias filed under the night bound at raidArg,
// reportArg and cutoffArg (awardNight.args), awardOnNight's rule: the same raid; else, both
// raid-less, the same report; else a raid-less, report-less row whose atColumn is past the cutoff.
func nightMatchSQL(alias, raidArg, reportArg, cutoffArg, atColumn string) string {
	return fmt.Sprintf(`((%[2]s::bigint is not null and %[1]s.raid_id = %[2]s)
	 or (%[2]s::bigint is null and %[1]s.raid_id is null
	     and ((%[3]s::text is not null and %[1]s.report_id = %[3]s)
	          or (%[3]s::text is null and %[1]s.report_id is null and %[1]s.%[5]s >= %[4]s))))`,
		alias, raidArg, reportArg, cutoffArg, atColumn)
}
