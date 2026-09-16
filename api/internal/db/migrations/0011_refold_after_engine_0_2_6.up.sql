-- Engine 0.2.6 measures a closed fight by its own wall length rather than by
-- its last event, which changes every per-second figure a bracket holds, and
-- a re-parse never re-folds a fight into the brackets: they are emptied again
-- and refilled by the next parse of each report. Pre-launch data only.
truncate percentile_digests, fight_metrics;
