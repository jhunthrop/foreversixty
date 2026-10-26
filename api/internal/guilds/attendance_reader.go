// api/internal/guilds/attendance_reader.go
//
// Attendance read from matched raids. A guild switches to this reading once it has one done
// raid with attendance rows; until then the roster keeps counting reports (roster_extended.go).
// In raid mode nights are per character: the raids the character signed in/late for or
// attended, among the guild's last N logged raids. An excused row is neither a night nor
// present.
package guilds

import (
	"context"
	"fmt"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

const (
	// AttendanceSourceRaids and AttendanceSourceReports are HomeView.AttendanceSource's values.
	AttendanceSourceRaids   = "raids"
	AttendanceSourceReports = "reports"

	// attendanceWindowLong is the second window of raid mode (attendance_30).
	attendanceWindowLong = 30
)

// loggedRaids is the guild's done raids that carry attendance rows, newest first. A raid whose
// log is not evidence (private or flagged) is left out: its rows count for nobody until the log
// is public again.
var loggedRaids = `select r.id from guild_raids r
	join reports rp on rp.id = r.report_id and ` + presence.EvidenceReport("rp") + `
	where r.guild_id = $1 and r.status = 'done'
	  and exists (select 1 from guild_raid_attendance a where a.raid_id = r.id)
	order by r.starts_at desc, r.id desc
	limit $2`

// attendanceSource is "raids" once the guild has a logged raid, else "reports".
func (s *Store) attendanceSource(ctx context.Context, guildID int64) (string, error) {
	var logged bool
	if err := s.Pool.QueryRow(ctx, `select exists (`+loggedRaids+`)`, guildID, 1).Scan(&logged); err != nil {
		return "", fmt.Errorf("guilds: attendance source: %w", err)
	}
	if logged {
		return AttendanceSourceRaids, nil
	}
	return AttendanceSourceReports, nil
}

// raidAttendance is, per key in keys, the raids it attended over the raids it counted a night
// for, within the guild's last window logged raids. A key with no night is absent from the map.
func (s *Store) raidAttendance(ctx context.Context, guildID int64, keys []string, window int) (map[string]Attendance, error) {
	out := map[string]Attendance{}
	if len(keys) == 0 {
		return out, nil
	}
	rows, err := s.Pool.Query(ctx, `
		select a.character_key,
		       count(*) filter (where a.outcome in ('present', 'late')),
		       count(*)
		from guild_raid_attendance a
		where a.raid_id in (`+loggedRaids+`)
		  and a.character_key = any($3) and not a.excused
		  and (a.signed in ('in', 'late') or a.outcome in ('present', 'late'))
		group by a.character_key`, guildID, window, keys)
	if err != nil {
		return nil, fmt.Errorf("guilds: raid attendance: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var key string
		var a Attendance
		if err := rows.Scan(&key, &a.Present, &a.Nights); err != nil {
			return nil, fmt.Errorf("guilds: raid attendance: %w", err)
		}
		out[key] = a
	}
	return out, rows.Err()
}

// fillRaidAttendance sets Attendance (last 8) and Attendance30 on every row of out.
func (s *Store) fillRaidAttendance(ctx context.Context, guildID int64, out []RosterRow, keys []string) error {
	short, err := s.raidAttendance(ctx, guildID, keys, attendanceWindow)
	if err != nil {
		return err
	}
	long, err := s.raidAttendance(ctx, guildID, keys, attendanceWindowLong)
	if err != nil {
		return err
	}
	for i := range out {
		key := out[i].CharacterKey
		out[i].Attendance = short[key]
		l := long[key]
		out[i].Attendance30 = &l
	}
	return nil
}

// CharacterAttendance answers /v1/me (auth.GuildAttendanceReader): the last 8 logged raids'
// attendance per key, empty while the guild still counts reports.
func (s *Store) CharacterAttendance(ctx context.Context, guildID int64, keys []string) (map[string]auth.Attendance, error) {
	source, err := s.attendanceSource(ctx, guildID)
	if err != nil {
		return nil, err
	}
	out := map[string]auth.Attendance{}
	if source != AttendanceSourceRaids {
		return out, nil
	}
	got, err := s.raidAttendance(ctx, guildID, keys, attendanceWindow)
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		a := got[key]
		out[key] = auth.Attendance{Present: a.Present, Nights: a.Nights}
	}
	return out, nil
}
