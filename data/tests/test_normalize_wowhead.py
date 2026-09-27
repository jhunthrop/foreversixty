"""normalize_build merges wowhead's Forever gear-planner supplement into its own
item outputs -- items.json, items/<class-slug>.json and sets.json -- for exactly the
ids the client's own tables lack (see
docs/superpowers/specs/2026-09-27-wowhead-item-supplement-design.md and
pipeline/normalize/wowhead.py).

tests/fixtures/wowhead-gear-planner.js (also test_wowhead_items.py's fixture) carries
six items: a cloak the client lacks (279865), a one-hand sword restricted to warrior/
hunter/rogue by its class mask, with damage and an untracked "hastertng" stat
(271218), a quality-1 ring (264908), a belt (11726), a QA-named epic (777777) and a
level-70 chest (888888). tests/fixtures/ChrClasses.csv (shared with
test_normalize_build.py) names Warrior, Paladin and Mage -- enough to prove the sword
lands only where its mask allows and the cloak lands wherever cloth is worn.
"""

import json
import logging
import shutil
from pathlib import Path

from test_normalize_build import HERE, prepare

from pipeline import wowhead_items as wh
from pipeline.models import ClassItems, Item, ItemSetRecord
from pipeline.normalize import normalize_build
from pipeline.normalize.wowhead import load_supplement, merge_class_items, merge_items, merge_sets

FIXTURE = HERE / "fixtures" / "wowhead-gear-planner.js"
#: The two ids test_wowhead_items.py already proved supplement() excludes: the belt
#: (already the client's, in this module's own tests) and the quality-1 ring.
CLIENT_IDS = {11726, 264908}


def _picked() -> list[wh.WowheadItem]:
    return wh.supplement(wh.load_items(FIXTURE), CLIENT_IDS)


def _add_client_belt(root: Path) -> None:
    """Give the "1.0.0.1" fixture build a real ItemSparse/Item row for id 11726 --
    the same id the wowhead fixture also carries, as "a belt the client does
    carry" -- so a full normalize_build run proves `supplement`'s client_ids
    filter reaches all the way down, not just the isolated tests in
    test_wowhead_items.py that pass client_ids by hand."""
    raw = root / "1.0.0.1" / "raw"
    sparse = raw / "ItemSparse.csv"
    sparse.write_text(
        sparse.read_text()
        + "11726,Gladiator's Leather Belt,3,50,45,6,0,0,-1,0,0,0,0,0,0,0,7,-1,10,0,0,0,0,0\n",
        encoding="utf-8",
    )
    item = raw / "Item.csv"
    item.write_text(item.read_text() + "11726,4,2,0\n", encoding="utf-8")


# --- Unit tests for the small merge helpers -----------------------------------


def test_load_supplement_is_none_and_logs_when_the_build_has_no_payload(
    tmp_path: Path, caplog
) -> None:
    caplog.set_level(logging.INFO)
    build_dir = tmp_path / "1.60.1.70009"
    assert load_supplement(build_dir, client_ids=set()) is None
    assert "no payload" in caplog.text


def test_load_supplement_wraps_the_planner_gates(tmp_path: Path) -> None:
    build_dir = tmp_path / "1.60.1.70009"
    raw = build_dir / "raw"
    raw.mkdir(parents=True)
    shutil.copy(FIXTURE, raw / wh.RAW_FILE)
    picked = load_supplement(build_dir, client_ids=CLIENT_IDS)
    assert [item.id for item in picked] == [271218, 279865]


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
    merged = merge_items(base, _picked())
    assert [i.id for i in merged] == [1, 271218, 279865]
    assert "merged 2 supplement items into items.json" in caplog.text


def test_merge_class_items_places_gear_only_where_its_class_may_equip_it(caplog) -> None:
    caplog.set_level(logging.INFO)
    class_rows = [
        {"ID": "1", "Name_lang": "Warrior"},
        {"ID": "8", "Name_lang": "Mage"},
    ]
    records = [
        ClassItems(build="1.60.1.70009", class_slug="warrior", items=[]),
        ClassItems(build="1.60.1.70009", class_slug="mage", items=[]),
    ]
    merged = merge_class_items(records, _picked(), class_rows)
    by_slug = {r.class_slug: r for r in merged}
    assert {i.id for i in by_slug["warrior"].items} == {271218, 279865}
    assert {i.id for i in by_slug["mage"].items} == {279865}
    # The sword's own class mask carries a stat wowhead states that the planner does
    # not track (hastertng); it is only placed in warrior's file, so the untracked
    # counter fires exactly once across the whole merge.
    assert "merged 3 supplement item placements into items/*.json" in caplog.text
    assert "hastertng" in caplog.text


def test_merge_class_items_leaves_a_class_absent_from_class_rows_unplaced() -> None:
    """A class_slug on a ClassItems record with no matching row in class_rows (should
    not happen in practice, since both come from the same ChrClasses table) gets no
    supplement gear rather than raising."""
    records = [ClassItems(build="b", class_slug="ghost", items=[])]
    merged = merge_class_items(records, _picked(), class_rows=[])
    assert merged[0].items == []


def test_merge_sets_joins_a_known_set_and_counts_an_unknown_one(caplog) -> None:
    caplog.set_level(logging.INFO)
    sets = [ItemSetRecord(id=200, name="Test Set", item_ids=[1, 2], bonuses=[])]
    sword = next(i for i in wh.load_items(FIXTURE) if i.id == 271218)
    joined = sword.model_copy(update={"set_id": 200})
    unknown = sword.model_copy(update={"id": 999, "set_id": 9999})
    merged = merge_sets(sets, [joined, unknown])
    assert merged[0].item_ids == [1, 2, 271218]
    assert "merged 1 supplement items into existing sets" in caplog.text
    assert "9999" in caplog.text


def test_merge_sets_is_idempotent_on_an_item_the_set_already_carries() -> None:
    sets = [ItemSetRecord(id=200, name="Test Set", item_ids=[271218], bonuses=[])]
    sword = next(i for i in wh.load_items(FIXTURE) if i.id == 271218).model_copy(
        update={"set_id": 200}
    )
    merged = merge_sets(sets, [sword])
    assert merged[0].item_ids == [271218]


def test_merge_sets_leaves_a_set_with_no_supplement_additions_untouched() -> None:
    sets = [ItemSetRecord(id=200, name="Test Set", item_ids=[1, 2], bonuses=[])]
    merged = merge_sets(sets, [])
    assert merged == sets


# --- Full normalize_build integration tests -----------------------------------


def test_a_build_with_no_wowhead_payload_matches_todays_golden_output(
    tmp_path: Path, caplog
) -> None:
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
    assert "no payload" in caplog.text


def test_normalize_build_merges_the_wowhead_supplement(tmp_path: Path, caplog) -> None:
    caplog.set_level(logging.INFO)
    root = prepare(tmp_path)
    _add_client_belt(root)
    shutil.copy(FIXTURE, root / "1.0.0.1" / "raw" / wh.RAW_FILE)

    result = normalize_build("1.0.0.1", root=root, curated_dir=HERE / "fixtures/curated")
    assert result.skipped == ()
    out = result.build_dir

    items_payload = json.loads((out / "items.json").read_text())
    items = {item["id"] for item in items_payload}
    # Gains the supplement ids that pass the planner's gates...
    assert {271218, 279865} <= items
    # ...not the ones the planner gates filter out...
    assert not {264908, 777777, 888888} & items
    # ...and not a second copy of an id the client already carries.
    assert [item["id"] for item in items_payload].count(11726) == 1

    warrior = json.loads((out / "items" / "warrior.json").read_text())
    mage = json.loads((out / "items" / "mage.json").read_text())
    warrior_ids = {item["id"] for item in warrior["items"]}
    mage_ids = {item["id"] for item in mage["items"]}
    assert {271218, 279865} <= warrior_ids  # the sword and the cloak
    assert mage_ids & {271218, 279865} == {279865}  # only the cloak
    assert not {264908, 777777, 888888} & (warrior_ids | mage_ids)

    sets = json.loads((out / "sets.json").read_text())
    # The sword's itemset (9001) names no set the client fixture has; it is
    # counted and skipped, not silently created.
    assert all(s["id"] != 9001 for s in sets)

    assert "merged 2 supplement items into items.json" in caplog.text
    # Warrior gets the sword and the cloak, mage and paladin each get the cloak only
    # (paladin's own class mask does not carry the sword's bit; see the sword's
    # classMask in the fixture docstring).
    assert "merged 4 supplement item placements into items/*.json" in caplog.text
    assert "hastertng" in caplog.text
    assert "merged 0 supplement items into existing sets" in caplog.text
    assert "9001" in caplog.text


def test_normalize_build_joins_a_supplement_item_into_a_set_the_client_has(
    tmp_path: Path,
) -> None:
    """End to end: a supplement item whose set id names an existing ItemSetRecord
    (fixture set 209, "Battlegear of Might") joins its item_ids, sorted and deduped."""
    root = prepare(tmp_path)
    raw = root / "1.0.0.1" / "raw"
    payload = FIXTURE.read_text(encoding="utf-8").replace('"itemset":9001', '"itemset":209')
    (raw / wh.RAW_FILE).write_text(payload, encoding="utf-8")

    result = normalize_build("1.0.0.1", root=root, curated_dir=HERE / "fixtures/curated")
    sets = json.loads((result.build_dir / "sets.json").read_text())
    battlegear = next(s for s in sets if s["id"] == 209)
    assert battlegear["item_ids"] == [16866, 16867, 271218]
