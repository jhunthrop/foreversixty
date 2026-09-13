from pathlib import Path

import pytest

from pipeline.csvio import read_csv
from pipeline.icons import icon_names
from pipeline.normalize import write_json, write_model
from pipeline.normalize.gear import ItemDataError, build_class_items, build_item_sets
from pipeline.spelltext import load_spell_text

HERE = Path(__file__).parent


def fixture_icons():
    return icon_names(read_csv(HERE / "fixtures/ManifestInterfaceData.csv"))


def fixture_spell_text():
    return load_spell_text(
        read_csv(HERE / "fixtures/Spell.csv"),
        read_csv(HERE / "fixtures/SpellMisc.csv"),
        read_csv(HERE / "fixtures/SpellEffect.csv"),
        read_csv(HERE / "fixtures/SpellDuration.csv"),
    )


def build_all():
    return build_class_items(
        read_csv(HERE / "fixtures/ItemSparse.csv"),
        read_csv(HERE / "fixtures/Item.csv"),
        read_csv(HERE / "fixtures/ChrClasses.csv"),
        fixture_icons(),
        "1.0.0.1",
    )


def by_slug():
    return {record.class_slug: record for record in build_all()}


def test_warrior_items_match_golden(tmp_path: Path):
    out = tmp_path / "warrior.json"
    write_model(by_slug()["warrior"], out)
    assert out.read_text() == (HERE / "golden/items_warrior.json").read_text()


def test_slot_comes_from_inventory_type():
    items = {i.id: i for i in by_slug()["warrior"].items}
    assert items[16866].slot == "head"
    assert items[19325].slot == "finger"
    assert items[19019].slot == "main_hand"


def test_armour_and_resistances_come_from_the_resistance_columns():
    helm = {i.id: i for i in by_slug()["warrior"].items}[16866]
    assert helm.armor == 608
    assert helm.stats == {"stamina": 35, "strength": 15, "fire_res": 10}


def test_unique_is_max_count_one():
    items = {i.id: i for i in by_slug()["warrior"].items}
    assert items[19325].unique is True
    assert items[16866].unique is False


def test_set_id_is_none_when_the_item_is_not_in_a_set():
    items = {i.id: i for i in by_slug()["warrior"].items}
    assert items[16866].set_id == 209
    assert items[19325].set_id is None


def test_class_restriction_and_proficiency_are_both_applied():
    warrior = {i.id for i in by_slug()["warrior"].items}
    mage = {i.id for i in by_slug()["mage"].items}
    assert 16866 in warrior and 16866 not in mage  # plate: proficiency excludes the mage
    assert 14152 in mage and 14152 not in warrior  # AllowableClass mask excludes the warrior
    assert 2825 in warrior and 2825 not in mage  # bow: proficiency excludes the mage
    assert 19019 in warrior and 19019 in mage  # one-hand sword: both classes can use it


def test_every_class_gets_a_record_even_when_the_list_is_short():
    assert sorted(record.class_slug for record in build_all()) == ["mage", "paladin", "warrior"]


def test_non_equipment_and_overlevelled_items_are_dropped():
    for record in build_all():
        ids = {i.id for i in record.items}
        assert 2589 not in ids  # InventoryType 0
        assert 12345 not in ids  # RequiredLevel 70


def test_items_are_sorted_by_required_level_then_name():
    items = by_slug()["warrior"].items
    assert [(i.required_level, i.name) for i in items] == sorted(
        (i.required_level, i.name) for i in items
    )


def test_an_unmapped_stat_id_is_an_error_not_a_guess():
    rows = read_csv(HERE / "fixtures/ItemSparse.csv")
    rows[0]["StatModifier_bonusStat_0"] = "99"
    with pytest.raises(ItemDataError, match="99"):
        build_class_items(
            rows,
            read_csv(HERE / "fixtures/Item.csv"),
            read_csv(HERE / "fixtures/ChrClasses.csv"),
            fixture_icons(),
            "1.0.0.1",
        )


def test_sets_match_golden(tmp_path: Path):
    records = build_item_sets(
        read_csv(HERE / "fixtures/ItemSet.csv"),
        read_csv(HERE / "fixtures/ItemSetSpell.csv"),
        fixture_spell_text(),
    )
    out = tmp_path / "sets.json"
    write_json(records, out)
    assert out.read_text() == (HERE / "golden/sets.json").read_text()


def test_set_bonuses_are_sorted_by_piece_count_with_resolved_text():
    records = build_item_sets(
        read_csv(HERE / "fixtures/ItemSet.csv"),
        read_csv(HERE / "fixtures/ItemSetSpell.csv"),
        fixture_spell_text(),
    )
    assert records[0].item_ids == [16866, 16867]
    assert [(b.pieces, b.description) for b in records[0].bonuses] == [
        (3, "Increases the block value of your shield by 31."),
        (5, "Gives you a $h% chance to generate an additional Rage point."),
    ]
