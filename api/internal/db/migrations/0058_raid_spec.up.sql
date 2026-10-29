-- The spec a character brings to this guild's raids (spec 2026-10-28 §P.10.1): the spec key of
-- data/curated/specs.json ("paladin-holy"), written only by the spec card, Change spec and the
-- raid page's spec select. Null until the player picks one. Per membership, because a
-- character can heal for one guild and tank for another.
alter table guild_characters
  add column raid_spec text check (raid_spec ~ '^[a-z]+(-[a-z]+)+$');
