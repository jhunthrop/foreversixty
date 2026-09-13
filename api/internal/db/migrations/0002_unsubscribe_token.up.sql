create extension if not exists pgcrypto;
alter table subscribers add column unsubscribe_token text not null default encode(gen_random_bytes(24), 'hex');
create unique index if not exists subscribers_unsubscribe_token_idx on subscribers (unsubscribe_token);
