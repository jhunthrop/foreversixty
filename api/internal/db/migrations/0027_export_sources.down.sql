-- api/internal/db/migrations/0027_export_sources.down.sql
alter table characters drop column if exists bnet_talents;
alter table addon_exports drop column if exists captured_at;
alter table addon_exports drop column if exists source;
