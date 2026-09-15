-- Every kill row now feeds every metric's bracket for its spec, not only the
-- metric its role is sorted by. A rewrite of a fight never re-folds it, so the
-- brackets and the rows are emptied here and refilled by the next parse of each
-- report, which counts as a first write again. Pre-launch data only.
truncate percentile_digests, fight_metrics;
