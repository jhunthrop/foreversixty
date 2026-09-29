# data/tests/test_classic_sources.py
"""cmangos/classic-db's loot/vendor/quest tables as a third `loot.json`
source (src-classicdb lane, 2026-09-29) -- a small, synthetic mysqldump
snippet covering every table `pipeline.classic_sources` reads: a plain
creature drop, a NEGATIVE `mincountOrRef` reference (recursively, one
level deep), a gameobject (chest) drop, a rep-gated vendor slot (via
`npc_vendor_template`, joined through `creature_template.
VendorTemplateId`), and a quest reward with `RequiredRaces` restricted to
one faction. See `pipeline.classic_quest_levels`'s own test for the same
"comma inside a quoted field" concern this shares.
"""

import gzip

import httpx

from pipeline import classic_sources as cs

# Column counts below only need to be internally consistent with each
# table's own CREATE TABLE block -- `pipeline.sqldump.iter_table_records`
# checks the row's field count against the declared column count, not
# against the real dump's 130-column quest_template.
SAMPLE_SQL = """
CREATE TABLE `creature_template` (
  `Entry` mediumint,
  `Name` char(100),
  `VendorTemplateId` mediumint
) ENGINE=MyISAM;
INSERT INTO `creature_template` VALUES (100,'Lady Anacondra',0),(200,'Reagent Vendor',900);

CREATE TABLE `creature` (
  `guid` int,
  `id` mediumint,
  `map` smallint
) ENGINE=MyISAM;
INSERT INTO `creature` VALUES (1,100,43),(2,100,43),(3,200,0);

CREATE TABLE `gameobject_template` (
  `entry` mediumint,
  `name` varchar(100)
) ENGINE=MyISAM;
INSERT INTO `gameobject_template` VALUES (500,'Abercrombie\\'s Crate');

CREATE TABLE `gameobject` (
  `guid` int,
  `id` mediumint,
  `map` smallint
) ENGINE=MyISAM;
INSERT INTO `gameobject` VALUES (1,500,43);

CREATE TABLE `creature_loot_template` (
  `entry` mediumint,
  `item` mediumint,
  `ChanceOrQuestChance` float,
  `groupid` tinyint,
  `mincountOrRef` mediumint,
  `maxcount` tinyint,
  `condition_id` mediumint,
  `comments` varchar(300)
) ENGINE=MyISAM;
INSERT INTO `creature_loot_template` VALUES (100,5404,40,0,1,1,0,'Serpent''s Shoulders'),(100,999,0.8,0,-999,1,0,'Referenced loot group'),(300,60321,0.02,0,-60321,1,0,'NPC LOOT (Grey World Drop) - (Item Levels: 51-60)');

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
INSERT INTO `reference_loot_template` VALUES (999,7777,5,0,1,1,0,'Referenced item'),(60321,8888,3,0,1,1,0,'Some grey junk item');

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
INSERT INTO `gameobject_loot_template` VALUES (500,1951,6,0,1,1,0,'Blackwater Cutlass');

CREATE TABLE `skinning_loot_template` (
  `entry` mediumint,
  `item` mediumint,
  `ChanceOrQuestChance` float,
  `groupid` tinyint,
  `mincountOrRef` mediumint,
  `maxcount` tinyint,
  `condition_id` mediumint,
  `comments` varchar(300)
) ENGINE=MyISAM;
INSERT INTO `skinning_loot_template` VALUES (100,2934,60,1,1,1,0,'Ruined Leather Scraps');

CREATE TABLE `pickpocketing_loot_template` (
  `entry` mediumint,
  `item` mediumint,
  `ChanceOrQuestChance` float,
  `groupid` tinyint,
  `mincountOrRef` mediumint,
  `maxcount` tinyint,
  `condition_id` mediumint,
  `comments` varchar(300)
) ENGINE=MyISAM;
INSERT INTO `pickpocketing_loot_template` VALUES (100,1206,2,0,1,1,0,'Moss Agate');

CREATE TABLE `fishing_loot_template` (
  `entry` mediumint,
  `item` mediumint,
  `ChanceOrQuestChance` float,
  `groupid` tinyint,
  `mincountOrRef` mediumint,
  `maxcount` tinyint,
  `condition_id` mediumint,
  `comments` varchar(300)
) ENGINE=MyISAM;
INSERT INTO `fishing_loot_template` VALUES (1,6303,100,1,1,1,0,'Raw Bristle Whisker Catfish');

CREATE TABLE `npc_vendor` (
  `entry` mediumint,
  `item` mediumint,
  `maxcount` tinyint,
  `incrtime` int,
  `slot` tinyint,
  `condition_id` mediumint,
  `comments` text
) ENGINE=MyISAM;
INSERT INTO `npc_vendor` VALUES (100,6260,0,0,15,0,'Blue Dye');

CREATE TABLE `npc_vendor_template` (
  `entry` mediumint,
  `item` mediumint,
  `maxcount` tinyint,
  `incrtime` int,
  `slot` tinyint,
  `condition_id` mediumint,
  `comments` text
) ENGINE=MyISAM;
INSERT INTO `npc_vendor_template` VALUES (900,8925,0,0,125,1,'Crystal Vial');

CREATE TABLE `conditions` (
  `condition_entry` mediumint,
  `type` tinyint,
  `value1` mediumint,
  `value2` mediumint,
  `value3` mediumint,
  `value4` mediumint,
  `flags` tinyint,
  `comments` varchar(500)
) ENGINE=MyISAM;
INSERT INTO `conditions` VALUES (1,5,76,7,0,0,0,'Has Minimum Rank Exalted With Faction ID: 76');

CREATE TABLE `quest_template` (
  `entry` mediumint,
  `MinLevel` tinyint,
  `QuestLevel` smallint,
  `RequiredRaces` smallint,
  `Title` text,
  `RewChoiceItemId1` mediumint,
  `RewChoiceItemId2` mediumint,
  `RewChoiceItemId3` mediumint,
  `RewChoiceItemId4` mediumint,
  `RewChoiceItemId5` mediumint,
  `RewChoiceItemId6` mediumint,
  `RewItemId1` mediumint,
  `RewItemId2` mediumint,
  `RewItemId3` mediumint,
  `RewItemId4` mediumint
) ENGINE=MyISAM;
INSERT INTO `quest_template` VALUES
  (53,40,44,77,'Sweet Amber',744,0,0,0,0,0,0,0,0,0),
  (8,1,5,178,'A Rogues Deal',0,0,0,0,0,0,159,0,0,0);
"""  # noqa: E501


def test_parse_creature_drop_reads_name_map_and_chance():
    items = cs.parse_classic_db_sources(SAMPLE_SQL)
    records = items[5404]
    assert len(records) == 1
    record = records[0]
    assert record.kind == "creature_drop"
    assert record.npc_id == 100
    assert record.name == "Lady Anacondra"
    assert record.map_id == 43  # both her spawn rows are on map 43
    assert record.chance == 40.0


def test_negative_mincount_or_ref_expands_the_reference_loot_template():
    items = cs.parse_classic_db_sources(SAMPLE_SQL)
    records = items[7777]
    assert len(records) == 1
    assert records[0].kind == "creature_drop"
    assert records[0].npc_id == 100
    assert records[0].chance == 5.0  # the REFERENCE row's own chance, not the outer 0.8


def test_a_world_drop_commented_reference_contributes_nothing():
    """cmangos' own generic level-gated pool (any qualifying creature can
    drop from it -- see pipeline.classic_sources' own doc for the 300+
    item Onyxia's-Lair-trash-mob regression this prevents): item 8888,
    behind reference 60321, is never attributed to creature 300 at all."""
    items = cs.parse_classic_db_sources(SAMPLE_SQL)
    assert 8888 not in items


def test_gameobject_drop_reads_the_chest_name_and_spawn_map():
    items = cs.parse_classic_db_sources(SAMPLE_SQL)
    records = items[1951]
    assert len(records) == 1
    assert records[0].kind == "object_drop"
    assert records[0].object_id == 500
    assert records[0].name == "Abercrombie's Crate"
    assert records[0].map_id == 43
    assert records[0].chance == 6.0


def test_skinning_and_pickpocketing_are_kept_distinct_from_a_plain_drop():
    items = cs.parse_classic_db_sources(SAMPLE_SQL)
    assert items[2934][0].kind == "skinning"
    assert items[2934][0].npc_id == 100
    assert items[1206][0].kind == "pickpocketing"


def test_fishing_has_no_npc_but_still_names_a_flat_bucket():
    items = cs.parse_classic_db_sources(SAMPLE_SQL)
    record = items[6303][0]
    assert record.kind == "fishing"
    assert record.name == "Fishing"
    assert record.npc_id is None
    assert record.chance == 100.0


def test_direct_vendor_row_has_no_condition():
    items = cs.parse_classic_db_sources(SAMPLE_SQL)
    record = next(r for r in items[6260] if r.kind == "vendor")
    assert record.npc_id == 100
    assert record.name == "Lady Anacondra"
    assert record.condition is None


def test_vendor_template_join_resolves_a_rep_gated_slot():
    """Npc 200's own VendorTemplateId (900) pulls in
    npc_vendor_template's row for entry 900 -- item 8925, condition_id 1,
    which `conditions` names as CONDITION_REPUTATION_RANK (type 5):
    faction 76, rank 7 (cmangos' own 0-indexed Exalted) -> this module's
    1-indexed `REP_LEVELS[8]` == "exalted"."""
    items = cs.parse_classic_db_sources(SAMPLE_SQL)
    record = next(r for r in items[8925] if r.kind == "vendor")
    assert record.npc_id == 200
    assert record.name == "Reagent Vendor"
    assert record.condition is not None
    assert record.condition.faction_id == 76
    assert record.condition.standing == "exalted"


def test_quest_reward_reads_level_fields_and_alliance_only_faction():
    items = cs.parse_classic_db_sources(SAMPLE_SQL)
    record = next(r for r in items[744] if r.kind == "quest_reward")
    assert record.name == "Sweet Amber"
    assert record.quest.quest_id == 53
    assert record.quest.min_level == 40
    assert record.quest.level == 44
    assert record.quest.faction == "alliance"  # RequiredRaces 77 = Human|Dwarf|NightElf|Gnome


def test_quest_reward_horde_only_faction_and_rew_item_id_slot():
    items = cs.parse_classic_db_sources(SAMPLE_SQL)
    record = next(r for r in items[159] if r.kind == "quest_reward")
    assert record.quest.quest_id == 8
    assert record.quest.faction == "horde"  # RequiredRaces 178 = Orc|Undead|Tauren|Troll


def test_excluded_reference_ids_catches_a_generic_pool_by_high_fan_out_even_without_the_marker():
    """id 60446 (the real dump's own "16 Slot Bag - NPC Levels: 48+",
    shared by 777 creatures with no "World Drop" in its comment at all)
    is why `_SHARED_REFERENCE_MAX_USERS` exists as a second signal
    alongside `_WORLD_DROP_MARKER`."""
    rows = ",".join(f"({100 + n},500,1,0,-500,1,0,'Generic Bag')" for n in range(60))
    loot_columns = (
        "\n  `entry` mediumint,\n  `item` mediumint,\n  `ChanceOrQuestChance` float,\n"
        "  `groupid` tinyint,\n  `mincountOrRef` mediumint,\n  `maxcount` tinyint,\n"
        "  `condition_id` mediumint,\n  `comments` varchar(300)\n"
    )
    empty_tables = "".join(
        f"CREATE TABLE `{table}` ({loot_columns}) ENGINE=MyISAM;\n"
        for table in (
            "gameobject_loot_template", "skinning_loot_template",
            "pickpocketing_loot_template", "fishing_loot_template",
        )
    )
    sql = f"""
CREATE TABLE `creature_loot_template` ({loot_columns}) ENGINE=MyISAM;
INSERT INTO `creature_loot_template` VALUES {rows};
CREATE TABLE `reference_loot_template` ({loot_columns}) ENGINE=MyISAM;
INSERT INTO `reference_loot_template` VALUES (500,9001,1,0,1,1,0,'Bag contents');
{empty_tables}
"""
    assert cs._excluded_reference_ids(sql) == {500}


def test_faction_from_required_races_zero_means_both():
    assert cs._faction_from_required_races(0) == "both"
    assert cs._faction_from_required_races(1) == "alliance"
    assert cs._faction_from_required_races(2) == "horde"
    assert cs._faction_from_required_races(1 | 2) == "both"  # mixed bits, never guessed


def test_write_and_load_classic_sources_round_trip(tmp_path):
    build_dir = tmp_path / "1.60.1.70009"
    items = cs.parse_classic_db_sources(SAMPLE_SQL)
    cs.write_classic_sources(build_dir, items)
    loaded = cs.load_classic_sources(build_dir)
    assert loaded[5404][0].name == "Lady Anacondra"
    assert loaded[744][0].quest.faction == "alliance"


def test_load_classic_sources_with_no_file_warns_and_returns_empty(tmp_path):
    assert cs.load_classic_sources(tmp_path / "1.60.1.70009") == {}


def test_fetch_classic_db_sources_downloads_and_gunzips_and_parses():
    payload = gzip.compress(SAMPLE_SQL.encode("utf-8"))

    def handler(request: httpx.Request) -> httpx.Response:
        assert str(request.url) == cs.SOURCE_URL
        return httpx.Response(200, content=payload)

    client = httpx.Client(transport=httpx.MockTransport(handler))
    items = cs.fetch_classic_db_sources(client=client)
    assert items[5404][0].name == "Lady Anacondra"
