-- Log Coach (design 2026-10-27 sections 11, 11.1, 15): the per-fight profile basis, the
-- sim reference cache, built packs, generated reviews with their leases, and the spend and
-- cap counters. Migration 0054 belongs to the launch-pricing lane; this is 0055.
--
-- coach_profiles is partitioned on fought_at like fight_metrics. This migration creates
-- only the parent: partitions come from db.EnsurePartitions / PartitionJob (section 11.1).

create table coach_profiles (
  report_id text not null,
  fight_index int not null,
  player_key text not null,
  class text not null,
  spec text not null,
  role text not null,
  encounter_id bigint not null,
  difficulty bigint not null,
  size int not null,
  kill boolean not null,
  duration_ms bigint not null,
  phase text not null,
  metric double precision not null,
  fought_at timestamptz not null,
  model_version text not null,
  profile jsonb not null,
  primary key (report_id, fight_index, player_key, fought_at)
) partition by range (fought_at);

create index coach_profiles_basis_idx
  on coach_profiles (encounter_id, difficulty, size, spec, phase, kill, metric desc);
create index coach_profiles_player_idx on coach_profiles (player_key, fought_at desc);

create table coach_sim_refs (
  engine_version text,
  build text,
  spec text,
  band int,
  duration_bucket int,
  summary jsonb not null,
  computed_at timestamptz not null default now(),
  primary key (engine_version, build, spec, band, duration_bucket)
);

create table coach_packs (
  report_id text references reports(id) on delete cascade,
  fight_index int,
  player_key text,
  pack_version text,
  pack jsonb not null,
  built_at timestamptz not null default now(),
  primary key (report_id, fight_index, player_key, pack_version)
);

create table coach_reviews (
  report_id text references reports(id) on delete cascade,
  fight_index int,
  player_key text,
  pack_version text,
  model_version text,
  persona_version text,
  review jsonb not null,
  source text not null check (source in ('model', 'rules')),
  reason text,
  retryable boolean not null default false,
  validation jsonb not null default '[]',
  tokens_in int not null default 0,
  tokens_out int not null default 0,
  cost_usd numeric(12,6) not null default 0,
  cost_estimated boolean not null default false,
  model_slug text,
  gateway_request_id text,
  account_id bigint references users(id) on delete set null,
  generated_at timestamptz not null default now(),
  primary key (report_id, fight_index, player_key, pack_version, model_version, persona_version)
);

create index coach_reviews_account_idx on coach_reviews (account_id, generated_at desc);
create index coach_reviews_generated_idx on coach_reviews (generated_at);

create table coach_review_leases (
  report_id text references reports(id) on delete cascade,
  fight_index int,
  player_key text,
  pack_version text,
  model_version text,
  persona_version text,
  holder text not null,
  leased_until timestamptz not null,
  primary key (report_id, fight_index, player_key, pack_version, model_version, persona_version)
);

create index coach_review_leases_expiry_idx on coach_review_leases (leased_until);

create table llm_spend (
  day date primary key,
  usd numeric(12,6) not null default 0,
  calls int not null default 0,
  estimated_calls int not null default 0
);

create table coach_account_day (
  day date,
  user_id bigint references users(id) on delete cascade,
  generations int not null default 0,
  primary key (day, user_id)
);

create index coach_account_day_user_idx on coach_account_day (user_id);

create table coach_review_samples (
  day date,
  report_id text references reports(id) on delete cascade,
  fight_index int,
  player_key text,
  character_name text not null,
  spec text not null,
  source text not null,
  validation jsonb not null,
  pack jsonb not null,
  review jsonb not null,
  sampled_at timestamptz not null default now(),
  primary key (day, report_id, fight_index, player_key)
);
