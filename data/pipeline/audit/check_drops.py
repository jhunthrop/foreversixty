"""Check C: every dungeon/raid boss and trash drop against cmangos/
classic-db's own loot templates.

The item-level comparison (is the item in this boss's own `creature_loot`,
direct or reference, at the chance we publish; is classic-db's own
generic world-drop pool the real source instead) reads the committed
`raw/classicdb/sources.json` cache alone. The boss-level comparison (does
the npc exist at all; does it spawn on the instance's own map) needs the
raw dump's `creature_template`/`creature` tables, which that cache does not
carry -- pass `--classicdb-dump` a copy of the pinned dump for that half;
without one it is labelled, not skipped outright (the item-level half still
runs).

A "boss" is sometimes really a reward CHEST -- Dire Maul's "Tribute",
Blackrock Depths' "Chest of The Seven" -- whose own `npc_id` is a
fork-internal placeholder with no relation to any classic-db id at all
(not the object's own GAMEOBJECT entry, and its NAME does not match
classic-db's own either -- "Tribute" vs "Gordok Tribute"). Its real loot
lives under `gameobject_template.data1` (a loot id distinct from any
object's own entry), never `creature_loot_template`. `_resolve_chest_loot`
finds it by zone map plus item overlap instead (`ClassicDbDump.
gameobject_chest_loot`/`object_spawn_map`), which also needs
`--classicdb-dump` -- without one these two bosses' items fall back to the
(false) "no such drop" finding, same as the rest of the boss-level half.
"""

from __future__ import annotations

from collections import Counter

from pipeline.audit.context import AuditContext
from pipeline.audit.dumpdb import ClassicDbDump
from pipeline.audit.findings import CategoryResult
from pipeline.loot.wowhead import FOREVER_NEW_ID_THRESHOLD

CODE = "C"
LABEL = "Drops vs classic-db"
PRIMARY_SOURCE = "raw/classicdb/sources.json (creature_loot_template, direct + reference)"

#: Above this many distinct sources for one item, the brief's own signal
#: for "probably over-attributed" (a reitemisation or a merge that fanned a
#: real single source out too far).
MAX_EXPECTED_SOURCES = 3


def _instance_sources(loot: dict) -> list[dict]:
    return [s for s in loot.get("sources", []) if s.get("kind") in ("dungeon", "raid")]


def _boss_existence_and_map(
    dump: ClassicDbDump | None,
    npc_id: int,
    boss_name: str,
    zone_id: int | None,
    zone_map_id: dict[int, int],
    result: CategoryResult,
) -> None:
    if dump is None:
        return
    if npc_id not in dump.creature_names:
        result.add(
            "major",
            npc_id,
            f"boss {boss_name!r} (npc {npc_id}) has no creature_template row in classic-db at all",
            source="classic-db dump: creature_template",
        )
        return
    spawn_map = dump.creature_spawn_map.get(npc_id)
    expected_map = zone_map_id.get(zone_id) if zone_id is not None else None
    if spawn_map is not None and expected_map is not None and spawn_map != expected_map:
        result.add(
            "major",
            npc_id,
            f"boss {boss_name!r} (npc {npc_id}) spawns on map {spawn_map} in classic-db, "
            f"not this source's own zone's map {expected_map}",
            ours=str(expected_map),
            theirs=str(spawn_map),
            source="classic-db dump: creature",
        )


#: drop-sources-2 lane, 2026-09-29: a boss item classic-db has no row for
#: at all is never DELETED (tenet 8: "unverifiable is labelled or left
#: out, never shown as fact" -- this pipeline labels it, via `LootBoss.
#: item_source_origin`, rather than dropping a real fork/wowhead
#: attribution on no more evidence than "classic-db doesn't say so
#: either"). Three reasons this finding is `minor` rather than `major` --
#: any ONE is enough, matching `_boss_item_reason`'s own order:
#:
#: * A Forever-new item id (`FOREVER_NEW_ID_THRESHOLD`) -- classic-db can
#:   never corroborate an id it predates, pre-existing rule.
#: * The instance is raid-gated (`LootSource.opens` is set, "opens
#:   later") -- nobody can verify it against the live game yet either.
#: * The attribution's own origin is `"wowhead"` (`LootBoss.item_source_
#:   origin`) -- already labelled unverified right on the data, so this
#:   finding is confirming a label already there, not surfacing a silent
#:   gap.
#:
#: `major` is left for the one case worth a person's judgment: a
#: `"fork"`-origin (or unlabelled -- the fork's OWN sources predate this
#: lane's `item_source_origin` tagging) attribution, in a LAUNCH
#: (non-raid-gated) instance, that neither classic-db NOR wowhead
#: corroborates.
def _boss_item_reason(
    item_id: int, source: dict, item_source_origin: dict[str, str]
) -> tuple[str, str] | None:
    if item_id >= FOREVER_NEW_ID_THRESHOLD:
        return "minor", " (Forever-new item id)"
    if source.get("opens") is not None:
        return "minor", " (raid-gated instance, opens later -- unverifiable yet)"
    origin = item_source_origin.get(str(item_id))
    if origin == "wowhead":
        return "minor", " (wowhead-only attribution, already labelled unverified)"
    return "major", ""


def _resolve_chest_loot(
    dump: ClassicDbDump | None,
    npc_id: int,
    boss_name: str,
    zone_id: int | None,
    zone_map_id: dict[int, int],
    boss_items: list[int],
) -> dict[int, float] | None:
    """`ClassicDbDump.gameobject_chest_loot`'s own entry for a "boss"
    that is really a reward CHEST (Dire Maul's "Tribute", Blackrock
    Depths' "Chest of The Seven"). Its own `npc_id` is a fork-internal
    placeholder that names no classic-db GAMEOBJECT at all, so this
    instead finds every chest classic-db spawns on the SAME instance map
    whose OWN name contains -- or is contained by -- `boss_name`
    (case-insensitively: "Tribute" is a whole word inside classic-db's
    own "Gordok Tribute"; "Chest of The Seven" matches its own classic-db
    name exactly) and picks whichever one's resolved loot overlaps this
    boss's own item list the most: measured on the pinned dump, Blackrock
    Depths' one name-and-map candidate matches 8 of 8 items, and Dire
    Maul's matches 20 of 25. The name gate is load-bearing, not cosmetic:
    without it, a same-map chest that merely happens to ALSO carry some
    ubiquitous item (Greater Healing Potion, item 1710) produces a false
    "chance disagrees" major on an entirely unrelated, ordinary creature
    boss -- measured while building this fix, 13 of them, on bosses with
    no chest involved at all. Never a guess when no name-and-map
    candidate overlaps at all (`None`, same as no dump).

    The instance map itself comes from `ClassicDbDump.creature_spawn_map.
    get(npc_id)` FIRST, not `zone_map_id.get(zone_id)`: measured on the
    pinned dump, `npc_id` (9034/14324) happens to also be a real classic-
    db CREATURE entry from the SAME zone ("Hate'rel"/"Cho'Rush the
    Observer", both trash mobs in the right instance, not this boss) --
    an accident of how the fork assigns a placeholder id for a
    chest-shaped "boss", but a reliable one across both measured cases,
    and it sidesteps a genuine `zones.json` defect this lane did not
    introduce and is not the one to fix: Blackrock Depths carries a
    SECOND zone row (id 17803, map 230, real) alongside its own zone id's
    row (1584, map 0, a bare continent id) -- and `dungeon:blackrock-
    depths`'s own `zone_id` in `loot.json` is 1584, the wrong one.
    `zone_map_id.get(zone_id)` is still tried as a fallback for a chest
    whose own `npc_id` names no classic-db creature at all."""
    if dump is None or not boss_name:
        return None
    expected_map = dump.creature_spawn_map.get(npc_id)
    if expected_map is None and zone_id is not None:
        expected_map = zone_map_id.get(zone_id)
    if expected_map is None or expected_map in (0, 1):
        return None
    want = set(boss_items)
    wanted_name = boss_name.lower()
    best_entry, best_overlap = None, 0
    for entry, items in dump.gameobject_chest_loot.items():
        if dump.object_spawn_map.get(entry) != expected_map:
            continue
        candidate_name = dump.gameobject_names.get(entry, "").lower()
        names_match = candidate_name and (
            wanted_name in candidate_name or candidate_name in wanted_name
        )
        if not names_match:
            continue
        overlap = len(want & items.keys())
        if overlap > best_overlap:
            best_entry, best_overlap = entry, overlap
    return dump.gameobject_chest_loot[best_entry] if best_entry is not None else None


def _creature_records(
    classic_sources: dict[int, list], item_id: int, npc_id: int
) -> list:
    return [
        r
        for r in classic_sources.get(item_id, [])
        if r.kind == "creature_drop" and r.npc_id == npc_id
    ]


def _boss_items(
    ctx: AuditContext,
    source: dict,
    boss: dict,
    result: CategoryResult,
) -> None:
    npc_id = boss["npc_id"]
    item_chances = boss.get("item_chances") or {}
    item_source_origin = boss.get("item_source_origin") or {}
    classic_sources = ctx.classic_sources_by_item
    boss_items = boss.get("items", [])
    # A "boss" is sometimes really a reward CHEST -- its real loot lives
    # under `gameobject_template.data1`, which `creature_loot_template`
    # never names at all (item-level, drop-sources-2 lane's own
    # follow-up, this lane's brief item 3). Resolved once per boss (not
    # per item: `_resolve_chest_loot` is a zone/map-scoped search, not a
    # per-item lookup) but only ever CONSULTED below for an item with no
    # real `creature_drop` record -- Dire Maul's own "Tribute" measured
    # case: `npc_id` 14324 happens to also be a real classic-db creature
    # ("Cho'Rush the Observer") that genuinely drops 19 of this boss's 39
    # items, so a boss-level "any real match at all" gate would wrongly
    # skip chest resolution for the OTHER 20, which are real chest loot
    # with no creature record at all. `None` (no `--classicdb-dump`, or
    # no same-map chest overlaps this boss's own items) leaves every item
    # to the creature-only check below, same as every other check this
    # module skips without a dump.
    chest_loot = _resolve_chest_loot(
        ctx.dump, npc_id, boss.get("name") or "", source.get("zone_id"), ctx.zone_map_id, boss_items
    )
    for item_id in boss_items:
        result.checked += 1
        records = _creature_records(classic_sources, item_id, npc_id)
        if records:
            # A real `creature_drop` record already verifies this item --
            # never mixed with `chest_loot` (which may itself be resolved
            # from an entirely unrelated same-map object; see this
            # function's own doc), so the comparison below is the SAME
            # one this module always ran before chest resolution existed.
            chances = [r.chance for r in records if r.chance is not None]
            our_chance = item_chances.get(str(item_id))
            if chances and our_chance is not None and abs(max(chances) - our_chance) > 0.01:
                result.add(
                    "major",
                    item_id,
                    f"item {item_id} drop chance on npc {npc_id} disagrees with classic-db",
                    ours=str(our_chance),
                    theirs=str(max(chances)),
                )
        elif chest_loot is not None and item_id in chest_loot:
            # classic-db's own chest loot rows commonly state chance 0
            # for a guaranteed/always-drop item (the same "0 is
            # classic-db's 'unknown' sentinel, never a real 0%" rule
            # `pipeline.loot.classicdb` already applies) -- only a
            # truthy chance is usable for the disagreement check, but
            # membership alone still counts as a match.
            chest_chance = chest_loot[item_id]
            our_chance = item_chances.get(str(item_id))
            if chest_chance and our_chance is not None and abs(chest_chance - our_chance) > 0.01:
                result.add(
                    "major",
                    item_id,
                    f"item {item_id} drop chance on npc {npc_id} disagrees with classic-db",
                    ours=str(our_chance),
                    theirs=str(chest_chance),
                )
        else:
            severity, reason = _boss_item_reason(item_id, source, item_source_origin)
            result.add(
                severity,
                item_id,
                f"item {item_id} is on boss {boss.get('name') or npc_id!r}'s list but classic-db's "
                "creature_loot_template (direct or reference) names no such drop for that npc"
                + reason,
            )
        world_pool = [r for r in classic_sources.get(item_id, []) if r.kind == "world_drop"]
        if world_pool:
            result.add(
                "minor",
                item_id,
                f"item {item_id} is attributed to boss {boss.get('name') or npc_id!r} but "
                "classic-db classifies it as a generic world-drop pool, not specific to that boss",
                source="classic-db dump: reference_loot_template (world-drop-pool classification)",
            )


def _trash(ctx: AuditContext, source: dict, result: CategoryResult) -> None:
    classic_sources = ctx.classic_sources_by_item
    for item_id in source.get("trash") or []:
        result.checked += 1
        kinds = {r.kind for r in classic_sources.get(item_id, [])}
        if not kinds & {"creature_drop", "skinning", "pickpocketing", "object_drop"}:
            severity = "minor" if item_id >= FOREVER_NEW_ID_THRESHOLD else "major"
            result.add(
                severity,
                item_id,
                f"trash item {item_id} in {source['id']} has no creature/object drop record "
                "in classic-db at all"
                + (" (Forever-new item id)" if severity == "minor" else ""),
            )


def _over_attributed(loot: dict, result: CategoryResult) -> None:
    counts: Counter[int] = Counter()
    for source in loot.get("sources", []):
        for boss in source.get("bosses") or []:
            for item_id in boss.get("items", []):
                counts[item_id] += 1
        for item_id in (source.get("trash") or []) + (source.get("items") or []):
            counts[item_id] += 1
    for item_id, count in sorted(counts.items()):
        if count > MAX_EXPECTED_SOURCES:
            result.add(
                "minor",
                item_id,
                f"item {item_id} is attributed to {count} distinct drop sources across "
                f"loot.json (> {MAX_EXPECTED_SOURCES})",
                ours=str(count),
                source="loot.json (source fan-out)",
            )


def check(ctx: AuditContext) -> CategoryResult:
    result = CategoryResult(code=CODE, label=LABEL, primary_source=PRIMARY_SOURCE)
    if not ctx.classic_sources_by_item:
        result.skipped = (
            "raw/classicdb/sources.json has no entries -- boss/trash drop verification skipped"
        )
        return result
    if ctx.dump is None:
        result.skipped = (
            "no --classicdb-dump given: boss npc existence and instance spawn-map checks "
            "skipped (item-level creature_loot/chance checks below do not need it and still ran)"
        )
    loot = ctx.loot
    for source in _instance_sources(loot):
        for boss in source.get("bosses") or []:
            _boss_existence_and_map(
                ctx.dump, boss["npc_id"], boss.get("name", ""), source.get("zone_id"),
                ctx.zone_map_id, result,
            )  # fmt: skip
            _boss_items(ctx, source, boss, result)
        _trash(ctx, source, result)
    _over_attributed(loot, result)
    return result
