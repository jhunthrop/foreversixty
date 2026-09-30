# data/tests/test_audit_drops.py
"""Check C (drops vs classic-db), item-level half only (no --classicdb-dump,
so boss existence/spawn-map is exercised in test_audit_dumpdb.py instead)."""

import json
from pathlib import Path

from pipeline.audit import check_drops
from pipeline.audit.context import AuditContext
from pipeline.classic_sources import ClassicDbSourceRecord, write_classic_sources


def _write_build(root: Path, build: str) -> Path:
    build_dir = root / build
    build_dir.mkdir(parents=True)
    write_classic_sources(
        build_dir,
        {
            300: [
                ClassicDbSourceRecord(
                    kind="creature_drop", npc_id=9001, name="Test Boss", chance=10.0
                ),
                ClassicDbSourceRecord(kind="world_drop", name="World drop", chance=10.0),
            ],
            301: [
                ClassicDbSourceRecord(
                    kind="creature_drop", npc_id=9001, name="Test Boss", chance=50.0
                ),
            ],
            # 302: no classic-db record at all.
            # 303 (trash): no classic-db record at all.
        },
    )  # fmt: skip
    loot = {
        "sources": [
            {
                "id": "dungeon:test:9001", "kind": "dungeon", "name": "Test Dungeon",
                "zone_id": 1,
                "bosses": [
                    {
                        "id": "dungeon:test:9001", "name": "Test Boss", "npc_id": 9001,
                        "items": [300, 301, 302],
                        "item_chances": {"300": 10.0, "301": 5.0, "302": 2.0},
                    }
                ],
                "trash": [303],
            }
        ],
        "quests": {},
        "factions": {},
    }  # fmt: skip
    (build_dir / "loot.json").write_text(json.dumps(loot))
    return build_dir


def test_matching_drop_is_clean(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_drops.check(ctx)
    assert not any(f.subject == "300" and "chance" in f.message for f in result.findings)


def test_chance_mismatch_is_major(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_drops.check(ctx)
    finding = next(f for f in result.findings if f.subject == "301" and "chance" in f.message)
    assert finding.severity == "major"
    assert finding.ours == "5.0" and finding.theirs == "50.0"


def test_item_missing_from_classic_db_is_major(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_drops.check(ctx)
    finding = next(f for f in result.findings if f.subject == "302")
    assert finding.severity == "major"
    assert "no such drop" in finding.message


def test_world_drop_pool_item_is_flagged(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_drops.check(ctx)
    finding = next(
        f for f in result.findings if f.subject == "300" and "world-drop pool" in f.message
    )
    assert finding.severity == "minor"


def test_trash_item_missing_is_flagged(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_drops.check(ctx)
    finding = next(f for f in result.findings if f.subject == "303")
    assert finding.severity == "major"


def test_item_missing_from_classic_db_in_a_raid_gated_instance_is_minor(tmp_path):
    """drop-sources-2 lane, 2026-09-29: a raid-gated instance (`opens` set
    -- the phase calendar has not opened it yet) cannot be verified
    against the live game either, so its own unconfirmed boss items are
    `minor`, not `major`, alongside the pre-existing Forever-new-id
    exception (`test_item_missing_from_classic_db_is_major`'s own
    item 302, a LAUNCH dungeon, stays `major`)."""
    root = tmp_path / "builds"
    build_dir = root / "testbuild"
    build_dir.mkdir(parents=True)
    write_classic_sources(
        build_dir,
        {
            999: [
                ClassicDbSourceRecord(
                    kind="creature_drop", npc_id=1, name="Unrelated", chance=1.0
                )
            ]
        },
    )
    loot = {
        "sources": [
            {
                "id": "raid:test", "kind": "raid", "name": "Test Raid", "zone_id": 2,
                "opens": "2026-12-09",
                "bosses": [
                    {
                        "id": "raid:test:9002", "name": "Gated Boss", "npc_id": 9002,
                        "items": [400],
                    },
                ],
            }
        ],
        "quests": {}, "factions": {},
    }  # fmt: skip
    (build_dir / "loot.json").write_text(json.dumps(loot))
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_drops.check(ctx)
    finding = next(f for f in result.findings if f.subject == "400")
    assert finding.severity == "minor"
    assert "raid-gated" in finding.message


def test_item_missing_from_classic_db_with_wowhead_origin_is_minor(tmp_path):
    """A `LootBoss.item_source_origin` of `"wowhead"` is already labelled
    unverified right on the data (`item_source_origin`'s own doc), so the
    audit's own finding for it is `minor`, confirming a label already
    there rather than surfacing a silent gap."""
    root = tmp_path / "builds"
    build_dir = root / "testbuild"
    build_dir.mkdir(parents=True)
    write_classic_sources(
        build_dir,
        {
            999: [
                ClassicDbSourceRecord(
                    kind="creature_drop", npc_id=1, name="Unrelated", chance=1.0
                )
            ]
        },
    )
    loot = {
        "sources": [
            {
                "id": "dungeon:test", "kind": "dungeon", "name": "Test Dungeon", "zone_id": 3,
                "bosses": [
                    {
                        "id": "dungeon:test:9003", "name": "Wowhead Boss", "npc_id": 9003,
                        "items": [401], "item_source_origin": {"401": "wowhead"},
                    },
                ],
            }
        ],
        "quests": {}, "factions": {},
    }  # fmt: skip
    (build_dir / "loot.json").write_text(json.dumps(loot))
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_drops.check(ctx)
    finding = next(f for f in result.findings if f.subject == "401")
    assert finding.severity == "minor"
    assert "wowhead" in finding.message


def test_item_missing_from_classic_db_with_fork_origin_in_a_launch_dungeon_is_major(tmp_path):
    """The one case worth a person's judgment: a `"fork"`-origin
    attribution (explicitly labelled, not just unlabelled) in a launch,
    non-raid-gated instance that neither classic-db nor wowhead
    corroborates stays `major`."""
    root = tmp_path / "builds"
    build_dir = root / "testbuild"
    build_dir.mkdir(parents=True)
    write_classic_sources(
        build_dir,
        {
            999: [
                ClassicDbSourceRecord(
                    kind="creature_drop", npc_id=1, name="Unrelated", chance=1.0
                )
            ]
        },
    )
    loot = {
        "sources": [
            {
                "id": "dungeon:test", "kind": "dungeon", "name": "Test Dungeon", "zone_id": 3,
                "bosses": [
                    {
                        "id": "dungeon:test:9004", "name": "Fork Boss", "npc_id": 9004,
                        "items": [402], "item_source_origin": {"402": "fork"},
                    },
                ],
            }
        ],
        "quests": {}, "factions": {},
    }  # fmt: skip
    (build_dir / "loot.json").write_text(json.dumps(loot))
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_drops.check(ctx)
    finding = next(f for f in result.findings if f.subject == "402")
    assert finding.severity == "major"


def test_skipped_note_when_no_dump_given(tmp_path):
    root = tmp_path / "builds"
    _write_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_drops.check(ctx)
    assert "classicdb-dump" in result.skipped
