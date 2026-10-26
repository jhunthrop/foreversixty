alter table guilds add column raid_timezone text;
create table guild_raids (
  id bigserial primary key,
  guild_id bigint not null references guilds(id) on delete cascade,
  zone text not null,
  title text not null check (char_length(title) between 1 and 80),
  starts_at timestamptz not null,
  duration_minutes int not null check (duration_minutes between 60 and 240),
  notes text not null default '' check (char_length(notes) <= 2000),
  status text not null default 'planned' check (status in ('planned','locked','done','cancelled')),
  report_id text references reports(id) on delete set null,
  post_default boolean not null default false,
  created_by bigint not null references users(id) on delete restrict,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);
create index guild_raids_guild_starts_idx on guild_raids (guild_id, starts_at);
create table guild_raid_signups (
  raid_id bigint not null references guild_raids(id) on delete cascade,
  character_key text not null,
  name text not null, class text not null, spec text not null,
  role text not null check (role in ('tank','healer','dps')),
  state text not null check (state in ('in','late','bench','out')),
  note text not null default '' check (char_length(note) <= 140),
  signed_by bigint not null references users(id) on delete restrict,
  updated_at timestamptz not null default now(),
  primary key (raid_id, character_key)
);
