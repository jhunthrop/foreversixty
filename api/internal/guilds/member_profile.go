// api/internal/guilds/member_profile.go
//
// One member's profile (design/specs/2026-11-05-web-layout.md §7.2 Member profile): the
// roster row the home already builds for the character, plus attendance over the last 4 and
// last 30 nights, the character's loot this tier, the readiness board's row (its gear gap in the
// BiS page's unit) and, for officers only, the officer notes. Officer or self.
package guilds

import (
	"context"
	"fmt"
	"time"
)

// profileShortWindow is the profile's short attendance window, the mock's "Last 4 raids".
const profileShortWindow = 4

// ProfileNight is one of the last profileShortWindow logged raids as the member's attendance
// row reads: Outcome present | late | absent | no_show, null when the raid has no row for them.
type ProfileNight struct {
	RaidID   int64     `json:"raid_id"`
	StartsAt time.Time `json:"starts_at"`
	Outcome  *string   `json:"outcome"`
	Excused  bool      `json:"excused"`
}

// ProfileAttendance is the attendance section. Source is the guild's attendance source (raids
// once a raid is logged, else reports); Nights is set in raids mode only, Last30 likewise.
type ProfileAttendance struct {
	Source string         `json:"source"`
	Last4  Attendance     `json:"last_4"`
	Last30 *Attendance    `json:"last_30"`
	Nights []ProfileNight `json:"nights"`
}

// ProfileLootItem is one award to the character this tier.
type ProfileLootItem struct {
	ItemID    int       `json:"item_id"`
	Name      *string   `json:"name"`
	Icon      *string   `json:"icon"`
	Quality   *int      `json:"quality"`
	RaidID    *int64    `json:"raid_id"`
	AwardedAt time.Time `json:"awarded_at"`
}

// ProfileLoot is the loot this tier section: every award row counts, as the loot council's own.
type ProfileLoot struct {
	Count int               `json:"count"`
	Items []ProfileLootItem `json:"items"`
}

// MemberProfile is GET /v1/guilds/{id}/members/{account}/profile. Member is null when none of
// the account's characters is on the roster with an export; then the sections are empty.
// Notes is present for an officer only.
type MemberProfile struct {
	AccountKey string             `json:"account_key"`
	Member     *RosterRow         `json:"member"`
	Characters []ProfileCharacter `json:"characters"`
	Attendance *ProfileAttendance `json:"attendance"`
	Loot       ProfileLoot        `json:"loot"`
	Readiness  *ReadinessRow      `json:"readiness"`
	Notes      *[]MemberNote      `json:"notes,omitempty"`
}

// ProfileCharacter is one of the account's roster characters, to switch the profile between.
type ProfileCharacter struct {
	CharacterKey string  `json:"character_key"`
	Name         string  `json:"name"`
	Class        *string `json:"class"`
	IsMain       bool    `json:"is_main"`
}

// MemberProfile assembles subjectID's profile in guildID on characterKey ("" picks the main,
// else the first roster row).
func (s *Store) MemberProfile(ctx context.Context, guildID, subjectID int64, characterKey string) (MemberProfile, error) {
	out := MemberProfile{AccountKey: accountKeyFor(subjectID), Characters: []ProfileCharacter{},
		Loot: ProfileLoot{Items: []ProfileLootItem{}}}
	roster, median, err := s.profileRoster(ctx, guildID)
	if err != nil {
		return out, err
	}
	row, chars := pickProfileRow(roster, subjectID, characterKey)
	out.Characters = chars
	if row == nil {
		return out, nil
	}
	out.Member = row
	if out.Attendance, err = s.profileAttendance(ctx, guildID, *row); err != nil {
		return out, err
	}
	if out.Loot, err = s.profileLoot(ctx, guildID, awardKeyOf(*row)); err != nil {
		return out, err
	}
	if row.Verified {
		rr := readinessRowOf(*row, s.readinessOf(*row, median))
		out.Readiness = &rr
	}
	return out, nil
}

// profileRoster is the readiness board's roster plus the invited accounts' picked characters,
// and the board's median item level.
func (s *Store) profileRoster(ctx context.Context, guildID int64) ([]RosterRow, *int, error) {
	roster, err := s.readinessRoster(ctx, guildID)
	if err != nil {
		return nil, nil, err
	}
	median := medianItemLevel(roster)
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return nil, nil, err
	}
	invited, err := s.invitedRosterRows(ctx, guildID, g)
	if err != nil {
		return nil, nil, err
	}
	return append(roster, invited...), median, nil
}

// pickProfileRow is subjectID's row for characterKey (else its main, else its first) and the
// account's characters on the roster.
func pickProfileRow(roster []RosterRow, subjectID int64, characterKey string) (*RosterRow, []ProfileCharacter) {
	var picked *RosterRow
	chars := []ProfileCharacter{}
	for i := range roster {
		row := roster[i]
		if row.ownerUserID != subjectID {
			continue
		}
		chars = append(chars, ProfileCharacter{CharacterKey: row.CharacterKey, Name: row.Name, Class: row.Class, IsMain: row.IsMain})
		switch {
		case characterKey != "":
			if row.CharacterKey == characterKey {
				picked = &roster[i]
			}
		case picked == nil, row.IsMain && !picked.IsMain:
			picked = &roster[i]
		}
	}
	return picked, chars
}

// profileAttendance reads the guild's attendance source for the row's character: in raids mode
// the last 4 and last 30 logged raids and the last 4 nights one by one, in reports mode the last
// 4 reports.
func (s *Store) profileAttendance(ctx context.Context, guildID int64, row RosterRow) (*ProfileAttendance, error) {
	source, err := s.attendanceSource(ctx, guildID)
	if err != nil {
		return nil, err
	}
	out := &ProfileAttendance{Source: source, Nights: []ProfileNight{}}
	keys := []string{row.CharacterKey}
	if source != AttendanceSourceRaids {
		nights, present, err := s.attendance(ctx, guildID, keys, profileShortWindow)
		if err != nil {
			return nil, err
		}
		out.Last4 = Attendance{Present: present[row.CharacterKey], Nights: nights}
		return out, nil
	}
	short, err := s.raidAttendance(ctx, guildID, keys, profileShortWindow)
	if err != nil {
		return nil, err
	}
	long, err := s.raidAttendance(ctx, guildID, keys, attendanceWindowLong)
	if err != nil {
		return nil, err
	}
	l := long[row.CharacterKey]
	out.Last4, out.Last30 = short[row.CharacterKey], &l
	if out.Nights, err = s.profileNights(ctx, guildID, row.CharacterKey); err != nil {
		return nil, err
	}
	return out, nil
}

// profileNights is the guild's last profileShortWindow logged raids, newest first, with the
// character's attendance row on each.
func (s *Store) profileNights(ctx context.Context, guildID int64, key string) ([]ProfileNight, error) {
	rows, err := s.Pool.Query(ctx, `
		select r.id, r.starts_at, a.outcome, coalesce(a.excused, false)
		from (`+loggedRaids+`) lr
		join guild_raids r on r.id = lr.id
		left join guild_raid_attendance a on a.raid_id = r.id and a.character_key = $3
		order by r.starts_at desc, r.id desc`, guildID, profileShortWindow, key)
	if err != nil {
		return nil, fmt.Errorf("guilds: profile nights: %w", err)
	}
	defer rows.Close()
	out := []ProfileNight{}
	for rows.Next() {
		var n ProfileNight
		if err := rows.Scan(&n.RaidID, &n.StartsAt, &n.Outcome, &n.Excused); err != nil {
			return nil, fmt.Errorf("guilds: profile nights: %w", err)
		}
		n.StartsAt = n.StartsAt.UTC()
		out = append(out, n)
	}
	return out, rows.Err()
}

// profileLoot is every award to key in guildID, newest first, named from the loot table.
func (s *Store) profileLoot(ctx context.Context, guildID int64, key string) (ProfileLoot, error) {
	out := ProfileLoot{Items: []ProfileLootItem{}}
	rows, err := s.Pool.Query(ctx, `select item_id, raid_id, awarded_at from loot_awards
		where guild_id = $1 and character_key = $2 order by awarded_at desc, id desc`, guildID, key)
	if err != nil {
		return out, fmt.Errorf("guilds: profile loot: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var it ProfileLootItem
		if err := rows.Scan(&it.ItemID, &it.RaidID, &it.AwardedAt); err != nil {
			return out, fmt.Errorf("guilds: profile loot: %w", err)
		}
		it.AwardedAt = it.AwardedAt.UTC()
		if item, ok := onyxiaItem(it.ItemID); ok {
			it.Name, it.Icon, it.Quality = &item.Name, &item.Icon, &item.Quality
		}
		out.Items = append(out.Items, it)
	}
	out.Count = len(out.Items)
	return out, rows.Err()
}
