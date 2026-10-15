-- Search characters by the start of their name (GET /v1/characters/search;
-- api/internal/rankings/character_search.go).
--
-- The lookup filters on lower(name) like 'prefix%', which only an index with the pattern
--   operator class can serve under a non-C collation. characters was indexed by user only.
create index if not exists characters_name_prefix_idx on characters (lower(name) text_pattern_ops);
