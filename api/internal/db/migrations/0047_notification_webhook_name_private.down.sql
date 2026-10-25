-- The upload_private flag stays in the events jsonb; it is harmless there.
alter table notification_webhooks drop column webhook_name;
