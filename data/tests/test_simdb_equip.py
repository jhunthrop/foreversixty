from pathlib import Path

import pytest

from pipeline.csvio import read_csv
from pipeline.simdb.equip import (
    EquipEffectError,
    equip_bonuses,
    index_spell_effects,
    item_effect_spells,
    spell_bonus,
)

FIXTURES = Path(__file__).parent / "fixtures" / "sim"

#: The items a caller would have kept. Item 100007 is deliberately absent: its
#: on-equip spell applies an unclassified aura, so any helper that included it
#: by default would raise in eight tests that mean to assert something else.
#: The two tests that want the raise pass an explicit set.
KEPT_ITEMS = frozenset({21407, 100001, 100002, 100003, 100004, 100005, 100006})


def effects():
    return index_spell_effects(read_csv(FIXTURES / "SpellEffect.csv"))


def bonuses(item_ids=KEPT_ITEMS):
    return equip_bonuses(
        read_csv(FIXTURES / "ItemEffect.csv"),
        read_csv(FIXTURES / "ItemXItemEffect.csv"),
        effects(),
        item_ids,
    )


def test_a_flat_stat_aura_is_read_as_the_client_states_it():
    """Mace of Unending Life (21407) grants +140 attack power through spell
    26153, which exists in no ItemSparse column."""
    assert bonuses()[21407].stats == {"attack_power": 140.0}


def test_two_effects_of_one_spell_and_two_spells_of_one_item_both_land():
    assert bonuses()[100001].stats == {
        "attack_power": 24.0,
        "ranged_attack_power": 24.0,
        "spell_power": 85.0,
        "healing": 85.0,
    }


def test_all_school_damage_done_is_spell_power():
    assert bonuses()[100001].stats["spell_power"] == 85.0


def test_physical_damage_done_is_not_a_stat():
    item = bonuses()[100004]
    assert item.stats == {}
    assert item.bonus_physical_damage == 2.0


def test_mod_stat_reads_its_stat_index_from_misc_value():
    assert spell_bonus([9103], effects()).stats == {"strength": 8.0}


def test_mod_stat_minus_one_grants_all_five():
    assert spell_bonus([900008], effects()).stats == {
        "strength": 5.0,
        "agility": 5.0,
        "stamina": 5.0,
        "intellect": 5.0,
        "spirit": 5.0,
    }


def test_mod_skill_defense_is_a_stat_and_a_weapon_line_is_a_weapon_skill():
    item = bonuses()[100002]
    assert item.stats == {"defense": 7.0}
    assert item.weapon_skills == {"WeaponSkillSwords": 3.0}


def test_a_profession_skill_grants_nothing():
    assert 100003 not in bonuses()


def test_a_resistance_aura_reads_its_school_mask():
    item = bonuses()[100005]
    assert item.stats == {"frost_res": 10.0}


def test_a_proc_effect_is_not_an_equip_bonus():
    """Item 100006's only item effect is TriggerType 2. Its resistance shred is
    on-proc behaviour the engine hand-writes, not a stat the wearer carries."""
    assert 100006 not in bonuses()


def test_a_reviewed_non_stat_aura_grants_nothing_and_does_not_raise():
    """Aura 155 is water breathing: real, classified, and not a stat."""
    assert spell_bonus([5227], effects()).is_empty()


def test_flat_block_value_is_a_stat():
    """Aura 274 is the flat 'Block Value NN' the 1.60 client uses."""
    assert spell_bonus([23731], effects()).stats == {"block_value": 19.0}


def test_an_item_the_caller_filtered_out_is_never_looked_at():
    """Item 100007's spell applies an unclassified aura. Scoping to the kept
    items is what stops a gamemaster row failing the whole run."""
    assert bonuses(item_ids={21407}).keys() == {21407}
    with pytest.raises(EquipEffectError, match="9999"):
        bonuses(item_ids={21407, 100007})


def test_spell_bonus_is_reusable_for_enchants():
    """Enchant 352 '+8 Strength' is an equip-spell enchant pointing at 9103."""
    assert spell_bonus([9103], effects()).stats == {"strength": 8.0}


def test_an_unclassified_aura_is_an_error_not_a_guess():
    with pytest.raises(EquipEffectError, match="9999"):
        spell_bonus([900006], effects())


def test_an_unclassified_skill_line_is_an_error_not_a_guess():
    with pytest.raises(EquipEffectError, match="7777"):
        spell_bonus([900007], effects())


def test_the_legacy_schema_links_through_parent_item_id():
    """Classic Era's ItemEffect carries ParentItemID and there is no
    ItemXItemEffect table to fetch."""
    legacy = [
        {"TriggerType": "1", "SpellID": "26153", "ParentItemID": "21407"},
        {"TriggerType": "2", "SpellID": "900005", "ParentItemID": "21407"},
    ]
    assert item_effect_spells(legacy, [], {21407}) == {21407: [26153]}


def test_a_difficulty_variant_effect_row_is_ignored():
    rows = [
        {
            "SpellID": "700",
            "DifficultyID": "1",
            "EffectIndex": "0",
            "Effect": "6",
            "EffectAura": "99",
            "EffectBasePointsF": "99",
            "EffectMiscValue_0": "0",
        }
    ]
    assert spell_bonus([700], index_spell_effects(rows)).is_empty()
