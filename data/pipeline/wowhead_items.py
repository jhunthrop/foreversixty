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

The live feed is not a stable source: between 28 and 30 Oct 2026 it lost 292 of the 688
items only it carried, and they fell back to 1.12 stats
(design/reviews/2026-10-30-70707-dps-drop.md). So ``fetch_wowhead`` merges the live
feed into a committed record set, ``raw/wowhead/items.json``, seeded from the build's
own file or the previous build's, and every consumer reads that merged file
(``raw_path``). ``carry_forward`` has the rules.
"""

from __future__ import annotations

import hashlib
import json
import logging
import re
from collections import Counter
from collections.abc import Iterable, Mapping
from dataclasses import dataclass
from datetime import UTC, datetime
from pathlib import Path

import httpx
from pydantic import BaseModel

from pipeline.build_floors import MIN_LIVE_COVERAGE, previous_build_dir
from pipeline.csvio import read_csv
from pipeline.models import GearItem, Item
from pipeline.normalize.gear import (
    MAX_PLAYER_LEVEL,
    PLANNER_QUALITIES,
    SLOT_BY_INVENTORY_TYPE,
    is_junk_name,
    is_weapon_row,
    resolve_required_level,
    weapon_type_for,
)
from pipeline.proficiency import ARMOR, can_equip
from pipeline.wago import USER_AGENT

logger = logging.getLogger(__name__)

WOWHEAD_URL = "https://nether.wowhead.com/forever/data/gear-planner"
WOWHEAD_PARAMS = {"dv": "100"}
#: The live payload as downloaded (6.8 MB, gitignored). ``levels`` reads its base
#: stats; nothing reads its items but ``fetch_wowhead``.
PAYLOAD_FILE = "wowhead-gear-planner.js"
META_FILE = "wowhead-gear-planner.meta.json"
#: The merged, committed item records every consumer reads (``raw_path``).
CARRY_FILE = Path("wowhead") / "items.json"
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


#: The record fields ``_item_from_record`` reads. Everything else is wowhead's
#: bookkeeping and stays out of the committed file.
RECORD_KEYS = frozenset(
    {
        "id",
        "class",
        "subclass",
        "name",
        "quality",
        "icon",
        "versionNum",
        "inventoryType",
        "itemLevel",
        "requiredLevel",
        "classMask",
    }
)
#: Numeric ``stats`` keys that are wowhead's own bookkeeping (prices, display ids,
#: durability, its sort slot), not item stats.
BOOKKEEPING_STATS = frozenset(
    {
        "avgbuyout",
        "buyprice",
        "displayid",
        "dura",
        "itemSquishEraId",
        "sellprice",
        "sheathtype",
        "slotbak",
    }
)


class WowheadPayloadError(ValueError):
    """The payload is not the Forever item list this module expects."""


class WowheadCoverageError(ValueError):
    """The live feed lists too little of the seeded supplement to be trusted."""


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


def _check_versions(records: Iterable[dict]) -> None:
    versions = Counter(int(r.get("versionNum") or 0) for r in records)
    if versions and set(versions) != {FOREVER_VERSION}:
        raise WowheadPayloadError(
            f"payload versionNum {dict(versions)} is not Forever's {FOREVER_VERSION}"
        )


def trim_record(record: dict) -> dict:
    """One item record with only what the pipeline reads: ``RECORD_KEYS`` and the
    numeric stats that are not ``BOOKKEEPING_STATS``."""
    stats = {
        key: value
        for key, value in (record.get("stats") or {}).items()
        if isinstance(value, int | float) and key not in BOOKKEEPING_STATS
    }
    return {**{k: v for k, v in record.items() if k in RECORD_KEYS}, "stats": stats}


def payload_records(text: str) -> dict[int, dict]:
    """A downloaded payload's item records, trimmed and keyed by id, refusing a
    payload that is not Forever's."""
    records = parse_page_data(text, ITEM_PAGE_DATA)
    _check_versions(records.values())
    return {int(r["id"]): trim_record(r) for r in records.values()}


def read_records(path: Path) -> dict[int, dict]:
    """The records of a committed carry file, keyed by id."""
    document = json.loads(path.read_text(encoding="utf-8"))
    records = {int(item_id): record for item_id, record in document["items"].items()}
    _check_versions(records.values())
    return records


def write_records(
    path: Path,
    records: Mapping[int, dict],
    *,
    build: str,
    seeded_from: str | None,
    live: int | None = None,
    carried: Iterable[int] = (),
) -> None:
    """The carry file: a small header, then one record per line, by id, so a fetch
    diffs only on the records that changed."""
    header = {
        "build": build,
        "seeded_from": seeded_from,
        "live": live if live is not None else len(records),
        "carried": sorted(carried),
    }
    lines = [
        f"  {json.dumps(str(item_id))}: "
        f"{json.dumps(records[item_id], sort_keys=True, ensure_ascii=False)}"
        for item_id in sorted(records)
    ]
    head = json.dumps(header, indent=2, ensure_ascii=False)[:-2]
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(f'{head},\n  "items": {{\n' + ",\n".join(lines) + "\n  }\n}\n")


def load_items(path: Path) -> list[WowheadItem]:
    """Every item in a carry file (``.json``) or a downloaded payload (``.js``),
    refusing records that are not Forever's."""
    if path.suffix == ".json":
        records = read_records(path)
    else:
        records = payload_records(path.read_text(encoding="utf-8"))
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
        icon_source="wowhead",
        weapon_type=weapon_type_for(item.class_id, item.subclass_id, item.inventory_type),
    )


def to_item(item: WowheadItem) -> Item:
    # loot-parity-2 lane, 2026-09-30: the SAME resolution `to_gear_item`
    # above already runs for this id's per-class row -- reused here so the
    # flat catalogue's own `required_level`/`required_level_source` always
    # agree with it, never a second, independent guess.
    required_level, required_level_source = resolve_required_level(
        0, item.item_level, item.required_level
    )
    return Item(
        id=item.id,
        name=item.name,
        quality=item.quality,
        item_level=item.item_level,
        required_level=required_level,
        class_id=item.class_id,
        subclass_id=item.subclass_id,
        inventory_type=item.inventory_type,
        required_level_source=required_level_source,
        stats_source="wowhead",
    )


def raw_path(build_dir: Path) -> Path:
    """The merged item records every consumer reads."""
    return build_dir / "raw" / CARRY_FILE


def payload_path(build_dir: Path) -> Path:
    """The live payload as downloaded."""
    return build_dir / "raw" / PAYLOAD_FILE


@dataclass(frozen=True)
class CarryResult:
    records: dict[int, dict]
    #: Seeded ids the live feed no longer lists, kept from the seed.
    carried: tuple[int, ...]
    #: The share of the seeded supplement the live feed still lists; None when the
    #: seed has no supplement to measure.
    coverage: float | None
    #: The seeded supplement's size and how much of it the live feed lists.
    seeded: int
    listed: int


def carry_forward(
    live: Mapping[int, dict],
    seed: Mapping[int, dict],
    *,
    item_ids: set[int],
    sparse_ids: set[int],
) -> CarryResult:
    """The live feed merged over the seed.

    The live feed wins wherever it has the item. A seeded record the feed lacks is
    kept, unless the client no longer has the item (no ``Item.csv`` row) or now has
    its own ``ItemSparse`` row for it, in which case the client is the truth.

    Coverage is measured on the seeded supplement, the planner gear the client still
    lacks: those are the items whose stats depend on this feed alone."""
    carried = tuple(
        sorted(
            item_id
            for item_id in seed
            if item_id not in live and item_id in item_ids and item_id not in sparse_ids
        )
    )
    seeded = {
        item.id
        for item in supplement((_item_from_record(r) for r in seed.values()), sparse_ids)
        if item.id in item_ids
    }
    listed = len(seeded & live.keys())
    return CarryResult(
        records={**{item_id: seed[item_id] for item_id in carried}, **live},
        carried=carried,
        coverage=listed / len(seeded) if seeded else None,
        seeded=len(seeded),
        listed=listed,
    )


def _seed_for(build_dir: Path) -> tuple[dict[int, dict], str | None]:
    """The records a fetch merges over: the build's own carry file, else the newest
    other build's, else none."""
    own = raw_path(build_dir)
    if own.exists():
        return read_records(own), build_dir.name
    previous = previous_build_dir(build_dir, str(Path("raw") / CARRY_FILE))
    if previous is None:
        return {}, None
    return read_records(raw_path(previous)), previous.name


def _client_ids(build_dir: Path) -> tuple[set[int], set[int]]:
    """The ids of the client's ``Item.csv`` and ``ItemSparse.csv``."""
    raw = build_dir / "raw"
    tables = [raw / "Item.csv", raw / "ItemSparse.csv"]
    for table in tables:
        if not table.exists():
            raise SystemExit(
                f"no {table}: carrying the wowhead records forward needs the client's "
                "item tables; run `python -m pipeline fetch` for this build first"
            )
    item_ids, sparse_ids = ({int(row["ID"]) for row in read_csv(t)} for t in tables)
    return item_ids, sparse_ids


def merge_live_feed(
    build_dir: Path, live: Mapping[int, dict], *, allow_shrink: bool = False
) -> Path:
    """Merge a live feed into the build's carry file and write it, refusing a feed
    below ``MIN_LIVE_COVERAGE`` of the seeded supplement unless ``allow_shrink``.
    Nothing is written when it refuses."""
    seed, seeded_from = _seed_for(build_dir)
    if seed:
        item_ids, sparse_ids = _client_ids(build_dir)
        result = carry_forward(live, seed, item_ids=item_ids, sparse_ids=sparse_ids)
    else:
        result = CarryResult(dict(live), (), None, 0, 0)
    if result.coverage is not None and result.coverage < MIN_LIVE_COVERAGE:
        message = (
            f"wowhead's live feed lists {result.listed} of {result.seeded} seeded supplement "
            f"items ({result.coverage:.0%}, need {MIN_LIVE_COVERAGE:.0%}) from {seeded_from}"
        )
        if not allow_shrink:
            raise WowheadCoverageError(
                f"{message}; the feed has lost items, and the seeded records would carry "
                "them -- pass --allow-shrink to accept the feed and carry them"
            )
        logger.warning("%s; --allow-shrink: carrying %d records", message, len(result.carried))
    path = raw_path(build_dir)
    write_records(
        path,
        result.records,
        build=build_dir.name,
        seeded_from=seeded_from,
        live=len(live),
        carried=result.carried,
    )
    logger.info(
        "wowhead: %d live records, %d carried from %s, %d written to %s",
        len(live),
        len(result.carried),
        seeded_from,
        len(result.records),
        path,
    )
    return path


def fetch_wowhead(
    build: str,
    root: Path = Path("builds"),
    client: httpx.Client | None = None,
    *,
    allow_shrink: bool = False,
) -> Path:
    """Download the live payload, merge it into the build's carry file and record when
    it was taken. Returns the carry file."""
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
        build_dir = root / build
        carry = merge_live_feed(build_dir, payload_records(text), allow_shrink=allow_shrink)
        path = payload_path(build_dir)
        path.write_text(text, encoding="utf-8")
        meta = {
            "url": str(response.url),
            "fetched_at": datetime.now(UTC).isoformat(),
            "sha256": hashlib.sha256(text.encode("utf-8")).hexdigest(),
            "items": len(records),
        }
        (path.parent / META_FILE).write_text(json.dumps(meta, indent=1) + "\n")
        logger.info("wowhead: %d Forever items saved to %s", len(records), path)
        return carry
    finally:
        if own:
            client.close()
