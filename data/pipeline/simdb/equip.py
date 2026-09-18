"""What a spell grants its caster, for items and enchants alike.

Classic grants crit, hit, spell power, mp5 and defence through an on-equip
spell rather than through an ItemSparse column, which `data/README.md` records
as a gap in the planner's item data. Measured on build 1.60.1.69893:

* 285 of the 4,986 items simdb keeps carry an on-equip spell, and 18 of those
  grant a stat. Forever's re-itemised world states the rest in
  StatPercentEditor_<n>, which pipeline/normalize/item_curves.py resolves. The
  eighteen -- Mace of Unending Life's +140 attack power, Atiesh's +420, Rune of
  the Guard Captain's +42/+42 -- exist in no column at all.
* 1,995 of the 2,216 SpellItemEnchantment rows name an equip spell
  (`Effect_<n> == 3` puts its id in `EffectArg_<n>`), and 1,334 end up carrying
  stats. The enchant table would be nearly empty without this module, which is
  now the main reason it exists.

Two client-shape facts, both verified:

* The amount is `pipeline.spelltext.effect_amount`, not a local expression.
  Classic Era states "Attack Power 24" as base 23 plus one die side; the 1.60
  client has no EffectDieSides and states 24 outright. One function owns that.
* The 1.60 ItemEffect has no ParentItemID: the item-to-effect link is the
  separate ItemXItemEffect table, which Classic Era 404s on. Both are accepted.

Conventions:

* `EffectMiscValue_0` is a school mask where the aura takes one: 1 physical,
  2 holy, 4 fire, 8 nature, 16 frost, 32 shadow, 64 arcane, 126 all magic.
* Aura 29 takes a stat index: 0 strength, 1 agility, 2 stamina, 3 intellect,
  4 spirit, -1 all five.
* Aura 30 takes a SkillLine id: 95 is Defense, the weapon lines are the
  engine's `weapon_skills`, professions grant nothing.

The kept items' equip spells use 31 distinct aura types on this build. Every
one is in STAT_AURAS or IGNORED_AURAS below; the rest of IGNORED_AURAS is
behaviour the engine hand-writes in `sim/common/item_effects.go` and
`enchant_effects.go`. An aura in neither raises -- a dropped stat is a wrong
sim and a guessed one is worse. A later build with a new aura fails loudly,
which is the new-build workflow the README describes.
"""

from __future__ import annotations

from collections import defaultdict
from collections.abc import Container, Iterable
from dataclasses import dataclass, field

from pipeline.spelltext import effect_amount

TRIGGER_ON_EQUIP = "1"
APPLY_AURA = "6"

AURA_DAMAGE_DONE = 13
AURA_RESISTANCE = 22
AURA_MOD_STAT = 29
AURA_MOD_SKILL = 30
AURA_SPELL_CRIT_SCHOOL = 552

SCHOOL_PHYSICAL = 1
SCHOOL_ALL_MAGIC = 126

#: School mask bit -> school-specific spell power stat.
SCHOOL_POWER: dict[int, str] = {
    2: "holy_power",
    4: "fire_power",
    8: "nature_power",
    16: "frost_power",
    32: "shadow_power",
    64: "arcane_power",
}

#: School mask bit -> resistance stat. Bit 1 is armour and bit 2 is holy
#: resistance; armour is handled in the branch, holy resistance has no Stat.
SCHOOL_RESISTANCE: dict[int, str] = {
    4: "fire_res",
    8: "nature_res",
    16: "frost_res",
    32: "shadow_res",
    64: "arcane_res",
}

#: Aura 29's EffectMiscValue_0. -1 means all five.
MOD_STAT_BY_INDEX: dict[int, str] = {
    0: "strength",
    1: "agility",
    2: "stamina",
    3: "intellect",
    4: "spirit",
}
MOD_STAT_ALL = -1

DEFENSE_SKILL_LINE = 95

#: SkillLine id -> WeaponSkill enum name, or None for a line the sim ignores.
#: Staves, Polearms, Thrown and Feral Combat are included because Forever's
#: re-itemised world may start using them.
WEAPON_SKILL_BY_SKILL_LINE: dict[int, str | None] = {
    43: "WeaponSkillSwords",
    44: "WeaponSkillAxes",
    45: "WeaponSkillBows",
    46: "WeaponSkillGuns",
    54: "WeaponSkillMaces",
    55: "WeaponSkillTwoHandedSwords",
    136: "WeaponSkillStaves",
    160: "WeaponSkillTwoHandedMaces",
    162: "WeaponSkillUnarmed",
    172: "WeaponSkillTwoHandedAxes",
    173: "WeaponSkillDaggers",
    176: "WeaponSkillThrown",
    182: None,  # herbalism
    186: None,  # mining
    226: "WeaponSkillCrossbows",
    228: None,  # wands carry no weapon skill in the engine
    229: "WeaponSkillPolearms",
    356: None,  # fishing
    393: None,  # skinning
    473: "WeaponSkillFeralCombat",
    633: None,  # lockpicking
}

#: EffectAura -> stat key, for the auras whose amount is a flat stat. The four
#: ids that read EffectMiscValue_0 (13, 22, 29, 30) have their own branches.
#: 274 is the 1.60 client's flat "Block Value NN"; 564 is the older tables'.
STAT_AURAS: dict[int, str] = {
    47: "parry",
    49: "dodge",
    51: "block",
    52: "crit",
    54: "hit",
    55: "spell_hit",
    65: "spell_haste",
    85: "mp5",
    99: "attack_power",
    123: "spell_penetration",
    124: "ranged_attack_power",
    135: "healing",
    138: "melee_haste",
    274: "block_value",
    AURA_SPELL_CRIT_SCHOOL: "spell_crit",
    564: "block_value",
}

#: Auras reviewed and known not to be stats: procs, immunities, spell-specific
#: modifiers, percentage modifiers, pet and shapeshift behaviour, and the
#: per-school hit and ranged haste values the engine models as pseudo-stats.
#: Seven were added for build 1.60.1.69893 and each is named after its evidence:
#: 155 water breathing, 272 percentage block ("Shield Specialization"), 319
#: percentage melee speed ("Rapid Fire"), 342 percentage haste ("Gyroscopic
#: Acceleration"), 470 a food effect ("Wastewanderer Rations"), 598 a
#: percentage stat conversion ("Careful Aim"). Every aura this build puts on a
#: kept item's or an enchant's equip spell is in this set or in STAT_AURAS.
IGNORED_AURAS = frozenset(
    {
        3, 4, 8, 10, 14, 15, 17, 19, 23, 31, 33, 34, 35, 42, 43, 56, 57, 58,
        59, 64, 69, 77, 79, 82, 87, 89, 98, 102, 107, 108, 109, 112, 117, 122,
        129, 130, 131, 134, 139, 140, 144, 154, 155, 161, 180, 187, 194, 197,
        213, 226, 234, 262, 272, 275, 290, 319, 328, 332, 342, 395, 470, 561,
        576, 593, 598, 601, 608,
    }
)

_BY_MISC_VALUE = frozenset({AURA_DAMAGE_DONE, AURA_RESISTANCE, AURA_MOD_STAT, AURA_MOD_SKILL})


class EquipEffectError(ValueError):
    """A spell effect this module will not guess at."""


@dataclass(frozen=True)
class SpellBonus:
    """What one or more spells add to the unit that carries them."""

    stats: dict[str, float] = field(default_factory=dict)
    weapon_skills: dict[str, float] = field(default_factory=dict)
    bonus_physical_damage: float = 0.0

    def is_empty(self) -> bool:
        return not self.stats and not self.weapon_skills and not self.bonus_physical_damage


def index_spell_effects(spell_effect_rows: list[dict[str, str]]) -> dict[int, list[dict[str, str]]]:
    """Group SpellEffect rows by spell id, dropping non-default difficulties."""
    grouped: dict[int, list[dict[str, str]]] = defaultdict(list)
    for row in spell_effect_rows:
        if row.get("DifficultyID", "0") not in ("0", ""):
            continue
        grouped[int(row["SpellID"])].append(row)
    return grouped


def item_effect_spells(
    item_effect_rows: list[dict[str, str]],
    item_x_item_effect_rows: list[dict[str, str]],
    item_ids: Container[int],
    trigger_types: Container[str] | None = frozenset({TRIGGER_ON_EQUIP}),
) -> dict[int, list[int]]:
    """Item id -> the spell ids its item effects cast, on either schema.

    Classic Era's ItemEffect names its item in ParentItemID. The 1.60 client
    dropped that column and links through ItemXItemEffect (ID, ItemEffectID,
    ItemID) instead, which Era in turn 404s on. Choosing per row means one
    reader works on both and neither build needs a flag.

    `trigger_types` keeps only rows whose TriggerType is in it; the default,
    on-equip only, is what `equip_bonuses` needs -- only an on-equip spell is
    a stat the wearer carries. `build_consumables` passes `None` to keep
    every trigger type: a potion's spell is on-use (TriggerType 2), and every
    trigger type is consumable behaviour the engine needs to know about.
    """
    spells: dict[int, list[int]] = defaultdict(list)
    by_effect_id: dict[int, dict[str, str]] = {}
    for row in item_effect_rows:
        if "ParentItemID" in row:
            if trigger_types is not None and row.get("TriggerType") not in trigger_types:
                continue
            item_id = int(row["ParentItemID"])
            if item_id in item_ids:
                spells[item_id].append(int(row["SpellID"]))
        else:
            by_effect_id[int(row["ID"])] = row
    for link in item_x_item_effect_rows:
        item_id = int(link["ItemID"])
        if item_id not in item_ids:
            continue
        effect = by_effect_id.get(int(link["ItemEffectID"]))
        if effect is None:
            continue
        if trigger_types is not None and effect.get("TriggerType") not in trigger_types:
            continue
        spells[item_id].append(int(effect["SpellID"]))
    return dict(spells)


def _bits(mask: int) -> list[int]:
    return [1 << shift for shift in range(8) if mask & (1 << shift)]


def _add(target: dict[str, float], key: str, amount: float) -> None:
    target[key] = target.get(key, 0.0) + amount


def spell_bonus(
    spell_ids: Iterable[int],
    effects_by_spell: dict[int, list[dict[str, str]]],
) -> SpellBonus:
    """Everything the given spells grant, summed."""
    stats: dict[str, float] = {}
    skills: dict[str, float] = {}
    physical = 0.0
    for spell_id in spell_ids:
        for row in effects_by_spell.get(spell_id, []):
            if row["Effect"] != APPLY_AURA:
                continue
            aura = int(row["EffectAura"])
            if aura not in STAT_AURAS and aura not in IGNORED_AURAS and aura not in _BY_MISC_VALUE:
                raise EquipEffectError(
                    f"spell {spell_id} applies unclassified aura {aura}; classify it in "
                    f"STAT_AURAS or IGNORED_AURAS in pipeline/simdb/equip.py"
                )
            amount = float(effect_amount(row))
            misc = int(row["EffectMiscValue_0"])
            if aura == AURA_DAMAGE_DONE:
                if misc == SCHOOL_ALL_MAGIC:
                    _add(stats, "spell_power", amount)
                    continue
                for bit in _bits(misc):
                    if bit == SCHOOL_PHYSICAL:
                        physical += amount
                    elif bit in SCHOOL_POWER:
                        _add(stats, SCHOOL_POWER[bit], amount)
            elif aura == AURA_RESISTANCE:
                for bit in _bits(misc):
                    if bit == SCHOOL_PHYSICAL:
                        _add(stats, "armor", amount)
                    elif bit in SCHOOL_RESISTANCE:
                        _add(stats, SCHOOL_RESISTANCE[bit], amount)
            elif aura == AURA_MOD_STAT:
                if misc == MOD_STAT_ALL:
                    for key in MOD_STAT_BY_INDEX.values():
                        _add(stats, key, amount)
                elif misc in MOD_STAT_BY_INDEX:
                    _add(stats, MOD_STAT_BY_INDEX[misc], amount)
            elif aura == AURA_MOD_SKILL:
                if misc == DEFENSE_SKILL_LINE:
                    _add(stats, "defense", amount)
                elif misc in WEAPON_SKILL_BY_SKILL_LINE:
                    name = WEAPON_SKILL_BY_SKILL_LINE[misc]
                    if name is not None:
                        _add(skills, name, amount)
                else:
                    raise EquipEffectError(
                        f"spell {spell_id} modifies unknown skill line {misc}; add it to "
                        f"WEAPON_SKILL_BY_SKILL_LINE in pipeline/simdb/equip.py"
                    )
            elif aura in STAT_AURAS and amount:
                if aura == AURA_SPELL_CRIT_SCHOOL and misc != SCHOOL_ALL_MAGIC:
                    # Per-school spell crit has no Stat; the engine models it as
                    # an item effect.
                    continue
                _add(stats, STAT_AURAS[aura], amount)
    return SpellBonus(
        stats={key: value for key, value in stats.items() if value},
        weapon_skills={key: value for key, value in skills.items() if value},
        bonus_physical_damage=physical,
    )


def equip_bonuses(
    item_effect_rows: list[dict[str, str]],
    item_x_item_effect_rows: list[dict[str, str]],
    effects_by_spell: dict[int, list[dict[str, str]]],
    item_ids: Container[int],
) -> dict[int, SpellBonus]:
    """Item id -> what its on-equip spells grant, for the items the caller kept.

    `item_ids` is not an optimisation: the caller has already dropped the
    gamemaster, test and monster rows, and those carry auras nobody has
    reviewed. Checking them would fail the run over data that is never shipped.
    """
    bonuses: dict[int, SpellBonus] = {}
    for item_id, spell_ids in item_effect_spells(
        item_effect_rows, item_x_item_effect_rows, item_ids
    ).items():
        bonus = spell_bonus(spell_ids, effects_by_spell)
        if not bonus.is_empty():
            bonuses[item_id] = bonus
    return bonuses
