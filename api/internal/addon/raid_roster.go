// api/internal/addon/raid_roster.go
//
// The raid_roster inbox message (docs/contracts/addon-v2.md contract 3; spec
// design/specs/2026-11-13-addon-hub.md section 9): per open raid of a character's guild in the
// next week, the answers, who is coming by role against the raid's comp target, and how many of
// the raid leader's jobs are held (guilds.Service.AddonRaidRosters). The site decides who reads
// names: an officer of the guild also gets every member's answer and every job's holders,
// everyone else the counts. It is nested, so the companion renders it as a type it knows.
package addon

import (
	"context"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
)

const messageRaidRoster = "raid_roster"

// RaidRosterSource is guilds.Service's raid_roster read.
type RaidRosterSource interface {
	AddonRaidRosters(ctx context.Context, userID int64, keys []string, now time.Time) ([]guilds.AddonRaidRoster, error)
}

// RosterSignups are a raid's answers, one per account, and the accounts that have not answered.
type RosterSignups struct {
	In        int `json:"in"`
	Late      int `json:"late"`
	Bench     int `json:"bench"`
	Out       int `json:"out"`
	NotSigned int `json:"none"`
}

// RosterRole is one roster group's In and Late rows against the raid's comp target; Wanted is
// absent for a zone with no target.
type RosterRole struct {
	Role   string `json:"role"`
	Coming int    `json:"coming"`
	Wanted *int   `json:"wanted,omitempty"`
}

// RosterJobs is how many of the raid leader's job places are held, of how many.
type RosterJobs struct {
	Held  int `json:"held"`
	Total int `json:"total"`
}

// RaidRosterMember is one row of an officer's roster: Answer is in, late, bench, out or none.
type RaidRosterMember struct {
	Name   string `json:"name"`
	Class  string `json:"class"`
	Answer string `json:"answer"`
	Role   string `json:"role"`
}

// RosterAssignment is one job of the raid leader's: the boss, the label, how many names it takes
// and who holds it (an empty list when open).
type RosterAssignment struct {
	Encounter string   `json:"encounter"`
	Label     string   `json:"label"`
	Count     int      `json:"count"`
	Holders   []string `json:"holders"`
}

// raidRosterMessages is one raid_roster per open raid of each of keys' guilds; none on a failed
// read, which is logged (the sync never fails over it).
func (s *Service) raidRosterMessages(ctx context.Context, userID int64, keys []string, now time.Time) []MessageEntry {
	if s.RaidRosters == nil || len(keys) == 0 {
		return nil
	}
	rows, err := s.RaidRosters.AddonRaidRosters(ctx, userID, keys, now)
	if err != nil {
		s.logger().Warn("addon", "op", "messages_raid_roster", "err", err)
		return nil
	}
	out := make([]MessageEntry, 0, len(rows))
	for _, r := range rows {
		out = append(out, raidRosterEntry(r))
	}
	return out
}

// raidRosterEntry is one row as its message; members and assignments only for an officer.
func raidRosterEntry(r guilds.AddonRaidRoster) MessageEntry {
	jobs := RosterJobs{Held: r.JobsFilled, Total: r.Jobs}
	m := MessageEntry{
		Type: messageRaidRoster, Character: r.CharacterKey, RaidID: r.RaidID, Raid: r.ZoneName, Title: r.Title,
		StartsAt: r.StartsAt.UTC().Format(time.RFC3339), Size: r.Size,
		Signups: &RosterSignups{In: r.Counts.In, Late: r.Counts.Late, Bench: r.Counts.Bench, Out: r.Counts.Out,
			NotSigned: r.Counts.NotSigned},
		Roles: rosterRoles(r.Coming, r.Wanted), Jobs: &jobs,
	}
	if !r.Officer {
		return m
	}
	members := make([]RaidRosterMember, len(r.Members))
	for i, mb := range r.Members {
		members[i] = RaidRosterMember{Name: mb.Name, Class: mb.Class, Answer: mb.Answer, Role: mb.Group}
	}
	assignments := make([]RosterAssignment, len(r.JobList))
	for i, j := range r.JobList {
		holders := append([]string{}, j.Holders...)
		assignments[i] = RosterAssignment{Encounter: j.Encounter, Label: j.Label, Count: j.Count, Holders: holders}
	}
	m.Members, m.Assignments = &members, &assignments
	return m
}

// rosterRoles is the four roster groups in the raid page's order: tank, healer, melee, ranged.
func rosterRoles(coming guilds.RoleCounts, wanted *guilds.RoleCounts) []RosterRole {
	pick := func(c guilds.RoleCounts, group string) int {
		switch group {
		case raidspec.GroupTank:
			return c.Tanks
		case raidspec.GroupHealer:
			return c.Healers
		case raidspec.GroupMelee:
			return c.Melee
		default:
			return c.Ranged
		}
	}
	groups := []string{raidspec.GroupTank, raidspec.GroupHealer, raidspec.GroupMelee, raidspec.GroupRanged}
	out := make([]RosterRole, len(groups))
	for i, g := range groups {
		out[i] = RosterRole{Role: g, Coming: pick(coming, g)}
		if wanted != nil {
			n := pick(*wanted, g)
			out[i].Wanted = &n
		}
	}
	return out
}
