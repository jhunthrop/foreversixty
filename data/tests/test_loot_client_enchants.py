from pathlib import Path

import pytest

from pipeline.loot.client_enchants import (
    ClientEnchantError,
    ClientTables,
    CuratedEnchant,
    items_by_spell,
    load_curated,
    rebind_to_client,
    taught_enchants,
    with_client_enchants,
)
from pipeline.loot.gear import EquipRestriction
from pipeline.models import EnchantRecord

CURATED = Path("curated")


def record(enchant_id: int, spell_id: int, item_id: int = 0) -> EnchantRecord:
    return EnchantRecord(
        id=enchant_id,
        name=f"spell {spell_id}",
        icon="icon",
        slots=["hands"],
        item_types=["normal"],
        classes=[],
        stats={"agility": 7.0},
        phase=0,
        spell_id=spell_id,
        item_id=item_id,
    )


def tables(**overrides) -> ClientTables:
    base = {
        "taught": {20012: 8206, 1249057: 8218, 1254770: 8492, 1219580: 7662},
        "stats": {
            8206: {"agility": 10.0},
            8218: {"spell_power": 6.0},
            8492: {"armor": 40.0, "spell_power": 6.0},
            7662: {"spell_power": 65.0},
        },
        "spell_names": {
            1249057: "Enchant Necklace - Spell Power",
            1254770: "Mystic Rugged Armor Kit",
            1219580: "Enchant 2H Weapon - Spellblasting",
        },
        "restrictions": {
            1249057: EquipRestriction(4, 1, 4),
            1254770: EquipRestriction(4, 30, 1050016),
            1219580: EquipRestriction(2, 1378, 0),
        },
        "items_by_spell": {1254770: frozenset({252487})},
    }
    return ClientTables(**{**base, **overrides})


def icon(item_id: int, spell_id: int) -> str:
    return f"icon-{item_id}-{spell_id}"


def test_taught_enchants_reads_effect_53_only():
    rows = [
        {"SpellID": "20012", "Effect": "53", "EffectMiscValue_0": "8206"},
        {"SpellID": "20012", "Effect": "6", "EffectMiscValue_0": "5"},
        {"SpellID": "1", "Effect": "53", "EffectMiscValue_0": "0"},
    ]
    assert taught_enchants(rows) == {20012: 8206}


def test_a_fork_row_takes_the_clients_id_and_stats():
    [rebound] = rebind_to_client([record(1887, 20012)], tables().taught, tables().stats)
    assert (rebound.id, rebound.stats) == (8206, {"agility": 10.0})


def test_a_fork_row_the_client_does_not_teach_keeps_its_id():
    [kept] = rebind_to_client([record(904, 13815)], tables().taught, tables().stats)
    assert kept.id == 904


def test_a_taught_id_with_no_client_row_is_refused():
    with pytest.raises(ClientEnchantError, match="8206"):
        rebind_to_client([record(1887, 20012)], {20012: 8206}, {})


def test_curated_additions_read_slots_and_shape_from_the_client():
    curated = [
        CuratedEnchant(1249057, 0),
        CuratedEnchant(1254770, 252487),
        CuratedEnchant(1219580, 0),
    ]
    rows = {r.spell_id: r for r in with_client_enchants([], curated, tables(), icon)}
    assert (rows[1249057].id, rows[1249057].slots, rows[1249057].item_types) == (
        8218,
        ["neck"],
        ["normal"],
    )
    kit = rows[1254770]
    assert (kit.slots, kit.item_types, kit.icon) == (
        ["chest", "feet", "hands", "legs"],
        ["kit"],
        "icon-252487-1254770",
    )
    weapon = rows[1219580]
    assert (weapon.slots, weapon.item_types, weapon.stats) == (
        ["main_hand", "off_hand"],
        ["two_hand"],
        {"spell_power": 65.0},
    )


def test_a_curated_spell_the_fork_already_has_is_refused():
    with pytest.raises(ClientEnchantError, match="already"):
        with_client_enchants([record(8218, 1249057)], [CuratedEnchant(1249057, 0)], tables(), icon)


def test_a_curated_spell_that_applies_no_enchant_is_refused():
    with pytest.raises(ClientEnchantError, match="no enchant"):
        with_client_enchants([], [CuratedEnchant(999, 0)], tables(), icon)


def test_a_curated_kit_item_that_does_not_cast_the_spell_is_refused():
    with pytest.raises(ClientEnchantError, match="does not cast"):
        with_client_enchants([], [CuratedEnchant(1254770, 1)], tables(), icon)


def test_items_by_spell_joins_the_item_effect_links():
    effects = [{"ID": "5", "SpellID": "1254770"}]
    links = [{"ItemEffectID": "5", "ItemID": "252487"}, {"ItemEffectID": "6", "ItemID": "1"}]
    assert items_by_spell(effects, links) == {1254770: frozenset({252487})}


def test_the_curated_list_names_teaching_spells_with_a_source():
    """Floors, not counts: every row has a spell and says where it comes from."""
    entries = load_curated(CURATED)
    assert len(entries) >= 4
    assert {1219580, 1220624, 1249057, 1254770} <= {entry.spell_id for entry in entries}
