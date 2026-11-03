-- Tracked short links (foreversixty.gg/go/<slug>). A link is a slug, a path on our own site
-- and a label for the owner's table. Clicks are one counter per slug per day, people and bots
-- apart; nothing about a visitor is stored. An account records the first link it arrived from,
-- once, at sign-up (users.first_touch_slug), which is all the per-link funnel reads.
create table go_links (
  slug       text primary key check (slug ~ '^[a-z0-9][a-z0-9-]{0,39}$'),
  -- A path on this site only, so a link can never be an open redirect: one leading slash,
  -- never two, no backslash, no whitespace.
  target     text not null check (target ~ '^/([^/\\[:space:]][^\\[:space:]]*)?$' and length(target) <= 200),
  label      text not null check (length(label) between 1 and 80),
  created_at timestamptz not null default now()
);

create table go_link_clicks (
  slug       text not null references go_links (slug) on delete cascade,
  day        date not null,
  clicks     bigint not null default 0 check (clicks >= 0),
  bot_clicks bigint not null default 0 check (bot_clicks >= 0),
  primary key (slug, day)
);

alter table users add column first_touch_slug text references go_links (slug) on delete set null;
create index users_first_touch_idx on users (first_touch_slug) where first_touch_slug is not null;

insert into go_links (slug, target, label) values
  ('reddit', '/', 'Reddit, any subreddit'),
  ('reddit-classicwow', '/', 'Reddit: r/classicwow'),
  ('reddit-wowforever', '/', 'Reddit: r/wowforever'),
  ('reddit-wowclassic', '/', 'Reddit: r/wowclassic'),
  ('reddit-guides', '/guides', 'Reddit post linking the guides'),
  ('reddit-sim', '/sim', 'Reddit post linking the sim'),
  ('discord', '/', 'Discord'),
  ('youtube', '/', 'YouTube'),
  ('curseforge', '/', 'CurseForge'),
  ('wago', '/', 'Wago'),
  ('twitter', '/', 'Twitter')
on conflict (slug) do nothing;

-- The owner's account (production user 78, hunthrop#1894) is the one admin: the links page
-- and its funnel are theirs alone. Matching on the battletag as well keeps this a no-op on
-- every other database.
update users set role = 'admin' where id = 78 and battletag = 'hunthrop#1894';
