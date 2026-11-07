alter table guild_raids drop constraint if exists guild_raids_created_by_fkey;
delete from guild_raids where created_by is null;
alter table guild_raids alter column created_by set not null;
alter table guild_raids add constraint guild_raids_created_by_fkey
  foreign key (created_by) references users (id) on delete restrict;
