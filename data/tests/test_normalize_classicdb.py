# data/tests/test_normalize_classicdb.py
"""normalize_build merges the classic-db item supplement into its own item
outputs -- items.json, items/<class-slug>.json and sets.json -- for exactly
the ids that are missing BOTH from the client's own tables and from
wowhead's Forever gear-planner supplement (catalogue-universe lane,
2026-09-30; see pipeline/classicdb_items.py and pipeline/normalize/
classicdb.py for the defect this closes).
"""

import json
import logging
import shutil
from pathlib import Path

from test_normalize_build import HERE, prepare
from test_normalize_wowhead import FIXTURE as WOWHEAD_FIXTURE

from pipeline import wowhead_items as wh
from pipeline.classicdb_items import ClassicDbItem, write_extract
from pipeline.models import ClassItems, Item, ItemSetRecord
from pipeline.normalize import normalize_build
from pipeline.normalize.classicdb import (
    load_classicdb_supplement,
    merge_class_items,
    merge_items,
    merge_sets,
)
from pipeline.spelltext import SpellText, load_spell_text


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


def _empty_spell_text() -> SpellText:
    return load_spell_text([], [], [], [], extra=None)


# --- Unit tests for the small merge helpers ---------------------------------


def test_load_classicdb_supplement_is_none_and_logs_when_the_build_has_no_extract(
    tmp_path: Path, caplog
) -> None:
    caplog.set_level(logging.INFO)
    build_dir = tmp_path / "1.60.1.70009"
    assert load_classicdb_supplement(build_dir, known_ids=set()) is None
    assert "no extract" in caplog.text


def test_load_classicdb_supplement_excludes_known_ids() -> None:
    build_dir_records = [_item(id=11815), _item(id=647, name="Destiny")]
    import tempfile

    with tempfile.TemporaryDirectory() as tmp:
        build_dir = Path(tmp) / "1.60.1.70009"
        write_extract(build_dir, build_dir_records, {}, source_commit="deadbeef")
        picked, _spells = load_classicdb_supplement(build_dir, known_ids={647})
    assert [item.id for item in picked] == [11815]


def test_merge_items_appends_a_flat_item_per_picked_and_logs_the_count(caplog) -> None:
    caplog.set_level(logging.INFO)
    base = [
        Item(
            id=1,
            name="Placeholder",
            quality=3,
            item_level=10,
            required_level=1,
            class_id=4,
            subclass_id=0,
            inventory_type=1,
        )
    ]
    merged = merge_items(base, [_item(id=11815)])
    assert [i.id for i in merged] == [1, 11815]
    assert "merged 1 supplement items into items.json" in caplog.text


def test_merge_items_replaces_a_client_row_that_has_no_real_inventory_type(caplog) -> None:
    """A client row can name an id without ever giving it real equippable
    data (Orb of Deception, 1973, on build 1.60.1.70009 -- InventoryType 0
    in the hotfix cache); such an id is still pick-able from classic-db
    (`known_ids` in `pipeline.normalize` only counts a client row with a real
    InventoryType as covering the id), so its flat items.json row must be
    replaced, not duplicated."""
    caplog.set_level(logging.INFO)
    stale = Item(
        id=1973,
        name="Orb of Deception",
        quality=3,
        item_level=59,
        required_level=54,
        class_id=4,
        subclass_id=0,
        inventory_type=0,  # the client's own broken row
    )
    merged = merge_items([stale], [_item(id=1973, name="Orb of Deception", inventory_type=12)])
    assert [i.id for i in merged] == [1973]
    assert merged[0].inventory_type == 12


def test_merge_class_items_places_gear_only_where_its_allowable_class_mask_permits(
    caplog,
) -> None:
    caplog.set_level(logging.INFO)
    class_rows = [
        {"ID": "1", "Name_lang": "Warrior"},
        {"ID": "8", "Name_lang": "Mage"},
    ]
    records = [
        ClassItems(build="1.60.1.70009", class_slug="warrior", items=[]),
        ClassItems(build="1.60.1.70009", class_slug="mage", items=[]),
    ]
    picked = [_item(id=11815, allowable_class=-1), _item(id=647, name="Destiny", allowable_class=1)]
    merged = merge_class_items(
        records, picked, class_rows, _empty_spell_text(), fork_icons={}, wowhead_icons={}
    )
    by_slug = {r.class_slug: r for r in merged}
    assert {i.id for i in by_slug["warrior"].items} == {11815, 647}
    assert {i.id for i in by_slug["mage"].items} == {11815}
    assert "merged 3 supplement item placements into items/*.json" in caplog.text
    # Every placed row is tagged classic-db, client_unconfirmed.
    for record in merged:
        for gear in record.items:
            assert gear.stats_source == "classic-db"
            assert gear.client_unconfirmed is True


def test_merge_sets_joins_a_known_set_and_counts_an_unknown_one(caplog) -> None:
    caplog.set_level(logging.INFO)
    sets = [ItemSetRecord(id=200, name="Test Set", item_ids=[1, 2], bonuses=[])]
    joined = _item(id=11815, set_id=200)
    unknown = _item(id=647, name="Destiny", set_id=9999)
    merged = merge_sets(sets, [joined, unknown])
    assert merged[0].item_ids == [1, 2, 11815]
    assert "merged 1 supplement items into existing sets" in caplog.text


# --- Full normalize_build integration --------------------------------------


def _write_item_template_extract(root: Path) -> None:
    """Hand of Justice (real, uncommon-or-better, missing from both the
    client and the wowhead fixture) plus an id (271218) that collides with
    wowhead's own supplement's sword, to prove classic-db never re-adds an
    id wowhead already placed."""
    build_dir = root / "1.0.0.1"
    write_extract(
        build_dir,
        [
            _item(id=11815, allowable_class=-1),
            _item(id=271218, name="Vileblood Scimitar", allowable_class=-1),
        ],
        {},
        source_commit="deadbeef",
    )


def test_normalize_build_merges_the_classicdb_supplement_after_wowhead(
    tmp_path: Path, caplog
) -> None:
    caplog.set_level(logging.INFO)
    root = prepare(tmp_path)
    shutil.copy(WOWHEAD_FIXTURE, root / "1.0.0.1" / "raw" / wh.RAW_FILE)
    _write_item_template_extract(root)

    result = normalize_build("1.0.0.1", root=root, curated_dir=HERE / "fixtures/curated")
    assert result.skipped == ()
    out = result.build_dir

    items_payload = json.loads((out / "items.json").read_text())
    ids = [item["id"] for item in items_payload]
    # Hand of Justice is added once...
    assert ids.count(11815) == 1
    # ...but 271218 is wowhead's own sword id already, so classic-db's copy
    # of it is never appended a second time.
    assert ids.count(271218) == 1

    warrior = json.loads((out / "items" / "warrior.json").read_text())
    hand_of_justice = next(item for item in warrior["items"] if item["id"] == 11815)
    assert hand_of_justice["stats_source"] == "classic-db"
    assert hand_of_justice["required_level_source"] == "classic-db"
    assert hand_of_justice["client_unconfirmed"] is True
    assert hand_of_justice["slot"] == "trinket"


def test_a_build_with_no_classicdb_extract_matches_todays_golden_output(
    tmp_path: Path, caplog
) -> None:
    """No raw/classicdb/item_template.json at all -- the same "soft gap"
    contract as `pipeline.normalize.wowhead.load_supplement`: every output
    is unchanged and today's golden files still match byte for byte."""
    caplog.set_level(logging.INFO)
    root = prepare(tmp_path)
    result = normalize_build("1.0.0.1", root=root, curated_dir=HERE / "fixtures/curated")
    assert result.skipped == ()
    out = result.build_dir
    assert (out / "items.json").read_text() == (HERE / "golden/items.json").read_text()
    assert (out / "items" / "warrior.json").read_text() == (
        HERE / "golden/items_warrior.json"
    ).read_text()
    assert (out / "sets.json").read_text() == (HERE / "golden/sets.json").read_text()
    assert "no extract" in caplog.text
