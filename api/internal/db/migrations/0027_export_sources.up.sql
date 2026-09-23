-- api/internal/db/migrations/0027_export_sources.up.sql
-- Battle.net-first (spec docs/superpowers/specs/2026-09-22-battlenet-first-design.md §2.1,
-- §2.3): addon_exports learns which path wrote it and when it was true in the game, and
-- characters gains a raw capture column for the specializations read, alongside
-- bnet_profile/bnet_equipment (migration 0025).
alter table addon_exports
  add column if not exists source text not null default 'addon' check (source in ('addon', 'blizzard')),
  add column if not exists captured_at timestamptz;
update addon_exports set captured_at = updated_at where captured_at is null;
alter table addon_exports alter column captured_at set not null;

alter table characters
  add column if not exists bnet_talents jsonb;
