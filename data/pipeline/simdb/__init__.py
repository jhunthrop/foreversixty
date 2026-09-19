"""The engine's SimDatabase, built from the same tables the planner uses.

`simdb.bin` is exactly `proto.SimDatabase`: items, enchants and random
suffixes. It is what the web and the API hand the engine as `Player.Database`
(see the engine's `sim/core/database.go`), and what the engine lane embeds for
its own tests.

Two things the interface contract lists cannot be fields of that message:

* **Item sets.** `SimDatabase` has no set message. A set lives on each item as
  `set_id` + `set_name`, which is what the engine's `sim/common/item_sets/`
  matches on, and that is what `build_sim_items` fills in.
* **Consumables.** The engine models them as enums in a hand-written
  `sim/core/consumes.go`, with no database home at all. So the raw material
  for regenerating that file is written beside the protobuf as
  `simconsumes.json`; `simdb.bin` stays exactly the engine's message.

`random_suffixes` is emitted empty: `ItemRandomSuffix` 404s on build
1.60.1.69893, Forever re-itemises the world anyway, and the contract does not
ask for suffixes.
"""

from __future__ import annotations

import json
import logging
from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.manifest import refresh_manifest
from pipeline.models import ConsumableRecord
from pipeline.normalize import write_json
from pipeline.normalize.gear import MAX_PLAYER_LEVEL, column_value, int_column, is_junk_name
from pipeline.normalize.item_curves import load_item_curves
from pipeline.simdb.enchants import build_sim_enchants
from pipeline.simdb.equip import equip_bonuses, index_spell_effects, item_effect_spells
from pipeline.simdb.items import build_sim_items, simdb_item_rows
from pipeline.simdb.weapons import load_weapon_curves
from pipeline.simproto import pb

logger = logging.getLogger(__name__)

SIMDB = "simdb.bin"
SIMCONSUMES = "simconsumes.json"
ITEM_CLASS_CONSUMABLE = 0


def build_consumables(
    sparse_rows: list[dict[str, str]],
    item_rows: list[dict[str, str]],
    item_effect_rows: list[dict[str, str]],
    item_x_item_effect_rows: list[dict[str, str]],
) -> list[ConsumableRecord]:
    """Consumable items and the spells they cast, for the engine's consumes.go.

    On build 1.60.1.69893 this is 1,579 rows. The engine hand-writes the
    effects; what it cannot hand-write is which item ids and spell ids Forever
    ships. Unlike the equip-bonus path this keeps every trigger type -- a
    potion's spell is an on-use one -- so it calls `item_effect_spells` with
    `trigger_types=None` rather than that helper's on-equip-only default.
    """
    consumable_ids = {
        int_column(row, "ID")
        for row in item_rows
        if int_column(row, "ClassID") == ITEM_CLASS_CONSUMABLE
    }
    spells = item_effect_spells(
        item_effect_rows, item_x_item_effect_rows, consumable_ids, trigger_types=None
    )
    records = []
    for row in sparse_rows:
        item_id = int_column(row, "ID")
        if item_id not in spells:
            continue
        if int_column(row, "RequiredLevel") > MAX_PLAYER_LEVEL:
            continue
        name = column_value(row, "Display_lang")
        if is_junk_name(name):
            continue
        records.append(
            ConsumableRecord(
                id=item_id,
                name=name,
                quality=int_column(row, "OverallQualityID"),
                required_level=int_column(row, "RequiredLevel"),
                spell_ids=sorted(set(spells[item_id])),
            )
        )
    return sorted(records, key=lambda record: record.id)


def _set_names(build_dir: Path) -> dict[int, str]:
    path = build_dir / "sets.json"
    if not path.exists():
        raise SystemExit(f"no {path}; run `python -m pipeline normalize` for this build first")
    return {int(row["id"]): row["name"] for row in json.loads(path.read_text(encoding="utf-8"))}


def _optional(raw: Path, name: str) -> list[dict[str, str]]:
    """A table this build's client may not have (see wago.OPTIONAL_TABLES)."""
    path = raw / f"{name}.csv"
    return read_csv(path) if path.exists() else []


def build_sim_database(build_dir: Path) -> tuple[pb.SimDatabase, list[ConsumableRecord]]:
    raw = build_dir / "raw"
    if not raw.exists():
        raise SystemExit(f"no raw data at {raw}; run `python -m pipeline fetch` first")
    set_names = _set_names(build_dir)
    sparse_rows = read_csv(raw / "ItemSparse.csv")
    item_rows = read_csv(raw / "Item.csv")
    item_effect_rows = read_csv(raw / "ItemEffect.csv")
    link_rows = _optional(raw, "ItemXItemEffect")
    effects_by_spell = index_spell_effects(read_csv(raw / "SpellEffect.csv"))
    curves = load_item_curves(
        read_csv(raw / "ItemArmorTotal.csv"),
        read_csv(raw / "ItemArmorQuality.csv"),
        read_csv(raw / "ItemArmorShield.csv"),
        read_csv(raw / "ArmorLocation.csv"),
        read_csv(raw / "RandPropPoints.csv"),
    )
    weapon_curves = load_weapon_curves(
        *(_optional(raw, f"ItemDamage{name}") for name in
          ("OneHand", "TwoHand", "Ranged", "Wand", "Thrown"))
    )

    pairs = simdb_item_rows(sparse_rows, item_rows)
    kept_ids = {int_column(sparse, "ID") for sparse, _ in pairs}
    equip = equip_bonuses(item_effect_rows, link_rows, effects_by_spell, kept_ids)
    database = pb.SimDatabase(
        items=build_sim_items(pairs, set_names, equip, curves, weapon_curves),
        enchants=build_sim_enchants(read_csv(raw / "SpellItemEnchantment.csv"), effects_by_spell),
    )
    consumables = build_consumables(sparse_rows, item_rows, item_effect_rows, link_rows)
    logger.info(
        "simdb: %d items (%d with an on-equip bonus, %d with weapon damage), "
        "%d enchants, %d consumables",
        len(database.items),
        len(equip),
        sum(1 for item in database.items if item.weapon_speed),
        len(database.enchants),
        len(consumables),
    )
    return database, consumables


def write_sim_database(build: str, root: Path = Path("builds")) -> Path:
    build_dir = root / build
    database, consumables = build_sim_database(build_dir)
    path = build_dir / SIMDB
    path.write_bytes(database.SerializeToString(deterministic=True))
    write_json(consumables, build_dir / SIMCONSUMES)
    refresh_manifest(build_dir)
    logger.info("wrote %s (%d bytes)", path, path.stat().st_size)
    return path
