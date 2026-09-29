# data/tests/test_audit_bis.py
"""Check F (leveling BiS vs rules) against a tiny synthetic
bis/warrior-fury.json + items/warrior.json + loot.json."""

import json
from pathlib import Path

from pipeline.audit import check_bis
from pipeline.audit.context import AuditContext


def _item(item_id, name, required_level=20, two_hand=False, unique=False):
    return {
        "id": item_id, "name": name, "icon": "inv_misc", "slot": "head", "quality": 2,
        "required_level": required_level, "item_level": 20, "armor": 0, "stats": {},
        "damage_min": 0, "damage_max": 0, "speed": 0.0, "dps": 0.0, "two_hand": two_hand,
        "effect_text": "", "set_id": None, "unique": unique,
    }  # fmt: skip


def _write_build(root: Path, build: str) -> Path:
    build_dir = root / build
    (build_dir / "items").mkdir(parents=True)
    (build_dir / "bis").mkdir(parents=True)
    (build_dir / "classes.json").write_text(json.dumps([{"slug": "warrior"}]))
    items = [
        _item(100, "Overleveled Helm", required_level=99),
        _item(101, "Two-Hand Weapon", required_level=20, two_hand=True),
        _item(102, "Legit Off Hand", required_level=20),
        _item(103, "Unique Trinket", required_level=20, unique=True),
        _item(200, "Real Crafted Item", required_level=20),
    ]
    (build_dir / "items" / "warrior.json").write_text(
        json.dumps({"build": build, "class_slug": "warrior", "items": items})
    )
    bands = [
        {
            "spec": "warrior-fury", "band": 20, "faction": "alliance", "race": "human",
            "slots": [
                {
                    "slot": "head", "item_id": 100, "item_name": "Overleveled Helm",
                    "source": "", "source_kind": "world",
                },
                {
                    "slot": "main_hand", "item_id": 101, "item_name": "Two-Hand Weapon",
                    "source": "Blacksmithing", "source_kind": "crafted",
                },
                {
                    "slot": "off_hand", "item_id": 102, "item_name": "Legit Off Hand",
                    "source": "Ghost Source", "source_kind": "world",
                },
                {
                    "slot": "trinket1", "item_id": 103, "item_name": "Unique Trinket",
                    "source": "", "source_kind": "world",
                },
                {
                    "slot": "trinket2", "item_id": 103, "item_name": "Unique Trinket",
                    "source": "", "source_kind": "world",
                },
            ],
        }
    ]  # fmt: skip
    (build_dir / "bis" / "warrior-fury.json").write_text(json.dumps({"bands": bands}))
    loot = {
        "sources": [
            {
                "id": "crafted:blacksmithing", "kind": "crafted", "name": "Blacksmithing",
                "profession": "blacksmithing", "items": [101, 200],
            }
        ],
        "quests": {},
        "factions": {"100": "horde"},
    }
    (build_dir / "loot.json").write_text(json.dumps(loot))
    return build_dir


def test_required_level_above_band_is_blocker(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_bis.check(ctx)
    finding = next(
        f for f in result.findings if f.subject == "100" and "required_level" in f.message
    )
    assert finding.severity == "blocker"


def test_faction_mismatch_is_blocker(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_bis.check(ctx)
    finding = next(f for f in result.findings if f.subject == "100" and "horde-only" in f.message)
    assert finding.severity == "blocker"


def test_off_hand_with_two_hand_main_hand_is_flagged(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_bis.check(ctx)
    finding = next(f for f in result.findings if f.subject == "102" and "two-handed" in f.message)
    assert "two-handed" in finding.message


def test_source_string_not_listing_item_is_flagged(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_bis.check(ctx)
    assert any(
        f.subject == "102" and "does not list the item" in f.message for f in result.findings
    )


def test_valid_crafted_source_is_not_flagged(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_bis.check(ctx)
    assert not any(
        f.subject == "101" and "does not list the item" in f.message for f in result.findings
    )


def test_duplicate_unique_trinket_is_blocker(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_bis.check(ctx)
    finding = next(
        f for f in result.findings if f.subject == "103" and "more than one slot" in f.message
    )
    assert finding.severity == "blocker"


def test_duplicate_trinket_pair_is_flagged(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_bis.check(ctx)
    assert any("not a distinct pair" in f.message for f in result.findings)
