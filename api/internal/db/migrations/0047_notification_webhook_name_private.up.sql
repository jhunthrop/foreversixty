-- The webhook's own name (Discord's `name` on the same GET that gives the
-- channel id), and the third event flag. Webhooks saved before this change
-- keep posting private logs, so their flag is set true once; new saves
-- default it to false.
alter table notification_webhooks add column webhook_name text;
-- Idempotent and opt-out-safe: a row that already has the key is left alone.
update notification_webhooks set events = '{"upload_private": true}'::jsonb || events
  where not (events ? 'upload_private');
