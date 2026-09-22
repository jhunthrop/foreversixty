-- api/internal/db/migrations/0024_bnet_characters.up.sql
-- Battle.net character import, guild discovery (spec
-- docs/superpowers/specs/2026-09-22-battlenet-character-import-design.md §2).

-- characters gains the fields Blizzard's profile API can tell us that an
-- addon export cannot: the realm the character actually lives on (kept
-- for collision diagnostics even though the key stays realm-less), level,
-- faction, Blizzard's own numeric character id (for the nightly refresh's
-- WHERE clause), which path wrote the row, and when it was imported.
alter table characters
  add column if not exists realm_slug text,
  add column if not exists realm_name text,
  add column if not exists level int,
  add column if not exists faction text check (faction in ('alliance', 'horde')),
  add column if not exists bnet_character_id bigint,
  add column if not exists source text not null default 'export'
    check (source in ('export', 'bnet')),
  add column if not exists imported_at timestamptz;
create index if not exists characters_bnet_idx on characters (bnet_character_id) where bnet_character_id is not null;

-- guild_characters.source and .verified_by (both from 0018) gain a third
-- corroboration path: a Blizzard guild roster.
alter table guild_characters drop constraint if exists guild_characters_source_check;
alter table guild_characters add constraint guild_characters_source_check
  check (source in ('export', 'invite', 'claim', 'bnet'));
alter table guild_characters drop constraint if exists guild_characters_verified_by_check;
alter table guild_characters add constraint guild_characters_verified_by_check
  check (verified_by in ('claim', 'officer', 'invite', 'logs', 'bnet'));

-- guilds gains Blizzard's own guild id (so the nightly refresh can key off
-- it without re-resolving the name), the realm the roster was read from,
-- and when that roster was last fetched.
alter table guilds
  add column if not exists bnet_guild_id bigint,
  add column if not exists realm_slug text,
  add column if not exists roster_refreshed_at timestamptz;

-- users.bnet_imported_at is GET /v1/me's top-level "when did the last
-- Battle.net import happen" fact (spec §6); null means never imported.
alter table users add column if not exists bnet_imported_at timestamptz;
