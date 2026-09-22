// api/internal/bnetimport/character.go
package bnetimport

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// importOneCharacter writes one character's characters row and syncs its
// guild membership, in one transaction (spec §4.2 step 6: "so a failure
// on the 7th leaves the first six written"). wrote is false when key
// already belongs to a different account — a collision (spec §2's key-
// format RULING): the existing row is left exactly as it was.
func (s *Service) importOneCharacter(ctx context.Context, userID int64, region, ruleset, key string,
	ch bnetapi.AccountCharacter, rosterCache map[string]bnetapi.Roster) (wrote, guildWritten bool, err error) {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return false, false, fmt.Errorf("bnetimport: begin %s: %w", key, err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	faction := ch.Faction
	if faction != "alliance" && faction != "horde" {
		faction = ""
	}
	tag, err := tx.Exec(ctx,
		`insert into characters (key, region, ruleset, name, class, user_id, realm_slug, realm_name,
		    level, faction, bnet_character_id, source, imported_at, refreshed_at)
		 values ($1, $2, $3, $4, $5, $6, $7, $8, $9, nullif($10, ''), $11, 'bnet', now(), now())
		 on conflict (key) do update set
		   name = excluded.name, class = excluded.class, user_id = excluded.user_id,
		   realm_slug = excluded.realm_slug, realm_name = excluded.realm_name,
		   level = excluded.level, faction = excluded.faction, bnet_character_id = excluded.bnet_character_id,
		   source = 'bnet', imported_at = now(), refreshed_at = now()
		 where characters.user_id is null or characters.user_id = excluded.user_id`,
		key, region, ruleset, ch.Name, ch.ClassSlug, userID, ch.RealmSlug, ch.RealmName,
		ch.Level, faction, ch.ID)
	if err != nil {
		return false, false, fmt.Errorf("bnetimport: write character %s: %w", key, err)
	}
	if tag.RowsAffected() == 0 {
		return false, false, nil
	}

	guildWritten, err = s.syncCharacterGuild(ctx, tx, userID, region, ruleset, key, ch.RealmSlug, ch.Name, rosterCache)
	if err != nil {
		return false, false, err
	}
	if err := tx.Commit(ctx); err != nil {
		return false, false, fmt.Errorf("bnetimport: commit %s: %w", key, err)
	}
	return true, guildWritten, nil
}

// syncCharacterGuild reads the character's public profile (app token)
// and makes guild_characters agree with it: unguilded clears a 'bnet'-
// sourced row only (spec §4.2 step 4 — a row from another source is that
// path's own evidence, left alone); guilded resolves/creates the guild,
// finds the character's roster rank, and upserts a Reverify=true
// membership row (spec §4.2 step 5). A character whose profile or guild
// roster cannot be read (403/404) is left exactly as it was — this is
// public data that is sometimes simply private or gone, not an error.
func (s *Service) syncCharacterGuild(ctx context.Context, tx pgx.Tx, userID int64, region, ruleset, key, realmSlug, name string,
	rosterCache map[string]bnetapi.Roster) (guildWritten bool, err error) {
	var prevGuildID *int64
	var prevUserID int64
	var prevSource string
	err = tx.QueryRow(ctx,
		`select guild_id, user_id, source from guild_characters where character_key = $1`, key).
		Scan(&prevGuildID, &prevUserID, &prevSource)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return false, fmt.Errorf("bnetimport: read previous guild for %s: %w", key, err)
	}

	profile, err := s.Client.Character(ctx, region, realmSlug, name)
	if err != nil {
		if errors.Is(err, bnetapi.ErrNotFound) || errors.Is(err, bnetapi.ErrForbidden) {
			return false, nil
		}
		return false, fmt.Errorf("bnetimport: character profile %s: %w", key, err)
	}

	if !profile.HasGuild {
		if prevGuildID != nil && prevSource == "bnet" {
			if _, err := tx.Exec(ctx, `delete from guild_characters where character_key = $1`, key); err != nil {
				return false, fmt.Errorf("bnetimport: clear guild for %s: %w", key, err)
			}
			return false, guilds.AfterGuildChange(ctx, tx, *prevGuildID, prevUserID)
		}
		return false, nil
	}

	cacheKey := strings.ToLower(region + "/" + realmSlug + "/" + profile.GuildName)
	roster, ok := rosterCache[cacheKey]
	if !ok {
		roster, err = s.Client.GuildRoster(ctx, region, realmSlug, profile.GuildName)
		if err != nil {
			if errors.Is(err, bnetapi.ErrNotFound) || errors.Is(err, bnetapi.ErrForbidden) {
				return false, nil
			}
			return false, fmt.Errorf("bnetimport: guild roster %s: %w", key, err)
		}
		rosterCache[cacheKey] = roster
	}
	rankIndex, found := findRank(roster.Members, name, realmSlug)
	if !found {
		s.logger().Warn("bnetimport", "op", "roster_rank_not_found", "key", key, "guild", profile.GuildName)
		return false, nil
	}

	guildID, officerMax, err := guilds.ResolveGuild(ctx, tx, region, ruleset, profile.GuildName)
	if err != nil {
		return false, err
	}
	if err := guilds.StampBnetRoster(ctx, tx, guildID, roster.GuildID, realmSlug); err != nil {
		return false, err
	}
	rank := guilds.DeriveRank(rankIndex, officerMax)

	if prevGuildID != nil && *prevGuildID != guildID {
		if err := guilds.LockGuilds(ctx, tx, guildID, *prevGuildID); err != nil {
			return false, err
		}
		if _, err := tx.Exec(ctx,
			`delete from guild_characters where character_key = $1 and guild_id = $2`, key, *prevGuildID); err != nil {
			return false, fmt.Errorf("bnetimport: clear previous guild for %s: %w", key, err)
		}
	}

	if err := guilds.UpsertCharacterMembership(ctx, tx, guilds.MembershipRow{
		GuildID: guildID, CharacterKey: key, UserID: userID,
		RankIndex: rankIndex, Rank: rank, Source: "bnet", VerifiedBy: "bnet",
		VerifiedAt: time.Now(), Reverify: true,
	}); err != nil {
		return false, err
	}
	if rank == "leader" {
		if err := guilds.AutoConfirmClaimIfPending(ctx, tx, guildID, userID); err != nil {
			return false, err
		}
	}
	if err := guilds.AfterGuildChange(ctx, tx, guildID, userID); err != nil {
		return false, err
	}
	if prevGuildID != nil && *prevGuildID != guildID {
		if err := guilds.AfterGuildChange(ctx, tx, *prevGuildID, prevUserID); err != nil {
			return false, err
		}
	}
	return true, nil
}

// findRank locates name's own roster rank by name and realm (case-
// insensitive), since a guild roster can — rarely, on Era — carry two
// same-named characters on different realms.
func findRank(members []bnetapi.RosterMember, name, realmSlug string) (int, bool) {
	for _, m := range members {
		if strings.EqualFold(m.Name, name) && strings.EqualFold(m.RealmSlug, realmSlug) {
			return m.Rank, true
		}
	}
	return 0, false
}

// rulesetForRealm maps a character's realm slug to a ruleset via the
// region's realm list (bnetapi.RulesetOf); an unresolvable slug (a stale
// or unlisted realm) falls back to normal, the same default RulesetOf
// itself uses for an unrecognised realm type.
func rulesetForRealm(realms []bnetapi.Realm, slug string) string {
	for _, r := range realms {
		if strings.EqualFold(r.Slug, slug) {
			return bnetapi.RulesetOf(r)
		}
	}
	return character.RulesetNormal
}
