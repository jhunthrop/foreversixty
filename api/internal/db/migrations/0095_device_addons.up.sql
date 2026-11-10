-- The Forever Sixty addons a companion found in its game folder's Interface/AddOns, with each
-- TOC's version, as the companion last reported them (PUT /v1/devices/self/addons). Null until a
-- companion that reports them has; an empty array is a report of none installed. /setup's addon
-- step reads it. Additive: nothing that exists changes.
alter table devices add column addons jsonb;
