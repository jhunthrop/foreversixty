// api/internal/directory/realms.go
package directory

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Querier is what both a pool and a transaction answer.
type Querier interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

// RealmRuleset is the learned ruleset of a realm id: ok only while every piece of evidence
// for the realm names the same ruleset.
func RealmRuleset(ctx context.Context, q Querier, region string, realmID int64) (string, bool, error) {
	if realmID == 0 {
		return "", false, nil
	}
	var ruleset string
	err := q.QueryRow(ctx,
		`select min(ruleset) from realm_rulesets where region = $1 and realm_id = $2
		  having count(*) = 1`, region, realmID).Scan(&ruleset)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", false, nil
	}
	if err != nil {
		return "", false, fmt.Errorf("directory: realm ruleset %s/%d: %w", region, realmID, err)
	}
	return ruleset, true, nil
}

// CharacterRuleset is the directory's ruleset for a Blizzard character id: from its own
// profile, its guild (a roster lists members who are not public) or its realm.
func CharacterRuleset(ctx context.Context, q Querier, region string, bnetID int64) (string, bool, error) {
	if bnetID == 0 {
		return "", false, nil
	}
	var ruleset string
	err := q.QueryRow(ctx,
		`select ruleset from character_directory
		  where region = $1 and bnet_character_id = $2 and ruleset is not null`, region, bnetID).Scan(&ruleset)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", false, nil
	}
	if err != nil {
		return "", false, fmt.Errorf("directory: character ruleset %s/%d: %w", region, bnetID, err)
	}
	return ruleset, true, nil
}

// realmEvidence is every (region, realm id, character id) the site has seen: combat-log
// GUIDs (the report's region from its guild, else from its owner's characters when they are
// all in one region) and sign-ins' protected-character realm ids.
const realmEvidence = `
	with sightings as (
	  select coalesce(g.region, owner.region) as region,
	         split_part(p, '-', 2)::bigint as realm_id,
	         ('x' || lpad(split_part(p, '-', 3), 16, '0'))::bit(64)::bigint as character_id
	    from fights f
	    join reports r on r.id = f.report_id
	    left join guilds g on g.id = r.guild_id
	    left join lateral (
	      select min(c.region) as region from characters c where c.user_id = r.owner_id
	      having count(distinct c.region) = 1
	    ) owner on true
	    cross join unnest(f.players) as p
	   where p ~ '^Player-[0-9]{1,10}-[0-9A-Fa-f]{1,15}$'
	  union
	  select region, bnet_realm_id, bnet_character_id from characters
	   where bnet_realm_id is not null and bnet_character_id is not null
	)`

// learnRealms matches every sighting to a directory character whose ruleset its profile or
// guild named, stamps the character's realm id, raises each realm's evidence, and gives the
// realm's ruleset to directory characters that have none. No Blizzard call.
func (s *Service) learnRealms(ctx context.Context) (int64, error) {
	if _, err := s.Pool.Exec(ctx, realmEvidence+`
		update character_directory d set realm_id = e.realm_id
		  from sightings e
		 where e.region = d.region and e.character_id = d.bnet_character_id
		   and d.realm_id is distinct from e.realm_id`); err != nil {
		return 0, fmt.Errorf("directory: stamp realm ids: %w", err)
	}
	if _, err := s.Pool.Exec(ctx, `
		insert into realm_rulesets (region, realm_id, ruleset, evidence)
		select region, realm_id, ruleset, count(*) from character_directory
		 where realm_id is not null and ruleset is not null and ruleset_source in ('profile', 'guild')
		 group by region, realm_id, ruleset
		on conflict (region, realm_id, ruleset) do update set
		  evidence = greatest(realm_rulesets.evidence, excluded.evidence),
		  updated_at = case when excluded.evidence > realm_rulesets.evidence then now() else realm_rulesets.updated_at end`); err != nil {
		return 0, fmt.Errorf("directory: learn realms: %w", err)
	}
	tag, err := s.Pool.Exec(ctx, `
		update character_directory d set ruleset = r.ruleset, ruleset_source = 'realm'
		  from (select region, realm_id, min(ruleset) as ruleset from realm_rulesets
		         group by region, realm_id having count(*) = 1) r
		 where d.ruleset is null and d.region = r.region and d.realm_id = r.realm_id`)
	if err != nil {
		return 0, fmt.Errorf("directory: rulesets from realms: %w", err)
	}
	return tag.RowsAffected(), nil
}
