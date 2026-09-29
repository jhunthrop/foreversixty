"""Constants `pipeline.loot.sources` and `pipeline.loot.wowhead` both
need -- split out to avoid a circular import between them (`sources`
calls into `wowhead` to merge in a scrape; `wowhead` needs the same
`Map.InstanceType` -> kind mapping `sources` uses for the fork's own
drops)."""

from __future__ import annotations

#: `Map.InstanceType`. 3 (battleground) and 4 (arena) are instances whose
#: loot the contract has no kind for -- a battleground's rewards are
#: reputation and rank, which are their own kinds -- so only these two
#: become drop sources.
INSTANCE_KIND: dict[int, str] = {1: "dungeon", 2: "raid"}
