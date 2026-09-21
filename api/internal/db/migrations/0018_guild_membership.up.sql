-- api/internal/db/migrations/0018_guild_membership.up.sql
-- One row per character currently reporting membership in a guild. rank_index is the raw
-- GetGuildInfo value (kept so an officer-threshold change can re-derive rank with no addon
-- round trip); rank is the derived label; source records how the row was made;
-- verified_at is null until corroborated.
create table if not exists guild_characters (
  guild_id     bigint not null references guilds (id) on delete cascade,
  character_key text not null,
  user_id      bigint not null references users (id) on delete cascade,
  rank_index   smallint,
  rank         text not null default 'member' check (rank in ('member', 'officer', 'leader')),
  source       text not null default 'export' check (source in ('export', 'invite', 'claim')),
  verified_at  timestamptz,
  refreshed_at timestamptz not null default now(),
  primary key (guild_id, character_key)
);
-- A character belongs to at most one guild at a time - this is both that invariant and
-- the index PutExports uses to find a character's *previous* guild row.
create unique index if not exists guild_characters_character_idx on guild_characters (character_key);
create index if not exists guild_characters_user_idx on guild_characters (guild_id, user_id);
create index if not exists guild_characters_unverified_idx on guild_characters (guild_id)
  where verified_at is null;

-- guild_members stays the account-level access row every existing GuildRank/mayView/
-- mayEdit call site reads; it gains per-account-per-guild consent and the derived
-- verified_at recomputeMembership writes. Its rank column already exists (migration
-- 0005); this spec stops writing it directly and starts deriving it.
alter table guild_members add column if not exists consent text not null default 'gear'
  check (consent in ('roster', 'gear', 'gear_bags'));
alter table guild_members add column if not exists verified_at timestamptz;
create index if not exists guild_members_user_idx on guild_members (user_id);

-- guilds gains: the officer-rank threshold (a claimed guild's own setting), the pending
-- half of the two-step claim, and the invite link (hashed like login_tokens and devices -
-- never stored in the clear).
alter table guilds add column if not exists officer_max_rank_index integer not null default 1;
alter table guilds add column if not exists claim_pending_by bigint references users (id) on delete set null;
alter table guilds add column if not exists claim_requested_at timestamptz;
alter table guilds add column if not exists invite_token_hash bytea;
alter table guilds add column if not exists invite_token_rotated_at timestamptz;
