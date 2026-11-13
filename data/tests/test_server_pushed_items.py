"""Item values only the server's hotfix push states, pinned on the active build.

wago.tools' hotfixed CSV export dropped push 16778086 between the 1.60.1.70790
and 70938 data runs, and the 70938 run committed Fiery War Axe (870) at its
client 2.9 s where the server pushes 3.6 s (issue #27). Nothing failed: the
sim's own weapon-speed test was skipped under LAUNCH_MODE. `pipeline.item_push`
now reads the push from the hotfix page; this pins what it must produce, so a
future change of wago's export or page cannot pass silently. Every value is
checked against Wowhead Forever's tooltip (2026-11-13).
"""

from __future__ import annotations

import json
from functools import cache
from pathlib import Path

import pytest

from pipeline.manifest import newest_build
from pipeline.simproto import pb

#: item id -> (name, swing speed, damage min, damage max, the item's stats by name).
#: 870, 5201, 271215, 286750 and 287063 are rows the 70938 run lost with push
#: 16778086 and its neighbours; 2801 is the push that took Blade of Hanna's
#: stats away (the 70790 run still carried them).
PUSHED_WEAPONS: dict[int, tuple[str, float, int, int, dict[str, int]]] = {
    870: ("Fiery War Axe", 3.6, 116, 175, {"strength": 15}),
    2801: ("Blade of Hanna", 2.0, 96, 145, {}),
    5201: (
        "Emberstone Staff",
        3.0,
        34,
        51,
        {"stamina": 8, "fire_res": 8, "spell_power": 24, "healing": 24},
    ),
    271215: ("Coldspire Staff", 3.6, 43, 66, {"frost_res": 13, "frost_power": 32}),
    286750: ("Wisesight Wand", 1.7, 10, 20, {"intellect": 1}),
    287063: ("Serrated Knife", 2.0, 10, 20, {"agility": 1, "spirit": 1}),
}
#: item id -> (name, quality, armor, the item's stats by name), for pushed armor.
PUSHED_ARMOR: dict[int, tuple[str, int, int, dict[str, int]]] = {
    6480: ("Slick Deviate Leggings", 2, 68, {"agility": 4, "stamina": 4}),
}


def build_dir() -> Path:
    return Path("builds") / newest_build()


@cache
def sim_items() -> dict[int, pb.SimItem]:
    database = pb.SimDatabase()
    database.ParseFromString((build_dir() / "simdb.bin").read_bytes())
    return {item.id: item for item in database.items}


@cache
def published_items() -> dict[int, dict]:
    items: dict[int, dict] = {}
    for path in sorted((build_dir() / "items").glob("*.json")):
        for item in json.loads(path.read_text(encoding="utf-8"))["items"]:
            items[item["id"]] = item
    return items


@pytest.mark.parametrize("item_id", sorted(PUSHED_WEAPONS))
def test_the_sim_swings_a_pushed_weapon_at_the_servers_speed_and_damage(item_id):
    name, speed, low, high, _ = PUSHED_WEAPONS[item_id]
    item = sim_items()[item_id]
    assert item.name == name
    assert (round(item.weapon_speed, 2), item.weapon_damage_min, item.weapon_damage_max) == (
        speed,
        low,
        high,
    ), f"{name} in simdb.bin is not the server's pushed row"


@pytest.mark.parametrize("item_id", sorted(PUSHED_WEAPONS))
def test_the_site_publishes_a_pushed_weapon_as_the_server_states_it(item_id):
    name, speed, low, high, stats = PUSHED_WEAPONS[item_id]
    item = published_items()[item_id]
    assert (item["name"], item["speed"], item["damage_min"], item["damage_max"]) == (
        name,
        speed,
        low,
        high,
    )
    assert item["stats"] == stats


@pytest.mark.parametrize("item_id", sorted(PUSHED_ARMOR))
def test_the_site_publishes_pushed_armor_as_the_server_states_it(item_id):
    name, quality, armor, stats = PUSHED_ARMOR[item_id]
    item = published_items()[item_id]
    assert (item["name"], item["quality"], item["armor"]) == (name, quality, armor)
    assert item["stats"] == stats
