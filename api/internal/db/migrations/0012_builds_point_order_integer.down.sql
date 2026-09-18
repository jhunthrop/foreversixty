-- Rolling back after beta-era builds exist fails outright: the smallint cast raises
-- "smallint out of range" for trait node ids above 32767 and leaves the migration
-- dirty. Delete or rewrite those rows first if this down ever has to run.
alter table builds alter column point_order type smallint[] using point_order::smallint[];
