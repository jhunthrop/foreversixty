"""Constants `pipeline.loot.sources`, `pipeline.loot.wowhead` and
`pipeline.loot.classicdb` all need -- split out to avoid a circular
import between them (`sources` calls into `wowhead` to merge in a
scrape; `wowhead` needs the same `Map.InstanceType` -> kind mapping
`sources` uses for the fork's own drops; `classicdb` needs the same
world-drop-pattern rule `wowhead` applies to its own scrape, over
classic-db's differently-shaped rows, drop-sources-2 lane 2026-09-29)."""

from __future__ import annotations

from collections.abc import Iterable

#: `Map.InstanceType`. 3 (battleground) and 4 (arena) are instances whose
#: loot the contract has no kind for -- a battleground's rewards are
#: reputation and rank, which are their own kinds -- so only these two
#: become drop sources.
INSTANCE_KIND: dict[int, str] = {1: "dungeon", 2: "raid"}

#: drop-sources-2 lane, 2026-09-29: the world-drop-pattern rule
#: `pipeline.loot.wowhead` applies to a wowhead item page's own
#: `dropped-by` list (see that module's own doc for the full
#: measurement/rationale -- item 4706, 111 distinct creatures across 15
#: zones) and `pipeline.loot.classicdb` now applies identically to
#: classic-db's DIRECT `creature_loot_template` rows for one item (the
#: gap `pipeline.classic_sources._world_drop_pools` cannot close on its
#: own: that classifier only catches a NEGATIVE `mincountOrRef` row
#: pointing at a SHARED `reference_loot_template` id, never a POSITIVE
#: row naming the item directly on each of many creatures' own entries
#: -- items 7909/7910/4306, ~900-1,500 direct per-creature rows each,
#: measured on build 1.60.1.70009's audit). Three independent signals,
#: any ONE of which is enough (a real, single-source drop practically
#: never needs all three to agree): naming this many distinct creatures,
#: spanning this many distinct zones/maps, or -- only when every row
#: states a chance -- every one of them being below the "this could
#: plausibly be someone's normal kill" floor. An unknown chance never
#: counts toward "every chance is low"; this pipeline never guesses one
#: to justify the classification.
WORLD_DROP_MIN_CREATURES = 6
WORLD_DROP_MIN_ZONES = 2
WORLD_DROP_MAX_CHANCE_PERCENT = 1.0

#: A row inside an item classified as a world-drop pattern above this
#: chance, resolving to a dungeon/raid zone, still reads as a real,
#: intentional boss/instance kill (tenet 7's own exception: "a named
#: boss with a chance >= 5% ... keeps its own source alongside") and is
#: kept as its own source next to the pool -- only when BOTH hold; an
#: open-world creature at any chance, or an instance one below the
#: floor, folds into the pool like every other row. See
#: `is_confirmed_boss_drop`.
WORLD_DROP_BOSS_MIN_CHANCE_PERCENT = 5.0


def is_world_drop_pattern(
    rows: Iterable[tuple[int | None, int | None, float | None]],
) -> bool:
    """Whether one item's own (creature_id, zone_or_map_id, chance)
    triples read as a generic world-drop pool rather than a set of real,
    individually-sourced kills -- `WORLD_DROP_MIN_CREATURES`'s own doc
    for the three signals. `pipeline.loot.wowhead` calls this with a
    wowhead `dropped-by` row's `(npc_id, zone_ids[0], chance)`;
    `pipeline.loot.classicdb` calls it with a classic-db direct
    `creature_loot_template` row's `(npc_id, map_id, chance)` -- the
    SAME rule over whichever shape each origin's own rows come in,
    which is the whole point of sharing it here rather than each origin
    keeping its own near-identical copy.
    """
    rows = list(rows)
    if not rows:
        return False
    distinct_creatures = len({creature_id for creature_id, _, _ in rows if creature_id})
    distinct_zones = len({zone_id for _, zone_id, _ in rows if zone_id})
    every_chance_known_and_low = all(
        chance is not None and chance < WORLD_DROP_MAX_CHANCE_PERCENT for _, _, chance in rows
    )
    return (
        distinct_creatures >= WORLD_DROP_MIN_CREATURES
        or distinct_zones >= WORLD_DROP_MIN_ZONES
        or every_chance_known_and_low
    )


def is_confirmed_boss_drop(is_instance_zone: bool, chance: float | None) -> bool:
    """Whether one row inside an item already classified a world-drop
    pattern (`is_world_drop_pattern`) still keeps its own attribution
    rather than folding into the item's synthetic `world_drop` pool --
    `WORLD_DROP_BOSS_MIN_CHANCE_PERCENT`'s own doc: only when the row
    resolves to a dungeon/raid zone AND states a chance at or above that
    floor."""
    return (
        is_instance_zone and chance is not None and chance >= WORLD_DROP_BOSS_MIN_CHANCE_PERCENT
    )


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
