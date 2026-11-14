// api/internal/directory/claims.go
package directory

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// SourceDiscord marks a characters or guild_characters row a Discord-only account claimed by
// picking it from its guild's directory roster.
const SourceDiscord = "discord"

// pickPrefix starts a /signup character value that names a directory character, not a key.
const pickPrefix = "dir:"

var (
	// ErrTaken is a character another account already holds.
	ErrTaken = errors.New("directory: character held by another account")
	// ErrNotInGuild is a pick the guild's directory roster does not list.
	ErrNotInGuild = errors.New("directory: character not on the guild's roster")
	// ErrNotDiscordOnly is a pick asked by an account with a Battle.net identity: its own
	// characters come from its sign-in, never from a pick.
	ErrNotDiscordOnly = errors.New("directory: account is not Discord-only")
	// ErrTooManyPicks is a pick past MaxDiscordPicks.
	ErrTooManyPicks = errors.New("directory: too many characters picked")
)

// MaxDiscordPicks is how many characters one Discord-only account may hold: a main and alts,
// never half a roster (a pick proves nothing until a Battle.net sign-in).
const MaxDiscordPicks = 3

// Claims binds directory characters to Discord-only accounts and hands them over to the
// Battle.net account that owns them.
type Claims struct {
	Pool *pgxpool.Pool
}

// Pick is one character of a guild's directory roster as /signup's autocomplete offers it.
type Pick struct {
	BnetID int64
	Name   string
	Class  string
	Level  int
	Spec   string
}

// PickValue is the /signup character value of a directory character.
func PickValue(bnetID int64) string { return pickPrefix + strconv.FormatInt(bnetID, 10) }

// PickID reads a /signup character value back; ok is false for anything but a pick.
func PickID(value string) (int64, bool) {
	rest, ok := strings.CutPrefix(value, pickPrefix)
	if !ok {
		return 0, false
	}
	id, err := strconv.ParseInt(rest, 10, 64)
	return id, err == nil && id > 0
}

// EnsureDiscordAccount is the account behind a Discord user: the linked one, else a new
// Discord-only account (no Battle.net identity, no email) and its link. created is true when
// this call made it.
func (c *Claims) EnsureDiscordAccount(ctx context.Context, discordUserID, displayName string) (userID int64, created bool, err error) {
	if discordUserID == "" {
		return 0, false, errors.New("directory: no Discord user")
	}
	if id, ok, err := c.linked(ctx, discordUserID); err != nil || ok {
		return id, false, err
	}
	tx, err := c.Pool.Begin(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("directory: begin discord account: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if err := tx.QueryRow(ctx, `insert into users (discord_only) values (true) returning id`).Scan(&userID); err != nil {
		return 0, false, fmt.Errorf("directory: create discord account: %w", err)
	}
	_, err = tx.Exec(ctx, `insert into discord_links (user_id, discord_user_id, display_name) values ($1, $2, $3)`,
		userID, discordUserID, displayName)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		// Another press made it first: use that one.
		_ = tx.Rollback(ctx)
		id, _, err := c.linked(ctx, discordUserID)
		return id, false, err
	}
	if err != nil {
		return 0, false, fmt.Errorf("directory: link discord account: %w", err)
	}
	if err := tx.Commit(ctx); err != nil {
		return 0, false, fmt.Errorf("directory: commit discord account: %w", err)
	}
	return userID, true, nil
}

func (c *Claims) linked(ctx context.Context, discordUserID string) (int64, bool, error) {
	var id int64
	err := c.Pool.QueryRow(ctx, `select user_id from discord_links where discord_user_id = $1`, discordUserID).Scan(&id)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, false, nil
	}
	if err != nil {
		return 0, false, fmt.Errorf("directory: discord link: %w", err)
	}
	return id, true, nil
}

// IsDiscordOnly reports whether the account was made by the bot and has no Battle.net identity.
func (c *Claims) IsDiscordOnly(ctx context.Context, userID int64) (bool, error) {
	var only bool
	err := c.Pool.QueryRow(ctx, `select discord_only from users where id = $1`, userID).Scan(&only)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("directory: discord-only: %w", err)
	}
	return only, nil
}

// guildDirectoryKey is how a guilds row finds its directory roster: region, ruleset, its
// officer threshold and its Blizzard guild id (its own, else the directory's for its slug).
type guildDirectoryKey struct {
	region, ruleset string
	officerMax      int
	bnetID          int64
}

func guildKey(ctx context.Context, q Querier, guildID int64) (guildDirectoryKey, error) {
	var k guildDirectoryKey
	var name string
	var own *int64
	err := q.QueryRow(ctx, `select region, ruleset, name, officer_max_rank_index, bnet_guild_id from guilds where id = $1`, guildID).
		Scan(&k.region, &k.ruleset, &name, &k.officerMax, &own)
	if err != nil {
		return k, fmt.Errorf("directory: guild %d: %w", guildID, err)
	}
	if own != nil {
		k.bnetID = *own
		return k, nil
	}
	err = q.QueryRow(ctx, `select coalesce(bnet_guild_id, 0) from directory_guilds where region = $1 and slug = $2`,
		k.region, bnetapi.GuildSlug(name)).Scan(&k.bnetID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return k, fmt.Errorf("directory: guild %d slug: %w", guildID, err)
	}
	return k, nil
}

// GuildPicks is up to limit characters of the guild's directory roster whose name contains
// typed, highest level first, leaving out every character an account already holds.
func (c *Claims) GuildPicks(ctx context.Context, guildID int64, typed string, limit int) ([]Pick, error) {
	k, err := guildKey(ctx, c.Pool, guildID)
	if err != nil || k.bnetID == 0 {
		return []Pick{}, err
	}
	rows, err := c.Pool.Query(ctx, `
		select d.bnet_character_id, d.name, coalesce(d.class, ''), coalesce(d.level, 0), coalesce(d.spec, '')
		  from character_directory d
		 where d.region = $1 and d.guild_bnet_id = $2 and strpos(lower(d.name), lower($3)) > 0
		   and not exists (select 1 from characters ch
		                    where ch.region = d.region and ch.bnet_character_id = d.bnet_character_id and ch.user_id is not null)
		 order by d.level desc nulls last, d.name limit $4`,
		k.region, k.bnetID, strings.TrimSpace(typed), limit)
	if err != nil {
		return nil, fmt.Errorf("directory: guild picks: %w", err)
	}
	return pgx.CollectRows(rows, func(r pgx.CollectableRow) (Pick, error) {
		var p Pick
		err := r.Scan(&p.BnetID, &p.Name, &p.Class, &p.Level, &p.Spec)
		return p, err
	})
}

// pickedCharacter is a directory character as a claim writes it.
type pickedCharacter struct {
	name, class, race, faction string
	level                      *int
	ruleset                    string
	rankIndex                  *int
}

// Claim is a character a Discord-only account picked: its key and name, and whether this pick
// made the claim (false for a character the account already held).
type Claim struct {
	Key, Name string
	Created   bool
}

// ClaimForDiscord gives a Discord-only account one character of its guild's directory roster:
// a characters row and a verified guild_characters row, both source 'discord', at member rank
// whatever the character's guild rank (a pick proves nothing an officer's powers could rest
// on). ErrTaken (with the claim's name) when another account holds it.
func (c *Claims) ClaimForDiscord(ctx context.Context, guildID, userID, bnetID int64) (Claim, error) {
	only, err := c.IsDiscordOnly(ctx, userID)
	if err != nil {
		return Claim{}, err
	}
	if !only {
		return Claim{}, ErrNotDiscordOnly
	}
	tx, err := c.Pool.Begin(ctx)
	if err != nil {
		return Claim{}, fmt.Errorf("directory: begin claim: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	k, err := guildKey(ctx, tx, guildID)
	if err != nil {
		return Claim{}, err
	}
	pc, err := pickedOnRoster(ctx, tx, k, bnetID)
	if err != nil {
		return Claim{}, err
	}
	claim := Claim{Key: character.Key(k.region, pc.ruleset, pc.name), Name: pc.name}
	if err := tx.QueryRow(ctx, `select not exists (select 1 from characters where key = $1 and user_id = $2)`,
		claim.Key, userID).Scan(&claim.Created); err != nil {
		return Claim{}, fmt.Errorf("directory: claim held: %w", err)
	}
	if claim.Created {
		if err := underPickCap(ctx, tx, userID); err != nil {
			return claim, err
		}
	}
	key := claim.Key
	if err := claimCharacterRow(ctx, tx, key, k.region, userID, bnetID, pc); err != nil {
		return claim, err
	}
	if err := guilds.UpsertCharacterMembership(ctx, tx, guilds.MembershipRow{
		GuildID: guildID, CharacterKey: key, UserID: userID, RankIndex: rankIndexOf(pc),
		Rank: "member", Source: SourceDiscord, VerifiedBy: SourceDiscord, VerifiedAt: time.Now(), Reverify: true,
	}); err != nil {
		return Claim{}, err
	}
	if err := guilds.AfterGuildChange(ctx, tx, guildID, userID); err != nil {
		return Claim{}, err
	}
	if err := tx.Commit(ctx); err != nil {
		return Claim{}, fmt.Errorf("directory: commit claim: %w", err)
	}
	return claim, nil
}

// underPickCap is ErrTooManyPicks when the account already holds MaxDiscordPicks characters.
// Every characters row counts: a Discord-only account holds nothing but its picks.
func underPickCap(ctx context.Context, tx pgx.Tx, userID int64) error {
	var held int
	if err := tx.QueryRow(ctx, `select count(*) from characters where user_id = $1`, userID).Scan(&held); err != nil {
		return fmt.Errorf("directory: picks held: %w", err)
	}
	if held >= MaxDiscordPicks {
		return ErrTooManyPicks
	}
	return nil
}

func rankIndexOf(pc pickedCharacter) int {
	if pc.rankIndex == nil {
		return 0
	}
	return *pc.rankIndex
}

// pickedOnRoster reads a directory character of the guild's roster; ErrNotInGuild when the
// roster does not list it. Its ruleset is its own, else the guild's.
func pickedOnRoster(ctx context.Context, tx pgx.Tx, k guildDirectoryKey, bnetID int64) (pickedCharacter, error) {
	var pc pickedCharacter
	var ruleset *string
	err := tx.QueryRow(ctx,
		`select name, coalesce(class, ''), coalesce(race, ''), coalesce(faction, ''), level, ruleset, guild_rank
		   from character_directory where region = $1 and bnet_character_id = $2 and guild_bnet_id = $3`,
		k.region, bnetID, k.bnetID).Scan(&pc.name, &pc.class, &pc.race, &pc.faction, &pc.level, &ruleset, &pc.rankIndex)
	if errors.Is(err, pgx.ErrNoRows) || k.bnetID == 0 {
		return pc, ErrNotInGuild
	}
	if err != nil {
		return pc, fmt.Errorf("directory: picked character: %w", err)
	}
	pc.ruleset = k.ruleset
	if ruleset != nil {
		pc.ruleset = *ruleset
	}
	return pc, nil
}

// claimCharacterRow writes the claimed characters row for userID; ErrTaken when the key or the
// Blizzard id already belongs to another account.
func claimCharacterRow(ctx context.Context, tx pgx.Tx, key, region string, userID, bnetID int64, pc pickedCharacter) error {
	var other bool
	if err := tx.QueryRow(ctx,
		`select exists (select 1 from characters where region = $1 and bnet_character_id = $2
		                  and user_id is not null and user_id <> $3)`, region, bnetID, userID).Scan(&other); err != nil {
		return fmt.Errorf("directory: claim holder: %w", err)
	}
	if other {
		return ErrTaken
	}
	tag, err := tx.Exec(ctx,
		`insert into characters (key, region, ruleset, name, class, user_id, level, faction, race, bnet_character_id, source)
		 values ($1, $2, $3, $4, nullif($5, ''), $6, $7, nullif($8, ''), nullif($9, ''), $10, $11)
		 on conflict (key) do update set user_id = excluded.user_id,
		   class = coalesce(characters.class, excluded.class), level = coalesce(excluded.level, characters.level),
		   bnet_character_id = excluded.bnet_character_id
		 where characters.user_id is null or characters.user_id = excluded.user_id`,
		key, region, pc.ruleset, pc.name, pc.class, userID, pc.level, pc.faction, pc.race, bnetID, SourceDiscord)
	if err != nil {
		return fmt.Errorf("directory: claim character %s: %w", key, err)
	}
	if tag.RowsAffected() == 0 {
		return ErrTaken
	}
	return nil
}
