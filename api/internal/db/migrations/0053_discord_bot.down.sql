drop table discord_shares;
drop table discord_command_sets;
drop table guild_raid_discord_posts;
drop table guild_discord_refusals;
drop table guild_discord_installs;
drop table discord_links;

delete from notify_connect_states where kind <> 'guild_webhook';
alter table notify_connect_states
  drop constraint notify_connect_states_guild_check,
  drop constraint notify_connect_states_kind_check,
  drop column kind,
  alter column guild_id set not null;
