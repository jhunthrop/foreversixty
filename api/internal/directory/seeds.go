// api/internal/directory/seeds.go
package directory

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
)

// seed adds what the site already knows to the sweep, with no Blizzard call: every guild in
// the guilds table, and every signed-in character by its Blizzard id.
func (s *Service) seed(ctx context.Context) error {
	if err := s.seedGuilds(ctx); err != nil {
		return err
	}
	return s.seedSignIns(ctx)
}

// seedGuilds adds every guild of the guilds table, by Blizzard's slug of its name.
func (s *Service) seedGuilds(ctx context.Context) error {
	rows, err := s.Pool.Query(ctx,
		`select distinct region, name, coalesce(bnet_guild_id, 0) from guilds where region = any($1)`, s.Regions)
	if err != nil {
		return fmt.Errorf("directory: seed guilds: %w", err)
	}
	type known struct {
		region, name string
		bnetID       int64
	}
	guilds, err := pgx.CollectRows(rows, func(r pgx.CollectableRow) (known, error) {
		var k known
		err := r.Scan(&k.region, &k.name, &k.bnetID)
		return k, err
	})
	if err != nil {
		return fmt.Errorf("directory: seed guilds: %w", err)
	}
	for _, g := range guilds {
		if slug := bnetapi.GuildSlug(g.name); slug != "" {
			if err := s.addGuild(ctx, g.region, slug, g.bnetID, g.name); err != nil {
				return err
			}
		}
	}
	return nil
}

// seedSignIns adds every character a Battle.net sign-in imported, by its Blizzard id. A name
// that answered 404 lately is left out until its miss expires.
func (s *Service) seedSignIns(ctx context.Context) error {
	_, err := s.Pool.Exec(ctx, `
		insert into character_directory (region, bnet_character_id, name, slug, class, level, faction, sources)
		select distinct on (c.region, c.bnet_character_id)
		       c.region, c.bnet_character_id, c.name, split_part(c.key, '/', 3), c.class, c.level, c.faction, array[$2]
		  from characters c
		 where c.bnet_character_id is not null and c.region = any($1)
		   and not exists (select 1 from directory_name_misses m where m.region = c.region and m.slug = split_part(c.key, '/', 3))
		   and not exists (select 1 from directory_optouts o where o.region = c.region and o.bnet_character_id = c.bnet_character_id)
		 order by c.region, c.bnet_character_id, c.refreshed_at desc
		on conflict (region, bnet_character_id) do update set sources = `+mergeSources+`
		 where not ($2 = any(character_directory.sources))`,
		s.Regions, SourceSignIn)
	if err != nil {
		return fmt.Errorf("directory: seed sign-ins: %w", err)
	}
	return nil
}

// seedNames looks up, by name, characters the site knows only by name: addon rows without a
// Blizzard id, and players in the last month's logs and ratings. A name that answers is
// added with its full read; one that answers 404 is remembered as a miss for a week.
func (s *Service) seedNames(ctx context.Context, res *PassResult) error {
	rows, err := s.Pool.Query(ctx, `
		with names as (
		  select c.region, split_part(c.key, '/', 3) as slug, $3::text as src
		    from characters c where c.bnet_character_id is null
		  union
		  select split_part(m.player_key, '/', 1), split_part(m.player_key, '/', 3), $4
		    from fight_metrics m where m.fought_at > now() - $6::interval and m.player_key ~ '^[a-z]+/[a-z]+/[^/]+$'
		  union
		  select split_part(r.player_key, '/', 1), split_part(r.player_key, '/', 3), $5
		    from rating_scores r where r.fought_at > now() - $6::interval and r.player_key ~ '^[a-z]+/[a-z]+/[^/]+$'
		)
		select n.region, n.slug, array_agg(distinct n.src order by n.src)
		  from names n
		 where n.region = any($1) and n.slug <> ''
		   and not exists (select 1 from character_directory d where d.region = n.region and d.slug = n.slug)
		   and not exists (select 1 from directory_name_misses x where x.region = n.region and x.slug = n.slug)
		 group by n.region, n.slug order by n.region, n.slug limit $2`,
		s.Regions, nameSeedsPerPass, SourceAddon, SourceLog, SourceRating, interval(staleAfter))
	if err != nil {
		return fmt.Errorf("directory: seed names: %w", err)
	}
	seeds, err := pgx.CollectRows(rows, func(r pgx.CollectableRow) (dueRow, error) {
		var d dueRow
		err := r.Scan(&d.Region, &d.Slug, &d.Sources)
		return d, err
	})
	if err != nil || len(seeds) == 0 {
		return err
	}
	return s.readAll(ctx, seeds, res)
}
