import csv
import json
import logging
from pathlib import Path

logger = logging.getLogger(__name__)


def read_csv(path: Path) -> list[dict[str, str]]:
    with path.open(newline="", encoding="utf-8") as f:
        return [dict(row) for row in csv.DictReader(f)]


def _committed_item_count(build_dir: Path) -> int:
    """The last committed `items.json`'s own id count, or 0 when this build
    has never had one committed -- a first-ever run for a build is never a
    shrink, the same rule `normalize._check_class_items_not_shrunk` applies
    per class."""
    path = build_dir / "items.json"
    if not path.exists():
        return 0
    return len(json.loads(path.read_text(encoding="utf-8")))


def check_item_sparse_completeness(
    build_dir: Path,
    new_count: int,
    *,
    allow_shrink: bool = False,
) -> None:
    """Refuse a `normalize`/`loot` run whose item catalog would shrink
    against the build's own last committed `items.json`.

    2026-09-29 (night-fetch-guard): wago.tools served build 1.60.1.70009's
    export with only 19,226 of `Item.csv`'s 31,819 `ItemSparse` rows --
    reproduced on two separate fetches -- which silently halved every
    class's committed `items/<class>.json` when `normalize`/`loot` ran
    against it. The original fix compared `ItemSparse.csv` to `Item.csv`
    and demanded 90% -- but Forever's own client ships that shipped table at
    genuinely ~60% by design (most of the catalog arrives as server
    hotfixes the client's own `.db2` export never carries; see
    `pipeline.hotfix_cache`), so a fixed ratio of `Item.csv` cannot tell a
    broken export from Forever's actual shape. The real invariant this gate
    protects is narrower and needs no ratio at all: whatever `normalize`
    or `loot` is about to commit must not have FEWER item ids than the
    build's own last good run already committed (normalize-levels lane,
    2026-09-29).

    `new_count` is the caller's own count of what it is about to write --
    `normalize_build` (the only caller; `loot.write_loot_files` does not
    recompute `items.json` itself and no longer calls this) passes its
    final, wowhead-merged `items` list length, checked (and, on success,
    written) only after any narrower per-class regression
    (`_check_class_items_not_shrunk`) has already had first refusal.

    `allow_shrink=True` (the CLI's `--allow-shrink`) skips the check for a
    deliberate re-baseline, and always logs that it did, loudly, so a
    silent flip of this flag can't hide a real regression in CI logs.
    """
    previous_count = _committed_item_count(build_dir)
    if previous_count == 0:
        return
    if allow_shrink:
        logger.warning(
            "ItemSparse completeness gate skipped (--allow-shrink) for %s", build_dir
        )
        return
    if new_count < previous_count:
        raise SystemExit(
            f"{build_dir / 'items.json'} would shrink from {previous_count} to "
            f"{new_count} item ids; refusing to overwrite the committed build -- "
            f"re-fetch the build (and merge any hotfixes), or pass --allow-shrink "
            f"to force a deliberate re-baseline"
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
