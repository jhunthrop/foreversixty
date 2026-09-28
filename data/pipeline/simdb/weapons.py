"""Weapon damage and swing speed, from the client's ItemDamage* curves.

A weapon's whole value to the sim is its damage and its speed. The curve
tables, the dispatch by SubclassID/InventoryType and the pure damage formula
live in `pipeline.normalize.weapon_curves` -- this module re-exports
`WeaponCurves`/`load_weapon_curves`/`WeaponDamage` for its existing callers
and adds the one thing that is specific to the simulator's own item rows:
reading ItemSparse's columns (`ItemDelay`, `DmgVariance`, `MinDamage_0`/
`MaxDamage_0`) through `pipeline.normalize.gear`'s row helpers.

An older-schema build states the damage in a column; `weapon_damage` reads
the columns when the row has them (`row_has_literal_weapon_damage`), the
same per-row branch `gear.py`'s `_row_has_literal_amounts` makes for stats.
"""

from __future__ import annotations

from pipeline.normalize.gear import column_value, int_column
from pipeline.normalize.weapon_curves import (
    WeaponCurves,
    WeaponDamage,
    load_weapon_curves,
    resolve_weapon_damage,
    row_has_literal_weapon_damage,
)

__all__ = [
    "WeaponCurves",
    "WeaponDamage",
    "load_weapon_curves",
    "weapon_damage",
]

_MS_PER_SECOND = 1000.0


def weapon_damage(
    sparse_row: dict[str, str],
    subclass_id: int,
    curves: WeaponCurves,
) -> WeaponDamage | None:
    """The row's damage and swing speed, or None when it is not a weapon or the
    curve tables cannot answer."""
    if "ItemDelay" not in sparse_row or "DmgVariance" not in sparse_row:
        # A build whose ItemSparse states neither a delay nor a spread states
        # no weapon at all. There is nothing to resolve and nothing to guess,
        # so this is None rather than an ItemDataError: an absent column is a
        # schema, a present-but-empty one is a truncated row and still raises.
        return None
    delay_ms = int_column(sparse_row, "ItemDelay")
    if delay_ms <= 0:
        return None
    speed = delay_ms / _MS_PER_SECOND
    if row_has_literal_weapon_damage(sparse_row):
        return WeaponDamage(
            minimum=float(int_column(sparse_row, "MinDamage_0")),
            maximum=float(int_column(sparse_row, "MaxDamage_0")),
            speed=speed,
        )
    return resolve_weapon_damage(
        curves,
        subclass_id,
        inventory_type=int_column(sparse_row, "InventoryType"),
        item_level=int_column(sparse_row, "ItemLevel"),
        quality=int_column(sparse_row, "OverallQualityID"),
        speed=speed,
        variance=float(column_value(sparse_row, "DmgVariance")),
    )
