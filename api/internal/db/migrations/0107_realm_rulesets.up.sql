-- Learned realm rulesets (docs/superpowers/specs/2026-11-14-character-directory-discovery.md §2):
-- Forever realms have no names and Blizzard lists none, so a realm id's ruleset is learned
-- from characters whose ruleset is known without a realm: their own public profile or their
-- guild (character_directory.ruleset_source 'profile' or 'guild'), matched by region and
-- Blizzard character id to a realm id seen in a combat-log GUID (Player-{realmId}-{hex}, the
-- hex being the character id) or a sign-in's protected-character link. One row per ruleset a
-- realm's evidence names; a realm resolves only while exactly one ruleset has evidence.
create table realm_rulesets (
  region     text   not null check (region in ('us', 'eu')),
  realm_id   bigint not null,
  ruleset    text   not null check (ruleset in ('normal', 'pvp', 'rp', 'hardcore')),
  -- evidence is the most distinct characters any pass matched for this realm and ruleset.
  evidence   int    not null,
  updated_at timestamptz not null default now(),
  primary key (region, realm_id, ruleset)
);
