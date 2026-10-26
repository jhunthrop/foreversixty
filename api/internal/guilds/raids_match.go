// api/internal/guilds/raids_match.go
//
// Lane G2: a raid learns its log. AutoMatch ties a freshly completed report to the guild raid
// it was for; SetMatch is an officer's hand pick; ClearMatch undoes either and teaches the
// matcher never to pick that log for that raid again. Candidates are in raids_match_candidates.go,
// the attendance rows a match feeds in raids_attendance.go.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

// How a raid came by its log (guild_raids.matched_by).
const (
	MatchedAuto = "auto"
	MatchedHand = "hand"
)

// The match window around a raid: a log counts from before the pull to well after the end.
const (
	matchLeadHours = 1
	matchTailHours = 6
)

// matchGuildOf is the SQL for the guild the reports row aliased alias is read as by the matcher:
// the officer's guild_id, else the guild its owner's standing named (matched_guild_id). Only
// Lane G2's readers use it; everything older reads guild_id alone.
func matchGuildOf(alias string) string {
	return "coalesce(" + alias + ".guild_id, " + alias + ".matched_guild_id)"
}

// matchGuildIs is the condition that the reports row aliased alias is read by the matcher as the
// guild in the parameter param, written so either column's index serves it.
func matchGuildIs(alias, param string) string {
	return "(" + alias + ".guild_id = " + param + " or (" + alias + ".guild_id is null and " +
		alias + ".matched_guild_id = " + param + "))"
}

// pgUniqueViolation is the SQLSTATE of a unique index refusing a row.
const pgUniqueViolation = "23505"

var (
	// ErrReportMatched is a log already matched to another raid (one log, one raid).
	ErrReportMatched = errors.New("guilds: report already matched to a raid")
	// ErrRaidNotEnded is a hand match on a planned or locked raid that has not ended.
	ErrRaidNotEnded = errors.New("guilds: raid has not ended")
)

// RaidRef names a raid in a list of logs: the raid a log is already matched to.
type RaidRef struct {
	ID       int64
	Title    string
	StartsAt time.Time
}

// logAnchor is what the matcher knows about a completed report.
type logAnchor struct {
	guildID int64
	logZone string
	at      time.Time
}

// reportAnchor reads a report's guild, log zone and anchor time (completed_at on the live
// path, created_at on an upload). ok is false when the report cannot be matched at all: it is
// unknown, has no guild or zone yet, or is not evidence (private or flagged).
func (s *Store) reportAnchor(ctx context.Context, reportID string) (logAnchor, bool, error) {
	var a logAnchor
	var guildID *int64
	var zone *string
	err := s.Pool.QueryRow(ctx, `
		select `+matchGuildOf("r")+`, r.log_zone, coalesce(r.completed_at, r.created_at)
		from reports r where r.id = $1 and `+presence.EvidenceReport("r"), reportID).Scan(&guildID, &zone, &a.at)
	if errors.Is(err, pgx.ErrNoRows) {
		return logAnchor{}, false, nil
	}
	if err != nil {
		return logAnchor{}, false, fmt.Errorf("guilds: report anchor: %w", err)
	}
	if guildID == nil || zone == nil || *zone == "" {
		return logAnchor{}, false, nil
	}
	a.guildID, a.logZone = *guildID, *zone
	return a, true, nil
}

// zoneKeysForLog lists the curated raid zones whose log name is logZone.
func zoneKeysForLog(logZone string) []string {
	keys := []string{}
	for _, z := range curatedZones {
		if z.LogZone != nil && *z.LogZone == logZone {
			keys = append(keys, z.Key)
		}
	}
	return keys
}

// AutoMatch ties a completed report to the guild's open raid it was for: a planned or locked
// raid with no log, of the zone the log names, whose window (an hour before the start to six
// after the end) holds the log's anchor, that an officer has not cleared this log from; the
// raid whose start is nearest the anchor wins. It reports false, with no error, when the
// report has no guild or zone, no raid fits, or the log is already some raid's. The candidate
// subselect skips raids another match holds locked, so two reports racing for the guild's raids
// each find one rather than the second giving up on a raid that was still free.
// The raid's attendance and the loot awarded during it are written in the same transaction; a
// failure there rolls the match back and is returned.
func (s *Store) AutoMatch(ctx context.Context, reportID string) (Raid, bool, error) {
	a, ok, err := s.reportAnchor(ctx, reportID)
	if err != nil || !ok {
		return Raid{}, false, err
	}
	keys := zoneKeysForLog(a.logZone)
	if len(keys) == 0 {
		return Raid{}, false, nil
	}
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return Raid{}, false, fmt.Errorf("guilds: auto match: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	raid, err := scanRaid(tx.QueryRow(ctx, `
		update guild_raids set report_id = $1, pre_done_status = status, status = 'done',
			matched_by = $5, matched_at = now(), updated_at = now()
		where id = (
			select c.id from guild_raids c
			where c.guild_id = $2 and c.status in ('planned','locked') and c.report_id is null
			  and c.zone = any($3)
			  and $4 >= c.starts_at - make_interval(hours => $6)
			  and $4 <= c.starts_at + make_interval(mins => c.duration_minutes) + make_interval(hours => $7)
			  and not exists (select 1 from guild_raid_rejected_reports x where x.raid_id = c.id and x.report_id = $1)
			order by abs(extract(epoch from (c.starts_at - $4))), c.id limit 1
			for update skip locked)
		  and status in ('planned','locked') and report_id is null
		  and not exists (select 1 from guild_raids o where o.report_id = $1)
		returning `+raidColumns,
		reportID, a.guildID, keys, a.at, MatchedAuto, matchLeadHours, matchTailHours))
	if errors.Is(err, pgx.ErrNoRows) || isUniqueViolation(err) {
		return Raid{}, false, nil
	}
	if err != nil {
		return Raid{}, false, fmt.Errorf("guilds: auto match: %w", err)
	}
	if err := s.finishMatch(ctx, tx, raid); err != nil {
		return Raid{}, false, err
	}
	if err := tx.Commit(ctx); err != nil {
		return Raid{}, false, fmt.Errorf("guilds: auto match: commit: %w", err)
	}
	return raid, true, nil
}

// finishMatch is the rest of every match, in the match's own transaction: the raid's
// attendance against its new log and the loot awarded during it. A failure rolls the match back.
func (s *Store) finishMatch(ctx context.Context, tx pgx.Tx, raid Raid) error {
	rows, err := s.attendanceRowsOf(ctx, raid.GuildID, raid.ID, *raid.ReportID)
	if err != nil {
		return err
	}
	if err := writeAttendanceRows(ctx, tx, raid.ID, rows); err != nil {
		return err
	}
	return attachMatchLoot(ctx, tx, raid.ID)
}

func isUniqueViolation(err error) bool {
	var pg *pgconn.PgError
	return errors.As(err, &pg) && pg.Code == pgUniqueViolation
}

// SetMatch makes reportID the log of guildID's raid raidID and, in the same transaction,
// recomputes the raid's attendance and attaches the loot awarded during it. how is MatchedAuto or MatchedHand. A planned or locked raid must have ended
// (ErrRaidNotEnded); a done raid may be re-matched to another log and keeps the status it will
// return to when cleared; a cancelled one is ErrRaidStatus. ErrNotFound when the raid is not
// the guild's or the report is not a complete report of the guild; ErrReportMatched when the
// report is already another raid's.
func (s *Store) SetMatch(ctx context.Context, guildID, raidID int64, reportID, how string) (Raid, error) {
	if how != MatchedAuto && how != MatchedHand {
		return Raid{}, fmt.Errorf("guilds: set match: unknown matched_by %q", how)
	}
	if err := s.requireGuildReport(ctx, guildID, reportID); err != nil {
		return Raid{}, err
	}
	return s.writeMatch(ctx, guildID, raidID, reportID, how)
}

func (s *Store) requireGuildReport(ctx context.Context, guildID int64, reportID string) error {
	var found bool
	if err := s.Pool.QueryRow(ctx,
		`select exists (select 1 from reports r where r.id = $1 and `+matchGuildIs("r", "$2")+` and r.status = 'complete'
		   and `+presence.EvidenceReport("r")+`)`,
		reportID, guildID).Scan(&found); err != nil {
		return fmt.Errorf("guilds: set match: read report: %w", err)
	}
	if !found {
		return ErrNotFound
	}
	return nil
}

// writeMatch is SetMatch's guarded write: one transaction reading the raid under lock, updating
// it only while its status is still one a match may replace, then finishing the match.
func (s *Store) writeMatch(ctx context.Context, guildID, raidID int64, reportID, how string) (Raid, error) {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return Raid{}, fmt.Errorf("guilds: set match: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	var status string
	var ended bool
	err = tx.QueryRow(ctx, `
		select status, starts_at + make_interval(mins => duration_minutes) <= now()
		from guild_raids where id = $1 and guild_id = $2 for update`, raidID, guildID).Scan(&status, &ended)
	if errors.Is(err, pgx.ErrNoRows) {
		return Raid{}, ErrNotFound
	}
	if err != nil {
		return Raid{}, fmt.Errorf("guilds: set match: read raid: %w", err)
	}
	if err := matchableStatus(status, ended); err != nil {
		return Raid{}, err
	}
	raid, err := scanRaid(tx.QueryRow(ctx, `
		update guild_raids set report_id = $3,
			pre_done_status = case when status = 'done' then pre_done_status else status end,
			status = 'done', matched_by = $4, matched_at = now(), updated_at = now()
		where id = $1 and guild_id = $2 and status = any($5) returning `+raidColumns,
		raidID, guildID, reportID, how, []string{RaidPlanned, RaidLocked, RaidDone}))
	if isUniqueViolation(err) {
		return Raid{}, ErrReportMatched
	}
	if err != nil {
		return Raid{}, fmt.Errorf("guilds: set match: write: %w", err)
	}
	if err := s.finishMatch(ctx, tx, raid); err != nil {
		return Raid{}, err
	}
	if err := tx.Commit(ctx); err != nil {
		return Raid{}, fmt.Errorf("guilds: set match: commit: %w", err)
	}
	return raid, nil
}

func matchableStatus(status string, ended bool) error {
	switch {
	case status == RaidCancelled:
		return ErrRaidStatus
	case status != RaidDone && !ended:
		return ErrRaidNotEnded
	}
	return nil
}

// ClearMatch removes the log from guildID's done raid: the status returns to what it was before
// the match, the attendance rows go, and the cleared log is recorded (by is the officer) so
// the auto-matcher never picks it for this raid again. ErrNotFound when the guild has no such
// raid; ErrRaidStatus when the raid has no match.
func (s *Store) ClearMatch(ctx context.Context, guildID, raidID, by int64) (Raid, error) {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return Raid{}, fmt.Errorf("guilds: clear match: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	var reportID *string
	err = tx.QueryRow(ctx, `select report_id from guild_raids where id = $1 and guild_id = $2 and status = 'done' for update`,
		raidID, guildID).Scan(&reportID)
	if errors.Is(err, pgx.ErrNoRows) {
		return Raid{}, s.raidWriteErr(ctx, "clear match", guildID, raidID, err)
	}
	if err != nil {
		return Raid{}, fmt.Errorf("guilds: clear match: read raid: %w", err)
	}
	if reportID == nil {
		return Raid{}, ErrRaidStatus
	}
	raid, err := releaseLog(ctx, tx, raidID, *reportID, by)
	if err != nil {
		return Raid{}, err
	}
	if err := tx.Commit(ctx); err != nil {
		return Raid{}, fmt.Errorf("guilds: clear match: commit: %w", err)
	}
	return raid, nil
}

func releaseLog(ctx context.Context, tx pgx.Tx, raidID int64, reportID string, by int64) (Raid, error) {
	raid, err := scanRaid(tx.QueryRow(ctx, `
		update guild_raids set status = coalesce(pre_done_status, 'planned'), pre_done_status = null,
			report_id = null, matched_by = null, matched_at = null, updated_at = now()
		where id = $1 returning `+raidColumns, raidID))
	if err != nil {
		return Raid{}, fmt.Errorf("guilds: clear match: restore: %w", err)
	}
	if _, err := tx.Exec(ctx, `
		insert into guild_raid_rejected_reports (raid_id, report_id, rejected_by) values ($1, $2, $3)
		on conflict (raid_id, report_id) do nothing`, raidID, reportID, by); err != nil {
		return Raid{}, fmt.Errorf("guilds: clear match: reject: %w", err)
	}
	if _, err := tx.Exec(ctx, `delete from guild_raid_attendance where raid_id = $1`, raidID); err != nil {
		return Raid{}, fmt.Errorf("guilds: clear match: attendance: %w", err)
	}
	if err := detachMatchLoot(ctx, tx, raidID); err != nil {
		return Raid{}, err
	}
	return raid, nil
}
