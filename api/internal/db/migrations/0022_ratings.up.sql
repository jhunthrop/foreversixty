-- api/internal/db/migrations/0022_ratings.up.sql
-- Player performance ratings (docs/superpowers/specs/2026-09-21-performance-rating-
-- design.md §4.2). Numbered 0022, not the spec's own 0019: 0018 (guild membership) is
-- already on main, 0020 is the entitlements lane's migration and 0021 is reserved for its
-- follow-up, both in flight in other worktrees as this lane starts. golang-migrate only
-- ever moves a database forward from its currently-applied version, so a lower number
-- landing after a higher one has already deployed would be silently skipped forever -
-- 0022 clears all three. migration_order_test.go guards this class of mistake mechanically.

create table if not exists rating_scores (
  report_id      text not null,
  fight_index    int not null,
  player_key     text not null,
  player_name    text not null default '',
  class          text,
  spec           text,
  role           text,
  encounter_id   int,
  difficulty     int,
  size           int,
  duration_ms    int,
  kill           boolean not null,
  kill_time_band text not null default '',
  overall           numeric,
  overall_uncapped  numeric,
  overall_capped    boolean not null default false,
  components     jsonb not null,
  model_version  text not null,
  fought_at      timestamptz not null,
  computed_at    timestamptz not null default now(),
  primary key (report_id, fight_index, player_key, fought_at)
) partition by range (fought_at);

create index if not exists rating_scores_player_idx on rating_scores (player_key, fought_at desc);
create index if not exists rating_scores_report_idx on rating_scores (report_id, fight_index);
create index if not exists rating_scores_stale_idx on rating_scores (model_version) where model_version <> '';

create table if not exists rating_percentile_digests (
  encounter_id   int not null,
  difficulty     int not null,
  spec           text not null,
  role           text not null,
  kill_time_band text not null,
  kill           boolean not null,
  component      text not null,
  digest         bytea not null,
  n              bigint not null default 0,
  updated_at     timestamptz not null default now(),
  primary key (encounter_id, difficulty, spec, role, kill_time_band, kill, component)
);

create table if not exists kill_duration_digests (
  encounter_id int not null,
  difficulty   int not null,
  digest       bytea not null,
  n            bigint not null default 0,
  updated_at   timestamptz not null default now(),
  primary key (encounter_id, difficulty)
);
