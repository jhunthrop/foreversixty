-- api/internal/db/migrations/0073_loot_council.up.sql
-- The loot council (design/specs/2026-11-05-web-layout.md §7.2, §7.4, decision D13): officers
-- vote per drop and candidate, a raider flags a drop main spec or off spec, and an award keeps
-- its decision record. A drop is keyed the way loot_awards keys it: the guild, the night (a
-- planned raid, else a logged night's report, else no night yet), the encounter and the item.
-- Additive only: the running API never reads these columns or tables.

-- The award's decision record. Null on every award made before the council: not recorded.
alter table loot_awards
  add column ms_os          text check (ms_os in ('ms', 'os')),
  add column reason         text check (char_length(reason) <= 200),
  add column votes_yes      integer check (votes_yes >= 0),
  add column votes_no       integer check (votes_no >= 0),
  add column votes_officers integer check (votes_officers >= 0);

-- One officer's yes or no on one candidate for one drop of one night; changeable, removable.
create table loot_votes (
  id            bigserial primary key,
  guild_id      bigint not null references guilds (id) on delete cascade,
  raid_id       bigint references guild_raids (id) on delete cascade,
  report_id     text references reports (id) on delete cascade,
  encounter_id  bigint not null,
  item_id       bigint not null,
  character_key text not null,
  voter_id      bigint not null references users (id) on delete cascade,
  vote          text not null check (vote in ('yes', 'no')),
  voted_at      timestamptz not null default now()
);
create unique index loot_votes_one_idx on loot_votes
  (guild_id, coalesce(raid_id, 0), coalesce(report_id, ''), encounter_id, item_id, character_key, voter_id);

-- A character's main spec or off spec interest in one drop of one night, set by the raider or
-- by an officer on their behalf (set_by).
create table loot_interest (
  id            bigserial primary key,
  guild_id      bigint not null references guilds (id) on delete cascade,
  raid_id       bigint references guild_raids (id) on delete cascade,
  report_id     text references reports (id) on delete cascade,
  encounter_id  bigint not null,
  item_id       bigint not null,
  character_key text not null,
  ms_os         text not null check (ms_os in ('ms', 'os')),
  set_by        bigint references users (id) on delete set null,
  set_at        timestamptz not null default now()
);
create unique index loot_interest_one_idx on loot_interest
  (guild_id, coalesce(raid_id, 0), coalesce(report_id, ''), encounter_id, item_id, character_key);
