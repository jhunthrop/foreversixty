"""The committed 1.60 build must be the trees measured off the client.

test_build_conformance.py checks the Classic Era build against the Phase 1
interface contract. This one checks the beta build against the facts recorded
in docs/superpowers/plans/2026-09-17-forever-trees.md: the numbers are the
measurement, so a regeneration that moves one has to be looked at.
"""

import json
from functools import cache
from pathlib import Path

BUILD = "1.60.1.69893"
BUILD_DIR = Path("builds") / BUILD

#: class slug -> the three tree names left to right, and the talent count in each.
TABS = {
    "warrior": [("Arms", 17), ("Fury", 18), ("Protection", 18)],
    "paladin": [("Holy", 18), ("Protection", 16), ("Retribution", 18)],
    "hunter": [("Beast Mastery", 16), ("Marksmanship", 17), ("Survival", 18)],
    "rogue": [("Assassination", 17), ("Combat", 17), ("Subtlety", 19)],
    "priest": [("Discipline", 18), ("Holy", 17), ("Shadow", 18)],
    "shaman": [("Elemental", 16), ("Enhancement", 18), ("Restoration", 16)],
    "mage": [("Arcane", 18), ("Fire", 17), ("Frost", 19)],
    "warlock": [("Affliction", 17), ("Demonology", 19), ("Destruction", 16)],
    "druid": [("Balance", 16), ("Feral Combat", 19), ("Restoration", 16)],
}
TALENT_TOTAL = 469
PREREQUISITE_TOTAL = 69
GRID_ROWS = 7
GRID_COLUMNS = 4


@cache
def talents(slug: str) -> dict:
    return json.loads((BUILD_DIR / "talents" / f"{slug}.json").read_text(encoding="utf-8"))


def test_every_class_has_its_three_tabs_in_the_games_order():
    for slug, want in TABS.items():
        trees = talents(slug)["trees"]
        assert [t["position"] for t in trees] == [0, 1, 2], slug
        assert [(t["name"], len(t["talents"])) for t in trees] == want, slug


def test_the_build_carries_every_talent_measured_off_the_client():
    total = sum(len(t["talents"]) for slug in TABS for t in talents(slug)["trees"])
    assert total == TALENT_TOTAL


def test_every_tree_names_a_background_and_every_talent_a_spell():
    for slug in TABS:
        for tree in talents(slug)["trees"]:
            assert tree["background"], slug
            assert tree["background"] == tree["background"].lower(), slug
            for talent in tree["talents"]:
                assert talent["spell_id"] > 0, (slug, talent["id"])
                assert talent["ranks"][0]["spell_id"] == talent["spell_id"]


def test_every_talent_sits_alone_on_the_four_by_seven_grid():
    for slug in TABS:
        for tree in talents(slug)["trees"]:
            cells = [(t["tier"], t["column"]) for t in tree["talents"]]
            assert len(set(cells)) == len(cells), (slug, tree["name"])
            for tier, column in cells:
                assert 0 <= tier < GRID_ROWS and 0 <= column < GRID_COLUMNS, (slug, tier, column)


def test_every_prerequisite_is_above_or_beside_its_dependent_in_the_same_tree():
    total = 0
    for slug in TABS:
        for tree in talents(slug)["trees"]:
            by_id = {t["id"]: t for t in tree["talents"]}
            for talent in tree["talents"]:
                if talent["prereq_talent_id"] is None:
                    assert talent["prereq_rank"] is None, (slug, talent["id"])
                    continue
                total += 1
                prereq = by_id[talent["prereq_talent_id"]]
                assert prereq["tier"] <= talent["tier"], (slug, talent["name"])
                assert talent["prereq_rank"] == prereq["max_rank"], (slug, talent["name"])
    assert total == PREREQUISITE_TOTAL
