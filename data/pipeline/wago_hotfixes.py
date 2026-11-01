"""wago.tools' record of Blizzard's item hotfix pushes, layered newest build first.

The committed `raw/hotfixes/` dump is one DBCache read (1.60.1.70009), seeded
forward add-only through every build since, so every push after that session
was missing: on 1.60.1.70707, 201 of its rows disagreed with what the server
sends (Fiery War Axe 870 at 2.9 s against 3.6 s, Blade of Hanna 2801 with five
stats it no longer has, The Stitcher 279874 rare instead of uncommon), and 569
hotfix-only items were absent. Wowhead Forever's tooltips agree with wago on
every one checked (sim-accuracy round 5, physical lane).

wago.tools serves a table with the pushes it saw for a build applied
(`/db2/<Table>/csv?build=<b>&useHotfixes=1`). It records full push sets for some
builds only (70532 yes, 70685 none, 70707 a few hundred), so one build's answer
is not enough: `fetch_wago_hotfixes` reads every build of the product's patch
line up to the target, keeps the rows each build's hotfixes changed, and lets
the newest build win per id. The result goes to `raw/wago-hotfixes/<Table>.csv`
(ignored, like the other raw exports: it is re-fetched on every data run) with
`HOTFIX_BUILD_COLUMN` naming the build each row came from, and
`pipeline.hotfix_merge.merge_hotfix_table` layers it over the dump.
"""

from __future__ import annotations

import csv
import io
from collections.abc import Iterable
from pathlib import Path
from typing import Any

import httpx

from pipeline.wago import BASE_URL, USER_AGENT

#: Column naming the build whose hotfix push a row came from.
HOTFIX_BUILD_COLUMN = "_HotfixBuild"
#: Where the layered rows land, beside the other raw exports.
WAGO_HOTFIX_DIR = "wago-hotfixes"
#: The tables `merge_hotfix_table` merges hotfixes for.
TABLES: tuple[str, ...] = ("ItemSparse", "Item")

_TIMEOUT_SECONDS = 180


def _normalised(value: str | None) -> object:
    """A cell compared by value: wago prints floats at two widths
    (0.4000000059604645 plain, 0.40000000596046 hotfixed)."""
    try:
        return round(float(value), 6) if value is not None else None
    except ValueError:
        return value


def _same(a: dict[str, str], b: dict[str, str]) -> bool:
    return all(_normalised(a.get(key)) == _normalised(value) for key, value in b.items())


def changed_rows(
    plain: Iterable[dict[str, str]], hotfixed: Iterable[dict[str, str]]
) -> dict[str, dict[str, str]]:
    """The rows of `hotfixed` that `plain` lacks or states differently, by id."""
    shipped = {row["ID"]: row for row in plain}
    return {
        row["ID"]: row
        for row in hotfixed
        if row["ID"] not in shipped or not _same(shipped[row["ID"]], row)
    }


def layer_newest(
    by_build: Iterable[tuple[str, dict[str, dict[str, str]]]],
) -> list[dict[str, str]]:
    """Each build's changed rows (oldest build first), the later build winning
    per id, every row tagged with its build; sorted by id."""
    layered: dict[str, dict[str, str]] = {}
    for build, rows in by_build:
        for row_id, row in rows.items():
            layered[row_id] = {**row, HOTFIX_BUILD_COLUMN: build}
    return [layered[row_id] for row_id in sorted(layered, key=int)]


def _version(version: str) -> tuple[int, ...]:
    return tuple(int(part) for part in version.split("."))


def hotfix_builds(api_builds: dict[str, Any], product: str, build: str) -> list[str]:
    """`product`'s builds on `build`'s patch line (the same first three version
    parts) up to and including `build`, oldest first."""
    target = _version(build)
    versions = {
        entry["version"]
        for entry in api_builds.get(product, [])
        if _version(entry["version"])[:3] == target[:3] and _version(entry["version"]) <= target
    }
    if build not in versions:
        raise SystemExit(f"wago.tools lists no build {build!r} for product {product!r}")
    return sorted(versions, key=_version)


def _table(client: httpx.Client, table: str, build: str, hotfixed: bool) -> list[dict[str, str]]:
    params = {"build": build, **({"useHotfixes": "1"} if hotfixed else {})}
    response = client.get(f"/db2/{table}/csv", params=params, timeout=_TIMEOUT_SECONDS)
    response.raise_for_status()
    rows = [dict(row) for row in csv.DictReader(io.StringIO(response.text))]
    if rows and "ID" not in rows[0]:
        raise SystemExit(f"{table} for build {build} came back without an ID column")
    return rows


def _write(rows: list[dict[str, str]], path: Path) -> None:
    fieldnames: list[str] = []
    for row in rows:
        fieldnames.extend(key for key in row if key not in fieldnames)
    path.parent.mkdir(parents=True, exist_ok=True)
    with path.open("w", newline="", encoding="utf-8") as f:
        writer = csv.DictWriter(f, fieldnames=fieldnames, restval="")
        writer.writeheader()
        writer.writerows(rows)


def fetch_wago_hotfixes(
    build: str,
    product: str,
    root: Path = Path("builds"),
    client: httpx.Client | None = None,
    tables: tuple[str, ...] = TABLES,
) -> dict[str, int]:
    """`python -m pipeline fetch-wago-hotfixes`: write
    `<root>/<build>/raw/wago-hotfixes/<Table>.csv` for each table and return
    how many rows each carries."""
    own = client is None
    client = client or httpx.Client(base_url=BASE_URL, headers={"User-Agent": USER_AGENT})
    try:
        response = client.get("/api/builds")
        response.raise_for_status()
        builds = hotfix_builds(response.json(), product, build)
        out_dir = root / build / "raw" / WAGO_HOTFIX_DIR
        written: dict[str, int] = {}
        for table in tables:
            rows = layer_newest(
                (b, changed_rows(_table(client, table, b, False), _table(client, table, b, True)))
                for b in builds
            )
            _write(rows, out_dir / f"{table}.csv")
            written[table] = len(rows)
        return written
    finally:
        if own:
            client.close()
