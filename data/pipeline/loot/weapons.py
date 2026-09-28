"""The fork's own weapon damage, overlaid onto items/<class-slug>.json.

The engine fork's `assets/database/db.json` (`pipeline.forkdb.ForkDatabase`)
carries a `weaponDamageMin`/`weaponDamageMax`/`weaponSpeed` per item it has
itemised by hand -- numbers independent of the client's own `ItemDamage*`
curve tables (`pipeline.normalize.weapon_curves`, which `normalize` already
used to fill every weapon `items/<class-slug>.json` had no damage for). A
30-row spot check against the pinned fork agreed with the curve formula on
29 of 30 (see data/README.md); the one miss (Balanced Fighting Stick, 6215:
derived 15-24, fork 18-21) is real hand-authored data the curve's generic
formula cannot reproduce, not a bug in either side.

This module is `loot`'s, not `normalize`'s, for the same reason
`pipeline/loot/gear.py`'s suffixes and faction restrictions are: nothing but
`loot` reads an engine checkout, and `normalize` alone (no `--engine`) has
no fork database to prefer.
"""

from __future__ import annotations

import json
from pathlib import Path

from pipeline.forkdb import ForkDatabase
from pipeline.models import ClassItems
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


def apply_fork_weapon_damage(build_dir: Path, damage: dict[int, tuple[int, int, float]]) -> int:
    """Overwrite `items/<class-slug>.json`'s damage_min/damage_max/speed/dps
    with the fork's own numbers for every weapon id it itemises.

    Every other row -- `normalize`'s curve-derived value, or an honest zero
    the client's own `ItemDamage*` tables state no dps for -- is left exactly
    as `normalize` wrote it: "keep the derived value only for rows the fork
    lacks" (rotation accuracy program design, item 1). Returns the number of
    (item, class) rows the fork's numbers won, across every class file.

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
            if fork_damage is None:
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
