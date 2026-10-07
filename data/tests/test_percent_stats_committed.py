"""The committed planner items carry `percent_stats`: the aura-sourced (literal
percent) share of each rating-family stat, so the ranker converts only the
combat-rating share (pipeline/simdb/ratings.py) and never an aura percent."""

import json
from pathlib import Path

import pytest

from pipeline.normalize.gear import RATING_FAMILY_STAT_KEYS

BUILD = Path(__file__).parent.parent / "builds" / "1.60.1.70009"
LIONHEART_HELM = 12640  # 20 hit / 28 crit rating points
FURY_VISOR = 20521  # on-equip aura: a literal 1% hit / 1% crit


def _warrior_items() -> dict[int, dict]:
    path = BUILD / "items" / "warrior.json"
    if not path.exists():
        pytest.skip("build 1.60.1.70009 is not committed here")
    return {item["id"]: item for item in json.loads(path.read_text())["items"]}


def test_a_rating_item_carries_no_percent_stats():
    helm = _warrior_items()[LIONHEART_HELM]
    assert helm["stats"]["hit"] == 20 and helm["stats"]["crit"] == 28
    assert helm["percent_stats"] == {}


def test_an_aura_item_names_its_hit_and_crit_as_percent():
    visor = _warrior_items()[FURY_VISOR]
    assert visor["percent_stats"] == {"hit": 1, "crit": 1}


def test_percent_stats_is_a_rating_family_share_of_stats_on_every_item():
    for item in _warrior_items().values():
        for key, amount in item["percent_stats"].items():
            assert key in RATING_FAMILY_STAT_KEYS, item["id"]
            assert 0 < amount <= item["stats"][key], item["id"]
