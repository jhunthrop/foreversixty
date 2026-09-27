"""itemnames.json: a name for every wearable item the planner files leave out.

`items/<class-slug>.json` carries only gear the planner recommends (uncommon or
better, level 60 or under, a slot the planner tracks). A character can still be
wearing something outside that set -- a quality-1 keepsake ring, a shaman's
totem in the relic slot -- and the site used to show "Item 264908, not in this
build's data yet" for it, which reads as a data bug when the item is simply not
gear the sim scores. This file is the small lookup (about 2,600 rows, ~80 KB)
the simulator strip and the planner fetch only when they meet such an id, so
they can name it and say plainly that it is not simmed.
"""

from __future__ import annotations

import json
import logging
from pathlib import Path

from pipeline.models import ClassItems, Item

logger = logging.getLogger(__name__)

ITEMNAMES_FILE = "itemnames.json"

#: InventoryType 0 is not wearable at all (reagents, quest items, recipes).
NOT_WEARABLE = 0


def item_names_outside(items: list[Item], class_items: list[ClassItems]) -> dict[int, str]:
    """id -> name for every wearable item no per-class file carries, sorted by id."""
    planner_ids = {item.id for record in class_items for item in record.items}
    return {
        item.id: item.name
        for item in sorted(items, key=lambda i: i.id)
        if item.inventory_type != NOT_WEARABLE and item.id not in planner_ids
    }


def write_item_names(
    build: str, items: list[Item], class_items: list[ClassItems], build_dir: Path
) -> Path:
    names = item_names_outside(items, class_items)
    path = build_dir / ITEMNAMES_FILE
    path.write_text(
        json.dumps(
            {"build": build, "names": {str(k): v for k, v in names.items()}},
            indent=1,
            sort_keys=True,
        )
        + "\n",
        encoding="utf-8",
    )
    logger.info("wrote %s: %d wearable items outside the planner files", path, len(names))
    return path


def write_item_names_from_build(build: str, root: Path = Path("builds")) -> Path:
    """The same file from a build directory's committed items.json and items/*.json, so it
    can be (re)generated without another normalize pass."""
    build_dir = root / build
    items = [Item.model_validate(row) for row in json.loads((build_dir / "items.json").read_text())]
    class_items = [
        ClassItems.model_validate_json(path.read_text(encoding="utf-8"))
        for path in sorted((build_dir / "items").glob("*.json"))
    ]
    return write_item_names(build, items, class_items, build_dir)
