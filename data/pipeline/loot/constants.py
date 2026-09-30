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

#: raid-loot-regression lane, 2026-09-29: `WORLD_DROP_BOSS_MIN_CHANCE_
#: PERCENT` (a >= 5% floor a dungeon/raid row also had to clear to keep
#: its own attribution) is retired. Tier armour in cmangos routinely
#: drops from several bosses -- or many trash creatures -- of ONE
#: instance, sometimes at a chance under 1% or none stated at all
#: (unknown, not zero: `ClassicDbSourceRecord.chance`'s own doc); gating
#: a dungeon/raid row's attribution on that chance folded real,
#: intentional boss/instance kills into the item's generic `world_drop`
#: pool instead -- measured on build 1.60.1.70009: raid distinct items
#: fell 767 -> 350 (Molten Core 174 -> 38, Blackwing Lair 324 -> 46,
#: Ahn'Qiraj 178 -> 93, Naxxramas 152 -> 90) once drop-sources-2's own
#: direct-row rule started flagging these items' pooled classification
#: at all. See `is_confirmed_boss_drop`.


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


def is_confirmed_boss_drop(
    is_instance_zone: bool, has_world_drop_record: bool, chance: float | None
) -> bool:
    """Whether one dungeon/raid creature row keeps its own boss/trash
    attribution rather than folding into the item's `world_drop` pool.

    raid-loot-regression lane, 2026-09-29: a creature that spawns in a
    dungeon or raid instance is NEVER a world-pool member on the strength
    of its OWN chance alone -- `WORLD_DROP_BOSS_MIN_CHANCE_PERCENT`'s own
    retirement doc, above, has the measured regression a chance floor by
    itself caused (a tier-armour row folding into the pool because its
    dump-stated chance was low or unknown). An open-world row
    (`is_instance_zone` false) is never confirmed here -- the world-pool
    rule (`is_world_drop_pattern`) owns those.

    pooled-boss-greens lane, 2026-09-29's own addendum: that "whatever
    its chance" reach is too wide for a dungeon/raid row whose item is
    ALSO a known world-drop pool member by another route
    (`has_world_drop_record` -- the item carries its own separate
    `world_drop` `ClassicDbSourceRecord`/wowhead pool, from a signal this
    one row's own creature never triggered on its own). cmangos gives
    many dungeon bosses a "(Boss Loot)" `reference_loot_template` used by
    that ONE boss alone -- e.g. id 35009, Maraudon's Princess Theradras,
    265 rows, one referencing creature -- holding the boss's real drops
    AND ~260 equal-weight random BoE greens at `ChanceOrQuestChance` 0
    (an equal-share group). `_world_drop_pools`' own three signals
    (marked comment, >= `_SHARED_REFERENCE_MAX_USERS` users, multi-map)
    all key off the REFERENCING creature count, so a one-user reference
    like this one is invisible to them even though every item inside is,
    separately, a well-known world drop with its OWN `world_drop` record
    from a differently-shaped (marked/multi-map/fan-out) pool elsewhere
    in the dump -- measured on build 1.60.1.70009: Princess Theradras 273
    items, Spawn of Hakkar 277, Darkmaster Gandling 164, Anvilrage
    Overseer/Warden/Guardsman 107 each, Mangled Cadaver 101, audit
    category C minors 1,817 -> 27,927.
    `WORLD_DROP_MAX_CHANCE_PERCENT` is the same "this could plausibly be
    a normal kill" floor `is_world_drop_pattern` already uses -- a row
    whose OWN stated chance clears it is a real, specific boss kill
    regardless of what else the item is also known for, so it still
    keeps its attribution; only an unknown (`None`, or cmangos' own `0`
    sentinel) or sub-floor chance folds. A folded row needs no new pool:
    the item's own existing `world_drop` source already lists it
    (`classicdb_additions`'s own doc). An item with NO `world_drop`
    record keeps every instance attribution whatever its chance -- the
    raid-tier invariant raid-loot-regression established, unchanged."""
    if not is_instance_zone:
        return False
    if not has_world_drop_record:
        return True
    return chance is not None and chance >= WORLD_DROP_MAX_CHANCE_PERCENT


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
