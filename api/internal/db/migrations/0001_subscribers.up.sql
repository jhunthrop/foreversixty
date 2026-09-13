create table if not exists subscribers (
  id bigserial primary key,
  email text not null,
  email_normalized text not null unique,
  token text not null unique,
  created_at timestamptz not null default now(),
  confirmed_at timestamptz,
  unsubscribed_at timestamptz
);
create index if not exists subscribers_confirmed_idx on subscribers (confirmed_at) where confirmed_at is not null;
