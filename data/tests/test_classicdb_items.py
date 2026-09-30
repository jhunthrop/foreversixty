# data/tests/test_classicdb_items.py
"""cmangos/classic-db's 1.12 `item_template` as the item source neither the
client's `ItemSparse` nor wowhead's Forever gear-planner scrape cover
(catalogue-universe lane, 2026-09-30; see pipeline/classicdb_items.py's own
doc for the defect this closes: Hand of Justice, Devilsaur Eye and nine more
real, uncommon-or-better items missing from every hunter/rogue/warrior
catalogue because they never shipped in the client's ItemSparse at all).

`_row_sql` builds one `item_template` row from a small, minimal schema (not
the full ~90-column production table -- same convention
test_audit_dumpdb.py's own SAMPLE_SQL already uses) so each test only states
the columns it cares about.
"""

from pipeline.classicdb_items import (
    ClassicDbItem,
    class_allowed,
    effect_text,
    extract_records,
    is_gm_class_mask,
    planner_stats,
    supplement,
    to_gear_item,
    to_item,
)

ITEM_TEMPLATE_COLUMNS = [
    "entry",
    "class",
    "subclass",
    "name",
    "Quality",
    "InventoryType",
    "AllowableClass",
    "AllowableRace",
    "ItemLevel",
    "RequiredLevel",
    "maxcount",
    "stat_type1",
    "stat_value1",
    "stat_type2",
    "stat_value2",
    "stat_type3",
    "stat_value3",
    "stat_type4",
    "stat_value4",
    "stat_type5",
    "stat_value5",
    "stat_type6",
    "stat_value6",
    "stat_type7",
    "stat_value7",
    "stat_type8",
    "stat_value8",
    "stat_type9",
    "stat_value9",
    "stat_type10",
    "stat_value10",
    "dmg_min1",
    "dmg_max1",
    "armor",
    "fire_res",
    "nature_res",
    "frost_res",
    "shadow_res",
    "arcane_res",
    "delay",
    "spellid_1",
    "spellid_2",
    "spellid_3",
    "spellid_4",
    "spellid_5",
    "itemset",
]

_SCHEMA = (
    "CREATE TABLE `item_template` (\n"
    + ",\n".join(f"  `{c}` varchar(255)" for c in ITEM_TEMPLATE_COLUMNS)
    + "\n) ENGINE=MyISAM;\n"
)

_SPELL_SCHEMA = (
    "CREATE TABLE `spell_template` (\n  `Id` int,\n  `SpellName` text\n) ENGINE=MyISAM;\n"
)


def _row_sql(**overrides: object) -> str:
    values = {c: 0 for c in ITEM_TEMPLATE_COLUMNS}
    values["name"] = "Test Item"
    values["maxcount"] = 0
    values.update(overrides)
    fields = []
    for column in ITEM_TEMPLATE_COLUMNS:
        value = values[column]
        fields.append(f"'{value}'" if column == "name" else str(value))
    return "(" + ",".join(fields) + ")"


def _sql(rows: list[str], spell_rows: list[str] | None = None) -> str:
    text = _SCHEMA + "INSERT INTO `item_template` VALUES\n" + ",\n".join(rows) + ";\n"
    text += _SPELL_SCHEMA
    text += "INSERT INTO `spell_template` VALUES " + ",".join(spell_rows or ["(0,'')"]) + ";\n"
    return text


# --- is_gm_class_mask -----------------------------------------------------


def test_is_gm_class_mask_flags_only_bits_above_the_real_class_range():
    # -1 ("any class"), the full real-class mask, and two harmless slack
    # bits (Death Knight/Monk) seen on confirmed real items are all fine.
    assert is_gm_class_mask(-1) is False
    assert is_gm_class_mask(1503) is False
    assert is_gm_class_mask(1535) is False  # 1503 | bit5 (Stoneslayer, entry 9418)
    assert is_gm_class_mask(2047) is False  # 1503 | bit5 | bit9 (Dented Buckler)
    # 31233 = 1503-ish low bits plus bits 9/11/12/13/14 -- no real or slack
    # class enumeration reaches that high ("90 Epic Warrior Neck", a GM row).
    assert is_gm_class_mask(31233) is True


# --- extract_records: the raw-universe reader + planner-quality narrowing ---


def test_extract_records_keeps_only_equippable_planner_quality_rows():
    rows = [
        # A real uncommon trinket with no stats, only an on-equip spell --
        # exactly Hand of Justice's own shape (class 4, no stat_type/value,
        # one spellid).
        _row_sql(
            entry=11815,
            name="Hand of Justice",
            **{"class": 4, "subclass": 0},
            Quality=3,
            InventoryType=12,
            AllowableClass=-1,
            AllowableRace=-1,
            ItemLevel=58,
            RequiredLevel=53,
            spellid_1=7001,
        ),
        # Poor quality -- PLANNER_QUALITIES excludes it before it ever
        # reaches supplement()'s own gates.
        _row_sql(
            entry=1, name="Poor Item", **{"class": 4, "subclass": 0}, Quality=0, InventoryType=1
        ),
        # Not equippable at all (InventoryType 0) -- dropped at the reader.
        _row_sql(
            entry=2, name="Not Gear", **{"class": 4, "subclass": 0}, Quality=3, InventoryType=0
        ),
    ]
    records = extract_records(_sql(rows, spell_rows=["(7001,'Hand of Justice')"]))
    assert [r.id for r in records] == [11815]
    hand_of_justice = records[0]
    assert hand_of_justice.name == "Hand of Justice"
    assert hand_of_justice.spells == [(7001, "Hand of Justice")]


# --- supplement / is_planner_gear ------------------------------------------


def _item(**overrides: object) -> ClassicDbItem:
    base = dict(
        id=1,
        name="Test Item",
        quality=3,
        item_level=50,
        required_level=45,
        class_id=4,
        subclass_id=0,
        inventory_type=12,
        allowable_class=-1,
        allowable_race=-1,
        armor=0,
        raw_stats={},
        resistances={},
        damage_min=0,
        damage_max=0,
        delay=0,
        set_id=None,
        unique=False,
        spells=[],
    )
    base.update(overrides)
    return ClassicDbItem(**base)


def test_supplement_excludes_known_ids_junk_names_and_gm_masks():
    items = [
        _item(id=1, name="Hand of Justice"),
        _item(id=2, name="Deprecated Amulet"),  # junk-named
        _item(id=3, name="90 Epic Warrior Neck", allowable_class=31233),  # GM mask
        _item(id=4, name="Devilsaur Eye", required_level=61),  # over the level-60 cap
        _item(id=5, name="Already Known"),  # excluded via known_ids, not its own gates
    ]
    picked = supplement(items, known_ids={5})
    assert [item.id for item in picked] == [1]


# --- class_allowed / planner_stats / effect_text ----------------------------


def test_class_allowed_respects_allowable_class_mask_and_proficiency():
    ring = _item(class_id=4, subclass_id=0, allowable_class=1)  # warrior only
    assert class_allowed(ring, class_id=1) is True  # warrior
    assert class_allowed(ring, class_id=8) is False  # mage: mask excludes it
    idol = _item(class_id=4, subclass_id=8, allowable_class=-1)  # druid relic
    assert class_allowed(idol, class_id=11) is True  # druid
    assert class_allowed(idol, class_id=1) is False  # warrior can't equip an idol


def test_planner_stats_maps_raw_ids_and_folds_in_resistances():
    item = _item(raw_stats={4: 10, 7: 15}, resistances={"fire_res": 5})
    assert planner_stats(item) == {"strength": 10, "stamina": 15, "fire_res": 5}


class _FakeSpellText:
    def __init__(self, texts: dict[int, str]):
        self._texts = texts

    def describe(self, spell_id: int) -> str:
        return self._texts.get(spell_id, "")


def test_effect_text_prefers_the_client_description_over_the_classicdb_name():
    item = _item(spells=[(1, "Increase Spell Dam 29")])
    spell_text = _FakeSpellText({1: "Equip: Increases damage and healing done by up to 29."})
    assert effect_text(item, spell_text) == "Equip: Increases damage and healing done by up to 29."


def test_effect_text_falls_back_to_the_classicdb_name_when_the_client_lacks_the_spell():
    item = _item(spells=[(1, "Increase Spell Dam 29")])
    assert effect_text(item, _FakeSpellText({})) == "Increase Spell Dam 29"


def test_effect_text_joins_multiple_spells_and_drops_empty_ones():
    item = _item(spells=[(1, ""), (2, "Second Effect")])
    spell_text = _FakeSpellText({1: "First effect."})
    assert effect_text(item, spell_text) == "First effect. Second Effect"


# --- to_gear_item / to_item --------------------------------------------------


def test_to_gear_item_tags_classicdb_provenance_and_client_unconfirmed():
    item = _item(id=11815, raw_stats={}, spells=[(1, "fallback")])
    spell_text = _FakeSpellText({1: "Equip: does a thing."})
    gear = to_gear_item(item, spell_text, fork_icons={}, wowhead_icons={})
    assert gear.stats_source == "classic-db"
    assert gear.required_level_source == "classic-db"
    assert gear.client_unconfirmed is True
    assert gear.effect_text == "Equip: does a thing."
    assert gear.slot == "trinket"


def test_to_gear_item_resolves_icon_through_the_fork_wowhead_fallback_chain():
    item = _item(id=11815)
    gear = to_gear_item(
        item, _FakeSpellText({}), fork_icons={11815: "inv_jewelry_ring_03"}, wowhead_icons={}
    )
    assert gear.icon == "inv_jewelry_ring_03"


def test_to_gear_item_computes_weapon_damage_and_two_hand():
    sword = _item(
        id=647,
        name="Destiny",
        class_id=2,
        subclass_id=8,
        inventory_type=17,  # two-hand
        damage_min=112,
        damage_max=168,
        delay=2600,
    )
    gear = to_gear_item(sword, _FakeSpellText({}), fork_icons={}, wowhead_icons={})
    assert (gear.damage_min, gear.damage_max, gear.speed) == (112, 168, 2.6)
    assert gear.dps == round((112 + 168) / 2 / 2.6, 2)
    assert gear.two_hand is True


def test_to_item_carries_the_flat_entity_fields():
    item = _item(id=11815, name="Hand of Justice", class_id=4, subclass_id=0, inventory_type=12)
    flat = to_item(item)
    assert (flat.id, flat.name, flat.class_id, flat.inventory_type) == (
        11815,
        "Hand of Justice",
        4,
        12,
    )


def test_to_item_carries_the_same_provenance_columns_the_per_class_row_gets():
    """loot-parity-2 lane, 2026-09-30: the flat catalogue's own row for a
    classic-db-only item now carries the same three provenance columns
    `to_gear_item`'s per-class row already did -- so items.json ALONE (not
    just items/<class>.json) can tell a 1.12 row from a client one."""
    item = _item(id=11815, name="Hand of Justice", class_id=4, subclass_id=0, inventory_type=12)
    flat = to_item(item)
    assert flat.required_level_source == "classic-db"
    assert flat.stats_source == "classic-db"
    assert flat.client_unconfirmed is True
