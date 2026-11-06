-- api/internal/db/migrations/0075_guild_recruitment.down.sql
drop table if exists guild_member_notes;
drop table if exists guild_application_votes;
drop table if exists guild_applications;
alter table guilds
  drop column if exists recruit_updated_at,
  drop column if exists recruit_discord_role_id,
  drop column if exists recruit_wanted,
  drop column if exists recruit_nights,
  drop column if exists recruit_pitch,
  drop column if exists recruit_open;
