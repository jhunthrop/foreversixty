// api/internal/guilds/raids_attendance.go
//
// Attendance rows (guild_raid_attendance): read the sign-ups and the matched log, apply the
// rule in raids_attendance_compute.go, store the rows. An officer's excused flag survives
// every recompute for a character who still has a row.
package guilds

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

// RecomputeAttendance rebuilds raidID's attendance rows from its sign-ups and its log. A raid
// with no log has no rows. Excused flags are kept for characters that still have a row. A log
// that is not evidence (made private or flagged after the match) leaves the stored rows as they
// are: no reader counts them while it is not, and they count again once it is public.
func (s *Store) RecomputeAttendance(ctx context.Context, raidID int64) error {
	var guildID int64
	var reportID *string
	var evidence bool
	err := s.Pool.QueryRow(ctx, `
		select g.guild_id, g.report_id, coalesce(`+presence.EvidenceReport("r")+`, true)
		from guild_raids g left join reports r on r.id = g.report_id where g.id = $1`, raidID).Scan(&guildID, &reportID, &evidence)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotFound
	}
	if err != nil {
		return fmt.Errorf("guilds: recompute attendance: read raid: %w", err)
	}
	if reportID == nil {
		return s.replaceAttendance(ctx, raidID, nil, nil)
	}
	if !evidence {
		return nil
	}
	rows, err := s.attendanceRowsOf(ctx, guildID, raidID, *reportID)
	if err != nil {
		return err
	}
	return s.replaceAttendance(ctx, raidID, reportID, rows)
}

// attendanceRowsOf computes raidID's rows against reportID, writing nothing.
func (s *Store) attendanceRowsOf(ctx context.Context, guildID, raidID int64, reportID string) ([]AttendanceRow, error) {
	in, err := s.attendanceInputFor(ctx, guildID, raidID, reportID)
	if err != nil {
		return nil, err
	}
	return computeAttendance(in), nil
}

func (s *Store) attendanceInputFor(ctx context.Context, guildID, raidID int64, reportID string) (attendanceInput, error) {
	signups, err := s.RaidSignups(ctx, raidID)
	if err != nil {
		return attendanceInput{}, err
	}
	chars, err := s.raidRoster(ctx, guildID, "")
	if err != nil {
		return attendanceInput{}, err
	}
	roster := make(map[string]raidRosterChar, len(chars))
	for _, c := range chars {
		roster[c.CharacterKey] = c
	}
	present, err := s.logPresentKeys(ctx, reportID)
	if err != nil {
		return attendanceInput{}, err
	}
	if err := s.fillLogFigures(ctx, reportID, present); err != nil {
		return attendanceInput{}, err
	}
	total, err := s.encounterPulls(ctx, reportID)
	if err != nil {
		return attendanceInput{}, err
	}
	return attendanceInput{signups: signups, roster: roster, present: present, pullsTotal: total}, nil
}

// logPresentKeys is who the log shows, by the presence rules (a private or flagged log shows
// nobody).
func (s *Store) logPresentKeys(ctx context.Context, reportID string) (map[string]logPresence, error) {
	rows, err := presence.In(ctx, s.Pool, []string{reportID}, nil)
	if err != nil {
		return nil, fmt.Errorf("guilds: attendance presence: %w", err)
	}
	out := map[string]logPresence{}
	for _, r := range rows {
		out[r.Key] = logPresence{}
	}
	return out, nil
}

// fillLogFigures adds each present character's name, class, spec, pulls and log-relative
// first sighting from their fight_metrics rows.
func (s *Store) fillLogFigures(ctx context.Context, reportID string, present map[string]logPresence) error {
	if len(present) == 0 {
		return nil
	}
	rows, err := s.Pool.Query(ctx, `
		select m.player_key, max(m.player_name), max(coalesce(m.class, '')), max(coalesce(m.spec, '')),
		       count(distinct m.fight_index),
		       ((min(f.start_ms) - (select min(start_ms) from fights where report_id = $1 and encounter_id is not null)) / 1000)::int
		from fight_metrics m
		join fights f on f.report_id = m.report_id and f.fight_index = m.fight_index
		join reports r on r.id = m.report_id
		where m.report_id = $1 and m.state = 'ok' and `+presence.EvidenceReport("r")+`
		group by m.player_key`, reportID)
	if err != nil {
		return fmt.Errorf("guilds: attendance figures: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var key string
		var p logPresence
		var offset *int // null when the log has no encounter fight to measure from
		if err := rows.Scan(&key, &p.Name, &p.Class, &p.Spec, &p.PullsPresent, &offset); err != nil {
			return fmt.Errorf("guilds: attendance figures: %w", err)
		}
		if _, ok := present[key]; ok {
			p.FirstSeenOffsetS = offset
			present[key] = p
		}
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("guilds: attendance figures: %w", err)
	}
	return nil
}

func (s *Store) encounterPulls(ctx context.Context, reportID string) (int, error) {
	var n int
	if err := s.Pool.QueryRow(ctx,
		`select count(*) from fights where report_id = $1 and encounter_id is not null`, reportID).Scan(&n); err != nil {
		return 0, fmt.Errorf("guilds: attendance pulls: %w", err)
	}
	return n, nil
}

// replaceAttendance swaps raidID's rows for rows in one transaction, keeping excused flags. It
// locks the raid and writes only while the raid still holds the log the rows were computed
// from (want; nil for none), so a clear or re-match that lands first is never followed by
// stale rows: the later recompute is the one that counts.
func (s *Store) replaceAttendance(ctx context.Context, raidID int64, want *string, rows []AttendanceRow) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("guilds: attendance: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	var current *string
	if err := tx.QueryRow(ctx, `select report_id from guild_raids where id = $1 for update`, raidID).Scan(&current); err != nil {
		return fmt.Errorf("guilds: attendance: lock raid: %w", err)
	}
	if (current == nil) != (want == nil) || (current != nil && *current != *want) {
		return nil
	}
	if err := writeAttendanceRows(ctx, tx, raidID, rows); err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("guilds: attendance: commit: %w", err)
	}
	return nil
}

// writeAttendanceRows swaps raidID's rows for rows inside tx, keeping excused flags of the
// characters that keep a row.
func writeAttendanceRows(ctx context.Context, tx pgx.Tx, raidID int64, rows []AttendanceRow) error {
	keys := make([]string, len(rows))
	for i, r := range rows {
		keys[i] = r.CharacterKey
	}
	if _, err := tx.Exec(ctx, `delete from guild_raid_attendance where raid_id = $1 and not (character_key = any($2))`,
		raidID, keys); err != nil {
		return fmt.Errorf("guilds: attendance: prune: %w", err)
	}
	for _, r := range rows {
		if _, err := tx.Exec(ctx, `
			insert into guild_raid_attendance (raid_id, character_key, name, class, spec, signed, role, outcome,
				came_as, came_as_key, first_seen_offset_s, pulls_present, pulls_total)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
			on conflict (raid_id, character_key) do update set name = excluded.name, class = excluded.class,
				spec = excluded.spec, signed = excluded.signed, role = excluded.role, outcome = excluded.outcome,
				came_as = excluded.came_as, came_as_key = excluded.came_as_key, first_seen_offset_s = excluded.first_seen_offset_s,
				pulls_present = excluded.pulls_present, pulls_total = excluded.pulls_total, computed_at = now()`,
			raidID, r.CharacterKey, r.Name, r.Class, r.Spec, r.Signed, r.Role, r.Outcome, r.CameAs, r.CameAsKey,
			r.FirstSeenOffsetS, r.PullsPresent, r.PullsTotal); err != nil {
			return fmt.Errorf("guilds: attendance: write: %w", err)
		}
	}
	return nil
}

// SetExcused sets the officer's excused flag on one attendance row. ErrNotFound when the raid
// has no row for the character.
func (s *Store) SetExcused(ctx context.Context, raidID int64, characterKey string, excused bool) error {
	tag, err := s.Pool.Exec(ctx,
		`update guild_raid_attendance set excused = $3 where raid_id = $1 and character_key = $2`,
		raidID, characterKey, excused)
	if err != nil {
		return fmt.Errorf("guilds: set excused: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

// AttendanceRows reads raidID's stored rows, ordered by name then key.
func (s *Store) AttendanceRows(ctx context.Context, raidID int64) ([]AttendanceRow, error) {
	byRaid, err := s.attendanceRowsFor(ctx, []int64{raidID})
	if err != nil {
		return nil, err
	}
	return byRaid[raidID], nil
}

// attendanceRowsFor reads the stored rows of every raid in raidIDs in one query, each raid's
// ordered by name then key. A raid with no rows maps to an empty, non-nil slice.
func (s *Store) attendanceRowsFor(ctx context.Context, raidIDs []int64) (map[int64][]AttendanceRow, error) {
	rows, err := s.Pool.Query(ctx, `
		select raid_id, character_key, name, class, spec, signed, role, outcome, came_as, came_as_key, first_seen_offset_s,
		       pulls_present, pulls_total, excused
		from guild_raid_attendance where raid_id = any($1) order by raid_id, name, character_key`, raidIDs)
	if err != nil {
		return nil, fmt.Errorf("guilds: attendance rows: %w", err)
	}
	defer rows.Close()
	out := make(map[int64][]AttendanceRow, len(raidIDs))
	for _, id := range raidIDs {
		out[id] = []AttendanceRow{}
	}
	for rows.Next() {
		var raidID int64
		var r AttendanceRow
		if err := rows.Scan(&raidID, &r.CharacterKey, &r.Name, &r.Class, &r.Spec, &r.Signed, &r.Role, &r.Outcome,
			&r.CameAs, &r.CameAsKey, &r.FirstSeenOffsetS, &r.PullsPresent, &r.PullsTotal, &r.Excused); err != nil {
			return nil, fmt.Errorf("guilds: attendance rows: %w", err)
		}
		out[raidID] = append(out[raidID], r)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: attendance rows: %w", err)
	}
	return out, nil
}
