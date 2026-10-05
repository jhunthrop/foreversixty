-- api/internal/db/migrations/0031_guild_faction.down.sql
alter table guilds drop column if exists faction_updated_at;
alter table guilds drop column if exists faction;
