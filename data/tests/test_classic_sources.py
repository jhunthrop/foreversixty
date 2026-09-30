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
  `RequiredMinRepFaction` smallint,
  `RequiredMinRepValue` mediumint,
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
  (53,40,44,77,0,0,'Sweet Amber',744,0,0,0,0,0,0,0,0,0),
  (8,1,5,178,0,0,'A Rogues Deal',0,0,0,0,0,0,159,0,0,0),
  (8573,60,60,255,609,42000,'Champion\\'s Battlegear',21180,0,0,0,0,0,0,0,0,0);
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


def test_a_world_drop_commented_reference_contributes_nothing_to_the_creature():
    """cmangos' own generic level-gated pool (any qualifying creature can
    drop from it -- see pipeline.classic_sources' own doc for the 300+
    item Onyxia's-Lair-trash-mob regression this prevents): item 8888,
    behind reference 60321, is never attributed to creature 300 itself."""
    items = cs.parse_classic_db_sources(SAMPLE_SQL)
    assert not any(r.kind != "world_drop" for r in items[8888])


def test_a_world_drop_commented_reference_becomes_one_world_drop_record_with_its_own_levels():
    """World-drop-pool lane, 2026-09-29: rather than simply vanishing
    (the old behaviour the test above once pinned in full), item 8888
    gets ONE `world_drop` record, named "World drop" and carrying the
    level range reference 60321's own referencing row states in its
    comments ("(Item Levels: 51-60)" -- no "NPC Levels" text on this one,
    so `_level_range_from_comment` falls back to the item-level figure),
    with the REFERENCE row's own chance (3.0), not the outer creature
    row's (0.02)."""
    items = cs.parse_classic_db_sources(SAMPLE_SQL)
    records = [r for r in items[8888] if r.kind == "world_drop"]
    assert len(records) == 1
    record = records[0]
    assert record.name == "World drop"
    assert record.npc_id is None
    assert record.level_min == 51
    assert record.level_max == 60
    assert record.chance == 3.0


def test_gameobject_drop_reads_the_chest_name_and_spawn_map():
    items = cs.parse_classic_db_sources(SAMPLE_SQL)
    records = items[1951]
    assert len(records) == 1
    assert records[0].kind == "object_drop"
    assert records[0].object_id == 500
    assert records[0].name == "Abercrombie's Crate"
    assert records[0].map_id == 43
    assert records[0].chance == 6.0


#: day3 data-followups-6 lane, 2026-09-30, this lane's brief item (a): a
#: `_GAMEOBJECT_TYPE_CHEST` row (179703, "Cache of the Firelord") whose
#: own `data1` (16719) differs from its own entry -- the real shape
#: `gameobject_loot_template` is keyed by `data1`, not the chest's own
#: entry, the exact defect this lane's report measured on the pinned
#: dump. A second row (400, "Locked Chest") has NO `data1` at all (an
#: ordinary instance object whose own entry IS its own loot key) --
#: `_parse_object_drops` must resolve the first through `chest_owner` and
#: leave the second exactly as before.
_CHEST_OWNER_LOOT_COLUMNS = (
    "\n  `entry` mediumint,\n  `item` mediumint,\n  `ChanceOrQuestChance` float,\n"
    "  `groupid` tinyint,\n  `mincountOrRef` mediumint,\n  `maxcount` tinyint,\n"
    "  `condition_id` mediumint,\n  `comments` varchar(300)\n"
)

_CHEST_OWNER_SQL = f"""
CREATE TABLE `creature_template` (
  `Entry` mediumint,
  `Name` char(100),
  `VendorTemplateId` mediumint
) ENGINE=MyISAM;
CREATE TABLE `creature` (
  `guid` int,
  `id` mediumint,
  `map` smallint
) ENGINE=MyISAM;
CREATE TABLE `creature_loot_template` ({_CHEST_OWNER_LOOT_COLUMNS}) ENGINE=MyISAM;
CREATE TABLE `reference_loot_template` ({_CHEST_OWNER_LOOT_COLUMNS}) ENGINE=MyISAM;
CREATE TABLE `skinning_loot_template` ({_CHEST_OWNER_LOOT_COLUMNS}) ENGINE=MyISAM;
CREATE TABLE `pickpocketing_loot_template` ({_CHEST_OWNER_LOOT_COLUMNS}) ENGINE=MyISAM;
CREATE TABLE `fishing_loot_template` ({_CHEST_OWNER_LOOT_COLUMNS}) ENGINE=MyISAM;

CREATE TABLE `gameobject_template` (
  `entry` mediumint,
  `type` tinyint,
  `name` varchar(100),
  `data1` int
) ENGINE=MyISAM;
INSERT INTO `gameobject_template` VALUES
  (179703,3,'Cache of the Firelord',16719),
  (400,3,'Locked Chest',0);

CREATE TABLE `gameobject` (
  `guid` int,
  `id` mediumint,
  `map` smallint
) ENGINE=MyISAM;
INSERT INTO `gameobject` VALUES (1,179703,409),(2,400,409);

CREATE TABLE `gameobject_loot_template` ({_CHEST_OWNER_LOOT_COLUMNS}) ENGINE=MyISAM;
INSERT INTO `gameobject_loot_template` VALUES
  (16719,18803,0,0,1,1,0,'Cache loot'),
  (400,90050,50,0,1,1,0,'Chest loot');

CREATE TABLE `npc_vendor` (
  `entry` mediumint, `item` mediumint, `maxcount` tinyint, `incrtime` int,
  `slot` tinyint, `condition_id` mediumint, `comments` text
) ENGINE=MyISAM;
CREATE TABLE `npc_vendor_template` (
  `entry` mediumint, `item` mediumint, `maxcount` tinyint, `incrtime` int,
  `slot` tinyint, `condition_id` mediumint, `comments` text
) ENGINE=MyISAM;
CREATE TABLE `conditions` (
  `condition_entry` mediumint, `type` tinyint, `value1` mediumint,
  `value2` mediumint, `value3` mediumint, `value4` mediumint,
  `flags` tinyint, `comments` varchar(500)
) ENGINE=MyISAM;
CREATE TABLE `quest_template` (
  `entry` mediumint, `MinLevel` tinyint, `QuestLevel` smallint,
  `RequiredRaces` smallint, `Title` text,
  `RewChoiceItemId1` mediumint, `RewChoiceItemId2` mediumint,
  `RewChoiceItemId3` mediumint, `RewChoiceItemId4` mediumint,
  `RewChoiceItemId5` mediumint, `RewChoiceItemId6` mediumint,
  `RewItemId1` mediumint, `RewItemId2` mediumint,
  `RewItemId3` mediumint, `RewItemId4` mediumint
) ENGINE=MyISAM;
"""


def test_a_chest_whose_data1_differs_from_its_entry_is_named_and_mapped():
    items = cs.parse_classic_db_sources(_CHEST_OWNER_SQL)
    record = items[18803][0]
    assert record.kind == "object_drop"
    assert record.object_id == 179703
    assert record.name == "Cache of the Firelord"
    assert record.map_id == 409
    assert record.is_reward_chest is True


def test_an_ordinary_chest_with_no_data1_keeps_its_own_entry_and_is_not_a_reward_chest():
    items = cs.parse_classic_db_sources(_CHEST_OWNER_SQL)
    record = items[90050][0]
    assert record.object_id == 400
    assert record.name == "Locked Chest"
    assert record.is_reward_chest is False


def test_chest_owner_by_loot_id_inverts_data1_to_the_owning_entry():
    owner = cs.chest_owner_by_loot_id(_CHEST_OWNER_SQL)
    assert owner[16719] == 179703
    assert 400 not in owner  # data1 unset -- never in the domain at all


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


def test_a_vendor_record_carries_its_npcs_own_spawn_map():
    """pvp-faction lane, 2026-09-29: a vendor record's `map_id` is the
    SAME `creature.map` a creature/skinning/pickpocketing record already
    reads -- `pipeline.loot.pvp_faction.vendor_npc_factions`' own primary
    signal for a rank quartermaster's side. Lady Anacondra's own two
    spawns (guid 1 and 2) are both map 43; Reagent Vendor's one spawn is
    map 0."""
    items = cs.parse_classic_db_sources(SAMPLE_SQL)
    assert next(r for r in items[6260] if r.kind == "vendor").map_id == 43
    assert next(r for r in items[8925] if r.kind == "vendor").map_id == 0


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


def test_quest_reward_reads_required_min_rep_faction_and_value_as_a_standing():
    """rep-gate lane, 2026-09-30, this lane's brief item 1: quest 8573's
    own RequiredMinRepFaction/RequiredMinRepValue (609, 42000) become
    ClassicDbQuestInfo.required_rep_faction/required_rep_standing --
    42000 is the Exalted cutoff, not a raw number."""
    items = cs.parse_classic_db_sources(SAMPLE_SQL)
    record = next(r for r in items[21180] if r.kind == "quest_reward")
    assert record.quest.quest_id == 8573
    assert record.quest.required_rep_faction == 609
    assert record.quest.required_rep_standing == "exalted"


def test_quest_reward_with_no_rep_requirement_leaves_both_fields_none():
    items = cs.parse_classic_db_sources(SAMPLE_SQL)
    record = next(r for r in items[744] if r.kind == "quest_reward")
    assert record.quest.required_rep_faction is None
    assert record.quest.required_rep_standing is None

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


def test_level_range_from_comment_prefers_npc_over_item_levels():
    """The pinned dump states both on the SAME row ("NPC LOOT (Green
    World Drop) - (Item Levels: 20-25) - (NPC Levels: 21-22)") --
    `_level_range_from_comment` reads the NPC figure, closer to this
    lane's brief's own "creature level range" alternative."""
    assert cs._level_range_from_comment(
        "NPC LOOT (Green World Drop) - (Item Levels: 20-25) - (NPC Levels: 21-22)"
    ) == (21, 22)
    assert cs._level_range_from_comment("NPC LOOT (Grey World Drop) - (Item Levels: 51-60)") == (
        51,
        60,
    )
    assert cs._level_range_from_comment("16 Slot Bag") is None


#: `table_columns` (pipeline.sqldump) requires one column per LINE
#: (`^\`name\``, `re.MULTILINE`) -- every synthetic fixture below a table
#: from this string, never the single-line form, for that reason.
_LOOT_COLUMNS = (
    "\n  `entry` mediumint,\n  `item` mediumint,\n  `ChanceOrQuestChance` float,\n"
    "  `groupid` tinyint,\n  `mincountOrRef` mediumint,\n  `maxcount` tinyint,\n"
    "  `condition_id` mediumint,\n  `comments` varchar(300)\n"
)

#: The tables `parse_classic_db_sources` always reads but neither
#: `_world_drop_pool_sql` nor the merge-range test below needs any rows
#: in -- declared empty so the parser does not raise on a missing table.
_EMPTY_SUPPORTING_TABLES = "".join(
    f"CREATE TABLE `{table}` ({_LOOT_COLUMNS}) ENGINE=MyISAM;\n"
    for table in (
        "gameobject_loot_template", "skinning_loot_template",
        "pickpocketing_loot_template", "fishing_loot_template",
    )
) + """
CREATE TABLE `gameobject_template` (
  `entry` mediumint,
  `name` varchar(100)
) ENGINE=MyISAM;
CREATE TABLE `gameobject` (
  `guid` int,
  `id` mediumint,
  `map` smallint
) ENGINE=MyISAM;
CREATE TABLE `npc_vendor` (
  `entry` mediumint,
  `item` mediumint,
  `maxcount` tinyint,
  `incrtime` int,
  `slot` tinyint,
  `condition_id` mediumint,
  `comments` text
) ENGINE=MyISAM;
CREATE TABLE `npc_vendor_template` (
  `entry` mediumint,
  `item` mediumint,
  `maxcount` tinyint,
  `incrtime` int,
  `slot` tinyint,
  `condition_id` mediumint,
  `comments` text
) ENGINE=MyISAM;
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
"""


# quest-gates lane, 2026-09-29: a synthetic dump mirroring the pinned
# dump's own Onyxia/Ossirian reward quests exactly (see this lane's own
# report) -- three real-shaped rows plus two edge cases
# (`_quest_chain_turn_in_items`'s own doc: a negative PrevQuestId is
# cmangos' "any quest in this ExclusiveGroup" marker, not a literal id,
# and a chain must not infinite-loop on a row that points at itself).
_QUEST_CHAIN_SQL = f"""
CREATE TABLE `creature_template` (
  `Entry` mediumint,
  `Name` char(100),
  `VendorTemplateId` mediumint
) ENGINE=MyISAM;
CREATE TABLE `creature` (
  `guid` int,
  `id` mediumint,
  `map` smallint
) ENGINE=MyISAM;
CREATE TABLE `creature_loot_template` ({_LOOT_COLUMNS}) ENGINE=MyISAM;
CREATE TABLE `reference_loot_template` ({_LOOT_COLUMNS}) ENGINE=MyISAM;
CREATE TABLE `gameobject_loot_template` ({_LOOT_COLUMNS}) ENGINE=MyISAM;
CREATE TABLE `skinning_loot_template` ({_LOOT_COLUMNS}) ENGINE=MyISAM;
CREATE TABLE `pickpocketing_loot_template` ({_LOOT_COLUMNS}) ENGINE=MyISAM;
CREATE TABLE `fishing_loot_template` ({_LOOT_COLUMNS}) ENGINE=MyISAM;
CREATE TABLE `gameobject_template` (
  `entry` mediumint,
  `name` varchar(100)
) ENGINE=MyISAM;
CREATE TABLE `gameobject` (
  `guid` int,
  `id` mediumint,
  `map` smallint
) ENGINE=MyISAM;
CREATE TABLE `npc_vendor` (
  `entry` mediumint,
  `item` mediumint,
  `maxcount` tinyint,
  `incrtime` int,
  `slot` tinyint,
  `condition_id` mediumint,
  `comments` text
) ENGINE=MyISAM;
CREATE TABLE `npc_vendor_template` (
  `entry` mediumint,
  `item` mediumint,
  `maxcount` tinyint,
  `incrtime` int,
  `slot` tinyint,
  `condition_id` mediumint,
  `comments` text
) ENGINE=MyISAM;
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
CREATE TABLE `quest_template` (
  `entry` mediumint,
  `MinLevel` tinyint,
  `QuestLevel` smallint,
  `RequiredRaces` smallint,
  `SrcItemId` mediumint,
  `ReqItemId1` mediumint,
  `ReqItemId2` mediumint,
  `ReqItemId3` mediumint,
  `ReqItemId4` mediumint,
  `PrevQuestId` mediumint,
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
  (7490,60,60,0,18422,18422,0,0,0,0,'Victory for the Horde',0,0,0,0,0,0,0,0,0,0),
  (7491,60,60,178,0,0,0,0,0,7490,'For All To See',18406,18403,18404,0,0,0,0,0,0,0),
  (8791,60,60,0,21220,21220,0,0,0,0,'The Fall of Ossirian',21504,0,0,0,0,0,0,0,0,0),
  (9000,60,60,0,0,0,0,0,0,-5,'Exclusive Group Quest',777,0,0,0,0,0,0,0,0,0),
  (100,60,60,0,555,0,0,0,0,100,'Self Loop',666,0,0,0,0,0,0,0,0,0);
"""


def test_a_chain_quest_with_no_turn_in_of_its_own_inherits_its_predecessors():
    """"For All To See" (7491) itself states SrcItemId/ReqItemId all
    zero -- its own PrevQuestId (7490, "Victory for the Horde") is what
    actually turns in Head of Onyxia (18422), an Onyxia's Lair boss drop.
    Mirrors the pinned dump's own shape exactly (this lane's report)."""
    items = cs.parse_classic_db_sources(_QUEST_CHAIN_SQL)
    record = next(r for r in items[18406] if r.kind == "quest_reward")
    assert record.quest.quest_id == 7491
    assert record.quest.turn_in_item_ids == [18422]


def test_a_quest_with_its_own_turn_in_item_needs_no_chain_walk():
    """"The Fall of Ossirian" turns in Head of Ossirian (21220) directly
    -- SrcItemId and ReqItemId1 both name it on the SAME row that hands
    out the reward, no PrevQuestId chain involved."""
    items = cs.parse_classic_db_sources(_QUEST_CHAIN_SQL)
    record = next(r for r in items[21504] if r.kind == "quest_reward")
    assert record.quest.quest_id == 8791
    assert record.quest.turn_in_item_ids == [21220]


def test_a_negative_prev_quest_id_is_an_exclusive_group_marker_not_a_literal_id():
    """cmangos' own convention: a negative PrevQuestId means "any quest
    in this ExclusiveGroup", not a literal quest id to chase -- chasing
    it as one would misattribute an unrelated quest's own turn-in item."""
    items = cs.parse_classic_db_sources(_QUEST_CHAIN_SQL)
    record = next(r for r in items[777] if r.kind == "quest_reward")
    assert record.quest.turn_in_item_ids == []


def test_a_self_referencing_prev_quest_id_does_not_infinite_loop():
    items = cs.parse_classic_db_sources(_QUEST_CHAIN_SQL)
    record = next(r for r in items[666] if r.kind == "quest_reward")
    assert record.quest.turn_in_item_ids == [555]


def test_quest_turn_in_items_from_classic_sources_extracts_every_quest():
    items = cs.parse_classic_db_sources(_QUEST_CHAIN_SQL)
    turn_ins = cs.quest_turn_in_items_from_classic_sources(items)
    assert turn_ins[7491] == [18422]
    assert turn_ins[8791] == [21220]
    assert turn_ins[9000] == []




def _world_drop_pool_sql(*creature_rows: tuple[int, int], comment: str, ref_id: int = 700) -> str:
    """A minimal dump -- `creature_template`/`creature`/
    `creature_loot_template`/`reference_loot_template` plus every table
    `parse_classic_db_sources` unconditionally reads -- for
    `_world_drop_pools`' own map-spanning and `_ZONE_DROP_MARKER` tests.
    `creature_rows` is one `(npc_id, map_id)` pair per referencing
    creature, all pointing at the SAME `ref_id` with the SAME `comment`.
    """
    template_rows = ",".join(f"({npc},'Trash Mob {npc}',0)" for npc, _ in creature_rows)
    spawn_rows = ",".join(f"({npc},{npc},{map_id})" for npc, map_id in creature_rows)
    loot_rows = ",".join(f"({npc},1,2,0,-{ref_id},1,0,'{comment}')" for npc, _ in creature_rows)
    return f"""
CREATE TABLE `creature_template` (
  `Entry` mediumint,
  `Name` char(100),
  `VendorTemplateId` mediumint
) ENGINE=MyISAM;
INSERT INTO `creature_template` VALUES {template_rows};
CREATE TABLE `creature` (
  `guid` int,
  `id` mediumint,
  `map` smallint
) ENGINE=MyISAM;
INSERT INTO `creature` VALUES {spawn_rows};
CREATE TABLE `creature_loot_template` ({_LOOT_COLUMNS}) ENGINE=MyISAM;
INSERT INTO `creature_loot_template` VALUES {loot_rows};
CREATE TABLE `reference_loot_template` ({_LOOT_COLUMNS}) ENGINE=MyISAM;
INSERT INTO `reference_loot_template` VALUES ({ref_id},9500,10,0,1,1,0,'Shared item');
{_EMPTY_SUPPORTING_TABLES}"""


def test_a_pool_spanning_two_distinct_maps_is_classified_a_world_drop_with_no_marker_needed():
    """The reported defect's own shape: Lambent Scale Cloak's real
    reference ids are each under `_SHARED_REFERENCE_MAX_USERS` and ALSO
    carry `_WORLD_DROP_MARKER` already -- but a pool with NEITHER of the
    two original signals, referenced by creatures on two different
    instance maps, is exactly what a Gnomeregan+Scarlet-Monastery+
    Shadowfang-Keep+Stockade fan-out looks like with the marker/count
    signals set aside. `_world_drop_pools` catches it on map-spanning
    alone."""
    sql = _world_drop_pool_sql(
        (601, 90), (602, 91), comment="Generic Trash Table",
    )
    items = cs.parse_classic_db_sources(sql)
    records = items[9500]
    assert [r.kind for r in records] == ["world_drop"]
    assert all(r.npc_id is None for r in records)


def test_a_zone_drop_labelled_pool_is_exempt_from_the_multi_map_signal():
    """`_ZONE_DROP_MARKER`'s own doc: a stale/duplicate `map` column
    value (the same quirk `instance_zone_by_map` already guards against
    for Onyxia's Lair) must not turn a real, hand-labelled per-instance
    "Zone Drop" table -- Gnomeregan's, Scarlet Monastery's and
    Blackfathom Deeps' own pools on the pinned dump -- into a world drop
    just because it spans two recorded map ids."""
    sql = _world_drop_pool_sql(
        (601, 90), (602, 91), comment="Some Instance - Zone Drop",
    )
    items = cs.parse_classic_db_sources(sql)
    # Not reclassified: both creatures still carry their own ordinary
    # creature_drop record for the item (`pipeline.loot.classicdb` is what
    # later resolves 601/602's own map to a real instance zone -- this
    # module's own job stops at "not a world drop").
    assert {r.kind for r in items[9500]} == {"creature_drop"}
    assert {r.npc_id for r in items[9500]} == {601, 602}


def test_world_drop_records_merges_the_widest_range_across_every_pool_naming_the_same_item():
    """Real dump shape: Lambent Scale Cloak's own item id sits behind 11
    sibling reference ids, each a narrow slice of the same green-quality
    band. `_world_drop_records` merges them into the one range a player
    actually sees rather than one near-duplicate `world_drop` source per
    slice."""
    sql = f"""
CREATE TABLE `creature_template` (
  `Entry` mediumint,
  `Name` char(100),
  `VendorTemplateId` mediumint
) ENGINE=MyISAM;
INSERT INTO `creature_template` VALUES (610,'Narrow Band Mob A',0),(611,'Narrow Band Mob B',0);
CREATE TABLE `creature` (
  `guid` int,
  `id` mediumint,
  `map` smallint
) ENGINE=MyISAM;
INSERT INTO `creature` VALUES (1,610,50),(2,611,50);
CREATE TABLE `creature_loot_template` ({_LOOT_COLUMNS}) ENGINE=MyISAM;
INSERT INTO `creature_loot_template` VALUES
  (610,1,2,0,-701,1,0,'NPC LOOT (Green World Drop) - (NPC Levels: 20-22)'),
  (611,1,2,0,-702,1,0,'NPC LOOT (Green World Drop) - (NPC Levels: 24-26)');
CREATE TABLE `reference_loot_template` ({_LOOT_COLUMNS}) ENGINE=MyISAM;
INSERT INTO `reference_loot_template` VALUES (701,9600,4,0,1,1,0,'x'),(702,9600,6,0,1,1,0,'x');
{_EMPTY_SUPPORTING_TABLES}"""
    items = cs.parse_classic_db_sources(sql)
    records = [r for r in items[9600] if r.kind == "world_drop"]
    assert len(records) == 1
    assert (records[0].level_min, records[0].level_max) == (20, 26)
    assert records[0].chance == 6.0  # the higher of the two bands' own chances


#: drop-sources-2 lane, 2026-09-29: items 7909 (Aquamarine), 7910 (Star
#: Ruby) and 4306 (Silk Cloth) on build 1.60.1.70009's audit -- each
#: creature below names item 9700 on its OWN DIRECT (positive
#: `mincountOrRef`) `creature_loot_template` row, never through a shared
#: `reference_loot_template` id, which is exactly the shape
#: `_world_drop_pools`/`_excluded_reference_ids` cannot see at all (both
#: only ever look at a NEGATIVE `mincountOrRef` row). Six distinct
#: creatures across two maps (`WORLD_DROP_MIN_CREATURES`/
#: `WORLD_DROP_MIN_ZONES` both fire independently), each with its own
#: `creature_template.MinLevel`/`MaxLevel`.
_DIRECT_WORLD_DROP_SQL = f"""
CREATE TABLE `creature_template` (
  `Entry` mediumint,
  `Name` char(100),
  `VendorTemplateId` mediumint,
  `MinLevel` tinyint,
  `MaxLevel` tinyint
) ENGINE=MyISAM;
INSERT INTO `creature_template` VALUES
  (901,'Trash One',0,20,22),(902,'Trash Two',0,21,23),(903,'Trash Three',0,22,24),
  (904,'Trash Four',0,23,25),(905,'Trash Five',0,24,26),(906,'Trash Six',0,25,35);
CREATE TABLE `creature` (
  `guid` int,
  `id` mediumint,
  `map` smallint
) ENGINE=MyISAM;
INSERT INTO `creature` VALUES
  (1,901,1),(2,902,1),(3,903,1),(4,904,2),(5,905,2),(6,906,2);
CREATE TABLE `creature_loot_template` ({_LOOT_COLUMNS}) ENGINE=MyISAM;
INSERT INTO `creature_loot_template` VALUES
  (901,9700,0.2,0,1,1,0,'x'),(902,9700,0.3,0,1,1,0,'x'),(903,9700,0.1,0,1,1,0,'x'),
  (904,9700,0.4,0,1,1,0,'x'),(905,9700,0.05,0,1,1,0,'x'),(906,9700,0.6,0,1,1,0,'x');
CREATE TABLE `reference_loot_template` ({_LOOT_COLUMNS}) ENGINE=MyISAM;
{_EMPTY_SUPPORTING_TABLES}"""


def test_a_direct_creature_loot_row_still_carries_the_npcs_own_level_range():
    """No reference id is involved here at all (every row's own
    `mincountOrRef` is POSITIVE) -- `creature_template.MinLevel`/
    `MaxLevel` reaches the record through the SAME `level_min`/
    `level_max` fields a reference-pool `world_drop` record already
    uses, so `pipeline.loot.classicdb.classicdb_additions` can merge a
    direct-row pool's own range the identical way."""
    items = cs.parse_classic_db_sources(_DIRECT_WORLD_DROP_SQL)
    records = items[9700]
    assert len(records) == 6
    assert all(r.kind == "creature_drop" for r in records)
    assert {(r.npc_id, r.level_min, r.level_max) for r in records} == {
        (901, 20, 22), (902, 21, 23), (903, 22, 24),
        (904, 23, 25), (905, 24, 26), (906, 25, 35),
    }  # fmt: skip


def test_direct_row_world_drop_items_flags_the_item_six_creatures_two_maps_name_directly():
    items = cs.parse_classic_db_sources(_DIRECT_WORLD_DROP_SQL)
    assert cs.direct_row_world_drop_items(items) == {9700}


#: day3 data-followups-6 lane, 2026-09-30, this lane's brief item (d):
#: Molten Core's Cache of the Firelord's own real shape -- exactly two
#: creatures (Flamewaker Healer/Elite), no static spawn map for either,
#: every stated chance under 1%. `WORLD_DROP_CHANCE_SIGNAL_MIN_CREATURES`
#: (3) must keep this OFF the "every known chance is low" signal --
#: two creatures is a rare drop off two specific mobs, never a pool.
_TWO_CREATURE_LOW_CHANCE_SQL = f"""
CREATE TABLE `creature_template` (
  `Entry` mediumint,
  `Name` char(100),
  `VendorTemplateId` mediumint
) ENGINE=MyISAM;
INSERT INTO `creature_template` VALUES (11663,'Flamewaker Healer',0),(11664,'Flamewaker Elite',0);
CREATE TABLE `creature` (
  `guid` int,
  `id` mediumint,
  `map` smallint
) ENGINE=MyISAM;
CREATE TABLE `creature_loot_template` ({_LOOT_COLUMNS}) ENGINE=MyISAM;
INSERT INTO `creature_loot_template` VALUES
  (11663,18803,0.14,0,1,1,0,'x'),(11664,18803,0.54,0,1,1,0,'x');
CREATE TABLE `reference_loot_template` ({_LOOT_COLUMNS}) ENGINE=MyISAM;
{_EMPTY_SUPPORTING_TABLES}"""


def test_direct_row_world_drop_items_does_not_flag_a_two_creature_low_chance_item():
    items = cs.parse_classic_db_sources(_TWO_CREATURE_LOW_CHANCE_SQL)
    assert cs.direct_row_world_drop_items(items) == set()


def test_is_world_drop_pattern_still_flags_three_distinct_low_chance_creatures():
    """The floor sits AT 3, not above it -- three unrelated creatures all
    under the chance ceiling is still a real generic pool (item 3829's own
    shape on the pinned dump: Tar Creeper/Deadwood Gardener/Vilebranch
    Witch Doctor, three unrelated trash mobs, each under 1%)."""
    from pipeline.loot.constants import is_world_drop_pattern

    assert is_world_drop_pattern([(1, None, 0.02), (2, None, 0.06), (3, None, 0.02)]) is True
    assert is_world_drop_pattern([(1, None, 0.14), (2, None, 0.54)]) is False


def test_direct_row_world_drop_items_does_not_flag_an_ordinary_single_boss_drop():
    """The base fixture's own item 5404 (one creature, one map, a real
    40% chance) clears none of the three signals -- confirms this new
    check does not false-positive on the shape every other test in this
    file already exercises."""
    items = cs.parse_classic_db_sources(SAMPLE_SQL)
    assert cs.direct_row_world_drop_items(items) == set()


def test_direct_row_world_drop_items_treats_classic_dbs_own_zero_chance_sentinel_as_unknown():
    """A single creature/single map row with `chance=0.0` must NOT
    trivially satisfy `is_world_drop_pattern`'s "every known chance is
    low" signal on its own -- classic-db's own `ChanceOrQuestChance`
    uses 0 as ITS sentinel for "no chance recorded", never a real 0%
    (`ClassicDbSourceRecord.chance`'s own doc), so treating it as a
    real, known, low chance would misclassify a real single-boss kill as
    a world drop."""
    sql = f"""
CREATE TABLE `creature_template` (
  `Entry` mediumint,
  `Name` char(100),
  `VendorTemplateId` mediumint
) ENGINE=MyISAM;
INSERT INTO `creature_template` VALUES (657,'Defias Pirate',0);
CREATE TABLE `creature` (
  `guid` int,
  `id` mediumint,
  `map` smallint
) ENGINE=MyISAM;
INSERT INTO `creature` VALUES (1,657,36);
CREATE TABLE `creature_loot_template` ({_LOOT_COLUMNS}) ENGINE=MyISAM;
INSERT INTO `creature_loot_template` VALUES (657,9800,0.0,0,1,1,0,'x');
CREATE TABLE `reference_loot_template` ({_LOOT_COLUMNS}) ENGINE=MyISAM;
{_EMPTY_SUPPORTING_TABLES}"""
    items = cs.parse_classic_db_sources(sql)
    assert cs.direct_row_world_drop_items(items) == set()


#: this lane's brief item 3: a gameobject CHEST (type 3) whose own
#: `data1` (16577) is a DIFFERENT numeric id than its own entry (179564,
#: Dire Maul's real "Gordok Tribute" chest) -- one direct row, one
#: NEGATIVE `mincountOrRef` reference row, so `gameobject_chest_loot`
#: proves both the `data1` indirection AND `_expand_loot_template`'s own
#: reference expansion at once.
_CHEST_SQL = """
CREATE TABLE `gameobject_template` (
  `entry` mediumint,
  `type` tinyint,
  `name` varchar(100),
  `data1` int
) ENGINE=MyISAM;
INSERT INTO `gameobject_template` VALUES
  (179564,3,'Gordok Tribute',16577),
  (500,0,'Not A Chest',999);

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
  (16577,0,100,0,-35033,1,0,'Reference');

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


def test_gameobject_chest_loot_resolves_data1_not_the_objects_own_entry():
    chests = cs.gameobject_chest_loot(_CHEST_SQL)
    assert set(chests[179564]) == {8952, 35033}
    assert chests[179564][35033] == 3.0  # the REFERENCE row's own chance


def test_gameobject_chest_loot_skips_a_non_chest_type():
    """`type` 0 (not `_GAMEOBJECT_TYPE_CHEST`) is never read for `data1`,
    even when it happens to have one -- entry 500 names no loot here."""
    chests = cs.gameobject_chest_loot(_CHEST_SQL)
    assert 500 not in chests


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
