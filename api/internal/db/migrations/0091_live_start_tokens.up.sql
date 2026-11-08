-- Push-to-start tokens (spec 2026-11-08-app-flows F2): one per app install and kind, the
-- ActivityKit push-to-start token on iOS (the FCM token on Android), so the server can start a
-- raid's Live Activity on raid.starting without the app running. Kept as sent because the
-- senders need it; never logged and never answered back. A token the push service refuses goes.
create table if not exists live_activity_start_tokens (
  user_id    bigint not null references users (id) on delete cascade,
  device_id  text not null references devices (id) on delete cascade,
  kind       text not null check (kind in ('raid')),
  platform   text not null check (platform in ('ios', 'android')),
  token      text not null,
  updated_at timestamptz not null default now(),
  primary key (user_id, device_id, kind)
);
