-- api/internal/db/migrations/0070_guild_invite_shown.down.sql
alter table guilds
  drop column if exists invite_rotated_by,
  drop column if exists invite_token_enc;
