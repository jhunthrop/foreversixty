"""Fixes the `icon` field already committed in `items/<class>.json` for an
item whose client-table resolution (`pipeline.icons.resolve_icon`, run at
normalize time) landed on the placeholder, using the same fallback chain
`pipeline.icons.resolve_icon_name` applies: the engine fork's own icon, then
wowhead's gear-planner payload's.

A standalone pass rather than a rerun of `normalize`: `normalize` needs a
complete raw `ItemSparse` export, and wago.tools has been serving a
truncated one for build 1.60.1.70009 since 2026-09-27 (the night-relics
finding `pipeline.csvio.check_item_sparse_completeness` guards against) --
the completeness gate refuses it. This module never reads `ItemSparse` at
all: it only rewrites the `icon` field of the `items/<class>.json` rows
`normalize` already produced, from the two sources that do not need it --
the fork checkout's `assets/database/db.json` (`pipeline.forkdb`) and the
committed `raw/wowhead-gear-planner.js` (`fetch-wowhead` writes it;
`python -m pipeline icons --build <build> --engine <path>` reads it if
present and otherwise skips that source, same as `normalize` does for the
wowhead item supplement).

771 of the 9,391 real items on build 1.60.1.70009 carry this placeholder;
see `resolve_icon_name`'s own doc for the cause (a Forever hotfix item the
client's own `Item` table genuinely states no icon for, not a join this
pipeline is missing).
"""

from __future__ import annotations

import json
import logging
from collections import Counter
from dataclasses import dataclass
from pathlib import Path

from pipeline.icons import PLACEHOLDER_ICON, resolve_icon_name
from pipeline.models import ClassItems

logger = logging.getLogger(__name__)


@dataclass(frozen=True)
class ClassFixResult:
    """One items/<class>.json's placeholder count before and after this
    pass, and how many of the fixed ones came from each source."""

    class_slug: str
    total_items: int
    before_placeholder: int
    after_placeholder: int
    fixed_by_fork: int
    fixed_by_wowhead: int


def load_fork_icons(engine_dir: Path) -> dict[int, str]:
    """Item id -> icon from the fork's own db.json.

    Imported locally: nothing else in this module's own import chain needs
    `pipeline.forkdb`, and keeping the cross-module import local here (the
    same way `pipeline.wowhead_items` keeps its `pipeline.icons` import
    local, per that module's own doc) costs nothing and avoids ever having
    to reason about import order between the two.
    """
    from pipeline.forkdb import icon_by_item_id, load_fork_database

    return icon_by_item_id(load_fork_database(engine_dir))


def load_wowhead_icons(build_dir: Path) -> dict[int, str]:
    """Item id -> icon from the build's committed wowhead gear-planner
    payload, or `{}` when the build has none (`fetch-wowhead` has not run
    for it yet).

    Unlike `pipeline.normalize.wowhead.load_supplement`, every id the
    payload names is wanted here, not only the ones the client's
    `ItemSparse` lacks entirely: a placeholder-icon item IS in `ItemSparse`
    -- it just has no icon of its own -- so the supplement's own id filter
    would drop exactly the rows this module exists to fix.
    """
    from pipeline.wowhead_items import load_items, raw_path

    path = raw_path(build_dir)
    if not path.exists():
        return {}
    return {item.id: item.icon for item in load_items(path) if item.icon}


def _write_class_items(record: ClassItems, path: Path) -> None:
    """The same serialization `pipeline.normalize`'s `write_model` uses, so a
    class this pass does not touch stays byte-identical and a class it does
    touch diffs only on the `icon` fields that actually changed."""
    text = json.dumps(record.model_dump(), indent=2, ensure_ascii=False)
    path.write_text(text + "\n", encoding="utf-8")


def fix_placeholder_icons(
    build_dir: Path,
    fork_icons: dict[int, str],
    wowhead_icons: dict[int, str],
) -> list[ClassFixResult]:
    """Rewrite `icon` on every `items/<class>.json` row still on
    PLACEHOLDER_ICON that `fork_icons` or `wowhead_icons` names a real one
    for. A class with nothing to fix is reported but left unwritten.
    """
    results: list[ClassFixResult] = []
    for path in sorted((build_dir / "items").glob("*.json")):
        record = ClassItems.model_validate_json(path.read_text(encoding="utf-8"))
        before = sum(1 for item in record.items if item.icon in (PLACEHOLDER_ICON, "0", ""))
        if before == 0:
            results.append(ClassFixResult(record.class_slug, len(record.items), 0, 0, 0, 0))
            continue
        origins: Counter[str] = Counter()
        new_items = []
        changed = False
        for item in record.items:
            icon, origin = resolve_icon_name(item.icon, item.id, fork_icons, wowhead_icons)
            if origin != "client":
                origins[origin] += 1
            if icon != item.icon:
                changed = True
                item = item.model_copy(update={"icon": icon})
            new_items.append(item)
        after = sum(1 for item in new_items if item.icon in (PLACEHOLDER_ICON, "0", ""))
        if changed:
            _write_class_items(record.model_copy(update={"items": new_items}), path)
        results.append(
            ClassFixResult(
                record.class_slug,
                len(record.items),
                before,
                after,
                origins["fork"],
                origins["wowhead"],
            )
        )
        logger.info(
            "icons: %s %d -> %d placeholder icons (%d fixed via fork db, %d via wowhead)",
            record.class_slug,
            before,
            after,
            origins["fork"],
            origins["wowhead"],
        )
    return results


def fix_icons_for_build(
    build: str,
    root: Path = Path("builds"),
    engine: Path | None = None,
) -> list[ClassFixResult]:
    """The `python -m pipeline icons --build <build> [--engine <path>]` entry
    point for this pass. `engine` is optional the same way `normalize`'s
    wowhead supplement is: without it, the fork-db source is simply skipped
    (an empty `fork_icons`), not an error -- `resolve_icon_name` falls
    straight through to the wowhead payload.
    """
    build_dir = root / build
    fork_icons = load_fork_icons(engine) if engine is not None else {}
    wowhead_icons = load_wowhead_icons(build_dir)
    if not fork_icons and not wowhead_icons:
        logger.warning(
            "icons: no fork checkout (--engine) and no wowhead payload "
            "(raw/wowhead-gear-planner.js) for build %s; no placeholder icon "
            "names can be fixed this run",
            build,
        )
    return fix_placeholder_icons(build_dir, fork_icons, wowhead_icons)
