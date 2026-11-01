"""The fork's own weapon damage, overlaid onto items/<class-slug>.json.

The engine fork's `assets/database/db.json` (`pipeline.forkdb.ForkDatabase`)
carries a `weaponDamageMin`/`weaponDamageMax`/`weaponSpeed` per item it has
itemised by hand -- numbers independent of the client's own `ItemDamage*`
curve tables (`pipeline.normalize.weapon_curves`, which `normalize` already
used to fill every weapon `items/<class-slug>.json` had no damage for). The fork's
numbers are Classic Era's, so they only fill a weapon `normalize` could not
resolve (`apply_fork_weapon_damage`).

This module is `loot`'s, not `normalize`'s, for the same reason
`pipeline/loot/gear.py`'s suffixes and faction restrictions are: nothing but
`loot` reads an engine checkout, and `normalize` alone (no `--engine`) has
no fork database to prefer.
"""

from __future__ import annotations

import json
from pathlib import Path

from pipeline.forkdb import ForkDatabase
from pipeline.models import ClassItems, GearItem
from pipeline.normalize import write_model


def fork_weapon_damage(fork: ForkDatabase) -> dict[int, tuple[int, int, float]]:
    """Item id -> (damage_min, damage_max, speed), for every fork item that
    itemises real weapon damage.

    A fork row with `weaponType` set but no damage (403 of the pinned
    fork's 7,553 items) is a real weapon the fork has not itemised yet --
    exactly the gap the curve derivation exists to fill, so it is left out
    of this dict rather than mapped to (0, 0, 0.0); `apply_fork_weapon_damage`
    below only overwrites a row this dict actually names.
    """
    damage: dict[int, tuple[int, int, float]] = {}
    for row in fork.items:
        damage_min, damage_max, speed = (
            row.get("weaponDamageMin"),
            row.get("weaponDamageMax"),
            row.get("weaponSpeed"),
        )
        if damage_min and damage_max and speed:
            damage[int(row["id"])] = (int(damage_min), int(damage_max), float(speed))
    return damage


def _has_resolved_damage(item: GearItem) -> bool:
    """True when `normalize` already gave the row real damage (the client's
    curve, wowhead Forever or classic-db), which the fork must not replace."""
    return item.damage_min > 0 and item.damage_max > 0


def apply_fork_weapon_damage(build_dir: Path, damage: dict[int, tuple[int, int, float]]) -> int:
    """Fill `items/<class-slug>.json`'s damage_min/damage_max/speed/dps with
    the fork's own numbers for every weapon id it itemises that `normalize`
    left without damage.

    The fork's db.json is Classic Era's itemisation, not Forever's. Until
    sim-audit round 4 it won over every row, and 225 weapons on 1.60.1.70707
    showed Era damage the sim never swings (Elemental Mage Staff 147 - 221
    against Forever's 109 - 164; Thunderfury 60 - 145 against 44 - 115). A
    row `normalize` resolved keeps its value; the fork fills only the zeros.
    Returns the number of (item, class) rows the fork's numbers filled.

    Read-modify-write through `ClassItems`/`GearItem` rather than the raw
    dicts, the same discipline `pipeline.loot.gear.apply_fork_columns` uses
    for `items.json`, so a build the fork does not change re-writes the
    identical bytes.
    """
    won = 0
    items_dir = build_dir / "items"
    for path in sorted(items_dir.glob("*.json")):
        record = ClassItems(**json.loads(path.read_text(encoding="utf-8")))
        updated = []
        for item in record.items:
            fork_damage = damage.get(item.id)
            if fork_damage is None or _has_resolved_damage(item):
                updated.append(item)
                continue
            damage_min, damage_max, speed = fork_damage
            dps = round((damage_min + damage_max) / 2 / speed, 2) if speed > 0 else 0.0
            updated.append(
                item.model_copy(
                    update={
                        "damage_min": damage_min,
                        "damage_max": damage_max,
                        "speed": speed,
                        "dps": dps,
                    }
                )
            )
            won += 1
        write_model(record.model_copy(update={"items": updated}), path)
    return won
