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
            "quality": 2, "required_level": 99, "required_level_source": "client",
            "item_level": 20, "armor": 0, "stats": {},
            "damage_min": 0, "damage_max": 0, "speed": 0.0, "dps": 0.0, "two_hand": False,
            "effect_text": "", "set_id": None, "unique": False,
        },
        {
            "id": 101, "name": "Placeholder Trinket", "icon": PLACEHOLDER_ICON,
            "slot": "trinket", "quality": 1, "required_level": 20,
            "required_level_source": "client", "item_level": 20,
            "armor": 0, "stats": {}, "damage_min": 0, "damage_max": 0, "speed": 0.0,
            "dps": 0.0, "two_hand": False, "effect_text": "", "set_id": None, "unique": False,
        },
        {
            "id": 102, "name": "Quest Reward Ring", "icon": "inv_ring_01", "slot": "finger",
            "quality": 2, "required_level": 0, "required_level_source": "none",
            "item_level": 20, "armor": 0, "stats": {},
            "damage_min": 0, "damage_max": 0, "speed": 0.0, "dps": 0.0, "two_hand": False,
            "effect_text": "", "set_id": None, "unique": False,
        },
        {
            "id": 103, "name": "Wowhead-Sourced Belt", "icon": "inv_belt_01", "slot": "waist",
            "quality": 2, "required_level": 25, "required_level_source": "wowhead",
            "item_level": 20, "armor": 0, "stats": {},
            "damage_min": 0, "damage_max": 0, "speed": 0.0, "dps": 0.0, "two_hand": False,
            "effect_text": "", "set_id": None, "unique": False,
        },
        {
            "id": 104, "name": "Wowhead-Only Cloak", "icon": "inv_cloak_01", "slot": "back",
            "quality": 2, "required_level": 30, "required_level_source": "wowhead",
            "item_level": 20, "armor": 0, "stats": {},
            "damage_min": 0, "damage_max": 0, "speed": 0.0, "dps": 0.0, "two_hand": False,
            "effect_text": "", "set_id": None, "unique": False,
        },
        {
            "id": 107, "name": "Classic-DB Only Trinket", "icon": "inv_trinket_01",
            "slot": "trinket", "quality": 3, "required_level": 53,
            "required_level_source": "classic-db", "stats_source": "classic-db",
            "client_unconfirmed": True, "item_level": 58, "armor": 0, "stats": {},
            "damage_min": 0, "damage_max": 0, "speed": 0.0, "dps": 0.0, "two_hand": False,
            "effect_text": "", "set_id": None, "unique": False,
        },
        {
            "id": 105, "name": "Proxy-Estimated Wristguard", "icon": "inv_bracer_01",
            "slot": "wrist", "quality": 2, "required_level": 15,
            "required_level_source": "item_level_proxy", "item_level": 20, "armor": 0,
            "stats": {}, "damage_min": 0, "damage_max": 0, "speed": 0.0, "dps": 0.0,
            "two_hand": False, "effect_text": "", "set_id": None, "unique": False,
        },
        {
            "id": 106, "name": "Proxy-Only Signet", "icon": "inv_ring_02", "slot": "finger",
            "quality": 2, "required_level": 15, "required_level_source": "item_level_proxy",
            "item_level": 20, "armor": 0, "stats": {},
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
        "103,0,20,2,6\n"
        "105,0,20,2,9\n"
    )
    (build_dir / "raw" / "Item.csv").write_text(
        "ID,ClassID,SubclassID\n100,4,1\n101,4,0\n102,4,0\n103,4,0\n105,4,0\n"
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


def test_required_level_mismatch_is_minor_when_source_is_wowhead(tmp_path):
    """A row present in the client's own raw tables (RequiredLevel 0) whose
    published level came from wowhead (`resolve_required_level`'s
    `"wowhead"` branch) is a deliberate, sourced departure from the client's
    stated 0 -- not a disagreement, so it stays off the blocker path."""
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_items.check(ctx)
    findings = [
        f for f in result.findings if f.subject == "103" and "required_level" in f.message
    ]
    assert len(findings) == 1
    assert findings[0].severity == "minor"
    assert "wowhead estimate" in findings[0].message
    assert not any(f.subject == "103" and f.severity == "blocker" for f in result.findings)


def test_required_level_mismatch_is_minor_when_source_is_item_level_proxy(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_items.check(ctx)
    findings = [
        f for f in result.findings if f.subject == "105" and "required_level" in f.message
    ]
    assert len(findings) == 1
    assert findings[0].severity == "minor"
    assert "item_level_proxy estimate" in findings[0].message


def test_missing_from_raw_is_minor_when_source_is_wowhead(tmp_path):
    """An item with no client row at all (a wowhead-supplement item) is
    never going to be in the raw tables; that is what makes it a supplement,
    not the "major" defect a real client id going missing would be."""
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_items.check(ctx)
    findings = [f for f in result.findings if f.subject == "104"]
    assert len(findings) == 1
    assert findings[0].severity == "minor"
    assert "not in the build's own raw" in findings[0].message


def test_missing_from_raw_is_minor_for_a_classic_db_row(tmp_path):
    """A classic-db 1.12 row (catalogue-universe, `client_unconfirmed`) has no
    client row by construction and says so on the site; it is unverifiable
    here, not a missing client id."""
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_items.check(ctx)
    findings = [f for f in result.findings if f.subject == "107"]
    assert len(findings) == 1
    assert findings[0].severity == "minor"
    assert "classic-db 1.12 row" in findings[0].message


def test_missing_from_raw_stays_major_when_source_is_not_wowhead(tmp_path):
    """A supplement item wowhead itself names no level for (source fell
    through to the item-level proxy) is unverified against every source,
    client or wowhead -- the softer treatment above is for the id wowhead
    specifically corroborates, not for every supplement id."""
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_items.check(ctx)
    findings = [f for f in result.findings if f.subject == "106"]
    assert len(findings) == 1
    assert findings[0].severity == "major"


def test_a_hotfix_only_item_counts_as_in_the_clients_raw_table(tmp_path):
    """`AuditContext.raw_item_sparse`/`raw_item` merge `raw/hotfixes/*.csv`
    over the shipped tables (`pipeline.hotfix_merge.merge_hotfix_table`,
    normalize-levels lane, 2026-09-29) the same way `normalize` itself does:
    an id the client only carries as a hotfix is "in the client's raw table"
    here too, not a false "not in raw" finding -- item 200 below has no row
    in `raw/ItemSparse.csv`/`Item.csv` at all, only in `raw/hotfixes/`."""
    root = tmp_path / "builds"
    build_dir = _write_build(root, "testbuild")
    items = json.loads((build_dir / "items" / "warrior.json").read_text())
    items["items"].append(
        {
            "id": 200, "name": "Hotfix-Only Girdle", "icon": "inv_belt_02", "slot": "waist",
            "quality": 2, "required_level": 18, "required_level_source": "client",
            "item_level": 20, "armor": 0, "stats": {},
            "damage_min": 0, "damage_max": 0, "speed": 0.0, "dps": 0.0, "two_hand": False,
            "effect_text": "", "set_id": None, "unique": False,
        }
    )
    (build_dir / "items" / "warrior.json").write_text(json.dumps(items))
    (build_dir / "raw" / "hotfixes" / "ItemSparse.csv").parent.mkdir(
        parents=True, exist_ok=True
    )
    (build_dir / "raw" / "hotfixes" / "ItemSparse.csv").write_text(
        "ID,RequiredLevel,ItemLevel,OverallQualityID,InventoryType\n200,18,20,2,6\n"
    )
    (build_dir / "raw" / "hotfixes" / "Item.csv").write_text("ID,ClassID,SubclassID\n200,4,0\n")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_items.check(ctx)
    assert not any(f.subject == "200" for f in result.findings)


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
