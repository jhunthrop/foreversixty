drop table if exists realm_ruleset_answers;
alter table characters drop column if exists ruleset_answered;
alter table characters drop column if exists bnet_realm_id;
alter table characters drop constraint if exists characters_ruleset_check;
alter table characters add constraint characters_ruleset_check
  check (ruleset in ('normal', 'pvp', 'rp', 'hardcore')) not valid;
