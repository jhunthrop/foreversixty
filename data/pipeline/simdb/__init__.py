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

`SimItem.random_suffix_options` and `.faction_restriction` (contract 10.3)
come from `items.json`'s fork-derived columns instead, which `python -m
pipeline loot` writes -- so `loot` must run before this command. `_fork_columns`
below cannot detect every way that ordering was skipped (nothing here can tell
"loot never ran" from "loot ran and genuinely found nothing" once `items.json`
is the only input), but it does fail fast on the two shapes that are
detectable: the columns being entirely absent (an older schema, or a
hand-built build directory), and `loot`'s own output files sitting in the
build directory while every row reads as unrestricted -- which is what
re-running `normalize` after `loot` looks like, since that overwrites
`items.json` from the model defaults and blanks both columns back out.
"""

from __future__ import annotations

import json
import logging
from collections import Counter
from collections.abc import Mapping
from pathlib import Path

from pipeline import classicdb_items as cdb
from pipeline import wowhead_items as wh
from pipeline.csvio import read_csv
from pipeline.manifest import refresh_manifest
from pipeline.models import ConsumableRecord
from pipeline.normalize import write_json
from pipeline.normalize.gear import MAX_PLAYER_LEVEL, column_value, int_column, is_junk_name
from pipeline.normalize.item_curves import load_item_curves
from pipeline.simdb.enchants import build_sim_enchants
from pipeline.simdb.equip import equip_bonuses, index_spell_effects, item_effect_spells
from pipeline.simdb.items import (
    build_classicdb_sim_items,
    build_sim_items,
    build_wowhead_sim_items,
    simdb_item_rows,
)
from pipeline.simdb.ratings import load_rating_factors
from pipeline.simdb.weapons import load_weapon_curves
from pipeline.simproto import pb

logger = logging.getLogger(__name__)

SIMDB = "simdb.bin"
SIMCONSUMES = "simconsumes.json"
SIMITEMS = "simitems.json"
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


#: `items.json`'s two fork-derived columns (parity contract 10.3/10.8).
FORK_COLUMNS = ("suffixes", "faction_restriction")


def _fork_columns(build_dir: Path) -> dict[int, tuple[list[int], str]]:
    """`items.json`'s two fork-derived columns, for `SimItem`.

    They are read from `items.json` and not from the fork database directly
    so that `simdb` never needs an engine checkout; `python -m pipeline loot`
    is what puts them on `items.json`'s rows, which is why it runs first
    (contract 10.8). This catches the two shapes of "loot did not run" that
    are actually detectable from `items.json` and the build directory alone:

    * The columns are missing entirely -- an older schema, or a hand-built
      build directory that never went through `normalize`.
    * `items.json` is empty -- the same "no restrictions" shape as a build
      that genuinely has none, but on zero rows it cannot be genuine.
    * `loot`'s own output (`loot.json` or `suffixes.json`) sits in the build
      directory, yet every row's `suffixes` and `faction_restriction` reads
      empty -- which is what running `normalize` again *after* `loot` looks
      like, since that overwrites `items.json` from the model defaults and
      blanks both columns back out even though loot already ran once.

    An individual item with `[]` and `""` is legitimate -- most items have
    no random suffix and no faction restriction. What these checks refuse is
    the whole build looking un-looted, not any one row looking unrestricted.
    """
    path = build_dir / "items.json"
    if not path.exists():
        raise SystemExit(f"no {path}; run `python -m pipeline normalize` for this build first")
    rows = json.loads(path.read_text(encoding="utf-8"))
    if not rows:
        raise SystemExit(
            f"{path} has no items; run `python -m pipeline normalize` for this build first"
        )
    missing = sorted(set(FORK_COLUMNS) - set(rows[0]))
    if missing:
        raise SystemExit(
            f"{path} is missing {missing}; run `python -m pipeline loot` for this build first"
        )
    columns = {
        int(row["id"]): (row.get("suffixes", []), row.get("faction_restriction", ""))
        for row in rows
    }
    # Local import: pipeline.loot imports pipeline.simdb.statmap, so importing
    # it at module scope here would be circular. By the time this runs,
    # pipeline.simdb has already finished importing, so the cycle resolves.
    from pipeline.loot import LOOT, SUFFIXES

    loot_ran = any((build_dir / name).exists() for name in (LOOT, SUFFIXES))
    all_unrestricted = all(not suffixes and not faction for suffixes, faction in columns.values())
    if loot_ran and all_unrestricted:
        raise SystemExit(
            f"{path} has no suffixes or faction restrictions even though {LOOT} or "
            f"{SUFFIXES} exists in {build_dir}; `normalize` likely ran again after "
            "`loot` and blanked both columns -- re-run `python -m pipeline loot` for "
            "this build"
        )
    return columns


def _optional(raw: Path, name: str) -> list[dict[str, str]]:
    """A table this build's client may not have (see wago.OPTIONAL_TABLES)."""
    path = raw / f"{name}.csv"
    return read_csv(path) if path.exists() else []


def build_sim_database(
    build_dir: Path,
) -> tuple[pb.SimDatabase, list[ConsumableRecord], dict[int, str]]:
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
    rating_factors = load_rating_factors(build_dir)
    fork_columns = _fork_columns(build_dir)
    items = build_sim_items(
        pairs, set_names, equip, curves, weapon_curves, rating_factors, fork_columns
    )
    sources: dict[int, str] = {item.id: "client" for item in items}

    # Every id ItemSparse states outright, junk or not, gate-passing or not --
    # "the client already has this id" for a supplement's own purposes
    # (`wowhead_items.supplement`'s and `classicdb_items.supplement`'s own
    # docs), not merely the narrower `kept_ids` this function's client path
    # itself keeps.
    client_ids = {int_column(row, "ID") for row in sparse_rows}

    # The wowhead supplement (docs/superpowers/specs/2026-09-27-wowhead-item-
    # supplement-design.md): items the client's own ItemSparse lacks
    # entirely, added on top of the client universe rather than in place of
    # any of it. A build whose raw/ has no payload -- every build before
    # `fetch-wowhead` became a workflow step, and any build fetched without
    # it -- is byte-identical to what this function produced before this
    # branch existed.
    wowhead_path = raw / wh.RAW_FILE
    wowhead_ids: set[int] = set()
    if wowhead_path.exists():
        supplement = wh.supplement(wh.load_items(wowhead_path), client_ids)
        untracked: Counter[str] = Counter()
        wowhead_items = build_wowhead_sim_items(supplement, rating_factors, set_names, untracked)
        wowhead_ids = {item.id for item in wowhead_items}
        sources.update({item_id: "wowhead" for item_id in wowhead_ids})
        items = sorted((*items, *wowhead_items), key=lambda row: row.id)
        logger.info(
            "simdb: wowhead supplement added %d items with no on-equip effect "
            "(untracked stats: %s)",
            len(wowhead_items),
            dict(untracked),
        )

    # The classic-db supplement (simdb-supplement lane, 2026-09-30; see
    # pipeline/classicdb_items.py's own doc): cmangos/classic-db's own 1.12
    # item_template, for the ids neither the client's ItemSparse nor
    # wowhead's own supplement carry at all -- Hand of Justice (11815),
    # Blackhand's Breadth (13965) and the rest. Runs after the wowhead block,
    # over the union of ids either already placed
    # (`classicdb_items.supplement`'s own doc), so it never re-adds an id
    # wowhead already covered. A build with no committed extract
    # (`fetch-classic-sources` never ran for it) is unchanged, the same
    # optional-source contract the wowhead branch above already has.
    classicdb_extract = cdb.load_extract(build_dir)
    if classicdb_extract is not None:
        classicdb_records, classicdb_spells = classicdb_extract
        known_ids = client_ids | wowhead_ids
        classicdb_supplement = cdb.supplement(classicdb_records, known_ids)
        classicdb_untracked: Counter[str] = Counter()
        classicdb_sim_items = build_classicdb_sim_items(
            classicdb_supplement,
            classicdb_spells,
            rating_factors,
            set_names,
            fork_columns,
            classicdb_untracked,
        )
        sources.update({item.id: "classic-db" for item in classicdb_sim_items})
        items = sorted((*items, *classicdb_sim_items), key=lambda row: row.id)
        logger.info(
            "simdb: classic-db supplement added %d items (dropped: %s)",
            len(classicdb_sim_items),
            dict(classicdb_untracked),
        )

    database = pb.SimDatabase(
        items=items,
        enchants=build_sim_enchants(
            read_csv(raw / "SpellItemEnchantment.csv"), effects_by_spell, rating_factors
        ),
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
    return database, consumables, sources


def write_sim_items(
    build: str, database: pb.SimDatabase, sources: Mapping[int, str], build_dir: Path
) -> Path:
    """`simitems.json`: the item ids `simdb_item_rows` kept, for the web lane.

    `sources` is `build_sim_database`'s own id -> `"client"`/`"wowhead"`/
    `"classic-db"` mapping (simdb-supplement lane, 2026-09-30), written
    alongside the unchanged `items` list as `sim_source` so an audit can
    count how many published picks land in each bucket without re-deriving
    it -- purely additive: the web lane's `loadSimItems`/`knownItemIds`
    (`web/src/lib/sim/sim-items.ts`) reads only `build`/`items` and ignores
    a key it does not know, so this never needs a web-lane change to stay
    correct.

    `build_sim_items` emits exactly one `SimItem` per kept `(ItemSparse, Item)`
    pair (see its module docstring) and never reorders them, so
    `database.items`'s ids are already `simdb_item_rows`'s kept set in its
    sorted order -- this does not need the raw tables or the intermediate
    `pairs` list a caller built `database` from, only `database` itself.

    The web's bulk pages (item search, bag, bank, sets, Droptimizer sources)
    load this to filter candidates down to items the embedded engine
    database actually carries: `ItemSparse` and `Item` disagree on which ids
    exist (`items.py`'s module docstring), and a candidate id the planner
    knows but the engine does not makes `simCount` fail outright.
    """
    path = build_dir / SIMITEMS
    payload = {
        "build": build,
        "items": [item.id for item in database.items],
        "sim_source": {
            str(item.id): sources.get(item.id, "client") for item in database.items
        },
    }
    path.write_text(json.dumps(payload, indent=2, ensure_ascii=False) + "\n", encoding="utf-8")
    return path


def write_sim_database(build: str, root: Path = Path("builds")) -> Path:
    build_dir = root / build
    database, consumables, sources = build_sim_database(build_dir)
    path = build_dir / SIMDB
    path.write_bytes(database.SerializeToString(deterministic=True))
    write_json(consumables, build_dir / SIMCONSUMES)
    write_sim_items(build, database, sources, build_dir)
    refresh_manifest(build_dir)
    logger.info("wrote %s (%d bytes)", path, path.stat().st_size)
    return path
