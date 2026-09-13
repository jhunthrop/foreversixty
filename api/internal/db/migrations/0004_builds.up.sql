create table if not exists builds (
  id            text primary key,
  class_id      smallint not null,
  race_id       smallint not null,
  tree_version  text not null,
  point_order   smallint[] not null,
  gear          jsonb not null default '{}',
  title         text,
  created_at    timestamptz not null default now(),
  views         bigint not null default 0
);
