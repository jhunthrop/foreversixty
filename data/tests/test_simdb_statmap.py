import pytest

from pipeline.normalize.gear import STAT_BY_MODIFIER_ID
from pipeline.simdb.statmap import (
    PROTO_STAT_ALIASES,
    StatMapError,
    stat_array,
    stat_index,
    stat_keys,
    weapon_skill_array,
    weapon_skill_index,
)
from pipeline.simproto import pb


def test_a_plain_stat_resolves_to_the_proto_index():
    assert stat_index("strength") == pb.Stat.Value("StatStrength")
    assert stat_index("agility") == pb.Stat.Value("StatAgility")
    assert stat_index("attack_power") == pb.Stat.Value("StatAttackPower")


def test_every_planner_stat_key_has_an_engine_stat():
    """gear.py's vocabulary is the pipeline's one set of stat names. A key the
    planner emits that simdb cannot map would be a stat the sim silently loses."""
    for key in STAT_BY_MODIFIER_ID.values():
        if key is not None:
            assert stat_index(key) >= 0, key


def test_hit_and_crit_resolve_through_the_alias_list():
    """The engine lane renames MeleeHit/SpellHit to Hit and MeleeCrit/SpellCrit to
    Crit. Both spellings resolve here, so the rename lands without a pipeline edit."""
    known = set(pb.Stat.keys())
    for key in ("hit", "crit", "spell_hit", "spell_crit"):
        candidates = [name for name in PROTO_STAT_ALIASES[key] if name in known]
        assert candidates, key
        assert stat_index(key) == pb.Stat.Value(candidates[0])


def test_an_unknown_key_is_an_error_not_a_guess():
    with pytest.raises(StatMapError, match="cleverness"):
        stat_index("cleverness")


def test_a_key_whose_engine_stats_all_vanished_is_an_error():
    with pytest.raises(StatMapError, match="StatNope"):
        stat_index("__probe__")


def test_stat_array_is_indexed_by_the_proto_enum():
    array = stat_array({"strength": 10, "stamina": 4})
    assert array[pb.Stat.Value("StatStrength")] == 10.0
    assert array[pb.Stat.Value("StatStamina")] == 4.0


def test_stat_array_is_truncated_after_the_last_non_zero_entry():
    array = stat_array({"strength": 10})
    assert len(array) == pb.Stat.Value("StatStrength") + 1


def test_an_all_zero_stat_array_is_empty():
    assert stat_array({"strength": 0, "agility": 0}) == []
    assert stat_array({}) == []


def test_repeated_keys_are_summed():
    """`stat_array` accumulates rather than assigns, so it takes pairs as well
    as a mapping: a dict literal cannot hold the same key twice, and a test
    that passes one proves nothing about the `+=`."""
    array = stat_array([("strength", 3.0), ("strength", 4.0), ("agility", 1.0)])
    assert array[pb.Stat.Value("StatStrength")] == 7.0
    assert array[pb.Stat.Value("StatAgility")] == 1.0


def test_two_keys_that_may_alias_to_one_stat_do_not_lose_a_point():
    """This is why the accumulation exists. Today `hit` and `spell_hit` resolve
    to StatMeleeHit and StatSpellHit; after the engine lane's collapse they
    both resolve to StatHit and land in one slot. Summing the array holds
    either way, so the test does not have to be rewritten on the day it lands."""
    assert sum(stat_array([("hit", 1.0), ("spell_hit", 2.0)])) == 3.0


def test_weapon_skill_pairs_are_summed_too():
    array = weapon_skill_array([("WeaponSkillSwords", 2.0), ("WeaponSkillSwords", 3.0)])
    assert array[pb.WeaponSkill.Value("WeaponSkillSwords")] == 5.0


def test_weapon_skills_use_their_own_enum():
    assert weapon_skill_index("WeaponSkillSwords") == pb.WeaponSkill.Value("WeaponSkillSwords")
    array = weapon_skill_array({"WeaponSkillSwords": 3})
    assert array[pb.WeaponSkill.Value("WeaponSkillSwords")] == 3.0
    assert weapon_skill_array({}) == []


def test_an_unknown_weapon_skill_is_an_error():
    with pytest.raises(StatMapError, match="WeaponSkillSpoons"):
        weapon_skill_index("WeaponSkillSpoons")


def test_stat_keys_is_the_inverse_of_stat_array():
    array = stat_array({"strength": 10, "crit": 2.5, "attack_power": 40})
    assert stat_keys(array) == {"strength": 10.0, "crit": 2.5, "attack_power": 40.0}


def test_stat_keys_drops_zero_amounts():
    assert stat_keys([0, 0, 0, 0]) == {}


def test_stat_keys_reads_a_shared_index_back_as_the_first_name():
    """Forever merges spell and melee hit into one stat, so `hit` and
    `spell_hit` are the same index; the array can only be read back as one
    of them, and PROTO_STAT_ALIASES' order decides which."""
    assert stat_keys(stat_array({"spell_hit": 3})) == {"hit": 3.0}


def test_stat_keys_refuses_an_index_no_planner_key_covers():
    # StatEnergy: a resource stat no enchant, suffix or item column has ever
    # granted a flat bonus of, so PROTO_STAT_ALIASES has never needed a key
    # for it (unlike StatMana, StatHealth and StatBonusArmor, which the real
    # fork database's enchants and suffixes do state).
    array = [0.0] * (len(pb.Stat.keys()))
    array[pb.Stat.Value("StatEnergy")] = 5
    with pytest.raises(StatMapError, match="no planner key"):
        stat_keys(array)
