# data/tests/test_normalize_itemnames.py
import json
from pathlib import Path

from pipeline.models import ClassItems, GearItem, Item
from pipeline.normalize.itemnames import (
    item_names_outside,
    write_item_names,
    write_item_names_from_build,
)


def _item(id_: int, name: str, inventory_type: int, quality: int = 1) -> Item:
    return Item(
        id=id_,
        name=name,
        quality=quality,
        item_level=1,
        required_level=0,
        class_id=4,
        subclass_id=0,
        inventory_type=inventory_type,
    )


def _gear(id_: int, name: str) -> GearItem:
    return GearItem(
        id=id_,
        name=name,
        icon="inv_misc_questionmark",
        slot="finger",
        quality=3,
        required_level=0,
        item_level=1,
        armor=0,
        stats={},
        set_id=None,
        unique=False,
    )


def test_names_only_wearable_items_the_planner_files_leave_out():
    items = [
        _item(264908, "Ancient Heirloom", 11),  # quality-1 ring: wearable, not planner gear
        _item(263412, "Windcarved Effigy", 28),  # relic slot: wearable, no planner slot
        _item(21182, "Band of Earthen Might", 11, quality=3),  # in a class file below
        _item(2589, "Linen Cloth", 0),  # not wearable at all
    ]
    class_items = [
        ClassItems(build="b", class_slug="warrior", items=[_gear(21182, "Band of Earthen Might")])
    ]
    assert item_names_outside(items, class_items) == {
        263412: "Windcarved Effigy",
        264908: "Ancient Heirloom",
    }


def test_write_and_regenerate_from_a_build_directory(tmp_path: Path):
    build = tmp_path / "1.0.0.1"
    (build / "items").mkdir(parents=True)
    items = [_item(264908, "Ancient Heirloom", 11), _item(21182, "Band", 11, quality=3)]
    class_items = [ClassItems(build="1.0.0.1", class_slug="warrior", items=[_gear(21182, "Band")])]
    (build / "items.json").write_text(json.dumps([i.model_dump() for i in items]))
    (build / "items" / "warrior.json").write_text(class_items[0].model_dump_json())

    first = write_item_names("1.0.0.1", items, class_items, build)
    written = json.loads(first.read_text())
    assert written == {"build": "1.0.0.1", "names": {"264908": "Ancient Heirloom"}}

    again = write_item_names_from_build("1.0.0.1", root=tmp_path)
    assert again.read_text() == first.read_text()
