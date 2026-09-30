# data/tests/test_audit_dumpdb.py
"""`pipeline.audit.dumpdb.ClassicDbDump` -- a small synthetic mysqldump
snippet covering `creature_template`/`creature` (boss spawn map),
`item_template` (RequiredReputationFaction/Rank + spellid_1-5) and
`spell_template` (SPELL_EFFECT_CREATE_ITEM's own EffectItemType)."""

from pipeline.audit.dumpdb import SPELL_EFFECT_CREATE_ITEM, ClassicDbDump

SAMPLE_SQL = f"""
CREATE TABLE `creature_template` (
  `Entry` mediumint,
  `Name` char(100)
) ENGINE=MyISAM;
INSERT INTO `creature_template` VALUES (9001,'Test Boss');

CREATE TABLE `creature` (
  `guid` int,
  `id` mediumint,
  `map` smallint
) ENGINE=MyISAM;
INSERT INTO `creature` VALUES (1,9001,43),(2,9001,43),(3,9001,0);

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
INSERT INTO `item_template` VALUES
  (500,87,4,0,0,0,0,0),
  (600,0,0,7001,0,0,0,0);

CREATE TABLE `spell_template` (
  `Id` int,
  `Effect1` int,
  `Effect2` int,
  `Effect3` int,
  `EffectItemType1` int,
  `EffectItemType2` int,
  `EffectItemType3` int
) ENGINE=MyISAM;
INSERT INTO `spell_template` VALUES
  (7001,{SPELL_EFFECT_CREATE_ITEM},0,0,700,0,0),
  (7002,0,0,0,0,0,0);
"""


def _write_dump(tmp_path):
    path = tmp_path / "classicdb.sql"
    path.write_text(SAMPLE_SQL, encoding="utf-8")
    return ClassicDbDump(path)


def test_creature_names_and_spawn_map(tmp_path):
    dump = _write_dump(tmp_path)
    assert dump.creature_names[9001] == "Test Boss"
    # Spawns twice on map 43, once on map 0 -- 43 is the most common.
    assert dump.creature_spawn_map[9001] == 43


def test_item_template_reputation_and_spell_ids(tmp_path):
    dump = _write_dump(tmp_path)
    assert dump.item_template[500]["required_reputation_faction"] == 87
    assert dump.item_template[500]["required_reputation_rank"] == 4
    assert dump.item_template[600]["spell_ids"] == [7001]


def test_created_item_to_spells(tmp_path):
    dump = _write_dump(tmp_path)
    assert dump.created_item_to_spells[700] == [7001]
    assert 0 not in dump.created_item_to_spells


def test_item_template_spell_ids(tmp_path):
    """Every non-zero `spellid_1..5`, keyed by item id -- NOT narrowed
    to equippable rows the way `item_template` (RequiredReputation*) is;
    `pipeline.classicdb_crafted`'s own reverse lookup needs recipe items
    (Item.ClassID 9, never equippable) too."""
    dump = _write_dump(tmp_path)
    assert dump.item_template_spell_ids[500] == []
    assert dump.item_template_spell_ids[600] == [7001]


# --- spell_reagents (data-followups-3 lane, 2026-09-30, item 1) ---

REAGENT_SQL = """
CREATE TABLE `spell_template` (
  `Id` int,
  `Reagent1` int,
  `Reagent2` int,
  `Reagent3` int,
  `Reagent4` int,
  `Reagent5` int,
  `Reagent6` int,
  `Reagent7` int,
  `Reagent8` int,
  `ReagentCount1` int,
  `ReagentCount2` int,
  `ReagentCount3` int,
  `ReagentCount4` int,
  `ReagentCount5` int,
  `ReagentCount6` int,
  `ReagentCount7` int,
  `ReagentCount8` int
) ENGINE=MyISAM;
INSERT INTO `spell_template` VALUES
  (21161,17203,11371,0,0,0,0,0,0,8,20,0,0,0,0,0,0),
  (7002,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0);
"""


def test_spell_reagents_reads_every_nonzero_reagent_and_its_count(tmp_path):
    dump = ClassicDbDump.from_text(REAGENT_SQL)
    assert dump.spell_reagents[21161] == [(17203, 8), (11371, 20)]
    assert 7002 not in dump.spell_reagents  # every reagent slot is zero


# --- equippable_item_template_rows / spell_names (catalogue-universe lane) ---

EQUIPPABLE_SQL = """
CREATE TABLE `item_template` (
  `entry` mediumint,
  `class` tinyint,
  `InventoryType` tinyint
) ENGINE=MyISAM;
INSERT INTO `item_template` VALUES
  (100,2,13),
  (200,4,1),
  (300,4,0),
  (400,9,5);

CREATE TABLE `spell_template` (
  `Id` int,
  `SpellName` text
) ENGINE=MyISAM;
INSERT INTO `spell_template` VALUES (7001,'Increase Spell Dam 29');
"""


def test_equippable_item_template_rows_keeps_only_weapon_or_armor_with_a_real_slot():
    dump = ClassicDbDump.from_text(EQUIPPABLE_SQL)
    ids = {int(row["entry"]) for row in dump.equippable_item_template_rows}
    # 100 is a weapon with a real slot, 200 is armour with a real slot; 300 is
    # armour with InventoryType 0 (not equippable) and 400 is neither class
    # (consumable, ClassID 9) -- both excluded.
    assert ids == {100, 200}


def test_spell_names_reads_the_internal_cmangos_label():
    dump = ClassicDbDump.from_text(EQUIPPABLE_SQL)
    assert dump.spell_names[7001] == "Increase Spell Dam 29"


def test_from_text_builds_a_dump_with_no_backing_file():
    dump = ClassicDbDump.from_text(EQUIPPABLE_SQL)
    assert dump.path is None
    assert len(dump.equippable_item_template_rows) == 2
