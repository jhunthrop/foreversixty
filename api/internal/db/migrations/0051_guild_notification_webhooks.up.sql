-- The guild's one Discord channel for raid posts (design/specs/2026-10-26-guild-raids.md §3.8).
-- token_box is the sealed webhook URL; it is never returned or logged.
create table guild_notification_webhooks (
  guild_id bigint primary key references guilds(id) on delete cascade,
  token_box bytea not null,
  channel_id text not null,
  discord_guild_id text,
  failures int not null default 0,
  -- disabled_reason is set when five counted failures switch posting off; a successful
  -- test or a reconnect clears it.
  disabled_reason text,
  last_error text,
  connected_by bigint references users(id) on delete set null,
  -- post_default is the raid form's remembered "Post to Discord" box.
  post_default boolean not null default false,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);
-- The guild connect's OAuth state (the account connect keeps its cookie): single use,
-- ten minutes, bound to the officer who started it.
create table notify_connect_states (
  state text primary key,
  user_id bigint not null references users(id) on delete cascade,
  guild_id bigint not null references guilds(id) on delete cascade,
  expires_at timestamptz not null
);
create index notify_connect_states_expires_idx on notify_connect_states (expires_at);
