-- api/internal/db/migrations/0032_guild_crest.down.sql
alter table guilds drop column if exists crest_updated_at;
alter table guilds drop column if exists crest_key;
