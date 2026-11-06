-- The role-aware guild Overview (design/specs/2026-11-05-web-layout.md §7.1, §7.4).
--
-- guild_raid_comp_targets is an officer's comp target for one raid (D13): tanks, healers,
-- melee and ranged. A raid with no row reads its zone's default (Go: raids_comp_target.go).
create table if not exists guild_raid_comp_targets (
  raid_id    bigint primary key references guild_raids (id) on delete cascade,
  tanks      smallint not null check (tanks >= 0),
  healers    smallint not null check (healers >= 0),
  melee      smallint not null check (melee >= 0),
  ranged     smallint not null check (ranged >= 0),
  updated_by bigint references users (id) on delete set null,
  updated_at timestamptz not null default now()
);

-- guild_raid_pings is the last "Ping on Discord" of a raid's non-signers (D16): one row per
-- raid, so the six-hour throttle is one atomic upsert.
create table if not exists guild_raid_pings (
  raid_id        bigint primary key references guild_raids (id) on delete cascade,
  last_pinged_at timestamptz not null,
  pinged_by      bigint references users (id) on delete set null,
  mentioned      int not null default 0
);

-- When the account joined the guild, so the attendance grid can tell a night before it from a
-- missed one. Null for every account that joined before this column (the grid then reads the
-- account's first sign-up or attendance row instead); every new row stamps now().
alter table guild_members add column if not exists joined_at timestamptz;
alter table guild_members alter column joined_at set default now();
