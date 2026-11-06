// api/internal/guilds/loot_raid.go
//
// Lane G2: a loot award remembers the raid it dropped in. The body may name a raid (a done raid
// of the guild); without one the award goes to the raid being played now, else to the latest
// matched night that killed the boss this week (round 3), else to no raid until a match attaches
// it. Re-awarding the same drop in a later raid is a new row.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

const (
	// onyxiaZoneKey is the curated zone every award's encounter belongs to.
	onyxiaZoneKey = "raid:onyxias-lair"
)

// AwardRaidError is an award that names a raid that cannot take it. It is an ErrInvalidAward.
type AwardRaidError struct{ RaidID int64 }

func (e *AwardRaidError) Error() string {
	return fmt.Sprintf("guilds: raid %d is not a done Onyxia raid of this guild", e.RaidID)
}

func (e *AwardRaidError) Is(target error) bool { return target == ErrInvalidAward }

// awardNight is the night an award is filed under: a planned raid (raidID), or a logged night
// with no planned raid (reportID, raidID nil). Both nil is no night yet.
type awardNight struct {
	raidID   *int64
	reportID *string
}

// awardNightLookback is how far back the latest night that killed the boss still takes an award
// made outside every raid's match window: the morning after, or the next evening, the officer
// hands out what dropped last night (round 3 e2e). A week covers any raid lockout; an older kill
// is a past reset, and its drops are long gone.
const awardNightLookback = 7 * 24 * time.Hour

// resolveAwardNight is the night an award is filed under: the named done raid of the guild (an
// *AwardRaidError when it is not a done raid of the Onyxia zone); else the raid being played now
// (raidInPlay); else the latest matched night (a planned raid holding its log, or a logged night
// of the guild with no planned raid) whose log killed encounterID within awardNightLookback;
// else no night, and the award joins its raid when that raid's log matches (attachMatchLoot).
// The CSV, the raid page and /report all read the night this files the award under.
func (s *Store) resolveAwardNight(ctx context.Context, guildID, encounterID int64, named *int64, now time.Time) (awardNight, error) {
	if named != nil {
		var ok bool
		err := s.Pool.QueryRow(ctx,
			`select exists(select 1 from guild_raids where id = $1 and guild_id = $2 and zone = $3 and status = 'done')`,
			*named, guildID, onyxiaZoneKey).Scan(&ok)
		if err != nil {
			return awardNight{}, fmt.Errorf("guilds: award raid: %w", err)
		}
		if !ok {
			return awardNight{}, &AwardRaidError{RaidID: *named}
		}
		return awardNight{raidID: named}, nil
	}
	playing, err := s.raidInPlay(ctx, guildID, now)
	if err != nil || playing != nil {
		return awardNight{raidID: playing}, err
	}
	return s.latestKillNight(ctx, guildID, encounterID, now)
}

// raidInPlay is the raid being played now: the guild's Onyxia raid, of any status but
// cancelled, whose match window (an hour before the start to six after the end) holds now, the
// start nearest now winning; nil when none.
func (s *Store) raidInPlay(ctx context.Context, guildID int64, now time.Time) (*int64, error) {
	var id int64
	err := s.Pool.QueryRow(ctx, `
		select id from guild_raids
		where guild_id = $1 and zone = $2 and status <> 'cancelled'
		  and $3 >= starts_at - make_interval(hours => $4)
		  and $3 <= starts_at + make_interval(mins => duration_minutes) + make_interval(hours => $5)
		order by abs(extract(epoch from (starts_at - $3))), id limit 1`,
		guildID, onyxiaZoneKey, now, matchLeadHours, matchTailHours).Scan(&id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("guilds: award raid: %w", err)
	}
	return &id, nil
}

// latestKillNight is the guild's latest night whose log (never a flagged one) killed encounterID
// within awardNightLookback of now: the planned Onyxia raid holding that log, else the log
// itself when no planned raid holds it. No night when there is none.
func (s *Store) latestKillNight(ctx context.Context, guildID, encounterID int64, now time.Time) (awardNight, error) {
	var night awardNight
	err := s.Pool.QueryRow(ctx, `
		select g.id, r.id
		from fights f
		join reports r on r.id = f.report_id
		left join lateral (
		  select g.id from guild_raids g
		  where g.report_id = r.id and g.guild_id = $1 and g.zone = $3 and g.status <> 'cancelled'
		  order by g.starts_at desc limit 1
		) g on true
		where f.kill and f.encounter_id = $2 and `+presence.NotFlagged("r")+`
		  and (g.id is not null or r.guild_id = $1)
		  and to_timestamp(f.start_ms / 1000.0) between $4::timestamptz - make_interval(secs => $5) and $4::timestamptz
		order by f.start_ms desc limit 1`,
		guildID, encounterID, onyxiaZoneKey, now, awardNightLookback.Seconds()).Scan(&night.raidID, &night.reportID)
	if errors.Is(err, pgx.ErrNoRows) {
		return awardNight{}, nil
	}
	if err != nil {
		return awardNight{}, fmt.Errorf("guilds: award night: %w", err)
	}
	if night.raidID != nil {
		night.reportID = nil // the raid holds the log; the award is the raid's
	}
	return night, nil
}

// awardConflict is the ON CONFLICT clause for an award filed under night: on a raid, the one
// award per drop per raid (loot_awards_drop_idx); on a logged night with no planned raid, the one
// per drop per log (loot_awards_drop_night_idx, migration 0071). An award with no night has no
// conflict target: it is its own row, so the same item awarded on two unmatched nights keeps both.
func awardConflict(night awardNight) string {
	const update = ` do update set
		  character_key = excluded.character_key, awarded_by = excluded.awarded_by,
		  awarded_at = now(), report_id = excluded.report_id`
	switch {
	case night.raidID != nil:
		return `on conflict (guild_id, raid_id, encounter_id, item_id)` + update
	case night.reportID != nil:
		return `on conflict (guild_id, report_id, encounter_id, item_id) where raid_id is null and report_id is not null` + update
	default:
		return ``
	}
}

// upsertAward writes the award, replacing the earlier one for the same drop on the same night
// (awardConflict); an award with no night is always a new row.
func (s *Store) upsertAward(ctx context.Context, guildID, encounterID int64, itemID int, characterKey string, awardedBy int64, night awardNight) (int64, error) {
	var id int64
	err := s.Pool.QueryRow(ctx, `
		insert into loot_awards (guild_id, encounter_id, item_id, character_key, awarded_by, raid_id, report_id)
		values ($1, $2, $3, $4, $5, $6, $7)
		`+awardConflict(night)+`
		returning id`, guildID, encounterID, itemID, characterKey, awardedBy, night.raidID, night.reportID).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("guilds: create loot award: %w", err)
	}
	return id, nil
}
