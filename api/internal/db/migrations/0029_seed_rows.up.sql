-- api/internal/db/migrations/0029_seed_rows.up.sql
-- Bookkeeping for api/cmd/seedguild: the one place every row the mock-guild seed tool
-- creates or mutates is recorded, so --remove can delete exactly what --apply wrote and
-- restore exactly what it overwrote, without hand-matching a battletag prefix or any
-- other heuristic against the real data a guild accumulates over time.
--
-- row_key is table-specific (a users.id, a characters.key, a "guild_id:character_key"
-- pair, a reports.id, ...) - whatever uniquely names the row within table_name. prior is
-- null for a row the seed created outright (remove = delete it); it holds a JSON
-- snapshot of the row's previous column values for a row the seed only mutated (the
-- owner's own guild_characters/guild_members rows, the guilds row it claims) - remove
-- restores those columns instead of deleting the row.
create table if not exists seed_rows (
  tag        text not null,
  table_name text not null,
  row_key    text not null,
  prior      jsonb,
  created_at timestamptz not null default now(),
  primary key (tag, table_name, row_key)
);
create index if not exists seed_rows_tag_idx on seed_rows (tag, table_name);
