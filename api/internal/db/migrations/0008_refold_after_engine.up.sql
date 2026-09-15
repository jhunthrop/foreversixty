-- Engine 0.2.1 changed every damage figure (overkill is no longer counted), and
-- a re-parse never re-folds a fight into the brackets, so they are emptied again
-- and refilled by the next parse of each report. Pre-launch data only.
truncate percentile_digests, fight_metrics;
