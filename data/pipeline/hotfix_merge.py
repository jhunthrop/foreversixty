"""Merge a table's hotfix rows (`pipeline.hotfix_cache`, written to
`raw/hotfixes/<Table>.csv` by `python -m pipeline hotfixes`) over its shipped
`raw/<Table>.csv` -- hotfix wins per id, a `RECORD_REMOVED` hotfix drops the
row entirely, everything else passes through untouched.

This is deliberately its own module rather than living in `pipeline/csvio.py`:
`csvio.read_csv`/`check_item_sparse_completeness` know nothing about hotfixes
and are used by tables (`SpellName`, `AreaTable`, ...) this lane never
touches, so keeping the merge here means a build with no `raw/hotfixes/`
directory (every build before this lane, and every table this lane's CLI
does not write) behaves exactly as before -- `merge_hotfix_table` is a
passthrough the moment `raw/hotfixes/<Table>.csv` does not exist.
"""

from __future__ import annotations

import csv
from dataclasses import dataclass
from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.db2 import ITEM_SPARSE_SCHEMA, rows_to_csv_dicts
from pipeline.hotfix_cache import (
    ITEM_SCHEMA,
    ITEM_SPARSE_TABLE_HASH,
    ITEM_TABLE_HASH,
    HotfixCacheError,
    TableHotfixes,
    decode_table_hotfixes,
    read_cache,
)

#: Column `hotfix_csv_rows` adds to mark a row's hotfix status. Not a real
#: client column -- `merge_hotfix_rows` strips it before a hotfix row
#: replaces a shipped one, so a downstream reader that only knows the real
#: schema (`pipeline.normalize.gear`, etc.) never sees it.
HOTFIX_STATUS_COLUMN = "_HotfixStatus"
STATUS_VALID = "valid"
STATUS_REMOVED = "removed"


def hotfix_csv_rows(hotfixes: TableHotfixes) -> list[dict[str, str]]:
    """`hotfixes` flattened to the CSV shape `write_hotfix_csv`/`merge_hotfix_rows`
    read: one row per valid id (`pipeline.db2.rows_to_csv_dicts`'s column
    layout, plus `HOTFIX_STATUS_COLUMN`), plus one minimal `ID`-only row per
    removed id, sorted by id.
    """
    rows = rows_to_csv_dicts(hotfixes.rows)
    for row in rows:
        row[HOTFIX_STATUS_COLUMN] = STATUS_VALID
    for removed_id in hotfixes.removed_ids:
        rows.append({"ID": str(removed_id), HOTFIX_STATUS_COLUMN: STATUS_REMOVED})
    rows.sort(key=lambda row: int(row["ID"]))
    return rows


def write_hotfix_csv(rows: list[dict[str, str]], path: Path) -> None:
    """Write `rows` (`hotfix_csv_rows`'s shape) as a CSV `pipeline.csvio.read_csv`
    can read back, and `merge_hotfix_rows` can merge over the shipped table.

    The header is the widest row's own key order (`ID` first, then the
    schema's fields in `pipeline.db2.rows_to_csv_dicts`' order, then
    `HOTFIX_STATUS_COLUMN` last) plus any key only a narrower row carries --
    a removed-id stub has no schema fields at all, so `restval=""` fills them
    in blank rather than raising.
    """
    fieldnames: list[str] = list(max(rows, key=len, default={}).keys())
    for row in rows:
        for key in row:
            if key not in fieldnames:
                fieldnames.append(key)
    path.parent.mkdir(parents=True, exist_ok=True)
    with path.open("w", newline="", encoding="utf-8") as f:
        writer = csv.DictWriter(f, fieldnames=fieldnames, restval="")
        writer.writeheader()
        writer.writerows(rows)


def merge_hotfix_rows(
    shipped: list[dict[str, str]], hotfix: list[dict[str, str]]
) -> list[dict[str, str]]:
    """`shipped` (a raw table's own rows) with `hotfix` (`hotfix_csv_rows`'
    shape) merged over it: a `valid` hotfix id replaces the shipped row (or
    adds one, for an id `shipped` never had) and a `removed` hotfix id drops
    the row, whether or not `shipped` had it. Rows keep `shipped`'s column
    set; `HOTFIX_STATUS_COLUMN` never appears in the result.
    """
    by_id = {row["ID"]: row for row in shipped}
    for row in hotfix:
        row_id = row["ID"]
        if row.get(HOTFIX_STATUS_COLUMN) == STATUS_REMOVED:
            by_id.pop(row_id, None)
            continue
        by_id[row_id] = {k: v for k, v in row.items() if k != HOTFIX_STATUS_COLUMN}
    return [by_id[row_id] for row_id in sorted(by_id, key=int)]


def merge_hotfix_table(raw: Path, table: str) -> list[dict[str, str]]:
    """`raw/<table>.csv` merged with `raw/hotfixes/<table>.csv`, or the
    shipped rows unchanged when no hotfix file exists for this table."""
    shipped = read_csv(raw / f"{table}.csv")
    hotfix_path = raw / "hotfixes" / f"{table}.csv"
    if not hotfix_path.exists():
        return shipped
    return merge_hotfix_rows(shipped, read_csv(hotfix_path))


#: (table name, table hash, schema) for every table `python -m pipeline
#: hotfixes` writes. Both hashes/schemas are `pipeline.db2`'s or
#: `pipeline.hotfix_cache`'s own (module docs there carry the verification);
#: this tuple is the one place that has to list every table, so adding a
#: third one later is a one-line change here plus a schema, not a rewrite of
#: `write_build_hotfix_tables`.
_TABLES: tuple[tuple[str, int, list], ...] = (
    ("ItemSparse", ITEM_SPARSE_TABLE_HASH, ITEM_SPARSE_SCHEMA),
    ("Item", ITEM_TABLE_HASH, ITEM_SCHEMA),
)


@dataclass(frozen=True)
class HotfixWriteResult:
    """One table's `write_build_hotfix_tables` outcome, for the CLI's printed
    summary line. `valid`/`removed`/`invalid`/`not_public` are hotfix
    *records* seen for this table (module doc on `TableHotfixes`: a
    record_id can be counted more than once if it was hotfixed twice).
    `new_ids` is how many valid ids `raw/<table>.csv` did not already have;
    `overriding_ids` is how many replace a shipped row.
    """

    table: str
    path: Path
    valid: int
    removed: int
    invalid: int
    not_public: int
    new_ids: int
    overriding_ids: int


def _trailing_build_number(build: str) -> int:
    """`"1.60.1.70009"` -> `70009` -- `HotfixHeader.build` is just that
    trailing integer, not the dotted version string the rest of the pipeline
    keys builds by."""
    return int(build.rsplit(".", 1)[-1])


def write_build_hotfix_tables(
    build: str, cache_path: Path, root: Path = Path("builds")
) -> list[HotfixWriteResult]:
    """`python -m pipeline hotfixes`: decode `cache_path` (a `DBCache.bin`
    the controller copied off the beta client -- never fetched by this
    pipeline itself) and write `builds/<build>/raw/hotfixes/{ItemSparse,
    Item}.csv` for `pipeline.normalize.normalize_build` to merge in.

    Raises `HotfixCacheError` when `cache_path`'s own `HotfixHeader.build`
    does not match `build`'s trailing number: hotfixes are pushed per build
    (module doc), so a cache from a play session on a different build would
    apply that build's row edits under this one's name.
    """
    table_hashes = {table_hash for _, table_hash, _ in _TABLES}
    header, records = read_cache(cache_path, table_hashes=table_hashes)
    wanted_build = _trailing_build_number(build)
    if header.build != wanted_build:
        raise HotfixCacheError(
            f"{cache_path} is a hotfix cache for build {header.build}, but --build "
            f"{build!r} names build {wanted_build}; a cache from a different build's "
            f"play session would apply the wrong build's row edits"
        )

    raw = root / build / "raw"
    results = []
    for table, table_hash, schema in _TABLES:
        hotfixes = decode_table_hotfixes(header, records, table_hash, schema)
        rows = hotfix_csv_rows(hotfixes)
        path = raw / "hotfixes" / f"{table}.csv"
        write_hotfix_csv(rows, path)

        shipped_path = raw / f"{table}.csv"
        shipped_ids = (
            {row["ID"] for row in read_csv(shipped_path)} if shipped_path.exists() else set()
        )
        hotfix_ids = {str(record_id) for record_id in hotfixes.rows}
        results.append(
            HotfixWriteResult(
                table=table,
                path=path,
                valid=hotfixes.valid_count,
                removed=hotfixes.removed_count,
                invalid=hotfixes.invalid_count,
                not_public=hotfixes.not_public_count,
                new_ids=len(hotfix_ids - shipped_ids),
                overriding_ids=len(hotfix_ids & shipped_ids),
            )
        )
    return results
