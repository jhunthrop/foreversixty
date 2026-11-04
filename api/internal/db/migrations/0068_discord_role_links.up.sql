-- Discord Linked Roles: the Premium role in the Forever Sixty server follows the account's
-- Premium. One row per account that authorized the role connection; the refresh token is
-- sealed with NOTIFY_ENCRYPTION_KEY (AES-256-GCM), never stored in the clear.

alter table notify_connect_states
  drop constraint notify_connect_states_kind_check,
  add constraint notify_connect_states_kind_check
    check (kind in ('guild_webhook', 'bot_install', 'link', 'linked_role')),
  drop constraint notify_connect_states_guild_check,
  add constraint notify_connect_states_guild_check
    check (kind in ('link', 'linked_role') or guild_id is not null);

create table if not exists discord_role_links (
  user_id bigint primary key references users(id) on delete cascade,
  discord_user_id text not null unique,
  refresh_token_box bytea not null,
  -- What Discord was last told; null until the first push lands.
  pushed_premium boolean,
  pushed_at timestamptz,
  linked_at timestamptz not null default now()
);
