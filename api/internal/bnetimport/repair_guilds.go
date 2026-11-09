package bnetimport

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// GuildFix is one guild the repair moved to the ruleset its members are
// on.
type GuildFix struct {
	GuildID  int64
	From, To string
}

// misrulesettedGuildsSQL is every guild whose Battle.net-verified members
// are all on one ruleset that is not the guild's own: a guild the import
// created under a guessed ruleset.
const misrulesettedGuildsSQL = `
select g.id, g.ruleset, min(c.ruleset),
       exists (select 1 from guilds t where t.id <> g.id and t.region = g.region
                 and t.ruleset = min(c.ruleset) and lower(t.name) = lower(g.name))
  from guilds g
  join guild_characters gc on gc.guild_id = g.id and gc.source = 'bnet'
  join characters c on c.key = gc.character_key
 group by g.id, g.ruleset, g.region, g.name
having count(distinct c.ruleset) = 1 and min(c.ruleset) <> g.ruleset
 order by g.id`

// repairGuilds moves each such guild to its members' ruleset. A guild
// whose right-ruleset twin already exists is not merged here: it is
// returned as an error naming both, so a merge is never done blind.
func (s *Service) repairGuilds(ctx context.Context, tx pgx.Tx, sum *RepairSummary) error {
	rows, err := tx.Query(ctx, misrulesettedGuildsSQL)
	if err != nil {
		return fmt.Errorf("bnetimport: repair guilds: %w", err)
	}
	var fixes []GuildFix
	for rows.Next() {
		var f GuildFix
		var twin bool
		if err := rows.Scan(&f.GuildID, &f.From, &f.To, &twin); err != nil {
			rows.Close()
			return fmt.Errorf("bnetimport: repair guild row: %w", err)
		}
		if twin {
			rows.Close()
			return fmt.Errorf("bnetimport: guild %d (%s) has a %s twin: merge it by hand", f.GuildID, f.From, f.To)
		}
		fixes = append(fixes, f)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return err
	}
	for _, f := range fixes {
		if _, err := tx.Exec(ctx, `update guilds set ruleset = $2 where id = $1`, f.GuildID, f.To); err != nil {
			return fmt.Errorf("bnetimport: move guild %d to %s: %w", f.GuildID, f.To, err)
		}
		sum.Guilds = append(sum.Guilds, f)
	}
	return nil
}
