-- One outbound notification channel per member and kind (Discord now). The
-- webhook's token is the credential, so it is stored sealed (AES-256-GCM);
-- the channel id stays in clear for display.
create table notification_webhooks (
  user_id     bigint      not null references users(id) on delete cascade,
  kind        text        not null,
  channel_id  text        not null,
  token_box   bytea       not null,
  events      jsonb       not null default '{}',
  failures    int         not null default 0,
  last_error  text        not null default '',
  disabled_at timestamptz,
  created_at  timestamptz not null default now(),
  updated_at  timestamptz not null default now(),
  primary key (user_id, kind)
);
