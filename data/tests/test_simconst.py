from pathlib import Path

import pytest

from pipeline.simconst import SPELL_FAMILY_BY_CLASS_SLUG, build_spell_constants

FIXTURES = Path(__file__).parent / "fixtures" / "sim"


def by_slug():
    return {record.class_slug: record for record in build_spell_constants("9.9.9.9", FIXTURES)}


def test_there_is_one_record_per_class_even_when_it_has_no_spells():
    records = by_slug()
    assert set(records) == set(SPELL_FAMILY_BY_CLASS_SLUG)
    assert records["warrior"].family == SPELL_FAMILY_BY_CLASS_SLUG["warrior"] == 4
    assert records["mage"].family == 3
    assert records["warrior"].build == "9.9.9.9"
    assert records["paladin"].spells == {}


def test_an_unknown_spell_class_set_is_ignored():
    """Sets 0, 1 and 13 are not player classes; spell 99999 must vanish."""
    assert not any("99999" in record.spells for record in by_slug().values())


def test_bloodthirst_carries_the_numbers_the_engine_needs():
    spell = by_slug()["warrior"].spells["23894"]
    assert spell.name == "Bloodthirst"
    assert spell.rank == 4
    assert spell.school_mask == 1
    assert spell.cast_time_ms == 0
    assert spell.gcd_ms == 1500
    assert spell.cooldown_ms == 0
    assert spell.category_cooldown_ms == 6000
    assert spell.duration_ms == 10000
    assert spell.cost == 300  # rage is stored times ten
    assert spell.cost_type == 1
    assert spell.spell_level == 60
    assert spell.family_mask == [33554432, 1024, 0, 0]
    assert [effect.index for effect in spell.effects] == [0, 1, 2]
    assert spell.effects[0].amount == 48
    assert spell.effects[0].effect == 2
    assert spell.effects[0].sp_coefficient == pytest.approx(1.0)
    assert spell.effects[0].ap_coefficient == 0.0


def test_frostbolt_carries_its_cast_time_duration_and_coefficient():
    spell = by_slug()["mage"].spells["25304"]
    assert spell.rank == 11
    assert spell.school_mask == 16
    assert spell.cast_time_ms == 3000
    assert spell.duration_ms == 9000
    assert spell.cost == 290
    assert spell.cost_type == 0
    assert spell.family_mask == [1075314720, 0, 0, 0]
    assert spell.effects[1].amount == 475
    assert spell.effects[1].sp_coefficient == pytest.approx(0.814)


def test_a_negative_effect_amount_is_kept_as_stated():
    """Frostbolt's slow is -40 percent; a constants file that dropped the sign
    would make the engine speed the target up."""
    assert by_slug()["mage"].spells["25304"].effects[0].amount == -40


def test_a_spell_present_in_no_other_table_is_all_zeros_not_an_error():
    """Most passives appear in SpellClassOptions and in no cooldown, power,
    level, misc or effect table. A missing row is zero, not a KeyError."""
    spell = by_slug()["warrior"].spells["88888"]
    assert spell.name == "Deep Wounds"
    assert spell.rank == 0
    assert (spell.cast_time_ms, spell.gcd_ms, spell.cooldown_ms) == (0, 0, 0)
    assert (spell.category_cooldown_ms, spell.duration_ms, spell.school_mask) == (0, 0, 0)
    assert (spell.cost, spell.cost_type, spell.spell_level) == (0, 0, 0)
    assert spell.family_mask == [1, 0, 0, 0]
    assert spell.effects == []


def test_spell_ids_are_written_in_a_stable_numeric_order():
    for record in by_slug().values():
        keys = list(record.spells)
        assert keys == sorted(keys, key=int)
