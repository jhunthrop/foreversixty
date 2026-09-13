from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.normalize import write_json
from pipeline.normalize.spells import normalize_spells

HERE = Path(__file__).parent


def test_spells_match_golden(tmp_path: Path):
    spells = normalize_spells(read_csv(HERE / "fixtures/SpellName.csv"))
    out = tmp_path / "spells.json"
    write_json(spells, out)
    assert out.read_text() == (HERE / "golden/spells.json").read_text()
