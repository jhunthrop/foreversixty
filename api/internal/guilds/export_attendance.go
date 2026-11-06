// api/internal/guilds/export_attendance.go
//
// The attendance CSV reads the nights the Raids tab counts: a planned raid's attendance rows
// where a matched raid has them, and for every other logged night the characters its log shows
// present, the same rows behind the tab's "N of N present".
package guilds

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

// exportOutcomePresent is a logged night's row: the log shows the character, no plan to judge.
const exportOutcomePresent = "present"

// attendanceExportRow is one CSV row with what it sorts by.
type attendanceExportRow struct {
	at        time.Time
	name, key string
	cells     []string
}

// loggedNight is a guild report that no planned raid's attendance rows cover.
type loggedNight struct {
	id, label string
	at        time.Time
}

func (s *Store) streamAttendance(ctx context.Context, guildID int64, emit func([]string) error) error {
	loc, err := s.exportLocation(ctx, guildID)
	if err != nil {
		return err
	}
	planned, err := s.plannedAttendanceRows(ctx, guildID, loc)
	if err != nil {
		return err
	}
	logged, err := s.loggedNightRows(ctx, guildID, loc)
	if err != nil {
		return err
	}
	all := make([]attendanceExportRow, 0, len(planned)+len(logged))
	all = append(append(all, planned...), logged...)
	sort.SliceStable(all, func(i, j int) bool { return attendanceRowBefore(all[i], all[j]) })
	for _, r := range all {
		if err := emit(r.cells); err != nil {
			return err
		}
	}
	return nil
}

// attendanceRowBefore orders the file by night, then name, then character key.
func attendanceRowBefore(a, b attendanceExportRow) bool {
	if !a.at.Equal(b.at) {
		return a.at.Before(b.at)
	}
	if a.name != b.name {
		return a.name < b.name
	}
	return a.key < b.key
}

// plannedAttendanceRows is every planned raid's attendance rows.
func (s *Store) plannedAttendanceRows(ctx context.Context, guildID int64, loc *time.Location) ([]attendanceExportRow, error) {
	rows, err := s.Pool.Query(ctx, `
		select r.starts_at, r.title, a.name, a.character_key, a.class, coalesce(a.signed, ''), a.outcome, a.excused,
		       a.pulls_present, a.pulls_total, coalesce(o.status, '')
		from guild_raid_attendance a join guild_raids r on r.id = a.raid_id
		left join guild_raid_attendance_overrides o on o.raid_id = a.raid_id and o.character_key = a.character_key
		where r.guild_id = $1`, guildID)
	if err != nil {
		return nil, fmt.Errorf("guilds: export attendance: %w", err)
	}
	defer rows.Close()
	var out []attendanceExportRow
	for rows.Next() {
		var (
			startsAt                              time.Time
			title, name, key, class, signed, kind string
			fix                                   string
			excused                               bool
			present, total                        int
		)
		if err := rows.Scan(&startsAt, &title, &name, &key, &class, &signed, &kind, &excused, &present, &total, &fix); err != nil {
			return nil, fmt.Errorf("guilds: export attendance: %w", err)
		}
		out = append(out, attendanceExportRow{at: startsAt, name: name, key: key,
			cells: exportAttendanceCells(startsAt.In(loc), title, name, class, signed, kind, excused, present, total, fix)})
	}
	return out, rows.Err()
}

// loggedNightRows is, for each logged night no planned raid's rows cover, one present row per
// character its log shows (the presence rules: a private or flagged log shows nobody).
func (s *Store) loggedNightRows(ctx context.Context, guildID int64, loc *time.Location) ([]attendanceExportRow, error) {
	nights, err := s.loggedNights(ctx, guildID)
	if err != nil {
		return nil, err
	}
	var out []attendanceExportRow
	for _, n := range nights {
		present, err := logPresentKeys(ctx, s.Pool, n.id)
		if err != nil {
			return nil, err
		}
		if err := fillLogFigures(ctx, s.Pool, n.id, present); err != nil {
			return nil, err
		}
		total, err := encounterPulls(ctx, s.Pool, n.id)
		if err != nil {
			return nil, err
		}
		for key, p := range present {
			name := p.Name
			if name == "" {
				name = key
			}
			out = append(out, attendanceExportRow{at: n.at, name: name, key: key,
				cells: exportAttendanceCells(n.at.In(loc), n.label, name, p.Class, "", exportOutcomePresent, false, p.PullsPresent, total, "")})
		}
	}
	return out, nil
}

// loggedNights is the guild's evidence reports (the ones the Raids tab lists anyone present
// for) that no planned raid with attendance rows holds, each at the night it was played.
func (s *Store) loggedNights(ctx context.Context, guildID int64) ([]loggedNight, error) {
	rows, err := s.Pool.Query(ctx, `
		select r.id, coalesce(nullif(r.zone, ''), r.title), `+logAnchorSQL("r")+`
		from reports r
		where r.guild_id = $1 and `+presence.EvidenceReport("r")+`
		  and not exists (select 1 from guild_raids gr
		                  where gr.report_id = r.id and gr.guild_id = $1
		                    and exists (select 1 from guild_raid_attendance a where a.raid_id = gr.id))`, guildID)
	if err != nil {
		return nil, fmt.Errorf("guilds: export logged nights: %w", err)
	}
	defer rows.Close()
	var out []loggedNight
	for rows.Next() {
		var n loggedNight
		if err := rows.Scan(&n.id, &n.label, &n.at); err != nil {
			return nil, fmt.Errorf("guilds: export logged nights: %w", err)
		}
		out = append(out, n)
	}
	return out, rows.Err()
}

// exportAttendanceCells is one attendance CSV row, formula-safe; at is already in the guild's
// zone (exportLocation). outcome and excused are the effective values; fix is the officer's
// manual fix ("" for none).
func exportAttendanceCells(at time.Time, raid, name, class, signed, outcome string, excused bool, present, total int, fix string) []string {
	return safeRow(at.Format(exportDateLayout), raid, name, class, signed, outcome,
		strconv.FormatBool(excused), strconv.Itoa(present), strconv.Itoa(total), fix)
}
