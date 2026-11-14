-- Directory opt-out (docs/superpowers/specs/2026-11-14-character-directory-discovery.md §5,
-- Blizzard API terms 11(c)): a player who asks is deleted from the character directory and
-- never stored again. directory_optouts holds the region and Blizzard character id only.
create table directory_optouts (
  region            text   not null check (region in ('us', 'eu')),
  bnet_character_id bigint not null,
  method            text   not null check (method in ('battlenet', 'note')),
  created_at        timestamptz not null default now(),
  primary key (region, bnet_character_id)
);

-- A code the "Remove my character" page gave for a character, until the character's public
-- guild note carries it (read by a guildmate's addon) or it expires. slug is kept only so the
-- page can tell the asker how it went; the row goes 30 days after it expires.
create table directory_optout_codes (
  region            text   not null check (region in ('us', 'eu')),
  bnet_character_id bigint not null,
  slug              text   not null,
  code              text   not null unique,
  created_at        timestamptz not null default now(),
  expires_at        timestamptz not null,
  verified_at       timestamptz,
  primary key (region, bnet_character_id)
);
create index directory_optout_codes_slug_idx on directory_optout_codes (region, slug);
