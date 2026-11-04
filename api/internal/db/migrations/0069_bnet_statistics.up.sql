-- api/internal/db/migrations/0069_bnet_statistics.up.sql
-- The profile API's character statistics resource (base and effective strength, agility,
-- stamina, intellect, spirit, ...), captured verbatim beside bnet_profile/bnet_equipment/
-- bnet_talents. A Forever race's base stats (Skyborne) are read nowhere else (sim audit
-- round 17).
alter table characters
  add column if not exists bnet_statistics jsonb;
