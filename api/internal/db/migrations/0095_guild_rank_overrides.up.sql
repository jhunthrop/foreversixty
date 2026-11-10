-- An operator's time-boxed guild rank override (api grant-rank / revoke-rank). Officer rank
-- otherwise comes only from the in-game rank index on every sync; an unexpired row here raises
-- that account's guild_members.rank in guilds.RecomputeMembership, and only while the account
-- still has a character in the guild. guild_characters is never touched.
create table guild_rank_overrides (
  guild_id   bigint not null references guilds(id) on delete cascade,
  user_id    bigint not null references users(id) on delete cascade,
  rank       text not null check (rank in ('officer', 'leader')),
  expires_at timestamptz not null,
  note       text not null,
  -- The operator who ran the CLI (no account id: the CLI runs outside any session).
  created_by text not null,
  created_at timestamptz not null default now(),
  primary key (guild_id, user_id)
);
