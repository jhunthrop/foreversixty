-- Home B, "the world is moving" (design/specs/2026-11-07-home-b.md section 2).
--
-- activity_events is the site feed (GET /v1/feed): one row per event, written by
-- api/internal/activity from the report completion, the guild raid planner, the hourly board
-- job, the BiS startup diff and the sim job. guild_id set means the row is its guild's only.
-- The hourly loop deletes rows older than 30 days.
create table if not exists activity_events (
  id           bigserial primary key,
  kind         text not null,
  occurred_at  timestamptz not null,
  spec         text,
  band         int,
  guild_id     bigint references guilds (id) on delete cascade,
  character_key text,
  tag_base     text not null,
  icon         text not null default '',
  crest        text not null default '',
  quality      int,
  href         text not null default '',
  date_only    boolean not null default false,
  segments     jsonb not null,
  dedupe_key   text not null unique,
  created_at   timestamptz not null default now()
);
create index if not exists activity_events_time_idx on activity_events (occurred_at desc, id desc);
create index if not exists activity_events_guild_idx on activity_events (guild_id, occurred_at desc)
  where guild_id is not null;
create index if not exists activity_events_spec_idx on activity_events (spec, band, occurred_at desc)
  where spec is not null;

-- The last seen pick per spec, band, faction, preset and slot: the API's startup diff
-- against the deployed BiS files writes the bis_change rows from it.
create table if not exists bis_pick_state (
  spec       text not null,
  band       int not null,
  faction    text not null,
  preset     text not null,
  slot       text not null,
  item_id    int not null,
  updated_at timestamptz not null default now(),
  primary key (spec, band, faction, preset, slot)
);

-- The tier grade per reset week (Tuesday 15:00 UTC), spec, band and faction, upserted
-- hourly so a week's row is its last observed state. "Up from B last week" reads the
-- previous week's row; with none the home shows no delta.
create table if not exists tier_grade_weeks (
  week       date not null,
  spec       text not null,
  band       int not null,
  faction    text not null,
  role       text not null,
  grade      text,
  place      int not null,
  of_n       int not null,
  updated_at timestamptz not null default now(),
  primary key (week, spec, band, faction)
);

-- The spec board per reset week: each character's place on its spec and level band,
-- replaced hourly for the current week. Movement compares with the previous week.
create table if not exists spec_board_weeks (
  week          date not null,
  spec          text not null,
  band          int not null,
  character_key text not null,
  name          text not null,
  class         text not null,
  rank          int not null,
  score         double precision not null,
  updated_at    timestamptz not null default now(),
  primary key (week, spec, band, character_key)
);
create index if not exists spec_board_weeks_rank_idx on spec_board_weeks (week, spec, band, rank);

-- A character's server sims with what the guild ladder compares: the import sim (one per
-- gear change) and the addon's in-game requests. The figure is the sims row's.
create table if not exists character_sims (
  sim_id        text primary key references sims (id) on delete cascade,
  character_key text not null,
  user_id       bigint not null references users (id) on delete cascade,
  kind          text not null check (kind in ('import', 'addon', 'saved')),
  spec          text not null,
  role          text not null,
  band          int not null,
  preset        text not null,
  gear_hash     text not null default '',
  posted_value  double precision,
  created_at    timestamptz not null default now()
);
create index if not exists character_sims_character_idx on character_sims (character_key, created_at desc);
create index if not exists character_sims_user_day_idx on character_sims (user_id, created_at desc);
