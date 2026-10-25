-- The webhook's own name (Discord's `name` on the same GET that gives the
-- channel id), and the third event flag. Webhooks saved before this change
-- keep posting private logs, so their flag is set true once; new saves
-- default it to false.
alter table notification_webhooks add column webhook_name text;
update notification_webhooks set events = events || '{"upload_private": true}'::jsonb;
