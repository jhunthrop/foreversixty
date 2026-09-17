"""Armor and stat-amount curves for a client whose ItemSparse states no literal
amounts (see gear.py's ``_row_has_literal_amounts``).

The 1.60 client (Forever beta) computes armour and stat amounts client-side from
curve tables instead of stating them in ``ItemSparse`` columns: ``ItemArmorTotal``,
``ItemArmorQuality``, ``ItemArmorShield``, ``ArmorLocation`` and ``RandPropPoints``.
This module resolves the same numbers from those tables. The formula and every
constant below were verified against build 1.60.1.69893's own tables and 2,830
items whose id, name, item level and quality are unchanged from
``builds/1.15.9.69722``: armour matched on 2,197 of 2,200 real (non-test) armour
pieces (the 3 misses are synthetic QA-named test rows, not player gear), and
stat amounts matched on 1,357 of 1,364 items whose stat *keys* are also
unchanged (99.5%; the rest are ambiguous single-stat items where more than one
slot group rounds to the same integer). See data/README.md for the full check.
"""

from __future__ import annotations

import bisect
from dataclasses import dataclass, field

#: Item.SubclassID values for ClassID == ARMOR (pipeline.proficiency.ARMOR) that
#: this module knows a curve for. See proficiency.py's docstring for the full
#: armour/weapon subclass vocabulary.
CLOTH = 1
LEATHER = 2
MAIL = 3
PLATE = 4
SHIELD = 6

ARMOR_TYPE_COLUMN = {CLOTH: "Cloth", LEATHER: "Leather", MAIL: "Mail", PLATE: "Plate"}
LOCATION_COLUMN = {
    CLOTH: "Clothmodifier",
    LEATHER: "Leathermodifier",
    MAIL: "Chainmodifier",
    PLATE: "Platemodifier",
}

#: A robe (InventoryType 20) has no row of its own in ArmorLocation (every
#: column reads 0): the client scores a robe's armour the same as a chest
#: piece (5). Confirmed against every robe in the 2,830-item cross-check --
#: without this substitution every cloth/leather robe predicts 0 armour.
ROBE_INVENTORY_TYPE = 20
CHEST_INVENTORY_TYPE = 5

#: RandPropPoints has no explicit legendary column; a legendary (OverallQualityID
#: 5) item draws the same stat budget as an epic (4) one. Confirmed against
#: Talisman of Binding Shard/Fragment (17782/17783) -- the only quality-5 armour
#: rows in build 1.60.1.69893 -- whose budget matches EpicF_<group> exactly.
QUALITY_BUDGET_COLUMN = {2: "Good", 3: "Superior", 4: "Epic", 5: "Epic"}

#: InventoryType -> the RandPropPoints stat-budget slot group (0-4). Fit against
#: the 1,364-item stat-key-stable subset of the cross-check: every group here
#: was each inventory type's overwhelming majority (>=86%, most 100%) match.
STAT_BUDGET_GROUP_BY_INVENTORY_TYPE: dict[int, int] = {
    1: 0,
    5: 0,
    7: 0,
    17: 0,
    20: 0,  # head, chest, legs, 2H weapon, robe
    3: 1,
    6: 1,
    8: 1,
    10: 1,
    12: 1,  # shoulder, waist, feet, hands, trinket
    2: 2,
    9: 2,
    11: 2,
    14: 2,
    16: 2,
    23: 2,  # neck, wrist, finger, held off-hand, back, holdable
    13: 3,
    21: 3,
    22: 3,  # one-hand weapon (main hand or off hand)
    15: 4,
    25: 4,
    26: 4,  # ranged, thrown, wand
}

_QUALITY_COLUMNS = range(7)
_GROUP_COLUMNS = range(5)


@dataclass(frozen=True)
class ItemCurves:
    """The five curve tables, each keyed by item level (or InventoryType for
    ArmorLocation), loaded once per normalize run."""

    armor_total: dict[int, dict[str, float]] = field(default_factory=dict)
    armor_quality: dict[int, dict[int, float]] = field(default_factory=dict)
    armor_shield: dict[int, dict[int, float]] = field(default_factory=dict)
    armor_location: dict[int, dict[str, float]] = field(default_factory=dict)
    rand_prop_points: dict[int, dict[str, list[float]]] = field(default_factory=dict)

    @property
    def available(self) -> bool:
        """False when this build's fetch could not supply every curve table.

        Resolving armour from some tables and leaving stats at 0 (or vice
        versa) would be a guess dressed up as data; per the project's rule
        ("only what the client states is emitted"), an incomplete curve set
        resolves nothing rather than half of it.
        """
        return bool(
            self.armor_total
            and self.armor_quality
            and self.armor_shield
            and self.armor_location
            and self.rand_prop_points
        )


def load_item_curves(
    armor_total_rows: list[dict[str, str]],
    armor_quality_rows: list[dict[str, str]],
    armor_shield_rows: list[dict[str, str]],
    armor_location_rows: list[dict[str, str]],
    rand_prop_points_rows: list[dict[str, str]],
) -> ItemCurves:
    armor_total = {
        int(row["ItemLevel"]): {col: float(row[col]) for col in ARMOR_TYPE_COLUMN.values()}
        for row in armor_total_rows
    }
    armor_quality = {
        int(row["ID"]): {q: float(row[f"Qualitymod_{q}"]) for q in _QUALITY_COLUMNS}
        for row in armor_quality_rows
    }
    armor_shield = {
        int(row["ItemLevel"]): {q: float(row[f"Quality_{q}"]) for q in _QUALITY_COLUMNS}
        for row in armor_shield_rows
    }
    armor_location = {
        int(row["ID"]): {col: float(row[col]) for col in LOCATION_COLUMN.values()}
        for row in armor_location_rows
    }
    # Classic Era's RandPropPoints has only the integer Good_0.. columns; the
    # 1.60 client's adds float GoodF_0.. twins carrying the same values (the
    # client evidently switched to float storage at some point). Preferring
    # the float column where it exists, and falling back to the integer one
    # where it does not, resolves the same numbers on both schemas.
    def _budget(row: dict[str, str], budget_column: str, group: int) -> float:
        key = f"{budget_column}F_{group}"
        return float(row[key] if key in row else row[f"{budget_column}_{group}"])

    rand_prop_points = {
        int(row["ID"]): {
            budget_column: [_budget(row, budget_column, g) for g in _GROUP_COLUMNS]
            for budget_column in ("Good", "Superior", "Epic")
        }
        for row in rand_prop_points_rows
    }
    return ItemCurves(
        armor_total=armor_total,
        armor_quality=armor_quality,
        armor_shield=armor_shield,
        armor_location=armor_location,
        rand_prop_points=rand_prop_points,
    )


def _nearest(table: dict[int, object], item_level: int) -> object | None:
    """The value for the largest key <= `item_level`, clamped to the table's own
    range. The real curve tables are gapless from 1 to their max item level, so
    this is an exact lookup in production; a unit-test fixture only needs the
    specific item levels its cases use, not every level in between."""
    if not table:
        return None
    keys = sorted(table)
    index = max(0, min(bisect.bisect_right(keys, item_level) - 1, len(keys) - 1))
    return table[keys[index]]


def resolve_armor(
    curves: ItemCurves,
    item_level: int,
    quality: int,
    inventory_type: int,
    armor_subclass_id: int,
) -> int:
    """The item's armour, or 0 when its subclass carries none (rings, necks,
    trinkets, librams, idols, totems -- see proficiency.py) or a curve table
    has no row for this item level."""
    if armor_subclass_id == SHIELD:
        shield_row = _nearest(curves.armor_shield, item_level)
        if shield_row is None:
            return 0
        return round(shield_row[quality])
    type_column = ARMOR_TYPE_COLUMN.get(armor_subclass_id)
    if type_column is None:
        return 0
    total_row = _nearest(curves.armor_total, item_level)
    quality_row = _nearest(curves.armor_quality, item_level)
    lookup_inventory_type = (
        CHEST_INVENTORY_TYPE if inventory_type == ROBE_INVENTORY_TYPE else inventory_type
    )
    location_row = curves.armor_location.get(lookup_inventory_type)
    if total_row is None or quality_row is None or location_row is None:
        return 0
    base = total_row[type_column]
    quality_modifier = quality_row[quality]
    location_modifier = location_row[LOCATION_COLUMN[armor_subclass_id]]
    return round(base * quality_modifier * location_modifier)


def stat_budget(
    curves: ItemCurves, item_level: int, quality: int, inventory_type: int
) -> float | None:
    """The RandPropPoints stat-point budget for one item, or None when its
    quality or slot carries no budget (poor/common gear the planner never
    reaches, or an inventory type with no equip-stat convention)."""
    budget_column = QUALITY_BUDGET_COLUMN.get(quality)
    group = STAT_BUDGET_GROUP_BY_INVENTORY_TYPE.get(inventory_type)
    if budget_column is None or group is None:
        return None
    budgets_row = _nearest(curves.rand_prop_points, item_level)
    if budgets_row is None:
        return None
    return budgets_row[budget_column][group]
