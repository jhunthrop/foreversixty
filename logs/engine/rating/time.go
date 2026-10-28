// logs/engine/rating/time.go
package rating

import "github.com/jhunthrop/foreversixty/logs/engine/summary"

// TimeAliveMS is the player's own time alive this fight: their first
// death, or the whole fight if they never died (spec §1.3's Activity
// formula defines TimeAliveMS this way; Mechanics' interrupt/dispel
// per-second rates use the same denominator, since a dead player cannot
// interrupt or dispel anything either).
func TimeAliveMS(fight summary.Summary, player string) int64 {
	alive := fight.DurationMS
	for _, d := range fight.Deaths {
		if d.GUID == player && d.AtMS < alive {
			alive = d.AtMS
		}
	}
	if alive < 0 {
		alive = 0
	}
	return alive
}
