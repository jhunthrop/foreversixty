-- Calendar links (mobile Wave 3, W3-C17). GET /v1/me/calendar-feed makes a player's link on
-- first call; GET /v1/calendar/{token}.ics serves the planned raids of their guilds to any
-- calendar app holding it. The token is a 256-bit random secret shown once: only its SHA-256
-- is written down, like devices and login_tokens. A reset stamps revoked_at on the live row
-- and inserts a new one, so an old link answers 404 at once. One live link per account.
create table calendar_feeds (
  id         bigserial primary key,
  user_id    bigint not null references users (id) on delete cascade,
  token_hash bytea not null unique check (octet_length(token_hash) = 32),
  created_at timestamptz not null default now(),
  revoked_at timestamptz
);
create unique index calendar_feeds_live_idx on calendar_feeds (user_id) where revoked_at is null;
