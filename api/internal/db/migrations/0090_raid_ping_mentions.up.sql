-- Who a raid's "Ping on Discord" last mentioned, per member (D16): a ping mentions only the
-- not-signed raiders not mentioned for that raid in the last six hours, so an officer can mention
-- the healers and later everyone without pinging anyone twice. guild_raid_pings keeps the raid's
-- last ping for the raid page.
create table if not exists guild_raid_ping_mentions (
  raid_id      bigint not null references guild_raids (id) on delete cascade,
  user_id      bigint not null references users (id) on delete cascade,
  mentioned_at timestamptz not null,
  primary key (raid_id, user_id)
);
