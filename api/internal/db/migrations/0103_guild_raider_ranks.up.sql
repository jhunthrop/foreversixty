-- Guild raider ranks: which guild rank indexes (0 is the guild master, as rank_index) a guild
-- counts as its raiders. The addon's Officer roster lists raiders only by default and the
-- Guild tab places a player among them (design/specs/2026-11-13-addon-hub.md section 9); the
-- nightly Data addon carries the list in the guild's record (contract 5's raider_ranks). Null
-- until an officer says, and the addon then counts every level 60 as a raider. Only
-- api/internal/guilds writes it (PATCH /v1/guilds/{id}/settings).
alter table guilds add column if not exists raider_rank_indexes smallint[];
