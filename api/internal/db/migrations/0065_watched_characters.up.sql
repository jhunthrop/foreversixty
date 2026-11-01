-- Watching (mobile Wave 3 lane F11, contract W3-C7). One row per character an account watches
-- from the apps' character page; the apps show the list in Search and Spotlight. character_key
-- is the rankings' player key (region/ruleset/name-slug) and carries no foreign key: a
-- character known only from logs has no characters row. An account keeps at most 50 (the
-- handler's rule); a watched character whose account turns anonymous is left out on read.
create table watched_characters (
  user_id       bigint not null references users (id) on delete cascade,
  character_key text not null check (length(character_key) between 5 and 120),
  added_at      timestamptz not null default now(),
  primary key (user_id, character_key)
);
