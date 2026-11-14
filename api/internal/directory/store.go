// api/internal/directory/store.go
package directory

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
)

// dirGuild is one guild the directory sweeps.
type dirGuild struct {
	Region  string
	Slug    string
	BnetID  int64
	Name    string
	Ruleset string
}

// dueRow is one character a pass reads: by id when the directory holds it, by slug alone
// (ID 0) for a name seed.
type dueRow struct {
	Region       string
	ID           int64
	Slug         string
	Sources      []string
	RosterRecent bool
}

// profileRead is what one character's public endpoints answered.
type profileRead struct {
	Profile     bnetapi.CharacterProfile
	Equipment   []EquippedItem
	Spec        string
	Professions []Profession
	AvatarURL   string
}

// mergeSources is the SQL that unions a row's sources with an insert's, sorted.
const mergeSources = `array(select distinct s from unnest(character_directory.sources || excluded.sources) s order by s)`

// upsertRoster writes every member of one roster read and clears the guild from rows of
// members no longer on it, in one transaction. A member's ruleset is the guild's unless its
// own profile named one.
func (s *Service) upsertRoster(ctx context.Context, g dirGuild, roster bnetapi.Roster) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("directory: begin roster %s/%s: %w", g.Region, g.Slug, err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	batch := &pgx.Batch{}
	ids := make([]int64, 0, len(roster.Members))
	for _, m := range roster.Members {
		if m.ID == 0 || m.Name == "" {
			continue
		}
		ids = append(ids, m.ID)
		batch.Queue(`
			insert into character_directory (region, bnet_character_id, name, slug, class, race_id, faction, level,
			    guild_bnet_id, guild_name, guild_rank, ruleset, ruleset_source, sources, roster_seen_at, refreshed_at)
			select $1, $2, $3, $4, nullif($5, ''), nullif($6::int, 0), nullif($7, ''), $8, $9, $10, $11,
			       nullif($12, ''), case when $12 = '' then null else 'guild' end, array[$13], now(), now()
			on conflict (region, bnet_character_id) do update set
			  name = excluded.name, slug = excluded.slug,
			  class = coalesce(excluded.class, character_directory.class),
			  race_id = coalesce(excluded.race_id, character_directory.race_id),
			  faction = coalesce(excluded.faction, character_directory.faction),
			  level = excluded.level, guild_bnet_id = excluded.guild_bnet_id, guild_name = excluded.guild_name,
			  guild_rank = excluded.guild_rank,
			  ruleset = case when character_directory.ruleset_source = 'profile' or excluded.ruleset is null
			                 then character_directory.ruleset else excluded.ruleset end,
			  ruleset_source = case when character_directory.ruleset_source = 'profile' or excluded.ruleset is null
			                 then character_directory.ruleset_source else 'guild' end,
			  sources = `+mergeSources+`, roster_seen_at = now(), refreshed_at = now()`,
			g.Region, m.ID, m.Name, bnetapi.CharacterSlug(m.Name), bnetapi.ClassSlugByID(m.ClassID), m.RaceID,
			m.Faction, m.Level, roster.GuildID, roster.GuildName, m.Rank, g.Ruleset, SourceRoster)
	}
	if err := tx.SendBatch(ctx, batch).Close(); err != nil {
		return fmt.Errorf("directory: roster members %s/%s: %w", g.Region, g.Slug, err)
	}
	if _, err := tx.Exec(ctx,
		`update character_directory set guild_bnet_id = null, guild_name = null, guild_rank = null
		  where region = $1 and guild_bnet_id = $2 and not (bnet_character_id = any($3))`,
		g.Region, roster.GuildID, ids); err != nil {
		return fmt.Errorf("directory: roster leavers %s/%s: %w", g.Region, g.Slug, err)
	}
	if _, err := tx.Exec(ctx,
		`update directory_guilds set bnet_guild_id = $3, name = $4, member_count = $5,
		        roster_refreshed_at = now(), missing_since = null
		  where region = $1 and slug = $2`,
		g.Region, g.Slug, roster.GuildID, roster.GuildName, len(ids)); err != nil {
		return fmt.Errorf("directory: stamp guild %s/%s: %w", g.Region, g.Slug, err)
	}
	return tx.Commit(ctx)
}

// stampGuild records a guild summary read: its id, name, faction and ruleset.
func (s *Service) stampGuild(ctx context.Context, g dirGuild, summary bnetapi.Guild) error {
	_, err := s.Pool.Exec(ctx,
		`update directory_guilds set bnet_guild_id = $3, name = $4, faction = nullif($5, ''), ruleset = nullif($6, '')
		  where region = $1 and slug = $2`,
		g.Region, g.Slug, summary.ID, summary.Name, summary.Faction, summary.Ruleset)
	if err != nil {
		return fmt.Errorf("directory: guild summary %s/%s: %w", g.Region, g.Slug, err)
	}
	return nil
}

// guildMissing records a guild that answered 404, keeping the first 404's time.
func (s *Service) guildMissing(ctx context.Context, g dirGuild) error {
	_, err := s.Pool.Exec(ctx,
		`update directory_guilds set roster_refreshed_at = now(), missing_since = coalesce(missing_since, now())
		  where region = $1 and slug = $2`, g.Region, g.Slug)
	if err != nil {
		return fmt.Errorf("directory: guild missing %s/%s: %w", g.Region, g.Slug, err)
	}
	return nil
}

// dueGuilds is up to limit guilds whose roster is older than rosterAfter, oldest first.
func (s *Service) dueGuilds(ctx context.Context, limit int) ([]dirGuild, error) {
	rows, err := s.Pool.Query(ctx,
		`select region, slug, coalesce(bnet_guild_id, 0), coalesce(name, ''), coalesce(ruleset, '')
		   from directory_guilds
		  where region = any($1) and (roster_refreshed_at is null or roster_refreshed_at < now() - $2::interval)
		  order by roster_refreshed_at nulls first, slug limit $3`,
		s.Regions, interval(rosterAfter), limit)
	if err != nil {
		return nil, fmt.Errorf("directory: due guilds: %w", err)
	}
	return pgx.CollectRows(rows, func(r pgx.CollectableRow) (dirGuild, error) {
		var g dirGuild
		err := r.Scan(&g.Region, &g.Slug, &g.BnetID, &g.Name, &g.Ruleset)
		return g, err
	})
}

// dueCharacters is up to limit rows whose profile read is older than profileAfter, never
// read first, then oldest. RosterRecent marks a row a roster confirmed within staleAfter.
func (s *Service) dueCharacters(ctx context.Context, limit int) ([]dueRow, error) {
	rows, err := s.Pool.Query(ctx,
		`select region, bnet_character_id, slug, sources, coalesce(roster_seen_at > now() - $3::interval, false)
		   from character_directory
		  where region = any($1) and (profile_refreshed_at is null or profile_refreshed_at < now() - $2::interval)
		  order by profile_refreshed_at nulls first, refreshed_at limit $4`,
		s.Regions, interval(profileAfter), interval(staleAfter), limit)
	if err != nil {
		return nil, fmt.Errorf("directory: due characters: %w", err)
	}
	return pgx.CollectRows(rows, func(r pgx.CollectableRow) (dueRow, error) {
		var d dueRow
		err := r.Scan(&d.Region, &d.ID, &d.Slug, &d.Sources, &d.RosterRecent)
		return d, err
	})
}

// saveProfile writes one character's public read, inserting the row when a name seed found
// it, and adds its guild to the sweep.
func (s *Service) saveProfile(ctx context.Context, region string, sources []string, read profileRead) error {
	p := read.Profile
	equipment, err := jsonOrNil(read.Equipment)
	if err != nil {
		return err
	}
	professions, err := jsonOrNil(read.Professions)
	if err != nil {
		return err
	}
	var lastLogin *time.Time
	if p.LastLoginTimestamp > 0 {
		t := time.UnixMilli(p.LastLoginTimestamp).UTC()
		lastLogin = &t
	}
	itemLevel := p.EquippedItemLevel
	if itemLevel == nil {
		itemLevel = p.AverageItemLevel
	}
	var guildID *int64
	var guildName *string
	if p.HasGuild && p.GuildID != 0 {
		guildID, guildName = &p.GuildID, &p.GuildName
	}
	_, err = s.Pool.Exec(ctx, `
		insert into character_directory (region, bnet_character_id, name, slug, class, race, faction, level,
		    guild_bnet_id, guild_name, ruleset, ruleset_source, item_level, spec, last_login_at, equipment,
		    professions, avatar_url, public, sources, profile_refreshed_at, refreshed_at)
		select $1, $2, $3, $4, nullif($5, ''), nullif($6, ''), nullif($7, ''), $8, $9, $10, nullif($11, ''),
		       case when $11 = '' then null else 'profile' end, $12, nullif($13, ''), $14, $15, $16, nullif($17, ''),
		       true, $18, now(), now()
		on conflict (region, bnet_character_id) do update set
		  name = excluded.name, slug = excluded.slug, class = coalesce(excluded.class, character_directory.class),
		  race = coalesce(excluded.race, character_directory.race),
		  faction = coalesce(excluded.faction, character_directory.faction), level = excluded.level,
		  guild_bnet_id = excluded.guild_bnet_id, guild_name = excluded.guild_name,
		  guild_rank = case when excluded.guild_bnet_id is distinct from character_directory.guild_bnet_id
		                    then null else character_directory.guild_rank end,
		  ruleset = coalesce(excluded.ruleset, character_directory.ruleset),
		  ruleset_source = coalesce(excluded.ruleset_source, character_directory.ruleset_source),
		  item_level = excluded.item_level, spec = excluded.spec, last_login_at = excluded.last_login_at,
		  equipment = excluded.equipment, professions = excluded.professions, avatar_url = excluded.avatar_url,
		  public = true, sources = `+mergeSources+`, profile_refreshed_at = now(), refreshed_at = now()`,
		region, p.ID, p.Name, bnetapi.CharacterSlug(p.Name), p.ClassSlug, p.RaceName, p.Faction, p.Level,
		guildID, guildName, p.Ruleset, itemLevel, read.Spec, lastLogin, equipment, professions, read.AvatarURL, sources)
	if err != nil {
		return fmt.Errorf("directory: save profile %s/%d: %w", region, p.ID, err)
	}
	if p.HasGuild && p.GuildSlug != "" {
		return s.addGuild(ctx, region, p.GuildSlug, p.GuildID, p.GuildName)
	}
	return nil
}

// addGuild adds a guild to the sweep; one already there is left as it is.
func (s *Service) addGuild(ctx context.Context, region, slug string, bnetID int64, name string) error {
	_, err := s.Pool.Exec(ctx,
		`insert into directory_guilds (region, slug, bnet_guild_id, name)
		 values ($1, $2, nullif($3::bigint, 0), nullif($4, ''))
		 on conflict do nothing`, region, slug, bnetID, name)
	if err != nil {
		return fmt.Errorf("directory: add guild %s/%s: %w", region, slug, err)
	}
	return nil
}

// markPrivate records a character whose profile is not public (any more): the profile's own
// fields are dropped, since they can no longer be refreshed, and the roster keeps the row.
func (s *Service) markPrivate(ctx context.Context, region string, id int64) error {
	_, err := s.Pool.Exec(ctx,
		`update character_directory set public = false, item_level = null, spec = null, last_login_at = null,
		        equipment = null, professions = null, avatar_url = null, race = null,
		        ruleset = case when ruleset_source = 'profile' then null else ruleset end,
		        ruleset_source = case when ruleset_source = 'profile' then null else ruleset_source end,
		        profile_refreshed_at = now()
		  where region = $1 and bnet_character_id = $2`, region, id)
	if err != nil {
		return fmt.Errorf("directory: mark private %s/%d: %w", region, id, err)
	}
	return nil
}

// deleteCharacter removes one character from the directory.
func (s *Service) deleteCharacter(ctx context.Context, region string, id int64) error {
	if _, err := s.Pool.Exec(ctx,
		`delete from character_directory where region = $1 and bnet_character_id = $2`, region, id); err != nil {
		return fmt.Errorf("directory: delete %s/%d: %w", region, id, err)
	}
	return nil
}

// retryLater puts a row whose read failed back in the queue for retryAfter, so one broken
// answer never holds up the rows behind it. A name seed (ID 0) is retried by its seed.
func (s *Service) retryLater(ctx context.Context, row dueRow) error {
	if row.ID == 0 {
		return nil
	}
	_, err := s.Pool.Exec(ctx,
		`update character_directory set profile_refreshed_at = now() - $3::interval + $4::interval
		  where region = $1 and bnet_character_id = $2`,
		row.Region, row.ID, interval(profileAfter), interval(retryAfter))
	if err != nil {
		return fmt.Errorf("directory: retry later %s/%d: %w", row.Region, row.ID, err)
	}
	return nil
}

// recordMiss remembers a name that answered 404.
func (s *Service) recordMiss(ctx context.Context, region, slug string) error {
	_, err := s.Pool.Exec(ctx,
		`insert into directory_name_misses (region, slug) values ($1, $2)
		 on conflict (region, slug) do update set checked_at = now()`, region, slug)
	if err != nil {
		return fmt.Errorf("directory: record miss %s/%s: %w", region, slug, err)
	}
	return nil
}

// expire deletes every row no source confirmed within staleAfter, guilds missing that long,
// and name misses past missAfter. Returns how many characters it deleted.
func (s *Service) expire(ctx context.Context) (int64, error) {
	tag, err := s.Pool.Exec(ctx,
		`delete from character_directory where refreshed_at < now() - $1::interval`, interval(staleAfter))
	if err != nil {
		return 0, fmt.Errorf("directory: expire characters: %w", err)
	}
	if _, err := s.Pool.Exec(ctx,
		`delete from directory_guilds where missing_since < now() - $1::interval`, interval(staleAfter)); err != nil {
		return 0, fmt.Errorf("directory: expire guilds: %w", err)
	}
	if _, err := s.Pool.Exec(ctx,
		`delete from directory_name_misses where checked_at < now() - $1::interval`, interval(missAfter)); err != nil {
		return 0, fmt.Errorf("directory: expire misses: %w", err)
	}
	return tag.RowsAffected(), nil
}

func jsonOrNil[T any](v []T) ([]byte, error) {
	if v == nil {
		return nil, nil
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("directory: encode: %w", err)
	}
	return b, nil
}

// interval is d as a Postgres interval literal.
func interval(d time.Duration) string { return fmt.Sprintf("%d seconds", int64(d.Seconds())) }
