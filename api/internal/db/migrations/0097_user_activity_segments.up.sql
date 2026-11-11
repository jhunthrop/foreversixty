-- An Activity row's styled sentence, as the feed's rows carry it (GET /v1/me/activity
-- rows[].segments): names in class colour, items with quality and id. crest is the class of
-- the other player the row is about; figure the row's number (a sim's DPS). All null on a row
-- without them and on every row written before.
alter table user_activity
  add column segments jsonb check (segments is null or jsonb_typeof(segments) = 'array'),
  add column crest    text,
  add column figure   double precision;
