"""Weapon damage and swing speed, from the client's ItemDamage* curves.

A weapon's whole value to the sim is its damage and its speed, and the 1.60
client (Forever beta) states neither: its ItemSparse has no MinDamage_<n> or
MaxDamage_<n> columns at all, only ItemDelay in milliseconds and DmgVariance.
The damage comes off a curve, exactly as armour does in
`pipeline/normalize/item_curves.py`:

    dps   = ItemDamage<kind>[item level][quality]
    avg   = dps * ItemDelay / 1000
    min   = int(avg * (1 - DmgVariance / 2))     # truncated
    max   = round(avg * (1 + DmgVariance / 2))
    speed = ItemDelay / 1000

Verified against Classic Era's own literal MinDamage_0/MaxDamage_0 on the 478
quality-2-and-better weapons whose id, name, item level, quality and delay are
identical across builds 1.15.9.69722 and 1.60.1.69893: 471 exact (98.5%). The
seven misses are Era rows whose own DmgVariance is the unset 1.0. Truncating
the minimum and rounding the maximum is load-bearing -- rounding both ends
matches 47% of the same set instead of 98.5%.

ItemDamageOneHandCaster and ItemDamageTwoHandCaster are byte-identical to their
non-caster twins on this build, so the caster distinction is not modelled and
those two tables are not fetched.

An older-schema build states the damage in a column; this reads the columns
when the row has them, the per-row branch `gear.py`'s `_row_has_literal_amounts`
makes for stats.
"""

from __future__ import annotations

from dataclasses import dataclass, field

from pipeline.normalize.gear import column_value, int_column
from pipeline.normalize.item_curves import nearest_row

#: Item.SubclassID -> the damage curve the client scores that weapon on, for
#: Item.ClassID 2. Anything absent is melee and uses the one- or two-hand curve
#: chosen by InventoryType. See pipeline/proficiency.py for the vocabulary.
DAMAGE_KIND_BY_SUBCLASS: dict[int, str] = {
    2: "ranged",  # bow
    3: "ranged",  # gun
    16: "thrown",
    18: "ranged",  # crossbow
    19: "wand",
}

TWO_HAND_INVENTORY_TYPE = 17
_MS_PER_SECOND = 1000.0
_QUALITY_COLUMNS = range(7)


@dataclass(frozen=True)
class WeaponCurves:
    """The five ItemDamage tables, each item level -> a DPS per quality."""

    one_hand: dict[int, list[float]] = field(default_factory=dict)
    two_hand: dict[int, list[float]] = field(default_factory=dict)
    ranged: dict[int, list[float]] = field(default_factory=dict)
    wand: dict[int, list[float]] = field(default_factory=dict)
    thrown: dict[int, list[float]] = field(default_factory=dict)

    @property
    def available(self) -> bool:
        """False when this build's fetch could not supply every damage table.

        Resolving two-handers and leaving wands at zero would be a guess
        dressed up as data, so an incomplete set resolves nothing, exactly as
        `ItemCurves.available` does for armour.
        """
        return bool(self.one_hand and self.two_hand and self.ranged and self.wand and self.thrown)

    def table(self, kind: str) -> dict[int, list[float]]:
        return getattr(self, kind)


@dataclass(frozen=True)
class WeaponDamage:
    minimum: float
    maximum: float
    speed: float


def _by_item_level(rows: list[dict[str, str]]) -> dict[int, list[float]]:
    return {
        int(row["ItemLevel"]): [float(row[f"Quality_{q}"]) for q in _QUALITY_COLUMNS]
        for row in rows
    }


def load_weapon_curves(
    one_hand_rows: list[dict[str, str]],
    two_hand_rows: list[dict[str, str]],
    ranged_rows: list[dict[str, str]],
    wand_rows: list[dict[str, str]],
    thrown_rows: list[dict[str, str]],
) -> WeaponCurves:
    return WeaponCurves(
        one_hand=_by_item_level(one_hand_rows),
        two_hand=_by_item_level(two_hand_rows),
        ranged=_by_item_level(ranged_rows),
        wand=_by_item_level(wand_rows),
        thrown=_by_item_level(thrown_rows),
    )


def _damage_kind(subclass_id: int, inventory_type: int) -> str:
    kind = DAMAGE_KIND_BY_SUBCLASS.get(subclass_id)
    if kind is not None:
        return kind
    return "two_hand" if inventory_type == TWO_HAND_INVENTORY_TYPE else "one_hand"


def _row_has_literal_damage(row: dict[str, str]) -> bool:
    """True when this build's ItemSparse states the damage outright (Classic
    Era's shape) rather than leaving it to the curve (the 1.60 client's)."""
    return "MinDamage_0" in row


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
    if _row_has_literal_damage(sparse_row):
        return WeaponDamage(
            minimum=float(int_column(sparse_row, "MinDamage_0")),
            maximum=float(int_column(sparse_row, "MaxDamage_0")),
            speed=speed,
        )
    if not curves.available:
        return None
    inventory_type = int_column(sparse_row, "InventoryType")
    quality = int_column(sparse_row, "OverallQualityID")
    row = nearest_row(
        curves.table(_damage_kind(subclass_id, inventory_type)),
        int_column(sparse_row, "ItemLevel"),
    )
    if row is None or not 0 <= quality < len(row):
        return None
    average = row[quality] * speed
    variance = float(column_value(sparse_row, "DmgVariance"))
    return WeaponDamage(
        minimum=float(int(average * (1 - variance / 2))),
        maximum=float(round(average * (1 + variance / 2))),
        speed=speed,
    )
