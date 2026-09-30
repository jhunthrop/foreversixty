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

from pipeline.icons import PLACEHOLDER_ICON, resolve_icon_name
from pipeline.models import GearItem, Item
from pipeline.normalize.gear import (
    MAX_PLAYER_LEVEL,
    PLANNER_QUALITIES,
    SLOT_BY_INVENTORY_TYPE,
    STAT_BY_MODIFIER_ID,
    TWO_HAND_INVENTORY_TYPES,
    is_junk_name,
    is_weapon_row,
)
from pipeline.proficiency import ARMOR, can_equip
from pipeline.spelltext import SpellText
from pipeline.sqldump import unquote

logger = logging.getLogger(__name__)

FILE_NAME = "item_template.json"

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


class ClassicDbItem(BaseModel):
    """One `item_template` row, narrowed to what `to_gear_item`/`to_item`
    need -- the committed extract's own record shape."""

    id: int
    name: str
    quality: int
    item_level: int
    required_level: int
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
    #: Every non-zero `spellid_<n>` (1-5) paired with classic-db's own
    #: `spell_template.SpellName` for it, as a last-resort effect-text
    #: fallback -- see `effect_text`'s own doc for why the client's own
    #: description is always tried first.
    spells: list[tuple[int, str]]


class ClassicDbPayloadError(ValueError):
    """The committed extract is not the shape this module expects."""


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


def _spells(row: dict[str, str], spell_names: dict[int, str]) -> list[tuple[int, str]]:
    out: list[tuple[int, str]] = []
    for n in range(1, 6):
        spell_id = int(row[f"spellid_{n}"])
        if spell_id:
            out.append((spell_id, spell_names.get(spell_id, "")))
    return out


def _item_from_row(row: dict[str, str], spell_names: dict[int, str]) -> ClassicDbItem:
    return ClassicDbItem(
        id=int(row["entry"]),
        name=unquote(row["name"]) or "",
        quality=int(row["Quality"]),
        item_level=int(row["ItemLevel"]),
        required_level=int(row["RequiredLevel"]),
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


def extract_records(sql_text: str) -> list[ClassicDbItem]:
    """Every equippable, planner-quality `item_template` row in a pinned
    mysqldump's text, as the committed extract's own records.

    Narrowed to `PLANNER_QUALITIES` (uncommon/rare/epic/legendary) here, at
    extraction time, not left for `supplement` to filter at every normalize
    run: a poor/common/artifact row can never pass that gate (every other
    source's candidates are narrowed the same way before they reach the
    planner), so committing it would only inflate the extract for rows that
    can never ship.
    """
    from pipeline.audit.dumpdb import ClassicDbDump

    dump = ClassicDbDump.from_text(sql_text)
    spell_names = dump.spell_names
    return sorted(
        (
            _item_from_row(row, spell_names)
            for row in dump.equippable_item_template_rows
            if int(row["Quality"]) in PLANNER_QUALITIES
        ),
        key=lambda item: item.id,
    )


def raw_path(build_dir: Path) -> Path:
    return build_dir / "raw" / "classicdb" / FILE_NAME


def write_extract(build_dir: Path, records: list[ClassicDbItem], *, source_commit: str) -> Path:
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
                "InventoryType) and PLANNER_QUALITIES -- see pipeline.classicdb_items' own doc.",
            },
        },
        "items": [record.model_dump() for record in records],
    }
    path.write_text(json.dumps(document, indent=1) + "\n", encoding="utf-8")
    return path


def load_extract(build_dir: Path) -> list[ClassicDbItem] | None:
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
    return [ClassicDbItem(**record) for record in items]


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


def effect_text(item: ClassicDbItem, spell_text: SpellText) -> str:
    """The item's use/proc/equip descriptions, client text preferred.

    Every non-zero `spellid_<n>` is tried against the Forever client's own
    `SpellText` first (`spell_text.describe`) -- the client carries virtually
    every Classic-era spell even when it lacks the ITEM that used to grant
    it, so this is not a rare path (all 12 spell ids the lane's own 11
    example items reference resolved this way against build 1.60.1.70009's
    own `Spell.csv`). Only a spell id the client has genuinely never heard of
    falls back to classic-db's own internal `spell_template.SpellName` label
    (`item.spells`' own second element) -- not real tooltip text (compare
    "Increase Spell Dam 29" to the client's resolved "Equip: Increases damage
    and healing done by magical spells and effects by up to 29"), but better
    than an empty effect line and the only thing classic-db itself states.
    """
    parts = [
        spell_text.describe(spell_id) or classic_db_name
        for spell_id, classic_db_name in item.spells
    ]
    return " ".join(part for part in parts if part)


def to_gear_item(
    item: ClassicDbItem,
    spell_text: SpellText,
    fork_icons: dict[int, str],
    wowhead_icons: dict[int, str],
) -> GearItem:
    """classic-db states no icon at all (`item_template` has no such column);
    `icon` falls back through the same fork-db/wowhead chain a client row's
    placeholder icon does (`pipeline.icons.resolve_icon_name`), starting from
    `PLACEHOLDER_ICON` rather than a client-resolved base."""
    is_weapon = is_weapon_row(item.class_id, item.inventory_type)
    speed = round(item.delay / 1000, 2)
    dps = (
        round((item.damage_min + item.damage_max) / 2 / speed, 2)
        if is_weapon and speed > 0
        else 0.0
    )
    icon, _origin = resolve_icon_name(PLACEHOLDER_ICON, item.id, fork_icons, wowhead_icons)
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
        stats=planner_stats(item),
        damage_min=item.damage_min if is_weapon else 0,
        damage_max=item.damage_max if is_weapon else 0,
        speed=speed if is_weapon else 0.0,
        dps=dps,
        two_hand=is_weapon and item.inventory_type in TWO_HAND_INVENTORY_TYPES,
        effect_text=effect_text(item, spell_text),
        stats_source="classic-db",
        client_unconfirmed=True,
        set_id=item.set_id,
        unique=item.unique,
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
    )
