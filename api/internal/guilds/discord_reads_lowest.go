// api/internal/guilds/discord_reads_lowest.go
//
// The read behind an officer's guided /attendance (spec 2026-10-28-discord-guided-cards §G.4.5):
// the guild's members who came least over the roster chip's last-8 window.
package guilds

import (
	"context"
	"sort"
	"strings"
)

// LowestAttendance is the guild's verified characters that have a roster chip (at least one
// night), lowest first by how many of the last-8 nights they came to, then by name; at most
// limit. Officers only, by seesEveryAttendanceRow's rule: anyone else gets ErrOfficerOnly.
// Records carry no raid lines.
func (s *Service) LowestAttendance(ctx context.Context, guildID, viewerID int64, limit int) ([]AttendanceRecord, error) {
	officer, err := s.seesEveryAttendanceRow(ctx, guildID, viewerID)
	if err != nil {
		return nil, err
	}
	if !officer {
		return nil, ErrOfficerOnly
	}
	roster, source, err := s.attendanceRoster(ctx, guildID)
	if err != nil {
		return nil, err
	}
	return lowestOf(roster, source, limit), nil
}

// lowestOf is the chipped rows of roster, fewest nights come first, then by name and key.
func lowestOf(roster map[string]RosterRow, source string, limit int) []AttendanceRecord {
	rows := make([]RosterRow, 0, len(roster))
	for _, r := range roster {
		if attendanceChip(r, source) != "" {
			rows = append(rows, r)
		}
	}
	sort.Slice(rows, func(i, j int) bool {
		a, b := rows[i], rows[j]
		if a.Attendance.Present != b.Attendance.Present {
			return a.Attendance.Present < b.Attendance.Present
		}
		if an, bn := strings.ToLower(a.Name), strings.ToLower(b.Name); an != bn {
			return an < bn
		}
		return a.CharacterKey < b.CharacterKey
	})
	rows = rows[:min(max(limit, 0), len(rows))]
	out := make([]AttendanceRecord, 0, len(rows))
	for _, r := range rows {
		rec := AttendanceRecord{Name: r.Name, Chip: attendanceChip(r, source)}
		if r.Class != nil {
			rec.Class = strings.ToLower(*r.Class)
		}
		out = append(out, rec)
	}
	return out
}
