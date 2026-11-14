// api/internal/bnetimport/character.go
package bnetimport

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/directory"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/synclog"
)

// characterRef is one imported character as every enrichment step reads
// it: whose it is, its key, and what its Blizzard profile is addressed by.
// RealmSlug is empty for a Forever character, whose profile is addressed
// by {name}-{surname} alone.
type characterRef struct {
	UserID    int64
	Region    string
	Ruleset   string
	Key       string
	RealmSlug string
	Name      string
	BnetID    int64
	// RealmID is the realm id the account profile's protected-character link
	// carries; 0 when it carried none.
	RealmID int64
	// Answered is true when Ruleset is the player's own answer to the
	// unknown-realm question (realm_answer.go), not a realm source's.
	Answered bool
}

// writeCharacterRow is the first import pass for one character (spec §4.2
// step 3): the characters row, straight from the account profile, in its
// own transaction. wrote is false when key already belongs to a different
// account, a collision (spec §2's key-format RULING): the existing row is
// left exactly as it was. A new row's refreshed_at starts a day back, so
// the nightly refresh enriches it if no sign-in run gets to it first.
func (s *Service) writeCharacterRow(ctx context.Context, ref characterRef, ch bnetapi.AccountCharacter) (bool, error) {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return false, fmt.Errorf("bnetimport: begin %s: %w", ref.Key, err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	// A character a Discord-only account picked on Discord is this account's: the Battle.net
	// owner takes it over, with its guild row and sign-ups, before the row is written.
	if moved, err := directory.TakeOverDiscordClaims(ctx, tx, ref.UserID, ref.Region, ch.ID); err != nil {
		return false, err
	} else if moved > 0 {
		s.logger().Info("bnetimport", "op", "discord_claim_taken_over", "user_id", ref.UserID, "rows", moved)
	}
	if err := s.carryProvisionalKey(ctx, tx, ref); err != nil {
		return false, err
	}
	if err := s.rekeyIfNeeded(ctx, tx, ref.UserID, ref.Region, ref.Key, ch.ID); err != nil {
		return false, err
	}
	faction := ch.Faction
	if faction != "alliance" && faction != "horde" {
		faction = ""
	}
	accountRaw := s.capCapture("bnet_account", ref.Key, ch.Raw)
	tag, err := tx.Exec(ctx,
		`insert into characters (key, region, ruleset, name, class, user_id, realm_slug, realm_name,
		    level, faction, bnet_character_id, source, imported_at, refreshed_at,
		    race, gender, bnet_account, bnet_captured_at, bnet_realm_id, ruleset_answered)
		 values ($1, $2, $3, $4, $5, $6, nullif($7, ''), nullif($8, ''), $9, nullif($10, ''), $11, 'bnet', now(),
		    now() - interval '1 day', nullif($12, ''), nullif($13, ''), $14, now(), nullif($15::bigint, 0), $16)
		 on conflict (key) do update set
		   name = excluded.name, class = excluded.class, user_id = excluded.user_id,
		   realm_slug = excluded.realm_slug, realm_name = excluded.realm_name,
		   level = excluded.level, faction = excluded.faction, bnet_character_id = excluded.bnet_character_id,
		   source = 'bnet', imported_at = now(),
		   race = excluded.race, gender = excluded.gender,
		   bnet_account = coalesce(excluded.bnet_account, characters.bnet_account), bnet_captured_at = now(),
		   bnet_realm_id = coalesce(excluded.bnet_realm_id, characters.bnet_realm_id),
		   ruleset_answered = excluded.ruleset_answered
		 where characters.user_id is null or characters.user_id = excluded.user_id`,
		ref.Key, ref.Region, ref.Ruleset, ch.Name, ch.ClassSlug, ref.UserID, ch.RealmSlug, ch.RealmName,
		ch.Level, faction, ch.ID, ch.RaceName, ch.GenderType, rawOrNil(accountRaw), ref.RealmID, ref.Answered)
	if err != nil {
		return false, fmt.Errorf("bnetimport: write character %s: %w", ref.Key, err)
	}
	if tag.RowsAffected() == 0 {
		return false, nil
	}
	if err := tx.Commit(ctx); err != nil {
		return false, fmt.Errorf("bnetimport: commit %s: %w", ref.Key, err)
	}
	return true, nil
}

// recentlyEnriched reports whether key was enriched within enrichFreshness.
func (s *Service) recentlyEnriched(ctx context.Context, key string) (bool, error) {
	var fresh bool
	err := s.Pool.QueryRow(ctx,
		`select refreshed_at > now() - $2::interval from characters where key = $1`,
		key, fmt.Sprintf("%d seconds", int(enrichFreshness.Seconds()))).Scan(&fresh)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("bnetimport: freshness %s: %w", key, err)
	}
	return fresh, nil
}

// enrichCharacter is the second import pass for one character, and the
// whole of the nightly refresh's per-character work: the public profile and
// guild membership, gear, talents and media, then the Blizzard-sourced
// build, in one transaction (spec §4.2 steps 4-6). unavailable is true when
// the public profile answered 404 (spec A2).
func (s *Service) enrichCharacter(ctx context.Context, ref characterRef, rosterCache map[string]bnetapi.Roster) (guildWritten, unavailable bool, err error) {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return false, false, fmt.Errorf("bnetimport: begin enrich %s: %w", ref.Key, err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	if err := s.enrichInTx(ctx, tx, &ref, rosterCache, &guildWritten, &unavailable); err != nil {
		return false, false, err
	}
	if err := tx.Commit(ctx); err != nil {
		return false, false, fmt.Errorf("bnetimport: commit enrich %s: %w", ref.Key, err)
	}
	s.simImport(ctx, ref, true)
	return guildWritten, unavailable, nil
}

// enrichInTx runs the enrichment steps inside tx, which the caller owns;
// the nightly refresh calls it inside the transaction that also re-resolved
// the row's key.
func (s *Service) enrichInTx(ctx context.Context, tx pgx.Tx, ref *characterRef, rosterCache map[string]bnetapi.Roster,
	guildWritten, unavailable *bool) error {
	if _, err := tx.Exec(ctx, `update characters set refreshed_at = now() where key = $1`, ref.Key); err != nil {
		return fmt.Errorf("bnetimport: stamp %s: %w", ref.Key, err)
	}
	profile, wroteGuild, gone, err := s.syncCharacterGuild(ctx, tx, ref, rosterCache)
	if err != nil {
		return err
	}
	*guildWritten, *unavailable = wroteGuild, gone
	rawEquipment, err := s.captureEquipment(ctx, tx, ref.Key, ref.Region, ref.RealmSlug, ref.Name)
	if err != nil {
		return err
	}
	rawSpecializations, err := s.captureSpecializations(ctx, tx, ref.Key, ref.Region, ref.RealmSlug, ref.Name)
	if err != nil {
		return err
	}
	if err := s.captureMedia(ctx, tx, ref.Key, ref.Region, ref.RealmSlug, ref.Name); err != nil {
		return err
	}
	if err := s.captureStatistics(ctx, tx, ref.Key, ref.Region, ref.RealmSlug, ref.Name); err != nil {
		return err
	}
	if err := s.buildAndWriteExport(ctx, tx, ref.UserID, ref.Key, ref.Region, ref.Ruleset, profile, rawEquipment, rawSpecializations); err != nil {
		return err
	}
	return synclog.Record(ctx, tx, ref.Key, synclog.SourceBlizzard, synclog.OutcomeOK)
}

// rekeyIfNeeded implements spec A3: a character whose Blizzard id already
// belongs to this account under a *different* key (a realm's ruleset
// resolved differently between imports, or a realm transfer) is moved,
// not duplicated. The old guild_characters row is deleted (running
// AfterGuildChange for its guild and user) and the old characters row is
// deleted, inside the same transaction the caller commits after writing
// the new row.
func (s *Service) rekeyIfNeeded(ctx context.Context, tx pgx.Tx, userID int64, region, newKey string, bnetCharacterID int64) error {
	var oldKey string
	err := tx.QueryRow(ctx,
		`select key from characters where bnet_character_id = $1 and user_id = $2 and key != $3 and region = $4`,
		bnetCharacterID, userID, newKey, region).Scan(&oldKey)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("bnetimport: rekey lookup for %s: %w", newKey, err)
	}

	var guildID *int64
	err = tx.QueryRow(ctx, `select guild_id from guild_characters where character_key = $1`, oldKey).Scan(&guildID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return fmt.Errorf("bnetimport: rekey guild lookup for %s: %w", oldKey, err)
	}
	if guildID != nil {
		if _, err := tx.Exec(ctx, `delete from guild_characters where character_key = $1`, oldKey); err != nil {
			return fmt.Errorf("bnetimport: rekey clear guild for %s: %w", oldKey, err)
		}
		if err := guilds.AfterGuildChange(ctx, tx, *guildID, userID); err != nil {
			return err
		}
	}
	if _, err := tx.Exec(ctx, `delete from characters where key = $1`, oldKey); err != nil {
		return fmt.Errorf("bnetimport: rekey delete old row %s: %w", oldKey, err)
	}
	// The account's chosen main follows its character to the new key.
	if _, err := tx.Exec(ctx, `update users set main_character_key = $1 where main_character_key = $2`, newKey, oldKey); err != nil {
		return fmt.Errorf("bnetimport: rekey main %s: %w", oldKey, err)
	}
	s.logger().Info("bnetimport", "op", "rekey", "old_key", oldKey, "new_key", newKey)
	return nil
}

// syncCharacterGuild reads the character's public profile (app token),
// captures it (spec §B), and makes guild_characters agree with it:
// unguilded clears a 'bnet'-sourced row only (spec §4.2 step 4 — a row
// from another source is that path's own evidence, left alone); guilded
// resolves/creates the guild, finds the character's roster rank, and
// upserts a Reverify=true membership row (spec §4.2 step 5). A character
// whose guild roster cannot be read (403/404) is left exactly as it was
// — this is public data that is sometimes simply private or gone, not an
// error. A profile that answers 404 (spec A2 — a classic1x/Season of
// Discovery character) is logged and reported via unavailable, and 403
// (a private profile) is left silent, as before.
//
// A profile that names a ruleset other than the one ref is keyed under moves
// the row to the profile's key first (adoptProfileRuleset), and ref follows,
// so the guild below resolves under the right ruleset.
func (s *Service) syncCharacterGuild(ctx context.Context, tx pgx.Tx, ref *characterRef,
	rosterCache map[string]bnetapi.Roster) (profile bnetapi.CharacterProfile, guildWritten, unavailable bool, err error) {
	userID, region, ruleset, key, realmSlug, name := ref.UserID, ref.Region, ref.Ruleset, ref.Key, ref.RealmSlug, ref.Name
	var prevGuildID *int64
	var prevUserID int64
	var prevSource string
	err = tx.QueryRow(ctx,
		`select guild_id, user_id, source from guild_characters where character_key = $1`, key).
		Scan(&prevGuildID, &prevUserID, &prevSource)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return bnetapi.CharacterProfile{}, false, false, fmt.Errorf("bnetimport: read previous guild for %s: %w", key, err)
	}

	profile, rawProfile, err := s.Client.Character(ctx, region, realmSlug, name)
	if err != nil {
		if errors.Is(err, bnetapi.ErrNotFound) {
			s.logger().Info("bnetimport", "op", "profile_unavailable", "key", key, "status", http.StatusNotFound)
			return bnetapi.CharacterProfile{}, false, true, nil
		}
		if errors.Is(err, bnetapi.ErrForbidden) {
			return bnetapi.CharacterProfile{}, false, false, nil
		}
		return bnetapi.CharacterProfile{}, false, false, fmt.Errorf("bnetimport: character profile %s: %w", key, err)
	}
	if err := s.captureProfile(ctx, tx, key, profile, rawProfile); err != nil {
		return profile, false, false, err
	}
	s.noteStaticNamespace(region, rawProfile)
	if err := s.adoptProfileRuleset(ctx, tx, ref, profile.Ruleset); err != nil {
		return profile, false, false, err
	}
	ruleset, key = ref.Ruleset, ref.Key
	// A character on a realm no source has typed has no ruleset to resolve its
	// guild under; the guild follows once the player answers or the realm resolves.
	if ruleset == character.RulesetUnknown {
		return profile, false, false, nil
	}

	if !profile.HasGuild {
		if prevGuildID != nil && prevSource == "bnet" {
			if _, err := tx.Exec(ctx, `delete from guild_characters where character_key = $1`, key); err != nil {
				return profile, false, false, fmt.Errorf("bnetimport: clear guild for %s: %w", key, err)
			}
			return profile, false, false, guilds.AfterGuildChange(ctx, tx, *prevGuildID, prevUserID)
		}
		return profile, false, false, nil
	}

	guildRealm := profile.GuildRealmSlug
	if guildRealm == "" {
		guildRealm = realmSlug
	}
	cacheKey := strings.ToLower(region + "/" + guildRealm + "/" + profile.GuildName)
	roster, ok := rosterCache[cacheKey]
	if !ok {
		roster, err = s.Client.GuildRoster(ctx, region, guildRealm, profile.GuildName)
		if err != nil {
			if errors.Is(err, bnetapi.ErrNotFound) || errors.Is(err, bnetapi.ErrForbidden) {
				return profile, false, false, nil
			}
			return profile, false, false, fmt.Errorf("bnetimport: guild roster %s: %w", key, err)
		}
		rosterCache[cacheKey] = roster
	}
	rankIndex, found := findRank(roster.Members, ref.BnetID, name, realmSlug)
	if !found {
		s.logger().Warn("bnetimport", "op", "roster_rank_not_found", "key", key, "guild", profile.GuildName)
		return profile, false, false, nil
	}

	guildID, officerMax, err := guilds.ResolveGuild(ctx, tx, region, ruleset, profile.GuildName)
	if err != nil {
		return profile, false, false, err
	}
	if err := guilds.StampBnetRoster(ctx, tx, guildID, roster.GuildID, guildRealm); err != nil {
		return profile, false, false, err
	}
	rank := guilds.DeriveRank(rankIndex, officerMax)

	if prevGuildID != nil && *prevGuildID != guildID {
		if err := guilds.LockGuilds(ctx, tx, guildID, *prevGuildID); err != nil {
			return profile, false, false, err
		}
		if _, err := tx.Exec(ctx,
			`delete from guild_characters where character_key = $1 and guild_id = $2`, key, *prevGuildID); err != nil {
			return profile, false, false, fmt.Errorf("bnetimport: clear previous guild for %s: %w", key, err)
		}
	}

	if err := guilds.UpsertCharacterMembership(ctx, tx, guilds.MembershipRow{
		GuildID: guildID, CharacterKey: key, UserID: userID,
		RankIndex: rankIndex, Rank: rank, Source: "bnet", VerifiedBy: "bnet",
		VerifiedAt: time.Now(), Reverify: true,
	}); err != nil {
		return profile, false, false, err
	}
	if rank == "leader" {
		if err := guilds.AutoConfirmClaimIfPending(ctx, tx, guildID, userID); err != nil {
			return profile, false, false, err
		}
	}
	if err := guilds.AfterGuildChange(ctx, tx, guildID, userID); err != nil {
		return profile, false, false, err
	}
	if prevGuildID != nil && *prevGuildID != guildID {
		if err := guilds.AfterGuildChange(ctx, tx, *prevGuildID, prevUserID); err != nil {
			return profile, false, false, err
		}
	}
	return profile, true, false, nil
}

// adoptProfileRuleset re-keys ref's row under the ruleset its own profile
// names when that differs from the one it was keyed under: the realm-id path
// in rulesetFor falls back to normal whenever Blizzard cannot list the
// region's realms, as it has not for Forever since 2026-11-05. An empty
// profile ruleset, or a new key another row holds, leaves ref as it was.
func (s *Service) adoptProfileRuleset(ctx context.Context, tx pgx.Tx, ref *characterRef, profileRuleset string) error {
	if profileRuleset == "" || profileRuleset == ref.Ruleset {
		return nil
	}
	newKey := character.Key(ref.Region, profileRuleset, ref.Name)
	_, moved, err := s.moveKey(ctx, tx, ref.Key, newKey, profileRuleset)
	if err != nil || !moved {
		return err
	}
	ref.Key, ref.Ruleset = newKey, profileRuleset
	return nil
}

// findRank locates the character's own roster rank: by Blizzard character
// id when both sides carry one (a realmless Forever roster), otherwise by
// name and realm (case-insensitive), since a guild roster can, rarely, on
// Era, carry two same-named characters on different realms.
func findRank(members []bnetapi.RosterMember, bnetID int64, name, realmSlug string) (int, bool) {
	if bnetID != 0 {
		for _, m := range members {
			if m.ID == bnetID {
				return m.Rank, true
			}
		}
	}
	for _, m := range members {
		if strings.EqualFold(m.Name, name) && strings.EqualFold(m.RealmSlug, realmSlug) {
			return m.Rank, true
		}
	}
	return 0, false
}

// rulesetForRealm maps a character's realm slug to a ruleset via the
// region's realm list (bnetapi.RulesetOf). ok is false for a slug the list
// does not hold or a realm type RulesetOf does not know: nothing is guessed.
func rulesetForRealm(realms []bnetapi.Realm, slug string) (ruleset string, ok bool) {
	for _, r := range realms {
		if strings.EqualFold(r.Slug, slug) {
			return bnetapi.RulesetOf(r)
		}
	}
	return "", false
}

// rulesetSources names, for the needs_ruleset log line, every field
// rulesetFor reads, in order.
const rulesetSources = "realm.slug,learned_realm,directory,protected_character.realm_id,realm_table,profile.ruleset,addon_row,own_row,realm_answers"

// resolvedRuleset is what rulesetFor found: the ruleset, and whether it is the
// player's own answer to the unknown-realm question.
type resolvedRuleset struct {
	Ruleset  string
	Answered bool
}

// rulesetFor resolves an account character's ruleset, the middle segment
// of its key, from the first source that names one:
//
//  1. the realm slug, for a realmed game;
//  2. the realm id's learned ruleset (directory.RealmRuleset: realms the
//     character directory typed from log GUIDs and sign-ins), then the
//     character directory's own row for this character id, whose ruleset its
//     profile or its guild named (a roster lists members who are not public);
//  3. the realm id from the protected-character link, whose realm type
//     (PvE, PvP, RP, RP PvP) is the ruleset: first in the region's realm
//     list, while Blizzard lists one (Forever's has answered 404 since
//     2026-11-05), then in the committed Forever realm table
//     (bnetapi.Client.KnownRealm), which needs no call at all;
//  4. the character profile's own ruleset ("PvP", "Roleplay PvP");
//  5. this account's own addon row of the same name (rulesetFromAddonRow);
//  6. this character's own Battle.net row, when it is keyed by a ruleset:
//     the player's answer to the unknown-realm question, or what an earlier
//     source named (a profile since made private never demotes it);
//  7. the realm's agreed answers (realmConsensus): two or more players, none
//     disagreeing.
//
// ok is false when none does. There is no default: a guessed ruleset keys the
// character, and its guild, wrongly, which is how every sign-up from
// 2026-11-05 to 2026-11-09 landed on normal. The caller keys such a character
// under character.RulesetUnknown and the setup page asks its player.
func (s *Service) rulesetFor(ctx context.Context, userID int64, region string, realms []bnetapi.Realm,
	ch bnetapi.AccountCharacter) (resolved resolvedRuleset, ok bool, err error) {
	if ch.RealmSlug != "" {
		if ruleset, ok := rulesetForRealm(realms, ch.RealmSlug); ok {
			return resolvedRuleset{Ruleset: ruleset}, true, nil
		}
	}
	if ruleset, ok, err := directoryRuleset(ctx, s.Pool, region, ch.RealmID, ch.ID); err != nil || ok {
		return resolvedRuleset{Ruleset: ruleset}, ok, err
	}
	if ruleset, ok := s.rulesetForRealmID(region, realms, ch.RealmID); ok {
		return resolvedRuleset{Ruleset: ruleset}, true, nil
	}
	profile, _, err := s.Client.Character(ctx, region, ch.RealmSlug, ch.Name)
	switch {
	case errors.Is(err, bnetapi.ErrNotFound) || errors.Is(err, bnetapi.ErrForbidden):
	case err != nil:
		return resolvedRuleset{}, false, fmt.Errorf("bnetimport: ruleset from profile of %s: %w", ch.Name, err)
	case profile.Ruleset != "":
		s.noteUnlistedRealm(region, ch.RealmID, profile.Ruleset)
		return resolvedRuleset{Ruleset: profile.Ruleset}, true, nil
	}
	if ruleset, ok, err := s.rulesetFromAddonRow(ctx, userID, region, ch.Name); err != nil || ok {
		return resolvedRuleset{Ruleset: ruleset}, ok, err
	}
	if resolved, ok, err := s.rulesetFromOwnRow(ctx, userID, region, ch.ID); err != nil || ok {
		return resolved, ok, err
	}
	ruleset, ok, err := realmConsensus(ctx, s.Pool, region, ch.RealmID)
	return resolvedRuleset{Ruleset: ruleset}, ok, err
}

// rulesetForRealmID is the ruleset of the realm with this id: from the
// region's realm list when Blizzard serves one, else from the committed
// Forever realm table. ok is false for id 0, an id neither holds, or a realm
// type RulesetOf does not know.
func (s *Service) rulesetForRealmID(region string, realms []bnetapi.Realm, id int64) (string, bool) {
	if id == 0 {
		return "", false
	}
	for _, r := range realms {
		if r.ID == id {
			if ruleset, ok := bnetapi.RulesetOf(r); ok {
				return ruleset, true
			}
		}
	}
	if r, ok := s.Client.KnownRealm(region, id); ok {
		return bnetapi.RulesetOf(r)
	}
	return "", false
}

// noteUnlistedRealm logs, once per realm per process, a realm id the profile
// named a ruleset for but the committed realm table does not hold: the line
// is the evidence the next table row is added from (bnetapi/forever_realms.go).
func (s *Service) noteUnlistedRealm(region string, realmID int64, ruleset string) {
	if realmID == 0 {
		return
	}
	if _, known := s.Client.KnownRealm(region, realmID); known {
		return
	}
	if _, seen := s.unlistedRealms.LoadOrStore(region+"/"+strconv.FormatInt(realmID, 10), true); seen {
		return
	}
	s.logger().Info("bnetimport", "op", "realm_unlisted", "region", region, "realm_id", realmID, "ruleset", ruleset)
}

// rulesetFromAddonRow is the ruleset this account's own addon or companion
// row of the same region and name was keyed under: the addon reads the
// client's realm, so that is evidence. A row the Battle.net import wrote is
// not: it may hold the guess this file no longer makes.
func (s *Service) rulesetFromAddonRow(ctx context.Context, userID int64, region, name string) (string, bool, error) {
	var ruleset string
	err := s.Pool.QueryRow(ctx,
		`select ruleset from characters
		  where region = $1 and user_id = $2 and split_part(key, '/', 3) = $3 and source <> 'bnet'
		  order by key limit 1`,
		region, userID, character.Slug(name)).Scan(&ruleset)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", false, nil
	}
	if err != nil {
		return "", false, fmt.Errorf("bnetimport: ruleset lookup for %s: %w", name, err)
	}
	return ruleset, true, nil
}

// noteStaticNamespace logs the versioned static namespace a profile links
// to, once per region per process: the signal for when Blizzard's static
// data moves to a new client build (bnetapi.StaticNamespace).
func (s *Service) noteStaticNamespace(region string, raw []byte) {
	ns := bnetapi.StaticNamespace(raw)
	if ns == "" {
		return
	}
	if _, seen := s.staticSeen.LoadOrStore(region+"|"+ns, true); !seen {
		s.logger().Info("bnetimport", "op", "static_namespace", "region", region, "namespace", ns)
	}
}
