from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.normalize import write_json
from pipeline.normalize.classes import normalize_classes, normalize_races

HERE = Path(__file__).parent


def test_classes_match_golden(tmp_path: Path):
    classes = normalize_classes(read_csv(HERE / "fixtures/ChrClasses.csv"))
    out = tmp_path / "classes.json"
    write_json(classes, out)
    assert out.read_text() == (HERE / "golden/classes.json").read_text()


def test_races_match_golden(tmp_path: Path):
    races = normalize_races(read_csv(HERE / "fixtures/ChrRaces.csv"))
    out = tmp_path / "races.json"
    write_json(races, out)
    assert out.read_text() == (HERE / "golden/races.json").read_text()


def test_a_playable_race_bit_of_minus_one_is_dropped_not_a_duplicate():
    """The 1.60 client (Forever beta) exports PlayableRaceBit, and some rows carry -1
    in it: an internal duplicate the client itself never lets a player pick (id 33 here
    is a real row, an unused "ThinHuman" body-type copy of Human). Keeping it would
    collide by slug with the real playable Human (id 1) in curated.merge_curated's
    by-slug lookup, so it is dropped before a slug is ever computed. The two real
    Skyborne rows (95, 96) both carry a non-negative bit and both survive, one per
    faction."""
    races = normalize_races(read_csv(HERE / "fixtures/ChrRaces_1_60.csv"))
    assert [r.id for r in races] == [1, 95, 96]
    assert {r.slug for r in races} == {"human", "high-order-skyborne", "windshaper-skyborne"}
