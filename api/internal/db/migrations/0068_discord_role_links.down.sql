drop table if exists discord_role_links;

delete from notify_connect_states where kind = 'linked_role';
alter table notify_connect_states
  drop constraint notify_connect_states_guild_check,
  add constraint notify_connect_states_guild_check check (kind = 'link' or guild_id is not null),
  drop constraint notify_connect_states_kind_check,
  add constraint notify_connect_states_kind_check check (kind in ('guild_webhook', 'bot_install', 'link'));
