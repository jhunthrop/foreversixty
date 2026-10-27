-- Discord bot: install, account link, refusals, durable raid posts, command sets, shares.

-- Connect states gain a kind; an account link has no guild.
alter table notify_connect_states
  add column kind text not null default 'guild_webhook',
  add constraint notify_connect_states_kind_check check (kind in ('guild_webhook', 'bot_install', 'link')),
  alter column guild_id drop not null,
  add constraint notify_connect_states_guild_check check (kind = 'link' or guild_id is not null);

create table discord_links (
  user_id bigint primary key references users(id) on delete cascade,
  discord_user_id text not null unique,
  display_name text not null,
  linked_at timestamptz not null default now()
);

create table guild_discord_installs (
  guild_id bigint primary key references guilds(id) on delete cascade,
  server_id text not null unique,
  server_name text not null,
  channel_id text,
  channel_name text,
  failures int not null default 0,
  last_error text,
  removed_in_discord boolean not null default false,
  installed_by bigint references users(id) on delete set null,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);

create table guild_discord_refusals (
  user_id bigint not null references users(id) on delete cascade,
  guild_id bigint not null references guilds(id) on delete cascade,
  held_by_name text not null,
  expires_at timestamptz not null,
  primary key (user_id, guild_id)
);
create index guild_discord_refusals_guild_idx on guild_discord_refusals (guild_id);

create table guild_raid_discord_posts (
  raid_id bigint primary key references guild_raids(id) on delete cascade,
  guild_id bigint not null references guilds(id) on delete cascade,
  channel_id text not null,
  message_id text,
  moved_from timestamptz,
  dirty_seq bigint not null default 1,
  dirty_at timestamptz default now(),
  wake_at timestamptz,
  edited_at timestamptz,
  attempts int not null default 0,
  next_attempt_at timestamptz not null default now(),
  lease_until timestamptz,
  last_error text,
  gone boolean not null default false,
  created_at timestamptz not null default now()
);
create index guild_raid_discord_posts_guild_idx on guild_raid_discord_posts (guild_id);
create index guild_raid_discord_posts_due_idx on guild_raid_discord_posts (next_attempt_at)
  where not gone and (dirty_at is not null or wake_at is not null);

create table discord_command_sets (
  application_id text not null,
  scope text not null,
  hash text not null,
  registered_at timestamptz not null default now(),
  primary key (application_id, scope)
);

create table discord_shares (
  token text primary key,
  guild_id bigint not null references guilds(id) on delete cascade,
  command text not null,
  options jsonb not null,
  used_at timestamptz,
  expires_at timestamptz not null
);
create index discord_shares_guild_idx on discord_shares (guild_id);
create index discord_shares_expires_idx on discord_shares (expires_at);
