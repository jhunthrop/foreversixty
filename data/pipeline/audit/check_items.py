"""Check A: every item in items/<class>.json against the client's own raw
tables (`raw/Item.csv` + `raw/ItemSparse.csv`) -- required_level, item_level,
quality, two_hand, icon always; armor/stats/damage/speed only where the
client states them literally (this build's Forever beta client instead
computes them from curve tables the audit does not carry, so a curve-based
item's stats are labelled unverifiable here, never guessed at -- see
`pipeline.normalize.gear.resolve_item_values`'s own doc).
"""

from __future__ import annotations

from pipeline.audit.context import AuditContext
from pipeline.audit.findings import CategoryResult
from pipeline.icons import PLACEHOLDER_ICON
from pipeline.normalize.gear import (
    ItemDataError,
    int_column,
    is_weapon_row,
    resolve_item_values,
    weapon_fields,
)
from pipeline.normalize.weapon_curves import row_has_literal_weapon_damage

CODE = "A"
LABEL = "Items vs client"
PRIMARY_SOURCE = "raw/Item.csv + raw/ItemSparse.csv (client DB2 export)"


def _row_has_literal_amounts(row: dict[str, str]) -> bool:
    # `pipeline.normalize.gear._row_has_literal_amounts`'s own rule: Classic
    # Era's ItemSparse states flat Resistances_0/StatModifier_* columns;
    # the 1.60 (Forever beta) client's curve-only shape drops the column
    # entirely rather than zeroing it, so presence is what tells the two
    # shapes apart, not the value.
    return "Resistances_0" in row


def _icons_and_required_level_zero(ctx: AuditContext, result: CategoryResult) -> None:
    planner = ctx.wowhead_gear_planner
    zero_required_level = 0
    zero_with_planner_level = 0
    placeholder = 0
    for class_slug, doc in ctx.items_by_class.items():
        for row in doc.get("items", []):
            result.checked += 1
            item_id = int(row["id"])
            if row.get("icon") == PLACEHOLDER_ICON:
                placeholder += 1
                result.add(
                    "minor",
                    item_id,
                    f"{row['name']!r} ({class_slug}) keeps the placeholder icon "
                    f"{PLACEHOLDER_ICON!r}",
                    ours=row["icon"],
                    source="items/<class>.json vs pipeline.icons.PLACEHOLDER_ICON",
                )
            if row.get("required_level") == 0:
                zero_required_level += 1
                if planner is not None and item_id in planner and planner[item_id].required_level:
                    zero_with_planner_level += 1
                    result.add(
                        "major",
                        item_id,
                        f"{row['name']!r} ({class_slug}) is required_level 0 but wowhead's "
                        "Forever gear-planner payload names a real level for it",
                        ours="0",
                        theirs=str(planner[item_id].required_level),
                        source="raw/wowhead-gear-planner.js",
                    )
    if zero_required_level:
        result.add(
            "minor",
            "*",
            f"{zero_required_level} item(s) across every class are required_level 0"
            + (
                f"; {zero_with_planner_level} of those already have their own finding above "
                "because wowhead names a real level"
                if planner is not None
                else " (wowhead gear-planner payload not available locally to cross-check)"
            ),
        )


def _compare_against_raw(ctx: AuditContext, result: CategoryResult) -> None:
    sparse_by_id = {int_column(row, "ID"): row for row in ctx.raw_item_sparse}
    item_by_id = {int_column(row, "ID"): row for row in ctx.raw_item}
    for class_slug, doc in ctx.items_by_class.items():
        for ours in doc.get("items", []):
            item_id = int(ours["id"])
            sparse = sparse_by_id.get(item_id)
            item_row = item_by_id.get(item_id)
            if sparse is None or item_row is None:
                # normalize-levels lane, 2026-09-29: a wowhead-supplement row
                # (required_level_source "wowhead" on an item with no client
                # row at all -- see pipeline.wowhead_items.to_gear_item) was
                # never going to be in the client's own raw tables; that is
                # what makes it a supplement, not a defect the way a real
                # client id going missing would be. Downgraded to "minor"
                # ("unverified against the client") rather than "major".
                result.add(
                    "minor" if ours.get("required_level_source") == "wowhead" else "major",
                    item_id,
                    f"{ours['name']!r} ({class_slug}) is in items/{class_slug}.json but not in "
                    "the build's own raw ItemSparse.csv/Item.csv",
                )
                continue
            result.checked += 1
            try:
                for field, column in (
                    ("item_level", "ItemLevel"),
                    ("quality", "OverallQualityID"),
                ):
                    theirs = int_column(sparse, column)
                    if ours.get(field) != theirs:
                        result.add(
                            "blocker",
                            item_id,
                            f"{ours['name']!r} ({class_slug}) {field} disagrees with the client",
                            ours=str(ours.get(field)),
                            theirs=str(theirs),
                        )
                _compare_required_level(ours, sparse, item_id, class_slug, result)
                _compare_weapon(ours, sparse, item_row, item_id, class_slug, result)
                _compare_stats(ours, sparse, item_row, item_id, class_slug, result)
            except ItemDataError as error:
                result.add(
                    "minor",
                    item_id,
                    f"{ours['name']!r} ({class_slug}) raw row could not be read: {error}",
                )


def _compare_required_level(ours, sparse, item_id, class_slug, result) -> None:
    """`required_level`'s own comparison, separated from the other raw-column
    fields above: `pipeline.normalize.gear.resolve_required_level` (normalize-
    levels lane, 2026-09-29) can deliberately publish a level the client's own
    `RequiredLevel` column does not state -- wowhead's or the item-level
    proxy's, for a client row that states 0 -- which is the fix, not a
    disagreement with the client. Only a `"client"`-sourced row is checked
    against the raw column as a blocker; the other two sources get a `minor`
    finding instead, naming the source, so the report still says the number
    is unverified against the client rather than silently agreeing with it.
    """
    theirs = int_column(sparse, "RequiredLevel")
    ours_level = ours.get("required_level")
    if ours_level == theirs:
        return
    source = ours.get("required_level_source")
    if source == "client":
        result.add(
            "blocker",
            item_id,
            f"{ours['name']!r} ({class_slug}) required_level disagrees with the client",
            ours=str(ours_level),
            theirs=str(theirs),
        )
        return
    result.add(
        "minor",
        item_id,
        f"{ours['name']!r} ({class_slug}) required_level {ours_level} is a "
        f"{source} estimate, unverified against the client's own RequiredLevel {theirs}",
        ours=str(ours_level),
        theirs=str(theirs),
    )


def _compare_weapon(ours, sparse, item_row, item_id, class_slug, result) -> None:
    inventory_type = int_column(sparse, "InventoryType")
    item_class_id = int_column(item_row, "ClassID")
    subclass_id = int_column(item_row, "SubclassID")
    if not is_weapon_row(item_class_id, inventory_type):
        return
    literal_damage = row_has_literal_weapon_damage(sparse)
    fields = weapon_fields(sparse, subclass_id, curves=None)
    if ours.get("two_hand") != fields.two_hand:
        result.add(
            "major",
            item_id,
            f"{ours['name']!r} ({class_slug}) two_hand disagrees with the client",
            ours=str(ours.get("two_hand")),
            theirs=str(fields.two_hand),
        )
    if not literal_damage:
        return  # curve-based weapon damage: unverifiable without the fork's curve tables
    for field in ("damage_min", "damage_max", "speed"):
        theirs = getattr(fields, field)
        if ours.get(field) != theirs:
            result.add(
                "major",
                item_id,
                f"{ours['name']!r} ({class_slug}) {field} disagrees with the client",
                ours=str(ours.get(field)),
                theirs=str(theirs),
            )


def _compare_stats(ours, sparse, item_row, item_id, class_slug, result) -> None:
    if not _row_has_literal_amounts(sparse):
        return  # curve-based item: unverifiable without the fork's curve tables
    armor, stats = resolve_item_values(sparse, item_row, curves=None)
    if ours.get("armor") != armor:
        result.add(
            "major",
            item_id,
            f"{ours['name']!r} ({class_slug}) armor disagrees with the client",
            ours=str(ours.get("armor")),
            theirs=str(armor),
        )
    if ours.get("stats") != stats:
        result.add(
            "major",
            item_id,
            f"{ours['name']!r} ({class_slug}) stats disagree with the client",
            ours=str(ours.get("stats")),
            theirs=str(stats),
        )


def check(ctx: AuditContext) -> CategoryResult:
    result = CategoryResult(code=CODE, label=LABEL, primary_source=PRIMARY_SOURCE)
    _icons_and_required_level_zero(ctx, result)
    if not ctx.raw_tables_available:
        result.skipped = (
            "raw/Item.csv and raw/ItemSparse.csv are not committed for this build "
            "(gitignored -- `python -m pipeline fetch` was not run in this read-only audit); "
            "required_level/item_level/quality/two_hand/damage/armor/stats verification "
            "against the client skipped (icon placeholder and required_level==0 checks above "
            "do not need them and still ran)"
        )
        return result
    _compare_against_raw(ctx, result)
    return result
