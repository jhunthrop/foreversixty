alter table reports drop column if exists heartbeat_at;
alter table fights drop column if exists uploaded_at;
alter table addon_exports drop column if exists seen_at;
