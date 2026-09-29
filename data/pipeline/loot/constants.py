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


def world_drop_id(level_min: int | None, level_max: int | None) -> str:
    """`world_drop:<level_min>-<level_max>`, or `world_drop:unknown` for a
    pool no origin states a level range for at all -- shared by
    `pipeline.loot.classicdb` (the pinned dump's own `world_drop`
    records) and `pipeline.loot.wowhead` (wowhead-world-drops lane,
    2026-09-29's own pattern-detected pools), so the SAME item's pool
    from either origin lands on the SAME source id and merges into one,
    never a raw reference-template id or npc id, which is an
    implementation detail of whichever origin found it, not a fact the
    site should ever show or key a URL on."""
    if level_min is None or level_max is None:
        return "world_drop:unknown"
    return f"world_drop:{level_min}-{level_max}"
