"""Data contract: every item a spec's BiS file proposes (pick or alternative) is one that
class can equip at that band's level.

One rule decides who can use an item, upstream of every surface that proposes one: the
client's `ItemSparse.AllowableClass` and the class's armour and weapon skills
(`pipeline.normalize.gear.build_class_items` writes items/<class>.json through it, and
`sim/cmd/leveling-bis` sources its candidates from that file and gates armour by level).
The API's Next upgrades, the Discord /char panel and the addon's Data.lua all read the BiS
files, so holding the files to the rule holds every surface to it.

Dungeon Set 1 and 2 pieces carry no class restriction in Forever's client (Beaststalker's
Pants is AllowableClass -1 in ItemSparse, in classic-db's item_template and on Wowhead's
Forever tooltip), so a mail-wearing warrior may be offered one; Tier 1 and up are
class-locked and never cross classes.
"""

from __future__ import annotations

import json
from collections.abc import Iterator
from pathlib import Path

import pytest

from conftest import ACTIVE_BUILD

BUILD_DIR = Path("builds") / ACTIVE_BUILD

#: Item.ClassID values with a skill gate.
ITEM_CLASS_WEAPON = 2
ITEM_CLASS_ARMOR = 4

#: Armour subclasses no skill gates: miscellaneous (rings, necks, trinkets, held in off
#: hand) and the relics (libram, idol, totem), which the class mask alone gates.
UNGATED_ARMOR_SUBCLASSES = frozenset({0, 7, 8, 9})

BEASTSTALKERS_PANTS = 16678  # Dungeon Set 1, hunter-themed, AllowableClass -1
DRAGONSTALKERS_HELM = 16939  # Tier 2, AllowableClass 4 (hunter)
HELM_OF_WRATH = 16963  # Tier 2, AllowableClass 1 (warrior)


def _json(path: Path) -> dict:
    return json.loads(path.read_text())


def _class_items(class_slug: str) -> dict[int, dict]:
    return {item["id"]: item for item in _json(BUILD_DIR / "items" / f"{class_slug}.json")["items"]}


def _skill_levels(class_slug: str) -> dict[tuple[int, int], int]:
    skills = _json(BUILD_DIR / "equip-skills.json")["classes"][class_slug]
    levels = {(ITEM_CLASS_ARMOR, row["subclass"]): row["level"] for row in skills["armor"]}
    levels.update({(ITEM_CLASS_WEAPON, row["subclass"]): row["level"] for row in skills["weapons"]})
    return levels


def _proposals(bis: dict) -> Iterator[tuple[dict, str, int, str]]:
    """Every (band, slot, item id, kind) a BiS file proposes."""
    for band in bis["bands"]:
        for slot in band["slots"]:
            if slot.get("item_id") is not None:
                yield band, slot["slot"], slot["item_id"], "pick"
            for alt in slot.get("alternatives") or []:
                yield band, slot["slot"], alt["item_id"], "alternative"


def _unusable(spec: str, bis: dict) -> list[str]:
    class_slug = spec.split("-")[0]
    items = _class_items(class_slug)
    skills = _skill_levels(class_slug)
    problems: list[str] = []
    for band, slot, item_id, kind in _proposals(bis):
        where = f"{spec} {band['band']} {band['preset']} {band['faction']} {slot} {kind} {item_id}"
        item = items.get(item_id)
        if item is None:
            problems.append(f"{where}: not equippable by a {class_slug} (class mask or skill)")
            continue
        key = (item["class_id"], item["subclass_id"])
        gated = item["class_id"] == ITEM_CLASS_WEAPON or (
            item["class_id"] == ITEM_CLASS_ARMOR
            and item["subclass_id"] not in UNGATED_ARMOR_SUBCLASSES
        )
        if gated and skills.get(key, 10**6) > band["band"]:
            problems.append(
                f"{where} {item['name']}: skill learned at {skills.get(key)}, band {band['band']}"
            )
    return problems


BIS_FILES = sorted((BUILD_DIR / "bis").glob("*.json"))


@pytest.mark.parametrize("path", BIS_FILES, ids=[p.stem for p in BIS_FILES])
def test_every_bis_item_is_usable_by_the_spec_class_at_its_band(path: Path) -> None:
    bis = _json(path)
    assert _unusable(bis["spec"], bis) == []


def test_bis_files_exist_for_the_active_build() -> None:
    assert BIS_FILES, f"no BiS files under {BUILD_DIR / 'bis'}"


def test_the_rule_rejects_another_class_s_tier_piece() -> None:
    """The contract itself catches a warrior offered a hunter's class-locked set piece."""
    fake = {
        "bands": [
            {
                "band": 60,
                "preset": "raid",
                "faction": "horde",
                "slots": [{"slot": "head", "item_id": DRAGONSTALKERS_HELM, "alternatives": []}],
            }
        ]
    }
    assert len(_unusable("warrior-fury", fake)) == 1
    assert _unusable("hunter-marksmanship", fake) == []


def test_the_rule_rejects_armour_the_class_has_not_learned_at_the_band() -> None:
    """Plate opens at 40 for a warrior: a level-30 band offering plate is caught."""

    def at(band: int) -> dict:
        return {
            "bands": [
                {
                    "band": band,
                    "preset": "bare",
                    "faction": "horde",
                    "slots": [{"slot": "head", "item_id": HELM_OF_WRATH, "alternatives": []}],
                }
            ]
        }

    assert len(_unusable("warrior-fury", at(30))) == 1
    assert _unusable("warrior-fury", at(40)) == []


def test_class_locked_tier_pieces_stay_with_their_class() -> None:
    warrior, hunter = _class_items("warrior"), _class_items("hunter")
    assert DRAGONSTALKERS_HELM in hunter and DRAGONSTALKERS_HELM not in warrior
    assert HELM_OF_WRATH in warrior and HELM_OF_WRATH not in hunter


def test_dungeon_set_one_is_not_class_locked_in_the_client() -> None:
    """Beaststalker's Pants is mail with no class mask: a hunter and a warrior may both wear it."""
    assert BEASTSTALKERS_PANTS in _class_items("hunter")
    assert BEASTSTALKERS_PANTS in _class_items("warrior")
    assert BEASTSTALKERS_PANTS not in _class_items("rogue")  # leather wearer: no mail skill
