from pathlib import Path

import pytest

from pipeline.csvio import read_csv
from pipeline.icons import PLACEHOLDER_ICON, icon_names
from pipeline.models import ItemSetBonus
from pipeline.normalize import write_json, write_model
from pipeline.normalize.gear import (
    ItemDataError,
    build_class_items,
    build_item_sets,
    is_junk_name,
)
from pipeline.normalize.item_curves import load_item_curves
from pipeline.proficiency import WEAPON
from pipeline.spelltext import SpellTextError, load_spell_text

HERE = Path(__file__).parent


def fixture_icons():
    return icon_names(read_csv(HERE / "fixtures/ManifestInterfaceData.csv"))


def fixture_curves():
    return load_item_curves(
        read_csv(HERE / "fixtures/ItemArmorTotal.csv"),
        read_csv(HERE / "fixtures/ItemArmorQuality.csv"),
        read_csv(HERE / "fixtures/ItemArmorShield.csv"),
        read_csv(HERE / "fixtures/ArmorLocation.csv"),
        read_csv(HERE / "fixtures/RandPropPoints.csv"),
    )


def build_all_1_60(curves=None):
    return build_class_items(
        read_csv(HERE / "fixtures/ItemSparse_1_60.csv"),
        read_csv(HERE / "fixtures/Item.csv"),
        read_csv(HERE / "fixtures/ChrClasses.csv"),
        fixture_icons(),
        "1.60.1.69893",
        curves,
    )


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
    assert 16866 in warrior and 16866 not in mage  # plate: both filters exclude the mage
    assert 14152 in mage and 14152 not in warrior  # AllowableClass mask excludes the warrior
    assert 17066 in warrior and 17066 not in mage  # shield: proficiency excludes the mage
    assert 19019 in warrior and 19019 in mage  # one-hand sword: both classes can use it


def test_every_class_gets_a_record_even_when_the_list_is_short():
    assert sorted(record.class_slug for record in build_all()) == ["mage", "paladin", "warrior"]


def test_non_equipment_and_overlevelled_items_are_dropped():
    for record in build_all():
        ids = {i.id for i in record.items}
        assert 2589 not in ids  # InventoryType 0
        assert 12345 not in ids  # RequiredLevel 70


@pytest.mark.parametrize(
    "name",
    [
        "Gamemaster Hood",
        "GM Robe",
        "AHNQIRAJ TEST ITEM",
        "Cloaked Hood TEST",
        "Deprecated Old Belt",
        "Monster - Axe, 2H Arcanite Reaper",
        "(OLD)Heavy Throwing Axe",
        # Found leaking into items/ once the curve resolver gave them real stat
        # values (build 1.60.1.69893) -- "test"/"gm"/etc alone did not catch them.
        "Ring of Critical Testing",  # id 18968: "testing" is not a \btest\b match
        "Ring of Critical Testing 2",  # id 18970
        "Ring of Critical Testing 4",  # id 18982
        "QATest +1000 Spell Dmg Ring",  # id 24358: "qatest" has no space or "test\b"
        "QATest Darkmoon Faire Tickets",
        "UNUSED Electrified Mithril Gauntlets",  # id 213391
        "UNUSED - Cloak of Arcane Insulation",  # id 215112
        "UNUSED - Razor-Lined Shoulderpads",  # id 215113
        "Blessed Qiraji Naturalist Staff UNUSED",  # id 21276
        "Ahn'Qiraj Mace [PH]",  # id 21127; a whole unshipped tier set ("[PH] ...
        "Naxxramas Polearm [PH]",  # id 22817   Brilliant/Rising/Shining Dawn")
        "[PH] Brilliant Dawn Gauntlets",  # is also marked this way in raw ItemSparse
        "Placeholder",  # id 274027
        "[DNT] Crafted Tier Piece Placeholder",  # id 273878
        "Magic Knucklebone (DND)",
        "zzOLDCodex of Prayer of Fortitude",
    ],
)
def test_the_junk_name_matcher_catches_every_pattern_class(name: str):
    assert is_junk_name(name) is True


@pytest.mark.parametrize(
    "name",
    [
        "Testament of Hope",  # real uncommon off-hand; contains "Test"
        "Magma Forged Band",  # real rare ring; contains "gm"
        "Old Blunderbuss",  # real gun; "old" without the client's "(OLD)" marker
        "Grasp of the Old God",
        "Buru's Skull Fragment",
        "Contest Winner's Tabard",  # contains "test" mid-word, no boundary
        "The Greatest Race of Hunters",  # contains "test" mid-word ("Greatest")
        "Rexxar's Testament",  # "Testament" is not a \btest\b or \btesting\b match
        "Corrupt Tested Sample",  # "Tested" is not "testing"; no pattern matches it
        "Un'Goro Tested Sample",
    ],
)
def test_the_junk_name_matcher_keeps_real_items_that_merely_contain_the_letters(name: str):
    assert is_junk_name(name) is False


@pytest.mark.parametrize(
    "name",
    [
        "Field Testing Kit",
        "Sealed Field Testing Kit",
        "Potion of Tradeskill Testing - Level 60",
    ],
)
def test_the_testing_pattern_has_a_known_false_positive_on_real_consumables(name: str):
    """These are real quest items, not QA junk -- but the "testing"/"qatest"
    alternatives added to catch "Ring of Critical Testing" and "QATest ..."
    (leaked into items/ once curve support gave build 1.60.1.69893's items real
    values; see JUNK_NAME_PATTERN) cannot tell a QA marker's "Testing" from
    ordinary English. Accepted rather than narrowed further: every name here is
    InventoryType 0 in the raw client data (not equippable gear), so
    build_class_items's own slot filter drops each one before is_junk_name is
    ever called on it -- the false positive here has no effect on emitted
    items/."""
    assert is_junk_name(name) is True


def test_only_uncommon_through_legendary_items_are_emitted():
    """Linen Belt is quality 1 and real armour, so only the quality clause drops it."""
    ids = {i.id for record in build_all() for i in record.items}
    assert 7026 not in ids
    assert {i.quality for record in build_all() for i in record.items} <= {2, 3, 4, 5}


def test_junk_named_rows_are_dropped_but_a_near_miss_name_is_kept():
    """Cloaked Hood TEST is uncommon plate-free armour with stats: only the name drops it."""
    ids = {i.id for record in build_all() for i in record.items}
    assert 19743 not in ids
    assert 13315 in ids  # Testament of Hope survives the "test" pattern


def test_a_damage_only_weapon_survives_the_no_armour_no_stats_clause():
    """Annihilator is a real one-hand axe carrying no armour and no stat: its whole
    value is its damage, which this pipeline does not emit yet. `Item.ClassID` 2
    exempts it from the clause, so it survives on quality and name alone. Bow of
    Searing Arrows is the same case in a second slot."""
    warrior = {i.id: i for i in by_slug()["warrior"].items}
    annihilator = warrior[12798]
    assert annihilator.armor == 0
    assert annihilator.stats == {}
    assert annihilator.slot == "main_hand"
    assert warrior[2825].slot == "ranged"


def test_a_stat_less_armour_piece_is_still_dropped():
    """The exemption is weapons only. Featureless Cloth Vest is a quality-2 cloth
    chest with no armour and no stat, so the planner has nothing to compare it on and
    the clause still drops it."""
    ids = {i.id for record in build_all() for i in record.items}
    assert 99002 not in ids


def test_only_weapons_may_be_emitted_with_neither_armour_nor_stats():
    weapon_ids = {
        int(row["ID"])
        for row in read_csv(HERE / "fixtures/Item.csv")
        if int(row["ClassID"]) == WEAPON
    }
    for record in build_all():
        for item in record.items:
            if item.id in weapon_ids:
                continue
            assert item.armor != 0 or any(item.stats.values()), item.id


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


def test_a_row_truncated_inside_the_stat_block_is_an_error_not_partial_stats():
    """csv.DictReader pads a short row with None, so the key is there but the value is not."""
    rows = read_csv(HERE / "fixtures/ItemSparse.csv")
    rows[0]["StatModifier_bonusStat_1"] = None  # type: ignore[assignment]
    with pytest.raises(ItemDataError, match="StatModifier_bonusStat_1"):
        build_class_items(
            rows,
            read_csv(HERE / "fixtures/Item.csv"),
            read_csv(HERE / "fixtures/ChrClasses.csv"),
            fixture_icons(),
            "1.0.0.1",
        )


def test_a_missing_bonus_amount_column_reads_as_no_data_for_that_stat():
    """The 1.60 client (Forever beta) has no StatModifier_bonusAmount_* columns at all:
    a stat type with no paired amount column is a build whose schema states nothing
    about the amount, not a malformed row, so it must not raise. See
    test_a_1_60_shaped_armour_item_has_no_armour_or_stats for the whole-build shape;
    this covers just the one column going missing."""
    rows = read_csv(HERE / "fixtures/ItemSparse.csv")
    del rows[0]["StatModifier_bonusAmount_0"]  # the pair for StatModifier_bonusStat_0
    records = build_class_items(
        rows,
        read_csv(HERE / "fixtures/Item.csv"),
        read_csv(HERE / "fixtures/ChrClasses.csv"),
        fixture_icons(),
        "1.0.0.1",
    )
    helm = {i.id: i for i in {r.class_slug: r for r in records}["warrior"].items}[16866]
    # StatModifier_bonusStat_0 was stamina (see the golden fixture); with no amount
    # column for it, stamina is silently absent rather than raising, while the other
    # stat pair and the resistance columns -- untouched -- still come through.
    assert helm.stats == {"strength": 15, "fire_res": 10}
    assert helm.armor == 608


def test_a_1_60_shaped_armour_item_has_no_armour_or_stats_when_no_curves_are_given():
    """The 1.60 client's ItemSparse carries no Resistances_* or
    StatModifier_bonusAmount_* columns at all -- armour and stat amounts are computed
    from curve tables (RandPropPoints, ItemArmorTotal, ItemArmorQuality, ItemArmorShield,
    ArmorLocation) instead. Without those tables (curves=None, the default -- a future
    product whose fetch could not supply them, see pipeline.wago.OPTIONAL_TABLES), the
    client states nothing usable in a column for them: an armour piece has no armour and
    no stats to report and is dropped, exactly like a stat-less armour piece on the old
    schema (test_a_stat_less_armour_piece_is_still_dropped); a weapon is exempt from
    that clause and still survives on quality and name alone. See
    test_curve_tables_resolve_armour_and_stats_for_every_armour_type below for the same
    build with curves supplied."""
    items = {i.id: i for r in build_all_1_60() for i in r.items}
    assert 16866 not in items  # armour with nothing to compare it on: dropped
    annihilator = items[12798]
    assert annihilator.armor == 0
    assert annihilator.stats == {}


def test_curve_tables_resolve_armour_for_every_armour_type_and_a_shield():
    """armour = ItemArmorTotal[ilvl][type] * ItemArmorQuality[ilvl][quality] *
    ArmorLocation[slot][type] (ItemArmorShield[ilvl][quality] for a shield, no
    location term). Every number here is the real 1.60.1.69893 curve tables' own
    value at item level 40/66 -- see data/README.md for how they were pulled and
    cross-checked against 2,830 items unchanged from builds/1.15.9.69722. The plate
    helm (16866) is the same item id: 608 armour is the *literal* Resistances_0
    Era emits for it too, confirming the curve formula reproduces it exactly."""
    items = {i.id: i for r in build_all_1_60(fixture_curves()) for i in r.items}
    assert items[16866].armor == 608  # plate, ilvl66, epic, head
    assert items[30001].armor == 58  # cloth, ilvl40, rare, robe (scored as chest)
    assert items[30002].armor == 121  # leather, ilvl40, rare, chest
    assert items[30003].armor == 254  # mail, ilvl40, rare, chest
    assert items[30005].armor == 1078  # shield, ilvl40, rare -- ItemArmorShield only


def test_curve_tables_resolve_a_primary_and_a_secondary_stat():
    """amount = RandPropPoints[ilvl][quality][slot group] * StatPercentEditor / 10000.
    The helm's stats also come from the curve now: budget 47 (Epic, ilvl66, head's
    group 0) at editor 7440/3200 rounds to 35 stamina and 15 strength -- matching
    Era's own literal 35/15 for the same item id exactly. The ring carries one of
    each kind of stat the planner tracks: a primary (strength) and what the game
    calls a secondary (crit)."""
    items = {i.id: i for r in build_all_1_60(fixture_curves()) for i in r.items}
    assert items[16866].stats == {"stamina": 35, "strength": 15}
    assert items[30006].stats == {"strength": 8, "crit": 5}
    assert items[30006].armor == 0  # misc subclass (rings): no armour curve applies


def test_an_era_shaped_row_still_uses_the_literal_path_even_when_curves_are_given():
    """Curves are only a fallback for a row with no literal columns: passing them
    alongside an Era-shaped fixture must not change Era's own golden output."""
    records = build_class_items(
        read_csv(HERE / "fixtures/ItemSparse.csv"),
        read_csv(HERE / "fixtures/Item.csv"),
        read_csv(HERE / "fixtures/ChrClasses.csv"),
        fixture_icons(),
        "1.0.0.1",
        fixture_curves(),
    )
    helm = {i.id: i for i in {r.class_slug: r for r in records}["warrior"].items}[16866]
    assert helm.armor == 608
    assert helm.stats == {"stamina": 35, "strength": 15, "fire_res": 10}


def test_curves_whose_own_tables_are_incomplete_resolve_nothing():
    """An ItemCurves loaded from a build that 404s some but not all of the five
    curve tables (OPTIONAL_TABLES) is not `available`; a 1.60-shaped row then
    gets no armour and no stats, the same honest gap as curves=None -- resolving
    armour from some tables and stats from none (or vice versa) would be a guess,
    not data."""
    incomplete = load_item_curves(
        read_csv(HERE / "fixtures/ItemArmorTotal.csv"),
        [],  # ItemArmorQuality 404'd for this hypothetical product
        read_csv(HERE / "fixtures/ItemArmorShield.csv"),
        read_csv(HERE / "fixtures/ArmorLocation.csv"),
        read_csv(HERE / "fixtures/RandPropPoints.csv"),
    )
    assert incomplete.available is False
    items = {i.id: i for r in build_all_1_60(incomplete) for i in r.items}
    assert 16866 not in items
    assert 30001 not in items


def test_a_stat_over_the_level_60_sanity_cap_is_an_error_not_a_guess():
    """A backstop behind JUNK_NAME_PATTERN, on the literal-amount path: three
    QA rows (Ring of Critical Testing x2, QATest +1000 Spell Dmg Ring) leaked
    into items/ with a 700-1000 stat once the curve resolver gave 1.60's items
    real values, because none of their names matched JUNK_NAME_PATTERN at the
    time. This asserts the cap independently of any name, on Era's literal path,
    so the next unnamed leak fails normalize instead of shipping a number no
    real level-60 item has (the highest real one in build 1.60.1.69893 is 46)."""
    rows = read_csv(HERE / "fixtures/ItemSparse.csv")
    row = next(r for r in rows if r["ID"] == "19325")  # Don Julio's Band: stat0 stamina
    row["ItemLevel"] = "60"
    row["StatModifier_bonusAmount_0"] = "999"
    with pytest.raises(ItemDataError, match="999 stamina"):
        build_class_items(
            rows,
            read_csv(HERE / "fixtures/Item.csv"),
            read_csv(HERE / "fixtures/ChrClasses.csv"),
            fixture_icons(),
            "1.0.0.1",
        )


def test_armour_over_the_level_60_sanity_cap_is_an_error_not_a_guess():
    """Same backstop, for armour (the highest real value in build 1.60.1.69893
    is 1,833)."""
    rows = read_csv(HERE / "fixtures/ItemSparse.csv")
    row = next(r for r in rows if r["ID"] == "16866")  # Helm of Might
    row["ItemLevel"] = "60"
    row["Resistances_0"] = "9999"
    with pytest.raises(ItemDataError, match="9999 armour"):
        build_class_items(
            rows,
            read_csv(HERE / "fixtures/Item.csv"),
            read_csv(HERE / "fixtures/ChrClasses.csv"),
            fixture_icons(),
            "1.0.0.1",
        )


def test_the_level_60_armour_cap_admits_the_real_item_that_motivated_raising_it():
    """MAX_LEVEL_60_ARMOR was raised from 2,000 to 2,100 (gear.py) because Era's
    own ItemSparse carries item 13375, "Crest of Retribution" -- a real rare
    shield, RequiredLevel 55, quality 3, with 2,057 armour -- which the old cap
    rejected. Pinned here at 2,057 so the next tuning pass can't silently
    re-break Era without a test noticing (see test_normalize_gear.py's other
    cap test for the "still rejects nonsense" side of this).

    Reuses fixtures/ItemSparse.csv's own 16866 row (present in both fixture
    CSVs already) rather than adding item 13375 to the shared fixture files:
    those are read by several other tests' golden-output comparisons, and a
    third fixture item would need every one of those goldens updated for a
    change that has nothing to do with them.
    """
    rows = read_csv(HERE / "fixtures/ItemSparse.csv")
    row = next(r for r in rows if r["ID"] == "16866")  # Helm of Might, AllowableClass warrior
    row["ItemLevel"] = "60"
    row["RequiredLevel"] = "55"
    row["OverallQualityID"] = "3"
    row["Resistances_0"] = "2057"
    records = build_class_items(
        rows,
        read_csv(HERE / "fixtures/Item.csv"),
        read_csv(HERE / "fixtures/ChrClasses.csv"),
        fixture_icons(),
        "1.0.0.1",
    )
    (warrior,) = (r for r in records if r.class_slug == "warrior")
    (item,) = (i for i in warrior.items if i.id == 16866)
    assert item.armor == 2057


def test_the_sanity_cap_also_guards_the_curve_resolved_path():
    """The same guard applies whether armour/stats came from a literal column or
    from item_curves.py's formula: a malformed StatPercentEditor on a 1.60-shaped
    row is caught the same way."""
    rows = read_csv(HERE / "fixtures/ItemSparse_1_60.csv")
    row = next(r for r in rows if r["ID"] == "16866")  # Helm of Might
    row["ItemLevel"] = "60"  # RandPropPoints fixture clamps to its ilvl-40 row
    row["StatPercentEditor_0"] = "200000"  # an absurd 2000%, not a real percent
    with pytest.raises(ItemDataError, match="stamina"):
        build_class_items(
            rows,
            read_csv(HERE / "fixtures/Item.csv"),
            read_csv(HERE / "fixtures/ChrClasses.csv"),
            fixture_icons(),
            "1.60.1.69893",
            fixture_curves(),
        )


def test_an_item_missing_from_the_item_table_is_skipped_with_a_warning(caplog):
    rows = read_csv(HERE / "fixtures/ItemSparse.csv")
    item_rows = [r for r in read_csv(HERE / "fixtures/Item.csv") if r["ID"] != "16866"]
    with caplog.at_level("WARNING"):
        records = build_class_items(
            rows,
            item_rows,
            read_csv(HERE / "fixtures/ChrClasses.csv"),
            fixture_icons(),
            "1.0.0.1",
        )
    assert 16866 not in {i.id for r in records for i in r.items}
    assert any("16866" in record.getMessage() for record in caplog.records)


def test_a_set_bonus_with_no_spell_text_is_emitted_but_warned_about(caplog):
    spell_rows = [{"ItemSetID": "209", "SpellID": "99999", "Threshold": "8"}]
    with caplog.at_level("WARNING"):
        records = build_item_sets(
            read_csv(HERE / "fixtures/ItemSet.csv"), spell_rows, fixture_spell_text()
        )
    assert records[0].bonuses == [ItemSetBonus(pieces=8, description="")]
    messages = [record.getMessage() for record in caplog.records]
    assert any("209" in m and "99999" in m for m in messages)


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


def _item_rows_with_icon(item_id: str, icon_file_data_id: str) -> list[dict[str, str]]:
    """The fixture Item rows with one row's IconFileDataID overridden in memory."""
    rows = read_csv(HERE / "fixtures/Item.csv")
    for row in rows:
        if row["ID"] == item_id:
            row["IconFileDataID"] = icon_file_data_id
    return rows


def test_an_item_the_client_has_no_icon_for_falls_back_to_the_placeholder(caplog):
    """IconFileDataID 0 means the client itself has no art; "" would 404 as icons/.webp."""
    with caplog.at_level("WARNING"):
        records = build_class_items(
            read_csv(HERE / "fixtures/ItemSparse.csv"),
            _item_rows_with_icon("16866", "0"),
            read_csv(HERE / "fixtures/ChrClasses.csv"),
            fixture_icons(),
            "1.0.0.1",
        )
    helm = {i.id: i for r in records for i in r.items}[16866]
    assert helm.icon == PLACEHOLDER_ICON
    assert helm.icon != ""
    messages = [record.getMessage() for record in caplog.records]
    assert any("16866" in m and "Helm of Might" in m for m in messages)


def test_a_nonzero_icon_id_that_names_no_file_also_falls_back(caplog):
    """Same branch: the client gave an id, but nothing in the manifest resolves it."""
    with caplog.at_level("WARNING"):
        records = build_class_items(
            read_csv(HERE / "fixtures/ItemSparse.csv"),
            _item_rows_with_icon("16866", "99999999"),
            read_csv(HERE / "fixtures/ChrClasses.csv"),
            fixture_icons(),
            "1.0.0.1",
        )
    helm = {i.id: i for r in records for i in r.items}[16866]
    assert helm.icon == PLACEHOLDER_ICON
    assert any("99999999" in record.getMessage() for record in caplog.records)


def test_every_emitted_item_icon_is_a_usable_file_name():
    """No item may carry an empty icon: the site builds icons/<icon>.webp from it."""
    assert all(i.icon for record in build_all() for i in record.items)


def build_with_sparse(rows: list[dict[str, str]]):
    return build_class_items(
        rows,
        read_csv(HERE / "fixtures/Item.csv"),
        read_csv(HERE / "fixtures/ChrClasses.csv"),
        fixture_icons(),
        "1.0.0.1",
    )


def sparse_rows_with(column: str, value: str | None) -> list[dict[str, str]]:
    """The fixture ItemSparse rows with one column of the Helm of Might row overridden."""
    rows = read_csv(HERE / "fixtures/ItemSparse.csv")
    for row in rows:
        if row["ID"] == "16866":
            row[column] = value  # type: ignore[assignment]
    return rows


@pytest.mark.parametrize(
    "column",
    [
        "Display_lang",
        "InventoryType",
        "RequiredLevel",
        "OverallQualityID",
        "ItemLevel",
        "MaxCount",
        "ItemSet",
        "AllowableClass",
        "Resistances_0",
        "Resistances_4",
    ],
)
def test_a_row_truncated_in_any_read_column_is_an_item_data_error(column: str):
    """Every column this module reads goes through _column, so an unreadable row is the
    single-valued ItemDataError the orchestrator catches, never a KeyError or TypeError
    that takes the whole run down with it."""
    with pytest.raises(ItemDataError, match=column):
        build_with_sparse(sparse_rows_with(column, None))


def test_a_non_numeric_value_in_a_numeric_column_is_an_item_data_error():
    with pytest.raises(ItemDataError, match="ItemLevel"):
        build_with_sparse(sparse_rows_with("ItemLevel", "sixty-six"))


def test_a_malformed_item_table_row_is_an_item_data_error_too():
    """The Item side of the join is read the same way as the ItemSparse side."""
    item_rows = read_csv(HERE / "fixtures/Item.csv")
    for row in item_rows:
        if row["ID"] == "16866":
            row["SubclassID"] = None  # type: ignore[assignment]
    with pytest.raises(ItemDataError, match="SubclassID"):
        build_class_items(
            read_csv(HERE / "fixtures/ItemSparse.csv"),
            item_rows,
            read_csv(HERE / "fixtures/ChrClasses.csv"),
            fixture_icons(),
            "1.0.0.1",
        )


def test_either_base_points_column_lands_on_the_same_number():
    # Classic Era exports EffectBasePoints as an integer; the 1.60 (Forever beta) client
    # exports EffectBasePointsF as a float. Both must land on the same whole number.
    #
    # Renamed from test_effect_base_points_read_the_float_column_when_the_build_has_it:
    # that name asserted the opposite of what _base_points actually does (it prefers
    # the INT column, per the bug described in its docstring -- reading the float
    # column whenever it merely existed silently zeroed every Era description, since
    # Era's own EffectBasePointsF is unused "0" padding). The two fixture rows below
    # are still mutually exclusive (each carries only one of the two columns), so this
    # test alone cannot tell the two orderings apart; see
    # test_the_int_column_wins_when_both_are_populated and
    # test_the_float_column_is_used_when_only_it_is_populated below for that.
    spell = [
        {
            "ID": "10",
            "NameSubtext_lang": "",
            "Description_lang": "$s1 dmg",
            "AuraDescription_lang": "",
        }
    ]
    misc = [
        {"SpellID": "10", "DurationIndex": "0", "SpellIconFileDataID": "0", "DifficultyID": "0"}
    ]
    era = [
        {
            "SpellID": "10",
            "EffectIndex": "0",
            "EffectBasePoints": "41",
            "EffectDieSides": "1",
            "EffectAuraPeriod": "0",
            "DifficultyID": "0",
        }
    ]
    beta = [
        {
            "SpellID": "10",
            "EffectIndex": "0",
            "EffectBasePointsF": "41.0",
            "EffectDieSides": "1",
            "EffectAuraPeriod": "0",
            "DifficultyID": "0",
        }
    ]
    assert load_spell_text(spell, misc, era, []).describe(10) == load_spell_text(
        spell, misc, beta, []
    ).describe(10)


def _effect_row(**overrides) -> dict:
    # EffectDieSides 0: _bounds() adds it to base_points, so a nonzero value
    # would shift the rendered number away from what these tests assert.
    row = {
        "SpellID": "10",
        "EffectIndex": "0",
        "EffectDieSides": "0",
        "EffectAuraPeriod": "0",
        "DifficultyID": "0",
    }
    row.update(overrides)
    return row


def _spell_and_misc() -> tuple[list[dict], list[dict]]:
    spell = [
        {
            "ID": "10",
            "NameSubtext_lang": "",
            "Description_lang": "$s1 dmg",
            "AuraDescription_lang": "",
        }
    ]
    misc = [
        {"SpellID": "10", "DurationIndex": "0", "SpellIconFileDataID": "0", "DifficultyID": "0"}
    ]
    return spell, misc


def test_the_int_column_wins_when_both_are_populated():
    """A build shaped like neither Era nor beta -- both EffectBasePoints and
    EffectBasePointsF populated and agreeing -- reads the int column, matching
    Era's own rule and item_curves.py's sibling _budget fallback."""
    spell, misc = _spell_and_misc()
    row = _effect_row(EffectBasePoints="41", EffectBasePointsF="41.0")
    assert load_spell_text(spell, misc, [row], []).describe(10) == "41 dmg"


def test_the_float_column_is_used_when_only_it_is_populated():
    spell, misc = _spell_and_misc()
    row = _effect_row(EffectBasePoints="", EffectBasePointsF="41.0")
    assert load_spell_text(spell, misc, [row], []).describe(10) == "41 dmg"


def test_base_points_columns_that_disagree_are_an_error_not_a_guess():
    """The whole premise of preferring one column over the other is that only
    one is ever real on a given build; a row where both are populated and
    disagree means that premise is false for this build, so this must not
    silently pick a side."""
    spell, misc = _spell_and_misc()
    row = _effect_row(EffectBasePoints="41", EffectBasePointsF="99.0")
    with pytest.raises(SpellTextError, match="disagreeing"):
        load_spell_text(spell, misc, [row], [])


def test_a_literal_zero_float_column_is_eras_padding_not_a_disagreement():
    """Era's own EffectBasePointsF is never truly empty -- the client exports the
    literal string "0" as unused padding on every single Era row, not "". A real,
    nonzero EffectBasePoints alongside that "0" padding is not a disagreement (see
    spell 543 effect 0 in build 1.15.9.69722's own SpellEffect export, which is
    exactly this shape and must not raise)."""
    spell, misc = _spell_and_misc()
    row = _effect_row(EffectBasePoints="165", EffectBasePointsF="0")
    assert load_spell_text(spell, misc, [row], []).describe(10) == "165 dmg"
