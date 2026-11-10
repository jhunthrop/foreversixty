// api/internal/guilds/member_profile_history.go
//
// What the member profile reads off a raider's last 30 logged raids and the guild's loot: the
// attendance streak, the nights they came since their last item, and the guild's median item
// count (guild-officer review, lane h-edges: "a status card" until the numbers have a baseline).
// The night arithmetic is pure, so a unit test pins it without a database.
package guilds

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
)

// attended is a night the raider came to: present or late. An excused night is neither.
func attended(n ProfileNight) bool {
	return n.Outcome != nil && (*n.Outcome == "present" || *n.Outcome == "late")
}

// attendanceStreak is the nights in a row, newest first, the raider came to; an excused night is
// skipped, any other night ends the run.
func attendanceStreak(newestFirst []ProfileNight) int {
	streak := 0
	for _, n := range newestFirst {
		switch {
		case attended(n):
			streak++
		case n.Excused:
			continue
		default:
			return streak
		}
	}
	return streak
}

// nightsSinceAward is how many of the nights the raider came to started after their last award;
// with no award (last nil) every night they came to counts.
func nightsSinceAward(newestFirst []ProfileNight, last *time.Time) int {
	count := 0
	for _, n := range newestFirst {
		if last != nil && !n.StartsAt.After(*last) {
			break
		}
		if attended(n) {
			count++
		}
	}
	return count
}

// lastAwardAt is the newest award's time (items are newest first), nil with none.
func lastAwardAt(newestFirst []ProfileLootItem) *time.Time {
	if len(newestFirst) == 0 {
		return nil
	}
	return &newestFirst[0].AwardedAt
}

// medianOf is the median of values (the lower middle pair's mean, rounded down); 0 for none.
func medianOf(values []int) int {
	if len(values) == 0 {
		return 0
	}
	sorted := append([]int(nil), values...)
	sort.Ints(sorted)
	mid := len(sorted) / 2
	if len(sorted)%2 == 0 {
		return (sorted[mid-1] + sorted[mid]) / 2
	}
	return sorted[mid]
}

// guildMedianAwards is the median award count across roster's verified characters, a character
// with no award counting as zero.
func (s *Store) guildMedianAwards(ctx context.Context, guildID int64, roster []RosterRow) (int, error) {
	rows, err := s.Pool.Query(ctx, `select character_key, count(*) from loot_awards where guild_id = $1
		group by character_key`, guildID)
	if err != nil {
		return 0, fmt.Errorf("guilds: guild median awards: %w", err)
	}
	defer rows.Close()
	awards := map[string]int{}
	for rows.Next() {
		var key string
		var n int
		if err := rows.Scan(&key, &n); err != nil {
			return 0, fmt.Errorf("guilds: guild median awards: %w", err)
		}
		awards[key] = n
	}
	if err := rows.Err(); err != nil {
		return 0, fmt.Errorf("guilds: guild median awards: %w", err)
	}
	counts := []int{}
	for _, row := range roster {
		if row.Verified {
			counts = append(counts, awards[awardKeyOf(row)])
		}
	}
	return medianOf(counts), nil
}

// ProfileUpgrade is one slot of the raider's gear gap: the BiS pick, where it comes from and its
// DPS gain over what they wear (null when no listed alternative measured the worn item).
type ProfileUpgrade struct {
	Slot    string   `json:"slot"`
	ItemID  int      `json:"item_id"`
	Name    string   `json:"name"`
	Source  string   `json:"source"`
	GainDps *float64 `json:"gain_dps"`
}

// profileUpgrades is the gap's pieces, the biggest known gain first, the unmeasured after them in
// the band's slot order.
func profileUpgrades(pieces []bis.GapPiece) []ProfileUpgrade {
	out := make([]ProfileUpgrade, 0, len(pieces))
	for _, p := range pieces {
		out = append(out, ProfileUpgrade{Slot: p.Slot, ItemID: p.ItemID, Name: p.ItemName, Source: p.Source, GainDps: p.GainDps})
	}
	sort.SliceStable(out, func(i, j int) bool {
		a, b := out[i].GainDps, out[j].GainDps
		if a == nil || b == nil {
			return a != nil && b == nil
		}
		return *a > *b
	})
	return out
}
