import pytest

from pipeline.normalize.effects import EQUIP_TRIGGER, EffectIndex, UnknownAuraError
from pipeline.spelltext import SpellRow, SpellText


def effect(spell_id, aura, amount, misc=0):
    return {
        "SpellID": str(spell_id),
        "Effect": "6",
        "EffectAura": str(aura),
        "EffectBasePointsF": str(amount),
        "EffectMiscValue_0": str(misc),
    }


def _spell_text(descriptions):
    return SpellText(
        {
            spell_id: SpellRow(description=text, duration_ms=None, icon_file_id=0)
            for spell_id, text in descriptions.items()
        }
    )


def index(item_effects, spell_effects, descriptions=None):
    return EffectIndex(item_effects, [], spell_effects, _spell_text(descriptions or {}))


def test_an_equip_aura_becomes_a_stat():
    """SPELL_AURA_MOD_DAMAGE_DONE (13) with an all-magic school mask is spell power."""
    idx = index(
        [{"ParentItemID": "9", "SpellID": "100", "TriggerType": str(EQUIP_TRIGGER)}],
        [effect(100, 13, 23, misc=126)],
    )
    assert idx.stats(9) == {"spell_power": 23}


def test_a_stat_aura_uses_its_misc_value():
    """SPELL_AURA_MOD_STAT (29) with misc 1 is agility."""
    idx = index(
        [{"ParentItemID": "9", "SpellID": "101", "TriggerType": str(EQUIP_TRIGGER)}],
        [effect(101, 29, 8, misc=1)],
    )
    assert idx.stats(9) == {"agility": 8}


def test_two_effects_on_one_item_accumulate():
    idx = index(
        [
            {"ParentItemID": "9", "SpellID": "100", "TriggerType": str(EQUIP_TRIGGER)},
            {"ParentItemID": "9", "SpellID": "102", "TriggerType": str(EQUIP_TRIGGER)},
        ],
        [effect(100, 13, 10, misc=126), effect(102, 13, 5, misc=126)],
    )
    assert idx.stats(9) == {"spell_power": 15}


def test_a_use_effect_contributes_text_and_no_stats():
    idx = index(
        [{"ParentItemID": "9", "SpellID": "103", "TriggerType": "0"}],
        [effect(103, 0, 0)],
        {103: "Restores 400 mana."},
    )
    assert idx.stats(9) == {}
    assert idx.text(9) == "Restores 400 mana."


def test_an_unmapped_aura_raises_rather_than_being_guessed():
    idx = index(
        [{"ParentItemID": "9", "SpellID": "104", "TriggerType": str(EQUIP_TRIGGER)}],
        [effect(104, 999, 3)],
    )
    with pytest.raises(UnknownAuraError, match="999"):
        idx.stats(9)


def test_every_stat_key_equip_can_produce_is_in_the_planners_vocabulary():
    """R4: item `stats` keys are the planner's own vocabulary -- exactly the
    keys of `pipeline.simdb.statmap.PROTO_STAT_ALIASES`, which
    `pipeline/simdb/statmap.py` documents as `STAT_BY_MODIFIER_ID`'s values
    plus the keys `pipeline.simdb.equip` reads out of an equip spell.

    `equip.py` is the only source of the stat keys `EffectIndex.stats` can
    ever produce -- this module deliberately keeps no aura table of its own
    (see effects.py's module docstring) -- so this asserts directly against
    `equip.py`'s real tables (`STAT_AURAS`, the by-misc-value branches'
    outputs) rather than against a second table that could drift from them,
    the way the brief's own `AURA_STAT_KEYS` did.
    """
    from pipeline.simdb import equip
    from pipeline.simdb.statmap import PROTO_STAT_ALIASES

    produced = (
        set(equip.STAT_AURAS.values())
        | set(equip.SCHOOL_POWER.values())
        | set(equip.SCHOOL_RESISTANCE.values())
        | set(equip.MOD_STAT_BY_INDEX.values())
        | {"spell_power", "armor", "defense"}
    )
    for key in produced:
        assert key in PROTO_STAT_ALIASES, key


def test_an_item_with_no_effects_is_empty():
    idx = index([], [])
    assert idx.stats(9) == {}
    assert idx.text(9) == ""
