-- The deleted classic1x rows are not restored; only the index is undone.
drop index if exists characters_region_bnet_id_key;
