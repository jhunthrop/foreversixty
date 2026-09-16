-- Engine 0.2.5 stops counting friendly fire (a hit on a player's own totem)
-- as damage done, which changes ranked damage figures, and a re-parse never
-- re-folds a fight into the brackets: they are emptied again and refilled by
-- the next parse of each report. Pre-launch data only.
truncate percentile_digests, fight_metrics;
