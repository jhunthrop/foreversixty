// api/internal/guilds/raid_drops_match.go
//
// Recorded drops follow their night when the raid learns its log, the way its awards do
// (attachMatchLoot): every raid-less drop of the guild filed under the raid's log, or looted inside
// the raid's match window, of a table boss of the raid's zone, joins the raid, marked
// attached_by_match; a clear sends exactly those back. The drop keeps its log, so a detached drop
// reads as that logged night again. Drops have no per-raid uniqueness, so nothing is swapped.
package guilds

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// attachMatchDrops runs inside a match's transaction, after attachMatchLoot.
func attachMatchDrops(ctx context.Context, tx pgx.Tx, raidID int64) error {
	if _, err := tx.Exec(ctx, `
		update raid_drops d set raid_id = $1, attached_by_match = true
		from guild_raids g
		join reports r on r.id = g.report_id
		where g.id = $1 and d.guild_id = g.guild_id and d.raid_id is null
		  and (d.report_id = r.id
		       or (d.report_id is null
		           and d.dropped_at >= g.starts_at - make_interval(hours => $2)
		           and d.dropped_at <= least(coalesce(r.completed_at, r.created_at),
		                                     g.starts_at + make_interval(mins => g.duration_minutes)) + make_interval(hours => $3)))
		  and `+zoneEncounterSQL("g.zone", "d.encounter_id", "$4", "$5"),
		raidID, matchLeadHours, matchTailHours, bossZones(), bossEncounters()); err != nil {
		return fmt.Errorf("guilds: attach drops: %w", err)
	}
	return nil
}

// detachMatchDrops runs inside a clear's transaction: the drops raidID's match attached go back to
// no raid (their log, if they had one, stands).
func detachMatchDrops(ctx context.Context, tx pgx.Tx, raidID int64) error {
	if _, err := tx.Exec(ctx, `update raid_drops set raid_id = null, attached_by_match = false
		where raid_id = $1 and attached_by_match`, raidID); err != nil {
		return fmt.Errorf("guilds: detach drops: %w", err)
	}
	return nil
}
