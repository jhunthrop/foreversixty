drop table guild_raid_discord_schedule;

alter table guild_raid_discord_posts
  drop column reminded_at,
  drop column reminder_message_id,
  drop column remind_due;

alter table guild_notification_webhooks
  drop column remind_default,
  alter column post_default set default false;
alter table guild_discord_installs
  drop column remind_default,
  alter column post_default set default false;
