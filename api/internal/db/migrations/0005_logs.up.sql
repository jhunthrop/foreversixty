-- The Phase 3 logs product: accounts, devices, reports, fights, ranking
-- rows, percentile digests, moderation, and the small stores the uploads
-- and the addon need.

create table if not exists users (
  id          bigserial primary key,
  bnet_sub    text unique,
  battletag   text,
  email       text unique,
  role        text not null default 'user',
  anonymize   boolean not null default false,
  created_at  timestamptz not null default now(),
  constraint users_role_check check (role in ('user', 'moderator', 'admin'))
);

create table if not exists sessions (
  id         text primary key,
  user_id    bigint not null references users (id) on delete cascade,
  method     text not null,
  created_at timestamptz not null default now(),
  expires_at timestamptz not null
);
create index if not exists sessions_user_idx on sessions (user_id);
create index if not exists sessions_expiry_idx on sessions (expires_at);

-- Magic-link tokens. The contract's table list does not name this one,
-- but the email sign-in it specifies needs somewhere to keep a 20-minute
-- single-use token and somewhere to count the five-per-hour budget.
create table if not exists login_tokens (
  token_hash bytea primary key,
  email      text not null,
  created_at timestamptz not null default now(),
  expires_at timestamptz not null,
  used_at    timestamptz
);
create index if not exists login_tokens_email_idx on login_tokens (email, created_at desc);

create table if not exists devices (
  id           text primary key,
  user_id      bigint not null references users (id) on delete cascade,
  name         text not null default '',
  platform     text not null default '',
  token_hash   bytea not null unique,
  created_at   timestamptz not null default now(),
  last_seen_at timestamptz,
  revoked_at   timestamptz
);
create index if not exists devices_user_idx on devices (user_id);

-- Pairing codes, the other half of the device flow the contract
-- specifies: the site issues one, the companion claims it within ten
-- minutes.
create table if not exists pairing_codes (
  code       text primary key,
  user_id    bigint not null references users (id) on delete cascade,
  created_at timestamptz not null default now(),
  expires_at timestamptz not null,
  claimed_at timestamptz
);

create table if not exists guilds (
  id                 bigserial primary key,
  region             text not null,
  ruleset            text not null,
  name               text not null,
  claimed_by         bigint references users (id) on delete set null,
  default_visibility text not null default 'public',
  created_at         timestamptz not null default now(),
  unique (region, ruleset, name),
  constraint guilds_ruleset_check check (ruleset in ('normal', 'pvp', 'rp', 'hardcore'))
);

create table if not exists guild_members (
  guild_id     bigint not null references guilds (id) on delete cascade,
  user_id      bigint not null references users (id) on delete cascade,
  rank         text not null default 'member',
  refreshed_at timestamptz not null default now(),
  primary key (guild_id, user_id)
);

create table if not exists characters (
  key          text primary key,
  region       text not null,
  ruleset      text not null,
  name         text not null,
  class        text,
  user_id      bigint references users (id) on delete set null,
  refreshed_at timestamptz not null default now(),
  constraint characters_ruleset_check check (ruleset in ('normal', 'pvp', 'rp', 'hardcore'))
);
create index if not exists characters_user_idx on characters (user_id);

-- Whole-file uploads: one row per multipart upload, so the complete call
-- can find the object key and the parse job can find the upload.
create table if not exists uploads (
  id           text primary key,
  user_id      bigint references users (id) on delete set null,
  object_key   text not null,
  r2_upload_id text not null,
  size_bytes   bigint not null,
  filename     text not null default '',
  created_at   timestamptz not null default now(),
  completed_at timestamptz,
  report_id    text
);

create table if not exists reports (
  id                text primary key,
  owner_id          bigint references users (id) on delete set null,
  guild_id          bigint references guilds (id) on delete set null,
  title             text not null default '',
  visibility        text not null,
  zone              text not null default '',
  status            text not null,
  engine_version    text not null default '',
  upload_id         text,
  logging_character text,
  health            jsonb,
  flagged           text,
  created_at        timestamptz not null default now(),
  completed_at      timestamptz,
  constraint reports_visibility_check check (visibility in ('public', 'unlisted', 'private', 'guild')),
  constraint reports_status_check check (status in ('live', 'processing', 'complete', 'failed'))
);
create index if not exists reports_owner_idx on reports (owner_id, created_at desc);
create index if not exists reports_guild_idx on reports (guild_id, created_at desc);

-- raw_start_offset, raw_end_offset, raw_sha256, deaths and npc_kills are
-- beyond the contract's column list: the fight PUT answers 200 rather than 201 when
-- the same fight arrives again "with the same sha", and the raw-sample
-- job needs the byte range a fight came from. Both need the range and
-- the hash stored beside the fight.
create table if not exists fights (
  report_id        text not null references reports (id) on delete cascade,
  fight_index      int not null,
  encounter_id     int,
  name             text not null default '',
  difficulty       int,
  size             int,
  kill             boolean not null default false,
  duration_ms      int not null default 0,
  start_ms         bigint not null default 0,
  verified         boolean not null default false,
  players          text[] not null default '{}',
  deaths           int not null default 0,
  npc_kills        int not null default 0,
  raw_start_offset bigint,
  raw_end_offset   bigint,
  raw_sha256       bytea,
  primary key (report_id, fight_index)
);

-- One row per stored raw chunk, so a re-sent offset is recognised and an
-- offset that overlaps a different stored range is refused.
create table if not exists raw_chunks (
  report_id    text not null references reports (id) on delete cascade,
  start_offset bigint not null,
  end_offset   bigint not null,
  sha256       bytea not null,
  created_at   timestamptz not null default now(),
  primary key (report_id, start_offset)
);

-- talent_split, trinkets, buff_count and faction are beyond the
-- contract's column list for the same reason as the fight columns
-- above: the rankings row shape includes the first three and the
-- rankings filters include the fourth, and reading them per row from
-- the summaries on R2 would cost one object read per ranked row.
create table if not exists fight_metrics (
  id           bigint generated always as identity,
  report_id    text not null,
  fight_index  int not null,
  player_key   text not null,
  player_name  text not null default '',
  class        text,
  spec         text,
  role         text,
  ilvl         int,
  metric_dps   numeric,
  metric_hps   numeric,
  damage_taken bigint,
  active_ms    int,
  deaths       int,
  encounter_id int,
  difficulty   int,
  size         int,
  duration_ms  int,
  kill         boolean,
  phase        text,
  fought_at    timestamptz not null,
  talent_split text,
  trinkets     bigint[] not null default '{}',
  buff_count   int not null default 0,
  faction      text,
  state        text not null default 'ok',
  primary key (id, fought_at),
  unique (report_id, fight_index, player_key, fought_at)
) partition by range (fought_at);

create index if not exists fight_metrics_dps_idx
  on fight_metrics (encounter_id, difficulty, spec, phase, metric_dps desc);
create index if not exists fight_metrics_hps_idx
  on fight_metrics (encounter_id, difficulty, spec, phase, metric_hps desc);
create index if not exists fight_metrics_taken_idx
  on fight_metrics (encounter_id, difficulty, spec, phase, damage_taken desc);
create index if not exists fight_metrics_player_idx on fight_metrics (player_key, fought_at desc);
create index if not exists fight_metrics_report_idx on fight_metrics (report_id, fight_index);

create table if not exists percentile_digests (
  encounter_id int not null,
  difficulty   int not null,
  spec         text not null,
  phase        text not null,
  metric       text not null,
  digest       bytea not null,
  n            bigint not null default 0,
  updated_at   timestamptz not null default now(),
  primary key (encounter_id, difficulty, spec, phase, metric)
);

create table if not exists moderation (
  id          bigserial primary key,
  target_kind text not null,
  target_id   text not null,
  state       text not null,
  reason      text not null default '',
  actor_id    bigint references users (id) on delete set null,
  created_at  timestamptz not null default now(),
  constraint moderation_state_check check (state in ('ok', 'at_risk', 'removed'))
);
create index if not exists moderation_target_idx on moderation (target_kind, target_id, created_at desc);

create table if not exists addon_exports (
  character_key text primary key,
  user_id       bigint not null references users (id) on delete cascade,
  region        text not null,
  ruleset       text not null,
  name          text not null,
  export        text not null,
  updated_at    timestamptz not null default now()
);
create index if not exists addon_exports_user_idx on addon_exports (user_id);

create table if not exists addon_inbox (
  id            bigserial primary key,
  user_id       bigint not null references users (id) on delete cascade,
  character_key text not null default '',
  build_id      text not null,
  created_at    timestamptz not null default now(),
  unique (user_id, character_key, build_id)
);
