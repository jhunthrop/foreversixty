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
        for boss in source.get("bosses") or []:
            bucket.update(boss.get("items", []))
            boss_name = f"{name}: {boss.get('name', '')}"
            index.setdefault(boss_name, set()).update(boss.get("items", []))
    return index


def _quest_level(loot: dict, item_id: int) -> int | None:
    quests = loot.get("quests", {}).get(str(item_id))
    if not quests:
        return None
    return max(q["level"] for q in quests)


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
        quest_level = _quest_level(loot, item_id)
        if quest_level is not None and quest_level - QUEST_FLOOR_OFFSET > band_level:
            result.add(
                severity,
                item_id,
                f"{label} item {item_id} ({row['name']})'s quest floor "
                f"(level {quest_level} - {QUEST_FLOOR_OFFSET}) exceeds band {band_level}",
                ours=str(band_level),
                theirs=str(quest_level - QUEST_FLOOR_OFFSET),
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
            result, "blocker", slot_entry, band_level, faction, items_by_id, factions, loot,
            name_index, f"{label} {slot_entry['slot']}",
        )
        if row is not None:
            rows_by_base_slot[_base_slot(slot_entry["slot"])] = {**row, "slot": slot_entry["slot"]}
            if row.get("unique"):
                unique_counts[item_id] += 1
        for alt in slot_entry.get("alternatives", []):
            _check_pick(
                result, "minor", alt, band_level, faction, items_by_id, factions, loot,
                name_index, f"{label} {slot_entry['slot']} alternative",
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
            _check_band(result, spec, band, items_by_id, factions, loot, name_index)
    return result
