"""SpellItemEnchantment -> the engine's SimEnchant rows.

Each row has three effect slots. On build 1.60.1.69893, over 2,216 rows and
their 6,648 slots: 2,473 slots are type 3 -- an equip spell whose id is in the
paired EffectArg -- against 71 type 5 (a direct ITEM_MOD stat) and 48 type 4
(a direct resistance). 1,995 rows name at least one equip spell and 1,334 end
up carrying stats, so the equip-spell path is the enchant table, not an edge
case; `pipeline.simdb.equip.spell_bonus` does that work, the same function the
items use.

Types 1 (proc spell), 2 (flat weapon damage) and 7 (use spell) are behaviour
the engine hand-writes in `sim/common/enchant_effects.go`. Their rows are still
emitted, with no stats, because the engine resolves an enchant by effect id and
a missing row is an unknown enchant.

The 1.60 table has no EffectPointsMax_<n>; EffectPointsMin_<n> is the amount.
"""

from __future__ import annotations

from pipeline.normalize.gear import RESISTANCE_KEYS, STAT_BY_MODIFIER_ID
from pipeline.simdb.equip import spell_bonus
from pipeline.simdb.statmap import stat_array
from pipeline.simproto import pb

EFFECT_SLOTS = range(3)
EFFECT_EQUIP_SPELL = 3
EFFECT_RESISTANCE = 4
EFFECT_STAT = 5

#: SpellItemEnchantment's resistance index is ItemSparse's Resistances_<n>
#: index, so the five schools are `gear.RESISTANCE_KEYS` and not a second copy
#: of it -- index 1 is holy resistance, which has no engine Stat, and index 0
#: is armour, which the planner reads from its own column and so is the one
#: entry gear.RESISTANCE_KEYS does not carry.
ARMOR_RESISTANCE_INDEX = 0
ENCHANT_RESISTANCE_BY_INDEX: dict[int, str] = {
    ARMOR_RESISTANCE_INDEX: "armor",
    **RESISTANCE_KEYS,
}


class EnchantDataError(ValueError):
    """An enchant row states something this module will not guess at."""


def build_sim_enchants(
    enchant_rows: list[dict[str, str]],
    effects_by_spell: dict[int, list[dict[str, str]]],
) -> list[pb.SimEnchant]:
    enchants: list[pb.SimEnchant] = []
    for row in sorted(enchant_rows, key=lambda r: int(r["ID"])):
        stats: dict[str, float] = {}
        for slot in EFFECT_SLOTS:
            effect = int(row[f"Effect_{slot}"])
            amount = float(int(row[f"EffectPointsMin_{slot}"]))
            arg = int(row[f"EffectArg_{slot}"])
            if effect == EFFECT_EQUIP_SPELL and arg:
                for key, value in spell_bonus([arg], effects_by_spell).stats.items():
                    stats[key] = stats.get(key, 0.0) + value
            elif effect == EFFECT_STAT and amount:
                if arg not in STAT_BY_MODIFIER_ID:
                    raise EnchantDataError(
                        f"enchant {row['ID']} grants unknown stat modifier id {arg}; "
                        f"add it to STAT_BY_MODIFIER_ID in pipeline/normalize/gear.py"
                    )
                key = STAT_BY_MODIFIER_ID[arg]
                if key is not None:
                    stats[key] = stats.get(key, 0.0) + amount
            elif effect == EFFECT_RESISTANCE and amount:
                key = ENCHANT_RESISTANCE_BY_INDEX.get(arg)
                if key:
                    stats[key] = stats.get(key, 0.0) + amount
        enchants.append(pb.SimEnchant(effect_id=int(row["ID"]), stats=stat_array(stats)))
    return enchants
