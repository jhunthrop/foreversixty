from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.normalize import write_json
from pipeline.normalize.zones import normalize_zones

HERE = Path(__file__).parent


def test_zones_match_golden(tmp_path: Path):
    area_table = read_csv(HERE / "fixtures/AreaTable.csv")
    map_table = read_csv(HERE / "fixtures/Map.csv")
    zones = normalize_zones(area_table, map_table)
    out = tmp_path / "zones.json"
    write_json(zones, out)
    assert out.read_text() == (HERE / "golden/zones.json").read_text()


def test_write_json_is_deterministic(tmp_path: Path):
    area_table = read_csv(HERE / "fixtures/AreaTable.csv")
    map_table = read_csv(HERE / "fixtures/Map.csv")
    zones = normalize_zones(area_table, map_table)
    a, b = tmp_path / "a.json", tmp_path / "b.json"
    write_json(list(reversed(zones)), a)
    write_json(zones, b)
    assert a.read_bytes() == b.read_bytes()
