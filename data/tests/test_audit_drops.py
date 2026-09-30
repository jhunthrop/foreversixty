# data/tests/test_audit_drops.py
"""Check C (drops vs classic-db), item-level half only (no --classicdb-dump,
so boss existence/spawn-map is exercised in test_audit_dumpdb.py instead)."""

import json
from pathlib import Path

from pipeline.audit import check_drops
from pipeline.audit.context import AuditContext
from pipeline.classic_sources import ClassicDbSourceRecord, write_classic_sources

#: A synthetic dump slice covering the "boss is really a gameobject chest"
#: shape check_drops.py's own item-level fix resolves (this lane's brief
#: item 3). Two CHEST (type 3) `gameobject_template` rows both spawn on
#: map 429 (`gameobject`): 179564 "Gordok Tribute", the real one, whose
#: own `data1` names a DIFFERENT `gameobject_loot_template` entry (16577,
#: one direct row plus one NEGATIVE `mincountOrRef` reference row -- the
#: same recursive shape `test_classic_sources.py`'s own
#: `test_gameobject_chest_loot_*` coverage already pins), and 179999 "A
#: Decoy Chest", an UNRELATED chest sharing the same map but naming
#: neither of the boss's own items -- proving `_resolve_chest_loot`
#: picks the real candidate by item OVERLAP, not just "any same-map
#: chest". The `LootBoss.npc_id` for the real one, 14324, is neither: a
#: fork-internal placeholder matching no classic-db creature OR object
#: entry at all, the actual shape measured on the pinned dump.
CHEST_DUMP_SQL = """
CREATE TABLE `creature_template` (
  `Entry` mediumint,
  `Name` char(100)
) ENGINE=MyISAM;

CREATE TABLE `creature` (
  `guid` int,
  `id` mediumint,
  `map` smallint
) ENGINE=MyISAM;

CREATE TABLE `gameobject_template` (
  `entry` mediumint,
  `type` tinyint,
  `name` varchar(100),
  `data1` int
) ENGINE=MyISAM;
INSERT INTO `gameobject_template` VALUES
  (179564,3,'Gordok Tribute',16577),
  (179999,3,'A Decoy Chest',16999);

CREATE TABLE `gameobject` (
  `guid` int,
  `id` mediumint,
  `map` smallint
) ENGINE=MyISAM;
INSERT INTO `gameobject` VALUES (1,179564,429),(2,179999,429);

CREATE TABLE `gameobject_loot_template` (
  `entry` mediumint,
  `item` mediumint,
  `ChanceOrQuestChance` float,
  `groupid` tinyint,
  `mincountOrRef` mediumint,
  `maxcount` tinyint,
  `condition_id` mediumint,
  `comments` varchar(300)
) ENGINE=MyISAM;
INSERT INTO `gameobject_loot_template` VALUES
  (16577,8952,0,0,15,1,0,'Direct'),
  (16577,0,100,0,-35033,1,0,'Reference'),
  (16999,90001,0,0,15,1,0,'Direct');

CREATE TABLE `reference_loot_template` (
  `entry` mediumint,
  `item` mediumint,
  `ChanceOrQuestChance` float,
  `groupid` tinyint,
  `mincountOrRef` mediumint,
  `maxcount` tinyint,
  `condition_id` mediumint,
  `comments` varchar(300)
) ENGINE=MyISAM;
INSERT INTO `reference_loot_template` VALUES (35033,35033,3.0,0,1,1,0,'');
"""


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


def _write_chest_build(root: Path, build: str) -> Path:
    """A "boss" that is really Dire Maul's Gordok Tribute chest: its
    `npc_id` (14324) is a fork-internal placeholder matching no
    classic-db creature OR object entry at all (the actual pinned-dump
    shape, `CHEST_DUMP_SQL`'s own doc) -- classic-db's
    `raw/classicdb/sources.json` cache alone names no `creature_drop`/
    `object_drop` record for either item (their real loot is keyed by
    `data1`, 16577, under object entry 179564, not `npc_id` at all), so
    exercising this lane's brief item 3's fix needs a
    `--classicdb-dump` AND a `zones.json` (`zone_id` 4 -> map 429, the
    same map `CHEST_DUMP_SQL`'s own `gameobject` spawns 179564/179999
    on)."""
    build_dir = root / build
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
    )  # fmt: skip
    loot = {
        "sources": [
            {
                "id": "dungeon:dire-maul", "kind": "dungeon", "name": "Dire Maul", "zone_id": 4,
                "bosses": [
                    {
                        "id": "dungeon:dire-maul:14324", "name": "Tribute", "npc_id": 14324,
                        "items": [8952, 35033],
                        "item_chances": {"8952": 15.0, "35033": 3.0},
                    },
                ],
            }
        ],
        "quests": {}, "factions": {},
    }  # fmt: skip
    (build_dir / "loot.json").write_text(json.dumps(loot))
    (build_dir / "zones.json").write_text(json.dumps([{"id": 4, "map_id": 429}]))
    return build_dir


def test_a_boss_whose_npc_id_is_really_a_gameobject_chest_is_not_falsely_flagged(tmp_path):
    """This lane's brief item 3: with a `--classicdb-dump`, both the
    direct row (8952) and the reference-expanded row (35033) resolve
    through `gameobject_template.data1`, so neither is a false "no such
    drop"."""
    root = tmp_path / "builds"
    _write_chest_build(root, "testbuild")
    dump_path = tmp_path / "classicdb.sql"
    dump_path.write_text(CHEST_DUMP_SQL, encoding="utf-8")
    ctx = AuditContext(
        "testbuild", root=root, curated_dir=tmp_path / "curated", classicdb_dump=dump_path
    )
    result = check_drops.check(ctx)
    assert not any(f.subject in ("8952", "35033") for f in result.findings)


def test_a_boss_with_both_a_real_creature_drop_and_chest_items_resolves_both_independently(
    tmp_path,
):
    """Dire Maul's own measured shape: `npc_id` 14324 is ALSO a real
    classic-db creature ("Cho'Rush the Observer") that genuinely drops
    SOME of this boss's own items alongside the chest-only ones -- a
    boss-level "any real match at all" gate would wrongly skip chest
    resolution for 8952/35033 just because ANOTHER item on the same boss
    has a real `creature_drop` record; `_boss_items` must resolve each
    item independently and never mix the two signals (a genuine
    `creature_drop` chance is never diluted by an unrelated chest's own
    chance, or vice versa)."""
    root = tmp_path / "builds"
    build_dir = _write_chest_build(root, "testbuild")
    write_classic_sources(
        build_dir,
        {
            999: [
                ClassicDbSourceRecord(
                    kind="creature_drop", npc_id=1, name="Unrelated", chance=1.0
                )
            ],
            90099: [
                ClassicDbSourceRecord(
                    kind="creature_drop", npc_id=14324, name="Cho'Rush the Observer",
                    chance=25.0,
                )
            ],
        },
    )  # fmt: skip
    loot = json.loads((build_dir / "loot.json").read_text())
    boss = loot["sources"][0]["bosses"][0]
    boss["items"].append(90099)
    boss["item_chances"]["90099"] = 25.0
    (build_dir / "loot.json").write_text(json.dumps(loot))
    dump_path = tmp_path / "classicdb.sql"
    dump_path.write_text(CHEST_DUMP_SQL, encoding="utf-8")
    ctx = AuditContext(
        "testbuild", root=root, curated_dir=tmp_path / "curated", classicdb_dump=dump_path
    )
    result = check_drops.check(ctx)
    assert not any(f.subject in ("8952", "35033", "90099") for f in result.findings)


def test_a_boss_whose_npc_id_is_really_a_gameobject_chest_is_falsely_flagged_without_a_dump(
    tmp_path,
):
    """Without `--classicdb-dump`, the indirection cannot be resolved at
    all (same "labelled, not skipped outright" rule the rest of the
    boss-level half already follows) -- both items are still flagged, the
    pre-fix behaviour this lane's brief measured as 20 + 8 false majors."""
    root = tmp_path / "builds"
    _write_chest_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_drops.check(ctx)
    assert {f.subject for f in result.findings if "no such drop" in f.message} == {
        "8952", "35033",
    }  # fmt: skip


def _write_world_drop_leak_build(
    root: Path, build: str, *, item_in_boss_list: bool = False
) -> Path:
    """day3 data-followups-6 lane, 2026-09-30: a published BiS pick (item
    90200) whose own classic-db drop row resolves to a known instance
    zone (raid:test-raid, zone 5000, map 500), but `loot.json`'s own
    published sources carry it under NOTHING but `world_drop:60-62` --
    the Cache of the Firelord shape this lane's fix closes. `raid:
    test-raid` itself names a DIFFERENT item (1) as its own boss's real
    drop, so the instance zone is genuinely known without item 90200
    needing to be correctly attributed there first."""
    build_dir = root / build
    build_dir.mkdir(parents=True)
    write_classic_sources(
        build_dir,
        {
            1: [ClassicDbSourceRecord(kind="creature_drop", npc_id=1, name="Boss", map_id=500)],
            90200: [
                ClassicDbSourceRecord(
                    kind="object_drop", object_id=9001, name="Reward Chest", map_id=500,
                )
            ],
        },
    )  # fmt: skip
    boss_source = {
        "id": "raid:test-raid", "kind": "raid", "name": "Test Raid", "zone_id": 5000,
        "bosses": [{"id": "raid:test-raid:1", "name": "Boss", "npc_id": 1, "items": [1]}],
    }  # fmt: skip
    if item_in_boss_list:
        boss_source["bosses"][0]["items"].append(90200)
    loot = {
        "sources": [
            boss_source,
            {
                "id": "world_drop:60-62", "kind": "world_drop", "name": "World drop",
                "items": [90200], "level_min": 60, "level_max": 62,
            },
        ],
        "quests": {}, "factions": {},
    }  # fmt: skip
    (build_dir / "loot.json").write_text(json.dumps(loot))
    (build_dir / "zones.json").write_text(json.dumps([{"id": 5000, "map_id": 500}]))
    (build_dir / "items").mkdir()
    (build_dir / "items" / "warrior.json").write_text(
        json.dumps(
            {
                "items": [
                    {"id": 90200, "name": "Leaked Item", "required_level": 1, "slot": "trinket1"}
                ]
            }
        )
    )
    (build_dir / "bis").mkdir()
    (build_dir / "bis" / "warrior-arms.json").write_text(
        json.dumps(
            {
                "spec": "warrior-arms",
                "bands": [
                    {
                        "band": 60,
                        "faction": "alliance",
                        "slots": [
                            {
                                "slot": "trinket1",
                                "item_id": 90200,
                                "item_name": "Leaked Item",
                                "source": "World drop",
                                "source_kind": "world_drop",
                            }
                        ],
                    }
                ],
            }
        )
    )
    return build_dir


def test_a_world_drop_only_bis_pick_is_flagged_when_classic_db_places_it_in_an_instance(
    tmp_path,
):
    root = tmp_path / "builds"
    _write_world_drop_leak_build(root, "testbuild")
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_drops.check(ctx)
    leak_findings = [f for f in result.findings if f.subject == "90200"]
    assert leak_findings
    assert leak_findings[0].severity == "major"
    assert "world/world_drop" in leak_findings[0].message


def test_the_same_item_is_not_flagged_once_it_also_has_a_real_boss_listing(tmp_path):
    """The same item, corroborated at all outside `world`/`world_drop` --
    the leak this check exists for is specifically "ONLY a world bucket",
    not "also has one". Filters by this check's own message text, not
    subject alone: check C's own pre-existing `_boss_items` half may add
    its own, unrelated "no such drop" finding for the same item id (no
    `--classicdb-dump` given here), which is not this test's concern."""
    root = tmp_path / "builds"
    _write_world_drop_leak_build(root, "testbuild", item_in_boss_list=True)
    ctx = AuditContext("testbuild", root=root, curated_dir=tmp_path / "curated")
    result = check_drops.check(ctx)
    assert not any("world/world_drop" in f.message for f in result.findings)
