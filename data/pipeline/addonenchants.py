"""The enchant each spec's BiS picks wear, by slot, for the addon's Gear tab.

The BiS ranker's raid pass wears data/curated/raid-enchants.json's plan (a role profile,
overlaid by the spec's own slots; a null slot takes none: sim/cmd/leveling-bis/enchants.go's
parseRaidEnchants). The Gear tab names that enchant where the player's slot has none, or a
weaker one: "No enchant: Enchant Chest - Greater Stats".

Data.lua carries two tables (pipeline.addonlua._enchants_lua):

* `enchantPlans[spec][key] = index`, key one of PLAN_KEYS (main_hand and off_hand for a
  one-hander, two_hand for a two-hander in the main hand, shield for a shield);
* `enchantRecords[index] = "id|name|icon|spellId|itemId|requiredLevel|stat=amount,..."`, one
  per distinct (enchant, name) the plans use, so a record is stored once.

The name is enchants.json's row whose slots hold the key's slot (one id names one enchant
per slot: 904 is "Enchant Necklace - Agility" on a neck, "Enchant Boots - Agility" on feet).
The required level is the applied item's (an armor kit, a scope, an arcanum), 0 for an
enchanter's spell. The stats are in the addon's stat words (STAT_WORDS), so the Gear tab can
weigh the plan's enchant against the one worn.
"""

from __future__ import annotations

import json
from pathlib import Path

#: The plan's keys (enchants.go's enchantPlanKeys) -> the slot whose enchants.json name it takes.
PLAN_KEYS: dict[str, str] = {
    "head": "head",
    "neck": "neck",
    "shoulder": "shoulder",
    "back": "back",
    "chest": "chest",
    "wrist": "wrist",
    "hands": "hands",
    "waist": "waist",
    "legs": "legs",
    "feet": "feet",
    "ranged": "ranged",
    "main_hand": "main_hand",
    "off_hand": "off_hand",
    "two_hand": "main_hand",
    "shield": "off_hand",
}

#: enchants.json's stat names -> the addon's (Rules.STAT_KEYS' values) where they differ.
STAT_WORDS: dict[str, str] = {"healing": "healing_power", "bonus_armor": "armor"}

RECORD_SEPARATOR = "|"
STAT_SEPARATOR = ","


class AddonEnchantError(SystemExit):
    """raid-enchants.json or enchants.json says something the addon will not publish."""


def _plan(file: dict, spec: str) -> dict[str, int]:
    """One spec's plan: its profile overlaid by its own slots, a null slot dropped."""
    entry = file["specs"][spec]
    profile = file["profiles"].get(entry.get("profile"))
    if profile is None:
        raise AddonEnchantError(
            f"raid-enchants: {spec} names unknown profile {entry.get('profile')!r}"
        )
    merged = {**profile, **(entry.get("slots") or {})}
    plan: dict[str, int] = {}
    for key, value in sorted(merged.items()):
        if key not in PLAN_KEYS:
            raise AddonEnchantError(f"raid-enchants: {spec}: {key!r} is not a plan key")
        if value is not None:
            plan[key] = int(value["id"])
    return plan


def _stats(stats: dict) -> str:
    parts = []
    for name, amount in sorted(stats.items()):
        value = int(amount) if float(amount).is_integer() else amount
        parts.append(f"{STAT_WORDS.get(name, name)}={value}")
    return STAT_SEPARATOR.join(parts)


def _record(rows: list[dict], enchant_id: int, slot: str, levels: dict[int, int]) -> str:
    """The packed record of `enchant_id` as `slot` names it."""
    named = [row for row in rows if row["id"] == enchant_id]
    row = next((row for row in named if slot in row.get("slots", [])), None)
    if row is None:
        raise AddonEnchantError(f"enchants.json has no enchant {enchant_id} for {slot}")
    item_id = int(row.get("item_id") or 0)
    fields = [
        str(enchant_id),
        row["name"],
        row.get("icon") or "",
        str(int(row.get("spell_id") or 0)),
        str(item_id),
        str(levels.get(item_id, 0)),
        _stats(row.get("stats") or {}),
    ]
    for field in fields:
        if RECORD_SEPARATOR in field:
            raise AddonEnchantError(f"enchant {enchant_id}: {field!r} holds {RECORD_SEPARATOR!r}")
    return RECORD_SEPARATOR.join(fields)


def _required_levels(root: Path, build: str) -> dict[int, int]:
    path = root / build / "items.json"
    if not path.exists():
        return {}
    return {
        row["id"]: int(row.get("required_level") or 0)
        for row in json.loads(path.read_text(encoding="utf-8"))
    }


def build_enchant_plans(
    root: Path, build: str, curated_dir: Path, specs: list[str]
) -> tuple[dict[str, dict[str, int]], list[str]]:
    """(spec -> plan key -> 1-based record index, the records) for every spec in `specs`
    raid-enchants.json names. Empty for a build with no enchants.json."""
    enchants_path = root / build / "enchants.json"
    curated_path = curated_dir / "raid-enchants.json"
    if not enchants_path.exists() or not curated_path.exists():
        return {}, []
    rows = json.loads(enchants_path.read_text(encoding="utf-8"))
    file = json.loads(curated_path.read_text(encoding="utf-8"))
    levels = _required_levels(root, build)
    records: list[str] = []
    index: dict[str, int] = {}
    plans: dict[str, dict[str, int]] = {}
    for spec in sorted(set(specs) & set(file.get("specs", {}))):
        plan: dict[str, int] = {}
        for key, enchant_id in _plan(file, spec).items():
            record = _record(rows, enchant_id, PLAN_KEYS[key], levels)
            if record not in index:
                records.append(record)
                index[record] = len(records)
            plan[key] = index[record]
        plans[spec] = plan
    return plans, records
