-- The raid-day reminder hour (spec 2026-11-08-app-flows F1, appendix B): minutes after local
-- midnight in the account's time_zone, on a 15-minute step, 15:00 by default. raid.signup_nudge
-- goes out at it on the raid day to an account that has not answered.
alter table notification_settings
  add column if not exists raid_reminder smallint not null default 900
    check (raid_reminder between 0 and 1425 and raid_reminder % 15 = 0);
