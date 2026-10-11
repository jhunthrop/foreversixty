import json
from pathlib import Path

import pytest

from pipeline.encounters import EncounterNeedsError, load_encounter_needs

BUILD = "1.60.1.70291"


def _build(tmp_path: Path, needs: dict) -> tuple[Path, Path]:
    """A build directory with one raid boss (Lucifron), one dungeon boss and
    two spell families, and a curated directory holding `needs`."""
    root = tmp_path / "builds"
    (root / BUILD).mkdir(parents=True)
    (root / BUILD / "loot.json").write_text(
        json.dumps(
            {
                "sources": [
                    {
                        "id": "raid:molten-core",
                        "kind": "raid",
                        "name": "Molten Core",
                        "bosses": [
                            {
                                "id": "raid:molten-core:12118",
                                "name": "Lucifron",
                                "npc_id": 12118,
                                "items": [],
                            }
                        ],
                    },
                    {
                        "id": "dungeon:x",
                        "kind": "dungeon",
                        "name": "X",
                        "bosses": [
                            {
                                "id": "dungeon:x:1",
                                "name": "Not a raid boss",
                                "npc_id": 1,
                                "items": [],
                            }
                        ],
                    },
                ]
            }
        ),
        encoding="utf-8",
    )
    (root / BUILD / "spellranks.json").write_text(
        json.dumps(
            {
                "build": BUILD,
                "classes": {
                    "warrior": {
                        "Sunder Armor": [
                            {"id": 7386, "level": 10, "rank": 1},
                            {"id": 11597, "level": 58, "rank": 5},
                        ],
                        "Battle Shout": [{"id": 11551, "level": 52, "rank": 6}],
                    }
                },
            }
        ),
        encoding="utf-8",
    )
    curated = tmp_path / "curated"
    curated.mkdir()
    (curated / "encounter_needs.json").write_text(json.dumps(needs), encoding="utf-8")
    return root, curated


def _boss(npc_id=12118, buffs=(11551,), debuffs=(7386,), resist="fire"):
    return {
        "npc_id": npc_id,
        "name": "Lucifron",
        "instance": "Molten Core",
        "buffs": list(buffs),
        "debuffs": list(debuffs),
        "resist": resist,
    }


def _needs(*bosses):
    return {"sources": [], "notes": "", "bosses": list(bosses)}


def test_a_valid_file_loads_keyed_by_npc_id(tmp_path):
    root, curated = _build(tmp_path, _needs(_boss()))
    needs = load_encounter_needs(root, BUILD, curated)
    assert list(needs) == [12118]
    assert needs[12118].buffs == [11551]
    assert needs[12118].debuffs == [7386]  # any rank of a known family
    assert needs[12118].resist == "fire"


def test_an_npc_id_that_is_not_a_raid_boss_in_the_build_is_refused(tmp_path):
    root, curated = _build(tmp_path, _needs(_boss(npc_id=1)))
    with pytest.raises(EncounterNeedsError, match="npc 1 "):
        load_encounter_needs(root, BUILD, curated)


def test_a_spell_id_spellranks_does_not_know_is_refused(tmp_path):
    root, curated = _build(tmp_path, _needs(_boss(debuffs=(17937,))))
    with pytest.raises(EncounterNeedsError, match="17937"):
        load_encounter_needs(root, BUILD, curated)


def test_an_unknown_resist_school_is_refused(tmp_path):
    root, curated = _build(tmp_path, _needs(_boss(resist="lava")))
    with pytest.raises(EncounterNeedsError, match="lava"):
        load_encounter_needs(root, BUILD, curated)


def test_a_boss_listed_twice_is_refused(tmp_path):
    root, curated = _build(tmp_path, _needs(_boss(), _boss()))
    with pytest.raises(EncounterNeedsError, match="twice"):
        load_encounter_needs(root, BUILD, curated)


def test_a_build_with_no_loot_record_ships_no_needs(tmp_path):
    root, curated = _build(tmp_path, _needs(_boss()))
    (root / BUILD / "loot.json").unlink()
    assert load_encounter_needs(root, BUILD, curated) == {}


def test_the_committed_file_validates_against_the_committed_build():
    needs = load_encounter_needs(Path("builds"), BUILD, Path("curated"))
    assert 12118 in needs and 10184 in needs
