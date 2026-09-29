import csv
import logging
from pathlib import Path

logger = logging.getLogger(__name__)

#: Minimum fraction of `Item.csv`'s row count that `ItemSparse.csv` must have
#: in the same `raw/` directory. wago.tools served build 1.60.1.70009's
#: export with only 19,226 of Item.csv's 31,819 ItemSparse rows -- reproduced
#: on two separate fetches -- which silently halved every class's committed
#: `items/<class>.json` when `normalize`/`loot` ran against it (night-relics
#: finding, 2026-09-29). Below this ratio the export is treated as broken,
#: not as Blizzard genuinely holding half the catalog back.
ITEM_SPARSE_COMPLETENESS_RATIO = 0.9


def read_csv(path: Path) -> list[dict[str, str]]:
    with path.open(newline="", encoding="utf-8") as f:
        return [dict(row) for row in csv.DictReader(f)]


def check_item_sparse_completeness(
    raw: Path,
    *,
    allow_shrink: bool = False,
    item_rows: list[dict[str, str]] | None = None,
    sparse_rows: list[dict[str, str]] | None = None,
) -> None:
    """Fail fast when `raw/ItemSparse.csv` is a truncated wago.tools export.

    Both `normalize.normalize_build` and `loot.write_loot_files` build the
    committed item catalog from this exact pair of raw tables, so this is
    the one place both call, before either reads `ItemSparse.csv` for real
    (night-fetch-guard, 2026-09-29). A build whose `Item.csv` has zero rows
    is not this function's problem -- something upstream of it already
    failed -- so it is left alone here rather than divided by zero.

    `item_rows`/`sparse_rows`, when given, are counted instead of re-reading
    `raw/Item.csv`/`raw/ItemSparse.csv` from disk -- `normalize_build` passes
    its hotfix-merged rows (`pipeline.hotfix_merge.merge_hotfix_table`) so a
    row the client only has as a hotfix counts toward completeness too
    (hotfix-cache lane, 2026-09-29), the same rows the rest of the build
    actually uses.

    `allow_shrink=True` (the CLI's `--allow-shrink`) skips the check for a
    deliberate re-baseline, and always logs that it did, loudly, so a
    silent flip of this flag can't hide a real regression in CI logs.
    """
    if allow_shrink:
        logger.warning("ItemSparse completeness gate skipped (--allow-shrink) for %s", raw)
        return
    item_count = len(item_rows) if item_rows is not None else len(read_csv(raw / "Item.csv"))
    sparse_count = (
        len(sparse_rows) if sparse_rows is not None else len(read_csv(raw / "ItemSparse.csv"))
    )
    if item_count and sparse_count < item_count * ITEM_SPARSE_COMPLETENESS_RATIO:
        raise SystemExit(
            f"{raw / 'ItemSparse.csv'} has only {sparse_count} rows against "
            f"{raw / 'Item.csv'}'s {item_count} rows ({sparse_count / item_count:.0%}, need "
            f"{ITEM_SPARSE_COMPLETENESS_RATIO:.0%}); wago.tools likely served a truncated "
            f"export (see the night-relics finding, 2026-09-29) -- re-fetch the build, or "
            f"pass --allow-shrink to force a deliberate re-baseline"
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
