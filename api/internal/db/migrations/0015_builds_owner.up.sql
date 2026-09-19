-- A saved build remembers who saved it, so a signed-in player can list
-- their own (contract 10.6). Nullable: an anonymous save is still saved
-- and still shareable, it simply has no owner.
alter table builds add column if not exists user_id bigint references users(id) on delete set null;
create index if not exists builds_user_idx on builds (user_id, created_at desc);
