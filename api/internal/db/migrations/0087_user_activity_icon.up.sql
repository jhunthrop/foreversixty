-- An Activity row's icon (spec 2026-11-08-app-flows 4C): a name under /data/{build}/icons, or
-- crest:<class> for a character; null on rows written before it, which the apps draw with their
-- group's fixed icon.
alter table user_activity add column if not exists icon text;
