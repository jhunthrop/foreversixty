# data/tests/test_classicdb_items.py
"""cmangos/classic-db's 1.12 `item_template` as the item source neither the
client's `ItemSparse` nor wowhead's Forever gear-planner scrape cover
(catalogue-universe lane, 2026-09-30; see pipeline/classicdb_items.py's own
doc for the defect this closes: Hand of Justice, Devilsaur Eye and nine more
real, uncommon-or-better items missing from every hunter/rogue/warrior
catalogue because they never shipped in the client's ItemSparse at all).

classicdb-fidelity lane, 2026-09-30: `_row_sql`/`_SPELL_SCHEMA` grew
`spelltrigger_<n>` and the `spell_template` effect columns
(`_equip_stats`/`effect_text` read classic-db's own structured aura data
now, not just a spell's bare name) -- same minimal-schema convention
test_audit_dumpdb.py's own SAMPLE_SQL already uses, so each test only states
the columns it cares about.
"""

from pipeline.classicdb_items import (
    AURA_MOD_ATTACK_POWER,
    AURA_MOD_CRIT_PERCENT,
    AURA_MOD_STAT,
    AURA_PROC_TRIGGER_SPELL,
    EFFECT_ADD_EXTRA_ATTACKS,
    EFFECT_APPLY_AURA,
    TRIGGER_CHANCE_ON_HIT,
    TRIGGER_ON_EQUIP,
    TRIGGER_ON_USE,
    ClassicDbItem,
    ClassicDbSpell,
    ClassicDbSpellEffect,
    ItemSpellSlot,
    class_allowed,
    classic_honor_ranks,
    effect_text,
    equip_stats,
    extract_records,
    honor_rank_for_name,
    honor_ranks_by_title_for_untitled_items,
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
    "spelltrigger_1",
    "spellid_2",
    "spelltrigger_2",
    "spellid_3",
    "spelltrigger_3",
    "spellid_4",
    "spelltrigger_4",
    "spellid_5",
    "spelltrigger_5",
    "itemset",
    "requiredhonorrank",
]

_SCHEMA = (
    "CREATE TABLE `item_template` (\n"
    + ",\n".join(f"  `{c}` varchar(255)" for c in ITEM_TEMPLATE_COLUMNS)
    + "\n) ENGINE=MyISAM;\n"
)

SPELL_TEMPLATE_COLUMNS = [
    "Id",
    "SpellName",
    "ProcChance",
    "Effect1",
    "Effect2",
    "Effect3",
    "EffectApplyAuraName1",
    "EffectApplyAuraName2",
    "EffectApplyAuraName3",
    "EffectBasePoints1",
    "EffectBasePoints2",
    "EffectBasePoints3",
    "EffectDieSides1",
    "EffectDieSides2",
    "EffectDieSides3",
    "EffectMiscValue1",
    "EffectMiscValue2",
    "EffectMiscValue3",
    "EffectTriggerSpell1",
    "EffectTriggerSpell2",
    "EffectTriggerSpell3",
]

_SPELL_SCHEMA = (
    "CREATE TABLE `spell_template` (\n"
    + ",\n".join(f"  `{c}` varchar(255)" for c in SPELL_TEMPLATE_COLUMNS)
    + "\n) ENGINE=MyISAM;\n"
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


def _spell_sql(spell_id: int, name: str, **overrides: object) -> str:
    values = {c: 0 for c in SPELL_TEMPLATE_COLUMNS}
    values["Id"] = spell_id
    values["SpellName"] = name
    values.update(overrides)
    fields = []
    for column in SPELL_TEMPLATE_COLUMNS:
        value = values[column]
        fields.append(f"'{value}'" if column == "SpellName" else str(value))
    return "(" + ",".join(fields) + ")"


def _sql(rows: list[str], spell_rows: list[str] | None = None) -> str:
    text = _SCHEMA + "INSERT INTO `item_template` VALUES\n" + ",\n".join(rows) + ";\n"
    text += _SPELL_SCHEMA
    spell_values = ",".join(spell_rows or [_spell_sql(0, "")])
    text += "INSERT INTO `spell_template` VALUES " + spell_values + ";\n"
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
            spelltrigger_1=1,
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
    records, spells = extract_records(_sql(rows, spell_rows=[_spell_sql(7001, "Hand of Justice")]))
    assert [r.id for r in records] == [11815]
    hand_of_justice = records[0]
    assert hand_of_justice.name == "Hand of Justice"
    assert hand_of_justice.spells == [
        ItemSpellSlot(spell_id=7001, trigger=1, classic_db_name="Hand of Justice")
    ]
    assert 7001 in spells


def test_extract_records_follows_effect_trigger_spell_one_level_deep():
    """Hand of Justice's own shape: spell 15600 (the equip aura) grants
    spell 15601 (`EFFECT_ADD_EXTRA_ATTACKS`) through `EffectTriggerSpell1` --
    15601 is never a direct `spellid_<n>` on any item, so it only reaches
    the extract through this second hop."""
    rows = [
        _row_sql(
            entry=11815,
            name="Hand of Justice",
            **{"class": 4, "subclass": 0},
            Quality=3,
            InventoryType=12,
            spellid_1=15600,
            spelltrigger_1=1,
        ),
    ]
    spell_rows = [
        _spell_sql(
            15600,
            "Hand of Justice",
            ProcChance=2,
            Effect1=EFFECT_APPLY_AURA,
            EffectApplyAuraName1=AURA_PROC_TRIGGER_SPELL,
            EffectTriggerSpell1=15601,
        ),
        _spell_sql(15601, "Hand of Justice", Effect1=EFFECT_ADD_EXTRA_ATTACKS, EffectBasePoints1=0),
    ]
    _records, spells = extract_records(_sql(rows, spell_rows=spell_rows))
    assert set(spells) == {15600, 15601}


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


def _slot(
    spell_id: int, trigger: int = TRIGGER_ON_EQUIP, classic_db_name: str = ""
) -> ItemSpellSlot:
    return ItemSpellSlot(spell_id=spell_id, trigger=trigger, classic_db_name=classic_db_name)


def _spell(spell_id: int, *effects: ClassicDbSpellEffect, proc_chance: int = 101) -> ClassicDbSpell:
    return ClassicDbSpell(id=spell_id, proc_chance=proc_chance, effects=list(effects))


def _effect(
    aura: int, base_points: int, die_sides: int = 0, misc_value: int = 0, trigger_spell: int = 0
):
    return ClassicDbSpellEffect(
        effect=EFFECT_APPLY_AURA,
        aura=aura,
        base_points=base_points,
        die_sides=die_sides,
        misc_value=misc_value,
        trigger_spell=trigger_spell,
    )


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


# --- class_allowed / planner_stats -------------------------------------------


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


# --- equip_stats: on-equip spell auras -> the planner's stat vocabulary -----


def test_equip_stats_maps_attack_power_and_drops_ranged_attack_power():
    """Devilsaur Eye's own shape, minus the trigger: a single spell granting
    both AURA_MOD_ATTACK_POWER and its ranged mirror -- only the melee key
    is tracked (STAT_BY_MODIFIER_ID's own id-39 precedent)."""
    item = _item(spells=[_slot(24352, trigger=TRIGGER_ON_EQUIP)])
    spells = {24352: _spell(24352, _effect(AURA_MOD_ATTACK_POWER, 19, 1))}
    assert equip_stats(item, spells) == {"attack_power": 20}


def test_equip_stats_ignores_use_and_chance_on_hit_spells():
    """A TRIGGER_ON_USE/TRIGGER_CHANCE_ON_HIT spell's aura is a temporary
    effect, never a flat equipped stat -- Devilsaur Eye's real +150 use
    effect must not inflate `stats`, and Destiny's real strength proc must
    not turn a two-handed sword into a permanent +200 Strength item."""
    item = _item(
        spells=[
            _slot(1, trigger=TRIGGER_ON_USE),
            _slot(2, trigger=TRIGGER_CHANCE_ON_HIT),
        ]
    )
    spells = {
        1: _spell(1, _effect(AURA_MOD_ATTACK_POWER, 149, 1)),
        2: _spell(2, _effect(AURA_MOD_STAT, 199, 1)),
    }
    assert equip_stats(item, spells) == {}


def test_equip_stats_drops_the_proc_trigger_aura_but_keeps_a_sibling_stat():
    """Hand of Justice's own shape: one equip spell is pure proc (dropped),
    the other grants +20 Attack Power (kept)."""
    item = _item(
        spells=[
            _slot(15600, trigger=TRIGGER_ON_EQUIP),
            _slot(9331, trigger=TRIGGER_ON_EQUIP),
        ]
    )
    spells = {
        15600: _spell(15600, _effect(AURA_PROC_TRIGGER_SPELL, 0, trigger_spell=15601)),
        9331: _spell(9331, _effect(AURA_MOD_ATTACK_POWER, 19, 1)),
    }
    assert equip_stats(item, spells) == {"attack_power": 20}


def test_equip_stats_raises_on_an_unclassified_aura():
    from pipeline.classicdb_items import EquipEffectError

    item = _item(spells=[_slot(1, trigger=TRIGGER_ON_EQUIP)])
    spells = {1: _spell(1, _effect(aura=9999, base_points=1))}
    import pytest

    with pytest.raises(EquipEffectError):
        equip_stats(item, spells)


# --- effect_text: classic-db structured render, then client, then bare name -


class _FakeSpellText:
    def __init__(self, texts: dict[int, str]):
        self._texts = texts

    def describe(self, spell_id: int) -> str:
        return self._texts.get(spell_id, "")


def test_effect_text_folds_a_flat_equip_stat_out_of_the_text():
    """Blackhand's Breadth: `AURA_MOD_CRIT_PERCENT` is fully classified and
    folds into the item's `stats` (see `equip_stats`'s own test) -- eleventh
    wow-player sweep, 2026-09-30: repeating "+2% Critical Strike" as prose
    here would double-count a number `stats` already carries, so a flat
    on-equip stat aura with nothing residual contributes no text at all
    (`equip_stat_folds_out_text`), the same rule the client schema's
    `EffectIndex.text` applies."""
    slot = _slot(7598, trigger=TRIGGER_ON_EQUIP, classic_db_name="Increased Critical 2")
    item = _item(spells=[slot])
    spells = {7598: _spell(7598, _effect(AURA_MOD_CRIT_PERCENT, 1, 1))}
    assert effect_text(item, _FakeSpellText({}), spells) == ""


def test_effect_text_prefers_classicdb_structured_render_over_a_name_matched_client_text():
    """Hand of Justice: the client's own spell 15600 has the SAME name as
    classic-db's ("Hand of Justice") yet its own text states the wrong
    percentage (a $h/3 divisor evaluating to 1%, not classic-db's own
    verified 2%) -- this module must not prefer client text just because
    the name matches when it can fully classify the spell itself."""
    item = _item(spells=[_slot(15600, trigger=TRIGGER_ON_EQUIP, classic_db_name="Hand of Justice")])
    spells = {
        15600: _spell(
            15600,
            _effect(AURA_PROC_TRIGGER_SPELL, 0, trigger_spell=15601),
            proc_chance=2,
        ),
        15601: _spell(15601, ClassicDbSpellEffect(
            effect=EFFECT_ADD_EXTRA_ATTACKS, aura=0, base_points=0, die_sides=0,
            misc_value=0, trigger_spell=0,
        )),
    }
    client_spell_names = {15600: "Hand of Justice"}
    spell_text = _FakeSpellText({15600: "1% chance ... Attacks against Dwarves ..."})
    text = effect_text(item, spell_text, spells, client_spell_names)
    assert text == "Chance on hit (2%): Gain 1 extra attack."
    assert "Dwarves" not in text


def test_effect_text_falls_back_to_client_text_when_name_matches_and_aura_is_unclassified():
    slot = _slot(1, trigger=TRIGGER_ON_EQUIP, classic_db_name="Increase Spell Dam 29")
    item = _item(spells=[slot])
    spells = {1: _spell(1, _effect(aura=9999, base_points=28, die_sides=1))}
    client_spell_names = {1: "Increase Spell Dam 29"}
    spell_text = _FakeSpellText({1: "Equip: Increases damage and healing done by up to 29."})
    assert (
        effect_text(item, spell_text, spells, client_spell_names)
        == "Equip: Increases damage and healing done by up to 29."
    )


def test_effect_text_falls_back_to_the_classicdb_name_when_names_disagree():
    """Devilsaur Eye's own defect: classic-db's spell 24352 is "Devilsaur
    Fury", the client's own spell 24352 is a different ability ("Devilsaur
    Glare", a Root effect) -- an unclassified aura here must never borrow
    the mismatched client's text."""
    item = _item(spells=[_slot(24352, trigger=TRIGGER_ON_USE, classic_db_name="Devilsaur Fury")])
    spells = {24352: _spell(24352, _effect(aura=9999, base_points=1))}
    client_spell_names = {24352: "Devilsaur Glare"}
    spell_text = _FakeSpellText({24352: "Your next attack will Root the target."})
    assert effect_text(item, spell_text, spells, client_spell_names) == "Devilsaur Fury"


def test_effect_text_falls_back_to_the_classicdb_name_when_the_spell_is_not_in_the_extract():
    item = _item(spells=[_slot(1, classic_db_name="Increase Spell Dam 29")])
    assert effect_text(item, _FakeSpellText({}), {}) == "Increase Spell Dam 29"


def test_effect_text_keeps_a_residual_proc_line_when_the_same_spell_also_folds_a_stat():
    """A single equip spell can carry both a flat stat aura (folded into
    `stats`, so its own prose is dropped) and a proc aura with no stat
    representation at all (the extra-attack line is the only place that
    residual effect shows up). The fold-out rule (`equip_stat_folds_out_text`)
    operates on the spell's stats as a whole, not by discarding the entire
    spell -- a real proc riding the same row must still surface."""
    item = _item(spells=[_slot(9000, trigger=TRIGGER_ON_EQUIP, classic_db_name="Test Proc")])
    spells = {
        9000: _spell(
            9000,
            _effect(AURA_MOD_ATTACK_POWER, 19, 1),
            _effect(AURA_PROC_TRIGGER_SPELL, 0, trigger_spell=15601),
            proc_chance=2,
        ),
        15601: _spell(
            15601,
            ClassicDbSpellEffect(
                effect=EFFECT_ADD_EXTRA_ATTACKS,
                aura=0,
                base_points=0,
                die_sides=0,
                misc_value=0,
                trigger_spell=0,
            ),
        ),
    }
    assert (
        effect_text(item, _FakeSpellText({}), spells)
        == "Chance on hit (2%): Gain 1 extra attack."
    )


def test_effect_text_joins_multiple_spells_and_drops_empty_ones():
    item = _item(
        spells=[
            _slot(1, classic_db_name=""),
            _slot(2, classic_db_name="Second Effect"),
        ]
    )
    spell_text = _FakeSpellText({1: "First effect."})
    client_spell_names = {1: ""}
    assert effect_text(item, spell_text, {}, client_spell_names) == "First effect. Second Effect"


# --- to_gear_item / to_item --------------------------------------------------


def test_to_gear_item_tags_classicdb_provenance_and_client_unconfirmed():
    item = _item(id=11815, raw_stats={}, spells=[_slot(1, classic_db_name="fallback")])
    gear = to_gear_item(item, _FakeSpellText({}), {}, fork_icons={}, wowhead_icons={})
    assert gear.stats_source == "classic-db"
    assert gear.required_level_source == "classic-db"
    assert gear.client_unconfirmed is True
    assert gear.effect_text == "fallback"
    assert gear.slot == "trinket"


def test_to_gear_item_merges_equip_stats_into_planner_stats():
    """The stat lands in `gear.stats`; `gear.effect_text` stays empty --
    the flat aura is fully accounted for by `stats` alone, so `to_gear_item`
    must not also carry a restating "Equip: +2% Critical Strike." string
    (eleventh wow-player sweep, 2026-09-30; see
    `test_effect_text_folds_a_flat_equip_stat_out_of_the_text`)."""
    item = _item(
        id=13965,
        raw_stats={},
        spells=[_slot(7598, trigger=TRIGGER_ON_EQUIP, classic_db_name="Increased Critical 2")],
    )
    spells = {7598: _spell(7598, _effect(AURA_MOD_CRIT_PERCENT, 1, 1))}
    gear = to_gear_item(item, _FakeSpellText({}), spells, fork_icons={}, wowhead_icons={})
    assert gear.stats == {"crit": 2}
    assert gear.effect_text == ""


def test_to_gear_item_resolves_icon_through_the_fork_wowhead_fallback_chain():
    item = _item(id=11815)
    gear = to_gear_item(
        item, _FakeSpellText({}), {}, fork_icons={11815: "inv_jewelry_ring_03"}, wowhead_icons={}
    )
    assert gear.icon == "inv_jewelry_ring_03"
    assert gear.icon_source == "fork"


def test_to_gear_item_prefers_the_clients_own_item_csv_icon_over_fork_and_wowhead():
    """Player-review sweep 15/16, 2026-09-30: "First Sergeant's Cloak"
    (16340) still resolved the placeholder icon even though the client's
    own Item.csv carries a real IconFileDataID for that id -- ItemSparse
    (this build's stats source) is only ~60% populated, but Item.csv is a
    separate table, populated independently (module doc). The client icon
    wins over BOTH fork and wowhead once it resolves at all."""
    item = _item(id=16340)
    gear = to_gear_item(
        item,
        _FakeSpellText({}),
        {},
        fork_icons={16340: "inv_jewelry_ring_03"},
        wowhead_icons={16340: "inv_misc_questionmark"},
        client_icon_file_ids={16340: 133759},
        icon_names={133759: "inv_misc_cape_07"},
    )
    assert gear.icon == "inv_misc_cape_07"
    assert gear.icon_source == "client"


def test_to_gear_item_falls_back_to_fork_wowhead_when_the_client_has_no_icon_for_the_id():
    """The overwhelming majority of classic-db-sourced ids: no Item.csv row
    at all (or one with IconFileDataID 0). Unchanged from before this
    lane: fork, then wowhead, then the placeholder."""
    item = _item(id=11815)
    gear = to_gear_item(
        item,
        _FakeSpellText({}),
        {},
        fork_icons={11815: "inv_jewelry_ring_03"},
        wowhead_icons={},
        client_icon_file_ids={},
        icon_names={},
    )
    assert gear.icon == "inv_jewelry_ring_03"
    assert gear.icon_source == "fork"


def test_to_gear_item_computes_weapon_damage_two_hand_and_weapon_type():
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
    gear = to_gear_item(sword, _FakeSpellText({}), {}, fork_icons={}, wowhead_icons={})
    assert (gear.damage_min, gear.damage_max, gear.speed) == (112, 168, 2.6)
    assert gear.dps == round((112 + 168) / 2 / 2.6, 2)
    assert gear.two_hand is True
    assert gear.weapon_type == "sword"


def test_to_gear_item_weapon_type_is_none_for_a_non_weapon_row():
    trinket = _item(id=1, class_id=4, subclass_id=0, inventory_type=12)
    gear = to_gear_item(trinket, _FakeSpellText({}), {}, fork_icons={}, wowhead_icons={})
    assert gear.weapon_type is None


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


def test_item_from_row_reads_required_honor_rank():
    """data-followups-10 lane, 2026-09-30, item 1: the original Classic
    honor-rank PvP sets (Lady Palanseer's, Captain Dirgehammer's, ...)
    carry no ItemSparse.RequiredPVPRank at all in this build (that
    hotfix table only ever covers Forever-new PvP items) -- classic-db's
    own `requiredhonorrank` column, verified against the pinned dump
    directly (entry 16465 "Field Marshal's Chain Helm": requiredhonorrank
    17), is this fact's one other primary source."""
    row = _row_sql(
        entry=16465,
        name="Field Marshal Chain Helm",
        **{"class": 4, "subclass": 4},
        Quality=4,
        InventoryType=1,
        AllowableClass=-1,
        AllowableRace=-1,
        ItemLevel=71,
        RequiredLevel=60,
        requiredhonorrank=17,
    )
    records, _ = extract_records(_sql([row]))
    assert [r.id for r in records] == [16465]
    assert records[0].required_honor_rank == 17


def test_item_defaults_required_honor_rank_to_zero():
    """Every OTHER fixture in this file builds a `ClassicDbItem` via
    `_item(...)` without naming this field at all - the model's own
    default (matching `_item_from_row`'s `.get(..., "0")` fallback for a
    hand-written SQL fixture that predates this column) must be 0, "no
    rank requirement", not a validation error."""
    assert _item(id=1).required_honor_rank == 0


def test_classic_honor_ranks_keeps_only_items_with_a_real_rank():
    items = [
        _item(id=16465, required_honor_rank=17),
        _item(id=16437, required_honor_rank=16),
        _item(id=9999, required_honor_rank=0),
    ]
    assert classic_honor_ranks(items) == {16465: 17, 16437: 16}


def test_honor_rank_by_title_table_agrees_with_the_committed_classic_db_extract():
    """HONOR_RANK_BY_TITLE's own doc: every item in this build's committed
    `raw/classicdb/item_template.json` whose name starts with one of these
    titles and carries a nonzero `required_honor_rank` must agree with the
    table -- a real regression guard against the table drifting from the
    primary source it was built from, not just a snapshot of today's data."""
    from pathlib import Path

    from pipeline.classicdb_items import load_extract

    build_dir = Path("builds/1.60.1.70009")
    extract = load_extract(build_dir)
    assert extract is not None, f"no committed classic-db extract at {build_dir}"
    items, _ = extract
    checked = 0
    for item in items:
        if not item.required_honor_rank:
            continue
        rank = honor_rank_for_name(item.name)
        if rank is None:
            continue
        checked += 1
        assert rank == item.required_honor_rank, (item.id, item.name)
    assert checked > 0, "no committed item exercised the title table at all"


def test_honor_rank_for_name_matches_the_longest_title_first():
    # "Field Marshal's" (17) must win over the shorter "Marshal's" (16)
    # for a name that starts with both.
    assert honor_rank_for_name("Field Marshal's Chain Greathelm") == 17
    assert honor_rank_for_name("Marshal's Chain Vices") == 16
    assert honor_rank_for_name("Warlord's Chain Helm") == 17
    assert honor_rank_for_name("General's Chain Vices") == 16
    assert honor_rank_for_name("Just a Regular Item") is None


def test_honor_ranks_by_title_for_untitled_items_fills_only_the_gap():
    """bis-ranker-integrity-16 lane, item 7: the exact live repro -- Forever's
    own re-itemised "Field Marshal's Chain Greathelm" (231562) shares
    neither id nor exact name with classic-db's "Field Marshal's Chain
    Helm" (16465, rank 17), so it must be filled in by TITLE alone. An id
    already in `already_ranked` (classic_honor_ranks/pvp_ranks already
    covered it) is never touched, and a name matching no known title is
    left out entirely."""
    item_rows = [
        {"id": 231562, "name": "Field Marshal's Chain Greathelm"},
        {"id": 231578, "name": "Marshal's Chain Vices"},
        {"id": 16465, "name": "Field Marshal's Chain Helm"},
        {"id": 12345, "name": "Ordinary Leather Belt"},
    ]
    already_ranked = {16465}
    assert honor_ranks_by_title_for_untitled_items(item_rows, already_ranked) == {
        231562: 17,
        231578: 16,
    }
