"""night-item-sources lane, 2026-09-28: folding `pipeline.item_sources`'
wowhead item-page scrape into `loot.json`, for an item the engine fork's
own database (`pipeline.forkdb`) names no source for at all. Split out
of `pipeline.loot.sources` (which was pushing 900 lines) as its own
cohesive unit: item enumeration/placeholder detection
(`unsourced_real_item_ids`, what `python -m pipeline item-sources`
fetches pages for) and the bucket-then-emit/union-merge logic
(`wowhead_additions`/`merge_wowhead_sources`, what `pipeline.loot.
sources.build_loot` folds the scrape's results into). See `pipeline.
wowhead_item_sources`'s own doc for what gets scraped and why only four
of wowhead's item-page listviews are read.
"""

from __future__ import annotations

import json
import re
from collections import defaultdict
from pathlib import Path

from pipeline.item_sources import ItemSourceEntry
from pipeline.loot.constants import (
    INSTANCE_KIND,
    is_confirmed_boss_drop,
    is_world_drop_pattern,
    world_drop_id,
)
from pipeline.models import LootBoss, LootSource, QuestSource
from pipeline.normalize.classes import slugify
from pipeline.wowhead_item_sources import NpcSource

#: A client placeholder row -- night-item-sources' 2026-09-28 measurement:
#: "90 Epic Rogue Dagger", "Bland Dagger", "Copy of X" and the like, none
#: of which are a real item a player will ever loot, buy or craft, so
#: fetching a wowhead page for one would waste the crawler's politeness
#: budget on something that can never resolve to a source. A name-based
#: check catches the two client naming conventions Forever itself uses;
#: `is_placeholder_item` also treats a row with no armor, no weapon
#: damage, no stats, no effect text and no set as a placeholder even when
#: its name does not match, since a real item always has at least one of
#: those.
_PLACEHOLDER_NAME_RE = re.compile(
    r"^(Bland |Copy of |DNT |Template Item|\d+ (Poor|Common|Uncommon|Rare|Epic|Legendary) )"
)

#: src-crawl-order lane, 2026-09-29: wowhead's OWN item-page crawl (the
#: one `pipeline.wowhead_item_sources.fetch_item_sources` walks, and
#: what `unsourced_real_item_ids` used to hand it in whatever order
#: `load_class_item_rows` happened to read the per-class files in, which
#: `fetch_item_sources`/`fetch_missing_from_wowhead` then flattened with
#: a plain ascending `sorted()`) works its backlog in ascending item id.
#: Build 1.60.1.70009's Classic ids top out at 24222 and every
#: Forever-new id starts at 202256, so an ascending-id crawl finishes
#: every Classic id -- 425/440 of the first pages fetched named a source
#: -- long before it ever reaches a Forever-new id, which the fork's own
#: database cannot possibly have itemised and this pipeline exists to
#: fill in. 200,000 sits cleanly in that gap and is used as the
#: Forever-new/Classic boundary throughout this pipeline.
FOREVER_NEW_ID_THRESHOLD = 200_000

#: Sorts after every real Classic `required_level` (max 60) so an item
#: with no required_level on file (0, or the key missing entirely --
#: bind-on-account/token rows commonly have neither) crawls LAST within
#: its own Forever-new/Classic bucket, never ahead of a levelling item
#: this pipeline actually knows the level of.
_UNSET_REQUIRED_LEVEL_RANK = 10_000


def _crawl_priority_key(item_id: int, required_level: int | None) -> tuple[int, int, int]:
    """(a) every Forever-new id before every Classic id, (b) ascending
    `required_level` within each bucket (unset last), (c) item id as a
    final, deterministic tiebreaker. `unsourced_real_item_ids` sorts by
    this; `fetch_missing_from_wowhead`/`fetch_item_sources` then have to
    preserve that order rather than re-sorting ascending by id
    themselves for this to reach the live crawl at all."""
    is_classic_bucket = 0 if item_id >= FOREVER_NEW_ID_THRESHOLD else 1
    level_rank = required_level if required_level else _UNSET_REQUIRED_LEVEL_RANK
    return (is_classic_bucket, level_rank, item_id)


def is_placeholder_item(item_row: dict) -> bool:
    """`item_row` is one entry of `builds/<build>/items/<class>.json`'s
    `items` list (or `items.json` itself, same shape)."""
    if _PLACEHOLDER_NAME_RE.match(item_row.get("name", "")):
        return True
    return (
        not item_row.get("armor")
        and not item_row.get("damage_max")
        and not item_row.get("stats")
        and not item_row.get("effect_text")
        and item_row.get("set_id") is None
    )


def load_class_item_rows(build_dir: Path) -> list[dict]:
    """Every row `build_dir/items/<class>.json` names, deduplicated by id
    (the same item is usable by more than one class and appears in each
    one's file identically) -- the scope `unsourced_real_item_ids` and
    night-item-sources' own measurement both use. `items.json` itself is
    the wrong scope for this: it carries thousands of non-equippable ids
    (quest tokens, reagents, recipe scrolls) `items/<class>.json` never
    lists, that `is_placeholder_item` would misjudge (they carry none of
    armor/damage/stats/effect_text/set_id either, being real but simply
    not gear) and that a wowhead scrape has no reason to chase for a
    leveling-BiS pick pool.
    """
    by_id: dict[int, dict] = {}
    for path in sorted((build_dir / "items").glob("*.json")):
        document = json.loads(path.read_text(encoding="utf-8"))
        for row in document["items"]:
            by_id.setdefault(int(row["id"]), row)
    return list(by_id.values())


def named_items_in_committed_loot(build_dir: Path) -> set[int]:
    """Every item id the build's already-committed `loot.json` names,
    fork- and wowhead-sourced alike -- `python -m pipeline item-sources`'
    own `named` set. Reading the COMMITTED file (rather than re-running
    `build_loot` fresh) is deliberate: it needs no raw/ CSVs and no
    engine checkout, only files this build already has checked in, and
    "already sourced by any means, from any previous night" is exactly
    what should not be re-fetched -- there is no need to tell a fork
    source from a previous night's wowhead one for this purpose, only
    `build_loot`'s own union-not-replace merge cares about that
    distinction. Returns `set()` (nothing excluded) when the build has
    no committed loot.json yet.
    """
    path = build_dir / "loot.json"
    if not path.exists():
        return set()
    document = json.loads(path.read_text(encoding="utf-8"))
    named: set[int] = set()
    for entry in document.get("sources", []):
        named.update(entry.get("items") or [])
        named.update(entry.get("trash") or [])
        for boss in entry.get("bosses") or []:
            named.update(boss.get("items") or [])
    return named


def unsourced_real_item_ids(item_rows: list[dict], named: set[int]) -> list[int]:
    """Real (non-placeholder) item ids in `item_rows` that `named` (the
    set `source_item_ids` names across a fork-only `loot.json`'s sources)
    does not cover -- what `python -m pipeline item-sources` fetches
    wowhead pages for, and what night-item-sources' measurement reports
    as the gap. A placeholder is excluded from BOTH the numerator and the
    denominator: it was never going to have a real source, fork or
    wowhead, so counting it as "unsourced" would overstate the gap this
    pipeline can actually close.

    Ordered by `_crawl_priority_key` (src-crawl-order lane, 2026-09-29):
    every Forever-new id before every Classic id, ascending
    `required_level` within each bucket -- NOT the ascending-item-id
    order `item_rows` arrives in or wowhead's own crawl defaults to. See
    that function's doc for why. Callers that fetch live pages
    (`pipeline.item_sources.fetch_missing_from_wowhead`,
    `pipeline.wowhead_item_sources.fetch_item_sources`) must preserve
    this order rather than re-sorting ascending by id themselves.
    """
    seen: set[int] = set()
    unsourced: list[dict] = []
    for row in item_rows:
        item_id = int(row["id"])
        if item_id in seen or is_placeholder_item(row) or item_id in named:
            continue
        seen.add(item_id)
        unsourced.append(row)
    unsourced.sort(key=lambda row: _crawl_priority_key(int(row["id"]), row.get("required_level")))
    return [int(row["id"]) for row in unsourced]


#: wowhead-world-drops lane, 2026-09-29: a wowhead item page's own
#: `dropped-by` listview names EVERY creature wowhead's community
#: database has ever recorded a BoE world drop from -- for a generic
#: green/white drop that is dozens to hundreds of unrelated trash mobs
#: across a dozen zones, each individually correct (that creature CAN
#: drop it) but together not a source a player recognises (tenet 7:
#: "one line a player recognises... never a dump of every mob"). Item
#: 4706 (Lambent Scale Cloak) measured on build 1.60.1.70009's committed
#: `item-sources.json`: 111 distinct creatures across 15 zones -- the
#: owner's own screenshot of this item's per-creature listing is the
#: report this lane closes. Three independent signals, any ONE of which
#: is enough (a real pool practically never needs all three to agree):
#: naming this many distinct creatures, spanning this many distinct
#: zones/instances (dungeon and open-world alike -- a pool does not stay
#: confined to one), or -- only when the scrape carries a chance for
#: every single row -- every one of them being below the "this could
#: plausibly be someone's normal boss kill" floor. An unknown chance
#: never counts toward "every listed chance is low"; this pipeline never
#: guesses one to justify the classification.
#:
#: The threshold constants and the classification logic itself now live
#: in `pipeline.loot.constants` (`WORLD_DROP_MIN_CREATURES`/
#: `WORLD_DROP_MIN_ZONES`/`WORLD_DROP_MAX_CHANCE_PERCENT`,
#: `is_world_drop_pattern`/`is_confirmed_boss_drop`), drop-sources-2
#: lane 2026-09-29: `pipeline.loot.classicdb` applies the SAME rule to
#: classic-db's own direct `creature_loot_template` rows, and the two
#: origins must stay identical rather than drift apart as two
#: near-copies. raid-loot-regression lane, 2026-09-29:
#: `is_confirmed_boss_drop` exempts a dungeon/raid row from the pool
#: whatever its chance -- the world-pool rule applies to open-world
#: creatures only, so the retired `WORLD_DROP_BOSS_MIN_CHANCE_PERCENT`
#: floor is gone from both origins together.

#: `pipeline.classic_sources._world_drop_records`' own doc: when neither
#: origin states a per-creature level for a world-drop pool, the item's
#: own client `required_level` (to +10) is the honest proxy a player
#: would use for "when do I outlevel this" -- this module's version of
#: that same fallback, for when wowhead's own `dropped-by` rows carry no
#: `minlevel`/`maxlevel` either (every entry fetched before this lane
#: never does: `NpcSource.min_level`/`max_level` postdate it).
WORLD_DROP_LEVEL_RANGE_PAD = 10


def _is_world_drop_pattern(rows: list[NpcSource]) -> bool:
    """Whether one item's own `dropped-by` rows read as a generic
    world-drop pool rather than a set of real, individually-sourced
    kills -- `pipeline.loot.constants.is_world_drop_pattern`, over each
    row's own `(npc_id, zone_ids[0], chance)`; see that function's own
    doc for the three signals."""
    return is_world_drop_pattern(
        (row.npc_id, row.zone_ids[0] if row.zone_ids else None, row.chance) for row in rows
    )


def _split_world_drop_rows(
    rows: list[NpcSource], types: dict[int, int]
) -> tuple[list[NpcSource], list[NpcSource]]:
    """`(boss_rows, pool_rows)` for a `dropped-by` list already
    classified a world-drop pattern -- `pipeline.loot.constants.
    is_confirmed_boss_drop`'s own doc: a row resolving to a dungeon/raid
    zone is a real boss/instance kill and stays its own source, whatever
    its chance; every other row (an open-world zone, or none at all)
    folds into the pool."""
    boss_rows: list[NpcSource] = []
    pool_rows: list[NpcSource] = []
    for row in rows:
        zone_id = row.zone_ids[0] if row.zone_ids else 0
        kind = INSTANCE_KIND.get(types.get(zone_id, 0)) if zone_id else None
        if is_confirmed_boss_drop(kind is not None):
            boss_rows.append(row)
        else:
            pool_rows.append(row)
    return boss_rows, pool_rows


def _world_drop_level_range(
    rows: list[NpcSource], required_level: int | None
) -> tuple[int | None, int | None]:
    """The pool's own level range: the min/max of every row's own
    `min_level`/`max_level` wowhead states, when any row states one; the
    item's own `required_level` (to `required_level +
    WORLD_DROP_LEVEL_RANGE_PAD`) when no row does; `(None, None)`
    (`world_drop:unknown`) when neither this pipeline nor the client
    states a required level for the item either."""
    mins = [row.min_level for row in rows if row.min_level is not None]
    maxes = [row.max_level for row in rows if row.max_level is not None]
    if mins or maxes:
        return (min(mins) if mins else None, max(maxes) if maxes else None)
    if required_level:
        return required_level, required_level + WORLD_DROP_LEVEL_RANGE_PAD
    return None, None


def wowhead_additions(
    item_sources: dict[int, ItemSourceEntry],
    build_items: set[int],
    equippable: set[int],
    zone_names: dict[int, str],
    types: dict[int, int],
    item_factions: dict[int, str],
    required_levels: dict[int, int] | None = None,
    classicdb_world_drop_items: set[int] | None = None,
) -> tuple[list[LootSource], list[int], dict[int, list[QuestSource]]]:
    """Extra sources `pipeline.item_sources`' wowhead scrape names for an
    item the fork database itself named NO source for at all. Every
    `LootSource` this returns carries `source_origin="wowhead"`;
    `merge_wowhead_sources` is what folds them into the fork-derived
    list, unioning items into an existing id rather than duplicating it.

    Mirrors `pipeline.loot.sources`' own `_drop_sources`/`_keyed_sources`
    bucket-then-emit shape, over `ItemSourceEntry` rows instead of the
    fork's `sources` list. A `dropped-by` row's `zone_ids` can name more
    than one zone (a mob that roams); only the FIRST is used, the same
    one-zone-per-drop limit the fork's own shape already has.

    `item_factions` (`pipeline.loot.sources.item_factions`' own result)
    is what a quest reward's `QuestSource.faction` is built from, same
    fix and same reasoning as `pipeline.loot.classicdb.
    classicdb_additions`' own doc: the scrape's own `row.faction` (which
    SIDE'S version of a quest chain this page names) can legitimately
    disagree with the item's own client-stated `factionRestriction` for
    a neutral item two faction-mirrored quest chains both award (item
    1490, "Guardian Talisman", quests 1445/1475, measured while
    regenerating loot.json for 1.60.1.70009) -- `QuestSource.faction`'s
    own contract is always the item's side, every origin agreeing.

    `required_levels` (item id -> the client's own `required_level`) is
    `_world_drop_level_range`'s fallback when wowhead states no
    per-creature level for a world-drop pool. `classicdb_world_drop_items`
    (every item id `pipeline.loot.classicdb.classicdb_additions` already
    put in a `world_drop` source THIS `build_loot` call) is
    wowhead-world-drops lane, 2026-09-29's own precedence rule: an item
    classic-db already classified as `world_drop` is skipped here
    entirely (not even re-bucketed into wowhead's OWN world_drop pool) --
    one `world_drop` source per item after merge, its level range from
    classic-db, never a second wowhead-derived one under a different id
    that `merge_wowhead_sources` (id-keyed) would append rather than
    union.
    """
    classicdb_world_drop_items = classicdb_world_drop_items or set()
    required_levels = required_levels or {}
    bosses: dict[tuple[int, int], set[int]] = defaultdict(set)
    boss_names: dict[int, str] = {}
    world: dict[int, set[int]] = defaultdict(set)
    world_names: dict[int, str] = {}
    zone: dict[int, set[int]] = defaultdict(set)
    world_drop: dict[tuple[int | None, int | None], set[int]] = defaultdict(set)
    vendor_items: dict[int, set[int]] = defaultdict(set)
    vendor_names: dict[int, str] = {}
    crafted: dict[str, set[int]] = defaultdict(set)
    quest: set[int] = set()
    quest_detail: dict[int, list[QuestSource]] = defaultdict(list)

    for item_id, page in item_sources.items():
        if item_id not in build_items:
            continue
        if item_id not in classicdb_world_drop_items and _is_world_drop_pattern(page.dropped_by):
            boss_rows, pool_rows = _split_world_drop_rows(page.dropped_by, types)
            for row in boss_rows:
                zone_id = row.zone_ids[0] if row.zone_ids else 0
                bosses[(zone_id, row.npc_id)].add(item_id)
                boss_names[row.npc_id] = row.name
            if pool_rows:
                level_range = _world_drop_level_range(pool_rows, required_levels.get(item_id))
                world_drop[level_range].add(item_id)
        elif item_id not in classicdb_world_drop_items:
            for row in page.dropped_by:
                zone_id = row.zone_ids[0] if row.zone_ids else 0
                kind = INSTANCE_KIND.get(types.get(zone_id, 0)) if zone_id else None
                if kind is not None:
                    bosses[(zone_id, row.npc_id)].add(item_id)
                    boss_names[row.npc_id] = row.name
                    continue
                if row.npc_id:
                    world[row.npc_id].add(item_id)
                    world_names[row.npc_id] = row.name
                if zone_id:
                    zone[zone_id].add(item_id)
        for row in page.sold_by:
            if not row.npc_id or item_id not in equippable:
                continue
            vendor_items[row.npc_id].add(item_id)
            vendor_names.setdefault(row.npc_id, row.name)
        for row in page.crafted_by:
            crafted[row.profession].add(item_id)
        for row in page.quest_rewards:
            quest.add(item_id)
            quest_detail[item_id].append(
                QuestSource(
                    quest_id=row.quest_id,
                    name=row.name,
                    faction=item_factions.get(item_id, "both"),
                    min_level=row.min_level,
                    level=row.level,
                    level_source="wowhead",
                )
            )

    out: list[LootSource] = []
    for zone_id in sorted(
        {zone_id for zone_id, _ in bosses}, key=lambda z: slugify(zone_names.get(z, str(z)))
    ):
        kind = INSTANCE_KIND[types[zone_id]]
        # A zone id `zone_names` does not have (should not happen for a
        # real dungeon/raid, but `.get` over `[]` matches this module's
        # own "never invent, never crash on a gap" policy elsewhere).
        slug = f"{kind}:{slugify(zone_names.get(zone_id) or str(zone_id))}"
        in_zone = sorted(npc for zone, npc in bosses if zone == zone_id)
        out.append(
            LootSource(
                id=slug,
                kind=kind,
                name=zone_names.get(zone_id, ""),
                zone_id=zone_id,
                bosses=[
                    LootBoss(
                        id=f"{slug}:{npc_id}",
                        name=boss_names.get(npc_id, ""),
                        npc_id=npc_id,
                        items=sorted(bosses[(zone_id, npc_id)]),
                    )
                    for npc_id in in_zone
                ]
                or None,
                source_origin="wowhead",
            )
        )
    out.extend(
        LootSource(
            id=f"world:{slugify(world_names[npc_id])}",
            kind="world",
            name=world_names[npc_id],
            items=sorted(items),
            source_origin="wowhead",
        )
        for npc_id, items in sorted(world.items(), key=lambda pair: slugify(world_names[pair[0]]))
    )
    out.extend(
        LootSource(
            id=f"zone:{zone_id}",
            kind="zone",
            name=zone_names.get(zone_id, ""),
            zone_id=zone_id,
            items=sorted(items),
            source_origin="wowhead",
        )
        for zone_id, items in sorted(zone.items())
    )
    out.extend(
        LootSource(
            id=world_drop_id(level_min, level_max),
            kind="world_drop",
            name="World drop",
            items=sorted(items),
            level_min=level_min,
            level_max=level_max,
            source_origin="wowhead",
        )
        # Sorted low-to-high by level_min (unknown-level buckets, `None`,
        # last), matching `pipeline.loot.classicdb.classicdb_additions`'
        # own `world_drop` sort -- the picker's per-kind list reads the
        # same way regardless of which origin found the pool.
        for (level_min, level_max), items in sorted(
            world_drop.items(),
            key=lambda pair: (
                pair[0][0] is None, pair[0][0] or 0,
                pair[0][1] is None, pair[0][1] or 0,
            ),
        )
    )
    out.extend(
        LootSource(
            id=f"vendor:{npc_id}",
            kind="vendor",
            name=vendor_names[npc_id],
            npc_id=npc_id,
            items=sorted(items),
            source_origin="wowhead",
        )
        for npc_id, items in sorted(vendor_items.items())
    )
    out.extend(
        LootSource(
            id=f"crafted:{profession}",
            kind="crafted",
            name=profession.replace("-", " ").title(),
            profession=profession,
            items=sorted(items),
            source_origin="wowhead",
        )
        for profession, items in sorted(crafted.items())
    )
    return out, sorted(quest), dict(quest_detail)


def _merge_item_dict(base: dict | None, extra: dict | None) -> dict | None:
    """`base`'s own item-keyed dict (`item_chances`/`reitemised_from`)
    with `extra`'s entries folded in for a key `base` does not already
    have -- `base` wins on a shared key (it is the primary side, same
    rule `_union_source`/`_union_boss` apply to everything else), never
    silently dropped the way a plain `base.model_copy` without this
    merge would (src-classicdb lane, 2026-09-29 -- a boss the fork
    already named lost every classic-db drop-chance for its own items
    until this existed, since only `items` was unioned before)."""
    if not base and not extra:
        return None
    return {**(extra or {}), **(base or {})}


def _union_boss(base: LootBoss, extra: LootBoss) -> LootBoss:
    return base.model_copy(
        update={
            "items": sorted(set(base.items) | set(extra.items)),
            "item_chances": _merge_item_dict(base.item_chances, extra.item_chances),
            "reitemised_from": _merge_item_dict(base.reitemised_from, extra.reitemised_from),
        }
    )


def _union_source(base: LootSource, extra: LootSource) -> LootSource:
    """`base` (fork-derived, kept as the primary record -- its own
    `source_origin` is left alone) with `extra`'s (wowhead-derived) item
    ids folded in wherever the two share a bucket. Only called for a
    `base`/`extra` pair that already share an `id`, which -- given how
    both sides build ids (kind + slugified zone/npc/profession/faction
    name) -- means they name the same real-world place, so unioning
    rather than picking one side is correct, not a guess.
    """
    extra_bosses = {(boss.npc_id): boss for boss in (extra.bosses or [])}
    base_bosses = {(boss.npc_id): boss for boss in (base.bosses or [])}
    merged_bosses = [
        _union_boss(base_bosses[npc_id], extra_bosses[npc_id])
        if npc_id in extra_bosses
        else base_bosses[npc_id]
        for npc_id in base_bosses
    ] + [boss for npc_id, boss in extra_bosses.items() if npc_id not in base_bosses]
    return base.model_copy(
        update={
            "items": sorted(set(base.items or []) | set(extra.items or [])) or None,
            "trash": sorted(set(base.trash or []) | set(extra.trash or [])) or None,
            "bosses": sorted(merged_bosses, key=lambda b: b.npc_id) or None,
            "item_chances": _merge_item_dict(base.item_chances, extra.item_chances),
            "reitemised_from": _merge_item_dict(base.reitemised_from, extra.reitemised_from),
        }
    )


def merge_wowhead_sources(base: list[LootSource], extra: list[LootSource]) -> list[LootSource]:
    """`base` (the fork-derived sources `pipeline.loot.sources.build_loot`
    already built) with `extra` (`wowhead_additions`' result) folded in:
    a source id both name gets its items unioned (see `_union_source`)
    and stays fork-primary; a source id only `extra` names is appended
    as-is, `source_origin="wowhead"` intact. Order is preserved for
    `base`; new sources from `extra` are appended (`build_loot`'s own
    final `sources.sort` puts everything in `KIND_ORDER` afterwards, so
    the append order here does not matter).
    """
    by_id = {source.id: index for index, source in enumerate(base)}
    merged = list(base)
    for source in extra:
        index = by_id.get(source.id)
        if index is None:
            by_id[source.id] = len(merged)
            merged.append(source)
        else:
            merged[index] = _union_source(merged[index], source)
    return merged
