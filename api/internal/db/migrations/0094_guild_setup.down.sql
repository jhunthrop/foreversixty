delete from notify_connect_states where kind = 'bot_install_setup';
alter table notify_connect_states
  drop constraint notify_connect_states_kind_check,
  add constraint notify_connect_states_kind_check
    check (kind in ('guild_webhook', 'bot_install', 'link', 'linked_role'));

alter table guild_discord_installs
  drop column if exists recruit_channel_id,
  drop column if exists recruit_channel_name;

drop table if exists guild_setup_marks;
