"""SpellItemEnchantment -> the engine's SimEnchant rows.

Each row has three effect slots. On build 1.60.1.69893, over 2,216 rows and
their 6,648 slots: 2,473 slots are type 3 -- an equip spell whose id is in the
paired EffectArg -- against 71 type 5 (a direct ITEM_MOD stat) and 48 type 4
(a direct resistance). 1,995 rows name at least one equip spell and 1,329 end
up carrying stats, so the equip-spell path is the enchant table, not an edge
case; `pipeline.simdb.equip.spell_bonus` does that work, the same function the
items use.

Types 1 (proc spell), 2 (flat weapon damage) and 7 (use spell) are behaviour
the engine hand-writes in `sim/common/enchant_effects.go`. Their rows are still
emitted, with no stats, because the engine resolves an enchant by effect id and
a missing row is an unknown enchant.

The 1.60 table has no EffectPointsMax_<n>; EffectPointsMin_<n> is the amount.

`pb.SimEnchant` (proto/common.proto:897-900) carries only `effect_id` and
`stats` -- no field for a weapon skill or a flat weapon-damage bonus, both of
which a spell in the equip-spell path can grant (`SpellBonus.weapon_skills`,
`.bonus_physical_damage`). "Sword Skill +1" and its kin are real, named
enchants whose equip spell grants nothing else, so silently keeping only
`.stats` would ship them as empty rows with no trace. `_warn_unrepresentable`
below makes that drop loud instead, the same way an unmapped
STAT_BY_MODIFIER_ID id raises rather than vanishes.
"""

from __future__ import annotations

import logging

from pipeline.normalize.gear import RESISTANCE_KEYS, STAT_BY_MODIFIER_ID
from pipeline.simdb.equip import SpellBonus, spell_bonus
from pipeline.simdb.ratings import convert_rating_stats
from pipeline.simdb.statmap import stat_array
from pipeline.simproto import pb

logger = logging.getLogger(__name__)

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


def _warn_unrepresentable(enchant_id: str, bonus: SpellBonus) -> None:
    """Log what `pb.SimEnchant` has no field to carry, naming the enchant and
    the skill or bonus, rather than letting it vanish inside `bonus.stats`.
    """
    for skill, amount in bonus.weapon_skills.items():
        logger.warning(
            "enchant %s grants %s %s through its equip spell, which SimEnchant "
            "has no field for; dropping it",
            enchant_id,
            amount,
            skill,
        )
    if bonus.bonus_physical_damage:
        logger.warning(
            "enchant %s grants %s bonus physical damage through its equip spell, "
            "which SimEnchant has no field for; dropping it",
            enchant_id,
            bonus.bonus_physical_damage,
        )


def build_sim_enchants(
    enchant_rows: list[dict[str, str]],
    effects_by_spell: dict[int, list[dict[str, str]]],
    rating_factors: dict[str, float],
) -> list[pb.SimEnchant]:
    """`rating_factors` is `ratings.load_rating_factors`'s output. A direct
    `EFFECT_STAT` slot states a combat-rating amount for hit, crit, dodge,
    parry, block and defense (see `pipeline/simdb/ratings.py`); an equip
    spell's aura (`spell_bonus`, below) already states a flat percentage and
    is never converted.
    """
    enchants: list[pb.SimEnchant] = []
    for row in sorted(enchant_rows, key=lambda r: int(r["ID"])):
        pairs: list[tuple[str, float]] = []
        equip_spell_ids: list[int] = []
        for slot in EFFECT_SLOTS:
            effect = int(row[f"Effect_{slot}"])
            amount = float(int(row[f"EffectPointsMin_{slot}"]))
            arg = int(row[f"EffectArg_{slot}"])
            if effect == EFFECT_EQUIP_SPELL and arg:
                equip_spell_ids.append(arg)
            elif effect == EFFECT_STAT and amount:
                if arg not in STAT_BY_MODIFIER_ID:
                    raise EnchantDataError(
                        f"enchant {row['ID']} grants unknown stat modifier id {arg}; "
                        f"add it to STAT_BY_MODIFIER_ID in pipeline/normalize/gear.py"
                    )
                key = STAT_BY_MODIFIER_ID[arg]
                if key is not None:
                    pairs.append((key, convert_rating_stats({key: amount}, rating_factors)[key]))
            elif effect == EFFECT_RESISTANCE and amount:
                key = ENCHANT_RESISTANCE_BY_INDEX.get(arg)
                if key:
                    pairs.append((key, amount))
        if equip_spell_ids:
            bonus = spell_bonus(equip_spell_ids, effects_by_spell)
            pairs.extend(bonus.stats.items())
            _warn_unrepresentable(row["ID"], bonus)
        enchants.append(pb.SimEnchant(effect_id=int(row["ID"]), stats=stat_array(pairs)))
    return enchants
