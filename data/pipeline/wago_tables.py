"""wago.tools' DB2 CSV exports, kept between runs so a table is downloaded once.

A build's shipped tables never change: `/db2/<Table>/csv?build=<b>` is the
same bytes every time it is asked for, yet data.yml, the nightly ranking and
the engine follow each downloaded all of them again, and
`pipeline.wago_hotfixes` asked for the same ItemSparse once per build of the
patch line on every run. The hotfixed export (`useHotfixes=1`) does move, as
wago records new pushes, but wago says how fresh it is: `cache-control:
max-age=14400` and a `last-modified` it honours with a 304.

So a `TableCache` keeps each answer under `<dir>/<build>/<plain|hotfixed>/`:

* a shipped table, once cached, is never asked for again;
* a hotfixed table is reused while younger than its max-age, then revalidated
  with If-Modified-Since (a 304 costs no body), and replaced on a 200;
* a "no such table" answer (400 or 404) is cached like a body.

The directory comes from WAGO_TABLE_CACHE; the workflows keep it in
actions/cache. Unset, nothing is cached and every call downloads.
"""

from __future__ import annotations

import json
import logging
import os
import re
import time
from collections.abc import Callable
from dataclasses import dataclass
from pathlib import Path

import httpx

logger = logging.getLogger(__name__)

CACHE_ENV = "WAGO_TABLE_CACHE"
#: wago's own cache-control for a table export; used when an answer carries none.
DEFAULT_MAX_AGE_S = 14_400
_MAX_AGE = re.compile(r"max-age=(\d+)")


@dataclass(frozen=True)
class TableAnswer:
    """A table export: its HTTP status (200, or the 400/404 of a missing table) and body."""

    status: int
    text: str


def _variant(hotfixed: bool) -> str:
    return "hotfixed" if hotfixed else "plain"


def _params(build: str, hotfixed: bool) -> dict[str, str]:
    return {"build": build, **({"useHotfixes": "1"} if hotfixed else {})}


def _max_age(response: httpx.Response) -> int:
    match = _MAX_AGE.search(response.headers.get("cache-control", ""))
    return int(match.group(1)) if match else DEFAULT_MAX_AGE_S


@dataclass(frozen=True)
class TableCache:
    """Table exports on disk under `root`, keyed by (build, hotfixed, table)."""

    root: Path
    clock: Callable[[], float] = time.time

    def _paths(self, table: str, build: str, hotfixed: bool) -> tuple[Path, Path]:
        base = self.root / build / _variant(hotfixed)
        return base / f"{table}.csv", base / f"{table}.json"

    def read(self, table: str, build: str, hotfixed: bool) -> tuple[TableAnswer, dict] | None:
        body, meta_path = self._paths(table, build, hotfixed)
        if not (body.exists() and meta_path.exists()):
            return None
        meta = json.loads(meta_path.read_text(encoding="utf-8"))
        return TableAnswer(int(meta["status"]), body.read_text(encoding="utf-8")), meta

    def write(
        self, table: str, build: str, hotfixed: bool, answer: TableAnswer, meta: dict
    ) -> None:
        body, meta_path = self._paths(table, build, hotfixed)
        body.parent.mkdir(parents=True, exist_ok=True)
        body.write_text(answer.text, encoding="utf-8")
        meta_path.write_text(
            json.dumps({**meta, "status": answer.status}, sort_keys=True), encoding="utf-8"
        )

    def touch(self, table: str, build: str, hotfixed: bool, meta: dict) -> None:
        _, meta_path = self._paths(table, build, hotfixed)
        meta_path.write_text(json.dumps(meta, sort_keys=True), encoding="utf-8")

    def fresh(self, meta: dict, hotfixed: bool) -> bool:
        if not hotfixed:
            return True
        return self.clock() - float(meta["fetched_at"]) < float(meta["max_age"])


def table_cache_from_env() -> TableCache | None:
    value = os.environ.get(CACHE_ENV, "").strip()
    return TableCache(Path(value)) if value else None


#: How wago answers for a table a build does not have: 404 for most, and 400 for the
#: 1.60.1.70009 beta build's JournalInstance (2026-09-27). Both mean "no such table here".
MISSING_TABLE_STATUSES = frozenset({400, 404})


def get_table(
    client: httpx.Client,
    table: str,
    build: str,
    hotfixed: bool = False,
    cache: TableCache | None = None,
) -> TableAnswer:
    """`table`'s export for `build`, from `cache` when it still holds a usable copy.

    Any status but 200, 304 and a missing table's 400/404 raises."""
    cached = cache.read(table, build, hotfixed) if cache is not None else None
    if cached is not None and cache is not None and cache.fresh(cached[1], hotfixed):
        return cached[0]
    headers = {}
    if cached is not None and cached[1].get("last_modified"):
        headers["If-Modified-Since"] = cached[1]["last_modified"]
    response = client.get(f"/db2/{table}/csv", params=_params(build, hotfixed), headers=headers)
    now = cache.clock() if cache is not None else time.time()
    if response.status_code == 304 and cached is not None and cache is not None:
        cache.touch(table, build, hotfixed, {**cached[1], "fetched_at": now})
        return cached[0]
    if response.status_code not in MISSING_TABLE_STATUSES:
        response.raise_for_status()
    answer = TableAnswer(response.status_code, response.text)
    if cache is not None:
        meta = {
            "fetched_at": now,
            "max_age": _max_age(response),
            "last_modified": response.headers.get("last-modified", ""),
        }
        cache.write(table, build, hotfixed, answer, meta)
    return answer
