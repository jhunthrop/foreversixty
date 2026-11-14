// api/internal/directory/handover.go
package directory

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// ErrNotMergeable is a merge asked of an account that is not Discord-only.
var ErrNotMergeable = errors.New("directory: only a Discord-only account merges")

// heldRow is one character row a Discord-only account holds.
type heldRow struct {
	key    string
	userID int64
}

// TakeOverDiscordClaims runs inside the Battle.net import's transaction, before it writes the
// character's row: every row a Discord-only account holds for this Blizzard character moves to
// the Battle.net account that owns it, with its guild membership and its raid sign-ups, so
// nothing is duplicated. The Battle.net owner always wins. Returns how many rows moved.
func TakeOverDiscordClaims(ctx context.Context, tx pgx.Tx, userID int64, region string, bnetID int64) (int, error) {
	if bnetID == 0 {
		return 0, nil
	}
	rows, err := tx.Query(ctx,
		`select c.key, c.user_id from characters c join users u on u.id = c.user_id
		  where c.region = $1 and c.bnet_character_id = $2 and u.discord_only and c.user_id <> $3
		  order by c.key for update of c`, region, bnetID, userID)
	if err != nil {
		return 0, fmt.Errorf("directory: discord claims of %d: %w", bnetID, err)
	}
	held, err := pgx.CollectRows(rows, func(r pgx.CollectableRow) (heldRow, error) {
		var h heldRow
		err := r.Scan(&h.key, &h.userID)
		return h, err
	})
	if err != nil {
		return 0, fmt.Errorf("directory: discord claims of %d: %w", bnetID, err)
	}
	for _, h := range held {
		if err := moveCharacter(ctx, tx, h.key, h.userID, userID); err != nil {
			return 0, err
		}
	}
	return len(held), nil
}

// moveCharacter moves one character, its guild row and its sign-ups from one account to
// another, and recomputes both accounts' membership of its guild.
func moveCharacter(ctx context.Context, tx pgx.Tx, key string, from, to int64) error {
	if _, err := tx.Exec(ctx, `update characters set user_id = $2 where key = $1`, key, to); err != nil {
		return fmt.Errorf("directory: move character %s: %w", key, err)
	}
	if _, err := tx.Exec(ctx, `update users set main_character_key = null where id = $1 and main_character_key = $2`, from, key); err != nil {
		return fmt.Errorf("directory: move main %s: %w", key, err)
	}
	if err := moveSignups(ctx, tx, `character_key = $3`, from, to, key); err != nil {
		return err
	}
	var guildID int64
	err := tx.QueryRow(ctx,
		`update guild_characters set user_id = $2 where character_key = $1 returning guild_id`, key, to).Scan(&guildID)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("directory: move guild row %s: %w", key, err)
	}
	if err := guilds.AfterGuildChange(ctx, tx, guildID, to); err != nil {
		return err
	}
	return guilds.AfterGuildChange(ctx, tx, guildID, from)
}

// moveSignups rewrites signed_by and acted_by from one account to another on the sign-ups
// where matches (extra arguments from $3 on).
func moveSignups(ctx context.Context, tx pgx.Tx, where string, from, to int64, args ...any) error {
	all := append([]any{from, to}, args...)
	if _, err := tx.Exec(ctx,
		`update guild_raid_signups set signed_by = $2 where signed_by = $1 and `+where, all...); err != nil {
		return fmt.Errorf("directory: move sign-ups: %w", err)
	}
	if _, err := tx.Exec(ctx,
		`update guild_raid_signups set acted_by = $2 where acted_by = $1 and `+where, all...); err != nil {
		return fmt.Errorf("directory: move sign-ups: %w", err)
	}
	return nil
}

// MergeDiscordOnly folds a Discord-only account into the account its Discord user just linked
// from the site: its characters, guild rows and sign-ups move across and the empty account is
// deleted (its Discord link with it), so the caller can save the link on the real account.
func (c *Claims) MergeDiscordOnly(ctx context.Context, from, into int64) error {
	tx, err := c.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("directory: begin merge: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	var only bool
	if err := tx.QueryRow(ctx, `select discord_only from users where id = $1 for update`, from).Scan(&only); err != nil {
		return fmt.Errorf("directory: merge source %d: %w", from, err)
	}
	if !only || from == into {
		return ErrNotMergeable
	}
	rows, err := tx.Query(ctx, `select key from characters where user_id = $1 order by key`, from)
	if err != nil {
		return fmt.Errorf("directory: merge characters: %w", err)
	}
	keys, err := pgx.CollectRows(rows, pgx.RowTo[string])
	if err != nil {
		return fmt.Errorf("directory: merge characters: %w", err)
	}
	for _, key := range keys {
		if err := moveCharacter(ctx, tx, key, from, into); err != nil {
			return err
		}
	}
	// Anything it answered for a character it never held (none today) moves too.
	if err := moveSignups(ctx, tx, `true`, from, into); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx, `delete from users where id = $1`, from); err != nil {
		return fmt.Errorf("directory: delete merged account: %w", err)
	}
	return tx.Commit(ctx)
}
