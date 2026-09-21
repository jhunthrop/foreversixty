-- api/internal/db/migrations/0018_guild_membership.down.sql
alter table guilds drop column if exists invite_token_rotated_at;
alter table guilds drop column if exists invite_token_hash;
alter table guilds drop column if exists claim_requested_at;
alter table guilds drop column if exists claim_pending_by;
alter table guilds drop column if exists officer_max_rank_index;

drop index if exists guild_members_user_idx;
alter table guild_members drop column if exists verified_at;
alter table guild_members drop column if exists consent;

drop index if exists guild_characters_unverified_idx;
drop index if exists guild_characters_user_idx;
drop index if exists guild_characters_character_idx;
drop table if exists guild_characters;
