-- The beta client's trees use trait node ids above 100,000, which do not
-- fit a smallint. Widen point_order to integer so those builds can save.
alter table builds alter column point_order type integer[] using point_order::integer[];
