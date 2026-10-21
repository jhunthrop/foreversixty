-- Re-key the rows that were keyed by half of a two-word name (api/tools/rekey-two-word-units.sql
-- is the same repair for a manual run). Before character.SplitUnit learned the export's
-- First-Last-Realm form, a logged unit "Bow-Jackzon-ClassicBetaPvP-" was keyed us/pvp/bow (the
-- first word) instead of the key the addon export and the character page use, us/pvp/bow-jackzon,
-- so the player's rankings and ratings never met their character. The owner cannot run SQL
-- against production, so the repair rides the deploy like any other data migration (0016,
-- 0027, 0036).
--
-- The tables keyed by a logged player are fight_metrics and rating_scores; they are the only
-- two that character.KeyFromUnit wrote a player_key into. Nothing else holds a key built from a
-- logged unit: characters, guild_characters, loot_awards and addon_exports take their keys from
-- a verified character, the addon or the guild roster, never from a log line.
--
-- A row is half-keyed when its player_name has the three-segment First-Last-Realm shape (an
-- optional trailing dash is the empty second realm word) and the key's name is still the first
-- word alone. A two-segment name ("Baelgrim-Nightslayer") never matches, so a single-word
-- character is left as it is. The corrected key keeps the row's own region and ruleset
-- segments, which were right all along.
--
-- A row whose corrected key already exists for the same fight (a re-ingest after the fix wrote
-- it) is left alone rather than collide with the primary key; the correct row is the one
-- rankings and ratings now read. Both updates only touch rows that still match, so a second run
-- changes nothing. golang-migrate sends a file as one multi-statement query, which Postgres
-- runs as a single implicit transaction, so no begin/commit is needed and a failure leaves
-- both tables as they were.

update fight_metrics m
set player_key = split_part(m.player_key, '/', 1) || '/' || split_part(m.player_key, '/', 2) || '/'
  || lower(split_part(m.player_name, '-', 1)) || '-' || lower(split_part(m.player_name, '-', 2))
where m.player_name ~ '^[^-]+-[^-]+-[^-]+-?$'
  and split_part(m.player_key, '/', 3) = lower(split_part(m.player_name, '-', 1))
  and not exists (
    select 1 from fight_metrics o
    where o.report_id = m.report_id and o.fight_index = m.fight_index and o.fought_at = m.fought_at
      and o.player_key = split_part(m.player_key, '/', 1) || '/' || split_part(m.player_key, '/', 2) || '/'
        || lower(split_part(m.player_name, '-', 1)) || '-' || lower(split_part(m.player_name, '-', 2)));

update rating_scores m
set player_key = split_part(m.player_key, '/', 1) || '/' || split_part(m.player_key, '/', 2) || '/'
  || lower(split_part(m.player_name, '-', 1)) || '-' || lower(split_part(m.player_name, '-', 2))
where m.player_name ~ '^[^-]+-[^-]+-[^-]+-?$'
  and split_part(m.player_key, '/', 3) = lower(split_part(m.player_name, '-', 1))
  and not exists (
    select 1 from rating_scores o
    where o.report_id = m.report_id and o.fight_index = m.fight_index and o.fought_at = m.fought_at
      and o.player_key = split_part(m.player_key, '/', 1) || '/' || split_part(m.player_key, '/', 2) || '/'
        || lower(split_part(m.player_name, '-', 1)) || '-' || lower(split_part(m.player_name, '-', 2)));
