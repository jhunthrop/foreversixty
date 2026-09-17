"""The 1.60 client's trait tables, read as Classic's three-tab talent trees.

The beta client keeps Forever's talents in the modern trait tables. The legacy
`Talent` table it still ships holds Classic Era's talents, so reading that
would draw the wrong trees; `pipeline/normalize/talent_trees.py` stays for
Classic Era, which has no class trait trees at all (`SkillLineXTraitTree` 404s
there). `has_trait_trees` is the switch between the two.

Every rule here was measured on build 1.60.1.69893 and is asserted rather than
assumed, because the only thing that makes any of them true is that Blizzard's
data says so today:

* a class tree is one `SkillLineXTraitTree` names, and the skill line's display
  name and icon identify the `TalentTab` whose `ClassMask` gives the class;
* a tree's three tabs are its `TraitNodeGroup` rows that are maximal by
  inclusion -- every other group of the tree (a row group, or a cumulative
  rows-0..k points gate) is a proper subset of one of them;
* a tab's left-to-right position is the `PosX` band most of its nodes sit in,
  and that equals the tab's `TalentTab.OrderIndex`;
* positions quantise onto a 4x7 grid in steps of 600, and six coordinates in
  the build are wrong -- three by a stray factor of ten and four by +-10;
* two nodes are stale twins of a live node in the same tab: same spell id, and
  the coordinate that needs the stray digit removed.
"""

from __future__ import annotations

import logging
from collections import Counter
from collections.abc import Iterable
from dataclasses import dataclass, replace

from pipeline.normalize.talents import class_id_from_mask

logger = logging.getLogger(__name__)

#: `TraitCurrency` row for a class's talent points. Its `SourcedMax` is 51,
#: the same budget `web/src/lib/planner/types.ts` calls MAX_POINTS.
CLASS_CURRENCY_ID = 3820
CLASS_POINT_BUDGET = 51

GRID_COLUMNS = 4
GRID_ROWS = 7
#: The distance between two neighbouring cells in `PosX`/`PosY` units.
CELL_STEP = 600
#: `PosX` of column 0, one per tab, left to right. The gaps between the three
#: are not equal (4000 then 4060), so a tab's origin cannot be derived from
#: tab index times a constant.
COLUMN_ORIGINS = (1020, 5020, 9080)
ROW_ORIGIN = 2130
#: How far off a cell's exact coordinate the client is allowed to be. Six rows
#: in build 1.60.1.69893 are off by exactly 10 (e.g. Paladin `PosX` 5030 for
#: column 5020); anything further out is a different kind of wrong.
POSITION_TOLERANCE = 10
#: Points that must already sit in a tab before its rows 1..6 open.
TIER_GATES = (5, 10, 15, 20, 25, 30)

_LAST_COLUMN_POS = COLUMN_ORIGINS[-1] + (GRID_COLUMNS - 1) * CELL_STEP
_LAST_ROW_POS = ROW_ORIGIN + (GRID_ROWS - 1) * CELL_STEP


class TraitDataError(ValueError):
    """The trait tables do not have the shape three-tab talent trees need."""


@dataclass(frozen=True)
class TraitTalent:
    node_id: int
    definition_id: int
    spell_id: int
    max_rank: int
    row: int
    column: int
    prereq_node_id: int | None = None
    prereq_rank: int | None = None


@dataclass(frozen=True)
class TraitTab:
    tab_id: int
    name: str
    position: int
    background: str
    talents: tuple[TraitTalent, ...]


@dataclass(frozen=True)
class TraitClassTree:
    class_id: int
    tree_id: int
    tabs: tuple[TraitTab, ...]


@dataclass(frozen=True)
class TraitRows:
    """The thirteen tables the reader needs, straight off `read_csv`."""

    skill_line: list[dict[str, str]]
    skill_line_x_trait_tree: list[dict[str, str]]
    talent_tab: list[dict[str, str]]
    chr_classes: list[dict[str, str]]
    node: list[dict[str, str]]
    node_entry: list[dict[str, str]]
    node_x_entry: list[dict[str, str]]
    definition: list[dict[str, str]]
    edge: list[dict[str, str]]
    cond: list[dict[str, str]]
    currency: list[dict[str, str]]
    node_group: list[dict[str, str]]
    node_group_x_node: list[dict[str, str]]


def has_trait_trees(skill_line_x_rows: list[dict[str, str]]) -> bool:
    """True when this build's tables name class trait trees.

    Classic Era serves eight of the thirteen trait tables -- their rows all
    belong to non-class trees -- and 404s on `SkillLineXTraitTree`, which the
    fetch writes as an empty table. One non-empty row there is the whole test.
    """
    return bool(skill_line_x_rows)


def _repair(value: int, last: int) -> int:
    """Strip a stray factor of ten from a coordinate.

    Three coordinates in build 1.60.1.69893 carry one: `PosX` 102800 for
    10280, `PosY` 39300 for 3930, `PosY` 21300 for 2130. Dividing while the
    value is past the grid's last coordinate (plus the +-10 the client is
    allowed) removes it and leaves every honest coordinate alone -- `PosY`
    5740 stays 5740 rather than becoming 574.
    """
    while value > last + POSITION_TOLERANCE:
        value //= 10
    return value


def _quantise(value: int, origin: int, count: int, axis: str, raw: int) -> int:
    offset = value - origin
    index = round(offset / CELL_STEP)
    if not 0 <= index < count or abs(offset - index * CELL_STEP) > POSITION_TOLERANCE:
        raise TraitDataError(
            f"{axis} {raw} is not on the grid: it resolves to {value}, which is "
            f"{offset} from origin {origin} and so no multiple of {CELL_STEP} "
            f"inside 0..{count - 1}"
        )
    return index


def grid_cell(pos_x: int, pos_y: int, tab_index: int) -> tuple[int, int]:
    """The (row, column) a node's raw `PosX`/`PosY` name in its tab's grid."""
    column = _quantise(
        _repair(pos_x, _LAST_COLUMN_POS), COLUMN_ORIGINS[tab_index], GRID_COLUMNS, "PosX", pos_x
    )
    row = _quantise(_repair(pos_y, _LAST_ROW_POS), ROW_ORIGIN, GRID_ROWS, "PosY", pos_y)
    return row, column


def _class_of_tree(rows: TraitRows) -> dict[int, int]:
    """Trait tree id -> class id, through the skill line's own talent tab."""
    lines = {int(r["ID"]): r for r in rows.skill_line}
    tabs = {(r["Name_lang"], int(r["SpellIconID"])): r for r in rows.talent_tab}
    out: dict[int, int] = {}
    for link in rows.skill_line_x_trait_tree:
        line = lines.get(int(link["SkillLineID"]))
        if line is None:
            raise TraitDataError(f"skill line {link['SkillLineID']} is not in SkillLine")
        key = (line["DisplayName_lang"], int(line["SpellIconFileID"]))
        tab = tabs.get(key)
        if tab is None:
            raise TraitDataError(f"no TalentTab matches skill line {key}")
        out[int(link["TraitTreeID"])] = class_id_from_mask(int(tab["ID"]), int(tab["ClassMask"]))
    return out


#: The `PosX` halfway between the end of one tab's columns and the start of
#: the next, used only to say which band a node is in.
_BAND_MIDPOINTS = tuple(
    (COLUMN_ORIGINS[i] + (GRID_COLUMNS - 1) * CELL_STEP + COLUMN_ORIGINS[i + 1]) // 2
    for i in range(len(COLUMN_ORIGINS) - 1)
)


def _band(pos_x: int) -> int:
    """Which of the three column bands a `PosX` falls in."""
    repaired = _repair(pos_x, _LAST_COLUMN_POS)
    return sum(repaired > midpoint for midpoint in _BAND_MIDPOINTS)


def _verify_tab_order(tree_id: int, tab_rows: list[dict[str, str]]) -> None:
    """Confirm `TalentTab.OrderIndex` names exactly the positions 0..2, in order.

    `read_trait_trees` pairs `tab_rows` (sorted by `OrderIndex`) with
    `_tab_node_sets`'s result (ordered by `PosX` band) purely by list index --
    nothing in the tables links a `TraitNodeGroup` to a `TalentTab` id
    directly, so `OrderIndex` naming the same left-to-right position as the
    band is the only thing that makes the pairing correct. A tab whose
    `OrderIndex` is not its own rank among the tree's tabs (a gap, a
    duplicate, anything but 0, 1, 2 in order) would zip against the wrong
    band silently; this catches that before the zip happens.
    """
    for index, tab_row in enumerate(tab_rows):
        order_index = int(tab_row["OrderIndex"])
        if order_index != index:
            raise TraitDataError(
                f"tree {tree_id} tab {tab_row['ID']} ({tab_row['Name_lang']}) has "
                f"OrderIndex {order_index}, not {index}; its TalentTab.OrderIndex "
                "and its PosX band position disagree on where it sits"
            )


def _tab_node_sets(
    tree_id: int,
    tree_nodes: set[int],
    rows: TraitRows,
    node_rows: dict[int, dict[str, str]],
) -> list[set[int]]:
    """The tree's three tab node sets, ordered left to right.

    A tab group is a `TraitNodeGroup` of this tree that no other group of the
    tree is a proper superset of. Measured on all nine class trees: exactly
    three per tree, always disjoint, always covering every node, and their
    majority `PosX` bands are always 0, 1 and 2.
    """
    members: dict[int, set[int]] = {}
    group_ids = {int(g["ID"]) for g in rows.node_group if int(g["TraitTreeID"]) == tree_id}
    for link in rows.node_group_x_node:
        group = int(link["TraitNodeGroupID"])
        if group in group_ids:
            members.setdefault(group, set()).add(int(link["TraitNodeID"]))
    maximal = [s for s in members.values() if not any(s < other for other in members.values())]
    covered: set[int] = set()
    for group in maximal:
        covered |= group
    if len(maximal) != len(COLUMN_ORIGINS) or covered != tree_nodes:
        raise TraitDataError(
            f"tree {tree_id} splits into {len(maximal)} tabs covering "
            f"{len(covered)} of its {len(tree_nodes)} nodes; a class tree has "
            f"{len(COLUMN_ORIGINS)} tabs covering all of them"
        )
    if sum(len(group) for group in maximal) != len(covered):
        raise TraitDataError(f"tree {tree_id}'s tabs overlap")
    bands = {}
    for group in maximal:
        band, _ = Counter(_band(int(node_rows[node]["PosX"])) for node in group).most_common(1)[0]
        if band in bands:
            raise TraitDataError(f"tree {tree_id} has two tabs in column band {band}")
        bands[band] = group
    return [bands[index] for index in range(len(COLUMN_ORIGINS))]


def _talents_for_tab(
    node_ids: set[int],
    tab_index: int,
    tab_name: str,
    node_rows: dict[int, dict[str, str]],
    entry_of: dict[int, dict[str, str]],
) -> tuple[TraitTalent, ...]:
    kept = _drop_stale_twins(node_ids, tab_name, node_rows, entry_of)
    talents: list[TraitTalent] = []
    occupied: dict[tuple[int, int], int] = {}
    for node_id in sorted(kept):
        node = node_rows[node_id]
        entry = entry_of[node_id]
        row, column = grid_cell(int(node["PosX"]), int(node["PosY"]), tab_index)
        if (row, column) in occupied:
            raise TraitDataError(
                f"{tab_name} has nodes {occupied[(row, column)]} and {node_id} "
                f"both at row {row} column {column}"
            )
        occupied[(row, column)] = node_id
        talents.append(
            TraitTalent(
                node_id=node_id,
                definition_id=int(entry["TraitDefinitionID"]),
                spell_id=int(entry["SpellID"]),
                max_rank=int(entry["MaxRanks"]),
                row=row,
                column=column,
            )
        )
    return tuple(talents)


def _needs_repair(node: dict[str, str]) -> bool:
    return (
        int(node["PosX"]) > _LAST_COLUMN_POS + POSITION_TOLERANCE
        or int(node["PosY"]) > _LAST_ROW_POS + POSITION_TOLERANCE
    )


def _drop_stale_twins(
    node_ids: set[int],
    tab_name: str,
    node_rows: dict[int, dict[str, str]],
    entry_of: dict[int, dict[str, str]],
) -> set[int]:
    """Keep one node per spell id in a tab.

    Build 1.60.1.69893 leaves two superseded nodes behind -- Hunter 104982 and
    Priest 105865 -- each carrying the same spell as a live node in the same
    tab and each with a coordinate that needs a stray digit removed, which the
    live node never does. Nothing else separates them: both are in the tab's
    group, neither is named by an edge or a condition, and the flags do not
    distinguish them.
    """
    by_spell: dict[int, list[int]] = {}
    for node_id in node_ids:
        by_spell.setdefault(int(entry_of[node_id]["SpellID"]), []).append(node_id)
    kept: set[int] = set()
    for spell_id, ids in by_spell.items():
        if len(ids) == 1:
            kept.add(ids[0])
            continue
        live = [node_id for node_id in ids if not _needs_repair(node_rows[node_id])]
        if len(live) != 1:
            raise TraitDataError(
                f"{tab_name} has {len(ids)} nodes for spell {spell_id} "
                f"({sorted(ids)}) and {len(live)} of them are on the grid; "
                "which one the client draws cannot be decided from the tables"
            )
        logger.info(
            "%s: dropping stale node %s, superseded by %s (spell %s)",
            tab_name,
            sorted(set(ids) - set(live)),
            live[0],
            spell_id,
        )
        kept.add(live[0])
    return kept


def check_tier_gates(rows: TraitRows, tree_ids: Iterable[int]) -> None:
    """Assert the client still gates a tab's rows at 5 points per tier.

    Every `TraitCond` row in build 1.60.1.69893 is `CondType` 0 and says only
    "this group needs N points spent in this tab's currency"; none of them
    carries a prerequisite rank. Filtering to the class currency, no pinned
    node and a non-zero amount leaves exactly eighteen rows per class tree --
    three tabs times 5/10/15/20/25/30 -- for all nine classes. The planner and
    the API both hardcode that ladder (`POINTS_PER_TIER`), so a build that
    changed it has to stop the pipeline rather than quietly disagree.
    """
    budgets = {int(c["ID"]): int(c["SourcedMax"]) for c in rows.currency}
    budget = budgets.get(CLASS_CURRENCY_ID)
    if budget != CLASS_POINT_BUDGET:
        raise TraitDataError(
            f"TraitCurrency {CLASS_CURRENCY_ID} allows {budget} points, "
            f"not the {CLASS_POINT_BUDGET} the planner spends"
        )
    want = sorted(TIER_GATES * len(COLUMN_ORIGINS))
    for tree_id in sorted(tree_ids):
        got = sorted(
            int(c["SpentAmountRequired"])
            for c in rows.cond
            if int(c["TraitTreeID"]) == tree_id
            and int(c["TraitCurrencyID"]) == CLASS_CURRENCY_ID
            and int(c["TraitNodeID"]) == 0
            and int(c["SpentAmountRequired"]) > 0
        )
        if got != want:
            raise TraitDataError(
                f"tree {tree_id}'s points gates are {got}, not {want}: the "
                "planner's five-points-per-tier rule would not match the client"
            )


def _with_prerequisites(tree: TraitClassTree, edge_rows: list[dict[str, str]]) -> TraitClassTree:
    """Fill in each talent's one prerequisite from `TraitEdge`.

    Left is the prerequisite and right the dependent, checked against the
    Wowhead snapshot's own `requires` arrays: 69 of the tree edges match it
    exactly. Two edges are the reverse leg of a two-way pair (Druid
    Nature's Splendor/Nature's Majesty, Hunter Bestial Wrath/Intimidation);
    dropping the leg whose prerequisite sits further down the tab leaves no
    cycles, at most one prerequisite per talent, and the direction the
    snapshot records. The required rank is not in the tables at all -- it is
    the prerequisite's own rank cap in all 69 cases.
    """
    by_node = {t.node_id: t for tab in tree.tabs for t in tab.talents}
    tab_of = {t.node_id: tab.tab_id for tab in tree.tabs for t in tab.talents}
    links: dict[int, tuple[int, int]] = {}
    for edge in edge_rows:
        left, right = int(edge["LeftTraitNodeID"]), int(edge["RightTraitNodeID"])
        # Edges of other trees, and the two edges into dropped stale nodes.
        if left not in by_node or right not in by_node:
            continue
        if tab_of[left] != tab_of[right]:
            raise TraitDataError(
                f"edge {edge['ID']} joins nodes {left} and {right} in different tabs"
            )
        prerequisite, dependent = by_node[left], by_node[right]
        if prerequisite.row > dependent.row:
            logger.info(
                "dropping edge %s: node %s is below its dependent %s",
                edge["ID"],
                left,
                right,
            )
            continue
        if right in links:
            raise TraitDataError(
                f"node {right} has two prerequisites, {links[right][0]} and {left}; "
                "the planner's rules carry one"
            )
        links[right] = (left, prerequisite.max_rank)
    return TraitClassTree(
        class_id=tree.class_id,
        tree_id=tree.tree_id,
        tabs=tuple(
            TraitTab(
                tab_id=tab.tab_id,
                name=tab.name,
                position=tab.position,
                background=tab.background,
                talents=tuple(
                    replace(
                        talent,
                        prereq_node_id=links[talent.node_id][0],
                        prereq_rank=links[talent.node_id][1],
                    )
                    if talent.node_id in links
                    else talent
                    for talent in tab.talents
                ),
            )
            for tab in tree.tabs
        ),
    )


def read_trait_trees(rows: TraitRows) -> list[TraitClassTree]:
    """One `TraitClassTree` per class, tabs left to right, sorted by class id."""
    class_of_tree = _class_of_tree(rows)
    node_rows = {int(n["ID"]): n for n in rows.node}
    entries = {int(e["ID"]): e for e in rows.node_entry}
    definitions = {int(d["ID"]): d for d in rows.definition}
    # The full build carries non-class trees too (covenant/soulbind trees and
    # the like), and some of their nodes are choice nodes with more than one
    # entry -- fine for them, meaningless for us. Scoping to nodes that
    # actually belong to a class tree before the one-entry-per-node check
    # keeps this reader from tripping over data it never reads.
    class_tree_node_ids = {
        node_id for node_id, node in node_rows.items() if int(node["TraitTreeID"]) in class_of_tree
    }
    entry_of: dict[int, dict[str, str]] = {}
    for link in rows.node_x_entry:
        node_id = int(link["TraitNodeID"])
        if node_id not in class_tree_node_ids:
            continue
        if node_id in entry_of:
            raise TraitDataError(f"node {node_id} has more than one entry")
        entry = entries[int(link["TraitNodeEntryID"])]
        definition = definitions[int(entry["TraitDefinitionID"])]
        entry_of[node_id] = {
            "TraitDefinitionID": entry["TraitDefinitionID"],
            "MaxRanks": entry["MaxRanks"],
            "SpellID": definition["SpellID"],
        }

    tabs_by_class: dict[int, list[dict[str, str]]] = {}
    for tab in rows.talent_tab:
        class_id = class_id_from_mask(int(tab["ID"]), int(tab["ClassMask"]))
        tabs_by_class.setdefault(class_id, []).append(tab)

    trees: list[TraitClassTree] = []
    for tree_id, class_id in sorted(class_of_tree.items()):
        tree_nodes = {n for n, node in node_rows.items() if int(node["TraitTreeID"]) == tree_id}
        missing = sorted(tree_nodes - set(entry_of))
        if missing:
            raise TraitDataError(f"tree {tree_id} nodes {missing} have no entry")
        tab_rows = sorted(tabs_by_class.get(class_id, []), key=lambda t: int(t["OrderIndex"]))
        if len(tab_rows) != len(COLUMN_ORIGINS):
            raise TraitDataError(
                f"class {class_id} has {len(tab_rows)} TalentTab rows, "
                f"not {len(COLUMN_ORIGINS)}"
            )
        _verify_tab_order(tree_id, tab_rows)
        node_sets = _tab_node_sets(tree_id, tree_nodes, rows, node_rows)
        tabs = tuple(
            TraitTab(
                tab_id=int(tab_row["ID"]),
                name=tab_row["Name_lang"],
                position=index,
                background=tab_row["BackgroundFile"],
                talents=_talents_for_tab(
                    node_sets[index], index, tab_row["Name_lang"], node_rows, entry_of
                ),
            )
            for index, tab_row in enumerate(tab_rows)
        )
        trees.append(TraitClassTree(class_id=class_id, tree_id=tree_id, tabs=tabs))
    check_tier_gates(rows, class_of_tree)
    linked = [_with_prerequisites(tree, rows.edge) for tree in trees]
    return sorted(linked, key=lambda t: t.class_id)
