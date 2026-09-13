import json
import shutil
from pathlib import Path

from pipeline.normalize import normalize_build
from pipeline.wago import TABLES

HERE = Path(__file__).parent
# classes.json and races.json are compared against the merged goldens below,
# because the orchestrator folds data/curated into them.
ENTITY_NAMES = ["zones", "dungeons", "items", "spells", "talents"]


def prepare(tmp_path: Path) -> Path:
    raw = tmp_path / "1.0.0.1" / "raw"
    raw.mkdir(parents=True)
    for table in TABLES:
        shutil.copy(HERE / "fixtures" / f"{table}.csv", raw / f"{table}.csv")
    meta = {"product": "test", "build": "1.0.0.1", "fetched_at": "t"}
    (raw / "_meta.json").write_text(json.dumps(meta))
    return tmp_path


def run(tmp_path: Path) -> Path:
    return normalize_build("1.0.0.1", root=prepare(tmp_path), curated_dir=HERE / "fixtures/curated")


def test_phase_0_entities_still_match_golden(tmp_path: Path):
    out = run(tmp_path)
    for name in ENTITY_NAMES:
        assert (out / f"{name}.json").read_text() == (HERE / "golden" / f"{name}.json").read_text()


def test_classes_and_races_are_merged_with_the_curated_facts(tmp_path: Path):
    out = run(tmp_path)
    assert (out / "classes.json").read_text() == (
        HERE / "golden/classes_merged.json"
    ).read_text()
    assert (out / "races.json").read_text() == (HERE / "golden/races_merged.json").read_text()


def test_phase_1_files_are_written(tmp_path: Path):
    out = run(tmp_path)
    assert (out / "talents" / "warrior.json").read_text() == (
        HERE / "golden/talents_warrior.json"
    ).read_text()
    assert (out / "items" / "warrior.json").read_text() == (
        HERE / "golden/items_warrior.json"
    ).read_text()
    assert (out / "sets.json").read_text() == (HERE / "golden/sets.json").read_text()
    assert (out / "combos.json").read_text() == (HERE / "golden/combos.json").read_text()


def test_manifest_lists_every_emitted_file(tmp_path: Path):
    out = run(tmp_path)
    files = set(json.loads((out / "manifest.json").read_text())["files"])
    assert {f"{name}.json" for name in ENTITY_NAMES} <= files
    assert {"classes.json", "races.json"} <= files
    assert {"sets.json", "combos.json", "talents/warrior.json", "items/warrior.json"} <= files


def test_a_rerun_is_byte_identical_and_clears_stale_class_files(tmp_path: Path):
    out = run(tmp_path)
    before = {p.relative_to(out).as_posix(): p.read_bytes() for p in sorted(out.rglob("*.json"))}
    (out / "talents" / "necromancer.json").write_text("{}\n")
    normalize_build("1.0.0.1", root=tmp_path, curated_dir=HERE / "fixtures/curated")
    after = {p.relative_to(out).as_posix(): p.read_bytes() for p in sorted(out.rglob("*.json"))}
    assert after == before


def test_items_are_skipped_when_the_item_table_is_unreadable(tmp_path: Path, caplog):
    root = prepare(tmp_path)
    sparse = root / "1.0.0.1" / "raw" / "ItemSparse.csv"
    sparse.write_text(sparse.read_text().replace(",7,4,35,15", ",99,4,35,15"))
    out = normalize_build("1.0.0.1", root=root, curated_dir=HERE / "fixtures/curated")
    assert not (out / "items").exists()
    assert (out / "sets.json").exists()
    assert "unknown stat modifier id 99" in caplog.text
