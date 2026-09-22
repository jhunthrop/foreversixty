-- api/internal/db/migrations/0025_bnet_capture.up.sql
-- Battle.net capture: store what Blizzard's account, character profile, and
-- equipment endpoints answer, verbatim, alongside the typed fields already
-- derived from them (spec docs/superpowers/specs/2026-09-22-battlenet-character-import-design.md,
-- amended by .superpowers/bnet-capture-brief.md §B).
alter table characters
  add column if not exists race text,
  add column if not exists gender text check (gender in ('male', 'female')),
  add column if not exists average_item_level int,
  add column if not exists equipped_item_level int,
  add column if not exists last_login_at timestamptz,
  add column if not exists bnet_account jsonb,    -- the account-profile entry, verbatim
  add column if not exists bnet_profile jsonb,    -- GET /profile/wow/character/{realm}/{name}, verbatim
  add column if not exists bnet_equipment jsonb,  -- .../equipment, verbatim
  add column if not exists bnet_captured_at timestamptz;
