import json
import shutil
from pathlib import Path

from pipeline.normalize import normalize_build

HERE = Path(__file__).parent
TABLES = [
    "Map",
    "AreaTable",
    "JournalInstance",
    "ItemSparse",
    "Item",
    "SpellName",
    "ChrClasses",
    "ChrRaces",
    "Talent",
    "TalentTab",
]
ENTITY_NAMES = ["zones", "dungeons", "items", "spells", "classes", "races", "talents"]


def test_normalize_build_writes_all_entities_and_manifest(tmp_path: Path):
    raw = tmp_path / "1.0.0.1" / "raw"
    raw.mkdir(parents=True)
    for t in TABLES:
        shutil.copy(HERE / "fixtures" / f"{t}.csv", raw / f"{t}.csv")
    meta = {"product": "test", "build": "1.0.0.1", "fetched_at": "t"}
    (raw / "_meta.json").write_text(json.dumps(meta))
    out = normalize_build("1.0.0.1", root=tmp_path)
    for name in ENTITY_NAMES:
        assert (out / f"{name}.json").read_text() == (HERE / "golden" / f"{name}.json").read_text()
    m = json.loads((out / "manifest.json").read_text())
    assert m["build"] == "1.0.0.1"
    assert set(m["files"]) == {f"{n}.json" for n in ENTITY_NAMES}
