-- api/internal/db/migrations/0033_character_syncs.up.sql
-- One row per character sync: an addon export taken, or a Battle.net import or refresh run.
-- addon_exports keeps only the newest capture, so this is the history GET /v1/me needs for
-- build.median_sync_gap_sec (the median gap between successive 'ok' rows) and
-- build.sync_error (the outcome of the latest 'blizzard' row when it is not 'ok').
-- A character re-keyed by the Battle.net refresh carries its history along (on update
-- cascade); a deleted character takes it with it.
create table if not exists character_syncs (
  id            bigserial primary key,
  character_key text not null references characters (key) on update cascade on delete cascade,
  source        text not null check (source in ('addon', 'blizzard')),
  outcome       text not null,
  created_at    timestamptz not null default now()
);
create index if not exists character_syncs_key_idx on character_syncs (character_key, created_at);
