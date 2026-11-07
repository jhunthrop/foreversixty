alter table users
  drop column if exists bnet_import_state_at,
  drop column if exists bnet_import_state;
