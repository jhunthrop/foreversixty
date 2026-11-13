import json
from pathlib import Path

import pytest

from pipeline.raid_supplies import RaidSuppliesError, load_consumables

BUILD = "1.60.1.70291"


def _build(tmp_path: Path, supplies: list, protection: list) -> tuple[Path, Path]:
    """A build whose simconsumes.json knows a flask, a food and a fire protection
    potion, and a curated directory holding the given lists."""
    root = tmp_path / "builds"
    (root / BUILD).mkdir(parents=True)
    (root / BUILD / "simconsumes.json").write_text(
        json.dumps(
            [
                {"id": 13510, "name": "Flask of the Titans"},
                {"id": 13931, "name": "Nightfin Soup"},
                {"id": 13457, "name": "Greater Fire Protection Potion"},
            ]
        ),
        encoding="utf-8",
    )
    curated = tmp_path / "curated"
    curated.mkdir()
    (curated / "raid-supplies.json").write_text(
        json.dumps({"notes": "", "supplies": supplies, "protection": protection}), encoding="utf-8"
    )
    return root, curated


FLASK = {"item_id": 13510, "name": "Flask of the Titans", "kind": "flask"}
FOOD = {"item_id": 13931, "name": "Nightfin Soup", "kind": "food"}
FIRE = {"item_id": 13457, "name": "Greater Fire Protection Potion", "kind": "fire_protection"}


def test_both_lists_load_keyed_by_item_id_ascending(tmp_path):
    root, curated = _build(tmp_path, [FOOD, FLASK], [FIRE])
    assert load_consumables(root, BUILD, curated) == {
        13457: "fire_protection",
        13510: "flask",
        13931: "food",
    }


def test_a_name_the_build_does_not_give_the_id_is_refused(tmp_path):
    root, curated = _build(tmp_path, [{**FLASK, "name": "Flask of the Giants"}], [])
    with pytest.raises(RaidSuppliesError, match="13510"):
        load_consumables(root, BUILD, curated)


def test_a_kind_out_of_its_list_is_refused(tmp_path):
    root, curated = _build(tmp_path, [FIRE], [])
    with pytest.raises(RaidSuppliesError, match="fire_protection"):
        load_consumables(root, BUILD, curated)


def test_an_item_listed_twice_is_refused(tmp_path):
    root, curated = _build(tmp_path, [FLASK, FLASK], [])
    with pytest.raises(RaidSuppliesError, match="twice"):
        load_consumables(root, BUILD, curated)


def test_a_build_with_no_simconsumes_ships_no_catalogue(tmp_path):
    root, curated = _build(tmp_path, [FLASK], [])
    (root / BUILD / "simconsumes.json").unlink()
    assert load_consumables(root, BUILD, curated) == {}


def test_the_committed_file_validates_against_the_active_build():
    build = json.loads(Path("../web/src/data/active-build.json").read_text())["build"]
    consumables = load_consumables(Path("builds"), build, Path("curated"))
    assert consumables[13510] == "flask" and consumables[13457] == "fire_protection"
