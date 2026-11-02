"""The food each spec eats in the published raid preset, the one the starter names.

sim/request/raid-consumes.json (generated from data/curated/presets.json) lists each spec's
consumables by band; the last row is the top band's. presets.json labels its ids.
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


def preset_food(slug: str) -> str:
    """The name of the food the spec eats at its top raid band; fails loudly when it has none."""
    rows = _load(RAID_CONSUMES_PATH)["specs"].get(slug)
    if not rows:
        raise SystemExit(f"{slug!r} has no rows in {RAID_CONSUMES_PATH}")
    foods = [c for c in rows[-1]["consumes"] if c.startswith(FOOD_PREFIX)]
    if len(foods) != 1:
        raise SystemExit(f"{slug!r} eats {foods} at its top raid band, not one food")
    try:
        return _labels()[foods[0]]
    except KeyError:
        raise SystemExit(f"{foods[0]!r} has no label in {PRESETS_PATH}") from None
