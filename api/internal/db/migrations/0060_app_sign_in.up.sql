-- Native app sign-in (mobile lane A1). The app opens GET /v1/auth/app/start in the system
-- browser with a PKCE challenge; Battle.net sends the browser back to the shared callback,
-- which mints a one-time code and hands it to the app link; POST /v1/auth/app/exchange
-- swaps that code and the PKCE verifier for a device token.
--
-- One row per sign-in, in two phases. Start writes the Battle.net state (hashed), the
-- app's PKCE challenge, its platform and its own state. The callback fills code_hash,
-- user_id and code_expires_at once. The exchange stamps used_at, so a code spends once.
-- Every secret is stored as a SHA-256, like login_tokens and devices.
create table app_auth_flows (
  state_hash      bytea primary key,
  platform        text not null check (platform in ('ios', 'android')),
  code_challenge  text not null check (code_challenge ~ '^[A-Za-z0-9_-]{43}$'),
  app_state       text not null,
  created_at      timestamptz not null default now(),
  expires_at      timestamptz not null,
  code_hash       bytea unique,
  user_id         bigint references users (id) on delete cascade,
  code_expires_at timestamptz,
  used_at         timestamptz
);
create index app_auth_flows_expiry_idx on app_auth_flows (expires_at);

-- What kind of install holds a device token: the desktop companion (every row so far) or a
-- native app. platform alone cannot say, because the companion names its own platform.
alter table devices
  add column kind text not null default 'companion' check (kind in ('companion', 'app'));
