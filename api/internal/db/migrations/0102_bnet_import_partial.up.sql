-- Battle.net import: a partial end state and how many characters it could not key.
--
-- A run that keys some of the account's characters and skips the rest (no source named
-- their ruleset while Blizzard lists no Forever realm, or a key collision) used to read as
-- failed whenever it also ran out of budget, with the keyed characters in place. It now
-- reads as partial, and bnet_import_unresolved counts the skipped characters so GET /v1/me
-- can say how many are missing (bnet_import.unresolved). Only api/internal/auth writes
-- them; the count is null unless the last run skipped any.
alter table users drop constraint if exists users_bnet_import_state_check;
alter table users
  add constraint users_bnet_import_state_check
    check (bnet_import_state in ('running', 'done', 'partial', 'failed')),
  add column if not exists bnet_import_unresolved integer;
