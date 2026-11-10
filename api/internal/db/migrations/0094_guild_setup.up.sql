-- Guild setup checklist (design/specs/2026-11-09-guild-onboarding.md). Every step's state is
-- derived from the guild's own data; a mark records only what cannot be derived: an officer
-- confirming the current value (the officer ranks, the report visibility), saying the invite
-- link was posted or the Discord linked role made, or skipping a step.
create table guild_setup_marks (
  guild_id bigint not null references guilds(id) on delete cascade,
  step text not null check (step in ('officer_ranks', 'discord_bot', 'recruiting', 'raid_schedule',
    'invite_link', 'addon_sync', 'report_visibility', 'linked_role')),
  state text not null check (state in ('done', 'skipped')),
  set_by bigint references users(id) on delete set null,
  set_at timestamptz not null default now(),
  primary key (guild_id, step)
);

-- The recruiting channel: the bot posts each new application there.
alter table guild_discord_installs
  add column recruit_channel_id text,
  add column recruit_channel_name text;

-- A bot install started from the setup checklist returns to it.
alter table notify_connect_states
  drop constraint notify_connect_states_kind_check,
  add constraint notify_connect_states_kind_check
    check (kind in ('guild_webhook', 'bot_install', 'bot_install_setup', 'link', 'linked_role'));
