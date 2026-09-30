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
    for item_id in boss.get("items", []):
        result.checked += 1
        records = [
            r
            for r in classic_sources.get(item_id, [])
            if r.kind == "creature_drop" and r.npc_id == npc_id
        ]
        if not records:
            severity, reason = _boss_item_reason(item_id, source, item_source_origin)
            result.add(
                severity,
                item_id,
                f"item {item_id} is on boss {boss.get('name') or npc_id!r}'s list but classic-db's "
                "creature_loot_template (direct or reference) names no such drop for that npc"
                + reason,
            )
        else:
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
