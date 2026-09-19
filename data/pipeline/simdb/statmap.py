"""Stat keys to engine `Stat` indexes, resolved by name rather than by number.

The engine's `Stat` enum is renumbering: the interface contract has the engine
lane collapsing MeleeHit+SpellHit into Hit and MeleeCrit+SpellCrit into Crit,
keeping `sim/core/stats` and `proto/common.proto` in sync. A table of integers
here would silently mis-map the day that lands, so every key carries the enum
names it may resolve to, most specific first, and the first one the generated
proto actually has wins.

The keys are the planner's own stat vocabulary -- the values of
`pipeline.normalize.gear.STAT_BY_MODIFIER_ID`, plus the keys
`pipeline.simdb.equip` reads out of an equip spell -- so the repository has one
name for a stat, not one per consumer.
"""

from __future__ import annotations

from collections.abc import Iterable, Mapping

from pipeline.simproto import pb

#: What the two array builders accept. A mapping is what every caller in this
#: package passes; the pair form exists because amounts accumulate and a dict
#: cannot express the same key twice, so nothing else could test the `+=`.
type StatPairs = Mapping[str, float] | Iterable[tuple[str, float]]


class StatMapError(ValueError):
    """A stat key has no engine stat, or the engine dropped one we relied on."""


#: Planner stat key -> candidate `Stat` enum names, most specific first.
PROTO_STAT_ALIASES: dict[str, tuple[str, ...]] = {
    "strength": ("StatStrength",),
    "agility": ("StatAgility",),
    "stamina": ("StatStamina",),
    "intellect": ("StatIntellect",),
    "spirit": ("StatSpirit",),
    "armor": ("StatArmor",),
    "attack_power": ("StatAttackPower",),
    "ranged_attack_power": ("StatRangedAttackPower",),
    "feral_attack_power": ("StatFeralAttackPower",),
    "spell_power": ("StatSpellPower",),
    "spell_damage": ("StatSpellDamage",),
    "healing": ("StatHealingPower",),
    "mp5": ("StatMP5",),
    "arcane_power": ("StatArcanePower",),
    "fire_power": ("StatFirePower",),
    "frost_power": ("StatFrostPower",),
    "holy_power": ("StatHolyPower",),
    "nature_power": ("StatNaturePower",),
    "shadow_power": ("StatShadowPower",),
    # Forever unifies hit and crit; the vanilla-lineage names are the fallback.
    "hit": ("StatHit", "StatMeleeHit"),
    "crit": ("StatCrit", "StatMeleeCrit"),
    "spell_hit": ("StatHit", "StatSpellHit"),
    "spell_crit": ("StatCrit", "StatSpellCrit"),
    "melee_haste": ("StatMeleeHaste",),
    "spell_haste": ("StatSpellHaste",),
    "spell_penetration": ("StatSpellPenetration",),
    "expertise": ("StatExpertise",),
    "armor_penetration": ("StatArmorPenetration",),
    "defense": ("StatDefense",),
    "dodge": ("StatDodge",),
    "parry": ("StatParry",),
    "block": ("StatBlock",),
    "block_value": ("StatBlockValue",),
    "arcane_res": ("StatArcaneResistance",),
    "fire_res": ("StatFireResistance",),
    "frost_res": ("StatFrostResistance",),
    "nature_res": ("StatNatureResistance",),
    "shadow_res": ("StatShadowResistance",),
    #: Never emitted. The one key whose engine stats deliberately do not exist,
    #: so the "the engine dropped a stat" branch below stays covered by a test
    #: rather than by waiting for the engine to drop one.
    "__probe__": ("StatNope",),
}

_STAT_COUNT = len(pb.Stat.keys())
_WEAPON_SKILL_COUNT = len(pb.WeaponSkill.keys())


def _truncate(array: list[float]) -> list[float]:
    """Drop the trailing zeros.

    `stats.FromFloatArray` and `stats.WeaponSkillsFloatArray` both `copy` into a
    fixed-size Go array, so a short slice is read as zero-padded. Over the whole
    item universe that is most of simdb.bin's bytes.
    """
    last = -1
    for index, value in enumerate(array):
        if value:
            last = index
    return array[: last + 1]


def stat_index(key: str) -> int:
    names = PROTO_STAT_ALIASES.get(key)
    if names is None:
        raise StatMapError(
            f"no engine stat for {key!r}; add it to PROTO_STAT_ALIASES in "
            f"pipeline/simdb/statmap.py"
        )
    known = set(pb.Stat.keys())
    for name in names:
        if name in known:
            return pb.Stat.Value(name)
    raise StatMapError(
        f"{key!r} maps to {names}, none of which the engine's Stat enum has; "
        f"the engine renamed a stat -- update PROTO_STAT_ALIASES"
    )


def _pairs(stats: StatPairs) -> Iterable[tuple[str, float]]:
    return stats.items() if isinstance(stats, Mapping) else stats


def stat_array(stats: StatPairs) -> list[float]:
    array = [0.0] * _STAT_COUNT
    for key, amount in _pairs(stats):
        array[stat_index(key)] += float(amount)
    return _truncate(array)


def weapon_skill_index(name: str) -> int:
    if name not in set(pb.WeaponSkill.keys()):
        raise StatMapError(f"{name!r} is not in the engine's WeaponSkill enum")
    return pb.WeaponSkill.Value(name)


def weapon_skill_array(skills: StatPairs) -> list[float]:
    array = [0.0] * _WEAPON_SKILL_COUNT
    for name, amount in _pairs(skills):
        array[weapon_skill_index(name)] += float(amount)
    return _truncate(array)
