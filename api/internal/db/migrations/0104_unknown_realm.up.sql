-- Unknown realms (2026-11-13): Blizzard serves no Forever realm data, so a Battle.net
-- character on a realm the committed table (bnetapi/forever_realms.go) does not know, with a
-- private profile and no addon row, is imported under the ruleset 'unknown' and its player is
-- asked which realm type it is on. bnet_realm_id is the realm id its protected-character link
-- carries; ruleset_answered marks a ruleset the player chose. Each answer is also evidence for
-- the realm: realm_ruleset_answers holds one row per player per realm, and once two or more
-- players agree with no one disagreeing the import keys that realm's characters by it.
alter table characters drop constraint if exists characters_ruleset_check;
alter table characters add constraint characters_ruleset_check
  check (ruleset in ('normal', 'pvp', 'rp', 'hardcore', 'unknown'));
alter table characters add column if not exists bnet_realm_id bigint;
alter table characters add column if not exists ruleset_answered boolean not null default false;

create table if not exists realm_ruleset_answers (
  region      text   not null,
  realm_id    bigint not null,
  user_id     bigint not null references users (id) on delete cascade,
  ruleset     text   not null check (ruleset in ('normal', 'pvp', 'rp', 'hardcore')),
  answered_at timestamptz not null default now(),
  primary key (region, realm_id, user_id)
);
