-- api/internal/db/migrations/0025_bnet_capture.down.sql
alter table characters drop column if exists bnet_captured_at;
alter table characters drop column if exists bnet_equipment;
alter table characters drop column if exists bnet_profile;
alter table characters drop column if exists bnet_account;
alter table characters drop column if exists last_login_at;
alter table characters drop column if exists equipped_item_level;
alter table characters drop column if exists average_item_level;
alter table characters drop column if exists gender;
alter table characters drop column if exists race;
