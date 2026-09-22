-- api/internal/db/migrations/0026_character_media.up.sql
-- Character render/avatar images from Blizzard's character-media endpoint
-- (.superpowers/account-visual-brief.md §A), alongside the raw response
-- for the same 256 KiB cap / bnet_captured_at stamp the other captures use.
alter table characters
  add column if not exists avatar_url text,
  add column if not exists render_url text,
  add column if not exists bnet_media jsonb;
