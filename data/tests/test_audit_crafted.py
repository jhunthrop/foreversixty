# data/tests/test_audit_crafted.py
"""Check E (crafted) against a synthetic classic-db dump: one Classic-id
item with a resolvable recipe spell, one Classic-id item with none, and one
Forever-new item id (>= FOREVER_NEW_ID_THRESHOLD) with none (expected, and
counted separately)."""

import json
from pathlib import Path

from pipeline.audit import check_crafted
from pipeline.audit.context import AuditContext
from pipeline.audit.dumpdb import SPELL_EFFECT_CREATE_ITEM
from pipeline.loot.wowhead import FOREVER_NEW_ID_THRESHOLD

DUMP_SQL = f"""
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
INSERT INTO `item_template` VALUES (700,0,0,0,0,0,0,0);

CREATE TABLE `spell_template` (
  `Id` int,
  `Effect1` int,
  `Effect2` int,
  `Effect3` int,
  `EffectItemType1` int,
  `EffectItemType2` int,
  `EffectItemType3` int
) ENGINE=MyISAM;
INSERT INTO `spell_template` VALUES (9001,{SPELL_EFFECT_CREATE_ITEM},0,0,700,0,0);
"""


def _write_build(root: Path, build: str, forever_item_id: int) -> Path:
    build_dir = root / build
    build_dir.mkdir(parents=True)
    loot = {
        "sources": [
            {
                "id": "crafted:blacksmithing", "kind": "crafted", "name": "Blacksmithing",
                "profession": "blacksmithing", "items": [700, 701, forever_item_id],
            }
        ],
        "quests": {},
        "factions": {},
    }
    (build_dir / "loot.json").write_text(json.dumps(loot))
    return build_dir


def _ctx(tmp_path, forever_item_id):
    root = tmp_path / "builds"
    _write_build(root, "testbuild", forever_item_id)
    dump_path = tmp_path / "classicdb.sql"
    dump_path.write_text(DUMP_SQL, encoding="utf-8")
    return AuditContext(
        "testbuild", root=root, curated_dir=tmp_path / "curated", classicdb_dump=dump_path
    )


def test_item_with_recipe_spell_is_clean(tmp_path):
    ctx = _ctx(tmp_path, FOREVER_NEW_ID_THRESHOLD + 1)
    result = check_crafted.check(ctx)
    assert not any(f.subject == "700" for f in result.findings)


def test_classic_item_with_no_recipe_is_major(tmp_path):
    ctx = _ctx(tmp_path, FOREVER_NEW_ID_THRESHOLD + 1)
    result = check_crafted.check(ctx)
    finding = next(f for f in result.findings if f.subject == "701")
    assert finding.severity == "major"


def test_forever_new_item_with_no_recipe_is_expected(tmp_path):
    forever_id = FOREVER_NEW_ID_THRESHOLD + 1
    ctx = _ctx(tmp_path, forever_id)
    result = check_crafted.check(ctx)
    assert not any(f.subject == str(forever_id) for f in result.findings)
    summary = next(f for f in result.findings if "Forever-new" in f.message)
    assert "1 Forever-new" in summary.message


def test_skipped_without_dump(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild", FOREVER_NEW_ID_THRESHOLD + 1)
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_crafted.check(ctx)
    assert "classicdb-dump" in result.skipped
    assert result.findings == []
