-- The apps' push tokens (APNs device tokens, FCM registration tokens). Each is tied to the
-- device token that registered it, so it goes when the device does: deleted with the device
-- row, and dropped by the trigger below the moment the device is revoked. A token is unique:
-- registering it again (another account signing in on the phone) moves it.
create table push_tokens (
  token       text primary key check (char_length(token) between 16 and 4096),
  platform    text not null check (platform in ('ios', 'android')),
  user_id     bigint not null references users (id) on delete cascade,
  device_id   text not null references devices (id) on delete cascade,
  app_version text not null default '' check (char_length(app_version) <= 32),
  created_at  timestamptz not null default now(),
  updated_at  timestamptz not null default now()
);
create index push_tokens_user_idx on push_tokens (user_id);
create index push_tokens_device_idx on push_tokens (device_id);

create function push_tokens_drop_revoked() returns trigger language plpgsql as $$
begin
  delete from push_tokens where device_id = new.id;
  return new;
end
$$;

create trigger devices_revoked_drop_push_tokens
  after update of revoked_at on devices
  for each row when (new.revoked_at is not null and old.revoked_at is null)
  execute function push_tokens_drop_revoked();
