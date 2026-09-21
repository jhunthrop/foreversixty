-- api/internal/db/migrations/0018_guild_membership.down.sql
-- Third security review response additions (2026-09-21), reversed
-- first, in exact reverse order of the up migration's appended block.
drop index if exists guilds_claim_contested_by_idx;

-- Second security review response additions (2026-09-21), reversed
-- first, in exact reverse order of the up migration's appended block.
drop index if exists guild_claim_resolutions_guild_contester_idx;
drop table if exists guild_claim_resolutions;
alter table guild_claim_attempts drop column if exists kind;
alter table guilds drop column if exists claimed_at;

-- Security hardening additions (2026-09-21), reversed first, in exact
-- reverse order of the up migration's appended block.
drop index if exists guilds_claimed_by_idx;
drop index if exists guilds_region_ruleset_lower_name_idx;
drop table if exists guild_claim_attempts;
alter table guilds drop column if exists claim_contested_by;
alter table guilds drop column if exists claim_contested_at;
alter table guild_characters drop column if exists verified_by;

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
