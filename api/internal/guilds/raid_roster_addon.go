// api/internal/guilds/raid_roster_addon.go
//
// The inbox's raid_roster message (docs/contracts/addon-v2.md contract 3; spec
// design/specs/2026-11-13-addon-hub.md section 9): per open raid of a character's guild in the
// next week, who is coming by role against the raid's comp target, how many have not answered,
// and how many of the raid leader's jobs are held. It reads the raid page's own view
// (raidViewsIn), so the addon and the site never count a raid two ways. Officers (the site's
// verified officer or leader rank, seesEveryAttendanceRow) also get the names: each sign-up's
// answer and every job's holders. Everyone else gets the counts only.
package guilds

import (
	"context"
	"fmt"
	"sort"
	"time"
)

// AddonRosterCounts are a raid's answers, one per account (the raid page's own counting), and
// the verified accounts of the guild that have not answered at all.
type AddonRosterCounts struct {
	In, Late, Bench, Out, NotSigned int
}

// AddonRosterMember is one row of an officer's roster: a sign-up (the account's best place) or a
// verified character whose account has not answered (Answer AddonAnswerNone). Group is the roster
// group it fights in (raidspec's tank, healer, melee or ranged).
type AddonRosterMember struct {
	Name, Class, Answer, Group string
}

// AddonRosterJob is one slot of the raid leader's assignments: the boss, the slot's label, how
// many names it takes and who holds it (empty: open).
type AddonRosterJob struct {
	Encounter, Label string
	Count            int
	Holders          []string
}

// AddonRaidRoster is one raid_roster row for one of the account's characters in the raid's guild.
// Wanted is nil for a zone with no comp target. Members and Jobs are filled for an officer only.
type AddonRaidRoster struct {
	CharacterKey string
	RaidID       int64
	ZoneName     string
	Title        string
	StartsAt     time.Time
	Size         *int
	Officer      bool
	Counts       AddonRosterCounts
	// Coming is the In and Late rows by roster group, so the four add up to In + Late.
	Coming     RoleCounts
	Wanted     *RoleCounts
	JobsFilled int
	Jobs       int
	Members    []AddonRosterMember
	JobList    []AddonRosterJob
}

// AddonRaidRosters is, for each of keys verified in a guild, one row per raid of that guild that
// is planned or locked, not over at now, and starting within AddonSignupWindow (the raid_signup
// rule), by character then start. userID is the account asking: its rank decides the names.
func (s *Service) AddonRaidRosters(ctx context.Context, userID int64, keys []string, now time.Time) ([]AddonRaidRoster, error) {
	byGuild, err := s.Store.verifiedGuildsOf(ctx, keys)
	if err != nil {
		return nil, err
	}
	var out []AddonRaidRoster
	for _, guildID := range sortedGuildIDs(byGuild) {
		rows, err := s.guildRosters(ctx, guildID, userID, now)
		if err != nil {
			return nil, err
		}
		for _, key := range byGuild[guildID] {
			for _, r := range rows {
				r.CharacterKey = key
				out = append(out, r)
			}
		}
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].CharacterKey < out[j].CharacterKey })
	return out, nil
}

// guildRosters is the rows of one guild's open raids as userID may read them (no character key).
func (s *Service) guildRosters(ctx context.Context, guildID, userID int64, now time.Time) ([]AddonRaidRoster, error) {
	planned, err := s.Store.PlannedRaids(ctx, guildID, now)
	if err != nil {
		return nil, err
	}
	raids := addonWindowRaids(planned, now)
	if len(raids) == 0 {
		return nil, nil
	}
	vc, err := s.Store.loadRaidViewContext(ctx, guildID, userID)
	if err != nil {
		return nil, err
	}
	views, byRaid, err := s.Store.raidViewsIn(ctx, vc, raids)
	if err != nil {
		return nil, err
	}
	targets, err := s.Store.compTargetsOf(ctx, raidIDsOf(views))
	if err != nil {
		return nil, err
	}
	officer, err := s.seesEveryAttendanceRow(ctx, guildID, userID)
	if err != nil {
		return nil, err
	}
	out := make([]AddonRaidRoster, len(views))
	for i, v := range views {
		var own *RoleCounts
		if c, ok := targets[v.ID]; ok {
			own = &c
		}
		out[i] = addonRoster(v, compTargetOf(zoneOf(v.Zone), own), notSignedAccounts(vc, byRaid[v.ID]), officer)
	}
	return out, nil
}

// addonWindowRaids keeps the raids raid_signup lists: planned or locked, not over, starting
// within AddonSignupWindow of now.
func addonWindowRaids(raids []Raid, now time.Time) []Raid {
	var out []Raid
	for _, r := range raids {
		open := r.Status == RaidPlanned || r.Status == RaidLocked
		if open && r.EndsAt().After(now) && r.StartsAt.Before(now.Add(AddonSignupWindow)) {
			out = append(out, r)
		}
	}
	return out
}

// notSignedAccounts counts the verified accounts on the roster with no sign-up for the raid.
func notSignedAccounts(vc raidViewContext, signups []Signup) int {
	answered := make(map[int64]bool, len(signups))
	for _, sg := range signups {
		answered[sg.SignedBy] = true
	}
	waiting := map[int64]bool{}
	for _, c := range vc.roster {
		if c.verified && !answered[c.userID] {
			waiting[c.userID] = true
		}
	}
	return len(waiting)
}

// addonRoster reads one raid view into a row; the names only for an officer.
func addonRoster(v RaidView, target *CompTarget, notSigned int, officer bool) AddonRaidRoster {
	r := AddonRaidRoster{
		RaidID: v.ID, ZoneName: v.ZoneName, Title: v.Title, StartsAt: v.StartsAt, Size: v.Size, Officer: officer,
		Counts: AddonRosterCounts{In: v.Counts.In, Late: v.Counts.Late, Bench: v.Counts.Bench, Out: v.Counts.Out,
			NotSigned: notSigned},
		Coming: RoleCounts{
			Tanks: v.Counts.Tanks + v.Counts.LateGroups.Tanks, Healers: v.Counts.Healers + v.Counts.LateGroups.Healers,
			Melee: v.Counts.Melee + v.Counts.LateGroups.Melee, Ranged: v.Counts.Ranged + v.Counts.LateGroups.Ranged,
		},
	}
	if target != nil {
		wanted := target.RoleCounts
		r.Wanted = &wanted
	}
	jobs := rosterJobs(v.Assignments)
	for _, j := range jobs {
		r.Jobs += j.Count
		r.JobsFilled += min(len(j.Holders), j.Count)
	}
	if officer {
		r.Members = rosterMembers(v)
		r.JobList = jobs
	}
	return r
}

// rosterMembers is every sign-up row the raid page lists, then every verified character whose
// account has not answered, each in the page's own order.
func rosterMembers(v RaidView) []AddonRosterMember {
	out := make([]AddonRosterMember, 0, len(v.Signups)+len(v.NotSigned))
	for _, sg := range v.Signups {
		out = append(out, AddonRosterMember{Name: sg.Name, Class: sg.Class, Answer: sg.State, Group: sg.Group})
	}
	for _, c := range v.NotSigned {
		out = append(out, AddonRosterMember{Name: c.Name, Class: c.Class, Answer: AddonAnswerNone, Group: c.Group})
	}
	return out
}

// rosterJobs is every slot of every encounter the raid leader can fill, in the page's order.
func rosterJobs(a *RaidAssignments) []AddonRosterJob {
	if a == nil {
		return nil
	}
	var out []AddonRosterJob
	for _, e := range a.Encounters {
		for _, slot := range e.Slots {
			holders := make([]string, len(slot.Names))
			for i, n := range slot.Names {
				holders[i] = n.Name
			}
			out = append(out, AddonRosterJob{Encounter: e.Name, Label: slot.Label, Count: slot.Count, Holders: holders})
		}
	}
	return out
}

// sortedGuildIDs is byGuild's keys, ascending.
func sortedGuildIDs(byGuild map[int64][]string) []int64 {
	ids := make([]int64, 0, len(byGuild))
	for id := range byGuild {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	return ids
}

// verifiedGuildsOf maps each guild to the keys among keys verified in it, keys sorted.
func (s *Store) verifiedGuildsOf(ctx context.Context, keys []string) (map[int64][]string, error) {
	out := map[int64][]string{}
	if len(keys) == 0 {
		return out, nil
	}
	rows, err := s.Pool.Query(ctx, `
		select guild_id, character_key from guild_characters
		where character_key = any($1) and verified_at is not null
		order by guild_id, character_key`, keys)
	if err != nil {
		return nil, fmt.Errorf("guilds: addon roster guilds: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		var key string
		if err := rows.Scan(&id, &key); err != nil {
			return nil, fmt.Errorf("guilds: addon roster guilds: %w", err)
		}
		out[id] = append(out[id], key)
	}
	return out, rows.Err()
}
