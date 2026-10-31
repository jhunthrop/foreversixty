import csv
import logging
from pathlib import Path

from pipeline.build_floors import (
    WOWHEAD_ITEMS_FLOOR_RATIO,
    committed_items,
    reference_for,
    wowhead_item_count,
)

logger = logging.getLogger(__name__)


def read_csv(path: Path) -> list[dict[str, str]]:
    with path.open(newline="", encoding="utf-8") as f:
        return [dict(row) for row in csv.DictReader(f)]


def check_item_sparse_completeness(
    build_dir: Path,
    new_count: int,
    *,
    allow_shrink: bool = False,
) -> None:
    """Refuse a `normalize` run whose item catalog would shrink against the
    build's own last committed `items.json`, or, on a build's first run,
    fall below `build_floors.CROSS_BUILD_ITEMS_RATIO` of the newest other
    build's.

    2026-09-29 (night-fetch-guard): wago.tools served build 1.60.1.70009's
    export with only 19,226 of `Item.csv`'s 31,819 `ItemSparse` rows --
    reproduced on two separate fetches -- which silently halved every
    class's committed `items/<class>.json` when `normalize`/`loot` ran
    against it. A fixed ratio of `Item.csv` cannot tell a broken export from
    Forever's actual shape (the shipped table is ~60% by design; most of the
    catalog arrives as server hotfixes, see `pipeline.hotfix_cache`), so the
    invariant is narrower: whatever `normalize` is about to commit must not
    have FEWER item ids than the build's own last good run already committed
    (normalize-levels lane, 2026-09-29).

    2026-10-30: a first run used to have nothing to compare with and always
    passed, which is how 1.60.1.70685 committed 313 fewer items than 70658
    (design/reviews/2026-10-30-70707-dps-drop.md). It is now compared with
    the previous build, with a small tolerance for a client that really did
    drop items.

    `new_count` is `normalize_build`'s final, wowhead-merged `items` list
    length, checked only after any narrower per-class regression
    (`_check_class_items_not_shrunk`) has already had first refusal.
    `allow_shrink=True` (the CLI's `--allow-shrink`) skips the check for a
    deliberate re-baseline, and logs loudly that it did.
    """
    reference = reference_for(build_dir, "items.json")
    previous_count = len(committed_items(reference.build_dir)) if reference else 0
    if previous_count == 0:
        return
    if allow_shrink:
        logger.warning(
            "ItemSparse completeness gate skipped (--allow-shrink) for %s", build_dir
        )
        return
    if new_count < previous_count * reference.ratio:
        against = (
            ""
            if reference.own
            else f" (below {reference.ratio:.0%} of {reference.build_dir.name}'s)"
        )
        raise SystemExit(
            f"{build_dir / 'items.json'} would shrink from {previous_count} to "
            f"{new_count} item ids{against}; refusing to write the build -- "
            f"re-fetch the build (and merge any hotfixes), or pass --allow-shrink "
            f"to force a deliberate re-baseline"
        )


def check_wowhead_items_floor(
    build_dir: Path,
    new_count: int,
    *,
    allow_shrink: bool = False,
) -> None:
    """Refuse a `normalize` run in which fewer than
    `build_floors.WOWHEAD_ITEMS_FLOOR_RATIO` of the reference run's items
    take their stats from wowhead's feed (the build's own last run, else the
    previous build's). A wowhead item that falls back to classic-db keeps
    the item count level while its stats roll back to 1.12, so the count
    gate above cannot see it: 70685 kept 7 of 70658's 688. A floor, never an
    exact count. `allow_shrink=True` skips it, loudly."""
    reference = reference_for(build_dir, "items.json")
    previous = wowhead_item_count(committed_items(reference.build_dir)) if reference else 0
    if previous == 0:
        return
    if allow_shrink:
        logger.warning("wowhead items floor skipped (--allow-shrink) for %s", build_dir)
        return
    if new_count < previous * WOWHEAD_ITEMS_FLOOR_RATIO:
        raise SystemExit(
            f"{build_dir.name} would take {new_count} items' stats from wowhead, below "
            f"{WOWHEAD_ITEMS_FLOOR_RATIO:.0%} of {reference.build_dir.name}'s {previous}; "
            "wowhead's feed has likely lost items -- check raw/wowhead/items.json's "
            "carry, or pass --allow-shrink to force a deliberate re-baseline"
        )


def populated(row: dict[str, str], key: str) -> str | None:
    """`row[key]`, or `None` when the column is absent or the client exported
    it empty.

    A DB2 export can carry a column header that every row leaves blank
    rather than dropping the column outright (see `pipeline/spelltext.py`'s
    `_base_points` and `pipeline/normalize/item_curves.py`'s `_budget`, both
    of which pick between two columns that are never both real on the same
    build). Testing for the key's mere presence, as an earlier version of
    each of those functions did, does not tell the two cases apart and
    raises deep inside `float()`/`int()` instead of falling back.
    """
    value = row.get(key)
    return value if value not in (None, "") else None
