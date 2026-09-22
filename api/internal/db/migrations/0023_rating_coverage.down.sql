-- api/internal/db/migrations/0023_rating_coverage.down.sql
alter table rating_scores
  drop column if exists insufficient_reason,
  drop column if exists insufficient,
  drop column if exists coverage;
