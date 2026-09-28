"""Weapon damage curve tables and the pure curve math both normalize stages need.

The 1.60 client (Forever beta) states no literal weapon damage in ItemSparse
(no `MinDamage_0`/`MaxDamage_0`, see `row_has_literal_weapon_damage`): damage
comes off a curve, exactly as armour and stats do in `item_curves.py`:

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

ItemDamageOneHandCaster and ItemDamageTwoHandCaster are byte-identical to
their non-caster twins on this build, so the caster distinction is not
modelled and those two tables are not fetched.

Two callers share this module rather than one importing the other:
`pipeline.simdb.weapons` (the simulator's SimItem rows) and
`pipeline.normalize.gear` (`items/<class-slug>.json`, this program's item 1).
`simdb/weapons.py` already imports `column_value`/`int_column` from
`normalize/gear.py`, so `gear.py` importing back from `simdb/weapons.py`
would be circular; putting the curve tables and the pure curve formula here
-- a module neither of those two imports -- lets both read one copy of it.
"""

from __future__ import annotations

from dataclasses import dataclass, field

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

#: The InventoryType that puts a melee weapon on the two-hand curve rather
#: than the one-hand one. Distinct from `gear.TWO_HAND_INVENTORY_TYPES`,
#: which answers a different question (does this item leave the off-hand
#: free); the two happen to agree today because both are 17 in this build.
TWO_HAND_INVENTORY_TYPE = 17

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


def damage_kind(subclass_id: int, inventory_type: int) -> str:
    """Which of `WeaponCurves`' five tables this weapon scores its damage on."""
    kind = DAMAGE_KIND_BY_SUBCLASS.get(subclass_id)
    if kind is not None:
        return kind
    return "two_hand" if inventory_type == TWO_HAND_INVENTORY_TYPE else "one_hand"


def row_has_literal_weapon_damage(row: dict[str, str]) -> bool:
    """True when this build's ItemSparse states the damage outright (Classic
    Era's shape) rather than leaving it to the curve (the 1.60 client's)."""
    return "MinDamage_0" in row


def resolve_weapon_damage(
    curves: WeaponCurves,
    subclass_id: int,
    inventory_type: int,
    item_level: int,
    quality: int,
    speed: float,
    variance: float,
) -> WeaponDamage | None:
    """The curve-resolved damage for one weapon, or None when the curve
    tables cannot answer (unavailable, no row for this item level, or a
    quality past the table's own columns).

    Pure: every value is already read off the row by the caller (`gear.py`'s
    `weapon_fields` and `simdb/weapons.py`'s `weapon_damage` each read
    ItemSparse's own columns their own way -- one through `int_column`/
    `column_value`, the other through the planner's -- so this function
    takes the resolved numbers rather than a row, and stays the one place
    that reads the curve.
    """
    if not curves.available or speed <= 0:
        return None
    row = nearest_row(curves.table(damage_kind(subclass_id, inventory_type)), item_level)
    if row is None or not 0 <= quality < len(row):
        return None
    average = row[quality] * speed
    return WeaponDamage(
        minimum=float(int(average * (1 - variance / 2))),
        maximum=float(round(average * (1 + variance / 2))),
        speed=speed,
    )
