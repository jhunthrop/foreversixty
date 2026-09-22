-- api/internal/db/migrations/0023_rating_coverage.up.sql
-- Spec §1.5's coverage ruling (dated 2026-09-21, found live in production): a card built
-- from too little of a role's own weight table published an overall that read as a
-- judgement resting on almost nothing. Card gains Coverage (the fraction of the role's
-- weight actually scored), Insufficient and InsufficientReason; 0022_ratings is already
-- live, so these land as nullable additive columns rather than a rewrite of that
-- migration. Nullable because a row written by a pre-coverage model version has none of
-- the three to backfill without a recompute -- the rating-backfill job re-rates it under
-- the bumped DefaultModelVersion instead (staleFights already selects on model_version, so
-- bumping the constant is the only change that mechanism itself needs).

alter table rating_scores
  add column if not exists coverage numeric,
  add column if not exists insufficient boolean,
  add column if not exists insufficient_reason text;
