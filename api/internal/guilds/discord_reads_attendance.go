// api/internal/guilds/discord_reads_attendance.go
//
// The read behind the Discord bot's /attendance (spec 2026-10-27-discord-bot §3.3.3): a
// character's roster chip, read as the Roster tab reads it (fillAttendance), and its last
// logged raids. A member reads their own characters; an officer may name anyone, by the same
// rank rule the raid page's attendance rows follow (seesEveryAttendanceRow: reads are never
// plan-gated).
package guilds

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"
)

const (
	// maxAttendanceRecords is how many own characters one answer lists.
	maxAttendanceRecords = 5
	// attendanceLastRaids is how many of the character's last logged raids it lists.
	attendanceLastRaids = 5
	// chipJoin joins the chip's two windows, as the web's SEP.
	chipJoin = " · "
)

// AttendanceLine is one logged raid of a character: when it started and its outcome
// (OutcomePresent, OutcomeLate, OutcomeNoShow); CameAs names the other character the account
// came as, "" when none.
type AttendanceLine struct {
	StartsAt time.Time
	Outcome  string
	Excused  bool
	CameAs   string
}

// AttendanceRecord is one character's attendance. Chip is the Roster tab's chip ("Came 7 of 8 ·
// 26 of 30 raids", report mode "5 of 8 nights"), "" when the character has no night yet.
type AttendanceRecord struct {
	Name, Class, Chip string
	Last              []AttendanceLine
}

// AttendanceFor is viewerID's own verified characters' records (name ""; main first, at most
// five), or the named character's. ErrNotMember for a viewer with no verified character,
// ErrOfficerOnly for a member naming a character not their own, ErrNoSuchMember for a name no
// verified character carries.
func (s *Service) AttendanceFor(ctx context.Context, guildID, viewerID int64, name string) ([]AttendanceRecord, error) {
	own, err := s.ownVerifiedKeys(ctx, guildID, viewerID)
	if err != nil {
		return nil, err
	}
	if len(own) == 0 {
		return nil, ErrNotMember
	}
	roster, source, err := s.attendanceRoster(ctx, guildID)
	if err != nil {
		return nil, err
	}
	keys, err := s.attendanceKeys(ctx, guildID, viewerID, own, roster, strings.TrimSpace(name))
	if err != nil {
		return nil, err
	}
	out := make([]AttendanceRecord, 0, len(keys))
	for _, key := range keys {
		row := roster[key]
		rec := AttendanceRecord{Name: row.Name, Chip: attendanceChip(row, source)}
		if row.Class != nil {
			rec.Class = strings.ToLower(*row.Class)
		}
		if rec.Last, err = s.Store.lastAttendance(ctx, guildID, key, attendanceLastRaids); err != nil {
			return nil, err
		}
		out = append(out, rec)
	}
	return out, nil
}

// ownVerifiedKeys is the viewer's verified characters in the guild, main first.
func (s *Service) ownVerifiedKeys(ctx context.Context, guildID, viewerID int64) ([]string, error) {
	cands, err := s.SignupCandidates(ctx, guildID, viewerID)
	if err != nil {
		return nil, fmt.Errorf("guilds: attendance candidates: %w", err)
	}
	var keys []string
	for _, c := range cands {
		if c.Verified {
			keys = append(keys, c.CharacterKey)
		}
	}
	return keys, nil
}

// attendanceRoster is the guild's verified characters by key, their attendance filled by the
// Roster tab's own read (fillAttendance) in the guild's attendance mode.
func (s *Service) attendanceRoster(ctx context.Context, guildID int64) (map[string]RosterRow, string, error) {
	chars, err := s.Store.raidRoster(ctx, guildID, "")
	if err != nil {
		return nil, "", fmt.Errorf("guilds: attendance roster: %w", err)
	}
	source, err := s.Store.attendanceSource(ctx, guildID)
	if err != nil {
		return nil, "", err
	}
	rows, keys := []RosterRow{}, []string{}
	for _, c := range chars {
		if c.verified {
			class := c.Class
			rows = append(rows, RosterRow{CharacterKey: c.CharacterKey, Name: c.Name, Class: &class, Verified: true})
			keys = append(keys, c.CharacterKey)
		}
	}
	if err := s.Store.fillAttendance(ctx, guildID, source, rows, keys); err != nil {
		return nil, "", err
	}
	out := make(map[string]RosterRow, len(rows))
	for _, r := range rows {
		out[r.CharacterKey] = r
	}
	return out, source, nil
}

// attendanceKeys is which characters the answer lists, applying the officer rule to a name.
func (s *Service) attendanceKeys(ctx context.Context, guildID, viewerID int64, own []string, roster map[string]RosterRow, name string) ([]string, error) {
	if name == "" {
		kept := []string{}
		for _, k := range own {
			if _, ok := roster[k]; ok && len(kept) < maxAttendanceRecords {
				kept = append(kept, k)
			}
		}
		return kept, nil
	}
	for _, k := range own {
		if r, ok := roster[k]; ok && strings.EqualFold(r.Name, name) {
			return []string{k}, nil
		}
	}
	officer, err := s.seesEveryAttendanceRow(ctx, guildID, viewerID)
	if err != nil {
		return nil, err
	}
	if !officer {
		return nil, ErrOfficerOnly
	}
	keys := make([]string, 0, len(roster))
	for k := range roster {
		keys = append(keys, k)
	}
	sort.Strings(keys) // two characters of one name (two realms) resolve the same way every time
	for _, k := range keys {
		if strings.EqualFold(roster[k].Name, name) {
			return []string{k}, nil
		}
	}
	return nil, ErrNoSuchMember
}

// attendanceChip is the Roster tab's chip verbatim (web/src/lib/guild/roster-facts.ts
// attendanceChip): report mode nightsLabel ("5 of 8 nights"), raid mode rosterChip ("Came 7 of
// 8 · 26 of 30 raids"); "" with no night yet.
func attendanceChip(row RosterRow, source string) string {
	last8 := row.Attendance
	if last8.Nights == 0 {
		return ""
	}
	if source != AttendanceSourceRaids {
		return fmt.Sprintf("%d of %d night%s", last8.Present, last8.Nights, plural(last8.Nights))
	}
	head := fmt.Sprintf("Came %d of %d", last8.Present, last8.Nights)
	if row.Attendance30 == nil || row.Attendance30.Nights <= last8.Nights {
		return head + " raids"
	}
	return fmt.Sprintf("%s%s%d of %d raids", head, chipJoin, row.Attendance30.Present, row.Attendance30.Nights)
}

// lastAttendance is key's rows over the guild's logged raids that count a night (signed In or
// Late, or came; not as another character's came_as, the chip's rule), newest first, at most limit.
func (s *Store) lastAttendance(ctx context.Context, guildID int64, key string, limit int) ([]AttendanceLine, error) {
	rows, err := s.Pool.Query(ctx, `
		select g.starts_at, a.outcome, a.excused, coalesce(a.came_as, '')
		from guild_raid_attendance a join guild_raids g on g.id = a.raid_id
		where a.raid_id in (`+loggedRaids+`)
		  and a.character_key = $3
		  and (a.signed in ('in', 'late') or a.outcome in ('present', 'late'))
		  and not exists (select 1 from guild_raid_attendance m
		                  where m.raid_id = a.raid_id and m.came_as_key = a.character_key)
		order by g.starts_at desc, g.id desc limit $4`, guildID, attendanceWindowLong, key, limit)
	if err != nil {
		return nil, fmt.Errorf("guilds: last attendance: %w", err)
	}
	defer rows.Close()
	out := []AttendanceLine{}
	for rows.Next() {
		var l AttendanceLine
		if err := rows.Scan(&l.StartsAt, &l.Outcome, &l.Excused, &l.CameAs); err != nil {
			return nil, fmt.Errorf("guilds: last attendance: %w", err)
		}
		l.StartsAt = l.StartsAt.UTC()
		out = append(out, l)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: last attendance: %w", err)
	}
	return out, nil
}

// AttendanceNames is the guild's verified character names containing typed, for an officer's
// /attendance name; empty for anyone else (another member's name is the officers' to look up).
func (s *Service) AttendanceNames(ctx context.Context, guildID, viewerID int64, typed string, limit int) ([]string, error) {
	officer, err := s.seesEveryAttendanceRow(ctx, guildID, viewerID)
	if err != nil || !officer {
		return []string{}, err
	}
	rows, err := s.Store.Pool.Query(ctx, `
		select c.name from guild_characters gc join characters c on c.key = gc.character_key
		where gc.guild_id = $1 and gc.verified_at is not null and c.name <> ''
		  and strpos(lower(c.name), lower($2)) > 0
		order by lower(c.name) limit $3`, guildID, strings.TrimSpace(typed), limit)
	if err != nil {
		return nil, fmt.Errorf("guilds: attendance names: %w", err)
	}
	defer rows.Close()
	out := []string{}
	for rows.Next() {
		var n string
		if err := rows.Scan(&n); err != nil {
			return nil, fmt.Errorf("guilds: attendance names: %w", err)
		}
		out = append(out, n)
	}
	return out, rows.Err()
}
