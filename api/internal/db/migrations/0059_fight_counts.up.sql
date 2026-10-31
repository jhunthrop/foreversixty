-- Engine 0.6.0 counts a fight's deaths honestly (logs/engine/fight/counts.go):
-- each hostile NPC that died once, how many of those the logger's group
-- killed, the group's own deaths, and who was in the group for this fight.
-- All four are null on a fight indexed before 0.6.0, which the API answers
-- by leaving the fields out and the report page by showing no kill count.
--
-- npc_kills is no longer written or read: it counted a group kill twice. It
-- stays (default 0) until every running revision has stopped writing it, and
-- a later migration drops it.
alter table fights
  add column mob_deaths    int,
  add column group_kills   int,
  add column group_deaths  int,
  add column group_players text[];
