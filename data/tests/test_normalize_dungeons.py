from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.normalize import write_json
from pipeline.normalize.dungeons import normalize_dungeons

HERE = Path(__file__).parent


def test_dungeons_match_golden(tmp_path: Path):
    dungeons = normalize_dungeons(
        read_csv(HERE / "fixtures/JournalInstance.csv"),
        read_csv(HERE / "fixtures/Map.csv"),
    )
    out = tmp_path / "dungeons.json"
    write_json(dungeons, out)
    assert out.read_text() == (HERE / "golden/dungeons.json").read_text()
