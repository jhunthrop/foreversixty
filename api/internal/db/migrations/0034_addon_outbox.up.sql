-- What the addon's outbox (docs/contracts/addon-v2.md, contract 4) leaves behind
-- beyond the export itself.
--
-- addon_sim_requests: one row per in-game sim request that started a run, keyed by
-- the character and the request time the addon wrote. requested_at is the player's
-- local wall time ("2026-11-04 14:10"), kept as text because it carries no zone; the
-- hour-per-character dedupe window reads created_at, the server's own clock. sim_id
-- is the queued sims row the request started.
create table if not exists addon_sim_requests (
  character_key text not null references characters (key) on update cascade on delete cascade,
  requested_at  text not null,
  user_id       bigint not null references users (id) on delete cascade,
  sim_id        text references sims (id) on delete set null,
  created_at    timestamptz not null default now(),
  primary key (character_key, requested_at)
);
create index if not exists addon_sim_requests_recent_idx on addon_sim_requests (character_key, created_at desc);

-- guild_roster_snapshots: the newest guild-tab roster a member's addon saw, one row
-- per guild, for the guild page. members is the contract's list as sent (name,
-- level, spec, gear_gap, build, seen_at); character_key is the character whose
-- export carried it, which belongs to the guild.
create table if not exists guild_roster_snapshots (
  guild_id      bigint primary key references guilds (id) on delete cascade,
  character_key text not null references characters (key) on update cascade on delete cascade,
  members       jsonb not null,
  received_at   timestamptz not null default now()
);
