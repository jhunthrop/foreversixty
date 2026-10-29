-- Discord guided posting: raids post themselves on a schedule, and the reminder default is remembered.

-- Ruling 5: the opt-in false the guilds hold is the default, not a choice; both boxes start on.
alter table guild_discord_installs
  alter column post_default set default true,
  add column remind_default boolean not null default true;
alter table guild_notification_webhooks
  alter column post_default set default true,
  add column remind_default boolean not null default true;
update guild_discord_installs set post_default = true;
update guild_notification_webhooks set post_default = true;

-- The bot's drainer needs only the reminder's own flag and reference.
alter table guild_raid_discord_posts
  add column remind_due boolean not null default false,
  add column reminder_message_id text,
  add column reminded_at timestamptz;

-- One schedule row per raid, destination-neutral: the bot or the webhook is picked at send time.
create table guild_raid_discord_schedule (
  raid_id bigint primary key references guild_raids(id) on delete cascade,
  guild_id bigint not null references guilds(id) on delete cascade,
  post_at timestamptz not null,
  remind boolean not null,
  posted_at timestamptz,
  reminded_at timestamptz,
  reminder_via text check (reminder_via in ('bot','webhook')),
  reminder_message_id text,
  lease_until timestamptz,
  attempts int not null default 0,
  next_attempt_at timestamptz not null default now(),
  last_error text,
  created_at timestamptz not null default now()
);
create index guild_raid_discord_schedule_due_idx on guild_raid_discord_schedule (next_attempt_at)
  where posted_at is null or reminded_at is null;

-- Ruling 6: every post the bot already made counts as posted, so an edited old raid never posts twice.
insert into guild_raid_discord_schedule (raid_id, guild_id, post_at, remind, posted_at, created_at)
select raid_id, guild_id, created_at, false, created_at, created_at from guild_raid_discord_posts;
