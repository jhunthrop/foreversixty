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

-- Security hardening (2026-09-21 review response — see the spec's dated
-- amendment blocks in §2.4 and §3.3).

-- Which of the four corroboration paths actually set verified_at, so a
-- claim release/transfer can un-verify precisely the rows the claim
-- itself vouched for and nothing a character separately earned through
-- officer approval, an invite, or log corroboration.
alter table guild_characters add column if not exists verified_by text
  check (verified_by in ('claim', 'officer', 'invite', 'logs'));

-- A claim may be contested while pending or already claimed; a
-- moderator resolves it (uphold, release, or transfer).
alter table guilds add column if not exists claim_contested_at timestamptz;
alter table guilds add column if not exists claim_contested_by bigint references users (id) on delete set null;

-- Rate-limits an account to one claim (successful or pending) per
-- rolling 30 days, across every guild - a small, guilds-owned table
-- rather than a column on auth's own users table.
create table if not exists guild_claim_attempts (
  id           bigserial primary key,
  user_id      bigint not null references users (id) on delete cascade,
  attempted_at timestamptz not null default now()
);
create index if not exists guild_claim_attempts_user_idx on guild_claim_attempts (user_id, attempted_at desc);

-- Guild identity is case-insensitive: one guild per (region, ruleset,
-- lower(name)). The pre-existing exact-text unique(region, ruleset,
-- name) constraint (migration 0005) stays - the new index is strictly
-- stricter and subsumes it, so both coexist harmlessly.
create unique index if not exists guilds_region_ruleset_lower_name_idx
  on guilds (region, ruleset, lower(name));

-- At most one currently-claimed guild per account: converts the
-- application-level claim-rate-limit's TOCTOU race (two concurrent
-- Claim() calls from the same account on two different unclaimed
-- guilds could otherwise both succeed before either's pre-transaction
-- check sees the other) into a safe commit-time unique-violation for
-- the loser (2026-09-21 whole-round review, round 2).
create unique index if not exists guilds_claimed_by_idx on guilds (claimed_by) where claimed_by is not null;
