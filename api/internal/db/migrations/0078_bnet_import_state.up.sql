-- Battle.net import state (design/specs/2026-11-07-setup.md section 9, G1).
--
-- Sign-in runs the Battle.net import for up to 5 s and may continue it in the background
-- for up to 2 min (api/internal/auth/importer.go). GET /v1/me reads these two columns as
-- bnet_import { state, at } so /setup can say "still importing" from any Cloud Run instance.
-- Both are null until the first tracked import. Only api/internal/auth writes them.
alter table users
  add column if not exists bnet_import_state text
    check (bnet_import_state in ('running', 'done', 'failed')),
  add column if not exists bnet_import_state_at timestamptz;
