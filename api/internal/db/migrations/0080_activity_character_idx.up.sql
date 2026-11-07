-- The character page's activity (GET /v1/feed?character=<key>, design/specs/2026-11-06-character.md
-- section 6): one character's rows, newest first.
create index if not exists activity_events_character_idx on activity_events (character_key, occurred_at desc)
  where character_key is not null;
