-- Forever Score v2 (api/internal/rating/score_v2.go; docs/contracts/addon-v2.md contract 5).
-- Every column is nullable and additive: rows written before this migration have none of
-- them, and a NULL reads as "not known" rather than as a zero.
--
-- discipline: this fight's own fact - whether the player pulled with full consumables
--   (flask, food and, where the role's catalogue lists one, a weapon stone or oil) and every
--   equipped enchantable slot enchanted. Written by Store.RateFight from the fight summary;
--   the score's discipline component is the share of a character's fights where it is true.
-- gear_score, experience, attendance, tier: the character's Forever Score as it stood when
--   this fight was rated, rolled up by the rater (Store.RollUpScores) onto the player's row.
--   gear_score is the raw sims.GearScore of the character's gear, rounded; experience and
--   attendance are those components' points; tier is the tier name.
alter table rating_scores
  add column if not exists discipline boolean,
  add column if not exists gear_score integer,
  add column if not exists experience integer,
  add column if not exists attendance integer,
  add column if not exists tier text;
