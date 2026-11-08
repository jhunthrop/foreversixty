// api/internal/guilds/raids_last_week.go
//
// The not-answered rows' last_week (spec 2026-11-08-app-flows F3): each account's answer for
// the guild's raid of the week before (its latest raid that started in the 8 days before this
// one, cancelled ones aside), its best place when it answered with more than one character,
// missed when it gave none; null when the guild had no such raid.
package guilds

import (
	"context"
	"fmt"
	"time"
)

// lastWeekMissed is an account that gave no answer for the raid the week before.
const lastWeekMissed = "missed"

// lastWeekWindow is how far back the raid the week before may have started.
const lastWeekWindow = 8 * 24 * time.Hour

// signupRank orders an account's answers best first, as its row reads them.
var signupRank = map[string]int{SignupIn: 0, SignupLate: 1, SignupBench: 2, SignupOut: 3}

// groupRoles is the sign-up role a roster group plays.
var groupRoles = map[string]string{"tank": RoleTank, "healer": RoleHealer, "melee": RoleDPS, "ranged": RoleDPS}

// fillLastWeek sets each row's last_week for the raid v.
func (s *Store) fillLastWeek(ctx context.Context, guildID int64, v RaidView, rows []NotSignedRaider) error {
	answers, found, err := s.lastWeekAnswers(ctx, guildID, v.ID, v.StartsAt)
	if err != nil || !found {
		return err
	}
	for i := range rows {
		state, ok := answers[rows[i].userID]
		if !ok {
			state = lastWeekMissed
		}
		rows[i].LastWeek = &state
	}
	return nil
}

// lastWeekAnswers is each account's best answer for the guild's raid the week before the raid
// starting at start; found is false when there was none.
func (s *Store) lastWeekAnswers(ctx context.Context, guildID, raidID int64, start time.Time) (map[int64]string, bool, error) {
	var prev int64
	err := s.Pool.QueryRow(ctx, `select coalesce((select id from guild_raids
		where guild_id = $1 and id <> $2 and status <> 'cancelled' and starts_at < $3 and starts_at >= $4
		order by starts_at desc limit 1), 0)`, guildID, raidID, start, start.Add(-lastWeekWindow)).Scan(&prev)
	if err != nil {
		return nil, false, fmt.Errorf("guilds: raid last week: %w", err)
	}
	if prev == 0 {
		return nil, false, nil
	}
	signups, err := s.RaidSignups(ctx, prev)
	if err != nil {
		return nil, false, err
	}
	out := map[int64]string{}
	for _, sg := range signups {
		if cur, ok := out[sg.SignedBy]; !ok || signupRank[sg.State] < signupRank[cur] {
			out[sg.SignedBy] = sg.State
		}
	}
	return out, true, nil
}
