from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.normalize import write_json
from pipeline.normalize.talents import normalize_talents

HERE = Path(__file__).parent


def test_talents_match_golden(tmp_path: Path):
    talent_rows = read_csv(HERE / "fixtures/Talent.csv")
    tab_rows = read_csv(HERE / "fixtures/TalentTab.csv")
    nodes = normalize_talents(talent_rows, tab_rows)
    out = tmp_path / "talents.json"
    write_json(nodes, out)
    assert out.read_text() == (HERE / "golden/talents.json").read_text()
