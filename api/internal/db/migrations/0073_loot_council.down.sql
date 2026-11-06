drop table if exists loot_interest;
drop table if exists loot_votes;
alter table loot_awards
  drop column if exists votes_officers,
  drop column if exists votes_no,
  drop column if exists votes_yes,
  drop column if exists reason,
  drop column if exists ms_os;
