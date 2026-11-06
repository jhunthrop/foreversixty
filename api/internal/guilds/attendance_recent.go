// api/internal/guilds/attendance_recent.go
//
// GET /v1/guilds/{id}/attendance/recent?n=4[&all=1] (spec 2026-11-05-web-layout §7.1
// "Attendance, last 4 raids"): an officer's grid of the guild's last n logged raids, one row
// per verified account, read from the stored attendance rows (raids_attendance.go). A night
// counts as expected exactly as the roster's attendance counts it (signed In or Late, not
// excused), so "slipping" agrees with the roster's "3 of 4". By default only the slipping
// rows and the rows with a Bench night are listed; all=1 lists everyone.
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// Attendance grid cells.
const (
	CellPresent = "present"
	CellAbsent  = "absent"
	CellBench   = "bench"
	CellNone    = "none"
)

const (
	recentNightsDefault = 4
	recentNightsMax     = 12
	// slippingMisses is how many expected nights missed make a raider slipping.
	slippingMisses = 2
)

// RecentRaid is one column of the grid.
type RecentRaid struct {
	ID       int64     `json:"id"`
	StartsAt time.Time `json:"starts_at"`
	ZoneName string    `json:"zone_name"`
}

// RecentAttendanceRow is one account's nights, oldest raid first like the columns.
type RecentAttendanceRow struct {
	Name       string   `json:"name"`
	Class      string   `json:"class"`
	AccountKey string   `json:"account_key"`
	Cells      []string `json:"cells"`
	Slipping   bool     `json:"slipping"`
}

// RecentAttendanceView is the grid.
type RecentAttendanceView struct {
	Raids []RecentRaid          `json:"raids"`
	Rows  []RecentAttendanceRow `json:"rows"`
}

// accountNight is what one account's attendance rows say of one raid.
type accountNight struct {
	came, bench, expected bool
}

func (s *Service) recentAttendance(w http.ResponseWriter, r *http.Request) {
	guildID, ok := s.officerRead(w, r, "attendance_recent")
	if !ok {
		return
	}
	n := recentNightsDefault
	if q := r.URL.Query().Get("n"); q != "" {
		v, err := strconv.Atoi(q)
		if err != nil || v < 1 || v > recentNightsMax {
			writeInvalid(w, r, map[string]string{"n": fmt.Sprintf("Pick 1 to %d raids.", recentNightsMax)})
			return
		}
		n = v
	}
	view, err := s.Store.RecentAttendance(r.Context(), guildID, n, r.URL.Query().Get("all") == "1")
	if err != nil {
		s.fail(w, r, "attendance_recent", err, "could not load attendance just now")
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, view)
}

// officerRead gates an officer-only read by rank (no plan needed: a read is never locked): 404
// for a bad id, 403 for anyone but a verified officer or leader.
func (s *Service) officerRead(w http.ResponseWriter, r *http.Request, op string) (int64, bool) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return 0, false
	}
	officer, err := s.seesEveryAttendanceRow(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID)
	if err != nil {
		s.fail(w, r, op, err, "could not load that just now")
		return 0, false
	}
	if !officer {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "you must be a verified officer of this guild", nil)
		return 0, false
	}
	return guildID, true
}

// RecentAttendance is the grid over the guild's last n logged raids.
func (s *Store) RecentAttendance(ctx context.Context, guildID int64, n int, all bool) (RecentAttendanceView, error) {
	raids, err := s.recentLoggedRaids(ctx, guildID, n)
	if err != nil {
		return RecentAttendanceView{}, err
	}
	view := RecentAttendanceView{Raids: raids, Rows: []RecentAttendanceRow{}}
	if len(raids) == 0 {
		return view, nil
	}
	nights, err := s.accountNights(ctx, guildID, raids)
	if err != nil {
		return RecentAttendanceView{}, err
	}
	joined, err := s.accountJoins(ctx, guildID)
	if err != nil {
		return RecentAttendanceView{}, err
	}
	people, err := s.accountFaces(ctx, guildID)
	if err != nil {
		return RecentAttendanceView{}, err
	}
	for uid, face := range people {
		row := gridRow(raids, nights[uid], joined[uid])
		row.Name, row.Class, row.AccountKey = face.Name, face.Class, accountKeyFor(uid)
		if all || row.Slipping || hasCell(row.Cells, CellBench) {
			view.Rows = append(view.Rows, row)
		}
	}
	sort.Slice(view.Rows, func(i, j int) bool { return view.Rows[i].Name < view.Rows[j].Name })
	return view, nil
}

// gridRow is one account's cells: present when any character of it came, bench when it was
// benched, none before it joined, else absent; slipping when it missed slippingMisses or more
// nights it was expected for.
func gridRow(raids []RecentRaid, nights map[int64]accountNight, joined *time.Time) RecentAttendanceRow {
	row := RecentAttendanceRow{Cells: make([]string, len(raids))}
	missed := 0
	for i, raid := range raids {
		night := nights[raid.ID]
		switch {
		case night.came:
			row.Cells[i] = CellPresent
		case night.bench:
			row.Cells[i] = CellBench
		case joined != nil && raid.StartsAt.Before(*joined):
			row.Cells[i] = CellNone
		default:
			row.Cells[i] = CellAbsent
			if night.expected {
				missed++
			}
		}
	}
	row.Slipping = missed >= slippingMisses
	return row
}

func hasCell(cells []string, want string) bool {
	for _, c := range cells {
		if c == want {
			return true
		}
	}
	return false
}

// recentLoggedRaids is the guild's last n logged raids (loggedRaids), oldest first.
func (s *Store) recentLoggedRaids(ctx context.Context, guildID int64, n int) ([]RecentRaid, error) {
	rows, err := s.Pool.Query(ctx, `select g.id, g.starts_at, g.zone from guild_raids g
		where g.id in (`+loggedRaids+`) order by g.starts_at, g.id`, guildID, n)
	if err != nil {
		return nil, fmt.Errorf("guilds: recent raids: %w", err)
	}
	defer rows.Close()
	out := []RecentRaid{}
	for rows.Next() {
		var r RecentRaid
		var zone string
		if err := rows.Scan(&r.ID, &r.StartsAt, &zone); err != nil {
			return nil, fmt.Errorf("guilds: recent raids: %w", err)
		}
		r.StartsAt, r.ZoneName = r.StartsAt.UTC(), zoneOf(zone).Name
		out = append(out, r)
	}
	return out, rows.Err()
}

// accountNights reads, per account and raid, what its characters' attendance rows say.
func (s *Store) accountNights(ctx context.Context, guildID int64, raids []RecentRaid) (map[int64]map[int64]accountNight, error) {
	ids := make([]int64, len(raids))
	for i, r := range raids {
		ids[i] = r.ID
	}
	rows, err := s.Pool.Query(ctx, `
		select gc.user_id, a.raid_id,
		       bool_or(a.outcome in ('present', 'late')),
		       coalesce(bool_or(a.signed = 'bench'), false),
		       coalesce(bool_or(a.signed in ('in', 'late') and not a.excused), false)
		from guild_raid_attendance a
		join guild_characters gc on gc.guild_id = $1 and gc.character_key = a.character_key
		where a.raid_id = any($2)
		group by gc.user_id, a.raid_id`, guildID, ids)
	if err != nil {
		return nil, fmt.Errorf("guilds: account nights: %w", err)
	}
	defer rows.Close()
	out := map[int64]map[int64]accountNight{}
	for rows.Next() {
		var uid, raidID int64
		var n accountNight
		if err := rows.Scan(&uid, &raidID, &n.came, &n.bench, &n.expected); err != nil {
			return nil, fmt.Errorf("guilds: account nights: %w", err)
		}
		if out[uid] == nil {
			out[uid] = map[int64]accountNight{}
		}
		out[uid][raidID] = n
	}
	return out, rows.Err()
}

// accountJoins is when each verified account joined: guild_members.joined_at, else (an account
// from before that column) the start of its first raid sign-up in the guild; absent when
// neither is known, which reads as a member for every night.
func (s *Store) accountJoins(ctx context.Context, guildID int64) (map[int64]*time.Time, error) {
	rows, err := s.Pool.Query(ctx, `
		select gm.user_id, coalesce(gm.joined_at, (
		  select min(r.starts_at) from guild_raid_signups sg join guild_raids r on r.id = sg.raid_id
		  where r.guild_id = gm.guild_id and sg.signed_by = gm.user_id))
		from guild_members gm where gm.guild_id = $1 and gm.verified_at is not null`, guildID)
	if err != nil {
		return nil, fmt.Errorf("guilds: account joins: %w", err)
	}
	defer rows.Close()
	out := map[int64]*time.Time{}
	for rows.Next() {
		var uid int64
		var at *time.Time
		if err := rows.Scan(&uid, &at); err != nil {
			return nil, fmt.Errorf("guilds: account joins: %w", err)
		}
		if at != nil {
			out[uid] = at
		}
	}
	return out, rows.Err()
}

// accountFaces is the character each verified account is shown as: its main when verified on
// the roster, else its first verified character by name.
func (s *Store) accountFaces(ctx context.Context, guildID int64) (map[int64]RaidCharacter, error) {
	roster, err := s.raidRoster(ctx, guildID, "")
	if err != nil {
		return nil, err
	}
	mains, err := s.mainKeys(ctx, guildID)
	if err != nil {
		return nil, err
	}
	out := map[int64]RaidCharacter{}
	for _, c := range roster {
		if !c.verified {
			continue
		}
		cur, seen := out[c.userID]
		isMain := mains[c.userID] == c.CharacterKey
		if !seen || isMain || (mains[c.userID] != cur.CharacterKey && c.Name < cur.Name) {
			out[c.userID] = c.RaidCharacter
		}
	}
	return out, nil
}

// mainKeys is each guild account's main character key, where it has one.
func (s *Store) mainKeys(ctx context.Context, guildID int64) (map[int64]string, error) {
	rows, err := s.Pool.Query(ctx, `select u.id, u.main_character_key from users u
		join guild_members gm on gm.user_id = u.id and gm.guild_id = $1
		where u.main_character_key is not null`, guildID)
	if err != nil {
		return nil, fmt.Errorf("guilds: main keys: %w", err)
	}
	defer rows.Close()
	out := map[int64]string{}
	for rows.Next() {
		var uid int64
		var key string
		if err := rows.Scan(&uid, &key); err != nil {
			return nil, fmt.Errorf("guilds: main keys: %w", err)
		}
		out[uid] = key
	}
	return out, rows.Err()
}
