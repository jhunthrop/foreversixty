import json
import shutil
from pathlib import Path

from pipeline.forkdb import load_fork_database
from pipeline.loot.gear import (
    apply_fork_columns,
    build_enchants,
    build_suffixes,
    faction_restrictions,
    suffix_options,
)

HERE = Path(__file__).parent
ENGINE = HERE / "fixtures/loot"


def fork():
    return load_fork_database(ENGINE)


def test_every_suffix_is_emitted_with_its_stats_named():
    records = build_suffixes(fork())
    assert [(r.id, r.name) for r in records] == [(5, "of Intellect"), (6, "of Strength")]
    assert records[0].stats == {"intellect": 4.0}
    assert records[1].stats == {"strength": 7.0}


def test_suffix_options_index_only_the_items_that_roll_one():
    assert suffix_options(fork()) == {110: [5, 6]}


def test_faction_restrictions_index_only_the_restricted_items():
    assert faction_restrictions(fork()) == {100: "alliance_only"}


def prepared(tmp_path):
    build_dir = tmp_path / "1.60.1.69893"
    build_dir.mkdir()
    shutil.copy(ENGINE / "items.json", build_dir / "items.json")
    return build_dir


def test_apply_fork_columns_fills_both_and_leaves_the_rest_empty(tmp_path):
    build_dir = prepared(tmp_path)
    with_suffixes, restricted = apply_fork_columns(
        build_dir, suffix_options(fork()), faction_restrictions(fork())
    )
    assert (with_suffixes, restricted) == (1, 1)
    rows = json.loads((build_dir / "items.json").read_text(encoding="utf-8"))
    by_id = {row["id"]: row for row in rows}
    assert by_id[110]["suffixes"] == [5, 6]
    assert by_id[100]["suffixes"] == []
    assert by_id[100]["faction_restriction"] == "alliance_only"
    assert by_id[110]["faction_restriction"] == ""
    assert len(rows) == 11


def test_apply_fork_columns_keeps_the_key_order_and_appends_the_new_ones(tmp_path):
    build_dir = prepared(tmp_path)
    apply_fork_columns(build_dir, suffix_options(fork()), faction_restrictions(fork()))
    rows = json.loads((build_dir / "items.json").read_text(encoding="utf-8"))
    assert list(rows[0]) == [
        "id",
        "name",
        "quality",
        "item_level",
        "required_level",
        "class_id",
        "subclass_id",
        "inventory_type",
        "suffixes",
        "faction_restriction",
    ]


def test_apply_fork_columns_is_idempotent(tmp_path):
    build_dir = prepared(tmp_path)
    options, restrictions = suffix_options(fork()), faction_restrictions(fork())
    apply_fork_columns(build_dir, options, restrictions)
    once = (build_dir / "items.json").read_bytes()
    apply_fork_columns(build_dir, options, restrictions)
    assert (build_dir / "items.json").read_bytes() == once


def enchant(effect_id: int, spell_id: int):
    for record in build_enchants(fork()):
        if record.id == effect_id and record.spell_id == spell_id:
            return record
    raise AssertionError(f"no enchant {effect_id}/{spell_id}")


def test_every_fork_enchant_row_is_emitted_even_when_it_shares_an_effect_id():
    records = build_enchants(fork())
    assert len(records) == 4
    assert [r.id for r in records] == [15, 41, 41, 241]
    assert [r.spell_id for r in records if r.id == 41] == [7418, 7420]


def test_an_enchants_slots_come_from_its_item_type_and_extra_types():
    assert enchant(41, 7420).slots == ["chest"]
    assert enchant(41, 7418).slots == ["wrist"]
    assert enchant(15, 2831).slots == ["chest", "feet", "hands", "legs"]
    assert enchant(241, 7745).slots == ["main_hand", "off_hand"]


def test_item_types_is_the_shape_restriction_not_the_slot():
    assert enchant(41, 7420).item_types == ["normal"]
    assert enchant(15, 2831).item_types == ["kit"]
    assert enchant(241, 7745).item_types == ["two_hand"]


def test_classes_are_slugs_and_sorted_and_empty_means_anyone():
    assert enchant(15, 2831).classes == ["paladin", "warrior"]
    assert enchant(41, 7420).classes == []


def test_icons_come_from_the_item_when_there_is_one_and_the_spell_otherwise():
    assert enchant(15, 2831).icon == "inv_misc_armorkit_17"
    assert enchant(41, 7420).icon == "spell_holy_chest"


def test_stats_and_phase_are_read_off_the_row():
    # The fork encodes this kit's bonus at the engine's StatBonusArmor
    # index, distinct from StatArmor -- see the note on `PROTO_STAT_ALIASES`
    # in pipeline/simdb/statmap.py. The brief's draft of this test read
    # "armor" here; that was measured against the wrong stat index.
    assert enchant(15, 2831).stats == {"bonus_armor": 8.0}
    assert enchant(241, 7745).stats == {"attack_power": 12.0}
    assert enchant(241, 7745).phase == 2
    assert enchant(41, 7418).phase == 0
