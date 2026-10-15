-- Rotation adherence on the addon's last_raid message (docs/contracts/addon-v2.md contract 3;
-- api/internal/rating/adherence.go).
--
-- rating_scores.adherence: the share of the player's ability casts in the fight that are
--   abilities the spec's shipped rotation lists for the character's level band, 0 to 1 to
--   two decimals. NULL when the spec has no shipped rotation, when the player cast fewer
--   than ten abilities, and on every row rated before this migration.
alter table rating_scores add column if not exists adherence numeric;
