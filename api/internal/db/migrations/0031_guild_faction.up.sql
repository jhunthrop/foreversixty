-- api/internal/db/migrations/0031_guild_faction.up.sql
-- Stores each guild's majority faction (guilds.RecomputeFaction, api/internal/guilds/
-- faction.go) so the guild page header can paint the faction emblem and watermark from
-- one read, rather than decoding every roster character's FS1 export per request
-- (design/specs/2026-10-04-guild-page.md §12.2). Nullable: a brand-new guild with no
-- exports yet, or one whose roster ties or carries no countable race, has no faction to
-- show rather than a guessed one.
alter table guilds add column if not exists faction text check (faction in ('alliance', 'horde'));
alter table guilds add column if not exists faction_updated_at timestamptz;
