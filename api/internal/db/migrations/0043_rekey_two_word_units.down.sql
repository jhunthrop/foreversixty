-- Nothing to undo. The keys 0043 replaced were wrong (half of a two-word name), the half-keyed
-- rows were not kept anywhere, and putting the first word back would re-break every player the
-- migration repaired. Rolling back to 0042 leaves the corrected keys in place.
select 1;
