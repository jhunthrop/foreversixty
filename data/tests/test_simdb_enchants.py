from pathlib import Path

import pytest

from pipeline.csvio import read_csv
from pipeline.simdb.enchants import EnchantDataError, build_sim_enchants
from pipeline.simdb.equip import index_spell_effects
from pipeline.simproto import pb

SIM = Path(__file__).parent / "fixtures" / "sim"
SHIPPABLE = [352, 803, 930, 931, 1900, 2504]


def rows():
    return read_csv(SIM / "SpellItemEnchantment.csv")


def effects():
    return index_spell_effects(read_csv(SIM / "SpellEffect.csv"))


def enchants():
    return {row.effect_id: row for row in build_sim_enchants(rows(), effects())}


def test_an_equip_spell_enchant_resolves_through_its_spell():
    """352 '+8 Strength' is effect type 3 pointing at spell 9103, which is one
    apply-aura effect: aura 29, amount 8, misc 0."""
    assert enchants()[352].stats[pb.Stat.Value("StatStrength")] == 8.0


def test_a_direct_stat_enchant_uses_its_own_columns():
    assert enchants()[930].stats[pb.Stat.Value("StatAgility")] == 15.0


def test_a_direct_resistance_enchant_uses_the_resistance_index():
    assert enchants()[931].stats[pb.Stat.Value("StatFireResistance")] == 5.0


def test_a_proc_enchant_carries_no_stats():
    """Fiery Weapon is effect type 1. Its behaviour lives in the engine's
    enchant_effects.go, and it must still be present so the id resolves."""
    assert list(enchants()[803].stats) == []


def test_an_enchant_that_mixes_a_stat_spell_with_a_proc_keeps_the_stat():
    assert enchants()[1900].stats[pb.Stat.Value("StatDefense")] == 7.0


def test_a_stat_the_planner_does_not_track_is_dropped_not_guessed():
    """ITEM_MOD 0 is mana; gear.py maps it to None on purpose."""
    assert list(enchants()[2504].stats) == []


def test_an_unknown_stat_modifier_id_is_an_error():
    with pytest.raises(EnchantDataError, match="999"):
        build_sim_enchants(read_csv(SIM / "SpellItemEnchantment_bad.csv"), effects())


def test_every_enchant_row_is_emitted_and_sorted():
    ids = [row.effect_id for row in build_sim_enchants(rows(), effects())]
    assert ids == sorted(ids) == SHIPPABLE
