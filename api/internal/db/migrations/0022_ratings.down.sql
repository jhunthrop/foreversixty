-- api/internal/db/migrations/0022_ratings.down.sql
drop table if exists kill_duration_digests;
drop table if exists rating_percentile_digests;
drop index if exists rating_scores_stale_idx;
drop index if exists rating_scores_report_idx;
drop index if exists rating_scores_player_idx;
drop table if exists rating_scores;
