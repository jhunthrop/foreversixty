update users set bnet_import_state = 'done' where bnet_import_state = 'partial';
alter table users drop constraint if exists users_bnet_import_state_check;
alter table users
  add constraint users_bnet_import_state_check
    check (bnet_import_state in ('running', 'done', 'failed')),
  drop column if exists bnet_import_unresolved;
