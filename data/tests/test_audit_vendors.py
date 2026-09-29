# data/tests/test_audit_vendors.py
"""Check D (vendors and rep) against a synthetic classic-db cache, with and
without a `--classicdb-dump` for the item_template rep-column half."""

import json
from pathlib import Path

from pipeline.audit import check_vendors
from pipeline.audit.context import AuditContext
from pipeline.audit.dumpdb import ClassicDbDump
from pipeline.classic_sources import (
    ClassicDbCondition,
    ClassicDbSourceRecord,
    write_classic_sources,
)

DUMP_SQL = """
CREATE TABLE `item_template` (
  `entry` mediumint,
  `RequiredReputationFaction` smallint,
  `RequiredReputationRank` smallint,
  `spellid_1` mediumint,
  `spellid_2` mediumint,
  `spellid_3` mediumint,
  `spellid_4` mediumint,
  `spellid_5` mediumint
) ENGINE=MyISAM;
INSERT INTO `item_template` VALUES (400,50,2,0,0,0,0,0);

CREATE TABLE `spell_template` (
  `Id` int, `Effect1` int, `Effect2` int, `Effect3` int,
  `EffectItemType1` int, `EffectItemType2` int, `EffectItemType3` int
) ENGINE=MyISAM;
"""


def _write_build(root: Path, build: str) -> Path:
    build_dir = root / build
    build_dir.mkdir(parents=True)
    write_classic_sources(
        build_dir,
        {
            400: [ClassicDbSourceRecord(kind="vendor", npc_id=8001, name="Test Vendor")],
            401: [
                ClassicDbSourceRecord(
                    kind="vendor", npc_id=8001, name="Test Vendor",
                    condition=ClassicDbCondition(faction_id=87, standing="honored"),
                )
            ],
            # 402: no classic-db vendor row at all.
        },
    )  # fmt: skip
    loot = {
        "sources": [
            {
                "id": "vendor:8001", "kind": "vendor", "name": "Test Vendor", "npc_id": 8001,
                "items": [400, 401, 402],
            }
        ],
        "quests": {},
        "factions": {},
    }
    (build_dir / "loot.json").write_text(json.dumps(loot))
    return build_dir


def test_matching_vendor_item_is_clean(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_vendors.check(ctx)
    assert not any(f.subject == "400" and "no such row" in f.message for f in result.findings)


def test_rep_gate_mismatch_is_blocker(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_vendors.check(ctx)
    finding = next(f for f in result.findings if f.subject == "401")
    assert finding.severity == "blocker"
    assert "honored" in finding.theirs


def test_item_missing_from_npc_vendor_is_major(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_vendors.check(ctx)
    finding = next(f for f in result.findings if f.subject == "402")
    assert finding.severity == "major"


def test_item_template_reputation_column_checked_with_dump(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    dump_path = tmp_path / "classicdb.sql"
    dump_path.write_text(DUMP_SQL, encoding="utf-8")
    ctx = AuditContext(
        "testbuild", root=root, curated_dir=tmp_path / "curated", classicdb_dump=dump_path
    )
    assert isinstance(ctx.dump, ClassicDbDump)
    result = check_vendors.check(ctx)
    finding = next(f for f in result.findings if f.subject == "400")
    assert "RequiredReputationFaction=50" in finding.message
