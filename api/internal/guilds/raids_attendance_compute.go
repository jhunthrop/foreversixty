// api/internal/guilds/raids_attendance_compute.go
//
// The attendance rule, with no database: a raid's sign-ups against the characters its log
// shows. raids_attendance.go reads the inputs and stores the rows.
package guilds

import "sort"

// Attendance outcomes (guild_raid_attendance.outcome). "no_show" is a JSON value only; no
// player-facing string says it.
const (
	OutcomePresent = "present"
	OutcomeLate    = "late"
	OutcomeAbsent  = "absent"
	OutcomeNoShow  = "no_show"
)

// lateGraceSeconds is how long after the log's first fight a signed-up raider may first
// appear before the night calls them late.
const lateGraceSeconds = 15 * 60

// AttendanceRow is one character's night.
type AttendanceRow struct {
	CharacterKey, Name, Class, Spec string
	Signed, Role                    *string
	Outcome                         string
	CameAs                          *string
	FirstSeenOffsetS                *int
	PullsPresent, PullsTotal        int
	Excused                         bool
}

// logPresence is one character as the log shows it. FirstSeenOffsetS is nil when the
// character is present with no fight rows to place them.
type logPresence struct {
	Name, Class, Spec string
	FirstSeenOffsetS  *int
	PullsPresent      int
}

// attendanceInput is everything the rule reads.
type attendanceInput struct {
	signups    []Signup
	roster     map[string]raidRosterChar
	present    map[string]logPresence
	pullsTotal int
}

func textPtr(s string) *string { return &s }

// computeAttendance returns one row per signed-up character and per roster character the log
// shows, ordered by name then key. Characters in the log who are not on the roster are left out.
func computeAttendance(in attendanceInput) []AttendanceRow {
	rows := map[string]AttendanceRow{}
	signed := map[string]Signup{}
	for _, sg := range in.signups {
		signed[sg.CharacterKey] = sg
		rows[sg.CharacterKey] = AttendanceRow{CharacterKey: sg.CharacterKey, Name: sg.Name, Class: sg.Class,
			Spec: sg.Spec, Signed: textPtr(sg.State), Role: textPtr(sg.Role)}
	}
	for key := range in.present {
		rc, onRoster := in.roster[key]
		if _, ok := rows[key]; !ok && onRoster {
			rows[key] = AttendanceRow{CharacterKey: key, Name: rc.Name, Class: rc.Class, Spec: rc.Spec}
		}
	}
	out := make([]AttendanceRow, 0, len(rows))
	for key, row := range rows {
		out = append(out, in.decide(row, signed[key]))
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Name != out[j].Name {
			return out[i].Name < out[j].Name
		}
		return out[i].CharacterKey < out[j].CharacterKey
	})
	return out
}

// decide fills a row's outcome and log figures.
func (in attendanceInput) decide(row AttendanceRow, sg Signup) AttendanceRow {
	row.PullsTotal = in.pullsTotal
	attending := row.Signed != nil && isAttending(*row.Signed)
	if p, ok := in.present[row.CharacterKey]; ok {
		return withPresence(row, p, attending)
	}
	if !attending {
		row.Outcome = OutcomeAbsent
		return row
	}
	if key, ok := in.sameAccountInLog(row.CharacterKey, sg); ok {
		p := in.present[key]
		row = withPresence(row, p, true)
		row.CameAs = textPtr(in.nameOf(key, p))
		return row
	}
	row.Outcome = OutcomeNoShow
	return row
}

func withPresence(row AttendanceRow, p logPresence, attending bool) AttendanceRow {
	row.Outcome = OutcomePresent
	row.FirstSeenOffsetS = p.FirstSeenOffsetS
	row.PullsPresent = p.PullsPresent
	if attending && p.FirstSeenOffsetS != nil && *p.FirstSeenOffsetS > lateGraceSeconds {
		row.Outcome = OutcomeLate
	}
	return row
}

// sameAccountInLog finds another roster character of key's account that the log shows, the
// one with the most pulls (then the lowest key).
func (in attendanceInput) sameAccountInLog(key string, sg Signup) (string, bool) {
	owner := sg.SignedBy
	if rc, ok := in.roster[key]; ok {
		owner = rc.userID
	}
	best, found := "", false
	for other, rc := range in.roster {
		p, here := in.present[other]
		if other == key || !here || rc.userID != owner {
			continue
		}
		if !found || p.PullsPresent > in.present[best].PullsPresent ||
			(p.PullsPresent == in.present[best].PullsPresent && other < best) {
			best, found = other, true
		}
	}
	return best, found
}

func (in attendanceInput) nameOf(key string, p logPresence) string {
	if rc, ok := in.roster[key]; ok && rc.Name != "" {
		return rc.Name
	}
	return p.Name
}
