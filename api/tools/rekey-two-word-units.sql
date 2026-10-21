-- One-off repair for rows keyed by half a two-word name. Before the SplitUnit fix a logged
-- unit "Bow-Jackzon-ClassicBetaPvP-" was keyed us/pvp/bow (the first word) instead of the
-- export's us/pvp/bow-jackzon. This re-keys fight_metrics and rating_scores rows whose
-- player_name has the three-segment First-Last-Realm form and whose key is still the first
-- word alone. A row whose corrected key already exists for the same fight is left for the
-- report's next re-ingest to replace. Migration 0043_rekey_two_word_units applies this on
-- deploy, so production needs no manual run; this file stays for another database. Take a
-- dump first and run:
--   psql "$DIRECT_URL" -v ON_ERROR_STOP=1 -f api/tools/rekey-two-word-units.sql
begin;

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

commit;
