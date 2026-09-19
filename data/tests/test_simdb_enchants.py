import logging
import re
from pathlib import Path

import pytest

from pipeline.csvio import read_csv
from pipeline.manifest import newest_build
from pipeline.simdb.enchants import EnchantDataError, build_sim_enchants
from pipeline.simdb.equip import index_spell_effects
from pipeline.simproto import pb

SIM = Path(__file__).parent / "fixtures" / "sim"
SHIPPABLE = [352, 803, 930, 931, 934, 1900, 2504]

#: The committed build's own level-60 combatratings.txt row (see
#: tests/fixtures/sim/combatratings.txt).
RATING_FACTORS = {
    "hit": 10.0,
    "crit": 14.0,
    "dodge": 12.0,
    "parry": 15.0,
    "block": 5.0,
    "defense": 1.0,
}


def rows():
    return read_csv(SIM / "SpellItemEnchantment.csv")


def effects():
    return index_spell_effects(read_csv(SIM / "SpellEffect.csv"))


def enchants():
    return {row.effect_id: row for row in build_sim_enchants(rows(), effects(), RATING_FACTORS)}


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
    with pytest.raises(EnchantDataError, match=r"enchant 2505.*999"):
        build_sim_enchants(
            read_csv(SIM / "SpellItemEnchantment_bad.csv"), effects(), RATING_FACTORS
        )


def test_every_enchant_row_is_emitted_and_sorted():
    ids = [row.effect_id for row in build_sim_enchants(rows(), effects(), RATING_FACTORS)]
    assert ids == sorted(ids) == SHIPPABLE


def test_a_direct_stat_enchant_that_is_a_combat_rating_is_converted():
    """A synthetic row granting 20 crit rating (mod id 32) through its own
    EFFECT_STAT slot, not an equip spell -- the same rating-to-percentage
    conversion items.py applies, exercised here on the enchant path. At this
    fixture's level-60 crit factor (14) that is 20 / 14 %."""
    rating_row = dict(rows()[0])
    rating_row.update(
        {
            "ID": "9001",
            "Effect_0": "5",
            "EffectPointsMin_0": "20",
            "EffectArg_0": "32",
            "Effect_1": "0",
            "EffectPointsMin_1": "0",
            "EffectArg_1": "0",
            "Effect_2": "0",
            "EffectPointsMin_2": "0",
            "EffectArg_2": "0",
        }
    )
    built = {
        row.effect_id: row
        for row in build_sim_enchants([rating_row], effects(), RATING_FACTORS)
    }
    assert built[9001].stats[pb.Stat.Value("StatCrit")] == pytest.approx(20.0 / 14.0)


def test_a_weapon_skill_only_equip_spell_is_dropped_with_a_warning(caplog):
    """`pb.SimEnchant` has no weapon-skill field (proto/common.proto:897-900:
    effect_id, stats only). 934 'Sword Skill +3' grants WeaponSkillSwords
    through spell 900001 and nothing else, so it must still emit an (empty)
    row -- loudly, not silently."""
    with caplog.at_level(logging.WARNING, logger="pipeline.simdb.enchants"):
        result = enchants()
    assert list(result[934].stats) == []
    assert "934" in caplog.text
    assert "WeaponSkillSwords" in caplog.text


def test_the_live_build_warns_about_every_weapon_skill_only_enchant(caplog):
    """Pinned against the committed build so a table move that adds or drops
    a weapon-skill-only enchant (e.g. 'Sword Skill +1' and kin) is noticed,
    not silently absorbed into an empty stats array."""
    raw = Path("builds") / newest_build() / "raw"
    if not (raw / "SpellItemEnchantment.csv").exists():
        pytest.skip("raw client tables have not been fetched for this build")
    with caplog.at_level(logging.WARNING, logger="pipeline.simdb.enchants"):
        build_sim_enchants(
            read_csv(raw / "SpellItemEnchantment.csv"),
            index_spell_effects(read_csv(raw / "SpellEffect.csv")),
            RATING_FACTORS,
        )
    dropped_ids = {
        match.group(1)
        for record in caplog.records
        if (match := re.search(r"enchant (\d+) grants", record.getMessage()))
    }
    assert len(dropped_ids) == 40
