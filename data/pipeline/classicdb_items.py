"""cmangos/classic-db's 1.12 `item_template` as an item source, for the ids
neither the client's `ItemSparse` (shipped or hotfix-merged) nor wowhead's
Forever gear-planner scrape carry at all.

Fifth wow-player sweep, 2026-09-30: the client's shipped `ItemSparse` is only
~60% populated in this build by design, and the beta box's hotfix cache only
ever names an id the client has actually seen. cmangos/classic-db's own
`item_template` -- the same pinned dump `pipeline.classic_sources` already
reads for loot -- has 1,960 uncommon-or-better weapons/armor (Item.ClassID
2/4, `Quality >= 2`, a real equip slot) that are in NEITHER: Hand of Justice
(11815), Blackhand's Breadth (13965), Devilsaur Eye (19991) and eight more,
every one of them a real trinket/ring/weapon whose entire value is an
on-equip or on-use spell effect and not a flat stat -- so leaving them out
does not just shrink the catalogue, it empties a hunter's trinket slots at
level 60.

Same shape as `pipeline.wowhead_items`: the client (then wowhead) stays the
source of truth for every id it has; this module only ever adds the ids
BOTH of those lack, through the same planner gates `build_class_items`
applies to a client row, plus one classic-db-specific one (a GM-only
`AllowableClass` bitmask -- see `is_gm_class_mask`).

Every row this module ever emits sets `stats_source`/`required_level_source`
to `"classic-db"` and `client_unconfirmed` to `True`: it is real 1.12
itemization, not yet confirmed by this Forever client build (`data/README.md`,
"classic-db-only items and client_unconfirmed").

Two-stage pipeline, same split `pipeline.classic_sources`/`pipeline.
wowhead_items` already use:

1. `extract_records` (occasional, network -- `fetch-classic-sources` CLI
   command) reads the full pinned mysqldump via
   `pipeline.audit.dumpdb.ClassicDbDump.equippable_item_template_rows` and
   writes the committed, ~2 MB `raw/classicdb/item_template.json` extract
   (`write_extract`/`load_extract`) -- the one parser `pipeline.audit.
   dumpdb` already owns, reused rather than duplicated.
2. `supplement`/`to_gear_item`/`to_item` (every `normalize` run, offline --
   `pipeline.normalize.classicdb` calls these) turn the committed extract
   into the ids the catalogue is still missing.
"""

from __future__ import annotations

import json
import logging
from collections.abc import Iterable
from datetime import UTC, datetime
from pathlib import Path

from pydantic import BaseModel

from pipeline.icons import PLACEHOLDER_ICON, resolve_icon, resolve_icon_name
from pipeline.models import GearItem, Item
from pipeline.normalize.effects import equip_stat_folds_out_text
from pipeline.normalize.gear import (
    MAX_PLAYER_LEVEL,
    PLANNER_QUALITIES,
    SLOT_BY_INVENTORY_TYPE,
    STAT_BY_MODIFIER_ID,
    TWO_HAND_INVENTORY_TYPES,
    is_junk_name,
    is_weapon_row,
    weapon_type_for,
)
from pipeline.proficiency import ARMOR, can_equip
from pipeline.simdb.equip import (
    DEFENSE_SKILL_LINE,
    MOD_STAT_ALL,
    MOD_STAT_BY_INDEX,
    SCHOOL_ALL_MAGIC,
    SCHOOL_RESISTANCE,
)
from pipeline.spelltext import SpellText
from pipeline.sqldump import unquote

logger = logging.getLogger(__name__)

FILE_NAME = "item_template.json"

#: cmangos' own `ItemSpelltriggerType` (item_template's `spelltrigger_<n>`).
#: Only these three appear on any equippable, planner-quality row in the
#: pinned dump (measured, classicdb-fidelity lane 2026-09-30).
TRIGGER_ON_USE = 0
TRIGGER_ON_EQUIP = 1
TRIGGER_CHANCE_ON_HIT = 2

#: cmangos' own `spell_template.Effect<n>` enum (the effect's TYPE, distinct
#: from `EffectApplyAuraName<n>`, the aura it applies when the type is
#: APPLY_AURA). `SPELL_EFFECT_CREATE_ITEM` (24) is already used by
#: `pipeline.audit.dumpdb`; these two are this module's own.
EFFECT_APPLY_AURA = 6
#: cmangos' own SPELL_EFFECT_ADD_EXTRA_ATTACKS -- verified against Hand of
#: Justice's own nested spell 15601 (`EffectTriggerSpell1` off spell 15600),
#: `Effect1` 19, `EffectBasePoints1` 0 (+1 real value: "gain 1 extra attack"),
#: matching wowhead classic's real text ("chance on melee hit to gain 1
#: extra attack") -- see `_extra_attacks`' own doc.
EFFECT_ADD_EXTRA_ATTACKS = 19

#: `EffectApplyAuraName<n>` ids verified against real Classic Wowhead
#: tooltips for the six items the sixth wow-player sweep named (see
#: `effect_text`'s own doc for the citation on each). classic-db's own 1.12
#: aura numbering is NOT always the Forever client's modern `SpellEffect`
#: numbering -- `pipeline.simdb.equip`'s own aura ids are calibrated
#: against the CLIENT's table and must never be reused here blindly:
#: `AURA_MOD_SPELL_CRIT_CHANCE` below is the clearest case (71 here, 552 on
#: the client, confirmed by Eye of the Beast disagreeing with a same-number
#: guess).
AURA_MOD_STAT = 29
#: Not a stat -- the equip aura a "chance on hit" trinket itself carries
#: (Hand of Justice's spell 15600); the real effect is its own
#: `EffectTriggerSpell1` (see `_extra_attacks`).
AURA_PROC_TRIGGER_SPELL = 42
AURA_MOD_RESISTANCE = 22
#: SkillLine-indexed (misc_value); only `DEFENSE_SKILL_LINE` (95, reused
#: from `pipeline.simdb.equip`, a SkillLine.dbc id and not aura-numbering at
#: all) is a planner stat -- a weapon-skill line has no `GearItem.stats` key
#: either, on a client row or here (verified: Legplates of Might's own
#: "+7 Defense" equip line, alongside its separately-verified `parry`).
AURA_MOD_SKILL = 30
AURA_MOD_DAMAGE_DONE = 13
#: Verified: Legplates of Might (+1% parry), Bloodfang Spaulders (+12
#: dodge rating), and -- by the same 47/49/51 sequence, not independently
#: cited -- block (Breastplate of Might's set-bonus text corroborates block
#: exists on this gear tier but not this exact line).
AURA_MOD_PARRY_PERCENT = 47
AURA_MOD_DODGE_PERCENT = 49
AURA_MOD_BLOCK_PERCENT = 51
AURA_MOD_CRIT_PERCENT = 52
AURA_MOD_HIT_CHANCE = 54
#: Not independently verified against a classic-db item (no spell-hit-only
#: trinket in this build's six named items); kept adjacent to the verified
#: 54 (hit) and unified into the same `hit` key regardless, per the lane
#: brief's own "Forever's unified hit/crit" instruction.
AURA_MOD_SPELL_HIT_CHANCE = 55
AURA_MOD_SPELL_CRIT_CHANCE = 71
AURA_MOD_POWER_REGEN = 85
AURA_MOD_HEALING_DONE = 135
#: Verified: Hand of Justice / Devilsaur Eye's own use effect (both +20/+150
#: Attack Power). `AURA_MOD_RANGED_ATTACK_POWER` mirrors it on the same
#: spells (Blizzard grants both together so melee and ranged classes see the
#: same tooltip number) but is dropped, never summed into `attack_power`,
#: matching `pipeline.normalize.gear.STAT_BY_MODIFIER_ID`'s own id-39 (ranged
#: attack power) precedent.
AURA_MOD_ATTACK_POWER = 99
AURA_MOD_RANGED_ATTACK_POWER = 124

#: `EffectApplyAuraName<n>` -> the planner's stat key, for a "flat number,
#: no school mask" aura -- `AURA_MOD_STAT`/`AURA_MOD_DAMAGE_DONE`/
#: `AURA_MOD_RESISTANCE`/`AURA_MOD_SKILL`/`AURA_MOD_SPELL_CRIT_CHANCE` have
#: their own branches in `_equip_stats` (an index or a school mask to
#: decode first). `crit`/`hit` intentionally collect two aura ids each
#: (melee 52 + spell-school-masked 71; physical 54 + spell 55): Forever's
#: unified hit/crit system means both land on the one key
#: (`pipeline.simdb.statmap`).
SIMPLE_STAT_AURAS: dict[int, str] = {
    AURA_MOD_PARRY_PERCENT: "parry",
    AURA_MOD_DODGE_PERCENT: "dodge",
    AURA_MOD_BLOCK_PERCENT: "block",
    AURA_MOD_CRIT_PERCENT: "crit",
    AURA_MOD_HIT_CHANCE: "hit",
    AURA_MOD_SPELL_HIT_CHANCE: "hit",
    AURA_MOD_POWER_REGEN: "mp5",
    AURA_MOD_HEALING_DONE: "healing",
}

#: Auras seen on an on-equip classic-db spell in this build's 1,498-row
#: supplement (classicdb-fidelity lane measurement, 2026-09-30) that are
#: reviewed and known NOT to be a planner stat, by category -- the same
#: "IGNORED_AURAS" shape `pipeline.simdb.equip` uses for the client's own
#: SpellEffect table, at the same rigor (a reasoned bucket, not each one
#: individually wowhead-cited; 144 IS individually cited: Duskbat Drape's
#: own "reduces damage from falling", not a stat). 124 (ranged attack
#: power, mirrors 99) is handled in its own branch, not this set, since it
#: is deliberately dropped rather than ignored-as-a-category.
IGNORED_STAT_AURAS: frozenset[int] = frozenset(
    {8, 15, 19, 23, 31, 43, 77, 89, 102, 107, 109, 117, 123, 131, 139, 144, 154, 158, 161, 180}
)

#: classic-db's own resistance columns -> the planner's stat key. Holy
#: resistance is dropped, same as `pipeline.normalize.gear.RESISTANCE_KEYS`
#: drops its index 1.
RESISTANCE_COLUMNS: dict[str, str] = {
    "fire_res": "fire_res",
    "nature_res": "nature_res",
    "frost_res": "frost_res",
    "shadow_res": "shadow_res",
    "arcane_res": "arcane_res",
}

#: `AllowableClass` bits real vanilla data ever sets, plus two bits (5 Death
#: Knight, 9 Monk) seen on confirmed real items in the pinned dump despite
#: naming a class Classic Era never shipped -- an apparent exporter/authoring
#: quirk, not evidence of a GM row (Stoneslayer, entry 9418, quality 3, mask
#: 1535 = 1503 | bit5; Dented Buckler, entry 1166, mask 2047 = 1503 | bit5 |
#: bit9). A mask with any bit BEYOND this one -- 31233/31236/31240/31248 in
#: the pinned dump, on rows named "90 Epic Warrior Neck", "Arcane Infused
#: Gem", "Boots of Darkness" and "Truefaith Vestments" -- has no such
#: explanation: no class or monk/DK-style bit above bit 10 exists in any
#: vanilla-era class enumeration, real or slack. `-1` ("any class") is never
#: flagged regardless.
ALLOWED_CLASS_MASK = 0x7FF


def is_gm_class_mask(allowable_class: int) -> bool:
    """True when `AllowableClass` names a bit no real vanilla class (or the
    two harmless slack bits above) ever sets -- the signature of a
    gamemaster/test row `is_junk_name` does not catch by name alone. See
    `ALLOWED_CLASS_MASK`'s own doc for the four rows this was measured
    against."""
    return allowable_class >= 0 and (allowable_class & ~ALLOWED_CLASS_MASK) != 0


class ItemSpellSlot(BaseModel):
    """One `spellid_<n>`/`spelltrigger_<n>` pair off an `item_template` row."""

    spell_id: int
    trigger: int
    classic_db_name: str


class ClassicDbSpellEffect(BaseModel):
    """One `spell_template.Effect<n>` slot (n in 1..3, classic-db's own cap)."""

    effect: int
    aura: int
    base_points: int
    die_sides: int
    misc_value: int
    trigger_spell: int


class ClassicDbSpell(BaseModel):
    """One `spell_template` row, narrowed to what `_equip_stats`/`effect_text`
    need -- the committed extract's own `spells` record shape."""

    id: int
    proc_chance: int
    effects: list[ClassicDbSpellEffect]


class ClassicDbPayloadError(ValueError):
    """The committed extract is not the shape this module expects."""


class EquipEffectError(ClassicDbPayloadError):
    """An on-equip spell effect uses an aura `_equip_stats` will not guess at."""


class ClassicDbItem(BaseModel):
    """One `item_template` row, narrowed to what `to_gear_item`/`to_item`
    need -- the committed extract's own record shape."""

    id: int
    name: str
    quality: int
    item_level: int
    required_level: int
    #: `requiredhonorrank` -- the item's OWN PvP-rank floor (Blizzard's
    #: 1.12 vendor-purchase gate for the honor-rank armor/weapon sets;
    #: 0 for every item that needs none). Data-followups-10 lane,
    #: 2026-09-30: `pipeline.loot.sources.pvp_ranks()` only ever reads
    #: the CLIENT's `ItemSparse.RequiredPVPRank`, which this build's own
    #: hotfix table populates for Forever-new PvP items only -- the
    #: original Classic honor-rank sets (Lady Palanseer's, Captain
    #: Dirgehammer's, ...) are untouched client rows with no hotfix, so
    #: `ItemSparse.csv` names no rank for them at all and they fell
    #: through to their bare `vendor:<npc_id>` source with no gate.
    #: `pipeline.classicdb_items.classic_honor_ranks` is this field's
    #: one reader, merged into the client's own ranks before either
    #: `write_loot_files` or `merge_loot_files` builds `pvp:rank-N`
    #: sources.
    required_honor_rank: int = 0
    class_id: int
    subclass_id: int
    inventory_type: int
    #: `AllowableClass`; `-1` means any.
    allowable_class: int
    #: `AllowableRace`; `-1` means any. Carried for completeness but not
    #: enforced anywhere today -- the client's own `AllowableRace` is not
    #: enforced by this pipeline either (`pipeline/forkdb.py`'s own doc notes
    #: 19,171 of the client's own `ItemSparse` rows carry `-1/-1`).
    allowable_race: int
    armor: int
    #: `stat_type<n>` -> `stat_value<n>`, non-zero pairs only, keyed by the
    #: RAW modifier id (`STAT_BY_MODIFIER_ID` maps it at read time -- the
    #: same Blizzard `ItemModType` enum a client row's own
    #: `StatModifier_bonusStat_<n>` column states, unchanged since Classic).
    raw_stats: dict[int, int]
    resistances: dict[str, int]
    damage_min: int
    damage_max: int
    #: milliseconds, same unit as the client's own `ItemDelay`.
    delay: int
    set_id: int | None
    unique: bool
    #: Every non-zero `spellid_<n>` (1-5), classic-db's own `spelltrigger_<n>`
    #: for it (`TRIGGER_ON_USE`/`TRIGGER_ON_EQUIP`/`TRIGGER_CHANCE_ON_HIT`)
    #: and classic-db's own `spell_template.SpellName` -- the last-resort
    #: effect-text fallback and the name-match check `effect_text`'s own doc
    #: describes.
    spells: list[ItemSpellSlot]


def _stats(row: dict[str, str]) -> dict[int, int]:
    stats: dict[int, int] = {}
    for n in range(1, 11):
        stat_id = int(row[f"stat_type{n}"])
        amount = int(row[f"stat_value{n}"])
        if stat_id == 0 or amount == 0:
            continue
        stats[stat_id] = stats.get(stat_id, 0) + amount
    return stats


def _resistances(row: dict[str, str]) -> dict[str, int]:
    return {key: int(row[column]) for column, key in RESISTANCE_COLUMNS.items() if int(row[column])}


def _spells(row: dict[str, str], spell_names: dict[int, str]) -> list[ItemSpellSlot]:
    out: list[ItemSpellSlot] = []
    for n in range(1, 6):
        spell_id = int(row[f"spellid_{n}"])
        if spell_id:
            out.append(
                ItemSpellSlot(
                    spell_id=spell_id,
                    trigger=int(row[f"spelltrigger_{n}"]),
                    classic_db_name=spell_names.get(spell_id, ""),
                )
            )
    return out


def _item_from_row(row: dict[str, str], spell_names: dict[int, str]) -> ClassicDbItem:
    return ClassicDbItem(
        id=int(row["entry"]),
        name=unquote(row["name"]) or "",
        quality=int(row["Quality"]),
        item_level=int(row["ItemLevel"]),
        required_level=int(row["RequiredLevel"]),
        # `.get(..., "0")`, not a bare index: data-followups-10 lane,
        # 2026-09-30 -- this column is new to `_item_from_row` and a
        # handful of this module's own small, hand-written SQL test
        # fixtures predate it (the real pinned dump always carries it,
        # confirmed against the scratchpad's own classicdb.sql); a
        # fixture that never named it means "no rank requirement",
        # matching the field's own zero-value default on ClassicDbItem.
        required_honor_rank=int(row.get("requiredhonorrank", "0") or 0),
        class_id=int(row["class"]),
        subclass_id=int(row["subclass"]),
        inventory_type=int(row["InventoryType"]),
        allowable_class=int(row["AllowableClass"]),
        allowable_race=int(row["AllowableRace"]),
        armor=int(row["armor"]),
        raw_stats=_stats(row),
        resistances=_resistances(row),
        damage_min=round(float(row["dmg_min1"])),
        damage_max=round(float(row["dmg_max1"])),
        delay=int(row["delay"]),
        set_id=int(row["itemset"]) or None,
        unique=int(row["maxcount"]) == 1,
        spells=_spells(row, spell_names),
    )


def _spell_from_row(row: dict[str, str]) -> ClassicDbSpell:
    effects = []
    for n in (1, 2, 3):
        effect = int(row[f"Effect{n}"])
        if effect == 0:
            continue
        effects.append(
            ClassicDbSpellEffect(
                effect=effect,
                aura=int(row[f"EffectApplyAuraName{n}"]),
                base_points=int(row[f"EffectBasePoints{n}"]),
                die_sides=int(row[f"EffectDieSides{n}"]),
                misc_value=int(row[f"EffectMiscValue{n}"]),
                trigger_spell=int(row[f"EffectTriggerSpell{n}"]),
            )
        )
    return ClassicDbSpell(id=int(row["Id"]), proc_chance=int(row["ProcChance"]), effects=effects)


def extract_records(sql_text: str) -> tuple[list[ClassicDbItem], dict[int, ClassicDbSpell]]:
    """Every equippable, planner-quality `item_template` row in a pinned
    mysqldump's text, as the committed extract's own records, plus every
    `spell_template` row any of those items' `spellid_<n>` slots references
    (classicdb-fidelity lane, 2026-09-30: `effect_text`/`_equip_stats` need
    the referenced spell's own structured effects, not only its name --
    see `effect_text`'s own doc for why).

    Narrowed to `PLANNER_QUALITIES` (uncommon/rare/epic/legendary) here, at
    extraction time, not left for `supplement` to filter at every normalize
    run: a poor/common/artifact row can never pass that gate (every other
    source's candidates are narrowed the same way before they reach the
    planner), so committing it would only inflate the extract for rows that
    can never ship. The referenced-spell set is narrowed the same way: only
    spells a kept item's own `spellid_<n>` names, plus (one level deep) any
    `EffectTriggerSpell` those name -- Hand of Justice's own nested "extra
    attack" grant (`_extra_attacks`'s own doc) needs that second hop.
    """
    from pipeline.audit.dumpdb import ClassicDbDump

    dump = ClassicDbDump.from_text(sql_text)
    spell_names = dump.spell_names
    items = sorted(
        (
            _item_from_row(row, spell_names)
            for row in dump.equippable_item_template_rows
            if int(row["Quality"]) in PLANNER_QUALITIES
        ),
        key=lambda item: item.id,
    )
    all_effects = dump.spell_effects
    wanted: set[int] = set()
    for item in items:
        for slot in item.spells:
            wanted.add(slot.spell_id)
    for spell_id in list(wanted):
        for effect in all_effects.get(spell_id, {}).get("effects", []):
            if effect["trigger_spell"]:
                wanted.add(effect["trigger_spell"])
    spells = {
        spell_id: ClassicDbSpell(
            id=spell_id,
            proc_chance=all_effects[spell_id]["proc_chance"],
            effects=[ClassicDbSpellEffect(**effect) for effect in all_effects[spell_id]["effects"]],
        )
        for spell_id in wanted
        if spell_id in all_effects
    }
    return items, spells


def raw_path(build_dir: Path) -> Path:
    return build_dir / "raw" / "classicdb" / FILE_NAME


def write_extract(
    build_dir: Path,
    records: list[ClassicDbItem],
    spells: dict[int, ClassicDbSpell],
    *,
    source_commit: str,
) -> Path:
    path = raw_path(build_dir)
    path.parent.mkdir(parents=True, exist_ok=True)
    document = {
        "generated_at": datetime.now(UTC).isoformat(),
        "source": {
            "classic-db": {
                "repo": "cmangos/classic-db",
                "commit": source_commit,
                "license": "GPL-3.0",
                "note": "item_template, narrowed to equippable (Item.ClassID 2/4, a real "
                "InventoryType) and PLANNER_QUALITIES, plus every spell_template row "
                "referenced by one of those items' spellid_<n> (directly or through an "
                "EffectTriggerSpell) -- see pipeline.classicdb_items' own doc.",
            },
        },
        "items": [record.model_dump() for record in records],
        "spells": [
            spell.model_dump() for spell in sorted(spells.values(), key=lambda spell: spell.id)
        ],
    }
    path.write_text(json.dumps(document, indent=1) + "\n", encoding="utf-8")
    return path


def load_extract(build_dir: Path) -> tuple[list[ClassicDbItem], dict[int, ClassicDbSpell]] | None:
    """The committed extract for one build, or `None` when
    `fetch-classic-sources` has not written one yet -- the caller then
    leaves every output unchanged, same contract as
    `pipeline.normalize.wowhead.load_supplement`."""
    path = raw_path(build_dir)
    if not path.exists():
        logger.info("classic-db items: no extract at %s; items unchanged", path)
        return None
    document = json.loads(path.read_text(encoding="utf-8"))
    items = document.get("items")
    if not isinstance(items, list):
        raise ClassicDbPayloadError(f"{path} has no 'items' list")
    spell_records = document.get("spells", [])
    if not isinstance(spell_records, list):
        raise ClassicDbPayloadError(f"{path} has a 'spells' entry that is not a list")
    spells = {spell.id: spell for spell in (ClassicDbSpell(**record) for record in spell_records)}
    return [ClassicDbItem(**record) for record in items], spells


def classic_honor_ranks(items: list[ClassicDbItem]) -> dict[int, int]:
    """Item id -> `required_honor_rank`, for every item the committed
    extract names a nonzero one for -- the classic-db-sourced twin of
    `pipeline.loot.sources.pvp_ranks()` (which reads only the client's
    own `ItemSparse.RequiredPVPRank`, populated for Forever-new PvP
    items only). `pipeline.loot.__init__` merges this into that
    function's own dict, client wins on an id both name (none measured
    on build 1.60.1.70009 -- the two tables cover disjoint item sets),
    before either `write_loot_files` or `merge_loot_files` builds
    `pvp:rank-N` sources -- data-followups-10 lane, 2026-09-30 (the
    original Classic honor-rank sets had fallen through to their bare
    `vendor:<npc_id>` source with no rank gate at all, see
    `ClassicDbItem.required_honor_rank`'s own doc)."""
    return {item.id: item.required_honor_rank for item in items if item.required_honor_rank}


#: Vanilla Classic's honor-rank title ladder, title prefix -> required
#: rank. This is fixed Classic game design, not a build-specific fact
#: (unlike this module's neighbours, which all read a build's own
#: client/classic-db export) -- verified here against this build's own
#: committed `raw/classicdb/item_template.json`: every item name
#: starting with one of these titles that ALSO carries a nonzero
#: `required_honor_rank` agrees with the rank below (see
#: test_classicdb_items.py's own check, which walks the committed
#: extract and asserts exactly that).
#:
#: bis-ranker-integrity-16 lane, 2026-09-30 (fourteenth player sweep,
#: melee-ranged.md finding 7): all three hunter specs published
#: rank-16/17 honor armor -- "Field Marshal's Chain Greathelm",
#: "Marshal's Chain Vices", "Warlord's Chain Helm", "General's Chain
#: Vices" among them -- with no rank shown and no cap applied at all.
#: `classic_honor_ranks` above only ever keys by the CLASSIC item's own
#: id (16465 "Field Marshal's Chain Helm", rank 17), but Forever's own
#: re-itemised twin renamed every slot word (Helm -> Greathelm,
#: Spaulders -> Pauldrons, Boots -> Greaves/Sabatons, Legguards ->
#: Legplates, Gloves/Grips -> Vices) under a BRAND NEW id (231562) that
#: shares neither the classic id nor the exact name -- so neither the
#: id-keyed rank dict nor `pipeline.loot.sources.apply_reitemisation`'s
#: own exact-name inheritance ever found it. The title prefix alone is
#: a safe, general match: nothing in this catalogue names an item
#: starting with "Field Marshal's " that is not that same honor-rank
#: reward, whatever Forever renamed the rest of the string to.
HONOR_RANK_BY_TITLE: dict[str, int] = {
    "Sergeant's": 7,
    "Senior Sergeant's": 8,
    "Master Sergeant's": 8,
    "First Sergeant's": 9,
    "Sergeant Major's": 9,
    "Blood Guard's": 11,
    "Legionnaire's": 12,
    "Champion's": 14,
    "Lieutenant Commander's": 14,
    "Marshal's": 16,
    "General's": 16,
    "Field Marshal's": 17,
    "Warlord's": 17,
    "Grand Marshal's": 18,
    "High Warlord's": 18,
}


def honor_rank_for_name(name: str) -> int | None:
    """`HONOR_RANK_BY_TITLE`'s own rank for `name`'s title prefix, or
    `None` when `name` starts with none of them. Checked longest title
    first, so "Field Marshal's" (rank 17) matches before the shorter
    "Marshal's" (rank 16) would otherwise also match "Field Marshal's
    Chain Greathelm"."""
    for title in sorted(HONOR_RANK_BY_TITLE, key=len, reverse=True):
        if name.startswith(f"{title} "):
            return HONOR_RANK_BY_TITLE[title]
    return None


def honor_ranks_by_title_for_untitled_items(
    item_rows: Iterable[dict[str, object]], already_ranked: set[int]
) -> dict[int, int]:
    """`honor_rank_for_name` applied to every one of this build's own
    `item_rows` (items.json, Forever's flat catalogue -- id and name,
    whatever else each row carries) not already in `already_ranked`
    (the id-keyed union `classic_honor_ranks`/`pipeline.loot.sources.
    pvp_ranks` already cover) -- `pipeline.loot.__init__.
    _pvp_ranks_with_classic_db`'s own gap-filler for a Forever
    re-itemised honor reward under a new id and a renamed slot word
    (HONOR_RANK_BY_TITLE's own doc)."""
    out: dict[int, int] = {}
    for row in item_rows:
        item_id = int(row["id"])  # type: ignore[arg-type]
        if item_id in already_ranked:
            continue
        rank = honor_rank_for_name(str(row["name"]))
        if rank is not None:
            out[item_id] = rank
    return out


def is_planner_gear(item: ClassicDbItem) -> bool:
    """The gates `normalize/gear.py`'s `build_class_items` applies to a
    client row, plus the classic-db-specific GM-mask check -- see
    `is_gm_class_mask`."""
    return (
        item.inventory_type in SLOT_BY_INVENTORY_TYPE
        and item.quality in PLANNER_QUALITIES
        and item.required_level <= MAX_PLAYER_LEVEL
        and not is_junk_name(item.name)
        and not is_gm_class_mask(item.allowable_class)
    )


def supplement(items: Iterable[ClassicDbItem], known_ids: set[int]) -> list[ClassicDbItem]:
    """The planner gear classic-db has that neither the client nor wowhead's
    supplement already cover. `known_ids` is the union of both -- the
    client's own ids plus every id `pipeline.wowhead_items.supplement`
    picked, so classic-db never re-adds an id wowhead already placed with
    client-corroborated stats (`pipeline.wowhead_items`'s own doc)."""
    return [item for item in items if item.id not in known_ids and is_planner_gear(item)]


def class_allowed(item: ClassicDbItem, class_id: int) -> bool:
    if item.allowable_class >= 0 and not item.allowable_class & (1 << (class_id - 1)):
        return False
    return can_equip(class_id, item.class_id, item.subclass_id)


def planner_stats(item: ClassicDbItem) -> dict[str, int]:
    """`raw_stats` plus `resistances`, in the planner's own stat vocabulary
    (`STAT_BY_MODIFIER_ID`)."""
    stats: dict[str, int] = dict(item.resistances)
    for stat_id, amount in item.raw_stats.items():
        if stat_id not in STAT_BY_MODIFIER_ID:
            raise ValueError(
                f"classic-db item {item.id} ({item.name}) uses unknown stat modifier id "
                f"{stat_id}; add it to STAT_BY_MODIFIER_ID in pipeline/normalize/gear.py"
            )
        key = STAT_BY_MODIFIER_ID[stat_id]
        if key is None:
            continue
        stats[key] = stats.get(key, 0) + amount
    return stats


def _mask_bits(mask: int) -> list[int]:
    """Same convention as `pipeline.simdb.equip`'s own private helper --
    duplicated (3 lines) rather than imported: that name is private to its
    module, and school-bitmask decoding is schema-agnostic enough that a
    tiny local copy costs less than coupling to another module's internals.
    """
    return [1 << shift for shift in range(8) if mask & (1 << shift)]


def _equip_stats(spell: ClassicDbSpell, item_id: int, item_name: str) -> dict[str, int]:
    """One on-equip spell's `APPLY_AURA` effects, in the planner's stat
    vocabulary -- the classic-db-side twin of `pipeline.normalize.gear.
    _apply_stat`/`pipeline.simdb.equip.spell_bonus`. Raises `EquipEffectError`
    for an aura this build has never seen and verified (a classic-db row and
    a client row must agree on a shared id's stats, so a silently-dropped
    aura here is exactly the wrong-scoring defect this lane exists to
    close); an aura reviewed and known NOT to be a stat (`IGNORED_STAT_AURAS`,
    or a school/index this aura tracks that just isn't one of the planner's)
    is dropped, not raised.
    """
    stats: dict[str, int] = {}
    for effect in spell.effects:
        if effect.effect != EFFECT_APPLY_AURA:
            continue
        aura = effect.aura
        amount = effect.base_points + effect.die_sides
        misc = effect.misc_value
        if aura == AURA_MOD_STAT:
            if misc == MOD_STAT_ALL:
                for key in MOD_STAT_BY_INDEX.values():
                    stats[key] = stats.get(key, 0) + amount
            elif misc in MOD_STAT_BY_INDEX:
                key = MOD_STAT_BY_INDEX[misc]
                stats[key] = stats.get(key, 0) + amount
            else:
                raise EquipEffectError(
                    f"item {item_id} ({item_name}) has an on-equip MOD_STAT aura naming "
                    f"unknown stat index {misc}"
                )
        elif aura == AURA_MOD_DAMAGE_DONE:
            if misc == SCHOOL_ALL_MAGIC:
                stats["spell_power"] = stats.get("spell_power", 0) + amount
            # A single-school or physical damage-done bonus has no planner
            # stat key either on a client row (pipeline.simdb.equip's own
            # per-school SCHOOL_POWER split is sim-only); dropped, not raised.
        elif aura == AURA_MOD_RESISTANCE:
            for bit in _mask_bits(misc):
                key = SCHOOL_RESISTANCE.get(bit)
                if key:
                    stats[key] = stats.get(key, 0) + amount
        elif aura == AURA_MOD_SKILL:
            if misc == DEFENSE_SKILL_LINE:
                stats["defense"] = stats.get("defense", 0) + amount
            # A weapon-skill line has no planner stat key either on a client
            # row (only the sim's separate weapon_skills tracks it); dropped.
        elif aura == AURA_MOD_SPELL_CRIT_CHANCE:
            if misc == SCHOOL_ALL_MAGIC:
                stats["crit"] = stats.get("crit", 0) + amount
            # A per-school spell-crit bonus is not verified for classic-db's
            # own aura numbering; dropped rather than guessed.
        elif aura == AURA_MOD_ATTACK_POWER:
            stats["attack_power"] = stats.get("attack_power", 0) + amount
        elif aura in SIMPLE_STAT_AURAS:
            key = SIMPLE_STAT_AURAS[aura]
            stats[key] = stats.get(key, 0) + amount
        elif aura in (AURA_MOD_RANGED_ATTACK_POWER, AURA_PROC_TRIGGER_SPELL):
            continue
        elif aura in IGNORED_STAT_AURAS:
            continue
        else:
            raise EquipEffectError(
                f"item {item_id} ({item_name}) has an on-equip spell using aura {aura}, "
                f"which pipeline/classicdb_items.py does not classify; add it to "
                f"SIMPLE_STAT_AURAS or IGNORED_STAT_AURAS once verified against wowhead"
            )
    return stats


def equip_stats(item: ClassicDbItem, spells: dict[int, ClassicDbSpell]) -> dict[str, int]:
    """`_equip_stats`, summed over every `TRIGGER_ON_EQUIP` spell this item
    carries -- a `TRIGGER_ON_USE`/`TRIGGER_CHANCE_ON_HIT` spell's aura is a
    temporary effect, not an equipped stat (Destiny's own +200 Strength
    proc, verified against wowhead classic, is never a flat +200 Strength
    item -- see `effect_text`'s own doc), so only `TRIGGER_ON_EQUIP` feeds
    the structured `stats` a ranker scores.
    """
    stats: dict[str, int] = {}
    for slot in item.spells:
        if slot.trigger != TRIGGER_ON_EQUIP:
            continue
        spell = spells.get(slot.spell_id)
        if spell is None:
            continue
        for key, amount in _equip_stats(spell, item.id, item.name).items():
            stats[key] = stats.get(key, 0) + amount
    return stats


#: `_STAT_PHRASE`'s values are a plain, honest label -- not a transcription
#: of Wowhead's own copy-editing (this lane verified the NUMBERS against
#: Wowhead classic for the six items its brief names, not the exact prose;
#: see this module's own report for that distinction).
_STAT_PHRASE: dict[str, str] = {
    "strength": "Strength",
    "agility": "Agility",
    "stamina": "Stamina",
    "intellect": "Intellect",
    "spirit": "Spirit",
    "attack_power": "Attack Power",
    "spell_power": "Spell Damage",
    "healing": "Healing",
    "mp5": "Mana per 5 sec",
    "defense": "Defense",
    "parry": "Parry Rating",
    "dodge": "Dodge Rating",
    "block": "Block Rating",
    "fire_res": "Fire Resistance",
    "nature_res": "Nature Resistance",
    "frost_res": "Frost Resistance",
    "shadow_res": "Shadow Resistance",
    "arcane_res": "Arcane Resistance",
    "crit": "Critical Strike",
    "hit": "Hit",
}
_PERCENT_STATS = frozenset({"crit", "hit"})

_TRIGGER_PREFIX: dict[int, str] = {
    TRIGGER_ON_EQUIP: "Equip: ",
    TRIGGER_ON_USE: "Use: ",
    TRIGGER_CHANCE_ON_HIT: "Chance on hit: ",
}


def _extra_attacks(spell: ClassicDbSpell, spells: dict[int, ClassicDbSpell]) -> int | None:
    """The number of extra attacks a `AURA_PROC_TRIGGER_SPELL` effect grants,
    or None when this spell does not grant any (or the nested spell is not
    in the extract). Verified against Hand of Justice (spell 15600's own
    `EffectTriggerSpell1` names 15601, whose `Effect1` is
    `EFFECT_ADD_EXTRA_ATTACKS` with `EffectBasePoints1` 0 -- the real "gain 1
    extra attack" wowhead classic states, `+1` for the same base-points
    convention every other effect here uses)."""
    for effect in spell.effects:
        if effect.effect != EFFECT_APPLY_AURA or effect.aura != AURA_PROC_TRIGGER_SPELL:
            continue
        nested = spells.get(effect.trigger_spell)
        if nested is None:
            continue
        for nested_effect in nested.effects:
            if nested_effect.effect == EFFECT_ADD_EXTRA_ATTACKS:
                return nested_effect.base_points + 1
    return None


def _proc_line(spell: ClassicDbSpell, extra: int) -> str:
    """The "Chance on hit: Gain N extra attack(s)." line for an
    `AURA_PROC_TRIGGER_SPELL` effect -- shared by the two `_render_spell_text`
    branches below so neither has to repeat the phrasing."""
    chance = min(spell.proc_chance, 100)
    plural = "s" if extra != 1 else ""
    return f"Chance on hit ({chance}%): Gain {extra} extra attack{plural}."


def _render_spell_text(
    spell: ClassicDbSpell,
    trigger: int,
    item_id: int,
    item_name: str,
    spells: dict[int, ClassicDbSpell],
) -> str | None:
    """One spell slot's text, built entirely from classic-db's own
    structured fields.

    Returns `None` when this spell has an aura `_equip_stats` does not
    classify -- that spell falls back to the client/bare-name path in
    `effect_text` instead of a half-built sentence. Returns `""` (not
    `None`) when `equip_stat_folds_out_text` says this spell's own stats are
    already fully accounted for by the item's structured `stats` (`.stats`
    fed to `to_gear_item`'s own field) and nothing is left to say -- an
    empty string here means "nothing to add", not "try the client/bare-name
    fallback", because the fallback would just re-describe the same stat
    that already has a home; see that predicate's own doc for the full
    rule. A residual, non-stat effect riding the SAME spell (an extra-attack
    proc such as Hand of Justice's) still gets its own line even when the
    spell's stat portion is folded out.
    """
    try:
        stats = _equip_stats(spell, item_id, item_name)
    except EquipEffectError:
        return None
    extra = _extra_attacks(spell, spells)
    if equip_stat_folds_out_text(trigger == TRIGGER_ON_EQUIP, bool(stats)):
        return _proc_line(spell, extra) if extra is not None else ""
    if not stats and extra is None:
        return None
    pieces = []
    for key, amount in stats.items():
        label = _STAT_PHRASE.get(key, key)
        pieces.append(f"+{amount}% {label}" if key in _PERCENT_STATS else f"+{amount} {label}")
    sentence = _TRIGGER_PREFIX.get(trigger, "") + ", ".join(pieces) + "." if pieces else ""
    if extra is not None:
        proc_line = _proc_line(spell, extra)
        sentence = f"{sentence} {proc_line}".strip() if sentence else proc_line
    return sentence or None


def effect_text(
    item: ClassicDbItem,
    spell_text: SpellText,
    spells: dict[int, ClassicDbSpell],
    client_spell_names: dict[int, str] | None = None,
) -> str:
    """The item's use/proc/equip description.

    Rule (sixth wow-player sweep, 2026-09-30): classic-db's own
    `spell_template` is tried FIRST, structurally -- every `APPLY_AURA`
    effect this module classifies (`_equip_stats`) is rendered straight from
    classic-db's own numbers, with no dependence on the Forever client's
    modern `Spell.csv` at all. 1.12 spell ids are not stable across clients:
    Devilsaur Eye (19991) references classic-db spell 24352 ("Devilsaur
    Fury"), but the CLIENT's own spell 24352 is a different ability entirely
    ("Devilsaur Glare", a Root effect) -- the client's own `SpellName.csv`
    disagrees with classic-db's, the collision this rule exists to catch.
    Verified against Wowhead Classic for all six items the sweep named:
    Devilsaur Eye (https://www.wowhead.com/classic/item=19991/devilsaur-eye,
    "Use: Increases attack power by 150 and your chance to hit by 2%"),
    Hand of Justice (https://www.wowhead.com/classic/item=11815/hand-of-justice,
    "2% chance on melee hit to gain 1 extra attack" plus "+20 Attack Power" --
    classic-db's own ProcChance (2) also disagrees with the client's, which
    states a $h/3 divisor that evaluates to 1%; the client's spell NAME
    matches here ("Hand of Justice" both sides) yet its own text is still
    wrong, so this module never trusts client text over a spell it can
    itself fully classify), Briarwood Reed
    (https://www.wowhead.com/classic/item=12930/briarwood-reed, "+29 spell
    damage and healing"), and Destiny
    (https://www.wowhead.com/classic/item=647/destiny, "Chance on hit:
    Increases Strength by 200" -- item-level `spelltrigger` 2, so this
    never reaches `equip_stats`, only this text).

    An on-equip aura that fully classifies AND became a structured stat
    contributes no text of its own -- `equip_stat_folds_out_text` (eleventh
    wow-player sweep, 2026-09-30), shared with the client schema's
    `EffectIndex.text`, so the two paths cannot silently diverge on when a
    flat "Equip: +X Stat" restates a number the item's own `stats` block
    already carries. Blackhand's Breadth (https://www.wowhead.com/classic/
    item=13965/blackhands-breadth, a flat "+2% critical strike with melee
    attacks") and Eye of the Beast (https://www.wowhead.com/classic/
    item=13968/eye-of-the-beast, a flat "+2% critical strike with spells" --
    classic-db's own aura for this is 71, NOT the client's own 552 for the
    same concept) are both this case now: `_equip_stats` still classifies
    their aura and folds it into `stats`, but `effect_text` no longer
    repeats it as prose, since nothing is left unaccounted for.

    A spell this module cannot fully classify (an aura outside
    `_equip_stats`'s table) falls back to the Forever client's own
    `SpellText` (`spell_text.describe`) when the client's spell NAME equals
    classic-db's own (a real match found for most of the 1,498-row
    supplement's less exotic spells, per data/README.md's own note on this
    lane); a name mismatch, or a spell id the client has no row for at all,
    falls back to classic-db's own internal `SpellName` label -- not real
    tooltip text, but the only thing classic-db itself states, and never
    worse than this field was before this lane.
    """
    client_spell_names = client_spell_names or {}
    parts = []
    for slot in item.spells:
        spell = spells.get(slot.spell_id)
        rendered = (
            _render_spell_text(spell, slot.trigger, item.id, item.name, spells)
            if spell
            else None
        )
        if rendered is None:
            client_name = client_spell_names.get(slot.spell_id)
            if client_name is not None and client_name == slot.classic_db_name:
                rendered = spell_text.describe(slot.spell_id) or slot.classic_db_name
            else:
                rendered = slot.classic_db_name
        parts.append(rendered)
    return " ".join(part for part in parts if part)


def _classicdb_client_icon(
    item_id: int, client_icon_file_ids: dict[int, int], icon_names: dict[int, str]
) -> str:
    """The client's own icon for a classic-db-sourced item id, straight off
    `Item.csv`'s `IconFileDataID` through `ManifestInterfaceData` -- the same
    join `pipeline.normalize.gear.build_class_items` already runs for a
    client row, tried here too before this item ever reaches the fork-db/
    wowhead fallback chain.

    player-review sweep 15/16, 2026-09-30: "First Sergeant's Cloak" (16340,
    a real 1.12 PvP reward, `client_unconfirmed: true` because this beta's
    own `ItemSparse` is only ~60% populated -- this module's own doc -- and
    has no row for this id at all) still resolved a placeholder icon even
    though `Item.csv` (a smaller, separate table `normalize_items` also
    reads, populated independently of `ItemSparse`) carries a real,
    nonzero `IconFileDataID` (133759, `inv_misc_cape_07.blp`, confirmed
    against a fresh wago.tools fetch for this exact build). Item ids are a
    permanent registry across Blizzard's whole catalogue -- an id missing
    from `ItemSparse` does not mean `Item.csv` has forgotten it too, only
    that this ONE client row lacks usable equip data, which is exactly why
    stats/required_level still come from classic-db regardless of what
    this function returns.

    `client_icon_file_ids`/`icon_names` default to empty (every caller
    before this existed, and any test that does not care) so a classic-db
    item with no client Item.csv row at all -- the overwhelming majority --
    falls straight through to `PLACEHOLDER_ICON`, exactly as before."""
    file_id = client_icon_file_ids.get(item_id, 0)
    if not file_id:
        return PLACEHOLDER_ICON
    return resolve_icon(file_id, icon_names, f"classic-db item {item_id}")


def to_gear_item(
    item: ClassicDbItem,
    spell_text: SpellText,
    spells: dict[int, ClassicDbSpell],
    fork_icons: dict[int, str],
    wowhead_icons: dict[int, str],
    client_spell_names: dict[int, str] | None = None,
    client_icon_file_ids: dict[int, int] | None = None,
    icon_names: dict[int, str] | None = None,
) -> GearItem:
    """classic-db's own `item_template` states no icon at all (no such
    column); `icon` first tries the client's OWN `Item.csv` table for this
    exact id (`_classicdb_client_icon`'s own doc -- `Item.csv` is populated
    independently of the `ItemSparse` gap that put this item here at all),
    then falls back through the same fork-db/wowhead chain a client row's
    placeholder icon does (`pipeline.icons.resolve_icon_name`)."""
    is_weapon = is_weapon_row(item.class_id, item.inventory_type)
    speed = round(item.delay / 1000, 2)
    dps = (
        round((item.damage_min + item.damage_max) / 2 / speed, 2)
        if is_weapon and speed > 0
        else 0.0
    )
    base_icon = _classicdb_client_icon(
        item.id, client_icon_file_ids or {}, icon_names or {}
    )
    icon, icon_source = resolve_icon_name(base_icon, item.id, fork_icons, wowhead_icons)
    stats = planner_stats(item)
    for key, amount in equip_stats(item, spells).items():
        stats[key] = stats.get(key, 0) + amount
    return GearItem(
        id=item.id,
        name=item.name,
        icon=icon,
        slot=SLOT_BY_INVENTORY_TYPE[item.inventory_type],
        quality=item.quality,
        required_level=item.required_level,
        required_level_source="classic-db",
        item_level=item.item_level,
        armor=item.armor if item.class_id == ARMOR else 0,
        stats=stats,
        damage_min=item.damage_min if is_weapon else 0,
        damage_max=item.damage_max if is_weapon else 0,
        speed=speed if is_weapon else 0.0,
        dps=dps,
        two_hand=is_weapon and item.inventory_type in TWO_HAND_INVENTORY_TYPES,
        effect_text=effect_text(item, spell_text, spells, client_spell_names),
        stats_source="classic-db",
        client_unconfirmed=True,
        set_id=item.set_id,
        unique=item.unique,
        icon_source=icon_source,
        weapon_type=weapon_type_for(item.class_id, item.subclass_id, item.inventory_type),
    )


def to_item(item: ClassicDbItem) -> Item:
    return Item(
        id=item.id,
        name=item.name,
        quality=item.quality,
        item_level=item.item_level,
        required_level=item.required_level,
        class_id=item.class_id,
        subclass_id=item.subclass_id,
        inventory_type=item.inventory_type,
        # loot-parity-2 lane, 2026-09-30: same labels `to_gear_item` above
        # already sets on the per-class row for this same item -- the flat
        # catalogue's own copy, so a consumer reading items.json alone can
        # tell this is real 1.12 itemization with no client counterpart.
        required_level_source="classic-db",
        stats_source="classic-db",
        client_unconfirmed=True,
    )
