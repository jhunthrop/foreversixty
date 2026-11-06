// api/internal/guilds/loot_council_match.go
//
// Council votes and MS/OS flags follow their night when the raid learns its log, the way its
// awards do (attachMatchLoot): the raid-less rows filed under the raid's own log, or made with no
// night inside the raid's match window, join the raid. Where the raid already holds a row for the
// same drop and character (and voter), it keeps it; of two raid-less rows the newer moves.
package guilds

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// councilMatchTable is one council table as attaching moves it: its timestamp column and the
// columns one row is unique on beyond the guild, night and drop.
type councilMatchTable struct {
	name, at, keyColumns string
}

var councilMatchTables = []councilMatchTable{
	{name: "loot_votes", at: "voted_at", keyColumns: "character_key, voter_id"},
	{name: "loot_interest", at: "set_at", keyColumns: "character_key"},
}

// attachCouncilRows runs inside a match's transaction, after attachMatchLoot.
func attachCouncilRows(ctx context.Context, tx pgx.Tx, raidID int64) error {
	for _, t := range councilMatchTables {
		if _, err := tx.Exec(ctx, attachCouncilSQL(t), raidID, matchLeadHours, matchTailHours, onyxiaZoneKey); err != nil {
			return fmt.Errorf("guilds: attach %s: %w", t.name, err)
		}
	}
	return nil
}

func attachCouncilSQL(t councilMatchTable) string {
	return fmt.Sprintf(`
		update %[1]s x set raid_id = $1, report_id = null
		where x.id in (
		  select distinct on (y.encounter_id, y.item_id, %[3]s) y.id
		  from guild_raids g
		  join reports r on r.id = g.report_id
		  join %[1]s y on y.guild_id = g.guild_id and y.raid_id is null
		   and (y.report_id = r.id
		        or (y.report_id is null
		            and y.%[2]s >= g.starts_at - make_interval(hours => $2)
		            and y.%[2]s <= least(coalesce(r.completed_at, r.created_at),
		                                 g.starts_at + make_interval(mins => g.duration_minutes)) + make_interval(hours => $3)))
		  where g.id = $1 and g.zone = $4
		  order by y.encounter_id, y.item_id, %[3]s, y.%[2]s desc, y.id desc)
		and not exists (
		  select 1 from %[1]s o
		  where o.guild_id = x.guild_id and o.raid_id = $1 and o.encounter_id = x.encounter_id
		    and o.item_id = x.item_id and (%[4]s))`,
		t.name, t.at, prefixed("y", t.keyColumns), sameColumns("o", "x", t.keyColumns))
}

// sameColumns is the row comparison "(a.c1, a.c2) = (b.c1, b.c2)" over a comma-separated column list.
func sameColumns(a, b, columns string) string {
	left, right := prefixed(a, columns), prefixed(b, columns)
	return fmt.Sprintf("(%s) = (%s)", left, right)
}
