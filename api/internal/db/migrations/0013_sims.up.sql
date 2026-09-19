-- The simulator: the premium flag, the execution score on every ranked
-- fight, saved sim results, and the nightly validation figure per spec.
--
-- fight_metrics is partitioned by fought_at; both statements below
-- propagate to every partition, which is what is wanted here.

alter table users add column if not exists premium boolean not null default false;

alter table fight_metrics add column if not exists execution_score numeric;
create index if not exists fight_metrics_execution_idx
  on fight_metrics (encounter_id, difficulty, spec, phase, execution_score desc);

create table if not exists sims (
  id             text primary key,
  user_id        bigint references users(id) on delete set null,
  spec           text not null,
  engine_version text not null,
  lane           text not null,
  dps_mean       numeric not null,
  dps_error      numeric not null,
  iterations     int not null,
  title          text,
  result         jsonb not null,       -- the whole SimResult, request included
  state          text not null default 'done',
  created_at     timestamptz not null default now()
);
create index if not exists sims_user_idx on sims (user_id, created_at desc);

create table if not exists sim_specs (
  spec           text primary key,
  state          text not null,        -- validated | in_progress | unsupported
  median_gap     numeric,
  parses         int not null default 0,
  worst_actions  jsonb not null default '[]',
  engine_version text,
  updated_at     timestamptz not null default now()
);
