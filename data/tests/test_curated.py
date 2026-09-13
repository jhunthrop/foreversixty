import json
from pathlib import Path

import pytest

from pipeline.csvio import read_csv
from pipeline.curated import CuratedError, merge_curated
from pipeline.models import PlayableClass, PlayableRace
from pipeline.normalize import write_json, write_records
from pipeline.normalize.classes import normalize_classes, normalize_races

HERE = Path(__file__).parent
ERA_BUILD = Path("builds/1.15.9.69722")


def client_rows():
    return (
        normalize_classes(read_csv(HERE / "fixtures/ChrClasses.csv")),
        normalize_races(read_csv(HERE / "fixtures/ChrRaces.csv")),
    )


def merged():
    classes, races = client_rows()
    return merge_curated(classes, races, HERE / "fixtures/curated")


def test_forever_changes_are_attached_by_slug():
    classes, _races, _combos = merged()
    warrior = {c.slug: c for c in classes}["warrior"]
    assert [change.text for change in warrior.forever_changes] == ["Warriors gain a fourth tree."]
    assert warrior.forever_changes[0].sources[0].url == "https://example.test/forever-warrior"


def test_classes_without_curated_entries_keep_an_empty_list():
    classes, _races, _combos = merged()
    assert {c.slug: c for c in classes}["paladin"].forever_changes == []


def test_a_placeholder_race_is_appended_and_marked():
    _classes, races, _combos = merged()
    skyborne = {r.slug: r for r in races}["skyborne"]
    assert (skyborne.id, skyborne.faction, skyborne.placeholder) == (900, "neutral", True)
    assert {r.slug: r for r in races}["human"].placeholder is False


def test_combos_are_emitted_sorted(tmp_path: Path):
    _classes, _races, combos = merged()
    out = tmp_path / "combos.json"
    write_records(combos, out)
    assert json.loads(out.read_text()) == [
        {"race_id": 1, "class_id": 1, "new_in_forever": False},
        {"race_id": 1, "class_id": 2, "new_in_forever": True},
    ]
    assert out.read_text() == (HERE / "golden/combos.json").read_text()


def test_a_change_with_no_source_fails(tmp_path: Path):
    (tmp_path / "classes.json").write_text(
        json.dumps(
            [{"slug": "warrior", "forever_changes": [{"text": "Trust me.", "sources": []}]}]
        )
    )
    (tmp_path / "races.json").write_text("[]")
    (tmp_path / "combos.json").write_text("[]")
    classes, races = client_rows()
    with pytest.raises(CuratedError, match="at least one source"):
        merge_curated(classes, races, tmp_path)


def test_an_unknown_source_kind_fails(tmp_path: Path):
    (tmp_path / "classes.json").write_text(
        json.dumps(
            [
                {
                    "slug": "warrior",
                    "forever_changes": [
                        {
                            "text": "x",
                            "sources": [{"label": "l", "url": "https://e.test", "kind": "rumour"}],
                        }
                    ],
                }
            ]
        )
    )
    (tmp_path / "races.json").write_text("[]")
    (tmp_path / "combos.json").write_text("[]")
    classes, races = client_rows()
    with pytest.raises(CuratedError, match="rumour"):
        merge_curated(classes, races, tmp_path)


def test_an_unknown_slug_fails(tmp_path: Path):
    (tmp_path / "classes.json").write_text(
        json.dumps([{"slug": "necromancer", "forever_changes": []}])
    )
    (tmp_path / "races.json").write_text("[]")
    (tmp_path / "combos.json").write_text("[]")
    classes, races = client_rows()
    with pytest.raises(CuratedError, match="necromancer"):
        merge_curated(classes, races, tmp_path)


def test_a_combo_naming_an_unknown_id_fails(tmp_path: Path):
    (tmp_path / "classes.json").write_text("[]")
    (tmp_path / "races.json").write_text("[]")
    (tmp_path / "combos.json").write_text(
        json.dumps([{"race_id": 1, "class_id": 99, "new_in_forever": False}])
    )
    classes, races = client_rows()
    with pytest.raises(CuratedError, match="class_id 99"):
        merge_curated(classes, races, tmp_path)


def test_a_new_combo_without_a_source_fails(tmp_path: Path):
    (tmp_path / "classes.json").write_text("[]")
    (tmp_path / "races.json").write_text("[]")
    (tmp_path / "combos.json").write_text(
        json.dumps([{"race_id": 1, "class_id": 1, "new_in_forever": True}])
    )
    classes, races = client_rows()
    with pytest.raises(CuratedError, match="at least one source"):
        merge_curated(classes, races, tmp_path)


def test_a_placeholder_race_may_not_reuse_a_client_race_id(tmp_path: Path):
    (tmp_path / "classes.json").write_text("[]")
    (tmp_path / "races.json").write_text(
        json.dumps(
            [
                {
                    "id": 1,
                    "name": "Skyborne",
                    "slug": "skyborne",
                    "faction": "neutral",
                    "placeholder": True,
                    "forever_changes": [],
                }
            ]
        )
    )
    (tmp_path / "combos.json").write_text("[]")
    classes, races = client_rows()
    with pytest.raises(CuratedError, match="race id 1"):
        merge_curated(classes, races, tmp_path)


def test_merged_classes_and_races_match_golden(tmp_path: Path):
    classes, races, _combos = merged()
    write_json(classes, tmp_path / "classes.json")
    write_json(races, tmp_path / "races.json")
    assert (tmp_path / "classes.json").read_text() == (
        HERE / "golden/classes_merged.json"
    ).read_text()
    assert (tmp_path / "races.json").read_text() == (
        HERE / "golden/races_merged.json"
    ).read_text()


def test_the_real_curated_directory_is_valid():
    """data/curated must merge cleanly onto the committed Classic Era rows."""
    classes = [PlayableClass(**row) for row in json.loads((ERA_BUILD / "classes.json").read_text())]
    races = [PlayableRace(**row) for row in json.loads((ERA_BUILD / "races.json").read_text())]
    _classes, merged_races, combos = merge_curated(classes, races, Path("curated"))
    assert "skyborne" in {race.slug for race in merged_races}
    assert len(combos) == 40
