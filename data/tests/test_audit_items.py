# data/tests/test_audit_items.py
"""Check A (items vs client) -- a tiny synthetic build directory: one item
whose published required_level disagrees with the raw ItemSparse.csv, one
kept at the placeholder icon, and one required_level 0 with no wowhead
payload to cross-check (so it is counted, not individually flagged)."""

import json
from pathlib import Path

from pipeline.audit import check_items
from pipeline.audit.context import AuditContext
from pipeline.icons import PLACEHOLDER_ICON


def _write_build(root: Path, build: str) -> Path:
    build_dir = root / build
    (build_dir / "items").mkdir(parents=True)
    (build_dir / "raw").mkdir(parents=True)
    (build_dir / "classes.json").write_text(json.dumps([{"slug": "warrior"}]))
    items = [
        {
            "id": 100, "name": "Mismatched Helm", "icon": "inv_helmet_01", "slot": "head",
            "quality": 2, "required_level": 99, "item_level": 20, "armor": 0, "stats": {},
            "damage_min": 0, "damage_max": 0, "speed": 0.0, "dps": 0.0, "two_hand": False,
            "effect_text": "", "set_id": None, "unique": False,
        },
        {
            "id": 101, "name": "Placeholder Trinket", "icon": PLACEHOLDER_ICON,
            "slot": "trinket", "quality": 1, "required_level": 20, "item_level": 20,
            "armor": 0, "stats": {}, "damage_min": 0, "damage_max": 0, "speed": 0.0,
            "dps": 0.0, "two_hand": False, "effect_text": "", "set_id": None, "unique": False,
        },
        {
            "id": 102, "name": "Quest Reward Ring", "icon": "inv_ring_01", "slot": "finger",
            "quality": 2, "required_level": 0, "item_level": 20, "armor": 0, "stats": {},
            "damage_min": 0, "damage_max": 0, "speed": 0.0, "dps": 0.0, "two_hand": False,
            "effect_text": "", "set_id": None, "unique": False,
        },
    ]  # fmt: skip
    (build_dir / "items" / "warrior.json").write_text(
        json.dumps({"build": build, "class_slug": "warrior", "items": items})
    )
    (build_dir / "raw" / "ItemSparse.csv").write_text(
        "ID,RequiredLevel,ItemLevel,OverallQualityID,InventoryType\n"
        "100,10,20,2,1\n"
        "101,20,20,1,12\n"
        "102,0,20,2,11\n"
    )
    (build_dir / "raw" / "Item.csv").write_text(
        "ID,ClassID,SubclassID\n100,4,1\n101,4,0\n102,4,0\n"
    )
    return build_dir


def test_required_level_mismatch_is_a_blocker(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_items.check(ctx)
    mismatches = [
        f for f in result.findings if f.subject == "100" and "required_level" in f.message
    ]
    assert len(mismatches) == 1
    assert mismatches[0].severity == "blocker"
    assert mismatches[0].ours == "99"
    assert mismatches[0].theirs == "10"


def test_placeholder_icon_is_flagged(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_items.check(ctx)
    assert any(f.subject == "101" and "placeholder" in f.message for f in result.findings)


def test_required_level_zero_is_counted_not_individually_flagged(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_items.check(ctx)
    summary = next(f for f in result.findings if f.subject == "*")
    assert "1 item(s)" in summary.message
    assert not any(f.subject == "102" for f in result.findings)


def test_skips_when_raw_tables_absent(tmp_path):
    root = tmp_path / "builds"
    build_dir = _write_build(root, "testbuild")
    (build_dir / "raw" / "ItemSparse.csv").unlink()
    (build_dir / "raw" / "Item.csv").unlink()
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_items.check(ctx)
    assert "not committed" in result.skipped
    # The checks that do not need raw tables still ran.
    assert any(f.subject == "101" for f in result.findings)
