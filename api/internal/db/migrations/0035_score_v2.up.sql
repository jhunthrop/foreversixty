-- Forever Score v2 (api/internal/rating/score_v2.go; docs/contracts/addon-v2.md contract 5).
-- Every column is nullable and additive: rows written before this migration have none of
-- them, and a NULL reads as "not known" rather than as a zero.
--
-- discipline: this fight's own fact - whether the player pulled with a flask, a food and
--   every equipped enchantable slot enchanted. Written by Store.RateFight from the fight
--   summary; the score's discipline component is the share of a character's fights where it
--   is true.
-- score, gear_score, experience, attendance, tier: the character's Forever Score as it
--   stood when this fight was rated, rolled up by the rater (Store.RollUpScores) onto the
--   player's row, with contract 5's meanings: gear_score is the gear component in points
--   (0-500), experience the experience points, attendance the fraction (0-1) of guild
--   nights attended. A snapshot, not a history source: a backfill re-rate leaves them NULL.
alter table rating_scores
  add column if not exists discipline boolean,
  add column if not exists score integer,
  add column if not exists gear_score integer,
  add column if not exists experience integer,
  add column if not exists attendance numeric,
  add column if not exists tier text;
