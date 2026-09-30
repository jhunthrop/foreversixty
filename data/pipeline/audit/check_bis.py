"""Check F: every band/faction/slot pick in `bis/*.json` against the rules
`sim/cmd/leveling-bis` itself is supposed to have already enforced --
existence, level, faction, source provenance, two-hand/off-hand and
uniqueness. Re-checking a generator's own output after the fact is exactly
what this audit is for: `pipeline.proficiency` (already the ported form of
`sim/cmd/leveling-bis/eligible.go`'s own armour/weapon tables, per that
module's own doc) is not re-applied item-by-item here because
items/<class-slug>.json is itself already filtered through
`pipeline.proficiency.can_equip` at generation time (`pipeline.normalize.
gear.build_class_items`) -- a bis pick's item id being present in that
file at all is proficiency already having passed once, upstream; the "item
exists" check below is the same fact restated as this category's own check.
"""

from __future__ import annotations

from collections import Counter

from pipeline.audit.context import AuditContext
from pipeline.audit.findings import CategoryResult

CODE = "F"
LABEL = "Leveling BiS vs rules"
PRIMARY_SOURCE = (
    "items/<class>.json (existence + required_level + two_hand; already proficiency-filtered "
    "by pipeline.proficiency at generation time), loot.json (faction restriction + quest floor "
    "+ source-name index), leveling.QuestFloor rule (quest level - 3)"
)

#: `sim/cmd/leveling-bis`'s own leveling.QuestFloor rule, restated here
#: (the brief's own instruction) rather than imported -- this is a Go
#: package the pipeline's Python audit cannot import.
QUEST_FLOOR_OFFSET = 3


def _base_slot(slot: str) -> str:
    return slot[:-1] if slot[-1].isdigit() else slot


def _pvp_labels(source: dict) -> list[str]:
    """Every form of the ranker's pvp label one `pvp` source can appear
    under in a bis pick's `source`: with its title when loot.json states
    one (`rank_title`/`title`), and the title-less form the ranker falls
    back to when it does not."""
    if source.get("kind") != "pvp" or not source.get("rank"):
        return []
    faction = str(source.get("faction") or "").capitalize()
    if not faction:
        return []
    rank = int(source["rank"])
    title = source.get("rank_title") or source.get("title")
    labels = [f"PvP rank {rank} · {faction}"]
    if title:
        labels.append(f"PvP rank {rank} · {title} · {faction}")
    return labels


def _source_name_index(loot: dict) -> dict[str, set[int]]:
    """Human-readable source name -> item ids it lists, built the same way
    `leveling-bis`'s own `source` strings are built (`crafted:<profession>`
    and `rep:*`/`vendor:*` sources by their own `name`; a boss drop by
    `"<instance name>: <boss name>"`) -- what a bis pick's own `source`
    string is checked against.

    A raid/dungeon pick's own `source` is not always the full "<instance>:
    <boss>" form -- `sim/cmd/leveling-bis`'s own report sometimes shortens
    it to just the instance name (observed on the real build: Dark Storm
    Gauntlets, item 21585, source `"Ahn'Qiraj"` alone, even though the
    item is only on one (unnamed) boss's own list). The plain instance name
    key below therefore covers every boss's items too, not just the
    source's own flat `items`/`trash`, so both display conventions match
    without a false "source does not list the item".
    """
    index: dict[str, set[int]] = {}
    for source in loot.get("sources", []):
        name = source.get("name", "")
        bucket = index.setdefault(name, set())
        bucket.update(source.get("items") or [])
        bucket.update(source.get("trash") or [])
        # bis-ranker-integrity-4, 2026-09-30: a pvp pick's `source` is the
        # ranker's own "PvP rank N · <title> · <Faction>" label
        # (`pvpSourceLabel`, sim/cmd/leveling-bis/band.go), not the
        # source's bare `name` ("Rank 9 (Alliance)"), so that label keys
        # the same bucket; a vendor row that inherited the rank shows the
        # same label and its item is on the pvp source too.
        for label in _pvp_labels(source):
            index.setdefault(label, set()).update(source.get("items") or [])
        for boss in source.get("bosses") or []:
            bucket.update(boss.get("items", []))
            boss_name = f"{name}: {boss.get('name', '')}"
            index.setdefault(boss_name, set()).update(boss.get("items", []))
    # quest-faction lane, 2026-09-29: a quest pick's `source` is now the
    # quest's own name (one itemSource per quest in leveling-bis's loader),
    # not the flat "Quests" bucket, so every quest name maps to the items
    # its records reward.
    for item_id, quests in (loot.get("quests") or {}).items():
        for quest in quests:
            index.setdefault(quest.get("name", ""), set()).add(int(item_id))
    return index


def _quest_only_item_ids(loot: dict) -> set[int]:
    """Every item id whose ONLY `loot.json` source bucket, across the
    whole build, is the flat `quest` kind -- day3 data-followups-7 lane,
    2026-09-30's own signal for `_class_excludes_every_quest` below: an
    item with another, independently-verified source is never blocked
    by what one quest's own `RequiredClasses` says."""
    kinds: dict[int, set[str]] = {}
    for source in loot.get("sources", []):
        kind = source.get("kind", "")
        ids = set(source.get("items") or []) | set(source.get("trash") or [])
        for boss in source.get("bosses") or []:
            ids.update(boss.get("items", []))
        for item_id in ids:
            kinds.setdefault(int(item_id), set()).add(kind)
    return {item_id for item_id, item_kinds in kinds.items() if item_kinds == {"quest"}}


def _class_excludes_every_quest(loot: dict, item_id: int, class_slug: str) -> bool:
    """`True` when EVERY quest `loot.json`'s own `quests` map names for
    `item_id` excludes `class_slug` (`QuestSource.classes` -- `None`
    means any class, so a single entry with no restriction, or one that
    admits the class, clears it). `False` (never blocks) for an item
    absent from `quests` entirely."""
    quests = loot.get("quests", {}).get(str(item_id))
    if not quests:
        return False
    return all(
        q.get("classes") is not None and class_slug not in q["classes"] for q in quests
    )


def _quest_floor(loot: dict, item_id: int, faction: str) -> int | None:
    """The level a character of `faction` must reach before ANY quest that
    rewards `item_id` counts as obtainable: the lowest QuestFloor across
    the item's quests open to that faction (sim/leveling's LowestFloor,
    the ranker's own rule), where QuestFloor is the accept level raised to
    the quest level minus the slack when the level is known. Silver Star
    (3463) is rewarded by Stealing Supplies (30/35) AND the Forever quest
    Scramble (14/24): a level-21 character can finish Scramble, so the
    floor is 21, not 32."""
    quests = loot.get("quests", {}).get(str(item_id))
    if not quests:
        return None
    floors = []
    for q in quests:
        # data-followups-3, 2026-09-30: a Forever quest whose side no primary
        # source states publishes "unknown"; the ranker reads that as both,
        # so the audit's floor does too.
        if q.get("faction", "both") not in ("both", "unknown", faction):
            continue
        min_level = int(q.get("min_level") or 0)
        level = int(q.get("level") or 0)
        floors.append(max(min_level, level - QUEST_FLOOR_OFFSET) if level > 0 else min_level)
    return min(floors) if floors else None


def _check_pick(
    result: CategoryResult,
    severity: str,
    pick: dict,
    band_level: int,
    faction: str,
    items_by_id: dict[int, dict],
    factions: dict[str, str],
    loot: dict,
    name_index: dict[str, set[int]],
    label: str,
    class_slug: str,
    quest_only_items: set[int],
) -> dict | None:
    item_id = pick["item_id"]
    result.checked += 1
    row = items_by_id.get(item_id)
    if row is None:
        result.add(
            "blocker" if severity == "blocker" else "minor",
            item_id,
            f"{label} item {item_id} ({pick.get('item_name')}) is not in items/<class>.json "
            "for this spec's class",
        )
        return None
    # Day3 data-followups-7 lane, 2026-09-30: BLOCKER regardless of `severity`
    # (a pick or an alternative) -- a quest-only item every one of whose
    # quests excludes this class (Fire Ruby/Destroy Morphaz's own defect)
    # is not obtainable at all, not merely a soft downgrade.
    if item_id in quest_only_items and _class_excludes_every_quest(loot, item_id, class_slug):
        result.add(
            "blocker",
            item_id,
            f"{label} item {item_id} ({row['name']}) is quest-only and every quest that "
            f"rewards it excludes {class_slug} (classic-db RequiredClasses)",
        )
    if row["required_level"] > band_level:
        result.add(
            severity,
            item_id,
            f"{label} item {item_id} ({row['name']}) required_level {row['required_level']} "
            f"> band {band_level}",
            ours=str(band_level),
            theirs=str(row["required_level"]),
        )
    if pick.get("source_kind") == "quest":
        floor = _quest_floor(loot, item_id, faction)
        if floor is not None and floor > band_level:
            result.add(
                severity,
                item_id,
                f"{label} item {item_id} ({row['name']})'s lowest quest floor "
                f"(quest level - {QUEST_FLOOR_OFFSET}, or the accept level) is {floor}, "
                f"above band {band_level}",
                ours=str(band_level),
                theirs=str(floor),
            )
    item_faction = factions.get(str(item_id))
    if item_faction and item_faction != "both" and item_faction != faction:
        result.add(
            "blocker" if severity == "blocker" else "major",
            item_id,
            f"{label} item {item_id} ({row['name']}) is {item_faction}-only, band is {faction}",
            ours=faction,
            theirs=item_faction,
        )
    source_name = pick.get("source") or ""
    if source_name and item_id not in name_index.get(source_name, set()):
        result.add(
            severity,
            item_id,
            f"{label} item {item_id} ({row['name']})'s source {source_name!r} does not list "
            "the item in loot.json",
        )
    return row


def _check_band(
    result: CategoryResult, spec: str, band: dict, items_by_id: dict[int, dict],
    factions: dict[str, str], loot: dict, name_index: dict[str, set[int]],
    class_slug: str, quest_only_items: set[int],
) -> None:  # fmt: skip
    band_level = band["band"]
    faction = band["faction"]
    label = f"{spec} band {band_level} {faction}"
    rows_by_base_slot: dict[str, dict] = {}
    unique_counts: Counter[int] = Counter()
    for slot_entry in band.get("slots", []):
        item_id = slot_entry.get("item_id")
        if item_id is None:
            continue
        row = _check_pick(
            result,
            "blocker",
            slot_entry,
            band_level,
            faction,
            items_by_id,
            factions,
            loot,
            name_index,
            f"{label} {slot_entry['slot']}",
            class_slug,
            quest_only_items,
        )
        if row is not None:
            rows_by_base_slot[_base_slot(slot_entry["slot"])] = {**row, "slot": slot_entry["slot"]}
            if row.get("unique"):
                unique_counts[item_id] += 1
        for alt in slot_entry.get("alternatives", []):
            _check_pick(
                result,
                "minor",
                alt,
                band_level,
                faction,
                items_by_id,
                factions,
                loot,
                name_index,
                f"{label} {slot_entry['slot']} alternative",
                class_slug,
                quest_only_items,
            )

    main_hand = rows_by_base_slot.get("main_hand")
    off_hand = rows_by_base_slot.get("off_hand")
    if main_hand and main_hand.get("two_hand") and off_hand:
        result.add(
            "major",
            off_hand["id"],
            f"{label}: off_hand ({off_hand['name']}) is filled while main_hand "
            f"({main_hand['name']}) is two-handed",
        )
    for item_id, count in unique_counts.items():
        if count > 1:
            result.add(
                "blocker",
                item_id,
                f"{label}: unique item {item_id} is picked for more than one slot",
            )
    slots_by_name = {s["slot"]: s.get("item_id") for s in band.get("slots", [])}
    for pair in (("finger1", "finger2"), ("trinket1", "trinket2")):
        a, b = slots_by_name.get(pair[0]), slots_by_name.get(pair[1])
        if a is not None and a == b:
            result.add(
                "minor",
                a,
                f"{label}: {pair[0]}/{pair[1]} both pick item {a} -- not a distinct pair",
            )


def check(ctx: AuditContext) -> CategoryResult:
    result = CategoryResult(code=CODE, label=LABEL, primary_source=PRIMARY_SOURCE)
    if not ctx.bis_specs:
        result.skipped = "no bis/*.json for this build"
        return result
    loot = ctx.loot
    factions = loot.get("factions", {})
    name_index = _source_name_index(loot)
    quest_only_items = _quest_only_item_ids(loot)
    for spec, doc in ctx.bis_by_spec.items():
        class_slug = spec.split("-")[0]
        items_by_id = {
            row["id"]: row for row in ctx.items_by_class.get(class_slug, {}).get("items", [])
        }
        if not items_by_id:
            result.add(
                "major", spec, f"bis/{spec}.json exists but items/{class_slug}.json is missing"
            )
            continue
        for band in doc.get("bands", []):
            _check_band(
                result, spec, band, items_by_id, factions, loot, name_index,
                class_slug, quest_only_items,
            )  # fmt: skip
    return result
