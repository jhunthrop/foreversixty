-- The phone apps' beta list (foreversixty.gg/app, internal/apptesters): one row per address and
-- platform, the address trimmed and lowercased by the API. Its own table, not subscribers:
-- that one is the double opt-in newsletter, unique per address, and mailed.
create table if not exists app_testers (
  id bigserial primary key,
  email text not null check (char_length(email) between 3 and 254),
  platform text not null check (platform in ('ios', 'android')),
  created_at timestamptz not null default now(),
  unique (email, platform)
);
