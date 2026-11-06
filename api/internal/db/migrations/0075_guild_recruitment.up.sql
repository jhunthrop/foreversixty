-- api/internal/db/migrations/0075_guild_recruitment.up.sql
-- Recruitment, applications and private officer notes (design/specs/2026-11-05-web-layout.md
-- §7.2, §7.4; owner decisions D13 and D16). Additive: a running build that does not know these
-- columns and tables keeps working, every new guilds column has a default.

-- The public recruiting card, edited by officers. recruit_wanted is a list of
-- {class, spec?, note?}; the API validates it, the check here only keeps it a list.
-- recruit_discord_role_id is the Discord role the bot grants an accepted applicant, in the
-- guild's own server (guild_discord_installs); null grants nothing.
alter table guilds
  add column if not exists recruit_open boolean not null default false,
  add column if not exists recruit_pitch text not null default ''
    check (char_length(recruit_pitch) <= 500),
  add column if not exists recruit_nights text not null default ''
    check (char_length(recruit_nights) <= 120),
  add column if not exists recruit_wanted jsonb not null default '[]'::jsonb
    check (jsonb_typeof(recruit_wanted) = 'array'),
  add column if not exists recruit_discord_role_id text
    check (recruit_discord_role_id ~ '^[0-9]{17,20}$'),
  add column if not exists recruit_updated_at timestamptz;

-- One application per row. discord_dm is the applicant's opt-in ("Message me on Discord about
-- this application"); the bot DMs the decision only when it is true. role_outcome and dm_outcome
-- record what the bot did on the decision, null until one is made.
create table if not exists guild_applications (
  id            bigserial primary key,
  guild_id      bigint not null references guilds (id) on delete cascade,
  user_id       bigint not null references users (id) on delete cascade,
  character_key text not null,
  message       text not null default '' check (char_length(message) <= 1000),
  availability  text not null default '' check (char_length(availability) <= 200),
  discord_dm    boolean not null default false,
  status        text not null default 'pending'
    check (status in ('pending', 'accepted', 'declined', 'withdrawn')),
  created_at    timestamptz not null default now(),
  decided_by    bigint references users (id) on delete set null,
  decided_at    timestamptz,
  role_outcome  text,
  dm_outcome    text
);
-- One open application per (guild, account).
create unique index if not exists guild_applications_open_idx
  on guild_applications (guild_id, user_id) where status = 'pending';
create index if not exists guild_applications_guild_idx
  on guild_applications (guild_id, created_at desc);
create index if not exists guild_applications_user_idx on guild_applications (user_id);

-- One officer's yes or no on one application; changeable (upsert) and removable (delete).
create table if not exists guild_application_votes (
  application_id bigint not null references guild_applications (id) on delete cascade,
  voter_id       bigint not null references users (id) on delete cascade,
  yes            boolean not null,
  voted_at       timestamptz not null default now(),
  primary key (application_id, voter_id)
);

-- Private officer notes on a member, keyed by the member's account (subject_user_id), which
-- survives a character rename and an alt swap; never read by a member-visible route.
create table if not exists guild_member_notes (
  id              bigserial primary key,
  guild_id        bigint not null references guilds (id) on delete cascade,
  subject_user_id bigint not null references users (id) on delete cascade,
  author_id       bigint references users (id) on delete set null,
  body            text not null check (char_length(body) between 1 and 1000),
  created_at      timestamptz not null default now()
);
create index if not exists guild_member_notes_subject_idx
  on guild_member_notes (guild_id, subject_user_id, created_at desc);
