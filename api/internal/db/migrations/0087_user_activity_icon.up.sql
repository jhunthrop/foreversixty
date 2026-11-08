-- An Activity row's icon (spec 2026-11-08-app-flows 4C): a name under /data/{build}/icons, set
-- by its producer or its group's (events.IconOf). Rows written before it are null and read as
-- their group's icon.
alter table user_activity add column if not exists icon text;
