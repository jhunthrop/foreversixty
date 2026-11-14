-- Character directory (docs/superpowers/specs/2026-11-14-character-directory-discovery.md §6):
-- every Forever character the site can learn about before its player signs in, keyed by
-- region and Blizzard character id. Public game data only: no BattleTag, no account profile,
-- nothing behind a user token. Filled from guild rosters (directory_guilds), sign-ins, addon
-- exports and logged names; enriched from the public profile endpoints. Blizzard's API terms
-- (section 2(s)) bound it: a row is validated through /status and refreshed, or deleted,
-- within 30 days (refreshed_at is the last time any source confirmed it).
create table character_directory (
  region               text    not null check (region in ('us', 'eu')),
  bnet_character_id    bigint  not null,
  name                 text    not null,
  slug                 text    not null,
  class                text,
  race                 text,
  race_id              int,
  faction              text    check (faction in ('alliance', 'horde')),
  level                int,
  guild_bnet_id        bigint,
  guild_name           text,
  guild_rank           int,
  ruleset              text    check (ruleset in ('normal', 'pvp', 'rp', 'hardcore')),
  -- Where ruleset came from, strongest first: the character's own profile, its guild, its realm.
  ruleset_source       text    check (ruleset_source in ('profile', 'guild', 'realm')),
  realm_id             bigint,
  item_level           int,
  spec                 text,
  last_login_at        timestamptz,
  equipment            jsonb,
  professions          jsonb,
  avatar_url           text,
  -- public is true once the profile answered; a roster lists members who are not public yet.
  public               boolean not null default false,
  sources              text[]  not null default '{}',
  roster_seen_at       timestamptz,
  profile_refreshed_at timestamptz,
  refreshed_at         timestamptz not null default now(),
  created_at           timestamptz not null default now(),
  primary key (region, bnet_character_id)
);
create index character_directory_guild_idx on character_directory (region, guild_bnet_id) where guild_bnet_id is not null;
create index character_directory_slug_idx on character_directory (region, slug);
create index character_directory_due_idx on character_directory (profile_refreshed_at nulls first);
create index character_directory_refreshed_idx on character_directory (refreshed_at);

-- The guilds the directory sweeps, by the realmless slug Blizzard answers for Forever. Seeded
-- from the guilds table and from every guild a directory profile names; one roster read a day.
create table directory_guilds (
  region              text   not null check (region in ('us', 'eu')),
  slug                text   not null,
  bnet_guild_id       bigint,
  name                text,
  faction             text   check (faction in ('alliance', 'horde')),
  ruleset             text   check (ruleset in ('normal', 'pvp', 'rp', 'hardcore')),
  member_count        int,
  roster_refreshed_at timestamptz,
  -- missing_since is the first 404 of a run of them: a renamed or disbanded guild, or a stored
  -- name that does not slug the way Blizzard's does. Dropped after 30 days of 404s.
  missing_since       timestamptz,
  created_at          timestamptz not null default now(),
  primary key (region, slug)
);
create index directory_guilds_due_idx on directory_guilds (roster_refreshed_at nulls first);
create unique index directory_guilds_bnet_idx on directory_guilds (region, bnet_guild_id) where bnet_guild_id is not null;

-- Names a seed asked Blizzard about that answered 404, so a pass does not ask again for a week.
create table directory_name_misses (
  region     text not null,
  slug       text not null,
  checked_at timestamptz not null default now(),
  primary key (region, slug)
);
