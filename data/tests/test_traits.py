import pytest
from traits_fixture import TRAITS, trait_rows

from pipeline.csvio import read_csv
from pipeline.normalize.traits import (
    COLUMN_ORIGINS,
    TraitDataError,
    grid_cell,
    has_trait_trees,
    read_trait_trees,
)


def test_a_build_without_the_class_link_table_has_no_trait_trees():
    assert has_trait_trees(read_csv(TRAITS / "SkillLineXTraitTree.csv")) is True
    assert has_trait_trees([]) is False


@pytest.mark.parametrize(
    ("pos_x", "pos_y", "tab", "cell"),
    [
        (1020, 2130, 0, (0, 0)),
        (2820, 5730, 0, (6, 3)),
        (6820, 3330, 1, (2, 3)),
        (10880, 5730, 2, (6, 3)),
        # The client's own +-10 wobble: Paladin PosX 5030, Warlock PosY 2120.
        (5030, 2730, 1, (1, 0)),
        (1020, 2120, 0, (0, 0)),
        # A stray factor of ten: Hunter PosX 102800, Priest PosY 21300.
        (102800, 5730, 2, (6, 2)),
        (9080, 21300, 2, (0, 0)),
    ],
)
def test_grid_cell_quantises_and_repairs_the_clients_typos(pos_x, pos_y, tab, cell):
    assert grid_cell(pos_x, pos_y, tab) == cell


def test_grid_cell_refuses_a_position_that_is_not_on_its_tabs_grid():
    # Priest node 105865 sits at PosX 9280 while its tab's columns are 5020..6820.
    with pytest.raises(TraitDataError, match="9280"):
        grid_cell(9280, 2130, 1)


def test_the_tree_is_split_into_the_three_tabs_in_the_games_order():
    (warrior,) = read_trait_trees(trait_rows())
    assert warrior.class_id == 1
    assert warrior.tree_id == 1117
    assert [(t.tab_id, t.name, t.position, t.background) for t in warrior.tabs] == [
        (161, "Arms", 0, "WarriorArms"),
        (164, "Fury", 1, "WarriorFury"),
        (163, "Protection", 2, "WarriorProtection"),
    ]


def test_every_talent_lands_on_its_own_cell_with_its_spell_and_rank_cap():
    (warrior,) = read_trait_trees(trait_rows())
    arms, fury, protection = warrior.tabs
    assert [(t.node_id, t.spell_id, t.max_rank, t.row, t.column) for t in arms.talents] == [
        (900001, 12282, 3, 0, 0),
        (900002, 16462, 5, 0, 1),
        (900003, 12295, 5, 1, 1),
        (900004, 12296, 3, 1, 2),
        (900005, 12163, 2, 2, 1),
    ]
    assert [(t.node_id, t.row, t.column) for t in fury.talents] == [
        (900011, 0, 1),
        (900012, 0, 2),
        (900013, 1, 0),
    ]
    assert [(t.node_id, t.row, t.column) for t in protection.talents] == [
        (900021, 0, 0),
        (900022, 2, 0),
        (900023, 1, 1),
    ]


def test_the_stale_twin_of_a_talent_is_dropped_and_the_live_one_kept():
    (warrior,) = read_trait_trees(trait_rows())
    arms = warrior.tabs[0]
    deflection = [t for t in arms.talents if t.spell_id == 16462]
    assert [t.node_id for t in deflection] == [900002]


def test_two_twins_that_both_need_repair_are_refused_rather_than_guessed():
    nodes = read_csv(TRAITS / "TraitNode.csv")
    for node in nodes:
        if node["ID"] == "900002":
            node["PosX"] = "16200"
    with pytest.raises(TraitDataError, match="16462"):
        read_trait_trees(trait_rows(node=nodes))


def test_two_talents_in_one_cell_are_refused():
    nodes = read_csv(TRAITS / "TraitNode.csv")
    for node in nodes:
        if node["ID"] == "900004":
            node["PosX"] = str(COLUMN_ORIGINS[0] + 600)  # onto 900003's cell
    with pytest.raises(TraitDataError, match="row 1 column 1"):
        read_trait_trees(trait_rows(node=nodes))


def test_a_tree_whose_groups_do_not_make_three_tabs_is_refused():
    groups = [g for g in read_csv(TRAITS / "TraitNodeGroup.csv") if g["ID"] != "600002"]
    members = [
        m
        for m in read_csv(TRAITS / "TraitNodeGroupXTraitNode.csv")
        if m["TraitNodeGroupID"] != "600002"
    ]
    with pytest.raises(TraitDataError, match="2 tabs"):
        read_trait_trees(trait_rows(node_group=groups, node_group_x_node=members))


def test_a_tabs_order_index_disagreeing_with_its_column_band_is_refused():
    # Protection's real OrderIndex is 2; break the 0..2 contiguity so its
    # position in the OrderIndex-sorted list can no longer be trusted to line
    # up with the PosX band _tab_node_sets computed for it.
    tabs = read_csv(TRAITS / "TalentTab.csv")
    for tab in tabs:
        if tab["ID"] == "163":  # Protection
            tab["OrderIndex"] = "5"
    with pytest.raises(TraitDataError, match="163"):
        read_trait_trees(trait_rows(talent_tab=tabs))


def test_a_non_class_trees_choice_node_with_two_entries_is_ignored():
    # A choice node with more than one TraitNodeEntry is ordinary in a
    # non-class tree (covenant/soulbind-style trees in the full build); tree
    # 9999 is never named by SkillLineXTraitTree.csv, so it is not one of the
    # class trees this reader resolves and its node should simply be skipped.
    nodes = [
        *read_csv(TRAITS / "TraitNode.csv"),
        {
            "ID": "999001",
            "TraitTreeID": "9999",
            "PosX": "0",
            "PosY": "0",
            "Type": "0",
            "Flags": "0",
            "TraitSubTreeID": "0",
        },
    ]
    choice_entry = {
        "TraitDefinitionID": "700001",
        "MaxRanks": "1",
        "NodeEntryType": "0",
        "TraitSubTreeID": "0",
    }
    entries = [
        *read_csv(TRAITS / "TraitNodeEntry.csv"),
        {"ID": "999101", **choice_entry},
        {"ID": "999102", **choice_entry},
    ]
    links = [
        *read_csv(TRAITS / "TraitNodeXTraitNodeEntry.csv"),
        {"ID": "999201", "TraitNodeID": "999001", "TraitNodeEntryID": "999101", "_Index": "0"},
        {"ID": "999202", "TraitNodeID": "999001", "TraitNodeEntryID": "999102", "_Index": "1"},
    ]
    (warrior,) = read_trait_trees(trait_rows(node=nodes, node_entry=entries, node_x_entry=links))
    assert warrior.class_id == 1
