-- The player's notification preferences and delivery log (docs/superpowers/plans/
-- 2026-10-30-mobile-apps-lanes.md, "The player's lifecycle and its notifications").
-- notification_settings holds the zone and the quiet hours, read in that zone (minutes after
-- midnight; a start after the end wraps past midnight). notification_prefs holds only a
-- player's overrides of the catalogue's per-event defaults, which live in code
-- (api/internal/events/catalogue.go), so a new event needs no migration.
create table notification_settings (
  user_id       bigint primary key references users (id) on delete cascade,
  time_zone     text not null default 'UTC' check (char_length(time_zone) between 1 and 64),
  quiet_enabled boolean not null default false,
  quiet_start   smallint not null default 1320 check (quiet_start between 0 and 1439),
  quiet_end     smallint not null default 420 check (quiet_end between 0 and 1439),
  updated_at    timestamptz not null default now()
);

create table notification_prefs (
  user_id    bigint not null references users (id) on delete cascade,
  kind       text not null check (kind ~ '^[a-z]+\.[a-z_]+$'),
  channel    text not null check (channel in ('push', 'discord')),
  enabled    boolean not null,
  updated_at timestamptz not null default now(),
  primary key (user_id, kind, channel)
);

-- One row per event sent to an account on a channel: the dedupe key makes every publish and
-- every scheduler run idempotent per event, raid and player. Rows past the retention are
-- pruned by the raid-posts job.
create table notification_deliveries (
  user_id    bigint not null references users (id) on delete cascade,
  dedupe_key text not null check (char_length(dedupe_key) between 1 and 200),
  channel    text not null check (channel in ('push', 'discord')),
  sent_at    timestamptz not null default now(),
  primary key (user_id, dedupe_key, channel)
);
create index notification_deliveries_sent_idx on notification_deliveries (sent_at);
