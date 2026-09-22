-- api/internal/db/migrations/0026_character_media.down.sql
alter table characters drop column if exists bnet_media;
alter table characters drop column if exists render_url;
alter table characters drop column if exists avatar_url;
