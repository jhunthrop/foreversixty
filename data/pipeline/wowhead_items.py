"""Wowhead's Forever gear planner as an item source, for the ids the client lacks.

The site's items come from wago.tools' export of the client's ``ItemSparse``. Blizzard
delivers much of Forever's itemization as server hotfixes, which live in the client's
hotfix cache and never reach the shipped DB2 files wago exports: on 2026-09-27, 4,712 of
wowhead's 11,239 Forever items (4,687 of them equippable, uncommon or better) had no
``ItemSparse`` row in the newest export, 1.60.1.70009. Seven of one hunter's ten equipped
items were among them. Wowhead reads the live game and carries them all, with stat values
that match what ``normalize/item_curves.py`` resolves from the client's curves one for one
on the 6,527 items both sources have (spec
``docs/superpowers/specs/2026-09-27-wowhead-item-supplement-design.md``).

This module owns that source. The client stays the truth for every id it has; the
``supplement`` is exactly the ids it lacks, passed through the planner's own gates.
"""

from __future__ import annotations

import hashlib
import json
import logging
import re
from collections import Counter
from collections.abc import Iterable
from datetime import UTC, datetime
from pathlib import Path

import httpx
from pydantic import BaseModel

from pipeline.models import GearItem, Item
from pipeline.normalize.gear import (
    MAX_PLAYER_LEVEL,
    PLANNER_QUALITIES,
    SLOT_BY_INVENTORY_TYPE,
    is_junk_name,
    is_weapon_row,
    resolve_required_level,
)
from pipeline.proficiency import ARMOR, can_equip
from pipeline.wago import USER_AGENT

logger = logging.getLogger(__name__)

WOWHEAD_URL = "https://nether.wowhead.com/forever/data/gear-planner"
WOWHEAD_PARAMS = {"dv": "100"}
RAW_FILE = "wowhead-gear-planner.js"
META_FILE = "wowhead-gear-planner.meta.json"
ITEM_PAGE_DATA = "wow.gearPlanner.classicplus.item"
ITEM_SET_PAGE_DATA = "wow.gearPlanner.classicplus.itemSet"

#: Wowhead's ``versionNum`` for the 1.60 client. A payload stating another version is
#: not Forever data (the same URL served Classic Era items until the beta opened).
FOREVER_VERSION = 16001

#: Two-handed melee: the one inventory type that takes the main hand and leaves no
#: off-hand free (``normalize/gear.py``'s TWO_HAND_INVENTORY_TYPES).
TWO_HAND_INVENTORY_TYPE = 17

#: Wowhead stat key -> the planner's stat key, or None for a stat the planner does not
#: track. Calibrated against the client's own resolved values on shared items; every
#: mapped key agreed one for one. Keys absent here are wowhead's own bookkeeping
#: (``slotbak``, ``sellprice``, ``appearances``...) and are ignored without a count.
STAT_KEYS: dict[str, str | None] = {
    "sta": "stamina",
    "agi": "agility",
    "str": "strength",
    "int": "intellect",
    "spi": "spirit",
    "splpwr": "spell_power",
    "spldmg": "spell_power",
    "splheal": "healing",
    "atkpwr": "attack_power",
    "rgdatkpwr": None,  # ranged attack power: STAT_BY_MODIFIER_ID maps 39 to None too
    "feratkpwr": None,
    "manargn": "mp5",
    "healthrgn": None,
    "critstrkrtng": "crit",
    "splcritstrkpct": "crit",
    "hitrtng": "hit",
    "hastertng": None,
    "exprtng": None,
    "armorpenrtng": None,
    "splpen": None,
    "defrtng": "defense",
    "dodgertng": "dodge",
    "parryrtng": "parry",
    "blockrtng": "block",
    "resipct": None,
    "firres": "fire_res",
    "frores": "frost_res",
    "natres": "nature_res",
    "arcres": "arcane_res",
    "shares": "shadow_res",
    "holres": None,
}

#: Wowhead files rings, necks, trinkets, held items, cloaks, tabards and shirts under
#: its own negative armour subclasses; the client files cloaks as cloth (1) and the rest
#: as Miscellaneous (0), which is what ``proficiency.can_equip`` expects.
ARMOR_SUBCLASS_BY_WOWHEAD: dict[int, int] = {
    -2: 0,
    -3: 0,
    -4: 0,
    -5: 0,
    -6: 1,
    -7: 0,
    -8: 0,
}


class WowheadItem(BaseModel):
    id: int
    name: str
    quality: int
    item_level: int
    required_level: int
    class_id: int
    subclass_id: int
    inventory_type: int
    #: Lowercase icon name without extension, the site's own icon vocabulary.
    icon: str
    #: Wowhead's class bitmask, bit ``ChrClasses.ID - 1``; None when any class may equip.
    class_mask: int | None
    #: Wowhead's own stat keys and values, as sent. ``planner_stats`` maps them.
    stats: dict[str, float]
    set_id: int | None
    unique: bool
    armor: int
    damage_min: int
    damage_max: int
    speed: float
    dps: float


class WowheadPayloadError(ValueError):
    """The payload is not the Forever item list this module expects."""


_TRAILING_COMMA = re.compile(r",(\s*[}\]])")


def parse_page_data(text: str, key: str) -> dict:
    """The object of ``WH.setPageData("<key>", {...})``, tolerant of trailing commas."""
    marker = f'WH.setPageData("{key}"'
    at = text.find(marker)
    if at < 0:
        raise WowheadPayloadError(f"payload has no setPageData for {key!r}")
    start = text.index("{", at)
    body = _TRAILING_COMMA.sub(r"\1", text[start:])
    value, _ = json.JSONDecoder().raw_decode(body)
    if not isinstance(value, dict):
        raise WowheadPayloadError(f"setPageData for {key!r} is not an object")
    return value


def _client_subclass(class_id: int, subclass_id: int) -> int:
    if class_id == ARMOR and subclass_id < 0:
        return ARMOR_SUBCLASS_BY_WOWHEAD.get(subclass_id, 0)
    return subclass_id


def _item_from_record(record: dict) -> WowheadItem:
    stats = dict(record.get("stats") or {})
    numeric = {k: float(v) for k, v in stats.items() if isinstance(v, int | float)}
    class_id = int(record["class"])
    armor = int(numeric.get("armor", 0)) + int(numeric.get("armorbonus", 0))
    return WowheadItem(
        id=int(record["id"]),
        name=str(record["name"]),
        quality=int(record["quality"]),
        item_level=int(record.get("itemLevel") or 0),
        required_level=int(record.get("requiredLevel") or numeric.get("reqlevel", 0)),
        class_id=class_id,
        subclass_id=_client_subclass(class_id, int(record["subclass"])),
        inventory_type=int(record.get("inventoryType") or 0),
        icon=str(record.get("icon") or "").lower(),
        class_mask=int(record["classMask"]) if record.get("classMask") else None,
        stats=numeric,
        set_id=int(numeric["itemset"]) if numeric.get("itemset") else None,
        unique=numeric.get("maxcount") == 1,
        armor=armor,
        damage_min=int(numeric.get("dmgmin1", 0)),
        damage_max=int(numeric.get("dmgmax1", 0)),
        speed=float(numeric.get("speed", 0.0)),
        dps=round(float(numeric.get("dps", 0.0)), 2),
    )


def load_items(path: Path) -> list[WowheadItem]:
    """Every item in a saved payload, refusing a payload that is not Forever's."""
    records = parse_page_data(path.read_text(encoding="utf-8"), ITEM_PAGE_DATA)
    versions = Counter(int(r.get("versionNum") or 0) for r in records.values())
    if versions and set(versions) != {FOREVER_VERSION}:
        raise WowheadPayloadError(
            f"payload versionNum {dict(versions)} is not Forever's {FOREVER_VERSION}"
        )
    return sorted((_item_from_record(r) for r in records.values()), key=lambda i: i.id)


def is_planner_gear(item: WowheadItem) -> bool:
    """The gates ``normalize/gear.py``'s build_class_items applies to a client row."""
    return (
        item.inventory_type in SLOT_BY_INVENTORY_TYPE
        and item.quality in PLANNER_QUALITIES
        and item.required_level <= MAX_PLAYER_LEVEL
        and not is_junk_name(item.name)
    )


def supplement(items: Iterable[WowheadItem], client_ids: set[int]) -> list[WowheadItem]:
    """The planner gear wowhead has and the client's ItemSparse does not."""
    return [item for item in items if item.id not in client_ids and is_planner_gear(item)]


def class_allowed(item: WowheadItem, class_id: int) -> bool:
    """Whether ``ChrClasses.ID`` ``class_id`` may equip the item: wowhead's mask and the
    proficiency table, the same two tests a client row gets."""
    if item.class_mask is not None and not item.class_mask & (1 << (class_id - 1)):
        return False
    return can_equip(class_id, item.class_id, item.subclass_id)


def planner_stats(item: WowheadItem, untracked: Counter[str] | None = None) -> dict[str, int]:
    """The item's stats in the planner's vocabulary. A wowhead key the planner does not
    track is counted in ``untracked`` (when given) and dropped, never raised on."""
    out: dict[str, int] = {}
    for key, value in item.stats.items():
        if key not in STAT_KEYS:
            continue
        planner_key = STAT_KEYS[key]
        if planner_key is None:
            if untracked is not None:
                untracked[key] += 1
            continue
        out[planner_key] = out.get(planner_key, 0) + int(round(value))
    return out


def to_gear_item(item: WowheadItem, untracked: Counter[str] | None = None) -> GearItem:
    """The one gate `normalize/gear.py`'s `build_class_items` does not need a
    second time here: wowhead itself states a nonzero "speed" stat for real
    non-weapon items sitting at InventoryType 23 HOLDABLE (Antipodean Rod,
    2879, and Orb of Mistmantle, 13031, are both real class-4 off-hand items
    wowhead's own scrape gives `speed: 1.6` with no damage at all) -- the
    wowhead-side twin of the client-row defect `is_weapon_row` exists to
    fix, so it gates weapon fields here too rather than trusting wowhead's
    numbers at face value the way every other field on this row is.
    """
    is_weapon = is_weapon_row(item.class_id, item.inventory_type)
    # This item has no client row at all (it exists in items/<class>.json
    # only because the supplement added it), so `item.required_level` IS
    # wowhead's own number -- `resolve_required_level`'s `client_level` is 0
    # and its `wowhead_level` is this same value, exactly the precedence
    # `build_class_items` applies for a client row wowhead corroborates.
    required_level, required_level_source = resolve_required_level(
        0, item.item_level, item.required_level
    )
    return GearItem(
        id=item.id,
        name=item.name,
        icon=item.icon,
        slot=SLOT_BY_INVENTORY_TYPE[item.inventory_type],
        quality=item.quality,
        required_level=required_level,
        required_level_source=required_level_source,
        item_level=item.item_level,
        armor=item.armor,
        stats=planner_stats(item, untracked),
        damage_min=item.damage_min if is_weapon else 0,
        damage_max=item.damage_max if is_weapon else 0,
        speed=item.speed if is_weapon else 0.0,
        dps=item.dps if is_weapon else 0.0,
        two_hand=is_weapon and item.inventory_type == TWO_HAND_INVENTORY_TYPE,
        effect_text="",
        stats_source="wowhead",
        set_id=item.set_id,
        unique=item.unique,
    )


def to_item(item: WowheadItem) -> Item:
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


def raw_path(build_dir: Path) -> Path:
    return build_dir / "raw" / RAW_FILE


def fetch_wowhead(
    build: str,
    root: Path = Path("builds"),
    client: httpx.Client | None = None,
) -> Path:
    """Download the payload for a build's raw directory and record when it was taken."""
    own = client is None
    client = client or httpx.Client(headers={"User-Agent": USER_AGENT})
    try:
        response = client.get(WOWHEAD_URL, params=WOWHEAD_PARAMS, timeout=120)
        response.raise_for_status()
        text = response.text
        # Validate before writing: a Classic Era payload must not replace a Forever one.
        records = parse_page_data(text, ITEM_PAGE_DATA)
        versions = {int(r.get("versionNum") or 0) for r in records.values()}
        if versions != {FOREVER_VERSION}:
            raise WowheadPayloadError(
                f"wowhead served versionNum {sorted(versions)}, not Forever's {FOREVER_VERSION}"
            )
        path = raw_path(root / build)
        path.parent.mkdir(parents=True, exist_ok=True)
        path.write_text(text, encoding="utf-8")
        meta = {
            "url": str(response.url),
            "fetched_at": datetime.now(UTC).isoformat(),
            "sha256": hashlib.sha256(text.encode("utf-8")).hexdigest(),
            "items": len(records),
        }
        (path.parent / META_FILE).write_text(json.dumps(meta, indent=1) + "\n")
        logger.info("wowhead: %d Forever items saved to %s", len(records), path)
        return path
    finally:
        if own:
            client.close()
