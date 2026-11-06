-- api/internal/db/migrations/0076_raid_drops.up.sql
-- What dropped (guild slice 7, owner approval): the boss corpses the addon's loot window read in a
-- raid, from the export's drops= section, one row per item of a corpse. A new table only, so the
-- running API keeps working against this schema.
--
-- A drop is filed under a night the way an award is (loot_council_night.go): a planned raid
-- (raid_id), else a logged night (report_id), else none yet, read by its time, zone and encounter
-- until a match attaches it (attached_by_match, so a clear can undo it). corpse_key is the corpse
-- GUID the client gave, else "kill:<encounter>:<UTC date>": several raiders export the same
-- corpse, and the unique index keeps it once. winners is who the loot chat says received a copy.
create table if not exists raid_drops (
  id                 bigserial primary key,
  guild_id           bigint not null references guilds (id) on delete cascade,
  raid_id            bigint references guild_raids (id) on delete set null,
  report_id          text references reports (id) on delete set null,
  attached_by_match  boolean not null default false,
  zone               text not null default '',
  instance_id        int not null default 0 check (instance_id >= 0),
  encounter_id       bigint not null check (encounter_id >= 0),
  item_id            int not null check (item_id > 0),
  count              int not null default 1 check (count between 1 and 40),
  winners            text[] not null default '{}' check (cardinality(winners) <= 40),
  loot_method        text not null default '' check (char_length(loot_method) <= 32),
  corpse_key         text not null check (char_length(corpse_key) between 1 and 128),
  dropped_at         timestamptz not null,
  reported_by        bigint references users (id) on delete set null,
  reported_character text not null,
  reported_at        timestamptz not null default now()
);

create unique index if not exists raid_drops_corpse_idx on raid_drops (guild_id, encounter_id, item_id, corpse_key);
create index if not exists raid_drops_raid_idx on raid_drops (raid_id) where raid_id is not null;
create index if not exists raid_drops_unmatched_idx on raid_drops (guild_id, dropped_at) where raid_id is null;
