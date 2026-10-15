-- The companion connect flow (docs/contracts: device authorization). The companion mints a
-- connect, the player approves it on the site, and the companion's poll collects the device
-- token once. id_hash is the SHA-256 of the connect id the companion holds, so no secret is
-- at rest. user_code is the six symbols the site shows, stored without the hyphen. The device
-- itself is created by the poll that collects an approved row, so no raw token is ever stored.
create table if not exists device_connects (
  id_hash      bytea primary key,
  user_code    text not null,
  name         text not null default '',
  platform     text not null default '',
  user_id      bigint references users (id) on delete cascade,
  status       text not null default 'pending' check (status in ('pending', 'approved', 'denied')),
  created_at   timestamptz not null default now(),
  expires_at   timestamptz not null,
  collected_at timestamptz
);
-- A code is unique among the rows that can still be decided or collected.
create unique index if not exists device_connects_live_code_idx
  on device_connects (user_code) where collected_at is null;
create index if not exists device_connects_expires_idx on device_connects (expires_at);
