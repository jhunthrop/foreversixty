-- One row per signed-in account per UTC day it touched the product, with which surfaces it
-- touched that day as bits: 1 the site (a browser session's API call), 2 the addon or the
-- desktop companion (a companion device token), 4 the phone app (an app device token). The
-- API's auth middleware writes it (internal/engagement), at most once per account, day and
-- surface per instance. Nothing else about the request is kept: no path, no address, no time.
-- The owner's /admin/activity reads it.
create table if not exists user_active_days (
  user_id    bigint not null references users (id) on delete cascade,
  day        date not null,
  surfaces   smallint not null check (surfaces between 1 and 7),
  -- true while the row holds only what the backfill below could reconstruct; the first live
  -- write of the day clears it. The page marks days before the first live row as partial.
  backfilled boolean not null default false,
  primary key (user_id, day)
);
create index if not exists user_active_days_day_idx on user_active_days (day);

-- The backfill: what history still says about who was here on which day, always a floor.
-- Sessions last weeks, so a returning visitor with a live session left no trace; addon_exports
-- keeps only each character's latest capture; devices keep their pairing day and last use.
-- Import sims are left out: the nightly Battle.net refresh starts them without the player.
insert into user_active_days (user_id, day, surfaces, backfilled)
select user_id, day, bit_or(surface)::smallint, true
from (
  select id as user_id, (created_at at time zone 'UTC')::date as day, 1 as surface from users
  union all
  select user_id, (created_at at time zone 'UTC')::date, 1 from sessions
  union all
  select s.user_id, (s.created_at at time zone 'UTC')::date, 1 from sims s
  where s.user_id is not null
    and not exists (select 1 from character_sims c where c.sim_id = s.id and c.kind in ('import', 'addon'))
  union all
  select user_id, (created_at at time zone 'UTC')::date, 2 from character_sims where kind = 'addon'
  union all
  select c.user_id, (cs.created_at at time zone 'UTC')::date, 2
  from character_syncs cs join characters c on c.key = cs.character_key
  where cs.source = 'addon' and c.user_id is not null
  union all
  select user_id, (seen_at at time zone 'UTC')::date, 2 from addon_exports
  union all
  select user_id, (created_at at time zone 'UTC')::date, case kind when 'app' then 4 else 2 end from devices
  union all
  select user_id, (last_seen_at at time zone 'UTC')::date, case kind when 'app' then 4 else 2 end
  from devices where last_seen_at is not null
) src
group by user_id, day
on conflict (user_id, day) do nothing;
