-- api/internal/db/migrations/0024_bnet_characters.down.sql
alter table users drop column if exists bnet_imported_at;

alter table guilds drop column if exists roster_refreshed_at;
alter table guilds drop column if exists realm_slug;
alter table guilds drop column if exists bnet_guild_id;

alter table guild_characters drop constraint if exists guild_characters_verified_by_check;
alter table guild_characters add constraint guild_characters_verified_by_check
  check (verified_by in ('claim', 'officer', 'invite', 'logs'));
alter table guild_characters drop constraint if exists guild_characters_source_check;
alter table guild_characters add constraint guild_characters_source_check
  check (source in ('export', 'invite', 'claim'));

drop index if exists characters_bnet_idx;
alter table characters drop column if exists imported_at;
alter table characters drop column if exists source;
alter table characters drop column if exists bnet_character_id;
alter table characters drop column if exists faction;
alter table characters drop column if exists level;
alter table characters drop column if exists realm_name;
alter table characters drop column if exists realm_slug;
