-- api/internal/db/migrations/0032_guild_crest.up.sql
-- A guild's uploaded crest (docs/contracts/2026-10-05-guild-crest-api.md): the R2 object
-- key the server produced it under, and when it was last set. Both nullable - a guild with
-- no crest falls back to its faction logo, never a guessed or placeholder image.
alter table guilds add column if not exists crest_key text;
alter table guilds add column if not exists crest_updated_at timestamptz;
