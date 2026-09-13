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
