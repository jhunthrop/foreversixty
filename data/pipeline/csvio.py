import csv
from pathlib import Path


def read_csv(path: Path) -> list[dict[str, str]]:
    with path.open(newline="", encoding="utf-8") as f:
        return [dict(row) for row in csv.DictReader(f)]


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
