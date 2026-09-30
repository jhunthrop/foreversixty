# data/tests/test_audit_quests.py
"""Check B (quests vs classic-db) against a small synthetic
`raw/classicdb/sources.json` cache and a matching `loot.json`."""

import json
from pathlib import Path

from pipeline.audit import check_quests
from pipeline.audit.context import AuditContext
from pipeline.classic_sources import (
    ClassicDbQuestInfo,
    ClassicDbSourceRecord,
    write_classic_sources,
)


def _write_build(root: Path, build: str) -> Path:
    build_dir = root / build
    build_dir.mkdir(parents=True)
    write_classic_sources(
        build_dir,
        {
            200: [
                ClassicDbSourceRecord(
                    kind="quest_reward", name="Level Mismatch Quest",
                    quest=ClassicDbQuestInfo(
                        quest_id=5000, min_level=10, level=15, faction="alliance"
                    ),
                )
            ],
            201: [
                ClassicDbSourceRecord(
                    kind="quest_reward", name="Faction Mismatch Quest",
                    quest=ClassicDbQuestInfo(
                        quest_id=6000, min_level=10, level=10, faction="alliance"
                    ),
                )
            ],
            210: [
                ClassicDbSourceRecord(
                    kind="quest_reward", name="Destroy Morphaz",
                    quest=ClassicDbQuestInfo(
                        quest_id=9000, min_level=50, level=52, faction="both",
                        classes=["mage"],
                    ),
                )
            ],
        },
    )  # fmt: skip
    loot = {
        "sources": [],
        "quests": {
            "200": [
                {
                    "quest_id": 5000, "name": "Level Mismatch Quest", "faction": "alliance",
                    "min_level": 10, "level": 12, "level_source": "classic-db",
                }
            ],
            "201": [
                {
                    "quest_id": 6000, "name": "Faction Mismatch Quest", "faction": "horde",
                    "min_level": 10, "level": 10, "level_source": "classic-db",
                }
            ],
            "202": [
                {
                    "quest_id": 7000, "name": "Forever New Quest", "faction": "both",
                    "min_level": 30, "level": 32, "level_source": "wowhead",
                }
            ],
            "203": [
                {
                    "quest_id": 5000, "name": "Level Mismatch Quest", "faction": "alliance",
                    "min_level": 10, "level": 15, "level_source": "classic-db",
                }
            ],
            # No `classes` key at all: classic-db's own RequiredClasses
            # (quest 9000, mage-only) says this should carry one.
            "210": [
                {
                    "quest_id": 9000, "name": "Destroy Morphaz", "faction": "both",
                    "min_level": 50, "level": 52, "level_source": "classic-db",
                }
            ],
        },
        "factions": {},
    }  # fmt: skip
    (build_dir / "loot.json").write_text(json.dumps(loot))
    return build_dir


def test_level_mismatch_is_major(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_quests.check(ctx)
    finding = next(
        f for f in result.findings if f.subject == "5000" and "level disagrees" in f.message
    )
    assert finding.severity == "major"
    assert finding.ours == "12" and finding.theirs == "15"


def test_faction_mismatch_is_blocker(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_quests.check(ctx)
    finding = next(
        f for f in result.findings if f.subject == "6000" and "faction disagrees" in f.message
    )
    assert finding.severity == "blocker"
    assert finding.ours == "horde" and finding.theirs == "alliance"


def test_quest_absent_from_classic_db_is_forever_new(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_quests.check(ctx)
    finding = next(f for f in result.findings if f.subject == "7000")
    assert finding.severity == "minor"
    assert "Forever-new" in finding.message


def test_reward_not_in_classic_reward_list(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_quests.check(ctx)
    finding = next(f for f in result.findings if f.subject == "203")
    assert "reward not in 1.12" in finding.message
    assert "does not name it either" in finding.message


def test_missing_classes_when_classic_db_states_required_classes_is_major(tmp_path):
    """Day3 data-followups-7 lane, 2026-09-30: quest 9000's own classic-db
    `RequiredClasses` (mage-only) is not reflected in loot.json's own
    `quests["210"]` entry at all -- a regression in `build_loot`'s own
    direct pass-through this check now guards, the same way it already
    guards `faction`/`min_level`/`level` drift."""
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_quests.check(ctx)
    finding = next(
        f
        for f in result.findings
        if f.subject == "9000" and "publishes no `classes`" in f.message
    )
    assert finding.severity == "major"
    assert finding.theirs == "mage"


def test_classes_present_and_agreeing_is_not_flagged(tmp_path):
    root = tmp_path / "builds"
    build_dir = _write_build(root, "testbuild")
    loot = json.loads((build_dir / "loot.json").read_text())
    loot["quests"]["210"][0]["classes"] = ["mage"]
    (build_dir / "loot.json").write_text(json.dumps(loot))
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_quests.check(ctx)
    assert not any(
        f.subject == "9000" and "publishes no `classes`" in f.message for f in result.findings
    )
