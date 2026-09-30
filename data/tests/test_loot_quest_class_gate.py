# data/tests/test_loot_quest_class_gate.py
"""Eleventh wow-player sweep finding, day3 data-followups-7 lane,
2026-09-30: Fire Ruby (item 20036) is published as a trinket pick or
alternative for hunter-beast-mastery, shaman-elemental and paladin-
retribution alike, but its ONLY source, quest 8253 "Destroy Morphaz", is
mage-only. `quest_only_class_exclusions`/`apply_quest_class_gate_to_
class_items` narrow `items/<class>.json` to match.
"""

import json

from pipeline.loot.quest_class_gate import (
    apply_quest_class_gate_to_class_items,
    quest_only_class_exclusions,
)
from pipeline.models import ClassItems, GearItem, LootFile, LootSource, QuestSource

FIRE_RUBY = 20036
OTHER_ITEM = 30000


def _gear_item(item_id: int) -> GearItem:
    return GearItem(
        id=item_id, name=f"Item {item_id}", icon="inv_misc_gem_01", slot="trinket1",
        quality=3, required_level=52, required_level_source="client", item_level=52,
        armor=0, stats={}, set_id=None, unique=False,
    )


def _quest_only_document(classes: list[str] | None, *, also_dropped: bool = False) -> LootFile:
    sources = [LootSource(id="quest", kind="quest", name="Quests", items=[FIRE_RUBY])]
    if also_dropped:
        sources.append(LootSource(id="zone:1", kind="zone", name="Somewhere", items=[FIRE_RUBY]))
    return LootFile(
        sources=sources,
        quests={
            str(FIRE_RUBY): [
                QuestSource(
                    quest_id=8253, name="Destroy Morphaz", faction="both",
                    min_level=50, level=52, level_source="classic-db", classes=classes,
                )
            ]
        },
    )


def test_quest_only_item_excluded_from_every_class_but_the_quests_own():
    document = _quest_only_document(["mage"])
    exclusions = quest_only_class_exclusions(document, ["hunter", "mage", "paladin", "shaman"])
    assert exclusions == {
        "hunter": {FIRE_RUBY},
        "paladin": {FIRE_RUBY},
        "shaman": {FIRE_RUBY},
    }
    assert "mage" not in exclusions


def test_quest_open_to_any_class_excludes_nobody():
    document = _quest_only_document(None)
    exclusions = quest_only_class_exclusions(document, ["hunter", "mage"])
    assert exclusions == {}


def test_an_item_with_another_source_besides_quest_is_never_excluded():
    """An item that also drops in the world stays in every class it can
    equip, whatever one particular quest's own RequiredClasses says."""
    document = _quest_only_document(["mage"], also_dropped=True)
    exclusions = quest_only_class_exclusions(document, ["hunter", "mage"])
    assert exclusions == {}


def test_an_item_with_one_open_quest_among_several_stays_for_every_class():
    """Two quests reward the same item: one mage-only, one open to any
    class -- the item is still reachable by every class through the
    second quest, so no class is excluded."""
    document = LootFile(
        sources=[LootSource(id="quest", kind="quest", name="Quests", items=[FIRE_RUBY])],
        quests={
            str(FIRE_RUBY): [
                QuestSource(
                    quest_id=8253, name="Destroy Morphaz", faction="both",
                    min_level=50, level=52, level_source="classic-db", classes=["mage"],
                ),
                QuestSource(
                    quest_id=9999, name="An Open Alternative", faction="both",
                    min_level=50, level=52, level_source="classic-db", classes=None,
                ),
            ]
        },
    )
    exclusions = quest_only_class_exclusions(document, ["hunter", "mage"])
    assert exclusions == {}


def test_apply_quest_class_gate_drops_the_row_from_the_excluded_classes_only(tmp_path):
    build_dir = tmp_path
    items_dir = build_dir / "items"
    items_dir.mkdir()
    for class_slug in ("mage", "hunter"):
        (items_dir / f"{class_slug}.json").write_text(
            json.dumps(
                ClassItems(
                    build="test", class_slug=class_slug,
                    items=[_gear_item(FIRE_RUBY), _gear_item(OTHER_ITEM)],
                ).model_dump()
            ),
            encoding="utf-8",
        )
    document = _quest_only_document(["mage"])

    dropped = apply_quest_class_gate_to_class_items(build_dir, document)

    assert dropped == {"hunter": 1}
    mage_ids = {
        row["id"]
        for row in json.loads((items_dir / "mage.json").read_text(encoding="utf-8"))["items"]
    }
    hunter_ids = {
        row["id"]
        for row in json.loads((items_dir / "hunter.json").read_text(encoding="utf-8"))["items"]
    }
    assert mage_ids == {FIRE_RUBY, OTHER_ITEM}
    assert hunter_ids == {OTHER_ITEM}


def test_apply_quest_class_gate_is_a_noop_with_nothing_to_drop(tmp_path):
    build_dir = tmp_path
    items_dir = build_dir / "items"
    items_dir.mkdir()
    path = items_dir / "mage.json"
    path.write_text(
        json.dumps(
            ClassItems(build="test", class_slug="mage", items=[_gear_item(OTHER_ITEM)]).model_dump()
        ),
        encoding="utf-8",
    )
    before = path.read_bytes()
    document = LootFile(sources=[], quests={})

    dropped = apply_quest_class_gate_to_class_items(build_dir, document)

    assert dropped == {}
    assert path.read_bytes() == before
