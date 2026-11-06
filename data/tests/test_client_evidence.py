"""The client's own Item and item-effect tables confirming classic-db-only items
(pipeline/client_evidence.py; client_unconfirmed lane, 2026-11-05)."""

import json
import logging
from pathlib import Path

from conftest import ACTIVE_BUILD
from pipeline.classicdb_items import (
    AURA_MOD_CRIT_PERCENT,
    AURA_PROC_TRIGGER_SPELL,
    EFFECT_APPLY_AURA,
    TRIGGER_CHANCE_ON_HIT,
    TRIGGER_ON_EQUIP,
    TRIGGER_ON_USE,
    ClassicDbItem,
    ClassicDbSpell,
    ClassicDbSpellEffect,
    ItemSpellSlot,
)
from pipeline.client_evidence import (
    CLIENT_EFFECTS_DIFFER,
    CLIENT_LACKS_EFFECT,
    ITEM_SHAPE_DIFFERS,
    NOT_IN_ITEM_TABLE,
    ClientItemRow,
    client_item_rows,
    confirmed_ids,
    unconfirmed_reason,
)

ACTIVE_BUILD_DIR = Path(__file__).resolve().parents[1] / "builds" / ACTIVE_BUILD

#: Phase 1 dungeon staples whose Item row and item effects match classic-db's
#: on 1.60.1.70790: the ranker may pick them.
CONFIRMED_STAPLES = {
    11815: "Hand of Justice",
    11832: "Burst of Knowledge",
    12590: "Felstriker",
    15411: "Mark of Fordring",
    12929: "Emberfury Talisman",
}
#: Staples Forever reworked: the client's item effects name new spells
#: (Briarwood Reed's swamp-doubled spell power, Blackhand's Breadth's mark,
#: the Royal Seal's new effect), so classic-db's 1.12 numbers are not theirs.
REWORKED_STAPLES = {
    12930: "Briarwood Reed",
    13965: "Blackhand's Breadth",
    18467: "Royal Seal of Eldre'Thalas",
}


def _item(**overrides: object) -> ClassicDbItem:
    base = dict(
        id=11815,
        name="Hand of Justice",
        quality=3,
        item_level=58,
        required_level=53,
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
        unique=True,
        spells=[],
    )
    base.update(overrides)
    return ClassicDbItem(**base)


def _slot(spell_id: int, trigger: int) -> ItemSpellSlot:
    return ItemSpellSlot(spell_id=spell_id, trigger=trigger, classic_db_name="")


def _spell(spell_id: int, aura: int) -> ClassicDbSpell:
    effect = ClassicDbSpellEffect(
        effect=EFFECT_APPLY_AURA,
        aura=aura,
        base_points=0,
        die_sides=1,
        misc_value=0,
        trigger_spell=0,
    )
    return ClassicDbSpell(id=spell_id, proc_chance=101, effects=[effect])


def _client(*spell_ids: int, class_id=4, subclass_id=0, inventory_type=12) -> ClientItemRow:
    return ClientItemRow(class_id, subclass_id, inventory_type, frozenset(spell_ids))


# Hand of Justice in 1.12: an on-equip proc (15600) and a flat AP spell (9331).
HOJ = _item(spells=[_slot(15600, TRIGGER_ON_EQUIP), _slot(9331, TRIGGER_ON_EQUIP)])
HOJ_SPELLS = {15600: _spell(15600, AURA_PROC_TRIGGER_SPELL), 9331: _spell(9331, 99)}


def test_an_item_the_client_names_in_classicdbs_shape_is_confirmed():
    """The flat AP spell need not be an item effect: the 1.60 client states
    those as ItemSparse stats."""
    assert unconfirmed_reason(HOJ, _client(15600), HOJ_SPELLS) is None


def test_an_item_missing_from_the_item_table_is_unconfirmed():
    assert unconfirmed_reason(HOJ, None, HOJ_SPELLS) == NOT_IN_ITEM_TABLE


def test_another_class_subclass_or_inventory_type_is_unconfirmed():
    """Orb of Deception is a toy in Forever (inventory type 0); Witchblade
    moved from one-hand (13) to main hand (21)."""
    assert unconfirmed_reason(HOJ, _client(15600, inventory_type=0), HOJ_SPELLS) == (
        ITEM_SHAPE_DIFFERS
    )
    assert unconfirmed_reason(HOJ, _client(15600, subclass_id=1), HOJ_SPELLS) == ITEM_SHAPE_DIFFERS


def test_a_forever_spell_on_the_item_is_a_rework():
    """Briarwood Reed's client effects are 1301707/1318326/1318327, not 1.12's 13881."""
    reed = _item(id=12930, spells=[_slot(13881, TRIGGER_ON_EQUIP)])
    assert unconfirmed_reason(reed, _client(1318326), {}) == CLIENT_EFFECTS_DIFFER


def test_a_dropped_proc_or_use_spell_is_unconfirmed():
    assert unconfirmed_reason(HOJ, _client(), HOJ_SPELLS) == CLIENT_LACKS_EFFECT
    use = _item(spells=[_slot(15646, TRIGGER_ON_USE)])
    assert unconfirmed_reason(use, _client(), {}) == CLIENT_LACKS_EFFECT
    on_hit = _item(spells=[_slot(16551, TRIGGER_CHANCE_ON_HIT)])
    assert unconfirmed_reason(on_hit, _client(16551), {}) is None


def test_a_flat_equip_stat_spell_may_be_absent():
    """Mark of Fordring's 1% crit and 26 AP are ItemSparse stats on the client."""
    mark = _item(id=15411, spells=[_slot(7597, TRIGGER_ON_EQUIP)])
    assert unconfirmed_reason(mark, _client(), {7597: _spell(7597, AURA_MOD_CRIT_PERCENT)}) is None


def test_client_item_rows_reads_the_item_table_and_every_trigger_of_its_effects():
    items = [
        {"ID": "11815", "ClassID": "4", "SubclassID": "0", "InventoryType": "12"},
        {"ID": "12930", "ClassID": "4", "SubclassID": "0", "InventoryType": "12"},
    ]
    effects = [
        {"ID": "99238", "TriggerType": "1", "SpellID": "15600"},
        {"ID": "5", "TriggerType": "0", "SpellID": "1318326"},
    ]
    links = [
        {"ID": "1", "ItemEffectID": "99238", "ItemID": "11815"},
        {"ID": "2", "ItemEffectID": "5", "ItemID": "12930"},
        {"ID": "3", "ItemEffectID": "5", "ItemID": "999"},
    ]
    rows = client_item_rows(items, effects, links)
    assert rows == {
        11815: _client(15600),
        12930: _client(1318326),
    }


def test_confirmed_ids_logs_each_reason(caplog):
    caplog.set_level(logging.INFO)
    reed = _item(id=12930, spells=[_slot(13881, TRIGGER_ON_EQUIP)])
    got = confirmed_ids([HOJ, reed], {11815: _client(15600), 12930: _client(1318326)}, HOJ_SPELLS)
    assert got == {11815}
    assert "the client confirms 1" in caplog.text
    assert f"1 stay client_unconfirmed: {CLIENT_EFFECTS_DIFFER}" in caplog.text


def _active_rows() -> dict[int, list[dict]]:
    rows: dict[int, list[dict]] = {}
    for path in [ACTIVE_BUILD_DIR / "items.json", *(ACTIVE_BUILD_DIR / "items").glob("*.json")]:
        payload = json.loads(path.read_text(encoding="utf-8"))
        for row in payload if isinstance(payload, list) else payload["items"]:
            rows.setdefault(row["id"], []).append(row)
    return rows


def test_the_active_build_ranks_the_confirmed_phase_1_staples():
    """Every copy of each staple, flat and per class, is pickable. A later
    client that states the item outright (stats_source None) passes too."""
    rows = _active_rows()
    for item_id, name in CONFIRMED_STAPLES.items():
        assert item_id in rows, name
        for row in rows[item_id]:
            assert row["client_unconfirmed"] is False, (item_id, name)


def test_the_active_build_never_ranks_a_reworked_staple_on_its_1_12_numbers():
    """A reworked staple is either still unconfirmed or the client's own row."""
    rows = _active_rows()
    for item_id, name in REWORKED_STAPLES.items():
        for row in rows.get(item_id, []):
            assert row["client_unconfirmed"] or row["stats_source"] != "classic-db", (item_id, name)
