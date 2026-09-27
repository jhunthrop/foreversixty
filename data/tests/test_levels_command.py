"""`python -m pipeline levels`: the CLI command that writes both levels.json and
spellranks.json for a build, run after `simconst` (see pipeline/levels.py and
pipeline/spellranks.py for the two files' own contracts).
"""

import json
import shutil
from pathlib import Path

from pipeline.__main__ import main

HERE = Path(__file__).parent
WOWHEAD_FIXTURE = HERE / "fixtures" / "wowhead-gear-planner.js"
SPELLCONST_FIXTURE = HERE / "fixtures" / "spellconst"


def _class_and_race_rows() -> tuple[list[dict], list[dict]]:
    classes = [
        {"id": 1, "name": "Warrior", "slug": "warrior"},
        {"id": 2, "name": "Paladin", "slug": "paladin"},
    ]
    races = [
        {"id": 1, "name": "Human", "slug": "human"},
        {"id": 95, "name": "High-Order Skyborne", "slug": "high-order-skyborne"},
    ]
    return classes, races


def _prepare_build(tmp_path: Path, monkeypatch, with_wowhead_payload: bool) -> Path:
    build_dir = tmp_path / "builds" / "1.0.0.1"
    (build_dir / "spellconst").mkdir(parents=True)
    for path in SPELLCONST_FIXTURE.glob("*.json"):
        shutil.copy(path, build_dir / "spellconst" / path.name)

    classes, races = _class_and_race_rows()
    (build_dir / "classes.json").write_text(json.dumps(classes))
    (build_dir / "races.json").write_text(json.dumps(races))
    manifest = {"build": "1.0.0.1", "product": "test", "fetched_at": "t", "files": {}}
    (build_dir / "manifest.json").write_text(json.dumps(manifest))

    if with_wowhead_payload:
        raw = build_dir / "raw"
        raw.mkdir()
        (raw / "wowhead-gear-planner.js").write_text(
            WOWHEAD_FIXTURE.read_text(encoding="utf-8")
        )
        meta = {"url": "https://example.test/gear-planner", "fetched_at": "t"}
        (raw / "wowhead-gear-planner.meta.json").write_text(json.dumps(meta))

    monkeypatch.chdir(tmp_path)
    return build_dir


def test_levels_command_writes_both_files_when_a_wowhead_payload_exists(tmp_path, monkeypatch):
    build_dir = _prepare_build(tmp_path, monkeypatch, with_wowhead_payload=True)
    assert main(["levels", "--build", "1.0.0.1"]) == 0
    assert (build_dir / "levels.json").exists()
    assert (build_dir / "spellranks.json").exists()
    levels = json.loads((build_dir / "levels.json").read_text())
    assert levels["classes"]["warrior"]["levels"][-1]["stamina"] == 110
    spellranks = json.loads((build_dir / "spellranks.json").read_text())
    assert "Sinister Strike" in spellranks["classes"]["rogue"]


def test_levels_command_writes_only_spellranks_without_a_wowhead_payload(
    tmp_path, monkeypatch, caplog
):
    build_dir = _prepare_build(tmp_path, monkeypatch, with_wowhead_payload=False)
    with caplog.at_level("WARNING"):
        assert main(["levels", "--build", "1.0.0.1"]) == 0
    assert not (build_dir / "levels.json").exists()
    assert (build_dir / "spellranks.json").exists()
    assert "no wowhead payload" in caplog.text
