"""School spell damage and spell penetration reach the planner and the sim.

The 1.60 client states "+N Frost spell damage" as ItemSparse
StatModifier_bonusStat 87 (Azure Silk set, Icicle Rod), and the other schools
beside it: 84 holy (Shining set), 85 fire (Inferno Robe), 86 nature
(Stormcloth Boots), 88 shadow (Shadow Hood, Orb of Soran'ruk), 89 arcane
(Nether Force Wand, Tome of Arcane Domination); 47 is spell penetration
(Doomcaller's Robes, the PvP silk spaulders). Wowhead's supplement names the
same stats `holsplpwr`/`firsplpwr`/`natsplpwr`/`frosplpwr`/`shasplpwr`/
`arcsplpwr` and `splpen` (Elemental Mage Staff: +100 fire, +100 frost, 8
penetration). All of these used to map to None and vanish
(sim-accuracy round 4, physical row H3; casters lane).
"""

from pipeline import wowhead_items as wh
from pipeline.normalize.gear import STAT_BY_MODIFIER_ID
from pipeline.simdb.statmap import stat_index

SCHOOL_POWER_BY_MODIFIER_ID = {
    84: "holy_power",
    85: "fire_power",
    86: "nature_power",
    87: "frost_power",
    88: "shadow_power",
    89: "arcane_power",
    47: "spell_penetration",
}

SCHOOL_POWER_BY_WOWHEAD_KEY = {
    "holsplpwr": "holy_power",
    "firsplpwr": "fire_power",
    "natsplpwr": "nature_power",
    "frosplpwr": "frost_power",
    "shasplpwr": "shadow_power",
    "arcsplpwr": "arcane_power",
    "splpen": "spell_penetration",
}


def test_item_sparse_school_damage_and_penetration_have_planner_keys() -> None:
    for modifier_id, key in SCHOOL_POWER_BY_MODIFIER_ID.items():
        assert STAT_BY_MODIFIER_ID[modifier_id] == key, modifier_id


def test_wowhead_school_damage_and_penetration_have_planner_keys() -> None:
    for wowhead_key, key in SCHOOL_POWER_BY_WOWHEAD_KEY.items():
        assert wh.STAT_KEYS.get(wowhead_key) == key, wowhead_key


def test_every_new_key_is_an_engine_stat() -> None:
    for key in {*SCHOOL_POWER_BY_MODIFIER_ID.values(), *SCHOOL_POWER_BY_WOWHEAD_KEY.values()}:
        stat_index(key)
