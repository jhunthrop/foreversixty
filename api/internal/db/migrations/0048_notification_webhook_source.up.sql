-- Where a saved webhook came from: pasted by the member or created by the
-- one-click Discord connect. Rows saved before this change were all pasted.
-- guild_id is the Discord server the channel belongs to; null when unknown.
alter table notification_webhooks
  add column source text not null default 'pasted' check (source in ('pasted', 'connected')),
  add column guild_id text;
