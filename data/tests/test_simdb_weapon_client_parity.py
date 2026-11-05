"""Every client-backed weapon in the committed simdb.bin against the client tables.

A weapon's whole value to the sim is its damage range and its speed. For an
item the client's ItemSparse carries, both come off the 1.60 client: speed is
`ItemDelay`, damage is the ItemDamage* curve by item level and quality, spread
by `DmgVariance`, offset by `QualityModifier`, cut by the caster bit
(`pipeline.normalize.weapon_curves`). The curve formula itself is pinned by
tests/test_simdb_weapons.py against Wowhead Forever tooltips; this test pins
that the committed database still says what the client says, row by row, so a
stale simdb.bin (hotfix pushes after the last data run) or a supplement source
shadowing a client row shows up as a named list of items.

Sim-accuracy round 22 (items lane) found eight weapons out of date this way on
1.60.1.70790: Earthshaker 288900 at its pre-hotfix item level 10 (40 - 61
against the client's 87 - 132, as Wowhead Forever shows it) and seven thrown
weapons whose 70790 hotfix rows now carry a 0.6 spread where the database
(and, so far, Wowhead) has none. Both are hotfix pushes wago.tools recorded
after the 4 Nov 20:43 UTC data run; the next data run clears them.

The raw client tables (builds/<build>/raw/) are never committed, so this skips
unless they have been fetched (`python -m pipeline fetch` plus
`fetch-wago-hotfixes`), as the data run does before it rebuilds simdb.bin.
"""

import json
from functools import cache
from pathlib import Path

import pytest

from pipeline.csvio import read_csv
from pipeline.hotfix_merge import merge_hotfix_table
from pipeline.manifest import newest_build
from pipeline.simdb import SIMITEMS
from pipeline.simdb.weapons import WeaponCurves, load_weapon_curves, weapon_damage
from pipeline.simproto import pb

CURVE_TABLES = ("OneHand", "TwoHand", "Ranged", "Wand", "Thrown")
CLIENT_SOURCE = "client"
RAW_TABLES = ("ItemSparse.csv", "Item.csv", *(f"ItemDamage{name}.csv" for name in CURVE_TABLES))
#: Every 70790 client weapon resolves; a floor, not a count (data contracts are floors).
MIN_CLIENT_WEAPONS = 1500


def build_dir() -> Path:
    return Path("builds") / newest_build()


def raw_dir() -> Path:
    raw = build_dir() / "raw"
    missing = [name for name in RAW_TABLES if not (raw / name).exists()]
    if missing:
        pytest.skip(f"raw client tables have not been fetched for this build: {missing}")
    return raw


@cache
def client_rows() -> tuple[dict[str, dict[str, str]], dict[str, dict[str, str]], WeaponCurves]:
    raw = raw_dir()
    sparse = {row["ID"]: row for row in merge_hotfix_table(raw, "ItemSparse")}
    items = {row["ID"]: row for row in merge_hotfix_table(raw, "Item")}
    curves = load_weapon_curves(*(read_csv(raw / f"ItemDamage{name}.csv") for name in CURVE_TABLES))
    return sparse, items, curves


def client_backed_weapons() -> list[pb.SimItem]:
    database = pb.SimDatabase()
    database.ParseFromString((build_dir() / "simdb.bin").read_bytes())
    sources = json.loads((build_dir() / SIMITEMS).read_text())["sim_source"]
    return [
        item
        for item in database.items
        if item.weapon_speed > 0 and sources.get(str(item.id)) == CLIENT_SOURCE
    ]


Damage = tuple[float, float, float]


def stored_damage(item: pb.SimItem) -> Damage:
    return (item.weapon_damage_min, item.weapon_damage_max, round(item.weapon_speed, 3))


def client_damage(
    sparse_row: dict[str, str], item_row: dict[str, str], curves: WeaponCurves
) -> Damage | None:
    damage = weapon_damage(sparse_row, int(item_row["SubclassID"]), curves)
    if damage is None:
        return None
    return (damage.minimum, damage.maximum, round(damage.speed, 3))


def test_every_client_weapon_carries_the_clients_damage_range_and_speed():
    sparse, items, curves = client_rows()
    weapons = client_backed_weapons()
    assert len(weapons) >= MIN_CLIENT_WEAPONS

    mismatches: list[str] = []
    for weapon in weapons:
        key = str(weapon.id)
        client = (
            client_damage(sparse[key], items[key], curves)
            if key in sparse and key in items
            else None
        )
        if client != stored_damage(weapon):
            mismatches.append(
                f"{weapon.id} {weapon.name}: simdb {stored_damage(weapon)}, client {client}"
            )
    assert not mismatches, "simdb.bin is out of step with the client:\n" + "\n".join(mismatches)
