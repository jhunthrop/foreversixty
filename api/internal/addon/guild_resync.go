package addon

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
)

// storedExport is one of an account's addon_exports rows, as syncGuild reads it.
type storedExport struct {
	key, region, ruleset, export string
}

// ResyncGuildExports runs syncGuild, within tx, for each of userID's stored exports on
// guildID's realm whose guild= section names guildID (guilds.InviteExports). An officer who
// removes a member deletes his guild_characters rows but not his exports, so a member who
// rejoins by invite gets his synced characters back through the very path his next sync would
// take, rank and verification rules included.
func (s *Store) ResyncGuildExports(ctx context.Context, tx pgx.Tx, userID, guildID int64) error {
	exports, guildName, err := exportsOnGuildRealm(ctx, tx, userID, guildID)
	if err != nil {
		return err
	}
	for _, e := range exports {
		name, _, ok := ParseFS1Guild(e.export)
		if !ok {
			continue
		}
		if name, ok = validateGuildName(name); !ok || !strings.EqualFold(name, guildName) {
			continue
		}
		if err := s.syncGuild(ctx, tx, userID, e.key, e.region, e.ruleset, e.export); err != nil {
			return err
		}
	}
	return nil
}

// exportsOnGuildRealm is userID's exports on guildID's region and ruleset, read in full before
// any sync writes on tx, and the guild's name.
func exportsOnGuildRealm(ctx context.Context, tx pgx.Tx, userID, guildID int64) ([]storedExport, string, error) {
	rows, err := tx.Query(ctx, `
		select ae.character_key, ae.region, ae.ruleset, ae.export, g.name
		from addon_exports ae
		join guilds g on g.id = $2 and g.region = ae.region and g.ruleset = ae.ruleset
		where ae.user_id = $1
		order by ae.character_key`, userID, guildID)
	if err != nil {
		return nil, "", fmt.Errorf("addon: read exports to resync for user %d: %w", userID, err)
	}
	defer rows.Close()
	var (
		out       []storedExport
		guildName string
	)
	for rows.Next() {
		var e storedExport
		if err := rows.Scan(&e.key, &e.region, &e.ruleset, &e.export, &guildName); err != nil {
			return nil, "", fmt.Errorf("addon: read exports to resync for user %d: %w", userID, err)
		}
		out = append(out, e)
	}
	return out, guildName, rows.Err()
}
