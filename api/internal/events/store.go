package events

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store is the events tables (migration 0060) and the audience reads.
type Store struct {
	Pool *pgxpool.Pool
}

var _ RouterStore = (*Store)(nil)

// verifiedMember is the clause every guild audience shares: the account is a verified member
// of the guild now, so a player who left hears nothing more.
const verifiedMember = `exists (select 1 from guild_members gm
	where gm.guild_id = $1 and gm.user_id = %s and gm.verified_at is not null)`

// audienceSQL reads each guild and raid audience. $1 is the guild, $2 the raid (the raid
// audiences only).
var audienceSQL = map[AudienceKind]string{
	ToGuildMembers: `select user_id from guild_members where guild_id = $1 and verified_at is not null`,
	ToGuildOfficers: `select user_id from guild_members
		where guild_id = $1 and verified_at is not null and rank in ('officer', 'leader')`,
	ToRaidSigned: `select distinct s.signed_by from guild_raid_signups s join guild_raids r on r.id = s.raid_id
		where r.guild_id = $1 and s.raid_id = $2 and ` + fmt.Sprintf(verifiedMember, "s.signed_by"),
	ToRaidInLate: `select distinct s.signed_by from guild_raid_signups s join guild_raids r on r.id = s.raid_id
		where r.guild_id = $1 and s.raid_id = $2 and s.state in ('in', 'late') and ` + fmt.Sprintf(verifiedMember, "s.signed_by"),
	ToUnanswered: `select gm.user_id from guild_members gm
		where gm.guild_id = $1 and gm.verified_at is not null
		  and exists (select 1 from guild_raids r where r.id = $2 and r.guild_id = $1)
		  and not exists (select 1 from guild_raid_signups s where s.raid_id = $2 and s.signed_by = gm.user_id)`,
}

// Recipients resolves a to account ids, each once, in ascending order.
func (s *Store) Recipients(ctx context.Context, a Audience) ([]int64, error) {
	if a.Kind == ToUsers {
		out := slices.Clone(a.UserIDs)
		slices.Sort(out)
		return slices.Compact(out), nil
	}
	q, ok := audienceSQL[a.Kind]
	if !ok {
		return nil, fmt.Errorf("events: unknown audience %q", a.Kind)
	}
	args := []any{a.GuildID}
	if a.Kind != ToGuildMembers && a.Kind != ToGuildOfficers {
		args = append(args, a.RaidID)
	}
	rows, err := s.Pool.Query(ctx, q+` order by 1`, args...)
	if err != nil {
		return nil, fmt.Errorf("events: read %s: %w", a.Kind, err)
	}
	ids, err := pgx.CollectRows(rows, pgx.RowTo[int64])
	if err != nil {
		return nil, fmt.Errorf("events: read %s: %w", a.Kind, err)
	}
	return ids, nil
}

// PrefsOf reads each account's settings and overrides; an account with neither has
// DefaultPrefs.
func (s *Store) PrefsOf(ctx context.Context, userIDs []int64) (map[int64]Prefs, error) {
	out := make(map[int64]Prefs, len(userIDs))
	for _, uid := range userIDs {
		out[uid] = DefaultPrefs()
	}
	rows, err := s.Pool.Query(ctx, `select user_id, time_zone, quiet_enabled, quiet_start, quiet_end
		from notification_settings where user_id = any($1)`, userIDs)
	if err != nil {
		return nil, fmt.Errorf("events: read settings: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var uid int64
		var p Prefs
		if err := rows.Scan(&uid, &p.TimeZone, &p.Quiet.Enabled, &p.Quiet.Start, &p.Quiet.End); err != nil {
			return nil, fmt.Errorf("events: read settings: %w", err)
		}
		out[uid] = p
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("events: read settings: %w", err)
	}
	return s.withOverrides(ctx, userIDs, out)
}

func (s *Store) withOverrides(ctx context.Context, userIDs []int64, out map[int64]Prefs) (map[int64]Prefs, error) {
	rows, err := s.Pool.Query(ctx, `select user_id, kind, channel, enabled from notification_prefs where user_id = any($1)`, userIDs)
	if err != nil {
		return nil, fmt.Errorf("events: read prefs: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var uid int64
		var kind, channel string
		var on bool
		if err := rows.Scan(&uid, &kind, &channel, &on); err != nil {
			return nil, fmt.Errorf("events: read prefs: %w", err)
		}
		out[uid] = out[uid].With(Kind(kind), Channel(channel), on)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("events: read prefs: %w", err)
	}
	return out, nil
}

// Prefs is one account's preferences.
func (s *Store) Prefs(ctx context.Context, userID int64) (Prefs, error) {
	all, err := s.PrefsOf(ctx, []int64{userID})
	if err != nil {
		return Prefs{}, err
	}
	return all[userID], nil
}

// PrefChange is one event on one channel switched on or off.
type PrefChange struct {
	Kind    Kind
	Channel Channel
	Enabled bool
}

// SettingsChange is a PATCH of the settings: a nil field keeps the stored value.
type SettingsChange struct {
	TimeZone *string
	Quiet    *QuietHours
	Prefs    []PrefChange
}

var errEmptyChange = errors.New("events: nothing to change")

// SaveSettings applies change to the account's preferences in one transaction.
func (s *Store) SaveSettings(ctx context.Context, userID int64, change SettingsChange) error {
	if change.TimeZone == nil && change.Quiet == nil && len(change.Prefs) == 0 {
		return errEmptyChange
	}
	return pgx.BeginFunc(ctx, s.Pool, func(tx pgx.Tx) error {
		if change.TimeZone != nil || change.Quiet != nil {
			if err := saveSettingsRow(ctx, tx, userID, change); err != nil {
				return err
			}
		}
		for _, p := range change.Prefs {
			if _, err := tx.Exec(ctx, `insert into notification_prefs (user_id, kind, channel, enabled) values ($1, $2, $3, $4)
				on conflict (user_id, kind, channel) do update set enabled = excluded.enabled, updated_at = now()`,
				userID, string(p.Kind), string(p.Channel), p.Enabled); err != nil {
				return fmt.Errorf("events: save pref: %w", err)
			}
		}
		return nil
	})
}

func saveSettingsRow(ctx context.Context, tx pgx.Tx, userID int64, change SettingsChange) error {
	zone := change.TimeZone
	var enabled *bool
	var start, end *int
	if q := change.Quiet; q != nil {
		enabled, start, end = &q.Enabled, &q.Start, &q.End
	}
	_, err := tx.Exec(ctx, `insert into notification_settings (user_id, time_zone, quiet_enabled, quiet_start, quiet_end)
		values ($1, coalesce($2::text, $6::text), coalesce($3::boolean, false), coalesce($4::smallint, $7::smallint), coalesce($5::smallint, $8::smallint))
		on conflict (user_id) do update set
		  time_zone = coalesce($2::text, notification_settings.time_zone),
		  quiet_enabled = coalesce($3::boolean, notification_settings.quiet_enabled),
		  quiet_start = coalesce($4::smallint, notification_settings.quiet_start),
		  quiet_end = coalesce($5::smallint, notification_settings.quiet_end),
		  updated_at = now()`, userID, zone, enabled, start, end, DefaultTimeZone, DefaultQuietStart, DefaultQuietEnd)
	if err != nil {
		return fmt.Errorf("events: save settings: %w", err)
	}
	return nil
}

// Claim records key as sent to the account on channel; false when it was already.
func (s *Store) Claim(ctx context.Context, userID int64, key string, channel Channel) (bool, error) {
	tag, err := s.Pool.Exec(ctx, `insert into notification_deliveries (user_id, dedupe_key, channel) values ($1, $2, $3)
		on conflict do nothing`, userID, key, string(channel))
	if err != nil {
		return false, fmt.Errorf("events: claim: %w", err)
	}
	return tag.RowsAffected() == 1, nil
}

// Unclaim drops a claim, so the key may be sent again.
func (s *Store) Unclaim(ctx context.Context, userID int64, key string, channel Channel) error {
	if _, err := s.Pool.Exec(ctx, `delete from notification_deliveries where user_id = $1 and dedupe_key = $2 and channel = $3`,
		userID, key, string(channel)); err != nil {
		return fmt.Errorf("events: unclaim: %w", err)
	}
	return nil
}

// DeliveryRetention is how long a delivery is remembered. It outlasts every dedupe window
// (the weekly digest is the longest).
const DeliveryRetention = 30 * 24 * time.Hour

// PruneDeliveries forgets deliveries older than the retention, at most limit rows per call.
func (s *Store) PruneDeliveries(ctx context.Context, now time.Time, limit int) (int64, error) {
	tag, err := s.Pool.Exec(ctx, `delete from notification_deliveries where ctid in (
		select ctid from notification_deliveries where sent_at < $1 limit $2)`, now.Add(-DeliveryRetention), limit)
	if err != nil {
		return 0, fmt.Errorf("events: prune deliveries: %w", err)
	}
	return tag.RowsAffected(), nil
}
