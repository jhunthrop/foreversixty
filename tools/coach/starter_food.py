"""The consumables each spec brings in the published raid preset, the ones the starter names.

sim/request/raid-consumes.json (generated from data/curated/presets.json) lists each spec's
consumables by band; the last row is the top band's. presets.json labels its ids. The flask,
the potion and the weapon enchant are read here the same way as the food, never by resource or
weapon style: a cat drinks a Frenzy Potion and brings no flask, a Frost mage opens on a
Spellblasting Potion (round 25).
"""

from __future__ import annotations

import json
from functools import cache
from pathlib import Path
from typing import Any

REPO = Path(__file__).resolve().parent.parent.parent
RAID_CONSUMES_PATH = REPO / "sim" / "request" / "raid-consumes.json"
PRESETS_PATH = REPO / "data" / "curated" / "presets.json"
FOOD_PREFIX = "food_"
FLASK_PREFIX = "flask_"
POTION_SUFFIX = "_potion"
IMBUE_MARK = "_imbue:"


def _load(path: Path) -> Any:
    return json.loads(path.read_text(encoding="utf-8"))


@cache
def _labels() -> dict[str, str]:
    """Every raid consumable's label: the role groups', and the class and spec additions'
    (a spec's own food, such as the tanks' Smoked Desert Dumplings, lives only in by_spec)."""
    consumes = _load(PRESETS_PATH)["presets"]["raid"]["consumes"]
    entries = [entry for group in consumes["groups"].values() for entry in group["ids"]]
    for section in ("by_class", "by_spec"):
        entries += [entry for rows in consumes[section].values() for entry in rows]
    return {entry["id"]: entry["label"] for entry in entries}


def top_band(slug: str) -> list[str]:
    """The consumable ids of the spec's top raid band; fails loudly when it has none."""
    rows = _load(RAID_CONSUMES_PATH)["specs"].get(slug)
    if not rows:
        raise SystemExit(f"{slug!r} has no rows in {RAID_CONSUMES_PATH}")
    return list(rows[-1]["consumes"])


def preset_labels(slug: str, kind: str) -> list[str]:
    """The labels of the top band's flasks, combat potions or weapon enchants, in preset order,
    each once (a stone on both hands is one name). An id with no label (a poison, an imbue spell)
    is not a consumable the starter names."""
    matches = {
        "flask": lambda c: c.startswith(FLASK_PREFIX),
        "combat_potion": lambda c: c.endswith(POTION_SUFFIX),
        "weapon_enchant": lambda c: IMBUE_MARK in c,
    }[kind]
    labels = [_labels().get(c) for c in top_band(slug) if matches(c)]
    return list(dict.fromkeys(label for label in labels if label))


def preset_food(slug: str) -> str:
    """The name of the food the spec eats at its top raid band; fails loudly when it has none."""
    foods = [c for c in top_band(slug) if c.startswith(FOOD_PREFIX)]
    if len(foods) != 1:
        raise SystemExit(f"{slug!r} eats {foods} at its top raid band, not one food")
    try:
        return _labels()[foods[0]]
    except KeyError:
        raise SystemExit(f"{foods[0]!r} has no label in {PRESETS_PATH}") from None
